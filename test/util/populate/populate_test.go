package populate

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestPopulatedAfterWalk(t *testing.T) {
	RegisterTestingT(t)

	type CertKeyPair struct {
		Key  *rsa.PrivateKey
		Cert *x509.Certificate
	}

	type Bar struct {
		Age          int
		Name         string
		CheckUpdates bool
		Certificates CertKeyPair
	}

	type Foo struct {
		Age  int
		Name string
		Bar  []*Bar
		Tags map[string]string
	}

	foo := &Foo{}
	Walk(foo)
	Expect(foo.Name).To(Not(Or(BeNil(), BeEmpty())))
	Expect(foo.Name).To(Equal("Name"))
	Expect(foo.Age).To(Equal(1))
	Expect(len(foo.Bar)).To(Equal(1))
	for i := 0; i < len(foo.Bar); i++ {
		Expect(foo.Bar[i].Age).To(Equal(1))
		Expect(foo.Bar[i].Name).To(Equal(fmt.Sprintf("Bar[%d].Name", i)))
		Expect(foo.Bar[i].CheckUpdates).To(Equal(true))
		Expect(foo.Bar[i].Certificates).NotTo(Or(BeNil()))
		Expect(*foo.Bar[i].Certificates.Key).To(Equal(DummyPrivateKey()))
	}
	Expect(foo.Tags).To(Not(BeNil()))
	Expect(len(foo.Tags)).To(Equal(1))
	keypath := fmt.Sprintf("Tags.%s", "key")
	valpath := fmt.Sprintf("Tags.%s", "val")
	Expect(foo.Tags[keypath]).To(Equal(valpath))
}
