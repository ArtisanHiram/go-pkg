// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_5efebb632ee092eb

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
	__obf_1065ec3412b827ec = "2006-01-02_15:04:05.000"
	__obf_c70831457306bb9e = ".gz"
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
	__obf_72ddced8d8c596ca Option
	__obf_74a95368ac22e29f *os.File
	__obf_a97b94dc369bacdb int64
	__obf_93a48dd857f4e9f8 string
	__obf_411f94c9b8381110 sync.Mutex
	__obf_0fd1dca82c62e2e3 sync.Once
	__obf_37c1640fffc0fa22 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_72ddced8d8c596ca Option) *LogWriter {
	__obf_93a48dd857f4e9f8 := filepath.Dir(__obf_72ddced8d8c596ca.LogPath)
	__obf_72ddced8d8c596ca.
		MaxAge = __obf_72ddced8d8c596ca.MaxAge * time.Hour * 24
	__obf_72ddced8d8c596ca.
		MaxSize = __obf_72ddced8d8c596ca.MaxSize * 1024 * 1024
	__obf_a14c3fd62650bb48 := &LogWriter{__obf_72ddced8d8c596ca: __obf_72ddced8d8c596ca, __obf_93a48dd857f4e9f8: __obf_93a48dd857f4e9f8}

	return __obf_a14c3fd62650bb48
}

// Write writes content into file.
func (__obf_442eeb2efef80bcb *LogWriter) Write(__obf_5241727993c46127 []byte) (__obf_7548592a71380bec int, __obf_6ab41df571d54712 error) {
	__obf_442eeb2efef80bcb.__obf_411f94c9b8381110.
		Lock()
	defer __obf_442eeb2efef80bcb.__obf_411f94c9b8381110.Unlock()
	__obf_45023b35f3860fa4 := int64(len(__obf_5241727993c46127))

	if __obf_442eeb2efef80bcb.__obf_74a95368ac22e29f == nil {
		if __obf_6ab41df571d54712 = __obf_442eeb2efef80bcb.__obf_decd1531315bbee7(__obf_45023b35f3860fa4); __obf_6ab41df571d54712 != nil {
			return
		}
	}
	if __obf_442eeb2efef80bcb.__obf_01b2c8787e321b30(__obf_45023b35f3860fa4 + __obf_442eeb2efef80bcb.__obf_a97b94dc369bacdb) {
		if __obf_6ab41df571d54712 = __obf_442eeb2efef80bcb.__obf_9cb70ad1fa83cdfd(); __obf_6ab41df571d54712 != nil {
			return
		}
	}
	__obf_7548592a71380bec, __obf_6ab41df571d54712 = __obf_442eeb2efef80bcb.__obf_74a95368ac22e29f.Write(__obf_5241727993c46127)
	__obf_442eeb2efef80bcb.__obf_a97b94dc369bacdb += int64(__obf_7548592a71380bec)
	return
}

// Close closes file resource
func (__obf_442eeb2efef80bcb *LogWriter) Close() error {
	__obf_442eeb2efef80bcb.__obf_411f94c9b8381110.
		Lock()
	defer __obf_442eeb2efef80bcb.__obf_411f94c9b8381110.Unlock()
	return __obf_442eeb2efef80bcb.close()
}

func (__obf_442eeb2efef80bcb *LogWriter) close() error {
	if __obf_442eeb2efef80bcb.__obf_74a95368ac22e29f == nil {
		return nil
	}
	__obf_6ab41df571d54712 := __obf_442eeb2efef80bcb.__obf_74a95368ac22e29f.Close()
	if __obf_6ab41df571d54712 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_6ab41df571d54712)
	}
	__obf_442eeb2efef80bcb.__obf_74a95368ac22e29f = nil
	__obf_442eeb2efef80bcb.__obf_a97b94dc369bacdb = 0
	return nil
}

func (__obf_442eeb2efef80bcb *LogWriter) __obf_01b2c8787e321b30(__obf_a97b94dc369bacdb int64) bool {
	return __obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.MaxSize > 0 && __obf_a97b94dc369bacdb > __obf_442eeb2efef80bcb.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_72ddced8d8c596ca.MaxSize
}

func (__obf_442eeb2efef80bcb *LogWriter) __obf_decd1531315bbee7(__obf_25cfed70f8e2ea77 int64) error {
	__obf_955dd122c8ddefa0, __obf_6ab41df571d54712 := os.Stat(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath)
	if os.IsNotExist(__obf_6ab41df571d54712) {
		return __obf_442eeb2efef80bcb.__obf_c40156e09f43e0a2()
	}
	if __obf_6ab41df571d54712 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_6ab41df571d54712)
	}
	if __obf_442eeb2efef80bcb.__obf_01b2c8787e321b30(__obf_955dd122c8ddefa0.Size() + __obf_25cfed70f8e2ea77) {
		return __obf_442eeb2efef80bcb.__obf_9cb70ad1fa83cdfd()
	}
	__obf_74a95368ac22e29f, __obf_6ab41df571d54712 := os.OpenFile(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_6ab41df571d54712 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_6ab41df571d54712)
	}
	__obf_442eeb2efef80bcb.__obf_74a95368ac22e29f = __obf_74a95368ac22e29f
	__obf_442eeb2efef80bcb.

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
		__obf_a97b94dc369bacdb = __obf_955dd122c8ddefa0.Size()
	return nil
}

func (__obf_442eeb2efef80bcb *LogWriter) __obf_c40156e09f43e0a2() error {
	if __obf_6ab41df571d54712 := os.MkdirAll(__obf_442eeb2efef80bcb.__obf_93a48dd857f4e9f8, os.ModePerm); __obf_6ab41df571d54712 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_6ab41df571d54712)
	}
	__obf_df8b0698dab9a53a := os.FileMode(0644)
	__obf_955dd122c8ddefa0, __obf_6ab41df571d54712 := os.Stat(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath)
	if __obf_6ab41df571d54712 == nil {
		__obf_df8b0698dab9a53a = __obf_955dd122c8ddefa0.Mode()
		__obf_95a5e54f64e06cb3, __obf_04ee4259ffd817ed := util.SplitFilename(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath)
		__obf_fcc3098738fa2b22 := time.Now().Format(__obf_1065ec3412b827ec)
		__obf_d16add31790654db := filepath.Join(__obf_442eeb2efef80bcb.__obf_93a48dd857f4e9f8, fmt.Sprintf("%s-%s%s", __obf_95a5e54f64e06cb3, __obf_fcc3098738fa2b22, __obf_04ee4259ffd817ed))

		if __obf_6ab41df571d54712 = os.Rename(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath, __obf_d16add31790654db); __obf_6ab41df571d54712 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_d16add31790654db, __obf_6ab41df571d54712)
		}
	}
	__obf_8a60791d671f4d91, __obf_6ab41df571d54712 := os.OpenFile(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_df8b0698dab9a53a)
	if __obf_6ab41df571d54712 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath, __obf_6ab41df571d54712)
	}
	__obf_442eeb2efef80bcb.__obf_74a95368ac22e29f = __obf_8a60791d671f4d91
	__obf_442eeb2efef80bcb.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_a97b94dc369bacdb = 0
	return nil
}

func (__obf_442eeb2efef80bcb *LogWriter) __obf_9cb70ad1fa83cdfd() error {
	if __obf_6ab41df571d54712 := __obf_442eeb2efef80bcb.close(); __obf_6ab41df571d54712 != nil {
		return __obf_6ab41df571d54712
	}
	if __obf_6ab41df571d54712 := __obf_442eeb2efef80bcb.__obf_c40156e09f43e0a2(); __obf_6ab41df571d54712 != nil {
		return __obf_6ab41df571d54712
	}
	__obf_442eeb2efef80bcb.__obf_0fd1dca82c62e2e3.
		Do(func() {
			__obf_442eeb2efef80bcb.__obf_37c1640fffc0fa22 = make(chan bool, 1)
			go func() {
				for range __obf_442eeb2efef80bcb.__obf_37c1640fffc0fa22 {
					if __obf_569f70144177db09 := __obf_442eeb2efef80bcb.__obf_d748f9612772b751(); __obf_569f70144177db09 != nil {
						println(__obf_569f70144177db09.Error())
					}
				}
			}()
		})
	select {
	case __obf_442eeb2efef80bcb.__obf_37c1640fffc0fa22 <- true:
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
func (__obf_442eeb2efef80bcb *LogWriter) __obf_d748f9612772b751() error {
	if __obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.MaxSize == 0 {
		return nil
	}
	__obf_8e1d0a28a6021dc9, __obf_6ab41df571d54712 := __obf_442eeb2efef80bcb.__obf_c7fbd043361d6093()
	if __obf_6ab41df571d54712 != nil {
		return __obf_6ab41df571d54712
	}
	var __obf_f11f1d47bf64eb62, __obf_62eb3329c3c866ad []__obf_7b109107f58439bb
	if __obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.MaxAge > 0 {
		__obf_8f79fc07529cb512 := time.Now().Add(-1 * __obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.MaxAge)
		for _, __obf_8a60791d671f4d91 := range __obf_8e1d0a28a6021dc9 {
			if __obf_8a60791d671f4d91.__obf_dfc07a1e53049328.Before(__obf_8f79fc07529cb512) {
				__obf_f11f1d47bf64eb62 = append(__obf_f11f1d47bf64eb62, __obf_8a60791d671f4d91)
			} else {
				__obf_62eb3329c3c866ad = append(__obf_62eb3329c3c866ad, __obf_8a60791d671f4d91)
			}
		}
	}

	var __obf_272966bb3e39763d []error
	for _, __obf_8a60791d671f4d91 := range __obf_f11f1d47bf64eb62 {
		if __obf_6ab41df571d54712 = os.Remove(filepath.Join(__obf_442eeb2efef80bcb.__obf_93a48dd857f4e9f8, __obf_8a60791d671f4d91.Name())); __obf_6ab41df571d54712 != nil {
			__obf_272966bb3e39763d = append(__obf_272966bb3e39763d, fmt.Errorf("remove %s: %w", __obf_8a60791d671f4d91.Name(), __obf_6ab41df571d54712))
		}
	}
	if len(__obf_272966bb3e39763d) > 0 {
		return errors.Join(__obf_272966bb3e39763d...)
	}

	for _, __obf_8a60791d671f4d91 := range __obf_62eb3329c3c866ad {
		if strings.HasSuffix(__obf_8a60791d671f4d91.Name(), __obf_c70831457306bb9e) {
			continue
		}
		__obf_60c46439cbee8d6e := filepath.Join(__obf_442eeb2efef80bcb.__obf_93a48dd857f4e9f8, __obf_8a60791d671f4d91.Name())
		__obf_6ab41df571d54712 = util.PathGzip(__obf_60c46439cbee8d6e)
		if __obf_6ab41df571d54712 != nil {
			__obf_272966bb3e39763d = append(__obf_272966bb3e39763d, fmt.Errorf("compress %s: %w", __obf_8a60791d671f4d91.Name(), __obf_6ab41df571d54712))
		} else {
			if __obf_6ab41df571d54712 = os.Remove(__obf_60c46439cbee8d6e); __obf_6ab41df571d54712 != nil {
				__obf_272966bb3e39763d = append(__obf_272966bb3e39763d, fmt.Errorf("remove %s: %w", __obf_60c46439cbee8d6e, __obf_6ab41df571d54712))
			}
		}
	}

	return errors.Join(__obf_272966bb3e39763d...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_442eeb2efef80bcb *LogWriter) __obf_c7fbd043361d6093() ([]__obf_7b109107f58439bb, error) {
	__obf_8e1d0a28a6021dc9, __obf_6ab41df571d54712 := os.ReadDir(__obf_442eeb2efef80bcb.__obf_93a48dd857f4e9f8)
	if __obf_6ab41df571d54712 != nil {
		return nil, __obf_6ab41df571d54712
	}
	var (
		__obf_fcc3098738fa2b22                         time.Time
		__obf_547ff96a8b482084                         []__obf_7b109107f58439bb
		__obf_95a5e54f64e06cb3, __obf_04ee4259ffd817ed = util.SplitFilename(__obf_442eeb2efef80bcb.__obf_72ddced8d8c596ca.LogPath)
	)
	__obf_95a5e54f64e06cb3 += "-"
	for _, __obf_8a60791d671f4d91 := range __obf_8e1d0a28a6021dc9 {
		if __obf_8a60791d671f4d91.IsDir() {
			continue
		}
		if __obf_fcc3098738fa2b22, __obf_6ab41df571d54712 = __obf_77d64450321d025a(__obf_1065ec3412b827ec, __obf_8a60791d671f4d91.Name(), __obf_95a5e54f64e06cb3, __obf_04ee4259ffd817ed); __obf_6ab41df571d54712 == nil {
			__obf_547ff96a8b482084 = append(__obf_547ff96a8b482084, __obf_7b109107f58439bb{__obf_8a60791d671f4d91, __obf_fcc3098738fa2b22})
			continue
		}
		if __obf_fcc3098738fa2b22, __obf_6ab41df571d54712 = __obf_77d64450321d025a(__obf_1065ec3412b827ec, __obf_8a60791d671f4d91.Name(), __obf_95a5e54f64e06cb3, __obf_04ee4259ffd817ed+__obf_c70831457306bb9e); __obf_6ab41df571d54712 == nil {
			__obf_547ff96a8b482084 = append(__obf_547ff96a8b482084, __obf_7b109107f58439bb{__obf_8a60791d671f4d91, __obf_fcc3098738fa2b22})
			continue
		}
	}
	sort.Sort(__obf_e4b9a303b19d811d(__obf_547ff96a8b482084))
	return __obf_547ff96a8b482084, nil
}

type __obf_7b109107f58439bb struct {
	os.DirEntry
	__obf_dfc07a1e53049328 time.Time
}

type __obf_e4b9a303b19d811d []__obf_7b109107f58439bb

func (__obf_1daad46a9ffc02be __obf_e4b9a303b19d811d) Less(__obf_89508e929b14f04c, __obf_dc5b8090d872df50 int) bool {
	return __obf_1daad46a9ffc02be[__obf_89508e929b14f04c].__obf_dfc07a1e53049328.After(__obf_1daad46a9ffc02be[__obf_dc5b8090d872df50].__obf_dfc07a1e53049328)
}

func (__obf_1daad46a9ffc02be __obf_e4b9a303b19d811d) Swap(__obf_89508e929b14f04c, __obf_dc5b8090d872df50 int) {
	__obf_1daad46a9ffc02be[__obf_89508e929b14f04c], __obf_1daad46a9ffc02be[__obf_dc5b8090d872df50] = __obf_1daad46a9ffc02be[__obf_dc5b8090d872df50], __obf_1daad46a9ffc02be[__obf_89508e929b14f04c]
}

func (__obf_1daad46a9ffc02be __obf_e4b9a303b19d811d) Len() int {
	return len(__obf_1daad46a9ffc02be)
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
func __obf_77d64450321d025a(__obf_1065ec3412b827ec, __obf_183892b8a1cc298c, __obf_95a5e54f64e06cb3, __obf_04ee4259ffd817ed string) (time.Time, error) {
	if !strings.HasPrefix(__obf_183892b8a1cc298c, __obf_95a5e54f64e06cb3) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_95a5e54f64e06cb3, __obf_183892b8a1cc298c)
	}
	if !strings.HasSuffix(__obf_183892b8a1cc298c, __obf_04ee4259ffd817ed) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_04ee4259ffd817ed, __obf_183892b8a1cc298c)
	}
	__obf_74f4fd2040ae1c40 := __obf_183892b8a1cc298c[len(__obf_95a5e54f64e06cb3) : len(__obf_183892b8a1cc298c)-len(__obf_04ee4259ffd817ed)]
	return time.Parse(__obf_1065ec3412b827ec, __obf_74f4fd2040ae1c40)
}
