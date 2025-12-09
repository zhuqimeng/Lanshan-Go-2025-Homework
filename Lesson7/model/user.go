package model

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type QueryRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	Token    string `json:"token"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64;unique;not null;comment:用户名"`
	Password string `json:"password" gorm:"size:128;not null;comment:密码哈希"` // 不返回密码
	Token    string `json:"token" gorm:"size:512;default:null;comment:认证令牌"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=6,max=128"`
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
