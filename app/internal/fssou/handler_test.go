package fssou

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestHandler",
			args: args{
				ctx: context.Background(),
				request: &events.APIGatewayProxyRequest{
					Body: `{"key":"value"}`,
				},
			},
			want: &events.APIGatewayProxyResponse{
				StatusCode:        200,
				Body:              `{"key":"value"}`,
				Headers:           map[string]string{},
				MultiValueHeaders: map[string][]string{},
				IsBase64Encoded:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}
