package googleauth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	auth "github.com/thaitanloi365/go-social-auth"
)

// SignInProvider provider
type SignInProvider string

// Providers
var (
	Phone SignInProvider = "phone"
	Email SignInProvider = "email"
)

// TokenResponse response
type TokenResponse struct {
	Aud            string         `json:"aud"`
	AuthTime       int64          `json:"auth_time"`
	CHash          string         `json:"c_hash"`
	Email          string         `json:"email"`
	EmailVerified  bool           `json:"email_verified"`
	Exp            int64          `json:"exp"`
	Iat            int64          `json:"iat"`
	Iss            string         `json:"iss"`
	NonceSupported bool           `json:"nonce_supported"`
	Sub            string         `json:"sub"`
	SignInProvider SignInProvider `json:"sign_in_provider"`
	PhoneNumber    string         `json:"phone_number"`
	UserID         string         `json:"user_id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Name           string         `json:"name"`
}

type tokenResponse struct {
	Iss            string `json:"iss"`
	Aud            string `json:"aud"`
	AuthTime       int64  `json:"auth_time"`
	NonceSupported bool   `json:"nonce_supported"`
	CHash          string `json:"c_hash"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
	UserID         string `json:"user_id"`
	Sub            string `json:"sub"`
	Iat            int64  `json:"iat"`
	Exp            int64  `json:"exp"`
	PhoneNumber    string `json:"phone_number"`
	Firebase       struct {
		Identities struct {
			Phone []string `json:"phone"`
		} `json:"identities"`
		SignInProvider string `json:"sign_in_provider"`
	} `json:"firebase"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
}

// Config provider
type Config struct {
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
	SkipExpiry bool   `json:"skip_expiry"`
}

// New new
func New() *Config {
	return &Config{
		Iss: "https://accounts.google.com",
		Aud: "",
	}
}

// WithIssuer override issuer
func (c *Config) WithIssuer(iss string) *Config {
	c.Iss = iss
	return c
}

// WithAudience override audience
func (c *Config) WithAudience(aud string) *Config {
	c.Aud = aud
	return c
}

// WithExpiry override expiry
func (c *Config) WithExpiry(skipExpiry bool) *Config {
	c.SkipExpiry = skipExpiry
	return c
}

// Login login
func (c *Config) Login(token string) (*TokenResponse, error) {
	var result tokenResponse
	var claims = jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, nil)
	if err != nil {

		switch e := err.(type) {
		case *jwt.ValidationError:
			if e.Errors == jwt.ValidationErrorMalformed {
				break
			}
		default:
			return nil, err
		}

	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &result,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(&claims)
	if err != nil {
		return nil, err
	}

	if c.Iss != "" {
		if result.Iss != c.Iss {
			return nil, auth.ErrIssuerInvalid
		}
	}

	if c.Aud != "" {
		if result.Aud != c.Aud {
			return nil, auth.ErrAudienceInvalid
		}
	}

	var response = TokenResponse{
		Aud:            result.Aud,
		Sub:            result.Sub,
		Iss:            result.Iss,
		PhoneNumber:    result.PhoneNumber,
		AuthTime:       result.AuthTime,
		CHash:          result.CHash,
		Email:          result.Email,
		EmailVerified:  result.EmailVerified,
		Exp:            result.Exp,
		Iat:            result.Iat,
		NonceSupported: result.NonceSupported,
		UserID:         result.UserID,
		SignInProvider: Email,
		FirstName:      result.GivenName,
		LastName:       result.FamilyName,
	}

	if result.Firebase.SignInProvider == "phone" {
		response.SignInProvider = Phone

	}

	if !c.SkipExpiry {
		if response.Exp < time.Now().Unix() {
			return nil, auth.ErrTokenExpired
		}
	}

	return &response, nil
}
