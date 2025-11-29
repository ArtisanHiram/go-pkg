package __obf_76e668a60a544837

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
func ParseFromNet(__obf_144f2a1b68ec0260 string) (*Result, error) {
	__obf_25b2c41890becc61, __obf_c440f080fd9d531f := url.Parse(__obf_144f2a1b68ec0260)
	if __obf_c440f080fd9d531f != nil {
		return nil, __obf_c440f080fd9d531f
	}
	__obf_bee355b491697956 := __obf_25b2c41890becc61.String()
	__obf_d0f76b271988266e, __obf_c440f080fd9d531f := tool.Get(__obf_bee355b491697956)
	if __obf_c440f080fd9d531f != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_c440f080fd9d531f.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_d0f76b271988266e.Close()
	__obf_2a9efb22c878897d, __obf_c440f080fd9d531f := __obf_76e668a60a544837(__obf_d0f76b271988266e)
	if __obf_c440f080fd9d531f != nil {
		return nil, __obf_c440f080fd9d531f
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_2a9efb22c878897d.MasterPlaylist) != 0 {
		__obf_09982148888cc64c := __obf_2a9efb22c878897d.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_25b2c41890becc61, __obf_09982148888cc64c.URI))
	}
	if len(__obf_2a9efb22c878897d.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_2a9efb22c878897d.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_aa25cbf5f4a8bdd9 := &Result{
		URL:  __obf_25b2c41890becc61,
		M3u8: __obf_2a9efb22c878897d,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_5c2155159514e979, __obf_1ffa1a359ec1627b := range __obf_2a9efb22c878897d.Keys {
		switch {
		case __obf_1ffa1a359ec1627b.Method == "" || __obf_1ffa1a359ec1627b.Method == CryptMethodNONE:
			continue
		case __obf_1ffa1a359ec1627b.Method == CryptMethodAES:
			__obf_70d8f11c90a6b0e6 := // Request URL to extract decryption key
				__obf_1ffa1a359ec1627b.URI
			__obf_70d8f11c90a6b0e6 = tool.ResolveURL(__obf_25b2c41890becc61, __obf_70d8f11c90a6b0e6)
			__obf_b206fe59ae3a773a, __obf_c440f080fd9d531f := tool.Get(__obf_70d8f11c90a6b0e6)
			if __obf_c440f080fd9d531f != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_c440f080fd9d531f.Error())
			}
			__obf_c242f0f516f0365d, __obf_c440f080fd9d531f := io.ReadAll(__obf_b206fe59ae3a773a)
			_ = __obf_b206fe59ae3a773a.Close()
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			fmt.Println("decryption key: ", string(__obf_c242f0f516f0365d))
			__obf_aa25cbf5f4a8bdd9.
				Keys[__obf_5c2155159514e979] = string(__obf_c242f0f516f0365d)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_1ffa1a359ec1627b.Method)
		}
	}
	return __obf_aa25cbf5f4a8bdd9, nil
}
