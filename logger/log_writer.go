// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_6664e1d521d1e765

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
	__obf_aad5512029c94f78 = "2006-01-02_15:04:05.000"
	__obf_83b12c7ec728a5a3 = ".gz"
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
	__obf_1056bcce33496c89 Option
	__obf_5459f006a6731245 *os.File
	__obf_5dd561b4d3957d72 int64
	__obf_273d1c9050580319 string
	__obf_30e40ed73af1945f sync.Mutex
	__obf_2cdf780121bf82eb sync.Once
	__obf_a55c8d738c6a879c chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_1056bcce33496c89 Option) *LogWriter {
	__obf_273d1c9050580319 := filepath.Dir(__obf_1056bcce33496c89.LogPath)

	__obf_1056bcce33496c89.MaxAge = __obf_1056bcce33496c89.MaxAge * time.Hour * 24
	__obf_1056bcce33496c89.MaxSize = __obf_1056bcce33496c89.MaxSize * 1024 * 1024

	__obf_ddf58715333f46a5 := &LogWriter{
		__obf_1056bcce33496c89: __obf_1056bcce33496c89,
		__obf_273d1c9050580319: __obf_273d1c9050580319,
	}

	return __obf_ddf58715333f46a5
}

// Write writes content into file.
func (__obf_426855a9c71dd95b *LogWriter) Write(__obf_6e06d57c9288f851 []byte) (__obf_432de3dc90edb40e int, __obf_7285451253707447 error) {
	__obf_426855a9c71dd95b.__obf_30e40ed73af1945f.Lock()
	defer __obf_426855a9c71dd95b.__obf_30e40ed73af1945f.Unlock()
	__obf_1ff7f42d79cc2cd6 := int64(len(__obf_6e06d57c9288f851))

	if __obf_426855a9c71dd95b.__obf_5459f006a6731245 == nil {
		if __obf_7285451253707447 = __obf_426855a9c71dd95b.__obf_46992d4ff46f8058(__obf_1ff7f42d79cc2cd6); __obf_7285451253707447 != nil {
			return
		}
	}
	if __obf_426855a9c71dd95b.__obf_4108da2e9024d646(__obf_1ff7f42d79cc2cd6 + __obf_426855a9c71dd95b.__obf_5dd561b4d3957d72) {
		if __obf_7285451253707447 = __obf_426855a9c71dd95b.__obf_b12f10bc154f7f73(); __obf_7285451253707447 != nil {
			return
		}
	}
	__obf_432de3dc90edb40e, __obf_7285451253707447 = __obf_426855a9c71dd95b.__obf_5459f006a6731245.Write(__obf_6e06d57c9288f851)
	__obf_426855a9c71dd95b.__obf_5dd561b4d3957d72 += int64(__obf_432de3dc90edb40e)
	return
}

// Close closes file resource
func (__obf_426855a9c71dd95b *LogWriter) Close() error {
	__obf_426855a9c71dd95b.__obf_30e40ed73af1945f.Lock()
	defer __obf_426855a9c71dd95b.__obf_30e40ed73af1945f.Unlock()
	return __obf_426855a9c71dd95b.close()
}

func (__obf_426855a9c71dd95b *LogWriter) close() error {
	if __obf_426855a9c71dd95b.__obf_5459f006a6731245 == nil {
		return nil
	}
	__obf_7285451253707447 := __obf_426855a9c71dd95b.__obf_5459f006a6731245.Close()
	if __obf_7285451253707447 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_7285451253707447)
	}
	__obf_426855a9c71dd95b.__obf_5459f006a6731245 = nil
	__obf_426855a9c71dd95b.__obf_5dd561b4d3957d72 = 0
	return nil
}

func (__obf_426855a9c71dd95b *LogWriter) __obf_4108da2e9024d646(__obf_5dd561b4d3957d72 int64) bool {
	return __obf_426855a9c71dd95b.__obf_1056bcce33496c89.MaxSize > 0 && __obf_5dd561b4d3957d72 > __obf_426855a9c71dd95b.__obf_1056bcce33496c89.MaxSize
}

// openFile 负责打开日志文件以供写入。
// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
// 如果文件存在，它会判断是否需要轮换日志文件。
func (__obf_426855a9c71dd95b *LogWriter) __obf_46992d4ff46f8058(__obf_9e539ca2abe55331 int64) error {
	__obf_3ed41f26f48b94b3, __obf_7285451253707447 := os.Stat(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath)
	if os.IsNotExist(__obf_7285451253707447) {
		return __obf_426855a9c71dd95b.__obf_f27be902f4fc2d15()
	}
	if __obf_7285451253707447 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_7285451253707447)
	}
	if __obf_426855a9c71dd95b.__obf_4108da2e9024d646(__obf_3ed41f26f48b94b3.Size() + __obf_9e539ca2abe55331) {
		return __obf_426855a9c71dd95b.__obf_b12f10bc154f7f73()
	}
	__obf_5459f006a6731245, __obf_7285451253707447 := os.OpenFile(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_7285451253707447 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_7285451253707447)
	}
	__obf_426855a9c71dd95b.__obf_5459f006a6731245 = __obf_5459f006a6731245
	__obf_426855a9c71dd95b.__obf_5dd561b4d3957d72 = __obf_3ed41f26f48b94b3.Size()
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
func (__obf_426855a9c71dd95b *LogWriter) __obf_f27be902f4fc2d15() error {
	if __obf_7285451253707447 := os.MkdirAll(__obf_426855a9c71dd95b.__obf_273d1c9050580319, os.ModePerm); __obf_7285451253707447 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_7285451253707447)
	}
	__obf_c5c6418b3c6e6c9a := os.FileMode(0644)
	__obf_3ed41f26f48b94b3, __obf_7285451253707447 := os.Stat(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath)
	if __obf_7285451253707447 == nil {
		__obf_c5c6418b3c6e6c9a = __obf_3ed41f26f48b94b3.Mode()
		__obf_d7475ad494cced56, __obf_ecc383f1785c8ace := util.SplitFilename(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath)
		__obf_5a02fc6daeb8c9be := time.Now().Format(__obf_aad5512029c94f78)
		__obf_74d887c0fdda735c := filepath.Join(__obf_426855a9c71dd95b.__obf_273d1c9050580319, fmt.Sprintf("%s-%s%s", __obf_d7475ad494cced56, __obf_5a02fc6daeb8c9be, __obf_ecc383f1785c8ace))

		if __obf_7285451253707447 = os.Rename(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath, __obf_74d887c0fdda735c); __obf_7285451253707447 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_74d887c0fdda735c, __obf_7285451253707447)
		}
	}
	__obf_07aa1e94da759dbe, __obf_7285451253707447 := os.OpenFile(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_c5c6418b3c6e6c9a)
	if __obf_7285451253707447 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath, __obf_7285451253707447)
	}
	__obf_426855a9c71dd95b.__obf_5459f006a6731245 = __obf_07aa1e94da759dbe
	__obf_426855a9c71dd95b.__obf_5dd561b4d3957d72 = 0
	return nil
}

// rotate 处理 LogWriter 的日志文件轮换过程。
// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
//
// 如果关闭当前文件或打开新文件失败，则返回错误。

func (__obf_426855a9c71dd95b *LogWriter) __obf_b12f10bc154f7f73() error {
	if __obf_7285451253707447 := __obf_426855a9c71dd95b.close(); __obf_7285451253707447 != nil {
		return __obf_7285451253707447
	}
	if __obf_7285451253707447 := __obf_426855a9c71dd95b.__obf_f27be902f4fc2d15(); __obf_7285451253707447 != nil {
		return __obf_7285451253707447
	}
	__obf_426855a9c71dd95b.__obf_2cdf780121bf82eb.Do(func() {
		__obf_426855a9c71dd95b.__obf_a55c8d738c6a879c = make(chan bool, 1)
		go func() {
			for range __obf_426855a9c71dd95b.__obf_a55c8d738c6a879c {
				if __obf_cbda6a52468d50c3 := __obf_426855a9c71dd95b.__obf_4a7e599de2b5e0e2(); __obf_cbda6a52468d50c3 != nil {
					println(__obf_cbda6a52468d50c3.Error())
				}
			}
		}()
	})
	select {
	case __obf_426855a9c71dd95b.__obf_a55c8d738c6a879c <- true:
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
func (__obf_426855a9c71dd95b *LogWriter) __obf_4a7e599de2b5e0e2() error {
	if __obf_426855a9c71dd95b.__obf_1056bcce33496c89.MaxSize == 0 {
		return nil
	}
	__obf_fc3d4e3681b972c1, __obf_7285451253707447 := __obf_426855a9c71dd95b.__obf_ca7f615e6aeaa38b()
	if __obf_7285451253707447 != nil {
		return __obf_7285451253707447
	}
	var __obf_547027b901f90026, __obf_1187daea557a02bc []__obf_1e82ff5d7c37cf98
	if __obf_426855a9c71dd95b.__obf_1056bcce33496c89.MaxAge > 0 {
		__obf_968f88d29763beb9 := time.Now().Add(-1 * __obf_426855a9c71dd95b.__obf_1056bcce33496c89.MaxAge)
		for _, __obf_07aa1e94da759dbe := range __obf_fc3d4e3681b972c1 {
			if __obf_07aa1e94da759dbe.__obf_d6f53121e723e5aa.Before(__obf_968f88d29763beb9) {
				__obf_547027b901f90026 = append(__obf_547027b901f90026, __obf_07aa1e94da759dbe)
			} else {
				__obf_1187daea557a02bc = append(__obf_1187daea557a02bc, __obf_07aa1e94da759dbe)
			}
		}
	}

	var __obf_b110d5db145879d7 []error
	for _, __obf_07aa1e94da759dbe := range __obf_547027b901f90026 {
		if __obf_7285451253707447 = os.Remove(filepath.Join(__obf_426855a9c71dd95b.__obf_273d1c9050580319, __obf_07aa1e94da759dbe.Name())); __obf_7285451253707447 != nil {
			__obf_b110d5db145879d7 = append(__obf_b110d5db145879d7, fmt.Errorf("remove %s: %w", __obf_07aa1e94da759dbe.Name(), __obf_7285451253707447))
		}
	}
	if len(__obf_b110d5db145879d7) > 0 {
		return errors.Join(__obf_b110d5db145879d7...)
	}

	for _, __obf_07aa1e94da759dbe := range __obf_1187daea557a02bc {
		if strings.HasSuffix(__obf_07aa1e94da759dbe.Name(), __obf_83b12c7ec728a5a3) {
			continue
		}
		__obf_62203d5556484182 := filepath.Join(__obf_426855a9c71dd95b.__obf_273d1c9050580319, __obf_07aa1e94da759dbe.Name())

		__obf_7285451253707447 = util.PathGzip(__obf_62203d5556484182)
		if __obf_7285451253707447 != nil {
			__obf_b110d5db145879d7 = append(__obf_b110d5db145879d7, fmt.Errorf("compress %s: %w", __obf_07aa1e94da759dbe.Name(), __obf_7285451253707447))
		} else {
			if __obf_7285451253707447 = os.Remove(__obf_62203d5556484182); __obf_7285451253707447 != nil {
				__obf_b110d5db145879d7 = append(__obf_b110d5db145879d7, fmt.Errorf("remove %s: %w", __obf_62203d5556484182, __obf_7285451253707447))
			}
		}
	}

	return errors.Join(__obf_b110d5db145879d7...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_426855a9c71dd95b *LogWriter) __obf_ca7f615e6aeaa38b() ([]__obf_1e82ff5d7c37cf98, error) {
	__obf_fc3d4e3681b972c1, __obf_7285451253707447 := os.ReadDir(__obf_426855a9c71dd95b.__obf_273d1c9050580319)
	if __obf_7285451253707447 != nil {
		return nil, __obf_7285451253707447
	}
	var (
		__obf_5a02fc6daeb8c9be                         time.Time
		__obf_4387a04e89a1ceb6                         []__obf_1e82ff5d7c37cf98
		__obf_d7475ad494cced56, __obf_ecc383f1785c8ace = util.SplitFilename(__obf_426855a9c71dd95b.__obf_1056bcce33496c89.LogPath)
	)
	__obf_d7475ad494cced56 += "-"
	for _, __obf_07aa1e94da759dbe := range __obf_fc3d4e3681b972c1 {
		if __obf_07aa1e94da759dbe.IsDir() {
			continue
		}
		if __obf_5a02fc6daeb8c9be, __obf_7285451253707447 = __obf_c71f5430e2151b17(__obf_aad5512029c94f78, __obf_07aa1e94da759dbe.Name(), __obf_d7475ad494cced56, __obf_ecc383f1785c8ace); __obf_7285451253707447 == nil {
			__obf_4387a04e89a1ceb6 = append(__obf_4387a04e89a1ceb6, __obf_1e82ff5d7c37cf98{__obf_07aa1e94da759dbe, __obf_5a02fc6daeb8c9be})
			continue
		}
		if __obf_5a02fc6daeb8c9be, __obf_7285451253707447 = __obf_c71f5430e2151b17(__obf_aad5512029c94f78, __obf_07aa1e94da759dbe.Name(), __obf_d7475ad494cced56, __obf_ecc383f1785c8ace+__obf_83b12c7ec728a5a3); __obf_7285451253707447 == nil {
			__obf_4387a04e89a1ceb6 = append(__obf_4387a04e89a1ceb6, __obf_1e82ff5d7c37cf98{__obf_07aa1e94da759dbe, __obf_5a02fc6daeb8c9be})
			continue
		}
	}
	sort.Sort(__obf_b59e6d9c2357a0a3(__obf_4387a04e89a1ceb6))
	return __obf_4387a04e89a1ceb6, nil
}

type __obf_1e82ff5d7c37cf98 struct {
	os.DirEntry
	__obf_d6f53121e723e5aa time.Time
}

type __obf_b59e6d9c2357a0a3 []__obf_1e82ff5d7c37cf98

func (__obf_2f88505312c29755 __obf_b59e6d9c2357a0a3) Less(__obf_4fd4e707bda8dc68, __obf_7c8597bc28106e89 int) bool {
	return __obf_2f88505312c29755[__obf_4fd4e707bda8dc68].__obf_d6f53121e723e5aa.After(__obf_2f88505312c29755[__obf_7c8597bc28106e89].__obf_d6f53121e723e5aa)
}

func (__obf_2f88505312c29755 __obf_b59e6d9c2357a0a3) Swap(__obf_4fd4e707bda8dc68, __obf_7c8597bc28106e89 int) {
	__obf_2f88505312c29755[__obf_4fd4e707bda8dc68], __obf_2f88505312c29755[__obf_7c8597bc28106e89] = __obf_2f88505312c29755[__obf_7c8597bc28106e89], __obf_2f88505312c29755[__obf_4fd4e707bda8dc68]
}

func (__obf_2f88505312c29755 __obf_b59e6d9c2357a0a3) Len() int {
	return len(__obf_2f88505312c29755)
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
func __obf_c71f5430e2151b17(__obf_aad5512029c94f78, __obf_28a967bc1226e476, __obf_d7475ad494cced56, __obf_ecc383f1785c8ace string) (time.Time, error) {
	if !strings.HasPrefix(__obf_28a967bc1226e476, __obf_d7475ad494cced56) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_d7475ad494cced56, __obf_28a967bc1226e476)
	}
	if !strings.HasSuffix(__obf_28a967bc1226e476, __obf_ecc383f1785c8ace) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_ecc383f1785c8ace, __obf_28a967bc1226e476)
	}
	__obf_91a256d9fd7db8bc := __obf_28a967bc1226e476[len(__obf_d7475ad494cced56) : len(__obf_28a967bc1226e476)-len(__obf_ecc383f1785c8ace)]
	return time.Parse(__obf_aad5512029c94f78, __obf_91a256d9fd7db8bc)
}
