package __obf_85fcc7bd65d30170

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
func SetTimeOut(__obf_34214d422ca61c59 time.Duration) {
	TimeOut = __obf_34214d422ca61c59
}

// httpClient() 带超时的http.Client
func __obf_7dd7a45197327cb3() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_ae0b02887150dacd string, __obf_37a13444ba0aee54 any) error {

	__obf_4edebc331d8e8496, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Get(__obf_ae0b02887150dacd)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_4edebc331d8e8496.Body.Close()
	return __obf_9c2b993d609fe396.NewDecoder(__obf_4edebc331d8e8496.Body).Decode(__obf_37a13444ba0aee54)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_ae0b02887150dacd string, __obf_37a13444ba0aee54 any) error {
	__obf_4edebc331d8e8496, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Get(__obf_ae0b02887150dacd)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_4edebc331d8e8496.Body.Close()
	return xml.NewDecoder(__obf_4edebc331d8e8496.Body).Decode(__obf_37a13444ba0aee54)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_ae0b02887150dacd string) ([]byte, error) {
	__obf_f844075f41f45d17, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Get(__obf_ae0b02887150dacd)
	if __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}
	defer __obf_f844075f41f45d17.Body.Close()

	if __obf_f844075f41f45d17.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_ae0b02887150dacd, __obf_f844075f41f45d17.StatusCode)
	}
	return io.ReadAll(__obf_f844075f41f45d17.Body)
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
func PostJson(__obf_ae0b02887150dacd string, __obf_fe34695ffe11948d any) ([]byte, error) {
	__obf_706407c46b1ce40f := new(bytes.Buffer)
	__obf_65d9ce0e64489fa6 := __obf_9c2b993d609fe396.NewEncoder(__obf_706407c46b1ce40f)
	__obf_65d9ce0e64489fa6.SetEscapeHTML(false)
	__obf_80a78f1bfaca43d3 := __obf_65d9ce0e64489fa6.Encode(__obf_fe34695ffe11948d)
	if __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}
	__obf_f844075f41f45d17, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Post(__obf_ae0b02887150dacd, "application/json;charset=utf-8", __obf_706407c46b1ce40f)
	if __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}
	defer __obf_f844075f41f45d17.Body.Close()

	if __obf_f844075f41f45d17.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_ae0b02887150dacd, __obf_f844075f41f45d17.StatusCode)
	}
	return io.ReadAll(__obf_f844075f41f45d17.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_ae0b02887150dacd string, __obf_fe34695ffe11948d any, __obf_521605711569ff86 any, __obf_fea5da1bee95b378 ...string) (__obf_80a78f1bfaca43d3 error) {
	__obf_706407c46b1ce40f := new(bytes.Buffer)
	__obf_65d9ce0e64489fa6 := __obf_9c2b993d609fe396.NewEncoder(__obf_706407c46b1ce40f)
	//	enc.SetEscapeHTML(false)
	__obf_80a78f1bfaca43d3 = __obf_65d9ce0e64489fa6.Encode(__obf_fe34695ffe11948d)
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	__obf_525303f3f56195a3 := "application/json;charset=utf-8"
	if len(__obf_fea5da1bee95b378) > 0 {
		__obf_525303f3f56195a3 = strings.Join(__obf_fea5da1bee95b378, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_f844075f41f45d17, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Post(__obf_ae0b02887150dacd, __obf_525303f3f56195a3, __obf_706407c46b1ce40f)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_f844075f41f45d17.Body.Close()

	if __obf_f844075f41f45d17.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_ae0b02887150dacd, __obf_f844075f41f45d17.StatusCode)
	}
	return __obf_9c2b993d609fe396.NewDecoder(__obf_f844075f41f45d17.Body).Decode(__obf_521605711569ff86)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_ae0b02887150dacd string, __obf_fe34695ffe11948d any, __obf_521605711569ff86 any) (__obf_80a78f1bfaca43d3 error) {
	__obf_706407c46b1ce40f := new(bytes.Buffer)
	__obf_65d9ce0e64489fa6 := xml.NewEncoder(__obf_706407c46b1ce40f)
	//	enc.SetEscapeHTML(false)
	__obf_80a78f1bfaca43d3 = __obf_65d9ce0e64489fa6.Encode(__obf_fe34695ffe11948d)
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}

	__obf_f844075f41f45d17, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Post(__obf_ae0b02887150dacd, "application/xml;charset=utf-8", __obf_706407c46b1ce40f)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_f844075f41f45d17.Body.Close()

	if __obf_f844075f41f45d17.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_ae0b02887150dacd, __obf_f844075f41f45d17.StatusCode)
	}
	return xml.NewDecoder(__obf_f844075f41f45d17.Body).Decode(__obf_521605711569ff86)
}

// PostFile 上传文件
func PostFile(__obf_39d692314fa38b41, __obf_6eee62656413907b, __obf_ae0b02887150dacd string) ([]byte, error) {
	__obf_8c9c68292225f855 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_39d692314fa38b41,
			Filename:  __obf_6eee62656413907b,
		},
	}
	return PostMultipartForm(__obf_8c9c68292225f855, __obf_ae0b02887150dacd)
}

// GetFile 下载文件
func GetFile(__obf_6eee62656413907b, __obf_ae0b02887150dacd string) error {
	__obf_f844075f41f45d17, __obf_80a78f1bfaca43d3 := __obf_7dd7a45197327cb3().Get(__obf_ae0b02887150dacd)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_f844075f41f45d17.Body.Close()
	__obf_e32d040a715bd08e, __obf_80a78f1bfaca43d3 := os.Create(__obf_6eee62656413907b)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_e32d040a715bd08e.Close()
	_, __obf_80a78f1bfaca43d3 = io.Copy(__obf_e32d040a715bd08e, __obf_f844075f41f45d17.Body)
	return __obf_80a78f1bfaca43d3
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_8c9c68292225f855 []MultipartFormField, __obf_ae0b02887150dacd string) (__obf_eae524014a1820ab []byte, __obf_80a78f1bfaca43d3 error) {
	__obf_edb4f54986fae55f := &bytes.Buffer{}
	__obf_170c758aec77c0c6 := multipart.NewWriter(__obf_edb4f54986fae55f)

	for _, __obf_cd437b447d65c60b := range __obf_8c9c68292225f855 {
		if __obf_cd437b447d65c60b.IsFile {
			__obf_3b9d0115c1c47903, __obf_75c519d5845e5d1e := __obf_170c758aec77c0c6.CreateFormFile(__obf_cd437b447d65c60b.Fieldname, filepath.Base(__obf_cd437b447d65c60b.Filename))
			if __obf_75c519d5845e5d1e != nil {
				__obf_80a78f1bfaca43d3 = fmt.Errorf("error writing to buffer , err=%v", __obf_75c519d5845e5d1e)
				return
			}

			__obf_3327ce99bab6e495, __obf_75c519d5845e5d1e := os.Open(__obf_cd437b447d65c60b.Filename)
			if __obf_75c519d5845e5d1e != nil {
				__obf_80a78f1bfaca43d3 = fmt.Errorf("error opening file , err=%v", __obf_75c519d5845e5d1e)
				return
			}
			defer __obf_3327ce99bab6e495.Close()

			if _, __obf_80a78f1bfaca43d3 = io.Copy(__obf_3b9d0115c1c47903, __obf_3327ce99bab6e495); __obf_80a78f1bfaca43d3 != nil {
				return
			}
		} else {
			__obf_aa34be1b2a0c7c2a, __obf_75c519d5845e5d1e := __obf_170c758aec77c0c6.CreateFormField(__obf_cd437b447d65c60b.Fieldname)
			if __obf_75c519d5845e5d1e != nil {
				__obf_80a78f1bfaca43d3 = __obf_75c519d5845e5d1e
				return
			}
			__obf_0f817ba9be0f5608 := bytes.NewReader(__obf_cd437b447d65c60b.Value)
			if _, __obf_80a78f1bfaca43d3 = io.Copy(__obf_aa34be1b2a0c7c2a, __obf_0f817ba9be0f5608); __obf_80a78f1bfaca43d3 != nil {
				return
			}
		}
	}

	__obf_fea5da1bee95b378 := __obf_170c758aec77c0c6.FormDataContentType()
	__obf_170c758aec77c0c6.Close()

	__obf_f844075f41f45d17, __obf_75c519d5845e5d1e := __obf_7dd7a45197327cb3().Post(__obf_ae0b02887150dacd, __obf_fea5da1bee95b378, __obf_edb4f54986fae55f)
	if __obf_75c519d5845e5d1e != nil {
		__obf_80a78f1bfaca43d3 = __obf_75c519d5845e5d1e
		return
	}
	defer __obf_f844075f41f45d17.Body.Close()

	if __obf_f844075f41f45d17.StatusCode != http.StatusOK {
		return nil, __obf_80a78f1bfaca43d3
	}
	return io.ReadAll(__obf_f844075f41f45d17.Body)
}
