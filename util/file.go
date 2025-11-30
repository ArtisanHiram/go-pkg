package __obf_e13b701bec415b48

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

func FileGzip(__obf_34a93e1eca3762b2 *os.File) error {
	__obf_9dbfc9492d2c6232 := __obf_34a93e1eca3762b2.Name() + ".gz"
	__obf_e79f64274a5bc361, __obf_b93c04d14e5c503f := os.Create(__obf_9dbfc9492d2c6232)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_b93c04d14e5c503f)
	}
	defer __obf_e79f64274a5bc361.Close()
	__obf_c838662fbf35d1a0 := gzip.NewWriter(__obf_e79f64274a5bc361)
	defer func() {
		__obf_c838662fbf35d1a0.
			Close()
		if __obf_b93c04d14e5c503f != nil {
			_ = os.Remove(__obf_9dbfc9492d2c6232)
			__obf_b93c04d14e5c503f = fmt.Errorf("failed to compress file, error(%v)", __obf_b93c04d14e5c503f)
		}
	}()

	_, __obf_b93c04d14e5c503f = io.Copy(__obf_c838662fbf35d1a0,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_34a93e1eca3762b2)
	return __obf_b93c04d14e5c503f

}

func PathGzip(__obf_b5479ea9362768d4 string) error {
	var __obf_b93c04d14e5c503f error
	__obf_b5479ea9362768d4, __obf_b93c04d14e5c503f = filepath.Abs(__obf_b5479ea9362768d4)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_b93c04d14e5c503f)
	}

	var __obf_34a93e1eca3762b2 *os.File
	__obf_34a93e1eca3762b2, __obf_b93c04d14e5c503f = os.Open(__obf_b5479ea9362768d4)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_b5479ea9362768d4, __obf_b93c04d14e5c503f)
	}
	defer __obf_34a93e1eca3762b2.Close()

	return FileGzip(__obf_34a93e1eca3762b2)
}

func PathCopy(__obf_822f63011c520ac3, __obf_8b35271a9e94c6a7 string) error {
	var __obf_b93c04d14e5c503f error
	__obf_822f63011c520ac3, __obf_b93c04d14e5c503f = filepath.Abs(__obf_822f63011c520ac3)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("srcPath error: %s", __obf_b93c04d14e5c503f)
	}
	__obf_8b35271a9e94c6a7, __obf_b93c04d14e5c503f = filepath.Abs(__obf_8b35271a9e94c6a7)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_b93c04d14e5c503f)
	}

	var __obf_0a5b3786e86302a1 []byte
	__obf_0a5b3786e86302a1, __obf_b93c04d14e5c503f = os.ReadFile(__obf_822f63011c520ac3)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_b93c04d14e5c503f)
	}

	return os.WriteFile(__obf_8b35271a9e94c6a7, __obf_0a5b3786e86302a1, 0644)
}

func FileCopy(__obf_a3d4a3c0ddfb0d83 *os.File, __obf_8b35271a9e94c6a7 string) error {
	var __obf_b93c04d14e5c503f error
	__obf_8b35271a9e94c6a7, __obf_b93c04d14e5c503f = filepath.Abs(__obf_8b35271a9e94c6a7)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_b93c04d14e5c503f)
	}
	__obf_6004f7d37045bf33, __obf_b93c04d14e5c503f := os.Create(__obf_8b35271a9e94c6a7)
	if __obf_b93c04d14e5c503f != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_b93c04d14e5c503f)
	}
	defer __obf_6004f7d37045bf33.Close()

	_, __obf_b93c04d14e5c503f = io.Copy(__obf_6004f7d37045bf33, __obf_a3d4a3c0ddfb0d83)
	return __obf_b93c04d14e5c503f
}

func FileZip(__obf_811f2224307efb3f io.Writer, __obf_695934164af1d8ff uint16, __obf_f7fcca366fccd9d2 ...struct {
	Name string
	Path string
},
) error {
	__obf_c396d7b8f82d8b2f := zip.NewWriter(__obf_811f2224307efb3f)
	defer __obf_c396d7b8f82d8b2f.Close()

	var (
		__obf_b93c04d14e5c503f error
		__obf_d04a0e199cdc03a3 *os.File
		__obf_5a0cea88d46ee29e io.Writer
		__obf_9ea53a08e18f0c98 fs.FileInfo
		__obf_7b2babfa0a9c09b1 *zip.FileHeader
	)
	for _, __obf_99fad6cd23d30931 := range __obf_f7fcca366fccd9d2 {
		__obf_d04a0e199cdc03a3, __obf_b93c04d14e5c503f = os.Open(__obf_99fad6cd23d30931.Path)
		if __obf_b93c04d14e5c503f != nil {
			return __obf_b93c04d14e5c503f
		}
		defer __obf_d04a0e199cdc03a3.Close()
		__obf_9ea53a08e18f0c98, __obf_b93c04d14e5c503f = __obf_d04a0e199cdc03a3.Stat()
		if __obf_b93c04d14e5c503f != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_99fad6cd23d30931.Path, __obf_b93c04d14e5c503f)
		}
		__obf_7b2babfa0a9c09b1, __obf_b93c04d14e5c503f = zip.FileInfoHeader(__obf_9ea53a08e18f0c98)
		if __obf_b93c04d14e5c503f != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_b93c04d14e5c503f)
		}
		__obf_7b2babfa0a9c09b1.
			// 压缩算法
			Method = __obf_695934164af1d8ff
		// header.Name = filepath.Base(v)
		__obf_7b2babfa0a9c09b1.
			Name = __obf_99fad6cd23d30931.Name
		__obf_5a0cea88d46ee29e, __obf_b93c04d14e5c503f = __obf_c396d7b8f82d8b2f.CreateHeader(__obf_7b2babfa0a9c09b1)
		if __obf_b93c04d14e5c503f != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_b93c04d14e5c503f)
		}
		_, __obf_b93c04d14e5c503f = io.Copy(__obf_5a0cea88d46ee29e, __obf_d04a0e199cdc03a3)
		if __obf_b93c04d14e5c503f != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_b93c04d14e5c503f)
		}
	}
	return nil
}

func WriteCSV(__obf_a5ab9a007e4ee13b string, __obf_1e49a709823fcdda [][]string) (__obf_b93c04d14e5c503f error) {
	var __obf_2f1d0e2db86d0a35 *os.File
	__obf_2f1d0e2db86d0a35, __obf_b93c04d14e5c503f = os.OpenFile(__obf_a5ab9a007e4ee13b+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	defer __obf_2f1d0e2db86d0a35.Close()
	_, __obf_b93c04d14e5c503f = __obf_2f1d0e2db86d0a35.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_b93c04d14e5c503f != nil {
		return
	}
	__obf_c838662fbf35d1a0 := csv.NewWriter(__obf_2f1d0e2db86d0a35)
	__obf_b93c04d14e5c503f = __obf_c838662fbf35d1a0.WriteAll(__obf_1e49a709823fcdda)
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	__obf_c838662fbf35d1a0.
		Flush()
	return
}

func ReadCSV(path string) (__obf_1e49a709823fcdda [][]string, __obf_b93c04d14e5c503f error) {
	var fs *os.File
	fs, __obf_b93c04d14e5c503f = os.Open(path)
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	defer fs.Close()
	__obf_af65a69017eda959 := csv.NewReader(fs)
	__obf_af65a69017eda959.
		TrimLeadingSpace = true
	__obf_1e49a709823fcdda, __obf_b93c04d14e5c503f = __obf_af65a69017eda959.ReadAll()
	if __obf_b93c04d14e5c503f != nil {
		return
	}
	return
}

func __obf_278483f25e73bbaf(__obf_5c9c2b1eb9a7f576 string) bool {
	_, __obf_b93c04d14e5c503f := os.Stat(__obf_5c9c2b1eb9a7f576)
	return !os.IsNotExist(__obf_b93c04d14e5c503f)
}

func GetFileExt(__obf_aff8991e6054649d string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_aff8991e6054649d), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_da8461c67018e131 string) error {
	if !__obf_278483f25e73bbaf(__obf_da8461c67018e131) {
		return os.MkdirAll(__obf_da8461c67018e131, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_a5ab9a007e4ee13b, __obf_991c6908209e9369 string) *os.File {
	if !__obf_278483f25e73bbaf(__obf_a5ab9a007e4ee13b) {
		__obf_af65a69017eda959 := []rune(__obf_a5ab9a007e4ee13b)
		if __obf_550d8612d04f5667 := strings.LastIndex(__obf_a5ab9a007e4ee13b, __obf_991c6908209e9369); __obf_550d8612d04f5667 != -1 {
			if __obf_b93c04d14e5c503f := os.MkdirAll(string(__obf_af65a69017eda959[:__obf_550d8612d04f5667]), os.ModePerm); __obf_b93c04d14e5c503f != nil {
				log.Fatalln(`Mkdir`, __obf_b93c04d14e5c503f)
			}
		}
	}
	__obf_34a93e1eca3762b2, __obf_b93c04d14e5c503f := os.Create(__obf_a5ab9a007e4ee13b) //创建文件
	if __obf_b93c04d14e5c503f != nil {
		log.Fatalln(`file create`, __obf_b93c04d14e5c503f)
	}
	return __obf_34a93e1eca3762b2
}

func OpenFile(__obf_a5ab9a007e4ee13b, __obf_991c6908209e9369 string) *os.File {
	if !__obf_278483f25e73bbaf(__obf_a5ab9a007e4ee13b) {
		__obf_af65a69017eda959 := []rune(__obf_a5ab9a007e4ee13b)
		if __obf_550d8612d04f5667 := strings.LastIndex(__obf_a5ab9a007e4ee13b, __obf_991c6908209e9369); __obf_550d8612d04f5667 != -1 {
			if __obf_b93c04d14e5c503f := os.MkdirAll(string(__obf_af65a69017eda959[:__obf_550d8612d04f5667]), os.ModePerm); __obf_b93c04d14e5c503f != nil {
				log.Fatalln(`Mkdir`, __obf_b93c04d14e5c503f)
			}
		}
	}
	__obf_34a93e1eca3762b2, __obf_b93c04d14e5c503f := os.OpenFile(__obf_a5ab9a007e4ee13b, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_b93c04d14e5c503f != nil {
		log.Fatalln(`file create`, __obf_b93c04d14e5c503f)
	}
	return __obf_34a93e1eca3762b2
}
