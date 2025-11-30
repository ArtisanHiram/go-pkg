package main

import (
	"bufio"
	"fmt"
	parse "github.com/ArtisanHiram/go-pkg/m3u8-dl/pkg/parse"
	tool "github.com/ArtisanHiram/go-pkg/m3u8-dl/pkg/tool"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	__obf_18b830a45925679b = ".ts"
	__obf_83fb9ced680e4257 = "ts"
	__obf_7533cb0b3f2c2fb3 = "main.ts"
	__obf_390bb9d573b17c87 = "_tmp"
	__obf_c3590534151a078c = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_8d00c8f1e3ce1eab sync.Mutex
	__obf_88997d3403f49d7e []int
	__obf_c573b889987f5cc1 string
	__obf_fc490b994714b516 string
	__obf_86cee42cbf250e94 int32
	__obf_340d7c8ea793e23a int
	__obf_1be2d16393c02b78 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_f153ab1806df0bbb string, __obf_1d184b1c1810901d string) (*Downloader, error) {
	__obf_1be2d16393c02b78, __obf_dc3122c5b8033a43 := parse.ParseFromNet(__obf_1d184b1c1810901d)
	if __obf_dc3122c5b8033a43 != nil {
		return nil, __obf_dc3122c5b8033a43
	}
	var __obf_c573b889987f5cc1 string
	// If no output folder specified, use current directory
	if __obf_f153ab1806df0bbb == "" {
		__obf_0f20e577907d7cfd, __obf_dc3122c5b8033a43 := tool.CurrentDir()
		if __obf_dc3122c5b8033a43 != nil {
			return nil, __obf_dc3122c5b8033a43
		}
		__obf_c573b889987f5cc1 = filepath.Join(__obf_0f20e577907d7cfd, __obf_f153ab1806df0bbb)
	} else {
		__obf_c573b889987f5cc1 = __obf_f153ab1806df0bbb
	}
	if __obf_dc3122c5b8033a43 := os.MkdirAll(__obf_c573b889987f5cc1, os.ModePerm); __obf_dc3122c5b8033a43 != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_dc3122c5b8033a43.Error())
	}
	__obf_fc490b994714b516 := filepath.Join(__obf_c573b889987f5cc1, __obf_83fb9ced680e4257)
	if __obf_dc3122c5b8033a43 := os.MkdirAll(__obf_fc490b994714b516, os.ModePerm); __obf_dc3122c5b8033a43 != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_fc490b994714b516, __obf_dc3122c5b8033a43.Error())
	}
	__obf_c12d68fa4f84f22a := &Downloader{__obf_c573b889987f5cc1: __obf_c573b889987f5cc1, __obf_fc490b994714b516: __obf_fc490b994714b516, __obf_1be2d16393c02b78: __obf_1be2d16393c02b78}
	__obf_c12d68fa4f84f22a.
		// segment length
		__obf_340d7c8ea793e23a = len(__obf_1be2d16393c02b78.M3u8.Segments)
	__obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e = __obf_ba97e4154965777f(__obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a)
	return __obf_c12d68fa4f84f22a, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_c12d68fa4f84f22a *Downloader) Start(__obf_40ea3edba43f3604 int) error {
	var __obf_4e7b7aa22ca4ec35 sync.WaitGroup
	for {
		__obf_1f9b15f2489fb658, __obf_d091c7cb40aa0fde, __obf_dc3122c5b8033a43 := __obf_c12d68fa4f84f22a.__obf_6bca1c3b374bd79a()
		if __obf_dc3122c5b8033a43 != nil {
			if __obf_d091c7cb40aa0fde {
				break
			}
			continue
		}
		__obf_4e7b7aa22ca4ec35.
			Add(1)
		go func(__obf_114e493d9ae4faa9 int) {
			defer __obf_4e7b7aa22ca4ec35.Done()
			if __obf_dc3122c5b8033a43 := __obf_c12d68fa4f84f22a.__obf_09819b41bf4da3bf(__obf_114e493d9ae4faa9); __obf_dc3122c5b8033a43 != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_dc3122c5b8033a43.Error())
				if __obf_dc3122c5b8033a43 := __obf_c12d68fa4f84f22a.__obf_157e1e6096fa8e37(__obf_114e493d9ae4faa9); __obf_dc3122c5b8033a43 != nil {
					fmt.Println(__obf_dc3122c5b8033a43.Error())
				}
			}
		}(__obf_1f9b15f2489fb658)
	}
	__obf_4e7b7aa22ca4ec35.
		Wait()
	if __obf_dc3122c5b8033a43 := __obf_c12d68fa4f84f22a.__obf_7cfceb6d1af62c56(); __obf_dc3122c5b8033a43 != nil {
		return __obf_dc3122c5b8033a43
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_c12d68fa4f84f22a *Downloader) __obf_09819b41bf4da3bf(__obf_c84e4f5ebf37ede9 int) error {
	__obf_d4be53f61ba5b763 := __obf_d4be53f61ba5b763(__obf_c84e4f5ebf37ede9)
	__obf_fecb1c377268ee27 := __obf_c12d68fa4f84f22a.__obf_82bbeea61c14ea37(__obf_c84e4f5ebf37ede9)
	__obf_16f02dbc898d4525, __obf_fb87e5baf0b3996c := tool.Get(__obf_fecb1c377268ee27)
	if __obf_fb87e5baf0b3996c != nil {
		return fmt.Errorf("request %s, %s", __obf_fecb1c377268ee27, __obf_fb87e5baf0b3996c.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_16f02dbc898d4525.Close()
	__obf_dbd3504758a00f00 := filepath.Join(__obf_c12d68fa4f84f22a.__obf_fc490b994714b516, __obf_d4be53f61ba5b763)
	__obf_824ce67eea461c85 := __obf_dbd3504758a00f00 + __obf_390bb9d573b17c87
	__obf_d616409e59de26e4, __obf_dc3122c5b8033a43 := os.Create(__obf_824ce67eea461c85)
	if __obf_dc3122c5b8033a43 != nil {
		return fmt.Errorf("create file: %s, %s", __obf_d4be53f61ba5b763, __obf_dc3122c5b8033a43.Error())
	}
	__obf_429e9503bb286115, __obf_dc3122c5b8033a43 := io.ReadAll(__obf_16f02dbc898d4525)
	if __obf_dc3122c5b8033a43 != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_fecb1c377268ee27, __obf_dc3122c5b8033a43.Error())
	}
	__obf_2e2fef70d510d2cc := __obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.M3u8.Segments[__obf_c84e4f5ebf37ede9]
	if __obf_2e2fef70d510d2cc == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_c84e4f5ebf37ede9)
	}
	__obf_66389a2da74c3988, __obf_a34014197a1e9d7d := __obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.Keys[__obf_2e2fef70d510d2cc.KeyIndex]
	if __obf_a34014197a1e9d7d && __obf_66389a2da74c3988 != "" {
		__obf_429e9503bb286115, __obf_dc3122c5b8033a43 = tool.AES128Decrypt(__obf_429e9503bb286115, []byte(__obf_66389a2da74c3988),
			[]byte(__obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.M3u8.Keys[__obf_2e2fef70d510d2cc.KeyIndex].IV))
		if __obf_dc3122c5b8033a43 != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_fecb1c377268ee27, __obf_dc3122c5b8033a43.Error())
		}
	}
	__obf_8b6eedb725760e55 := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_c213a750deef74e2 := //0x47
		len(__obf_429e9503bb286115)
	for __obf_ffbb1430cc4cf8e7 := 0; __obf_ffbb1430cc4cf8e7 < __obf_c213a750deef74e2; __obf_ffbb1430cc4cf8e7++ {
		if __obf_429e9503bb286115[__obf_ffbb1430cc4cf8e7] == __obf_8b6eedb725760e55 {
			__obf_429e9503bb286115 = __obf_429e9503bb286115[__obf_ffbb1430cc4cf8e7:]
			break
		}
	}
	__obf_490c58dd129697d8 := bufio.NewWriter(__obf_d616409e59de26e4)
	if _, __obf_dc3122c5b8033a43 := __obf_490c58dd129697d8.Write(__obf_429e9503bb286115); __obf_dc3122c5b8033a43 != nil {
		return fmt.Errorf("write to %s: %s", __obf_824ce67eea461c85, __obf_dc3122c5b8033a43.Error())
	}
	// Release file resource to rename file
	_ = __obf_d616409e59de26e4.Close()
	if __obf_dc3122c5b8033a43 = os.Rename(__obf_824ce67eea461c85, __obf_dbd3504758a00f00); __obf_dc3122c5b8033a43 != nil {
		return __obf_dc3122c5b8033a43
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_c12d68fa4f84f22a.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_86cee42cbf250e94, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_c12d68fa4f84f22a.__obf_86cee42cbf250e94)/float32(__obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a)*100, __obf_fecb1c377268ee27)
	return nil
}

// next Obtains a next segment index
func (__obf_c12d68fa4f84f22a *Downloader) __obf_6bca1c3b374bd79a() (__obf_c84e4f5ebf37ede9 int, __obf_d091c7cb40aa0fde bool, __obf_dc3122c5b8033a43 error) {
	__obf_c12d68fa4f84f22a.__obf_8d00c8f1e3ce1eab.
		Lock()
	defer __obf_c12d68fa4f84f22a.__obf_8d00c8f1e3ce1eab.Unlock()
	if len(__obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e) == 0 {
		__obf_dc3122c5b8033a43 = fmt.Errorf("queue empty")
		if __obf_c12d68fa4f84f22a.__obf_86cee42cbf250e94 == int32(__obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a) {
			__obf_d091c7cb40aa0fde = true
			return
		}
		__obf_d091c7cb40aa0fde = // Some segment indexes are still running.
			false
		return
	}
	__obf_c84e4f5ebf37ede9 = __obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e[0]
	__obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e = __obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e[1:]
	return
}

// back For retry while download failed
func (__obf_c12d68fa4f84f22a *Downloader) __obf_157e1e6096fa8e37(__obf_c84e4f5ebf37ede9 int) error {
	__obf_c12d68fa4f84f22a.__obf_8d00c8f1e3ce1eab.
		Lock()
	defer __obf_c12d68fa4f84f22a.__obf_8d00c8f1e3ce1eab.Unlock()
	if __obf_2e2fef70d510d2cc := __obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.M3u8.Segments[__obf_c84e4f5ebf37ede9]; __obf_2e2fef70d510d2cc == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_c84e4f5ebf37ede9)
	}
	__obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e = append(__obf_c12d68fa4f84f22a.__obf_88997d3403f49d7e,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_c84e4f5ebf37ede9)
	return nil
}

func (__obf_c12d68fa4f84f22a *Downloader) __obf_7cfceb6d1af62c56() error {
	__obf_d370445e18b76654 := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_114e493d9ae4faa9 := 0; __obf_114e493d9ae4faa9 < __obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a; __obf_114e493d9ae4faa9++ {
		__obf_d4be53f61ba5b763 := __obf_d4be53f61ba5b763(__obf_114e493d9ae4faa9)
		__obf_d616409e59de26e4 := filepath.Join(__obf_c12d68fa4f84f22a.__obf_fc490b994714b516, __obf_d4be53f61ba5b763)
		if _, __obf_dc3122c5b8033a43 := os.Stat(__obf_d616409e59de26e4); __obf_dc3122c5b8033a43 != nil {
			__obf_d370445e18b76654++
		}
	}
	if __obf_d370445e18b76654 > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_d370445e18b76654)
	}
	__obf_78f38697a4dfb24e := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_c12d68fa4f84f22a.__obf_c573b889987f5cc1, __obf_7533cb0b3f2c2fb3)
	__obf_96b3a6386f35091c, __obf_dc3122c5b8033a43 := os.Create(__obf_78f38697a4dfb24e)
	if __obf_dc3122c5b8033a43 != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_dc3122c5b8033a43.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_96b3a6386f35091c.Close()
	__obf_894a45cafb773718 := bufio.NewWriter(__obf_96b3a6386f35091c)
	__obf_76e7862352c3518c := 0
	for __obf_c84e4f5ebf37ede9 := 0; __obf_c84e4f5ebf37ede9 < __obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a; __obf_c84e4f5ebf37ede9++ {
		__obf_d4be53f61ba5b763 := __obf_d4be53f61ba5b763(__obf_c84e4f5ebf37ede9)
		__obf_429e9503bb286115, __obf_dc3122c5b8033a43 := os.ReadFile(filepath.Join(__obf_c12d68fa4f84f22a.__obf_fc490b994714b516, __obf_d4be53f61ba5b763))
		_, __obf_dc3122c5b8033a43 = __obf_894a45cafb773718.Write(__obf_429e9503bb286115)
		if __obf_dc3122c5b8033a43 != nil {
			continue
		}
		__obf_76e7862352c3518c++
		tool.DrawProgressBar("merge",
			float32(__obf_76e7862352c3518c)/float32(__obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a), __obf_c3590534151a078c)
	}
	_ = __obf_894a45cafb773718.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_c12d68fa4f84f22a.__obf_fc490b994714b516)

	if __obf_76e7862352c3518c != __obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a {
		fmt.Printf("[warning] \n%d files merge failed", __obf_c12d68fa4f84f22a.__obf_340d7c8ea793e23a-__obf_76e7862352c3518c)
	}

	fmt.Printf("\n[output] %s\n", __obf_78f38697a4dfb24e)

	return nil
}

func (__obf_c12d68fa4f84f22a *Downloader) __obf_82bbeea61c14ea37(__obf_c84e4f5ebf37ede9 int) string {
	__obf_7837b7611ffda168 := __obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.M3u8.Segments[__obf_c84e4f5ebf37ede9]
	return tool.ResolveURL(__obf_c12d68fa4f84f22a.__obf_1be2d16393c02b78.URL, __obf_7837b7611ffda168.URI)
}

func __obf_d4be53f61ba5b763(__obf_c19c9ddfbd67965a int) string {
	return strconv.Itoa(__obf_c19c9ddfbd67965a) + __obf_18b830a45925679b
}

func __obf_ba97e4154965777f(len int) []int {
	__obf_09bbbf524afff3fa := make([]int, 0)
	for __obf_f33cf166447779bc := 0; __obf_f33cf166447779bc < len; __obf_f33cf166447779bc++ {
		__obf_09bbbf524afff3fa = append(__obf_09bbbf524afff3fa, __obf_f33cf166447779bc)
	}
	return __obf_09bbbf524afff3fa
}
