package __obf_73049aff9f20feb6

import (
	"errors"
	"fmt"
	tool "github.com/ArtisanHiram/go-pkg/m3u8-dl/pkg/tool"
	"io"
	"net/url"
)

type Result struct {
	URL  *url.URL
	M3u8 *M3u8
	Keys map[int]string
}

// ParseFromNet Gets parse result from a url link
func ParseFromNet(__obf_4bc35b7dc9e02492 string) (*Result, error) {
	__obf_f749a15dcf21f1ea, __obf_5159af91c7c96fc8 := url.Parse(__obf_4bc35b7dc9e02492)
	if __obf_5159af91c7c96fc8 != nil {
		return nil, __obf_5159af91c7c96fc8
	}
	__obf_7517825490f20536 := __obf_f749a15dcf21f1ea.String()
	__obf_cdd2fcb7e7d58d95, __obf_5159af91c7c96fc8 := tool.Get(__obf_7517825490f20536)
	if __obf_5159af91c7c96fc8 != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_5159af91c7c96fc8.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_cdd2fcb7e7d58d95.Close()
	__obf_bae86baf40186d8d, __obf_5159af91c7c96fc8 := __obf_73049aff9f20feb6(__obf_cdd2fcb7e7d58d95)
	if __obf_5159af91c7c96fc8 != nil {
		return nil, __obf_5159af91c7c96fc8
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_bae86baf40186d8d.MasterPlaylist) != 0 {
		__obf_1fb58ee6220c7aaf := __obf_bae86baf40186d8d.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_f749a15dcf21f1ea, __obf_1fb58ee6220c7aaf.URI))
	}
	if len(__obf_bae86baf40186d8d.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_bae86baf40186d8d.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_306cf1600e4eb096 := &Result{
		URL:  __obf_f749a15dcf21f1ea,
		M3u8: __obf_bae86baf40186d8d,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_1555b6d5a4bc759e, __obf_e6b5c32ce5643ea2 := range __obf_bae86baf40186d8d.Keys {
		switch {
		case __obf_e6b5c32ce5643ea2.Method == "" || __obf_e6b5c32ce5643ea2.Method == CryptMethodNONE:
			continue
		case __obf_e6b5c32ce5643ea2.Method == CryptMethodAES:
			__obf_836bd3396ec007a4 := // Request URL to extract decryption key
				__obf_e6b5c32ce5643ea2.URI
			__obf_836bd3396ec007a4 = tool.ResolveURL(__obf_f749a15dcf21f1ea, __obf_836bd3396ec007a4)
			__obf_a2395d2f3018b4f8, __obf_5159af91c7c96fc8 := tool.Get(__obf_836bd3396ec007a4)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_5159af91c7c96fc8.Error())
			}
			__obf_b8b98f97cb76c6c3, __obf_5159af91c7c96fc8 := io.ReadAll(__obf_a2395d2f3018b4f8)
			_ = __obf_a2395d2f3018b4f8.Close()
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			fmt.Println("decryption key: ", string(__obf_b8b98f97cb76c6c3))
			__obf_306cf1600e4eb096.
				Keys[__obf_1555b6d5a4bc759e] = string(__obf_b8b98f97cb76c6c3)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_e6b5c32ce5643ea2.Method)
		}
	}
	return __obf_306cf1600e4eb096, nil
}
