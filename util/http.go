package __obf_077bcefb098a89af

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
func SetTimeOut(__obf_ba749bcf2f186fe5 time.Duration) {
	TimeOut = __obf_ba749bcf2f186fe5
}

// httpClient() 带超时的http.Client
func __obf_71ef16cbb6e99dc2() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_55239f5897d11831 string, __obf_1c4edfc2acf87521 any) error {

	__obf_c07a5c3f92c3565e, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Get(__obf_55239f5897d11831)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_c07a5c3f92c3565e.Body.Close()
	return __obf_24b8a17d9ed5e394.NewDecoder(__obf_c07a5c3f92c3565e.Body).Decode(__obf_1c4edfc2acf87521)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_55239f5897d11831 string, __obf_1c4edfc2acf87521 any) error {
	__obf_c07a5c3f92c3565e, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Get(__obf_55239f5897d11831)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_c07a5c3f92c3565e.Body.Close()
	return xml.NewDecoder(__obf_c07a5c3f92c3565e.Body).Decode(__obf_1c4edfc2acf87521)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_55239f5897d11831 string) ([]byte, error) {
	__obf_c6aed633ce9b381f, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Get(__obf_55239f5897d11831)
	if __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}
	defer __obf_c6aed633ce9b381f.Body.Close()

	if __obf_c6aed633ce9b381f.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_55239f5897d11831, __obf_c6aed633ce9b381f.StatusCode)
	}
	return io.ReadAll(__obf_c6aed633ce9b381f.Body)
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
func PostJson(__obf_55239f5897d11831 string, __obf_b74dbf6a86da485d any) ([]byte, error) {
	__obf_5083bc90f0cd4d2b := new(bytes.Buffer)
	__obf_b4d0e797e83e1033 := __obf_24b8a17d9ed5e394.NewEncoder(__obf_5083bc90f0cd4d2b)
	__obf_b4d0e797e83e1033.SetEscapeHTML(false)
	__obf_224415a75b186177 := __obf_b4d0e797e83e1033.Encode(__obf_b74dbf6a86da485d)
	if __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}
	__obf_c6aed633ce9b381f, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Post(__obf_55239f5897d11831, "application/json;charset=utf-8", __obf_5083bc90f0cd4d2b)
	if __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}
	defer __obf_c6aed633ce9b381f.Body.Close()

	if __obf_c6aed633ce9b381f.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_55239f5897d11831, __obf_c6aed633ce9b381f.StatusCode)
	}
	return io.ReadAll(__obf_c6aed633ce9b381f.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_55239f5897d11831 string, __obf_b74dbf6a86da485d any, __obf_84489a76bf4c5164 any, __obf_001dad27b3bf9f7f ...string) (__obf_224415a75b186177 error) {
	__obf_5083bc90f0cd4d2b := new(bytes.Buffer)
	__obf_b4d0e797e83e1033 := __obf_24b8a17d9ed5e394.NewEncoder(__obf_5083bc90f0cd4d2b)
	//	enc.SetEscapeHTML(false)
	__obf_224415a75b186177 = __obf_b4d0e797e83e1033.Encode(__obf_b74dbf6a86da485d)
	if __obf_224415a75b186177 != nil {
		return
	}
	__obf_506abc8e081e2bfb := "application/json;charset=utf-8"
	if len(__obf_001dad27b3bf9f7f) > 0 {
		__obf_506abc8e081e2bfb = strings.Join(__obf_001dad27b3bf9f7f, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_c6aed633ce9b381f, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Post(__obf_55239f5897d11831, __obf_506abc8e081e2bfb, __obf_5083bc90f0cd4d2b)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_c6aed633ce9b381f.Body.Close()

	if __obf_c6aed633ce9b381f.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_55239f5897d11831, __obf_c6aed633ce9b381f.StatusCode)
	}
	return __obf_24b8a17d9ed5e394.NewDecoder(__obf_c6aed633ce9b381f.Body).Decode(__obf_84489a76bf4c5164)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_55239f5897d11831 string, __obf_b74dbf6a86da485d any, __obf_84489a76bf4c5164 any) (__obf_224415a75b186177 error) {
	__obf_5083bc90f0cd4d2b := new(bytes.Buffer)
	__obf_b4d0e797e83e1033 := xml.NewEncoder(__obf_5083bc90f0cd4d2b)
	//	enc.SetEscapeHTML(false)
	__obf_224415a75b186177 = __obf_b4d0e797e83e1033.Encode(__obf_b74dbf6a86da485d)
	if __obf_224415a75b186177 != nil {
		return
	}

	__obf_c6aed633ce9b381f, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Post(__obf_55239f5897d11831, "application/xml;charset=utf-8", __obf_5083bc90f0cd4d2b)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_c6aed633ce9b381f.Body.Close()

	if __obf_c6aed633ce9b381f.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_55239f5897d11831, __obf_c6aed633ce9b381f.StatusCode)
	}
	return xml.NewDecoder(__obf_c6aed633ce9b381f.Body).Decode(__obf_84489a76bf4c5164)
}

// PostFile 上传文件
func PostFile(__obf_db2049bc55bde66a, __obf_075fd7b3346969fd, __obf_55239f5897d11831 string) ([]byte, error) {
	__obf_91ee077001f3e167 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_db2049bc55bde66a,
			Filename:  __obf_075fd7b3346969fd,
		},
	}
	return PostMultipartForm(__obf_91ee077001f3e167, __obf_55239f5897d11831)
}

// GetFile 下载文件
func GetFile(__obf_075fd7b3346969fd, __obf_55239f5897d11831 string) error {
	__obf_c6aed633ce9b381f, __obf_224415a75b186177 := __obf_71ef16cbb6e99dc2().Get(__obf_55239f5897d11831)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_c6aed633ce9b381f.Body.Close()
	__obf_d33124a4806e2498, __obf_224415a75b186177 := os.Create(__obf_075fd7b3346969fd)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_d33124a4806e2498.Close()
	_, __obf_224415a75b186177 = io.Copy(__obf_d33124a4806e2498, __obf_c6aed633ce9b381f.Body)
	return __obf_224415a75b186177
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_91ee077001f3e167 []MultipartFormField, __obf_55239f5897d11831 string) (__obf_e5dabeaa44051b1d []byte, __obf_224415a75b186177 error) {
	__obf_3709a078bff0bba7 := &bytes.Buffer{}
	__obf_28f5dd7a7a68334e := multipart.NewWriter(__obf_3709a078bff0bba7)

	for _, __obf_f93606ef09426383 := range __obf_91ee077001f3e167 {
		if __obf_f93606ef09426383.IsFile {
			__obf_943d5077f5318744, __obf_e05175036b1dc5fa := __obf_28f5dd7a7a68334e.CreateFormFile(__obf_f93606ef09426383.Fieldname, filepath.Base(__obf_f93606ef09426383.Filename))
			if __obf_e05175036b1dc5fa != nil {
				__obf_224415a75b186177 = fmt.Errorf("error writing to buffer , err=%v", __obf_e05175036b1dc5fa)
				return
			}

			__obf_ee608c776574c305, __obf_e05175036b1dc5fa := os.Open(__obf_f93606ef09426383.Filename)
			if __obf_e05175036b1dc5fa != nil {
				__obf_224415a75b186177 = fmt.Errorf("error opening file , err=%v", __obf_e05175036b1dc5fa)
				return
			}
			defer __obf_ee608c776574c305.Close()

			if _, __obf_224415a75b186177 = io.Copy(__obf_943d5077f5318744, __obf_ee608c776574c305); __obf_224415a75b186177 != nil {
				return
			}
		} else {
			__obf_a23a4ca8f8f4d63b, __obf_e05175036b1dc5fa := __obf_28f5dd7a7a68334e.CreateFormField(__obf_f93606ef09426383.Fieldname)
			if __obf_e05175036b1dc5fa != nil {
				__obf_224415a75b186177 = __obf_e05175036b1dc5fa
				return
			}
			__obf_7ca58f0965767bbc := bytes.NewReader(__obf_f93606ef09426383.Value)
			if _, __obf_224415a75b186177 = io.Copy(__obf_a23a4ca8f8f4d63b, __obf_7ca58f0965767bbc); __obf_224415a75b186177 != nil {
				return
			}
		}
	}

	__obf_001dad27b3bf9f7f := __obf_28f5dd7a7a68334e.FormDataContentType()
	__obf_28f5dd7a7a68334e.Close()

	__obf_c6aed633ce9b381f, __obf_e05175036b1dc5fa := __obf_71ef16cbb6e99dc2().Post(__obf_55239f5897d11831, __obf_001dad27b3bf9f7f, __obf_3709a078bff0bba7)
	if __obf_e05175036b1dc5fa != nil {
		__obf_224415a75b186177 = __obf_e05175036b1dc5fa
		return
	}
	defer __obf_c6aed633ce9b381f.Body.Close()

	if __obf_c6aed633ce9b381f.StatusCode != http.StatusOK {
		return nil, __obf_224415a75b186177
	}
	return io.ReadAll(__obf_c6aed633ce9b381f.Body)
}
