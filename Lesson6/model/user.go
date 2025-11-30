package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Token     string `json:"token"`
}

type TokenClaims struct {
	Username         string               `json:"username"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}

func (t TokenClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return t.RegisteredClaims.ExpiresAt, nil
}

func (t TokenClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return t.RegisteredClaims.NotBefore, nil
}

func (t TokenClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return t.RegisteredClaims.IssuedAt, nil
}

func (t TokenClaims) GetAudience() (jwt.ClaimStrings, error) {
	return t.RegisteredClaims.Audience, nil
}

func (t TokenClaims) GetIssuer() (string, error) {
	return t.RegisteredClaims.Issuer, nil
}

func (t TokenClaims) GetSubject() (string, error) {
	return t.RegisteredClaims.Subject, nil
}
