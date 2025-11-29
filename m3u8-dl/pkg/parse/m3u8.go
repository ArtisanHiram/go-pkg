package __obf_76e668a60a544837

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
var __obf_40c5d667d0dd16f2 = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

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
func __obf_76e668a60a544837(__obf_b6e8ba7d996185bc io.Reader) (*M3u8, error) {
	__obf_7ac355206ce4b581 := bufio.NewScanner(__obf_b6e8ba7d996185bc)
	var __obf_213673797538bb6e []string
	for __obf_7ac355206ce4b581.Scan() {
		__obf_213673797538bb6e = append(__obf_213673797538bb6e, __obf_7ac355206ce4b581.Text())
	}

	var (
		__obf_db79f8eeae8fdb44 = 0
		__obf_0790e8b53cee017a = len(__obf_213673797538bb6e)
		__obf_2a9efb22c878897d = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_38a325c3f435b257 = 0
		__obf_1ffa1a359ec1627b *Key
		__obf_fc3c7e67ebb963cb *Segment
		__obf_20b81692820a4a6e bool
		__obf_a3b0d71ddb9fef57 bool
	)

	for ; __obf_db79f8eeae8fdb44 < __obf_0790e8b53cee017a; __obf_db79f8eeae8fdb44++ {
		__obf_4e6c5f648da55012 := strings.TrimSpace(__obf_213673797538bb6e[__obf_db79f8eeae8fdb44])
		if __obf_db79f8eeae8fdb44 == 0 {
			if "#EXTM3U" != __obf_4e6c5f648da55012 {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_4e6c5f648da55012 == "":
			continue
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_2a9efb22c878897d.PlaylistType); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_7ed35f45e3d405d2 := __obf_2a9efb22c878897d.PlaylistType == "" || __obf_2a9efb22c878897d.PlaylistType == PlaylistTypeVOD || __obf_2a9efb22c878897d.PlaylistType == PlaylistTypeEvent
			if !__obf_7ed35f45e3d405d2 {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_2a9efb22c878897d.PlaylistType, __obf_db79f8eeae8fdb44+1)
			}
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-TARGETDURATION:"):
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXT-X-TARGETDURATION:%f", &__obf_2a9efb22c878897d.TargetDuration); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_2a9efb22c878897d.MediaSequence); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-VERSION:"):
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXT-X-VERSION:%d", &__obf_2a9efb22c878897d.Version); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-STREAM-INF:"):
			__obf_e6dce7476e25508a, __obf_c440f080fd9d531f := __obf_2cd064455919f542(__obf_4e6c5f648da55012)
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_db79f8eeae8fdb44++
			__obf_e6dce7476e25508a.
				URI = __obf_213673797538bb6e[__obf_db79f8eeae8fdb44]
			if __obf_e6dce7476e25508a.URI == "" || strings.HasPrefix(__obf_e6dce7476e25508a.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_db79f8eeae8fdb44+1)
			}
			__obf_2a9efb22c878897d.
				MasterPlaylist = append(__obf_2a9efb22c878897d.MasterPlaylist, __obf_e6dce7476e25508a)
			continue
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXTINF:"):
			if __obf_20b81692820a4a6e {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_4e6c5f648da55012, __obf_db79f8eeae8fdb44+1)
			}
			if __obf_fc3c7e67ebb963cb == nil {
				__obf_fc3c7e67ebb963cb = new(Segment)
			}
			var __obf_7ac355206ce4b581 string
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXTINF:%s", &__obf_7ac355206ce4b581); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			if strings.Contains(__obf_7ac355206ce4b581, ",") {
				__obf_73b6c1c2a4609dd1 := strings.Split(__obf_7ac355206ce4b581, ",")
				__obf_fc3c7e67ebb963cb.
					Title = __obf_73b6c1c2a4609dd1[1]
				__obf_7ac355206ce4b581 = __obf_73b6c1c2a4609dd1[0]
			}
			__obf_1d61d66f80f81b9d, __obf_c440f080fd9d531f := strconv.ParseFloat(__obf_7ac355206ce4b581, 32)
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_fc3c7e67ebb963cb.
				Duration = float32(__obf_1d61d66f80f81b9d)
			__obf_fc3c7e67ebb963cb.
				KeyIndex = __obf_38a325c3f435b257
			__obf_20b81692820a4a6e = true
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-BYTERANGE:"):
			if __obf_a3b0d71ddb9fef57 {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_4e6c5f648da55012, __obf_db79f8eeae8fdb44+1)
			}
			if __obf_fc3c7e67ebb963cb == nil {
				__obf_fc3c7e67ebb963cb = new(Segment)
			}
			var __obf_48ea1fdcae117386 string
			if _, __obf_c440f080fd9d531f := fmt.Sscanf(__obf_4e6c5f648da55012, "#EXT-X-BYTERANGE:%s", &__obf_48ea1fdcae117386); __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			if __obf_48ea1fdcae117386 == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_db79f8eeae8fdb44+1)
			}
			if strings.Contains(__obf_48ea1fdcae117386, "@") {
				__obf_73b6c1c2a4609dd1 := strings.Split(__obf_48ea1fdcae117386, "@")
				__obf_d42889cc787805d0, __obf_c440f080fd9d531f := strconv.ParseUint(__obf_73b6c1c2a4609dd1[1], 10, 64)
				if __obf_c440f080fd9d531f != nil {
					return nil, __obf_c440f080fd9d531f
				}
				__obf_fc3c7e67ebb963cb.
					Offset = uint64(__obf_d42889cc787805d0)
				__obf_48ea1fdcae117386 = __obf_73b6c1c2a4609dd1[0]
			}
			__obf_c4a82f19bf750efc, __obf_c440f080fd9d531f := strconv.ParseUint(__obf_48ea1fdcae117386, 10, 64)
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_fc3c7e67ebb963cb.
				Length = uint64(__obf_c4a82f19bf750efc)
			__obf_a3b0d71ddb9fef57 = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_4e6c5f648da55012, "#"):
			if __obf_20b81692820a4a6e {
				if __obf_fc3c7e67ebb963cb == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_4e6c5f648da55012)
				}
				__obf_fc3c7e67ebb963cb.
					URI = __obf_4e6c5f648da55012
				__obf_a3b0d71ddb9fef57 = false
				__obf_20b81692820a4a6e = false
				__obf_2a9efb22c878897d.
					Segments = append(__obf_2a9efb22c878897d.Segments, __obf_fc3c7e67ebb963cb)
				__obf_fc3c7e67ebb963cb = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_4e6c5f648da55012, "#EXT-X-KEY"):
			__obf_f3331bf7db99fbe8 := __obf_2a3656ca6fc9e991(__obf_4e6c5f648da55012)
			if len(__obf_f3331bf7db99fbe8) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_4e6c5f648da55012, __obf_db79f8eeae8fdb44+1)
			}
			__obf_aecc285674de1d14 := CryptMethod(__obf_f3331bf7db99fbe8["METHOD"])
			if __obf_aecc285674de1d14 != "" && __obf_aecc285674de1d14 != CryptMethodAES && __obf_aecc285674de1d14 != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_aecc285674de1d14, __obf_db79f8eeae8fdb44+1)
			}
			__obf_38a325c3f435b257++
			__obf_1ffa1a359ec1627b = new(Key)
			__obf_1ffa1a359ec1627b.
				Method = __obf_aecc285674de1d14
			__obf_1ffa1a359ec1627b.
				URI = __obf_f3331bf7db99fbe8["URI"]
			__obf_1ffa1a359ec1627b.
				IV = __obf_f3331bf7db99fbe8["IV"]
			__obf_2a9efb22c878897d.
				Keys[__obf_38a325c3f435b257] = __obf_1ffa1a359ec1627b
		case __obf_4e6c5f648da55012 == "#EndList":
			__obf_2a9efb22c878897d.
				EndList = true
		default:
			continue
		}
	}

	return __obf_2a9efb22c878897d, nil
}

func __obf_2cd064455919f542(__obf_4e6c5f648da55012 string) (*MasterPlaylist, error) {
	__obf_f3331bf7db99fbe8 := __obf_2a3656ca6fc9e991(__obf_4e6c5f648da55012)
	if len(__obf_f3331bf7db99fbe8) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_e6dce7476e25508a := new(MasterPlaylist)
	for __obf_da8df2e0d1d71a60, __obf_7fe1560818350e3c := range __obf_f3331bf7db99fbe8 {
		switch {
		case __obf_da8df2e0d1d71a60 == "BANDWIDTH":
			__obf_7fe1560818350e3c, __obf_c440f080fd9d531f := strconv.ParseUint(__obf_7fe1560818350e3c, 10, 32)
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_e6dce7476e25508a.
				BandWidth = uint32(__obf_7fe1560818350e3c)
		case __obf_da8df2e0d1d71a60 == "RESOLUTION":
			__obf_e6dce7476e25508a.
				Resolution = __obf_7fe1560818350e3c
		case __obf_da8df2e0d1d71a60 == "PROGRAM-ID":
			__obf_7fe1560818350e3c, __obf_c440f080fd9d531f := strconv.ParseUint(__obf_7fe1560818350e3c, 10, 32)
			if __obf_c440f080fd9d531f != nil {
				return nil, __obf_c440f080fd9d531f
			}
			__obf_e6dce7476e25508a.
				ProgramID = uint32(__obf_7fe1560818350e3c)
		case __obf_da8df2e0d1d71a60 == "CODECS":
			__obf_e6dce7476e25508a.
				Codecs = __obf_7fe1560818350e3c
		}
	}
	return __obf_e6dce7476e25508a, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_2a3656ca6fc9e991(__obf_4e6c5f648da55012 string) map[string]string {
	__obf_d8701e668f45fb1e := __obf_40c5d667d0dd16f2.FindAllStringSubmatch(__obf_4e6c5f648da55012, -1)
	__obf_f3331bf7db99fbe8 := make(map[string]string)
	for _, __obf_a878006af810bb5f := range __obf_d8701e668f45fb1e {
		__obf_f3331bf7db99fbe8[__obf_a878006af810bb5f[1]] = strings.Trim(__obf_a878006af810bb5f[2], "\"")
	}
	return __obf_f3331bf7db99fbe8
}
