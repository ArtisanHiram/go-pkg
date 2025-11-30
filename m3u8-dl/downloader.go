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
	__obf_b3f684960c08754a = ".ts"
	__obf_f895b56be5a7bc8f = "ts"
	__obf_e554de360e3cfb29 = "main.ts"
	__obf_1db7a3f0d8198c0e = "_tmp"
	__obf_6fb978fe522a73e6 = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_2b9d56b267b50c41 sync.Mutex
	__obf_97803d52ef8a4ae9 []int
	__obf_a8416c21393574e6 string
	__obf_3122dc308e76160e string
	__obf_2752942958276981 int32
	__obf_1b9eb5e1d985a818 int
	__obf_bb9b1fc3a10ee560 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_6050b974c19fc9b9 string, __obf_cbdb4071e572fbe6 string) (*Downloader, error) {
	__obf_bb9b1fc3a10ee560, __obf_31a042aea6f8bcd2 := parse.ParseFromNet(__obf_cbdb4071e572fbe6)
	if __obf_31a042aea6f8bcd2 != nil {
		return nil, __obf_31a042aea6f8bcd2
	}
	var __obf_a8416c21393574e6 string
	// If no output folder specified, use current directory
	if __obf_6050b974c19fc9b9 == "" {
		__obf_7f02d1bc411a32b8, __obf_31a042aea6f8bcd2 := tool.CurrentDir()
		if __obf_31a042aea6f8bcd2 != nil {
			return nil, __obf_31a042aea6f8bcd2
		}
		__obf_a8416c21393574e6 = filepath.Join(__obf_7f02d1bc411a32b8, __obf_6050b974c19fc9b9)
	} else {
		__obf_a8416c21393574e6 = __obf_6050b974c19fc9b9
	}
	if __obf_31a042aea6f8bcd2 := os.MkdirAll(__obf_a8416c21393574e6, os.ModePerm); __obf_31a042aea6f8bcd2 != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_31a042aea6f8bcd2.Error())
	}
	__obf_3122dc308e76160e := filepath.Join(__obf_a8416c21393574e6, __obf_f895b56be5a7bc8f)
	if __obf_31a042aea6f8bcd2 := os.MkdirAll(__obf_3122dc308e76160e, os.ModePerm); __obf_31a042aea6f8bcd2 != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_3122dc308e76160e, __obf_31a042aea6f8bcd2.Error())
	}
	__obf_1351f1973c47f7c5 := &Downloader{__obf_a8416c21393574e6: __obf_a8416c21393574e6, __obf_3122dc308e76160e: __obf_3122dc308e76160e, __obf_bb9b1fc3a10ee560: __obf_bb9b1fc3a10ee560}
	__obf_1351f1973c47f7c5.
		// segment length
		__obf_1b9eb5e1d985a818 = len(__obf_bb9b1fc3a10ee560.M3u8.Segments)
	__obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9 = __obf_e64fbe484672d4a2(__obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818)
	return __obf_1351f1973c47f7c5, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_1351f1973c47f7c5 *Downloader) Start(__obf_3fd813ce7a9e78db int) error {
	var __obf_29e6aa1b88f2a7f7 sync.WaitGroup
	for {
		__obf_98075bfe7dcb4ca9, __obf_6abe187615665e1f, __obf_31a042aea6f8bcd2 := __obf_1351f1973c47f7c5.__obf_403dc38ef01aa297()
		if __obf_31a042aea6f8bcd2 != nil {
			if __obf_6abe187615665e1f {
				break
			}
			continue
		}
		__obf_29e6aa1b88f2a7f7.
			Add(1)
		go func(__obf_64c3e7ad3b980280 int) {
			defer __obf_29e6aa1b88f2a7f7.Done()
			if __obf_31a042aea6f8bcd2 := __obf_1351f1973c47f7c5.__obf_f18e57b787ed036b(__obf_64c3e7ad3b980280); __obf_31a042aea6f8bcd2 != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_31a042aea6f8bcd2.Error())
				if __obf_31a042aea6f8bcd2 := __obf_1351f1973c47f7c5.__obf_f443636124ad07cd(__obf_64c3e7ad3b980280); __obf_31a042aea6f8bcd2 != nil {
					fmt.Println(__obf_31a042aea6f8bcd2.Error())
				}
			}
		}(__obf_98075bfe7dcb4ca9)
	}
	__obf_29e6aa1b88f2a7f7.
		Wait()
	if __obf_31a042aea6f8bcd2 := __obf_1351f1973c47f7c5.__obf_f475a41edd439924(); __obf_31a042aea6f8bcd2 != nil {
		return __obf_31a042aea6f8bcd2
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_1351f1973c47f7c5 *Downloader) __obf_f18e57b787ed036b(__obf_3b72347dfca6c90a int) error {
	__obf_b0a80c3ce5a1fc64 := __obf_b0a80c3ce5a1fc64(__obf_3b72347dfca6c90a)
	__obf_c9ac9a6b87ca681c := __obf_1351f1973c47f7c5.__obf_4a67b82d07469229(__obf_3b72347dfca6c90a)
	__obf_339cb2e165e99d30, __obf_e8cec24a0089b609 := tool.Get(__obf_c9ac9a6b87ca681c)
	if __obf_e8cec24a0089b609 != nil {
		return fmt.Errorf("request %s, %s", __obf_c9ac9a6b87ca681c, __obf_e8cec24a0089b609.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_339cb2e165e99d30.Close()
	__obf_5f9a5a6591ef20fd := filepath.Join(__obf_1351f1973c47f7c5.__obf_3122dc308e76160e, __obf_b0a80c3ce5a1fc64)
	__obf_a9aca0adfab8d602 := __obf_5f9a5a6591ef20fd + __obf_1db7a3f0d8198c0e
	__obf_14117a50f154fb12, __obf_31a042aea6f8bcd2 := os.Create(__obf_a9aca0adfab8d602)
	if __obf_31a042aea6f8bcd2 != nil {
		return fmt.Errorf("create file: %s, %s", __obf_b0a80c3ce5a1fc64, __obf_31a042aea6f8bcd2.Error())
	}
	__obf_f6aa9cefed36c9eb, __obf_31a042aea6f8bcd2 := io.ReadAll(__obf_339cb2e165e99d30)
	if __obf_31a042aea6f8bcd2 != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_c9ac9a6b87ca681c, __obf_31a042aea6f8bcd2.Error())
	}
	__obf_89d28c6b364e5392 := __obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.M3u8.Segments[__obf_3b72347dfca6c90a]
	if __obf_89d28c6b364e5392 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_3b72347dfca6c90a)
	}
	__obf_7600f6160cdf3b52, __obf_9a9f5372d175780c := __obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.Keys[__obf_89d28c6b364e5392.KeyIndex]
	if __obf_9a9f5372d175780c && __obf_7600f6160cdf3b52 != "" {
		__obf_f6aa9cefed36c9eb, __obf_31a042aea6f8bcd2 = tool.AES128Decrypt(__obf_f6aa9cefed36c9eb, []byte(__obf_7600f6160cdf3b52),
			[]byte(__obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.M3u8.Keys[__obf_89d28c6b364e5392.KeyIndex].IV))
		if __obf_31a042aea6f8bcd2 != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_c9ac9a6b87ca681c, __obf_31a042aea6f8bcd2.Error())
		}
	}
	__obf_1d9869f2b261f1be := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_0613793bd152ed2a := //0x47
		len(__obf_f6aa9cefed36c9eb)
	for __obf_6677360e629b08d6 := 0; __obf_6677360e629b08d6 < __obf_0613793bd152ed2a; __obf_6677360e629b08d6++ {
		if __obf_f6aa9cefed36c9eb[__obf_6677360e629b08d6] == __obf_1d9869f2b261f1be {
			__obf_f6aa9cefed36c9eb = __obf_f6aa9cefed36c9eb[__obf_6677360e629b08d6:]
			break
		}
	}
	__obf_ed2640f6ce27f175 := bufio.NewWriter(__obf_14117a50f154fb12)
	if _, __obf_31a042aea6f8bcd2 := __obf_ed2640f6ce27f175.Write(__obf_f6aa9cefed36c9eb); __obf_31a042aea6f8bcd2 != nil {
		return fmt.Errorf("write to %s: %s", __obf_a9aca0adfab8d602, __obf_31a042aea6f8bcd2.Error())
	}
	// Release file resource to rename file
	_ = __obf_14117a50f154fb12.Close()
	if __obf_31a042aea6f8bcd2 = os.Rename(__obf_a9aca0adfab8d602, __obf_5f9a5a6591ef20fd); __obf_31a042aea6f8bcd2 != nil {
		return __obf_31a042aea6f8bcd2
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_1351f1973c47f7c5.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_2752942958276981, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_1351f1973c47f7c5.__obf_2752942958276981)/float32(__obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818)*100, __obf_c9ac9a6b87ca681c)
	return nil
}

// next Obtains a next segment index
func (__obf_1351f1973c47f7c5 *Downloader) __obf_403dc38ef01aa297() (__obf_3b72347dfca6c90a int, __obf_6abe187615665e1f bool, __obf_31a042aea6f8bcd2 error) {
	__obf_1351f1973c47f7c5.__obf_2b9d56b267b50c41.
		Lock()
	defer __obf_1351f1973c47f7c5.__obf_2b9d56b267b50c41.Unlock()
	if len(__obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9) == 0 {
		__obf_31a042aea6f8bcd2 = fmt.Errorf("queue empty")
		if __obf_1351f1973c47f7c5.__obf_2752942958276981 == int32(__obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818) {
			__obf_6abe187615665e1f = true
			return
		}
		__obf_6abe187615665e1f = // Some segment indexes are still running.
			false
		return
	}
	__obf_3b72347dfca6c90a = __obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9[0]
	__obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9 = __obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9[1:]
	return
}

// back For retry while download failed
func (__obf_1351f1973c47f7c5 *Downloader) __obf_f443636124ad07cd(__obf_3b72347dfca6c90a int) error {
	__obf_1351f1973c47f7c5.__obf_2b9d56b267b50c41.
		Lock()
	defer __obf_1351f1973c47f7c5.__obf_2b9d56b267b50c41.Unlock()
	if __obf_89d28c6b364e5392 := __obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.M3u8.Segments[__obf_3b72347dfca6c90a]; __obf_89d28c6b364e5392 == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_3b72347dfca6c90a)
	}
	__obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9 = append(__obf_1351f1973c47f7c5.__obf_97803d52ef8a4ae9,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_3b72347dfca6c90a)
	return nil
}

func (__obf_1351f1973c47f7c5 *Downloader) __obf_f475a41edd439924() error {
	__obf_dde9638b67efc9cd := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_64c3e7ad3b980280 := 0; __obf_64c3e7ad3b980280 < __obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818; __obf_64c3e7ad3b980280++ {
		__obf_b0a80c3ce5a1fc64 := __obf_b0a80c3ce5a1fc64(__obf_64c3e7ad3b980280)
		__obf_14117a50f154fb12 := filepath.Join(__obf_1351f1973c47f7c5.__obf_3122dc308e76160e, __obf_b0a80c3ce5a1fc64)
		if _, __obf_31a042aea6f8bcd2 := os.Stat(__obf_14117a50f154fb12); __obf_31a042aea6f8bcd2 != nil {
			__obf_dde9638b67efc9cd++
		}
	}
	if __obf_dde9638b67efc9cd > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_dde9638b67efc9cd)
	}
	__obf_9e65f590a855664f := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_1351f1973c47f7c5.__obf_a8416c21393574e6, __obf_e554de360e3cfb29)
	__obf_0eba583ca96884e2, __obf_31a042aea6f8bcd2 := os.Create(__obf_9e65f590a855664f)
	if __obf_31a042aea6f8bcd2 != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_31a042aea6f8bcd2.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_0eba583ca96884e2.Close()
	__obf_3b84b549517f3909 := bufio.NewWriter(__obf_0eba583ca96884e2)
	__obf_d99217209b868a16 := 0
	for __obf_3b72347dfca6c90a := 0; __obf_3b72347dfca6c90a < __obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818; __obf_3b72347dfca6c90a++ {
		__obf_b0a80c3ce5a1fc64 := __obf_b0a80c3ce5a1fc64(__obf_3b72347dfca6c90a)
		__obf_f6aa9cefed36c9eb, __obf_31a042aea6f8bcd2 := os.ReadFile(filepath.Join(__obf_1351f1973c47f7c5.__obf_3122dc308e76160e, __obf_b0a80c3ce5a1fc64))
		_, __obf_31a042aea6f8bcd2 = __obf_3b84b549517f3909.Write(__obf_f6aa9cefed36c9eb)
		if __obf_31a042aea6f8bcd2 != nil {
			continue
		}
		__obf_d99217209b868a16++
		tool.DrawProgressBar("merge",
			float32(__obf_d99217209b868a16)/float32(__obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818), __obf_6fb978fe522a73e6)
	}
	_ = __obf_3b84b549517f3909.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_1351f1973c47f7c5.__obf_3122dc308e76160e)

	if __obf_d99217209b868a16 != __obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818 {
		fmt.Printf("[warning] \n%d files merge failed", __obf_1351f1973c47f7c5.__obf_1b9eb5e1d985a818-__obf_d99217209b868a16)
	}

	fmt.Printf("\n[output] %s\n", __obf_9e65f590a855664f)

	return nil
}

func (__obf_1351f1973c47f7c5 *Downloader) __obf_4a67b82d07469229(__obf_3b72347dfca6c90a int) string {
	__obf_d2dafbde3aab76d4 := __obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.M3u8.Segments[__obf_3b72347dfca6c90a]
	return tool.ResolveURL(__obf_1351f1973c47f7c5.__obf_bb9b1fc3a10ee560.URL, __obf_d2dafbde3aab76d4.URI)
}

func __obf_b0a80c3ce5a1fc64(__obf_e714bfe896a1efb9 int) string {
	return strconv.Itoa(__obf_e714bfe896a1efb9) + __obf_b3f684960c08754a
}

func __obf_e64fbe484672d4a2(len int) []int {
	__obf_b2e0e1c8fe29f928 := make([]int, 0)
	for __obf_4af225571639bdad := 0; __obf_4af225571639bdad < len; __obf_4af225571639bdad++ {
		__obf_b2e0e1c8fe29f928 = append(__obf_b2e0e1c8fe29f928, __obf_4af225571639bdad)
	}
	return __obf_b2e0e1c8fe29f928
}
