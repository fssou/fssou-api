package hello

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

// ListarProdutosHandler Handlers das rotas
type ListarProdutosHandler struct{}

func (h *ListarProdutosHandler) Handle(ctx context.Context, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error) {
	produtos := []map[string]interface{}{
		{"id": 1, "nome": "Produto A", "preco": 10.99},
		{"id": 2, "nome": "Produto B", "preco": 25.50},
	}

	body, err := json.Marshal(produtos)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Erro ao serializar produtos",
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

type CriarProdutoHandler struct{}

func (h *CriarProdutoHandler) Handle(ctx context.Context, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error) {
	var produto map[string]interface{}
	err := json.Unmarshal([]byte(request.Body), &produto)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Corpo da requisição inválido",
		}, nil
	}

	// Simula a criação do produto
	log.Printf("Produto criado: %v", produto)

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusCreated,
		Body:       "Produto criado com sucesso",
	}, nil
}

type ObterProdutoHandler struct{}

func (h *ObterProdutoHandler) Handle(ctx context.Context, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error) {
	id := pathParams["id"]

	// Simula a busca do produto por ID
	produto := map[string]interface{}{
		"id":    id,
		"nome":  "Produto C",
		"preco": 15.00,
	}

	body, err := json.Marshal(produto)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Erro ao serializar produto",
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}
