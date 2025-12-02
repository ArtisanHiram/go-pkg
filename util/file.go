package __obf_8b17832dd38bb602

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

func FileGzip(__obf_6754aaa26fa71907 *os.File) error {
	__obf_3baabfc8c2e47646 := __obf_6754aaa26fa71907.Name() + ".gz"
	__obf_c8ee636c993e4bc8, __obf_17881c470b1e6626 := os.Create(__obf_3baabfc8c2e47646)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_17881c470b1e6626)
	}
	defer __obf_c8ee636c993e4bc8.Close()
	__obf_32e6a0729889a343 := gzip.NewWriter(__obf_c8ee636c993e4bc8)
	defer func() {
		__obf_32e6a0729889a343.
			Close()
		if __obf_17881c470b1e6626 != nil {
			_ = os.Remove(__obf_3baabfc8c2e47646)
			__obf_17881c470b1e6626 = fmt.Errorf("failed to compress file, error(%v)", __obf_17881c470b1e6626)
		}
	}()

	_, __obf_17881c470b1e6626 = io.Copy(__obf_32e6a0729889a343,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_6754aaa26fa71907)
	return __obf_17881c470b1e6626

}

func PathGzip(__obf_def1d103abf8ef3d string) error {
	var __obf_17881c470b1e6626 error
	__obf_def1d103abf8ef3d, __obf_17881c470b1e6626 = filepath.Abs(__obf_def1d103abf8ef3d)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_17881c470b1e6626)
	}

	var __obf_6754aaa26fa71907 *os.File
	__obf_6754aaa26fa71907, __obf_17881c470b1e6626 = os.Open(__obf_def1d103abf8ef3d)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_def1d103abf8ef3d, __obf_17881c470b1e6626)
	}
	defer __obf_6754aaa26fa71907.Close()

	return FileGzip(__obf_6754aaa26fa71907)
}

func PathCopy(__obf_bd2df54ca69f3f8b, __obf_d2e6c429c951eb70 string) error {
	var __obf_17881c470b1e6626 error
	__obf_bd2df54ca69f3f8b, __obf_17881c470b1e6626 = filepath.Abs(__obf_bd2df54ca69f3f8b)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_17881c470b1e6626)
	}
	__obf_d2e6c429c951eb70, __obf_17881c470b1e6626 = filepath.Abs(__obf_d2e6c429c951eb70)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_17881c470b1e6626)
	}

	var __obf_3f61171b0204f03a []byte
	__obf_3f61171b0204f03a, __obf_17881c470b1e6626 = os.ReadFile(__obf_bd2df54ca69f3f8b)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_17881c470b1e6626)
	}

	return os.WriteFile(__obf_d2e6c429c951eb70, __obf_3f61171b0204f03a, 0644)
}

func FileCopy(__obf_0e6c3d368006f7e7 *os.File, __obf_d2e6c429c951eb70 string) error {
	var __obf_17881c470b1e6626 error
	__obf_d2e6c429c951eb70, __obf_17881c470b1e6626 = filepath.Abs(__obf_d2e6c429c951eb70)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_17881c470b1e6626)
	}
	__obf_e906b0ee3883f2b3, __obf_17881c470b1e6626 := os.Create(__obf_d2e6c429c951eb70)
	if __obf_17881c470b1e6626 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_17881c470b1e6626)
	}
	defer __obf_e906b0ee3883f2b3.Close()

	_, __obf_17881c470b1e6626 = io.Copy(__obf_e906b0ee3883f2b3, __obf_0e6c3d368006f7e7)
	return __obf_17881c470b1e6626
}

func FileZip(__obf_4517662b5afbf57c io.Writer, __obf_b827fdb7ccfd87b3 uint16, __obf_6e113e0dc133cb4d ...struct {
	Name string
	Path string
},
) error {
	__obf_da3a20800d84ce36 := zip.NewWriter(__obf_4517662b5afbf57c)
	defer __obf_da3a20800d84ce36.Close()

	var (
		__obf_17881c470b1e6626 error
		__obf_fda1617e22a16a24 *os.File
		__obf_b16e31e4574f487d io.Writer
		__obf_23296f9e58f5f8c7 fs.FileInfo
		__obf_49560f6bcd5717e9 *zip.FileHeader
	)
	for _, __obf_d4f65aef7baa5a20 := range __obf_6e113e0dc133cb4d {
		__obf_fda1617e22a16a24, __obf_17881c470b1e6626 = os.Open(__obf_d4f65aef7baa5a20.Path)
		if __obf_17881c470b1e6626 != nil {
			return __obf_17881c470b1e6626
		}
		defer __obf_fda1617e22a16a24.Close()
		__obf_23296f9e58f5f8c7, __obf_17881c470b1e6626 = __obf_fda1617e22a16a24.Stat()
		if __obf_17881c470b1e6626 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_d4f65aef7baa5a20.Path, __obf_17881c470b1e6626)
		}
		__obf_49560f6bcd5717e9, __obf_17881c470b1e6626 = zip.FileInfoHeader(__obf_23296f9e58f5f8c7)
		if __obf_17881c470b1e6626 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_17881c470b1e6626)
		}
		__obf_49560f6bcd5717e9.
			// 压缩算法
			Method = __obf_b827fdb7ccfd87b3
		// header.Name = filepath.Base(v)
		__obf_49560f6bcd5717e9.
			Name = __obf_d4f65aef7baa5a20.Name
		__obf_b16e31e4574f487d, __obf_17881c470b1e6626 = __obf_da3a20800d84ce36.CreateHeader(__obf_49560f6bcd5717e9)
		if __obf_17881c470b1e6626 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_17881c470b1e6626)
		}
		_, __obf_17881c470b1e6626 = io.Copy(__obf_b16e31e4574f487d, __obf_fda1617e22a16a24)
		if __obf_17881c470b1e6626 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_17881c470b1e6626)
		}
	}
	return nil
}

func WriteCSV(__obf_b8bb78e4260e7c51 string, __obf_cab20f4d57b74374 [][]string) (__obf_17881c470b1e6626 error) {
	var __obf_48d8acc85eb9b8a9 *os.File
	__obf_48d8acc85eb9b8a9, __obf_17881c470b1e6626 = os.OpenFile(__obf_b8bb78e4260e7c51+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_17881c470b1e6626 != nil {
		return
	}
	defer __obf_48d8acc85eb9b8a9.Close()
	_, __obf_17881c470b1e6626 = __obf_48d8acc85eb9b8a9.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_17881c470b1e6626 != nil {
		return
	}
	__obf_32e6a0729889a343 := csv.NewWriter(__obf_48d8acc85eb9b8a9)
	__obf_17881c470b1e6626 = __obf_32e6a0729889a343.WriteAll(__obf_cab20f4d57b74374)
	if __obf_17881c470b1e6626 != nil {
		return
	}
	__obf_32e6a0729889a343.
		Flush()
	return
}

func ReadCSV(path string) (__obf_cab20f4d57b74374 [][]string, __obf_17881c470b1e6626 error) {
	var fs *os.File
	fs, __obf_17881c470b1e6626 = os.Open(path)
	if __obf_17881c470b1e6626 != nil {
		return
	}
	defer fs.Close()
	__obf_274b76b7499b5c52 := csv.NewReader(fs)
	__obf_274b76b7499b5c52.
		TrimLeadingSpace = true
	__obf_cab20f4d57b74374, __obf_17881c470b1e6626 = __obf_274b76b7499b5c52.ReadAll()
	if __obf_17881c470b1e6626 != nil {
		return
	}
	return
}

func __obf_628ebea6a7a40d14(__obf_60b44cf0449a03d9 string) bool {
	_, __obf_17881c470b1e6626 := os.Stat(__obf_60b44cf0449a03d9)
	return !os.IsNotExist(__obf_17881c470b1e6626)
}

func GetFileExt(__obf_17da6955efc0f4cc string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_17da6955efc0f4cc), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_3c6f8f676162bd35 string) error {
	if !__obf_628ebea6a7a40d14(__obf_3c6f8f676162bd35) {
		return os.MkdirAll(__obf_3c6f8f676162bd35, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_b8bb78e4260e7c51, __obf_16a3a22302fb358d string) *os.File {
	if !__obf_628ebea6a7a40d14(__obf_b8bb78e4260e7c51) {
		__obf_274b76b7499b5c52 := []rune(__obf_b8bb78e4260e7c51)
		if __obf_786bd10169423f93 := strings.LastIndex(__obf_b8bb78e4260e7c51, __obf_16a3a22302fb358d); __obf_786bd10169423f93 != -1 {
			if __obf_17881c470b1e6626 := os.MkdirAll(string(__obf_274b76b7499b5c52[:__obf_786bd10169423f93]), os.ModePerm); __obf_17881c470b1e6626 != nil {
				log.Fatalln(`Mkdir`, __obf_17881c470b1e6626)
			}
		}
	}
	__obf_6754aaa26fa71907, __obf_17881c470b1e6626 := os.Create(__obf_b8bb78e4260e7c51) //创建文件
	if __obf_17881c470b1e6626 != nil {
		log.Fatalln(`file create`, __obf_17881c470b1e6626)
	}
	return __obf_6754aaa26fa71907
}

func OpenFile(__obf_b8bb78e4260e7c51, __obf_16a3a22302fb358d string) *os.File {
	if !__obf_628ebea6a7a40d14(__obf_b8bb78e4260e7c51) {
		__obf_274b76b7499b5c52 := []rune(__obf_b8bb78e4260e7c51)
		if __obf_786bd10169423f93 := strings.LastIndex(__obf_b8bb78e4260e7c51, __obf_16a3a22302fb358d); __obf_786bd10169423f93 != -1 {
			if __obf_17881c470b1e6626 := os.MkdirAll(string(__obf_274b76b7499b5c52[:__obf_786bd10169423f93]), os.ModePerm); __obf_17881c470b1e6626 != nil {
				log.Fatalln(`Mkdir`, __obf_17881c470b1e6626)
			}
		}
	}
	__obf_6754aaa26fa71907, __obf_17881c470b1e6626 := os.OpenFile(__obf_b8bb78e4260e7c51, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_17881c470b1e6626 != nil {
		log.Fatalln(`file create`, __obf_17881c470b1e6626)
	}
	return __obf_6754aaa26fa71907
}
