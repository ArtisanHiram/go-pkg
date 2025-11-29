package __obf_ce58dca6c0b3a695

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
var __obf_44a353e46bbd60b1 = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

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
func __obf_ce58dca6c0b3a695(__obf_b48006c49b1ccc26 io.Reader) (*M3u8, error) {
	__obf_9798ea289a86c4a9 := bufio.NewScanner(__obf_b48006c49b1ccc26)
	var __obf_c3cc67f6a5a7190a []string
	for __obf_9798ea289a86c4a9.Scan() {
		__obf_c3cc67f6a5a7190a = append(__obf_c3cc67f6a5a7190a, __obf_9798ea289a86c4a9.Text())
	}

	var (
		__obf_b104aac23fcea60d = 0
		__obf_391160686cdcb6e4 = len(__obf_c3cc67f6a5a7190a)
		__obf_4d6af585f15ca9ea = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_a890fff637ffe3e7 = 0
		__obf_8a7802f8f4246380 *Key
		__obf_2eed89abaa66cb55 *Segment
		__obf_85039c5ddfa11c46 bool
		__obf_13af3916a1a534ab bool
	)

	for ; __obf_b104aac23fcea60d < __obf_391160686cdcb6e4; __obf_b104aac23fcea60d++ {
		__obf_26d5bbae8304c589 := strings.TrimSpace(__obf_c3cc67f6a5a7190a[__obf_b104aac23fcea60d])
		if __obf_b104aac23fcea60d == 0 {
			if "#EXTM3U" != __obf_26d5bbae8304c589 {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_26d5bbae8304c589 == "":
			continue
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_4d6af585f15ca9ea.PlaylistType); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_812d973ec4dde6a9 := __obf_4d6af585f15ca9ea.PlaylistType == "" || __obf_4d6af585f15ca9ea.PlaylistType == PlaylistTypeVOD || __obf_4d6af585f15ca9ea.PlaylistType == PlaylistTypeEvent
			if !__obf_812d973ec4dde6a9 {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_4d6af585f15ca9ea.PlaylistType, __obf_b104aac23fcea60d+1)
			}
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-TARGETDURATION:"):
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXT-X-TARGETDURATION:%f", &__obf_4d6af585f15ca9ea.TargetDuration); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_4d6af585f15ca9ea.MediaSequence); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-VERSION:"):
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXT-X-VERSION:%d", &__obf_4d6af585f15ca9ea.Version); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-STREAM-INF:"):
			__obf_ce86970444015a2c, __obf_2f034306c99cf38e := __obf_8cf68c8e7e256f27(__obf_26d5bbae8304c589)
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_b104aac23fcea60d++
			__obf_ce86970444015a2c.
				URI = __obf_c3cc67f6a5a7190a[__obf_b104aac23fcea60d]
			if __obf_ce86970444015a2c.URI == "" || strings.HasPrefix(__obf_ce86970444015a2c.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_b104aac23fcea60d+1)
			}
			__obf_4d6af585f15ca9ea.
				MasterPlaylist = append(__obf_4d6af585f15ca9ea.MasterPlaylist, __obf_ce86970444015a2c)
			continue
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXTINF:"):
			if __obf_85039c5ddfa11c46 {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_26d5bbae8304c589, __obf_b104aac23fcea60d+1)
			}
			if __obf_2eed89abaa66cb55 == nil {
				__obf_2eed89abaa66cb55 = new(Segment)
			}
			var __obf_9798ea289a86c4a9 string
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXTINF:%s", &__obf_9798ea289a86c4a9); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			if strings.Contains(__obf_9798ea289a86c4a9, ",") {
				__obf_d0b2c1313dcfd750 := strings.Split(__obf_9798ea289a86c4a9, ",")
				__obf_2eed89abaa66cb55.
					Title = __obf_d0b2c1313dcfd750[1]
				__obf_9798ea289a86c4a9 = __obf_d0b2c1313dcfd750[0]
			}
			__obf_f51cf8064929aa8e, __obf_2f034306c99cf38e := strconv.ParseFloat(__obf_9798ea289a86c4a9, 32)
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_2eed89abaa66cb55.
				Duration = float32(__obf_f51cf8064929aa8e)
			__obf_2eed89abaa66cb55.
				KeyIndex = __obf_a890fff637ffe3e7
			__obf_85039c5ddfa11c46 = true
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-BYTERANGE:"):
			if __obf_13af3916a1a534ab {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_26d5bbae8304c589, __obf_b104aac23fcea60d+1)
			}
			if __obf_2eed89abaa66cb55 == nil {
				__obf_2eed89abaa66cb55 = new(Segment)
			}
			var __obf_8918a1308140e980 string
			if _, __obf_2f034306c99cf38e := fmt.Sscanf(__obf_26d5bbae8304c589, "#EXT-X-BYTERANGE:%s", &__obf_8918a1308140e980); __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			if __obf_8918a1308140e980 == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_b104aac23fcea60d+1)
			}
			if strings.Contains(__obf_8918a1308140e980, "@") {
				__obf_d0b2c1313dcfd750 := strings.Split(__obf_8918a1308140e980, "@")
				__obf_2994eb74c07e0985, __obf_2f034306c99cf38e := strconv.ParseUint(__obf_d0b2c1313dcfd750[1], 10, 64)
				if __obf_2f034306c99cf38e != nil {
					return nil, __obf_2f034306c99cf38e
				}
				__obf_2eed89abaa66cb55.
					Offset = uint64(__obf_2994eb74c07e0985)
				__obf_8918a1308140e980 = __obf_d0b2c1313dcfd750[0]
			}
			__obf_ed408eab5f803d5d, __obf_2f034306c99cf38e := strconv.ParseUint(__obf_8918a1308140e980, 10, 64)
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_2eed89abaa66cb55.
				Length = uint64(__obf_ed408eab5f803d5d)
			__obf_13af3916a1a534ab = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_26d5bbae8304c589, "#"):
			if __obf_85039c5ddfa11c46 {
				if __obf_2eed89abaa66cb55 == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_26d5bbae8304c589)
				}
				__obf_2eed89abaa66cb55.
					URI = __obf_26d5bbae8304c589
				__obf_13af3916a1a534ab = false
				__obf_85039c5ddfa11c46 = false
				__obf_4d6af585f15ca9ea.
					Segments = append(__obf_4d6af585f15ca9ea.Segments, __obf_2eed89abaa66cb55)
				__obf_2eed89abaa66cb55 = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_26d5bbae8304c589, "#EXT-X-KEY"):
			__obf_09c5f4b2fbdf9621 := __obf_fd258b3c819f3ebd(__obf_26d5bbae8304c589)
			if len(__obf_09c5f4b2fbdf9621) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_26d5bbae8304c589, __obf_b104aac23fcea60d+1)
			}
			__obf_a84fd95f4482180a := CryptMethod(__obf_09c5f4b2fbdf9621["METHOD"])
			if __obf_a84fd95f4482180a != "" && __obf_a84fd95f4482180a != CryptMethodAES && __obf_a84fd95f4482180a != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_a84fd95f4482180a, __obf_b104aac23fcea60d+1)
			}
			__obf_a890fff637ffe3e7++
			__obf_8a7802f8f4246380 = new(Key)
			__obf_8a7802f8f4246380.
				Method = __obf_a84fd95f4482180a
			__obf_8a7802f8f4246380.
				URI = __obf_09c5f4b2fbdf9621["URI"]
			__obf_8a7802f8f4246380.
				IV = __obf_09c5f4b2fbdf9621["IV"]
			__obf_4d6af585f15ca9ea.
				Keys[__obf_a890fff637ffe3e7] = __obf_8a7802f8f4246380
		case __obf_26d5bbae8304c589 == "#EndList":
			__obf_4d6af585f15ca9ea.
				EndList = true
		default:
			continue
		}
	}

	return __obf_4d6af585f15ca9ea, nil
}

func __obf_8cf68c8e7e256f27(__obf_26d5bbae8304c589 string) (*MasterPlaylist, error) {
	__obf_09c5f4b2fbdf9621 := __obf_fd258b3c819f3ebd(__obf_26d5bbae8304c589)
	if len(__obf_09c5f4b2fbdf9621) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_ce86970444015a2c := new(MasterPlaylist)
	for __obf_be52339d47e2f7a4, __obf_8e5c8bd1a7f37e72 := range __obf_09c5f4b2fbdf9621 {
		switch {
		case __obf_be52339d47e2f7a4 == "BANDWIDTH":
			__obf_8e5c8bd1a7f37e72, __obf_2f034306c99cf38e := strconv.ParseUint(__obf_8e5c8bd1a7f37e72, 10, 32)
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_ce86970444015a2c.
				BandWidth = uint32(__obf_8e5c8bd1a7f37e72)
		case __obf_be52339d47e2f7a4 == "RESOLUTION":
			__obf_ce86970444015a2c.
				Resolution = __obf_8e5c8bd1a7f37e72
		case __obf_be52339d47e2f7a4 == "PROGRAM-ID":
			__obf_8e5c8bd1a7f37e72, __obf_2f034306c99cf38e := strconv.ParseUint(__obf_8e5c8bd1a7f37e72, 10, 32)
			if __obf_2f034306c99cf38e != nil {
				return nil, __obf_2f034306c99cf38e
			}
			__obf_ce86970444015a2c.
				ProgramID = uint32(__obf_8e5c8bd1a7f37e72)
		case __obf_be52339d47e2f7a4 == "CODECS":
			__obf_ce86970444015a2c.
				Codecs = __obf_8e5c8bd1a7f37e72
		}
	}
	return __obf_ce86970444015a2c, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_fd258b3c819f3ebd(__obf_26d5bbae8304c589 string) map[string]string {
	__obf_95bfe72760a6bd42 := __obf_44a353e46bbd60b1.FindAllStringSubmatch(__obf_26d5bbae8304c589, -1)
	__obf_09c5f4b2fbdf9621 := make(map[string]string)
	for _, __obf_ad74011509a956b7 := range __obf_95bfe72760a6bd42 {
		__obf_09c5f4b2fbdf9621[__obf_ad74011509a956b7[1]] = strings.Trim(__obf_ad74011509a956b7[2], "\"")
	}
	return __obf_09c5f4b2fbdf9621
}
