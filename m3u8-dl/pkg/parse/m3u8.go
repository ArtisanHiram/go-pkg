package __obf_13e34d0e1e6d5a3a

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type (
	PlaylistType string
	CryptMethod  string
)

const (
	PlaylistTypeVOD   PlaylistType = "VOD"
	PlaylistTypeEvent PlaylistType = "EVENT"

	CryptMethodAES  CryptMethod = "AES-128"
	CryptMethodNONE CryptMethod = "NONE"
)

// extracting 'key = value' regex
var __obf_2d204c1a83dc7daf = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

// M3u8 m3u8 Data struct definition
type M3u8 struct {
	Version        int8              // EXT-X-VERSION:version
	MediaSequence  uint64            // Default 0, #EXT-X-MEDIA-SEQUENCE:sequence
	Segments       []*Segment        // Define a Play List
	MasterPlaylist []*MasterPlaylist // Define a Master Play List
	Keys           map[int]*Key      // Keys for per segment
	EndList        bool              // #EXT-X-ENDLIST
	PlaylistType   PlaylistType      // VOD or EVENT
	TargetDuration float64           // #EXT-X-TARGETDURATION:duration
}

// Segment
// #EXTINF:10.000000,
// 5dd92bfb879c6421d7281c769f0f8c93-4.ts
type Segment struct {
	URI      string
	KeyIndex int
	Title    string  // #EXTINF: duration,<title>
	Duration float32 // #EXTINF: duration,<title>
	Length   uint64  // #EXT-X-BYTERANGE: length[@offset]
	Offset   uint64  // #EXT-X-BYTERANGE: length[@offset]
}

// MasterPlaylist
// #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=240000,RESOLUTION=416x234,CODECS="avc1.42e00a,mp4a.40.2"
type MasterPlaylist struct {
	URI        string
	BandWidth  uint32
	Resolution string
	Codecs     string
	ProgramID  uint32
}

// Key
// #EXT-X-KEY:METHOD=AES-128,URI="key.key"
type Key struct {
	// 'AES-128' or 'NONE'
	// If the encryption method is NONE, the URI and the IV attributes MUST NOT be present
	Method CryptMethod
	URI    string
	IV     string
}

// @param reader
func __obf_13e34d0e1e6d5a3a(__obf_7c6259d5ef161e22 io.Reader) (*M3u8, error) {
	__obf_f912d17e4b4a93ba := bufio.NewScanner(__obf_7c6259d5ef161e22)
	var __obf_8ac279997f31e521 []string
	for __obf_f912d17e4b4a93ba.Scan() {
		__obf_8ac279997f31e521 = append(__obf_8ac279997f31e521, __obf_f912d17e4b4a93ba.Text())
	}

	var (
		__obf_e71f93ee091092d8 = 0
		__obf_8686a243937b4113 = len(__obf_8ac279997f31e521)
		__obf_5f60c6671d122931 = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_8f2722bc76623cdd = 0
		__obf_34815b5c820544ed *Key
		__obf_6581871659ad468c *Segment
		__obf_5a920a364053c1cf bool
		__obf_372dcfcfe36ba24a bool
	)

	for ; __obf_e71f93ee091092d8 < __obf_8686a243937b4113; __obf_e71f93ee091092d8++ {
		__obf_e859567c0f7014ee := strings.TrimSpace(__obf_8ac279997f31e521[__obf_e71f93ee091092d8])
		if __obf_e71f93ee091092d8 == 0 {
			if "#EXTM3U" != __obf_e859567c0f7014ee {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_e859567c0f7014ee == "":
			continue
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_5f60c6671d122931.PlaylistType); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_f34f4eb5cd29a9a3 := __obf_5f60c6671d122931.PlaylistType == "" || __obf_5f60c6671d122931.PlaylistType == PlaylistTypeVOD || __obf_5f60c6671d122931.PlaylistType == PlaylistTypeEvent
			if !__obf_f34f4eb5cd29a9a3 {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_5f60c6671d122931.PlaylistType, __obf_e71f93ee091092d8+1)
			}
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-TARGETDURATION:"):
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXT-X-TARGETDURATION:%f", &__obf_5f60c6671d122931.TargetDuration); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_5f60c6671d122931.MediaSequence); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-VERSION:"):
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXT-X-VERSION:%d", &__obf_5f60c6671d122931.Version); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-STREAM-INF:"):
			__obf_23652a372dd01652, __obf_2480b83375e43a4e := __obf_38f55bab0e4dd3ea(__obf_e859567c0f7014ee)
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_e71f93ee091092d8++
			__obf_23652a372dd01652.
				URI = __obf_8ac279997f31e521[__obf_e71f93ee091092d8]
			if __obf_23652a372dd01652.URI == "" || strings.HasPrefix(__obf_23652a372dd01652.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_e71f93ee091092d8+1)
			}
			__obf_5f60c6671d122931.
				MasterPlaylist = append(__obf_5f60c6671d122931.MasterPlaylist, __obf_23652a372dd01652)
			continue
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXTINF:"):
			if __obf_5a920a364053c1cf {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_e859567c0f7014ee, __obf_e71f93ee091092d8+1)
			}
			if __obf_6581871659ad468c == nil {
				__obf_6581871659ad468c = new(Segment)
			}
			var __obf_f912d17e4b4a93ba string
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXTINF:%s", &__obf_f912d17e4b4a93ba); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			if strings.Contains(__obf_f912d17e4b4a93ba, ",") {
				__obf_39444d873157f285 := strings.Split(__obf_f912d17e4b4a93ba, ",")
				__obf_6581871659ad468c.
					Title = __obf_39444d873157f285[1]
				__obf_f912d17e4b4a93ba = __obf_39444d873157f285[0]
			}
			__obf_7325f62b13e75e2b, __obf_2480b83375e43a4e := strconv.ParseFloat(__obf_f912d17e4b4a93ba, 32)
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_6581871659ad468c.
				Duration = float32(__obf_7325f62b13e75e2b)
			__obf_6581871659ad468c.
				KeyIndex = __obf_8f2722bc76623cdd
			__obf_5a920a364053c1cf = true
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-BYTERANGE:"):
			if __obf_372dcfcfe36ba24a {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_e859567c0f7014ee, __obf_e71f93ee091092d8+1)
			}
			if __obf_6581871659ad468c == nil {
				__obf_6581871659ad468c = new(Segment)
			}
			var __obf_4ace321aacb5e1fc string
			if _, __obf_2480b83375e43a4e := fmt.Sscanf(__obf_e859567c0f7014ee, "#EXT-X-BYTERANGE:%s", &__obf_4ace321aacb5e1fc); __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			if __obf_4ace321aacb5e1fc == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_e71f93ee091092d8+1)
			}
			if strings.Contains(__obf_4ace321aacb5e1fc, "@") {
				__obf_39444d873157f285 := strings.Split(__obf_4ace321aacb5e1fc, "@")
				__obf_5326cc5163392a32, __obf_2480b83375e43a4e := strconv.ParseUint(__obf_39444d873157f285[1], 10, 64)
				if __obf_2480b83375e43a4e != nil {
					return nil, __obf_2480b83375e43a4e
				}
				__obf_6581871659ad468c.
					Offset = uint64(__obf_5326cc5163392a32)
				__obf_4ace321aacb5e1fc = __obf_39444d873157f285[0]
			}
			__obf_3eec249983936482, __obf_2480b83375e43a4e := strconv.ParseUint(__obf_4ace321aacb5e1fc, 10, 64)
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_6581871659ad468c.
				Length = uint64(__obf_3eec249983936482)
			__obf_372dcfcfe36ba24a = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_e859567c0f7014ee, "#"):
			if __obf_5a920a364053c1cf {
				if __obf_6581871659ad468c == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_e859567c0f7014ee)
				}
				__obf_6581871659ad468c.
					URI = __obf_e859567c0f7014ee
				__obf_372dcfcfe36ba24a = false
				__obf_5a920a364053c1cf = false
				__obf_5f60c6671d122931.
					Segments = append(__obf_5f60c6671d122931.Segments, __obf_6581871659ad468c)
				__obf_6581871659ad468c = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_e859567c0f7014ee, "#EXT-X-KEY"):
			__obf_06f10035a9011a83 := __obf_a8c7444ee3045896(__obf_e859567c0f7014ee)
			if len(__obf_06f10035a9011a83) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_e859567c0f7014ee, __obf_e71f93ee091092d8+1)
			}
			__obf_e7245fba48f524aa := CryptMethod(__obf_06f10035a9011a83["METHOD"])
			if __obf_e7245fba48f524aa != "" && __obf_e7245fba48f524aa != CryptMethodAES && __obf_e7245fba48f524aa != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_e7245fba48f524aa, __obf_e71f93ee091092d8+1)
			}
			__obf_8f2722bc76623cdd++
			__obf_34815b5c820544ed = new(Key)
			__obf_34815b5c820544ed.
				Method = __obf_e7245fba48f524aa
			__obf_34815b5c820544ed.
				URI = __obf_06f10035a9011a83["URI"]
			__obf_34815b5c820544ed.
				IV = __obf_06f10035a9011a83["IV"]
			__obf_5f60c6671d122931.
				Keys[__obf_8f2722bc76623cdd] = __obf_34815b5c820544ed
		case __obf_e859567c0f7014ee == "#EndList":
			__obf_5f60c6671d122931.
				EndList = true
		default:
			continue
		}
	}

	return __obf_5f60c6671d122931, nil
}

func __obf_38f55bab0e4dd3ea(__obf_e859567c0f7014ee string) (*MasterPlaylist, error) {
	__obf_06f10035a9011a83 := __obf_a8c7444ee3045896(__obf_e859567c0f7014ee)
	if len(__obf_06f10035a9011a83) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_23652a372dd01652 := new(MasterPlaylist)
	for __obf_0931910163550cb2, __obf_16531318deccf55e := range __obf_06f10035a9011a83 {
		switch {
		case __obf_0931910163550cb2 == "BANDWIDTH":
			__obf_16531318deccf55e, __obf_2480b83375e43a4e := strconv.ParseUint(__obf_16531318deccf55e, 10, 32)
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_23652a372dd01652.
				BandWidth = uint32(__obf_16531318deccf55e)
		case __obf_0931910163550cb2 == "RESOLUTION":
			__obf_23652a372dd01652.
				Resolution = __obf_16531318deccf55e
		case __obf_0931910163550cb2 == "PROGRAM-ID":
			__obf_16531318deccf55e, __obf_2480b83375e43a4e := strconv.ParseUint(__obf_16531318deccf55e, 10, 32)
			if __obf_2480b83375e43a4e != nil {
				return nil, __obf_2480b83375e43a4e
			}
			__obf_23652a372dd01652.
				ProgramID = uint32(__obf_16531318deccf55e)
		case __obf_0931910163550cb2 == "CODECS":
			__obf_23652a372dd01652.
				Codecs = __obf_16531318deccf55e
		}
	}
	return __obf_23652a372dd01652, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_a8c7444ee3045896(__obf_e859567c0f7014ee string) map[string]string {
	__obf_7b0e689392466bf2 := __obf_2d204c1a83dc7daf.FindAllStringSubmatch(__obf_e859567c0f7014ee, -1)
	__obf_06f10035a9011a83 := make(map[string]string)
	for _, __obf_5b4e287cff60ee14 := range __obf_7b0e689392466bf2 {
		__obf_06f10035a9011a83[__obf_5b4e287cff60ee14[1]] = strings.Trim(__obf_5b4e287cff60ee14[2], "\"")
	}
	return __obf_06f10035a9011a83
}
