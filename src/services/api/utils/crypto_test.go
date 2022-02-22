package utils

import "testing"

const MockPassword = "my_pass"

func TestBlogchainPasswordCorrectness(t *testing.T) {
	t.Run("it should be correct password hash and verify", func(t *testing.T) {
		hash, err := GeneratePasswordHash(MockPassword)
		if err != nil {
			t.Fatal(err)
		}
		if CompareHashAndPassword(string(hash), MockPassword) != true {
			t.Fatal("failed, the password was expected to be valid")
		}
	})

	t.Run("it should be invalid password and verify also", func(t *testing.T) {
		hash, err := GeneratePasswordHash(MockPassword)
		if err != nil {
			t.Fatal(err)
		}
		if CompareHashAndPassword(string(hash), "my_another_pass") == true {
			t.Fatal("failed, the password was expected to be invalid")
		}
	})
}
