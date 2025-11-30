package main

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
	__obf_152d3d6821401fe0 := &x509.Certificate{
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
	__obf_b35fd3d30bda193e, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e, __obf_152d3d6821401fe0, nil)
	__obf_b04ef76e341283b9 := &x509.Certificate{
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
	__obf_447f061b6b9d3158 := []string{"localhost", "127.0.0.1"}
	for _, __obf_7e404a3606ede3f7 := range __obf_447f061b6b9d3158 {
		if __obf_fee41b11291ca394 := net.ParseIP(__obf_7e404a3606ede3f7); __obf_fee41b11291ca394 != nil {
			__obf_b04ef76e341283b9.
				IPAddresses = append(__obf_b04ef76e341283b9.IPAddresses, __obf_fee41b11291ca394)
		} else {
			__obf_b04ef76e341283b9.
				DNSNames = append(__obf_b04ef76e341283b9.DNSNames, __obf_7e404a3606ede3f7)
		}
	}
	__obf_288f66eab17b85e3, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_b04ef76e341283b9, __obf_288f66eab17b85e3, __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e)
	__obf_50f1524a0f572b1a := &x509.Certificate{
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
	__obf_1d049cf4e1477ec8, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_50f1524a0f572b1a, __obf_1d049cf4e1477ec8, __obf_152d3d6821401fe0, __obf_b35fd3d30bda193e)
}

func CreateCertificateFile(__obf_20afa3a20b4e7df3 string, __obf_b16eba1636bcfd7b *x509.Certificate, __obf_79f143f45af76681 *rsa.PrivateKey, __obf_9ae66557d0a6577d *x509.Certificate, __obf_89bc8a7c7f623f96 *rsa.PrivateKey) {
	__obf_a00b20fba5e0687b := __obf_79f143f45af76681
	__obf_0673cdb4bf5415c0 := &__obf_a00b20fba5e0687b.PublicKey
	__obf_78a4f3baeaa6935c := __obf_a00b20fba5e0687b
	if __obf_89bc8a7c7f623f96 != nil {
		__obf_78a4f3baeaa6935c = __obf_89bc8a7c7f623f96
	}
	__obf_b0b275f1493f4e8d, __obf_50d9d8acee20f6af := x509.CreateCertificate(rand.Reader, __obf_b16eba1636bcfd7b, __obf_9ae66557d0a6577d, __obf_0673cdb4bf5415c0, __obf_78a4f3baeaa6935c)
	if __obf_50d9d8acee20f6af != nil {
		log.Println("create failed", __obf_50d9d8acee20f6af)
		return
	}
	__obf_ca7a372696df7e5a := __obf_20afa3a20b4e7df3 + ".pem"
	log.Println("write to pem", __obf_ca7a372696df7e5a)
	var __obf_aae5fde6704b9a04 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_b0b275f1493f4e8d}
	__obf_adda7084495ff63a := pem.EncodeToMemory(__obf_aae5fde6704b9a04)
	os.WriteFile(__obf_ca7a372696df7e5a, __obf_adda7084495ff63a, 0777)
	__obf_cc13c4f18d5eb72a := __obf_20afa3a20b4e7df3 + ".key"
	__obf_0833ebb9145cf77e := x509.MarshalPKCS1PrivateKey(__obf_a00b20fba5e0687b)
	log.Println("write to key", __obf_cc13c4f18d5eb72a)
	os.WriteFile(__obf_cc13c4f18d5eb72a, __obf_0833ebb9145cf77e, 0777)
	var __obf_0d1e0742fb46681a = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_0833ebb9145cf77e}
	__obf_b584dab635b4689d := pem.EncodeToMemory(__obf_0d1e0742fb46681a)
	os.WriteFile(__obf_cc13c4f18d5eb72a, __obf_b584dab635b4689d, 0777)
}
