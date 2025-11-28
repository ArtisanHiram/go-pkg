package __obf_d984cff8712b1ee6

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
func SetTimeOut(__obf_968eb1e1c0ee776f time.Duration) {
	TimeOut = __obf_968eb1e1c0ee776f
}

// httpClient() 带超时的http.Client
func __obf_21a9a780cde04c30() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_3763ebebd613ad78 string, __obf_5ca54f845825fcc4 any) error {

	__obf_d7f59788250a1a1e, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Get(__obf_3763ebebd613ad78)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_d7f59788250a1a1e.Body.Close()
	return __obf_4f9564fbc13d6c68.NewDecoder(__obf_d7f59788250a1a1e.Body).Decode(__obf_5ca54f845825fcc4)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_3763ebebd613ad78 string, __obf_5ca54f845825fcc4 any) error {
	__obf_d7f59788250a1a1e, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Get(__obf_3763ebebd613ad78)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_d7f59788250a1a1e.Body.Close()
	return xml.NewDecoder(__obf_d7f59788250a1a1e.Body).Decode(__obf_5ca54f845825fcc4)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_3763ebebd613ad78 string) ([]byte, error) {
	__obf_db0a0fb3f35da2af, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Get(__obf_3763ebebd613ad78)
	if __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()

	if __obf_db0a0fb3f35da2af.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_3763ebebd613ad78, __obf_db0a0fb3f35da2af.StatusCode)
	}
	return io.ReadAll(__obf_db0a0fb3f35da2af.Body)
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
func PostJson(__obf_3763ebebd613ad78 string, __obf_73f879cfe33c419d any) ([]byte, error) {
	__obf_f2c34ab8aba39a73 := new(bytes.Buffer)
	__obf_7583e34f0c7e9350 := __obf_4f9564fbc13d6c68.NewEncoder(__obf_f2c34ab8aba39a73)
	__obf_7583e34f0c7e9350.SetEscapeHTML(false)
	__obf_44ebf351b5f3fef8 := __obf_7583e34f0c7e9350.Encode(__obf_73f879cfe33c419d)
	if __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}
	__obf_db0a0fb3f35da2af, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Post(__obf_3763ebebd613ad78, "application/json;charset=utf-8", __obf_f2c34ab8aba39a73)
	if __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()

	if __obf_db0a0fb3f35da2af.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_3763ebebd613ad78, __obf_db0a0fb3f35da2af.StatusCode)
	}
	return io.ReadAll(__obf_db0a0fb3f35da2af.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_3763ebebd613ad78 string, __obf_73f879cfe33c419d any, __obf_21e1d580f0acd4d1 any, __obf_0a58cee3165ac025 ...string) (__obf_44ebf351b5f3fef8 error) {
	__obf_f2c34ab8aba39a73 := new(bytes.Buffer)
	__obf_7583e34f0c7e9350 := __obf_4f9564fbc13d6c68.NewEncoder(__obf_f2c34ab8aba39a73)
	//	enc.SetEscapeHTML(false)
	__obf_44ebf351b5f3fef8 = __obf_7583e34f0c7e9350.Encode(__obf_73f879cfe33c419d)
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	__obf_57c8ec0445de48f1 := "application/json;charset=utf-8"
	if len(__obf_0a58cee3165ac025) > 0 {
		__obf_57c8ec0445de48f1 = strings.Join(__obf_0a58cee3165ac025, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_db0a0fb3f35da2af, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Post(__obf_3763ebebd613ad78, __obf_57c8ec0445de48f1, __obf_f2c34ab8aba39a73)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()

	if __obf_db0a0fb3f35da2af.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_3763ebebd613ad78, __obf_db0a0fb3f35da2af.StatusCode)
	}
	return __obf_4f9564fbc13d6c68.NewDecoder(__obf_db0a0fb3f35da2af.Body).Decode(__obf_21e1d580f0acd4d1)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_3763ebebd613ad78 string, __obf_73f879cfe33c419d any, __obf_21e1d580f0acd4d1 any) (__obf_44ebf351b5f3fef8 error) {
	__obf_f2c34ab8aba39a73 := new(bytes.Buffer)
	__obf_7583e34f0c7e9350 := xml.NewEncoder(__obf_f2c34ab8aba39a73)
	//	enc.SetEscapeHTML(false)
	__obf_44ebf351b5f3fef8 = __obf_7583e34f0c7e9350.Encode(__obf_73f879cfe33c419d)
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}

	__obf_db0a0fb3f35da2af, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Post(__obf_3763ebebd613ad78, "application/xml;charset=utf-8", __obf_f2c34ab8aba39a73)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()

	if __obf_db0a0fb3f35da2af.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_3763ebebd613ad78, __obf_db0a0fb3f35da2af.StatusCode)
	}
	return xml.NewDecoder(__obf_db0a0fb3f35da2af.Body).Decode(__obf_21e1d580f0acd4d1)
}

// PostFile 上传文件
func PostFile(__obf_2376fc683a6fcf97, __obf_31babcd135b23047, __obf_3763ebebd613ad78 string) ([]byte, error) {
	__obf_a95601c8435db021 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_2376fc683a6fcf97,
			Filename:  __obf_31babcd135b23047,
		},
	}
	return PostMultipartForm(__obf_a95601c8435db021, __obf_3763ebebd613ad78)
}

// GetFile 下载文件
func GetFile(__obf_31babcd135b23047, __obf_3763ebebd613ad78 string) error {
	__obf_db0a0fb3f35da2af, __obf_44ebf351b5f3fef8 := __obf_21a9a780cde04c30().Get(__obf_3763ebebd613ad78)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()
	__obf_c813cd1cdfeaed75, __obf_44ebf351b5f3fef8 := os.Create(__obf_31babcd135b23047)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_c813cd1cdfeaed75.Close()
	_, __obf_44ebf351b5f3fef8 = io.Copy(__obf_c813cd1cdfeaed75, __obf_db0a0fb3f35da2af.Body)
	return __obf_44ebf351b5f3fef8
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_a95601c8435db021 []MultipartFormField, __obf_3763ebebd613ad78 string) (__obf_b4082d07c3bc8dfb []byte, __obf_44ebf351b5f3fef8 error) {
	__obf_308575a6d6da703c := &bytes.Buffer{}
	__obf_b6d4a5f64247e47a := multipart.NewWriter(__obf_308575a6d6da703c)

	for _, __obf_bb033809a9bc2975 := range __obf_a95601c8435db021 {
		if __obf_bb033809a9bc2975.IsFile {
			__obf_2c517430022f314d, __obf_1a08474ebb9e961a := __obf_b6d4a5f64247e47a.CreateFormFile(__obf_bb033809a9bc2975.Fieldname, filepath.Base(__obf_bb033809a9bc2975.Filename))
			if __obf_1a08474ebb9e961a != nil {
				__obf_44ebf351b5f3fef8 = fmt.Errorf("error writing to buffer , err=%v", __obf_1a08474ebb9e961a)
				return
			}

			__obf_eb4a741cdc2b5a47, __obf_1a08474ebb9e961a := os.Open(__obf_bb033809a9bc2975.Filename)
			if __obf_1a08474ebb9e961a != nil {
				__obf_44ebf351b5f3fef8 = fmt.Errorf("error opening file , err=%v", __obf_1a08474ebb9e961a)
				return
			}
			defer __obf_eb4a741cdc2b5a47.Close()

			if _, __obf_44ebf351b5f3fef8 = io.Copy(__obf_2c517430022f314d, __obf_eb4a741cdc2b5a47); __obf_44ebf351b5f3fef8 != nil {
				return
			}
		} else {
			__obf_32fa93eddf6b8a37, __obf_1a08474ebb9e961a := __obf_b6d4a5f64247e47a.CreateFormField(__obf_bb033809a9bc2975.Fieldname)
			if __obf_1a08474ebb9e961a != nil {
				__obf_44ebf351b5f3fef8 = __obf_1a08474ebb9e961a
				return
			}
			__obf_8d67b742b1eb0d65 := bytes.NewReader(__obf_bb033809a9bc2975.Value)
			if _, __obf_44ebf351b5f3fef8 = io.Copy(__obf_32fa93eddf6b8a37, __obf_8d67b742b1eb0d65); __obf_44ebf351b5f3fef8 != nil {
				return
			}
		}
	}

	__obf_0a58cee3165ac025 := __obf_b6d4a5f64247e47a.FormDataContentType()
	__obf_b6d4a5f64247e47a.Close()

	__obf_db0a0fb3f35da2af, __obf_1a08474ebb9e961a := __obf_21a9a780cde04c30().Post(__obf_3763ebebd613ad78, __obf_0a58cee3165ac025, __obf_308575a6d6da703c)
	if __obf_1a08474ebb9e961a != nil {
		__obf_44ebf351b5f3fef8 = __obf_1a08474ebb9e961a
		return
	}
	defer __obf_db0a0fb3f35da2af.Body.Close()

	if __obf_db0a0fb3f35da2af.StatusCode != http.StatusOK {
		return nil, __obf_44ebf351b5f3fef8
	}
	return io.ReadAll(__obf_db0a0fb3f35da2af.Body)
}
