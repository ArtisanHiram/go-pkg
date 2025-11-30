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
	__obf_ec85745def2d4ae6 := &x509.Certificate{
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
	__obf_91759aa608c2ec39, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39, __obf_ec85745def2d4ae6, nil)
	__obf_e64c40cfb6bdad21 := &x509.Certificate{
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
	__obf_ad24a3c0621435b3 := []string{"localhost", "127.0.0.1"}
	for _, __obf_07c3f9870a15fb0b := range __obf_ad24a3c0621435b3 {
		if __obf_35fe568983671a61 := net.ParseIP(__obf_07c3f9870a15fb0b); __obf_35fe568983671a61 != nil {
			__obf_e64c40cfb6bdad21.
				IPAddresses = append(__obf_e64c40cfb6bdad21.IPAddresses, __obf_35fe568983671a61)
		} else {
			__obf_e64c40cfb6bdad21.
				DNSNames = append(__obf_e64c40cfb6bdad21.DNSNames, __obf_07c3f9870a15fb0b)
		}
	}
	__obf_34e7baa7d7852431, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_e64c40cfb6bdad21, __obf_34e7baa7d7852431, __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39)
	__obf_312e3abafc6822e1 := &x509.Certificate{
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
	__obf_6f6fbd67ae64bdcd, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_312e3abafc6822e1, __obf_6f6fbd67ae64bdcd, __obf_ec85745def2d4ae6, __obf_91759aa608c2ec39)
}

func CreateCertificateFile(__obf_09d528286129fda4 string, __obf_1118abe6002652cf *x509.Certificate, __obf_d073cba656c62c18 *rsa.PrivateKey, __obf_d77ba6307d0c3956 *x509.Certificate, __obf_aff132c3a457e217 *rsa.PrivateKey) {
	__obf_fe5ca72bb97c16b8 := __obf_d073cba656c62c18
	__obf_fb675a630c70f7b6 := &__obf_fe5ca72bb97c16b8.PublicKey
	__obf_291b188c843da91c := __obf_fe5ca72bb97c16b8
	if __obf_aff132c3a457e217 != nil {
		__obf_291b188c843da91c = __obf_aff132c3a457e217
	}
	__obf_f52e4c0d1cc74eae, __obf_dd7d7edd66f39226 := x509.CreateCertificate(rand.Reader, __obf_1118abe6002652cf, __obf_d77ba6307d0c3956, __obf_fb675a630c70f7b6, __obf_291b188c843da91c)
	if __obf_dd7d7edd66f39226 != nil {
		log.Println("create failed", __obf_dd7d7edd66f39226)
		return
	}
	__obf_e301f165b8df7abf := __obf_09d528286129fda4 + ".pem"
	log.Println("write to pem", __obf_e301f165b8df7abf)
	var __obf_5a22a508efb90dc7 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_f52e4c0d1cc74eae}
	__obf_296c04406cfe0986 := pem.EncodeToMemory(__obf_5a22a508efb90dc7)
	os.WriteFile(__obf_e301f165b8df7abf, __obf_296c04406cfe0986, 0777)
	__obf_9ee54992a4bf0357 := __obf_09d528286129fda4 + ".key"
	__obf_df3032d429e4c468 := x509.MarshalPKCS1PrivateKey(__obf_fe5ca72bb97c16b8)
	log.Println("write to key", __obf_9ee54992a4bf0357)
	os.WriteFile(__obf_9ee54992a4bf0357, __obf_df3032d429e4c468, 0777)
	var __obf_459ea24ea7c288c4 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_df3032d429e4c468}
	__obf_a47b9e6b4f4d8ea1 := pem.EncodeToMemory(__obf_459ea24ea7c288c4)
	os.WriteFile(__obf_9ee54992a4bf0357, __obf_a47b9e6b4f4d8ea1, 0777)
}
