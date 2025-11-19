package __obf_cad85c1a84aaff79

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func Main() {
	__obf_fb8d4a3a0c3f4ce6 := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_f5823a3b17eb143f, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f, __obf_fb8d4a3a0c3f4ce6, nil)
	__obf_7c90e3cc678b8812 := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"SERVER"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_b95b40fd5e0c8be7 := []string{"localhost", "127.0.0.1"}
	for _, __obf_df4ac6e7e1c79d0b := range __obf_b95b40fd5e0c8be7 {
		if __obf_ab36e3b42276b230 := net.ParseIP(__obf_df4ac6e7e1c79d0b); __obf_ab36e3b42276b230 != nil {
			__obf_7c90e3cc678b8812.IPAddresses = append(__obf_7c90e3cc678b8812.IPAddresses, __obf_ab36e3b42276b230)
		} else {
			__obf_7c90e3cc678b8812.DNSNames = append(__obf_7c90e3cc678b8812.DNSNames, __obf_df4ac6e7e1c79d0b)
		}
	}
	__obf_78af2f561b6d1b25, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_7c90e3cc678b8812, __obf_78af2f561b6d1b25, __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f)
	__obf_cb8a63fd2e71ecee := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"CLIENT"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	__obf_a4fa9ee49e8d1aa1, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_cb8a63fd2e71ecee, __obf_a4fa9ee49e8d1aa1, __obf_fb8d4a3a0c3f4ce6, __obf_f5823a3b17eb143f)
}

func CreateCertificateFile(__obf_0663138c5d22b02f string, __obf_c229ddd9636a5c4d *x509.Certificate, __obf_6d7fb6e36aca1063 *rsa.PrivateKey, __obf_708eed607e067687 *x509.Certificate, __obf_87b1ecc11fc358b0 *rsa.PrivateKey) {
	__obf_7713b116c8940aa0 := __obf_6d7fb6e36aca1063
	__obf_2c3dff0eafd5bcbf := &__obf_7713b116c8940aa0.PublicKey
	__obf_e3e734b16b9c09db := __obf_7713b116c8940aa0
	if __obf_87b1ecc11fc358b0 != nil {
		__obf_e3e734b16b9c09db = __obf_87b1ecc11fc358b0
	}
	__obf_e162a29f5c6f8be8, __obf_80c4653dad68e783 := x509.CreateCertificate(rand.Reader, __obf_c229ddd9636a5c4d, __obf_708eed607e067687, __obf_2c3dff0eafd5bcbf, __obf_e3e734b16b9c09db)
	if __obf_80c4653dad68e783 != nil {
		log.Println("create failed", __obf_80c4653dad68e783)
		return
	}
	__obf_40275c15341ef3b9 := __obf_0663138c5d22b02f + ".pem"
	log.Println("write to pem", __obf_40275c15341ef3b9)
	var __obf_deacfbce1461fa5e = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_e162a29f5c6f8be8}
	__obf_d1f61c7d02f317eb := pem.EncodeToMemory(__obf_deacfbce1461fa5e)
	os.WriteFile(__obf_40275c15341ef3b9, __obf_d1f61c7d02f317eb, 0777)

	__obf_643e704ddf1821e7 := __obf_0663138c5d22b02f + ".key"
	__obf_ec0f84cae0075790 := x509.MarshalPKCS1PrivateKey(__obf_7713b116c8940aa0)
	log.Println("write to key", __obf_643e704ddf1821e7)
	os.WriteFile(__obf_643e704ddf1821e7, __obf_ec0f84cae0075790, 0777)
	var __obf_93a6f8e7720a7e20 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_ec0f84cae0075790}
	__obf_ddb57861ca6f06c7 := pem.EncodeToMemory(__obf_93a6f8e7720a7e20)
	os.WriteFile(__obf_643e704ddf1821e7, __obf_ddb57861ca6f06c7, 0777)
}
