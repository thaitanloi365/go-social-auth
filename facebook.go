package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// FacebookLoginResponse response
type FacebookLoginResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`
}

// FacebookLogin provider
type FacebookLogin struct {
	Scope []string `json:"scope"`
	URL   string   `json:"url"`
}

// NewFacebookLogin new
func NewFacebookLogin() *FacebookLogin {
	return &FacebookLogin{
		URL:   "https://graph.facebook.com",
		Scope: []string{"id", "email", "first_name", "last_name", "name"},
	}
}

func isValidFacebookToken(accessToken string) bool {
	// TODO fixed me https://developers.facebook.com/tools/explorer/?method=GET&path=debug_token%3Finput_token%3D%257Binput-token%257D&version=v6.0

	return true
}

// Login login
func (f *FacebookLogin) Login(accessToken string) (*FacebookLoginResponse, error) {
	var result FacebookLoginResponse
	if !isValidFacebookToken(accessToken) {
		return nil, ErrTokenInvalid
	}
	var scope = strings.Join(f.Scope, ",")
	var url = fmt.Sprintf("%s/me?fields=%s&access_token=%s", scope, f.URL, url.QueryEscape(accessToken))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(responseData, &responseMap)
	if err != nil {
		return nil, err
	}

	if value, ok := responseMap["error_description"]; ok {
		return nil, fmt.Errorf("%s", value.(string))
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &result,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(&responseMap)
	if err != nil {
		return nil, err
	}

	return &result, err
}
