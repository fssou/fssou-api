package fssou

import (
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes"
)

func Handler(ctx events.LambdaFunctionURLRequestContext, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	return routes.RegisterRoutes(ctx, request)
}
