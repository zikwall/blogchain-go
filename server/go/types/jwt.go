package types

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
