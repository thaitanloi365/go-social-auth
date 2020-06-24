package facebookauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
	auth "github.com/thaitanloi365/go-social-auth"
)

// TokenResponse response
type TokenResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`
}

// Config config
type Config struct {
	Scopes []string `json:"scopes"`
	URL    string   `json:"url"`
}

// New new
func New() *Config {
	return &Config{
		URL:    "https://graph.facebook.com",
		Scopes: []string{"id", "email", "first_name", "last_name", "name"},
	}
}

// WithURL override url
func (c *Config) WithURL(url string) *Config {
	c.URL = url
	return c
}

// WithScopes override scopes
func (c *Config) WithScopes(scopes []string) *Config {
	c.Scopes = scopes
	return c
}

func isValidFacebookToken(accessToken string) bool {
	// TODO fixed me https://developers.facebook.com/tools/explorer/?method=GET&path=debug_token%3Finput_token%3D%257Binput-token%257D&version=v6.0

	return true
}

// Login login
func (c *Config) Login(accessToken string) (*TokenResponse, error) {
	var result TokenResponse
	if !isValidFacebookToken(accessToken) {
		return nil, auth.ErrTokenInvalid
	}
	var scopes = strings.Join(c.Scopes, ",")
	var url = fmt.Sprintf("%s/me?fields=%s&access_token=%s", c.URL, scopes, url.QueryEscape(accessToken))
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
