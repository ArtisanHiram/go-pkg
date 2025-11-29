// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_ad1b8b65c829ec46

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
	__obf_e6b7cf7168f3666a = "2006-01-02_15:04:05.000"
	__obf_128408006183829e = ".gz"
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
	__obf_b9186a3395106d9b Option
	__obf_f669dbd05fc7537a *os.File
	__obf_dae8a61c894d3064 int64
	__obf_d764b8d55a7df3d7 string
	__obf_61027a9c2b5d6596 sync.Mutex
	__obf_584ee8c58daa2672 sync.Once
	__obf_3db76e66c496a137 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_b9186a3395106d9b Option) *LogWriter {
	__obf_d764b8d55a7df3d7 := filepath.Dir(__obf_b9186a3395106d9b.LogPath)
	__obf_b9186a3395106d9b.
		MaxAge = __obf_b9186a3395106d9b.MaxAge * time.Hour * 24
	__obf_b9186a3395106d9b.
		MaxSize = __obf_b9186a3395106d9b.MaxSize * 1024 * 1024
	__obf_53ff32637d409ed6 := &LogWriter{__obf_b9186a3395106d9b: __obf_b9186a3395106d9b, __obf_d764b8d55a7df3d7: __obf_d764b8d55a7df3d7}

	return __obf_53ff32637d409ed6
}

// Write writes content into file.
func (__obf_fc2298c77aff69e0 *LogWriter) Write(__obf_4891b9fa46ac0274 []byte) (__obf_f46253eb5b555dec int, __obf_874f1b3a1c1536b7 error) {
	__obf_fc2298c77aff69e0.__obf_61027a9c2b5d6596.
		Lock()
	defer __obf_fc2298c77aff69e0.__obf_61027a9c2b5d6596.Unlock()
	__obf_f0532b4140c29178 := int64(len(__obf_4891b9fa46ac0274))

	if __obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a == nil {
		if __obf_874f1b3a1c1536b7 = __obf_fc2298c77aff69e0.__obf_1e6297755c556c9a(__obf_f0532b4140c29178); __obf_874f1b3a1c1536b7 != nil {
			return
		}
	}
	if __obf_fc2298c77aff69e0.__obf_c7bf831f40e72902(__obf_f0532b4140c29178 + __obf_fc2298c77aff69e0.__obf_dae8a61c894d3064) {
		if __obf_874f1b3a1c1536b7 = __obf_fc2298c77aff69e0.__obf_e77152193ad7d443(); __obf_874f1b3a1c1536b7 != nil {
			return
		}
	}
	__obf_f46253eb5b555dec, __obf_874f1b3a1c1536b7 = __obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a.Write(__obf_4891b9fa46ac0274)
	__obf_fc2298c77aff69e0.__obf_dae8a61c894d3064 += int64(__obf_f46253eb5b555dec)
	return
}

// Close closes file resource
func (__obf_fc2298c77aff69e0 *LogWriter) Close() error {
	__obf_fc2298c77aff69e0.__obf_61027a9c2b5d6596.
		Lock()
	defer __obf_fc2298c77aff69e0.__obf_61027a9c2b5d6596.Unlock()
	return __obf_fc2298c77aff69e0.close()
}

func (__obf_fc2298c77aff69e0 *LogWriter) close() error {
	if __obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a == nil {
		return nil
	}
	__obf_874f1b3a1c1536b7 := __obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a.Close()
	if __obf_874f1b3a1c1536b7 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_874f1b3a1c1536b7)
	}
	__obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a = nil
	__obf_fc2298c77aff69e0.__obf_dae8a61c894d3064 = 0
	return nil
}

func (__obf_fc2298c77aff69e0 *LogWriter) __obf_c7bf831f40e72902(__obf_dae8a61c894d3064 int64) bool {
	return __obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.MaxSize > 0 && __obf_dae8a61c894d3064 > __obf_fc2298c77aff69e0.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_b9186a3395106d9b.MaxSize
}

func (__obf_fc2298c77aff69e0 *LogWriter) __obf_1e6297755c556c9a(__obf_652e6717426f1e64 int64) error {
	__obf_4c8ff566671a1666, __obf_874f1b3a1c1536b7 := os.Stat(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath)
	if os.IsNotExist(__obf_874f1b3a1c1536b7) {
		return __obf_fc2298c77aff69e0.__obf_e3b197f9683349d8()
	}
	if __obf_874f1b3a1c1536b7 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_874f1b3a1c1536b7)
	}
	if __obf_fc2298c77aff69e0.__obf_c7bf831f40e72902(__obf_4c8ff566671a1666.Size() + __obf_652e6717426f1e64) {
		return __obf_fc2298c77aff69e0.__obf_e77152193ad7d443()
	}
	__obf_f669dbd05fc7537a, __obf_874f1b3a1c1536b7 := os.OpenFile(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_874f1b3a1c1536b7 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_874f1b3a1c1536b7)
	}
	__obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a = __obf_f669dbd05fc7537a
	__obf_fc2298c77aff69e0.

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
		__obf_dae8a61c894d3064 = __obf_4c8ff566671a1666.Size()
	return nil
}

func (__obf_fc2298c77aff69e0 *LogWriter) __obf_e3b197f9683349d8() error {
	if __obf_874f1b3a1c1536b7 := os.MkdirAll(__obf_fc2298c77aff69e0.__obf_d764b8d55a7df3d7, os.ModePerm); __obf_874f1b3a1c1536b7 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_874f1b3a1c1536b7)
	}
	__obf_0b7a5e8d978dd5da := os.FileMode(0644)
	__obf_4c8ff566671a1666, __obf_874f1b3a1c1536b7 := os.Stat(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath)
	if __obf_874f1b3a1c1536b7 == nil {
		__obf_0b7a5e8d978dd5da = __obf_4c8ff566671a1666.Mode()
		__obf_fc660d113a218a0a, __obf_d2a121bdff49e600 := util.SplitFilename(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath)
		__obf_a22b30d0e178bdad := time.Now().Format(__obf_e6b7cf7168f3666a)
		__obf_af10b9495ac508be := filepath.Join(__obf_fc2298c77aff69e0.__obf_d764b8d55a7df3d7, fmt.Sprintf("%s-%s%s", __obf_fc660d113a218a0a, __obf_a22b30d0e178bdad, __obf_d2a121bdff49e600))

		if __obf_874f1b3a1c1536b7 = os.Rename(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath, __obf_af10b9495ac508be); __obf_874f1b3a1c1536b7 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_af10b9495ac508be, __obf_874f1b3a1c1536b7)
		}
	}
	__obf_3b91508eac98d6e5, __obf_874f1b3a1c1536b7 := os.OpenFile(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_0b7a5e8d978dd5da)
	if __obf_874f1b3a1c1536b7 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath, __obf_874f1b3a1c1536b7)
	}
	__obf_fc2298c77aff69e0.__obf_f669dbd05fc7537a = __obf_3b91508eac98d6e5
	__obf_fc2298c77aff69e0.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_dae8a61c894d3064 = 0
	return nil
}

func (__obf_fc2298c77aff69e0 *LogWriter) __obf_e77152193ad7d443() error {
	if __obf_874f1b3a1c1536b7 := __obf_fc2298c77aff69e0.close(); __obf_874f1b3a1c1536b7 != nil {
		return __obf_874f1b3a1c1536b7
	}
	if __obf_874f1b3a1c1536b7 := __obf_fc2298c77aff69e0.__obf_e3b197f9683349d8(); __obf_874f1b3a1c1536b7 != nil {
		return __obf_874f1b3a1c1536b7
	}
	__obf_fc2298c77aff69e0.__obf_584ee8c58daa2672.
		Do(func() {
			__obf_fc2298c77aff69e0.__obf_3db76e66c496a137 = make(chan bool, 1)
			go func() {
				for range __obf_fc2298c77aff69e0.__obf_3db76e66c496a137 {
					if __obf_907f4045cda8b827 := __obf_fc2298c77aff69e0.__obf_8fe378946620e7af(); __obf_907f4045cda8b827 != nil {
						println(__obf_907f4045cda8b827.Error())
					}
				}
			}()
		})
	select {
	case __obf_fc2298c77aff69e0.__obf_3db76e66c496a137 <- true:
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
func (__obf_fc2298c77aff69e0 *LogWriter) __obf_8fe378946620e7af() error {
	if __obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.MaxSize == 0 {
		return nil
	}
	__obf_7d212fadcdc70a73, __obf_874f1b3a1c1536b7 := __obf_fc2298c77aff69e0.__obf_25a15534f674e145()
	if __obf_874f1b3a1c1536b7 != nil {
		return __obf_874f1b3a1c1536b7
	}
	var __obf_afe9fe24a9442d6f, __obf_b4458e8f8dbdcff6 []__obf_00900b790cede2ed
	if __obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.MaxAge > 0 {
		__obf_207f803ee0e348d9 := time.Now().Add(-1 * __obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.MaxAge)
		for _, __obf_3b91508eac98d6e5 := range __obf_7d212fadcdc70a73 {
			if __obf_3b91508eac98d6e5.__obf_6bc488781d67451d.Before(__obf_207f803ee0e348d9) {
				__obf_afe9fe24a9442d6f = append(__obf_afe9fe24a9442d6f, __obf_3b91508eac98d6e5)
			} else {
				__obf_b4458e8f8dbdcff6 = append(__obf_b4458e8f8dbdcff6, __obf_3b91508eac98d6e5)
			}
		}
	}

	var __obf_0f4b5fbdce447b1b []error
	for _, __obf_3b91508eac98d6e5 := range __obf_afe9fe24a9442d6f {
		if __obf_874f1b3a1c1536b7 = os.Remove(filepath.Join(__obf_fc2298c77aff69e0.__obf_d764b8d55a7df3d7, __obf_3b91508eac98d6e5.Name())); __obf_874f1b3a1c1536b7 != nil {
			__obf_0f4b5fbdce447b1b = append(__obf_0f4b5fbdce447b1b, fmt.Errorf("remove %s: %w", __obf_3b91508eac98d6e5.Name(), __obf_874f1b3a1c1536b7))
		}
	}
	if len(__obf_0f4b5fbdce447b1b) > 0 {
		return errors.Join(__obf_0f4b5fbdce447b1b...)
	}

	for _, __obf_3b91508eac98d6e5 := range __obf_b4458e8f8dbdcff6 {
		if strings.HasSuffix(__obf_3b91508eac98d6e5.Name(), __obf_128408006183829e) {
			continue
		}
		__obf_fcb4e40318c7e7bf := filepath.Join(__obf_fc2298c77aff69e0.__obf_d764b8d55a7df3d7, __obf_3b91508eac98d6e5.Name())
		__obf_874f1b3a1c1536b7 = util.PathGzip(__obf_fcb4e40318c7e7bf)
		if __obf_874f1b3a1c1536b7 != nil {
			__obf_0f4b5fbdce447b1b = append(__obf_0f4b5fbdce447b1b, fmt.Errorf("compress %s: %w", __obf_3b91508eac98d6e5.Name(), __obf_874f1b3a1c1536b7))
		} else {
			if __obf_874f1b3a1c1536b7 = os.Remove(__obf_fcb4e40318c7e7bf); __obf_874f1b3a1c1536b7 != nil {
				__obf_0f4b5fbdce447b1b = append(__obf_0f4b5fbdce447b1b, fmt.Errorf("remove %s: %w", __obf_fcb4e40318c7e7bf, __obf_874f1b3a1c1536b7))
			}
		}
	}

	return errors.Join(__obf_0f4b5fbdce447b1b...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_fc2298c77aff69e0 *LogWriter) __obf_25a15534f674e145() ([]__obf_00900b790cede2ed, error) {
	__obf_7d212fadcdc70a73, __obf_874f1b3a1c1536b7 := os.ReadDir(__obf_fc2298c77aff69e0.__obf_d764b8d55a7df3d7)
	if __obf_874f1b3a1c1536b7 != nil {
		return nil, __obf_874f1b3a1c1536b7
	}
	var (
		__obf_a22b30d0e178bdad                         time.Time
		__obf_4c2b12a769929582                         []__obf_00900b790cede2ed
		__obf_fc660d113a218a0a, __obf_d2a121bdff49e600 = util.SplitFilename(__obf_fc2298c77aff69e0.__obf_b9186a3395106d9b.LogPath)
	)
	__obf_fc660d113a218a0a += "-"
	for _, __obf_3b91508eac98d6e5 := range __obf_7d212fadcdc70a73 {
		if __obf_3b91508eac98d6e5.IsDir() {
			continue
		}
		if __obf_a22b30d0e178bdad, __obf_874f1b3a1c1536b7 = __obf_a66a3699375e3ff9(__obf_e6b7cf7168f3666a, __obf_3b91508eac98d6e5.Name(), __obf_fc660d113a218a0a, __obf_d2a121bdff49e600); __obf_874f1b3a1c1536b7 == nil {
			__obf_4c2b12a769929582 = append(__obf_4c2b12a769929582, __obf_00900b790cede2ed{__obf_3b91508eac98d6e5, __obf_a22b30d0e178bdad})
			continue
		}
		if __obf_a22b30d0e178bdad, __obf_874f1b3a1c1536b7 = __obf_a66a3699375e3ff9(__obf_e6b7cf7168f3666a, __obf_3b91508eac98d6e5.Name(), __obf_fc660d113a218a0a, __obf_d2a121bdff49e600+__obf_128408006183829e); __obf_874f1b3a1c1536b7 == nil {
			__obf_4c2b12a769929582 = append(__obf_4c2b12a769929582, __obf_00900b790cede2ed{__obf_3b91508eac98d6e5, __obf_a22b30d0e178bdad})
			continue
		}
	}
	sort.Sort(__obf_3088bd0e1cb80d03(__obf_4c2b12a769929582))
	return __obf_4c2b12a769929582, nil
}

type __obf_00900b790cede2ed struct {
	os.DirEntry
	__obf_6bc488781d67451d time.Time
}

type __obf_3088bd0e1cb80d03 []__obf_00900b790cede2ed

func (__obf_90113ee02640e18d __obf_3088bd0e1cb80d03) Less(__obf_33b15be3028c450c, __obf_54e6457239f5911f int) bool {
	return __obf_90113ee02640e18d[__obf_33b15be3028c450c].__obf_6bc488781d67451d.After(__obf_90113ee02640e18d[__obf_54e6457239f5911f].__obf_6bc488781d67451d)
}

func (__obf_90113ee02640e18d __obf_3088bd0e1cb80d03) Swap(__obf_33b15be3028c450c, __obf_54e6457239f5911f int) {
	__obf_90113ee02640e18d[__obf_33b15be3028c450c], __obf_90113ee02640e18d[__obf_54e6457239f5911f] = __obf_90113ee02640e18d[__obf_54e6457239f5911f], __obf_90113ee02640e18d[__obf_33b15be3028c450c]
}

func (__obf_90113ee02640e18d __obf_3088bd0e1cb80d03) Len() int {
	return len(__obf_90113ee02640e18d)
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
func __obf_a66a3699375e3ff9(__obf_e6b7cf7168f3666a, __obf_cedcf659ac37eb39, __obf_fc660d113a218a0a, __obf_d2a121bdff49e600 string) (time.Time, error) {
	if !strings.HasPrefix(__obf_cedcf659ac37eb39, __obf_fc660d113a218a0a) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_fc660d113a218a0a, __obf_cedcf659ac37eb39)
	}
	if !strings.HasSuffix(__obf_cedcf659ac37eb39, __obf_d2a121bdff49e600) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_d2a121bdff49e600, __obf_cedcf659ac37eb39)
	}
	__obf_b0afc1cc726de944 := __obf_cedcf659ac37eb39[len(__obf_fc660d113a218a0a) : len(__obf_cedcf659ac37eb39)-len(__obf_d2a121bdff49e600)]
	return time.Parse(__obf_e6b7cf7168f3666a, __obf_b0afc1cc726de944)
}
