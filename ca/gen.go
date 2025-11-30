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

func __obf_24158a440d100e75(__obf_812e873e97939aa4 string, __obf_e5ed06a0c82814d9 int) (__obf_168222c0b2cf6577 time.Time, __obf_4fca170c82fe1afc time.Time) {
	var __obf_2107c594e6166770 error
	if __obf_812e873e97939aa4 == "" {
		__obf_168222c0b2cf6577 = time.Now()
	} else {
		if __obf_168222c0b2cf6577, __obf_2107c594e6166770 = time.Parse("2006-01-02", __obf_812e873e97939aa4); __obf_2107c594e6166770 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_4fca170c82fe1afc = __obf_168222c0b2cf6577.Add(time.Duration(__obf_e5ed06a0c82814d9*24) * time.Hour)
		}
	}
	return
}

func __obf_959da6110d2640cd() any {
	var __obf_2107c594e6166770 error
	var __obf_173167651b5d15e3 any
	switch *__obf_340232d3b11c627a {
	case "":
		if *__obf_47b5e014df570fef {
			_, __obf_173167651b5d15e3, __obf_2107c594e6166770 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_173167651b5d15e3, __obf_2107c594e6166770 = rsa.GenerateKey(rand.Reader, *__obf_63740d0c822b9e42)
		}
	case "P224":
		__obf_173167651b5d15e3, __obf_2107c594e6166770 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_173167651b5d15e3, __obf_2107c594e6166770 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_173167651b5d15e3, __obf_2107c594e6166770 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_173167651b5d15e3, __obf_2107c594e6166770 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_340232d3b11c627a)
	}

	if __obf_2107c594e6166770 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_2107c594e6166770)
	}

	return __obf_173167651b5d15e3
}

func __obf_f2f005c6e08ec56f(__obf_d627e268fa6b1017, __obf_0d586dedae595ea7 string, __obf_ebb22adbcecc8d15 []byte, __obf_ee8d07de2ed39029 os.FileMode) error {

	if _, __obf_2107c594e6166770 := os.Stat(__obf_d627e268fa6b1017); __obf_2107c594e6166770 == nil {
		if __obf_2107c594e6166770 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_d627e268fa6b1017, string(os.PathSeparator), __obf_0d586dedae595ea7), __obf_ebb22adbcecc8d15, __obf_ee8d07de2ed39029); __obf_2107c594e6166770 != nil {
			return errors.Wrap(__obf_2107c594e6166770, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_2107c594e6166770) {
		return errors.Wrap(__obf_2107c594e6166770, "Directory not exist.")
	} else {
		return __obf_2107c594e6166770
	}
	return nil
}

func __obf_6b123c9194577d08(__obf_d6849ab3a95e4298 string, __obf_ec1cc1221f3f9dca *x509.Certificate, __obf_c1c829b3ef681f91 any, __obf_db9fe55105775ed2 *x509.Certificate, __obf_f758ce7f9074c0a3 any) {

	var __obf_aca6cb2251255f64 any
	switch __obf_5afb90efe43c8854 := __obf_c1c829b3ef681f91.(type) {
	case *rsa.PrivateKey:
		__obf_aca6cb2251255f64 = &__obf_5afb90efe43c8854.PublicKey
	case *ecdsa.PrivateKey:
		__obf_aca6cb2251255f64 = &__obf_5afb90efe43c8854.PublicKey
	case ed25519.PrivateKey:
		__obf_aca6cb2251255f64 = __obf_5afb90efe43c8854.Public().(ed25519.PublicKey)
	default:
		__obf_aca6cb2251255f64 = nil
	}

	var __obf_deacdf85c4285de9 any
	if __obf_f758ce7f9074c0a3 != nil {
		__obf_deacdf85c4285de9 = __obf_f758ce7f9074c0a3
	} else {
		__obf_deacdf85c4285de9 = __obf_c1c829b3ef681f91
	}
	__obf_18d5d6f844489868, __obf_2107c594e6166770 := x509.CreateCertificate(rand.Reader, __obf_ec1cc1221f3f9dca, __obf_db9fe55105775ed2, __obf_aca6cb2251255f64, __obf_deacdf85c4285de9)
	if __obf_2107c594e6166770 != nil {
		log.Println("create failed", __obf_2107c594e6166770)
		return
	}
	__obf_c3dd8c46e8dba5d5 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_18d5d6f844489868})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_2107c594e6166770 = __obf_f2f005c6e08ec56f(*__obf_b832ed210a876042, __obf_d6849ab3a95e4298+".pem", __obf_c3dd8c46e8dba5d5, 0777); __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
	__obf_8941cf408812dcc7, __obf_2107c594e6166770 := x509.MarshalPKCS8PrivateKey(__obf_c1c829b3ef681f91)
	if __obf_2107c594e6166770 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_214844b19375b2d9 := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_8941cf408812dcc7})

	if __obf_2107c594e6166770 = __obf_f2f005c6e08ec56f(*__obf_b832ed210a876042, __obf_d6849ab3a95e4298+".key", __obf_214844b19375b2d9, 0777); __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
}

// 公钥证书解析
func __obf_ca4f70ed37c87c67() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_304dce20a2343d37, __obf_2107c594e6166770 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_b832ed210a876042, string(os.PathSeparator))); __obf_2107c594e6166770 != nil {
		return nil, errors.Wrap(__obf_2107c594e6166770, "root.pem: Failed to read file.")
	} else {
		__obf_9d4dae3e360d9c62, __obf_3f65e84b06a2e11c := pem.Decode(__obf_304dce20a2343d37)
		if __obf_9d4dae3e360d9c62 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_3f65e84b06a2e11c)
		}

		return x509.ParseCertificate(__obf_9d4dae3e360d9c62.Bytes)
	}
}

// 私钥解析
func __obf_68dcd9229cef6443() (__obf_c1c829b3ef681f91 any, __obf_2107c594e6166770 error) {
	if __obf_304dce20a2343d37, __obf_2107c594e6166770 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_b832ed210a876042, string(os.PathSeparator))); __obf_2107c594e6166770 != nil {
		return nil, errors.Wrap(__obf_2107c594e6166770, "root.key: Failed to read file.")
	} else {
		__obf_9d4dae3e360d9c62, __obf_3f65e84b06a2e11c := pem.Decode(__obf_304dce20a2343d37)
		if __obf_9d4dae3e360d9c62 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_3f65e84b06a2e11c)
		}

		return x509.ParsePKCS8PrivateKey(__obf_9d4dae3e360d9c62.Bytes)
	}
}

func __obf_b871b5172c86b1b3() *big.Int {
	__obf_5cd32d5518776f18 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_5684133436720528, __obf_2107c594e6166770 := rand.Int(rand.Reader, __obf_5cd32d5518776f18); __obf_2107c594e6166770 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_5684133436720528
	}
}

func GenRootCA() {
	__obf_a014454ed0992bbd, __obf_7ca60a74ae4191fa := __obf_24158a440d100e75(*__obf_f179418df50a4ef1, *__obf_013e90306300e806)
	__obf_7f71517510974c14 := &x509.Certificate{
		SerialNumber: __obf_b871b5172c86b1b3(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_9f01f358f5287bb0},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_a014454ed0992bbd,
		NotAfter:              __obf_7ca60a74ae4191fa,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_fb5bb54bb36cf6fc := __obf_959da6110d2640cd()
	__obf_6b123c9194577d08("root", __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc, __obf_7f71517510974c14, nil)

}

func GenServerCA() {
	__obf_a014454ed0992bbd, __obf_7ca60a74ae4191fa := __obf_24158a440d100e75(*__obf_5dae3cf5185920ed, *__obf_14693b8ee36fd73b)
	__obf_c0692453855fca4d := &x509.Certificate{
		SerialNumber: __obf_b871b5172c86b1b3(),
		Subject: pkix.Name{
			Organization: []string{*__obf_83f5e2108cdf9dd8},
		},
		NotBefore: __obf_a014454ed0992bbd,
		NotAfter:  __obf_7ca60a74ae4191fa,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_1ddf20596d21c921 := strings.Split(*__obf_fb9853e4bd03851d, ",")

	for _, __obf_2e59dd0944bc9f24 := range __obf_1ddf20596d21c921 {
		if __obf_f5a5e5d5216c8535 := net.ParseIP(__obf_2e59dd0944bc9f24); __obf_f5a5e5d5216c8535 != nil {
			__obf_c0692453855fca4d.
				IPAddresses = append(__obf_c0692453855fca4d.IPAddresses, __obf_f5a5e5d5216c8535)
		} else {
			__obf_c0692453855fca4d.
				DNSNames = append(__obf_c0692453855fca4d.DNSNames, __obf_2e59dd0944bc9f24)
		}
	}
	__obf_84eaf4818e78c714 := __obf_959da6110d2640cd()

	var __obf_7f71517510974c14 *x509.Certificate
	var __obf_2107c594e6166770 error
	var __obf_fb5bb54bb36cf6fc any
	__obf_7f71517510974c14, __obf_2107c594e6166770 = __obf_ca4f70ed37c87c67()
	if __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
	__obf_fb5bb54bb36cf6fc, __obf_2107c594e6166770 = __obf_68dcd9229cef6443()
	if __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
	__obf_6b123c9194577d08("server", __obf_c0692453855fca4d, __obf_84eaf4818e78c714, __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc)
}

func GenClientCA() {
	__obf_a014454ed0992bbd, __obf_7ca60a74ae4191fa := __obf_24158a440d100e75(*__obf_f8a1ce35efbbcd8b, *__obf_6a6634cefa52690f)
	__obf_7fe838c84a1220fa := &x509.Certificate{
		SerialNumber: __obf_b871b5172c86b1b3(),
		Subject: pkix.Name{
			Organization: []string{*__obf_302bcd74543539f2},
		},
		NotBefore: __obf_a014454ed0992bbd,
		NotAfter:  __obf_7ca60a74ae4191fa,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_b63244d999491ec0 := __obf_959da6110d2640cd()

	var __obf_7f71517510974c14 *x509.Certificate
	var __obf_2107c594e6166770 error
	var __obf_fb5bb54bb36cf6fc any
	__obf_7f71517510974c14, __obf_2107c594e6166770 = __obf_ca4f70ed37c87c67()
	if __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
	__obf_fb5bb54bb36cf6fc, __obf_2107c594e6166770 = __obf_68dcd9229cef6443()
	if __obf_2107c594e6166770 != nil {
		log.Fatalln(__obf_2107c594e6166770)
	}
	__obf_6b123c9194577d08("client", __obf_7fe838c84a1220fa, __obf_b63244d999491ec0, __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc)

}
