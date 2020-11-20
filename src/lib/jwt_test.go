package lib

import (
	"github.com/zikwall/blogchain/src/constants"
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	t.Run("it should verify", func(t *testing.T) {
		claims := TokenClaims{
			UUID: 100,
		}

		createdToken, err := CreateJwtToken(claims, 99999999, constants.TestPrivateKey)

		if err != nil {
			t.Fatal(err)
		}

		parsedClaims, err := VerifyJwtToken(createdToken, &MockRSA{})

		if err != nil {
			t.Fatal(err)
		}

		if parsedClaims.UUID != 100 {
			t.Fatal("Invalid user data")
		}
	})

	t.Run("it should expired", func(t *testing.T) {
		claims := TokenClaims{
			UUID: 100,
		}

		createdToken, err := CreateJwtToken(claims, 1, constants.TestPrivateKey)

		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(2000 * time.Millisecond)

		_, err = VerifyJwtToken(createdToken, &MockRSA{})

		if err == nil {
			t.Fatal("The token should have expired")
		}
	})
}
