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

func __obf_57a9488a4123017e(__obf_9b7103436c78dbeb string, __obf_848b64619f7af332 int) (__obf_25585a6eaabc8f6c time.Time, __obf_7ee44325403218cd time.Time) {
	var __obf_55b9dfdec0b06a66 error
	if __obf_9b7103436c78dbeb == "" {
		__obf_25585a6eaabc8f6c = time.Now()
	} else {
		if __obf_25585a6eaabc8f6c, __obf_55b9dfdec0b06a66 = time.Parse("2006-01-02", __obf_9b7103436c78dbeb); __obf_55b9dfdec0b06a66 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_7ee44325403218cd = __obf_25585a6eaabc8f6c.Add(time.Duration(__obf_848b64619f7af332*24) * time.Hour)
		}
	}
	return
}

func __obf_c6b27f7bb6658d30() any {
	var __obf_55b9dfdec0b06a66 error
	var __obf_03eb2d2742e0f341 any
	switch *__obf_df6ed5643a391a04 {
	case "":
		if *__obf_7202109d1688abad {
			_, __obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = rsa.GenerateKey(rand.Reader, *__obf_9888c4a4b3aa3e56)
		}
	case "P224":
		__obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_03eb2d2742e0f341, __obf_55b9dfdec0b06a66 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_df6ed5643a391a04)
	}

	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_55b9dfdec0b06a66)
	}

	return __obf_03eb2d2742e0f341
}

func __obf_e9684a45242f62fb(__obf_dd2cd096056c0f1d, __obf_833e66cc24c12354 string, __obf_ba7199f872e9dadf []byte, __obf_ff0a9e378b7231a4 os.FileMode) error {

	if _, __obf_55b9dfdec0b06a66 := os.Stat(__obf_dd2cd096056c0f1d); __obf_55b9dfdec0b06a66 == nil {
		if __obf_55b9dfdec0b06a66 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_dd2cd096056c0f1d, string(os.PathSeparator), __obf_833e66cc24c12354), __obf_ba7199f872e9dadf, __obf_ff0a9e378b7231a4); __obf_55b9dfdec0b06a66 != nil {
			return errors.Wrap(__obf_55b9dfdec0b06a66, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_55b9dfdec0b06a66) {
		return errors.Wrap(__obf_55b9dfdec0b06a66, "Directory not exist.")
	} else {
		return __obf_55b9dfdec0b06a66
	}
	return nil
}

func __obf_ba50149987ebded8(__obf_88d290a111bdfb00 string, __obf_e4997f449948c743 *x509.Certificate, __obf_02b030c8d20c95c6 any, __obf_c5c0c1e601208157 *x509.Certificate, __obf_bbbb35f28334fa11 any) {

	var __obf_6a67a0eb73d11448 any
	switch __obf_d3f8f1027d8d1520 := __obf_02b030c8d20c95c6.(type) {
	case *rsa.PrivateKey:
		__obf_6a67a0eb73d11448 = &__obf_d3f8f1027d8d1520.PublicKey
	case *ecdsa.PrivateKey:
		__obf_6a67a0eb73d11448 = &__obf_d3f8f1027d8d1520.PublicKey
	case ed25519.PrivateKey:
		__obf_6a67a0eb73d11448 = __obf_d3f8f1027d8d1520.Public().(ed25519.PublicKey)
	default:
		__obf_6a67a0eb73d11448 = nil
	}

	var __obf_a663dcf260939188 any
	if __obf_bbbb35f28334fa11 != nil {
		__obf_a663dcf260939188 = __obf_bbbb35f28334fa11
	} else {
		__obf_a663dcf260939188 = __obf_02b030c8d20c95c6
	}
	__obf_4257599bfa9403eb, __obf_55b9dfdec0b06a66 := x509.CreateCertificate(rand.Reader, __obf_e4997f449948c743, __obf_c5c0c1e601208157, __obf_6a67a0eb73d11448, __obf_a663dcf260939188)
	if __obf_55b9dfdec0b06a66 != nil {
		log.Println("create failed", __obf_55b9dfdec0b06a66)
		return
	}
	__obf_df104a4558dc2ab3 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_4257599bfa9403eb})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_55b9dfdec0b06a66 = __obf_e9684a45242f62fb(*__obf_4b00946562018abc, __obf_88d290a111bdfb00+".pem", __obf_df104a4558dc2ab3, 0777); __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
	__obf_fb283a9f3eb2f789, __obf_55b9dfdec0b06a66 := x509.MarshalPKCS8PrivateKey(__obf_02b030c8d20c95c6)
	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_3414987a53346d72 := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_fb283a9f3eb2f789})

	if __obf_55b9dfdec0b06a66 = __obf_e9684a45242f62fb(*__obf_4b00946562018abc, __obf_88d290a111bdfb00+".key", __obf_3414987a53346d72, 0777); __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
}

// 公钥证书解析
func __obf_8b5ef384c128a54f() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_dc20c8bb2ec94dfc, __obf_55b9dfdec0b06a66 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_4b00946562018abc, string(os.PathSeparator))); __obf_55b9dfdec0b06a66 != nil {
		return nil, errors.Wrap(__obf_55b9dfdec0b06a66, "root.pem: Failed to read file.")
	} else {
		__obf_39cfa377b1b672ba, __obf_a607b0b1c6fe966e := pem.Decode(__obf_dc20c8bb2ec94dfc)
		if __obf_39cfa377b1b672ba == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_a607b0b1c6fe966e)
		}

		return x509.ParseCertificate(__obf_39cfa377b1b672ba.Bytes)
	}
}

// 私钥解析
func __obf_a43cddbc7b58ffce() (__obf_02b030c8d20c95c6 any, __obf_55b9dfdec0b06a66 error) {
	if __obf_dc20c8bb2ec94dfc, __obf_55b9dfdec0b06a66 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_4b00946562018abc, string(os.PathSeparator))); __obf_55b9dfdec0b06a66 != nil {
		return nil, errors.Wrap(__obf_55b9dfdec0b06a66, "root.key: Failed to read file.")
	} else {
		__obf_39cfa377b1b672ba, __obf_a607b0b1c6fe966e := pem.Decode(__obf_dc20c8bb2ec94dfc)
		if __obf_39cfa377b1b672ba == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_a607b0b1c6fe966e)
		}

		return x509.ParsePKCS8PrivateKey(__obf_39cfa377b1b672ba.Bytes)
	}
}

func __obf_a167439616cf2649() *big.Int {
	__obf_a354bd337aeff148 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_3923c321bcc62605, __obf_55b9dfdec0b06a66 := rand.Int(rand.Reader, __obf_a354bd337aeff148); __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_3923c321bcc62605
	}
}

func GenRootCA() {
	__obf_2b09718d038e93d3, __obf_124de5243edf747c := __obf_57a9488a4123017e(*__obf_6abe834fb3e25469, *__obf_8e04f1efba12895c)
	__obf_b95177042f904a4c := &x509.Certificate{
		SerialNumber: __obf_a167439616cf2649(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_b1f0196313c2663e},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_2b09718d038e93d3,
		NotAfter:              __obf_124de5243edf747c,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_e33fe71340c03c86 := __obf_c6b27f7bb6658d30()
	__obf_ba50149987ebded8("root", __obf_b95177042f904a4c, __obf_e33fe71340c03c86, __obf_b95177042f904a4c, nil)

}

func GenServerCA() {
	__obf_2b09718d038e93d3, __obf_124de5243edf747c := __obf_57a9488a4123017e(*__obf_eafe7ab8a89428b8, *__obf_e29bec54c690b2f9)
	__obf_1c097168084f7804 := &x509.Certificate{
		SerialNumber: __obf_a167439616cf2649(),
		Subject: pkix.Name{
			Organization: []string{*__obf_46d118ebb28c1025},
		},
		NotBefore: __obf_2b09718d038e93d3,
		NotAfter:  __obf_124de5243edf747c,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_fc6e2c0ac1d596f1 := strings.Split(*__obf_7927c321cb778076, ",")

	for _, __obf_2a5b387db77f9ace := range __obf_fc6e2c0ac1d596f1 {
		if __obf_fed4296046358797 := net.ParseIP(__obf_2a5b387db77f9ace); __obf_fed4296046358797 != nil {
			__obf_1c097168084f7804.
				IPAddresses = append(__obf_1c097168084f7804.IPAddresses, __obf_fed4296046358797)
		} else {
			__obf_1c097168084f7804.
				DNSNames = append(__obf_1c097168084f7804.DNSNames, __obf_2a5b387db77f9ace)
		}
	}
	__obf_ce3eb260a95f4fc4 := __obf_c6b27f7bb6658d30()

	var __obf_b95177042f904a4c *x509.Certificate
	var __obf_55b9dfdec0b06a66 error
	var __obf_e33fe71340c03c86 any
	__obf_b95177042f904a4c, __obf_55b9dfdec0b06a66 = __obf_8b5ef384c128a54f()
	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
	__obf_e33fe71340c03c86, __obf_55b9dfdec0b06a66 = __obf_a43cddbc7b58ffce()
	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
	__obf_ba50149987ebded8("server", __obf_1c097168084f7804, __obf_ce3eb260a95f4fc4, __obf_b95177042f904a4c, __obf_e33fe71340c03c86)
}

func GenClientCA() {
	__obf_2b09718d038e93d3, __obf_124de5243edf747c := __obf_57a9488a4123017e(*__obf_6315e9ca5d8be1e7, *__obf_892547ae07fe593a)
	__obf_c17d5b3ab639cab3 := &x509.Certificate{
		SerialNumber: __obf_a167439616cf2649(),
		Subject: pkix.Name{
			Organization: []string{*__obf_64c69c9a9ddb71c1},
		},
		NotBefore: __obf_2b09718d038e93d3,
		NotAfter:  __obf_124de5243edf747c,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_d6769bc7d39c18cc := __obf_c6b27f7bb6658d30()

	var __obf_b95177042f904a4c *x509.Certificate
	var __obf_55b9dfdec0b06a66 error
	var __obf_e33fe71340c03c86 any
	__obf_b95177042f904a4c, __obf_55b9dfdec0b06a66 = __obf_8b5ef384c128a54f()
	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
	__obf_e33fe71340c03c86, __obf_55b9dfdec0b06a66 = __obf_a43cddbc7b58ffce()
	if __obf_55b9dfdec0b06a66 != nil {
		log.Fatalln(__obf_55b9dfdec0b06a66)
	}
	__obf_ba50149987ebded8("client", __obf_c17d5b3ab639cab3, __obf_d6769bc7d39c18cc, __obf_b95177042f904a4c, __obf_e33fe71340c03c86)

}
