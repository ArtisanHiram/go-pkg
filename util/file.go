package __obf_f1e346e3aa5cc554

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

func FileGzip(__obf_1027c1dd0a81527d *os.File) error {
	__obf_452f3c367a889b92 := __obf_1027c1dd0a81527d.Name() + ".gz"
	__obf_d8ce902c843669fd, __obf_eec784b359ebf42f := os.Create(__obf_452f3c367a889b92)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_eec784b359ebf42f)
	}
	defer __obf_d8ce902c843669fd.Close()

	__obf_6e1c75e795080772 := gzip.NewWriter(__obf_d8ce902c843669fd)
	defer func() {
		__obf_6e1c75e795080772.Close()
		if __obf_eec784b359ebf42f != nil {
			_ = os.Remove(__obf_452f3c367a889b92)
			__obf_eec784b359ebf42f = fmt.Errorf("failed to compress file, error(%v)", __obf_eec784b359ebf42f)
		}
	}()

	_, __obf_eec784b359ebf42f = io.Copy(__obf_6e1c75e795080772, __obf_1027c1dd0a81527d)
	return __obf_eec784b359ebf42f
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_882e31e5ee8959a6 string) error {
	var __obf_eec784b359ebf42f error
	__obf_882e31e5ee8959a6, __obf_eec784b359ebf42f = filepath.Abs(__obf_882e31e5ee8959a6)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_eec784b359ebf42f)
	}

	var __obf_1027c1dd0a81527d *os.File
	__obf_1027c1dd0a81527d, __obf_eec784b359ebf42f = os.Open(__obf_882e31e5ee8959a6)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_882e31e5ee8959a6, __obf_eec784b359ebf42f)
	}
	defer __obf_1027c1dd0a81527d.Close()

	return FileGzip(__obf_1027c1dd0a81527d)
}

func PathCopy(__obf_75aa186c243e6ab1, __obf_0d0bb141d8e586ec string) error {
	var __obf_eec784b359ebf42f error
	__obf_75aa186c243e6ab1, __obf_eec784b359ebf42f = filepath.Abs(__obf_75aa186c243e6ab1)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("srcPath error: %s", __obf_eec784b359ebf42f)
	}

	__obf_0d0bb141d8e586ec, __obf_eec784b359ebf42f = filepath.Abs(__obf_0d0bb141d8e586ec)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_eec784b359ebf42f)
	}

	var __obf_7857bbfa11265353 []byte
	__obf_7857bbfa11265353, __obf_eec784b359ebf42f = os.ReadFile(__obf_75aa186c243e6ab1)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_eec784b359ebf42f)
	}

	return os.WriteFile(__obf_0d0bb141d8e586ec, __obf_7857bbfa11265353, 0644)
}

func FileCopy(__obf_0519622334e71da3 *os.File, __obf_0d0bb141d8e586ec string) error {
	var __obf_eec784b359ebf42f error
	__obf_0d0bb141d8e586ec, __obf_eec784b359ebf42f = filepath.Abs(__obf_0d0bb141d8e586ec)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("dstPath error: %s", __obf_eec784b359ebf42f)
	}

	__obf_eaf2753e8a77b6e9, __obf_eec784b359ebf42f := os.Create(__obf_0d0bb141d8e586ec)
	if __obf_eec784b359ebf42f != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_eec784b359ebf42f)
	}
	defer __obf_eaf2753e8a77b6e9.Close()

	_, __obf_eec784b359ebf42f = io.Copy(__obf_eaf2753e8a77b6e9, __obf_0519622334e71da3)
	return __obf_eec784b359ebf42f
}

func FileZip(
	__obf_71a5e0a1806f3baa io.Writer,
	__obf_1ba3303e8df762cc uint16,
	__obf_ac1f80c958d4d9f8 ...struct {
		Name string
		Path string
	},
) error {
	__obf_97d3ddc04463422a := zip.NewWriter(__obf_71a5e0a1806f3baa)
	defer __obf_97d3ddc04463422a.Close()

	var (
		__obf_eec784b359ebf42f error
		__obf_f411b526a092c63c *os.File
		__obf_ddf3a2be5adf066f io.Writer
		__obf_0399d349c9773b80 fs.FileInfo
		__obf_e4abfae33d7c744a *zip.FileHeader
	)
	for _, __obf_29f3385acfb70f69 := range __obf_ac1f80c958d4d9f8 {
		__obf_f411b526a092c63c, __obf_eec784b359ebf42f = os.Open(__obf_29f3385acfb70f69.Path)
		if __obf_eec784b359ebf42f != nil {
			return __obf_eec784b359ebf42f
		}
		defer __obf_f411b526a092c63c.Close()
		__obf_0399d349c9773b80, __obf_eec784b359ebf42f = __obf_f411b526a092c63c.Stat()
		if __obf_eec784b359ebf42f != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_29f3385acfb70f69.Path, __obf_eec784b359ebf42f)
		}
		__obf_e4abfae33d7c744a, __obf_eec784b359ebf42f = zip.FileInfoHeader(__obf_0399d349c9773b80)
		if __obf_eec784b359ebf42f != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_eec784b359ebf42f)
		}
		// 压缩算法
		__obf_e4abfae33d7c744a.Method = __obf_1ba3303e8df762cc
		// header.Name = filepath.Base(v)
		__obf_e4abfae33d7c744a.Name = __obf_29f3385acfb70f69.Name
		__obf_ddf3a2be5adf066f, __obf_eec784b359ebf42f = __obf_97d3ddc04463422a.CreateHeader(__obf_e4abfae33d7c744a)
		if __obf_eec784b359ebf42f != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_eec784b359ebf42f)
		}
		_, __obf_eec784b359ebf42f = io.Copy(__obf_ddf3a2be5adf066f, __obf_f411b526a092c63c)
		if __obf_eec784b359ebf42f != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_eec784b359ebf42f)
		}
	}
	return nil
}

func WriteCSV(__obf_1d62aee7a77ffb9c string, __obf_e8e5c333aaebfc88 [][]string) (__obf_eec784b359ebf42f error) {
	var __obf_fb654f3eabf3cd14 *os.File
	__obf_fb654f3eabf3cd14, __obf_eec784b359ebf42f = os.OpenFile(__obf_1d62aee7a77ffb9c+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_eec784b359ebf42f != nil {
		return
	}
	defer __obf_fb654f3eabf3cd14.Close()
	_, __obf_eec784b359ebf42f = __obf_fb654f3eabf3cd14.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_eec784b359ebf42f != nil {
		return
	}
	__obf_6e1c75e795080772 := csv.NewWriter(__obf_fb654f3eabf3cd14)
	__obf_eec784b359ebf42f = __obf_6e1c75e795080772.WriteAll(__obf_e8e5c333aaebfc88)
	if __obf_eec784b359ebf42f != nil {
		return
	}
	__obf_6e1c75e795080772.Flush()
	return
}

func ReadCSV(path string) (__obf_e8e5c333aaebfc88 [][]string, __obf_eec784b359ebf42f error) {
	var fs *os.File
	fs, __obf_eec784b359ebf42f = os.Open(path)
	if __obf_eec784b359ebf42f != nil {
		return
	}
	defer fs.Close()
	__obf_72fc7f85fed6cac3 := csv.NewReader(fs)
	__obf_72fc7f85fed6cac3.TrimLeadingSpace = true
	__obf_e8e5c333aaebfc88, __obf_eec784b359ebf42f = __obf_72fc7f85fed6cac3.ReadAll()
	if __obf_eec784b359ebf42f != nil {
		return
	}
	return
}

func __obf_603cc89171100346(__obf_4d2a16dc437c92be string) bool {
	_, __obf_eec784b359ebf42f := os.Stat(__obf_4d2a16dc437c92be)
	return !os.IsNotExist(__obf_eec784b359ebf42f)
}

func GetFileExt(__obf_d57dde5414214749 string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_d57dde5414214749), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_a43cf39ce0f643ae string) error {
	if !__obf_603cc89171100346(__obf_a43cf39ce0f643ae) {
		return os.MkdirAll(__obf_a43cf39ce0f643ae, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_1d62aee7a77ffb9c, __obf_41f2064bbc5b2f8c string) *os.File {
	if !__obf_603cc89171100346(__obf_1d62aee7a77ffb9c) {
		__obf_72fc7f85fed6cac3 := []rune(__obf_1d62aee7a77ffb9c)
		if __obf_719365bda976491c := strings.LastIndex(__obf_1d62aee7a77ffb9c, __obf_41f2064bbc5b2f8c); __obf_719365bda976491c != -1 {
			if __obf_eec784b359ebf42f := os.MkdirAll(string(__obf_72fc7f85fed6cac3[:__obf_719365bda976491c]), os.ModePerm); __obf_eec784b359ebf42f != nil {
				log.Fatalln(`Mkdir`, __obf_eec784b359ebf42f)
			}
		}
	}

	__obf_1027c1dd0a81527d, __obf_eec784b359ebf42f := os.Create(__obf_1d62aee7a77ffb9c) //创建文件
	if __obf_eec784b359ebf42f != nil {
		log.Fatalln(`file create`, __obf_eec784b359ebf42f)
	}
	return __obf_1027c1dd0a81527d
}

func OpenFile(__obf_1d62aee7a77ffb9c, __obf_41f2064bbc5b2f8c string) *os.File {
	if !__obf_603cc89171100346(__obf_1d62aee7a77ffb9c) {
		__obf_72fc7f85fed6cac3 := []rune(__obf_1d62aee7a77ffb9c)
		if __obf_719365bda976491c := strings.LastIndex(__obf_1d62aee7a77ffb9c, __obf_41f2064bbc5b2f8c); __obf_719365bda976491c != -1 {
			if __obf_eec784b359ebf42f := os.MkdirAll(string(__obf_72fc7f85fed6cac3[:__obf_719365bda976491c]), os.ModePerm); __obf_eec784b359ebf42f != nil {
				log.Fatalln(`Mkdir`, __obf_eec784b359ebf42f)
			}
		}
	}

	__obf_1027c1dd0a81527d, __obf_eec784b359ebf42f := os.OpenFile(__obf_1d62aee7a77ffb9c, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_eec784b359ebf42f != nil {
		log.Fatalln(`file create`, __obf_eec784b359ebf42f)
	}
	return __obf_1027c1dd0a81527d
}
