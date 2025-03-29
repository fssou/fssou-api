package routes

import (
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes/hello"
	"log"
	"net/http"
)

func RegisterRoutes(ctx events.LambdaFunctionURLRequestContext, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	log.Printf("Received request: %v", request)

	router := NewRouter()
	router.AddRoute(http.MethodGet, "/produtos", &hello.ListarProdutosHandler{})
	router.AddRoute(http.MethodPost, "/produtos", &hello.CriarProdutoHandler{})
	router.AddRoute(http.MethodGet, "/produtos/{id}", &hello.ObterProdutoHandler{})

	handler, pathParams := router.FindHandler(request.RequestContext.HTTP.Method, request.RawPath)

	if handler == nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Rota não encontrada",
		}, nil
	}

	return handler.Handle(ctx, request, pathParams)
}
