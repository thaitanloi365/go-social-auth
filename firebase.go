package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// SignInProvider provider
type SignInProvider string

// Providers
var (
	PhoneSignInProvider SignInProvider = "phone"
	EmailSignInProvider SignInProvider = "email"
)

// FirebaseLoginResponse response
type FirebaseLoginResponse struct {
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
	FirebaseLoginPhoneResponse
}

// FirebaseLoginInfo params
type FirebaseLoginInfo struct {
	SignInProvider SignInProvider          `json:"sign_in_provider"`
	Identities     FirebaseLoginIdentities `json:"identities"`
}

// FirebaseLoginIdentities params
type FirebaseLoginIdentities struct {
	Phone []string `json:"phone"`
}

// FirebaseLoginPhoneResponse params
type FirebaseLoginPhoneResponse struct {
	UserID            string             `json:"user_id"`
	PhoneNumber       string             `json:"phone_number"`
	FirebaseLoginInfo *FirebaseLoginInfo `json:"firebase"`
}

// FirebaseLogin provider
type FirebaseLogin struct {
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
	SkipExpiry bool   `json:"skip_expiry"`
}

// NewFirebaseLogin new
func NewFirebaseLogin() *FirebaseLogin {
	return &FirebaseLogin{
		Iss: "https://accounts.google.com",
		Aud: "",
	}
}

// Login login
func (a *FirebaseLogin) Login(token string) (*FirebaseLoginResponse, error) {
	var result FirebaseLoginResponse
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

	err = mapstructure.Decode(claims, &result)
	if err != nil {
		return nil, err
	}

	if result.Iss != a.Iss {
		return nil, ErrIssuerInvalid
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
