package __obf_b8dd0558a7deb9fa

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

func __obf_397b631981cf429a(__obf_b560d970db554348 string, __obf_051c2a02a505f856 int) (__obf_8b1b13dc542d306a time.Time, __obf_50cb4893ba0dc222 time.Time) {
	var __obf_b7fd5848550ca60c error
	if __obf_b560d970db554348 == "" {
		__obf_8b1b13dc542d306a = time.Now()
	} else {
		if __obf_8b1b13dc542d306a, __obf_b7fd5848550ca60c = time.Parse("2006-01-02", __obf_b560d970db554348); __obf_b7fd5848550ca60c != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_50cb4893ba0dc222 = __obf_8b1b13dc542d306a.Add(time.Duration(__obf_051c2a02a505f856*24) * time.Hour)
		}
	}
	return
}

func __obf_cfda8bfc21b81e73() any {
	var __obf_b7fd5848550ca60c error
	var __obf_200b0d942eda1a2a any
	switch *__obf_694c6c8758e0d040 {
	case "":
		if *__obf_eb00d4b9a27af664 {
			_, __obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = rsa.GenerateKey(rand.Reader, *__obf_911ff0576318f3c7)
		}
	case "P224":
		__obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_200b0d942eda1a2a, __obf_b7fd5848550ca60c = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_694c6c8758e0d040)
	}

	if __obf_b7fd5848550ca60c != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_b7fd5848550ca60c)
	}

	return __obf_200b0d942eda1a2a
}

func __obf_495d26f1466946e6(__obf_0fb227cea3c86a97, __obf_1c535ed6bfeaf860 string, __obf_81fdbcc13e296a9d []byte, __obf_eb63717f140f6061 os.FileMode) error {

	if _, __obf_b7fd5848550ca60c := os.Stat(__obf_0fb227cea3c86a97); __obf_b7fd5848550ca60c == nil {
		if __obf_b7fd5848550ca60c = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_0fb227cea3c86a97, string(os.PathSeparator), __obf_1c535ed6bfeaf860), __obf_81fdbcc13e296a9d, __obf_eb63717f140f6061); __obf_b7fd5848550ca60c != nil {
			return errors.Wrap(__obf_b7fd5848550ca60c, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_b7fd5848550ca60c) {
		return errors.Wrap(__obf_b7fd5848550ca60c, "Directory not exist.")
	} else {
		return __obf_b7fd5848550ca60c
	}
	return nil
}

func __obf_f3ff9550e4d47e55(__obf_e6d878397dd86eee string, __obf_35143b3a01b9448b *x509.Certificate, __obf_568d273391419dd1 any, __obf_751c8c53bc81fe11 *x509.Certificate, __obf_aa6de849c46a0f6e any) {

	var __obf_3d0c48eca5fc749e any
	switch __obf_0c573575c1a11316 := __obf_568d273391419dd1.(type) {
	case *rsa.PrivateKey:
		__obf_3d0c48eca5fc749e = &__obf_0c573575c1a11316.PublicKey
	case *ecdsa.PrivateKey:
		__obf_3d0c48eca5fc749e = &__obf_0c573575c1a11316.PublicKey
	case ed25519.PrivateKey:
		__obf_3d0c48eca5fc749e = __obf_0c573575c1a11316.Public().(ed25519.PublicKey)
	default:
		__obf_3d0c48eca5fc749e = nil
	}

	var __obf_00af8fb2c1c6de80 any
	if __obf_aa6de849c46a0f6e != nil {
		__obf_00af8fb2c1c6de80 = __obf_aa6de849c46a0f6e
	} else {
		__obf_00af8fb2c1c6de80 = __obf_568d273391419dd1
	}
	__obf_df65b962ca4d67b5, __obf_b7fd5848550ca60c := x509.CreateCertificate(rand.Reader, __obf_35143b3a01b9448b, __obf_751c8c53bc81fe11, __obf_3d0c48eca5fc749e, __obf_00af8fb2c1c6de80)
	if __obf_b7fd5848550ca60c != nil {
		log.Println("create failed", __obf_b7fd5848550ca60c)
		return
	}

	__obf_a7c70595d99cbd5a := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_df65b962ca4d67b5})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_b7fd5848550ca60c = __obf_495d26f1466946e6(*__obf_17f9ffe223fe4d58, __obf_e6d878397dd86eee+".pem", __obf_a7c70595d99cbd5a, 0777); __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}

	__obf_8bf120e25d7da167, __obf_b7fd5848550ca60c := x509.MarshalPKCS8PrivateKey(__obf_568d273391419dd1)
	if __obf_b7fd5848550ca60c != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_41614e79cb0c56d3 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_8bf120e25d7da167})

	if __obf_b7fd5848550ca60c = __obf_495d26f1466946e6(*__obf_17f9ffe223fe4d58, __obf_e6d878397dd86eee+".key", __obf_41614e79cb0c56d3, 0777); __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}
}

// 公钥证书解析
func __obf_85687d92b145c0ba() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_b1506d0d8917e964, __obf_b7fd5848550ca60c := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_17f9ffe223fe4d58, string(os.PathSeparator))); __obf_b7fd5848550ca60c != nil {
		return nil, errors.Wrap(__obf_b7fd5848550ca60c, "root.pem: Failed to read file.")
	} else {
		__obf_211bf31559b8aacb, __obf_7b662d5ef634c62b := pem.Decode(__obf_b1506d0d8917e964)
		if __obf_211bf31559b8aacb == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_7b662d5ef634c62b)
		}

		return x509.ParseCertificate(__obf_211bf31559b8aacb.Bytes)
	}
}

// 私钥解析
func __obf_242aaffa404c17b0() (__obf_568d273391419dd1 any, __obf_b7fd5848550ca60c error) {
	if __obf_b1506d0d8917e964, __obf_b7fd5848550ca60c := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_17f9ffe223fe4d58, string(os.PathSeparator))); __obf_b7fd5848550ca60c != nil {
		return nil, errors.Wrap(__obf_b7fd5848550ca60c, "root.key: Failed to read file.")
	} else {
		__obf_211bf31559b8aacb, __obf_7b662d5ef634c62b := pem.Decode(__obf_b1506d0d8917e964)
		if __obf_211bf31559b8aacb == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_7b662d5ef634c62b)
		}

		return x509.ParsePKCS8PrivateKey(__obf_211bf31559b8aacb.Bytes)
	}
}

func __obf_f2edc2d35a45d523() *big.Int {
	__obf_0e2e6c5949d43ea0 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_dfa1d0e76a164424, __obf_b7fd5848550ca60c := rand.Int(rand.Reader, __obf_0e2e6c5949d43ea0); __obf_b7fd5848550ca60c != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_dfa1d0e76a164424
	}
}

func GenRootCA() {

	__obf_66968245163fa68a, __obf_bfe877faf90c2d1d := __obf_397b631981cf429a(*__obf_4beb52f8769f794e, *__obf_8793139ab3849466)
	__obf_eb7ad60b95be299e := &x509.Certificate{
		SerialNumber: __obf_f2edc2d35a45d523(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_e0ad65cc4a7b9d78},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_66968245163fa68a,
		NotAfter:              __obf_bfe877faf90c2d1d,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_548f732c582bb879 := __obf_cfda8bfc21b81e73()

	__obf_f3ff9550e4d47e55("root", __obf_eb7ad60b95be299e, __obf_548f732c582bb879, __obf_eb7ad60b95be299e, nil)

}

func GenServerCA() {

	__obf_66968245163fa68a, __obf_bfe877faf90c2d1d := __obf_397b631981cf429a(*__obf_3498d038b99e93bc, *__obf_9a81352c2d8f646d)
	__obf_0c3f3422ff4ddc90 := &x509.Certificate{
		SerialNumber: __obf_f2edc2d35a45d523(),
		Subject: pkix.Name{
			Organization: []string{*__obf_a0b8eda05f033c7e},
		},
		NotBefore: __obf_66968245163fa68a,
		NotAfter:  __obf_bfe877faf90c2d1d,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_2ca22ef2f9c8e4f5 := strings.Split(*__obf_5a5eb3590804c3ce, ",")

	for _, __obf_3551ec09edcea1aa := range __obf_2ca22ef2f9c8e4f5 {
		if __obf_5c0de70f8ac4bb34 := net.ParseIP(__obf_3551ec09edcea1aa); __obf_5c0de70f8ac4bb34 != nil {
			__obf_0c3f3422ff4ddc90.IPAddresses = append(__obf_0c3f3422ff4ddc90.IPAddresses, __obf_5c0de70f8ac4bb34)
		} else {
			__obf_0c3f3422ff4ddc90.DNSNames = append(__obf_0c3f3422ff4ddc90.DNSNames, __obf_3551ec09edcea1aa)
		}
	}
	__obf_4c186a057a362197 := __obf_cfda8bfc21b81e73()

	var __obf_eb7ad60b95be299e *x509.Certificate
	var __obf_b7fd5848550ca60c error
	var __obf_548f732c582bb879 any

	__obf_eb7ad60b95be299e, __obf_b7fd5848550ca60c = __obf_85687d92b145c0ba()
	if __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}

	__obf_548f732c582bb879, __obf_b7fd5848550ca60c = __obf_242aaffa404c17b0()
	if __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}

	__obf_f3ff9550e4d47e55("server", __obf_0c3f3422ff4ddc90, __obf_4c186a057a362197, __obf_eb7ad60b95be299e, __obf_548f732c582bb879)
}

func GenClientCA() {

	__obf_66968245163fa68a, __obf_bfe877faf90c2d1d := __obf_397b631981cf429a(*__obf_c2b303d5acd15327, *__obf_e2a041976aba5300)
	__obf_7bf9b9def77e79a4 := &x509.Certificate{
		SerialNumber: __obf_f2edc2d35a45d523(),
		Subject: pkix.Name{
			Organization: []string{*__obf_d3a01009d8f993f2},
		},
		NotBefore: __obf_66968245163fa68a,
		NotAfter:  __obf_bfe877faf90c2d1d,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_7e4eec52bf5a4d32 := __obf_cfda8bfc21b81e73()

	var __obf_eb7ad60b95be299e *x509.Certificate
	var __obf_b7fd5848550ca60c error
	var __obf_548f732c582bb879 any

	__obf_eb7ad60b95be299e, __obf_b7fd5848550ca60c = __obf_85687d92b145c0ba()
	if __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}

	__obf_548f732c582bb879, __obf_b7fd5848550ca60c = __obf_242aaffa404c17b0()
	if __obf_b7fd5848550ca60c != nil {
		log.Fatalln(__obf_b7fd5848550ca60c)
	}

	__obf_f3ff9550e4d47e55("client", __obf_7bf9b9def77e79a4, __obf_7e4eec52bf5a4d32, __obf_eb7ad60b95be299e, __obf_548f732c582bb879)

}
