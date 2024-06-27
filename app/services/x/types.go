package x

import "net/http"

type X struct {
	secretsValue Credentials
	httpClient   *http.Client
}

type TwitterUserMe struct {
	Data Data `json:"data"`
}

type Data struct {
	Entities          Entities `json:"entities"`
	ID                string   `json:"id"`
	MostRecentTweetID string   `json:"most_recent_tweet_id"`
	Name              string   `json:"name"`
	PublicMetrics     Metrics  `json:"public_metrics"`
	VerifiedType      string   `json:"verified_type"`
	URL               string   `json:"url"`
	Username          string   `json:"username"`
	CreatedAt         string   `json:"created_at"`
	Location          string   `json:"location"`
	Verified          bool     `json:"verified"`
	Description       string   `json:"description"`
	ProfileImageURL   string   `json:"profile_image_url"`
	Protected         bool     `json:"protected"`
}

type Entities struct {
	URL         Entity `json:"url"`
	Description Entity `json:"description"`
}

type Entity struct {
	URLs []URL `json:"urls"`
}

type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}

type Metrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
	LikeCount      int `json:"like_count"`
}

type Credentials struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	TokenKey       string `json:"token_key"`
	TokenSecret    string `json:"token_secret"`
}
