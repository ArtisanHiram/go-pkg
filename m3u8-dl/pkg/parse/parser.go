package __obf_4f369713f5d562f0

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
func ParseFromNet(__obf_7f6b0c0e1dce0289 string) (*Result, error) {
	__obf_66969c1aa3c70317, __obf_cf22f4b7fd45f7b5 := url.Parse(__obf_7f6b0c0e1dce0289)
	if __obf_cf22f4b7fd45f7b5 != nil {
		return nil, __obf_cf22f4b7fd45f7b5
	}
	__obf_7f7c340462aeb80d := __obf_66969c1aa3c70317.String()
	__obf_4898e8900a78b050, __obf_cf22f4b7fd45f7b5 := tool.Get(__obf_7f7c340462aeb80d)
	if __obf_cf22f4b7fd45f7b5 != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %s", __obf_cf22f4b7fd45f7b5.Error())
	}

	//noinspection GoUnhandledErrorResult
	defer __obf_4898e8900a78b050.Close()
	__obf_569269c4aa626235, __obf_cf22f4b7fd45f7b5 := __obf_4f369713f5d562f0(__obf_4898e8900a78b050)
	if __obf_cf22f4b7fd45f7b5 != nil {
		return nil, __obf_cf22f4b7fd45f7b5
	}
	// 如果是master list,选取第一个作为play list的m3u8地址
	if len(__obf_569269c4aa626235.MasterPlaylist) != 0 {
		__obf_18a9949a3c07b915 := __obf_569269c4aa626235.MasterPlaylist[0]
		return ParseFromNet(tool.ResolveURL(__obf_66969c1aa3c70317, __obf_18a9949a3c07b915.URI))
	}
	if len(__obf_569269c4aa626235.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}

	if len(__obf_569269c4aa626235.Segments) == 0 {
		return nil, errors.New("can not found any TS file description")
	}
	__obf_2d5438eb15f5fac6 := &Result{
		URL:  __obf_66969c1aa3c70317,
		M3u8: __obf_569269c4aa626235,
		Keys: make(map[int]string),
	}

	// resolve keys for encrypt decode
	// 如果ts片段是加密的,获取每个片段的密钥
	for __obf_a36168649c7e1eaf, __obf_d31388392c335969 := range __obf_569269c4aa626235.Keys {
		switch {
		case __obf_d31388392c335969.Method == "" || __obf_d31388392c335969.Method == CryptMethodNONE:
			continue
		case __obf_d31388392c335969.Method == CryptMethodAES:
			__obf_37631d01f30cb866 := // Request URL to extract decryption key
				__obf_d31388392c335969.URI
			__obf_37631d01f30cb866 = tool.ResolveURL(__obf_66969c1aa3c70317, __obf_37631d01f30cb866)
			__obf_c2cfa0f8eb496276, __obf_cf22f4b7fd45f7b5 := tool.Get(__obf_37631d01f30cb866)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, fmt.Errorf("extract key failed: %s", __obf_cf22f4b7fd45f7b5.Error())
			}
			__obf_5f8cfd16f13bece1, __obf_cf22f4b7fd45f7b5 := io.ReadAll(__obf_c2cfa0f8eb496276)
			_ = __obf_c2cfa0f8eb496276.Close()
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			fmt.Println("decryption key: ", string(__obf_5f8cfd16f13bece1))
			__obf_2d5438eb15f5fac6.
				Keys[__obf_a36168649c7e1eaf] = string(__obf_5f8cfd16f13bece1)
		default:
			return nil, fmt.Errorf("unknown or unsupported cryption method: %s", __obf_d31388392c335969.Method)
		}
	}
	return __obf_2d5438eb15f5fac6, nil
}
