package __obf_cad85c1a84aaff79

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

func __obf_1693741ba18575bc(__obf_3c797b61155aa4de string, __obf_d22ac845a785740e int) (__obf_34973c1e53144c2c time.Time, __obf_777cdd3e535aff85 time.Time) {
	var __obf_80c4653dad68e783 error
	if __obf_3c797b61155aa4de == "" {
		__obf_34973c1e53144c2c = time.Now()
	} else {
		if __obf_34973c1e53144c2c, __obf_80c4653dad68e783 = time.Parse("2006-01-02", __obf_3c797b61155aa4de); __obf_80c4653dad68e783 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_777cdd3e535aff85 = __obf_34973c1e53144c2c.Add(time.Duration(__obf_d22ac845a785740e*24) * time.Hour)
		}
	}
	return
}

func __obf_31b68b98c549cee2() any {
	var __obf_80c4653dad68e783 error
	var __obf_7713b116c8940aa0 any
	switch *__obf_31a37a320eb83526 {
	case "":
		if *__obf_d8991a976dff20bf {
			_, __obf_7713b116c8940aa0, __obf_80c4653dad68e783 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_7713b116c8940aa0, __obf_80c4653dad68e783 = rsa.GenerateKey(rand.Reader, *__obf_3a3f21e4c08578dc)
		}
	case "P224":
		__obf_7713b116c8940aa0, __obf_80c4653dad68e783 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_7713b116c8940aa0, __obf_80c4653dad68e783 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_7713b116c8940aa0, __obf_80c4653dad68e783 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_7713b116c8940aa0, __obf_80c4653dad68e783 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_31a37a320eb83526)
	}

	if __obf_80c4653dad68e783 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_80c4653dad68e783)
	}

	return __obf_7713b116c8940aa0
}

func __obf_61ec5d1f5f9302c5(__obf_70179ce67faede8c, __obf_3c34a187283003b3 string, __obf_9efce1ed16feec63 []byte, __obf_71708c172a12ee0a os.FileMode) error {

	if _, __obf_80c4653dad68e783 := os.Stat(__obf_70179ce67faede8c); __obf_80c4653dad68e783 == nil {
		if __obf_80c4653dad68e783 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_70179ce67faede8c, string(os.PathSeparator), __obf_3c34a187283003b3), __obf_9efce1ed16feec63, __obf_71708c172a12ee0a); __obf_80c4653dad68e783 != nil {
			return errors.Wrap(__obf_80c4653dad68e783, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_80c4653dad68e783) {
		return errors.Wrap(__obf_80c4653dad68e783, "Directory not exist.")
	} else {
		return __obf_80c4653dad68e783
	}
	return nil
}

func __obf_7023345c9c093f7f(__obf_0663138c5d22b02f string, __obf_c229ddd9636a5c4d *x509.Certificate, __obf_6d7fb6e36aca1063 any, __obf_708eed607e067687 *x509.Certificate, __obf_87b1ecc11fc358b0 any) {

	var __obf_2c3dff0eafd5bcbf any
	switch __obf_ab3b0a8a48226fc7 := __obf_6d7fb6e36aca1063.(type) {
	case *rsa.PrivateKey:
		__obf_2c3dff0eafd5bcbf = &__obf_ab3b0a8a48226fc7.PublicKey
	case *ecdsa.PrivateKey:
		__obf_2c3dff0eafd5bcbf = &__obf_ab3b0a8a48226fc7.PublicKey
	case ed25519.PrivateKey:
		__obf_2c3dff0eafd5bcbf = __obf_ab3b0a8a48226fc7.Public().(ed25519.PublicKey)
	default:
		__obf_2c3dff0eafd5bcbf = nil
	}

	var __obf_e3e734b16b9c09db any
	if __obf_87b1ecc11fc358b0 != nil {
		__obf_e3e734b16b9c09db = __obf_87b1ecc11fc358b0
	} else {
		__obf_e3e734b16b9c09db = __obf_6d7fb6e36aca1063
	}
	__obf_e162a29f5c6f8be8, __obf_80c4653dad68e783 := x509.CreateCertificate(rand.Reader, __obf_c229ddd9636a5c4d, __obf_708eed607e067687, __obf_2c3dff0eafd5bcbf, __obf_e3e734b16b9c09db)
	if __obf_80c4653dad68e783 != nil {
		log.Println("create failed", __obf_80c4653dad68e783)
		return
	}

	__obf_d1f61c7d02f317eb := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_e162a29f5c6f8be8})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_80c4653dad68e783 = __obf_61ec5d1f5f9302c5(*__obf_d7fd1eaf7db608e6, __obf_0663138c5d22b02f+".pem", __obf_d1f61c7d02f317eb, 0777); __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}

	__obf_ec0f84cae0075790, __obf_80c4653dad68e783 := x509.MarshalPKCS8PrivateKey(__obf_6d7fb6e36aca1063)
	if __obf_80c4653dad68e783 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_ddb57861ca6f06c7 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_ec0f84cae0075790})

	if __obf_80c4653dad68e783 = __obf_61ec5d1f5f9302c5(*__obf_d7fd1eaf7db608e6, __obf_0663138c5d22b02f+".key", __obf_ddb57861ca6f06c7, 0777); __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}
}

// 公钥证书解析
func __obf_e59852709ce60404() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_6626ffe4ab4e050b, __obf_80c4653dad68e783 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_d7fd1eaf7db608e6, string(os.PathSeparator))); __obf_80c4653dad68e783 != nil {
		return nil, errors.Wrap(__obf_80c4653dad68e783, "root.pem: Failed to read file.")
	} else {
		__obf_d4000f8242b75e3f, __obf_704cfb87ccbacf57 := pem.Decode(__obf_6626ffe4ab4e050b)
		if __obf_d4000f8242b75e3f == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_704cfb87ccbacf57)
		}

		return x509.ParseCertificate(__obf_d4000f8242b75e3f.Bytes)
	}
}

// 私钥解析
func __obf_6a83a70899cbc5dd() (__obf_6d7fb6e36aca1063 any, __obf_80c4653dad68e783 error) {
	if __obf_6626ffe4ab4e050b, __obf_80c4653dad68e783 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_d7fd1eaf7db608e6, string(os.PathSeparator))); __obf_80c4653dad68e783 != nil {
		return nil, errors.Wrap(__obf_80c4653dad68e783, "root.key: Failed to read file.")
	} else {
		__obf_d4000f8242b75e3f, __obf_704cfb87ccbacf57 := pem.Decode(__obf_6626ffe4ab4e050b)
		if __obf_d4000f8242b75e3f == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_704cfb87ccbacf57)
		}

		return x509.ParsePKCS8PrivateKey(__obf_d4000f8242b75e3f.Bytes)
	}
}

func __obf_7d6b44187ea218fd() *big.Int {
	__obf_5001cee016b88020 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_a7990181ef54b398, __obf_80c4653dad68e783 := rand.Int(rand.Reader, __obf_5001cee016b88020); __obf_80c4653dad68e783 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_a7990181ef54b398
	}
}

func GenRootCA() {

	__obf_725c0422259534dc, __obf_74fe861d9b9faef0 := __obf_1693741ba18575bc(*__obf_f62a2d95847ff0f5, *__obf_eb61da3b8dc27fa4)
	__obf_fb8d4a3a0c3f4ce6 := &x509.Certificate{
		SerialNumber: __obf_7d6b44187ea218fd(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_dbbc139bf2b03050},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_725c0422259534dc,
		NotAfter:              __obf_74fe861d9b9faef0,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_f5823a3b17eb143f := __obf_31b68b98c549cee2()

	__obf_7023345c9c093f7f("root", __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f, __obf_fb8d4a3a0c3f4ce6, nil)

}

func GenServerCA() {

	__obf_725c0422259534dc, __obf_74fe861d9b9faef0 := __obf_1693741ba18575bc(*__obf_c1cc1283cf2bcf00, *__obf_6dc1feab678126da)
	__obf_7c90e3cc678b8812 := &x509.Certificate{
		SerialNumber: __obf_7d6b44187ea218fd(),
		Subject: pkix.Name{
			Organization: []string{*__obf_f782959824cfef0c},
		},
		NotBefore: __obf_725c0422259534dc,
		NotAfter:  __obf_74fe861d9b9faef0,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_b95b40fd5e0c8be7 := strings.Split(*__obf_967bcefeb68b3bc2, ",")

	for _, __obf_df4ac6e7e1c79d0b := range __obf_b95b40fd5e0c8be7 {
		if __obf_ab36e3b42276b230 := net.ParseIP(__obf_df4ac6e7e1c79d0b); __obf_ab36e3b42276b230 != nil {
			__obf_7c90e3cc678b8812.IPAddresses = append(__obf_7c90e3cc678b8812.IPAddresses, __obf_ab36e3b42276b230)
		} else {
			__obf_7c90e3cc678b8812.DNSNames = append(__obf_7c90e3cc678b8812.DNSNames, __obf_df4ac6e7e1c79d0b)
		}
	}
	__obf_78af2f561b6d1b25 := __obf_31b68b98c549cee2()

	var __obf_fb8d4a3a0c3f4ce6 *x509.Certificate
	var __obf_80c4653dad68e783 error
	var __obf_f5823a3b17eb143f any

	__obf_fb8d4a3a0c3f4ce6, __obf_80c4653dad68e783 = __obf_e59852709ce60404()
	if __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}

	__obf_f5823a3b17eb143f, __obf_80c4653dad68e783 = __obf_6a83a70899cbc5dd()
	if __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}

	__obf_7023345c9c093f7f("server", __obf_7c90e3cc678b8812, __obf_78af2f561b6d1b25, __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f)
}

func GenClientCA() {

	__obf_725c0422259534dc, __obf_74fe861d9b9faef0 := __obf_1693741ba18575bc(*__obf_29da549346963415, *__obf_3800c25206840554)
	__obf_cb8a63fd2e71ecee := &x509.Certificate{
		SerialNumber: __obf_7d6b44187ea218fd(),
		Subject: pkix.Name{
			Organization: []string{*__obf_5916aadfe17ba1cd},
		},
		NotBefore: __obf_725c0422259534dc,
		NotAfter:  __obf_74fe861d9b9faef0,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_a4fa9ee49e8d1aa1 := __obf_31b68b98c549cee2()

	var __obf_fb8d4a3a0c3f4ce6 *x509.Certificate
	var __obf_80c4653dad68e783 error
	var __obf_f5823a3b17eb143f any

	__obf_fb8d4a3a0c3f4ce6, __obf_80c4653dad68e783 = __obf_e59852709ce60404()
	if __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}

	__obf_f5823a3b17eb143f, __obf_80c4653dad68e783 = __obf_6a83a70899cbc5dd()
	if __obf_80c4653dad68e783 != nil {
		log.Fatalln(__obf_80c4653dad68e783)
	}

	__obf_7023345c9c093f7f("client", __obf_cb8a63fd2e71ecee, __obf_a4fa9ee49e8d1aa1, __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f)

}
