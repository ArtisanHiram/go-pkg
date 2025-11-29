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
	__obf_44cad197299e83ac = ".ts"
	__obf_54298c0c5b29da87 = "ts"
	__obf_4c3f62f65f8ffbca = "main.ts"
	__obf_219008868b79e7f7 = "_tmp"
	__obf_ef3f12dbd72a22bc = 40
)

// Downloader A task definition to download
type Downloader struct {
	__obf_ddcbc9ff400647ac sync.Mutex
	__obf_aafaa33efe1d7551 []int
	__obf_954290ac54c13ef9 string
	__obf_405c5e43008a542f string
	__obf_66d7a23d9404d1cf int32
	__obf_fdd5198ec3ef35d6 int
	__obf_d2bb40c8dd45eb71 *parse.Result
}

// NewTask Returns a Task instance
// @param output string  file output folder path
// @param url    string  m3u8 http url
func NewTask(__obf_719136bb3474572d string, __obf_3f814554f334ee3c string) (*Downloader, error) {
	__obf_d2bb40c8dd45eb71, __obf_8e97e8da92c78189 := parse.ParseFromNet(__obf_3f814554f334ee3c)
	if __obf_8e97e8da92c78189 != nil {
		return nil, __obf_8e97e8da92c78189
	}
	var __obf_954290ac54c13ef9 string
	// If no output folder specified, use current directory
	if __obf_719136bb3474572d == "" {
		__obf_ae280451fc8d8605, __obf_8e97e8da92c78189 := tool.CurrentDir()
		if __obf_8e97e8da92c78189 != nil {
			return nil, __obf_8e97e8da92c78189
		}
		__obf_954290ac54c13ef9 = filepath.Join(__obf_ae280451fc8d8605, __obf_719136bb3474572d)
	} else {
		__obf_954290ac54c13ef9 = __obf_719136bb3474572d
	}
	if __obf_8e97e8da92c78189 := os.MkdirAll(__obf_954290ac54c13ef9, os.ModePerm); __obf_8e97e8da92c78189 != nil {
		return nil, fmt.Errorf("create storage folder failed: %s", __obf_8e97e8da92c78189.Error())
	}
	__obf_405c5e43008a542f := filepath.Join(__obf_954290ac54c13ef9, __obf_54298c0c5b29da87)
	if __obf_8e97e8da92c78189 := os.MkdirAll(__obf_405c5e43008a542f, os.ModePerm); __obf_8e97e8da92c78189 != nil {
		return nil, fmt.Errorf("create ts folder '[%s]' failed: %s", __obf_405c5e43008a542f, __obf_8e97e8da92c78189.Error())
	}
	__obf_5f662c758ab042d5 := &Downloader{__obf_954290ac54c13ef9: __obf_954290ac54c13ef9, __obf_405c5e43008a542f: __obf_405c5e43008a542f, __obf_d2bb40c8dd45eb71: __obf_d2bb40c8dd45eb71}
	__obf_5f662c758ab042d5.
		// segment length
		__obf_fdd5198ec3ef35d6 = len(__obf_d2bb40c8dd45eb71.M3u8.Segments)
	__obf_5f662c758ab042d5.__obf_aafaa33efe1d7551 = __obf_eb5df934143f4773(__obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6)
	return __obf_5f662c758ab042d5, nil
}

// Start Current downloader runs
// Use goroutine to download parallel
func (__obf_5f662c758ab042d5 *Downloader) Start(__obf_e292df9c97aa521c int) error {
	var __obf_03df6f415d7d45c8 sync.WaitGroup
	for {
		__obf_d94da5a6b1d09c21, __obf_ecfda28f4344b612, __obf_8e97e8da92c78189 := __obf_5f662c758ab042d5.__obf_b8da08940016a6fb()
		if __obf_8e97e8da92c78189 != nil {
			if __obf_ecfda28f4344b612 {
				break
			}
			continue
		}
		__obf_03df6f415d7d45c8.
			Add(1)
		go func(__obf_b9ff44be4f8fd04b int) {
			defer __obf_03df6f415d7d45c8.Done()
			if __obf_8e97e8da92c78189 := __obf_5f662c758ab042d5.__obf_0976976d45ac764b(__obf_b9ff44be4f8fd04b); __obf_8e97e8da92c78189 != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", __obf_8e97e8da92c78189.Error())
				if __obf_8e97e8da92c78189 := __obf_5f662c758ab042d5.__obf_96a20db624a4c98a(__obf_b9ff44be4f8fd04b); __obf_8e97e8da92c78189 != nil {
					fmt.Println(__obf_8e97e8da92c78189.Error())
				}
			}
		}(__obf_d94da5a6b1d09c21)
	}
	__obf_03df6f415d7d45c8.
		Wait()
	if __obf_8e97e8da92c78189 := __obf_5f662c758ab042d5.__obf_f4236505c1471ac2(); __obf_8e97e8da92c78189 != nil {
		return __obf_8e97e8da92c78189
	}
	return nil
}

// download Current downloader to down specific index segment
func (__obf_5f662c758ab042d5 *Downloader) __obf_0976976d45ac764b(__obf_34dd02db8d7c15ae int) error {
	__obf_b02ec68d94caa84d := __obf_b02ec68d94caa84d(__obf_34dd02db8d7c15ae)
	__obf_f6ec21c760320ac5 := __obf_5f662c758ab042d5.__obf_d594178cfceb3d59(__obf_34dd02db8d7c15ae)
	__obf_c6384c323d207ea5, __obf_a9d679f8a5363602 := tool.Get(__obf_f6ec21c760320ac5)
	if __obf_a9d679f8a5363602 != nil {
		return fmt.Errorf("request %s, %s", __obf_f6ec21c760320ac5, __obf_a9d679f8a5363602.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_c6384c323d207ea5.Close()
	__obf_ba5820df02f83130 := filepath.Join(__obf_5f662c758ab042d5.__obf_405c5e43008a542f, __obf_b02ec68d94caa84d)
	__obf_021d27d1e3168540 := __obf_ba5820df02f83130 + __obf_219008868b79e7f7
	__obf_aa71cbbedaf47f66, __obf_8e97e8da92c78189 := os.Create(__obf_021d27d1e3168540)
	if __obf_8e97e8da92c78189 != nil {
		return fmt.Errorf("create file: %s, %s", __obf_b02ec68d94caa84d, __obf_8e97e8da92c78189.Error())
	}
	__obf_396f122c15d73950, __obf_8e97e8da92c78189 := io.ReadAll(__obf_c6384c323d207ea5)
	if __obf_8e97e8da92c78189 != nil {
		return fmt.Errorf("read bytes: %s, %s", __obf_f6ec21c760320ac5, __obf_8e97e8da92c78189.Error())
	}
	__obf_aaf55be1b62c0a0f := __obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.M3u8.Segments[__obf_34dd02db8d7c15ae]
	if __obf_aaf55be1b62c0a0f == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_34dd02db8d7c15ae)
	}
	__obf_cab91e687a21b37b, __obf_7ba8225146dc1d30 := __obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.Keys[__obf_aaf55be1b62c0a0f.KeyIndex]
	if __obf_7ba8225146dc1d30 && __obf_cab91e687a21b37b != "" {
		__obf_396f122c15d73950, __obf_8e97e8da92c78189 = tool.AES128Decrypt(__obf_396f122c15d73950, []byte(__obf_cab91e687a21b37b),
			[]byte(__obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.M3u8.Keys[__obf_aaf55be1b62c0a0f.KeyIndex].IV))
		if __obf_8e97e8da92c78189 != nil {
			return fmt.Errorf("decryt: %s, %s", __obf_f6ec21c760320ac5, __obf_8e97e8da92c78189.Error())
		}
	}
	__obf_97968661b1c73a67 := // https://en.wikipedia.org/wiki/MPEG_transport_stream
		// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
		// Need to remove the bytes before the SyncByte 0x47(71).
		uint8(71)
	__obf_afc4c8bd81f65a9d := //0x47
		len(__obf_396f122c15d73950)
	for __obf_ec73bc34e8661136 := 0; __obf_ec73bc34e8661136 < __obf_afc4c8bd81f65a9d; __obf_ec73bc34e8661136++ {
		if __obf_396f122c15d73950[__obf_ec73bc34e8661136] == __obf_97968661b1c73a67 {
			__obf_396f122c15d73950 = __obf_396f122c15d73950[__obf_ec73bc34e8661136:]
			break
		}
	}
	__obf_c0f14c17cb9146f6 := bufio.NewWriter(__obf_aa71cbbedaf47f66)
	if _, __obf_8e97e8da92c78189 := __obf_c0f14c17cb9146f6.Write(__obf_396f122c15d73950); __obf_8e97e8da92c78189 != nil {
		return fmt.Errorf("write to %s: %s", __obf_021d27d1e3168540, __obf_8e97e8da92c78189.Error())
	}
	// Release file resource to rename file
	_ = __obf_aa71cbbedaf47f66.Close()
	if __obf_8e97e8da92c78189 = os.Rename(__obf_021d27d1e3168540, __obf_ba5820df02f83130); __obf_8e97e8da92c78189 != nil {
		return __obf_8e97e8da92c78189
	}
	// Maybe it will be safer in this way...
	atomic.AddInt32(&__obf_5f662c758ab042d5.
		//tool.DrawProgressBar("Downloading", float32(d.finish)/float32(d.segLen), progressWidth)
		__obf_66d7a23d9404d1cf, 1)

	fmt.Printf("[download %6.2f%%] %s\n", float32(__obf_5f662c758ab042d5.__obf_66d7a23d9404d1cf)/float32(__obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6)*100, __obf_f6ec21c760320ac5)
	return nil
}

// next Obtains a next segment index
func (__obf_5f662c758ab042d5 *Downloader) __obf_b8da08940016a6fb() (__obf_34dd02db8d7c15ae int, __obf_ecfda28f4344b612 bool, __obf_8e97e8da92c78189 error) {
	__obf_5f662c758ab042d5.__obf_ddcbc9ff400647ac.
		Lock()
	defer __obf_5f662c758ab042d5.__obf_ddcbc9ff400647ac.Unlock()
	if len(__obf_5f662c758ab042d5.__obf_aafaa33efe1d7551) == 0 {
		__obf_8e97e8da92c78189 = fmt.Errorf("queue empty")
		if __obf_5f662c758ab042d5.__obf_66d7a23d9404d1cf == int32(__obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6) {
			__obf_ecfda28f4344b612 = true
			return
		}
		__obf_ecfda28f4344b612 = // Some segment indexes are still running.
			false
		return
	}
	__obf_34dd02db8d7c15ae = __obf_5f662c758ab042d5.__obf_aafaa33efe1d7551[0]
	__obf_5f662c758ab042d5.__obf_aafaa33efe1d7551 = __obf_5f662c758ab042d5.__obf_aafaa33efe1d7551[1:]
	return
}

// back For retry while download failed
func (__obf_5f662c758ab042d5 *Downloader) __obf_96a20db624a4c98a(__obf_34dd02db8d7c15ae int) error {
	__obf_5f662c758ab042d5.__obf_ddcbc9ff400647ac.
		Lock()
	defer __obf_5f662c758ab042d5.__obf_ddcbc9ff400647ac.Unlock()
	if __obf_aaf55be1b62c0a0f := __obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.M3u8.Segments[__obf_34dd02db8d7c15ae]; __obf_aaf55be1b62c0a0f == nil {
		return fmt.Errorf("invalid segment index: %d", __obf_34dd02db8d7c15ae)
	}
	__obf_5f662c758ab042d5.__obf_aafaa33efe1d7551 = append(__obf_5f662c758ab042d5.__obf_aafaa33efe1d7551,

		// merge Merge all ts segment after 'download' 'finish'
		__obf_34dd02db8d7c15ae)
	return nil
}

func (__obf_5f662c758ab042d5 *Downloader) __obf_f4236505c1471ac2() error {
	__obf_bf9bfde2894a0230 := // In fact, the number of downloaded segments should be equal to number of m3u8 segments
		0
	for __obf_b9ff44be4f8fd04b := 0; __obf_b9ff44be4f8fd04b < __obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6; __obf_b9ff44be4f8fd04b++ {
		__obf_b02ec68d94caa84d := __obf_b02ec68d94caa84d(__obf_b9ff44be4f8fd04b)
		__obf_aa71cbbedaf47f66 := filepath.Join(__obf_5f662c758ab042d5.__obf_405c5e43008a542f, __obf_b02ec68d94caa84d)
		if _, __obf_8e97e8da92c78189 := os.Stat(__obf_aa71cbbedaf47f66); __obf_8e97e8da92c78189 != nil {
			__obf_bf9bfde2894a0230++
		}
	}
	if __obf_bf9bfde2894a0230 > 0 {
		fmt.Printf("[warning] %d files missing\n", __obf_bf9bfde2894a0230)
	}
	__obf_fb519103bb392d18 := // Create a TS file for merging, all segment files will be written to this file.
		filepath.Join(__obf_5f662c758ab042d5.__obf_954290ac54c13ef9, __obf_4c3f62f65f8ffbca)
	__obf_f365fd8140379bd0, __obf_8e97e8da92c78189 := os.Create(__obf_fb519103bb392d18)
	if __obf_8e97e8da92c78189 != nil {
		return fmt.Errorf("create main TS file failed：%s", __obf_8e97e8da92c78189.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer __obf_f365fd8140379bd0.Close()
	__obf_1cd1796ffe93b6bc := bufio.NewWriter(__obf_f365fd8140379bd0)
	__obf_b0633e15d3031e5a := 0
	for __obf_34dd02db8d7c15ae := 0; __obf_34dd02db8d7c15ae < __obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6; __obf_34dd02db8d7c15ae++ {
		__obf_b02ec68d94caa84d := __obf_b02ec68d94caa84d(__obf_34dd02db8d7c15ae)
		__obf_396f122c15d73950, __obf_8e97e8da92c78189 := os.ReadFile(filepath.Join(__obf_5f662c758ab042d5.__obf_405c5e43008a542f, __obf_b02ec68d94caa84d))
		_, __obf_8e97e8da92c78189 = __obf_1cd1796ffe93b6bc.Write(__obf_396f122c15d73950)
		if __obf_8e97e8da92c78189 != nil {
			continue
		}
		__obf_b0633e15d3031e5a++
		tool.DrawProgressBar("merge",
			float32(__obf_b0633e15d3031e5a)/float32(__obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6), __obf_ef3f12dbd72a22bc)
	}
	_ = __obf_1cd1796ffe93b6bc.Flush()
	// Remove `ts` folder
	_ = os.RemoveAll(__obf_5f662c758ab042d5.__obf_405c5e43008a542f)

	if __obf_b0633e15d3031e5a != __obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6 {
		fmt.Printf("[warning] \n%d files merge failed", __obf_5f662c758ab042d5.__obf_fdd5198ec3ef35d6-__obf_b0633e15d3031e5a)
	}

	fmt.Printf("\n[output] %s\n", __obf_fb519103bb392d18)

	return nil
}

func (__obf_5f662c758ab042d5 *Downloader) __obf_d594178cfceb3d59(__obf_34dd02db8d7c15ae int) string {
	__obf_23dc070325fedcda := __obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.M3u8.Segments[__obf_34dd02db8d7c15ae]
	return tool.ResolveURL(__obf_5f662c758ab042d5.__obf_d2bb40c8dd45eb71.URL, __obf_23dc070325fedcda.URI)
}

func __obf_b02ec68d94caa84d(__obf_818e5c08ab6b965d int) string {
	return strconv.Itoa(__obf_818e5c08ab6b965d) + __obf_44cad197299e83ac
}

func __obf_eb5df934143f4773(len int) []int {
	__obf_f542dea8c1c5845c := make([]int, 0)
	for __obf_7410eb62a97a6ed9 := 0; __obf_7410eb62a97a6ed9 < len; __obf_7410eb62a97a6ed9++ {
		__obf_f542dea8c1c5845c = append(__obf_f542dea8c1c5845c, __obf_7410eb62a97a6ed9)
	}
	return __obf_f542dea8c1c5845c
}
