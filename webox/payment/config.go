package __obf_d2dbb168afcaf75d

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/pkcs12"
)

// Config 支付配置
type Config struct {
	AppID     string            `json:"app_id" yaml:"app_id"`         // 公众号或小程序APPID
	MchID     string            `json:"mch_id" yaml:"mch_id"`         // 商户号
	APIKey    string            `json:"api_key" yaml:"api_key"`       // 商户API密钥
	NotifyURL string            `json:"notify_url" yaml:"notify_url"` // 支付结果通知地址
	SignType  string            `json:"sign_type" yaml:"sign_type"`   // 签名加密类型，默认为MD5，可选HMAC-SHA256
	P12Path   string            `json:"p12_path" yaml:"p12_path"`     // API证书文件路径，如apiclient_cert.p12
	P12Key    string            `json:"p12_key" yaml:"p12_key"`       // API证书密码
	BankMap   map[string]string `json:"bank_map" yaml:"bank_map"`     // 银行对照表
	TLSCert   tls.Certificate   `json:"-" yaml:"-"`
}

func (__obf_b06a54262bb17581 *Config) LoadAPICert() (__obf_f172673a4c06dd64 error) {
	if __obf_b06a54262bb17581.P12Path == "" || __obf_b06a54262bb17581.P12Key == "" {
		return errors.New("p12 path or p12 key is not set")
	}
	__obf_a15f4c5acd9bf03e, __obf_f172673a4c06dd64 := os.ReadFile(__obf_b06a54262bb17581.P12Path)
	if __obf_f172673a4c06dd64 != nil {
		return fmt.Errorf("p12 read failed: %w", __obf_f172673a4c06dd64)
	}
	__obf_8819d7c867a69179, __obf_dbb194879d71a829, __obf_f172673a4c06dd64 := pkcs12.Decode(__obf_a15f4c5acd9bf03e, __obf_b06a54262bb17581.P12Key)
	if __obf_f172673a4c06dd64 != nil {
		return fmt.Errorf("p12 decode failed: %w", __obf_f172673a4c06dd64)
	}
	__obf_b06a54262bb17581.
		TLSCert, __obf_f172673a4c06dd64 = tls.X509KeyPair(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: __obf_dbb194879d71a829.Raw,
			},
		),
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(__obf_8819d7c867a69179.(*rsa.PrivateKey)),
			},
		),
	)
	return
}
