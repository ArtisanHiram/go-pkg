// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_10b299a6084cd5a7

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
	__obf_8ca40483052a410a = "2006-01-02_15:04:05.000"
	__obf_4db8fa8f3d59c873 = ".gz"
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
	__obf_67789ce2eb931f27 Option
	__obf_a142fe1a109f3f49 *os.File
	__obf_e853b9057c6d5b00 int64
	__obf_eba0e5c8d96035dc string
	__obf_18ec5a3164d017f4 sync.Mutex
	__obf_62cc36e79dec1d75 sync.Once
	__obf_242f58652c116d22 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_67789ce2eb931f27 Option) *LogWriter {
	__obf_eba0e5c8d96035dc := filepath.Dir(__obf_67789ce2eb931f27.LogPath)
	__obf_67789ce2eb931f27.
		MaxAge = __obf_67789ce2eb931f27.MaxAge * time.Hour * 24
	__obf_67789ce2eb931f27.
		MaxSize = __obf_67789ce2eb931f27.MaxSize * 1024 * 1024
	__obf_4decaaabe0b09b1c := &LogWriter{__obf_67789ce2eb931f27: __obf_67789ce2eb931f27, __obf_eba0e5c8d96035dc: __obf_eba0e5c8d96035dc}

	return __obf_4decaaabe0b09b1c
}

// Write writes content into file.
func (__obf_f04ff4cd068c6a02 *LogWriter) Write(__obf_f95b1d7cdffecebe []byte) (__obf_dbc5cb4fca09f913 int, __obf_88fc930ea296b3ad error) {
	__obf_f04ff4cd068c6a02.__obf_18ec5a3164d017f4.
		Lock()
	defer __obf_f04ff4cd068c6a02.__obf_18ec5a3164d017f4.Unlock()
	__obf_7f12b5626533c9a4 := int64(len(__obf_f95b1d7cdffecebe))

	if __obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49 == nil {
		if __obf_88fc930ea296b3ad = __obf_f04ff4cd068c6a02.__obf_701c9639f0fc0de0(__obf_7f12b5626533c9a4); __obf_88fc930ea296b3ad != nil {
			return
		}
	}
	if __obf_f04ff4cd068c6a02.__obf_5c58a4cae6324385(__obf_7f12b5626533c9a4 + __obf_f04ff4cd068c6a02.__obf_e853b9057c6d5b00) {
		if __obf_88fc930ea296b3ad = __obf_f04ff4cd068c6a02.__obf_25e27b727c8d6b3c(); __obf_88fc930ea296b3ad != nil {
			return
		}
	}
	__obf_dbc5cb4fca09f913, __obf_88fc930ea296b3ad = __obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49.Write(__obf_f95b1d7cdffecebe)
	__obf_f04ff4cd068c6a02.__obf_e853b9057c6d5b00 += int64(__obf_dbc5cb4fca09f913)
	return
}

// Close closes file resource
func (__obf_f04ff4cd068c6a02 *LogWriter) Close() error {
	__obf_f04ff4cd068c6a02.__obf_18ec5a3164d017f4.
		Lock()
	defer __obf_f04ff4cd068c6a02.__obf_18ec5a3164d017f4.Unlock()
	return __obf_f04ff4cd068c6a02.close()
}

func (__obf_f04ff4cd068c6a02 *LogWriter) close() error {
	if __obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49 == nil {
		return nil
	}
	__obf_88fc930ea296b3ad := __obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49.Close()
	if __obf_88fc930ea296b3ad != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_88fc930ea296b3ad)
	}
	__obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49 = nil
	__obf_f04ff4cd068c6a02.__obf_e853b9057c6d5b00 = 0
	return nil
}

func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_5c58a4cae6324385(__obf_e853b9057c6d5b00 int64) bool {
	return __obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.MaxSize > 0 && __obf_e853b9057c6d5b00 > __obf_f04ff4cd068c6a02.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_67789ce2eb931f27.MaxSize
}

func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_701c9639f0fc0de0(__obf_2d54451d2945dc30 int64) error {
	__obf_ebbe957de3c5593d, __obf_88fc930ea296b3ad := os.Stat(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath)
	if os.IsNotExist(__obf_88fc930ea296b3ad) {
		return __obf_f04ff4cd068c6a02.__obf_ee8ef983984f3a5c()
	}
	if __obf_88fc930ea296b3ad != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_88fc930ea296b3ad)
	}
	if __obf_f04ff4cd068c6a02.__obf_5c58a4cae6324385(__obf_ebbe957de3c5593d.Size() + __obf_2d54451d2945dc30) {
		return __obf_f04ff4cd068c6a02.__obf_25e27b727c8d6b3c()
	}
	__obf_a142fe1a109f3f49, __obf_88fc930ea296b3ad := os.OpenFile(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_88fc930ea296b3ad != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_88fc930ea296b3ad)
	}
	__obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49 = __obf_a142fe1a109f3f49
	__obf_f04ff4cd068c6a02.

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
		__obf_e853b9057c6d5b00 = __obf_ebbe957de3c5593d.Size()
	return nil
}

func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_ee8ef983984f3a5c() error {
	if __obf_88fc930ea296b3ad := os.MkdirAll(__obf_f04ff4cd068c6a02.__obf_eba0e5c8d96035dc, os.ModePerm); __obf_88fc930ea296b3ad != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_88fc930ea296b3ad)
	}
	__obf_56f91509125b253b := os.FileMode(0644)
	__obf_ebbe957de3c5593d, __obf_88fc930ea296b3ad := os.Stat(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath)
	if __obf_88fc930ea296b3ad == nil {
		__obf_56f91509125b253b = __obf_ebbe957de3c5593d.Mode()
		__obf_b76bcc3e40966c0a, __obf_f44031ec5e35abd3 := util.SplitFilename(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath)
		__obf_26addbed0cb5bd5b := time.Now().Format(__obf_8ca40483052a410a)
		__obf_22b39bbb73cd8e99 := filepath.Join(__obf_f04ff4cd068c6a02.__obf_eba0e5c8d96035dc, fmt.Sprintf("%s-%s%s", __obf_b76bcc3e40966c0a, __obf_26addbed0cb5bd5b, __obf_f44031ec5e35abd3))

		if __obf_88fc930ea296b3ad = os.Rename(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath, __obf_22b39bbb73cd8e99); __obf_88fc930ea296b3ad != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_22b39bbb73cd8e99, __obf_88fc930ea296b3ad)
		}
	}
	__obf_9caf613daadcc6e3, __obf_88fc930ea296b3ad := os.OpenFile(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_56f91509125b253b)
	if __obf_88fc930ea296b3ad != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath, __obf_88fc930ea296b3ad)
	}
	__obf_f04ff4cd068c6a02.__obf_a142fe1a109f3f49 = __obf_9caf613daadcc6e3
	__obf_f04ff4cd068c6a02.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_e853b9057c6d5b00 = 0
	return nil
}

func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_25e27b727c8d6b3c() error {
	if __obf_88fc930ea296b3ad := __obf_f04ff4cd068c6a02.close(); __obf_88fc930ea296b3ad != nil {
		return __obf_88fc930ea296b3ad
	}
	if __obf_88fc930ea296b3ad := __obf_f04ff4cd068c6a02.__obf_ee8ef983984f3a5c(); __obf_88fc930ea296b3ad != nil {
		return __obf_88fc930ea296b3ad
	}
	__obf_f04ff4cd068c6a02.__obf_62cc36e79dec1d75.
		Do(func() {
			__obf_f04ff4cd068c6a02.__obf_242f58652c116d22 = make(chan bool, 1)
			go func() {
				for range __obf_f04ff4cd068c6a02.__obf_242f58652c116d22 {
					if __obf_f1c585298a3ec26a := __obf_f04ff4cd068c6a02.__obf_ca4f914e28976e59(); __obf_f1c585298a3ec26a != nil {
						println(__obf_f1c585298a3ec26a.Error())
					}
				}
			}()
		})
	select {
	case __obf_f04ff4cd068c6a02.__obf_242f58652c116d22 <- true:
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
func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_ca4f914e28976e59() error {
	if __obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.MaxSize == 0 {
		return nil
	}
	__obf_bcd2bb7860a8fa09, __obf_88fc930ea296b3ad := __obf_f04ff4cd068c6a02.__obf_727e0b614dea7f4d()
	if __obf_88fc930ea296b3ad != nil {
		return __obf_88fc930ea296b3ad
	}
	var __obf_ccd89ac504de58c3, __obf_e3fa4b5e90f684b8 []__obf_a4fab20f59280310
	if __obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.MaxAge > 0 {
		__obf_d7505bd684c4edd8 := time.Now().Add(-1 * __obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.MaxAge)
		for _, __obf_9caf613daadcc6e3 := range __obf_bcd2bb7860a8fa09 {
			if __obf_9caf613daadcc6e3.__obf_726eedcb41ef526e.Before(__obf_d7505bd684c4edd8) {
				__obf_ccd89ac504de58c3 = append(__obf_ccd89ac504de58c3, __obf_9caf613daadcc6e3)
			} else {
				__obf_e3fa4b5e90f684b8 = append(__obf_e3fa4b5e90f684b8, __obf_9caf613daadcc6e3)
			}
		}
	}

	var __obf_487cba220a5f0852 []error
	for _, __obf_9caf613daadcc6e3 := range __obf_ccd89ac504de58c3 {
		if __obf_88fc930ea296b3ad = os.Remove(filepath.Join(__obf_f04ff4cd068c6a02.__obf_eba0e5c8d96035dc, __obf_9caf613daadcc6e3.Name())); __obf_88fc930ea296b3ad != nil {
			__obf_487cba220a5f0852 = append(__obf_487cba220a5f0852, fmt.Errorf("remove %s: %w", __obf_9caf613daadcc6e3.Name(), __obf_88fc930ea296b3ad))
		}
	}
	if len(__obf_487cba220a5f0852) > 0 {
		return errors.Join(__obf_487cba220a5f0852...)
	}

	for _, __obf_9caf613daadcc6e3 := range __obf_e3fa4b5e90f684b8 {
		if strings.HasSuffix(__obf_9caf613daadcc6e3.Name(), __obf_4db8fa8f3d59c873) {
			continue
		}
		__obf_ee1a4f9a119be494 := filepath.Join(__obf_f04ff4cd068c6a02.__obf_eba0e5c8d96035dc, __obf_9caf613daadcc6e3.Name())
		__obf_88fc930ea296b3ad = util.PathGzip(__obf_ee1a4f9a119be494)
		if __obf_88fc930ea296b3ad != nil {
			__obf_487cba220a5f0852 = append(__obf_487cba220a5f0852, fmt.Errorf("compress %s: %w", __obf_9caf613daadcc6e3.Name(), __obf_88fc930ea296b3ad))
		} else {
			if __obf_88fc930ea296b3ad = os.Remove(__obf_ee1a4f9a119be494); __obf_88fc930ea296b3ad != nil {
				__obf_487cba220a5f0852 = append(__obf_487cba220a5f0852, fmt.Errorf("remove %s: %w", __obf_ee1a4f9a119be494, __obf_88fc930ea296b3ad))
			}
		}
	}

	return errors.Join(__obf_487cba220a5f0852...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_f04ff4cd068c6a02 *LogWriter) __obf_727e0b614dea7f4d() ([]__obf_a4fab20f59280310, error) {
	__obf_bcd2bb7860a8fa09, __obf_88fc930ea296b3ad := os.ReadDir(__obf_f04ff4cd068c6a02.__obf_eba0e5c8d96035dc)
	if __obf_88fc930ea296b3ad != nil {
		return nil, __obf_88fc930ea296b3ad
	}
	var (
		__obf_26addbed0cb5bd5b                         time.Time
		__obf_df61cb30e896144d                         []__obf_a4fab20f59280310
		__obf_b76bcc3e40966c0a, __obf_f44031ec5e35abd3 = util.SplitFilename(__obf_f04ff4cd068c6a02.__obf_67789ce2eb931f27.LogPath)
	)
	__obf_b76bcc3e40966c0a += "-"
	for _, __obf_9caf613daadcc6e3 := range __obf_bcd2bb7860a8fa09 {
		if __obf_9caf613daadcc6e3.IsDir() {
			continue
		}
		if __obf_26addbed0cb5bd5b, __obf_88fc930ea296b3ad = __obf_c8bf81cdcd160e09(__obf_8ca40483052a410a, __obf_9caf613daadcc6e3.Name(), __obf_b76bcc3e40966c0a, __obf_f44031ec5e35abd3); __obf_88fc930ea296b3ad == nil {
			__obf_df61cb30e896144d = append(__obf_df61cb30e896144d, __obf_a4fab20f59280310{__obf_9caf613daadcc6e3, __obf_26addbed0cb5bd5b})
			continue
		}
		if __obf_26addbed0cb5bd5b, __obf_88fc930ea296b3ad = __obf_c8bf81cdcd160e09(__obf_8ca40483052a410a, __obf_9caf613daadcc6e3.Name(), __obf_b76bcc3e40966c0a, __obf_f44031ec5e35abd3+__obf_4db8fa8f3d59c873); __obf_88fc930ea296b3ad == nil {
			__obf_df61cb30e896144d = append(__obf_df61cb30e896144d, __obf_a4fab20f59280310{__obf_9caf613daadcc6e3, __obf_26addbed0cb5bd5b})
			continue
		}
	}
	sort.Sort(__obf_0dbd5c06f3a9230e(__obf_df61cb30e896144d))
	return __obf_df61cb30e896144d, nil
}

type __obf_a4fab20f59280310 struct {
	os.DirEntry
	__obf_726eedcb41ef526e time.Time
}

type __obf_0dbd5c06f3a9230e []__obf_a4fab20f59280310

func (__obf_4562e80108455e70 __obf_0dbd5c06f3a9230e) Less(__obf_47c9a29e8b2c582d, __obf_d91802ac36a2c1e5 int) bool {
	return __obf_4562e80108455e70[__obf_47c9a29e8b2c582d].__obf_726eedcb41ef526e.After(__obf_4562e80108455e70[__obf_d91802ac36a2c1e5].__obf_726eedcb41ef526e)
}

func (__obf_4562e80108455e70 __obf_0dbd5c06f3a9230e) Swap(__obf_47c9a29e8b2c582d, __obf_d91802ac36a2c1e5 int) {
	__obf_4562e80108455e70[__obf_47c9a29e8b2c582d], __obf_4562e80108455e70[__obf_d91802ac36a2c1e5] = __obf_4562e80108455e70[__obf_d91802ac36a2c1e5], __obf_4562e80108455e70[__obf_47c9a29e8b2c582d]
}

func (__obf_4562e80108455e70 __obf_0dbd5c06f3a9230e) Len() int {
	return len(__obf_4562e80108455e70)
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
func __obf_c8bf81cdcd160e09(__obf_8ca40483052a410a, __obf_4ce0d54bd043e736, __obf_b76bcc3e40966c0a, __obf_f44031ec5e35abd3 string) (time.Time, error) {
	if !strings.HasPrefix(__obf_4ce0d54bd043e736, __obf_b76bcc3e40966c0a) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_b76bcc3e40966c0a, __obf_4ce0d54bd043e736)
	}
	if !strings.HasSuffix(__obf_4ce0d54bd043e736, __obf_f44031ec5e35abd3) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_f44031ec5e35abd3, __obf_4ce0d54bd043e736)
	}
	__obf_5dc294347e067e22 := __obf_4ce0d54bd043e736[len(__obf_b76bcc3e40966c0a) : len(__obf_4ce0d54bd043e736)-len(__obf_f44031ec5e35abd3)]
	return time.Parse(__obf_8ca40483052a410a, __obf_5dc294347e067e22)
}
