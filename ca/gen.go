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

func __obf_704491f843d9aad9(__obf_8dfe9dc499acf7f1 string, __obf_8399d76af1dec67f int) (__obf_8952901113253fb5 time.Time, __obf_75c84bb8e3dcaba2 time.Time) {
	var __obf_8b98d286773a1eea error
	if __obf_8dfe9dc499acf7f1 == "" {
		__obf_8952901113253fb5 = time.Now()
	} else {
		if __obf_8952901113253fb5, __obf_8b98d286773a1eea = time.Parse("2006-01-02", __obf_8dfe9dc499acf7f1); __obf_8b98d286773a1eea != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_75c84bb8e3dcaba2 = __obf_8952901113253fb5.Add(time.Duration(__obf_8399d76af1dec67f*24) * time.Hour)
		}
	}
	return
}

func __obf_5bb458c8b36f678b() any {
	var __obf_8b98d286773a1eea error
	var __obf_61e0a98f3b5bf932 any
	switch *__obf_9c8c2e6a9197c48c {
	case "":
		if *__obf_88fc8a0958f00870 {
			_, __obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = rsa.GenerateKey(rand.Reader, *__obf_ed53a6cf1d81a796)
		}
	case "P224":
		__obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_61e0a98f3b5bf932, __obf_8b98d286773a1eea = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_9c8c2e6a9197c48c)
	}

	if __obf_8b98d286773a1eea != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_8b98d286773a1eea)
	}

	return __obf_61e0a98f3b5bf932
}

func __obf_899a6824827409e5(__obf_3777679e49fe3f7b, __obf_dcb050a7b1e33ef7 string, __obf_57f6fdf657df089b []byte, __obf_cd28ed9c8db9fa52 os.FileMode) error {

	if _, __obf_8b98d286773a1eea := os.Stat(__obf_3777679e49fe3f7b); __obf_8b98d286773a1eea == nil {
		if __obf_8b98d286773a1eea = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_3777679e49fe3f7b, string(os.PathSeparator), __obf_dcb050a7b1e33ef7), __obf_57f6fdf657df089b, __obf_cd28ed9c8db9fa52); __obf_8b98d286773a1eea != nil {
			return errors.Wrap(__obf_8b98d286773a1eea, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_8b98d286773a1eea) {
		return errors.Wrap(__obf_8b98d286773a1eea, "Directory not exist.")
	} else {
		return __obf_8b98d286773a1eea
	}
	return nil
}

func __obf_a0e6f9942171fb68(__obf_94fa2b121e9208a2 string, __obf_89b937c04944cb2b *x509.Certificate, __obf_28a277ccea818a74 any, __obf_19492af16f0f666d *x509.Certificate, __obf_a74fc724d9a30e04 any) {

	var __obf_4e690447a56017a0 any
	switch __obf_c9be9488dbae2196 := __obf_28a277ccea818a74.(type) {
	case *rsa.PrivateKey:
		__obf_4e690447a56017a0 = &__obf_c9be9488dbae2196.PublicKey
	case *ecdsa.PrivateKey:
		__obf_4e690447a56017a0 = &__obf_c9be9488dbae2196.PublicKey
	case ed25519.PrivateKey:
		__obf_4e690447a56017a0 = __obf_c9be9488dbae2196.Public().(ed25519.PublicKey)
	default:
		__obf_4e690447a56017a0 = nil
	}

	var __obf_f42003db463ece31 any
	if __obf_a74fc724d9a30e04 != nil {
		__obf_f42003db463ece31 = __obf_a74fc724d9a30e04
	} else {
		__obf_f42003db463ece31 = __obf_28a277ccea818a74
	}
	__obf_cc085b2bdcc9bc80, __obf_8b98d286773a1eea := x509.CreateCertificate(rand.Reader, __obf_89b937c04944cb2b, __obf_19492af16f0f666d, __obf_4e690447a56017a0, __obf_f42003db463ece31)
	if __obf_8b98d286773a1eea != nil {
		log.Println("create failed", __obf_8b98d286773a1eea)
		return
	}
	__obf_f865972c5b507991 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_cc085b2bdcc9bc80})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_8b98d286773a1eea = __obf_899a6824827409e5(*__obf_527251753bbe5828, __obf_94fa2b121e9208a2+".pem", __obf_f865972c5b507991, 0777); __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
	__obf_ae4fc2cbb7410041, __obf_8b98d286773a1eea := x509.MarshalPKCS8PrivateKey(__obf_28a277ccea818a74)
	if __obf_8b98d286773a1eea != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_cf30594748a1412b := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_ae4fc2cbb7410041})

	if __obf_8b98d286773a1eea = __obf_899a6824827409e5(*__obf_527251753bbe5828, __obf_94fa2b121e9208a2+".key", __obf_cf30594748a1412b, 0777); __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
}

// 公钥证书解析
func __obf_60e1b858637ea3fd() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_5dbd41a4908e0496, __obf_8b98d286773a1eea := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_527251753bbe5828, string(os.PathSeparator))); __obf_8b98d286773a1eea != nil {
		return nil, errors.Wrap(__obf_8b98d286773a1eea, "root.pem: Failed to read file.")
	} else {
		__obf_7adc7350872191bd, __obf_43a303256282a76e := pem.Decode(__obf_5dbd41a4908e0496)
		if __obf_7adc7350872191bd == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_43a303256282a76e)
		}

		return x509.ParseCertificate(__obf_7adc7350872191bd.Bytes)
	}
}

// 私钥解析
func __obf_7d94847bb46dec22() (__obf_28a277ccea818a74 any, __obf_8b98d286773a1eea error) {
	if __obf_5dbd41a4908e0496, __obf_8b98d286773a1eea := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_527251753bbe5828, string(os.PathSeparator))); __obf_8b98d286773a1eea != nil {
		return nil, errors.Wrap(__obf_8b98d286773a1eea, "root.key: Failed to read file.")
	} else {
		__obf_7adc7350872191bd, __obf_43a303256282a76e := pem.Decode(__obf_5dbd41a4908e0496)
		if __obf_7adc7350872191bd == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_43a303256282a76e)
		}

		return x509.ParsePKCS8PrivateKey(__obf_7adc7350872191bd.Bytes)
	}
}

func __obf_3a616818890e9136() *big.Int {
	__obf_5eca5ed17b1af004 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_65ba8c82ecf6784e, __obf_8b98d286773a1eea := rand.Int(rand.Reader, __obf_5eca5ed17b1af004); __obf_8b98d286773a1eea != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_65ba8c82ecf6784e
	}
}

func GenRootCA() {
	__obf_df8ecfd08547baa9, __obf_bd69b06a640448e5 := __obf_704491f843d9aad9(*__obf_07179264e258cdb6, *__obf_879e45fa0c7555fa)
	__obf_bd646ef02df2a803 := &x509.Certificate{
		SerialNumber: __obf_3a616818890e9136(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_c29edf7ae9e78e22},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_df8ecfd08547baa9,
		NotAfter:              __obf_bd69b06a640448e5,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_031842f1c2ed65a3 := __obf_5bb458c8b36f678b()
	__obf_a0e6f9942171fb68("root", __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3, __obf_bd646ef02df2a803, nil)

}

func GenServerCA() {
	__obf_df8ecfd08547baa9, __obf_bd69b06a640448e5 := __obf_704491f843d9aad9(*__obf_7b61af43155d6a85, *__obf_f035b210ead5b9d7)
	__obf_441b7d9118c61425 := &x509.Certificate{
		SerialNumber: __obf_3a616818890e9136(),
		Subject: pkix.Name{
			Organization: []string{*__obf_18060d85d772c392},
		},
		NotBefore: __obf_df8ecfd08547baa9,
		NotAfter:  __obf_bd69b06a640448e5,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_0f57c4cf881d27c7 := strings.Split(*__obf_e31bc42d46a4cadc, ",")

	for _, __obf_64a53c251efc3cd6 := range __obf_0f57c4cf881d27c7 {
		if __obf_663e92eef5c260d0 := net.ParseIP(__obf_64a53c251efc3cd6); __obf_663e92eef5c260d0 != nil {
			__obf_441b7d9118c61425.
				IPAddresses = append(__obf_441b7d9118c61425.IPAddresses, __obf_663e92eef5c260d0)
		} else {
			__obf_441b7d9118c61425.
				DNSNames = append(__obf_441b7d9118c61425.DNSNames, __obf_64a53c251efc3cd6)
		}
	}
	__obf_f02f58ad7a9d0c1f := __obf_5bb458c8b36f678b()

	var __obf_bd646ef02df2a803 *x509.Certificate
	var __obf_8b98d286773a1eea error
	var __obf_031842f1c2ed65a3 any
	__obf_bd646ef02df2a803, __obf_8b98d286773a1eea = __obf_60e1b858637ea3fd()
	if __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
	__obf_031842f1c2ed65a3, __obf_8b98d286773a1eea = __obf_7d94847bb46dec22()
	if __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
	__obf_a0e6f9942171fb68("server", __obf_441b7d9118c61425, __obf_f02f58ad7a9d0c1f, __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3)
}

func GenClientCA() {
	__obf_df8ecfd08547baa9, __obf_bd69b06a640448e5 := __obf_704491f843d9aad9(*__obf_f442c10694402f79, *__obf_09177447b8c2c966)
	__obf_15eb203631991e8f := &x509.Certificate{
		SerialNumber: __obf_3a616818890e9136(),
		Subject: pkix.Name{
			Organization: []string{*__obf_a0186db0e5770e5d},
		},
		NotBefore: __obf_df8ecfd08547baa9,
		NotAfter:  __obf_bd69b06a640448e5,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_a1788deecf5df7ee := __obf_5bb458c8b36f678b()

	var __obf_bd646ef02df2a803 *x509.Certificate
	var __obf_8b98d286773a1eea error
	var __obf_031842f1c2ed65a3 any
	__obf_bd646ef02df2a803, __obf_8b98d286773a1eea = __obf_60e1b858637ea3fd()
	if __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
	__obf_031842f1c2ed65a3, __obf_8b98d286773a1eea = __obf_7d94847bb46dec22()
	if __obf_8b98d286773a1eea != nil {
		log.Fatalln(__obf_8b98d286773a1eea)
	}
	__obf_a0e6f9942171fb68("client", __obf_15eb203631991e8f, __obf_a1788deecf5df7ee, __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3)

}
