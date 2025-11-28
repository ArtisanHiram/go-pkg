package __obf_f1e346e3aa5cc554

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
func SetTimeOut(__obf_8beab787517c6111 time.Duration) {
	TimeOut = __obf_8beab787517c6111
}

// httpClient() 带超时的http.Client
func __obf_eb674f05d452612b() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_a3d8a1d7a26a1b34 string, __obf_29f3385acfb70f69 any) error {

	__obf_72fc7f85fed6cac3, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Get(__obf_a3d8a1d7a26a1b34)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_72fc7f85fed6cac3.Body.Close()
	return __obf_f3d2a2a7fb490096.NewDecoder(__obf_72fc7f85fed6cac3.Body).Decode(__obf_29f3385acfb70f69)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_a3d8a1d7a26a1b34 string, __obf_29f3385acfb70f69 any) error {
	__obf_72fc7f85fed6cac3, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Get(__obf_a3d8a1d7a26a1b34)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_72fc7f85fed6cac3.Body.Close()
	return xml.NewDecoder(__obf_72fc7f85fed6cac3.Body).Decode(__obf_29f3385acfb70f69)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_a3d8a1d7a26a1b34 string) ([]byte, error) {
	__obf_a215c76839be61a5, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Get(__obf_a3d8a1d7a26a1b34)
	if __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}
	defer __obf_a215c76839be61a5.Body.Close()

	if __obf_a215c76839be61a5.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_a3d8a1d7a26a1b34, __obf_a215c76839be61a5.StatusCode)
	}
	return io.ReadAll(__obf_a215c76839be61a5.Body)
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
func PostJson(__obf_a3d8a1d7a26a1b34 string, __obf_8690295b52001a12 any) ([]byte, error) {
	__obf_92a6c1286d60d679 := new(bytes.Buffer)
	__obf_34006b956928274a := __obf_f3d2a2a7fb490096.NewEncoder(__obf_92a6c1286d60d679)
	__obf_34006b956928274a.SetEscapeHTML(false)
	__obf_eec784b359ebf42f := __obf_34006b956928274a.Encode(__obf_8690295b52001a12)
	if __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}
	__obf_a215c76839be61a5, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Post(__obf_a3d8a1d7a26a1b34, "application/json;charset=utf-8", __obf_92a6c1286d60d679)
	if __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}
	defer __obf_a215c76839be61a5.Body.Close()

	if __obf_a215c76839be61a5.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a3d8a1d7a26a1b34, __obf_a215c76839be61a5.StatusCode)
	}
	return io.ReadAll(__obf_a215c76839be61a5.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_a3d8a1d7a26a1b34 string, __obf_8690295b52001a12 any, __obf_969aede8eee9c705 any, __obf_66c3ab22b6b799e9 ...string) (__obf_eec784b359ebf42f error) {
	__obf_92a6c1286d60d679 := new(bytes.Buffer)
	__obf_34006b956928274a := __obf_f3d2a2a7fb490096.NewEncoder(__obf_92a6c1286d60d679)
	//	enc.SetEscapeHTML(false)
	__obf_eec784b359ebf42f = __obf_34006b956928274a.Encode(__obf_8690295b52001a12)
	if __obf_eec784b359ebf42f != nil {
		return
	}
	__obf_9557d38bd34367a6 := "application/json;charset=utf-8"
	if len(__obf_66c3ab22b6b799e9) > 0 {
		__obf_9557d38bd34367a6 = strings.Join(__obf_66c3ab22b6b799e9, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_a215c76839be61a5, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Post(__obf_a3d8a1d7a26a1b34, __obf_9557d38bd34367a6, __obf_92a6c1286d60d679)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_a215c76839be61a5.Body.Close()

	if __obf_a215c76839be61a5.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a3d8a1d7a26a1b34, __obf_a215c76839be61a5.StatusCode)
	}
	return __obf_f3d2a2a7fb490096.NewDecoder(__obf_a215c76839be61a5.Body).Decode(__obf_969aede8eee9c705)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_a3d8a1d7a26a1b34 string, __obf_8690295b52001a12 any, __obf_969aede8eee9c705 any) (__obf_eec784b359ebf42f error) {
	__obf_92a6c1286d60d679 := new(bytes.Buffer)
	__obf_34006b956928274a := xml.NewEncoder(__obf_92a6c1286d60d679)
	//	enc.SetEscapeHTML(false)
	__obf_eec784b359ebf42f = __obf_34006b956928274a.Encode(__obf_8690295b52001a12)
	if __obf_eec784b359ebf42f != nil {
		return
	}

	__obf_a215c76839be61a5, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Post(__obf_a3d8a1d7a26a1b34, "application/xml;charset=utf-8", __obf_92a6c1286d60d679)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_a215c76839be61a5.Body.Close()

	if __obf_a215c76839be61a5.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a3d8a1d7a26a1b34, __obf_a215c76839be61a5.StatusCode)
	}
	return xml.NewDecoder(__obf_a215c76839be61a5.Body).Decode(__obf_969aede8eee9c705)
}

// PostFile 上传文件
func PostFile(__obf_512f7ba4511bb22d, __obf_4d2a16dc437c92be, __obf_a3d8a1d7a26a1b34 string) ([]byte, error) {
	__obf_e762e70277dfa6bb := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_512f7ba4511bb22d,
			Filename:  __obf_4d2a16dc437c92be,
		},
	}
	return PostMultipartForm(__obf_e762e70277dfa6bb, __obf_a3d8a1d7a26a1b34)
}

// GetFile 下载文件
func GetFile(__obf_4d2a16dc437c92be, __obf_a3d8a1d7a26a1b34 string) error {
	__obf_a215c76839be61a5, __obf_eec784b359ebf42f := __obf_eb674f05d452612b().Get(__obf_a3d8a1d7a26a1b34)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_a215c76839be61a5.Body.Close()
	__obf_a7fa827cefd5f9da, __obf_eec784b359ebf42f := os.Create(__obf_4d2a16dc437c92be)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_a7fa827cefd5f9da.Close()
	_, __obf_eec784b359ebf42f = io.Copy(__obf_a7fa827cefd5f9da, __obf_a215c76839be61a5.Body)
	return __obf_eec784b359ebf42f
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_e762e70277dfa6bb []MultipartFormField, __obf_a3d8a1d7a26a1b34 string) (__obf_fd8156f051ce46a0 []byte, __obf_eec784b359ebf42f error) {
	__obf_2723fbc427b8679e := &bytes.Buffer{}
	__obf_6f34980bb723c1d2 := multipart.NewWriter(__obf_2723fbc427b8679e)

	for _, __obf_8880ca1fd5d68374 := range __obf_e762e70277dfa6bb {
		if __obf_8880ca1fd5d68374.IsFile {
			__obf_8987633b509a2b77, __obf_d4bd147170d96001 := __obf_6f34980bb723c1d2.CreateFormFile(__obf_8880ca1fd5d68374.Fieldname, filepath.Base(__obf_8880ca1fd5d68374.Filename))
			if __obf_d4bd147170d96001 != nil {
				__obf_eec784b359ebf42f = fmt.Errorf("error writing to buffer , err=%v", __obf_d4bd147170d96001)
				return
			}

			__obf_7be5d4fb8c22e328, __obf_d4bd147170d96001 := os.Open(__obf_8880ca1fd5d68374.Filename)
			if __obf_d4bd147170d96001 != nil {
				__obf_eec784b359ebf42f = fmt.Errorf("error opening file , err=%v", __obf_d4bd147170d96001)
				return
			}
			defer __obf_7be5d4fb8c22e328.Close()

			if _, __obf_eec784b359ebf42f = io.Copy(__obf_8987633b509a2b77, __obf_7be5d4fb8c22e328); __obf_eec784b359ebf42f != nil {
				return
			}
		} else {
			__obf_997a0d530d6e7c15, __obf_d4bd147170d96001 := __obf_6f34980bb723c1d2.CreateFormField(__obf_8880ca1fd5d68374.Fieldname)
			if __obf_d4bd147170d96001 != nil {
				__obf_eec784b359ebf42f = __obf_d4bd147170d96001
				return
			}
			__obf_d6b71428b1349375 := bytes.NewReader(__obf_8880ca1fd5d68374.Value)
			if _, __obf_eec784b359ebf42f = io.Copy(__obf_997a0d530d6e7c15, __obf_d6b71428b1349375); __obf_eec784b359ebf42f != nil {
				return
			}
		}
	}

	__obf_66c3ab22b6b799e9 := __obf_6f34980bb723c1d2.FormDataContentType()
	__obf_6f34980bb723c1d2.Close()

	__obf_a215c76839be61a5, __obf_d4bd147170d96001 := __obf_eb674f05d452612b().Post(__obf_a3d8a1d7a26a1b34, __obf_66c3ab22b6b799e9, __obf_2723fbc427b8679e)
	if __obf_d4bd147170d96001 != nil {
		__obf_eec784b359ebf42f = __obf_d4bd147170d96001
		return
	}
	defer __obf_a215c76839be61a5.Body.Close()

	if __obf_a215c76839be61a5.StatusCode != http.StatusOK {
		return nil, __obf_eec784b359ebf42f
	}
	return io.ReadAll(__obf_a215c76839be61a5.Body)
}
