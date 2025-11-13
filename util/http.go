package __obf_813b65e0329aecbf

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
func SetTimeOut(__obf_44b44226049e1da4 time.Duration) {
	TimeOut = __obf_44b44226049e1da4
}

// httpClient() 带超时的http.Client
func __obf_c2723d6fb4425b32() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_124958dbe5716ba2 string, __obf_2dfe4b44e04cc8de any) error {

	__obf_138ad5c24fa33841, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Get(__obf_124958dbe5716ba2)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_138ad5c24fa33841.Body.Close()
	return __obf_006321724ad61b24.NewDecoder(__obf_138ad5c24fa33841.Body).Decode(__obf_2dfe4b44e04cc8de)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_124958dbe5716ba2 string, __obf_2dfe4b44e04cc8de any) error {
	__obf_138ad5c24fa33841, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Get(__obf_124958dbe5716ba2)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_138ad5c24fa33841.Body.Close()
	return xml.NewDecoder(__obf_138ad5c24fa33841.Body).Decode(__obf_2dfe4b44e04cc8de)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_124958dbe5716ba2 string) ([]byte, error) {
	__obf_6bef80debccde8a6, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Get(__obf_124958dbe5716ba2)
	if __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}
	defer __obf_6bef80debccde8a6.Body.Close()

	if __obf_6bef80debccde8a6.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_124958dbe5716ba2, __obf_6bef80debccde8a6.StatusCode)
	}
	return io.ReadAll(__obf_6bef80debccde8a6.Body)
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
func PostJson(__obf_124958dbe5716ba2 string, __obf_349ecbd50f949190 any) ([]byte, error) {
	__obf_d7abab38c74206e1 := new(bytes.Buffer)
	__obf_b334ec5fe42448f9 := __obf_006321724ad61b24.NewEncoder(__obf_d7abab38c74206e1)
	__obf_b334ec5fe42448f9.SetEscapeHTML(false)
	__obf_54101b1325683820 := __obf_b334ec5fe42448f9.Encode(__obf_349ecbd50f949190)
	if __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}
	__obf_6bef80debccde8a6, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Post(__obf_124958dbe5716ba2, "application/json;charset=utf-8", __obf_d7abab38c74206e1)
	if __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}
	defer __obf_6bef80debccde8a6.Body.Close()

	if __obf_6bef80debccde8a6.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_124958dbe5716ba2, __obf_6bef80debccde8a6.StatusCode)
	}
	return io.ReadAll(__obf_6bef80debccde8a6.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_124958dbe5716ba2 string, __obf_349ecbd50f949190 any, __obf_0aa760b4b83b758d any, __obf_07625d30c87ed84e ...string) (__obf_54101b1325683820 error) {
	__obf_d7abab38c74206e1 := new(bytes.Buffer)
	__obf_b334ec5fe42448f9 := __obf_006321724ad61b24.NewEncoder(__obf_d7abab38c74206e1)
	//	enc.SetEscapeHTML(false)
	__obf_54101b1325683820 = __obf_b334ec5fe42448f9.Encode(__obf_349ecbd50f949190)
	if __obf_54101b1325683820 != nil {
		return
	}
	__obf_ec8707973b9401de := "application/json;charset=utf-8"
	if len(__obf_07625d30c87ed84e) > 0 {
		__obf_ec8707973b9401de = strings.Join(__obf_07625d30c87ed84e, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_6bef80debccde8a6, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Post(__obf_124958dbe5716ba2, __obf_ec8707973b9401de, __obf_d7abab38c74206e1)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_6bef80debccde8a6.Body.Close()

	if __obf_6bef80debccde8a6.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_124958dbe5716ba2, __obf_6bef80debccde8a6.StatusCode)
	}
	return __obf_006321724ad61b24.NewDecoder(__obf_6bef80debccde8a6.Body).Decode(__obf_0aa760b4b83b758d)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_124958dbe5716ba2 string, __obf_349ecbd50f949190 any, __obf_0aa760b4b83b758d any) (__obf_54101b1325683820 error) {
	__obf_d7abab38c74206e1 := new(bytes.Buffer)
	__obf_b334ec5fe42448f9 := xml.NewEncoder(__obf_d7abab38c74206e1)
	//	enc.SetEscapeHTML(false)
	__obf_54101b1325683820 = __obf_b334ec5fe42448f9.Encode(__obf_349ecbd50f949190)
	if __obf_54101b1325683820 != nil {
		return
	}

	__obf_6bef80debccde8a6, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Post(__obf_124958dbe5716ba2, "application/xml;charset=utf-8", __obf_d7abab38c74206e1)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_6bef80debccde8a6.Body.Close()

	if __obf_6bef80debccde8a6.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_124958dbe5716ba2, __obf_6bef80debccde8a6.StatusCode)
	}
	return xml.NewDecoder(__obf_6bef80debccde8a6.Body).Decode(__obf_0aa760b4b83b758d)
}

// PostFile 上传文件
func PostFile(__obf_07d41d4162995a34, __obf_999db615e71b11d1, __obf_124958dbe5716ba2 string) ([]byte, error) {
	__obf_54b8de609e638b80 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_07d41d4162995a34,
			Filename:  __obf_999db615e71b11d1,
		},
	}
	return PostMultipartForm(__obf_54b8de609e638b80, __obf_124958dbe5716ba2)
}

// GetFile 下载文件
func GetFile(__obf_999db615e71b11d1, __obf_124958dbe5716ba2 string) error {
	__obf_6bef80debccde8a6, __obf_54101b1325683820 := __obf_c2723d6fb4425b32().Get(__obf_124958dbe5716ba2)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_6bef80debccde8a6.Body.Close()
	__obf_d915d232ea3e4e3d, __obf_54101b1325683820 := os.Create(__obf_999db615e71b11d1)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_d915d232ea3e4e3d.Close()
	_, __obf_54101b1325683820 = io.Copy(__obf_d915d232ea3e4e3d, __obf_6bef80debccde8a6.Body)
	return __obf_54101b1325683820
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_54b8de609e638b80 []MultipartFormField, __obf_124958dbe5716ba2 string) (__obf_0ff39c7c4d479e24 []byte, __obf_54101b1325683820 error) {
	__obf_21a6f2efabcbdac1 := &bytes.Buffer{}
	__obf_167b42eb39cdf278 := multipart.NewWriter(__obf_21a6f2efabcbdac1)

	for _, __obf_1fb9c6ac5717e81d := range __obf_54b8de609e638b80 {
		if __obf_1fb9c6ac5717e81d.IsFile {
			__obf_faf1ee485755d831, __obf_fc343dd99409448f := __obf_167b42eb39cdf278.CreateFormFile(__obf_1fb9c6ac5717e81d.Fieldname, filepath.Base(__obf_1fb9c6ac5717e81d.Filename))
			if __obf_fc343dd99409448f != nil {
				__obf_54101b1325683820 = fmt.Errorf("error writing to buffer , err=%v", __obf_fc343dd99409448f)
				return
			}

			__obf_7475a259e6313e59, __obf_fc343dd99409448f := os.Open(__obf_1fb9c6ac5717e81d.Filename)
			if __obf_fc343dd99409448f != nil {
				__obf_54101b1325683820 = fmt.Errorf("error opening file , err=%v", __obf_fc343dd99409448f)
				return
			}
			defer __obf_7475a259e6313e59.Close()

			if _, __obf_54101b1325683820 = io.Copy(__obf_faf1ee485755d831, __obf_7475a259e6313e59); __obf_54101b1325683820 != nil {
				return
			}
		} else {
			__obf_f1d63570cc11335f, __obf_fc343dd99409448f := __obf_167b42eb39cdf278.CreateFormField(__obf_1fb9c6ac5717e81d.Fieldname)
			if __obf_fc343dd99409448f != nil {
				__obf_54101b1325683820 = __obf_fc343dd99409448f
				return
			}
			__obf_df44404d20cf73ba := bytes.NewReader(__obf_1fb9c6ac5717e81d.Value)
			if _, __obf_54101b1325683820 = io.Copy(__obf_f1d63570cc11335f, __obf_df44404d20cf73ba); __obf_54101b1325683820 != nil {
				return
			}
		}
	}

	__obf_07625d30c87ed84e := __obf_167b42eb39cdf278.FormDataContentType()
	__obf_167b42eb39cdf278.Close()

	__obf_6bef80debccde8a6, __obf_fc343dd99409448f := __obf_c2723d6fb4425b32().Post(__obf_124958dbe5716ba2, __obf_07625d30c87ed84e, __obf_21a6f2efabcbdac1)
	if __obf_fc343dd99409448f != nil {
		__obf_54101b1325683820 = __obf_fc343dd99409448f
		return
	}
	defer __obf_6bef80debccde8a6.Body.Close()

	if __obf_6bef80debccde8a6.StatusCode != http.StatusOK {
		return nil, __obf_54101b1325683820
	}
	return io.ReadAll(__obf_6bef80debccde8a6.Body)
}
