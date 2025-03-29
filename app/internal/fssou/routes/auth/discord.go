package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DiscordLogin struct{}

func (h *DiscordLogin) Handle(ctx context.Context, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error) {
	config, err := loadConfig()
	if err != nil {
		log.Printf("Erro ao carregar configurações: %v", err)
		body := &routes.ErrorResponse{
			Code:    "500",
			Message: "Internal Server Error",
		}
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       body.ToString(),
		}, nil
	}

	scope := "identify"
	authURL := fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		config.DiscordClientID,
		url.QueryEscape(config.DiscordRedirectURI),
		url.QueryEscape(scope),
	)

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": authURL,
		},
	}, nil
}

type DiscordCallback struct{}

func (h *DiscordCallback) Handle(ctx context.Context, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error) {
	// Extrai o código de autorização
	code, ok := request.QueryStringParameters["code"]
	if !ok || code == "" {
		body := &routes.ErrorResponse{
			Code:    "400",
			Message: "Código de autorização não fornecido",
		}
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       body.ToString(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Carrega configurações
	config, err := loadConfig()
	if err != nil {
		log.Printf("Erro ao carregar configurações: %v", err)
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Erro interno do servidor"}`,
		}, nil
	}

	// Prepara os dados para a requisição do token
	data := url.Values{}
	data.Set("client_id", config.DiscordClientID)
	data.Set("client_secret", config.DiscordClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.DiscordRedirectURI)

	// Realiza a requisição para obter o token
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://discord.com/api/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Erro ao criar requisição"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Erro na requisição ao Discord"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Erro ao fechar o corpo da resposta: %v", err)
		}
	}(resp.Body)

	// Verifica se a resposta foi bem-sucedida
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Erro na autenticação Discord: %s", string(bodyBytes))
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Falha na autenticação com Discord"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Processa a resposta
	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       `{"error": "Erro ao processar resposta do Discord"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Configuração de cookies para o access token
	accessTokenExpires := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).Format(time.RFC1123)
	accessTokenCookie := fmt.Sprintf("discord_access_token=%s; HttpOnly; Path=/; Expires=%s; Secure",
		tokenResp.AccessToken,
		accessTokenExpires,
	)

	// Configuração de cookies para o refresh token (30 dias)
	refreshTokenExpires := time.Now().Add(30 * 24 * time.Hour).Format(time.RFC1123)
	refreshTokenCookie := fmt.Sprintf("discord_refresh_token=%s; HttpOnly; Path=/; Expires=%s; Secure",
		tokenResp.RefreshToken,
		refreshTokenExpires,
	)

	// Configuração da resposta com redirecionamento e cookies
	return events.LambdaFunctionURLResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": config.FrontendURL,
		},
		Cookies: []string{
			accessTokenCookie,
			refreshTokenCookie,
		},
	}, nil
}
