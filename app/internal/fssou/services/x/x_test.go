package x

import (
	"context"
	"testing"

	"in.francl.api/internal/fssou/services/http"
)

func TestMe(t *testing.T) {
	mockClient := &MockHttpClient{
		GetFunc: func(path string, params *http.Params) (*http.Response, error) {
			// Simula uma resposta do servidor
			return &http.Response{
				StatusCode: 200,
				Body:       []byte(`{"data": {"username": "fssouoficial"}}`),
			}, nil
		},
	}
	x, err := New(context.Background(), &Credentials{
		ConsumerKey:    "",
		ConsumerSecret: "",
		TokenKey:       "",
		TokenSecret:    "",
	}, mockClient)
	if err != nil {
		t.Errorf("New() retornou um erro: %v", err)
	}
	me, err := x.UsersMe()
	if err != nil {
		t.Errorf("Me() retornou um erro: %v", err)
	}
	expectedUsername := "fssouoficial"
	if me.Data.Username != expectedUsername {
		t.Errorf("Expected username %v, but got %v", expectedUsername, me.Data.Username)
	}
}

type MockHttpClient struct {
	GetFunc func(path string, params *http.Params) (*http.Response, error)
}

func (m *MockHttpClient) Get(path string, params *http.Params) (*http.Response, error) {
	return m.GetFunc(path, params)
}
