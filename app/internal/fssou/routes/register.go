package routes

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes/auth"
	"in.francl.api/internal/fssou/routes/hello"
	"log"
	"net/http"
)

func RegisterRoutes(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	log.Printf("Received request: %v", request)

	router := NewRouter()
	router.AddRoute(http.MethodGet, "/produtos", &hello.ListarProdutosHandler{})
	router.AddRoute(http.MethodPost, "/produtos", &hello.CriarProdutoHandler{})
	router.AddRoute(http.MethodGet, "/produtos/{id}", &hello.ObterProdutoHandler{})
	router.AddRoute(http.MethodGet, "/auth/discord/login", &auth.DiscordLogin{})
	router.AddRoute(http.MethodGet, "/auth/discord/callback", &auth.DiscordCallback{})

	handler, pathParams := router.FindHandler(request.RequestContext.HTTP.Method, request.RawPath)

	if handler == nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Rota não encontrada",
		}, nil
	}

	return handler.Handle(ctx, request, pathParams)
}
