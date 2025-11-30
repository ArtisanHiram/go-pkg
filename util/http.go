package __obf_69af70389b6622a3

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// TimeOut 全局请求超时设置,默认1分钟
var TimeOut time.Duration = 60 * time.Second

// SetTimeOut 设置全局请求超时
func SetTimeOut(__obf_3cbd1ac0144c98a4 time.Duration) {
	TimeOut = __obf_3cbd1ac0144c98a4
}

// httpClient() 带超时的http.Client
func __obf_6ff21559e277adbc() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_e5f6a96b3d66cfb3 string, __obf_b1dd41bcd3423053 any) error {
	__obf_675dd561f1bbbbff, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Get(__obf_e5f6a96b3d66cfb3)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_675dd561f1bbbbff.Body.Close()
	return __obf_bd9eacaebdd04f5a.NewDecoder(__obf_675dd561f1bbbbff.Body).Decode(__obf_b1dd41bcd3423053)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_e5f6a96b3d66cfb3 string, __obf_b1dd41bcd3423053 any) error {
	__obf_675dd561f1bbbbff, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Get(__obf_e5f6a96b3d66cfb3)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_675dd561f1bbbbff.Body.Close()
	return xml.NewDecoder(__obf_675dd561f1bbbbff.Body).Decode(__obf_b1dd41bcd3423053)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_e5f6a96b3d66cfb3 string) ([]byte, error) {
	__obf_5a66cb0bf6c2afb0, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Get(__obf_e5f6a96b3d66cfb3)
	if __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()

	if __obf_5a66cb0bf6c2afb0.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_e5f6a96b3d66cfb3, __obf_5a66cb0bf6c2afb0.StatusCode)
	}
	return io.ReadAll(__obf_5a66cb0bf6c2afb0.Body)
}

// GetRawBody 发送GET请求，返回body字节
// func GetRawBody(uri string) (io.ReadCloser, error) {
// 	resp, err := httpClient().Get(uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", uri, resp.StatusCode)
// 	}
// 	return resp.Body, nil
// }

// PostJson 发送Json格式的POST请求
func PostJson(__obf_e5f6a96b3d66cfb3 string, __obf_5a4019b2f3aa4a12 any) ([]byte, error) {
	__obf_958555db3fcd9074 := new(bytes.Buffer)
	__obf_da00d5adbf65521e := __obf_bd9eacaebdd04f5a.NewEncoder(__obf_958555db3fcd9074)
	__obf_da00d5adbf65521e.
		SetEscapeHTML(false)
	__obf_e812cfd1203ed4f3 := __obf_da00d5adbf65521e.Encode(__obf_5a4019b2f3aa4a12)
	if __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	__obf_5a66cb0bf6c2afb0, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Post(__obf_e5f6a96b3d66cfb3, "application/json;charset=utf-8", __obf_958555db3fcd9074)
	if __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()

	if __obf_5a66cb0bf6c2afb0.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_e5f6a96b3d66cfb3, __obf_5a66cb0bf6c2afb0.StatusCode)
	}
	return io.ReadAll(__obf_5a66cb0bf6c2afb0.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_e5f6a96b3d66cfb3 string, __obf_5a4019b2f3aa4a12 any, __obf_f7c4b5f8a87c287c any, __obf_ba48b3d87c8e2bac ...string) (__obf_e812cfd1203ed4f3 error) {
	__obf_958555db3fcd9074 := new(bytes.Buffer)
	__obf_da00d5adbf65521e := __obf_bd9eacaebdd04f5a.NewEncoder(__obf_958555db3fcd9074)
	__obf_e812cfd1203ed4f3 = //	enc.SetEscapeHTML(false)
		__obf_da00d5adbf65521e.Encode(__obf_5a4019b2f3aa4a12)
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	__obf_16ecd6e522d6632a := "application/json;charset=utf-8"
	if len(__obf_ba48b3d87c8e2bac) > 0 {
		__obf_16ecd6e522d6632a = strings.Join(__obf_ba48b3d87c8e2bac, ";")
	}
	__obf_5a66cb0bf6c2afb0,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Post(__obf_e5f6a96b3d66cfb3, __obf_16ecd6e522d6632a, __obf_958555db3fcd9074)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()

	if __obf_5a66cb0bf6c2afb0.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_e5f6a96b3d66cfb3, __obf_5a66cb0bf6c2afb0.StatusCode)
	}
	return __obf_bd9eacaebdd04f5a.NewDecoder(__obf_5a66cb0bf6c2afb0.Body).Decode(__obf_f7c4b5f8a87c287c)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_e5f6a96b3d66cfb3 string, __obf_5a4019b2f3aa4a12 any, __obf_f7c4b5f8a87c287c any) (__obf_e812cfd1203ed4f3 error) {
	__obf_958555db3fcd9074 := new(bytes.Buffer)
	__obf_da00d5adbf65521e := xml.NewEncoder(__obf_958555db3fcd9074)
	__obf_e812cfd1203ed4f3 = //	enc.SetEscapeHTML(false)
		__obf_da00d5adbf65521e.Encode(__obf_5a4019b2f3aa4a12)
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	__obf_5a66cb0bf6c2afb0, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Post(__obf_e5f6a96b3d66cfb3, "application/xml;charset=utf-8", __obf_958555db3fcd9074)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()

	if __obf_5a66cb0bf6c2afb0.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_e5f6a96b3d66cfb3, __obf_5a66cb0bf6c2afb0.StatusCode)
	}
	return xml.NewDecoder(__obf_5a66cb0bf6c2afb0.Body).Decode(__obf_f7c4b5f8a87c287c)
}

// PostFile 上传文件
func PostFile(__obf_30e514ffa92b4154, __obf_8a7dd8c3f484bf31, __obf_e5f6a96b3d66cfb3 string) ([]byte, error) {
	__obf_6b37bdf034621f1c := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_30e514ffa92b4154,
			Filename:  __obf_8a7dd8c3f484bf31,
		},
	}
	return PostMultipartForm(__obf_6b37bdf034621f1c,

		// GetFile 下载文件
		__obf_e5f6a96b3d66cfb3)
}

func GetFile(__obf_8a7dd8c3f484bf31, __obf_e5f6a96b3d66cfb3 string) error {
	__obf_5a66cb0bf6c2afb0, __obf_e812cfd1203ed4f3 := __obf_6ff21559e277adbc().Get(__obf_e5f6a96b3d66cfb3)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()
	__obf_268050566cb23e9c, __obf_e812cfd1203ed4f3 := os.Create(__obf_8a7dd8c3f484bf31)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_268050566cb23e9c.Close()
	_, __obf_e812cfd1203ed4f3 = io.Copy(__obf_268050566cb23e9c, __obf_5a66cb0bf6c2afb0.Body)
	return __obf_e812cfd1203ed4f3
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_6b37bdf034621f1c []MultipartFormField, __obf_e5f6a96b3d66cfb3 string) (__obf_8aae42199ca3cfd0 []byte, __obf_e812cfd1203ed4f3 error) {
	__obf_b762dcf5d4d53022 := &bytes.Buffer{}
	__obf_10749b20ac94c72a := multipart.NewWriter(__obf_b762dcf5d4d53022)

	for _, __obf_5704ad6c16d70f41 := range __obf_6b37bdf034621f1c {
		if __obf_5704ad6c16d70f41.IsFile {
			__obf_0ddb873f84a42030, __obf_3ffba1b5be1ebf4c := __obf_10749b20ac94c72a.CreateFormFile(__obf_5704ad6c16d70f41.Fieldname, filepath.Base(__obf_5704ad6c16d70f41.Filename))
			if __obf_3ffba1b5be1ebf4c != nil {
				__obf_e812cfd1203ed4f3 = fmt.Errorf("error writing to buffer , err=%v", __obf_3ffba1b5be1ebf4c)
				return
			}
			__obf_d48d4d30b5d70fa4, __obf_3ffba1b5be1ebf4c := os.Open(__obf_5704ad6c16d70f41.Filename)
			if __obf_3ffba1b5be1ebf4c != nil {
				__obf_e812cfd1203ed4f3 = fmt.Errorf("error opening file , err=%v", __obf_3ffba1b5be1ebf4c)
				return
			}
			defer __obf_d48d4d30b5d70fa4.Close()

			if _, __obf_e812cfd1203ed4f3 = io.Copy(__obf_0ddb873f84a42030, __obf_d48d4d30b5d70fa4); __obf_e812cfd1203ed4f3 != nil {
				return
			}
		} else {
			__obf_e18796dcb31d626a, __obf_3ffba1b5be1ebf4c := __obf_10749b20ac94c72a.CreateFormField(__obf_5704ad6c16d70f41.Fieldname)
			if __obf_3ffba1b5be1ebf4c != nil {
				__obf_e812cfd1203ed4f3 = __obf_3ffba1b5be1ebf4c
				return
			}
			__obf_62d501287547d4e3 := bytes.NewReader(__obf_5704ad6c16d70f41.Value)
			if _, __obf_e812cfd1203ed4f3 = io.Copy(__obf_e18796dcb31d626a, __obf_62d501287547d4e3); __obf_e812cfd1203ed4f3 != nil {
				return
			}
		}
	}
	__obf_ba48b3d87c8e2bac := __obf_10749b20ac94c72a.FormDataContentType()
	__obf_10749b20ac94c72a.
		Close()
	__obf_5a66cb0bf6c2afb0, __obf_3ffba1b5be1ebf4c := __obf_6ff21559e277adbc().Post(__obf_e5f6a96b3d66cfb3, __obf_ba48b3d87c8e2bac, __obf_b762dcf5d4d53022)
	if __obf_3ffba1b5be1ebf4c != nil {
		__obf_e812cfd1203ed4f3 = __obf_3ffba1b5be1ebf4c
		return
	}
	defer __obf_5a66cb0bf6c2afb0.Body.Close()

	if __obf_5a66cb0bf6c2afb0.StatusCode != http.StatusOK {
		return nil, __obf_e812cfd1203ed4f3
	}
	return io.ReadAll(__obf_5a66cb0bf6c2afb0.Body)
}
