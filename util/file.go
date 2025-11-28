package __obf_d984cff8712b1ee6

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

func FileGzip(__obf_4621ced8fc91df0c *os.File) error {
	__obf_6fb796601d84de44 := __obf_4621ced8fc91df0c.Name() + ".gz"
	__obf_1d52cebcbef29401, __obf_44ebf351b5f3fef8 := os.Create(__obf_6fb796601d84de44)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_44ebf351b5f3fef8)
	}
	defer __obf_1d52cebcbef29401.Close()

	__obf_eedd3648c47f8eb2 := gzip.NewWriter(__obf_1d52cebcbef29401)
	defer func() {
		__obf_eedd3648c47f8eb2.Close()
		if __obf_44ebf351b5f3fef8 != nil {
			_ = os.Remove(__obf_6fb796601d84de44)
			__obf_44ebf351b5f3fef8 = fmt.Errorf("failed to compress file, error(%v)", __obf_44ebf351b5f3fef8)
		}
	}()

	_, __obf_44ebf351b5f3fef8 = io.Copy(__obf_eedd3648c47f8eb2, __obf_4621ced8fc91df0c)
	return __obf_44ebf351b5f3fef8
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_6a1bf30bd27d40b8 string) error {
	var __obf_44ebf351b5f3fef8 error
	__obf_6a1bf30bd27d40b8, __obf_44ebf351b5f3fef8 = filepath.Abs(__obf_6a1bf30bd27d40b8)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_44ebf351b5f3fef8)
	}

	var __obf_4621ced8fc91df0c *os.File
	__obf_4621ced8fc91df0c, __obf_44ebf351b5f3fef8 = os.Open(__obf_6a1bf30bd27d40b8)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_6a1bf30bd27d40b8, __obf_44ebf351b5f3fef8)
	}
	defer __obf_4621ced8fc91df0c.Close()

	return FileGzip(__obf_4621ced8fc91df0c)
}

func PathCopy(__obf_ece5fcad7eeac544, __obf_5b539fb1c022d93a string) error {
	var __obf_44ebf351b5f3fef8 error
	__obf_ece5fcad7eeac544, __obf_44ebf351b5f3fef8 = filepath.Abs(__obf_ece5fcad7eeac544)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_44ebf351b5f3fef8)
	}

	__obf_5b539fb1c022d93a, __obf_44ebf351b5f3fef8 = filepath.Abs(__obf_5b539fb1c022d93a)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_44ebf351b5f3fef8)
	}

	var __obf_7085eca3d8e93ef6 []byte
	__obf_7085eca3d8e93ef6, __obf_44ebf351b5f3fef8 = os.ReadFile(__obf_ece5fcad7eeac544)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_44ebf351b5f3fef8)
	}

	return os.WriteFile(__obf_5b539fb1c022d93a, __obf_7085eca3d8e93ef6, 0644)
}

func FileCopy(__obf_a1a837cc5ec66a9f *os.File, __obf_5b539fb1c022d93a string) error {
	var __obf_44ebf351b5f3fef8 error
	__obf_5b539fb1c022d93a, __obf_44ebf351b5f3fef8 = filepath.Abs(__obf_5b539fb1c022d93a)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_44ebf351b5f3fef8)
	}

	__obf_4ac08e5d9beb7887, __obf_44ebf351b5f3fef8 := os.Create(__obf_5b539fb1c022d93a)
	if __obf_44ebf351b5f3fef8 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_44ebf351b5f3fef8)
	}
	defer __obf_4ac08e5d9beb7887.Close()

	_, __obf_44ebf351b5f3fef8 = io.Copy(__obf_4ac08e5d9beb7887, __obf_a1a837cc5ec66a9f)
	return __obf_44ebf351b5f3fef8
}

func FileZip(
	__obf_c2bac87761c114f9 io.Writer,
	__obf_dd2a7a264fcfb27b uint16,
	__obf_43053eac9d2321a5 ...struct {
		Name string
		Path string
	},
) error {
	__obf_82c27b8717abe22b := zip.NewWriter(__obf_c2bac87761c114f9)
	defer __obf_82c27b8717abe22b.Close()

	var (
		__obf_44ebf351b5f3fef8 error
		__obf_791e03ba1e340788 *os.File
		__obf_e399c22409dc7aae io.Writer
		__obf_85c708c9f4c300f3 fs.FileInfo
		__obf_c87312d599dc75c6 *zip.FileHeader
	)
	for _, __obf_5ca54f845825fcc4 := range __obf_43053eac9d2321a5 {
		__obf_791e03ba1e340788, __obf_44ebf351b5f3fef8 = os.Open(__obf_5ca54f845825fcc4.Path)
		if __obf_44ebf351b5f3fef8 != nil {
			return __obf_44ebf351b5f3fef8
		}
		defer __obf_791e03ba1e340788.Close()
		__obf_85c708c9f4c300f3, __obf_44ebf351b5f3fef8 = __obf_791e03ba1e340788.Stat()
		if __obf_44ebf351b5f3fef8 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_5ca54f845825fcc4.Path, __obf_44ebf351b5f3fef8)
		}
		__obf_c87312d599dc75c6, __obf_44ebf351b5f3fef8 = zip.FileInfoHeader(__obf_85c708c9f4c300f3)
		if __obf_44ebf351b5f3fef8 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_44ebf351b5f3fef8)
		}
		// 压缩算法
		__obf_c87312d599dc75c6.Method = __obf_dd2a7a264fcfb27b
		// header.Name = filepath.Base(v)
		__obf_c87312d599dc75c6.Name = __obf_5ca54f845825fcc4.Name
		__obf_e399c22409dc7aae, __obf_44ebf351b5f3fef8 = __obf_82c27b8717abe22b.CreateHeader(__obf_c87312d599dc75c6)
		if __obf_44ebf351b5f3fef8 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_44ebf351b5f3fef8)
		}
		_, __obf_44ebf351b5f3fef8 = io.Copy(__obf_e399c22409dc7aae, __obf_791e03ba1e340788)
		if __obf_44ebf351b5f3fef8 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_44ebf351b5f3fef8)
		}
	}
	return nil
}

func WriteCSV(__obf_0a5084207611bf1f string, __obf_fff9dcb5e6c38578 [][]string) (__obf_44ebf351b5f3fef8 error) {
	var __obf_58c4f447f6829925 *os.File
	__obf_58c4f447f6829925, __obf_44ebf351b5f3fef8 = os.OpenFile(__obf_0a5084207611bf1f+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	defer __obf_58c4f447f6829925.Close()
	_, __obf_44ebf351b5f3fef8 = __obf_58c4f447f6829925.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	__obf_eedd3648c47f8eb2 := csv.NewWriter(__obf_58c4f447f6829925)
	__obf_44ebf351b5f3fef8 = __obf_eedd3648c47f8eb2.WriteAll(__obf_fff9dcb5e6c38578)
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	__obf_eedd3648c47f8eb2.Flush()
	return
}

func ReadCSV(path string) (__obf_fff9dcb5e6c38578 [][]string, __obf_44ebf351b5f3fef8 error) {
	var fs *os.File
	fs, __obf_44ebf351b5f3fef8 = os.Open(path)
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	defer fs.Close()
	__obf_d7f59788250a1a1e := csv.NewReader(fs)
	__obf_d7f59788250a1a1e.TrimLeadingSpace = true
	__obf_fff9dcb5e6c38578, __obf_44ebf351b5f3fef8 = __obf_d7f59788250a1a1e.ReadAll()
	if __obf_44ebf351b5f3fef8 != nil {
		return
	}
	return
}

func __obf_b506b1c6e1a2bb92(__obf_31babcd135b23047 string) bool {
	_, __obf_44ebf351b5f3fef8 := os.Stat(__obf_31babcd135b23047)
	return !os.IsNotExist(__obf_44ebf351b5f3fef8)
}

func GetFileExt(__obf_00ab4781e7b0cf5c string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_00ab4781e7b0cf5c), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_ac86c08c6ef63652 string) error {
	if !__obf_b506b1c6e1a2bb92(__obf_ac86c08c6ef63652) {
		return os.MkdirAll(__obf_ac86c08c6ef63652, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_0a5084207611bf1f, __obf_88ce94801855d048 string) *os.File {
	if !__obf_b506b1c6e1a2bb92(__obf_0a5084207611bf1f) {
		__obf_d7f59788250a1a1e := []rune(__obf_0a5084207611bf1f)
		if __obf_1b64762461a96f61 := strings.LastIndex(__obf_0a5084207611bf1f, __obf_88ce94801855d048); __obf_1b64762461a96f61 != -1 {
			if __obf_44ebf351b5f3fef8 := os.MkdirAll(string(__obf_d7f59788250a1a1e[:__obf_1b64762461a96f61]), os.ModePerm); __obf_44ebf351b5f3fef8 != nil {
				log.Fatalln(`Mkdir`, __obf_44ebf351b5f3fef8)
			}
		}
	}

	__obf_4621ced8fc91df0c, __obf_44ebf351b5f3fef8 := os.Create(__obf_0a5084207611bf1f) //创建文件
	if __obf_44ebf351b5f3fef8 != nil {
		log.Fatalln(`file create`, __obf_44ebf351b5f3fef8)
	}
	return __obf_4621ced8fc91df0c
}

func OpenFile(__obf_0a5084207611bf1f, __obf_88ce94801855d048 string) *os.File {
	if !__obf_b506b1c6e1a2bb92(__obf_0a5084207611bf1f) {
		__obf_d7f59788250a1a1e := []rune(__obf_0a5084207611bf1f)
		if __obf_1b64762461a96f61 := strings.LastIndex(__obf_0a5084207611bf1f, __obf_88ce94801855d048); __obf_1b64762461a96f61 != -1 {
			if __obf_44ebf351b5f3fef8 := os.MkdirAll(string(__obf_d7f59788250a1a1e[:__obf_1b64762461a96f61]), os.ModePerm); __obf_44ebf351b5f3fef8 != nil {
				log.Fatalln(`Mkdir`, __obf_44ebf351b5f3fef8)
			}
		}
	}

	__obf_4621ced8fc91df0c, __obf_44ebf351b5f3fef8 := os.OpenFile(__obf_0a5084207611bf1f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_44ebf351b5f3fef8 != nil {
		log.Fatalln(`file create`, __obf_44ebf351b5f3fef8)
	}
	return __obf_4621ced8fc91df0c
}
