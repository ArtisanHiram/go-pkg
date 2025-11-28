package __obf_be728ac087d82248

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
	__obf_b3faede65159f8cc := &x509.Certificate{
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
	__obf_799b6d3ad392bc41, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41, __obf_b3faede65159f8cc, nil)
	__obf_ec14c26d96fa147b := &x509.Certificate{
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
	__obf_f06ad913bede6373 := []string{"localhost", "127.0.0.1"}
	for _, __obf_562235314fa4d813 := range __obf_f06ad913bede6373 {
		if __obf_efdf1783571090b7 := net.ParseIP(__obf_562235314fa4d813); __obf_efdf1783571090b7 != nil {
			__obf_ec14c26d96fa147b.IPAddresses = append(__obf_ec14c26d96fa147b.IPAddresses, __obf_efdf1783571090b7)
		} else {
			__obf_ec14c26d96fa147b.DNSNames = append(__obf_ec14c26d96fa147b.DNSNames, __obf_562235314fa4d813)
		}
	}
	__obf_79e7664f49117a3c, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_ec14c26d96fa147b, __obf_79e7664f49117a3c, __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41)
	__obf_c048349241bd3fba := &x509.Certificate{
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
	__obf_f75e5e99dd49aa4b, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_c048349241bd3fba, __obf_f75e5e99dd49aa4b, __obf_b3faede65159f8cc, __obf_799b6d3ad392bc41)
}

func CreateCertificateFile(__obf_a6be010d3ba4ee10 string, __obf_133e099e3d467a1b *x509.Certificate, __obf_921c11678a889b4e *rsa.PrivateKey, __obf_d8789547d12e19b7 *x509.Certificate, __obf_5721445e51f1c16a *rsa.PrivateKey) {
	__obf_9b7b8b4ad20133b0 := __obf_921c11678a889b4e
	__obf_163b33d215165a24 := &__obf_9b7b8b4ad20133b0.PublicKey
	__obf_4841eb4e2e790500 := __obf_9b7b8b4ad20133b0
	if __obf_5721445e51f1c16a != nil {
		__obf_4841eb4e2e790500 = __obf_5721445e51f1c16a
	}
	__obf_682dc1578dc812d8, __obf_38d4e8c9703c399c := x509.CreateCertificate(rand.Reader, __obf_133e099e3d467a1b, __obf_d8789547d12e19b7, __obf_163b33d215165a24, __obf_4841eb4e2e790500)
	if __obf_38d4e8c9703c399c != nil {
		log.Println("create failed", __obf_38d4e8c9703c399c)
		return
	}
	__obf_8e65360d6881748a := __obf_a6be010d3ba4ee10 + ".pem"
	log.Println("write to pem", __obf_8e65360d6881748a)
	var __obf_4347f654c0434c49 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_682dc1578dc812d8}
	__obf_72bbb56c78c60fee := pem.EncodeToMemory(__obf_4347f654c0434c49)
	os.WriteFile(__obf_8e65360d6881748a, __obf_72bbb56c78c60fee, 0777)

	__obf_85840e9b63e465cc := __obf_a6be010d3ba4ee10 + ".key"
	__obf_4cfa010f200444c3 := x509.MarshalPKCS1PrivateKey(__obf_9b7b8b4ad20133b0)
	log.Println("write to key", __obf_85840e9b63e465cc)
	os.WriteFile(__obf_85840e9b63e465cc, __obf_4cfa010f200444c3, 0777)
	var __obf_5403d29b0ec254d6 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_4cfa010f200444c3}
	__obf_f075f6e8e46f2c96 := pem.EncodeToMemory(__obf_5403d29b0ec254d6)
	os.WriteFile(__obf_85840e9b63e465cc, __obf_f075f6e8e46f2c96, 0777)
}
