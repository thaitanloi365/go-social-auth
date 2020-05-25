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
}

type firebaseLoginResponse struct {
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
	var result firebaseLoginResponse
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

	var response = FirebaseLoginResponse{
		Aud:           result.Aud,
		Sub:           result.Sub,
		Iss:           result.Iss,
		PhoneNumber:   result.PhoneNumber,
		AuthTime:      result.AuthTime,
		CHash:         result.CHash,
		Email:         result.Email,
		EmailVerified: result.EmailVerified,
		Exp:           result.Exp, Iat: result.Iat, NonceSupported: result.NonceSupported,
		UserID:         result.UserID,
		SignInProvider: EmailSignInProvider,
	}

	if result.Firebase.SignInProvider == "phone" {
		response.SignInProvider = PhoneSignInProvider

	}

	if !a.SkipExpiry {
		if response.Exp < time.Now().Unix() {
			return nil, ErrTokenExpired
		}
	}

	return &response, nil
}
