package __obf_73049aff9f20feb6

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
var __obf_7e2dc12f0fb1d929 = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

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
func __obf_73049aff9f20feb6(__obf_6731eec6097ad6a3 io.Reader) (*M3u8, error) {
	__obf_ea385fc14e5ef8d0 := bufio.NewScanner(__obf_6731eec6097ad6a3)
	var __obf_f24fe726d92f0e38 []string
	for __obf_ea385fc14e5ef8d0.Scan() {
		__obf_f24fe726d92f0e38 = append(__obf_f24fe726d92f0e38, __obf_ea385fc14e5ef8d0.Text())
	}

	var (
		__obf_1de1feee1e49aa61 = 0
		__obf_409676e28dd0bd0b = len(__obf_f24fe726d92f0e38)
		__obf_bae86baf40186d8d = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_577459675089d47b = 0
		__obf_e6b5c32ce5643ea2 *Key
		__obf_bc50a2ddebe89d25 *Segment
		__obf_6f7e37d652728bca bool
		__obf_0a53e61d7a300cf2 bool
	)

	for ; __obf_1de1feee1e49aa61 < __obf_409676e28dd0bd0b; __obf_1de1feee1e49aa61++ {
		__obf_758d6b93b26fc650 := strings.TrimSpace(__obf_f24fe726d92f0e38[__obf_1de1feee1e49aa61])
		if __obf_1de1feee1e49aa61 == 0 {
			if "#EXTM3U" != __obf_758d6b93b26fc650 {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_758d6b93b26fc650 == "":
			continue
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_bae86baf40186d8d.PlaylistType); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_6f8f48f27c0a4367 := __obf_bae86baf40186d8d.PlaylistType == "" || __obf_bae86baf40186d8d.PlaylistType == PlaylistTypeVOD || __obf_bae86baf40186d8d.PlaylistType == PlaylistTypeEvent
			if !__obf_6f8f48f27c0a4367 {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_bae86baf40186d8d.PlaylistType, __obf_1de1feee1e49aa61+1)
			}
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-TARGETDURATION:"):
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXT-X-TARGETDURATION:%f", &__obf_bae86baf40186d8d.TargetDuration); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_bae86baf40186d8d.MediaSequence); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-VERSION:"):
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXT-X-VERSION:%d", &__obf_bae86baf40186d8d.Version); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-STREAM-INF:"):
			__obf_93895255bb6edd3b, __obf_5159af91c7c96fc8 := __obf_c90ec429c2fc3f58(__obf_758d6b93b26fc650)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_1de1feee1e49aa61++
			__obf_93895255bb6edd3b.
				URI = __obf_f24fe726d92f0e38[__obf_1de1feee1e49aa61]
			if __obf_93895255bb6edd3b.URI == "" || strings.HasPrefix(__obf_93895255bb6edd3b.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_1de1feee1e49aa61+1)
			}
			__obf_bae86baf40186d8d.
				MasterPlaylist = append(__obf_bae86baf40186d8d.MasterPlaylist, __obf_93895255bb6edd3b)
			continue
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXTINF:"):
			if __obf_6f7e37d652728bca {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_758d6b93b26fc650, __obf_1de1feee1e49aa61+1)
			}
			if __obf_bc50a2ddebe89d25 == nil {
				__obf_bc50a2ddebe89d25 = new(Segment)
			}
			var __obf_ea385fc14e5ef8d0 string
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXTINF:%s", &__obf_ea385fc14e5ef8d0); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			if strings.Contains(__obf_ea385fc14e5ef8d0, ",") {
				__obf_9959b249a3c51032 := strings.Split(__obf_ea385fc14e5ef8d0, ",")
				__obf_bc50a2ddebe89d25.
					Title = __obf_9959b249a3c51032[1]
				__obf_ea385fc14e5ef8d0 = __obf_9959b249a3c51032[0]
			}
			__obf_dcd0a29a5368c9d6, __obf_5159af91c7c96fc8 := strconv.ParseFloat(__obf_ea385fc14e5ef8d0, 32)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_bc50a2ddebe89d25.
				Duration = float32(__obf_dcd0a29a5368c9d6)
			__obf_bc50a2ddebe89d25.
				KeyIndex = __obf_577459675089d47b
			__obf_6f7e37d652728bca = true
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-BYTERANGE:"):
			if __obf_0a53e61d7a300cf2 {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_758d6b93b26fc650, __obf_1de1feee1e49aa61+1)
			}
			if __obf_bc50a2ddebe89d25 == nil {
				__obf_bc50a2ddebe89d25 = new(Segment)
			}
			var __obf_5ac92ddd0bc77437 string
			if _, __obf_5159af91c7c96fc8 := fmt.Sscanf(__obf_758d6b93b26fc650, "#EXT-X-BYTERANGE:%s", &__obf_5ac92ddd0bc77437); __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			if __obf_5ac92ddd0bc77437 == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_1de1feee1e49aa61+1)
			}
			if strings.Contains(__obf_5ac92ddd0bc77437, "@") {
				__obf_9959b249a3c51032 := strings.Split(__obf_5ac92ddd0bc77437, "@")
				__obf_7d49d4e739670bc7, __obf_5159af91c7c96fc8 := strconv.ParseUint(__obf_9959b249a3c51032[1], 10, 64)
				if __obf_5159af91c7c96fc8 != nil {
					return nil, __obf_5159af91c7c96fc8
				}
				__obf_bc50a2ddebe89d25.
					Offset = uint64(__obf_7d49d4e739670bc7)
				__obf_5ac92ddd0bc77437 = __obf_9959b249a3c51032[0]
			}
			__obf_1679a5bbeae80917, __obf_5159af91c7c96fc8 := strconv.ParseUint(__obf_5ac92ddd0bc77437, 10, 64)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_bc50a2ddebe89d25.
				Length = uint64(__obf_1679a5bbeae80917)
			__obf_0a53e61d7a300cf2 = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_758d6b93b26fc650, "#"):
			if __obf_6f7e37d652728bca {
				if __obf_bc50a2ddebe89d25 == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_758d6b93b26fc650)
				}
				__obf_bc50a2ddebe89d25.
					URI = __obf_758d6b93b26fc650
				__obf_0a53e61d7a300cf2 = false
				__obf_6f7e37d652728bca = false
				__obf_bae86baf40186d8d.
					Segments = append(__obf_bae86baf40186d8d.Segments, __obf_bc50a2ddebe89d25)
				__obf_bc50a2ddebe89d25 = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_758d6b93b26fc650, "#EXT-X-KEY"):
			__obf_a30a1ca4e0239f95 := __obf_303ed0550e240542(__obf_758d6b93b26fc650)
			if len(__obf_a30a1ca4e0239f95) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_758d6b93b26fc650, __obf_1de1feee1e49aa61+1)
			}
			__obf_73420cef7f380c22 := CryptMethod(__obf_a30a1ca4e0239f95["METHOD"])
			if __obf_73420cef7f380c22 != "" && __obf_73420cef7f380c22 != CryptMethodAES && __obf_73420cef7f380c22 != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_73420cef7f380c22, __obf_1de1feee1e49aa61+1)
			}
			__obf_577459675089d47b++
			__obf_e6b5c32ce5643ea2 = new(Key)
			__obf_e6b5c32ce5643ea2.
				Method = __obf_73420cef7f380c22
			__obf_e6b5c32ce5643ea2.
				URI = __obf_a30a1ca4e0239f95["URI"]
			__obf_e6b5c32ce5643ea2.
				IV = __obf_a30a1ca4e0239f95["IV"]
			__obf_bae86baf40186d8d.
				Keys[__obf_577459675089d47b] = __obf_e6b5c32ce5643ea2
		case __obf_758d6b93b26fc650 == "#EndList":
			__obf_bae86baf40186d8d.
				EndList = true
		default:
			continue
		}
	}

	return __obf_bae86baf40186d8d, nil
}

func __obf_c90ec429c2fc3f58(__obf_758d6b93b26fc650 string) (*MasterPlaylist, error) {
	__obf_a30a1ca4e0239f95 := __obf_303ed0550e240542(__obf_758d6b93b26fc650)
	if len(__obf_a30a1ca4e0239f95) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_93895255bb6edd3b := new(MasterPlaylist)
	for __obf_095464ee1993f2c9, __obf_201a5ce6f5a18794 := range __obf_a30a1ca4e0239f95 {
		switch {
		case __obf_095464ee1993f2c9 == "BANDWIDTH":
			__obf_201a5ce6f5a18794, __obf_5159af91c7c96fc8 := strconv.ParseUint(__obf_201a5ce6f5a18794, 10, 32)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_93895255bb6edd3b.
				BandWidth = uint32(__obf_201a5ce6f5a18794)
		case __obf_095464ee1993f2c9 == "RESOLUTION":
			__obf_93895255bb6edd3b.
				Resolution = __obf_201a5ce6f5a18794
		case __obf_095464ee1993f2c9 == "PROGRAM-ID":
			__obf_201a5ce6f5a18794, __obf_5159af91c7c96fc8 := strconv.ParseUint(__obf_201a5ce6f5a18794, 10, 32)
			if __obf_5159af91c7c96fc8 != nil {
				return nil, __obf_5159af91c7c96fc8
			}
			__obf_93895255bb6edd3b.
				ProgramID = uint32(__obf_201a5ce6f5a18794)
		case __obf_095464ee1993f2c9 == "CODECS":
			__obf_93895255bb6edd3b.
				Codecs = __obf_201a5ce6f5a18794
		}
	}
	return __obf_93895255bb6edd3b, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_303ed0550e240542(__obf_758d6b93b26fc650 string) map[string]string {
	__obf_acec7657b362e547 := __obf_7e2dc12f0fb1d929.FindAllStringSubmatch(__obf_758d6b93b26fc650, -1)
	__obf_a30a1ca4e0239f95 := make(map[string]string)
	for _, __obf_0f99f30fa91aca58 := range __obf_acec7657b362e547 {
		__obf_a30a1ca4e0239f95[__obf_0f99f30fa91aca58[1]] = strings.Trim(__obf_0f99f30fa91aca58[2], "\"")
	}
	return __obf_a30a1ca4e0239f95
}
