package fssou

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes"
)

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	return routes.RegisterRoutes(ctx, request)
}
