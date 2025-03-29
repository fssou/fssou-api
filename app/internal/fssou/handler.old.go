package fssou

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"in.francl.api/internal/fssou/services/x"
)

func HandlerOld(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("received event")
	secretsXJsonString, exists := os.LookupEnv("SECRET_CREDENTIALS_X")
	if !exists {
		log.Println("SECRET_CREDENTIALS_X not found")
		return nil, fmt.Errorf("SECRET_CREDENTIALS_X not found")
	}
	var credentials x.Credentials
	err := json.Unmarshal([]byte(secretsXJsonString), &credentials)
	if err != nil {
		log.Printf("error unmarshalling secrets: %v\n", err)
		return nil, err
	}
	if request == nil {
		log.Println("received nil event")
		return &events.APIGatewayProxyResponse{}, fmt.Errorf("received nil event")
	}
	twitter, err := x.New(ctx, &credentials, nil)
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
	me, err := twitter.UsersMe()
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
