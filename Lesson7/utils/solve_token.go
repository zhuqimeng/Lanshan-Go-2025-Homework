package utils

import (
	"Lesson07/model"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
)

func MakeToken(username string, expTime time.Time) (string, error) {
	snowflakeNode, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	tokenID := snowflakeNode.Generate()
	id := tokenID.String()
	claims := model.TokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "zqm",
			Subject:   id,
			Audience:  []string{"zqm"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	if err != nil || !token.Valid {
		fmt.Println("JWT无效:", err)
		return false, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		fmt.Println("JWT声明:", claims)
		return true, nil
	}
	return false, nil
}
