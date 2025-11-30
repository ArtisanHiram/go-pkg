package __obf_69af70389b6622a3

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

func FileGzip(__obf_fc0b0dc970c2bda4 *os.File) error {
	__obf_64845b343fd2a636 := __obf_fc0b0dc970c2bda4.Name() + ".gz"
	__obf_55e7c5b6c2c2b31d, __obf_e812cfd1203ed4f3 := os.Create(__obf_64845b343fd2a636)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_e812cfd1203ed4f3)
	}
	defer __obf_55e7c5b6c2c2b31d.Close()
	__obf_fcd1b865fb3f5116 := gzip.NewWriter(__obf_55e7c5b6c2c2b31d)
	defer func() {
		__obf_fcd1b865fb3f5116.
			Close()
		if __obf_e812cfd1203ed4f3 != nil {
			_ = os.Remove(__obf_64845b343fd2a636)
			__obf_e812cfd1203ed4f3 = fmt.Errorf("failed to compress file, error(%v)", __obf_e812cfd1203ed4f3)
		}
	}()

	_, __obf_e812cfd1203ed4f3 = io.Copy(__obf_fcd1b865fb3f5116,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_fc0b0dc970c2bda4)
	return __obf_e812cfd1203ed4f3

}

func PathGzip(__obf_6de26e5d74c01162 string) error {
	var __obf_e812cfd1203ed4f3 error
	__obf_6de26e5d74c01162, __obf_e812cfd1203ed4f3 = filepath.Abs(__obf_6de26e5d74c01162)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_e812cfd1203ed4f3)
	}

	var __obf_fc0b0dc970c2bda4 *os.File
	__obf_fc0b0dc970c2bda4, __obf_e812cfd1203ed4f3 = os.Open(__obf_6de26e5d74c01162)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_6de26e5d74c01162, __obf_e812cfd1203ed4f3)
	}
	defer __obf_fc0b0dc970c2bda4.Close()

	return FileGzip(__obf_fc0b0dc970c2bda4)
}

func PathCopy(__obf_7e4efe8ecf529bee, __obf_b3ddd6750b527257 string) error {
	var __obf_e812cfd1203ed4f3 error
	__obf_7e4efe8ecf529bee, __obf_e812cfd1203ed4f3 = filepath.Abs(__obf_7e4efe8ecf529bee)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_e812cfd1203ed4f3)
	}
	__obf_b3ddd6750b527257, __obf_e812cfd1203ed4f3 = filepath.Abs(__obf_b3ddd6750b527257)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_e812cfd1203ed4f3)
	}

	var __obf_0abbb6eb16723164 []byte
	__obf_0abbb6eb16723164, __obf_e812cfd1203ed4f3 = os.ReadFile(__obf_7e4efe8ecf529bee)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_e812cfd1203ed4f3)
	}

	return os.WriteFile(__obf_b3ddd6750b527257, __obf_0abbb6eb16723164, 0644)
}

func FileCopy(__obf_6c73df4fd7667826 *os.File, __obf_b3ddd6750b527257 string) error {
	var __obf_e812cfd1203ed4f3 error
	__obf_b3ddd6750b527257, __obf_e812cfd1203ed4f3 = filepath.Abs(__obf_b3ddd6750b527257)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_e812cfd1203ed4f3)
	}
	__obf_566697b3f9c80e4b, __obf_e812cfd1203ed4f3 := os.Create(__obf_b3ddd6750b527257)
	if __obf_e812cfd1203ed4f3 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_e812cfd1203ed4f3)
	}
	defer __obf_566697b3f9c80e4b.Close()

	_, __obf_e812cfd1203ed4f3 = io.Copy(__obf_566697b3f9c80e4b, __obf_6c73df4fd7667826)
	return __obf_e812cfd1203ed4f3
}

func FileZip(__obf_b770f62e5b686e22 io.Writer, __obf_6f82587df47de0ec uint16, __obf_6fb8112d553c288e ...struct {
	Name string
	Path string
},
) error {
	__obf_4e737df401389adb := zip.NewWriter(__obf_b770f62e5b686e22)
	defer __obf_4e737df401389adb.Close()

	var (
		__obf_e812cfd1203ed4f3 error
		__obf_baa972c590e4129b *os.File
		__obf_a29751187d96289e io.Writer
		__obf_7ea83908ffc18d2b fs.FileInfo
		__obf_40d41721ff4f9819 *zip.FileHeader
	)
	for _, __obf_b1dd41bcd3423053 := range __obf_6fb8112d553c288e {
		__obf_baa972c590e4129b, __obf_e812cfd1203ed4f3 = os.Open(__obf_b1dd41bcd3423053.Path)
		if __obf_e812cfd1203ed4f3 != nil {
			return __obf_e812cfd1203ed4f3
		}
		defer __obf_baa972c590e4129b.Close()
		__obf_7ea83908ffc18d2b, __obf_e812cfd1203ed4f3 = __obf_baa972c590e4129b.Stat()
		if __obf_e812cfd1203ed4f3 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_b1dd41bcd3423053.Path, __obf_e812cfd1203ed4f3)
		}
		__obf_40d41721ff4f9819, __obf_e812cfd1203ed4f3 = zip.FileInfoHeader(__obf_7ea83908ffc18d2b)
		if __obf_e812cfd1203ed4f3 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_e812cfd1203ed4f3)
		}
		__obf_40d41721ff4f9819.
			// 压缩算法
			Method = __obf_6f82587df47de0ec
		// header.Name = filepath.Base(v)
		__obf_40d41721ff4f9819.
			Name = __obf_b1dd41bcd3423053.Name
		__obf_a29751187d96289e, __obf_e812cfd1203ed4f3 = __obf_4e737df401389adb.CreateHeader(__obf_40d41721ff4f9819)
		if __obf_e812cfd1203ed4f3 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_e812cfd1203ed4f3)
		}
		_, __obf_e812cfd1203ed4f3 = io.Copy(__obf_a29751187d96289e, __obf_baa972c590e4129b)
		if __obf_e812cfd1203ed4f3 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_e812cfd1203ed4f3)
		}
	}
	return nil
}

func WriteCSV(__obf_04d9a2c744f64089 string, __obf_2801af39b525757c [][]string) (__obf_e812cfd1203ed4f3 error) {
	var __obf_f71ba171e91a4cf7 *os.File
	__obf_f71ba171e91a4cf7, __obf_e812cfd1203ed4f3 = os.OpenFile(__obf_04d9a2c744f64089+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	defer __obf_f71ba171e91a4cf7.Close()
	_, __obf_e812cfd1203ed4f3 = __obf_f71ba171e91a4cf7.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	__obf_fcd1b865fb3f5116 := csv.NewWriter(__obf_f71ba171e91a4cf7)
	__obf_e812cfd1203ed4f3 = __obf_fcd1b865fb3f5116.WriteAll(__obf_2801af39b525757c)
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	__obf_fcd1b865fb3f5116.
		Flush()
	return
}

func ReadCSV(path string) (__obf_2801af39b525757c [][]string, __obf_e812cfd1203ed4f3 error) {
	var fs *os.File
	fs, __obf_e812cfd1203ed4f3 = os.Open(path)
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	defer fs.Close()
	__obf_675dd561f1bbbbff := csv.NewReader(fs)
	__obf_675dd561f1bbbbff.
		TrimLeadingSpace = true
	__obf_2801af39b525757c, __obf_e812cfd1203ed4f3 = __obf_675dd561f1bbbbff.ReadAll()
	if __obf_e812cfd1203ed4f3 != nil {
		return
	}
	return
}

func __obf_20942562abdaa2f4(__obf_8a7dd8c3f484bf31 string) bool {
	_, __obf_e812cfd1203ed4f3 := os.Stat(__obf_8a7dd8c3f484bf31)
	return !os.IsNotExist(__obf_e812cfd1203ed4f3)
}

func GetFileExt(__obf_3dfd7aeec906cc0d string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_3dfd7aeec906cc0d), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_f4364007def7280f string) error {
	if !__obf_20942562abdaa2f4(__obf_f4364007def7280f) {
		return os.MkdirAll(__obf_f4364007def7280f, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_04d9a2c744f64089, __obf_5b24539688807f6c string) *os.File {
	if !__obf_20942562abdaa2f4(__obf_04d9a2c744f64089) {
		__obf_675dd561f1bbbbff := []rune(__obf_04d9a2c744f64089)
		if __obf_5b1c1c5422ae12c5 := strings.LastIndex(__obf_04d9a2c744f64089, __obf_5b24539688807f6c); __obf_5b1c1c5422ae12c5 != -1 {
			if __obf_e812cfd1203ed4f3 := os.MkdirAll(string(__obf_675dd561f1bbbbff[:__obf_5b1c1c5422ae12c5]), os.ModePerm); __obf_e812cfd1203ed4f3 != nil {
				log.Fatalln(`Mkdir`, __obf_e812cfd1203ed4f3)
			}
		}
	}
	__obf_fc0b0dc970c2bda4, __obf_e812cfd1203ed4f3 := os.Create(__obf_04d9a2c744f64089) //创建文件
	if __obf_e812cfd1203ed4f3 != nil {
		log.Fatalln(`file create`, __obf_e812cfd1203ed4f3)
	}
	return __obf_fc0b0dc970c2bda4
}

func OpenFile(__obf_04d9a2c744f64089, __obf_5b24539688807f6c string) *os.File {
	if !__obf_20942562abdaa2f4(__obf_04d9a2c744f64089) {
		__obf_675dd561f1bbbbff := []rune(__obf_04d9a2c744f64089)
		if __obf_5b1c1c5422ae12c5 := strings.LastIndex(__obf_04d9a2c744f64089, __obf_5b24539688807f6c); __obf_5b1c1c5422ae12c5 != -1 {
			if __obf_e812cfd1203ed4f3 := os.MkdirAll(string(__obf_675dd561f1bbbbff[:__obf_5b1c1c5422ae12c5]), os.ModePerm); __obf_e812cfd1203ed4f3 != nil {
				log.Fatalln(`Mkdir`, __obf_e812cfd1203ed4f3)
			}
		}
	}
	__obf_fc0b0dc970c2bda4, __obf_e812cfd1203ed4f3 := os.OpenFile(__obf_04d9a2c744f64089, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_e812cfd1203ed4f3 != nil {
		log.Fatalln(`file create`, __obf_e812cfd1203ed4f3)
	}
	return __obf_fc0b0dc970c2bda4
}
