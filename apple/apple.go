package apple

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thaitanloi365/go-social-auth/errs"
	"github.com/thaitanloi365/go-social-auth/utils"
)

// TokenResponse response
type TokenResponse struct {
	Aud            string `json:"aud"`
	AuthTime       int64  `json:"auth_time"`
	CHash          string `json:"c_hash"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
	Exp            int64  `json:"exp"`
	Iat            int64  `json:"iat"`
	Iss            string `json:"iss"`
	NonceSupported bool   `json:"nonce_supported"`
	Sub            string `json:"sub"`
}

// Config config
type Config struct {
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
	SkipExpiry bool   `json:"skip_expiry"`
}

// New new
func New() *Config {
	return &Config{
		Iss: "https://appleid.apple.com",
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
	var result TokenResponse
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

	if !c.SkipExpiry {
		if result.Exp < time.Now().Unix() {
			return nil, errs.ErrTokenExpired
		}
	}

	return &result, nil
}
