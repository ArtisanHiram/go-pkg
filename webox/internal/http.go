package __obf_6e18fdc479ab921c

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
	__obf_f1568dc294e85369 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_f1568dc294e85369: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_6b0046b7bc335469 tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_56d6ac9fb3d94c1f := &tls.Config{
		Certificates: []tls.Certificate{__obf_6b0046b7bc335469},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_1c2a6d144438c15a := &http.Transport{
		TLSClientConfig:       __obf_56d6ac9fb3d94c1f,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_f1568dc294e85369: &http.Client{
			Transport: __obf_1c2a6d144438c15a,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_8dd53513888e02ec *HttpClient) Get(__obf_7f7a63542215a305 string, __obf_dd163274f47d9625 map[string]string, __obf_7d181abfc08ebc9b any) error {
	__obf_11b74f644d976f78, __obf_d92b02392b1fb408 := http.NewRequest(http.MethodGet, __obf_7f7a63542215a305, nil)
	if __obf_d92b02392b1fb408 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_d92b02392b1fb408)
	}

	__obf_64374e64fc589e19 := __obf_11b74f644d976f78.URL.Query()
	for __obf_4dde8a7421899bf2, __obf_de9cd12b999685f6 := range __obf_dd163274f47d9625 {
		__obf_64374e64fc589e19.Add(__obf_4dde8a7421899bf2, __obf_de9cd12b999685f6)
	}
	__obf_11b74f644d976f78.URL.RawQuery = __obf_64374e64fc589e19.Encode()

	__obf_e79bf9ab11f94ab7, __obf_d92b02392b1fb408 := __obf_8dd53513888e02ec.__obf_f1568dc294e85369.Do(__obf_11b74f644d976f78)
	if __obf_d92b02392b1fb408 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_d92b02392b1fb408)
	}
	defer __obf_e79bf9ab11f94ab7.Body.Close()

	if __obf_e79bf9ab11f94ab7.StatusCode != http.StatusOK {
		__obf_3adad04f2ac6ade5, _ := io.ReadAll(__obf_e79bf9ab11f94ab7.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_e79bf9ab11f94ab7.StatusCode, string(__obf_3adad04f2ac6ade5))
	}

	if __obf_7d181abfc08ebc9b != nil {
		if __obf_d92b02392b1fb408 := json.NewDecoder(__obf_e79bf9ab11f94ab7.Body).Decode(__obf_7d181abfc08ebc9b); __obf_d92b02392b1fb408 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d92b02392b1fb408)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_8dd53513888e02ec *HttpClient) PostJSON(__obf_7f7a63542215a305 string, __obf_fdbe37be0f192b33 any, __obf_7d181abfc08ebc9b any) error {
	__obf_d945fc686e1ea16a, __obf_d92b02392b1fb408 := json.Marshal(__obf_fdbe37be0f192b33)
	if __obf_d92b02392b1fb408 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_d92b02392b1fb408)
	}

	__obf_11b74f644d976f78, __obf_d92b02392b1fb408 := http.NewRequest(http.MethodPost, __obf_7f7a63542215a305, bytes.NewBuffer(__obf_d945fc686e1ea16a))
	if __obf_d92b02392b1fb408 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_d92b02392b1fb408)
	}
	__obf_11b74f644d976f78.Header.Set("Content-Type", "application/json")

	__obf_e79bf9ab11f94ab7, __obf_d92b02392b1fb408 := __obf_8dd53513888e02ec.__obf_f1568dc294e85369.Do(__obf_11b74f644d976f78)
	if __obf_d92b02392b1fb408 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_d92b02392b1fb408)
	}
	defer __obf_e79bf9ab11f94ab7.Body.Close()

	if __obf_e79bf9ab11f94ab7.StatusCode != http.StatusOK {
		__obf_3adad04f2ac6ade5, _ := io.ReadAll(__obf_e79bf9ab11f94ab7.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_e79bf9ab11f94ab7.StatusCode, string(__obf_3adad04f2ac6ade5))
	}

	if __obf_7d181abfc08ebc9b != nil {
		if __obf_d92b02392b1fb408 := json.NewDecoder(__obf_e79bf9ab11f94ab7.Body).Decode(__obf_7d181abfc08ebc9b); __obf_d92b02392b1fb408 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_d92b02392b1fb408)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_8dd53513888e02ec *HttpClient) PostXML(__obf_7f7a63542215a305 string, __obf_fdbe37be0f192b33 any) ([]byte, error) { // <- 移除 response any
	__obf_d945fc686e1ea16a, __obf_d92b02392b1fb408 := xml.Marshal(__obf_fdbe37be0f192b33)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_d92b02392b1fb408)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_11b74f644d976f78, __obf_d92b02392b1fb408 := http.NewRequest(http.MethodPost, __obf_7f7a63542215a305, bytes.NewBuffer(__obf_d945fc686e1ea16a))
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_d92b02392b1fb408)
	}
	__obf_11b74f644d976f78.Header.Set("Content-Type", "application/xml")

	__obf_e79bf9ab11f94ab7, __obf_d92b02392b1fb408 := __obf_8dd53513888e02ec.__obf_f1568dc294e85369.Do(__obf_11b74f644d976f78)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_d92b02392b1fb408)
	}
	defer __obf_e79bf9ab11f94ab7.Body.Close()

	__obf_3adad04f2ac6ade5, __obf_d92b02392b1fb408 := io.ReadAll(__obf_e79bf9ab11f94ab7.Body) // <- 读取原始响应体
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_d92b02392b1fb408)
	}

	if __obf_e79bf9ab11f94ab7.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_e79bf9ab11f94ab7.StatusCode, string(__obf_3adad04f2ac6ade5))
	}

	return __obf_3adad04f2ac6ade5, nil // <- 返回原始响应体
}
