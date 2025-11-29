package __obf_ce58dca6c0b3a695

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
func ParseFromNet(__obf_770635dcfa23a864 string) (*Result, error) {
	__obf_13f9e980919d3fa3, __obf_2f034306c99cf38e := url.Parse(__obf_770635dcfa23a864)
	if __obf_2f034306c99cf38e != nil {
		return nil, __obf_2f034306c99cf38e
	}
	__obf_6cf539b72528b2fb := __obf_13f9e980919d3fa3.String()
	__obf_55bb0e070d36be7f, __obf_2f034306c99cf38e := tool.Get(__obf_6cf539b72528b2fb)
	if __obf_2f034306c99cf38e != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_2f034306c99cf38e.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_55bb0e070d36be7f.Close()
	__obf_4d6af585f15ca9ea, __obf_2f034306c99cf38e := __obf_ce58dca6c0b3a695(__obf_55bb0e070d36be7f)
	if __obf_2f034306c99cf38e != nil {
		return nil, __obf_2f034306c99cf38e
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_4d6af585f15ca9ea.MasterPlaylist) != 0 {
		__obf_6faa5a7352a16b5c := __obf_4d6af585f15ca9ea.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_13f9e980919d3fa3, __obf_6faa5a7352a16b5c.URI))
	}
	if len(__obf_4d6af585f15ca9ea.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_4d6af585f15ca9ea.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_c8e6a3deb67f354a := &Result{
		URL:  __obf_13f9e980919d3fa3,
		M3u8: __obf_4d6af585f15ca9ea,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_94957c36f5a12be8, __obf_8a7802f8f4246380 := range __obf_4d6af585f15ca9ea.Keys {
		switch {
		case __obf_8a7802f8f4246380.Method == "" || __obf_8a7802f8f4246380.Method == CryptMethodNONE:
			continue
		case __obf_8a7802f8f4246380.Method == CryptMethodAES:
			__obf_602ad791e3e025c4 := // Request URL to extract decryption key
				__obf_8a7802f8f4246380.URI
			__obf_602ad791e3e025c4 = tool.ResolveURL(__obf_13f9e980919d3fa3, __obf_602ad791e3e025c4)
			__obf_3ae3eb18a73a414e, __obf_2f034306c99cf38e := tool.Get(__obf_602ad791e3e025c4)
			if __obf_2f034306c99cf38e != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_2f034306c99cf38e.Error())
			}
			__obf_e4418a4ba326c93c, __obf_2f034306c99cf38e := io.ReadAll(__obf_3ae3eb18a73a414e)
			_ = __obf_3ae3eb18a73a414e.Close()
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			fmt.Println("decryption key: ", string(__obf_e4418a4ba326c93c))
			__obf_c8e6a3deb67f354a.
				Keys[__obf_94957c36f5a12be8] = string(__obf_e4418a4ba326c93c)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_8a7802f8f4246380.Method)
		}
	}
	return __obf_c8e6a3deb67f354a, nil
}
