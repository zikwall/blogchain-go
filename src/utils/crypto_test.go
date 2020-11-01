package utils

import "testing"

func TestBlogchainPasswordCorrectness(t *testing.T) {
	t.Run("it should be correct password hash and verify", func(t *testing.T) {
		password := "my_pass"
		hash, err := GenerateBlogchainPasswordHash(password)

		if err != nil {
			t.Fatal(err)
		}

		if BlogchainPasswordCorrectness(string(hash), password) != true {
			t.Fatal("Error, the password was expected to be valid")
		}
	})

	t.Run("it shoudl be invalid password and verify also", func(t *testing.T) {
		password := "my_pass"
		hash, err := GenerateBlogchainPasswordHash(password)

		if err != nil {
			t.Fatal(err)
		}

		if BlogchainPasswordCorrectness(string(hash), "my_another_pass") == true {
			t.Fatal("Error, the password was expected to be invalid")
		}
	})
}
