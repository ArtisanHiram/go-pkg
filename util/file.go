package __obf_077bcefb098a89af

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

func FileGzip(__obf_f5af2cf557c7f764 *os.File) error {
	__obf_9302312ace4ea232 := __obf_f5af2cf557c7f764.Name() + ".gz"
	__obf_af22803c8cca761e, __obf_224415a75b186177 := os.Create(__obf_9302312ace4ea232)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("failed to create file, error(%v)", __obf_224415a75b186177)
	}
	defer __obf_af22803c8cca761e.Close()

	__obf_e8fe267011335827 := gzip.NewWriter(__obf_af22803c8cca761e)
	defer func() {
		__obf_e8fe267011335827.Close()
		if __obf_224415a75b186177 != nil {
			_ = os.Remove(__obf_9302312ace4ea232)
			__obf_224415a75b186177 = fmt.Errorf("failed to compress file, error(%v)", __obf_224415a75b186177)
		}
	}()

	_, __obf_224415a75b186177 = io.Copy(__obf_e8fe267011335827, __obf_f5af2cf557c7f764)
	return __obf_224415a75b186177
	// 进行Sync读写时才需要使用Flush
	// return w.Flush()
}

func PathGzip(__obf_ef7ec8d7d744b4d2 string) error {
	var __obf_224415a75b186177 error
	__obf_ef7ec8d7d744b4d2, __obf_224415a75b186177 = filepath.Abs(__obf_ef7ec8d7d744b4d2)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("failed to get absolute representation of path, error(%v)", __obf_224415a75b186177)
	}

	var __obf_f5af2cf557c7f764 *os.File
	__obf_f5af2cf557c7f764, __obf_224415a75b186177 = os.Open(__obf_ef7ec8d7d744b4d2)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("failed to open the path: %s, error(%v)", __obf_ef7ec8d7d744b4d2, __obf_224415a75b186177)
	}
	defer __obf_f5af2cf557c7f764.Close()

	return FileGzip(__obf_f5af2cf557c7f764)
}

func PathCopy(__obf_aca93b5bbc1d58d6, __obf_c817e910fef94c17 string) error {
	var __obf_224415a75b186177 error
	__obf_aca93b5bbc1d58d6, __obf_224415a75b186177 = filepath.Abs(__obf_aca93b5bbc1d58d6)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("srcPath error: %s", __obf_224415a75b186177)
	}

	__obf_c817e910fef94c17, __obf_224415a75b186177 = filepath.Abs(__obf_c817e910fef94c17)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_224415a75b186177)
	}

	var __obf_8192b130f0ab95cd []byte
	__obf_8192b130f0ab95cd, __obf_224415a75b186177 = os.ReadFile(__obf_aca93b5bbc1d58d6)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("read srcPath error: %s", __obf_224415a75b186177)
	}

	return os.WriteFile(__obf_c817e910fef94c17, __obf_8192b130f0ab95cd, 0644)
}

func FileCopy(__obf_024d5eb28cf8baaf *os.File, __obf_c817e910fef94c17 string) error {
	var __obf_224415a75b186177 error
	__obf_c817e910fef94c17, __obf_224415a75b186177 = filepath.Abs(__obf_c817e910fef94c17)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("dstPath error: %s", __obf_224415a75b186177)
	}

	__obf_f19de3d25da94be7, __obf_224415a75b186177 := os.Create(__obf_c817e910fef94c17)
	if __obf_224415a75b186177 != nil {
		return fmt.Errorf("creat dstPath: %s", __obf_224415a75b186177)
	}
	defer __obf_f19de3d25da94be7.Close()

	_, __obf_224415a75b186177 = io.Copy(__obf_f19de3d25da94be7, __obf_024d5eb28cf8baaf)
	return __obf_224415a75b186177
}

func FileZip(
	__obf_b77b552e491d79f8 io.Writer,
	__obf_ec97a9e427800642 uint16,
	__obf_1da8943d4f351337 ...struct {
		Name string
		Path string
	},
) error {
	__obf_95e2a7ecd87c9f24 := zip.NewWriter(__obf_b77b552e491d79f8)
	defer __obf_95e2a7ecd87c9f24.Close()

	var (
		__obf_224415a75b186177 error
		__obf_b6cd8ba272207fe0 *os.File
		__obf_a4aef53a81224d79 io.Writer
		__obf_e3d74a0be2f4b93c fs.FileInfo
		__obf_a3fac713318e555d *zip.FileHeader
	)
	for _, __obf_1c4edfc2acf87521 := range __obf_1da8943d4f351337 {
		__obf_b6cd8ba272207fe0, __obf_224415a75b186177 = os.Open(__obf_1c4edfc2acf87521.Path)
		if __obf_224415a75b186177 != nil {
			return __obf_224415a75b186177
		}
		defer __obf_b6cd8ba272207fe0.Close()
		__obf_e3d74a0be2f4b93c, __obf_224415a75b186177 = __obf_b6cd8ba272207fe0.Stat()
		if __obf_224415a75b186177 != nil {
			return fmt.Errorf("failed to stat the file: %s, error(%v)", __obf_1c4edfc2acf87521.Path, __obf_224415a75b186177)
		}
		__obf_a3fac713318e555d, __obf_224415a75b186177 = zip.FileInfoHeader(__obf_e3d74a0be2f4b93c)
		if __obf_224415a75b186177 != nil {
			return fmt.Errorf("failed to get FileInfoHeader, error(%v)", __obf_224415a75b186177)
		}
		// 压缩算法
		__obf_a3fac713318e555d.Method = __obf_ec97a9e427800642
		// header.Name = filepath.Base(v)
		__obf_a3fac713318e555d.Name = __obf_1c4edfc2acf87521.Name
		__obf_a4aef53a81224d79, __obf_224415a75b186177 = __obf_95e2a7ecd87c9f24.CreateHeader(__obf_a3fac713318e555d)
		if __obf_224415a75b186177 != nil {
			return fmt.Errorf("failed to create zip header, error(%v)", __obf_224415a75b186177)
		}
		_, __obf_224415a75b186177 = io.Copy(__obf_a4aef53a81224d79, __obf_b6cd8ba272207fe0)
		if __obf_224415a75b186177 != nil {
			return fmt.Errorf("failed to copy the file, error(%v)", __obf_224415a75b186177)
		}
	}
	return nil
}

func WriteCSV(__obf_1e523056a8de9a9c string, __obf_f8bbf5dce647ada9 [][]string) (__obf_224415a75b186177 error) {
	var __obf_9260bb8a5d6f665a *os.File
	__obf_9260bb8a5d6f665a, __obf_224415a75b186177 = os.OpenFile(__obf_1e523056a8de9a9c+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) // 创建文件句柄
	if __obf_224415a75b186177 != nil {
		return
	}
	defer __obf_9260bb8a5d6f665a.Close()
	_, __obf_224415a75b186177 = __obf_9260bb8a5d6f665a.Write([]byte{0xEF, 0xBB, 0xBF}) // 写入UTF-8 BOM

	if __obf_224415a75b186177 != nil {
		return
	}
	__obf_e8fe267011335827 := csv.NewWriter(__obf_9260bb8a5d6f665a)
	__obf_224415a75b186177 = __obf_e8fe267011335827.WriteAll(__obf_f8bbf5dce647ada9)
	if __obf_224415a75b186177 != nil {
		return
	}
	__obf_e8fe267011335827.Flush()
	return
}

func ReadCSV(path string) (__obf_f8bbf5dce647ada9 [][]string, __obf_224415a75b186177 error) {
	var fs *os.File
	fs, __obf_224415a75b186177 = os.Open(path)
	if __obf_224415a75b186177 != nil {
		return
	}
	defer fs.Close()
	__obf_c07a5c3f92c3565e := csv.NewReader(fs)
	__obf_c07a5c3f92c3565e.TrimLeadingSpace = true
	__obf_f8bbf5dce647ada9, __obf_224415a75b186177 = __obf_c07a5c3f92c3565e.ReadAll()
	if __obf_224415a75b186177 != nil {
		return
	}
	return
}

func __obf_a478f024dfbaacdc(__obf_075fd7b3346969fd string) bool {
	_, __obf_224415a75b186177 := os.Stat(__obf_075fd7b3346969fd)
	return !os.IsNotExist(__obf_224415a75b186177)
}

func GetFileExt(__obf_33e6d02219ac6b6f string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(__obf_33e6d02219ac6b6f), "."))
}

// 创建目录（可以是多级目录，任何一级不存在则创建）
func TouchDir(__obf_a8c35f257405f0dd string) error {
	if !__obf_a478f024dfbaacdc(__obf_a8c35f257405f0dd) {
		return os.MkdirAll(__obf_a8c35f257405f0dd, os.ModePerm)
	}
	return nil
}

func CreateFile(__obf_1e523056a8de9a9c, __obf_53aa887d3a1a5d3b string) *os.File {
	if !__obf_a478f024dfbaacdc(__obf_1e523056a8de9a9c) {
		__obf_c07a5c3f92c3565e := []rune(__obf_1e523056a8de9a9c)
		if __obf_c705c49030e00b6f := strings.LastIndex(__obf_1e523056a8de9a9c, __obf_53aa887d3a1a5d3b); __obf_c705c49030e00b6f != -1 {
			if __obf_224415a75b186177 := os.MkdirAll(string(__obf_c07a5c3f92c3565e[:__obf_c705c49030e00b6f]), os.ModePerm); __obf_224415a75b186177 != nil {
				log.Fatalln(`Mkdir`, __obf_224415a75b186177)
			}
		}
	}

	__obf_f5af2cf557c7f764, __obf_224415a75b186177 := os.Create(__obf_1e523056a8de9a9c) //创建文件
	if __obf_224415a75b186177 != nil {
		log.Fatalln(`file create`, __obf_224415a75b186177)
	}
	return __obf_f5af2cf557c7f764
}

func OpenFile(__obf_1e523056a8de9a9c, __obf_53aa887d3a1a5d3b string) *os.File {
	if !__obf_a478f024dfbaacdc(__obf_1e523056a8de9a9c) {
		__obf_c07a5c3f92c3565e := []rune(__obf_1e523056a8de9a9c)
		if __obf_c705c49030e00b6f := strings.LastIndex(__obf_1e523056a8de9a9c, __obf_53aa887d3a1a5d3b); __obf_c705c49030e00b6f != -1 {
			if __obf_224415a75b186177 := os.MkdirAll(string(__obf_c07a5c3f92c3565e[:__obf_c705c49030e00b6f]), os.ModePerm); __obf_224415a75b186177 != nil {
				log.Fatalln(`Mkdir`, __obf_224415a75b186177)
			}
		}
	}

	__obf_f5af2cf557c7f764, __obf_224415a75b186177 := os.OpenFile(__obf_1e523056a8de9a9c, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if __obf_224415a75b186177 != nil {
		log.Fatalln(`file create`, __obf_224415a75b186177)
	}
	return __obf_f5af2cf557c7f764
}
