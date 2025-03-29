package routes

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes/hello"
	"log"
	"net/http"
)

func RegisterRoutes(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %v", request)

	router := NewRouter()
	router.AddRoute(http.MethodGet, "/produtos", &hello.ListarProdutosHandler{})
	router.AddRoute(http.MethodPost, "/produtos", &hello.CriarProdutoHandler{})
	router.AddRoute(http.MethodGet, "/produtos/{id}", &hello.ObterProdutoHandler{})

	handler, pathParams := router.FindHandler(request.HTTPMethod, request.Path)

	if handler == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Rota não encontrada",
		}, nil
	}

	return handler.Handle(ctx, request, pathParams)
}
