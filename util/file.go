package __obf_d7b39e56b82f7f57

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

func FileGzip(__obf_1615e38c49fb8375 *os.File) error {
	__obf_fc566d7a0c766786 := __obf_1615e38c49fb8375.Name() + ".gz"
	__obf_809a71b9939e3b59, __obf_3ccf8c32165eeb13 := os.Create(__obf_fc566d7a0c766786)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_3ccf8c32165eeb13)
	}
	defer __obf_809a71b9939e3b59.Close()
	__obf_e844bc825768c1a5 := gzip.NewWriter(__obf_809a71b9939e3b59)
	defer func() {
		__obf_e844bc825768c1a5.
			Close()
		if __obf_3ccf8c32165eeb13 != nil {
			_ = os.Remove(__obf_fc566d7a0c766786)
			__obf_3ccf8c32165eeb13 = fmt.Errorf("failed to compress file, error(%v)", __obf_3ccf8c32165eeb13)
		}
	}()

	_, __obf_3ccf8c32165eeb13 = io.Copy(__obf_e844bc825768c1a5,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_1615e38c49fb8375)
	return __obf_3ccf8c32165eeb13

}

func PathGzip(__obf_845324066d8df0dd string) error {
	var __obf_3ccf8c32165eeb13 error
	__obf_845324066d8df0dd, __obf_3ccf8c32165eeb13 = filepath.Abs(__obf_845324066d8df0dd)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_3ccf8c32165eeb13)
	}

	var __obf_1615e38c49fb8375 *os.File
	__obf_1615e38c49fb8375, __obf_3ccf8c32165eeb13 = os.Open(__obf_845324066d8df0dd)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_845324066d8df0dd, __obf_3ccf8c32165eeb13)
	}
	defer __obf_1615e38c49fb8375.Close()

	return FileGzip(__obf_1615e38c49fb8375)
}

func PathCopy(__obf_92088373c173e2ad, __obf_c1d49ab7c494ed14 string) error {
	var __obf_3ccf8c32165eeb13 error
	__obf_92088373c173e2ad, __obf_3ccf8c32165eeb13 = filepath.Abs(__obf_92088373c173e2ad)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_3ccf8c32165eeb13)
	}
	__obf_c1d49ab7c494ed14, __obf_3ccf8c32165eeb13 = filepath.Abs(__obf_c1d49ab7c494ed14)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_3ccf8c32165eeb13)
	}

	var __obf_574ecbf0cff9834b []byte
	__obf_574ecbf0cff9834b, __obf_3ccf8c32165eeb13 = os.ReadFile(__obf_92088373c173e2ad)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_3ccf8c32165eeb13)
	}

	return os.WriteFile(__obf_c1d49ab7c494ed14, __obf_574ecbf0cff9834b, 0644)
}

func FileCopy(__obf_db2f02741885afcf *os.File, __obf_c1d49ab7c494ed14 string) error {
	var __obf_3ccf8c32165eeb13 error
	__obf_c1d49ab7c494ed14, __obf_3ccf8c32165eeb13 = filepath.Abs(__obf_c1d49ab7c494ed14)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_3ccf8c32165eeb13)
	}
	__obf_2fb73aaa77231bda, __obf_3ccf8c32165eeb13 := os.Create(__obf_c1d49ab7c494ed14)
	if __obf_3ccf8c32165eeb13 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_3ccf8c32165eeb13)
	}
	defer __obf_2fb73aaa77231bda.Close()

	_, __obf_3ccf8c32165eeb13 = io.Copy(__obf_2fb73aaa77231bda, __obf_db2f02741885afcf)
	return __obf_3ccf8c32165eeb13
}

func FileZip(__obf_d8f7b6afe67ad13b io.Writer, __obf_321bfaf6a5b20ce9 uint16, __obf_def516632b2b7340 ...struct {
	Name string
	Path string
},
) error {
	__obf_75a67a5b4e76aa8e := zip.NewWriter(__obf_d8f7b6afe67ad13b)
	defer __obf_75a67a5b4e76aa8e.Close()

	var (
		__obf_3ccf8c32165eeb13 error
		__obf_0d7f9f32f0948310 *os.File
		__obf_51ee5cf062b26e8d io.Writer
		__obf_f3eca6dca5babea5 fs.FileInfo
		__obf_52f2d71bb987997f *zip.FileHeader
	)
	for _, __obf_a30d33f1a4aa6e72 := range __obf_def516632b2b7340 {
		__obf_0d7f9f32f0948310, __obf_3ccf8c32165eeb13 = os.Open(__obf_a30d33f1a4aa6e72.Path)
		if __obf_3ccf8c32165eeb13 != nil {
			return __obf_3ccf8c32165eeb13
		}
		defer __obf_0d7f9f32f0948310.Close()
		__obf_f3eca6dca5babea5, __obf_3ccf8c32165eeb13 = __obf_0d7f9f32f0948310.Stat()
		if __obf_3ccf8c32165eeb13 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_a30d33f1a4aa6e72.Path, __obf_3ccf8c32165eeb13)
		}
		__obf_52f2d71bb987997f, __obf_3ccf8c32165eeb13 = zip.FileInfoHeader(__obf_f3eca6dca5babea5)
		if __obf_3ccf8c32165eeb13 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_3ccf8c32165eeb13)
		}
		__obf_52f2d71bb987997f.
			// 压缩算法
			Method = __obf_321bfaf6a5b20ce9
		// header.Name = filepath.Base(v)
		__obf_52f2d71bb987997f.
			Name = __obf_a30d33f1a4aa6e72.Name
		__obf_51ee5cf062b26e8d, __obf_3ccf8c32165eeb13 = __obf_75a67a5b4e76aa8e.CreateHeader(__obf_52f2d71bb987997f)
		if __obf_3ccf8c32165eeb13 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_3ccf8c32165eeb13)
		}
		_, __obf_3ccf8c32165eeb13 = io.Copy(__obf_51ee5cf062b26e8d, __obf_0d7f9f32f0948310)
		if __obf_3ccf8c32165eeb13 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_3ccf8c32165eeb13)
		}
	}
	return nil
}

func WriteCSV(__obf_a65865c83781bbce string, __obf_c2744e465ce68a8d [][]string) (__obf_3ccf8c32165eeb13 error) {
	var __obf_703917c59b37f1b8 *os.File
	__obf_703917c59b37f1b8, __obf_3ccf8c32165eeb13 = os.OpenFile(__obf_a65865c83781bbce+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	defer __obf_703917c59b37f1b8.Close()
	_, __obf_3ccf8c32165eeb13 = __obf_703917c59b37f1b8.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	__obf_e844bc825768c1a5 := csv.NewWriter(__obf_703917c59b37f1b8)
	__obf_3ccf8c32165eeb13 = __obf_e844bc825768c1a5.WriteAll(__obf_c2744e465ce68a8d)
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	__obf_e844bc825768c1a5.
		Flush()
	return
}

func ReadCSV(path string) (__obf_c2744e465ce68a8d [][]string, __obf_3ccf8c32165eeb13 error) {
	var fs *os.File
	fs, __obf_3ccf8c32165eeb13 = os.Open(path)
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	defer fs.Close()
	__obf_403793979a29000e := csv.NewReader(fs)
	__obf_403793979a29000e.
		TrimLeadingSpace = true
	__obf_c2744e465ce68a8d, __obf_3ccf8c32165eeb13 = __obf_403793979a29000e.ReadAll()
	if __obf_3ccf8c32165eeb13 != nil {
		return
	}
	return
}

func __obf_299bd152e171723b(__obf_5ad072acfd05203c string) bool {
	_, __obf_3ccf8c32165eeb13 := os.Stat(__obf_5ad072acfd05203c)
	return !os.IsNotExist(__obf_3ccf8c32165eeb13)
}

func GetFileExt(__obf_625c67f1289ed5eb string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_625c67f1289ed5eb), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_b9f8b2cbd7c7e158 string) error {
	if !__obf_299bd152e171723b(__obf_b9f8b2cbd7c7e158) {
		return os.MkdirAll(__obf_b9f8b2cbd7c7e158, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_a65865c83781bbce, __obf_bbd2c7f8b6865e3b string) *os.File {
	if !__obf_299bd152e171723b(__obf_a65865c83781bbce) {
		__obf_403793979a29000e := []rune(__obf_a65865c83781bbce)
		if __obf_2688d2e7e665fec7 := strings.LastIndex(__obf_a65865c83781bbce, __obf_bbd2c7f8b6865e3b); __obf_2688d2e7e665fec7 != -1 {
			if __obf_3ccf8c32165eeb13 := os.MkdirAll(string(__obf_403793979a29000e[:__obf_2688d2e7e665fec7]), os.ModePerm); __obf_3ccf8c32165eeb13 != nil {
				log.Fatalln(`Mkdir`, __obf_3ccf8c32165eeb13)
			}
		}
	}
	__obf_1615e38c49fb8375, __obf_3ccf8c32165eeb13 := os.Create(__obf_a65865c83781bbce) //创建文件
	if __obf_3ccf8c32165eeb13 != nil {
		log.Fatalln(`file create`, __obf_3ccf8c32165eeb13)
	}
	return __obf_1615e38c49fb8375
}

func OpenFile(__obf_a65865c83781bbce, __obf_bbd2c7f8b6865e3b string) *os.File {
	if !__obf_299bd152e171723b(__obf_a65865c83781bbce) {
		__obf_403793979a29000e := []rune(__obf_a65865c83781bbce)
		if __obf_2688d2e7e665fec7 := strings.LastIndex(__obf_a65865c83781bbce, __obf_bbd2c7f8b6865e3b); __obf_2688d2e7e665fec7 != -1 {
			if __obf_3ccf8c32165eeb13 := os.MkdirAll(string(__obf_403793979a29000e[:__obf_2688d2e7e665fec7]), os.ModePerm); __obf_3ccf8c32165eeb13 != nil {
				log.Fatalln(`Mkdir`, __obf_3ccf8c32165eeb13)
			}
		}
	}
	__obf_1615e38c49fb8375, __obf_3ccf8c32165eeb13 := os.OpenFile(__obf_a65865c83781bbce, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_3ccf8c32165eeb13 != nil {
		log.Fatalln(`file create`, __obf_3ccf8c32165eeb13)
	}
	return __obf_1615e38c49fb8375
}
