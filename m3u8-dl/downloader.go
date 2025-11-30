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
	__obf_412f29766929b45a = ".ts"
	__obf_dd2bbf033b0e083f = "ts"
	__obf_0a27001219a6287c = "main.ts"
	__obf_14fb6f7b461646ca = "_tmp"
	__obf_74e246ed5dfcfa3b = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_7328558709def097 sync.Mutex
	__obf_443bdea9870aa393 []int
	__obf_8668dd1d7cbf68d4 string
	__obf_40d94e32155209e0 string
	__obf_4130e670c70f620a int32
	__obf_774318f00ce2acad int
	__obf_02b78213e8219b62 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_75dc0fa65ea3ee4a string, __obf_fba5742c8939d4d5 string) (*Downloader, error) {
	__obf_02b78213e8219b62, __obf_c3746095cbc77276 := parse.ParseFromNet(__obf_fba5742c8939d4d5)
	if __obf_c3746095cbc77276 != nil {
		return nil, __obf_c3746095cbc77276
	}
	var __obf_8668dd1d7cbf68d4 string
	// If no output folder specified, use current directory
	if __obf_75dc0fa65ea3ee4a == "" {
		__obf_d859d0f728498cfb, __obf_c3746095cbc77276 := tool.CurrentDir()
		if __obf_c3746095cbc77276 != nil {
			return nil, __obf_c3746095cbc77276
		}
		__obf_8668dd1d7cbf68d4 = filepath.Join(__obf_d859d0f728498cfb, __obf_75dc0fa65ea3ee4a)
	} else {
		__obf_8668dd1d7cbf68d4 = __obf_75dc0fa65ea3ee4a
	}
	if __obf_c3746095cbc77276 := os.MkdirAll(__obf_8668dd1d7cbf68d4, os.ModePerm); __obf_c3746095cbc77276 != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_c3746095cbc77276.Error())
	}
	__obf_40d94e32155209e0 := filepath.Join(__obf_8668dd1d7cbf68d4, __obf_dd2bbf033b0e083f)
	if __obf_c3746095cbc77276 := os.MkdirAll(__obf_40d94e32155209e0, os.ModePerm); __obf_c3746095cbc77276 != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_40d94e32155209e0, __obf_c3746095cbc77276.Error())
	}
	__obf_73861903bc064f83 := &Downloader{__obf_8668dd1d7cbf68d4: __obf_8668dd1d7cbf68d4, __obf_40d94e32155209e0: __obf_40d94e32155209e0, __obf_02b78213e8219b62: __obf_02b78213e8219b62}
	__obf_73861903bc064f83.
		// segment length
		__obf_774318f00ce2acad = len(__obf_02b78213e8219b62.M3u8.Segments)
	__obf_73861903bc064f83.__obf_443bdea9870aa393 = __obf_34bcc24b4504f48e(__obf_73861903bc064f83.__obf_774318f00ce2acad)
	return __obf_73861903bc064f83, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_73861903bc064f83 *Downloader) Start(__obf_f0bae1bec99f1840 int) error {
	var __obf_2b37cb05c9916766 sync.WaitGroup
	for {
		__obf_8359bb0a6d11dd9a, __obf_6ff5eadfa557e5dd, __obf_c3746095cbc77276 := __obf_73861903bc064f83.__obf_8d733be0011f8684()
		if __obf_c3746095cbc77276 != nil {
			if __obf_6ff5eadfa557e5dd {
				break
			}
			continue
		}
		__obf_2b37cb05c9916766.
			Add(1)
		go func(__obf_fd818d512f67597a int) {
			defer __obf_2b37cb05c9916766.Done()
			if __obf_c3746095cbc77276 := __obf_73861903bc064f83.__obf_dd7ccee14355ced4(__obf_fd818d512f67597a); __obf_c3746095cbc77276 != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_c3746095cbc77276.Error())
				if __obf_c3746095cbc77276 := __obf_73861903bc064f83.__obf_6315c945e3a63686(__obf_fd818d512f67597a); __obf_c3746095cbc77276 != nil {
					fmt.Println(__obf_c3746095cbc77276.Error())
				}
			}
		}(__obf_8359bb0a6d11dd9a)
	}
	__obf_2b37cb05c9916766.
		Wait()
	if __obf_c3746095cbc77276 := __obf_73861903bc064f83.__obf_0a8ab6aecda9e2f7(); __obf_c3746095cbc77276 != nil {
		return __obf_c3746095cbc77276
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_73861903bc064f83 *Downloader) __obf_dd7ccee14355ced4(__obf_1d028addd8c6408e int) error {
	__obf_8dfd8db98341fadd := __obf_8dfd8db98341fadd(__obf_1d028addd8c6408e)
	__obf_d44ae8bff844914a := __obf_73861903bc064f83.__obf_2dcb27a19875b1b3(__obf_1d028addd8c6408e)
	__obf_bbde3be857e9fe98, __obf_410b1db13ff670f6 := tool.Get(__obf_d44ae8bff844914a)
	if __obf_410b1db13ff670f6 != nil {
		return fmt.Errorf("request %s, %s", __obf_d44ae8bff844914a, __obf_410b1db13ff670f6.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_bbde3be857e9fe98.Close()
	__obf_da5f7b49cfcb3b0e := filepath.Join(__obf_73861903bc064f83.__obf_40d94e32155209e0, __obf_8dfd8db98341fadd)
	__obf_9e697a8fddc13b6f := __obf_da5f7b49cfcb3b0e + __obf_14fb6f7b461646ca
	__obf_79c36d25f225f02a, __obf_c3746095cbc77276 := os.Create(__obf_9e697a8fddc13b6f)
	if __obf_c3746095cbc77276 != nil {
		return fmt.Errorf("create file: %s, %s", __obf_8dfd8db98341fadd, __obf_c3746095cbc77276.Error())
	}
	__obf_b1079be0a1721c29, __obf_c3746095cbc77276 := io.ReadAll(__obf_bbde3be857e9fe98)
	if __obf_c3746095cbc77276 != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_d44ae8bff844914a, __obf_c3746095cbc77276.Error())
	}
	__obf_9d6240cf1635ca5e := __obf_73861903bc064f83.__obf_02b78213e8219b62.M3u8.Segments[__obf_1d028addd8c6408e]
	if __obf_9d6240cf1635ca5e == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_1d028addd8c6408e)
	}
	__obf_cb50115d219fd014, __obf_743eff2093de61dc := __obf_73861903bc064f83.__obf_02b78213e8219b62.Keys[__obf_9d6240cf1635ca5e.KeyIndex]
	if __obf_743eff2093de61dc && __obf_cb50115d219fd014 != "" {
		__obf_b1079be0a1721c29, __obf_c3746095cbc77276 = tool.AES128Decrypt(__obf_b1079be0a1721c29, []byte(__obf_cb50115d219fd014),
			[]byte(__obf_73861903bc064f83.__obf_02b78213e8219b62.M3u8.Keys[__obf_9d6240cf1635ca5e.KeyIndex].IV))
		if __obf_c3746095cbc77276 != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_d44ae8bff844914a, __obf_c3746095cbc77276.Error())
		}
	}
	__obf_ee37e9fa938219c9 := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_ddcc0d5c5fbb10f2 := //0x47
		len(__obf_b1079be0a1721c29)
	for __obf_f0ee9690de857e2f := 0; __obf_f0ee9690de857e2f < __obf_ddcc0d5c5fbb10f2; __obf_f0ee9690de857e2f++ {
		if __obf_b1079be0a1721c29[__obf_f0ee9690de857e2f] == __obf_ee37e9fa938219c9 {
			__obf_b1079be0a1721c29 = __obf_b1079be0a1721c29[__obf_f0ee9690de857e2f:]
			break
		}
	}
	__obf_385a02d97affe848 := bufio.NewWriter(__obf_79c36d25f225f02a)
	if _, __obf_c3746095cbc77276 := __obf_385a02d97affe848.Write(__obf_b1079be0a1721c29); __obf_c3746095cbc77276 != nil {
		return fmt.Errorf("write to %s: %s", __obf_9e697a8fddc13b6f, __obf_c3746095cbc77276.Error())
	}
	// Release file resource to rename file
	_ = __obf_79c36d25f225f02a.Close()
	if __obf_c3746095cbc77276 = os.Rename(__obf_9e697a8fddc13b6f, __obf_da5f7b49cfcb3b0e); __obf_c3746095cbc77276 != nil {
		return __obf_c3746095cbc77276
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_73861903bc064f83.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_4130e670c70f620a, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_73861903bc064f83.__obf_4130e670c70f620a)/float32(__obf_73861903bc064f83.__obf_774318f00ce2acad)*100, __obf_d44ae8bff844914a)
	return nil
}

// next Obtains a next segment index
func (__obf_73861903bc064f83 *Downloader) __obf_8d733be0011f8684() (__obf_1d028addd8c6408e int, __obf_6ff5eadfa557e5dd bool, __obf_c3746095cbc77276 error) {
	__obf_73861903bc064f83.__obf_7328558709def097.
		Lock()
	defer __obf_73861903bc064f83.__obf_7328558709def097.Unlock()
	if len(__obf_73861903bc064f83.__obf_443bdea9870aa393) == 0 {
		__obf_c3746095cbc77276 = fmt.Errorf("queue empty")
		if __obf_73861903bc064f83.__obf_4130e670c70f620a == int32(__obf_73861903bc064f83.__obf_774318f00ce2acad) {
			__obf_6ff5eadfa557e5dd = true
			return
		}
		__obf_6ff5eadfa557e5dd = // Some segment indexes are still running.
			false
		return
	}
	__obf_1d028addd8c6408e = __obf_73861903bc064f83.__obf_443bdea9870aa393[0]
	__obf_73861903bc064f83.__obf_443bdea9870aa393 = __obf_73861903bc064f83.__obf_443bdea9870aa393[1:]
	return
}

// back For retry while download failed
func (__obf_73861903bc064f83 *Downloader) __obf_6315c945e3a63686(__obf_1d028addd8c6408e int) error {
	__obf_73861903bc064f83.__obf_7328558709def097.
		Lock()
	defer __obf_73861903bc064f83.__obf_7328558709def097.Unlock()
	if __obf_9d6240cf1635ca5e := __obf_73861903bc064f83.__obf_02b78213e8219b62.M3u8.Segments[__obf_1d028addd8c6408e]; __obf_9d6240cf1635ca5e == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_1d028addd8c6408e)
	}
	__obf_73861903bc064f83.__obf_443bdea9870aa393 = append(__obf_73861903bc064f83.__obf_443bdea9870aa393,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_1d028addd8c6408e)
	return nil
}

func (__obf_73861903bc064f83 *Downloader) __obf_0a8ab6aecda9e2f7() error {
	__obf_4713e513031d785d := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_fd818d512f67597a := 0; __obf_fd818d512f67597a < __obf_73861903bc064f83.__obf_774318f00ce2acad; __obf_fd818d512f67597a++ {
		__obf_8dfd8db98341fadd := __obf_8dfd8db98341fadd(__obf_fd818d512f67597a)
		__obf_79c36d25f225f02a := filepath.Join(__obf_73861903bc064f83.__obf_40d94e32155209e0, __obf_8dfd8db98341fadd)
		if _, __obf_c3746095cbc77276 := os.Stat(__obf_79c36d25f225f02a); __obf_c3746095cbc77276 != nil {
			__obf_4713e513031d785d++
		}
	}
	if __obf_4713e513031d785d > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_4713e513031d785d)
	}
	__obf_f32ba27c948c9455 := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_73861903bc064f83.__obf_8668dd1d7cbf68d4, __obf_0a27001219a6287c)
	__obf_f357eb7c083fa4ea, __obf_c3746095cbc77276 := os.Create(__obf_f32ba27c948c9455)
	if __obf_c3746095cbc77276 != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_c3746095cbc77276.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_f357eb7c083fa4ea.Close()
	__obf_daa0704b419e8b20 := bufio.NewWriter(__obf_f357eb7c083fa4ea)
	__obf_068c735f3b6474c5 := 0
	for __obf_1d028addd8c6408e := 0; __obf_1d028addd8c6408e < __obf_73861903bc064f83.__obf_774318f00ce2acad; __obf_1d028addd8c6408e++ {
		__obf_8dfd8db98341fadd := __obf_8dfd8db98341fadd(__obf_1d028addd8c6408e)
		__obf_b1079be0a1721c29, __obf_c3746095cbc77276 := os.ReadFile(filepath.Join(__obf_73861903bc064f83.__obf_40d94e32155209e0, __obf_8dfd8db98341fadd))
		_, __obf_c3746095cbc77276 = __obf_daa0704b419e8b20.Write(__obf_b1079be0a1721c29)
		if __obf_c3746095cbc77276 != nil {
			continue
		}
		__obf_068c735f3b6474c5++
		tool.DrawProgressBar("merge",
			float32(__obf_068c735f3b6474c5)/float32(__obf_73861903bc064f83.__obf_774318f00ce2acad), __obf_74e246ed5dfcfa3b)
	}
	_ = __obf_daa0704b419e8b20.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_73861903bc064f83.__obf_40d94e32155209e0)

	if __obf_068c735f3b6474c5 != __obf_73861903bc064f83.__obf_774318f00ce2acad {
		fmt.Printf("[warning] \n%d files merge failed", __obf_73861903bc064f83.__obf_774318f00ce2acad-__obf_068c735f3b6474c5)
	}

	fmt.Printf("\n[output] %s\n", __obf_f32ba27c948c9455)

	return nil
}

func (__obf_73861903bc064f83 *Downloader) __obf_2dcb27a19875b1b3(__obf_1d028addd8c6408e int) string {
	__obf_dd29599e1291ed63 := __obf_73861903bc064f83.__obf_02b78213e8219b62.M3u8.Segments[__obf_1d028addd8c6408e]
	return tool.ResolveURL(__obf_73861903bc064f83.__obf_02b78213e8219b62.URL, __obf_dd29599e1291ed63.URI)
}

func __obf_8dfd8db98341fadd(__obf_5fb5eeb3b64381cb int) string {
	return strconv.Itoa(__obf_5fb5eeb3b64381cb) + __obf_412f29766929b45a
}

func __obf_34bcc24b4504f48e(len int) []int {
	__obf_9817e0163cbaa0f3 := make([]int, 0)
	for __obf_c4564db7753c2c1f := 0; __obf_c4564db7753c2c1f < len; __obf_c4564db7753c2c1f++ {
		__obf_9817e0163cbaa0f3 = append(__obf_9817e0163cbaa0f3, __obf_c4564db7753c2c1f)
	}
	return __obf_9817e0163cbaa0f3
}
