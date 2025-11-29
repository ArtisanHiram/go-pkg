// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_ddb9362593c4afda

import (
	"errors"
	"fmt"
	util "github.com/ArtisanHiram/go-pkg/util"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// ensure LogWriter always implement io.WriteCloser
var _ io.WriteCloser = (*LogWriter)(nil)

const (
	__obf_b92c039a12d148ca = "2006-01-02_15:04:05.000"
	__obf_fe4da2b16e24370c = ".gz"
)

type Option struct {
	// 日志存放路径
	LogPath string `yaml:"log-path"`
	// 日志保存期限（天）
	MaxAge time.Duration `yaml:"max-age"`
	// 单个日志文件大小（MB），超过则rotate
	MaxSize int64 `yaml:"max-size"`
	// 日志是否压缩
	// Compress bool `yaml:"compress"`
}

// LogWriter is a logger with rotation function
type LogWriter struct {
	__obf_4d53f6f79fd7e19e Option
	__obf_546b139714a6d61f *os.File
	__obf_5d4e76bb85d64194 int64
	__obf_e57e6de8bca69c4c string
	__obf_9a46d75de02f1711 sync.Mutex
	__obf_62f4a5b868867ee6 sync.Once
	__obf_98ef9e1b790b04b7 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_4d53f6f79fd7e19e Option) *LogWriter {
	__obf_e57e6de8bca69c4c := filepath.Dir(__obf_4d53f6f79fd7e19e.LogPath)
	__obf_4d53f6f79fd7e19e.
		MaxAge = __obf_4d53f6f79fd7e19e.MaxAge * time.Hour * 24
	__obf_4d53f6f79fd7e19e.
		MaxSize = __obf_4d53f6f79fd7e19e.MaxSize * 1024 * 1024
	__obf_5313a06ecef9be49 := &LogWriter{__obf_4d53f6f79fd7e19e: __obf_4d53f6f79fd7e19e, __obf_e57e6de8bca69c4c: __obf_e57e6de8bca69c4c}

	return __obf_5313a06ecef9be49
}

// Write writes content into file.
func (__obf_677d23f8777c6fc4 *LogWriter) Write(__obf_5e1c012c014d89c3 []byte) (__obf_7b9f687e6254a317 int, __obf_5ddb1aeb344ab2e0 error) {
	__obf_677d23f8777c6fc4.__obf_9a46d75de02f1711.
		Lock()
	defer __obf_677d23f8777c6fc4.__obf_9a46d75de02f1711.Unlock()
	__obf_86375dba3fa0b4f8 := int64(len(__obf_5e1c012c014d89c3))

	if __obf_677d23f8777c6fc4.__obf_546b139714a6d61f == nil {
		if __obf_5ddb1aeb344ab2e0 = __obf_677d23f8777c6fc4.__obf_bc371ceef1fee5b5(__obf_86375dba3fa0b4f8); __obf_5ddb1aeb344ab2e0 != nil {
			return
		}
	}
	if __obf_677d23f8777c6fc4.__obf_1a30a9251cbff51c(__obf_86375dba3fa0b4f8 + __obf_677d23f8777c6fc4.__obf_5d4e76bb85d64194) {
		if __obf_5ddb1aeb344ab2e0 = __obf_677d23f8777c6fc4.__obf_e253dc80364e401e(); __obf_5ddb1aeb344ab2e0 != nil {
			return
		}
	}
	__obf_7b9f687e6254a317, __obf_5ddb1aeb344ab2e0 = __obf_677d23f8777c6fc4.__obf_546b139714a6d61f.Write(__obf_5e1c012c014d89c3)
	__obf_677d23f8777c6fc4.__obf_5d4e76bb85d64194 += int64(__obf_7b9f687e6254a317)
	return
}

// Close closes file resource
func (__obf_677d23f8777c6fc4 *LogWriter) Close() error {
	__obf_677d23f8777c6fc4.__obf_9a46d75de02f1711.
		Lock()
	defer __obf_677d23f8777c6fc4.__obf_9a46d75de02f1711.Unlock()
	return __obf_677d23f8777c6fc4.close()
}

func (__obf_677d23f8777c6fc4 *LogWriter) close() error {
	if __obf_677d23f8777c6fc4.__obf_546b139714a6d61f == nil {
		return nil
	}
	__obf_5ddb1aeb344ab2e0 := __obf_677d23f8777c6fc4.__obf_546b139714a6d61f.Close()
	if __obf_5ddb1aeb344ab2e0 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_5ddb1aeb344ab2e0)
	}
	__obf_677d23f8777c6fc4.__obf_546b139714a6d61f = nil
	__obf_677d23f8777c6fc4.__obf_5d4e76bb85d64194 = 0
	return nil
}

func (__obf_677d23f8777c6fc4 *LogWriter) __obf_1a30a9251cbff51c(__obf_5d4e76bb85d64194 int64) bool {
	return __obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.MaxSize > 0 && __obf_5d4e76bb85d64194 > __obf_677d23f8777c6fc4.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_4d53f6f79fd7e19e.MaxSize
}

func (__obf_677d23f8777c6fc4 *LogWriter) __obf_bc371ceef1fee5b5(__obf_4119fdbf09c1a90d int64) error {
	__obf_06ec908cf4321a1b, __obf_5ddb1aeb344ab2e0 := os.Stat(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath)
	if os.IsNotExist(__obf_5ddb1aeb344ab2e0) {
		return __obf_677d23f8777c6fc4.__obf_9d5573ac971cfd15()
	}
	if __obf_5ddb1aeb344ab2e0 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_5ddb1aeb344ab2e0)
	}
	if __obf_677d23f8777c6fc4.__obf_1a30a9251cbff51c(__obf_06ec908cf4321a1b.Size() + __obf_4119fdbf09c1a90d) {
		return __obf_677d23f8777c6fc4.__obf_e253dc80364e401e()
	}
	__obf_546b139714a6d61f, __obf_5ddb1aeb344ab2e0 := os.OpenFile(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_5ddb1aeb344ab2e0 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_5ddb1aeb344ab2e0)
	}
	__obf_677d23f8777c6fc4.__obf_546b139714a6d61f = __obf_546b139714a6d61f
	__obf_677d23f8777c6fc4.

		// openNewFile 创建一个新的日志文件。
		// 如果日志文件已存在，它会通过重命名并附加时间戳的方式对现有文件进行归档。
		// 在创建新文件之前，它会确保必要的目录已存在。
		//
		// 如果发生以下情况，将返回错误：
		// - 无法创建所需的目录。
		// - 无法归档现有的日志文件。
		// - 无法创建或打开新的日志文件。
		//
		// 新日志文件的权限将继承自现有日志文件（如果存在），否则默认为 0644。
		__obf_5d4e76bb85d64194 = __obf_06ec908cf4321a1b.Size()
	return nil
}

func (__obf_677d23f8777c6fc4 *LogWriter) __obf_9d5573ac971cfd15() error {
	if __obf_5ddb1aeb344ab2e0 := os.MkdirAll(__obf_677d23f8777c6fc4.__obf_e57e6de8bca69c4c, os.ModePerm); __obf_5ddb1aeb344ab2e0 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_5ddb1aeb344ab2e0)
	}
	__obf_a1424517a7c05d76 := os.FileMode(0644)
	__obf_06ec908cf4321a1b, __obf_5ddb1aeb344ab2e0 := os.Stat(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath)
	if __obf_5ddb1aeb344ab2e0 == nil {
		__obf_a1424517a7c05d76 = __obf_06ec908cf4321a1b.Mode()
		__obf_2a48f2986f9e7da4, __obf_cdd3fd7f3fa49dfd := util.SplitFilename(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath)
		__obf_05f421a410836518 := time.Now().Format(__obf_b92c039a12d148ca)
		__obf_1176b934258a89c3 := filepath.Join(__obf_677d23f8777c6fc4.__obf_e57e6de8bca69c4c, fmt.Sprintf("%s-%s%s", __obf_2a48f2986f9e7da4, __obf_05f421a410836518, __obf_cdd3fd7f3fa49dfd))

		if __obf_5ddb1aeb344ab2e0 = os.Rename(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath, __obf_1176b934258a89c3); __obf_5ddb1aeb344ab2e0 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_1176b934258a89c3, __obf_5ddb1aeb344ab2e0)
		}
	}
	__obf_44850cb26b65dffd, __obf_5ddb1aeb344ab2e0 := os.OpenFile(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_a1424517a7c05d76)
	if __obf_5ddb1aeb344ab2e0 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath, __obf_5ddb1aeb344ab2e0)
	}
	__obf_677d23f8777c6fc4.__obf_546b139714a6d61f = __obf_44850cb26b65dffd
	__obf_677d23f8777c6fc4.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_5d4e76bb85d64194 = 0
	return nil
}

func (__obf_677d23f8777c6fc4 *LogWriter) __obf_e253dc80364e401e() error {
	if __obf_5ddb1aeb344ab2e0 := __obf_677d23f8777c6fc4.close(); __obf_5ddb1aeb344ab2e0 != nil {
		return __obf_5ddb1aeb344ab2e0
	}
	if __obf_5ddb1aeb344ab2e0 := __obf_677d23f8777c6fc4.__obf_9d5573ac971cfd15(); __obf_5ddb1aeb344ab2e0 != nil {
		return __obf_5ddb1aeb344ab2e0
	}
	__obf_677d23f8777c6fc4.__obf_62f4a5b868867ee6.
		Do(func() {
			__obf_677d23f8777c6fc4.__obf_98ef9e1b790b04b7 = make(chan bool, 1)
			go func() {
				for range __obf_677d23f8777c6fc4.__obf_98ef9e1b790b04b7 {
					if __obf_2c7475fb6bb599d1 := __obf_677d23f8777c6fc4.__obf_bc3e3b41028801ad(); __obf_2c7475fb6bb599d1 != nil {
						println(__obf_2c7475fb6bb599d1.Error())
					}
				}
			}()
		})
	select {
	case __obf_677d23f8777c6fc4.__obf_98ef9e1b790b04b7 <- true:
	default:
	}
	return nil
}

// handleArchives 负责根据 LogWriter 中配置的选项管理日志文件的归档和清理。
// 它执行以下任务：
//
// 1. 删除超过 MaxAge 选项指定的最大保存期限的日志文件。
// 2. 压缩尚未压缩的日志文件。
//
// 函数的操作流程如下：
// - 如果 MaxSize 设置为 0，则立即返回，不执行任何操作。
// - 检索归档日志文件的列表。
// - 将超过截止时间（根据 MaxAge 计算）的文件标记为待删除。
// - 检查剩余文件，并压缩未压缩的文件。
// - 在文件删除或压缩过程中遇到的任何错误都会被收集并作为单个聚合错误返回。
//
// 返回值：
// - 如果所有操作成功或不需要任何操作，则返回 nil。
// - 如果一个或多个操作失败，则返回聚合错误。
func (__obf_677d23f8777c6fc4 *LogWriter) __obf_bc3e3b41028801ad() error {
	if __obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.MaxSize == 0 {
		return nil
	}
	__obf_f8ae176e5b796c4c, __obf_5ddb1aeb344ab2e0 := __obf_677d23f8777c6fc4.__obf_8e7a4b8b1e4e9574()
	if __obf_5ddb1aeb344ab2e0 != nil {
		return __obf_5ddb1aeb344ab2e0
	}
	var __obf_3aa0dce30077ec8c, __obf_daa9772a2082598b []__obf_c755abfd19f5b475
	if __obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.MaxAge > 0 {
		__obf_1b28dee6ebdc9657 := time.Now().Add(-1 * __obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.MaxAge)
		for _, __obf_44850cb26b65dffd := range __obf_f8ae176e5b796c4c {
			if __obf_44850cb26b65dffd.__obf_1b92f09c365627ad.Before(__obf_1b28dee6ebdc9657) {
				__obf_3aa0dce30077ec8c = append(__obf_3aa0dce30077ec8c, __obf_44850cb26b65dffd)
			} else {
				__obf_daa9772a2082598b = append(__obf_daa9772a2082598b, __obf_44850cb26b65dffd)
			}
		}
	}

	var __obf_19e8b1471f655904 []error
	for _, __obf_44850cb26b65dffd := range __obf_3aa0dce30077ec8c {
		if __obf_5ddb1aeb344ab2e0 = os.Remove(filepath.Join(__obf_677d23f8777c6fc4.__obf_e57e6de8bca69c4c, __obf_44850cb26b65dffd.Name())); __obf_5ddb1aeb344ab2e0 != nil {
			__obf_19e8b1471f655904 = append(__obf_19e8b1471f655904, fmt.Errorf("remove %s: %w", __obf_44850cb26b65dffd.Name(), __obf_5ddb1aeb344ab2e0))
		}
	}
	if len(__obf_19e8b1471f655904) > 0 {
		return errors.Join(__obf_19e8b1471f655904...)
	}

	for _, __obf_44850cb26b65dffd := range __obf_daa9772a2082598b {
		if strings.HasSuffix(__obf_44850cb26b65dffd.Name(), __obf_fe4da2b16e24370c) {
			continue
		}
		__obf_2861e3e531577c0e := filepath.Join(__obf_677d23f8777c6fc4.__obf_e57e6de8bca69c4c, __obf_44850cb26b65dffd.Name())
		__obf_5ddb1aeb344ab2e0 = util.PathGzip(__obf_2861e3e531577c0e)
		if __obf_5ddb1aeb344ab2e0 != nil {
			__obf_19e8b1471f655904 = append(__obf_19e8b1471f655904, fmt.Errorf("compress %s: %w", __obf_44850cb26b65dffd.Name(), __obf_5ddb1aeb344ab2e0))
		} else {
			if __obf_5ddb1aeb344ab2e0 = os.Remove(__obf_2861e3e531577c0e); __obf_5ddb1aeb344ab2e0 != nil {
				__obf_19e8b1471f655904 = append(__obf_19e8b1471f655904, fmt.Errorf("remove %s: %w", __obf_2861e3e531577c0e, __obf_5ddb1aeb344ab2e0))
			}
		}
	}

	return errors.Join(__obf_19e8b1471f655904...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_677d23f8777c6fc4 *LogWriter) __obf_8e7a4b8b1e4e9574() ([]__obf_c755abfd19f5b475, error) {
	__obf_f8ae176e5b796c4c, __obf_5ddb1aeb344ab2e0 := os.ReadDir(__obf_677d23f8777c6fc4.__obf_e57e6de8bca69c4c)
	if __obf_5ddb1aeb344ab2e0 != nil {
		return nil, __obf_5ddb1aeb344ab2e0
	}
	var (
		__obf_05f421a410836518                         time.Time
		__obf_b4fd143ae1eedc43                         []__obf_c755abfd19f5b475
		__obf_2a48f2986f9e7da4, __obf_cdd3fd7f3fa49dfd = util.SplitFilename(__obf_677d23f8777c6fc4.__obf_4d53f6f79fd7e19e.LogPath)
	)
	__obf_2a48f2986f9e7da4 += "-"
	for _, __obf_44850cb26b65dffd := range __obf_f8ae176e5b796c4c {
		if __obf_44850cb26b65dffd.IsDir() {
			continue
		}
		if __obf_05f421a410836518, __obf_5ddb1aeb344ab2e0 = __obf_178eb7d79352c341(__obf_b92c039a12d148ca, __obf_44850cb26b65dffd.Name(), __obf_2a48f2986f9e7da4, __obf_cdd3fd7f3fa49dfd); __obf_5ddb1aeb344ab2e0 == nil {
			__obf_b4fd143ae1eedc43 = append(__obf_b4fd143ae1eedc43, __obf_c755abfd19f5b475{__obf_44850cb26b65dffd, __obf_05f421a410836518})
			continue
		}
		if __obf_05f421a410836518, __obf_5ddb1aeb344ab2e0 = __obf_178eb7d79352c341(__obf_b92c039a12d148ca, __obf_44850cb26b65dffd.Name(), __obf_2a48f2986f9e7da4, __obf_cdd3fd7f3fa49dfd+__obf_fe4da2b16e24370c); __obf_5ddb1aeb344ab2e0 == nil {
			__obf_b4fd143ae1eedc43 = append(__obf_b4fd143ae1eedc43, __obf_c755abfd19f5b475{__obf_44850cb26b65dffd, __obf_05f421a410836518})
			continue
		}
	}
	sort.Sort(__obf_eaee547770693572(__obf_b4fd143ae1eedc43))
	return __obf_b4fd143ae1eedc43, nil
}

type __obf_c755abfd19f5b475 struct {
	os.DirEntry
	__obf_1b92f09c365627ad time.Time
}

type __obf_eaee547770693572 []__obf_c755abfd19f5b475

func (__obf_2eecd50c72d1152c __obf_eaee547770693572) Less(__obf_1cc08a9c6ef42302, __obf_744dda69dba433cb int) bool {
	return __obf_2eecd50c72d1152c[__obf_1cc08a9c6ef42302].__obf_1b92f09c365627ad.After(__obf_2eecd50c72d1152c[__obf_744dda69dba433cb].__obf_1b92f09c365627ad)
}

func (__obf_2eecd50c72d1152c __obf_eaee547770693572) Swap(__obf_1cc08a9c6ef42302, __obf_744dda69dba433cb int) {
	__obf_2eecd50c72d1152c[__obf_1cc08a9c6ef42302], __obf_2eecd50c72d1152c[__obf_744dda69dba433cb] = __obf_2eecd50c72d1152c[__obf_744dda69dba433cb], __obf_2eecd50c72d1152c[__obf_1cc08a9c6ef42302]
}

func (__obf_2eecd50c72d1152c __obf_eaee547770693572) Len() int {
	return len(__obf_2eecd50c72d1152c)
}

// timeFromName 从文件名中提取并解析时间戳，根据提供的时间格式、前缀和扩展名进行验证。
// 它会检查文件名是否以指定的前缀开头并以指定的扩展名结尾，然后尝试解析时间戳。
//
// 参数:
//   - timeFormat: 用于解析时间戳的布局字符串 (例如: "2006-01-02")。
//   - filename: 包含时间戳的完整文件名。
//   - prefix: 文件名的预期前缀。
//   - ext: 文件名的预期扩展名。
//
// 返回值:
//   - time.Time: 如果解析成功，返回解析的时间戳。
//   - error: 如果文件名与前缀或扩展名不匹配，或者时间戳无法解析，则返回错误。
func __obf_178eb7d79352c341(__obf_b92c039a12d148ca, __obf_d0b43baab0c40a81, __obf_2a48f2986f9e7da4, __obf_cdd3fd7f3fa49dfd string) (time.Time, error) {
	if !strings.HasPrefix(__obf_d0b43baab0c40a81, __obf_2a48f2986f9e7da4) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_2a48f2986f9e7da4, __obf_d0b43baab0c40a81)
	}
	if !strings.HasSuffix(__obf_d0b43baab0c40a81, __obf_cdd3fd7f3fa49dfd) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_cdd3fd7f3fa49dfd, __obf_d0b43baab0c40a81)
	}
	__obf_f753ebedbba68eff := __obf_d0b43baab0c40a81[len(__obf_2a48f2986f9e7da4) : len(__obf_d0b43baab0c40a81)-len(__obf_cdd3fd7f3fa49dfd)]
	return time.Parse(__obf_b92c039a12d148ca, __obf_f753ebedbba68eff)
}
