package __obf_abe10294567bade2

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
	__obf_02aa4200ef510880 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{__obf_02aa4200ef510880: &http.Client{
		Timeout: 15 * time.Second, // 设置一个合理的超时时间
	},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_4918c583030326c5 tls.Certificate) (*HttpClient, error) {
	__obf_a35457144ba87221 := // 加载客户端证书和私钥
		// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		// if err != nil {
		// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
		// }

		// 构造TLS配置
		&tls.Config{
			Certificates: []tls.Certificate{__obf_4918c583030326c5},
			// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
			// RootCAs:            nil, // 如果有自定义CA，可以设置
		}
	__obf_84b958dfdd634ac1 := // 创建带证书的HTTP Transport
		&http.Transport{
			TLSClientConfig:       __obf_a35457144ba87221,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

	return &HttpClient{__obf_02aa4200ef510880: &http.Client{
		Transport: __obf_84b958dfdd634ac1,
		Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
	},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_9c841ff59930ecf7 *HttpClient) Get(__obf_5589c44e6fad90cb string, __obf_4283e9b020db61f4 map[string]string, __obf_d38cbe7bc815eb3d any) error {
	__obf_7b998eaf76b551b5, __obf_02af51ab07084000 := http.NewRequest(http.MethodGet, __obf_5589c44e6fad90cb, nil)
	if __obf_02af51ab07084000 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_02af51ab07084000)
	}
	__obf_dabbfc651c2ca4c1 := __obf_7b998eaf76b551b5.URL.Query()
	for __obf_685e68f3451766f3, __obf_7ef4adf3e75f965e := range __obf_4283e9b020db61f4 {
		__obf_dabbfc651c2ca4c1.
			Add(__obf_685e68f3451766f3, __obf_7ef4adf3e75f965e)
	}
	__obf_7b998eaf76b551b5.
		URL.RawQuery = __obf_dabbfc651c2ca4c1.Encode()
	__obf_1f92ef77fe1d922f, __obf_02af51ab07084000 := __obf_9c841ff59930ecf7.__obf_02aa4200ef510880.Do(__obf_7b998eaf76b551b5)
	if __obf_02af51ab07084000 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_02af51ab07084000)
	}
	defer __obf_1f92ef77fe1d922f.Body.Close()

	if __obf_1f92ef77fe1d922f.StatusCode != http.StatusOK {
		__obf_84fa132f511549dc, _ := io.ReadAll(__obf_1f92ef77fe1d922f.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_1f92ef77fe1d922f.StatusCode, string(__obf_84fa132f511549dc))
	}

	if __obf_d38cbe7bc815eb3d != nil {
		if __obf_02af51ab07084000 := json.NewDecoder(__obf_1f92ef77fe1d922f.Body).Decode(__obf_d38cbe7bc815eb3d); __obf_02af51ab07084000 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_02af51ab07084000)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_9c841ff59930ecf7 *HttpClient) PostJSON(__obf_5589c44e6fad90cb string, __obf_6630b95eacec778a any, __obf_d38cbe7bc815eb3d any) error {
	__obf_aecf8a7ab4334c27, __obf_02af51ab07084000 := json.Marshal(__obf_6630b95eacec778a)
	if __obf_02af51ab07084000 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_02af51ab07084000)
	}
	__obf_7b998eaf76b551b5, __obf_02af51ab07084000 := http.NewRequest(http.MethodPost, __obf_5589c44e6fad90cb, bytes.NewBuffer(__obf_aecf8a7ab4334c27))
	if __obf_02af51ab07084000 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_02af51ab07084000)
	}
	__obf_7b998eaf76b551b5.
		Header.Set("Content-Type", "application/json")
	__obf_1f92ef77fe1d922f, __obf_02af51ab07084000 := __obf_9c841ff59930ecf7.__obf_02aa4200ef510880.Do(__obf_7b998eaf76b551b5)
	if __obf_02af51ab07084000 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_02af51ab07084000)
	}
	defer __obf_1f92ef77fe1d922f.Body.Close()

	if __obf_1f92ef77fe1d922f.StatusCode != http.StatusOK {
		__obf_84fa132f511549dc, _ := io.ReadAll(__obf_1f92ef77fe1d922f.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_1f92ef77fe1d922f.StatusCode, string(__obf_84fa132f511549dc))
	}

	if __obf_d38cbe7bc815eb3d != nil {
		if __obf_02af51ab07084000 := json.NewDecoder(__obf_1f92ef77fe1d922f.Body).Decode(__obf_d38cbe7bc815eb3d); __obf_02af51ab07084000 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_02af51ab07084000)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_9c841ff59930ecf7 *HttpClient) PostXML(__obf_5589c44e6fad90cb string, __obf_6630b95eacec778a any) ([]byte, error) {
	__obf_aecf8a7ab4334c27, // <- 移除 response any
		__obf_02af51ab07084000 := xml.Marshal(__obf_6630b95eacec778a)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_02af51ab07084000)
	}
	__obf_7b998eaf76b551b5,
		// fmt.Println("签名报文：", string(reqBody))
		__obf_02af51ab07084000 := http.NewRequest(http.MethodPost, __obf_5589c44e6fad90cb, bytes.NewBuffer(__obf_aecf8a7ab4334c27))
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_02af51ab07084000)
	}
	__obf_7b998eaf76b551b5.
		Header.Set("Content-Type", "application/xml")
	__obf_1f92ef77fe1d922f, __obf_02af51ab07084000 := __obf_9c841ff59930ecf7.__obf_02aa4200ef510880.Do(__obf_7b998eaf76b551b5)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_02af51ab07084000)
	}
	defer __obf_1f92ef77fe1d922f.Body.Close()
	__obf_84fa132f511549dc, __obf_02af51ab07084000 := io.ReadAll(__obf_1f92ef77fe1d922f.Body) // <- 读取原始响应体
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_02af51ab07084000)
	}

	if __obf_1f92ef77fe1d922f.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_1f92ef77fe1d922f.StatusCode, string(__obf_84fa132f511549dc))
	}

	return __obf_84fa132f511549dc, nil // <- 返回原始响应体
}
