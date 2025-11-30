package main

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

func __obf_50861aeb6316df9d(__obf_eb1c4d7e77953d1e string, __obf_6f9c83ed61caf01d int) (__obf_79c5cb24f8272c7a time.Time, __obf_e0c4d448c79e0936 time.Time) {
	var __obf_dd7d7edd66f39226 error
	if __obf_eb1c4d7e77953d1e == "" {
		__obf_79c5cb24f8272c7a = time.Now()
	} else {
		if __obf_79c5cb24f8272c7a, __obf_dd7d7edd66f39226 = time.Parse("2006-01-02", __obf_eb1c4d7e77953d1e); __obf_dd7d7edd66f39226 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_e0c4d448c79e0936 = __obf_79c5cb24f8272c7a.Add(time.Duration(__obf_6f9c83ed61caf01d*24) * time.Hour)
		}
	}
	return
}

func __obf_1512141705373040() any {
	var __obf_dd7d7edd66f39226 error
	var __obf_fe5ca72bb97c16b8 any
	switch *__obf_65d9b57be59fef7f {
	case "":
		if *__obf_bda93e55c59138d6 {
			_, __obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = rsa.GenerateKey(rand.Reader, *__obf_676c29a914cffb78)
		}
	case "P224":
		__obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_fe5ca72bb97c16b8, __obf_dd7d7edd66f39226 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_65d9b57be59fef7f)
	}

	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_dd7d7edd66f39226)
	}

	return __obf_fe5ca72bb97c16b8
}

func __obf_275389c5a27bd707(__obf_3a78856844f33f61, __obf_236103f71d83e58c string, __obf_ab8b60b7bfc3fe49 []byte, __obf_961200f064226cba os.FileMode) error {

	if _, __obf_dd7d7edd66f39226 := os.Stat(__obf_3a78856844f33f61); __obf_dd7d7edd66f39226 == nil {
		if __obf_dd7d7edd66f39226 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_3a78856844f33f61, string(os.PathSeparator), __obf_236103f71d83e58c), __obf_ab8b60b7bfc3fe49, __obf_961200f064226cba); __obf_dd7d7edd66f39226 != nil {
			return errors.Wrap(__obf_dd7d7edd66f39226, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_dd7d7edd66f39226) {
		return errors.Wrap(__obf_dd7d7edd66f39226, "Directory not exist.")
	} else {
		return __obf_dd7d7edd66f39226
	}
	return nil
}

func __obf_83fe6121ae7b7075(__obf_09d528286129fda4 string, __obf_1118abe6002652cf *x509.Certificate, __obf_d073cba656c62c18 any, __obf_d77ba6307d0c3956 *x509.Certificate, __obf_aff132c3a457e217 any) {

	var __obf_fb675a630c70f7b6 any
	switch __obf_9ae28662bdf35e7f := __obf_d073cba656c62c18.(type) {
	case *rsa.PrivateKey:
		__obf_fb675a630c70f7b6 = &__obf_9ae28662bdf35e7f.PublicKey
	case *ecdsa.PrivateKey:
		__obf_fb675a630c70f7b6 = &__obf_9ae28662bdf35e7f.PublicKey
	case ed25519.PrivateKey:
		__obf_fb675a630c70f7b6 = __obf_9ae28662bdf35e7f.Public().(ed25519.PublicKey)
	default:
		__obf_fb675a630c70f7b6 = nil
	}

	var __obf_291b188c843da91c any
	if __obf_aff132c3a457e217 != nil {
		__obf_291b188c843da91c = __obf_aff132c3a457e217
	} else {
		__obf_291b188c843da91c = __obf_d073cba656c62c18
	}
	__obf_f52e4c0d1cc74eae, __obf_dd7d7edd66f39226 := x509.CreateCertificate(rand.Reader, __obf_1118abe6002652cf, __obf_d77ba6307d0c3956, __obf_fb675a630c70f7b6, __obf_291b188c843da91c)
	if __obf_dd7d7edd66f39226 != nil {
		log.Println("create failed", __obf_dd7d7edd66f39226)
		return
	}
	__obf_296c04406cfe0986 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_f52e4c0d1cc74eae})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_dd7d7edd66f39226 = __obf_275389c5a27bd707(*__obf_a210ca1feaeefbbf, __obf_09d528286129fda4+".pem", __obf_296c04406cfe0986, 0777); __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
	__obf_df3032d429e4c468, __obf_dd7d7edd66f39226 := x509.MarshalPKCS8PrivateKey(__obf_d073cba656c62c18)
	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_a47b9e6b4f4d8ea1 := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_df3032d429e4c468})

	if __obf_dd7d7edd66f39226 = __obf_275389c5a27bd707(*__obf_a210ca1feaeefbbf, __obf_09d528286129fda4+".key", __obf_a47b9e6b4f4d8ea1, 0777); __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
}

// 公钥证书解析
func __obf_350002914b8bd6af() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_a9a4adddfbfa1141, __obf_dd7d7edd66f39226 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_a210ca1feaeefbbf, string(os.PathSeparator))); __obf_dd7d7edd66f39226 != nil {
		return nil, errors.Wrap(__obf_dd7d7edd66f39226, "root.pem: Failed to read file.")
	} else {
		__obf_eae9bc071f3f7309, __obf_cafccb3acc1ca000 := pem.Decode(__obf_a9a4adddfbfa1141)
		if __obf_eae9bc071f3f7309 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_cafccb3acc1ca000)
		}

		return x509.ParseCertificate(__obf_eae9bc071f3f7309.Bytes)
	}
}

// 私钥解析
func __obf_fb529a535bcb04b4() (__obf_d073cba656c62c18 any, __obf_dd7d7edd66f39226 error) {
	if __obf_a9a4adddfbfa1141, __obf_dd7d7edd66f39226 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_a210ca1feaeefbbf, string(os.PathSeparator))); __obf_dd7d7edd66f39226 != nil {
		return nil, errors.Wrap(__obf_dd7d7edd66f39226, "root.key: Failed to read file.")
	} else {
		__obf_eae9bc071f3f7309, __obf_cafccb3acc1ca000 := pem.Decode(__obf_a9a4adddfbfa1141)
		if __obf_eae9bc071f3f7309 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_cafccb3acc1ca000)
		}

		return x509.ParsePKCS8PrivateKey(__obf_eae9bc071f3f7309.Bytes)
	}
}

func __obf_3f3050c01ce11b88() *big.Int {
	__obf_210a169d92dd72f0 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_7533060142a9b187, __obf_dd7d7edd66f39226 := rand.Int(rand.Reader, __obf_210a169d92dd72f0); __obf_dd7d7edd66f39226 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_7533060142a9b187
	}
}

func GenRootCA() {
	__obf_c7e8d5c9cc22e854, __obf_4994ee994dde0cd9 := __obf_50861aeb6316df9d(*__obf_2010a1d7a6b43947, *__obf_77175f9ff579aa0c)
	__obf_ec85745def2d4ae6 := &x509.Certificate{
		SerialNumber: __obf_3f3050c01ce11b88(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_c9dfe7ad685e91c6},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_c7e8d5c9cc22e854,
		NotAfter:              __obf_4994ee994dde0cd9,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_91759aa608c2ec39 := __obf_1512141705373040()
	__obf_83fe6121ae7b7075("root", __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39, __obf_ec85745def2d4ae6, nil)

}

func GenServerCA() {
	__obf_c7e8d5c9cc22e854, __obf_4994ee994dde0cd9 := __obf_50861aeb6316df9d(*__obf_cfa66aa3eaf60752, *__obf_04ef712d5ae945e6)
	__obf_e64c40cfb6bdad21 := &x509.Certificate{
		SerialNumber: __obf_3f3050c01ce11b88(),
		Subject: pkix.Name{
			Organization: []string{*__obf_d66092b40e18dd97},
		},
		NotBefore: __obf_c7e8d5c9cc22e854,
		NotAfter:  __obf_4994ee994dde0cd9,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_ad24a3c0621435b3 := strings.Split(*__obf_c60afb3a5311a7bd, ",")

	for _, __obf_07c3f9870a15fb0b := range __obf_ad24a3c0621435b3 {
		if __obf_35fe568983671a61 := net.ParseIP(__obf_07c3f9870a15fb0b); __obf_35fe568983671a61 != nil {
			__obf_e64c40cfb6bdad21.
				IPAddresses = append(__obf_e64c40cfb6bdad21.IPAddresses, __obf_35fe568983671a61)
		} else {
			__obf_e64c40cfb6bdad21.
				DNSNames = append(__obf_e64c40cfb6bdad21.DNSNames, __obf_07c3f9870a15fb0b)
		}
	}
	__obf_34e7baa7d7852431 := __obf_1512141705373040()

	var __obf_ec85745def2d4ae6 *x509.Certificate
	var __obf_dd7d7edd66f39226 error
	var __obf_91759aa608c2ec39 any
	__obf_ec85745def2d4ae6, __obf_dd7d7edd66f39226 = __obf_350002914b8bd6af()
	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
	__obf_91759aa608c2ec39, __obf_dd7d7edd66f39226 = __obf_fb529a535bcb04b4()
	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
	__obf_83fe6121ae7b7075("server", __obf_e64c40cfb6bdad21, __obf_34e7baa7d7852431, __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39)
}

func GenClientCA() {
	__obf_c7e8d5c9cc22e854, __obf_4994ee994dde0cd9 := __obf_50861aeb6316df9d(*__obf_3641c7ed101c51be, *__obf_60288d6a3a5662c8)
	__obf_312e3abafc6822e1 := &x509.Certificate{
		SerialNumber: __obf_3f3050c01ce11b88(),
		Subject: pkix.Name{
			Organization: []string{*__obf_bf1a0d93435070a5},
		},
		NotBefore: __obf_c7e8d5c9cc22e854,
		NotAfter:  __obf_4994ee994dde0cd9,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_6f6fbd67ae64bdcd := __obf_1512141705373040()

	var __obf_ec85745def2d4ae6 *x509.Certificate
	var __obf_dd7d7edd66f39226 error
	var __obf_91759aa608c2ec39 any
	__obf_ec85745def2d4ae6, __obf_dd7d7edd66f39226 = __obf_350002914b8bd6af()
	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
	__obf_91759aa608c2ec39, __obf_dd7d7edd66f39226 = __obf_fb529a535bcb04b4()
	if __obf_dd7d7edd66f39226 != nil {
		log.Fatalln(__obf_dd7d7edd66f39226)
	}
	__obf_83fe6121ae7b7075("client", __obf_312e3abafc6822e1, __obf_6f6fbd67ae64bdcd, __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39)

}
