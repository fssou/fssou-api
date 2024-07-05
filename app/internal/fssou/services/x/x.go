package x

import (
	"context"
	"encoding/json"
	"fmt"
	"in.francl.api/internal/fssou/services/http"
)

type X struct {
	credentials *Credentials
	baseURL     string
	httpClient  http.Client
}

func New(ctx context.Context, credentials *Credentials, httpClient http.Client) (*X, error) {
	baseURL := "https://api.twitter.com/2"
	if httpClient == nil {
		var err interface{} = nil
		httpClient, err = http.New(ctx, baseURL, &http.CredentialsOAuth1{
			ConsumerKey:    credentials.ConsumerKey,
			ConsumerSecret: credentials.ConsumerSecret,
			TokenKey:       credentials.TokenKey,
			TokenSecret:    credentials.TokenSecret,
		})
		if err != nil {
			return nil, fmt.Errorf("error creating http client: %v", err)
		}
	}
	return &X{
		baseURL:     baseURL,
		credentials: credentials,
		httpClient:  httpClient,
	}, nil
}

func (x *X) UsersMe() (*TwitterUserMe, error) {
	path := "/users/me"
	params := http.Params{
		"user.fields": "created_at,description,entities,id,location,most_recent_tweet_id,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,verified_type,withheld",
	}
	res, err := x.httpClient.Get(path, &params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	var me TwitterUserMe
	err = json.Unmarshal(res.Body, &me)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling body: %v", err)
	}
	return &me, nil
}
