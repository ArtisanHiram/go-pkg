package __obf_85fcc7bd65d30170

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

func FileGzip(__obf_b3f9ba974f202dc7 *os.File) error {
	__obf_b16422adcfb7544d := __obf_b3f9ba974f202dc7.Name() + ".gz"
	__obf_6a8351e90cb58d23, __obf_80a78f1bfaca43d3 := os.Create(__obf_b16422adcfb7544d)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_80a78f1bfaca43d3)
	}
	defer __obf_6a8351e90cb58d23.Close()

	__obf_2a9d3977ed95b56f := gzip.NewWriter(__obf_6a8351e90cb58d23)
	defer func() {
		__obf_2a9d3977ed95b56f.Close()
		if __obf_80a78f1bfaca43d3 != nil {
			_ = os.Remove(__obf_b16422adcfb7544d)
			__obf_80a78f1bfaca43d3 = fmt.Errorf("failed to compress file, error(%v)", __obf_80a78f1bfaca43d3)
		}
	}()

	_, __obf_80a78f1bfaca43d3 = io.Copy(__obf_2a9d3977ed95b56f, __obf_b3f9ba974f202dc7)
	return __obf_80a78f1bfaca43d3
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_456a20d6e9422e15 string) error {
	var __obf_80a78f1bfaca43d3 error
	__obf_456a20d6e9422e15, __obf_80a78f1bfaca43d3 = filepath.Abs(__obf_456a20d6e9422e15)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_80a78f1bfaca43d3)
	}

	var __obf_b3f9ba974f202dc7 *os.File
	__obf_b3f9ba974f202dc7, __obf_80a78f1bfaca43d3 = os.Open(__obf_456a20d6e9422e15)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_456a20d6e9422e15, __obf_80a78f1bfaca43d3)
	}
	defer __obf_b3f9ba974f202dc7.Close()

	return FileGzip(__obf_b3f9ba974f202dc7)
}

func PathCopy(__obf_0243249d763c02e5, __obf_7643f851b965922b string) error {
	var __obf_80a78f1bfaca43d3 error
	__obf_0243249d763c02e5, __obf_80a78f1bfaca43d3 = filepath.Abs(__obf_0243249d763c02e5)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_80a78f1bfaca43d3)
	}

	__obf_7643f851b965922b, __obf_80a78f1bfaca43d3 = filepath.Abs(__obf_7643f851b965922b)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_80a78f1bfaca43d3)
	}

	var __obf_b5322403a5f2d312 []byte
	__obf_b5322403a5f2d312, __obf_80a78f1bfaca43d3 = os.ReadFile(__obf_0243249d763c02e5)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_80a78f1bfaca43d3)
	}

	return os.WriteFile(__obf_7643f851b965922b, __obf_b5322403a5f2d312, 0644)
}

func FileCopy(__obf_9227f65be9f9cd0e *os.File, __obf_7643f851b965922b string) error {
	var __obf_80a78f1bfaca43d3 error
	__obf_7643f851b965922b, __obf_80a78f1bfaca43d3 = filepath.Abs(__obf_7643f851b965922b)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_80a78f1bfaca43d3)
	}

	__obf_bdbac1049fb4cb5a, __obf_80a78f1bfaca43d3 := os.Create(__obf_7643f851b965922b)
	if __obf_80a78f1bfaca43d3 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_80a78f1bfaca43d3)
	}
	defer __obf_bdbac1049fb4cb5a.Close()

	_, __obf_80a78f1bfaca43d3 = io.Copy(__obf_bdbac1049fb4cb5a, __obf_9227f65be9f9cd0e)
	return __obf_80a78f1bfaca43d3
}

func FileZip(
	__obf_1e692b2caedbe42f io.Writer,
	__obf_3d05926052056282 uint16,
	__obf_58ac4ace6141f6fb ...struct {
		Name string
		Path string
	},
) error {
	__obf_688a9f2f53bbab8e := zip.NewWriter(__obf_1e692b2caedbe42f)
	defer __obf_688a9f2f53bbab8e.Close()

	var (
		__obf_80a78f1bfaca43d3 error
		__obf_3d4b3bb26b0a9646 *os.File
		__obf_541d987341b55d2b io.Writer
		__obf_41c312ac426f65ee fs.FileInfo
		__obf_23f60d87108b1e23 *zip.FileHeader
	)
	for _, __obf_37a13444ba0aee54 := range __obf_58ac4ace6141f6fb {
		__obf_3d4b3bb26b0a9646, __obf_80a78f1bfaca43d3 = os.Open(__obf_37a13444ba0aee54.Path)
		if __obf_80a78f1bfaca43d3 != nil {
			return __obf_80a78f1bfaca43d3
		}
		defer __obf_3d4b3bb26b0a9646.Close()
		__obf_41c312ac426f65ee, __obf_80a78f1bfaca43d3 = __obf_3d4b3bb26b0a9646.Stat()
		if __obf_80a78f1bfaca43d3 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_37a13444ba0aee54.Path, __obf_80a78f1bfaca43d3)
		}
		__obf_23f60d87108b1e23, __obf_80a78f1bfaca43d3 = zip.FileInfoHeader(__obf_41c312ac426f65ee)
		if __obf_80a78f1bfaca43d3 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_80a78f1bfaca43d3)
		}
		// 压缩算法
		__obf_23f60d87108b1e23.Method = __obf_3d05926052056282
		// header.Name = filepath.Base(v)
		__obf_23f60d87108b1e23.Name = __obf_37a13444ba0aee54.Name
		__obf_541d987341b55d2b, __obf_80a78f1bfaca43d3 = __obf_688a9f2f53bbab8e.CreateHeader(__obf_23f60d87108b1e23)
		if __obf_80a78f1bfaca43d3 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_80a78f1bfaca43d3)
		}
		_, __obf_80a78f1bfaca43d3 = io.Copy(__obf_541d987341b55d2b, __obf_3d4b3bb26b0a9646)
		if __obf_80a78f1bfaca43d3 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_80a78f1bfaca43d3)
		}
	}
	return nil
}

func WriteCSV(__obf_ef2d507a631b5565 string, __obf_8b4207081337153d [][]string) (__obf_80a78f1bfaca43d3 error) {
	var __obf_daac877933277f10 *os.File
	__obf_daac877933277f10, __obf_80a78f1bfaca43d3 = os.OpenFile(__obf_ef2d507a631b5565+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	defer __obf_daac877933277f10.Close()
	_, __obf_80a78f1bfaca43d3 = __obf_daac877933277f10.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	__obf_2a9d3977ed95b56f := csv.NewWriter(__obf_daac877933277f10)
	__obf_80a78f1bfaca43d3 = __obf_2a9d3977ed95b56f.WriteAll(__obf_8b4207081337153d)
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	__obf_2a9d3977ed95b56f.Flush()
	return
}

func ReadCSV(path string) (__obf_8b4207081337153d [][]string, __obf_80a78f1bfaca43d3 error) {
	var fs *os.File
	fs, __obf_80a78f1bfaca43d3 = os.Open(path)
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	defer fs.Close()
	__obf_4edebc331d8e8496 := csv.NewReader(fs)
	__obf_4edebc331d8e8496.TrimLeadingSpace = true
	__obf_8b4207081337153d, __obf_80a78f1bfaca43d3 = __obf_4edebc331d8e8496.ReadAll()
	if __obf_80a78f1bfaca43d3 != nil {
		return
	}
	return
}

func __obf_8d1f63bc92ef6c7b(__obf_6eee62656413907b string) bool {
	_, __obf_80a78f1bfaca43d3 := os.Stat(__obf_6eee62656413907b)
	return !os.IsNotExist(__obf_80a78f1bfaca43d3)
}

func GetFileExt(__obf_7dd39572bc9871b5 string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_7dd39572bc9871b5), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_50ceaccf34b32b4b string) error {
	if !__obf_8d1f63bc92ef6c7b(__obf_50ceaccf34b32b4b) {
		return os.MkdirAll(__obf_50ceaccf34b32b4b, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_ef2d507a631b5565, __obf_c573de5b483263b2 string) *os.File {
	if !__obf_8d1f63bc92ef6c7b(__obf_ef2d507a631b5565) {
		__obf_4edebc331d8e8496 := []rune(__obf_ef2d507a631b5565)
		if __obf_0ccb3bbe367656b6 := strings.LastIndex(__obf_ef2d507a631b5565, __obf_c573de5b483263b2); __obf_0ccb3bbe367656b6 != -1 {
			if __obf_80a78f1bfaca43d3 := os.MkdirAll(string(__obf_4edebc331d8e8496[:__obf_0ccb3bbe367656b6]), os.ModePerm); __obf_80a78f1bfaca43d3 != nil {
				log.Fatalln(`Mkdir`, __obf_80a78f1bfaca43d3)
			}
		}
	}

	__obf_b3f9ba974f202dc7, __obf_80a78f1bfaca43d3 := os.Create(__obf_ef2d507a631b5565) //创建文件
	if __obf_80a78f1bfaca43d3 != nil {
		log.Fatalln(`file create`, __obf_80a78f1bfaca43d3)
	}
	return __obf_b3f9ba974f202dc7
}

func OpenFile(__obf_ef2d507a631b5565, __obf_c573de5b483263b2 string) *os.File {
	if !__obf_8d1f63bc92ef6c7b(__obf_ef2d507a631b5565) {
		__obf_4edebc331d8e8496 := []rune(__obf_ef2d507a631b5565)
		if __obf_0ccb3bbe367656b6 := strings.LastIndex(__obf_ef2d507a631b5565, __obf_c573de5b483263b2); __obf_0ccb3bbe367656b6 != -1 {
			if __obf_80a78f1bfaca43d3 := os.MkdirAll(string(__obf_4edebc331d8e8496[:__obf_0ccb3bbe367656b6]), os.ModePerm); __obf_80a78f1bfaca43d3 != nil {
				log.Fatalln(`Mkdir`, __obf_80a78f1bfaca43d3)
			}
		}
	}

	__obf_b3f9ba974f202dc7, __obf_80a78f1bfaca43d3 := os.OpenFile(__obf_ef2d507a631b5565, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_80a78f1bfaca43d3 != nil {
		log.Fatalln(`file create`, __obf_80a78f1bfaca43d3)
	}
	return __obf_b3f9ba974f202dc7
}
