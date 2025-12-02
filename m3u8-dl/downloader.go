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
	__obf_275560fef6ecadbd = ".ts"
	__obf_dcf01f9c274980af = "ts"
	__obf_4eebfa8f7b7e45e5 = "main.ts"
	__obf_ac489e873476f8ee = "_tmp"
	__obf_aa9349b594dc6e4f = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_4713ca93e91cd1b0 sync.Mutex
	__obf_9f8d055130477bb2 []int
	__obf_7c8db613d3d1c7a2 string
	__obf_0427b13a894d972f string
	__obf_726c89b5f4d9bd88 int32
	__obf_6fa4a1b2aaf955db int
	__obf_6b154a2b56ba7139 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_7b98c3dafb8ea7fa string, __obf_7179f2bbaa5772e8 string) (*Downloader, error) {
	__obf_6b154a2b56ba7139, __obf_66cc3fda1d4448dc := parse.ParseFromNet(__obf_7179f2bbaa5772e8)
	if __obf_66cc3fda1d4448dc != nil {
		return nil, __obf_66cc3fda1d4448dc
	}
	var __obf_7c8db613d3d1c7a2 string
	// If no output folder specified, use current directory
	if __obf_7b98c3dafb8ea7fa == "" {
		__obf_ab5ade3ebd72dba0, __obf_66cc3fda1d4448dc := tool.CurrentDir()
		if __obf_66cc3fda1d4448dc != nil {
			return nil, __obf_66cc3fda1d4448dc
		}
		__obf_7c8db613d3d1c7a2 = filepath.Join(__obf_ab5ade3ebd72dba0, __obf_7b98c3dafb8ea7fa)
	} else {
		__obf_7c8db613d3d1c7a2 = __obf_7b98c3dafb8ea7fa
	}
	if __obf_66cc3fda1d4448dc := os.MkdirAll(__obf_7c8db613d3d1c7a2, os.ModePerm); __obf_66cc3fda1d4448dc != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_66cc3fda1d4448dc.Error())
	}
	__obf_0427b13a894d972f := filepath.Join(__obf_7c8db613d3d1c7a2, __obf_dcf01f9c274980af)
	if __obf_66cc3fda1d4448dc := os.MkdirAll(__obf_0427b13a894d972f, os.ModePerm); __obf_66cc3fda1d4448dc != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_0427b13a894d972f, __obf_66cc3fda1d4448dc.Error())
	}
	__obf_c605dccfc005ea0a := &Downloader{__obf_7c8db613d3d1c7a2: __obf_7c8db613d3d1c7a2, __obf_0427b13a894d972f: __obf_0427b13a894d972f, __obf_6b154a2b56ba7139: __obf_6b154a2b56ba7139}
	__obf_c605dccfc005ea0a.
		// segment length
		__obf_6fa4a1b2aaf955db = len(__obf_6b154a2b56ba7139.M3u8.Segments)
	__obf_c605dccfc005ea0a.__obf_9f8d055130477bb2 = __obf_da2edbbfaeb12089(__obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db)
	return __obf_c605dccfc005ea0a, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_c605dccfc005ea0a *Downloader) Start(__obf_3f827e42eb9c9eb0 int) error {
	var __obf_e8bc6ec043f47d75 sync.WaitGroup
	for {
		__obf_5cc906306884a0fe, __obf_604c3ee8487df3d2, __obf_66cc3fda1d4448dc := __obf_c605dccfc005ea0a.__obf_610b00b59d3bc685()
		if __obf_66cc3fda1d4448dc != nil {
			if __obf_604c3ee8487df3d2 {
				break
			}
			continue
		}
		__obf_e8bc6ec043f47d75.
			Add(1)
		go func(__obf_3f2a3f016c6cf94c int) {
			defer __obf_e8bc6ec043f47d75.Done()
			if __obf_66cc3fda1d4448dc := __obf_c605dccfc005ea0a.__obf_9849278af07d46ed(__obf_3f2a3f016c6cf94c); __obf_66cc3fda1d4448dc != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_66cc3fda1d4448dc.Error())
				if __obf_66cc3fda1d4448dc := __obf_c605dccfc005ea0a.__obf_ba9ddfd50f807361(__obf_3f2a3f016c6cf94c); __obf_66cc3fda1d4448dc != nil {
					fmt.Println(__obf_66cc3fda1d4448dc.Error())
				}
			}
		}(__obf_5cc906306884a0fe)
	}
	__obf_e8bc6ec043f47d75.
		Wait()
	if __obf_66cc3fda1d4448dc := __obf_c605dccfc005ea0a.__obf_dc574b8c278d0aa4(); __obf_66cc3fda1d4448dc != nil {
		return __obf_66cc3fda1d4448dc
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_c605dccfc005ea0a *Downloader) __obf_9849278af07d46ed(__obf_7fefef9ba4758fef int) error {
	__obf_1e763796ace3c745 := __obf_1e763796ace3c745(__obf_7fefef9ba4758fef)
	__obf_0aa5c6142f41b153 := __obf_c605dccfc005ea0a.__obf_51b45711581877ef(__obf_7fefef9ba4758fef)
	__obf_426a32721166711a, __obf_a987f4ef3b46ea5f := tool.Get(__obf_0aa5c6142f41b153)
	if __obf_a987f4ef3b46ea5f != nil {
		return fmt.Errorf("request %s, %s", __obf_0aa5c6142f41b153, __obf_a987f4ef3b46ea5f.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_426a32721166711a.Close()
	__obf_3d2564b949ca1c3f := filepath.Join(__obf_c605dccfc005ea0a.__obf_0427b13a894d972f, __obf_1e763796ace3c745)
	__obf_4311e106dab3e73b := __obf_3d2564b949ca1c3f + __obf_ac489e873476f8ee
	__obf_0cd81129580ad0e8, __obf_66cc3fda1d4448dc := os.Create(__obf_4311e106dab3e73b)
	if __obf_66cc3fda1d4448dc != nil {
		return fmt.Errorf("create file: %s, %s", __obf_1e763796ace3c745, __obf_66cc3fda1d4448dc.Error())
	}
	__obf_9c9734c5af92f70c, __obf_66cc3fda1d4448dc := io.ReadAll(__obf_426a32721166711a)
	if __obf_66cc3fda1d4448dc != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_0aa5c6142f41b153, __obf_66cc3fda1d4448dc.Error())
	}
	__obf_6678afcb4fec74f8 := __obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.M3u8.Segments[__obf_7fefef9ba4758fef]
	if __obf_6678afcb4fec74f8 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_7fefef9ba4758fef)
	}
	__obf_b0893e7bca992153, __obf_3c36bc1953326614 := __obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.Keys[__obf_6678afcb4fec74f8.KeyIndex]
	if __obf_3c36bc1953326614 && __obf_b0893e7bca992153 != "" {
		__obf_9c9734c5af92f70c, __obf_66cc3fda1d4448dc = tool.AES128Decrypt(__obf_9c9734c5af92f70c, []byte(__obf_b0893e7bca992153),
			[]byte(__obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.M3u8.Keys[__obf_6678afcb4fec74f8.KeyIndex].IV))
		if __obf_66cc3fda1d4448dc != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_0aa5c6142f41b153, __obf_66cc3fda1d4448dc.Error())
		}
	}
	__obf_7d22b75e742108f7 := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_67853b57773208ed := //0x47
		len(__obf_9c9734c5af92f70c)
	for __obf_80877d9a960f199e := 0; __obf_80877d9a960f199e < __obf_67853b57773208ed; __obf_80877d9a960f199e++ {
		if __obf_9c9734c5af92f70c[__obf_80877d9a960f199e] == __obf_7d22b75e742108f7 {
			__obf_9c9734c5af92f70c = __obf_9c9734c5af92f70c[__obf_80877d9a960f199e:]
			break
		}
	}
	__obf_5e1e5f15c3fd7a8e := bufio.NewWriter(__obf_0cd81129580ad0e8)
	if _, __obf_66cc3fda1d4448dc := __obf_5e1e5f15c3fd7a8e.Write(__obf_9c9734c5af92f70c); __obf_66cc3fda1d4448dc != nil {
		return fmt.Errorf("write to %s: %s", __obf_4311e106dab3e73b, __obf_66cc3fda1d4448dc.Error())
	}
	// Release file resource to rename file
	_ = __obf_0cd81129580ad0e8.Close()
	if __obf_66cc3fda1d4448dc = os.Rename(__obf_4311e106dab3e73b, __obf_3d2564b949ca1c3f); __obf_66cc3fda1d4448dc != nil {
		return __obf_66cc3fda1d4448dc
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_c605dccfc005ea0a.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_726c89b5f4d9bd88, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_c605dccfc005ea0a.__obf_726c89b5f4d9bd88)/float32(__obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db)*100, __obf_0aa5c6142f41b153)
	return nil
}

// next Obtains a next segment index
func (__obf_c605dccfc005ea0a *Downloader) __obf_610b00b59d3bc685() (__obf_7fefef9ba4758fef int, __obf_604c3ee8487df3d2 bool, __obf_66cc3fda1d4448dc error) {
	__obf_c605dccfc005ea0a.__obf_4713ca93e91cd1b0.
		Lock()
	defer __obf_c605dccfc005ea0a.__obf_4713ca93e91cd1b0.Unlock()
	if len(__obf_c605dccfc005ea0a.__obf_9f8d055130477bb2) == 0 {
		__obf_66cc3fda1d4448dc = fmt.Errorf("queue empty")
		if __obf_c605dccfc005ea0a.__obf_726c89b5f4d9bd88 == int32(__obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db) {
			__obf_604c3ee8487df3d2 = true
			return
		}
		__obf_604c3ee8487df3d2 = // Some segment indexes are still running.
			false
		return
	}
	__obf_7fefef9ba4758fef = __obf_c605dccfc005ea0a.__obf_9f8d055130477bb2[0]
	__obf_c605dccfc005ea0a.__obf_9f8d055130477bb2 = __obf_c605dccfc005ea0a.__obf_9f8d055130477bb2[1:]
	return
}

// back For retry while download failed
func (__obf_c605dccfc005ea0a *Downloader) __obf_ba9ddfd50f807361(__obf_7fefef9ba4758fef int) error {
	__obf_c605dccfc005ea0a.__obf_4713ca93e91cd1b0.
		Lock()
	defer __obf_c605dccfc005ea0a.__obf_4713ca93e91cd1b0.Unlock()
	if __obf_6678afcb4fec74f8 := __obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.M3u8.Segments[__obf_7fefef9ba4758fef]; __obf_6678afcb4fec74f8 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_7fefef9ba4758fef)
	}
	__obf_c605dccfc005ea0a.__obf_9f8d055130477bb2 = append(__obf_c605dccfc005ea0a.__obf_9f8d055130477bb2,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_7fefef9ba4758fef)
	return nil
}

func (__obf_c605dccfc005ea0a *Downloader) __obf_dc574b8c278d0aa4() error {
	__obf_b75014e459d3e7b5 := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_3f2a3f016c6cf94c := 0; __obf_3f2a3f016c6cf94c < __obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db; __obf_3f2a3f016c6cf94c++ {
		__obf_1e763796ace3c745 := __obf_1e763796ace3c745(__obf_3f2a3f016c6cf94c)
		__obf_0cd81129580ad0e8 := filepath.Join(__obf_c605dccfc005ea0a.__obf_0427b13a894d972f, __obf_1e763796ace3c745)
		if _, __obf_66cc3fda1d4448dc := os.Stat(__obf_0cd81129580ad0e8); __obf_66cc3fda1d4448dc != nil {
			__obf_b75014e459d3e7b5++
		}
	}
	if __obf_b75014e459d3e7b5 > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_b75014e459d3e7b5)
	}
	__obf_c4a8999b8ecf7ded := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_c605dccfc005ea0a.__obf_7c8db613d3d1c7a2, __obf_4eebfa8f7b7e45e5)
	__obf_06c8b40600c50d8c, __obf_66cc3fda1d4448dc := os.Create(__obf_c4a8999b8ecf7ded)
	if __obf_66cc3fda1d4448dc != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_66cc3fda1d4448dc.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_06c8b40600c50d8c.Close()
	__obf_7f66351e3a00db69 := bufio.NewWriter(__obf_06c8b40600c50d8c)
	__obf_7eb765ce2ed6fb1d := 0
	for __obf_7fefef9ba4758fef := 0; __obf_7fefef9ba4758fef < __obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db; __obf_7fefef9ba4758fef++ {
		__obf_1e763796ace3c745 := __obf_1e763796ace3c745(__obf_7fefef9ba4758fef)
		__obf_9c9734c5af92f70c, __obf_66cc3fda1d4448dc := os.ReadFile(filepath.Join(__obf_c605dccfc005ea0a.__obf_0427b13a894d972f, __obf_1e763796ace3c745))
		_, __obf_66cc3fda1d4448dc = __obf_7f66351e3a00db69.Write(__obf_9c9734c5af92f70c)
		if __obf_66cc3fda1d4448dc != nil {
			continue
		}
		__obf_7eb765ce2ed6fb1d++
		tool.DrawProgressBar("merge",
			float32(__obf_7eb765ce2ed6fb1d)/float32(__obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db), __obf_aa9349b594dc6e4f)
	}
	_ = __obf_7f66351e3a00db69.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_c605dccfc005ea0a.__obf_0427b13a894d972f)

	if __obf_7eb765ce2ed6fb1d != __obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db {
		fmt.Printf("[warning] \n%d files merge failed", __obf_c605dccfc005ea0a.__obf_6fa4a1b2aaf955db-__obf_7eb765ce2ed6fb1d)
	}

	fmt.Printf("\n[output] %s\n", __obf_c4a8999b8ecf7ded)

	return nil
}

func (__obf_c605dccfc005ea0a *Downloader) __obf_51b45711581877ef(__obf_7fefef9ba4758fef int) string {
	__obf_be84f7f2fb1624ea := __obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.M3u8.Segments[__obf_7fefef9ba4758fef]
	return tool.ResolveURL(__obf_c605dccfc005ea0a.__obf_6b154a2b56ba7139.URL, __obf_be84f7f2fb1624ea.URI)
}

func __obf_1e763796ace3c745(__obf_47ca6f0eab9c068b int) string {
	return strconv.Itoa(__obf_47ca6f0eab9c068b) + __obf_275560fef6ecadbd
}

func __obf_da2edbbfaeb12089(len int) []int {
	__obf_d904bce20f8fd4d8 := make([]int, 0)
	for __obf_cac037f39ec875a2 := 0; __obf_cac037f39ec875a2 < len; __obf_cac037f39ec875a2++ {
		__obf_d904bce20f8fd4d8 = append(__obf_d904bce20f8fd4d8, __obf_cac037f39ec875a2)
	}
	return __obf_d904bce20f8fd4d8
}
