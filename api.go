package hsapi

import (
	"net/http"
)

//HsAPI is the main struct for our API
type HsAPI struct {
	token  string       `json:"token"`
	Debug  bool         `json:"debug"`
	client *http.Client `json:"-"`
}

//NewHsAPI creates a new API with a token
func NewHsAPI(token string) *HsAPI {
	return NewHsAPIWithClient(token, &http.Client{})
}

//NewHsAPIWithClient returns a new client with a custom client
func NewHsAPIWithClient(token string, client *http.Client) *HsAPI {
	return &HsAPI{
		token:  token,
		client: client,
	}
}
