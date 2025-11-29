package __obf_d7b39e56b82f7f57

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
func SetTimeOut(__obf_fff126686fc07d47 time.Duration) {
	TimeOut = __obf_fff126686fc07d47
}

// httpClient() 带超时的http.Client
func __obf_de7cfb7c399df8dc() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_374c3906c456f024 string, __obf_a30d33f1a4aa6e72 any) error {
	__obf_403793979a29000e, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Get(__obf_374c3906c456f024)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_403793979a29000e.Body.Close()
	return __obf_6e1ddda33ce54ba3.NewDecoder(__obf_403793979a29000e.Body).Decode(__obf_a30d33f1a4aa6e72)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_374c3906c456f024 string, __obf_a30d33f1a4aa6e72 any) error {
	__obf_403793979a29000e, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Get(__obf_374c3906c456f024)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_403793979a29000e.Body.Close()
	return xml.NewDecoder(__obf_403793979a29000e.Body).Decode(__obf_a30d33f1a4aa6e72)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_374c3906c456f024 string) ([]byte, error) {
	__obf_9e4dcfcdf0824ecf, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Get(__obf_374c3906c456f024)
	if __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()

	if __obf_9e4dcfcdf0824ecf.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_374c3906c456f024, __obf_9e4dcfcdf0824ecf.StatusCode)
	}
	return io.ReadAll(__obf_9e4dcfcdf0824ecf.Body)
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
func PostJson(__obf_374c3906c456f024 string, __obf_9b2baf23d79a3ad2 any) ([]byte, error) {
	__obf_0f9b3297f42db961 := new(bytes.Buffer)
	__obf_9614fb1e809fb48c := __obf_6e1ddda33ce54ba3.NewEncoder(__obf_0f9b3297f42db961)
	__obf_9614fb1e809fb48c.
		SetEscapeHTML(false)
	__obf_3ccf8c32165eeb13 := __obf_9614fb1e809fb48c.Encode(__obf_9b2baf23d79a3ad2)
	if __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	__obf_9e4dcfcdf0824ecf, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Post(__obf_374c3906c456f024, "application/json;charset=utf-8", __obf_0f9b3297f42db961)
	if __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()

	if __obf_9e4dcfcdf0824ecf.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_374c3906c456f024, __obf_9e4dcfcdf0824ecf.StatusCode)
	}
	return io.ReadAll(__obf_9e4dcfcdf0824ecf.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_374c3906c456f024 string, __obf_9b2baf23d79a3ad2 any, __obf_f5b0384ddb3578f7 any, __obf_fcf1cca9ce33ae82 ...string) (__obf_3ccf8c32165eeb13 error) {
	__obf_0f9b3297f42db961 := new(bytes.Buffer)
	__obf_9614fb1e809fb48c := __obf_6e1ddda33ce54ba3.NewEncoder(__obf_0f9b3297f42db961)
	__obf_3ccf8c32165eeb13 = //	enc.SetEscapeHTML(false)
		__obf_9614fb1e809fb48c.Encode(__obf_9b2baf23d79a3ad2)
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	__obf_bd4cc161c2216d4d := "application/json;charset=utf-8"
	if len(__obf_fcf1cca9ce33ae82) > 0 {
		__obf_bd4cc161c2216d4d = strings.Join(__obf_fcf1cca9ce33ae82, ";")
	}
	__obf_9e4dcfcdf0824ecf,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Post(__obf_374c3906c456f024, __obf_bd4cc161c2216d4d, __obf_0f9b3297f42db961)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()

	if __obf_9e4dcfcdf0824ecf.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_374c3906c456f024, __obf_9e4dcfcdf0824ecf.StatusCode)
	}
	return __obf_6e1ddda33ce54ba3.NewDecoder(__obf_9e4dcfcdf0824ecf.Body).Decode(__obf_f5b0384ddb3578f7)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_374c3906c456f024 string, __obf_9b2baf23d79a3ad2 any, __obf_f5b0384ddb3578f7 any) (__obf_3ccf8c32165eeb13 error) {
	__obf_0f9b3297f42db961 := new(bytes.Buffer)
	__obf_9614fb1e809fb48c := xml.NewEncoder(__obf_0f9b3297f42db961)
	__obf_3ccf8c32165eeb13 = //	enc.SetEscapeHTML(false)
		__obf_9614fb1e809fb48c.Encode(__obf_9b2baf23d79a3ad2)
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	__obf_9e4dcfcdf0824ecf, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Post(__obf_374c3906c456f024, "application/xml;charset=utf-8", __obf_0f9b3297f42db961)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()

	if __obf_9e4dcfcdf0824ecf.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_374c3906c456f024, __obf_9e4dcfcdf0824ecf.StatusCode)
	}
	return xml.NewDecoder(__obf_9e4dcfcdf0824ecf.Body).Decode(__obf_f5b0384ddb3578f7)
}

// PostFile 上传文件
func PostFile(__obf_0ab326b395497cf9, __obf_5ad072acfd05203c, __obf_374c3906c456f024 string) ([]byte, error) {
	__obf_7094740f02efb4b0 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_0ab326b395497cf9,
			Filename:  __obf_5ad072acfd05203c,
		},
	}
	return PostMultipartForm(__obf_7094740f02efb4b0,

		// GetFile 下载文件
		__obf_374c3906c456f024)
}

func GetFile(__obf_5ad072acfd05203c, __obf_374c3906c456f024 string) error {
	__obf_9e4dcfcdf0824ecf, __obf_3ccf8c32165eeb13 := __obf_de7cfb7c399df8dc().Get(__obf_374c3906c456f024)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()
	__obf_61843cfe4fcba7bf, __obf_3ccf8c32165eeb13 := os.Create(__obf_5ad072acfd05203c)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_61843cfe4fcba7bf.Close()
	_, __obf_3ccf8c32165eeb13 = io.Copy(__obf_61843cfe4fcba7bf, __obf_9e4dcfcdf0824ecf.Body)
	return __obf_3ccf8c32165eeb13
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_7094740f02efb4b0 []MultipartFormField, __obf_374c3906c456f024 string) (__obf_7d51ead17926b0a9 []byte, __obf_3ccf8c32165eeb13 error) {
	__obf_d636cb4e4e24e3a2 := &bytes.Buffer{}
	__obf_6631ced141e7277a := multipart.NewWriter(__obf_d636cb4e4e24e3a2)

	for _, __obf_b3c735e9ab0f55ed := range __obf_7094740f02efb4b0 {
		if __obf_b3c735e9ab0f55ed.IsFile {
			__obf_ef62cfb565e98408, __obf_36b2df5f73ce8ef9 := __obf_6631ced141e7277a.CreateFormFile(__obf_b3c735e9ab0f55ed.Fieldname, filepath.Base(__obf_b3c735e9ab0f55ed.Filename))
			if __obf_36b2df5f73ce8ef9 != nil {
				__obf_3ccf8c32165eeb13 = fmt.Errorf("error writing to buffer , err=%v", __obf_36b2df5f73ce8ef9)
				return
			}
			__obf_0135f3de1506ae9a, __obf_36b2df5f73ce8ef9 := os.Open(__obf_b3c735e9ab0f55ed.Filename)
			if __obf_36b2df5f73ce8ef9 != nil {
				__obf_3ccf8c32165eeb13 = fmt.Errorf("error opening file , err=%v", __obf_36b2df5f73ce8ef9)
				return
			}
			defer __obf_0135f3de1506ae9a.Close()

			if _, __obf_3ccf8c32165eeb13 = io.Copy(__obf_ef62cfb565e98408, __obf_0135f3de1506ae9a); __obf_3ccf8c32165eeb13 != nil {
				return
			}
		} else {
			__obf_54243131ca1f2d61, __obf_36b2df5f73ce8ef9 := __obf_6631ced141e7277a.CreateFormField(__obf_b3c735e9ab0f55ed.Fieldname)
			if __obf_36b2df5f73ce8ef9 != nil {
				__obf_3ccf8c32165eeb13 = __obf_36b2df5f73ce8ef9
				return
			}
			__obf_a3c3e085277c66a0 := bytes.NewReader(__obf_b3c735e9ab0f55ed.Value)
			if _, __obf_3ccf8c32165eeb13 = io.Copy(__obf_54243131ca1f2d61, __obf_a3c3e085277c66a0); __obf_3ccf8c32165eeb13 != nil {
				return
			}
		}
	}
	__obf_fcf1cca9ce33ae82 := __obf_6631ced141e7277a.FormDataContentType()
	__obf_6631ced141e7277a.
		Close()
	__obf_9e4dcfcdf0824ecf, __obf_36b2df5f73ce8ef9 := __obf_de7cfb7c399df8dc().Post(__obf_374c3906c456f024, __obf_fcf1cca9ce33ae82, __obf_d636cb4e4e24e3a2)
	if __obf_36b2df5f73ce8ef9 != nil {
		__obf_3ccf8c32165eeb13 = __obf_36b2df5f73ce8ef9
		return
	}
	defer __obf_9e4dcfcdf0824ecf.Body.Close()

	if __obf_9e4dcfcdf0824ecf.StatusCode != http.StatusOK {
		return nil, __obf_3ccf8c32165eeb13
	}
	return io.ReadAll(__obf_9e4dcfcdf0824ecf.Body)
}
