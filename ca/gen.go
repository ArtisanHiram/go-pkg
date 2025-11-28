package __obf_a7fac689e11862d9

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func __obf_a4cf4e726d10bbf1(__obf_038872964bc0d998 string, __obf_8f425e37b9888280 int) (__obf_137df713f0e246f1 time.Time, __obf_78c8a90e429eb100 time.Time) {
	var __obf_d388542c1816a2c3 error
	if __obf_038872964bc0d998 == "" {
		__obf_137df713f0e246f1 = time.Now()
	} else {
		if __obf_137df713f0e246f1, __obf_d388542c1816a2c3 = time.Parse("2006-01-02", __obf_038872964bc0d998); __obf_d388542c1816a2c3 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_78c8a90e429eb100 = __obf_137df713f0e246f1.Add(time.Duration(__obf_8f425e37b9888280*24) * time.Hour)
		}
	}
	return
}

func __obf_4d38a5fe05015f73() any {
	var __obf_d388542c1816a2c3 error
	var __obf_811665aa5ec324fa any
	switch *__obf_f279281a432f941c {
	case "":
		if *__obf_8f522c929d5b2e06 {
			_, __obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = rsa.GenerateKey(rand.Reader, *__obf_b416844f5e557892)
		}
	case "P224":
		__obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_811665aa5ec324fa, __obf_d388542c1816a2c3 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_f279281a432f941c)
	}

	if __obf_d388542c1816a2c3 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_d388542c1816a2c3)
	}

	return __obf_811665aa5ec324fa
}

func __obf_6372ec731b42de72(__obf_6ef9fadb1f69b363, __obf_b18684a5b76236f2 string, __obf_18bd8dc84f18c6ee []byte, __obf_cb6d6c8e8f097005 os.FileMode) error {

	if _, __obf_d388542c1816a2c3 := os.Stat(__obf_6ef9fadb1f69b363); __obf_d388542c1816a2c3 == nil {
		if __obf_d388542c1816a2c3 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_6ef9fadb1f69b363, string(os.PathSeparator), __obf_b18684a5b76236f2), __obf_18bd8dc84f18c6ee, __obf_cb6d6c8e8f097005); __obf_d388542c1816a2c3 != nil {
			return errors.Wrap(__obf_d388542c1816a2c3, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_d388542c1816a2c3) {
		return errors.Wrap(__obf_d388542c1816a2c3, "Directory not exist.")
	} else {
		return __obf_d388542c1816a2c3
	}
	return nil
}

func __obf_fc034c42d3bb7812(__obf_735f5b74487e5b11 string, __obf_7fc0261a12598c19 *x509.Certificate, __obf_a264d37264bcfd3c any, __obf_e21f6160f7309bde *x509.Certificate, __obf_e29c547489483b7f any) {

	var __obf_17b876ebc9671e6f any
	switch __obf_4c4934549b3772cd := __obf_a264d37264bcfd3c.(type) {
	case *rsa.PrivateKey:
		__obf_17b876ebc9671e6f = &__obf_4c4934549b3772cd.PublicKey
	case *ecdsa.PrivateKey:
		__obf_17b876ebc9671e6f = &__obf_4c4934549b3772cd.PublicKey
	case ed25519.PrivateKey:
		__obf_17b876ebc9671e6f = __obf_4c4934549b3772cd.Public().(ed25519.PublicKey)
	default:
		__obf_17b876ebc9671e6f = nil
	}

	var __obf_85a14dae0b94c317 any
	if __obf_e29c547489483b7f != nil {
		__obf_85a14dae0b94c317 = __obf_e29c547489483b7f
	} else {
		__obf_85a14dae0b94c317 = __obf_a264d37264bcfd3c
	}
	__obf_edd4a003880f9f62, __obf_d388542c1816a2c3 := x509.CreateCertificate(rand.Reader, __obf_7fc0261a12598c19, __obf_e21f6160f7309bde, __obf_17b876ebc9671e6f, __obf_85a14dae0b94c317)
	if __obf_d388542c1816a2c3 != nil {
		log.Println("create failed", __obf_d388542c1816a2c3)
		return
	}

	__obf_1b629f5ba33ef685 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_edd4a003880f9f62})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_d388542c1816a2c3 = __obf_6372ec731b42de72(*__obf_d04fc7665a661d58, __obf_735f5b74487e5b11+".pem", __obf_1b629f5ba33ef685, 0777); __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}

	__obf_4a34056e5abfcd15, __obf_d388542c1816a2c3 := x509.MarshalPKCS8PrivateKey(__obf_a264d37264bcfd3c)
	if __obf_d388542c1816a2c3 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_c2d7dd13537faa2d := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_4a34056e5abfcd15})

	if __obf_d388542c1816a2c3 = __obf_6372ec731b42de72(*__obf_d04fc7665a661d58, __obf_735f5b74487e5b11+".key", __obf_c2d7dd13537faa2d, 0777); __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}
}

// 公钥证书解析
func __obf_b5b73174ba0a6b2a() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_d223444df67bdb8e, __obf_d388542c1816a2c3 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_d04fc7665a661d58, string(os.PathSeparator))); __obf_d388542c1816a2c3 != nil {
		return nil, errors.Wrap(__obf_d388542c1816a2c3, "root.pem: Failed to read file.")
	} else {
		__obf_f56d695709096d01, __obf_df9437020cbcb6eb := pem.Decode(__obf_d223444df67bdb8e)
		if __obf_f56d695709096d01 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_df9437020cbcb6eb)
		}

		return x509.ParseCertificate(__obf_f56d695709096d01.Bytes)
	}
}

// 私钥解析
func __obf_33c0f9766cdb6954() (__obf_a264d37264bcfd3c any, __obf_d388542c1816a2c3 error) {
	if __obf_d223444df67bdb8e, __obf_d388542c1816a2c3 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_d04fc7665a661d58, string(os.PathSeparator))); __obf_d388542c1816a2c3 != nil {
		return nil, errors.Wrap(__obf_d388542c1816a2c3, "root.key: Failed to read file.")
	} else {
		__obf_f56d695709096d01, __obf_df9437020cbcb6eb := pem.Decode(__obf_d223444df67bdb8e)
		if __obf_f56d695709096d01 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_df9437020cbcb6eb)
		}

		return x509.ParsePKCS8PrivateKey(__obf_f56d695709096d01.Bytes)
	}
}

func __obf_f12470e1eee058d9() *big.Int {
	__obf_e7713dfa88d0d4b1 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_228301b9ee95546a, __obf_d388542c1816a2c3 := rand.Int(rand.Reader, __obf_e7713dfa88d0d4b1); __obf_d388542c1816a2c3 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_228301b9ee95546a
	}
}

func GenRootCA() {

	__obf_8e590bf093409809, __obf_8939e26b29a2a15c := __obf_a4cf4e726d10bbf1(*__obf_115c15a762a4cfbe, *__obf_64b88161bfe6ddb6)
	__obf_f4caa47498522a63 := &x509.Certificate{
		SerialNumber: __obf_f12470e1eee058d9(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_b0f27e785a5d7613},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_8e590bf093409809,
		NotAfter:              __obf_8939e26b29a2a15c,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_c3f16a5d19c0b5bc := __obf_4d38a5fe05015f73()

	__obf_fc034c42d3bb7812("root", __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc, __obf_f4caa47498522a63, nil)

}

func GenServerCA() {

	__obf_8e590bf093409809, __obf_8939e26b29a2a15c := __obf_a4cf4e726d10bbf1(*__obf_e85a2925f0461119, *__obf_f0a9e6b00a4f1eea)
	__obf_f98631a6dc322cfb := &x509.Certificate{
		SerialNumber: __obf_f12470e1eee058d9(),
		Subject: pkix.Name{
			Organization: []string{*__obf_8c0bc0ba949334fc},
		},
		NotBefore: __obf_8e590bf093409809,
		NotAfter:  __obf_8939e26b29a2a15c,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_51f4b0d669777373 := strings.Split(*__obf_757dc2f64e06a3fa, ",")

	for _, __obf_4def855968424027 := range __obf_51f4b0d669777373 {
		if __obf_86931e0d9c3a793c := net.ParseIP(__obf_4def855968424027); __obf_86931e0d9c3a793c != nil {
			__obf_f98631a6dc322cfb.IPAddresses = append(__obf_f98631a6dc322cfb.IPAddresses, __obf_86931e0d9c3a793c)
		} else {
			__obf_f98631a6dc322cfb.DNSNames = append(__obf_f98631a6dc322cfb.DNSNames, __obf_4def855968424027)
		}
	}
	__obf_f5e6430087d57a6f := __obf_4d38a5fe05015f73()

	var __obf_f4caa47498522a63 *x509.Certificate
	var __obf_d388542c1816a2c3 error
	var __obf_c3f16a5d19c0b5bc any

	__obf_f4caa47498522a63, __obf_d388542c1816a2c3 = __obf_b5b73174ba0a6b2a()
	if __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}

	__obf_c3f16a5d19c0b5bc, __obf_d388542c1816a2c3 = __obf_33c0f9766cdb6954()
	if __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}

	__obf_fc034c42d3bb7812("server", __obf_f98631a6dc322cfb, __obf_f5e6430087d57a6f, __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc)
}

func GenClientCA() {

	__obf_8e590bf093409809, __obf_8939e26b29a2a15c := __obf_a4cf4e726d10bbf1(*__obf_7497a241ee847b27, *__obf_6558bea5d760a593)
	__obf_440c82f5cba2bc11 := &x509.Certificate{
		SerialNumber: __obf_f12470e1eee058d9(),
		Subject: pkix.Name{
			Organization: []string{*__obf_0e47c303a8f77c9f},
		},
		NotBefore: __obf_8e590bf093409809,
		NotAfter:  __obf_8939e26b29a2a15c,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_02cb005039f8f182 := __obf_4d38a5fe05015f73()

	var __obf_f4caa47498522a63 *x509.Certificate
	var __obf_d388542c1816a2c3 error
	var __obf_c3f16a5d19c0b5bc any

	__obf_f4caa47498522a63, __obf_d388542c1816a2c3 = __obf_b5b73174ba0a6b2a()
	if __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}

	__obf_c3f16a5d19c0b5bc, __obf_d388542c1816a2c3 = __obf_33c0f9766cdb6954()
	if __obf_d388542c1816a2c3 != nil {
		log.Fatalln(__obf_d388542c1816a2c3)
	}

	__obf_fc034c42d3bb7812("client", __obf_440c82f5cba2bc11, __obf_02cb005039f8f182, __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc)

}
