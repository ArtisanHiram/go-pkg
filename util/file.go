package __obf_426da37e60cac670

import (
	"archive/zip"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func FileGzip(__obf_ea8da91a172a8512 *os.File) error {
	__obf_e7bf3b74b0653d37 := __obf_ea8da91a172a8512.Name() + ".gz"
	__obf_6317d2183b2a1e70, __obf_74916b80241ef1ff := os.Create(__obf_e7bf3b74b0653d37)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_74916b80241ef1ff)
	}
	defer __obf_6317d2183b2a1e70.Close()
	__obf_c7f97b6f9ca45d13 := gzip.NewWriter(__obf_6317d2183b2a1e70)
	defer func() {
		__obf_c7f97b6f9ca45d13.
			Close()
		if __obf_74916b80241ef1ff != nil {
			_ = os.Remove(__obf_e7bf3b74b0653d37)
			__obf_74916b80241ef1ff = fmt.Errorf("failed to compress file, error(%v)", __obf_74916b80241ef1ff)
		}
	}()

	_, __obf_74916b80241ef1ff = io.Copy(__obf_c7f97b6f9ca45d13,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_ea8da91a172a8512)
	return __obf_74916b80241ef1ff

}

func PathGzip(__obf_0caa525ab196f9a7 string) error {
	var __obf_74916b80241ef1ff error
	__obf_0caa525ab196f9a7, __obf_74916b80241ef1ff = filepath.Abs(__obf_0caa525ab196f9a7)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_74916b80241ef1ff)
	}

	var __obf_ea8da91a172a8512 *os.File
	__obf_ea8da91a172a8512, __obf_74916b80241ef1ff = os.Open(__obf_0caa525ab196f9a7)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_0caa525ab196f9a7, __obf_74916b80241ef1ff)
	}
	defer __obf_ea8da91a172a8512.Close()

	return FileGzip(__obf_ea8da91a172a8512)
}

func PathCopy(__obf_b9553aad25f4008f, __obf_b7333aa35a189ddf string) error {
	var __obf_74916b80241ef1ff error
	__obf_b9553aad25f4008f, __obf_74916b80241ef1ff = filepath.Abs(__obf_b9553aad25f4008f)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("srcPath error: %s", __obf_74916b80241ef1ff)
	}
	__obf_b7333aa35a189ddf, __obf_74916b80241ef1ff = filepath.Abs(__obf_b7333aa35a189ddf)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("dstPath error: %s", __obf_74916b80241ef1ff)
	}

	var __obf_f524680743378a06 []byte
	__obf_f524680743378a06, __obf_74916b80241ef1ff = os.ReadFile(__obf_b9553aad25f4008f)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_74916b80241ef1ff)
	}

	return os.WriteFile(__obf_b7333aa35a189ddf, __obf_f524680743378a06, 0644)
}

func FileCopy(__obf_9558c4b7c50309ef *os.File, __obf_b7333aa35a189ddf string) error {
	var __obf_74916b80241ef1ff error
	__obf_b7333aa35a189ddf, __obf_74916b80241ef1ff = filepath.Abs(__obf_b7333aa35a189ddf)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("dstPath error: %s", __obf_74916b80241ef1ff)
	}
	__obf_eac92198fd2cd08f, __obf_74916b80241ef1ff := os.Create(__obf_b7333aa35a189ddf)
	if __obf_74916b80241ef1ff != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_74916b80241ef1ff)
	}
	defer __obf_eac92198fd2cd08f.Close()

	_, __obf_74916b80241ef1ff = io.Copy(__obf_eac92198fd2cd08f, __obf_9558c4b7c50309ef)
	return __obf_74916b80241ef1ff
}

func FileZip(__obf_e38180aaf342b779 io.Writer, __obf_f45530bfb696670d uint16, __obf_18b7ce7a2be7b8cb ...struct {
	Name string
	Path string
},
) error {
	__obf_20ee94509a833934 := zip.NewWriter(__obf_e38180aaf342b779)
	defer __obf_20ee94509a833934.Close()

	var (
		__obf_74916b80241ef1ff error
		__obf_8e41d8667e33ef22 *os.File
		__obf_1b9b80411c9452b5 io.Writer
		__obf_90cf9b241b498b5d fs.FileInfo
		__obf_58ad5e2feee96de5 *zip.FileHeader
	)
	for _, __obf_c76af479674ee3d0 := range __obf_18b7ce7a2be7b8cb {
		__obf_8e41d8667e33ef22, __obf_74916b80241ef1ff = os.Open(__obf_c76af479674ee3d0.Path)
		if __obf_74916b80241ef1ff != nil {
			return __obf_74916b80241ef1ff
		}
		defer __obf_8e41d8667e33ef22.Close()
		__obf_90cf9b241b498b5d, __obf_74916b80241ef1ff = __obf_8e41d8667e33ef22.Stat()
		if __obf_74916b80241ef1ff != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_c76af479674ee3d0.Path, __obf_74916b80241ef1ff)
		}
		__obf_58ad5e2feee96de5, __obf_74916b80241ef1ff = zip.FileInfoHeader(__obf_90cf9b241b498b5d)
		if __obf_74916b80241ef1ff != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_74916b80241ef1ff)
		}
		__obf_58ad5e2feee96de5.
			// 压缩算法
			Method = __obf_f45530bfb696670d
		// header.Name = filepath.Base(v)
		__obf_58ad5e2feee96de5.
			Name = __obf_c76af479674ee3d0.Name
		__obf_1b9b80411c9452b5, __obf_74916b80241ef1ff = __obf_20ee94509a833934.CreateHeader(__obf_58ad5e2feee96de5)
		if __obf_74916b80241ef1ff != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_74916b80241ef1ff)
		}
		_, __obf_74916b80241ef1ff = io.Copy(__obf_1b9b80411c9452b5, __obf_8e41d8667e33ef22)
		if __obf_74916b80241ef1ff != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_74916b80241ef1ff)
		}
	}
	return nil
}

func WriteCSV(__obf_25ec75064e81d6e1 string, __obf_96894de1e88392eb [][]string) (__obf_74916b80241ef1ff error) {
	var __obf_b609bf81c0e36d47 *os.File
	__obf_b609bf81c0e36d47, __obf_74916b80241ef1ff = os.OpenFile(__obf_25ec75064e81d6e1+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_74916b80241ef1ff != nil {
		return
	}
	defer __obf_b609bf81c0e36d47.Close()
	_, __obf_74916b80241ef1ff = __obf_b609bf81c0e36d47.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_74916b80241ef1ff != nil {
		return
	}
	__obf_c7f97b6f9ca45d13 := csv.NewWriter(__obf_b609bf81c0e36d47)
	__obf_74916b80241ef1ff = __obf_c7f97b6f9ca45d13.WriteAll(__obf_96894de1e88392eb)
	if __obf_74916b80241ef1ff != nil {
		return
	}
	__obf_c7f97b6f9ca45d13.
		Flush()
	return
}

func ReadCSV(path string) (__obf_96894de1e88392eb [][]string, __obf_74916b80241ef1ff error) {
	var fs *os.File
	fs, __obf_74916b80241ef1ff = os.Open(path)
	if __obf_74916b80241ef1ff != nil {
		return
	}
	defer fs.Close()
	__obf_3be342a2e8ac84dc := csv.NewReader(fs)
	__obf_3be342a2e8ac84dc.
		TrimLeadingSpace = true
	__obf_96894de1e88392eb, __obf_74916b80241ef1ff = __obf_3be342a2e8ac84dc.ReadAll()
	if __obf_74916b80241ef1ff != nil {
		return
	}
	return
}

func __obf_16e88952e16d55a2(__obf_04310e04e67ea058 string) bool {
	_, __obf_74916b80241ef1ff := os.Stat(__obf_04310e04e67ea058)
	return !os.IsNotExist(__obf_74916b80241ef1ff)
}

func GetFileExt(__obf_e454ce3773624d31 string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_e454ce3773624d31), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_6a1bd54a37d9409b string) error {
	if !__obf_16e88952e16d55a2(__obf_6a1bd54a37d9409b) {
		return os.MkdirAll(__obf_6a1bd54a37d9409b, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_25ec75064e81d6e1, __obf_f442da269b639645 string) *os.File {
	if !__obf_16e88952e16d55a2(__obf_25ec75064e81d6e1) {
		__obf_3be342a2e8ac84dc := []rune(__obf_25ec75064e81d6e1)
		if __obf_f7eefc38acf35467 := strings.LastIndex(__obf_25ec75064e81d6e1, __obf_f442da269b639645); __obf_f7eefc38acf35467 != -1 {
			if __obf_74916b80241ef1ff := os.MkdirAll(string(__obf_3be342a2e8ac84dc[:__obf_f7eefc38acf35467]), os.ModePerm); __obf_74916b80241ef1ff != nil {
				log.Fatalln(`Mkdir`, __obf_74916b80241ef1ff)
			}
		}
	}
	__obf_ea8da91a172a8512, __obf_74916b80241ef1ff := os.Create(__obf_25ec75064e81d6e1) //创建文件
	if __obf_74916b80241ef1ff != nil {
		log.Fatalln(`file create`, __obf_74916b80241ef1ff)
	}
	return __obf_ea8da91a172a8512
}

func OpenFile(__obf_25ec75064e81d6e1, __obf_f442da269b639645 string) *os.File {
	if !__obf_16e88952e16d55a2(__obf_25ec75064e81d6e1) {
		__obf_3be342a2e8ac84dc := []rune(__obf_25ec75064e81d6e1)
		if __obf_f7eefc38acf35467 := strings.LastIndex(__obf_25ec75064e81d6e1, __obf_f442da269b639645); __obf_f7eefc38acf35467 != -1 {
			if __obf_74916b80241ef1ff := os.MkdirAll(string(__obf_3be342a2e8ac84dc[:__obf_f7eefc38acf35467]), os.ModePerm); __obf_74916b80241ef1ff != nil {
				log.Fatalln(`Mkdir`, __obf_74916b80241ef1ff)
			}
		}
	}
	__obf_ea8da91a172a8512, __obf_74916b80241ef1ff := os.OpenFile(__obf_25ec75064e81d6e1, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_74916b80241ef1ff != nil {
		log.Fatalln(`file create`, __obf_74916b80241ef1ff)
	}
	return __obf_ea8da91a172a8512
}
