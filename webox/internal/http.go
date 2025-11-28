package __obf_6a74725dc3694218

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
	__obf_7095a64d2d7c7e63 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{
		__obf_7095a64d2d7c7e63: &http.Client{
			Timeout: 15 * time.Second, // 设置一个合理的超时时间
		},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_9bf846c82da0ce93 tls.Certificate) (*HttpClient, error) {
	// 加载客户端证书和私钥
	// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
	// }

	// 构造TLS配置
	__obf_d2c9382e5ea95b66 := &tls.Config{
		Certificates: []tls.Certificate{__obf_9bf846c82da0ce93},
		// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
		// RootCAs:            nil, // 如果有自定义CA，可以设置
	}

	// 创建带证书的HTTP Transport
	__obf_6e6f34cd3e50047f := &http.Transport{
		TLSClientConfig:       __obf_d2c9382e5ea95b66,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &HttpClient{
		__obf_7095a64d2d7c7e63: &http.Client{
			Transport: __obf_6e6f34cd3e50047f,
			Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
		},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_d3e567063d420fd6 *HttpClient) Get(__obf_286899246539055b string, __obf_7a1adb5e66219327 map[string]string, __obf_b7cf6d719df113c6 any) error {
	__obf_af3f09d1476e7e1e, __obf_b5433c9ceabad42d := http.NewRequest(http.MethodGet, __obf_286899246539055b, nil)
	if __obf_b5433c9ceabad42d != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_b5433c9ceabad42d)
	}

	__obf_00aebead1508d594 := __obf_af3f09d1476e7e1e.URL.Query()
	for __obf_5014165c4e4961f3, __obf_2f5f1559bbfb1282 := range __obf_7a1adb5e66219327 {
		__obf_00aebead1508d594.Add(__obf_5014165c4e4961f3, __obf_2f5f1559bbfb1282)
	}
	__obf_af3f09d1476e7e1e.URL.RawQuery = __obf_00aebead1508d594.Encode()

	__obf_135fd88d69beb7f7, __obf_b5433c9ceabad42d := __obf_d3e567063d420fd6.__obf_7095a64d2d7c7e63.Do(__obf_af3f09d1476e7e1e)
	if __obf_b5433c9ceabad42d != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_b5433c9ceabad42d)
	}
	defer __obf_135fd88d69beb7f7.Body.Close()

	if __obf_135fd88d69beb7f7.StatusCode != http.StatusOK {
		__obf_4dd36dcd1427c8c1, _ := io.ReadAll(__obf_135fd88d69beb7f7.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_135fd88d69beb7f7.StatusCode, string(__obf_4dd36dcd1427c8c1))
	}

	if __obf_b7cf6d719df113c6 != nil {
		if __obf_b5433c9ceabad42d := json.NewDecoder(__obf_135fd88d69beb7f7.Body).Decode(__obf_b7cf6d719df113c6); __obf_b5433c9ceabad42d != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_b5433c9ceabad42d)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_d3e567063d420fd6 *HttpClient) PostJSON(__obf_286899246539055b string, __obf_6029ef0afd39bbd0 any, __obf_b7cf6d719df113c6 any) error {
	__obf_5f69bf2bf7c9eceb, __obf_b5433c9ceabad42d := json.Marshal(__obf_6029ef0afd39bbd0)
	if __obf_b5433c9ceabad42d != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_b5433c9ceabad42d)
	}

	__obf_af3f09d1476e7e1e, __obf_b5433c9ceabad42d := http.NewRequest(http.MethodPost, __obf_286899246539055b, bytes.NewBuffer(__obf_5f69bf2bf7c9eceb))
	if __obf_b5433c9ceabad42d != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_b5433c9ceabad42d)
	}
	__obf_af3f09d1476e7e1e.Header.Set("Content-Type", "application/json")

	__obf_135fd88d69beb7f7, __obf_b5433c9ceabad42d := __obf_d3e567063d420fd6.__obf_7095a64d2d7c7e63.Do(__obf_af3f09d1476e7e1e)
	if __obf_b5433c9ceabad42d != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_b5433c9ceabad42d)
	}
	defer __obf_135fd88d69beb7f7.Body.Close()

	if __obf_135fd88d69beb7f7.StatusCode != http.StatusOK {
		__obf_4dd36dcd1427c8c1, _ := io.ReadAll(__obf_135fd88d69beb7f7.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_135fd88d69beb7f7.StatusCode, string(__obf_4dd36dcd1427c8c1))
	}

	if __obf_b7cf6d719df113c6 != nil {
		if __obf_b5433c9ceabad42d := json.NewDecoder(__obf_135fd88d69beb7f7.Body).Decode(__obf_b7cf6d719df113c6); __obf_b5433c9ceabad42d != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_b5433c9ceabad42d)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_d3e567063d420fd6 *HttpClient) PostXML(__obf_286899246539055b string, __obf_6029ef0afd39bbd0 any) ([]byte, error) { // <- 移除 response any
	__obf_5f69bf2bf7c9eceb, __obf_b5433c9ceabad42d := xml.Marshal(__obf_6029ef0afd39bbd0)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_b5433c9ceabad42d)
	}
	// fmt.Println("签名报文：", string(reqBody))

	__obf_af3f09d1476e7e1e, __obf_b5433c9ceabad42d := http.NewRequest(http.MethodPost, __obf_286899246539055b, bytes.NewBuffer(__obf_5f69bf2bf7c9eceb))
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_b5433c9ceabad42d)
	}
	__obf_af3f09d1476e7e1e.Header.Set("Content-Type", "application/xml")

	__obf_135fd88d69beb7f7, __obf_b5433c9ceabad42d := __obf_d3e567063d420fd6.__obf_7095a64d2d7c7e63.Do(__obf_af3f09d1476e7e1e)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_b5433c9ceabad42d)
	}
	defer __obf_135fd88d69beb7f7.Body.Close()

	__obf_4dd36dcd1427c8c1, __obf_b5433c9ceabad42d := io.ReadAll(__obf_135fd88d69beb7f7.Body) // <- 读取原始响应体
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_b5433c9ceabad42d)
	}

	if __obf_135fd88d69beb7f7.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_135fd88d69beb7f7.StatusCode, string(__obf_4dd36dcd1427c8c1))
	}

	return __obf_4dd36dcd1427c8c1, nil // <- 返回原始响应体
}
