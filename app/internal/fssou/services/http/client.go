package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gomodule/oauth1/oauth"
)

type client struct {
	ctx              context.Context
	baseURL          *url.URL
	httpClient       *http.Client
	oauthClient      *oauth.Client
	credentialsToken *oauth.Credentials
}

func New(ctx context.Context, baseURL string, credentials *CredentialsOAuth1) (Client, error) {
	defaultHttpClient := http.DefaultClient
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	credentialsConsumer := oauth.Credentials{
		Token:  credentials.ConsumerKey,
		Secret: credentials.ConsumerSecret,
	}
	credentialsToken := oauth.Credentials{
		Token:  credentials.TokenKey,
		Secret: credentials.TokenSecret,
	}
	oauthClient := oauth.Client{
		Credentials:     credentialsConsumer,
		SignatureMethod: oauth.HMACSHA256,
	}
	return &client{
		ctx:              ctx,
		baseURL:          base,
		httpClient:       defaultHttpClient,
		oauthClient:      &oauthClient,
		credentialsToken: &credentialsToken,
	}, nil
}

func (c *client) Get(path string, params *Params) (*Response, error) {
	paramsUrl := url.Values{}
	for k, v := range *params {
		paramsUrl.Add(k, v)
	}
	relative, err := url.Parse(path)
	if err != nil {
		fmt.Println("Error parsing path:", err)
		return nil, fmt.Errorf("error parsing path: %v", err)
	}
	fullURL := c.baseURL.ResolveReference(relative)
	res, err := c.oauthClient.Get(c.httpClient, c.credentialsToken, fullURL.String(), paramsUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Println("error closing body")
		}
	}(res.Body)
	return &Response{
		Headers:    res.Header,
		StatusCode: res.StatusCode,
		Body:       body,
	}, nil
}
