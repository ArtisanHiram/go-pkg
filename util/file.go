package __obf_083c8deafa73f533

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

func FileGzip(__obf_2f555a3e46769181 *os.File) error {
	__obf_65440e514cf0946d := __obf_2f555a3e46769181.Name() + ".gz"
	__obf_e13e696a085f56b9, __obf_ab078048114898aa := os.Create(__obf_65440e514cf0946d)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_ab078048114898aa)
	}
	defer __obf_e13e696a085f56b9.Close()

	__obf_7c3afbd2ad3917b4 := gzip.NewWriter(__obf_e13e696a085f56b9)
	defer func() {
		__obf_7c3afbd2ad3917b4.Close()
		if __obf_ab078048114898aa != nil {
			_ = os.Remove(__obf_65440e514cf0946d)
			__obf_ab078048114898aa = fmt.Errorf("failed to compress file, error(%v)", __obf_ab078048114898aa)
		}
	}()

	_, __obf_ab078048114898aa = io.Copy(__obf_7c3afbd2ad3917b4, __obf_2f555a3e46769181)
	return __obf_ab078048114898aa
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_01540a72e39f0264 string) error {
	var __obf_ab078048114898aa error
	__obf_01540a72e39f0264, __obf_ab078048114898aa = filepath.Abs(__obf_01540a72e39f0264)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_ab078048114898aa)
	}

	var __obf_2f555a3e46769181 *os.File
	__obf_2f555a3e46769181, __obf_ab078048114898aa = os.Open(__obf_01540a72e39f0264)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_01540a72e39f0264, __obf_ab078048114898aa)
	}
	defer __obf_2f555a3e46769181.Close()

	return FileGzip(__obf_2f555a3e46769181)
}

func PathCopy(__obf_ebc6a1007f408692, __obf_5940879541654cee string) error {
	var __obf_ab078048114898aa error
	__obf_ebc6a1007f408692, __obf_ab078048114898aa = filepath.Abs(__obf_ebc6a1007f408692)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("srcPath error: %s", __obf_ab078048114898aa)
	}

	__obf_5940879541654cee, __obf_ab078048114898aa = filepath.Abs(__obf_5940879541654cee)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("dstPath error: %s", __obf_ab078048114898aa)
	}

	var __obf_85c12639215b6e15 []byte
	__obf_85c12639215b6e15, __obf_ab078048114898aa = os.ReadFile(__obf_ebc6a1007f408692)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_ab078048114898aa)
	}

	return os.WriteFile(__obf_5940879541654cee, __obf_85c12639215b6e15, 0644)
}

func FileCopy(__obf_472e68c1437061ec *os.File, __obf_5940879541654cee string) error {
	var __obf_ab078048114898aa error
	__obf_5940879541654cee, __obf_ab078048114898aa = filepath.Abs(__obf_5940879541654cee)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("dstPath error: %s", __obf_ab078048114898aa)
	}

	__obf_7f1a2faf30532967, __obf_ab078048114898aa := os.Create(__obf_5940879541654cee)
	if __obf_ab078048114898aa != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_ab078048114898aa)
	}
	defer __obf_7f1a2faf30532967.Close()

	_, __obf_ab078048114898aa = io.Copy(__obf_7f1a2faf30532967, __obf_472e68c1437061ec)
	return __obf_ab078048114898aa
}

func FileZip(
	__obf_a87a6682a9941d86 io.Writer,
	__obf_b6fcc61a905e10d7 uint16,
	__obf_641c50999baea69d ...struct {
		Name string
		Path string
	},
) error {
	__obf_eeeff3135f862241 := zip.NewWriter(__obf_a87a6682a9941d86)
	defer __obf_eeeff3135f862241.Close()

	var (
		__obf_ab078048114898aa error
		__obf_aee5a8ed8190bf75 *os.File
		__obf_a5af55db79596f10 io.Writer
		__obf_48b14d384a77d133 fs.FileInfo
		__obf_a8a4269787715c0a *zip.FileHeader
	)
	for _, __obf_bc97935e99909ed0 := range __obf_641c50999baea69d {
		__obf_aee5a8ed8190bf75, __obf_ab078048114898aa = os.Open(__obf_bc97935e99909ed0.Path)
		if __obf_ab078048114898aa != nil {
			return __obf_ab078048114898aa
		}
		defer __obf_aee5a8ed8190bf75.Close()
		__obf_48b14d384a77d133, __obf_ab078048114898aa = __obf_aee5a8ed8190bf75.Stat()
		if __obf_ab078048114898aa != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_bc97935e99909ed0.Path, __obf_ab078048114898aa)
		}
		__obf_a8a4269787715c0a, __obf_ab078048114898aa = zip.FileInfoHeader(__obf_48b14d384a77d133)
		if __obf_ab078048114898aa != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_ab078048114898aa)
		}
		// 压缩算法
		__obf_a8a4269787715c0a.Method = __obf_b6fcc61a905e10d7
		// header.Name = filepath.Base(v)
		__obf_a8a4269787715c0a.Name = __obf_bc97935e99909ed0.Name
		__obf_a5af55db79596f10, __obf_ab078048114898aa = __obf_eeeff3135f862241.CreateHeader(__obf_a8a4269787715c0a)
		if __obf_ab078048114898aa != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_ab078048114898aa)
		}
		_, __obf_ab078048114898aa = io.Copy(__obf_a5af55db79596f10, __obf_aee5a8ed8190bf75)
		if __obf_ab078048114898aa != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_ab078048114898aa)
		}
	}
	return nil
}

func WriteCSV(__obf_018da406c2219d4e string, __obf_9cf562b2e3294b16 [][]string) (__obf_ab078048114898aa error) {
	var __obf_83d6b45b831dcd02 *os.File
	__obf_83d6b45b831dcd02, __obf_ab078048114898aa = os.OpenFile(__obf_018da406c2219d4e+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_ab078048114898aa != nil {
		return
	}
	defer __obf_83d6b45b831dcd02.Close()
	_, __obf_ab078048114898aa = __obf_83d6b45b831dcd02.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_ab078048114898aa != nil {
		return
	}
	__obf_7c3afbd2ad3917b4 := csv.NewWriter(__obf_83d6b45b831dcd02)
	__obf_ab078048114898aa = __obf_7c3afbd2ad3917b4.WriteAll(__obf_9cf562b2e3294b16)
	if __obf_ab078048114898aa != nil {
		return
	}
	__obf_7c3afbd2ad3917b4.Flush()
	return
}

func ReadCSV(path string) (__obf_9cf562b2e3294b16 [][]string, __obf_ab078048114898aa error) {
	var fs *os.File
	fs, __obf_ab078048114898aa = os.Open(path)
	if __obf_ab078048114898aa != nil {
		return
	}
	defer fs.Close()
	__obf_c094a893233159c5 := csv.NewReader(fs)
	__obf_c094a893233159c5.TrimLeadingSpace = true
	__obf_9cf562b2e3294b16, __obf_ab078048114898aa = __obf_c094a893233159c5.ReadAll()
	if __obf_ab078048114898aa != nil {
		return
	}
	return
}

func __obf_aaa93b0c63bb31f5(__obf_efd6dcc9831c56ca string) bool {
	_, __obf_ab078048114898aa := os.Stat(__obf_efd6dcc9831c56ca)
	return !os.IsNotExist(__obf_ab078048114898aa)
}

func GetFileExt(__obf_76e233546cd73d16 string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_76e233546cd73d16), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_3b05cb334d57e29b string) error {
	if !__obf_aaa93b0c63bb31f5(__obf_3b05cb334d57e29b) {
		return os.MkdirAll(__obf_3b05cb334d57e29b, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_018da406c2219d4e, __obf_f80570083423b2d7 string) *os.File {
	if !__obf_aaa93b0c63bb31f5(__obf_018da406c2219d4e) {
		__obf_c094a893233159c5 := []rune(__obf_018da406c2219d4e)
		if __obf_d5bf5930f448d12c := strings.LastIndex(__obf_018da406c2219d4e, __obf_f80570083423b2d7); __obf_d5bf5930f448d12c != -1 {
			if __obf_ab078048114898aa := os.MkdirAll(string(__obf_c094a893233159c5[:__obf_d5bf5930f448d12c]), os.ModePerm); __obf_ab078048114898aa != nil {
				log.Fatalln(`Mkdir`, __obf_ab078048114898aa)
			}
		}
	}

	__obf_2f555a3e46769181, __obf_ab078048114898aa := os.Create(__obf_018da406c2219d4e) //创建文件
	if __obf_ab078048114898aa != nil {
		log.Fatalln(`file create`, __obf_ab078048114898aa)
	}
	return __obf_2f555a3e46769181
}

func OpenFile(__obf_018da406c2219d4e, __obf_f80570083423b2d7 string) *os.File {
	if !__obf_aaa93b0c63bb31f5(__obf_018da406c2219d4e) {
		__obf_c094a893233159c5 := []rune(__obf_018da406c2219d4e)
		if __obf_d5bf5930f448d12c := strings.LastIndex(__obf_018da406c2219d4e, __obf_f80570083423b2d7); __obf_d5bf5930f448d12c != -1 {
			if __obf_ab078048114898aa := os.MkdirAll(string(__obf_c094a893233159c5[:__obf_d5bf5930f448d12c]), os.ModePerm); __obf_ab078048114898aa != nil {
				log.Fatalln(`Mkdir`, __obf_ab078048114898aa)
			}
		}
	}

	__obf_2f555a3e46769181, __obf_ab078048114898aa := os.OpenFile(__obf_018da406c2219d4e, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_ab078048114898aa != nil {
		log.Fatalln(`file create`, __obf_ab078048114898aa)
	}
	return __obf_2f555a3e46769181
}
