package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	//address := pubKey.Address()

	msg := []byte("Hello World")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	fmt.Println(sig)
	fmt.Println(sig.Verify(pubKey, msg))
}
