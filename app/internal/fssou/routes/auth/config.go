package auth

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// TokenResponse representa a estrutura da resposta do Discord ao solicitar o token
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

// Configurações globais
type Config struct {
	DiscordClientID     string
	DiscordClientSecret string
	DiscordRedirectURI  string
	FrontendURL         string
}

// Carrega configurações do Parameter Store
func loadConfig() (*Config, error) {
	// Inicializa uma sessão AWS

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	// Cria um cliente SSM
	ssmClient := ssm.New(sess)

	// Define os parâmetros a serem recuperados
	paramNames := []string{
		"/fssou/discord/auth/client-id",
		"/fssou/discord/auth/client-secret",
		"/fssou/discord/auth/redirect-uri",
		"/fssou/discord/auth/frontend-url",
	}

	// Recupera os parâmetros
	params := make(map[string]string)
	for _, name := range paramNames {
		param, err := ssmClient.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(name),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			return nil, err
		}
		params[name] = *param.Parameter.Value
	}

	// Retorna a configuração
	return &Config{
		DiscordClientID:     params["/fssou/discord/auth/client-id"],
		DiscordClientSecret: params["/fssou/discord/auth/client-secret"],
		DiscordRedirectURI:  params["/fssou/discord/auth/redirect-uri"],
		FrontendURL:         params["/fssou/discord/auth/frontend-url"],
	}, nil
}
