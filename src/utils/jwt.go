package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	TokenClaims struct {
		UUID int64 `json:"uuid"`
		jwt.StandardClaims
	}
	TokenRequiredAttributes struct {
		Claims   TokenClaims
		Duration int64
		Secret   string
	}
)

func CreateJwtToken(attr TokenRequiredAttributes) (string, error) {
	sign := []byte(attr.Secret)

	claims := attr.Claims
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(attr.Duration) * time.Second).Unix(),
		Issuer:    "blogchain-go",
	}

	withClaimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return withClaimsToken.SignedString(sign)
}

func VerifyJwtToken(token string, secret []byte) (*TokenClaims, error) {
	withClaimsToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := withClaimsToken.Claims.(*TokenClaims); ok && withClaimsToken.Valid {
		return claims, nil
	}

	return nil, errors.New("Failed to get the source data from the JWT token")
}
