package __obf_b5bcac367b722f8a

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
	__obf_2c20c3187b60d1af *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{__obf_2c20c3187b60d1af: &http.Client{
		Timeout: 15 * time.Second, // 设置一个合理的超时时间
	},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_dee17fdde39832e9 tls.Certificate) (*HttpClient, error) {
	__obf_14000579e6973015 := // 加载客户端证书和私钥
		// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		// if err != nil {
		// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
		// }

		// 构造TLS配置
		&tls.Config{
			Certificates: []tls.Certificate{__obf_dee17fdde39832e9},
			// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
			// RootCAs:            nil, // 如果有自定义CA，可以设置
		}
	__obf_2c42ac4ac4d00852 := // 创建带证书的HTTP Transport
		&http.Transport{
			TLSClientConfig:       __obf_14000579e6973015,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

	return &HttpClient{__obf_2c20c3187b60d1af: &http.Client{
		Transport: __obf_2c42ac4ac4d00852,
		Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
	},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_743a5d9ae9ed3d6e *HttpClient) Get(__obf_5318b7d070036030 string, __obf_b48de75b986fba1b map[string]string, __obf_3a709bb0aff380c2 any) error {
	__obf_4e8d9d056fad54a9, __obf_d9b19f08bff09faf := http.NewRequest(http.MethodGet, __obf_5318b7d070036030, nil)
	if __obf_d9b19f08bff09faf != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_edbdb434eef08911 := __obf_4e8d9d056fad54a9.URL.Query()
	for __obf_a0009a4f0921242e, __obf_70d42ebd71a0d3b2 := range __obf_b48de75b986fba1b {
		__obf_edbdb434eef08911.
			Add(__obf_a0009a4f0921242e, __obf_70d42ebd71a0d3b2)
	}
	__obf_4e8d9d056fad54a9.
		URL.RawQuery = __obf_edbdb434eef08911.Encode()
	__obf_1fa9f007d7bd47be, __obf_d9b19f08bff09faf := __obf_743a5d9ae9ed3d6e.__obf_2c20c3187b60d1af.Do(__obf_4e8d9d056fad54a9)
	if __obf_d9b19f08bff09faf != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_d9b19f08bff09faf)
	}
	defer __obf_1fa9f007d7bd47be.Body.Close()

	if __obf_1fa9f007d7bd47be.StatusCode != http.StatusOK {
		__obf_05a65c198084ba0e, _ := io.ReadAll(__obf_1fa9f007d7bd47be.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_1fa9f007d7bd47be.StatusCode, string(__obf_05a65c198084ba0e))
	}

	if __obf_3a709bb0aff380c2 != nil {
		if __obf_d9b19f08bff09faf := json.NewDecoder(__obf_1fa9f007d7bd47be.Body).Decode(__obf_3a709bb0aff380c2); __obf_d9b19f08bff09faf != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d9b19f08bff09faf)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_743a5d9ae9ed3d6e *HttpClient) PostJSON(__obf_5318b7d070036030 string, __obf_a8eddbae47f701c1 any, __obf_3a709bb0aff380c2 any) error {
	__obf_b1881c73240a7045, __obf_d9b19f08bff09faf := json.Marshal(__obf_a8eddbae47f701c1)
	if __obf_d9b19f08bff09faf != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_4e8d9d056fad54a9, __obf_d9b19f08bff09faf := http.NewRequest(http.MethodPost, __obf_5318b7d070036030, bytes.NewBuffer(__obf_b1881c73240a7045))
	if __obf_d9b19f08bff09faf != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_4e8d9d056fad54a9.
		Header.Set("Content-Type", "application/json")
	__obf_1fa9f007d7bd47be, __obf_d9b19f08bff09faf := __obf_743a5d9ae9ed3d6e.__obf_2c20c3187b60d1af.Do(__obf_4e8d9d056fad54a9)
	if __obf_d9b19f08bff09faf != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_d9b19f08bff09faf)
	}
	defer __obf_1fa9f007d7bd47be.Body.Close()

	if __obf_1fa9f007d7bd47be.StatusCode != http.StatusOK {
		__obf_05a65c198084ba0e, _ := io.ReadAll(__obf_1fa9f007d7bd47be.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_1fa9f007d7bd47be.StatusCode, string(__obf_05a65c198084ba0e))
	}

	if __obf_3a709bb0aff380c2 != nil {
		if __obf_d9b19f08bff09faf := json.NewDecoder(__obf_1fa9f007d7bd47be.Body).Decode(__obf_3a709bb0aff380c2); __obf_d9b19f08bff09faf != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d9b19f08bff09faf)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_743a5d9ae9ed3d6e *HttpClient) PostXML(__obf_5318b7d070036030 string, __obf_a8eddbae47f701c1 any) ([]byte, error) {
	__obf_b1881c73240a7045, // <- 移除 response any
		__obf_d9b19f08bff09faf := xml.Marshal(__obf_a8eddbae47f701c1)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_4e8d9d056fad54a9,
		// fmt.Println("签名报文：", string(reqBody))
		__obf_d9b19f08bff09faf := http.NewRequest(http.MethodPost, __obf_5318b7d070036030, bytes.NewBuffer(__obf_b1881c73240a7045))
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_4e8d9d056fad54a9.
		Header.Set("Content-Type", "application/xml")
	__obf_1fa9f007d7bd47be, __obf_d9b19f08bff09faf := __obf_743a5d9ae9ed3d6e.__obf_2c20c3187b60d1af.Do(__obf_4e8d9d056fad54a9)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_d9b19f08bff09faf)
	}
	defer __obf_1fa9f007d7bd47be.Body.Close()
	__obf_05a65c198084ba0e, __obf_d9b19f08bff09faf := io.ReadAll(__obf_1fa9f007d7bd47be.Body) // <- 读取原始响应体
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_d9b19f08bff09faf)
	}

	if __obf_1fa9f007d7bd47be.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_1fa9f007d7bd47be.StatusCode, string(__obf_05a65c198084ba0e))
	}

	return __obf_05a65c198084ba0e, nil // <- 返回原始响应体
}
