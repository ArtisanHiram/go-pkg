package __obf_13e34d0e1e6d5a3a

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
func ParseFromNet(__obf_5950f83ade32dbb7 string) (*Result, error) {
	__obf_ead33d6d0f506de0, __obf_2480b83375e43a4e := url.Parse(__obf_5950f83ade32dbb7)
	if __obf_2480b83375e43a4e != nil {
		return nil, __obf_2480b83375e43a4e
	}
	__obf_33823dbc6e203e70 := __obf_ead33d6d0f506de0.String()
	__obf_80756e70f566fc22, __obf_2480b83375e43a4e := tool.Get(__obf_33823dbc6e203e70)
	if __obf_2480b83375e43a4e != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_2480b83375e43a4e.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_80756e70f566fc22.Close()
	__obf_5f60c6671d122931, __obf_2480b83375e43a4e := __obf_13e34d0e1e6d5a3a(__obf_80756e70f566fc22)
	if __obf_2480b83375e43a4e != nil {
		return nil, __obf_2480b83375e43a4e
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_5f60c6671d122931.MasterPlaylist) != 0 {
		__obf_fdcdca0bf177957c := __obf_5f60c6671d122931.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_ead33d6d0f506de0, __obf_fdcdca0bf177957c.URI))
	}
	if len(__obf_5f60c6671d122931.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_5f60c6671d122931.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_f48401f58e77b52a := &Result{
		URL:  __obf_ead33d6d0f506de0,
		M3u8: __obf_5f60c6671d122931,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_8e1243c2c520ee5b, __obf_34815b5c820544ed := range __obf_5f60c6671d122931.Keys {
		switch {
		case __obf_34815b5c820544ed.Method == "" || __obf_34815b5c820544ed.Method == CryptMethodNONE:
			continue
		case __obf_34815b5c820544ed.Method == CryptMethodAES:
			__obf_9fd123b78fb9229f := // Request URL to extract decryption key
				__obf_34815b5c820544ed.URI
			__obf_9fd123b78fb9229f = tool.ResolveURL(__obf_ead33d6d0f506de0, __obf_9fd123b78fb9229f)
			__obf_220585d665ce0dfb, __obf_2480b83375e43a4e := tool.Get(__obf_9fd123b78fb9229f)
			if __obf_2480b83375e43a4e != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_2480b83375e43a4e.Error())
			}
			__obf_081d38bf5f4fbad6, __obf_2480b83375e43a4e := io.ReadAll(__obf_220585d665ce0dfb)
			_ = __obf_220585d665ce0dfb.Close()
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			fmt.Println("decryption key: ", string(__obf_081d38bf5f4fbad6))
			__obf_f48401f58e77b52a.
				Keys[__obf_8e1243c2c520ee5b] = string(__obf_081d38bf5f4fbad6)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_34815b5c820544ed.Method)
		}
	}
	return __obf_f48401f58e77b52a, nil
}
