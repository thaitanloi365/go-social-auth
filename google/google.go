package google

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thaitanloi365/go-social-auth/errs"
	"github.com/thaitanloi365/go-social-auth/utils"
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
	Picture        string         `json:"picture"`
	Locale         string         `json:"vi"`
	AtHash         string         `json:"at_hash"`
}

type tokenResponse struct {
	// Login mail
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           int64  `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           int64  `json:"iat"`
	Iss           string `json:"iss"`
	Jti           string `json:"jti"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`

	// Login Phone
	AuthTime       int64  `json:"auth_time"`
	NonceSupported bool   `json:"nonce_supported"`
	CHash          string `json:"c_hash"`
	UserID         string `json:"user_id"`
	PhoneNumber    string `json:"phone_number"`
	Firebase       struct {
		Identities struct {
			Phone []string `json:"phone"`
		} `json:"identities"`
		SignInProvider string `json:"sign_in_provider"`
	} `json:"firebase"`
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

	err = utils.DecodeTypedWeakly(&claims, &result)
	if err != nil {
		return nil, err
	}

	if c.Iss != "" {
		if result.Iss != c.Iss {
			return nil, errs.ErrIssuerInvalid
		}
	}

	if c.Aud != "" {
		if result.Aud != c.Aud {
			return nil, errs.ErrAudienceInvalid
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
		Name:           result.Name,
		Picture:        result.Picture,
		Locale:         result.Locale,
		AtHash:         result.AtHash,
	}

	if result.Firebase.SignInProvider == "phone" {
		response.SignInProvider = Phone

	}

	if !c.SkipExpiry {
		if response.Exp < time.Now().Unix() {
			return nil, errs.ErrTokenExpired
		}
	}

	return &response, nil
}
