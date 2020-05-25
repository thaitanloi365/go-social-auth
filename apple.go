package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// AppleLoginResponse response
type AppleLoginResponse struct {
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

// AppleLogin provider
type AppleLogin struct {
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
	SkipExpiry bool   `json:"skip_expiry"`
}

// NewAppleLogin new
func NewAppleLogin() *AppleLogin {
	return &AppleLogin{
		Iss: "https://appleid.apple.com",
		Aud: "",
	}
}

// Login login
func (a *AppleLogin) Login(token string) (*AppleLoginResponse, error) {
	var result AppleLoginResponse
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

	if a.Iss != "" {
		if result.Iss != a.Iss {
			return nil, ErrIssuerInvalid
		}
	}

	if a.Aud != "" {
		if result.Aud != a.Aud {
			return nil, ErrAudienceInvalid
		}
	}

	if !a.SkipExpiry {
		if result.Exp < time.Now().Unix() {
			return nil, ErrTokenExpired
		}
	}

	return &result, nil
}
