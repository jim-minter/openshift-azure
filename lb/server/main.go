package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func generateCertAndKey() error {
	now := time.Now()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	req := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		NotBefore:             now,
		NotAfter:              now.AddDate(1, 0, 0),
		Subject:               pkix.Name{CommonName: "localhost"},
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	cert, err := x509.CreateCertificate(rand.Reader, req, req, &key.PublicKey, key)
	if err != nil {
		return err
	}
	fmt.Print(string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})))
	fmt.Print(string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})))
	return nil
}

var (
	key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA26XjX8xVgoEYMkG5weE/TPDq3v8OnlKNHjUMgbx2eJ9GH0Mn
nHHBXRhZkurWFMsCPHmf4Cqn9sXsQ4PKSZP/lTQ9NtyPSAh/SxKRb8JgStGcyBnW
jN+Jz76CEdy44HxO5I9tDTGZmJ6ca4BSLg5MuGEk9lp34frFWpho15tBhcPOZ4PE
JBRLxs5j+R3cgHdqX/DQ+v4L0oluD5xGG9DvQs3XSl1BgeGEPNudhrp3uUuwh1rp
uVFYv4d2/jtyRnfDtp+3bxQoRhW5lINsxQEbh9vpDAhlhfKosF4NU7LJQB0pNM4t
zQGlCS1PC0o6ZtMpuHwQQsoOSUjvHInvB95JVwIDAQABAoIBAQC2O1KCt3+2T66o
e5lHPr8K8dKbgpc5SZolFrQyqw7LkrFV3Jxvkn1v5HTkjItjIu7PB8VZ8Wn7NkiH
1z6sfuqMepPTAXiqtcoOmfAp/eVwDap65dz4cbnfrtoxQaPtM5Us0cYTLTSWx/lU
w1jrNxf13TsSXQqbZTf5qvtI7lmVRM2/PdFtZwKPZUEZiP614siY5sF6RuMVjG5M
vp8QScnjcr0Yyq0d43dTxYarMK3mtBZlMC6rLvUijS6hHZLWo4+C3mSAxQAe9llu
nuWzRpCznUE2sIfFSQS4qs20jIEAMNoohkcB2pMBrfL0el5qRWJCTfSsYSx9EdQ0
pryiJLBxAoGBAN9/jpsgPxCtAsh45vWrV0Qvcew35TC/d9lnto68rDKzCXaocTwG
kBfeTGm1G+3c3SA7+w/ccARLkRdGnFTpDpSI70TJ/oJ//8GxtiK8eaGizUtxvRZN
kY5HYSGQprkGILQWFgjiX2uoBob5LaU1SiymNTNfxDh/8PNBqaCgh74pAoGBAPuW
/hHZeJUr+rctPusrWrLE2GUi4R90YWPy0Jw5eBI7MfMEmQ7trsNDXYryHSDDLB7s
hk3A//gwNapnAHKsD07ykFrYE7B/0DzbVnw1PyJss3q9FZwsF4N/wSORF4jj4z0X
JX12oT2jpVhKeONPK2o4Ax3K8OOMJ3gR1ePth7t/AoGBANPbCyfa8kzxY0D68huP
9mHJA5liBpwl8wqfODqneCd69Q6IbwXyRqaJby+IoNfh065ZjQwk7f30T62bnlcS
sGJ2RzCStPGpOZu2xCq7NCTWuPm57/5zOvV+jgEOKCwdNeTfRrXXN5JKLR3Gl9ER
6aTXTHjNX6gbByDfbla3tNS5AoGBAPGOEFR881xuBGMZKv7J+mQHwSiha3oy2EsJ
WCeWueTvNs74TChcJl5N8KM2QKczHMp4F57RvjHBv9Ti3jg7YNtQ4y6FpanhncLA
aPIKgZqAuXYP047FerID2CFY7jq9anE+Jv2mB7vRwi/aGOVOHwX1z3AsaEphR4ft
v+n+JkLrAoGAMGXybHl61jjxTPbpjyaXpxuMYgnvf/a0v6N1vzjU6c2/Qj2hwLpT
7nASKVPrl0Wmb40oQ1mV/RVTykfWK+cZw/jmc6hn0RSeiXBXz0laIufYTL2XZt7N
+3sRLSjugIJ3ouDtEAcKvG9ndxSowqEQ+UjA3dsp5fTOCciwHMGR9NU=
-----END RSA PRIVATE KEY-----`)
	cert = []byte(`-----BEGIN CERTIFICATE-----
MIIC2DCCAcCgAwIBAgIBATANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDEwlsb2Nh
bGhvc3QwHhcNMTkwMzI4MTI0OTU1WhcNMjAwMzI4MTI0OTU1WjAUMRIwEAYDVQQD
Ewlsb2NhbGhvc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDbpeNf
zFWCgRgyQbnB4T9M8Ore/w6eUo0eNQyBvHZ4n0YfQyecccFdGFmS6tYUywI8eZ/g
Kqf2xexDg8pJk/+VND023I9ICH9LEpFvwmBK0ZzIGdaM34nPvoIR3LjgfE7kj20N
MZmYnpxrgFIuDky4YST2Wnfh+sVamGjXm0GFw85ng8QkFEvGzmP5HdyAd2pf8ND6
/gvSiW4PnEYb0O9CzddKXUGB4YQ8252Gune5S7CHWum5UVi/h3b+O3JGd8O2n7dv
FChGFbmUg2zFARuH2+kMCGWF8qiwXg1TsslAHSk0zi3NAaUJLU8LSjpm0ym4fBBC
yg5JSO8cie8H3klXAgMBAAGjNTAzMA4GA1UdDwEB/wQEAwIFoDATBgNVHSUEDDAK
BggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQBu4lDm
g8UMCDLbCcBQxnOx5zh02+uXo0Rvjnj0ECZ7GvDz9pMueDLTN7wPadkLk5hU0auc
VKRM4wMDDwTE+y7TkwCwXp0Faa3/BfG2XfuNdeBCG2B0ccN3BtybbzpCaMlgH5kz
HRPe4T84P+soR7fjx15jOpNLK5PsDIrz/LxlEMlTD5IaTyLgsIOrPq7wXDp5HIn3
BsNG9rCxdDQRNYTSKkA3jWGQYflz5p5romuwvnDXGPWuZ+dWmUAnt4G8lXKmqhu+
5hLLY8/dZyAZpkkt7Nl4c03pHHSQwc1GJS/4OuSChm2KCPbeIJFBcL6hbUCs80KS
4jxD9yG4zNM1inve
-----END CERTIFICATE-----`)
)

func run() error {
	cert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return err
	}
	l, err := tls.Listen("tcp", ":443", &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		return err
	}
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())
	}))
	return http.Serve(l, nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
