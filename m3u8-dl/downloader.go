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
	__obf_8134ed3b40be5ec0 = ".ts"
	__obf_2a4429725b53e4f1 = "ts"
	__obf_3efda21d9cb3dadc = "main.ts"
	__obf_cfdbcb71f0de2169 = "_tmp"
	__obf_13326195de1124fb = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_25d51c9d676b4ae0 sync.Mutex
	__obf_9d62fc2ea68dbc90 []int
	__obf_f2f5d889f82413c8 string
	__obf_89fa31665130ded7 string
	__obf_3a501d820e48bd04 int32
	__obf_40556a43ecb1a12f int
	__obf_9e5ac2f699f8eb94 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_10b0d572c2650427 string, __obf_c0fb6d8692497149 string) (*Downloader, error) {
	__obf_9e5ac2f699f8eb94, __obf_e4142ef6a9f97155 := parse.ParseFromNet(__obf_c0fb6d8692497149)
	if __obf_e4142ef6a9f97155 != nil {
		return nil, __obf_e4142ef6a9f97155
	}
	var __obf_f2f5d889f82413c8 string
	// If no output folder specified, use current directory
	if __obf_10b0d572c2650427 == "" {
		__obf_fa11f70f79569523, __obf_e4142ef6a9f97155 := tool.CurrentDir()
		if __obf_e4142ef6a9f97155 != nil {
			return nil, __obf_e4142ef6a9f97155
		}
		__obf_f2f5d889f82413c8 = filepath.Join(__obf_fa11f70f79569523, __obf_10b0d572c2650427)
	} else {
		__obf_f2f5d889f82413c8 = __obf_10b0d572c2650427
	}
	if __obf_e4142ef6a9f97155 := os.MkdirAll(__obf_f2f5d889f82413c8, os.ModePerm); __obf_e4142ef6a9f97155 != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_e4142ef6a9f97155.Error())
	}
	__obf_89fa31665130ded7 := filepath.Join(__obf_f2f5d889f82413c8, __obf_2a4429725b53e4f1)
	if __obf_e4142ef6a9f97155 := os.MkdirAll(__obf_89fa31665130ded7, os.ModePerm); __obf_e4142ef6a9f97155 != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_89fa31665130ded7, __obf_e4142ef6a9f97155.Error())
	}
	__obf_fd24043101b08638 := &Downloader{__obf_f2f5d889f82413c8: __obf_f2f5d889f82413c8, __obf_89fa31665130ded7: __obf_89fa31665130ded7, __obf_9e5ac2f699f8eb94: __obf_9e5ac2f699f8eb94}
	__obf_fd24043101b08638.
		// segment length
		__obf_40556a43ecb1a12f = len(__obf_9e5ac2f699f8eb94.M3u8.Segments)
	__obf_fd24043101b08638.__obf_9d62fc2ea68dbc90 = __obf_d7fefa73b06977df(__obf_fd24043101b08638.__obf_40556a43ecb1a12f)
	return __obf_fd24043101b08638, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_fd24043101b08638 *Downloader) Start(__obf_8e57269ee4925374 int) error {
	var __obf_53d040c6dd8b08e1 sync.WaitGroup
	for {
		__obf_2f55b442f88aae44, __obf_bfb95eebc5e13071, __obf_e4142ef6a9f97155 := __obf_fd24043101b08638.__obf_44b7fcc7ee626cad()
		if __obf_e4142ef6a9f97155 != nil {
			if __obf_bfb95eebc5e13071 {
				break
			}
			continue
		}
		__obf_53d040c6dd8b08e1.
			Add(1)
		go func(__obf_0a3fc1c23d8a6530 int) {
			defer __obf_53d040c6dd8b08e1.Done()
			if __obf_e4142ef6a9f97155 := __obf_fd24043101b08638.__obf_78c84d64eed36d3d(__obf_0a3fc1c23d8a6530); __obf_e4142ef6a9f97155 != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_e4142ef6a9f97155.Error())
				if __obf_e4142ef6a9f97155 := __obf_fd24043101b08638.__obf_b413c90d56ac4fb5(__obf_0a3fc1c23d8a6530); __obf_e4142ef6a9f97155 != nil {
					fmt.Println(__obf_e4142ef6a9f97155.Error())
				}
			}
		}(__obf_2f55b442f88aae44)
	}
	__obf_53d040c6dd8b08e1.
		Wait()
	if __obf_e4142ef6a9f97155 := __obf_fd24043101b08638.__obf_5b59d34954c034e7(); __obf_e4142ef6a9f97155 != nil {
		return __obf_e4142ef6a9f97155
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_fd24043101b08638 *Downloader) __obf_78c84d64eed36d3d(__obf_c0be86a4073a9ff6 int) error {
	__obf_0f053e4f910a2b8f := __obf_0f053e4f910a2b8f(__obf_c0be86a4073a9ff6)
	__obf_55146efecadaa805 := __obf_fd24043101b08638.__obf_3484741e53c56a80(__obf_c0be86a4073a9ff6)
	__obf_54375b0959069213, __obf_8f979500c83a748b := tool.Get(__obf_55146efecadaa805)
	if __obf_8f979500c83a748b != nil {
		return fmt.Errorf("request %s, %s", __obf_55146efecadaa805, __obf_8f979500c83a748b.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_54375b0959069213.Close()
	__obf_fc854cc0a8652d8d := filepath.Join(__obf_fd24043101b08638.__obf_89fa31665130ded7, __obf_0f053e4f910a2b8f)
	__obf_44e38527396dda6d := __obf_fc854cc0a8652d8d + __obf_cfdbcb71f0de2169
	__obf_51dbed2d2291578d, __obf_e4142ef6a9f97155 := os.Create(__obf_44e38527396dda6d)
	if __obf_e4142ef6a9f97155 != nil {
		return fmt.Errorf("create file: %s, %s", __obf_0f053e4f910a2b8f, __obf_e4142ef6a9f97155.Error())
	}
	__obf_55287880a927d1dc, __obf_e4142ef6a9f97155 := io.ReadAll(__obf_54375b0959069213)
	if __obf_e4142ef6a9f97155 != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_55146efecadaa805, __obf_e4142ef6a9f97155.Error())
	}
	__obf_eae1901d07da7ef1 := __obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.M3u8.Segments[__obf_c0be86a4073a9ff6]
	if __obf_eae1901d07da7ef1 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_c0be86a4073a9ff6)
	}
	__obf_2e138e70c2e8d0b7, __obf_9f4deb31efb767e9 := __obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.Keys[__obf_eae1901d07da7ef1.KeyIndex]
	if __obf_9f4deb31efb767e9 && __obf_2e138e70c2e8d0b7 != "" {
		__obf_55287880a927d1dc, __obf_e4142ef6a9f97155 = tool.AES128Decrypt(__obf_55287880a927d1dc, []byte(__obf_2e138e70c2e8d0b7),
			[]byte(__obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.M3u8.Keys[__obf_eae1901d07da7ef1.KeyIndex].IV))
		if __obf_e4142ef6a9f97155 != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_55146efecadaa805, __obf_e4142ef6a9f97155.Error())
		}
	}
	__obf_ee460da7247562d1 := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_6bf8f6457ac7abb7 := //0x47
		len(__obf_55287880a927d1dc)
	for __obf_8acc6b7163bcbd5c := 0; __obf_8acc6b7163bcbd5c < __obf_6bf8f6457ac7abb7; __obf_8acc6b7163bcbd5c++ {
		if __obf_55287880a927d1dc[__obf_8acc6b7163bcbd5c] == __obf_ee460da7247562d1 {
			__obf_55287880a927d1dc = __obf_55287880a927d1dc[__obf_8acc6b7163bcbd5c:]
			break
		}
	}
	__obf_be982d30b47873dc := bufio.NewWriter(__obf_51dbed2d2291578d)
	if _, __obf_e4142ef6a9f97155 := __obf_be982d30b47873dc.Write(__obf_55287880a927d1dc); __obf_e4142ef6a9f97155 != nil {
		return fmt.Errorf("write to %s: %s", __obf_44e38527396dda6d, __obf_e4142ef6a9f97155.Error())
	}
	// Release file resource to rename file
	_ = __obf_51dbed2d2291578d.Close()
	if __obf_e4142ef6a9f97155 = os.Rename(__obf_44e38527396dda6d, __obf_fc854cc0a8652d8d); __obf_e4142ef6a9f97155 != nil {
		return __obf_e4142ef6a9f97155
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_fd24043101b08638.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_3a501d820e48bd04, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_fd24043101b08638.__obf_3a501d820e48bd04)/float32(__obf_fd24043101b08638.__obf_40556a43ecb1a12f)*100, __obf_55146efecadaa805)
	return nil
}

// next Obtains a next segment index
func (__obf_fd24043101b08638 *Downloader) __obf_44b7fcc7ee626cad() (__obf_c0be86a4073a9ff6 int, __obf_bfb95eebc5e13071 bool, __obf_e4142ef6a9f97155 error) {
	__obf_fd24043101b08638.__obf_25d51c9d676b4ae0.
		Lock()
	defer __obf_fd24043101b08638.__obf_25d51c9d676b4ae0.Unlock()
	if len(__obf_fd24043101b08638.__obf_9d62fc2ea68dbc90) == 0 {
		__obf_e4142ef6a9f97155 = fmt.Errorf("queue empty")
		if __obf_fd24043101b08638.__obf_3a501d820e48bd04 == int32(__obf_fd24043101b08638.__obf_40556a43ecb1a12f) {
			__obf_bfb95eebc5e13071 = true
			return
		}
		__obf_bfb95eebc5e13071 = // Some segment indexes are still running.
			false
		return
	}
	__obf_c0be86a4073a9ff6 = __obf_fd24043101b08638.__obf_9d62fc2ea68dbc90[0]
	__obf_fd24043101b08638.__obf_9d62fc2ea68dbc90 = __obf_fd24043101b08638.__obf_9d62fc2ea68dbc90[1:]
	return
}

// back For retry while download failed
func (__obf_fd24043101b08638 *Downloader) __obf_b413c90d56ac4fb5(__obf_c0be86a4073a9ff6 int) error {
	__obf_fd24043101b08638.__obf_25d51c9d676b4ae0.
		Lock()
	defer __obf_fd24043101b08638.__obf_25d51c9d676b4ae0.Unlock()
	if __obf_eae1901d07da7ef1 := __obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.M3u8.Segments[__obf_c0be86a4073a9ff6]; __obf_eae1901d07da7ef1 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_c0be86a4073a9ff6)
	}
	__obf_fd24043101b08638.__obf_9d62fc2ea68dbc90 = append(__obf_fd24043101b08638.__obf_9d62fc2ea68dbc90,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_c0be86a4073a9ff6)
	return nil
}

func (__obf_fd24043101b08638 *Downloader) __obf_5b59d34954c034e7() error {
	__obf_39393a4f7f590b80 := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_0a3fc1c23d8a6530 := 0; __obf_0a3fc1c23d8a6530 < __obf_fd24043101b08638.__obf_40556a43ecb1a12f; __obf_0a3fc1c23d8a6530++ {
		__obf_0f053e4f910a2b8f := __obf_0f053e4f910a2b8f(__obf_0a3fc1c23d8a6530)
		__obf_51dbed2d2291578d := filepath.Join(__obf_fd24043101b08638.__obf_89fa31665130ded7, __obf_0f053e4f910a2b8f)
		if _, __obf_e4142ef6a9f97155 := os.Stat(__obf_51dbed2d2291578d); __obf_e4142ef6a9f97155 != nil {
			__obf_39393a4f7f590b80++
		}
	}
	if __obf_39393a4f7f590b80 > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_39393a4f7f590b80)
	}
	__obf_b756b21b369ce28a := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_fd24043101b08638.__obf_f2f5d889f82413c8, __obf_3efda21d9cb3dadc)
	__obf_14490e384dec455a, __obf_e4142ef6a9f97155 := os.Create(__obf_b756b21b369ce28a)
	if __obf_e4142ef6a9f97155 != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_e4142ef6a9f97155.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_14490e384dec455a.Close()
	__obf_7c48766fff85de69 := bufio.NewWriter(__obf_14490e384dec455a)
	__obf_96b1fcae764dba3e := 0
	for __obf_c0be86a4073a9ff6 := 0; __obf_c0be86a4073a9ff6 < __obf_fd24043101b08638.__obf_40556a43ecb1a12f; __obf_c0be86a4073a9ff6++ {
		__obf_0f053e4f910a2b8f := __obf_0f053e4f910a2b8f(__obf_c0be86a4073a9ff6)
		__obf_55287880a927d1dc, __obf_e4142ef6a9f97155 := os.ReadFile(filepath.Join(__obf_fd24043101b08638.__obf_89fa31665130ded7, __obf_0f053e4f910a2b8f))
		_, __obf_e4142ef6a9f97155 = __obf_7c48766fff85de69.Write(__obf_55287880a927d1dc)
		if __obf_e4142ef6a9f97155 != nil {
			continue
		}
		__obf_96b1fcae764dba3e++
		tool.DrawProgressBar("merge",
			float32(__obf_96b1fcae764dba3e)/float32(__obf_fd24043101b08638.__obf_40556a43ecb1a12f), __obf_13326195de1124fb)
	}
	_ = __obf_7c48766fff85de69.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_fd24043101b08638.__obf_89fa31665130ded7)

	if __obf_96b1fcae764dba3e != __obf_fd24043101b08638.__obf_40556a43ecb1a12f {
		fmt.Printf("[warning] \n%d files merge failed", __obf_fd24043101b08638.__obf_40556a43ecb1a12f-__obf_96b1fcae764dba3e)
	}

	fmt.Printf("\n[output] %s\n", __obf_b756b21b369ce28a)

	return nil
}

func (__obf_fd24043101b08638 *Downloader) __obf_3484741e53c56a80(__obf_c0be86a4073a9ff6 int) string {
	__obf_babfb9bf227c7bd1 := __obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.M3u8.Segments[__obf_c0be86a4073a9ff6]
	return tool.ResolveURL(__obf_fd24043101b08638.__obf_9e5ac2f699f8eb94.URL, __obf_babfb9bf227c7bd1.URI)
}

func __obf_0f053e4f910a2b8f(__obf_acb8fd4bf8e7bf04 int) string {
	return strconv.Itoa(__obf_acb8fd4bf8e7bf04) + __obf_8134ed3b40be5ec0
}

func __obf_d7fefa73b06977df(len int) []int {
	__obf_708f90e50471206e := make([]int, 0)
	for __obf_bff5d8ee2cb013b3 := 0; __obf_bff5d8ee2cb013b3 < len; __obf_bff5d8ee2cb013b3++ {
		__obf_708f90e50471206e = append(__obf_708f90e50471206e, __obf_bff5d8ee2cb013b3)
	}
	return __obf_708f90e50471206e
}
