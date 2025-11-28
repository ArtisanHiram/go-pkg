// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_1844c112ef6a527c

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
	__obf_57e6b916c9d46b69 = "2006-01-02_15:04:05.000"
	__obf_ecce23ca316f5b3f = ".gz"
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
	__obf_ae3ef9e53d756ab3 Option
	__obf_982af539c63d8b14 *os.File
	__obf_b104ef1bac148d8c int64
	__obf_1b48b36d46480a4f string
	__obf_5b21bb8b3485887c sync.Mutex
	__obf_c57d6874b5f2bb22 sync.Once
	__obf_1bcf8ae2cf890e24 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_ae3ef9e53d756ab3 Option) *LogWriter {
	__obf_1b48b36d46480a4f := filepath.Dir(__obf_ae3ef9e53d756ab3.LogPath)

	__obf_ae3ef9e53d756ab3.MaxAge = __obf_ae3ef9e53d756ab3.MaxAge * time.Hour * 24
	__obf_ae3ef9e53d756ab3.MaxSize = __obf_ae3ef9e53d756ab3.MaxSize * 1024 * 1024

	__obf_f3c45b5a2edabde8 := &LogWriter{
		__obf_ae3ef9e53d756ab3: __obf_ae3ef9e53d756ab3,
		__obf_1b48b36d46480a4f: __obf_1b48b36d46480a4f,
	}

	return __obf_f3c45b5a2edabde8
}

// Write writes content into file.
func (__obf_289021c5f8702f7a *LogWriter) Write(__obf_5262e7dd22f90800 []byte) (__obf_233064790c272bbc int, __obf_e2bff675d7b20762 error) {
	__obf_289021c5f8702f7a.__obf_5b21bb8b3485887c.Lock()
	defer __obf_289021c5f8702f7a.__obf_5b21bb8b3485887c.Unlock()
	__obf_de783213917fc07a := int64(len(__obf_5262e7dd22f90800))

	if __obf_289021c5f8702f7a.__obf_982af539c63d8b14 == nil {
		if __obf_e2bff675d7b20762 = __obf_289021c5f8702f7a.__obf_f340b089a08ea68d(__obf_de783213917fc07a); __obf_e2bff675d7b20762 != nil {
			return
		}
	}
	if __obf_289021c5f8702f7a.__obf_fcc32619c3f9ac87(__obf_de783213917fc07a + __obf_289021c5f8702f7a.__obf_b104ef1bac148d8c) {
		if __obf_e2bff675d7b20762 = __obf_289021c5f8702f7a.__obf_6f8794cb9cb8ec13(); __obf_e2bff675d7b20762 != nil {
			return
		}
	}
	__obf_233064790c272bbc, __obf_e2bff675d7b20762 = __obf_289021c5f8702f7a.__obf_982af539c63d8b14.Write(__obf_5262e7dd22f90800)
	__obf_289021c5f8702f7a.__obf_b104ef1bac148d8c += int64(__obf_233064790c272bbc)
	return
}

// Close closes file resource
func (__obf_289021c5f8702f7a *LogWriter) Close() error {
	__obf_289021c5f8702f7a.__obf_5b21bb8b3485887c.Lock()
	defer __obf_289021c5f8702f7a.__obf_5b21bb8b3485887c.Unlock()
	return __obf_289021c5f8702f7a.close()
}

func (__obf_289021c5f8702f7a *LogWriter) close() error {
	if __obf_289021c5f8702f7a.__obf_982af539c63d8b14 == nil {
		return nil
	}
	__obf_e2bff675d7b20762 := __obf_289021c5f8702f7a.__obf_982af539c63d8b14.Close()
	if __obf_e2bff675d7b20762 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_e2bff675d7b20762)
	}
	__obf_289021c5f8702f7a.__obf_982af539c63d8b14 = nil
	__obf_289021c5f8702f7a.__obf_b104ef1bac148d8c = 0
	return nil
}

func (__obf_289021c5f8702f7a *LogWriter) __obf_fcc32619c3f9ac87(__obf_b104ef1bac148d8c int64) bool {
	return __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.MaxSize > 0 && __obf_b104ef1bac148d8c > __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.MaxSize
}

// openFile 负责打开日志文件以供写入。
// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
// 如果文件存在，它会判断是否需要轮换日志文件。
func (__obf_289021c5f8702f7a *LogWriter) __obf_f340b089a08ea68d(__obf_885dd76a78baeb43 int64) error {
	__obf_d84a91629e5f6495, __obf_e2bff675d7b20762 := os.Stat(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath)
	if os.IsNotExist(__obf_e2bff675d7b20762) {
		return __obf_289021c5f8702f7a.__obf_793cd89287a3a6a9()
	}
	if __obf_e2bff675d7b20762 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_e2bff675d7b20762)
	}
	if __obf_289021c5f8702f7a.__obf_fcc32619c3f9ac87(__obf_d84a91629e5f6495.Size() + __obf_885dd76a78baeb43) {
		return __obf_289021c5f8702f7a.__obf_6f8794cb9cb8ec13()
	}
	__obf_982af539c63d8b14, __obf_e2bff675d7b20762 := os.OpenFile(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_e2bff675d7b20762 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_e2bff675d7b20762)
	}
	__obf_289021c5f8702f7a.__obf_982af539c63d8b14 = __obf_982af539c63d8b14
	__obf_289021c5f8702f7a.__obf_b104ef1bac148d8c = __obf_d84a91629e5f6495.Size()
	return nil
}

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
func (__obf_289021c5f8702f7a *LogWriter) __obf_793cd89287a3a6a9() error {
	if __obf_e2bff675d7b20762 := os.MkdirAll(__obf_289021c5f8702f7a.__obf_1b48b36d46480a4f, os.ModePerm); __obf_e2bff675d7b20762 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_e2bff675d7b20762)
	}
	__obf_051faa3fadb1ebd9 := os.FileMode(0644)
	__obf_d84a91629e5f6495, __obf_e2bff675d7b20762 := os.Stat(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath)
	if __obf_e2bff675d7b20762 == nil {
		__obf_051faa3fadb1ebd9 = __obf_d84a91629e5f6495.Mode()
		__obf_e1839ea9ae7ed391, __obf_64a6b917ff67ba1a := util.SplitFilename(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath)
		__obf_4b3e6a9a392e8396 := time.Now().Format(__obf_57e6b916c9d46b69)
		__obf_1737a1c8b5d2d81b := filepath.Join(__obf_289021c5f8702f7a.__obf_1b48b36d46480a4f, fmt.Sprintf("%s-%s%s", __obf_e1839ea9ae7ed391, __obf_4b3e6a9a392e8396, __obf_64a6b917ff67ba1a))

		if __obf_e2bff675d7b20762 = os.Rename(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath, __obf_1737a1c8b5d2d81b); __obf_e2bff675d7b20762 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_1737a1c8b5d2d81b, __obf_e2bff675d7b20762)
		}
	}
	__obf_4707931c0b678e4e, __obf_e2bff675d7b20762 := os.OpenFile(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_051faa3fadb1ebd9)
	if __obf_e2bff675d7b20762 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath, __obf_e2bff675d7b20762)
	}
	__obf_289021c5f8702f7a.__obf_982af539c63d8b14 = __obf_4707931c0b678e4e
	__obf_289021c5f8702f7a.__obf_b104ef1bac148d8c = 0
	return nil
}

// rotate 处理 LogWriter 的日志文件轮换过程。
// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
//
// 如果关闭当前文件或打开新文件失败，则返回错误。

func (__obf_289021c5f8702f7a *LogWriter) __obf_6f8794cb9cb8ec13() error {
	if __obf_e2bff675d7b20762 := __obf_289021c5f8702f7a.close(); __obf_e2bff675d7b20762 != nil {
		return __obf_e2bff675d7b20762
	}
	if __obf_e2bff675d7b20762 := __obf_289021c5f8702f7a.__obf_793cd89287a3a6a9(); __obf_e2bff675d7b20762 != nil {
		return __obf_e2bff675d7b20762
	}
	__obf_289021c5f8702f7a.__obf_c57d6874b5f2bb22.Do(func() {
		__obf_289021c5f8702f7a.__obf_1bcf8ae2cf890e24 = make(chan bool, 1)
		go func() {
			for range __obf_289021c5f8702f7a.__obf_1bcf8ae2cf890e24 {
				if __obf_56f11d83fde1bbf5 := __obf_289021c5f8702f7a.__obf_ac56a4f6ca6ae1b5(); __obf_56f11d83fde1bbf5 != nil {
					println(__obf_56f11d83fde1bbf5.Error())
				}
			}
		}()
	})
	select {
	case __obf_289021c5f8702f7a.__obf_1bcf8ae2cf890e24 <- true:
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
func (__obf_289021c5f8702f7a *LogWriter) __obf_ac56a4f6ca6ae1b5() error {
	if __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.MaxSize == 0 {
		return nil
	}
	__obf_4f8974523ebd2270, __obf_e2bff675d7b20762 := __obf_289021c5f8702f7a.__obf_21851e6eac336c68()
	if __obf_e2bff675d7b20762 != nil {
		return __obf_e2bff675d7b20762
	}
	var __obf_a71ee20fb8c2458e, __obf_e164e43c890642ee []__obf_ebe2c587fb12017a
	if __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.MaxAge > 0 {
		__obf_eefdaf6190729ae1 := time.Now().Add(-1 * __obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.MaxAge)
		for _, __obf_4707931c0b678e4e := range __obf_4f8974523ebd2270 {
			if __obf_4707931c0b678e4e.__obf_84ea61d500bad3a5.Before(__obf_eefdaf6190729ae1) {
				__obf_a71ee20fb8c2458e = append(__obf_a71ee20fb8c2458e, __obf_4707931c0b678e4e)
			} else {
				__obf_e164e43c890642ee = append(__obf_e164e43c890642ee, __obf_4707931c0b678e4e)
			}
		}
	}

	var __obf_5b75e4fe12a5f8fa []error
	for _, __obf_4707931c0b678e4e := range __obf_a71ee20fb8c2458e {
		if __obf_e2bff675d7b20762 = os.Remove(filepath.Join(__obf_289021c5f8702f7a.__obf_1b48b36d46480a4f, __obf_4707931c0b678e4e.Name())); __obf_e2bff675d7b20762 != nil {
			__obf_5b75e4fe12a5f8fa = append(__obf_5b75e4fe12a5f8fa, fmt.Errorf("remove %s: %w", __obf_4707931c0b678e4e.Name(), __obf_e2bff675d7b20762))
		}
	}
	if len(__obf_5b75e4fe12a5f8fa) > 0 {
		return errors.Join(__obf_5b75e4fe12a5f8fa...)
	}

	for _, __obf_4707931c0b678e4e := range __obf_e164e43c890642ee {
		if strings.HasSuffix(__obf_4707931c0b678e4e.Name(), __obf_ecce23ca316f5b3f) {
			continue
		}
		__obf_bd811118c19c3ac3 := filepath.Join(__obf_289021c5f8702f7a.__obf_1b48b36d46480a4f, __obf_4707931c0b678e4e.Name())

		__obf_e2bff675d7b20762 = util.PathGzip(__obf_bd811118c19c3ac3)
		if __obf_e2bff675d7b20762 != nil {
			__obf_5b75e4fe12a5f8fa = append(__obf_5b75e4fe12a5f8fa, fmt.Errorf("compress %s: %w", __obf_4707931c0b678e4e.Name(), __obf_e2bff675d7b20762))
		} else {
			if __obf_e2bff675d7b20762 = os.Remove(__obf_bd811118c19c3ac3); __obf_e2bff675d7b20762 != nil {
				__obf_5b75e4fe12a5f8fa = append(__obf_5b75e4fe12a5f8fa, fmt.Errorf("remove %s: %w", __obf_bd811118c19c3ac3, __obf_e2bff675d7b20762))
			}
		}
	}

	return errors.Join(__obf_5b75e4fe12a5f8fa...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_289021c5f8702f7a *LogWriter) __obf_21851e6eac336c68() ([]__obf_ebe2c587fb12017a, error) {
	__obf_4f8974523ebd2270, __obf_e2bff675d7b20762 := os.ReadDir(__obf_289021c5f8702f7a.__obf_1b48b36d46480a4f)
	if __obf_e2bff675d7b20762 != nil {
		return nil, __obf_e2bff675d7b20762
	}
	var (
		__obf_4b3e6a9a392e8396                         time.Time
		__obf_156ac46b5ed5137a                         []__obf_ebe2c587fb12017a
		__obf_e1839ea9ae7ed391, __obf_64a6b917ff67ba1a = util.SplitFilename(__obf_289021c5f8702f7a.__obf_ae3ef9e53d756ab3.LogPath)
	)
	__obf_e1839ea9ae7ed391 += "-"
	for _, __obf_4707931c0b678e4e := range __obf_4f8974523ebd2270 {
		if __obf_4707931c0b678e4e.IsDir() {
			continue
		}
		if __obf_4b3e6a9a392e8396, __obf_e2bff675d7b20762 = __obf_7ae2d898db8b43fb(__obf_57e6b916c9d46b69, __obf_4707931c0b678e4e.Name(), __obf_e1839ea9ae7ed391, __obf_64a6b917ff67ba1a); __obf_e2bff675d7b20762 == nil {
			__obf_156ac46b5ed5137a = append(__obf_156ac46b5ed5137a, __obf_ebe2c587fb12017a{__obf_4707931c0b678e4e, __obf_4b3e6a9a392e8396})
			continue
		}
		if __obf_4b3e6a9a392e8396, __obf_e2bff675d7b20762 = __obf_7ae2d898db8b43fb(__obf_57e6b916c9d46b69, __obf_4707931c0b678e4e.Name(), __obf_e1839ea9ae7ed391, __obf_64a6b917ff67ba1a+__obf_ecce23ca316f5b3f); __obf_e2bff675d7b20762 == nil {
			__obf_156ac46b5ed5137a = append(__obf_156ac46b5ed5137a, __obf_ebe2c587fb12017a{__obf_4707931c0b678e4e, __obf_4b3e6a9a392e8396})
			continue
		}
	}
	sort.Sort(__obf_067827d89e88322b(__obf_156ac46b5ed5137a))
	return __obf_156ac46b5ed5137a, nil
}

type __obf_ebe2c587fb12017a struct {
	os.DirEntry
	__obf_84ea61d500bad3a5 time.Time
}

type __obf_067827d89e88322b []__obf_ebe2c587fb12017a

func (__obf_421b42a519f9fb81 __obf_067827d89e88322b) Less(__obf_edd96f48d0e078e3, __obf_7bc8ac7e30e6aa42 int) bool {
	return __obf_421b42a519f9fb81[__obf_edd96f48d0e078e3].__obf_84ea61d500bad3a5.After(__obf_421b42a519f9fb81[__obf_7bc8ac7e30e6aa42].__obf_84ea61d500bad3a5)
}

func (__obf_421b42a519f9fb81 __obf_067827d89e88322b) Swap(__obf_edd96f48d0e078e3, __obf_7bc8ac7e30e6aa42 int) {
	__obf_421b42a519f9fb81[__obf_edd96f48d0e078e3], __obf_421b42a519f9fb81[__obf_7bc8ac7e30e6aa42] = __obf_421b42a519f9fb81[__obf_7bc8ac7e30e6aa42], __obf_421b42a519f9fb81[__obf_edd96f48d0e078e3]
}

func (__obf_421b42a519f9fb81 __obf_067827d89e88322b) Len() int {
	return len(__obf_421b42a519f9fb81)
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
func __obf_7ae2d898db8b43fb(__obf_57e6b916c9d46b69, __obf_c1cff5aca35a40e4, __obf_e1839ea9ae7ed391, __obf_64a6b917ff67ba1a string) (time.Time, error) {
	if !strings.HasPrefix(__obf_c1cff5aca35a40e4, __obf_e1839ea9ae7ed391) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_e1839ea9ae7ed391, __obf_c1cff5aca35a40e4)
	}
	if !strings.HasSuffix(__obf_c1cff5aca35a40e4, __obf_64a6b917ff67ba1a) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_64a6b917ff67ba1a, __obf_c1cff5aca35a40e4)
	}
	__obf_3cc119deebb9445c := __obf_c1cff5aca35a40e4[len(__obf_e1839ea9ae7ed391) : len(__obf_c1cff5aca35a40e4)-len(__obf_64a6b917ff67ba1a)]
	return time.Parse(__obf_57e6b916c9d46b69, __obf_3cc119deebb9445c)
}
