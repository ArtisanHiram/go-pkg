package __obf_426da37e60cac670

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
func SetTimeOut(__obf_ddc716cb6518a16a time.Duration) {
	TimeOut = __obf_ddc716cb6518a16a
}

// httpClient() 带超时的http.Client
func __obf_59c3debac748b990() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_a48e83c31cb47a6e string, __obf_c76af479674ee3d0 any) error {
	__obf_3be342a2e8ac84dc, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Get(__obf_a48e83c31cb47a6e)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_3be342a2e8ac84dc.Body.Close()
	return __obf_c77efeb3502803e0.NewDecoder(__obf_3be342a2e8ac84dc.Body).Decode(__obf_c76af479674ee3d0)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_a48e83c31cb47a6e string, __obf_c76af479674ee3d0 any) error {
	__obf_3be342a2e8ac84dc, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Get(__obf_a48e83c31cb47a6e)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_3be342a2e8ac84dc.Body.Close()
	return xml.NewDecoder(__obf_3be342a2e8ac84dc.Body).Decode(__obf_c76af479674ee3d0)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_a48e83c31cb47a6e string) ([]byte, error) {
	__obf_32eef76263192e6d, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Get(__obf_a48e83c31cb47a6e)
	if __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	defer __obf_32eef76263192e6d.Body.Close()

	if __obf_32eef76263192e6d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_a48e83c31cb47a6e, __obf_32eef76263192e6d.StatusCode)
	}
	return io.ReadAll(__obf_32eef76263192e6d.Body)
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
func PostJson(__obf_a48e83c31cb47a6e string, __obf_936376d3716db60e any) ([]byte, error) {
	__obf_210844b716d29f53 := new(bytes.Buffer)
	__obf_3aa68e4ceb6f78ad := __obf_c77efeb3502803e0.NewEncoder(__obf_210844b716d29f53)
	__obf_3aa68e4ceb6f78ad.
		SetEscapeHTML(false)
	__obf_74916b80241ef1ff := __obf_3aa68e4ceb6f78ad.Encode(__obf_936376d3716db60e)
	if __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	__obf_32eef76263192e6d, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Post(__obf_a48e83c31cb47a6e, "application/json;charset=utf-8", __obf_210844b716d29f53)
	if __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	defer __obf_32eef76263192e6d.Body.Close()

	if __obf_32eef76263192e6d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a48e83c31cb47a6e, __obf_32eef76263192e6d.StatusCode)
	}
	return io.ReadAll(__obf_32eef76263192e6d.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_a48e83c31cb47a6e string, __obf_936376d3716db60e any, __obf_9088057b0091059d any, __obf_52eed32857b9a9bc ...string) (__obf_74916b80241ef1ff error) {
	__obf_210844b716d29f53 := new(bytes.Buffer)
	__obf_3aa68e4ceb6f78ad := __obf_c77efeb3502803e0.NewEncoder(__obf_210844b716d29f53)
	__obf_74916b80241ef1ff = //	enc.SetEscapeHTML(false)
		__obf_3aa68e4ceb6f78ad.Encode(__obf_936376d3716db60e)
	if __obf_74916b80241ef1ff != nil {
		return
	}
	__obf_6fa317b6ee712443 := "application/json;charset=utf-8"
	if len(__obf_52eed32857b9a9bc) > 0 {
		__obf_6fa317b6ee712443 = strings.Join(__obf_52eed32857b9a9bc, ";")
	}
	__obf_32eef76263192e6d,
		// fmt.Println("post buf:", buf.String()) // Debug
		__obf_74916b80241ef1ff := __obf_59c3debac748b990().Post(__obf_a48e83c31cb47a6e, __obf_6fa317b6ee712443, __obf_210844b716d29f53)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_32eef76263192e6d.Body.Close()

	if __obf_32eef76263192e6d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a48e83c31cb47a6e, __obf_32eef76263192e6d.StatusCode)
	}
	return __obf_c77efeb3502803e0.NewDecoder(__obf_32eef76263192e6d.Body).Decode(__obf_9088057b0091059d)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_a48e83c31cb47a6e string, __obf_936376d3716db60e any, __obf_9088057b0091059d any) (__obf_74916b80241ef1ff error) {
	__obf_210844b716d29f53 := new(bytes.Buffer)
	__obf_3aa68e4ceb6f78ad := xml.NewEncoder(__obf_210844b716d29f53)
	__obf_74916b80241ef1ff = //	enc.SetEscapeHTML(false)
		__obf_3aa68e4ceb6f78ad.Encode(__obf_936376d3716db60e)
	if __obf_74916b80241ef1ff != nil {
		return
	}
	__obf_32eef76263192e6d, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Post(__obf_a48e83c31cb47a6e, "application/xml;charset=utf-8", __obf_210844b716d29f53)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_32eef76263192e6d.Body.Close()

	if __obf_32eef76263192e6d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_a48e83c31cb47a6e, __obf_32eef76263192e6d.StatusCode)
	}
	return xml.NewDecoder(__obf_32eef76263192e6d.Body).Decode(__obf_9088057b0091059d)
}

// PostFile 上传文件
func PostFile(__obf_de86f23dd7557f4a, __obf_04310e04e67ea058, __obf_a48e83c31cb47a6e string) ([]byte, error) {
	__obf_1fc78df60c5789b2 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_de86f23dd7557f4a,
			Filename:  __obf_04310e04e67ea058,
		},
	}
	return PostMultipartForm(__obf_1fc78df60c5789b2,

		// GetFile 下载文件
		__obf_a48e83c31cb47a6e)
}

func GetFile(__obf_04310e04e67ea058, __obf_a48e83c31cb47a6e string) error {
	__obf_32eef76263192e6d, __obf_74916b80241ef1ff := __obf_59c3debac748b990().Get(__obf_a48e83c31cb47a6e)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_32eef76263192e6d.Body.Close()
	__obf_298a7364dedfa681, __obf_74916b80241ef1ff := os.Create(__obf_04310e04e67ea058)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_298a7364dedfa681.Close()
	_, __obf_74916b80241ef1ff = io.Copy(__obf_298a7364dedfa681, __obf_32eef76263192e6d.Body)
	return __obf_74916b80241ef1ff
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_1fc78df60c5789b2 []MultipartFormField, __obf_a48e83c31cb47a6e string) (__obf_5617732a806f07d2 []byte, __obf_74916b80241ef1ff error) {
	__obf_1d7436302b77b85e := &bytes.Buffer{}
	__obf_922d6c80602b9d35 := multipart.NewWriter(__obf_1d7436302b77b85e)

	for _, __obf_33244153a2e3717f := range __obf_1fc78df60c5789b2 {
		if __obf_33244153a2e3717f.IsFile {
			__obf_e2a683542b42b30b, __obf_8f69d04a4f68fa13 := __obf_922d6c80602b9d35.CreateFormFile(__obf_33244153a2e3717f.Fieldname, filepath.Base(__obf_33244153a2e3717f.Filename))
			if __obf_8f69d04a4f68fa13 != nil {
				__obf_74916b80241ef1ff = fmt.Errorf("error writing to buffer , err=%v", __obf_8f69d04a4f68fa13)
				return
			}
			__obf_2daeb36d13e303d8, __obf_8f69d04a4f68fa13 := os.Open(__obf_33244153a2e3717f.Filename)
			if __obf_8f69d04a4f68fa13 != nil {
				__obf_74916b80241ef1ff = fmt.Errorf("error opening file , err=%v", __obf_8f69d04a4f68fa13)
				return
			}
			defer __obf_2daeb36d13e303d8.Close()

			if _, __obf_74916b80241ef1ff = io.Copy(__obf_e2a683542b42b30b, __obf_2daeb36d13e303d8); __obf_74916b80241ef1ff != nil {
				return
			}
		} else {
			__obf_8cc9cf6ab4776a37, __obf_8f69d04a4f68fa13 := __obf_922d6c80602b9d35.CreateFormField(__obf_33244153a2e3717f.Fieldname)
			if __obf_8f69d04a4f68fa13 != nil {
				__obf_74916b80241ef1ff = __obf_8f69d04a4f68fa13
				return
			}
			__obf_7dfc7e2061f8d524 := bytes.NewReader(__obf_33244153a2e3717f.Value)
			if _, __obf_74916b80241ef1ff = io.Copy(__obf_8cc9cf6ab4776a37, __obf_7dfc7e2061f8d524); __obf_74916b80241ef1ff != nil {
				return
			}
		}
	}
	__obf_52eed32857b9a9bc := __obf_922d6c80602b9d35.FormDataContentType()
	__obf_922d6c80602b9d35.
		Close()
	__obf_32eef76263192e6d, __obf_8f69d04a4f68fa13 := __obf_59c3debac748b990().Post(__obf_a48e83c31cb47a6e, __obf_52eed32857b9a9bc, __obf_1d7436302b77b85e)
	if __obf_8f69d04a4f68fa13 != nil {
		__obf_74916b80241ef1ff = __obf_8f69d04a4f68fa13
		return
	}
	defer __obf_32eef76263192e6d.Body.Close()

	if __obf_32eef76263192e6d.StatusCode != http.StatusOK {
		return nil, __obf_74916b80241ef1ff
	}
	return io.ReadAll(__obf_32eef76263192e6d.Body)
}
