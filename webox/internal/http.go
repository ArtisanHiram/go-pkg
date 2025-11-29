package __obf_18f8d68b9095d0e0

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
	__obf_a0584419335be4c8 *http.Client
}

// NewClient 创建一个新的HTTP客户端实例
func NewClient() *HttpClient {
	return &HttpClient{__obf_a0584419335be4c8: &http.Client{
		Timeout: 15 * time.Second, // 设置一个合理的超时时间
	},
	}
}

// NewClientWithTls 创建一个新的HTTP客户端实例 (带证书)
// certPath: API证书文件路径 (例如：apiclient_cert.pem)
// keyPath: API私钥文件路径 (例如：apiclient_key.pem)
func NewClientWithTLS(__obf_6f61a01fceaee6c5 tls.Certificate) (*HttpClient, error) {
	__obf_7f456eee6c75fe89 := // 加载客户端证书和私钥
		// cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		// if err != nil {
		// 	return nil, fmt.Errorf("load x509 key pair failed: %w", err)
		// }

		// 构造TLS配置
		&tls.Config{
			Certificates: []tls.Certificate{__obf_6f61a01fceaee6c5},
			// InsecureSkipVerify: true, // 生产环境不要开启！仅用于调试
			// RootCAs:            nil, // 如果有自定义CA，可以设置
		}
	__obf_33e6b4ab10ac8f40 := // 创建带证书的HTTP Transport
		&http.Transport{
			TLSClientConfig:       __obf_7f456eee6c75fe89,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

	return &HttpClient{__obf_a0584419335be4c8: &http.Client{
		Transport: __obf_33e6b4ab10ac8f40,
		Timeout:   30 * time.Second, // 证书请求可能更耗时，适当延长超时时间
	},
	}, nil
}

// Get 发送GET请求，并解析JSON响应
func (__obf_8477bb51362716f1 *HttpClient) Get(__obf_6013d898345a950d string, __obf_8f23c40566a8f981 map[string]string, __obf_df6337e15b306319 any) error {
	__obf_1c7c28e84bff3a31, __obf_f45e33988e88e5a2 := http.NewRequest(http.MethodGet, __obf_6013d898345a950d, nil)
	if __obf_f45e33988e88e5a2 != nil {
		return fmt.Errorf("create GET request failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_d25223237c4455ac := __obf_1c7c28e84bff3a31.URL.Query()
	for __obf_01074596fa5e7fd7, __obf_4b612b03fe7ab91c := range __obf_8f23c40566a8f981 {
		__obf_d25223237c4455ac.
			Add(__obf_01074596fa5e7fd7, __obf_4b612b03fe7ab91c)
	}
	__obf_1c7c28e84bff3a31.
		URL.RawQuery = __obf_d25223237c4455ac.Encode()
	__obf_557fca5e631bd53f, __obf_f45e33988e88e5a2 := __obf_8477bb51362716f1.__obf_a0584419335be4c8.Do(__obf_1c7c28e84bff3a31)
	if __obf_f45e33988e88e5a2 != nil {
		return fmt.Errorf("send GET request failed: %w", __obf_f45e33988e88e5a2)
	}
	defer __obf_557fca5e631bd53f.Body.Close()

	if __obf_557fca5e631bd53f.StatusCode != http.StatusOK {
		__obf_12144a8704dec170, _ := io.ReadAll(__obf_557fca5e631bd53f.Body)
		return fmt.Errorf("GET request failed with status %d: %s", __obf_557fca5e631bd53f.StatusCode, string(__obf_12144a8704dec170))
	}

	if __obf_df6337e15b306319 != nil {
		if __obf_f45e33988e88e5a2 := json.NewDecoder(__obf_557fca5e631bd53f.Body).Decode(__obf_df6337e15b306319); __obf_f45e33988e88e5a2 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_f45e33988e88e5a2)
		}
	}
	return nil
}

// PostJSON 发送POST请求，Body为JSON，并解析JSON响应
func (__obf_8477bb51362716f1 *HttpClient) PostJSON(__obf_6013d898345a950d string, __obf_7604ecffa2539ffb any, __obf_df6337e15b306319 any) error {
	__obf_ea1dc0d061ad80f0, __obf_f45e33988e88e5a2 := json.Marshal(__obf_7604ecffa2539ffb)
	if __obf_f45e33988e88e5a2 != nil {
		return fmt.Errorf("marshal JSON request failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_1c7c28e84bff3a31, __obf_f45e33988e88e5a2 := http.NewRequest(http.MethodPost, __obf_6013d898345a950d, bytes.NewBuffer(__obf_ea1dc0d061ad80f0))
	if __obf_f45e33988e88e5a2 != nil {
		return fmt.Errorf("create POST JSON request failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_1c7c28e84bff3a31.
		Header.Set("Content-Type", "application/json")
	__obf_557fca5e631bd53f, __obf_f45e33988e88e5a2 := __obf_8477bb51362716f1.__obf_a0584419335be4c8.Do(__obf_1c7c28e84bff3a31)
	if __obf_f45e33988e88e5a2 != nil {
		return fmt.Errorf("send POST JSON request failed: %w", __obf_f45e33988e88e5a2)
	}
	defer __obf_557fca5e631bd53f.Body.Close()

	if __obf_557fca5e631bd53f.StatusCode != http.StatusOK {
		__obf_12144a8704dec170, _ := io.ReadAll(__obf_557fca5e631bd53f.Body)
		return fmt.Errorf("POST JSON request failed with status %d: %s", __obf_557fca5e631bd53f.StatusCode, string(__obf_12144a8704dec170))
	}

	if __obf_df6337e15b306319 != nil {
		if __obf_f45e33988e88e5a2 := json.NewDecoder(__obf_557fca5e631bd53f.Body).Decode(__obf_df6337e15b306319); __obf_f45e33988e88e5a2 != nil {
			return fmt.Errorf("decode JSON response failed: %w", __obf_f45e33988e88e5a2)
		}
	}
	return nil
}

// PostXML 发送POST请求，Body为XML，并返回原始XML响应体字节数组
// 调用方需要自行解析和验签
func (__obf_8477bb51362716f1 *HttpClient) PostXML(__obf_6013d898345a950d string, __obf_7604ecffa2539ffb any) ([]byte, error) {
	__obf_ea1dc0d061ad80f0, // <- 移除 response any
		__obf_f45e33988e88e5a2 := xml.Marshal(__obf_7604ecffa2539ffb)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("marshal XML request failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_1c7c28e84bff3a31,
		// fmt.Println("签名报文：", string(reqBody))
		__obf_f45e33988e88e5a2 := http.NewRequest(http.MethodPost, __obf_6013d898345a950d, bytes.NewBuffer(__obf_ea1dc0d061ad80f0))
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("create POST XML request failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_1c7c28e84bff3a31.
		Header.Set("Content-Type", "application/xml")
	__obf_557fca5e631bd53f, __obf_f45e33988e88e5a2 := __obf_8477bb51362716f1.__obf_a0584419335be4c8.Do(__obf_1c7c28e84bff3a31)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("send POST XML request failed: %w", __obf_f45e33988e88e5a2)
	}
	defer __obf_557fca5e631bd53f.Body.Close()
	__obf_12144a8704dec170, __obf_f45e33988e88e5a2 := io.ReadAll(__obf_557fca5e631bd53f.Body) // <- 读取原始响应体
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("read response body failed: %w", __obf_f45e33988e88e5a2)
	}

	if __obf_557fca5e631bd53f.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST XML request failed with status %d: %s", __obf_557fca5e631bd53f.StatusCode, string(__obf_12144a8704dec170))
	}

	return __obf_12144a8704dec170, nil // <- 返回原始响应体
}
