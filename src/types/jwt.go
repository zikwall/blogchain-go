package types

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zikwall/blogchain/src/models/user"
	"time"
)

type TokenClaims struct {
	UUID int64 `json:"uuid"`
	jwt.StandardClaims
}

func CreateToken(u *user.User) (string, error) {
	mySigningKey := []byte("secret")

	claims := TokenClaims{
		UUID: u.GetId(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1000 * time.Second).Unix(),
			Issuer:    "blogchain-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
