// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_55b2846f9c90acae

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
	__obf_36c1f8151de4de15 = "2006-01-02_15:04:05.000"
	__obf_a7d3f0bb8643cd4d = ".gz"
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
	__obf_4b7d11bba726592a Option
	__obf_52c640a2618c9eec *os.File
	__obf_468bfd80a31af6fb int64
	__obf_04a94bed5a1cb1cd string
	__obf_a3a5f3b36cd29c37 sync.Mutex
	__obf_02ba292e4a3b324f sync.Once
	__obf_7076f7fa25e3b229 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_4b7d11bba726592a Option) *LogWriter {
	__obf_04a94bed5a1cb1cd := filepath.Dir(__obf_4b7d11bba726592a.LogPath)
	__obf_4b7d11bba726592a.
		MaxAge = __obf_4b7d11bba726592a.MaxAge * time.Hour * 24
	__obf_4b7d11bba726592a.
		MaxSize = __obf_4b7d11bba726592a.MaxSize * 1024 * 1024
	__obf_f103820b242abd88 := &LogWriter{__obf_4b7d11bba726592a: __obf_4b7d11bba726592a, __obf_04a94bed5a1cb1cd: __obf_04a94bed5a1cb1cd}

	return __obf_f103820b242abd88
}

// Write writes content into file.
func (__obf_266274e5ced83b4b *LogWriter) Write(__obf_d87391d8d7fca17b []byte) (__obf_24fea63b3182adf3 int, __obf_1b8178764dcad986 error) {
	__obf_266274e5ced83b4b.__obf_a3a5f3b36cd29c37.
		Lock()
	defer __obf_266274e5ced83b4b.__obf_a3a5f3b36cd29c37.Unlock()
	__obf_f1b8f2981778bcbc := int64(len(__obf_d87391d8d7fca17b))

	if __obf_266274e5ced83b4b.__obf_52c640a2618c9eec == nil {
		if __obf_1b8178764dcad986 = __obf_266274e5ced83b4b.__obf_817ef7d1560c8b50(__obf_f1b8f2981778bcbc); __obf_1b8178764dcad986 != nil {
			return
		}
	}
	if __obf_266274e5ced83b4b.__obf_2a1426bf9f736b74(__obf_f1b8f2981778bcbc + __obf_266274e5ced83b4b.__obf_468bfd80a31af6fb) {
		if __obf_1b8178764dcad986 = __obf_266274e5ced83b4b.__obf_034364590674dd70(); __obf_1b8178764dcad986 != nil {
			return
		}
	}
	__obf_24fea63b3182adf3, __obf_1b8178764dcad986 = __obf_266274e5ced83b4b.__obf_52c640a2618c9eec.Write(__obf_d87391d8d7fca17b)
	__obf_266274e5ced83b4b.__obf_468bfd80a31af6fb += int64(__obf_24fea63b3182adf3)
	return
}

// Close closes file resource
func (__obf_266274e5ced83b4b *LogWriter) Close() error {
	__obf_266274e5ced83b4b.__obf_a3a5f3b36cd29c37.
		Lock()
	defer __obf_266274e5ced83b4b.__obf_a3a5f3b36cd29c37.Unlock()
	return __obf_266274e5ced83b4b.close()
}

func (__obf_266274e5ced83b4b *LogWriter) close() error {
	if __obf_266274e5ced83b4b.__obf_52c640a2618c9eec == nil {
		return nil
	}
	__obf_1b8178764dcad986 := __obf_266274e5ced83b4b.__obf_52c640a2618c9eec.Close()
	if __obf_1b8178764dcad986 != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_1b8178764dcad986)
	}
	__obf_266274e5ced83b4b.__obf_52c640a2618c9eec = nil
	__obf_266274e5ced83b4b.__obf_468bfd80a31af6fb = 0
	return nil
}

func (__obf_266274e5ced83b4b *LogWriter) __obf_2a1426bf9f736b74(__obf_468bfd80a31af6fb int64) bool {
	return __obf_266274e5ced83b4b.__obf_4b7d11bba726592a.MaxSize > 0 && __obf_468bfd80a31af6fb > __obf_266274e5ced83b4b.

		// openFile 负责打开日志文件以供写入。
		// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
		// 如果文件存在，它会判断是否需要轮换日志文件。
		__obf_4b7d11bba726592a.MaxSize
}

func (__obf_266274e5ced83b4b *LogWriter) __obf_817ef7d1560c8b50(__obf_b7755e927c858a14 int64) error {
	__obf_64560a22e5626c96, __obf_1b8178764dcad986 := os.Stat(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath)
	if os.IsNotExist(__obf_1b8178764dcad986) {
		return __obf_266274e5ced83b4b.__obf_31c8790189327f1b()
	}
	if __obf_1b8178764dcad986 != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_1b8178764dcad986)
	}
	if __obf_266274e5ced83b4b.__obf_2a1426bf9f736b74(__obf_64560a22e5626c96.Size() + __obf_b7755e927c858a14) {
		return __obf_266274e5ced83b4b.__obf_034364590674dd70()
	}
	__obf_52c640a2618c9eec, __obf_1b8178764dcad986 := os.OpenFile(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_1b8178764dcad986 != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_1b8178764dcad986)
	}
	__obf_266274e5ced83b4b.__obf_52c640a2618c9eec = __obf_52c640a2618c9eec
	__obf_266274e5ced83b4b.

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
		__obf_468bfd80a31af6fb = __obf_64560a22e5626c96.Size()
	return nil
}

func (__obf_266274e5ced83b4b *LogWriter) __obf_31c8790189327f1b() error {
	if __obf_1b8178764dcad986 := os.MkdirAll(__obf_266274e5ced83b4b.__obf_04a94bed5a1cb1cd, os.ModePerm); __obf_1b8178764dcad986 != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_1b8178764dcad986)
	}
	__obf_81934ff7147246b2 := os.FileMode(0644)
	__obf_64560a22e5626c96, __obf_1b8178764dcad986 := os.Stat(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath)
	if __obf_1b8178764dcad986 == nil {
		__obf_81934ff7147246b2 = __obf_64560a22e5626c96.Mode()
		__obf_a220c97ffd18e618, __obf_82f00d564aa3355d := util.SplitFilename(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath)
		__obf_ba405ee15849d7c4 := time.Now().Format(__obf_36c1f8151de4de15)
		__obf_38ed478343e6bdd7 := filepath.Join(__obf_266274e5ced83b4b.__obf_04a94bed5a1cb1cd, fmt.Sprintf("%s-%s%s", __obf_a220c97ffd18e618, __obf_ba405ee15849d7c4, __obf_82f00d564aa3355d))

		if __obf_1b8178764dcad986 = os.Rename(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath, __obf_38ed478343e6bdd7); __obf_1b8178764dcad986 != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_38ed478343e6bdd7, __obf_1b8178764dcad986)
		}
	}
	__obf_8811e5e5ac892b93, __obf_1b8178764dcad986 := os.OpenFile(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_81934ff7147246b2)
	if __obf_1b8178764dcad986 != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath, __obf_1b8178764dcad986)
	}
	__obf_266274e5ced83b4b.__obf_52c640a2618c9eec = __obf_8811e5e5ac892b93
	__obf_266274e5ced83b4b.

		// rotate 处理 LogWriter 的日志文件轮换过程。
		// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
		// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
		// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
		//
		// 如果关闭当前文件或打开新文件失败，则返回错误。
		__obf_468bfd80a31af6fb = 0
	return nil
}

func (__obf_266274e5ced83b4b *LogWriter) __obf_034364590674dd70() error {
	if __obf_1b8178764dcad986 := __obf_266274e5ced83b4b.close(); __obf_1b8178764dcad986 != nil {
		return __obf_1b8178764dcad986
	}
	if __obf_1b8178764dcad986 := __obf_266274e5ced83b4b.__obf_31c8790189327f1b(); __obf_1b8178764dcad986 != nil {
		return __obf_1b8178764dcad986
	}
	__obf_266274e5ced83b4b.__obf_02ba292e4a3b324f.
		Do(func() {
			__obf_266274e5ced83b4b.__obf_7076f7fa25e3b229 = make(chan bool, 1)
			go func() {
				for range __obf_266274e5ced83b4b.__obf_7076f7fa25e3b229 {
					if __obf_2ea4a7bee9be9fe3 := __obf_266274e5ced83b4b.__obf_ebdfbc6b870ebc7d(); __obf_2ea4a7bee9be9fe3 != nil {
						println(__obf_2ea4a7bee9be9fe3.Error())
					}
				}
			}()
		})
	select {
	case __obf_266274e5ced83b4b.__obf_7076f7fa25e3b229 <- true:
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
func (__obf_266274e5ced83b4b *LogWriter) __obf_ebdfbc6b870ebc7d() error {
	if __obf_266274e5ced83b4b.__obf_4b7d11bba726592a.MaxSize == 0 {
		return nil
	}
	__obf_8a0586b86615a4ad, __obf_1b8178764dcad986 := __obf_266274e5ced83b4b.__obf_59a1423940106f55()
	if __obf_1b8178764dcad986 != nil {
		return __obf_1b8178764dcad986
	}
	var __obf_fab0c2129744aaed, __obf_31b558133f3fd78e []__obf_f4d6b7186f0ccacc
	if __obf_266274e5ced83b4b.__obf_4b7d11bba726592a.MaxAge > 0 {
		__obf_c4604952f5b0dec1 := time.Now().Add(-1 * __obf_266274e5ced83b4b.__obf_4b7d11bba726592a.MaxAge)
		for _, __obf_8811e5e5ac892b93 := range __obf_8a0586b86615a4ad {
			if __obf_8811e5e5ac892b93.__obf_01bcd8320974a9d0.Before(__obf_c4604952f5b0dec1) {
				__obf_fab0c2129744aaed = append(__obf_fab0c2129744aaed, __obf_8811e5e5ac892b93)
			} else {
				__obf_31b558133f3fd78e = append(__obf_31b558133f3fd78e, __obf_8811e5e5ac892b93)
			}
		}
	}

	var __obf_a5f1bb8c09f8e41c []error
	for _, __obf_8811e5e5ac892b93 := range __obf_fab0c2129744aaed {
		if __obf_1b8178764dcad986 = os.Remove(filepath.Join(__obf_266274e5ced83b4b.__obf_04a94bed5a1cb1cd, __obf_8811e5e5ac892b93.Name())); __obf_1b8178764dcad986 != nil {
			__obf_a5f1bb8c09f8e41c = append(__obf_a5f1bb8c09f8e41c, fmt.Errorf("remove %s: %w", __obf_8811e5e5ac892b93.Name(), __obf_1b8178764dcad986))
		}
	}
	if len(__obf_a5f1bb8c09f8e41c) > 0 {
		return errors.Join(__obf_a5f1bb8c09f8e41c...)
	}

	for _, __obf_8811e5e5ac892b93 := range __obf_31b558133f3fd78e {
		if strings.HasSuffix(__obf_8811e5e5ac892b93.Name(), __obf_a7d3f0bb8643cd4d) {
			continue
		}
		__obf_0d9e6fb4b39f9fe3 := filepath.Join(__obf_266274e5ced83b4b.__obf_04a94bed5a1cb1cd, __obf_8811e5e5ac892b93.Name())
		__obf_1b8178764dcad986 = util.PathGzip(__obf_0d9e6fb4b39f9fe3)
		if __obf_1b8178764dcad986 != nil {
			__obf_a5f1bb8c09f8e41c = append(__obf_a5f1bb8c09f8e41c, fmt.Errorf("compress %s: %w", __obf_8811e5e5ac892b93.Name(), __obf_1b8178764dcad986))
		} else {
			if __obf_1b8178764dcad986 = os.Remove(__obf_0d9e6fb4b39f9fe3); __obf_1b8178764dcad986 != nil {
				__obf_a5f1bb8c09f8e41c = append(__obf_a5f1bb8c09f8e41c, fmt.Errorf("remove %s: %w", __obf_0d9e6fb4b39f9fe3, __obf_1b8178764dcad986))
			}
		}
	}

	return errors.Join(__obf_a5f1bb8c09f8e41c...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_266274e5ced83b4b *LogWriter) __obf_59a1423940106f55() ([]__obf_f4d6b7186f0ccacc, error) {
	__obf_8a0586b86615a4ad, __obf_1b8178764dcad986 := os.ReadDir(__obf_266274e5ced83b4b.__obf_04a94bed5a1cb1cd)
	if __obf_1b8178764dcad986 != nil {
		return nil, __obf_1b8178764dcad986
	}
	var (
		__obf_ba405ee15849d7c4                         time.Time
		__obf_7f2cc9cd77036314                         []__obf_f4d6b7186f0ccacc
		__obf_a220c97ffd18e618, __obf_82f00d564aa3355d = util.SplitFilename(__obf_266274e5ced83b4b.__obf_4b7d11bba726592a.LogPath)
	)
	__obf_a220c97ffd18e618 += "-"
	for _, __obf_8811e5e5ac892b93 := range __obf_8a0586b86615a4ad {
		if __obf_8811e5e5ac892b93.IsDir() {
			continue
		}
		if __obf_ba405ee15849d7c4, __obf_1b8178764dcad986 = __obf_599574253d3c59df(__obf_36c1f8151de4de15, __obf_8811e5e5ac892b93.Name(), __obf_a220c97ffd18e618, __obf_82f00d564aa3355d); __obf_1b8178764dcad986 == nil {
			__obf_7f2cc9cd77036314 = append(__obf_7f2cc9cd77036314, __obf_f4d6b7186f0ccacc{__obf_8811e5e5ac892b93, __obf_ba405ee15849d7c4})
			continue
		}
		if __obf_ba405ee15849d7c4, __obf_1b8178764dcad986 = __obf_599574253d3c59df(__obf_36c1f8151de4de15, __obf_8811e5e5ac892b93.Name(), __obf_a220c97ffd18e618, __obf_82f00d564aa3355d+__obf_a7d3f0bb8643cd4d); __obf_1b8178764dcad986 == nil {
			__obf_7f2cc9cd77036314 = append(__obf_7f2cc9cd77036314, __obf_f4d6b7186f0ccacc{__obf_8811e5e5ac892b93, __obf_ba405ee15849d7c4})
			continue
		}
	}
	sort.Sort(__obf_c99c57ec5983b9a4(__obf_7f2cc9cd77036314))
	return __obf_7f2cc9cd77036314, nil
}

type __obf_f4d6b7186f0ccacc struct {
	os.DirEntry
	__obf_01bcd8320974a9d0 time.Time
}

type __obf_c99c57ec5983b9a4 []__obf_f4d6b7186f0ccacc

func (__obf_99a91eaedf7cdbfa __obf_c99c57ec5983b9a4) Less(__obf_062426eef3e13780, __obf_2c18afdddbe1ac76 int) bool {
	return __obf_99a91eaedf7cdbfa[__obf_062426eef3e13780].__obf_01bcd8320974a9d0.After(__obf_99a91eaedf7cdbfa[__obf_2c18afdddbe1ac76].__obf_01bcd8320974a9d0)
}

func (__obf_99a91eaedf7cdbfa __obf_c99c57ec5983b9a4) Swap(__obf_062426eef3e13780, __obf_2c18afdddbe1ac76 int) {
	__obf_99a91eaedf7cdbfa[__obf_062426eef3e13780], __obf_99a91eaedf7cdbfa[__obf_2c18afdddbe1ac76] = __obf_99a91eaedf7cdbfa[__obf_2c18afdddbe1ac76], __obf_99a91eaedf7cdbfa[__obf_062426eef3e13780]
}

func (__obf_99a91eaedf7cdbfa __obf_c99c57ec5983b9a4) Len() int {
	return len(__obf_99a91eaedf7cdbfa)
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
func __obf_599574253d3c59df(__obf_36c1f8151de4de15, __obf_ed379217e15668dd, __obf_a220c97ffd18e618, __obf_82f00d564aa3355d string) (time.Time, error) {
	if !strings.HasPrefix(__obf_ed379217e15668dd, __obf_a220c97ffd18e618) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_a220c97ffd18e618, __obf_ed379217e15668dd)
	}
	if !strings.HasSuffix(__obf_ed379217e15668dd, __obf_82f00d564aa3355d) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_82f00d564aa3355d, __obf_ed379217e15668dd)
	}
	__obf_23167c3706db5526 := __obf_ed379217e15668dd[len(__obf_a220c97ffd18e618) : len(__obf_ed379217e15668dd)-len(__obf_82f00d564aa3355d)]
	return time.Parse(__obf_36c1f8151de4de15, __obf_23167c3706db5526)
}
