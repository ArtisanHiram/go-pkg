package __obf_b6eef43f2d5b2785

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func __obf_6b6b330665324b95(__obf_12ddfaf8835b71e4 string, __obf_64ec15fd4dadc9fe int) (__obf_f08f0e66d42958a5 time.Time, __obf_2bbf5da0aeb44997 time.Time) {
	var __obf_415e6cf41a5191a6 error
	if __obf_12ddfaf8835b71e4 == "" {
		__obf_f08f0e66d42958a5 = time.Now()
	} else {
		if __obf_f08f0e66d42958a5, __obf_415e6cf41a5191a6 = time.Parse("2006-01-02", __obf_12ddfaf8835b71e4); __obf_415e6cf41a5191a6 != nil {
			log.Fatalln("time parse failed.")
		} else {
			__obf_2bbf5da0aeb44997 = __obf_f08f0e66d42958a5.Add(time.Duration(__obf_64ec15fd4dadc9fe*24) * time.Hour)
		}
	}
	return
}

func __obf_0b17e9abdbbf8505() any {
	var __obf_415e6cf41a5191a6 error
	var __obf_19c1bfdaf6027d07 any
	switch *__obf_bba0c595ec2b1e3d {
	case "":
		if *__obf_8188aca91d5a646c {
			_, __obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = ed25519.GenerateKey(rand.Reader)
		} else {
			__obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = rsa.GenerateKey(rand.Reader, *__obf_6cac191609a547f0)
		}
	case "P224":
		__obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		__obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		__obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		__obf_19c1bfdaf6027d07, __obf_415e6cf41a5191a6 = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", *__obf_bba0c595ec2b1e3d)
	}

	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalf("Failed to generate private key: %v", __obf_415e6cf41a5191a6)
	}

	return __obf_19c1bfdaf6027d07
}

func __obf_93d4391a9e1c34c7(__obf_2bb7aef23d1b6e1f, __obf_9b9102c5830c64d9 string, __obf_0661de7fa0c8957a []byte, __obf_42b9b50e45afbe86 os.FileMode) error {

	if _, __obf_415e6cf41a5191a6 := os.Stat(__obf_2bb7aef23d1b6e1f); __obf_415e6cf41a5191a6 == nil {
		if __obf_415e6cf41a5191a6 = os.WriteFile(fmt.Sprintf("%s%s%s", __obf_2bb7aef23d1b6e1f, string(os.PathSeparator), __obf_9b9102c5830c64d9), __obf_0661de7fa0c8957a, __obf_42b9b50e45afbe86); __obf_415e6cf41a5191a6 != nil {
			return errors.Wrap(__obf_415e6cf41a5191a6, "Write file failed.")
		}
	} else if os.IsNotExist(__obf_415e6cf41a5191a6) {
		return errors.Wrap(__obf_415e6cf41a5191a6, "Directory not exist.")
	} else {
		return __obf_415e6cf41a5191a6
	}
	return nil
}

func __obf_23d433731568035e(__obf_e3ff48e9fd78ac32 string, __obf_e474809d004eeffb *x509.Certificate, __obf_cd8dbb8ae8198672 any, __obf_5bdf75f3d3e18f9b *x509.Certificate, __obf_92bdfa6d1e458444 any) {

	var __obf_f67930f4e0a74fdc any
	switch __obf_a0496fd33c0dd52d := __obf_cd8dbb8ae8198672.(type) {
	case *rsa.PrivateKey:
		__obf_f67930f4e0a74fdc = &__obf_a0496fd33c0dd52d.PublicKey
	case *ecdsa.PrivateKey:
		__obf_f67930f4e0a74fdc = &__obf_a0496fd33c0dd52d.PublicKey
	case ed25519.PrivateKey:
		__obf_f67930f4e0a74fdc = __obf_a0496fd33c0dd52d.Public().(ed25519.PublicKey)
	default:
		__obf_f67930f4e0a74fdc = nil
	}

	var __obf_f54495018bb184d0 any
	if __obf_92bdfa6d1e458444 != nil {
		__obf_f54495018bb184d0 = __obf_92bdfa6d1e458444
	} else {
		__obf_f54495018bb184d0 = __obf_cd8dbb8ae8198672
	}
	__obf_cf8761499bbf6d1d, __obf_415e6cf41a5191a6 := x509.CreateCertificate(rand.Reader, __obf_e474809d004eeffb, __obf_5bdf75f3d3e18f9b, __obf_f67930f4e0a74fdc, __obf_f54495018bb184d0)
	if __obf_415e6cf41a5191a6 != nil {
		log.Println("create failed", __obf_415e6cf41a5191a6)
		return
	}

	__obf_aa7c8ddce2fc2cb9 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: __obf_cf8761499bbf6d1d})
	// fmt.Printf("%s private key： %s \n", name, ca_b)

	if __obf_415e6cf41a5191a6 = __obf_93d4391a9e1c34c7(*__obf_7662752a24f6c29a, __obf_e3ff48e9fd78ac32+".pem", __obf_aa7c8ddce2fc2cb9, 0777); __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}

	__obf_eb2cfa657b2a48cf, __obf_415e6cf41a5191a6 := x509.MarshalPKCS8PrivateKey(__obf_cd8dbb8ae8198672)
	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalln("marshal PKCS8 private key failed")
	}
	// os.WriteFile(name+".key", priv_b, 0777)

	__obf_3e35606181d43102 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: __obf_eb2cfa657b2a48cf})

	if __obf_415e6cf41a5191a6 = __obf_93d4391a9e1c34c7(*__obf_7662752a24f6c29a, __obf_e3ff48e9fd78ac32+".key", __obf_3e35606181d43102, 0777); __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}
}

// 公钥证书解析
func __obf_0bb2940fc6306c63() (*x509.Certificate, error) {
	//读取公钥证书并解码
	if __obf_6eb8407002b53b06, __obf_415e6cf41a5191a6 := os.ReadFile(fmt.Sprintf("%s%sroot.pem", *__obf_7662752a24f6c29a, string(os.PathSeparator))); __obf_415e6cf41a5191a6 != nil {
		return nil, errors.Wrap(__obf_415e6cf41a5191a6, "root.pem: Failed to read file.")
	} else {
		__obf_cb7a6ace424a49dd, __obf_706013395daf68fb := pem.Decode(__obf_6eb8407002b53b06)
		if __obf_cb7a6ace424a49dd == nil {
			return nil, errors.Errorf("root.pem: Cert block is nil, rest block is %s.", __obf_706013395daf68fb)
		}

		return x509.ParseCertificate(__obf_cb7a6ace424a49dd.Bytes)
	}
}

// 私钥解析
func __obf_069470c08ab56ed7() (__obf_cd8dbb8ae8198672 any, __obf_415e6cf41a5191a6 error) {
	if __obf_6eb8407002b53b06, __obf_415e6cf41a5191a6 := os.ReadFile(fmt.Sprintf("%s%sroot.key", *__obf_7662752a24f6c29a, string(os.PathSeparator))); __obf_415e6cf41a5191a6 != nil {
		return nil, errors.Wrap(__obf_415e6cf41a5191a6, "root.key: Failed to read file.")
	} else {
		__obf_cb7a6ace424a49dd, __obf_706013395daf68fb := pem.Decode(__obf_6eb8407002b53b06)
		if __obf_cb7a6ace424a49dd == nil {
			return nil, errors.Errorf("root.key: Cert block is nil, rest block is %s.", __obf_706013395daf68fb)
		}

		return x509.ParsePKCS8PrivateKey(__obf_cb7a6ace424a49dd.Bytes)
	}
}

func __obf_a3a0906b6aba92b5() *big.Int {
	__obf_9a1d4826b73b04d7 := new(big.Int).Lsh(big.NewInt(1), 128)
	if __obf_b5984421c5981de5, __obf_415e6cf41a5191a6 := rand.Int(rand.Reader, __obf_9a1d4826b73b04d7); __obf_415e6cf41a5191a6 != nil {
		log.Fatalln("generate serial number failed")
		return nil
	} else {
		return __obf_b5984421c5981de5
	}
}

func GenRootCA() {

	__obf_6aef97d325806a61, __obf_fe9595acdda6f828 := __obf_6b6b330665324b95(*__obf_f3d20ae37848f3fc, *__obf_ece9c7189be6d320)
	__obf_790b34ab736d83bf := &x509.Certificate{
		SerialNumber: __obf_a3a0906b6aba92b5(),
		Subject: pkix.Name{
			// Country:            []string{"China"},
			Organization: []string{*__obf_370baf1383967997},
			// OrganizationalUnit: []string{"Shit company Unit"},
		},
		NotBefore:             __obf_6aef97d325806a61,
		NotAfter:              __obf_fe9595acdda6f828,
		IsCA:                  true,
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	__obf_07b1594622df177b := __obf_0b17e9abdbbf8505()

	__obf_23d433731568035e("root", __obf_790b34ab736d83bf, __obf_07b1594622df177b, __obf_790b34ab736d83bf, nil)

}

func GenServerCA() {

	__obf_6aef97d325806a61, __obf_fe9595acdda6f828 := __obf_6b6b330665324b95(*__obf_d8c4f6c2d7602f25, *__obf_b42de427a3167a74)
	__obf_be5449785053321c := &x509.Certificate{
		SerialNumber: __obf_a3a0906b6aba92b5(),
		Subject: pkix.Name{
			Organization: []string{*__obf_5465625a1ceeaf75},
		},
		NotBefore: __obf_6aef97d325806a61,
		NotAfter:  __obf_fe9595acdda6f828,
		// SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	__obf_1ab5fba3a29354cf := strings.Split(*__obf_18118f4267561011, ",")

	for _, __obf_e31ba42bd8787f35 := range __obf_1ab5fba3a29354cf {
		if __obf_f89d1fc1f6ea6294 := net.ParseIP(__obf_e31ba42bd8787f35); __obf_f89d1fc1f6ea6294 != nil {
			__obf_be5449785053321c.IPAddresses = append(__obf_be5449785053321c.IPAddresses, __obf_f89d1fc1f6ea6294)
		} else {
			__obf_be5449785053321c.DNSNames = append(__obf_be5449785053321c.DNSNames, __obf_e31ba42bd8787f35)
		}
	}
	__obf_83fe129fe4adff85 := __obf_0b17e9abdbbf8505()

	var __obf_790b34ab736d83bf *x509.Certificate
	var __obf_415e6cf41a5191a6 error
	var __obf_07b1594622df177b any

	__obf_790b34ab736d83bf, __obf_415e6cf41a5191a6 = __obf_0bb2940fc6306c63()
	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}

	__obf_07b1594622df177b, __obf_415e6cf41a5191a6 = __obf_069470c08ab56ed7()
	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}

	__obf_23d433731568035e("server", __obf_be5449785053321c, __obf_83fe129fe4adff85, __obf_790b34ab736d83bf, __obf_07b1594622df177b)
}

func GenClientCA() {

	__obf_6aef97d325806a61, __obf_fe9595acdda6f828 := __obf_6b6b330665324b95(*__obf_f625282011c32525, *__obf_81b901915e137680)
	__obf_57c4e35144085644 := &x509.Certificate{
		SerialNumber: __obf_a3a0906b6aba92b5(),
		Subject: pkix.Name{
			Organization: []string{*__obf_4d65ba9ada5d6594},
		},
		NotBefore: __obf_6aef97d325806a61,
		NotAfter:  __obf_fe9595acdda6f828,
		// SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	__obf_828881dc6224de82 := __obf_0b17e9abdbbf8505()

	var __obf_790b34ab736d83bf *x509.Certificate
	var __obf_415e6cf41a5191a6 error
	var __obf_07b1594622df177b any

	__obf_790b34ab736d83bf, __obf_415e6cf41a5191a6 = __obf_0bb2940fc6306c63()
	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}

	__obf_07b1594622df177b, __obf_415e6cf41a5191a6 = __obf_069470c08ab56ed7()
	if __obf_415e6cf41a5191a6 != nil {
		log.Fatalln(__obf_415e6cf41a5191a6)
	}

	__obf_23d433731568035e("client", __obf_57c4e35144085644, __obf_828881dc6224de82, __obf_790b34ab736d83bf, __obf_07b1594622df177b)

}
