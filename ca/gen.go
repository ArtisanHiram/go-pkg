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

func __obf_834688e320d0c9af(__obf_d85d527966edead0 string, __obf_3308a34cf09359f2 int) (__obf_224d4644f28aa7d9 time.Time, __obf_cf0fcb3dc0fa6e8d time.Time) {
	var __obf_50d9d8acee20f6af error
	if __obf_d85d527966edead0 == "" {
		__obf_224d4644f28aa7d9 = time.Now()
	} else {
		if __obf_224d4644f28aa7d9, __obf_50d9d8acee20f6af = time.Parse("2006-01-02", __obf_d85d527966edead0); __obf_50d9d8acee20f6af != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_cf0fcb3dc0fa6e8d = __obf_224d4644f28aa7d9.Add(time.Duration(__obf_3308a34cf09359f2*24) * time.Hour)
		}
	}
	return
}

func __obf_e5d0be4e87dfa588() any {
	var __obf_50d9d8acee20f6af error
	var __obf_a00b20fba5e0687b any
	switch *__obf_3d01f66db0c6933a {
	case "":
		if *__obf_762a091af78eacdc {
			_, __obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = rsa.GenerateKey(rand.Reader, *__obf_0cb5ccb4d6fefda5)
		}
	case "P224":
		__obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_a00b20fba5e0687b, __obf_50d9d8acee20f6af = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_3d01f66db0c6933a)
	}

	if __obf_50d9d8acee20f6af != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_50d9d8acee20f6af)
	}

	return __obf_a00b20fba5e0687b
}

func __obf_78003f0e2407b45a(__obf_3c55cf4edba4b8ae, __obf_27a65f55daceb84c string, __obf_c604e6391eda24f3 []byte, __obf_a88f5c0ea593aad0 os.FileMode) error {

	if _, __obf_50d9d8acee20f6af := os.Stat(__obf_3c55cf4edba4b8ae); __obf_50d9d8acee20f6af == nil {
		if __obf_50d9d8acee20f6af = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_3c55cf4edba4b8ae, string(os.PathSeparator), __obf_27a65f55daceb84c), __obf_c604e6391eda24f3, __obf_a88f5c0ea593aad0); __obf_50d9d8acee20f6af != nil {
			return errors.Wrap(__obf_50d9d8acee20f6af, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_50d9d8acee20f6af) {
		return errors.Wrap(__obf_50d9d8acee20f6af, "Directory not exist.")
	} else {
		return __obf_50d9d8acee20f6af
	}
	return nil
}

func __obf_62822992eba5196a(__obf_20afa3a20b4e7df3 string, __obf_b16eba1636bcfd7b *x509.Certificate, __obf_79f143f45af76681 any, __obf_9ae66557d0a6577d *x509.Certificate, __obf_89bc8a7c7f623f96 any) {

	var __obf_0673cdb4bf5415c0 any
	switch __obf_3f44e40d8e468732 := __obf_79f143f45af76681.(type) {
	case *rsa.PrivateKey:
		__obf_0673cdb4bf5415c0 = &__obf_3f44e40d8e468732.PublicKey
	case *ecdsa.PrivateKey:
		__obf_0673cdb4bf5415c0 = &__obf_3f44e40d8e468732.PublicKey
	case ed25519.PrivateKey:
		__obf_0673cdb4bf5415c0 = __obf_3f44e40d8e468732.Public().(ed25519.PublicKey)
	default:
		__obf_0673cdb4bf5415c0 = nil
	}

	var __obf_78a4f3baeaa6935c any
	if __obf_89bc8a7c7f623f96 != nil {
		__obf_78a4f3baeaa6935c = __obf_89bc8a7c7f623f96
	} else {
		__obf_78a4f3baeaa6935c = __obf_79f143f45af76681
	}
	__obf_b0b275f1493f4e8d, __obf_50d9d8acee20f6af := x509.CreateCertificate(rand.Reader, __obf_b16eba1636bcfd7b, __obf_9ae66557d0a6577d, __obf_0673cdb4bf5415c0, __obf_78a4f3baeaa6935c)
	if __obf_50d9d8acee20f6af != nil {
		log.Println("create failed", __obf_50d9d8acee20f6af)
		return
	}
	__obf_adda7084495ff63a := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_b0b275f1493f4e8d})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_50d9d8acee20f6af = __obf_78003f0e2407b45a(*__obf_42e3c6717d306275, __obf_20afa3a20b4e7df3+".pem", __obf_adda7084495ff63a, 0777); __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
	__obf_0833ebb9145cf77e, __obf_50d9d8acee20f6af := x509.MarshalPKCS8PrivateKey(__obf_79f143f45af76681)
	if __obf_50d9d8acee20f6af != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_b584dab635b4689d := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_0833ebb9145cf77e})

	if __obf_50d9d8acee20f6af = __obf_78003f0e2407b45a(*__obf_42e3c6717d306275, __obf_20afa3a20b4e7df3+".key", __obf_b584dab635b4689d, 0777); __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
}

// 公钥证书解析
func __obf_8198211ff32e7b07() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_de17539f1cd51ed9, __obf_50d9d8acee20f6af := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_42e3c6717d306275, string(os.PathSeparator))); __obf_50d9d8acee20f6af != nil {
		return nil, errors.Wrap(__obf_50d9d8acee20f6af, "root.pem: Failed to read file.")
	} else {
		__obf_83677c32fae068d6, __obf_034efd1d17349b25 := pem.Decode(__obf_de17539f1cd51ed9)
		if __obf_83677c32fae068d6 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_034efd1d17349b25)
		}

		return x509.ParseCertificate(__obf_83677c32fae068d6.Bytes)
	}
}

// 私钥解析
func __obf_b0e5e6585d081719() (__obf_79f143f45af76681 any, __obf_50d9d8acee20f6af error) {
	if __obf_de17539f1cd51ed9, __obf_50d9d8acee20f6af := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_42e3c6717d306275, string(os.PathSeparator))); __obf_50d9d8acee20f6af != nil {
		return nil, errors.Wrap(__obf_50d9d8acee20f6af, "root.key: Failed to read file.")
	} else {
		__obf_83677c32fae068d6, __obf_034efd1d17349b25 := pem.Decode(__obf_de17539f1cd51ed9)
		if __obf_83677c32fae068d6 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_034efd1d17349b25)
		}

		return x509.ParsePKCS8PrivateKey(__obf_83677c32fae068d6.Bytes)
	}
}

func __obf_0349b2634f2531e9() *big.Int {
	__obf_4bd2feb38609c31d := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_5be567aaac1af9e6, __obf_50d9d8acee20f6af := rand.Int(rand.Reader, __obf_4bd2feb38609c31d); __obf_50d9d8acee20f6af != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_5be567aaac1af9e6
	}
}

func GenRootCA() {
	__obf_ad62fa64a348cbb3, __obf_ed7b277281995073 := __obf_834688e320d0c9af(*__obf_77b98a74ef4e3ee6, *__obf_7342674c21b8d9fe)
	__obf_152d3d6821401fe0 := &x509.Certificate{
		SerialNumber: __obf_0349b2634f2531e9(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_76e50b8d8497183b},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_ad62fa64a348cbb3,
		NotAfter:              __obf_ed7b277281995073,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_b35fd3d30bda193e := __obf_e5d0be4e87dfa588()
	__obf_62822992eba5196a("root", __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e, __obf_152d3d6821401fe0, nil)

}

func GenServerCA() {
	__obf_ad62fa64a348cbb3, __obf_ed7b277281995073 := __obf_834688e320d0c9af(*__obf_0ce26a566e9dee32, *__obf_f2a23f9f3af65936)
	__obf_b04ef76e341283b9 := &x509.Certificate{
		SerialNumber: __obf_0349b2634f2531e9(),
		Subject: pkix.Name{
			Organization: []string{*__obf_ca34a921c84f35c2},
		},
		NotBefore: __obf_ad62fa64a348cbb3,
		NotAfter:  __obf_ed7b277281995073,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_447f061b6b9d3158 := strings.Split(*__obf_1d133200ede8a085, ",")

	for _, __obf_7e404a3606ede3f7 := range __obf_447f061b6b9d3158 {
		if __obf_fee41b11291ca394 := net.ParseIP(__obf_7e404a3606ede3f7); __obf_fee41b11291ca394 != nil {
			__obf_b04ef76e341283b9.
				IPAddresses = append(__obf_b04ef76e341283b9.IPAddresses, __obf_fee41b11291ca394)
		} else {
			__obf_b04ef76e341283b9.
				DNSNames = append(__obf_b04ef76e341283b9.DNSNames, __obf_7e404a3606ede3f7)
		}
	}
	__obf_288f66eab17b85e3 := __obf_e5d0be4e87dfa588()

	var __obf_152d3d6821401fe0 *x509.Certificate
	var __obf_50d9d8acee20f6af error
	var __obf_b35fd3d30bda193e any
	__obf_152d3d6821401fe0, __obf_50d9d8acee20f6af = __obf_8198211ff32e7b07()
	if __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
	__obf_b35fd3d30bda193e, __obf_50d9d8acee20f6af = __obf_b0e5e6585d081719()
	if __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
	__obf_62822992eba5196a("server", __obf_b04ef76e341283b9, __obf_288f66eab17b85e3, __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e)
}

func GenClientCA() {
	__obf_ad62fa64a348cbb3, __obf_ed7b277281995073 := __obf_834688e320d0c9af(*__obf_7ce6c606cf2b1391, *__obf_626afac41336a16f)
	__obf_50f1524a0f572b1a := &x509.Certificate{
		SerialNumber: __obf_0349b2634f2531e9(),
		Subject: pkix.Name{
			Organization: []string{*__obf_e89e84dedb3f1ddb},
		},
		NotBefore: __obf_ad62fa64a348cbb3,
		NotAfter:  __obf_ed7b277281995073,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_1d049cf4e1477ec8 := __obf_e5d0be4e87dfa588()

	var __obf_152d3d6821401fe0 *x509.Certificate
	var __obf_50d9d8acee20f6af error
	var __obf_b35fd3d30bda193e any
	__obf_152d3d6821401fe0, __obf_50d9d8acee20f6af = __obf_8198211ff32e7b07()
	if __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
	__obf_b35fd3d30bda193e, __obf_50d9d8acee20f6af = __obf_b0e5e6585d081719()
	if __obf_50d9d8acee20f6af != nil {
		log.Fatalln(__obf_50d9d8acee20f6af)
	}
	__obf_62822992eba5196a("client", __obf_50f1524a0f572b1a, __obf_1d049cf4e1477ec8, __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e)

}
