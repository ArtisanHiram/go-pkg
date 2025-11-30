package __obf_495e3fa4e37cbc01

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
var __obf_ae4e82da1e024f10 = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

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
func __obf_495e3fa4e37cbc01(__obf_8f9d1185da9d6b40 io.Reader) (*M3u8, error) {
	__obf_b26f8fdaad3082a4 := bufio.NewScanner(__obf_8f9d1185da9d6b40)
	var __obf_a45d024102ac4cbc []string
	for __obf_b26f8fdaad3082a4.Scan() {
		__obf_a45d024102ac4cbc = append(__obf_a45d024102ac4cbc, __obf_b26f8fdaad3082a4.Text())
	}

	var (
		__obf_6ba5304fa75e503f = 0
		__obf_f7b09567bfd3ef46 = len(__obf_a45d024102ac4cbc)
		__obf_100c59b734f73232 = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_bf3a234884f6d894 = 0
		__obf_ec75b9484ec3cd6d *Key
		__obf_b91995eceaf7face *Segment
		__obf_c88f8b2de1baf29a bool
		__obf_f428be08182f889b bool
	)

	for ; __obf_6ba5304fa75e503f < __obf_f7b09567bfd3ef46; __obf_6ba5304fa75e503f++ {
		__obf_f7e4e1fd9ffc6d7a := strings.TrimSpace(__obf_a45d024102ac4cbc[__obf_6ba5304fa75e503f])
		if __obf_6ba5304fa75e503f == 0 {
			if "#EXTM3U" != __obf_f7e4e1fd9ffc6d7a {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_f7e4e1fd9ffc6d7a == "":
			continue
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_100c59b734f73232.PlaylistType); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_62469eb956550dc7 := __obf_100c59b734f73232.PlaylistType == "" || __obf_100c59b734f73232.PlaylistType == PlaylistTypeVOD || __obf_100c59b734f73232.PlaylistType == PlaylistTypeEvent
			if !__obf_62469eb956550dc7 {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_100c59b734f73232.PlaylistType, __obf_6ba5304fa75e503f+1)
			}
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-TARGETDURATION:"):
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-TARGETDURATION:%f", &__obf_100c59b734f73232.TargetDuration); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_100c59b734f73232.MediaSequence); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-VERSION:"):
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-VERSION:%d", &__obf_100c59b734f73232.Version); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-STREAM-INF:"):
			__obf_a3a2150e7a218e90, __obf_8a71cc83b0a472cc := __obf_46432c80ebf262c5(__obf_f7e4e1fd9ffc6d7a)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_6ba5304fa75e503f++
			__obf_a3a2150e7a218e90.
				URI = __obf_a45d024102ac4cbc[__obf_6ba5304fa75e503f]
			if __obf_a3a2150e7a218e90.URI == "" || strings.HasPrefix(__obf_a3a2150e7a218e90.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_6ba5304fa75e503f+1)
			}
			__obf_100c59b734f73232.
				MasterPlaylist = append(__obf_100c59b734f73232.MasterPlaylist, __obf_a3a2150e7a218e90)
			continue
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXTINF:"):
			if __obf_c88f8b2de1baf29a {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_f7e4e1fd9ffc6d7a, __obf_6ba5304fa75e503f+1)
			}
			if __obf_b91995eceaf7face == nil {
				__obf_b91995eceaf7face = new(Segment)
			}
			var __obf_b26f8fdaad3082a4 string
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXTINF:%s", &__obf_b26f8fdaad3082a4); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			if strings.Contains(__obf_b26f8fdaad3082a4, ",") {
				__obf_64cb7e4f69689d6f := strings.Split(__obf_b26f8fdaad3082a4, ",")
				__obf_b91995eceaf7face.
					Title = __obf_64cb7e4f69689d6f[1]
				__obf_b26f8fdaad3082a4 = __obf_64cb7e4f69689d6f[0]
			}
			__obf_6378b0c7568693d2, __obf_8a71cc83b0a472cc := strconv.ParseFloat(__obf_b26f8fdaad3082a4, 32)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_b91995eceaf7face.
				Duration = float32(__obf_6378b0c7568693d2)
			__obf_b91995eceaf7face.
				KeyIndex = __obf_bf3a234884f6d894
			__obf_c88f8b2de1baf29a = true
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-BYTERANGE:"):
			if __obf_f428be08182f889b {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_f7e4e1fd9ffc6d7a, __obf_6ba5304fa75e503f+1)
			}
			if __obf_b91995eceaf7face == nil {
				__obf_b91995eceaf7face = new(Segment)
			}
			var __obf_c2680eb4a1a7d243 string
			if _, __obf_8a71cc83b0a472cc := fmt.Sscanf(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-BYTERANGE:%s", &__obf_c2680eb4a1a7d243); __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			if __obf_c2680eb4a1a7d243 == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_6ba5304fa75e503f+1)
			}
			if strings.Contains(__obf_c2680eb4a1a7d243, "@") {
				__obf_64cb7e4f69689d6f := strings.Split(__obf_c2680eb4a1a7d243, "@")
				__obf_48cad7a75c785b03, __obf_8a71cc83b0a472cc := strconv.ParseUint(__obf_64cb7e4f69689d6f[1], 10, 64)
				if __obf_8a71cc83b0a472cc != nil {
					return nil, __obf_8a71cc83b0a472cc
				}
				__obf_b91995eceaf7face.
					Offset = uint64(__obf_48cad7a75c785b03)
				__obf_c2680eb4a1a7d243 = __obf_64cb7e4f69689d6f[0]
			}
			__obf_833c03aa9ed67cc4, __obf_8a71cc83b0a472cc := strconv.ParseUint(__obf_c2680eb4a1a7d243, 10, 64)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_b91995eceaf7face.
				Length = uint64(__obf_833c03aa9ed67cc4)
			__obf_f428be08182f889b = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#"):
			if __obf_c88f8b2de1baf29a {
				if __obf_b91995eceaf7face == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_f7e4e1fd9ffc6d7a)
				}
				__obf_b91995eceaf7face.
					URI = __obf_f7e4e1fd9ffc6d7a
				__obf_f428be08182f889b = false
				__obf_c88f8b2de1baf29a = false
				__obf_100c59b734f73232.
					Segments = append(__obf_100c59b734f73232.Segments, __obf_b91995eceaf7face)
				__obf_b91995eceaf7face = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_f7e4e1fd9ffc6d7a, "#EXT-X-KEY"):
			__obf_f8b91a58c46a06da := __obf_59be43baad2c5d6c(__obf_f7e4e1fd9ffc6d7a)
			if len(__obf_f8b91a58c46a06da) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_f7e4e1fd9ffc6d7a, __obf_6ba5304fa75e503f+1)
			}
			__obf_383067cac3a629f3 := CryptMethod(__obf_f8b91a58c46a06da["METHOD"])
			if __obf_383067cac3a629f3 != "" && __obf_383067cac3a629f3 != CryptMethodAES && __obf_383067cac3a629f3 != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_383067cac3a629f3, __obf_6ba5304fa75e503f+1)
			}
			__obf_bf3a234884f6d894++
			__obf_ec75b9484ec3cd6d = new(Key)
			__obf_ec75b9484ec3cd6d.
				Method = __obf_383067cac3a629f3
			__obf_ec75b9484ec3cd6d.
				URI = __obf_f8b91a58c46a06da["URI"]
			__obf_ec75b9484ec3cd6d.
				IV = __obf_f8b91a58c46a06da["IV"]
			__obf_100c59b734f73232.
				Keys[__obf_bf3a234884f6d894] = __obf_ec75b9484ec3cd6d
		case __obf_f7e4e1fd9ffc6d7a == "#EndList":
			__obf_100c59b734f73232.
				EndList = true
		default:
			continue
		}
	}

	return __obf_100c59b734f73232, nil
}

func __obf_46432c80ebf262c5(__obf_f7e4e1fd9ffc6d7a string) (*MasterPlaylist, error) {
	__obf_f8b91a58c46a06da := __obf_59be43baad2c5d6c(__obf_f7e4e1fd9ffc6d7a)
	if len(__obf_f8b91a58c46a06da) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_a3a2150e7a218e90 := new(MasterPlaylist)
	for __obf_116d2a087730814c, __obf_f0b8e8f53d37cf17 := range __obf_f8b91a58c46a06da {
		switch {
		case __obf_116d2a087730814c == "BANDWIDTH":
			__obf_f0b8e8f53d37cf17, __obf_8a71cc83b0a472cc := strconv.ParseUint(__obf_f0b8e8f53d37cf17, 10, 32)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_a3a2150e7a218e90.
				BandWidth = uint32(__obf_f0b8e8f53d37cf17)
		case __obf_116d2a087730814c == "RESOLUTION":
			__obf_a3a2150e7a218e90.
				Resolution = __obf_f0b8e8f53d37cf17
		case __obf_116d2a087730814c == "PROGRAM-ID":
			__obf_f0b8e8f53d37cf17, __obf_8a71cc83b0a472cc := strconv.ParseUint(__obf_f0b8e8f53d37cf17, 10, 32)
			if __obf_8a71cc83b0a472cc != nil {
				return nil, __obf_8a71cc83b0a472cc
			}
			__obf_a3a2150e7a218e90.
				ProgramID = uint32(__obf_f0b8e8f53d37cf17)
		case __obf_116d2a087730814c == "CODECS":
			__obf_a3a2150e7a218e90.
				Codecs = __obf_f0b8e8f53d37cf17
		}
	}
	return __obf_a3a2150e7a218e90, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_59be43baad2c5d6c(__obf_f7e4e1fd9ffc6d7a string) map[string]string {
	__obf_d32b868d604554a6 := __obf_ae4e82da1e024f10.FindAllStringSubmatch(__obf_f7e4e1fd9ffc6d7a, -1)
	__obf_f8b91a58c46a06da := make(map[string]string)
	for _, __obf_613594723f7db7b1 := range __obf_d32b868d604554a6 {
		__obf_f8b91a58c46a06da[__obf_613594723f7db7b1[1]] = strings.Trim(__obf_613594723f7db7b1[2], "\"")
	}
	return __obf_f8b91a58c46a06da
}
