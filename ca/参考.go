package __obf_a7fac689e11862d9

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
	__obf_f4caa47498522a63 := &x509.Certificate{
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
	__obf_c3f16a5d19c0b5bc, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc, __obf_f4caa47498522a63, nil)
	__obf_f98631a6dc322cfb := &x509.Certificate{
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
	__obf_51f4b0d669777373 := []string{"localhost", "127.0.0.1"}
	for _, __obf_4def855968424027 := range __obf_51f4b0d669777373 {
		if __obf_86931e0d9c3a793c := net.ParseIP(__obf_4def855968424027); __obf_86931e0d9c3a793c != nil {
			__obf_f98631a6dc322cfb.IPAddresses = append(__obf_f98631a6dc322cfb.IPAddresses, __obf_86931e0d9c3a793c)
		} else {
			__obf_f98631a6dc322cfb.DNSNames = append(__obf_f98631a6dc322cfb.DNSNames, __obf_4def855968424027)
		}
	}
	__obf_f5e6430087d57a6f, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_f98631a6dc322cfb, __obf_f5e6430087d57a6f, __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc)
	__obf_440c82f5cba2bc11 := &x509.Certificate{
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
	__obf_02cb005039f8f182, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_440c82f5cba2bc11, __obf_02cb005039f8f182, __obf_f4caa47498522a63, __obf_c3f16a5d19c0b5bc)
}

func CreateCertificateFile(__obf_735f5b74487e5b11 string, __obf_7fc0261a12598c19 *x509.Certificate, __obf_a264d37264bcfd3c *rsa.PrivateKey, __obf_e21f6160f7309bde *x509.Certificate, __obf_e29c547489483b7f *rsa.PrivateKey) {
	__obf_811665aa5ec324fa := __obf_a264d37264bcfd3c
	__obf_17b876ebc9671e6f := &__obf_811665aa5ec324fa.PublicKey
	__obf_85a14dae0b94c317 := __obf_811665aa5ec324fa
	if __obf_e29c547489483b7f != nil {
		__obf_85a14dae0b94c317 = __obf_e29c547489483b7f
	}
	__obf_edd4a003880f9f62, __obf_d388542c1816a2c3 := x509.CreateCertificate(rand.Reader, __obf_7fc0261a12598c19, __obf_e21f6160f7309bde, __obf_17b876ebc9671e6f, __obf_85a14dae0b94c317)
	if __obf_d388542c1816a2c3 != nil {
		log.Println("create failed", __obf_d388542c1816a2c3)
		return
	}
	__obf_0908f514e66dbc8f := __obf_735f5b74487e5b11 + ".pem"
	log.Println("write to pem", __obf_0908f514e66dbc8f)
	var __obf_1ba9d8a3516a43e8 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_edd4a003880f9f62}
	__obf_1b629f5ba33ef685 := pem.EncodeToMemory(__obf_1ba9d8a3516a43e8)
	os.WriteFile(__obf_0908f514e66dbc8f, __obf_1b629f5ba33ef685, 0777)

	__obf_25cbaf7597bc9a05 := __obf_735f5b74487e5b11 + ".key"
	__obf_4a34056e5abfcd15 := x509.MarshalPKCS1PrivateKey(__obf_811665aa5ec324fa)
	log.Println("write to key", __obf_25cbaf7597bc9a05)
	os.WriteFile(__obf_25cbaf7597bc9a05, __obf_4a34056e5abfcd15, 0777)
	var __obf_c980f55170ffc00f = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_4a34056e5abfcd15}
	__obf_c2d7dd13537faa2d := pem.EncodeToMemory(__obf_c980f55170ffc00f)
	os.WriteFile(__obf_25cbaf7597bc9a05, __obf_c2d7dd13537faa2d, 0777)
}
