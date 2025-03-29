package fssou

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/routes"
)

func Handler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return routes.RegisterRoutes(ctx, *request)
}
