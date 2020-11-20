package lib

import (
	"github.com/zikwall/blogchain/src/constants"
	"strings"
)

// ToDo: Create service for automatic synchronize & update RSA public keys
// ToDo: Check valid public key, if not valid -> update from service
type (
	RSA interface {
		GetPublicKey() string
		SetPublicKey(string)

		GetPrivateKey() string
		SetPrivateKey(string)
	}
	RSAContainer struct {
		publicKey  string
		privateKey string
	}
)

func NewBlogchainRSAContainer(public, private string) RSAContainer {
	rsa := RSAContainer{}
	rsa.SetPublicKey(public)
	rsa.SetPrivateKey(private)

	return rsa
}

func (r RSAContainer) GetPublicKey() string {
	return r.publicKey
}

func (r *RSAContainer) SetPublicKey(key string) {
	r.publicKey = key
}

func (r RSAContainer) GetPrivateKey() string {
	return r.privateKey
}

func (r *RSAContainer) SetPrivateKey(key string) {
	r.privateKey = key
}

type MockRSA struct{}

func (r MockRSA) GetPublicKey() string {
	// It's possible that your "\n" is actually the escaped version of a line break character.
	// You can replace these with real line breaks by searching for the escaped version
	// and replacing with the non escaped version
	escapedVersion := func(s string) string {
		return strings.ReplaceAll(s, `\n`, "\n")
	}

	publicKey := constants.TestPublicKey

	return escapedVersion(publicKey)
}

func (r MockRSA) GetPrivateKey() string {
	return constants.TestPrivateKey
}
func (r *MockRSA) SetPublicKey(key string)  {}
func (r *MockRSA) SetPrivateKey(key string) {}
