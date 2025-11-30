// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_0f05083f0caeeb4d

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
	__obf_9b4cf408039d10df = "2006-01-02_15:04:05.000"
	__obf_ce9e41b835caf00a = ".gz"
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
	__obf_1b9dd506cd45ee02 Option
	__obf_70f2d96a7c56d8e7 *os.File
	__obf_00f262893f08e1e1 int64
	__obf_c075281f4170a2bf string
	__obf_1534f2715af3f787 sync.Mutex
	__obf_e497ee5818e27459 sync.Once
	__obf_0c08a2ac3a79cedc chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_1b9dd506cd45ee02 Option) *LogWriter {
	__obf_c075281f4170a2bf := filepath.Dir(__obf_1b9dd506cd45ee02.LogPath)
	__obf_1b9dd506cd45ee02.
		MaxAge = __obf_1b9dd506cd45ee02.MaxAge * time.Hour * 24
	__obf_1b9dd506cd45ee02.
		MaxSize = __obf_1b9dd506cd45ee02.MaxSize * 1024 * 1024
	__obf_3d83cb1ce9baad05 := &LogWriter{__obf_1b9dd506cd45ee02: __obf_1b9dd506cd45ee02, __obf_c075281f4170a2bf: __obf_c075281f4170a2bf}

	return __obf_3d83cb1ce9baad05
}

// Write writes content into file.
func (__obf_71761bf6a852ebd6 *LogWriter) Write(__obf_8229f195d9b845b5 []byte) (__obf_d570e13cbf48109f int, __obf_c336006ef0dba3b9 error) {
	__obf_71761bf6a852ebd6.__obf_1534f2715af3f787.
		Lock()
	defer __obf_71761bf6a852ebd6.__obf_1534f2715af3f787.Unlock()
	__obf_39f8b0fa8b89a70b := int64(len(__obf_8229f195d9b845b5))

	if __obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7 == nil {
		if __obf_c336006ef0dba3b9 = __obf_71761bf6a852ebd6.__obf_545578855cf0fab8(__obf_39f8b0fa8b89a70b); __obf_c336006ef0dba3b9 != nil {
			return
		}
	}
	if __obf_71761bf6a852ebd6.__obf_d6536274e9cb0bc9(__obf_39f8b0fa8b89a70b + __obf_71761bf6a852ebd6.__obf_00f262893f08e1e1) {
		if __obf_c336006ef0dba3b9 = __obf_71761bf6a852ebd6.__obf_b72f6cb1fe0cabe5(); __obf_c336006ef0dba3b9 != nil {
			return
		}
	}
	__obf_d570e13cbf48109f, __obf_c336006ef0dba3b9 = __obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7.Write(__obf_8229f195d9b845b5)
	__obf_71761bf6a852ebd6.__obf_00f262893f08e1e1 += int64(__obf_d570e13cbf48109f)
	return
}

// Close closes file resource
func (__obf_71761bf6a852ebd6 *LogWriter) Close() error {
	__obf_71761bf6a852ebd6.__obf_1534f2715af3f787.
		Lock()
	defer __obf_71761bf6a852ebd6.__obf_1534f2715af3f787.Unlock()
	return __obf_71761bf6a852ebd6.close()
}

func (__obf_71761bf6a852ebd6 *LogWriter) close() error {
	if __obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7 == nil {
		return nil
	}
	__obf_c336006ef0dba3b9 := __obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7.Close()
	if __obf_c336006ef0dba3b9 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_c336006ef0dba3b9)
	}
	__obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7 = nil
	__obf_71761bf6a852ebd6.__obf_00f262893f08e1e1 = 0
	return nil
}

func (__obf_71761bf6a852ebd6 *LogWriter) __obf_d6536274e9cb0bc9(__obf_00f262893f08e1e1 int64) bool {
	return __obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.MaxSize > 0 && __obf_00f262893f08e1e1 > __obf_71761bf6a852ebd6.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_1b9dd506cd45ee02.MaxSize
}

func (__obf_71761bf6a852ebd6 *LogWriter) __obf_545578855cf0fab8(__obf_ea0ec17a7a34e6af int64) error {
	__obf_72f88a35a594bd11, __obf_c336006ef0dba3b9 := os.Stat(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath)
	if os.IsNotExist(__obf_c336006ef0dba3b9) {
		return __obf_71761bf6a852ebd6.__obf_2c904fb2a6e23a2c()
	}
	if __obf_c336006ef0dba3b9 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_c336006ef0dba3b9)
	}
	if __obf_71761bf6a852ebd6.__obf_d6536274e9cb0bc9(__obf_72f88a35a594bd11.Size() + __obf_ea0ec17a7a34e6af) {
		return __obf_71761bf6a852ebd6.__obf_b72f6cb1fe0cabe5()
	}
	__obf_70f2d96a7c56d8e7, __obf_c336006ef0dba3b9 := os.OpenFile(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_c336006ef0dba3b9 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_c336006ef0dba3b9)
	}
	__obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7 = __obf_70f2d96a7c56d8e7
	__obf_71761bf6a852ebd6.

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
		__obf_00f262893f08e1e1 = __obf_72f88a35a594bd11.Size()
	return nil
}

func (__obf_71761bf6a852ebd6 *LogWriter) __obf_2c904fb2a6e23a2c() error {
	if __obf_c336006ef0dba3b9 := os.MkdirAll(__obf_71761bf6a852ebd6.__obf_c075281f4170a2bf, os.ModePerm); __obf_c336006ef0dba3b9 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_c336006ef0dba3b9)
	}
	__obf_a8a474ec05171236 := os.FileMode(0644)
	__obf_72f88a35a594bd11, __obf_c336006ef0dba3b9 := os.Stat(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath)
	if __obf_c336006ef0dba3b9 == nil {
		__obf_a8a474ec05171236 = __obf_72f88a35a594bd11.Mode()
		__obf_f7a1a239ce328200, __obf_a08842b02ff18dcf := util.SplitFilename(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath)
		__obf_f26b469c3552f8f9 := time.Now().Format(__obf_9b4cf408039d10df)
		__obf_c700791ceb615649 := filepath.Join(__obf_71761bf6a852ebd6.__obf_c075281f4170a2bf, fmt.Sprintf("%s-%s%s", __obf_f7a1a239ce328200, __obf_f26b469c3552f8f9, __obf_a08842b02ff18dcf))

		if __obf_c336006ef0dba3b9 = os.Rename(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath, __obf_c700791ceb615649); __obf_c336006ef0dba3b9 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_c700791ceb615649, __obf_c336006ef0dba3b9)
		}
	}
	__obf_4f2668ae3f0a4dea, __obf_c336006ef0dba3b9 := os.OpenFile(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_a8a474ec05171236)
	if __obf_c336006ef0dba3b9 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath, __obf_c336006ef0dba3b9)
	}
	__obf_71761bf6a852ebd6.__obf_70f2d96a7c56d8e7 = __obf_4f2668ae3f0a4dea
	__obf_71761bf6a852ebd6.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_00f262893f08e1e1 = 0
	return nil
}

func (__obf_71761bf6a852ebd6 *LogWriter) __obf_b72f6cb1fe0cabe5() error {
	if __obf_c336006ef0dba3b9 := __obf_71761bf6a852ebd6.close(); __obf_c336006ef0dba3b9 != nil {
		return __obf_c336006ef0dba3b9
	}
	if __obf_c336006ef0dba3b9 := __obf_71761bf6a852ebd6.__obf_2c904fb2a6e23a2c(); __obf_c336006ef0dba3b9 != nil {
		return __obf_c336006ef0dba3b9
	}
	__obf_71761bf6a852ebd6.__obf_e497ee5818e27459.
		Do(func() {
			__obf_71761bf6a852ebd6.__obf_0c08a2ac3a79cedc = make(chan bool, 1)
			go func() {
				for range __obf_71761bf6a852ebd6.__obf_0c08a2ac3a79cedc {
					if __obf_fd88e5b40cbceecd := __obf_71761bf6a852ebd6.__obf_948959f414e7a2e2(); __obf_fd88e5b40cbceecd != nil {
						println(__obf_fd88e5b40cbceecd.Error())
					}
				}
			}()
		})
	select {
	case __obf_71761bf6a852ebd6.__obf_0c08a2ac3a79cedc <- true:
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
func (__obf_71761bf6a852ebd6 *LogWriter) __obf_948959f414e7a2e2() error {
	if __obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.MaxSize == 0 {
		return nil
	}
	__obf_3cac69f7a848b5ee, __obf_c336006ef0dba3b9 := __obf_71761bf6a852ebd6.__obf_02a574cdc7424d76()
	if __obf_c336006ef0dba3b9 != nil {
		return __obf_c336006ef0dba3b9
	}
	var __obf_fd356382eced3335, __obf_9571a42384b76f19 []__obf_09a57e372039e976
	if __obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.MaxAge > 0 {
		__obf_b51e84a88880d711 := time.Now().Add(-1 * __obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.MaxAge)
		for _, __obf_4f2668ae3f0a4dea := range __obf_3cac69f7a848b5ee {
			if __obf_4f2668ae3f0a4dea.__obf_4e58f20f26e0463d.Before(__obf_b51e84a88880d711) {
				__obf_fd356382eced3335 = append(__obf_fd356382eced3335, __obf_4f2668ae3f0a4dea)
			} else {
				__obf_9571a42384b76f19 = append(__obf_9571a42384b76f19, __obf_4f2668ae3f0a4dea)
			}
		}
	}

	var __obf_47638f41a4dc72d4 []error
	for _, __obf_4f2668ae3f0a4dea := range __obf_fd356382eced3335 {
		if __obf_c336006ef0dba3b9 = os.Remove(filepath.Join(__obf_71761bf6a852ebd6.__obf_c075281f4170a2bf, __obf_4f2668ae3f0a4dea.Name())); __obf_c336006ef0dba3b9 != nil {
			__obf_47638f41a4dc72d4 = append(__obf_47638f41a4dc72d4, fmt.Errorf("remove %s: %w", __obf_4f2668ae3f0a4dea.Name(), __obf_c336006ef0dba3b9))
		}
	}
	if len(__obf_47638f41a4dc72d4) > 0 {
		return errors.Join(__obf_47638f41a4dc72d4...)
	}

	for _, __obf_4f2668ae3f0a4dea := range __obf_9571a42384b76f19 {
		if strings.HasSuffix(__obf_4f2668ae3f0a4dea.Name(), __obf_ce9e41b835caf00a) {
			continue
		}
		__obf_444b53e517c97948 := filepath.Join(__obf_71761bf6a852ebd6.__obf_c075281f4170a2bf, __obf_4f2668ae3f0a4dea.Name())
		__obf_c336006ef0dba3b9 = util.PathGzip(__obf_444b53e517c97948)
		if __obf_c336006ef0dba3b9 != nil {
			__obf_47638f41a4dc72d4 = append(__obf_47638f41a4dc72d4, fmt.Errorf("compress %s: %w", __obf_4f2668ae3f0a4dea.Name(), __obf_c336006ef0dba3b9))
		} else {
			if __obf_c336006ef0dba3b9 = os.Remove(__obf_444b53e517c97948); __obf_c336006ef0dba3b9 != nil {
				__obf_47638f41a4dc72d4 = append(__obf_47638f41a4dc72d4, fmt.Errorf("remove %s: %w", __obf_444b53e517c97948, __obf_c336006ef0dba3b9))
			}
		}
	}

	return errors.Join(__obf_47638f41a4dc72d4...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_71761bf6a852ebd6 *LogWriter) __obf_02a574cdc7424d76() ([]__obf_09a57e372039e976, error) {
	__obf_3cac69f7a848b5ee, __obf_c336006ef0dba3b9 := os.ReadDir(__obf_71761bf6a852ebd6.__obf_c075281f4170a2bf)
	if __obf_c336006ef0dba3b9 != nil {
		return nil, __obf_c336006ef0dba3b9
	}
	var (
		__obf_f26b469c3552f8f9                         time.Time
		__obf_4c65daabf27a895b                         []__obf_09a57e372039e976
		__obf_f7a1a239ce328200, __obf_a08842b02ff18dcf = util.SplitFilename(__obf_71761bf6a852ebd6.__obf_1b9dd506cd45ee02.LogPath)
	)
	__obf_f7a1a239ce328200 += "-"
	for _, __obf_4f2668ae3f0a4dea := range __obf_3cac69f7a848b5ee {
		if __obf_4f2668ae3f0a4dea.IsDir() {
			continue
		}
		if __obf_f26b469c3552f8f9, __obf_c336006ef0dba3b9 = __obf_6536055a11acb4e2(__obf_9b4cf408039d10df, __obf_4f2668ae3f0a4dea.Name(), __obf_f7a1a239ce328200, __obf_a08842b02ff18dcf); __obf_c336006ef0dba3b9 == nil {
			__obf_4c65daabf27a895b = append(__obf_4c65daabf27a895b, __obf_09a57e372039e976{__obf_4f2668ae3f0a4dea, __obf_f26b469c3552f8f9})
			continue
		}
		if __obf_f26b469c3552f8f9, __obf_c336006ef0dba3b9 = __obf_6536055a11acb4e2(__obf_9b4cf408039d10df, __obf_4f2668ae3f0a4dea.Name(), __obf_f7a1a239ce328200, __obf_a08842b02ff18dcf+__obf_ce9e41b835caf00a); __obf_c336006ef0dba3b9 == nil {
			__obf_4c65daabf27a895b = append(__obf_4c65daabf27a895b, __obf_09a57e372039e976{__obf_4f2668ae3f0a4dea, __obf_f26b469c3552f8f9})
			continue
		}
	}
	sort.Sort(__obf_447c7a2a6ac43bce(__obf_4c65daabf27a895b))
	return __obf_4c65daabf27a895b, nil
}

type __obf_09a57e372039e976 struct {
	os.DirEntry
	__obf_4e58f20f26e0463d time.Time
}

type __obf_447c7a2a6ac43bce []__obf_09a57e372039e976

func (__obf_83381a02823ae02c __obf_447c7a2a6ac43bce) Less(__obf_d2fbf106aa46d62a, __obf_0ab578d52bdc16d1 int) bool {
	return __obf_83381a02823ae02c[__obf_d2fbf106aa46d62a].__obf_4e58f20f26e0463d.After(__obf_83381a02823ae02c[__obf_0ab578d52bdc16d1].__obf_4e58f20f26e0463d)
}

func (__obf_83381a02823ae02c __obf_447c7a2a6ac43bce) Swap(__obf_d2fbf106aa46d62a, __obf_0ab578d52bdc16d1 int) {
	__obf_83381a02823ae02c[__obf_d2fbf106aa46d62a], __obf_83381a02823ae02c[__obf_0ab578d52bdc16d1] = __obf_83381a02823ae02c[__obf_0ab578d52bdc16d1], __obf_83381a02823ae02c[__obf_d2fbf106aa46d62a]
}

func (__obf_83381a02823ae02c __obf_447c7a2a6ac43bce) Len() int {
	return len(__obf_83381a02823ae02c)
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
func __obf_6536055a11acb4e2(__obf_9b4cf408039d10df, __obf_6770791b53a0efa5, __obf_f7a1a239ce328200, __obf_a08842b02ff18dcf string) (time.Time, error) {
	if !strings.HasPrefix(__obf_6770791b53a0efa5, __obf_f7a1a239ce328200) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_f7a1a239ce328200, __obf_6770791b53a0efa5)
	}
	if !strings.HasSuffix(__obf_6770791b53a0efa5, __obf_a08842b02ff18dcf) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_a08842b02ff18dcf, __obf_6770791b53a0efa5)
	}
	__obf_e4d937abd52f1454 := __obf_6770791b53a0efa5[len(__obf_f7a1a239ce328200) : len(__obf_6770791b53a0efa5)-len(__obf_a08842b02ff18dcf)]
	return time.Parse(__obf_9b4cf408039d10df, __obf_e4d937abd52f1454)
}
