package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"in.francl.api/services/x"
	"log"
)

func init() {
	log.Println("Lambda cold start")
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("received event")
	if request == nil {
		log.Println("received nil event")
		return &events.APIGatewayProxyResponse{}, fmt.Errorf("received nil event")
	}
	twitter, err := x.New(ctx)
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
	log.Println("created x")
	me, err := twitter.Me()
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
	log.Println("got me")
	resp, _ := json.Marshal(me)
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Body:              string(resp),
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		IsBase64Encoded:   false,
	}, nil
}
