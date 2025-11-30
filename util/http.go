package __obf_e13b701bec415b48

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
func SetTimeOut(__obf_6ea3b2d60a754d22 time.Duration) {
	TimeOut = __obf_6ea3b2d60a754d22
}

// httpClient() 带超时的http.Client
func __obf_595323148bd835a9() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_d0b3f61e4d2f2802 string, __obf_99fad6cd23d30931 any) error {
	__obf_af65a69017eda959, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Get(__obf_d0b3f61e4d2f2802)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_af65a69017eda959.Body.Close()
	return __obf_8a0276aba040f633.NewDecoder(__obf_af65a69017eda959.Body).Decode(__obf_99fad6cd23d30931)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_d0b3f61e4d2f2802 string, __obf_99fad6cd23d30931 any) error {
	__obf_af65a69017eda959, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Get(__obf_d0b3f61e4d2f2802)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_af65a69017eda959.Body.Close()
	return xml.NewDecoder(__obf_af65a69017eda959.Body).Decode(__obf_99fad6cd23d30931)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_d0b3f61e4d2f2802 string) ([]byte, error) {
	__obf_a2bfdfa157d5962d, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Get(__obf_d0b3f61e4d2f2802)
	if __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()

	if __obf_a2bfdfa157d5962d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_d0b3f61e4d2f2802, __obf_a2bfdfa157d5962d.StatusCode)
	}
	return io.ReadAll(__obf_a2bfdfa157d5962d.Body)
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
func PostJson(__obf_d0b3f61e4d2f2802 string, __obf_2701e0e32229860d any) ([]byte, error) {
	__obf_84565b64d632fdd6 := new(bytes.Buffer)
	__obf_d9951a8bab9b5357 := __obf_8a0276aba040f633.NewEncoder(__obf_84565b64d632fdd6)
	__obf_d9951a8bab9b5357.
		SetEscapeHTML(false)
	__obf_b93c04d14e5c503f := __obf_d9951a8bab9b5357.Encode(__obf_2701e0e32229860d)
	if __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	__obf_a2bfdfa157d5962d, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Post(__obf_d0b3f61e4d2f2802, "application/json;charset=utf-8", __obf_84565b64d632fdd6)
	if __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()

	if __obf_a2bfdfa157d5962d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_d0b3f61e4d2f2802, __obf_a2bfdfa157d5962d.StatusCode)
	}
	return io.ReadAll(__obf_a2bfdfa157d5962d.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_d0b3f61e4d2f2802 string, __obf_2701e0e32229860d any, __obf_d230249dd0c37358 any, __obf_c4d1d323c3b638a6 ...string) (__obf_b93c04d14e5c503f error) {
	__obf_84565b64d632fdd6 := new(bytes.Buffer)
	__obf_d9951a8bab9b5357 := __obf_8a0276aba040f633.NewEncoder(__obf_84565b64d632fdd6)
	__obf_b93c04d14e5c503f = //	enc.SetEscapeHTML(false)
		__obf_d9951a8bab9b5357.Encode(__obf_2701e0e32229860d)
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	__obf_9f7dc2ed40c46725 := "application/json;charset=utf-8"
	if len(__obf_c4d1d323c3b638a6) > 0 {
		__obf_9f7dc2ed40c46725 = strings.Join(__obf_c4d1d323c3b638a6, ";")
	}
	__obf_a2bfdfa157d5962d,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_b93c04d14e5c503f := __obf_595323148bd835a9().Post(__obf_d0b3f61e4d2f2802, __obf_9f7dc2ed40c46725, __obf_84565b64d632fdd6)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()

	if __obf_a2bfdfa157d5962d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_d0b3f61e4d2f2802, __obf_a2bfdfa157d5962d.StatusCode)
	}
	return __obf_8a0276aba040f633.NewDecoder(__obf_a2bfdfa157d5962d.Body).Decode(__obf_d230249dd0c37358)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_d0b3f61e4d2f2802 string, __obf_2701e0e32229860d any, __obf_d230249dd0c37358 any) (__obf_b93c04d14e5c503f error) {
	__obf_84565b64d632fdd6 := new(bytes.Buffer)
	__obf_d9951a8bab9b5357 := xml.NewEncoder(__obf_84565b64d632fdd6)
	__obf_b93c04d14e5c503f = //	enc.SetEscapeHTML(false)
		__obf_d9951a8bab9b5357.Encode(__obf_2701e0e32229860d)
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	__obf_a2bfdfa157d5962d, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Post(__obf_d0b3f61e4d2f2802, "application/xml;charset=utf-8", __obf_84565b64d632fdd6)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()

	if __obf_a2bfdfa157d5962d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_d0b3f61e4d2f2802, __obf_a2bfdfa157d5962d.StatusCode)
	}
	return xml.NewDecoder(__obf_a2bfdfa157d5962d.Body).Decode(__obf_d230249dd0c37358)
}

// PostFile 上传文件
func PostFile(__obf_bd5053d7f1cf86fe, __obf_5c9c2b1eb9a7f576, __obf_d0b3f61e4d2f2802 string) ([]byte, error) {
	__obf_3c45d7160d05ff9c := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_bd5053d7f1cf86fe,
			Filename:  __obf_5c9c2b1eb9a7f576,
		},
	}
	return PostMultipartForm(__obf_3c45d7160d05ff9c,

		// GetFile 下载文件
		__obf_d0b3f61e4d2f2802)
}

func GetFile(__obf_5c9c2b1eb9a7f576, __obf_d0b3f61e4d2f2802 string) error {
	__obf_a2bfdfa157d5962d, __obf_b93c04d14e5c503f := __obf_595323148bd835a9().Get(__obf_d0b3f61e4d2f2802)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()
	__obf_8eb1bb7eab8d6404, __obf_b93c04d14e5c503f := os.Create(__obf_5c9c2b1eb9a7f576)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_8eb1bb7eab8d6404.Close()
	_, __obf_b93c04d14e5c503f = io.Copy(__obf_8eb1bb7eab8d6404, __obf_a2bfdfa157d5962d.Body)
	return __obf_b93c04d14e5c503f
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_3c45d7160d05ff9c []MultipartFormField, __obf_d0b3f61e4d2f2802 string) (__obf_162eed34eea5e136 []byte, __obf_b93c04d14e5c503f error) {
	__obf_56d78fab9dacc486 := &bytes.Buffer{}
	__obf_74dc4ccede9b9cc0 := multipart.NewWriter(__obf_56d78fab9dacc486)

	for _, __obf_63b9913831dd378d := range __obf_3c45d7160d05ff9c {
		if __obf_63b9913831dd378d.IsFile {
			__obf_cd18030411c3b3f4, __obf_24a198bd8c59d97f := __obf_74dc4ccede9b9cc0.CreateFormFile(__obf_63b9913831dd378d.Fieldname, filepath.Base(__obf_63b9913831dd378d.Filename))
			if __obf_24a198bd8c59d97f != nil {
				__obf_b93c04d14e5c503f = fmt.Errorf("error writing to buffer , err=%v", __obf_24a198bd8c59d97f)
				return
			}
			__obf_b4a77345e230b5c7, __obf_24a198bd8c59d97f := os.Open(__obf_63b9913831dd378d.Filename)
			if __obf_24a198bd8c59d97f != nil {
				__obf_b93c04d14e5c503f = fmt.Errorf("error opening file , err=%v", __obf_24a198bd8c59d97f)
				return
			}
			defer __obf_b4a77345e230b5c7.Close()

			if _, __obf_b93c04d14e5c503f = io.Copy(__obf_cd18030411c3b3f4, __obf_b4a77345e230b5c7); __obf_b93c04d14e5c503f != nil {
				return
			}
		} else {
			__obf_c9e242d6b1386e53, __obf_24a198bd8c59d97f := __obf_74dc4ccede9b9cc0.CreateFormField(__obf_63b9913831dd378d.Fieldname)
			if __obf_24a198bd8c59d97f != nil {
				__obf_b93c04d14e5c503f = __obf_24a198bd8c59d97f
				return
			}
			__obf_765f40e60c1ff9ef := bytes.NewReader(__obf_63b9913831dd378d.Value)
			if _, __obf_b93c04d14e5c503f = io.Copy(__obf_c9e242d6b1386e53, __obf_765f40e60c1ff9ef); __obf_b93c04d14e5c503f != nil {
				return
			}
		}
	}
	__obf_c4d1d323c3b638a6 := __obf_74dc4ccede9b9cc0.FormDataContentType()
	__obf_74dc4ccede9b9cc0.
		Close()
	__obf_a2bfdfa157d5962d, __obf_24a198bd8c59d97f := __obf_595323148bd835a9().Post(__obf_d0b3f61e4d2f2802, __obf_c4d1d323c3b638a6, __obf_56d78fab9dacc486)
	if __obf_24a198bd8c59d97f != nil {
		__obf_b93c04d14e5c503f = __obf_24a198bd8c59d97f
		return
	}
	defer __obf_a2bfdfa157d5962d.Body.Close()

	if __obf_a2bfdfa157d5962d.StatusCode != http.StatusOK {
		return nil, __obf_b93c04d14e5c503f
	}
	return io.ReadAll(__obf_a2bfdfa157d5962d.Body)
}
