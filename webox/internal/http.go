package __obf_dfacbdcb930f08e7

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
	__obf_8d976fc7dd8ff145 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{__obf_8d976fc7dd8ff145: &http.Client{
		Timeout: 15 * time.Second, // 设置一个合理的超时时间
	},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_92421723b2422f3c tls.Certificate) (*HttpClient, error) {
	__obf_96e527d4ed5d047b := // 加载客户端证书和私钥
		// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		// if err != nil {
		// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
		// }

		// 构造TLS配置
		&tls.Config{
			Certificates: []tls.Certificate{__obf_92421723b2422f3c},
			// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
			// RootCAs:            nil, // 如果有自定义CA，可以设置
		}
	__obf_1e02cd0e8da6d41e := // 创建带证书的HTTP Transport
		&http.Transport{
			TLSClientConfig:       __obf_96e527d4ed5d047b,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

	return &HttpClient{__obf_8d976fc7dd8ff145: &http.Client{
		Transport: __obf_1e02cd0e8da6d41e,
		Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
	},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_c072161b30da1f19 *HttpClient) Get(__obf_bd203f615f84affd string, __obf_f3eaf1026792ed90 map[string]string, __obf_66c10a7c7056ad07 any) error {
	__obf_ee6219419a712a37, __obf_f304d4a353574273 := http.NewRequest(http.MethodGet, __obf_bd203f615f84affd, nil)
	if __obf_f304d4a353574273 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_f304d4a353574273)
	}
	__obf_594fa1285fc8ab1e := __obf_ee6219419a712a37.URL.Query()
	for __obf_6ec0c0d7916f38a1, __obf_dc967b23178581e0 := range __obf_f3eaf1026792ed90 {
		__obf_594fa1285fc8ab1e.
			Add(__obf_6ec0c0d7916f38a1, __obf_dc967b23178581e0)
	}
	__obf_ee6219419a712a37.
		URL.RawQuery = __obf_594fa1285fc8ab1e.Encode()
	__obf_ac69f99159f4f61e, __obf_f304d4a353574273 := __obf_c072161b30da1f19.__obf_8d976fc7dd8ff145.Do(__obf_ee6219419a712a37)
	if __obf_f304d4a353574273 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_f304d4a353574273)
	}
	defer __obf_ac69f99159f4f61e.Body.Close()

	if __obf_ac69f99159f4f61e.StatusCode != http.StatusOK {
		__obf_93b6c9ccae25de12, _ := io.ReadAll(__obf_ac69f99159f4f61e.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_ac69f99159f4f61e.StatusCode, string(__obf_93b6c9ccae25de12))
	}

	if __obf_66c10a7c7056ad07 != nil {
		if __obf_f304d4a353574273 := json.NewDecoder(__obf_ac69f99159f4f61e.Body).Decode(__obf_66c10a7c7056ad07); __obf_f304d4a353574273 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_f304d4a353574273)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_c072161b30da1f19 *HttpClient) PostJSON(__obf_bd203f615f84affd string, __obf_c5f38b6ffd40036c any, __obf_66c10a7c7056ad07 any) error {
	__obf_906ba8325595ca9f, __obf_f304d4a353574273 := json.Marshal(__obf_c5f38b6ffd40036c)
	if __obf_f304d4a353574273 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_f304d4a353574273)
	}
	__obf_ee6219419a712a37, __obf_f304d4a353574273 := http.NewRequest(http.MethodPost, __obf_bd203f615f84affd, bytes.NewBuffer(__obf_906ba8325595ca9f))
	if __obf_f304d4a353574273 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_f304d4a353574273)
	}
	__obf_ee6219419a712a37.
		Header.Set("Content-Type", "application/json")
	__obf_ac69f99159f4f61e, __obf_f304d4a353574273 := __obf_c072161b30da1f19.__obf_8d976fc7dd8ff145.Do(__obf_ee6219419a712a37)
	if __obf_f304d4a353574273 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_f304d4a353574273)
	}
	defer __obf_ac69f99159f4f61e.Body.Close()

	if __obf_ac69f99159f4f61e.StatusCode != http.StatusOK {
		__obf_93b6c9ccae25de12, _ := io.ReadAll(__obf_ac69f99159f4f61e.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_ac69f99159f4f61e.StatusCode, string(__obf_93b6c9ccae25de12))
	}

	if __obf_66c10a7c7056ad07 != nil {
		if __obf_f304d4a353574273 := json.NewDecoder(__obf_ac69f99159f4f61e.Body).Decode(__obf_66c10a7c7056ad07); __obf_f304d4a353574273 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_f304d4a353574273)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_c072161b30da1f19 *HttpClient) PostXML(__obf_bd203f615f84affd string, __obf_c5f38b6ffd40036c any) ([]byte, error) {
	__obf_906ba8325595ca9f, // <- 移除 response any
		__obf_f304d4a353574273 := xml.Marshal(__obf_c5f38b6ffd40036c)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_f304d4a353574273)
	}
	__obf_ee6219419a712a37,
		// fmt.Println("签名报文：", string(reqBody))
		__obf_f304d4a353574273 := http.NewRequest(http.MethodPost, __obf_bd203f615f84affd, bytes.NewBuffer(__obf_906ba8325595ca9f))
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_f304d4a353574273)
	}
	__obf_ee6219419a712a37.
		Header.Set("Content-Type", "application/xml")
	__obf_ac69f99159f4f61e, __obf_f304d4a353574273 := __obf_c072161b30da1f19.__obf_8d976fc7dd8ff145.Do(__obf_ee6219419a712a37)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_f304d4a353574273)
	}
	defer __obf_ac69f99159f4f61e.Body.Close()
	__obf_93b6c9ccae25de12, __obf_f304d4a353574273 := io.ReadAll(__obf_ac69f99159f4f61e.Body) // <- 读取原始响应体
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_f304d4a353574273)
	}

	if __obf_ac69f99159f4f61e.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_ac69f99159f4f61e.StatusCode, string(__obf_93b6c9ccae25de12))
	}

	return __obf_93b6c9ccae25de12, nil // <- 返回原始响应体
}
