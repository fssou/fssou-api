package x

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/oauth1/oauth"
	"in.francl.api/services/aws/secretsmanager"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func init() {
	log.Println("X init")
}

func New(ctx context.Context) (*X, error) {
	httpClient := http.DefaultClient
	log.Println("Xray client set")
	var secretsValue Credentials
	secretsName, exists := os.LookupEnv("X_SECRETS_NAME")
	if !exists {
		log.Println("X_SECRETS_NAME not found")
		return nil, fmt.Errorf("X_SECRETS_NAME not found")
	}
	secretsManager := secretsmanager.New(ctx)
	secretsValueString, err := secretsManager.GetSecretValueWithCache(secretsName)
	if err != nil {
		log.Printf("error getting secrets: %v\n", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(secretsValueString), &secretsValue)
	if err != nil {
		log.Printf("error unmarshalling secrets: %v\n", err)
		return nil, err
	}
	return &X{
		secretsValue: secretsValue,
		httpClient:   httpClient,
	}, nil
}

func (x *X) Me() (*TwitterUserMe, error) {
	baseUrl := "https://api.twitter.com/2/users/me"
	params := url.Values{}
	params.Add("user.fields", "created_at,description,entities,id,location,most_recent_tweet_id,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,verified_type,withheld")
	credentialsConsumer := oauth.Credentials{
		Token:  x.secretsValue.ConsumerKey,
		Secret: x.secretsValue.ConsumerSecret,
	}
	client := oauth.Client{
		Credentials:     credentialsConsumer,
		SignatureMethod: oauth.HMACSHA256,
	}

	credentialsToken := oauth.Credentials{
		Token:  x.secretsValue.TokenKey,
		Secret: x.secretsValue.TokenSecret,
	}
	res, err := client.Get(x.httpClient, &credentialsToken, baseUrl, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	var me TwitterUserMe
	err = json.Unmarshal(body, &me)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling body: %v", err)
	}
	return &me, nil
}
