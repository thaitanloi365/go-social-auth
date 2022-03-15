package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/thaitanloi365/go-social-auth/errs"
	"github.com/thaitanloi365/go-social-auth/utils"
)

type TokenResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccountType string `json:"account_type"`
}

type Picture struct {
	Data PictureData `json:"data"`
}

type PictureData struct {
	Height       float64 `json:"height"`
	IsSilhouette string  `json:"is_silhouette"`
	URL          string  `json:"url"`
	Width        float64 `json:"width"`
}

type debugTokenResponse struct {
	Data struct {
		AppID               string `json:"app_id"`
		Type                string `json:"type"`
		Application         string `json:"application"`
		DataAccessExpiresAt int    `json:"data_access_expires_at"`
		ExpiresAt           int    `json:"expires_at"`
		IsValid             bool   `json:"is_valid"`
		Metadata            struct {
			AuthType string `json:"auth_type"`
		} `json:"metadata"`
		Scopes []string `json:"scopes"`
		UserID string   `json:"user_id"`
	} `json:"data"`
}

type Config struct {
	Scopes    []string `json:"scopes"`
	URL       string   `json:"url"`
	APIURL    string   `json:"api_url"`
	AppID     string   `json:"app_id"`
	AppSecret string   `json:"app_secret"`
}

func New() *Config {
	return &Config{
		URL:    "https://graph.instagram.com",
		APIURL: "https://api.instagram.com",
		Scopes: []string{"id", "account_type", "username"},
	}
}

func (c *Config) WithURL(url string) *Config {
	c.URL = url
	return c
}

func (c *Config) WithAppID(id string) *Config {
	c.AppID = id
	return c
}

func (c *Config) WithAppSecret(secret string) *Config {
	c.AppSecret = secret
	return c
}

func (c *Config) WithScopes(scopes []string) *Config {
	c.Scopes = scopes
	return c
}

func (c *Config) isValidInstagramToken(accessToken string) bool {
	if c.AppID != "" {
		var url = fmt.Sprintf("%s/debug_token?input_token=%s&access_token=%s", c.URL, url.QueryEscape(accessToken), url.QueryEscape(accessToken))
		resp, err := http.Get(url)
		if err != nil {
			return false
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false
		}

		var debugToken debugTokenResponse
		if err = json.Unmarshal(responseData, &debugToken); err == nil {
			return debugToken.Data.AppID == c.AppID
		}

	}
	return true
}

// Login login
func (c *Config) Login(accessToken string) (*TokenResponse, error) {
	var result TokenResponse
	if !c.isValidInstagramToken(accessToken) {
		return nil, errs.ErrTokenInvalid
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

	if value, ok := responseMap["error"].(map[string]interface{}); ok {
		var e Err
		err = utils.DecodeTypedWeakly(&value, &e)
		if err != nil {
			return nil, err
		}

		return nil, &e
	}

	if value, ok := responseMap["error_description"]; ok {
		return nil, fmt.Errorf("%s", value.(string))
	}

	err = utils.DecodeTypedWeakly(&responseMap, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

type GetAccessTokenResponse struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}

// GetAccessToken GetAccessToken
func (c *Config) GetAccessToken(code string, redirectUri string) (*GetAccessTokenResponse, error) {
	var result GetAccessTokenResponse

	var URL = fmt.Sprintf("%s/oauth/access_token", c.APIURL)
	var form = url.Values{
		"client_id":     {c.AppID},
		"client_secret": {c.AppSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirectUri},
	}
	resp, err := http.PostForm(URL, form)

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

	if value, ok := responseMap["code"].(map[string]interface{}); ok {
		var e Err
		err = utils.DecodeTypedWeakly(&value, &e)
		if err != nil {
			return nil, err
		}

		return nil, &e
	}

	if value, ok := responseMap["error_message"]; ok {
		return nil, fmt.Errorf("%s", value.(string))
	}

	err = utils.DecodeTypedWeakly(&responseMap, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}
