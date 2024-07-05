package http

type Params map[string]string

type Client interface {
	Get(path string, params *Params) (*Response, error)
}

type Response struct {
	Headers    map[string][]string
	StatusCode int
	Body       []byte
}

type CredentialsOAuth1 struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	TokenKey       string `json:"token_key"`
	TokenSecret    string `json:"token_secret"`
}
