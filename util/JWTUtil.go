package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"majooTest/dto"
	"majooTest/log"
	"time"
)

type Claims struct {
	dto.LoginDTO
	jwt.StandardClaims
}

func GenerateToken(username string, secretKey string) (string, error) {
	log.Info("username : ", username)

	expirationTime := time.Now().Add(10 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		LoginDTO: dto.LoginDTO{
			Username: username,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Info("token : ", token)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Error("Error", errors.Unwrap(err))
		return "", err
	}
	log.Info("Token created : ", tokenString)

	return tokenString, nil
}

func ValidateToken(tokenFromRequest string, secretKey string) (bool, *Claims, string) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenFromRequest, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			return false, claims, err.Error()
		}
		return false, claims, err.Error()
	}
	if !tkn.Valid {
		return false, claims, "Token not valid"
	}

	return true, claims, "Token is valid"
}
