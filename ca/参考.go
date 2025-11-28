package __obf_b6eef43f2d5b2785

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
	__obf_790b34ab736d83bf := &x509.Certificate{
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
	__obf_07b1594622df177b, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("ca", __obf_790b34ab736d83bf, __obf_07b1594622df177b, __obf_790b34ab736d83bf, nil)
	__obf_be5449785053321c := &x509.Certificate{
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
	__obf_1ab5fba3a29354cf := []string{"localhost", "127.0.0.1"}
	for _, __obf_e31ba42bd8787f35 := range __obf_1ab5fba3a29354cf {
		if __obf_f89d1fc1f6ea6294 := net.ParseIP(__obf_e31ba42bd8787f35); __obf_f89d1fc1f6ea6294 != nil {
			__obf_be5449785053321c.IPAddresses = append(__obf_be5449785053321c.IPAddresses, __obf_f89d1fc1f6ea6294)
		} else {
			__obf_be5449785053321c.DNSNames = append(__obf_be5449785053321c.DNSNames, __obf_e31ba42bd8787f35)
		}
	}
	__obf_83fe129fe4adff85, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("server", __obf_be5449785053321c, __obf_83fe129fe4adff85, __obf_790b34ab736d83bf, __obf_07b1594622df177b)
	__obf_57c4e35144085644 := &x509.Certificate{
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
	__obf_828881dc6224de82, _ := rsa.GenerateKey(rand.Reader, 1024)
	CreateCertificateFile("client", __obf_57c4e35144085644, __obf_828881dc6224de82, __obf_790b34ab736d83bf, __obf_07b1594622df177b)
}

func CreateCertificateFile(__obf_e3ff48e9fd78ac32 string, __obf_e474809d004eeffb *x509.Certificate, __obf_cd8dbb8ae8198672 *rsa.PrivateKey, __obf_5bdf75f3d3e18f9b *x509.Certificate, __obf_92bdfa6d1e458444 *rsa.PrivateKey) {
	__obf_19c1bfdaf6027d07 := __obf_cd8dbb8ae8198672
	__obf_f67930f4e0a74fdc := &__obf_19c1bfdaf6027d07.PublicKey
	__obf_f54495018bb184d0 := __obf_19c1bfdaf6027d07
	if __obf_92bdfa6d1e458444 != nil {
		__obf_f54495018bb184d0 = __obf_92bdfa6d1e458444
	}
	__obf_cf8761499bbf6d1d, __obf_415e6cf41a5191a6 := x509.CreateCertificate(rand.Reader, __obf_e474809d004eeffb, __obf_5bdf75f3d3e18f9b, __obf_f67930f4e0a74fdc, __obf_f54495018bb184d0)
	if __obf_415e6cf41a5191a6 != nil {
		log.Println("create failed", __obf_415e6cf41a5191a6)
		return
	}
	__obf_3cb694b9ac182f49 := __obf_e3ff48e9fd78ac32 + ".pem"
	log.Println("write to pem", __obf_3cb694b9ac182f49)
	var __obf_197c4fea5103cc4b = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   __obf_cf8761499bbf6d1d}
	__obf_aa7c8ddce2fc2cb9 := pem.EncodeToMemory(__obf_197c4fea5103cc4b)
	os.WriteFile(__obf_3cb694b9ac182f49, __obf_aa7c8ddce2fc2cb9, 0777)

	__obf_9399e8c09ad1b829 := __obf_e3ff48e9fd78ac32 + ".key"
	__obf_eb2cfa657b2a48cf := x509.MarshalPKCS1PrivateKey(__obf_19c1bfdaf6027d07)
	log.Println("write to key", __obf_9399e8c09ad1b829)
	os.WriteFile(__obf_9399e8c09ad1b829, __obf_eb2cfa657b2a48cf, 0777)
	var __obf_3df8e4d962cfe425 = &pem.Block{Type: "PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   __obf_eb2cfa657b2a48cf}
	__obf_3e35606181d43102 := pem.EncodeToMemory(__obf_3df8e4d962cfe425)
	os.WriteFile(__obf_9399e8c09ad1b829, __obf_3e35606181d43102, 0777)
}
