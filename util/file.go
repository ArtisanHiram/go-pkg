package __obf_813b65e0329aecbf

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

func FileGzip(__obf_49f027f6b6c03687 *os.File) error {
	__obf_4a5d7d9a4de34488 := __obf_49f027f6b6c03687.Name() + ".gz"
	__obf_d07f7077db35fbcc, __obf_54101b1325683820 := os.Create(__obf_4a5d7d9a4de34488)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_54101b1325683820)
	}
	defer __obf_d07f7077db35fbcc.Close()

	__obf_a1bfde50aa747e83 := gzip.NewWriter(__obf_d07f7077db35fbcc)
	defer func() {
		__obf_a1bfde50aa747e83.Close()
		if __obf_54101b1325683820 != nil {
			_ = os.Remove(__obf_4a5d7d9a4de34488)
			__obf_54101b1325683820 = fmt.Errorf("failed to compress file, error(%v)", __obf_54101b1325683820)
		}
	}()

	_, __obf_54101b1325683820 = io.Copy(__obf_a1bfde50aa747e83, __obf_49f027f6b6c03687)
	return __obf_54101b1325683820
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_303f08634c2b3966 string) error {
	var __obf_54101b1325683820 error
	__obf_303f08634c2b3966, __obf_54101b1325683820 = filepath.Abs(__obf_303f08634c2b3966)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_54101b1325683820)
	}

	var __obf_49f027f6b6c03687 *os.File
	__obf_49f027f6b6c03687, __obf_54101b1325683820 = os.Open(__obf_303f08634c2b3966)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_303f08634c2b3966, __obf_54101b1325683820)
	}
	defer __obf_49f027f6b6c03687.Close()

	return FileGzip(__obf_49f027f6b6c03687)
}

func PathCopy(__obf_a5395669ea668aaf, __obf_86d8604f6c248ad5 string) error {
	var __obf_54101b1325683820 error
	__obf_a5395669ea668aaf, __obf_54101b1325683820 = filepath.Abs(__obf_a5395669ea668aaf)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_54101b1325683820)
	}

	__obf_86d8604f6c248ad5, __obf_54101b1325683820 = filepath.Abs(__obf_86d8604f6c248ad5)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_54101b1325683820)
	}

	var __obf_cfc13f9f797e3385 []byte
	__obf_cfc13f9f797e3385, __obf_54101b1325683820 = os.ReadFile(__obf_a5395669ea668aaf)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_54101b1325683820)
	}

	return os.WriteFile(__obf_86d8604f6c248ad5, __obf_cfc13f9f797e3385, 0644)
}

func FileCopy(__obf_eb9f2e197d91e28c *os.File, __obf_86d8604f6c248ad5 string) error {
	var __obf_54101b1325683820 error
	__obf_86d8604f6c248ad5, __obf_54101b1325683820 = filepath.Abs(__obf_86d8604f6c248ad5)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_54101b1325683820)
	}

	__obf_071db135d0378af1, __obf_54101b1325683820 := os.Create(__obf_86d8604f6c248ad5)
	if __obf_54101b1325683820 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_54101b1325683820)
	}
	defer __obf_071db135d0378af1.Close()

	_, __obf_54101b1325683820 = io.Copy(__obf_071db135d0378af1, __obf_eb9f2e197d91e28c)
	return __obf_54101b1325683820
}

func FileZip(
	__obf_a61194f4b5a0990f io.Writer,
	__obf_36863999f4a56f28 uint16,
	__obf_36f511fc7589036d ...struct {
		Name string
		Path string
	},
) error {
	__obf_7b0b0303969a897b := zip.NewWriter(__obf_a61194f4b5a0990f)
	defer __obf_7b0b0303969a897b.Close()

	var (
		__obf_54101b1325683820 error
		__obf_9a6d458f604677cf *os.File
		__obf_e1402ac330412846 io.Writer
		__obf_62fb8424e69827f0 fs.FileInfo
		__obf_0de36dd4d0ae04f7 *zip.FileHeader
	)
	for _, __obf_2dfe4b44e04cc8de := range __obf_36f511fc7589036d {
		__obf_9a6d458f604677cf, __obf_54101b1325683820 = os.Open(__obf_2dfe4b44e04cc8de.Path)
		if __obf_54101b1325683820 != nil {
			return __obf_54101b1325683820
		}
		defer __obf_9a6d458f604677cf.Close()
		__obf_62fb8424e69827f0, __obf_54101b1325683820 = __obf_9a6d458f604677cf.Stat()
		if __obf_54101b1325683820 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_2dfe4b44e04cc8de.Path, __obf_54101b1325683820)
		}
		__obf_0de36dd4d0ae04f7, __obf_54101b1325683820 = zip.FileInfoHeader(__obf_62fb8424e69827f0)
		if __obf_54101b1325683820 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_54101b1325683820)
		}
		// 压缩算法
		__obf_0de36dd4d0ae04f7.Method = __obf_36863999f4a56f28
		// header.Name = filepath.Base(v)
		__obf_0de36dd4d0ae04f7.Name = __obf_2dfe4b44e04cc8de.Name
		__obf_e1402ac330412846, __obf_54101b1325683820 = __obf_7b0b0303969a897b.CreateHeader(__obf_0de36dd4d0ae04f7)
		if __obf_54101b1325683820 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_54101b1325683820)
		}
		_, __obf_54101b1325683820 = io.Copy(__obf_e1402ac330412846, __obf_9a6d458f604677cf)
		if __obf_54101b1325683820 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_54101b1325683820)
		}
	}
	return nil
}

func WriteCSV(__obf_794ec8367f651eb7 string, __obf_30267a212643ccd8 [][]string) (__obf_54101b1325683820 error) {
	var __obf_0a61369f6603c881 *os.File
	__obf_0a61369f6603c881, __obf_54101b1325683820 = os.OpenFile(__obf_794ec8367f651eb7+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_54101b1325683820 != nil {
		return
	}
	defer __obf_0a61369f6603c881.Close()
	_, __obf_54101b1325683820 = __obf_0a61369f6603c881.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_54101b1325683820 != nil {
		return
	}
	__obf_a1bfde50aa747e83 := csv.NewWriter(__obf_0a61369f6603c881)
	__obf_54101b1325683820 = __obf_a1bfde50aa747e83.WriteAll(__obf_30267a212643ccd8)
	if __obf_54101b1325683820 != nil {
		return
	}
	__obf_a1bfde50aa747e83.Flush()
	return
}

func ReadCSV(path string) (__obf_30267a212643ccd8 [][]string, __obf_54101b1325683820 error) {
	var fs *os.File
	fs, __obf_54101b1325683820 = os.Open(path)
	if __obf_54101b1325683820 != nil {
		return
	}
	defer fs.Close()
	__obf_138ad5c24fa33841 := csv.NewReader(fs)
	__obf_138ad5c24fa33841.TrimLeadingSpace = true
	__obf_30267a212643ccd8, __obf_54101b1325683820 = __obf_138ad5c24fa33841.ReadAll()
	if __obf_54101b1325683820 != nil {
		return
	}
	return
}

func __obf_fa331053a34d1fc3(__obf_999db615e71b11d1 string) bool {
	_, __obf_54101b1325683820 := os.Stat(__obf_999db615e71b11d1)
	return !os.IsNotExist(__obf_54101b1325683820)
}

func GetFileExt(__obf_92b42f8fc44f8e70 string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_92b42f8fc44f8e70), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_4eafc8cd582720c2 string) error {
	if !__obf_fa331053a34d1fc3(__obf_4eafc8cd582720c2) {
		return os.MkdirAll(__obf_4eafc8cd582720c2, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_794ec8367f651eb7, __obf_ea8c99a9b7b45ad7 string) *os.File {
	if !__obf_fa331053a34d1fc3(__obf_794ec8367f651eb7) {
		__obf_138ad5c24fa33841 := []rune(__obf_794ec8367f651eb7)
		if __obf_2c6b6b0f61b708f9 := strings.LastIndex(__obf_794ec8367f651eb7, __obf_ea8c99a9b7b45ad7); __obf_2c6b6b0f61b708f9 != -1 {
			if __obf_54101b1325683820 := os.MkdirAll(string(__obf_138ad5c24fa33841[:__obf_2c6b6b0f61b708f9]), os.ModePerm); __obf_54101b1325683820 != nil {
				log.Fatalln(`Mkdir`, __obf_54101b1325683820)
			}
		}
	}

	__obf_49f027f6b6c03687, __obf_54101b1325683820 := os.Create(__obf_794ec8367f651eb7) //创建文件
	if __obf_54101b1325683820 != nil {
		log.Fatalln(`file create`, __obf_54101b1325683820)
	}
	return __obf_49f027f6b6c03687
}

func OpenFile(__obf_794ec8367f651eb7, __obf_ea8c99a9b7b45ad7 string) *os.File {
	if !__obf_fa331053a34d1fc3(__obf_794ec8367f651eb7) {
		__obf_138ad5c24fa33841 := []rune(__obf_794ec8367f651eb7)
		if __obf_2c6b6b0f61b708f9 := strings.LastIndex(__obf_794ec8367f651eb7, __obf_ea8c99a9b7b45ad7); __obf_2c6b6b0f61b708f9 != -1 {
			if __obf_54101b1325683820 := os.MkdirAll(string(__obf_138ad5c24fa33841[:__obf_2c6b6b0f61b708f9]), os.ModePerm); __obf_54101b1325683820 != nil {
				log.Fatalln(`Mkdir`, __obf_54101b1325683820)
			}
		}
	}

	__obf_49f027f6b6c03687, __obf_54101b1325683820 := os.OpenFile(__obf_794ec8367f651eb7, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_54101b1325683820 != nil {
		log.Fatalln(`file create`, __obf_54101b1325683820)
	}
	return __obf_49f027f6b6c03687
}
