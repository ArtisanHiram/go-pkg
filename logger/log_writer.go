// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_0e25611156600e69

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
	__obf_acbbc716f349af37 = "2006-01-02_15:04:05.000"
	__obf_3823bd60ad7ed68a = ".gz"
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
	__obf_dd97462f5f12dd59 Option
	__obf_3ff7d69d4eeecb80 *os.File
	__obf_35d390460e4a34f5 int64
	__obf_ff9dadb72447a062 string
	__obf_3106f6c8095427b6 sync.Mutex
	__obf_25dc23699a5da108 sync.Once
	__obf_d08a1e7c1f3b5268 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_dd97462f5f12dd59 Option) *LogWriter {
	__obf_ff9dadb72447a062 := filepath.Dir(__obf_dd97462f5f12dd59.LogPath)

	__obf_dd97462f5f12dd59.MaxAge = __obf_dd97462f5f12dd59.MaxAge * time.Hour * 24
	__obf_dd97462f5f12dd59.MaxSize = __obf_dd97462f5f12dd59.MaxSize * 1024 * 1024

	__obf_875e5a4ecd76430a := &LogWriter{
		__obf_dd97462f5f12dd59: __obf_dd97462f5f12dd59,
		__obf_ff9dadb72447a062: __obf_ff9dadb72447a062,
	}

	return __obf_875e5a4ecd76430a
}

// Write writes content into file.
func (__obf_b151e3f928f75e3e *LogWriter) Write(__obf_aee9b408eb21f4ba []byte) (__obf_8ae6b19a31dd18fc int, __obf_881b15aa6c554968 error) {
	__obf_b151e3f928f75e3e.__obf_3106f6c8095427b6.Lock()
	defer __obf_b151e3f928f75e3e.__obf_3106f6c8095427b6.Unlock()
	__obf_c7df0c3a1f6f0795 := int64(len(__obf_aee9b408eb21f4ba))

	if __obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80 == nil {
		if __obf_881b15aa6c554968 = __obf_b151e3f928f75e3e.__obf_e944055c5b590f41(__obf_c7df0c3a1f6f0795); __obf_881b15aa6c554968 != nil {
			return
		}
	}
	if __obf_b151e3f928f75e3e.__obf_5e977ac2cd6c542c(__obf_c7df0c3a1f6f0795 + __obf_b151e3f928f75e3e.__obf_35d390460e4a34f5) {
		if __obf_881b15aa6c554968 = __obf_b151e3f928f75e3e.__obf_61b9cdcc0f6b78d6(); __obf_881b15aa6c554968 != nil {
			return
		}
	}
	__obf_8ae6b19a31dd18fc, __obf_881b15aa6c554968 = __obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80.Write(__obf_aee9b408eb21f4ba)
	__obf_b151e3f928f75e3e.__obf_35d390460e4a34f5 += int64(__obf_8ae6b19a31dd18fc)
	return
}

// Close closes file resource
func (__obf_b151e3f928f75e3e *LogWriter) Close() error {
	__obf_b151e3f928f75e3e.__obf_3106f6c8095427b6.Lock()
	defer __obf_b151e3f928f75e3e.__obf_3106f6c8095427b6.Unlock()
	return __obf_b151e3f928f75e3e.close()
}

func (__obf_b151e3f928f75e3e *LogWriter) close() error {
	if __obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80 == nil {
		return nil
	}
	__obf_881b15aa6c554968 := __obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80.Close()
	if __obf_881b15aa6c554968 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_881b15aa6c554968)
	}
	__obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80 = nil
	__obf_b151e3f928f75e3e.__obf_35d390460e4a34f5 = 0
	return nil
}

func (__obf_b151e3f928f75e3e *LogWriter) __obf_5e977ac2cd6c542c(__obf_35d390460e4a34f5 int64) bool {
	return __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.MaxSize > 0 && __obf_35d390460e4a34f5 > __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.MaxSize
}

// openFile 负责打开日志文件以供写入。
// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
// 如果文件存在，它会判断是否需要轮换日志文件。
func (__obf_b151e3f928f75e3e *LogWriter) __obf_e944055c5b590f41(__obf_4f0e530b0e6c8d65 int64) error {
	__obf_5293bbca45f9019a, __obf_881b15aa6c554968 := os.Stat(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath)
	if os.IsNotExist(__obf_881b15aa6c554968) {
		return __obf_b151e3f928f75e3e.__obf_13c8760309abf263()
	}
	if __obf_881b15aa6c554968 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_881b15aa6c554968)
	}
	if __obf_b151e3f928f75e3e.__obf_5e977ac2cd6c542c(__obf_5293bbca45f9019a.Size() + __obf_4f0e530b0e6c8d65) {
		return __obf_b151e3f928f75e3e.__obf_61b9cdcc0f6b78d6()
	}
	__obf_3ff7d69d4eeecb80, __obf_881b15aa6c554968 := os.OpenFile(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_881b15aa6c554968 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_881b15aa6c554968)
	}
	__obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80 = __obf_3ff7d69d4eeecb80
	__obf_b151e3f928f75e3e.__obf_35d390460e4a34f5 = __obf_5293bbca45f9019a.Size()
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
func (__obf_b151e3f928f75e3e *LogWriter) __obf_13c8760309abf263() error {
	if __obf_881b15aa6c554968 := os.MkdirAll(__obf_b151e3f928f75e3e.__obf_ff9dadb72447a062, os.ModePerm); __obf_881b15aa6c554968 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_881b15aa6c554968)
	}
	__obf_9d9cbd1acbedb3e0 := os.FileMode(0644)
	__obf_5293bbca45f9019a, __obf_881b15aa6c554968 := os.Stat(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath)
	if __obf_881b15aa6c554968 == nil {
		__obf_9d9cbd1acbedb3e0 = __obf_5293bbca45f9019a.Mode()
		__obf_45c7627a853f8d21, __obf_b183719997b8fdcd := util.SplitFilename(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath)
		__obf_1679b218c485ac30 := time.Now().Format(__obf_acbbc716f349af37)
		__obf_ea2e4e785701b618 := filepath.Join(__obf_b151e3f928f75e3e.__obf_ff9dadb72447a062, fmt.Sprintf("%s-%s%s", __obf_45c7627a853f8d21, __obf_1679b218c485ac30, __obf_b183719997b8fdcd))

		if __obf_881b15aa6c554968 = os.Rename(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath, __obf_ea2e4e785701b618); __obf_881b15aa6c554968 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_ea2e4e785701b618, __obf_881b15aa6c554968)
		}
	}
	__obf_22504fe46317cdf1, __obf_881b15aa6c554968 := os.OpenFile(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_9d9cbd1acbedb3e0)
	if __obf_881b15aa6c554968 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath, __obf_881b15aa6c554968)
	}
	__obf_b151e3f928f75e3e.__obf_3ff7d69d4eeecb80 = __obf_22504fe46317cdf1
	__obf_b151e3f928f75e3e.__obf_35d390460e4a34f5 = 0
	return nil
}

// rotate 处理 LogWriter 的日志文件轮换过程。
// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
//
// 如果关闭当前文件或打开新文件失败，则返回错误。

func (__obf_b151e3f928f75e3e *LogWriter) __obf_61b9cdcc0f6b78d6() error {
	if __obf_881b15aa6c554968 := __obf_b151e3f928f75e3e.close(); __obf_881b15aa6c554968 != nil {
		return __obf_881b15aa6c554968
	}
	if __obf_881b15aa6c554968 := __obf_b151e3f928f75e3e.__obf_13c8760309abf263(); __obf_881b15aa6c554968 != nil {
		return __obf_881b15aa6c554968
	}
	__obf_b151e3f928f75e3e.__obf_25dc23699a5da108.Do(func() {
		__obf_b151e3f928f75e3e.__obf_d08a1e7c1f3b5268 = make(chan bool, 1)
		go func() {
			for range __obf_b151e3f928f75e3e.__obf_d08a1e7c1f3b5268 {
				if __obf_0d5942aac14b8dcc := __obf_b151e3f928f75e3e.__obf_f34188ef384eb959(); __obf_0d5942aac14b8dcc != nil {
					println(__obf_0d5942aac14b8dcc.Error())
				}
			}
		}()
	})
	select {
	case __obf_b151e3f928f75e3e.__obf_d08a1e7c1f3b5268 <- true:
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
func (__obf_b151e3f928f75e3e *LogWriter) __obf_f34188ef384eb959() error {
	if __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.MaxSize == 0 {
		return nil
	}
	__obf_1bff92203c5f8783, __obf_881b15aa6c554968 := __obf_b151e3f928f75e3e.__obf_a7d2a17e330dc9d2()
	if __obf_881b15aa6c554968 != nil {
		return __obf_881b15aa6c554968
	}
	var __obf_7b2621889262be2a, __obf_af319cef428666be []__obf_2e43a57249cfa23c
	if __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.MaxAge > 0 {
		__obf_87e02bac613d845a := time.Now().Add(-1 * __obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.MaxAge)
		for _, __obf_22504fe46317cdf1 := range __obf_1bff92203c5f8783 {
			if __obf_22504fe46317cdf1.__obf_8cb9a791c75a846b.Before(__obf_87e02bac613d845a) {
				__obf_7b2621889262be2a = append(__obf_7b2621889262be2a, __obf_22504fe46317cdf1)
			} else {
				__obf_af319cef428666be = append(__obf_af319cef428666be, __obf_22504fe46317cdf1)
			}
		}
	}

	var __obf_725b607b8fed6610 []error
	for _, __obf_22504fe46317cdf1 := range __obf_7b2621889262be2a {
		if __obf_881b15aa6c554968 = os.Remove(filepath.Join(__obf_b151e3f928f75e3e.__obf_ff9dadb72447a062, __obf_22504fe46317cdf1.Name())); __obf_881b15aa6c554968 != nil {
			__obf_725b607b8fed6610 = append(__obf_725b607b8fed6610, fmt.Errorf("remove %s: %w", __obf_22504fe46317cdf1.Name(), __obf_881b15aa6c554968))
		}
	}
	if len(__obf_725b607b8fed6610) > 0 {
		return errors.Join(__obf_725b607b8fed6610...)
	}

	for _, __obf_22504fe46317cdf1 := range __obf_af319cef428666be {
		if strings.HasSuffix(__obf_22504fe46317cdf1.Name(), __obf_3823bd60ad7ed68a) {
			continue
		}
		__obf_3daf74da58d8a12a := filepath.Join(__obf_b151e3f928f75e3e.__obf_ff9dadb72447a062, __obf_22504fe46317cdf1.Name())

		__obf_881b15aa6c554968 = util.PathGzip(__obf_3daf74da58d8a12a)
		if __obf_881b15aa6c554968 != nil {
			__obf_725b607b8fed6610 = append(__obf_725b607b8fed6610, fmt.Errorf("compress %s: %w", __obf_22504fe46317cdf1.Name(), __obf_881b15aa6c554968))
		} else {
			if __obf_881b15aa6c554968 = os.Remove(__obf_3daf74da58d8a12a); __obf_881b15aa6c554968 != nil {
				__obf_725b607b8fed6610 = append(__obf_725b607b8fed6610, fmt.Errorf("remove %s: %w", __obf_3daf74da58d8a12a, __obf_881b15aa6c554968))
			}
		}
	}

	return errors.Join(__obf_725b607b8fed6610...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_b151e3f928f75e3e *LogWriter) __obf_a7d2a17e330dc9d2() ([]__obf_2e43a57249cfa23c, error) {
	__obf_1bff92203c5f8783, __obf_881b15aa6c554968 := os.ReadDir(__obf_b151e3f928f75e3e.__obf_ff9dadb72447a062)
	if __obf_881b15aa6c554968 != nil {
		return nil, __obf_881b15aa6c554968
	}
	var (
		__obf_1679b218c485ac30                         time.Time
		__obf_060d2064379f5c85                         []__obf_2e43a57249cfa23c
		__obf_45c7627a853f8d21, __obf_b183719997b8fdcd = util.SplitFilename(__obf_b151e3f928f75e3e.__obf_dd97462f5f12dd59.LogPath)
	)
	__obf_45c7627a853f8d21 += "-"
	for _, __obf_22504fe46317cdf1 := range __obf_1bff92203c5f8783 {
		if __obf_22504fe46317cdf1.IsDir() {
			continue
		}
		if __obf_1679b218c485ac30, __obf_881b15aa6c554968 = __obf_3427413971b445b0(__obf_acbbc716f349af37, __obf_22504fe46317cdf1.Name(), __obf_45c7627a853f8d21, __obf_b183719997b8fdcd); __obf_881b15aa6c554968 == nil {
			__obf_060d2064379f5c85 = append(__obf_060d2064379f5c85, __obf_2e43a57249cfa23c{__obf_22504fe46317cdf1, __obf_1679b218c485ac30})
			continue
		}
		if __obf_1679b218c485ac30, __obf_881b15aa6c554968 = __obf_3427413971b445b0(__obf_acbbc716f349af37, __obf_22504fe46317cdf1.Name(), __obf_45c7627a853f8d21, __obf_b183719997b8fdcd+__obf_3823bd60ad7ed68a); __obf_881b15aa6c554968 == nil {
			__obf_060d2064379f5c85 = append(__obf_060d2064379f5c85, __obf_2e43a57249cfa23c{__obf_22504fe46317cdf1, __obf_1679b218c485ac30})
			continue
		}
	}
	sort.Sort(__obf_f6b52dab5f2baf35(__obf_060d2064379f5c85))
	return __obf_060d2064379f5c85, nil
}

type __obf_2e43a57249cfa23c struct {
	os.DirEntry
	__obf_8cb9a791c75a846b time.Time
}

type __obf_f6b52dab5f2baf35 []__obf_2e43a57249cfa23c

func (__obf_749ad13e2bc2bb30 __obf_f6b52dab5f2baf35) Less(__obf_dd94ff2b4d44adc8, __obf_03deafba68e8ae75 int) bool {
	return __obf_749ad13e2bc2bb30[__obf_dd94ff2b4d44adc8].__obf_8cb9a791c75a846b.After(__obf_749ad13e2bc2bb30[__obf_03deafba68e8ae75].__obf_8cb9a791c75a846b)
}

func (__obf_749ad13e2bc2bb30 __obf_f6b52dab5f2baf35) Swap(__obf_dd94ff2b4d44adc8, __obf_03deafba68e8ae75 int) {
	__obf_749ad13e2bc2bb30[__obf_dd94ff2b4d44adc8], __obf_749ad13e2bc2bb30[__obf_03deafba68e8ae75] = __obf_749ad13e2bc2bb30[__obf_03deafba68e8ae75], __obf_749ad13e2bc2bb30[__obf_dd94ff2b4d44adc8]
}

func (__obf_749ad13e2bc2bb30 __obf_f6b52dab5f2baf35) Len() int {
	return len(__obf_749ad13e2bc2bb30)
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
func __obf_3427413971b445b0(__obf_acbbc716f349af37, __obf_e2431e001ba52f63, __obf_45c7627a853f8d21, __obf_b183719997b8fdcd string) (time.Time, error) {
	if !strings.HasPrefix(__obf_e2431e001ba52f63, __obf_45c7627a853f8d21) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_45c7627a853f8d21, __obf_e2431e001ba52f63)
	}
	if !strings.HasSuffix(__obf_e2431e001ba52f63, __obf_b183719997b8fdcd) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_b183719997b8fdcd, __obf_e2431e001ba52f63)
	}
	__obf_755e52980e0f922c := __obf_e2431e001ba52f63[len(__obf_45c7627a853f8d21) : len(__obf_e2431e001ba52f63)-len(__obf_b183719997b8fdcd)]
	return time.Parse(__obf_acbbc716f349af37, __obf_755e52980e0f922c)
}
