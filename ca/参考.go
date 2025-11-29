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
	__obf_cccdf1a27493dddc := &x509.Certificate{
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
	__obf_1d68b8dfdd86e587, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587, __obf_cccdf1a27493dddc, nil)
	__obf_28da83af6f8edbe6 := &x509.Certificate{
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
	__obf_1b90c90256dcecac := []string{"localhost", "127.0.0.1"}
	for _, __obf_56731192f777ebe3 := range __obf_1b90c90256dcecac {
		if __obf_34b6469a7cf51aaa := net.ParseIP(__obf_56731192f777ebe3); __obf_34b6469a7cf51aaa != nil {
			__obf_28da83af6f8edbe6.
				IPAddresses = append(__obf_28da83af6f8edbe6.IPAddresses, __obf_34b6469a7cf51aaa)
		} else {
			__obf_28da83af6f8edbe6.
				DNSNames = append(__obf_28da83af6f8edbe6.DNSNames, __obf_56731192f777ebe3)
		}
	}
	__obf_0dd6962322487601, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_28da83af6f8edbe6, __obf_0dd6962322487601, __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587)
	__obf_9c2efd023b154ad3 := &x509.Certificate{
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
	__obf_3a8c2f70431f5f7c, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_9c2efd023b154ad3, __obf_3a8c2f70431f5f7c, __obf_cccdf1a27493dddc, __obf_1d68b8dfdd86e587)
}

func CreateCertificateFile(__obf_1ed29960d2e595e7 string, __obf_c7be66b042b1a623 *x509.Certificate, __obf_2f56a8cdb16bd159 *rsa.PrivateKey, __obf_f5dae7569058304a *x509.Certificate, __obf_3b0f2a731719128f *rsa.PrivateKey) {
	__obf_3d6e42833081d2b9 := __obf_2f56a8cdb16bd159
	__obf_bfa26ca12c831807 := &__obf_3d6e42833081d2b9.PublicKey
	__obf_7515f5b04ab4fdd3 := __obf_3d6e42833081d2b9
	if __obf_3b0f2a731719128f != nil {
		__obf_7515f5b04ab4fdd3 = __obf_3b0f2a731719128f
	}
	__obf_ccfcf99a3da52331, __obf_2ec52a9654271348 := x509.CreateCertificate(rand.Reader, __obf_c7be66b042b1a623, __obf_f5dae7569058304a, __obf_bfa26ca12c831807, __obf_7515f5b04ab4fdd3)
	if __obf_2ec52a9654271348 != nil {
		log.Println("create failed", __obf_2ec52a9654271348)
		return
	}
	__obf_b6d4a36e61100fd8 := __obf_1ed29960d2e595e7 + ".pem"
	log.Println("write to pem", __obf_b6d4a36e61100fd8)
	var __obf_d1ea3d4d627fe20b = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_ccfcf99a3da52331}
	__obf_a8076ac87f764dce := pem.EncodeToMemory(__obf_d1ea3d4d627fe20b)
	os.WriteFile(__obf_b6d4a36e61100fd8, __obf_a8076ac87f764dce, 0777)
	__obf_7da87d2922d82302 := __obf_1ed29960d2e595e7 + ".key"
	__obf_c8c65e857103e73b := x509.MarshalPKCS1PrivateKey(__obf_3d6e42833081d2b9)
	log.Println("write to key", __obf_7da87d2922d82302)
	os.WriteFile(__obf_7da87d2922d82302, __obf_c8c65e857103e73b, 0777)
	var __obf_b651f30dc04c1b40 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_c8c65e857103e73b}
	__obf_ad33ff08cd765f4b := pem.EncodeToMemory(__obf_b651f30dc04c1b40)
	os.WriteFile(__obf_7da87d2922d82302, __obf_ad33ff08cd765f4b, 0777)
}
