package __obf_b8dd0558a7deb9fa

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
	__obf_eb7ad60b95be299e := &x509.Certificate{
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
	__obf_548f732c582bb879, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_eb7ad60b95be299e, __obf_548f732c582bb879, __obf_eb7ad60b95be299e, nil)
	__obf_0c3f3422ff4ddc90 := &x509.Certificate{
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
	__obf_2ca22ef2f9c8e4f5 := []string{"localhost", "127.0.0.1"}
	for _, __obf_3551ec09edcea1aa := range __obf_2ca22ef2f9c8e4f5 {
		if __obf_5c0de70f8ac4bb34 := net.ParseIP(__obf_3551ec09edcea1aa); __obf_5c0de70f8ac4bb34 != nil {
			__obf_0c3f3422ff4ddc90.IPAddresses = append(__obf_0c3f3422ff4ddc90.IPAddresses, __obf_5c0de70f8ac4bb34)
		} else {
			__obf_0c3f3422ff4ddc90.DNSNames = append(__obf_0c3f3422ff4ddc90.DNSNames, __obf_3551ec09edcea1aa)
		}
	}
	__obf_4c186a057a362197, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_0c3f3422ff4ddc90, __obf_4c186a057a362197, __obf_eb7ad60b95be299e, __obf_548f732c582bb879)
	__obf_7bf9b9def77e79a4 := &x509.Certificate{
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
	__obf_7e4eec52bf5a4d32, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_7bf9b9def77e79a4, __obf_7e4eec52bf5a4d32, __obf_eb7ad60b95be299e, __obf_548f732c582bb879)
}

func CreateCertificateFile(__obf_e6d878397dd86eee string, __obf_35143b3a01b9448b *x509.Certificate, __obf_568d273391419dd1 *rsa.PrivateKey, __obf_751c8c53bc81fe11 *x509.Certificate, __obf_aa6de849c46a0f6e *rsa.PrivateKey) {
	__obf_200b0d942eda1a2a := __obf_568d273391419dd1
	__obf_3d0c48eca5fc749e := &__obf_200b0d942eda1a2a.PublicKey
	__obf_00af8fb2c1c6de80 := __obf_200b0d942eda1a2a
	if __obf_aa6de849c46a0f6e != nil {
		__obf_00af8fb2c1c6de80 = __obf_aa6de849c46a0f6e
	}
	__obf_df65b962ca4d67b5, __obf_b7fd5848550ca60c := x509.CreateCertificate(rand.Reader, __obf_35143b3a01b9448b, __obf_751c8c53bc81fe11, __obf_3d0c48eca5fc749e, __obf_00af8fb2c1c6de80)
	if __obf_b7fd5848550ca60c != nil {
		log.Println("create failed", __obf_b7fd5848550ca60c)
		return
	}
	__obf_69bf83fbc2c73aae := __obf_e6d878397dd86eee + ".pem"
	log.Println("write to pem", __obf_69bf83fbc2c73aae)
	var __obf_203357a2d9ad3e9e = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_df65b962ca4d67b5}
	__obf_a7c70595d99cbd5a := pem.EncodeToMemory(__obf_203357a2d9ad3e9e)
	os.WriteFile(__obf_69bf83fbc2c73aae, __obf_a7c70595d99cbd5a, 0777)

	__obf_d8ebc380651d8269 := __obf_e6d878397dd86eee + ".key"
	__obf_8bf120e25d7da167 := x509.MarshalPKCS1PrivateKey(__obf_200b0d942eda1a2a)
	log.Println("write to key", __obf_d8ebc380651d8269)
	os.WriteFile(__obf_d8ebc380651d8269, __obf_8bf120e25d7da167, 0777)
	var __obf_1ae5c92cacbe0b2a = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_8bf120e25d7da167}
	__obf_41614e79cb0c56d3 := pem.EncodeToMemory(__obf_1ae5c92cacbe0b2a)
	os.WriteFile(__obf_d8ebc380651d8269, __obf_41614e79cb0c56d3, 0777)
}
