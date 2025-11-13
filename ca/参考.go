package __obf_3d5f73927e7527a3

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
	__obf_908b47d36bbdd2ed := &x509.Certificate{
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
	__obf_2cd6bb49ee83c5f8, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8, __obf_908b47d36bbdd2ed, nil)
	__obf_0456a1ca2826fea1 := &x509.Certificate{
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
	__obf_1c56ae222a3c2766 := []string{"localhost", "127.0.0.1"}
	for _, __obf_81d55c9a99174773 := range __obf_1c56ae222a3c2766 {
		if __obf_b98bdd9cf219874b := net.ParseIP(__obf_81d55c9a99174773); __obf_b98bdd9cf219874b != nil {
			__obf_0456a1ca2826fea1.IPAddresses = append(__obf_0456a1ca2826fea1.IPAddresses, __obf_b98bdd9cf219874b)
		} else {
			__obf_0456a1ca2826fea1.DNSNames = append(__obf_0456a1ca2826fea1.DNSNames, __obf_81d55c9a99174773)
		}
	}
	__obf_8f52ed7b17bc6019, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_0456a1ca2826fea1, __obf_8f52ed7b17bc6019, __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8)
	__obf_0f9b320858d91156 := &x509.Certificate{
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
	__obf_6755722962ef4f64, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_0f9b320858d91156, __obf_6755722962ef4f64, __obf_908b47d36bbdd2ed, __obf_2cd6bb49ee83c5f8)
}

func CreateCertificateFile(__obf_68a329dd0bc1d1a3 string, __obf_007dd92ffb4e580f *x509.Certificate, __obf_d22a56f6883a4549 *rsa.PrivateKey, __obf_8ed23350cdb49c87 *x509.Certificate, __obf_0e8e0972e62f0547 *rsa.PrivateKey) {
	__obf_b78fb0ff4f30c05e := __obf_d22a56f6883a4549
	__obf_5ce2e6a786e22558 := &__obf_b78fb0ff4f30c05e.PublicKey
	__obf_bdf472f121855193 := __obf_b78fb0ff4f30c05e
	if __obf_0e8e0972e62f0547 != nil {
		__obf_bdf472f121855193 = __obf_0e8e0972e62f0547
	}
	__obf_6fa9fdcf13d2e16d, __obf_2eef99cccb5deacc := x509.CreateCertificate(rand.Reader, __obf_007dd92ffb4e580f, __obf_8ed23350cdb49c87, __obf_5ce2e6a786e22558, __obf_bdf472f121855193)
	if __obf_2eef99cccb5deacc != nil {
		log.Println("create failed", __obf_2eef99cccb5deacc)
		return
	}
	__obf_a0e2dc722dee24e4 := __obf_68a329dd0bc1d1a3 + ".pem"
	log.Println("write to pem", __obf_a0e2dc722dee24e4)
	var __obf_89a03731c728864e = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_6fa9fdcf13d2e16d}
	__obf_4c473f20858b8d85 := pem.EncodeToMemory(__obf_89a03731c728864e)
	os.WriteFile(__obf_a0e2dc722dee24e4, __obf_4c473f20858b8d85, 0777)

	__obf_32e5cf4a26dd98a3 := __obf_68a329dd0bc1d1a3 + ".key"
	__obf_feb7f769524150f3 := x509.MarshalPKCS1PrivateKey(__obf_b78fb0ff4f30c05e)
	log.Println("write to key", __obf_32e5cf4a26dd98a3)
	os.WriteFile(__obf_32e5cf4a26dd98a3, __obf_feb7f769524150f3, 0777)
	var __obf_1424e2db13d1bb7d = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_feb7f769524150f3}
	__obf_7ad817bc29b678fa := pem.EncodeToMemory(__obf_1424e2db13d1bb7d)
	os.WriteFile(__obf_32e5cf4a26dd98a3, __obf_7ad817bc29b678fa, 0777)
}
