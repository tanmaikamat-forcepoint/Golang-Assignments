package helper

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = "This is Amazing Key"

type Claims struct {
	UserId  int  `json:userId`
	IsAdmin bool `json:isAdmin`
	jwt.StandardClaims
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CheckHashWithPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetJwtFromData(userId int, isAdmin bool) (string, error) {
	claims := &Claims{UserId: userId, IsAdmin: isAdmin}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	finalToken, err := token.SignedString([]byte(signingKey))
	return finalToken, err
}

func ValidateJwtToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}
