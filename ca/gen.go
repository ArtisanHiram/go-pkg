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

func __obf_032f2ec5a0eac80f(__obf_770f1ad1e1fd72d0 string, __obf_43806ac0f4f83280 int) (__obf_33889d1f7588078d time.Time, __obf_be228e6b4d822c54 time.Time) {
	var __obf_2ec52a9654271348 error
	if __obf_770f1ad1e1fd72d0 == "" {
		__obf_33889d1f7588078d = time.Now()
	} else {
		if __obf_33889d1f7588078d, __obf_2ec52a9654271348 = time.Parse("2006-01-02", __obf_770f1ad1e1fd72d0); __obf_2ec52a9654271348 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_be228e6b4d822c54 = __obf_33889d1f7588078d.Add(time.Duration(__obf_43806ac0f4f83280*24) * time.Hour)
		}
	}
	return
}

func __obf_4922cb44f2003694() any {
	var __obf_2ec52a9654271348 error
	var __obf_3d6e42833081d2b9 any
	switch *__obf_f5980e672145c29f {
	case "":
		if *__obf_687f16797dbe5480 {
			_, __obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = rsa.GenerateKey(rand.Reader, *__obf_244f5fea7ddf0c82)
		}
	case "P224":
		__obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_3d6e42833081d2b9, __obf_2ec52a9654271348 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_f5980e672145c29f)
	}

	if __obf_2ec52a9654271348 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_2ec52a9654271348)
	}

	return __obf_3d6e42833081d2b9
}

func __obf_ded3b7d03569731f(__obf_053a008c64910bdb, __obf_2fae55c305fb99f3 string, __obf_706a59d0e2ae4d94 []byte, __obf_9660f3c53970e2ec os.FileMode) error {

	if _, __obf_2ec52a9654271348 := os.Stat(__obf_053a008c64910bdb); __obf_2ec52a9654271348 == nil {
		if __obf_2ec52a9654271348 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_053a008c64910bdb, string(os.PathSeparator), __obf_2fae55c305fb99f3), __obf_706a59d0e2ae4d94, __obf_9660f3c53970e2ec); __obf_2ec52a9654271348 != nil {
			return errors.Wrap(__obf_2ec52a9654271348, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_2ec52a9654271348) {
		return errors.Wrap(__obf_2ec52a9654271348, "Directory not exist.")
	} else {
		return __obf_2ec52a9654271348
	}
	return nil
}

func __obf_b95ec46d49436b0b(__obf_1ed29960d2e595e7 string, __obf_c7be66b042b1a623 *x509.Certificate, __obf_2f56a8cdb16bd159 any, __obf_f5dae7569058304a *x509.Certificate, __obf_3b0f2a731719128f any) {

	var __obf_bfa26ca12c831807 any
	switch __obf_b0d88f72860f9741 := __obf_2f56a8cdb16bd159.(type) {
	case *rsa.PrivateKey:
		__obf_bfa26ca12c831807 = &__obf_b0d88f72860f9741.PublicKey
	case *ecdsa.PrivateKey:
		__obf_bfa26ca12c831807 = &__obf_b0d88f72860f9741.PublicKey
	case ed25519.PrivateKey:
		__obf_bfa26ca12c831807 = __obf_b0d88f72860f9741.Public().(ed25519.PublicKey)
	default:
		__obf_bfa26ca12c831807 = nil
	}

	var __obf_7515f5b04ab4fdd3 any
	if __obf_3b0f2a731719128f != nil {
		__obf_7515f5b04ab4fdd3 = __obf_3b0f2a731719128f
	} else {
		__obf_7515f5b04ab4fdd3 = __obf_2f56a8cdb16bd159
	}
	__obf_ccfcf99a3da52331, __obf_2ec52a9654271348 := x509.CreateCertificate(rand.Reader, __obf_c7be66b042b1a623, __obf_f5dae7569058304a, __obf_bfa26ca12c831807, __obf_7515f5b04ab4fdd3)
	if __obf_2ec52a9654271348 != nil {
		log.Println("create failed", __obf_2ec52a9654271348)
		return
	}
	__obf_a8076ac87f764dce := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_ccfcf99a3da52331})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_2ec52a9654271348 = __obf_ded3b7d03569731f(*__obf_e7dc96daf488b4d7, __obf_1ed29960d2e595e7+".pem", __obf_a8076ac87f764dce, 0777); __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
	__obf_c8c65e857103e73b, __obf_2ec52a9654271348 := x509.MarshalPKCS8PrivateKey(__obf_2f56a8cdb16bd159)
	if __obf_2ec52a9654271348 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	__obf_ad33ff08cd765f4b := // os.WriteFile(name+".key", priv_b, 0777)

		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_c8c65e857103e73b})

	if __obf_2ec52a9654271348 = __obf_ded3b7d03569731f(*__obf_e7dc96daf488b4d7, __obf_1ed29960d2e595e7+".key", __obf_ad33ff08cd765f4b, 0777); __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
}

// 公钥证书解析
func __obf_88f0f7635949dc5b() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_525403725fe58302, __obf_2ec52a9654271348 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_e7dc96daf488b4d7, string(os.PathSeparator))); __obf_2ec52a9654271348 != nil {
		return nil, errors.Wrap(__obf_2ec52a9654271348, "root.pem: Failed to read file.")
	} else {
		__obf_3939ce88260a2024, __obf_a68ed9dc7813bcf2 := pem.Decode(__obf_525403725fe58302)
		if __obf_3939ce88260a2024 == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_a68ed9dc7813bcf2)
		}

		return x509.ParseCertificate(__obf_3939ce88260a2024.Bytes)
	}
}

// 私钥解析
func __obf_70f12451da1937d1() (__obf_2f56a8cdb16bd159 any, __obf_2ec52a9654271348 error) {
	if __obf_525403725fe58302, __obf_2ec52a9654271348 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_e7dc96daf488b4d7, string(os.PathSeparator))); __obf_2ec52a9654271348 != nil {
		return nil, errors.Wrap(__obf_2ec52a9654271348, "root.key: Failed to read file.")
	} else {
		__obf_3939ce88260a2024, __obf_a68ed9dc7813bcf2 := pem.Decode(__obf_525403725fe58302)
		if __obf_3939ce88260a2024 == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_a68ed9dc7813bcf2)
		}

		return x509.ParsePKCS8PrivateKey(__obf_3939ce88260a2024.Bytes)
	}
}

func __obf_2f0d31a7316c7293() *big.Int {
	__obf_3178f3a90e5dc3b8 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_479c34d1d5c0389d, __obf_2ec52a9654271348 := rand.Int(rand.Reader, __obf_3178f3a90e5dc3b8); __obf_2ec52a9654271348 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_479c34d1d5c0389d
	}
}

func GenRootCA() {
	__obf_8ca71a6182159260, __obf_c2aab64f9a9a4357 := __obf_032f2ec5a0eac80f(*__obf_bf317b5f1f2b6382, *__obf_c82831f575e7670b)
	__obf_cccdf1a27493dddc := &x509.Certificate{
		SerialNumber: __obf_2f0d31a7316c7293(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_5723a8a1754a6ff7},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_8ca71a6182159260,
		NotAfter:              __obf_c2aab64f9a9a4357,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_1d68b8dfdd86e587 := __obf_4922cb44f2003694()
	__obf_b95ec46d49436b0b("root", __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587, __obf_cccdf1a27493dddc, nil)

}

func GenServerCA() {
	__obf_8ca71a6182159260, __obf_c2aab64f9a9a4357 := __obf_032f2ec5a0eac80f(*__obf_28837c6c18a9b3d9, *__obf_e1e272065c9d4c14)
	__obf_28da83af6f8edbe6 := &x509.Certificate{
		SerialNumber: __obf_2f0d31a7316c7293(),
		Subject: pkix.Name{
			Organization: []string{*__obf_fe2344381e73ffee},
		},
		NotBefore: __obf_8ca71a6182159260,
		NotAfter:  __obf_c2aab64f9a9a4357,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_1b90c90256dcecac := strings.Split(*__obf_7c9a6812bf2314f9, ",")

	for _, __obf_56731192f777ebe3 := range __obf_1b90c90256dcecac {
		if __obf_34b6469a7cf51aaa := net.ParseIP(__obf_56731192f777ebe3); __obf_34b6469a7cf51aaa != nil {
			__obf_28da83af6f8edbe6.
				IPAddresses = append(__obf_28da83af6f8edbe6.IPAddresses, __obf_34b6469a7cf51aaa)
		} else {
			__obf_28da83af6f8edbe6.
				DNSNames = append(__obf_28da83af6f8edbe6.DNSNames, __obf_56731192f777ebe3)
		}
	}
	__obf_0dd6962322487601 := __obf_4922cb44f2003694()

	var __obf_cccdf1a27493dddc *x509.Certificate
	var __obf_2ec52a9654271348 error
	var __obf_1d68b8dfdd86e587 any
	__obf_cccdf1a27493dddc, __obf_2ec52a9654271348 = __obf_88f0f7635949dc5b()
	if __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
	__obf_1d68b8dfdd86e587, __obf_2ec52a9654271348 = __obf_70f12451da1937d1()
	if __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
	__obf_b95ec46d49436b0b("server", __obf_28da83af6f8edbe6, __obf_0dd6962322487601, __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587)
}

func GenClientCA() {
	__obf_8ca71a6182159260, __obf_c2aab64f9a9a4357 := __obf_032f2ec5a0eac80f(*__obf_af7fc1426a450125, *__obf_b31dfc092fcc3a05)
	__obf_9c2efd023b154ad3 := &x509.Certificate{
		SerialNumber: __obf_2f0d31a7316c7293(),
		Subject: pkix.Name{
			Organization: []string{*__obf_dc0cb0d6e620a6ab},
		},
		NotBefore: __obf_8ca71a6182159260,
		NotAfter:  __obf_c2aab64f9a9a4357,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_3a8c2f70431f5f7c := __obf_4922cb44f2003694()

	var __obf_cccdf1a27493dddc *x509.Certificate
	var __obf_2ec52a9654271348 error
	var __obf_1d68b8dfdd86e587 any
	__obf_cccdf1a27493dddc, __obf_2ec52a9654271348 = __obf_88f0f7635949dc5b()
	if __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
	__obf_1d68b8dfdd86e587, __obf_2ec52a9654271348 = __obf_70f12451da1937d1()
	if __obf_2ec52a9654271348 != nil {
		log.Fatalln(__obf_2ec52a9654271348)
	}
	__obf_b95ec46d49436b0b("client", __obf_9c2efd023b154ad3, __obf_3a8c2f70431f5f7c, __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587)

}
