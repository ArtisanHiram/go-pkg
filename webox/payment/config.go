package __obf_b208cbe38c7e3325

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

func (__obf_114d56f910b3b284 *Config) LoadAPICert() (__obf_4a2f7560afdd36ef error) {
	if __obf_114d56f910b3b284.P12Path == "" || __obf_114d56f910b3b284.P12Key == "" {
		return errors.New("p12 path or p12 key is not set")
	}
	__obf_952a9f7d696cc7e9, __obf_4a2f7560afdd36ef := os.ReadFile(__obf_114d56f910b3b284.P12Path)
	if __obf_4a2f7560afdd36ef != nil {
		return fmt.Errorf("p12 read failed: %w", __obf_4a2f7560afdd36ef)
	}
	__obf_b3258e4e01c3274c, __obf_60b3d3dc2577a1dc, __obf_4a2f7560afdd36ef := pkcs12.Decode(__obf_952a9f7d696cc7e9, __obf_114d56f910b3b284.P12Key)
	if __obf_4a2f7560afdd36ef != nil {
		return fmt.Errorf("p12 decode failed: %w", __obf_4a2f7560afdd36ef)
	}
	__obf_114d56f910b3b284.
		TLSCert, __obf_4a2f7560afdd36ef = tls.X509KeyPair(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: __obf_60b3d3dc2577a1dc.Raw,
			},
		),
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(__obf_b3258e4e01c3274c.(*rsa.PrivateKey)),
			},
		),
	)
	return
}
