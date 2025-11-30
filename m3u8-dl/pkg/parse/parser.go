package __obf_495e3fa4e37cbc01

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
func ParseFromNet(__obf_424d003367187afc string) (*Result, error) {
	__obf_854d2316a28ac396, __obf_8a71cc83b0a472cc := url.Parse(__obf_424d003367187afc)
	if __obf_8a71cc83b0a472cc != nil {
		return nil, __obf_8a71cc83b0a472cc
	}
	__obf_9a5762a7f439139c := __obf_854d2316a28ac396.String()
	__obf_68f04193bbe075a6, __obf_8a71cc83b0a472cc := tool.Get(__obf_9a5762a7f439139c)
	if __obf_8a71cc83b0a472cc != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_8a71cc83b0a472cc.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_68f04193bbe075a6.Close()
	__obf_100c59b734f73232, __obf_8a71cc83b0a472cc := __obf_495e3fa4e37cbc01(__obf_68f04193bbe075a6)
	if __obf_8a71cc83b0a472cc != nil {
		return nil, __obf_8a71cc83b0a472cc
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_100c59b734f73232.MasterPlaylist) != 0 {
		__obf_c1af08b353089ba8 := __obf_100c59b734f73232.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_854d2316a28ac396, __obf_c1af08b353089ba8.URI))
	}
	if len(__obf_100c59b734f73232.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_100c59b734f73232.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_21f3cce4528f0bd1 := &Result{
		URL:  __obf_854d2316a28ac396,
		M3u8: __obf_100c59b734f73232,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_c33eaea796380b3e, __obf_ec75b9484ec3cd6d := range __obf_100c59b734f73232.Keys {
		switch {
		case __obf_ec75b9484ec3cd6d.Method == "" || __obf_ec75b9484ec3cd6d.Method == CryptMethodNONE:
			continue
		case __obf_ec75b9484ec3cd6d.Method == CryptMethodAES:
			__obf_57b07bd8257cb223 := // Request URL to extract decryption key
				__obf_ec75b9484ec3cd6d.URI
			__obf_57b07bd8257cb223 = tool.ResolveURL(__obf_854d2316a28ac396, __obf_57b07bd8257cb223)
			__obf_b999b3c86260297d, __obf_8a71cc83b0a472cc := tool.Get(__obf_57b07bd8257cb223)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_8a71cc83b0a472cc.Error())
			}
			__obf_fab70e7b25d2e0c1, __obf_8a71cc83b0a472cc := io.ReadAll(__obf_b999b3c86260297d)
			_ = __obf_b999b3c86260297d.Close()
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			fmt.Println("decryption key: ", string(__obf_fab70e7b25d2e0c1))
			__obf_21f3cce4528f0bd1.
				Keys[__obf_c33eaea796380b3e] = string(__obf_fab70e7b25d2e0c1)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_ec75b9484ec3cd6d.Method)
		}
	}
	return __obf_21f3cce4528f0bd1, nil
}
