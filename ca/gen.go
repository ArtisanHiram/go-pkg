package __obf_be728ac087d82248

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

func __obf_4db0a02ff633833b(__obf_aeff198ec4fa7677 string, __obf_3fd4e3d990ac0155 int) (__obf_46fd246225d6407b time.Time, __obf_6b997b3326e4e199 time.Time) {
	var __obf_38d4e8c9703c399c error
	if __obf_aeff198ec4fa7677 == "" {
		__obf_46fd246225d6407b = time.Now()
	} else {
		if __obf_46fd246225d6407b, __obf_38d4e8c9703c399c = time.Parse("2006-01-02", __obf_aeff198ec4fa7677); __obf_38d4e8c9703c399c != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_6b997b3326e4e199 = __obf_46fd246225d6407b.Add(time.Duration(__obf_3fd4e3d990ac0155*24) * time.Hour)
		}
	}
	return
}

func __obf_6759ed1384816cc1() any {
	var __obf_38d4e8c9703c399c error
	var __obf_9b7b8b4ad20133b0 any
	switch *__obf_f272f57faa9e35d7 {
	case "":
		if *__obf_da026a6582b22c97 {
			_, __obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = rsa.GenerateKey(rand.Reader, *__obf_2a90108bf71837ff)
		}
	case "P224":
		__obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_9b7b8b4ad20133b0, __obf_38d4e8c9703c399c = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_f272f57faa9e35d7)
	}

	if __obf_38d4e8c9703c399c != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_38d4e8c9703c399c)
	}

	return __obf_9b7b8b4ad20133b0
}

func __obf_36d86d5284f7867f(__obf_13eeff8e7ff39f79, __obf_54abbf28fc2482e8 string, __obf_1cec567b3907546b []byte, __obf_0a4addbc457fff12 os.FileMode) error {

	if _, __obf_38d4e8c9703c399c := os.Stat(__obf_13eeff8e7ff39f79); __obf_38d4e8c9703c399c == nil {
		if __obf_38d4e8c9703c399c = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_13eeff8e7ff39f79, string(os.PathSeparator), __obf_54abbf28fc2482e8), __obf_1cec567b3907546b, __obf_0a4addbc457fff12); __obf_38d4e8c9703c399c != nil {
			return errors.Wrap(__obf_38d4e8c9703c399c, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_38d4e8c9703c399c) {
		return errors.Wrap(__obf_38d4e8c9703c399c, "Directory not exist.")
	} else {
		return __obf_38d4e8c9703c399c
	}
	return nil
}

func __obf_8e8ff93062efde18(__obf_a6be010d3ba4ee10 string, __obf_133e099e3d467a1b *x509.Certificate, __obf_921c11678a889b4e any, __obf_d8789547d12e19b7 *x509.Certificate, __obf_5721445e51f1c16a any) {

	var __obf_163b33d215165a24 any
	switch __obf_68c922de29a178be := __obf_921c11678a889b4e.(type) {
	case *rsa.PrivateKey:
		__obf_163b33d215165a24 = &__obf_68c922de29a178be.PublicKey
	case *ecdsa.PrivateKey:
		__obf_163b33d215165a24 = &__obf_68c922de29a178be.PublicKey
	case ed25519.PrivateKey:
		__obf_163b33d215165a24 = __obf_68c922de29a178be.Public().(ed25519.PublicKey)
	default:
		__obf_163b33d215165a24 = nil
	}

	var __obf_4841eb4e2e790500 any
	if __obf_5721445e51f1c16a != nil {
		__obf_4841eb4e2e790500 = __obf_5721445e51f1c16a
	} else {
		__obf_4841eb4e2e790500 = __obf_921c11678a889b4e
	}
	__obf_682dc1578dc812d8, __obf_38d4e8c9703c399c := x509.CreateCertificate(rand.Reader, __obf_133e099e3d467a1b, __obf_d8789547d12e19b7, __obf_163b33d215165a24, __obf_4841eb4e2e790500)
	if __obf_38d4e8c9703c399c != nil {
		log.Println("create failed", __obf_38d4e8c9703c399c)
		return
	}

	__obf_72bbb56c78c60fee := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_682dc1578dc812d8})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_38d4e8c9703c399c = __obf_36d86d5284f7867f(*__obf_35c2f92426e7679c, __obf_a6be010d3ba4ee10+".pem", __obf_72bbb56c78c60fee, 0777); __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}

	__obf_4cfa010f200444c3, __obf_38d4e8c9703c399c := x509.MarshalPKCS8PrivateKey(__obf_921c11678a889b4e)
	if __obf_38d4e8c9703c399c != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_f075f6e8e46f2c96 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_4cfa010f200444c3})

	if __obf_38d4e8c9703c399c = __obf_36d86d5284f7867f(*__obf_35c2f92426e7679c, __obf_a6be010d3ba4ee10+".key", __obf_f075f6e8e46f2c96, 0777); __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}
}

// 公钥证书解析
func __obf_e0742f2b74948a3d() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_ee77e96f9210d626, __obf_38d4e8c9703c399c := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_35c2f92426e7679c, string(os.PathSeparator))); __obf_38d4e8c9703c399c != nil {
		return nil, errors.Wrap(__obf_38d4e8c9703c399c, "root.pem: Failed to read file.")
	} else {
		__obf_458d37eba2a092b8, __obf_5aa8a03a18c005aa := pem.Decode(__obf_ee77e96f9210d626)
		if __obf_458d37eba2a092b8 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_5aa8a03a18c005aa)
		}

		return x509.ParseCertificate(__obf_458d37eba2a092b8.Bytes)
	}
}

// 私钥解析
func __obf_6757449f11cb78bb() (__obf_921c11678a889b4e any, __obf_38d4e8c9703c399c error) {
	if __obf_ee77e96f9210d626, __obf_38d4e8c9703c399c := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_35c2f92426e7679c, string(os.PathSeparator))); __obf_38d4e8c9703c399c != nil {
		return nil, errors.Wrap(__obf_38d4e8c9703c399c, "root.key: Failed to read file.")
	} else {
		__obf_458d37eba2a092b8, __obf_5aa8a03a18c005aa := pem.Decode(__obf_ee77e96f9210d626)
		if __obf_458d37eba2a092b8 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_5aa8a03a18c005aa)
		}

		return x509.ParsePKCS8PrivateKey(__obf_458d37eba2a092b8.Bytes)
	}
}

func __obf_075c1507601bf544() *big.Int {
	__obf_471cc825482fd54b := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_ab228a0820e4b369, __obf_38d4e8c9703c399c := rand.Int(rand.Reader, __obf_471cc825482fd54b); __obf_38d4e8c9703c399c != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_ab228a0820e4b369
	}
}

func GenRootCA() {

	__obf_49bb6f6e69c96f91, __obf_f1f1f6d8a69097d2 := __obf_4db0a02ff633833b(*__obf_58c8553c2ef7748c, *__obf_c64022213fc6de55)
	__obf_b3faede65159f8cc := &x509.Certificate{
		SerialNumber: __obf_075c1507601bf544(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_534de083cb38fe22},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_49bb6f6e69c96f91,
		NotAfter:              __obf_f1f1f6d8a69097d2,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_799b6d3ad392bc41 := __obf_6759ed1384816cc1()

	__obf_8e8ff93062efde18("root", __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41, __obf_b3faede65159f8cc, nil)

}

func GenServerCA() {

	__obf_49bb6f6e69c96f91, __obf_f1f1f6d8a69097d2 := __obf_4db0a02ff633833b(*__obf_82052cf1888686b0, *__obf_fce9244442953147)
	__obf_ec14c26d96fa147b := &x509.Certificate{
		SerialNumber: __obf_075c1507601bf544(),
		Subject: pkix.Name{
			Organization: []string{*__obf_14b63fdf42420081},
		},
		NotBefore: __obf_49bb6f6e69c96f91,
		NotAfter:  __obf_f1f1f6d8a69097d2,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_f06ad913bede6373 := strings.Split(*__obf_2017486f76444d49, ",")

	for _, __obf_562235314fa4d813 := range __obf_f06ad913bede6373 {
		if __obf_efdf1783571090b7 := net.ParseIP(__obf_562235314fa4d813); __obf_efdf1783571090b7 != nil {
			__obf_ec14c26d96fa147b.IPAddresses = append(__obf_ec14c26d96fa147b.IPAddresses, __obf_efdf1783571090b7)
		} else {
			__obf_ec14c26d96fa147b.DNSNames = append(__obf_ec14c26d96fa147b.DNSNames, __obf_562235314fa4d813)
		}
	}
	__obf_79e7664f49117a3c := __obf_6759ed1384816cc1()

	var __obf_b3faede65159f8cc *x509.Certificate
	var __obf_38d4e8c9703c399c error
	var __obf_799b6d3ad392bc41 any

	__obf_b3faede65159f8cc, __obf_38d4e8c9703c399c = __obf_e0742f2b74948a3d()
	if __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}

	__obf_799b6d3ad392bc41, __obf_38d4e8c9703c399c = __obf_6757449f11cb78bb()
	if __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}

	__obf_8e8ff93062efde18("server", __obf_ec14c26d96fa147b, __obf_79e7664f49117a3c, __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41)
}

func GenClientCA() {

	__obf_49bb6f6e69c96f91, __obf_f1f1f6d8a69097d2 := __obf_4db0a02ff633833b(*__obf_e5be54f6dccae7ba, *__obf_ee52aae87c9ea350)
	__obf_c048349241bd3fba := &x509.Certificate{
		SerialNumber: __obf_075c1507601bf544(),
		Subject: pkix.Name{
			Organization: []string{*__obf_4ac727a060e9a62b},
		},
		NotBefore: __obf_49bb6f6e69c96f91,
		NotAfter:  __obf_f1f1f6d8a69097d2,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_f75e5e99dd49aa4b := __obf_6759ed1384816cc1()

	var __obf_b3faede65159f8cc *x509.Certificate
	var __obf_38d4e8c9703c399c error
	var __obf_799b6d3ad392bc41 any

	__obf_b3faede65159f8cc, __obf_38d4e8c9703c399c = __obf_e0742f2b74948a3d()
	if __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}

	__obf_799b6d3ad392bc41, __obf_38d4e8c9703c399c = __obf_6757449f11cb78bb()
	if __obf_38d4e8c9703c399c != nil {
		log.Fatalln(__obf_38d4e8c9703c399c)
	}

	__obf_8e8ff93062efde18("client", __obf_c048349241bd3fba, __obf_f75e5e99dd49aa4b, __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41)

}
