package __obf_3d5f73927e7527a3

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

func __obf_b79f1e02a5a8414b(__obf_01dcd9060f5d89a1 string, __obf_3cedcfbdaa2a85ae int) (__obf_8deda0e44aa0ab97 time.Time, __obf_b1fbc0d221044e35 time.Time) {
	var __obf_2eef99cccb5deacc error
	if __obf_01dcd9060f5d89a1 == "" {
		__obf_8deda0e44aa0ab97 = time.Now()
	} else {
		if __obf_8deda0e44aa0ab97, __obf_2eef99cccb5deacc = time.Parse("2006-01-02", __obf_01dcd9060f5d89a1); __obf_2eef99cccb5deacc != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_b1fbc0d221044e35 = __obf_8deda0e44aa0ab97.Add(time.Duration(__obf_3cedcfbdaa2a85ae*24) * time.Hour)
		}
	}
	return
}

func __obf_6f8bbcaef8b40ed6() any {
	var __obf_2eef99cccb5deacc error
	var __obf_b78fb0ff4f30c05e any
	switch *__obf_ca39108216f748ac {
	case "":
		if *__obf_3ac46f69661ae0d8 {
			_, __obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = rsa.GenerateKey(rand.Reader, *__obf_d20d08b4ac781958)
		}
	case "P224":
		__obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_b78fb0ff4f30c05e, __obf_2eef99cccb5deacc = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_ca39108216f748ac)
	}

	if __obf_2eef99cccb5deacc != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_2eef99cccb5deacc)
	}

	return __obf_b78fb0ff4f30c05e
}

func __obf_5bcb9d31eff6befb(__obf_76df4a3560fb1a28, __obf_979c0dc6d04409d2 string, __obf_de9d52eae290157f []byte, __obf_d1c43c805de05922 os.FileMode) error {

	if _, __obf_2eef99cccb5deacc := os.Stat(__obf_76df4a3560fb1a28); __obf_2eef99cccb5deacc == nil {
		if __obf_2eef99cccb5deacc = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_76df4a3560fb1a28, string(os.PathSeparator), __obf_979c0dc6d04409d2), __obf_de9d52eae290157f, __obf_d1c43c805de05922); __obf_2eef99cccb5deacc != nil {
			return errors.Wrap(__obf_2eef99cccb5deacc, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_2eef99cccb5deacc) {
		return errors.Wrap(__obf_2eef99cccb5deacc, "Directory not exist.")
	} else {
		return __obf_2eef99cccb5deacc
	}
	return nil
}

func __obf_952b9637f1aed208(__obf_68a329dd0bc1d1a3 string, __obf_007dd92ffb4e580f *x509.Certificate, __obf_d22a56f6883a4549 any, __obf_8ed23350cdb49c87 *x509.Certificate, __obf_0e8e0972e62f0547 any) {

	var __obf_5ce2e6a786e22558 any
	switch __obf_1fc958d8b2470f45 := __obf_d22a56f6883a4549.(type) {
	case *rsa.PrivateKey:
		__obf_5ce2e6a786e22558 = &__obf_1fc958d8b2470f45.PublicKey
	case *ecdsa.PrivateKey:
		__obf_5ce2e6a786e22558 = &__obf_1fc958d8b2470f45.PublicKey
	case ed25519.PrivateKey:
		__obf_5ce2e6a786e22558 = __obf_1fc958d8b2470f45.Public().(ed25519.PublicKey)
	default:
		__obf_5ce2e6a786e22558 = nil
	}

	var __obf_bdf472f121855193 any
	if __obf_0e8e0972e62f0547 != nil {
		__obf_bdf472f121855193 = __obf_0e8e0972e62f0547
	} else {
		__obf_bdf472f121855193 = __obf_d22a56f6883a4549
	}
	__obf_6fa9fdcf13d2e16d, __obf_2eef99cccb5deacc := x509.CreateCertificate(rand.Reader, __obf_007dd92ffb4e580f, __obf_8ed23350cdb49c87, __obf_5ce2e6a786e22558, __obf_bdf472f121855193)
	if __obf_2eef99cccb5deacc != nil {
		log.Println("create failed", __obf_2eef99cccb5deacc)
		return
	}

	__obf_4c473f20858b8d85 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_6fa9fdcf13d2e16d})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_2eef99cccb5deacc = __obf_5bcb9d31eff6befb(*__obf_99d81957a2a2485e, __obf_68a329dd0bc1d1a3+".pem", __obf_4c473f20858b8d85, 0777); __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}

	__obf_feb7f769524150f3, __obf_2eef99cccb5deacc := x509.MarshalPKCS8PrivateKey(__obf_d22a56f6883a4549)
	if __obf_2eef99cccb5deacc != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_7ad817bc29b678fa := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_feb7f769524150f3})

	if __obf_2eef99cccb5deacc = __obf_5bcb9d31eff6befb(*__obf_99d81957a2a2485e, __obf_68a329dd0bc1d1a3+".key", __obf_7ad817bc29b678fa, 0777); __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}
}

// 公钥证书解析
func __obf_0b3418c4bcccbfff() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_d964a41906ae924f, __obf_2eef99cccb5deacc := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_99d81957a2a2485e, string(os.PathSeparator))); __obf_2eef99cccb5deacc != nil {
		return nil, errors.Wrap(__obf_2eef99cccb5deacc, "root.pem: Failed to read file.")
	} else {
		__obf_dfa91ab2d4193315, __obf_c7ba3962e3bf470d := pem.Decode(__obf_d964a41906ae924f)
		if __obf_dfa91ab2d4193315 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_c7ba3962e3bf470d)
		}

		return x509.ParseCertificate(__obf_dfa91ab2d4193315.Bytes)
	}
}

// 私钥解析
func __obf_af554c954c66268c() (__obf_d22a56f6883a4549 any, __obf_2eef99cccb5deacc error) {
	if __obf_d964a41906ae924f, __obf_2eef99cccb5deacc := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_99d81957a2a2485e, string(os.PathSeparator))); __obf_2eef99cccb5deacc != nil {
		return nil, errors.Wrap(__obf_2eef99cccb5deacc, "root.key: Failed to read file.")
	} else {
		__obf_dfa91ab2d4193315, __obf_c7ba3962e3bf470d := pem.Decode(__obf_d964a41906ae924f)
		if __obf_dfa91ab2d4193315 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_c7ba3962e3bf470d)
		}

		return x509.ParsePKCS8PrivateKey(__obf_dfa91ab2d4193315.Bytes)
	}
}

func __obf_e80c0343acbd5c62() *big.Int {
	__obf_dd75ed26651edcb5 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_b585f770988471a4, __obf_2eef99cccb5deacc := rand.Int(rand.Reader, __obf_dd75ed26651edcb5); __obf_2eef99cccb5deacc != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_b585f770988471a4
	}
}

func GenRootCA() {

	__obf_7ebd6962ed75bdd3, __obf_7dbb9a30768110e2 := __obf_b79f1e02a5a8414b(*__obf_5670263c37016623, *__obf_3dfc5003a42402a4)
	__obf_908b47d36bbdd2ed := &x509.Certificate{
		SerialNumber: __obf_e80c0343acbd5c62(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_b5bc963dd5e57e51},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_7ebd6962ed75bdd3,
		NotAfter:              __obf_7dbb9a30768110e2,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_2cd6bb49ee83c5f8 := __obf_6f8bbcaef8b40ed6()

	__obf_952b9637f1aed208("root", __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8, __obf_908b47d36bbdd2ed, nil)

}

func GenServerCA() {

	__obf_7ebd6962ed75bdd3, __obf_7dbb9a30768110e2 := __obf_b79f1e02a5a8414b(*__obf_a10e2c032b2a681d, *__obf_697861048accd993)
	__obf_0456a1ca2826fea1 := &x509.Certificate{
		SerialNumber: __obf_e80c0343acbd5c62(),
		Subject: pkix.Name{
			Organization: []string{*__obf_d03c4ee6c7e51cb5},
		},
		NotBefore: __obf_7ebd6962ed75bdd3,
		NotAfter:  __obf_7dbb9a30768110e2,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_1c56ae222a3c2766 := strings.Split(*__obf_1a8588b7039b1147, ",")

	for _, __obf_81d55c9a99174773 := range __obf_1c56ae222a3c2766 {
		if __obf_b98bdd9cf219874b := net.ParseIP(__obf_81d55c9a99174773); __obf_b98bdd9cf219874b != nil {
			__obf_0456a1ca2826fea1.IPAddresses = append(__obf_0456a1ca2826fea1.IPAddresses, __obf_b98bdd9cf219874b)
		} else {
			__obf_0456a1ca2826fea1.DNSNames = append(__obf_0456a1ca2826fea1.DNSNames, __obf_81d55c9a99174773)
		}
	}
	__obf_8f52ed7b17bc6019 := __obf_6f8bbcaef8b40ed6()

	var __obf_908b47d36bbdd2ed *x509.Certificate
	var __obf_2eef99cccb5deacc error
	var __obf_2cd6bb49ee83c5f8 any

	__obf_908b47d36bbdd2ed, __obf_2eef99cccb5deacc = __obf_0b3418c4bcccbfff()
	if __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}

	__obf_2cd6bb49ee83c5f8, __obf_2eef99cccb5deacc = __obf_af554c954c66268c()
	if __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}

	__obf_952b9637f1aed208("server", __obf_0456a1ca2826fea1, __obf_8f52ed7b17bc6019, __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8)
}

func GenClientCA() {

	__obf_7ebd6962ed75bdd3, __obf_7dbb9a30768110e2 := __obf_b79f1e02a5a8414b(*__obf_a7d625f3ccb24dd6, *__obf_7350e0f0927a0cf7)
	__obf_0f9b320858d91156 := &x509.Certificate{
		SerialNumber: __obf_e80c0343acbd5c62(),
		Subject: pkix.Name{
			Organization: []string{*__obf_5aaf18bf1be127ce},
		},
		NotBefore: __obf_7ebd6962ed75bdd3,
		NotAfter:  __obf_7dbb9a30768110e2,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_6755722962ef4f64 := __obf_6f8bbcaef8b40ed6()

	var __obf_908b47d36bbdd2ed *x509.Certificate
	var __obf_2eef99cccb5deacc error
	var __obf_2cd6bb49ee83c5f8 any

	__obf_908b47d36bbdd2ed, __obf_2eef99cccb5deacc = __obf_0b3418c4bcccbfff()
	if __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}

	__obf_2cd6bb49ee83c5f8, __obf_2eef99cccb5deacc = __obf_af554c954c66268c()
	if __obf_2eef99cccb5deacc != nil {
		log.Fatalln(__obf_2eef99cccb5deacc)
	}

	__obf_952b9637f1aed208("client", __obf_0f9b320858d91156, __obf_6755722962ef4f64, __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8)

}
