package __obf_0d56cba5b7a269cc

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
	__obf_c42a04cc389c8fe8 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_c42a04cc389c8fe8: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_1babca7578c5b4de tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_733239a93ca42764 := &tls.Config{
		Certificates: []tls.Certificate{__obf_1babca7578c5b4de},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_687d86c9d991d4fe := &http.Transport{
		TLSClientConfig:       __obf_733239a93ca42764,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_c42a04cc389c8fe8: &http.Client{
			Transport: __obf_687d86c9d991d4fe,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_3348df8f53a60bb3 *HttpClient) Get(__obf_62a50bf6bca7744d string, __obf_d7e682e7f1fecceb map[string]string, __obf_0d81203afd78ff7f any) error {
	__obf_bd19768255d7a892, __obf_b8d74ed2e252a551 := http.NewRequest(http.MethodGet, __obf_62a50bf6bca7744d, nil)
	if __obf_b8d74ed2e252a551 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_b8d74ed2e252a551)
	}

	__obf_02cc1ab7217224ea := __obf_bd19768255d7a892.URL.Query()
	for __obf_9d120b736914ab35, __obf_24f87e0512d13675 := range __obf_d7e682e7f1fecceb {
		__obf_02cc1ab7217224ea.Add(__obf_9d120b736914ab35, __obf_24f87e0512d13675)
	}
	__obf_bd19768255d7a892.URL.RawQuery = __obf_02cc1ab7217224ea.Encode()

	__obf_47575840d540f50b, __obf_b8d74ed2e252a551 := __obf_3348df8f53a60bb3.__obf_c42a04cc389c8fe8.Do(__obf_bd19768255d7a892)
	if __obf_b8d74ed2e252a551 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_b8d74ed2e252a551)
	}
	defer __obf_47575840d540f50b.Body.Close()

	if __obf_47575840d540f50b.StatusCode != http.StatusOK {
		__obf_62a5e5be5a49772b, _ := io.ReadAll(__obf_47575840d540f50b.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_47575840d540f50b.StatusCode, string(__obf_62a5e5be5a49772b))
	}

	if __obf_0d81203afd78ff7f != nil {
		if __obf_b8d74ed2e252a551 := json.NewDecoder(__obf_47575840d540f50b.Body).Decode(__obf_0d81203afd78ff7f); __obf_b8d74ed2e252a551 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_b8d74ed2e252a551)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_3348df8f53a60bb3 *HttpClient) PostJSON(__obf_62a50bf6bca7744d string, __obf_0363cc6966fbafe1 any, __obf_0d81203afd78ff7f any) error {
	__obf_3d5e7aa495176b9e, __obf_b8d74ed2e252a551 := json.Marshal(__obf_0363cc6966fbafe1)
	if __obf_b8d74ed2e252a551 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_b8d74ed2e252a551)
	}

	__obf_bd19768255d7a892, __obf_b8d74ed2e252a551 := http.NewRequest(http.MethodPost, __obf_62a50bf6bca7744d, bytes.NewBuffer(__obf_3d5e7aa495176b9e))
	if __obf_b8d74ed2e252a551 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_b8d74ed2e252a551)
	}
	__obf_bd19768255d7a892.Header.Set("Content-Type", "application/json")

	__obf_47575840d540f50b, __obf_b8d74ed2e252a551 := __obf_3348df8f53a60bb3.__obf_c42a04cc389c8fe8.Do(__obf_bd19768255d7a892)
	if __obf_b8d74ed2e252a551 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_b8d74ed2e252a551)
	}
	defer __obf_47575840d540f50b.Body.Close()

	if __obf_47575840d540f50b.StatusCode != http.StatusOK {
		__obf_62a5e5be5a49772b, _ := io.ReadAll(__obf_47575840d540f50b.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_47575840d540f50b.StatusCode, string(__obf_62a5e5be5a49772b))
	}

	if __obf_0d81203afd78ff7f != nil {
		if __obf_b8d74ed2e252a551 := json.NewDecoder(__obf_47575840d540f50b.Body).Decode(__obf_0d81203afd78ff7f); __obf_b8d74ed2e252a551 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_b8d74ed2e252a551)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_3348df8f53a60bb3 *HttpClient) PostXML(__obf_62a50bf6bca7744d string, __obf_0363cc6966fbafe1 any) ([]byte, error) { // <- 移除 response any
	__obf_3d5e7aa495176b9e, __obf_b8d74ed2e252a551 := xml.Marshal(__obf_0363cc6966fbafe1)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_b8d74ed2e252a551)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_bd19768255d7a892, __obf_b8d74ed2e252a551 := http.NewRequest(http.MethodPost, __obf_62a50bf6bca7744d, bytes.NewBuffer(__obf_3d5e7aa495176b9e))
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_b8d74ed2e252a551)
	}
	__obf_bd19768255d7a892.Header.Set("Content-Type", "application/xml")

	__obf_47575840d540f50b, __obf_b8d74ed2e252a551 := __obf_3348df8f53a60bb3.__obf_c42a04cc389c8fe8.Do(__obf_bd19768255d7a892)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_b8d74ed2e252a551)
	}
	defer __obf_47575840d540f50b.Body.Close()

	__obf_62a5e5be5a49772b, __obf_b8d74ed2e252a551 := io.ReadAll(__obf_47575840d540f50b.Body) // <- 读取原始响应体
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_b8d74ed2e252a551)
	}

	if __obf_47575840d540f50b.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_47575840d540f50b.StatusCode, string(__obf_62a5e5be5a49772b))
	}

	return __obf_62a5e5be5a49772b, nil // <- 返回原始响应体
}
