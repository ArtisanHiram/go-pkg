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
	__obf_bd646ef02df2a803 := &x509.Certificate{
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
	__obf_031842f1c2ed65a3, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3, __obf_bd646ef02df2a803, nil)
	__obf_441b7d9118c61425 := &x509.Certificate{
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
	__obf_0f57c4cf881d27c7 := []string{"localhost", "127.0.0.1"}
	for _, __obf_64a53c251efc3cd6 := range __obf_0f57c4cf881d27c7 {
		if __obf_663e92eef5c260d0 := net.ParseIP(__obf_64a53c251efc3cd6); __obf_663e92eef5c260d0 != nil {
			__obf_441b7d9118c61425.
				IPAddresses = append(__obf_441b7d9118c61425.IPAddresses, __obf_663e92eef5c260d0)
		} else {
			__obf_441b7d9118c61425.
				DNSNames = append(__obf_441b7d9118c61425.DNSNames, __obf_64a53c251efc3cd6)
		}
	}
	__obf_f02f58ad7a9d0c1f, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_441b7d9118c61425, __obf_f02f58ad7a9d0c1f, __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3)
	__obf_15eb203631991e8f := &x509.Certificate{
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
	__obf_a1788deecf5df7ee, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_15eb203631991e8f, __obf_a1788deecf5df7ee, __obf_bd646ef02df2a803, __obf_031842f1c2ed65a3)
}

func CreateCertificateFile(__obf_94fa2b121e9208a2 string, __obf_89b937c04944cb2b *x509.Certificate, __obf_28a277ccea818a74 *rsa.PrivateKey, __obf_19492af16f0f666d *x509.Certificate, __obf_a74fc724d9a30e04 *rsa.PrivateKey) {
	__obf_61e0a98f3b5bf932 := __obf_28a277ccea818a74
	__obf_4e690447a56017a0 := &__obf_61e0a98f3b5bf932.PublicKey
	__obf_f42003db463ece31 := __obf_61e0a98f3b5bf932
	if __obf_a74fc724d9a30e04 != nil {
		__obf_f42003db463ece31 = __obf_a74fc724d9a30e04
	}
	__obf_cc085b2bdcc9bc80, __obf_8b98d286773a1eea := x509.CreateCertificate(rand.Reader, __obf_89b937c04944cb2b, __obf_19492af16f0f666d, __obf_4e690447a56017a0, __obf_f42003db463ece31)
	if __obf_8b98d286773a1eea != nil {
		log.Println("create failed", __obf_8b98d286773a1eea)
		return
	}
	__obf_feffd1b0b128641b := __obf_94fa2b121e9208a2 + ".pem"
	log.Println("write to pem", __obf_feffd1b0b128641b)
	var __obf_5a591975dd160c43 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_cc085b2bdcc9bc80}
	__obf_f865972c5b507991 := pem.EncodeToMemory(__obf_5a591975dd160c43)
	os.WriteFile(__obf_feffd1b0b128641b, __obf_f865972c5b507991, 0777)
	__obf_e12ec171d75a87f7 := __obf_94fa2b121e9208a2 + ".key"
	__obf_ae4fc2cbb7410041 := x509.MarshalPKCS1PrivateKey(__obf_61e0a98f3b5bf932)
	log.Println("write to key", __obf_e12ec171d75a87f7)
	os.WriteFile(__obf_e12ec171d75a87f7, __obf_ae4fc2cbb7410041, 0777)
	var __obf_6779f1ca0a40a80b = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_ae4fc2cbb7410041}
	__obf_cf30594748a1412b := pem.EncodeToMemory(__obf_6779f1ca0a40a80b)
	os.WriteFile(__obf_e12ec171d75a87f7, __obf_cf30594748a1412b, 0777)
}
