package __obf_b81118ac905f398e

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

func FileGzip(__obf_919144fdd1f21680 *os.File) error {
	__obf_b02e11d37a5bb073 := __obf_919144fdd1f21680.Name() + ".gz"
	__obf_6c83cff938bbe1a7, __obf_65d3bb18bf1bd15f := os.Create(__obf_b02e11d37a5bb073)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_65d3bb18bf1bd15f)
	}
	defer __obf_6c83cff938bbe1a7.Close()
	__obf_24e95199de4cb95f := gzip.NewWriter(__obf_6c83cff938bbe1a7)
	defer func() {
		__obf_24e95199de4cb95f.
			Close()
		if __obf_65d3bb18bf1bd15f != nil {
			_ = os.Remove(__obf_b02e11d37a5bb073)
			__obf_65d3bb18bf1bd15f = fmt.Errorf("failed to compress file, error(%v)", __obf_65d3bb18bf1bd15f)
		}
	}()

	_, __obf_65d3bb18bf1bd15f = io.Copy(__obf_24e95199de4cb95f,

		// 进行Sync读写时才需要使用Flush
		// return w.Flush()
		__obf_919144fdd1f21680)
	return __obf_65d3bb18bf1bd15f

}

func PathGzip(__obf_ef8864a60385f45f string) error {
	var __obf_65d3bb18bf1bd15f error
	__obf_ef8864a60385f45f, __obf_65d3bb18bf1bd15f = filepath.Abs(__obf_ef8864a60385f45f)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_65d3bb18bf1bd15f)
	}

	var __obf_919144fdd1f21680 *os.File
	__obf_919144fdd1f21680, __obf_65d3bb18bf1bd15f = os.Open(__obf_ef8864a60385f45f)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_ef8864a60385f45f, __obf_65d3bb18bf1bd15f)
	}
	defer __obf_919144fdd1f21680.Close()

	return FileGzip(__obf_919144fdd1f21680)
}

func PathCopy(__obf_c2f9b10899875aa3, __obf_bce21ea8c4a51aea string) error {
	var __obf_65d3bb18bf1bd15f error
	__obf_c2f9b10899875aa3, __obf_65d3bb18bf1bd15f = filepath.Abs(__obf_c2f9b10899875aa3)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("srcPath error: %s", __obf_65d3bb18bf1bd15f)
	}
	__obf_bce21ea8c4a51aea, __obf_65d3bb18bf1bd15f = filepath.Abs(__obf_bce21ea8c4a51aea)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_65d3bb18bf1bd15f)
	}

	var __obf_0ba22b9bf269478d []byte
	__obf_0ba22b9bf269478d, __obf_65d3bb18bf1bd15f = os.ReadFile(__obf_c2f9b10899875aa3)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_65d3bb18bf1bd15f)
	}

	return os.WriteFile(__obf_bce21ea8c4a51aea, __obf_0ba22b9bf269478d, 0644)
}

func FileCopy(__obf_37e6daad561ecb46 *os.File, __obf_bce21ea8c4a51aea string) error {
	var __obf_65d3bb18bf1bd15f error
	__obf_bce21ea8c4a51aea, __obf_65d3bb18bf1bd15f = filepath.Abs(__obf_bce21ea8c4a51aea)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_65d3bb18bf1bd15f)
	}
	__obf_bb69f2ec3ff72b96, __obf_65d3bb18bf1bd15f := os.Create(__obf_bce21ea8c4a51aea)
	if __obf_65d3bb18bf1bd15f != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_65d3bb18bf1bd15f)
	}
	defer __obf_bb69f2ec3ff72b96.Close()

	_, __obf_65d3bb18bf1bd15f = io.Copy(__obf_bb69f2ec3ff72b96, __obf_37e6daad561ecb46)
	return __obf_65d3bb18bf1bd15f
}

func FileZip(__obf_30397cf80bfef41f io.Writer, __obf_20d4feec1461acf9 uint16, __obf_16b08a5a1d6464d0 ...struct {
	Name string
	Path string
},
) error {
	__obf_164bd29a0c938240 := zip.NewWriter(__obf_30397cf80bfef41f)
	defer __obf_164bd29a0c938240.Close()

	var (
		__obf_65d3bb18bf1bd15f error
		__obf_3614a9a2333cb122 *os.File
		__obf_bc1941ac7dd0cdf3 io.Writer
		__obf_b377fc7201569f7c fs.FileInfo
		__obf_2d5a77b3945ebe73 *zip.FileHeader
	)
	for _, __obf_194e82db8912d873 := range __obf_16b08a5a1d6464d0 {
		__obf_3614a9a2333cb122, __obf_65d3bb18bf1bd15f = os.Open(__obf_194e82db8912d873.Path)
		if __obf_65d3bb18bf1bd15f != nil {
			return __obf_65d3bb18bf1bd15f
		}
		defer __obf_3614a9a2333cb122.Close()
		__obf_b377fc7201569f7c, __obf_65d3bb18bf1bd15f = __obf_3614a9a2333cb122.Stat()
		if __obf_65d3bb18bf1bd15f != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_194e82db8912d873.Path, __obf_65d3bb18bf1bd15f)
		}
		__obf_2d5a77b3945ebe73, __obf_65d3bb18bf1bd15f = zip.FileInfoHeader(__obf_b377fc7201569f7c)
		if __obf_65d3bb18bf1bd15f != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_65d3bb18bf1bd15f)
		}
		__obf_2d5a77b3945ebe73.
			// 压缩算法
			Method = __obf_20d4feec1461acf9
		// header.Name = filepath.Base(v)
		__obf_2d5a77b3945ebe73.
			Name = __obf_194e82db8912d873.Name
		__obf_bc1941ac7dd0cdf3, __obf_65d3bb18bf1bd15f = __obf_164bd29a0c938240.CreateHeader(__obf_2d5a77b3945ebe73)
		if __obf_65d3bb18bf1bd15f != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_65d3bb18bf1bd15f)
		}
		_, __obf_65d3bb18bf1bd15f = io.Copy(__obf_bc1941ac7dd0cdf3, __obf_3614a9a2333cb122)
		if __obf_65d3bb18bf1bd15f != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_65d3bb18bf1bd15f)
		}
	}
	return nil
}

func WriteCSV(__obf_783842bb1ca67ba6 string, __obf_703a701ac3bc0aab [][]string) (__obf_65d3bb18bf1bd15f error) {
	var __obf_64c3d960c4166345 *os.File
	__obf_64c3d960c4166345, __obf_65d3bb18bf1bd15f = os.OpenFile(__obf_783842bb1ca67ba6+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	defer __obf_64c3d960c4166345.Close()
	_, __obf_65d3bb18bf1bd15f = __obf_64c3d960c4166345.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	__obf_24e95199de4cb95f := csv.NewWriter(__obf_64c3d960c4166345)
	__obf_65d3bb18bf1bd15f = __obf_24e95199de4cb95f.WriteAll(__obf_703a701ac3bc0aab)
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	__obf_24e95199de4cb95f.
		Flush()
	return
}

func ReadCSV(path string) (__obf_703a701ac3bc0aab [][]string, __obf_65d3bb18bf1bd15f error) {
	var fs *os.File
	fs, __obf_65d3bb18bf1bd15f = os.Open(path)
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	defer fs.Close()
	__obf_8e8b497458f714e7 := csv.NewReader(fs)
	__obf_8e8b497458f714e7.
		TrimLeadingSpace = true
	__obf_703a701ac3bc0aab, __obf_65d3bb18bf1bd15f = __obf_8e8b497458f714e7.ReadAll()
	if __obf_65d3bb18bf1bd15f != nil {
		return
	}
	return
}

func __obf_913ad03bef5e9769(__obf_6427ba177c349111 string) bool {
	_, __obf_65d3bb18bf1bd15f := os.Stat(__obf_6427ba177c349111)
	return !os.IsNotExist(__obf_65d3bb18bf1bd15f)
}

func GetFileExt(__obf_3ca4633a0e300bce string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_3ca4633a0e300bce), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_2ad5dd4968465881 string) error {
	if !__obf_913ad03bef5e9769(__obf_2ad5dd4968465881) {
		return os.MkdirAll(__obf_2ad5dd4968465881, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_783842bb1ca67ba6, __obf_001d80c65dc4d93c string) *os.File {
	if !__obf_913ad03bef5e9769(__obf_783842bb1ca67ba6) {
		__obf_8e8b497458f714e7 := []rune(__obf_783842bb1ca67ba6)
		if __obf_cfef5ea6a21678e8 := strings.LastIndex(__obf_783842bb1ca67ba6, __obf_001d80c65dc4d93c); __obf_cfef5ea6a21678e8 != -1 {
			if __obf_65d3bb18bf1bd15f := os.MkdirAll(string(__obf_8e8b497458f714e7[:__obf_cfef5ea6a21678e8]), os.ModePerm); __obf_65d3bb18bf1bd15f != nil {
				log.Fatalln(`Mkdir`, __obf_65d3bb18bf1bd15f)
			}
		}
	}
	__obf_919144fdd1f21680, __obf_65d3bb18bf1bd15f := os.Create(__obf_783842bb1ca67ba6) //创建文件
	if __obf_65d3bb18bf1bd15f != nil {
		log.Fatalln(`file create`, __obf_65d3bb18bf1bd15f)
	}
	return __obf_919144fdd1f21680
}

func OpenFile(__obf_783842bb1ca67ba6, __obf_001d80c65dc4d93c string) *os.File {
	if !__obf_913ad03bef5e9769(__obf_783842bb1ca67ba6) {
		__obf_8e8b497458f714e7 := []rune(__obf_783842bb1ca67ba6)
		if __obf_cfef5ea6a21678e8 := strings.LastIndex(__obf_783842bb1ca67ba6, __obf_001d80c65dc4d93c); __obf_cfef5ea6a21678e8 != -1 {
			if __obf_65d3bb18bf1bd15f := os.MkdirAll(string(__obf_8e8b497458f714e7[:__obf_cfef5ea6a21678e8]), os.ModePerm); __obf_65d3bb18bf1bd15f != nil {
				log.Fatalln(`Mkdir`, __obf_65d3bb18bf1bd15f)
			}
		}
	}
	__obf_919144fdd1f21680, __obf_65d3bb18bf1bd15f := os.OpenFile(__obf_783842bb1ca67ba6, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_65d3bb18bf1bd15f != nil {
		log.Fatalln(`file create`, __obf_65d3bb18bf1bd15f)
	}
	return __obf_919144fdd1f21680
}
