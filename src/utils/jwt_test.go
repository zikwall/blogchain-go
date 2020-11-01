package utils

import (
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	const SECRET = "secret"

	t.Run("it should be success verified JWT signature", func(t *testing.T) {
		claims := TokenClaims{
			UUID: 100,
		}

		createdToken, err := CreateJwtToken(TokenRequiredAttributes{
			Claims:   claims,
			Duration: 100,
			Secret:   SECRET,
		})

		if err != nil {
			t.Fatal(err)
		}

		parsedClaims, err := VerifyJwtToken(createdToken, []byte(SECRET))

		if err != nil {
			t.Fatal(err)
		}

		if parsedClaims.UUID != 100 {
			t.Fatal("Invalid user data")
		}
	})

	t.Run("it should be expired JWT token", func(t *testing.T) {
		claims := TokenClaims{
			UUID: 100,
		}

		createdToken, err := CreateJwtToken(
			TokenRequiredAttributes{
				Claims:   claims,
				Duration: 1,
				Secret:   SECRET,
			},
		)

		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(2000 * time.Millisecond)

		_, err = VerifyJwtToken(createdToken, []byte(SECRET))

		if err == nil {
			t.Fatal("The token should have expired")
		}
	})
}
