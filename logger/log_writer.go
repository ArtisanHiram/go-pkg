// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_dc6fa34d54e9b538

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
	__obf_5bdd0d5a8c29c7b8 = "2006-01-02_15:04:05.000"
	__obf_780fd34263840cc8 = ".gz"
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
	__obf_6e107c26f1f08d2f Option
	__obf_577762d05139a620 *os.File
	__obf_8bd33f1a30b1e654 int64
	__obf_6eafb07713a9dff2 string
	__obf_876f9019d12e1217 sync.Mutex
	__obf_cd892f78fb27f225 sync.Once
	__obf_2fdb5894e174362f chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_6e107c26f1f08d2f Option) *LogWriter {
	__obf_6eafb07713a9dff2 := filepath.Dir(__obf_6e107c26f1f08d2f.LogPath)

	__obf_6e107c26f1f08d2f.MaxAge = __obf_6e107c26f1f08d2f.MaxAge * time.Hour * 24
	__obf_6e107c26f1f08d2f.MaxSize = __obf_6e107c26f1f08d2f.MaxSize * 1024 * 1024

	__obf_9328a7e8b4a28dbb := &LogWriter{
		__obf_6e107c26f1f08d2f: __obf_6e107c26f1f08d2f,
		__obf_6eafb07713a9dff2: __obf_6eafb07713a9dff2,
	}

	return __obf_9328a7e8b4a28dbb
}

// Write writes content into file.
func (__obf_ccab94b902292a52 *LogWriter) Write(__obf_c26deaf555a0b14a []byte) (__obf_0a01ff4187f6bf59 int, __obf_dba0f9c706c8eb7b error) {
	__obf_ccab94b902292a52.__obf_876f9019d12e1217.Lock()
	defer __obf_ccab94b902292a52.__obf_876f9019d12e1217.Unlock()
	__obf_869bb2277fe861e9 := int64(len(__obf_c26deaf555a0b14a))

	if __obf_ccab94b902292a52.__obf_577762d05139a620 == nil {
		if __obf_dba0f9c706c8eb7b = __obf_ccab94b902292a52.__obf_92996780f2445689(__obf_869bb2277fe861e9); __obf_dba0f9c706c8eb7b != nil {
			return
		}
	}
	if __obf_ccab94b902292a52.__obf_a41a7ec3f8cea1a5(__obf_869bb2277fe861e9 + __obf_ccab94b902292a52.__obf_8bd33f1a30b1e654) {
		if __obf_dba0f9c706c8eb7b = __obf_ccab94b902292a52.__obf_827e4937396472e6(); __obf_dba0f9c706c8eb7b != nil {
			return
		}
	}
	__obf_0a01ff4187f6bf59, __obf_dba0f9c706c8eb7b = __obf_ccab94b902292a52.__obf_577762d05139a620.Write(__obf_c26deaf555a0b14a)
	__obf_ccab94b902292a52.__obf_8bd33f1a30b1e654 += int64(__obf_0a01ff4187f6bf59)
	return
}

// Close closes file resource
func (__obf_ccab94b902292a52 *LogWriter) Close() error {
	__obf_ccab94b902292a52.__obf_876f9019d12e1217.Lock()
	defer __obf_ccab94b902292a52.__obf_876f9019d12e1217.Unlock()
	return __obf_ccab94b902292a52.close()
}

func (__obf_ccab94b902292a52 *LogWriter) close() error {
	if __obf_ccab94b902292a52.__obf_577762d05139a620 == nil {
		return nil
	}
	__obf_dba0f9c706c8eb7b := __obf_ccab94b902292a52.__obf_577762d05139a620.Close()
	if __obf_dba0f9c706c8eb7b != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_dba0f9c706c8eb7b)
	}
	__obf_ccab94b902292a52.__obf_577762d05139a620 = nil
	__obf_ccab94b902292a52.__obf_8bd33f1a30b1e654 = 0
	return nil
}

func (__obf_ccab94b902292a52 *LogWriter) __obf_a41a7ec3f8cea1a5(__obf_8bd33f1a30b1e654 int64) bool {
	return __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.MaxSize > 0 && __obf_8bd33f1a30b1e654 > __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.MaxSize
}

// openFile 负责打开日志文件以供写入。
// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
// 如果文件存在，它会判断是否需要轮换日志文件。
func (__obf_ccab94b902292a52 *LogWriter) __obf_92996780f2445689(__obf_84d9214bedb0e9a0 int64) error {
	__obf_04686acb36088867, __obf_dba0f9c706c8eb7b := os.Stat(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath)
	if os.IsNotExist(__obf_dba0f9c706c8eb7b) {
		return __obf_ccab94b902292a52.__obf_6ffb7b6efa6338c3()
	}
	if __obf_dba0f9c706c8eb7b != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_dba0f9c706c8eb7b)
	}
	if __obf_ccab94b902292a52.__obf_a41a7ec3f8cea1a5(__obf_04686acb36088867.Size() + __obf_84d9214bedb0e9a0) {
		return __obf_ccab94b902292a52.__obf_827e4937396472e6()
	}
	__obf_577762d05139a620, __obf_dba0f9c706c8eb7b := os.OpenFile(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_dba0f9c706c8eb7b != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_dba0f9c706c8eb7b)
	}
	__obf_ccab94b902292a52.__obf_577762d05139a620 = __obf_577762d05139a620
	__obf_ccab94b902292a52.__obf_8bd33f1a30b1e654 = __obf_04686acb36088867.Size()
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
func (__obf_ccab94b902292a52 *LogWriter) __obf_6ffb7b6efa6338c3() error {
	if __obf_dba0f9c706c8eb7b := os.MkdirAll(__obf_ccab94b902292a52.__obf_6eafb07713a9dff2, os.ModePerm); __obf_dba0f9c706c8eb7b != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_dba0f9c706c8eb7b)
	}
	__obf_ef62831884980785 := os.FileMode(0644)
	__obf_04686acb36088867, __obf_dba0f9c706c8eb7b := os.Stat(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath)
	if __obf_dba0f9c706c8eb7b == nil {
		__obf_ef62831884980785 = __obf_04686acb36088867.Mode()
		__obf_6939c8567f170eef, __obf_e23dac99407affec := util.SplitFilename(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath)
		__obf_2a858f5ae01f0cd9 := time.Now().Format(__obf_5bdd0d5a8c29c7b8)
		__obf_94fdf3945bc30d09 := filepath.Join(__obf_ccab94b902292a52.__obf_6eafb07713a9dff2, fmt.Sprintf("%s-%s%s", __obf_6939c8567f170eef, __obf_2a858f5ae01f0cd9, __obf_e23dac99407affec))

		if __obf_dba0f9c706c8eb7b = os.Rename(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath, __obf_94fdf3945bc30d09); __obf_dba0f9c706c8eb7b != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_94fdf3945bc30d09, __obf_dba0f9c706c8eb7b)
		}
	}
	__obf_c87ea26f026dba11, __obf_dba0f9c706c8eb7b := os.OpenFile(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_ef62831884980785)
	if __obf_dba0f9c706c8eb7b != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath, __obf_dba0f9c706c8eb7b)
	}
	__obf_ccab94b902292a52.__obf_577762d05139a620 = __obf_c87ea26f026dba11
	__obf_ccab94b902292a52.__obf_8bd33f1a30b1e654 = 0
	return nil
}

// rotate 处理 LogWriter 的日志文件轮换过程。
// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
//
// 如果关闭当前文件或打开新文件失败，则返回错误。

func (__obf_ccab94b902292a52 *LogWriter) __obf_827e4937396472e6() error {
	if __obf_dba0f9c706c8eb7b := __obf_ccab94b902292a52.close(); __obf_dba0f9c706c8eb7b != nil {
		return __obf_dba0f9c706c8eb7b
	}
	if __obf_dba0f9c706c8eb7b := __obf_ccab94b902292a52.__obf_6ffb7b6efa6338c3(); __obf_dba0f9c706c8eb7b != nil {
		return __obf_dba0f9c706c8eb7b
	}
	__obf_ccab94b902292a52.__obf_cd892f78fb27f225.Do(func() {
		__obf_ccab94b902292a52.__obf_2fdb5894e174362f = make(chan bool, 1)
		go func() {
			for range __obf_ccab94b902292a52.__obf_2fdb5894e174362f {
				if __obf_4f5e43ceba35278c := __obf_ccab94b902292a52.__obf_dd24d145747a57f0(); __obf_4f5e43ceba35278c != nil {
					println(__obf_4f5e43ceba35278c.Error())
				}
			}
		}()
	})
	select {
	case __obf_ccab94b902292a52.__obf_2fdb5894e174362f <- true:
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
func (__obf_ccab94b902292a52 *LogWriter) __obf_dd24d145747a57f0() error {
	if __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.MaxSize == 0 {
		return nil
	}
	__obf_efa67a0616d5efbd, __obf_dba0f9c706c8eb7b := __obf_ccab94b902292a52.__obf_e90c5b6ef9bd6fb7()
	if __obf_dba0f9c706c8eb7b != nil {
		return __obf_dba0f9c706c8eb7b
	}
	var __obf_b9949d9f57d1ee35, __obf_455a2eed75c50002 []__obf_5dfa4ff9e80af9c7
	if __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.MaxAge > 0 {
		__obf_7627ada88875e90b := time.Now().Add(-1 * __obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.MaxAge)
		for _, __obf_c87ea26f026dba11 := range __obf_efa67a0616d5efbd {
			if __obf_c87ea26f026dba11.__obf_1b8bc98d1f90165d.Before(__obf_7627ada88875e90b) {
				__obf_b9949d9f57d1ee35 = append(__obf_b9949d9f57d1ee35, __obf_c87ea26f026dba11)
			} else {
				__obf_455a2eed75c50002 = append(__obf_455a2eed75c50002, __obf_c87ea26f026dba11)
			}
		}
	}

	var __obf_8f956a99aabcd8ab []error
	for _, __obf_c87ea26f026dba11 := range __obf_b9949d9f57d1ee35 {
		if __obf_dba0f9c706c8eb7b = os.Remove(filepath.Join(__obf_ccab94b902292a52.__obf_6eafb07713a9dff2, __obf_c87ea26f026dba11.Name())); __obf_dba0f9c706c8eb7b != nil {
			__obf_8f956a99aabcd8ab = append(__obf_8f956a99aabcd8ab, fmt.Errorf("remove %s: %w", __obf_c87ea26f026dba11.Name(), __obf_dba0f9c706c8eb7b))
		}
	}
	if len(__obf_8f956a99aabcd8ab) > 0 {
		return errors.Join(__obf_8f956a99aabcd8ab...)
	}

	for _, __obf_c87ea26f026dba11 := range __obf_455a2eed75c50002 {
		if strings.HasSuffix(__obf_c87ea26f026dba11.Name(), __obf_780fd34263840cc8) {
			continue
		}
		__obf_2d4deb158a47d294 := filepath.Join(__obf_ccab94b902292a52.__obf_6eafb07713a9dff2, __obf_c87ea26f026dba11.Name())

		__obf_dba0f9c706c8eb7b = util.PathGzip(__obf_2d4deb158a47d294)
		if __obf_dba0f9c706c8eb7b != nil {
			__obf_8f956a99aabcd8ab = append(__obf_8f956a99aabcd8ab, fmt.Errorf("compress %s: %w", __obf_c87ea26f026dba11.Name(), __obf_dba0f9c706c8eb7b))
		} else {
			if __obf_dba0f9c706c8eb7b = os.Remove(__obf_2d4deb158a47d294); __obf_dba0f9c706c8eb7b != nil {
				__obf_8f956a99aabcd8ab = append(__obf_8f956a99aabcd8ab, fmt.Errorf("remove %s: %w", __obf_2d4deb158a47d294, __obf_dba0f9c706c8eb7b))
			}
		}
	}

	return errors.Join(__obf_8f956a99aabcd8ab...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_ccab94b902292a52 *LogWriter) __obf_e90c5b6ef9bd6fb7() ([]__obf_5dfa4ff9e80af9c7, error) {
	__obf_efa67a0616d5efbd, __obf_dba0f9c706c8eb7b := os.ReadDir(__obf_ccab94b902292a52.__obf_6eafb07713a9dff2)
	if __obf_dba0f9c706c8eb7b != nil {
		return nil, __obf_dba0f9c706c8eb7b
	}
	var (
		__obf_2a858f5ae01f0cd9                         time.Time
		__obf_17e377e63baeb79b                         []__obf_5dfa4ff9e80af9c7
		__obf_6939c8567f170eef, __obf_e23dac99407affec = util.SplitFilename(__obf_ccab94b902292a52.__obf_6e107c26f1f08d2f.LogPath)
	)
	__obf_6939c8567f170eef += "-"
	for _, __obf_c87ea26f026dba11 := range __obf_efa67a0616d5efbd {
		if __obf_c87ea26f026dba11.IsDir() {
			continue
		}
		if __obf_2a858f5ae01f0cd9, __obf_dba0f9c706c8eb7b = __obf_b684ad32989aa3c6(__obf_5bdd0d5a8c29c7b8, __obf_c87ea26f026dba11.Name(), __obf_6939c8567f170eef, __obf_e23dac99407affec); __obf_dba0f9c706c8eb7b == nil {
			__obf_17e377e63baeb79b = append(__obf_17e377e63baeb79b, __obf_5dfa4ff9e80af9c7{__obf_c87ea26f026dba11, __obf_2a858f5ae01f0cd9})
			continue
		}
		if __obf_2a858f5ae01f0cd9, __obf_dba0f9c706c8eb7b = __obf_b684ad32989aa3c6(__obf_5bdd0d5a8c29c7b8, __obf_c87ea26f026dba11.Name(), __obf_6939c8567f170eef, __obf_e23dac99407affec+__obf_780fd34263840cc8); __obf_dba0f9c706c8eb7b == nil {
			__obf_17e377e63baeb79b = append(__obf_17e377e63baeb79b, __obf_5dfa4ff9e80af9c7{__obf_c87ea26f026dba11, __obf_2a858f5ae01f0cd9})
			continue
		}
	}
	sort.Sort(__obf_0147d9cb927622cb(__obf_17e377e63baeb79b))
	return __obf_17e377e63baeb79b, nil
}

type __obf_5dfa4ff9e80af9c7 struct {
	os.DirEntry
	__obf_1b8bc98d1f90165d time.Time
}

type __obf_0147d9cb927622cb []__obf_5dfa4ff9e80af9c7

func (__obf_a014da64af2e944e __obf_0147d9cb927622cb) Less(__obf_75adab83c3e7d65f, __obf_226b003c21f0c9de int) bool {
	return __obf_a014da64af2e944e[__obf_75adab83c3e7d65f].__obf_1b8bc98d1f90165d.After(__obf_a014da64af2e944e[__obf_226b003c21f0c9de].__obf_1b8bc98d1f90165d)
}

func (__obf_a014da64af2e944e __obf_0147d9cb927622cb) Swap(__obf_75adab83c3e7d65f, __obf_226b003c21f0c9de int) {
	__obf_a014da64af2e944e[__obf_75adab83c3e7d65f], __obf_a014da64af2e944e[__obf_226b003c21f0c9de] = __obf_a014da64af2e944e[__obf_226b003c21f0c9de], __obf_a014da64af2e944e[__obf_75adab83c3e7d65f]
}

func (__obf_a014da64af2e944e __obf_0147d9cb927622cb) Len() int {
	return len(__obf_a014da64af2e944e)
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
func __obf_b684ad32989aa3c6(__obf_5bdd0d5a8c29c7b8, __obf_2371e5b7c83d51ac, __obf_6939c8567f170eef, __obf_e23dac99407affec string) (time.Time, error) {
	if !strings.HasPrefix(__obf_2371e5b7c83d51ac, __obf_6939c8567f170eef) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_6939c8567f170eef, __obf_2371e5b7c83d51ac)
	}
	if !strings.HasSuffix(__obf_2371e5b7c83d51ac, __obf_e23dac99407affec) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_e23dac99407affec, __obf_2371e5b7c83d51ac)
	}
	__obf_1d034721c66648b3 := __obf_2371e5b7c83d51ac[len(__obf_6939c8567f170eef) : len(__obf_2371e5b7c83d51ac)-len(__obf_e23dac99407affec)]
	return time.Parse(__obf_5bdd0d5a8c29c7b8, __obf_1d034721c66648b3)
}
