package __obf_79882becdca92e5e

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
	__obf_29c81bf1c5877bfa *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_29c81bf1c5877bfa: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_cc0d5c5af9e977ff tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_7f28cf0c58fdaa9e := &tls.Config{
		Certificates: []tls.Certificate{__obf_cc0d5c5af9e977ff},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_f1c7d5c00e0ab15f := &http.Transport{
		TLSClientConfig:       __obf_7f28cf0c58fdaa9e,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_29c81bf1c5877bfa: &http.Client{
			Transport: __obf_f1c7d5c00e0ab15f,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_9e00ff0ebddedf17 *HttpClient) Get(__obf_50fdf03c3486f894 string, __obf_b6a776f5c3ec2966 map[string]string, __obf_5634fa1c89be6fd1 any) error {
	__obf_1560934e6aff953a, __obf_d0090ed9a7614e3f := http.NewRequest(http.MethodGet, __obf_50fdf03c3486f894, nil)
	if __obf_d0090ed9a7614e3f != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_d0090ed9a7614e3f)
	}

	__obf_d65c5d5c84752a32 := __obf_1560934e6aff953a.URL.Query()
	for __obf_58a71da0eb799615, __obf_0be32d5ff1d7c891 := range __obf_b6a776f5c3ec2966 {
		__obf_d65c5d5c84752a32.Add(__obf_58a71da0eb799615, __obf_0be32d5ff1d7c891)
	}
	__obf_1560934e6aff953a.URL.RawQuery = __obf_d65c5d5c84752a32.Encode()

	__obf_7bfa6f69f22c40a3, __obf_d0090ed9a7614e3f := __obf_9e00ff0ebddedf17.__obf_29c81bf1c5877bfa.Do(__obf_1560934e6aff953a)
	if __obf_d0090ed9a7614e3f != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_d0090ed9a7614e3f)
	}
	defer __obf_7bfa6f69f22c40a3.Body.Close()

	if __obf_7bfa6f69f22c40a3.StatusCode != http.StatusOK {
		__obf_7971587456c51f68, _ := io.ReadAll(__obf_7bfa6f69f22c40a3.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_7bfa6f69f22c40a3.StatusCode, string(__obf_7971587456c51f68))
	}

	if __obf_5634fa1c89be6fd1 != nil {
		if __obf_d0090ed9a7614e3f := json.NewDecoder(__obf_7bfa6f69f22c40a3.Body).Decode(__obf_5634fa1c89be6fd1); __obf_d0090ed9a7614e3f != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d0090ed9a7614e3f)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_9e00ff0ebddedf17 *HttpClient) PostJSON(__obf_50fdf03c3486f894 string, __obf_a648625f69b72676 any, __obf_5634fa1c89be6fd1 any) error {
	__obf_833c3aa3fcc17f22, __obf_d0090ed9a7614e3f := json.Marshal(__obf_a648625f69b72676)
	if __obf_d0090ed9a7614e3f != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_d0090ed9a7614e3f)
	}

	__obf_1560934e6aff953a, __obf_d0090ed9a7614e3f := http.NewRequest(http.MethodPost, __obf_50fdf03c3486f894, bytes.NewBuffer(__obf_833c3aa3fcc17f22))
	if __obf_d0090ed9a7614e3f != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_d0090ed9a7614e3f)
	}
	__obf_1560934e6aff953a.Header.Set("Content-Type", "application/json")

	__obf_7bfa6f69f22c40a3, __obf_d0090ed9a7614e3f := __obf_9e00ff0ebddedf17.__obf_29c81bf1c5877bfa.Do(__obf_1560934e6aff953a)
	if __obf_d0090ed9a7614e3f != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_d0090ed9a7614e3f)
	}
	defer __obf_7bfa6f69f22c40a3.Body.Close()

	if __obf_7bfa6f69f22c40a3.StatusCode != http.StatusOK {
		__obf_7971587456c51f68, _ := io.ReadAll(__obf_7bfa6f69f22c40a3.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_7bfa6f69f22c40a3.StatusCode, string(__obf_7971587456c51f68))
	}

	if __obf_5634fa1c89be6fd1 != nil {
		if __obf_d0090ed9a7614e3f := json.NewDecoder(__obf_7bfa6f69f22c40a3.Body).Decode(__obf_5634fa1c89be6fd1); __obf_d0090ed9a7614e3f != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d0090ed9a7614e3f)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_9e00ff0ebddedf17 *HttpClient) PostXML(__obf_50fdf03c3486f894 string, __obf_a648625f69b72676 any) ([]byte, error) { // <- 移除 response any
	__obf_833c3aa3fcc17f22, __obf_d0090ed9a7614e3f := xml.Marshal(__obf_a648625f69b72676)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_d0090ed9a7614e3f)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_1560934e6aff953a, __obf_d0090ed9a7614e3f := http.NewRequest(http.MethodPost, __obf_50fdf03c3486f894, bytes.NewBuffer(__obf_833c3aa3fcc17f22))
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_d0090ed9a7614e3f)
	}
	__obf_1560934e6aff953a.Header.Set("Content-Type", "application/xml")

	__obf_7bfa6f69f22c40a3, __obf_d0090ed9a7614e3f := __obf_9e00ff0ebddedf17.__obf_29c81bf1c5877bfa.Do(__obf_1560934e6aff953a)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_d0090ed9a7614e3f)
	}
	defer __obf_7bfa6f69f22c40a3.Body.Close()

	__obf_7971587456c51f68, __obf_d0090ed9a7614e3f := io.ReadAll(__obf_7bfa6f69f22c40a3.Body) // <- 读取原始响应体
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_d0090ed9a7614e3f)
	}

	if __obf_7bfa6f69f22c40a3.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_7bfa6f69f22c40a3.StatusCode, string(__obf_7971587456c51f68))
	}

	return __obf_7971587456c51f68, nil // <- 返回原始响应体
}
