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
	__obf_b95177042f904a4c := &x509.Certificate{
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
	__obf_e33fe71340c03c86, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_b95177042f904a4c, __obf_e33fe71340c03c86, __obf_b95177042f904a4c, nil)
	__obf_1c097168084f7804 := &x509.Certificate{
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
	__obf_fc6e2c0ac1d596f1 := []string{"localhost", "127.0.0.1"}
	for _, __obf_2a5b387db77f9ace := range __obf_fc6e2c0ac1d596f1 {
		if __obf_fed4296046358797 := net.ParseIP(__obf_2a5b387db77f9ace); __obf_fed4296046358797 != nil {
			__obf_1c097168084f7804.
				IPAddresses = append(__obf_1c097168084f7804.IPAddresses, __obf_fed4296046358797)
		} else {
			__obf_1c097168084f7804.
				DNSNames = append(__obf_1c097168084f7804.DNSNames, __obf_2a5b387db77f9ace)
		}
	}
	__obf_ce3eb260a95f4fc4, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_1c097168084f7804, __obf_ce3eb260a95f4fc4, __obf_b95177042f904a4c, __obf_e33fe71340c03c86)
	__obf_c17d5b3ab639cab3 := &x509.Certificate{
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
	__obf_d6769bc7d39c18cc, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_c17d5b3ab639cab3, __obf_d6769bc7d39c18cc, __obf_b95177042f904a4c, __obf_e33fe71340c03c86)
}

func CreateCertificateFile(__obf_88d290a111bdfb00 string, __obf_e4997f449948c743 *x509.Certificate, __obf_02b030c8d20c95c6 *rsa.PrivateKey, __obf_c5c0c1e601208157 *x509.Certificate, __obf_bbbb35f28334fa11 *rsa.PrivateKey) {
	__obf_03eb2d2742e0f341 := __obf_02b030c8d20c95c6
	__obf_6a67a0eb73d11448 := &__obf_03eb2d2742e0f341.PublicKey
	__obf_a663dcf260939188 := __obf_03eb2d2742e0f341
	if __obf_bbbb35f28334fa11 != nil {
		__obf_a663dcf260939188 = __obf_bbbb35f28334fa11
	}
	__obf_4257599bfa9403eb, __obf_55b9dfdec0b06a66 := x509.CreateCertificate(rand.Reader, __obf_e4997f449948c743, __obf_c5c0c1e601208157, __obf_6a67a0eb73d11448, __obf_a663dcf260939188)
	if __obf_55b9dfdec0b06a66 != nil {
		log.Println("create failed", __obf_55b9dfdec0b06a66)
		return
	}
	__obf_041c49913fd23979 := __obf_88d290a111bdfb00 + ".pem"
	log.Println("write to pem", __obf_041c49913fd23979)
	var __obf_36885157f764ac30 = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_4257599bfa9403eb}
	__obf_df104a4558dc2ab3 := pem.EncodeToMemory(__obf_36885157f764ac30)
	os.WriteFile(__obf_041c49913fd23979, __obf_df104a4558dc2ab3, 0777)
	__obf_822a47813587116f := __obf_88d290a111bdfb00 + ".key"
	__obf_fb283a9f3eb2f789 := x509.MarshalPKCS1PrivateKey(__obf_03eb2d2742e0f341)
	log.Println("write to key", __obf_822a47813587116f)
	os.WriteFile(__obf_822a47813587116f, __obf_fb283a9f3eb2f789, 0777)
	var __obf_a89b7ed89c41e9ae = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_fb283a9f3eb2f789}
	__obf_3414987a53346d72 := pem.EncodeToMemory(__obf_a89b7ed89c41e9ae)
	os.WriteFile(__obf_822a47813587116f, __obf_3414987a53346d72, 0777)
}
