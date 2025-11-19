package __obf_083c8deafa73f533

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
func SetTimeOut(__obf_66f560b1a3148c0c time.Duration) {
	TimeOut = __obf_66f560b1a3148c0c
}

// httpClient() 带超时的http.Client
func __obf_4a60dda07f59cbef() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(__obf_22869fd63c164f49 string, __obf_bc97935e99909ed0 any) error {

	__obf_c094a893233159c5, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Get(__obf_22869fd63c164f49)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_c094a893233159c5.Body.Close()
	return __obf_283076f580f99f88.NewDecoder(__obf_c094a893233159c5.Body).Decode(__obf_bc97935e99909ed0)
}

// GetXml 发送GET请求并解析xml
func GetXml(__obf_22869fd63c164f49 string, __obf_bc97935e99909ed0 any) error {
	__obf_c094a893233159c5, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Get(__obf_22869fd63c164f49)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_c094a893233159c5.Body.Close()
	return xml.NewDecoder(__obf_c094a893233159c5.Body).Decode(__obf_bc97935e99909ed0)
}

// GetBody 发送GET请求，返回body字节
func GetBody(__obf_22869fd63c164f49 string) ([]byte, error) {
	__obf_1aeba3af5526ef6d, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Get(__obf_22869fd63c164f49)
	if __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()

	if __obf_1aeba3af5526ef6d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", __obf_22869fd63c164f49, __obf_1aeba3af5526ef6d.StatusCode)
	}
	return io.ReadAll(__obf_1aeba3af5526ef6d.Body)
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
func PostJson(__obf_22869fd63c164f49 string, __obf_34b84ecd45b601a3 any) ([]byte, error) {
	__obf_28b8108c6efefd30 := new(bytes.Buffer)
	__obf_207f5f9f1a2e7b5e := __obf_283076f580f99f88.NewEncoder(__obf_28b8108c6efefd30)
	__obf_207f5f9f1a2e7b5e.SetEscapeHTML(false)
	__obf_ab078048114898aa := __obf_207f5f9f1a2e7b5e.Encode(__obf_34b84ecd45b601a3)
	if __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}
	__obf_1aeba3af5526ef6d, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Post(__obf_22869fd63c164f49, "application/json;charset=utf-8", __obf_28b8108c6efefd30)
	if __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()

	if __obf_1aeba3af5526ef6d.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_22869fd63c164f49, __obf_1aeba3af5526ef6d.StatusCode)
	}
	return io.ReadAll(__obf_1aeba3af5526ef6d.Body)
}

// PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(__obf_22869fd63c164f49 string, __obf_34b84ecd45b601a3 any, __obf_fae510501b39e125 any, __obf_71b895a9ee0308a1 ...string) (__obf_ab078048114898aa error) {
	__obf_28b8108c6efefd30 := new(bytes.Buffer)
	__obf_207f5f9f1a2e7b5e := __obf_283076f580f99f88.NewEncoder(__obf_28b8108c6efefd30)
	//	enc.SetEscapeHTML(false)
	__obf_ab078048114898aa = __obf_207f5f9f1a2e7b5e.Encode(__obf_34b84ecd45b601a3)
	if __obf_ab078048114898aa != nil {
		return
	}
	__obf_3ca498a986584e5b := "application/json;charset=utf-8"
	if len(__obf_71b895a9ee0308a1) > 0 {
		__obf_3ca498a986584e5b = strings.Join(__obf_71b895a9ee0308a1, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	__obf_1aeba3af5526ef6d, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Post(__obf_22869fd63c164f49, __obf_3ca498a986584e5b, __obf_28b8108c6efefd30)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()

	if __obf_1aeba3af5526ef6d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_22869fd63c164f49, __obf_1aeba3af5526ef6d.StatusCode)
	}
	return __obf_283076f580f99f88.NewDecoder(__obf_1aeba3af5526ef6d.Body).Decode(__obf_fae510501b39e125)
}

// PostXmlPtr 发送Xml格式的POST请求并解析结果到result指针
func PostXmlPtr(__obf_22869fd63c164f49 string, __obf_34b84ecd45b601a3 any, __obf_fae510501b39e125 any) (__obf_ab078048114898aa error) {
	__obf_28b8108c6efefd30 := new(bytes.Buffer)
	__obf_207f5f9f1a2e7b5e := xml.NewEncoder(__obf_28b8108c6efefd30)
	//	enc.SetEscapeHTML(false)
	__obf_ab078048114898aa = __obf_207f5f9f1a2e7b5e.Encode(__obf_34b84ecd45b601a3)
	if __obf_ab078048114898aa != nil {
		return
	}

	__obf_1aeba3af5526ef6d, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Post(__obf_22869fd63c164f49, "application/xml;charset=utf-8", __obf_28b8108c6efefd30)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()

	if __obf_1aeba3af5526ef6d.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", __obf_22869fd63c164f49, __obf_1aeba3af5526ef6d.StatusCode)
	}
	return xml.NewDecoder(__obf_1aeba3af5526ef6d.Body).Decode(__obf_fae510501b39e125)
}

// PostFile 上传文件
func PostFile(__obf_c447bc0f2c0d0335, __obf_efd6dcc9831c56ca, __obf_22869fd63c164f49 string) ([]byte, error) {
	__obf_f489e5affddeda62 := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: __obf_c447bc0f2c0d0335,
			Filename:  __obf_efd6dcc9831c56ca,
		},
	}
	return PostMultipartForm(__obf_f489e5affddeda62, __obf_22869fd63c164f49)
}

// GetFile 下载文件
func GetFile(__obf_efd6dcc9831c56ca, __obf_22869fd63c164f49 string) error {
	__obf_1aeba3af5526ef6d, __obf_ab078048114898aa := __obf_4a60dda07f59cbef().Get(__obf_22869fd63c164f49)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()
	__obf_cddb0fc0ea27c776, __obf_ab078048114898aa := os.Create(__obf_efd6dcc9831c56ca)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_cddb0fc0ea27c776.Close()
	_, __obf_ab078048114898aa = io.Copy(__obf_cddb0fc0ea27c776, __obf_1aeba3af5526ef6d.Body)
	return __obf_ab078048114898aa
}

// MultipartFormField 文件或其他表单数据
type MultipartFormField struct {
	Fieldname string
	Filename  string
	Value     []byte
	IsFile    bool
}

// PostMultipartForm 上传文件或其他表单数据
func PostMultipartForm(__obf_f489e5affddeda62 []MultipartFormField, __obf_22869fd63c164f49 string) (__obf_0479429e7c43e116 []byte, __obf_ab078048114898aa error) {
	__obf_d95aaf2b0df07e00 := &bytes.Buffer{}
	__obf_64e83fc8b561be5f := multipart.NewWriter(__obf_d95aaf2b0df07e00)

	for _, __obf_17f9fa2b16b7f147 := range __obf_f489e5affddeda62 {
		if __obf_17f9fa2b16b7f147.IsFile {
			__obf_cb0e3c94f7a2749b, __obf_90ec48ccc47fd0ef := __obf_64e83fc8b561be5f.CreateFormFile(__obf_17f9fa2b16b7f147.Fieldname, filepath.Base(__obf_17f9fa2b16b7f147.Filename))
			if __obf_90ec48ccc47fd0ef != nil {
				__obf_ab078048114898aa = fmt.Errorf("error writing to buffer , err=%v", __obf_90ec48ccc47fd0ef)
				return
			}

			__obf_ae6e420b43b759e6, __obf_90ec48ccc47fd0ef := os.Open(__obf_17f9fa2b16b7f147.Filename)
			if __obf_90ec48ccc47fd0ef != nil {
				__obf_ab078048114898aa = fmt.Errorf("error opening file , err=%v", __obf_90ec48ccc47fd0ef)
				return
			}
			defer __obf_ae6e420b43b759e6.Close()

			if _, __obf_ab078048114898aa = io.Copy(__obf_cb0e3c94f7a2749b, __obf_ae6e420b43b759e6); __obf_ab078048114898aa != nil {
				return
			}
		} else {
			__obf_f9037ebc2e148dc5, __obf_90ec48ccc47fd0ef := __obf_64e83fc8b561be5f.CreateFormField(__obf_17f9fa2b16b7f147.Fieldname)
			if __obf_90ec48ccc47fd0ef != nil {
				__obf_ab078048114898aa = __obf_90ec48ccc47fd0ef
				return
			}
			__obf_31ed8c399bb554e9 := bytes.NewReader(__obf_17f9fa2b16b7f147.Value)
			if _, __obf_ab078048114898aa = io.Copy(__obf_f9037ebc2e148dc5, __obf_31ed8c399bb554e9); __obf_ab078048114898aa != nil {
				return
			}
		}
	}

	__obf_71b895a9ee0308a1 := __obf_64e83fc8b561be5f.FormDataContentType()
	__obf_64e83fc8b561be5f.Close()

	__obf_1aeba3af5526ef6d, __obf_90ec48ccc47fd0ef := __obf_4a60dda07f59cbef().Post(__obf_22869fd63c164f49, __obf_71b895a9ee0308a1, __obf_d95aaf2b0df07e00)
	if __obf_90ec48ccc47fd0ef != nil {
		__obf_ab078048114898aa = __obf_90ec48ccc47fd0ef
		return
	}
	defer __obf_1aeba3af5526ef6d.Body.Close()

	if __obf_1aeba3af5526ef6d.StatusCode != http.StatusOK {
		return nil, __obf_ab078048114898aa
	}
	return io.ReadAll(__obf_1aeba3af5526ef6d.Body)
}
