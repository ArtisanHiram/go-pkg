package __obf_c0d7bb2e04898f29

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
	__obf_d7e3a8a5eebdb96b *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_d7e3a8a5eebdb96b: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_1e2603d75e893450 tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_871bafc1b52ca1c8 := &tls.Config{
		Certificates: []tls.Certificate{__obf_1e2603d75e893450},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_b3b2be87ead018b9 := &http.Transport{
		TLSClientConfig:       __obf_871bafc1b52ca1c8,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_d7e3a8a5eebdb96b: &http.Client{
			Transport: __obf_b3b2be87ead018b9,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_b5c164dcca9f4e8d *HttpClient) Get(__obf_dca9797b5521c84f string, __obf_a0f9b06f567f969e map[string]string, __obf_a28ad583524f1d89 any) error {
	__obf_8c6dd339ebeafcbd, __obf_fb6d611ef291e586 := http.NewRequest(http.MethodGet, __obf_dca9797b5521c84f, nil)
	if __obf_fb6d611ef291e586 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_fb6d611ef291e586)
	}

	__obf_d451a38d4260f92c := __obf_8c6dd339ebeafcbd.URL.Query()
	for __obf_f8faaa11642d2945, __obf_1a8fb20ade71b8f0 := range __obf_a0f9b06f567f969e {
		__obf_d451a38d4260f92c.Add(__obf_f8faaa11642d2945, __obf_1a8fb20ade71b8f0)
	}
	__obf_8c6dd339ebeafcbd.URL.RawQuery = __obf_d451a38d4260f92c.Encode()

	__obf_697969e7314a2ef9, __obf_fb6d611ef291e586 := __obf_b5c164dcca9f4e8d.__obf_d7e3a8a5eebdb96b.Do(__obf_8c6dd339ebeafcbd)
	if __obf_fb6d611ef291e586 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_fb6d611ef291e586)
	}
	defer __obf_697969e7314a2ef9.Body.Close()

	if __obf_697969e7314a2ef9.StatusCode != http.StatusOK {
		__obf_54d42e5bf8df7934, _ := io.ReadAll(__obf_697969e7314a2ef9.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_697969e7314a2ef9.StatusCode, string(__obf_54d42e5bf8df7934))
	}

	if __obf_a28ad583524f1d89 != nil {
		if __obf_fb6d611ef291e586 := json.NewDecoder(__obf_697969e7314a2ef9.Body).Decode(__obf_a28ad583524f1d89); __obf_fb6d611ef291e586 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_fb6d611ef291e586)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_b5c164dcca9f4e8d *HttpClient) PostJSON(__obf_dca9797b5521c84f string, __obf_e561ad363a248750 any, __obf_a28ad583524f1d89 any) error {
	__obf_ba263a0a818cd60e, __obf_fb6d611ef291e586 := json.Marshal(__obf_e561ad363a248750)
	if __obf_fb6d611ef291e586 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_fb6d611ef291e586)
	}

	__obf_8c6dd339ebeafcbd, __obf_fb6d611ef291e586 := http.NewRequest(http.MethodPost, __obf_dca9797b5521c84f, bytes.NewBuffer(__obf_ba263a0a818cd60e))
	if __obf_fb6d611ef291e586 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_fb6d611ef291e586)
	}
	__obf_8c6dd339ebeafcbd.Header.Set("Content-Type", "application/json")

	__obf_697969e7314a2ef9, __obf_fb6d611ef291e586 := __obf_b5c164dcca9f4e8d.__obf_d7e3a8a5eebdb96b.Do(__obf_8c6dd339ebeafcbd)
	if __obf_fb6d611ef291e586 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_fb6d611ef291e586)
	}
	defer __obf_697969e7314a2ef9.Body.Close()

	if __obf_697969e7314a2ef9.StatusCode != http.StatusOK {
		__obf_54d42e5bf8df7934, _ := io.ReadAll(__obf_697969e7314a2ef9.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_697969e7314a2ef9.StatusCode, string(__obf_54d42e5bf8df7934))
	}

	if __obf_a28ad583524f1d89 != nil {
		if __obf_fb6d611ef291e586 := json.NewDecoder(__obf_697969e7314a2ef9.Body).Decode(__obf_a28ad583524f1d89); __obf_fb6d611ef291e586 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_fb6d611ef291e586)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_b5c164dcca9f4e8d *HttpClient) PostXML(__obf_dca9797b5521c84f string, __obf_e561ad363a248750 any) ([]byte, error) { // <- 移除 response any
	__obf_ba263a0a818cd60e, __obf_fb6d611ef291e586 := xml.Marshal(__obf_e561ad363a248750)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_fb6d611ef291e586)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_8c6dd339ebeafcbd, __obf_fb6d611ef291e586 := http.NewRequest(http.MethodPost, __obf_dca9797b5521c84f, bytes.NewBuffer(__obf_ba263a0a818cd60e))
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_fb6d611ef291e586)
	}
	__obf_8c6dd339ebeafcbd.Header.Set("Content-Type", "application/xml")

	__obf_697969e7314a2ef9, __obf_fb6d611ef291e586 := __obf_b5c164dcca9f4e8d.__obf_d7e3a8a5eebdb96b.Do(__obf_8c6dd339ebeafcbd)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_fb6d611ef291e586)
	}
	defer __obf_697969e7314a2ef9.Body.Close()

	__obf_54d42e5bf8df7934, __obf_fb6d611ef291e586 := io.ReadAll(__obf_697969e7314a2ef9.Body) // <- 读取原始响应体
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_fb6d611ef291e586)
	}

	if __obf_697969e7314a2ef9.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_697969e7314a2ef9.StatusCode, string(__obf_54d42e5bf8df7934))
	}

	return __obf_54d42e5bf8df7934, nil // <- 返回原始响应体
}
