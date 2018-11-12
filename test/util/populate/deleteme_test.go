package populate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 64)
	bs := x509.MarshalPKCS1PrivateKey(key)

	for _, b := range bs {
		fmt.Printf("0x%02x, ", b)
	}
}
