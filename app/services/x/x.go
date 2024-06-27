package x

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/gomodule/oauth1/oauth"
	"in.francl.api/services/aws/secretsmanager"
	"io"
	"net/http"
	"net/url"
	"os"
)

var secretsValue Credentials
var httpClient *http.Client

func init() {
	secretsName, exists := os.LookupEnv("X_SECRETS_NAME")
	if !exists {
		panic("X_SECRETS_NAME not found")
	}
	secretsManager := secretsmanager.New(context.Background())
	secretsValueString, err := secretsManager.GetSecretValueWithCache(secretsName)
	if err != nil {
		panic(fmt.Errorf("error getting secrets: %v", err))
	}
	err = json.Unmarshal([]byte(secretsValueString), &secretsValue)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling secrets: %v", err))
	}
	httpClient = xray.Client(http.DefaultClient)
}

func (x *X) Me() (*TwitterUserMe, error) {
	baseUrl := "https://api.twitter.com/2/users/me"
	params := url.Values{}
	params.Add("user.fields", "created_at,description,entities,id,location,most_recent_tweet_id,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,verified_type,withheld")
	credentialsConsumer := oauth.Credentials{
		Token:  secretsValue.ConsumerKey,
		Secret: secretsValue.ConsumerSecret,
	}
	client := oauth.Client{
		Credentials:     credentialsConsumer,
		SignatureMethod: oauth.HMACSHA256,
	}

	credentialsToken := oauth.Credentials{
		Token:  secretsValue.TokenKey,
		Secret: secretsValue.TokenSecret,
	}
	res, err := client.Get(httpClient, &credentialsToken, baseUrl, params)
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
