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
	__obf_7f71517510974c14 := &x509.Certificate{
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
	__obf_fb5bb54bb36cf6fc, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc, __obf_7f71517510974c14, nil)
	__obf_c0692453855fca4d := &x509.Certificate{
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
	__obf_1ddf20596d21c921 := []string{"localhost", "127.0.0.1"}
	for _, __obf_2e59dd0944bc9f24 := range __obf_1ddf20596d21c921 {
		if __obf_f5a5e5d5216c8535 := net.ParseIP(__obf_2e59dd0944bc9f24); __obf_f5a5e5d5216c8535 != nil {
			__obf_c0692453855fca4d.
				IPAddresses = append(__obf_c0692453855fca4d.IPAddresses, __obf_f5a5e5d5216c8535)
		} else {
			__obf_c0692453855fca4d.
				DNSNames = append(__obf_c0692453855fca4d.DNSNames, __obf_2e59dd0944bc9f24)
		}
	}
	__obf_84eaf4818e78c714, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_c0692453855fca4d, __obf_84eaf4818e78c714, __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc)
	__obf_7fe838c84a1220fa := &x509.Certificate{
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
	__obf_b63244d999491ec0, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_7fe838c84a1220fa, __obf_b63244d999491ec0, __obf_7f71517510974c14, __obf_fb5bb54bb36cf6fc)
}

func CreateCertificateFile(__obf_d6849ab3a95e4298 string, __obf_ec1cc1221f3f9dca *x509.Certificate, __obf_c1c829b3ef681f91 *rsa.PrivateKey, __obf_db9fe55105775ed2 *x509.Certificate, __obf_f758ce7f9074c0a3 *rsa.PrivateKey) {
	__obf_173167651b5d15e3 := __obf_c1c829b3ef681f91
	__obf_aca6cb2251255f64 := &__obf_173167651b5d15e3.PublicKey
	__obf_deacdf85c4285de9 := __obf_173167651b5d15e3
	if __obf_f758ce7f9074c0a3 != nil {
		__obf_deacdf85c4285de9 = __obf_f758ce7f9074c0a3
	}
	__obf_18d5d6f844489868, __obf_2107c594e6166770 := x509.CreateCertificate(rand.Reader, __obf_ec1cc1221f3f9dca, __obf_db9fe55105775ed2, __obf_aca6cb2251255f64, __obf_deacdf85c4285de9)
	if __obf_2107c594e6166770 != nil {
		log.Println("create failed", __obf_2107c594e6166770)
		return
	}
	__obf_67a35dc53ffa0d06 := __obf_d6849ab3a95e4298 + ".pem"
	log.Println("write to pem", __obf_67a35dc53ffa0d06)
	var __obf_e68bc445ae69964a = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_18d5d6f844489868}
	__obf_c3dd8c46e8dba5d5 := pem.EncodeToMemory(__obf_e68bc445ae69964a)
	os.WriteFile(__obf_67a35dc53ffa0d06, __obf_c3dd8c46e8dba5d5, 0777)
	__obf_5aba11a1837c53bc := __obf_d6849ab3a95e4298 + ".key"
	__obf_8941cf408812dcc7 := x509.MarshalPKCS1PrivateKey(__obf_173167651b5d15e3)
	log.Println("write to key", __obf_5aba11a1837c53bc)
	os.WriteFile(__obf_5aba11a1837c53bc, __obf_8941cf408812dcc7, 0777)
	var __obf_01a543acbe158da5 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_8941cf408812dcc7}
	__obf_214844b19375b2d9 := pem.EncodeToMemory(__obf_01a543acbe158da5)
	os.WriteFile(__obf_5aba11a1837c53bc, __obf_214844b19375b2d9, 0777)
}
