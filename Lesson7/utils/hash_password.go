package utils

import (
	"crypto/rand"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	s, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(s)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetRand(l, r int64) int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(r-l+1))
	return num.Int64() + l
}
