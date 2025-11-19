package __obf_2e4735ec379515a2

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
	__obf_aab07d5c71075d1b *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_aab07d5c71075d1b: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_a25314ceb42158aa tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_b82f20ba1ee67338 := &tls.Config{
		Certificates: []tls.Certificate{__obf_a25314ceb42158aa},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_d0661d2634c2b2bf := &http.Transport{
		TLSClientConfig:       __obf_b82f20ba1ee67338,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_aab07d5c71075d1b: &http.Client{
			Transport: __obf_d0661d2634c2b2bf,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_47ce5c326294d040 *HttpClient) Get(__obf_783ad16d817afd13 string, __obf_aaf231d4acfd72d0 map[string]string, __obf_0d6dbb07eb644b0f any) error {
	__obf_c259e7ec915f35d4, __obf_647657e89e0327f9 := http.NewRequest(http.MethodGet, __obf_783ad16d817afd13, nil)
	if __obf_647657e89e0327f9 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_647657e89e0327f9)
	}

	__obf_367bae35db10655a := __obf_c259e7ec915f35d4.URL.Query()
	for __obf_2df09d571b10e138, __obf_62acc15bb7c68aaa := range __obf_aaf231d4acfd72d0 {
		__obf_367bae35db10655a.Add(__obf_2df09d571b10e138, __obf_62acc15bb7c68aaa)
	}
	__obf_c259e7ec915f35d4.URL.RawQuery = __obf_367bae35db10655a.Encode()

	__obf_ac3f062913034a2c, __obf_647657e89e0327f9 := __obf_47ce5c326294d040.__obf_aab07d5c71075d1b.Do(__obf_c259e7ec915f35d4)
	if __obf_647657e89e0327f9 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_647657e89e0327f9)
	}
	defer __obf_ac3f062913034a2c.Body.Close()

	if __obf_ac3f062913034a2c.StatusCode != http.StatusOK {
		__obf_77cbc49753c8e9b5, _ := io.ReadAll(__obf_ac3f062913034a2c.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_ac3f062913034a2c.StatusCode, string(__obf_77cbc49753c8e9b5))
	}

	if __obf_0d6dbb07eb644b0f != nil {
		if __obf_647657e89e0327f9 := json.NewDecoder(__obf_ac3f062913034a2c.Body).Decode(__obf_0d6dbb07eb644b0f); __obf_647657e89e0327f9 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_647657e89e0327f9)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_47ce5c326294d040 *HttpClient) PostJSON(__obf_783ad16d817afd13 string, __obf_ed710df1b85302ba any, __obf_0d6dbb07eb644b0f any) error {
	__obf_1133a9bdfa0db7df, __obf_647657e89e0327f9 := json.Marshal(__obf_ed710df1b85302ba)
	if __obf_647657e89e0327f9 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_647657e89e0327f9)
	}

	__obf_c259e7ec915f35d4, __obf_647657e89e0327f9 := http.NewRequest(http.MethodPost, __obf_783ad16d817afd13, bytes.NewBuffer(__obf_1133a9bdfa0db7df))
	if __obf_647657e89e0327f9 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_647657e89e0327f9)
	}
	__obf_c259e7ec915f35d4.Header.Set("Content-Type", "application/json")

	__obf_ac3f062913034a2c, __obf_647657e89e0327f9 := __obf_47ce5c326294d040.__obf_aab07d5c71075d1b.Do(__obf_c259e7ec915f35d4)
	if __obf_647657e89e0327f9 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_647657e89e0327f9)
	}
	defer __obf_ac3f062913034a2c.Body.Close()

	if __obf_ac3f062913034a2c.StatusCode != http.StatusOK {
		__obf_77cbc49753c8e9b5, _ := io.ReadAll(__obf_ac3f062913034a2c.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_ac3f062913034a2c.StatusCode, string(__obf_77cbc49753c8e9b5))
	}

	if __obf_0d6dbb07eb644b0f != nil {
		if __obf_647657e89e0327f9 := json.NewDecoder(__obf_ac3f062913034a2c.Body).Decode(__obf_0d6dbb07eb644b0f); __obf_647657e89e0327f9 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_647657e89e0327f9)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_47ce5c326294d040 *HttpClient) PostXML(__obf_783ad16d817afd13 string, __obf_ed710df1b85302ba any) ([]byte, error) { // <- 移除 response any
	__obf_1133a9bdfa0db7df, __obf_647657e89e0327f9 := xml.Marshal(__obf_ed710df1b85302ba)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_647657e89e0327f9)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_c259e7ec915f35d4, __obf_647657e89e0327f9 := http.NewRequest(http.MethodPost, __obf_783ad16d817afd13, bytes.NewBuffer(__obf_1133a9bdfa0db7df))
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_647657e89e0327f9)
	}
	__obf_c259e7ec915f35d4.Header.Set("Content-Type", "application/xml")

	__obf_ac3f062913034a2c, __obf_647657e89e0327f9 := __obf_47ce5c326294d040.__obf_aab07d5c71075d1b.Do(__obf_c259e7ec915f35d4)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_647657e89e0327f9)
	}
	defer __obf_ac3f062913034a2c.Body.Close()

	__obf_77cbc49753c8e9b5, __obf_647657e89e0327f9 := io.ReadAll(__obf_ac3f062913034a2c.Body) // <- 读取原始响应体
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_647657e89e0327f9)
	}

	if __obf_ac3f062913034a2c.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_ac3f062913034a2c.StatusCode, string(__obf_77cbc49753c8e9b5))
	}

	return __obf_77cbc49753c8e9b5, nil // <- 返回原始响应体
}
