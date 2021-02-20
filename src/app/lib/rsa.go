package lib

import (
	"github.com/zikwall/blogchain/src/app/utils"
	"github.com/zikwall/blogchain/src/platform/constants"
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
	r.publicKey = utils.EscapeNewLine(key)
}

func (r RSAContainer) GetPrivateKey() string {
	return r.privateKey
}

func (r *RSAContainer) SetPrivateKey(key string) {
	r.privateKey = utils.EscapeNewLine(key)
}

type MockRSA struct{}

func (r MockRSA) GetPublicKey() string {
	return constants.TestPublicKey
}

func (r MockRSA) GetPrivateKey() string {
	return constants.TestPrivateKey
}
func (r *MockRSA) SetPublicKey(key string)  {}
func (r *MockRSA) SetPrivateKey(key string) {}
