package __obf_8b17832dd38bb602

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
func SetTimeOut(__obf_82958ae45e49795c time.Duration) {
	TimeOut = __obf_82958ae45e49795c
}

// httpClient() 带超时的http.Client
func __obf_ac6438675878ca20() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_c3427b9da52a5d20 string, __obf_d4f65aef7baa5a20 any) error {
	__obf_274b76b7499b5c52, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Get(__obf_c3427b9da52a5d20)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_274b76b7499b5c52.Body.Close()
	return __obf_1688e890467fd727.NewDecoder(__obf_274b76b7499b5c52.Body).Decode(__obf_d4f65aef7baa5a20)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_c3427b9da52a5d20 string, __obf_d4f65aef7baa5a20 any) error {
	__obf_274b76b7499b5c52, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Get(__obf_c3427b9da52a5d20)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_274b76b7499b5c52.Body.Close()
	return xml.NewDecoder(__obf_274b76b7499b5c52.Body).Decode(__obf_d4f65aef7baa5a20)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_c3427b9da52a5d20 string) ([]byte, error) {
	__obf_a4e85161536df34a, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Get(__obf_c3427b9da52a5d20)
	if __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	defer __obf_a4e85161536df34a.Body.Close()

	if __obf_a4e85161536df34a.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_c3427b9da52a5d20, __obf_a4e85161536df34a.StatusCode)
	}
	return io.ReadAll(__obf_a4e85161536df34a.Body)
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
func PostJson(__obf_c3427b9da52a5d20 string, __obf_ab7a93cf3f7c18cc any) ([]byte, error) {
	__obf_7b85f3d43771863b := new(bytes.Buffer)
	__obf_ed37dc65cf066a67 := __obf_1688e890467fd727.NewEncoder(__obf_7b85f3d43771863b)
	__obf_ed37dc65cf066a67.
		SetEscapeHTML(false)
	__obf_17881c470b1e6626 := __obf_ed37dc65cf066a67.Encode(__obf_ab7a93cf3f7c18cc)
	if __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	__obf_a4e85161536df34a, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Post(__obf_c3427b9da52a5d20, "application/json;charset=utf-8", __obf_7b85f3d43771863b)
	if __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	defer __obf_a4e85161536df34a.Body.Close()

	if __obf_a4e85161536df34a.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_c3427b9da52a5d20, __obf_a4e85161536df34a.StatusCode)
	}
	return io.ReadAll(__obf_a4e85161536df34a.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_c3427b9da52a5d20 string, __obf_ab7a93cf3f7c18cc any, __obf_5d1c6c1ddd8f37a7 any, __obf_eb2a4943afb07adb ...string) (__obf_17881c470b1e6626 error) {
	__obf_7b85f3d43771863b := new(bytes.Buffer)
	__obf_ed37dc65cf066a67 := __obf_1688e890467fd727.NewEncoder(__obf_7b85f3d43771863b)
	__obf_17881c470b1e6626 = //	enc.SetEscapeHTML(false)
		__obf_ed37dc65cf066a67.Encode(__obf_ab7a93cf3f7c18cc)
	if __obf_17881c470b1e6626 != nil {
		return
	}
	__obf_5305c1450b16676d := "application/json;charset=utf-8"
	if len(__obf_eb2a4943afb07adb) > 0 {
		__obf_5305c1450b16676d = strings.Join(__obf_eb2a4943afb07adb, ";")
	}
	__obf_a4e85161536df34a,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_17881c470b1e6626 := __obf_ac6438675878ca20().Post(__obf_c3427b9da52a5d20, __obf_5305c1450b16676d, __obf_7b85f3d43771863b)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_a4e85161536df34a.Body.Close()

	if __obf_a4e85161536df34a.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_c3427b9da52a5d20, __obf_a4e85161536df34a.StatusCode)
	}
	return __obf_1688e890467fd727.NewDecoder(__obf_a4e85161536df34a.Body).Decode(__obf_5d1c6c1ddd8f37a7)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_c3427b9da52a5d20 string, __obf_ab7a93cf3f7c18cc any, __obf_5d1c6c1ddd8f37a7 any) (__obf_17881c470b1e6626 error) {
	__obf_7b85f3d43771863b := new(bytes.Buffer)
	__obf_ed37dc65cf066a67 := xml.NewEncoder(__obf_7b85f3d43771863b)
	__obf_17881c470b1e6626 = //	enc.SetEscapeHTML(false)
		__obf_ed37dc65cf066a67.Encode(__obf_ab7a93cf3f7c18cc)
	if __obf_17881c470b1e6626 != nil {
		return
	}
	__obf_a4e85161536df34a, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Post(__obf_c3427b9da52a5d20, "application/xml;charset=utf-8", __obf_7b85f3d43771863b)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_a4e85161536df34a.Body.Close()

	if __obf_a4e85161536df34a.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_c3427b9da52a5d20, __obf_a4e85161536df34a.StatusCode)
	}
	return xml.NewDecoder(__obf_a4e85161536df34a.Body).Decode(__obf_5d1c6c1ddd8f37a7)
}

// PostFile 上传文件
func PostFile(__obf_cbd34a18b802e289, __obf_60b44cf0449a03d9, __obf_c3427b9da52a5d20 string) ([]byte, error) {
	__obf_66b13cc5b41e08e2 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_cbd34a18b802e289,
			Filename:  __obf_60b44cf0449a03d9,
		},
	}
	return PostMultipartForm(__obf_66b13cc5b41e08e2,

		// GetFile 下载文件
		__obf_c3427b9da52a5d20)
}

func GetFile(__obf_60b44cf0449a03d9, __obf_c3427b9da52a5d20 string) error {
	__obf_a4e85161536df34a, __obf_17881c470b1e6626 := __obf_ac6438675878ca20().Get(__obf_c3427b9da52a5d20)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_a4e85161536df34a.Body.Close()
	__obf_960187e3f99d50ce, __obf_17881c470b1e6626 := os.Create(__obf_60b44cf0449a03d9)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_960187e3f99d50ce.Close()
	_, __obf_17881c470b1e6626 = io.Copy(__obf_960187e3f99d50ce, __obf_a4e85161536df34a.Body)
	return __obf_17881c470b1e6626
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_66b13cc5b41e08e2 []MultipartFormField, __obf_c3427b9da52a5d20 string) (__obf_63c6f4c604d6446f []byte, __obf_17881c470b1e6626 error) {
	__obf_bb599ba804540f66 := &bytes.Buffer{}
	__obf_11fb77146ccca640 := multipart.NewWriter(__obf_bb599ba804540f66)

	for _, __obf_6247647f8d59ee66 := range __obf_66b13cc5b41e08e2 {
		if __obf_6247647f8d59ee66.IsFile {
			__obf_3a967a0cd1af1415, __obf_abf6da56a0f5212f := __obf_11fb77146ccca640.CreateFormFile(__obf_6247647f8d59ee66.Fieldname, filepath.Base(__obf_6247647f8d59ee66.Filename))
			if __obf_abf6da56a0f5212f != nil {
				__obf_17881c470b1e6626 = fmt.Errorf("error writing to buffer , err=%v", __obf_abf6da56a0f5212f)
				return
			}
			__obf_7e4defd6fe5e42ee, __obf_abf6da56a0f5212f := os.Open(__obf_6247647f8d59ee66.Filename)
			if __obf_abf6da56a0f5212f != nil {
				__obf_17881c470b1e6626 = fmt.Errorf("error opening file , err=%v", __obf_abf6da56a0f5212f)
				return
			}
			defer __obf_7e4defd6fe5e42ee.Close()

			if _, __obf_17881c470b1e6626 = io.Copy(__obf_3a967a0cd1af1415, __obf_7e4defd6fe5e42ee); __obf_17881c470b1e6626 != nil {
				return
			}
		} else {
			__obf_12976c6039d8ce80, __obf_abf6da56a0f5212f := __obf_11fb77146ccca640.CreateFormField(__obf_6247647f8d59ee66.Fieldname)
			if __obf_abf6da56a0f5212f != nil {
				__obf_17881c470b1e6626 = __obf_abf6da56a0f5212f
				return
			}
			__obf_20b915721f9309f8 := bytes.NewReader(__obf_6247647f8d59ee66.Value)
			if _, __obf_17881c470b1e6626 = io.Copy(__obf_12976c6039d8ce80, __obf_20b915721f9309f8); __obf_17881c470b1e6626 != nil {
				return
			}
		}
	}
	__obf_eb2a4943afb07adb := __obf_11fb77146ccca640.FormDataContentType()
	__obf_11fb77146ccca640.
		Close()
	__obf_a4e85161536df34a, __obf_abf6da56a0f5212f := __obf_ac6438675878ca20().Post(__obf_c3427b9da52a5d20, __obf_eb2a4943afb07adb, __obf_bb599ba804540f66)
	if __obf_abf6da56a0f5212f != nil {
		__obf_17881c470b1e6626 = __obf_abf6da56a0f5212f
		return
	}
	defer __obf_a4e85161536df34a.Body.Close()

	if __obf_a4e85161536df34a.StatusCode != http.StatusOK {
		return nil, __obf_17881c470b1e6626
	}
	return io.ReadAll(__obf_a4e85161536df34a.Body)
}
