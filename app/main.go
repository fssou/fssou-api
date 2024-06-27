package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"in.francl.api/services/x"
)

func init() {
	fmt.Println("Lambda cold start")
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("received nil event")
	}
	x := x.X{}
	me, err := x.Me()
	if err != nil {
		resp := map[string]interface{}{
			"error": err.Error(),
		}
		body, _ := json.Marshal(resp)
		return &events.APIGatewayProxyResponse{
			StatusCode:        500,
			Body:              string(body),
			Headers:           map[string]string{},
			MultiValueHeaders: map[string][]string{},
			IsBase64Encoded:   false,
		}, nil
	}
	resp, _ := json.Marshal(me)
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Body:              string(resp),
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		IsBase64Encoded:   false,
	}, nil
}
