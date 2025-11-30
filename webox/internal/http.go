package __obf_e79e1b0b12d02d0b

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client 封装了通用的HTTP请求逻辑
type HttpClient struct {
	__obf_4650a5d47e8654af *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{__obf_4650a5d47e8654af: &http.Client{
		Timeout: 15 * time.Second, // 设置一个合理的超时时间
	},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_0997d36f63f3eba3 tls.Certificate) (*HttpClient, error) {
	__obf_5e964a51416605a8 := // 加载客户端证书和私钥
		// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		// if err != nil {
		// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
		// }

		// 构造TLS配置
		&tls.Config{
			Certificates: []tls.Certificate{__obf_0997d36f63f3eba3},
			// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
			// RootCAs:            nil, // 如果有自定义CA，可以设置
		}
	__obf_222f4171e90b8d52 := // 创建带证书的HTTP Transport
		&http.Transport{
			TLSClientConfig:       __obf_5e964a51416605a8,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

	return &HttpClient{__obf_4650a5d47e8654af: &http.Client{
		Transport: __obf_222f4171e90b8d52,
		Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
	},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_e41c990542ded541 *HttpClient) Get(__obf_e30efbaedbdb33bc string, __obf_8b73ae11e1e90538 map[string]string, __obf_92da1cf9755f8a83 any) error {
	__obf_4f33d00a016cd8c0, __obf_d6bcde9761346c61 := http.NewRequest(http.MethodGet, __obf_e30efbaedbdb33bc, nil)
	if __obf_d6bcde9761346c61 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_c4369973a8ed61a8 := __obf_4f33d00a016cd8c0.URL.Query()
	for __obf_8004953ba26bf329, __obf_2ad96ff371650fa9 := range __obf_8b73ae11e1e90538 {
		__obf_c4369973a8ed61a8.
			Add(__obf_8004953ba26bf329, __obf_2ad96ff371650fa9)
	}
	__obf_4f33d00a016cd8c0.
		URL.RawQuery = __obf_c4369973a8ed61a8.Encode()
	__obf_f86e7ccc55a7708f, __obf_d6bcde9761346c61 := __obf_e41c990542ded541.__obf_4650a5d47e8654af.Do(__obf_4f33d00a016cd8c0)
	if __obf_d6bcde9761346c61 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_d6bcde9761346c61)
	}
	defer __obf_f86e7ccc55a7708f.Body.Close()

	if __obf_f86e7ccc55a7708f.StatusCode != http.StatusOK {
		__obf_325ae0f19cc81aa7, _ := io.ReadAll(__obf_f86e7ccc55a7708f.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_f86e7ccc55a7708f.StatusCode, string(__obf_325ae0f19cc81aa7))
	}

	if __obf_92da1cf9755f8a83 != nil {
		if __obf_d6bcde9761346c61 := json.NewDecoder(__obf_f86e7ccc55a7708f.Body).Decode(__obf_92da1cf9755f8a83); __obf_d6bcde9761346c61 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d6bcde9761346c61)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_e41c990542ded541 *HttpClient) PostJSON(__obf_e30efbaedbdb33bc string, __obf_0b08da8c797415e9 any, __obf_92da1cf9755f8a83 any) error {
	__obf_7576235c3ea3ff00, __obf_d6bcde9761346c61 := json.Marshal(__obf_0b08da8c797415e9)
	if __obf_d6bcde9761346c61 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_4f33d00a016cd8c0, __obf_d6bcde9761346c61 := http.NewRequest(http.MethodPost, __obf_e30efbaedbdb33bc, bytes.NewBuffer(__obf_7576235c3ea3ff00))
	if __obf_d6bcde9761346c61 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_4f33d00a016cd8c0.
		Header.Set("Content-Type", "application/json")
	__obf_f86e7ccc55a7708f, __obf_d6bcde9761346c61 := __obf_e41c990542ded541.__obf_4650a5d47e8654af.Do(__obf_4f33d00a016cd8c0)
	if __obf_d6bcde9761346c61 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_d6bcde9761346c61)
	}
	defer __obf_f86e7ccc55a7708f.Body.Close()

	if __obf_f86e7ccc55a7708f.StatusCode != http.StatusOK {
		__obf_325ae0f19cc81aa7, _ := io.ReadAll(__obf_f86e7ccc55a7708f.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_f86e7ccc55a7708f.StatusCode, string(__obf_325ae0f19cc81aa7))
	}

	if __obf_92da1cf9755f8a83 != nil {
		if __obf_d6bcde9761346c61 := json.NewDecoder(__obf_f86e7ccc55a7708f.Body).Decode(__obf_92da1cf9755f8a83); __obf_d6bcde9761346c61 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d6bcde9761346c61)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_e41c990542ded541 *HttpClient) PostXML(__obf_e30efbaedbdb33bc string, __obf_0b08da8c797415e9 any) ([]byte, error) {
	__obf_7576235c3ea3ff00, // <- 移除 response any
		__obf_d6bcde9761346c61 := xml.Marshal(__obf_0b08da8c797415e9)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_4f33d00a016cd8c0,
		// fmt.Println("签名报文：", string(reqBody))
		__obf_d6bcde9761346c61 := http.NewRequest(http.MethodPost, __obf_e30efbaedbdb33bc, bytes.NewBuffer(__obf_7576235c3ea3ff00))
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_4f33d00a016cd8c0.
		Header.Set("Content-Type", "application/xml")
	__obf_f86e7ccc55a7708f, __obf_d6bcde9761346c61 := __obf_e41c990542ded541.__obf_4650a5d47e8654af.Do(__obf_4f33d00a016cd8c0)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_d6bcde9761346c61)
	}
	defer __obf_f86e7ccc55a7708f.Body.Close()
	__obf_325ae0f19cc81aa7, __obf_d6bcde9761346c61 := io.ReadAll(__obf_f86e7ccc55a7708f.Body) // <- 读取原始响应体
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_d6bcde9761346c61)
	}

	if __obf_f86e7ccc55a7708f.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_f86e7ccc55a7708f.StatusCode, string(__obf_325ae0f19cc81aa7))
	}

	return __obf_325ae0f19cc81aa7, nil // <- 返回原始响应体
}
