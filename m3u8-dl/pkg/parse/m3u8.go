package __obf_4f369713f5d562f0

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
var __obf_8b1431749e1fd7d6 = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

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
func __obf_4f369713f5d562f0(__obf_d313294ac1c040d2 io.Reader) (*M3u8, error) {
	__obf_65acbad49ba354d4 := bufio.NewScanner(__obf_d313294ac1c040d2)
	var __obf_e9b40af8fd88f6e1 []string
	for __obf_65acbad49ba354d4.Scan() {
		__obf_e9b40af8fd88f6e1 = append(__obf_e9b40af8fd88f6e1, __obf_65acbad49ba354d4.Text())
	}

	var (
		__obf_9d48c15da9254119 = 0
		__obf_68c10d20ad16cc2c = len(__obf_e9b40af8fd88f6e1)
		__obf_569269c4aa626235 = &M3u8{
			Keys: make(map[int]*Key),
		}
		__obf_58da825b5512615b = 0
		__obf_d31388392c335969 *Key
		__obf_7eff3017258d2428 *Segment
		__obf_1be30dfdc4997490 bool
		__obf_037fde30d512d4d0 bool
	)

	for ; __obf_9d48c15da9254119 < __obf_68c10d20ad16cc2c; __obf_9d48c15da9254119++ {
		__obf_7895f4290c271bf7 := strings.TrimSpace(__obf_e9b40af8fd88f6e1[__obf_9d48c15da9254119])
		if __obf_9d48c15da9254119 == 0 {
			if "#EXTM3U" != __obf_7895f4290c271bf7 {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case __obf_7895f4290c271bf7 == "":
			continue
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-PLAYLIST-TYPE:"):
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXT-X-PLAYLIST-TYPE:%s", &__obf_569269c4aa626235.PlaylistType); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_27c428d645d616cb := __obf_569269c4aa626235.PlaylistType == "" || __obf_569269c4aa626235.PlaylistType == PlaylistTypeVOD || __obf_569269c4aa626235.PlaylistType == PlaylistTypeEvent
			if !__obf_27c428d645d616cb {
				return nil, fmt.Errorf("invalid playlist type: %s, line: %d", __obf_569269c4aa626235.PlaylistType, __obf_9d48c15da9254119+1)
			}
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-TARGETDURATION:"):
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXT-X-TARGETDURATION:%f", &__obf_569269c4aa626235.TargetDuration); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXT-X-MEDIA-SEQUENCE:%d", &__obf_569269c4aa626235.MediaSequence); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-VERSION:"):
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXT-X-VERSION:%d", &__obf_569269c4aa626235.Version); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
		// Parse master playlist
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-STREAM-INF:"):
			__obf_a53a117382b38118, __obf_cf22f4b7fd45f7b5 := __obf_723b98381f002e07(__obf_7895f4290c271bf7)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_9d48c15da9254119++
			__obf_a53a117382b38118.
				URI = __obf_e9b40af8fd88f6e1[__obf_9d48c15da9254119]
			if __obf_a53a117382b38118.URI == "" || strings.HasPrefix(__obf_a53a117382b38118.URI, "#") {
				return nil, fmt.Errorf("invalid EXT-X-STREAM-INF URI, line: %d", __obf_9d48c15da9254119+1)
			}
			__obf_569269c4aa626235.
				MasterPlaylist = append(__obf_569269c4aa626235.MasterPlaylist, __obf_a53a117382b38118)
			continue
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXTINF:"):
			if __obf_1be30dfdc4997490 {
				return nil, fmt.Errorf("duplicate EXTINF: %s, line: %d", __obf_7895f4290c271bf7, __obf_9d48c15da9254119+1)
			}
			if __obf_7eff3017258d2428 == nil {
				__obf_7eff3017258d2428 = new(Segment)
			}
			var __obf_65acbad49ba354d4 string
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXTINF:%s", &__obf_65acbad49ba354d4); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			if strings.Contains(__obf_65acbad49ba354d4, ",") {
				__obf_d048d8493f4114a5 := strings.Split(__obf_65acbad49ba354d4, ",")
				__obf_7eff3017258d2428.
					Title = __obf_d048d8493f4114a5[1]
				__obf_65acbad49ba354d4 = __obf_d048d8493f4114a5[0]
			}
			__obf_8df6d5ba81fc4ae8, __obf_cf22f4b7fd45f7b5 := strconv.ParseFloat(__obf_65acbad49ba354d4, 32)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_7eff3017258d2428.
				Duration = float32(__obf_8df6d5ba81fc4ae8)
			__obf_7eff3017258d2428.
				KeyIndex = __obf_58da825b5512615b
			__obf_1be30dfdc4997490 = true
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-BYTERANGE:"):
			if __obf_037fde30d512d4d0 {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE: %s, line: %d", __obf_7895f4290c271bf7, __obf_9d48c15da9254119+1)
			}
			if __obf_7eff3017258d2428 == nil {
				__obf_7eff3017258d2428 = new(Segment)
			}
			var __obf_4f07068c39b562ad string
			if _, __obf_cf22f4b7fd45f7b5 := fmt.Sscanf(__obf_7895f4290c271bf7, "#EXT-X-BYTERANGE:%s", &__obf_4f07068c39b562ad); __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			if __obf_4f07068c39b562ad == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE, line: %d", __obf_9d48c15da9254119+1)
			}
			if strings.Contains(__obf_4f07068c39b562ad, "@") {
				__obf_d048d8493f4114a5 := strings.Split(__obf_4f07068c39b562ad, "@")
				__obf_a33675ba14e7538e, __obf_cf22f4b7fd45f7b5 := strconv.ParseUint(__obf_d048d8493f4114a5[1], 10, 64)
				if __obf_cf22f4b7fd45f7b5 != nil {
					return nil, __obf_cf22f4b7fd45f7b5
				}
				__obf_7eff3017258d2428.
					Offset = uint64(__obf_a33675ba14e7538e)
				__obf_4f07068c39b562ad = __obf_d048d8493f4114a5[0]
			}
			__obf_5f494361e3518da7, __obf_cf22f4b7fd45f7b5 := strconv.ParseUint(__obf_4f07068c39b562ad, 10, 64)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_7eff3017258d2428.
				Length = uint64(__obf_5f494361e3518da7)
			__obf_037fde30d512d4d0 = true
		// Parse segments URI
		case !strings.HasPrefix(__obf_7895f4290c271bf7, "#"):
			if __obf_1be30dfdc4997490 {
				if __obf_7eff3017258d2428 == nil {
					return nil, fmt.Errorf("invalid line: %s", __obf_7895f4290c271bf7)
				}
				__obf_7eff3017258d2428.
					URI = __obf_7895f4290c271bf7
				__obf_037fde30d512d4d0 = false
				__obf_1be30dfdc4997490 = false
				__obf_569269c4aa626235.
					Segments = append(__obf_569269c4aa626235.Segments, __obf_7eff3017258d2428)
				__obf_7eff3017258d2428 = nil
				continue
			}
		// Parse key
		case strings.HasPrefix(__obf_7895f4290c271bf7, "#EXT-X-KEY"):
			__obf_15396aeb10623a59 := __obf_c62a53f7ff8ac42e(__obf_7895f4290c271bf7)
			if len(__obf_15396aeb10623a59) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", __obf_7895f4290c271bf7, __obf_9d48c15da9254119+1)
			}
			__obf_9edb956f31fd8467 := CryptMethod(__obf_15396aeb10623a59["METHOD"])
			if __obf_9edb956f31fd8467 != "" && __obf_9edb956f31fd8467 != CryptMethodAES && __obf_9edb956f31fd8467 != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", __obf_9edb956f31fd8467, __obf_9d48c15da9254119+1)
			}
			__obf_58da825b5512615b++
			__obf_d31388392c335969 = new(Key)
			__obf_d31388392c335969.
				Method = __obf_9edb956f31fd8467
			__obf_d31388392c335969.
				URI = __obf_15396aeb10623a59["URI"]
			__obf_d31388392c335969.
				IV = __obf_15396aeb10623a59["IV"]
			__obf_569269c4aa626235.
				Keys[__obf_58da825b5512615b] = __obf_d31388392c335969
		case __obf_7895f4290c271bf7 == "#EndList":
			__obf_569269c4aa626235.
				EndList = true
		default:
			continue
		}
	}

	return __obf_569269c4aa626235, nil
}

func __obf_723b98381f002e07(__obf_7895f4290c271bf7 string) (*MasterPlaylist, error) {
	__obf_15396aeb10623a59 := __obf_c62a53f7ff8ac42e(__obf_7895f4290c271bf7)
	if len(__obf_15396aeb10623a59) == 0 {
		return nil, errors.New("empty parameter")
	}
	__obf_a53a117382b38118 := new(MasterPlaylist)
	for __obf_7bd88fc4fd0d12b9, __obf_a7c314bef90e932a := range __obf_15396aeb10623a59 {
		switch {
		case __obf_7bd88fc4fd0d12b9 == "BANDWIDTH":
			__obf_a7c314bef90e932a, __obf_cf22f4b7fd45f7b5 := strconv.ParseUint(__obf_a7c314bef90e932a, 10, 32)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_a53a117382b38118.
				BandWidth = uint32(__obf_a7c314bef90e932a)
		case __obf_7bd88fc4fd0d12b9 == "RESOLUTION":
			__obf_a53a117382b38118.
				Resolution = __obf_a7c314bef90e932a
		case __obf_7bd88fc4fd0d12b9 == "PROGRAM-ID":
			__obf_a7c314bef90e932a, __obf_cf22f4b7fd45f7b5 := strconv.ParseUint(__obf_a7c314bef90e932a, 10, 32)
			if __obf_cf22f4b7fd45f7b5 != nil {
				return nil, __obf_cf22f4b7fd45f7b5
			}
			__obf_a53a117382b38118.
				ProgramID = uint32(__obf_a7c314bef90e932a)
		case __obf_7bd88fc4fd0d12b9 == "CODECS":
			__obf_a53a117382b38118.
				Codecs = __obf_a7c314bef90e932a
		}
	}
	return __obf_a53a117382b38118, nil
}

// parseLineParameters extra parameters in string `line`
func __obf_c62a53f7ff8ac42e(__obf_7895f4290c271bf7 string) map[string]string {
	__obf_94f53bdeb40d54c3 := __obf_8b1431749e1fd7d6.FindAllStringSubmatch(__obf_7895f4290c271bf7, -1)
	__obf_15396aeb10623a59 := make(map[string]string)
	for _, __obf_91ab9745ee46e8df := range __obf_94f53bdeb40d54c3 {
		__obf_15396aeb10623a59[__obf_91ab9745ee46e8df[1]] = strings.Trim(__obf_91ab9745ee46e8df[2], "\"")
	}
	return __obf_15396aeb10623a59
}
