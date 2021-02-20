package lib

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
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

const (
	AuthHeaderName = "Authorization"
	AuthTokenType  = "Bearer"
	ClaimsCtxKey   = "claims"
)

func ParseAuthHeader(header string) (string, bool) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || !strings.EqualFold(parts[0], AuthTokenType) {
		return "", false
	}

	return parts[1], true
}

func CreateJwtToken(claims TokenClaims, duration int64, private string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(private))

	if err != nil {
		return "", err
	}

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(duration) * time.Second).Unix(),
		Issuer:    "blogchain-go",
	}

	withClaimsToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)
	return withClaimsToken.SignedString(key)
}

func VerifyJwtToken(token string, r RSA) (*TokenClaims, error) {
	withClaimsToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(
			r.GetPublicKey(),
		))

		if err != nil {
			return nil, err
		}

		if key == nil || key.N == nil {
			return nil, errors.New("JWT token is not defined")
		}

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := withClaimsToken.Claims.(*TokenClaims); ok && withClaimsToken.Valid {
		return claims, nil
	}

	return nil, errors.New("Failed to get the source data from the JWT token")
}
