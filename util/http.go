package __obf_b81118ac905f398e

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
func SetTimeOut(__obf_1bd7c4bbab843f7b time.Duration) {
	TimeOut = __obf_1bd7c4bbab843f7b
}

// httpClient() 带超时的http.Client
func __obf_f7340d2bb8e3f25d() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_eba3dc4d8ebac647 string, __obf_194e82db8912d873 any) error {
	__obf_8e8b497458f714e7, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Get(__obf_eba3dc4d8ebac647)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_8e8b497458f714e7.Body.Close()
	return __obf_da4a17fce5fcc04c.NewDecoder(__obf_8e8b497458f714e7.Body).Decode(__obf_194e82db8912d873)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_eba3dc4d8ebac647 string, __obf_194e82db8912d873 any) error {
	__obf_8e8b497458f714e7, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Get(__obf_eba3dc4d8ebac647)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_8e8b497458f714e7.Body.Close()
	return xml.NewDecoder(__obf_8e8b497458f714e7.Body).Decode(__obf_194e82db8912d873)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_eba3dc4d8ebac647 string) ([]byte, error) {
	__obf_f67d838f59f91caf, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Get(__obf_eba3dc4d8ebac647)
	if __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	defer __obf_f67d838f59f91caf.Body.Close()

	if __obf_f67d838f59f91caf.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_eba3dc4d8ebac647, __obf_f67d838f59f91caf.StatusCode)
	}
	return io.ReadAll(__obf_f67d838f59f91caf.Body)
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
func PostJson(__obf_eba3dc4d8ebac647 string, __obf_b7c2ce33c3dfc411 any) ([]byte, error) {
	__obf_c3f1368b9a704e64 := new(bytes.Buffer)
	__obf_6bc2c6fa5d69c8f4 := __obf_da4a17fce5fcc04c.NewEncoder(__obf_c3f1368b9a704e64)
	__obf_6bc2c6fa5d69c8f4.
		SetEscapeHTML(false)
	__obf_65d3bb18bf1bd15f := __obf_6bc2c6fa5d69c8f4.Encode(__obf_b7c2ce33c3dfc411)
	if __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	__obf_f67d838f59f91caf, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Post(__obf_eba3dc4d8ebac647, "application/json;charset=utf-8", __obf_c3f1368b9a704e64)
	if __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	defer __obf_f67d838f59f91caf.Body.Close()

	if __obf_f67d838f59f91caf.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_eba3dc4d8ebac647, __obf_f67d838f59f91caf.StatusCode)
	}
	return io.ReadAll(__obf_f67d838f59f91caf.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_eba3dc4d8ebac647 string, __obf_b7c2ce33c3dfc411 any, __obf_e8db7e51ba1ec27d any, __obf_58ca2e096b61e34e ...string) (__obf_65d3bb18bf1bd15f error) {
	__obf_c3f1368b9a704e64 := new(bytes.Buffer)
	__obf_6bc2c6fa5d69c8f4 := __obf_da4a17fce5fcc04c.NewEncoder(__obf_c3f1368b9a704e64)
	__obf_65d3bb18bf1bd15f = //	enc.SetEscapeHTML(false)
		__obf_6bc2c6fa5d69c8f4.Encode(__obf_b7c2ce33c3dfc411)
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	__obf_8e6acc99c65ab7da := "application/json;charset=utf-8"
	if len(__obf_58ca2e096b61e34e) > 0 {
		__obf_8e6acc99c65ab7da = strings.Join(__obf_58ca2e096b61e34e, ";")
	}
	__obf_f67d838f59f91caf,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Post(__obf_eba3dc4d8ebac647, __obf_8e6acc99c65ab7da, __obf_c3f1368b9a704e64)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_f67d838f59f91caf.Body.Close()

	if __obf_f67d838f59f91caf.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_eba3dc4d8ebac647, __obf_f67d838f59f91caf.StatusCode)
	}
	return __obf_da4a17fce5fcc04c.NewDecoder(__obf_f67d838f59f91caf.Body).Decode(__obf_e8db7e51ba1ec27d)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_eba3dc4d8ebac647 string, __obf_b7c2ce33c3dfc411 any, __obf_e8db7e51ba1ec27d any) (__obf_65d3bb18bf1bd15f error) {
	__obf_c3f1368b9a704e64 := new(bytes.Buffer)
	__obf_6bc2c6fa5d69c8f4 := xml.NewEncoder(__obf_c3f1368b9a704e64)
	__obf_65d3bb18bf1bd15f = //	enc.SetEscapeHTML(false)
		__obf_6bc2c6fa5d69c8f4.Encode(__obf_b7c2ce33c3dfc411)
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	__obf_f67d838f59f91caf, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Post(__obf_eba3dc4d8ebac647, "application/xml;charset=utf-8", __obf_c3f1368b9a704e64)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_f67d838f59f91caf.Body.Close()

	if __obf_f67d838f59f91caf.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_eba3dc4d8ebac647, __obf_f67d838f59f91caf.StatusCode)
	}
	return xml.NewDecoder(__obf_f67d838f59f91caf.Body).Decode(__obf_e8db7e51ba1ec27d)
}

// PostFile 上传文件
func PostFile(__obf_5292dd6992f7e91b, __obf_6427ba177c349111, __obf_eba3dc4d8ebac647 string) ([]byte, error) {
	__obf_1dfa75fb0e530ca7 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_5292dd6992f7e91b,
			Filename:  __obf_6427ba177c349111,
		},
	}
	return PostMultipartForm(__obf_1dfa75fb0e530ca7,

		// GetFile 下载文件
		__obf_eba3dc4d8ebac647)
}

func GetFile(__obf_6427ba177c349111, __obf_eba3dc4d8ebac647 string) error {
	__obf_f67d838f59f91caf, __obf_65d3bb18bf1bd15f := __obf_f7340d2bb8e3f25d().Get(__obf_eba3dc4d8ebac647)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_f67d838f59f91caf.Body.Close()
	__obf_1544c103707bb2be, __obf_65d3bb18bf1bd15f := os.Create(__obf_6427ba177c349111)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_1544c103707bb2be.Close()
	_, __obf_65d3bb18bf1bd15f = io.Copy(__obf_1544c103707bb2be, __obf_f67d838f59f91caf.Body)
	return __obf_65d3bb18bf1bd15f
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_1dfa75fb0e530ca7 []MultipartFormField, __obf_eba3dc4d8ebac647 string) (__obf_190e1c6c8df71df0 []byte, __obf_65d3bb18bf1bd15f error) {
	__obf_de740ccc963fa366 := &bytes.Buffer{}
	__obf_45a7d6703a5355cc := multipart.NewWriter(__obf_de740ccc963fa366)

	for _, __obf_0d6064cac7db4dd7 := range __obf_1dfa75fb0e530ca7 {
		if __obf_0d6064cac7db4dd7.IsFile {
			__obf_6a238df77bc2daf6, __obf_c9c59dd2cdb7bc0f := __obf_45a7d6703a5355cc.CreateFormFile(__obf_0d6064cac7db4dd7.Fieldname, filepath.Base(__obf_0d6064cac7db4dd7.Filename))
			if __obf_c9c59dd2cdb7bc0f != nil {
				__obf_65d3bb18bf1bd15f = fmt.Errorf("error writing to buffer , err=%v", __obf_c9c59dd2cdb7bc0f)
				return
			}
			__obf_e9b4ae600466cf48, __obf_c9c59dd2cdb7bc0f := os.Open(__obf_0d6064cac7db4dd7.Filename)
			if __obf_c9c59dd2cdb7bc0f != nil {
				__obf_65d3bb18bf1bd15f = fmt.Errorf("error opening file , err=%v", __obf_c9c59dd2cdb7bc0f)
				return
			}
			defer __obf_e9b4ae600466cf48.Close()

			if _, __obf_65d3bb18bf1bd15f = io.Copy(__obf_6a238df77bc2daf6, __obf_e9b4ae600466cf48); __obf_65d3bb18bf1bd15f != nil {
				return
			}
		} else {
			__obf_41c522de72e0bbab, __obf_c9c59dd2cdb7bc0f := __obf_45a7d6703a5355cc.CreateFormField(__obf_0d6064cac7db4dd7.Fieldname)
			if __obf_c9c59dd2cdb7bc0f != nil {
				__obf_65d3bb18bf1bd15f = __obf_c9c59dd2cdb7bc0f
				return
			}
			__obf_afddd43de166394b := bytes.NewReader(__obf_0d6064cac7db4dd7.Value)
			if _, __obf_65d3bb18bf1bd15f = io.Copy(__obf_41c522de72e0bbab, __obf_afddd43de166394b); __obf_65d3bb18bf1bd15f != nil {
				return
			}
		}
	}
	__obf_58ca2e096b61e34e := __obf_45a7d6703a5355cc.FormDataContentType()
	__obf_45a7d6703a5355cc.
		Close()
	__obf_f67d838f59f91caf, __obf_c9c59dd2cdb7bc0f := __obf_f7340d2bb8e3f25d().Post(__obf_eba3dc4d8ebac647, __obf_58ca2e096b61e34e, __obf_de740ccc963fa366)
	if __obf_c9c59dd2cdb7bc0f != nil {
		__obf_65d3bb18bf1bd15f = __obf_c9c59dd2cdb7bc0f
		return
	}
	defer __obf_f67d838f59f91caf.Body.Close()

	if __obf_f67d838f59f91caf.StatusCode != http.StatusOK {
		return nil, __obf_65d3bb18bf1bd15f
	}
	return io.ReadAll(__obf_f67d838f59f91caf.Body)
}
