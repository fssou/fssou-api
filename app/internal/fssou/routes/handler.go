package routes

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

// RouteHandler Interface para os handlers de rota
type RouteHandler interface {
	Handle(ctx events.LambdaFunctionURLRequestContext, request events.LambdaFunctionURLRequest, pathParams map[string]string) (events.LambdaFunctionURLResponse, error)
}

// FindHandler Encontra o handler para uma determinada rota e extrai pathParams
func (r *Router) FindHandler(method, path string) (RouteHandler, map[string]string) {
	if handlers, ok := r.routes[method]; ok {
		for routePattern, handler := range handlers {
			pathParams := extractPathParams(routePattern, path)
			if pathParams != nil {
				return handler, pathParams
			}
		}
	}
	return nil, nil
}

// Extrai pathParams de uma rota
func extractPathParams(routePattern, path string) map[string]string {
	routeParts := strings.Split(routePattern, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return nil
	}

	pathParams := make(map[string]string)
	for i, part := range routeParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := strings.Trim(part, "{}")
			pathParams[paramName] = pathParts[i]
		} else if part != pathParts[i] {
			return nil
		}
	}

	return pathParams
}
