// 参考1: gopkg.in/natefinch/lumberjack.v2
// 参考2: https://github.com/gggwvg/logrotate

package __obf_290168e300ae1c1e

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
	__obf_99ac71c2db085776 = "2006-01-02_15:04:05.000"
	__obf_ba8e597c54e63c15 = ".gz"
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
	__obf_6f64510fd7abc1a8 Option
	__obf_0da9f7ddb1f44779 *os.File
	__obf_63cb38940d8d4849 int64
	__obf_da94489893334ea7 string
	__obf_d81c4629ffd25ab4 sync.Mutex
	__obf_271b3d5a10b94931 sync.Once
	__obf_1bd03746261c40c1 chan bool
}

// NewLogWriter creates a new logger
func NewLogWriter(__obf_6f64510fd7abc1a8 Option) *LogWriter {
	__obf_da94489893334ea7 := filepath.Dir(__obf_6f64510fd7abc1a8.LogPath)

	__obf_6f64510fd7abc1a8.MaxAge = __obf_6f64510fd7abc1a8.MaxAge * time.Hour * 24
	__obf_6f64510fd7abc1a8.MaxSize = __obf_6f64510fd7abc1a8.MaxSize * 1024 * 1024

	__obf_1e61fec0a73f1220 := &LogWriter{
		__obf_6f64510fd7abc1a8: __obf_6f64510fd7abc1a8,
		__obf_da94489893334ea7: __obf_da94489893334ea7,
	}

	return __obf_1e61fec0a73f1220
}

// Write writes content into file.
func (__obf_b802042ac7a13859 *LogWriter) Write(__obf_5dca8f0e76592e75 []byte) (__obf_05cc6afc89cce734 int, __obf_9c607119691e49da error) {
	__obf_b802042ac7a13859.__obf_d81c4629ffd25ab4.Lock()
	defer __obf_b802042ac7a13859.__obf_d81c4629ffd25ab4.Unlock()
	__obf_2a0f4271a1bb74ba := int64(len(__obf_5dca8f0e76592e75))

	if __obf_b802042ac7a13859.__obf_0da9f7ddb1f44779 == nil {
		if __obf_9c607119691e49da = __obf_b802042ac7a13859.__obf_760cbd88ff6053cd(__obf_2a0f4271a1bb74ba); __obf_9c607119691e49da != nil {
			return
		}
	}
	if __obf_b802042ac7a13859.__obf_0528a6ab16f01d9d(__obf_2a0f4271a1bb74ba + __obf_b802042ac7a13859.__obf_63cb38940d8d4849) {
		if __obf_9c607119691e49da = __obf_b802042ac7a13859.__obf_8cdedb372a7fff1c(); __obf_9c607119691e49da != nil {
			return
		}
	}
	__obf_05cc6afc89cce734, __obf_9c607119691e49da = __obf_b802042ac7a13859.__obf_0da9f7ddb1f44779.Write(__obf_5dca8f0e76592e75)
	__obf_b802042ac7a13859.__obf_63cb38940d8d4849 += int64(__obf_05cc6afc89cce734)
	return
}

// Close closes file resource
func (__obf_b802042ac7a13859 *LogWriter) Close() error {
	__obf_b802042ac7a13859.__obf_d81c4629ffd25ab4.Lock()
	defer __obf_b802042ac7a13859.__obf_d81c4629ffd25ab4.Unlock()
	return __obf_b802042ac7a13859.close()
}

func (__obf_b802042ac7a13859 *LogWriter) close() error {
	if __obf_b802042ac7a13859.__obf_0da9f7ddb1f44779 == nil {
		return nil
	}
	__obf_9c607119691e49da := __obf_b802042ac7a13859.__obf_0da9f7ddb1f44779.Close()
	if __obf_9c607119691e49da != nil {
		return fmt.Errorf("failed to close log file: %s", __obf_9c607119691e49da)
	}
	__obf_b802042ac7a13859.__obf_0da9f7ddb1f44779 = nil
	__obf_b802042ac7a13859.__obf_63cb38940d8d4849 = 0
	return nil
}

func (__obf_b802042ac7a13859 *LogWriter) __obf_0528a6ab16f01d9d(__obf_63cb38940d8d4849 int64) bool {
	return __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.MaxSize > 0 && __obf_63cb38940d8d4849 > __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.MaxSize
}

// openFile 负责打开日志文件以供写入。
// 它会检查日志文件是否存在，如果不存在则创建一个新文件。
// 如果文件存在，它会判断是否需要轮换日志文件。
func (__obf_b802042ac7a13859 *LogWriter) __obf_760cbd88ff6053cd(__obf_3501b9c640fa5996 int64) error {
	__obf_f83490257f420f6f, __obf_9c607119691e49da := os.Stat(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath)
	if os.IsNotExist(__obf_9c607119691e49da) {
		return __obf_b802042ac7a13859.__obf_1d0c1d15cb6ce0df()
	}
	if __obf_9c607119691e49da != nil {
		return fmt.Errorf("failed to get file info: %w", __obf_9c607119691e49da)
	}
	if __obf_b802042ac7a13859.__obf_0528a6ab16f01d9d(__obf_f83490257f420f6f.Size() + __obf_3501b9c640fa5996) {
		return __obf_b802042ac7a13859.__obf_8cdedb372a7fff1c()
	}
	__obf_0da9f7ddb1f44779, __obf_9c607119691e49da := os.OpenFile(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath, os.O_APPEND|os.O_WRONLY, 0644)
	if __obf_9c607119691e49da != nil {
		return fmt.Errorf("failed to open log file: %w", __obf_9c607119691e49da)
	}
	__obf_b802042ac7a13859.__obf_0da9f7ddb1f44779 = __obf_0da9f7ddb1f44779
	__obf_b802042ac7a13859.__obf_63cb38940d8d4849 = __obf_f83490257f420f6f.Size()
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
func (__obf_b802042ac7a13859 *LogWriter) __obf_1d0c1d15cb6ce0df() error {
	if __obf_9c607119691e49da := os.MkdirAll(__obf_b802042ac7a13859.__obf_da94489893334ea7, os.ModePerm); __obf_9c607119691e49da != nil {
		return fmt.Errorf("can't make directories for new logfile, error(%v)", __obf_9c607119691e49da)
	}
	__obf_ee81586153431bbb := os.FileMode(0644)
	__obf_f83490257f420f6f, __obf_9c607119691e49da := os.Stat(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath)
	if __obf_9c607119691e49da == nil {
		__obf_ee81586153431bbb = __obf_f83490257f420f6f.Mode()
		__obf_841443da5c31487a, __obf_a39c26a8085c5a08 := util.SplitFilename(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath)
		__obf_6d9ef399ff92733e := time.Now().Format(__obf_99ac71c2db085776)
		__obf_ef78a85913c40972 := filepath.Join(__obf_b802042ac7a13859.__obf_da94489893334ea7, fmt.Sprintf("%s-%s%s", __obf_841443da5c31487a, __obf_6d9ef399ff92733e, __obf_a39c26a8085c5a08))

		if __obf_9c607119691e49da = os.Rename(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath, __obf_ef78a85913c40972); __obf_9c607119691e49da != nil {
			return fmt.Errorf("can't archive(%s) error(%v)", __obf_ef78a85913c40972, __obf_9c607119691e49da)
		}
	}
	__obf_df4a5542fb2bd8f8, __obf_9c607119691e49da := os.OpenFile(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, __obf_ee81586153431bbb)
	if __obf_9c607119691e49da != nil {
		return fmt.Errorf("can't open new log file(%s) error(%v)", __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath, __obf_9c607119691e49da)
	}
	__obf_b802042ac7a13859.__obf_0da9f7ddb1f44779 = __obf_df4a5542fb2bd8f8
	__obf_b802042ac7a13859.__obf_63cb38940d8d4849 = 0
	return nil
}

// rotate 处理 LogWriter 的日志文件轮换过程。
// 它首先关闭当前日志文件并打开一个新文件。如果这是第一次执行轮换操作，
// 它会初始化一个通道并启动一个 goroutine 异步处理日志文件归档。
// 此方法确保在通道已被占用时，归档过程可以被触发而不会阻塞。
//
// 如果关闭当前文件或打开新文件失败，则返回错误。

func (__obf_b802042ac7a13859 *LogWriter) __obf_8cdedb372a7fff1c() error {
	if __obf_9c607119691e49da := __obf_b802042ac7a13859.close(); __obf_9c607119691e49da != nil {
		return __obf_9c607119691e49da
	}
	if __obf_9c607119691e49da := __obf_b802042ac7a13859.__obf_1d0c1d15cb6ce0df(); __obf_9c607119691e49da != nil {
		return __obf_9c607119691e49da
	}
	__obf_b802042ac7a13859.__obf_271b3d5a10b94931.Do(func() {
		__obf_b802042ac7a13859.__obf_1bd03746261c40c1 = make(chan bool, 1)
		go func() {
			for range __obf_b802042ac7a13859.__obf_1bd03746261c40c1 {
				if __obf_985f3d188423a86f := __obf_b802042ac7a13859.__obf_ce4c8eabc8c14a62(); __obf_985f3d188423a86f != nil {
					println(__obf_985f3d188423a86f.Error())
				}
			}
		}()
	})
	select {
	case __obf_b802042ac7a13859.__obf_1bd03746261c40c1 <- true:
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
func (__obf_b802042ac7a13859 *LogWriter) __obf_ce4c8eabc8c14a62() error {
	if __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.MaxSize == 0 {
		return nil
	}
	__obf_9fcb047ef26f5dd8, __obf_9c607119691e49da := __obf_b802042ac7a13859.__obf_9f00aef5ecceed2e()
	if __obf_9c607119691e49da != nil {
		return __obf_9c607119691e49da
	}
	var __obf_7c8648aaf779b9ad, __obf_bfb5be0752228e5b []__obf_f86003c7ac3e0dc0
	if __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.MaxAge > 0 {
		__obf_fa44842b3d84db7b := time.Now().Add(-1 * __obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.MaxAge)
		for _, __obf_df4a5542fb2bd8f8 := range __obf_9fcb047ef26f5dd8 {
			if __obf_df4a5542fb2bd8f8.__obf_c3d8b258b75a6c0e.Before(__obf_fa44842b3d84db7b) {
				__obf_7c8648aaf779b9ad = append(__obf_7c8648aaf779b9ad, __obf_df4a5542fb2bd8f8)
			} else {
				__obf_bfb5be0752228e5b = append(__obf_bfb5be0752228e5b, __obf_df4a5542fb2bd8f8)
			}
		}
	}

	var __obf_6825af3c9b349136 []error
	for _, __obf_df4a5542fb2bd8f8 := range __obf_7c8648aaf779b9ad {
		if __obf_9c607119691e49da = os.Remove(filepath.Join(__obf_b802042ac7a13859.__obf_da94489893334ea7, __obf_df4a5542fb2bd8f8.Name())); __obf_9c607119691e49da != nil {
			__obf_6825af3c9b349136 = append(__obf_6825af3c9b349136, fmt.Errorf("remove %s: %w", __obf_df4a5542fb2bd8f8.Name(), __obf_9c607119691e49da))
		}
	}
	if len(__obf_6825af3c9b349136) > 0 {
		return errors.Join(__obf_6825af3c9b349136...)
	}

	for _, __obf_df4a5542fb2bd8f8 := range __obf_bfb5be0752228e5b {
		if strings.HasSuffix(__obf_df4a5542fb2bd8f8.Name(), __obf_ba8e597c54e63c15) {
			continue
		}
		__obf_307376be6e7738bf := filepath.Join(__obf_b802042ac7a13859.__obf_da94489893334ea7, __obf_df4a5542fb2bd8f8.Name())

		__obf_9c607119691e49da = util.PathGzip(__obf_307376be6e7738bf)
		if __obf_9c607119691e49da != nil {
			__obf_6825af3c9b349136 = append(__obf_6825af3c9b349136, fmt.Errorf("compress %s: %w", __obf_df4a5542fb2bd8f8.Name(), __obf_9c607119691e49da))
		} else {
			if __obf_9c607119691e49da = os.Remove(__obf_307376be6e7738bf); __obf_9c607119691e49da != nil {
				__obf_6825af3c9b349136 = append(__obf_6825af3c9b349136, fmt.Errorf("remove %s: %w", __obf_307376be6e7738bf, __obf_9c607119691e49da))
			}
		}
	}

	return errors.Join(__obf_6825af3c9b349136...)
}

// archives 从 LogWriter 的配置中指定的目录中检索日志文件的排序列表。
// 它根据命名模式识别日志文件，该模式包括前缀、时间戳和扩展名。
// 此函数支持未压缩和压缩的日志文件。
//
// 返回值:
// - 一个 logFile 结构体切片，每个结构体包含一个文件及其关联的时间戳。
// - 如果目录无法读取或解析文件名时出现问题，则返回错误。
func (__obf_b802042ac7a13859 *LogWriter) __obf_9f00aef5ecceed2e() ([]__obf_f86003c7ac3e0dc0, error) {
	__obf_9fcb047ef26f5dd8, __obf_9c607119691e49da := os.ReadDir(__obf_b802042ac7a13859.__obf_da94489893334ea7)
	if __obf_9c607119691e49da != nil {
		return nil, __obf_9c607119691e49da
	}
	var (
		__obf_6d9ef399ff92733e                         time.Time
		__obf_326214dbcdb702b7                         []__obf_f86003c7ac3e0dc0
		__obf_841443da5c31487a, __obf_a39c26a8085c5a08 = util.SplitFilename(__obf_b802042ac7a13859.__obf_6f64510fd7abc1a8.LogPath)
	)
	__obf_841443da5c31487a += "-"
	for _, __obf_df4a5542fb2bd8f8 := range __obf_9fcb047ef26f5dd8 {
		if __obf_df4a5542fb2bd8f8.IsDir() {
			continue
		}
		if __obf_6d9ef399ff92733e, __obf_9c607119691e49da = __obf_e6e4dd1e6a593138(__obf_99ac71c2db085776, __obf_df4a5542fb2bd8f8.Name(), __obf_841443da5c31487a, __obf_a39c26a8085c5a08); __obf_9c607119691e49da == nil {
			__obf_326214dbcdb702b7 = append(__obf_326214dbcdb702b7, __obf_f86003c7ac3e0dc0{__obf_df4a5542fb2bd8f8, __obf_6d9ef399ff92733e})
			continue
		}
		if __obf_6d9ef399ff92733e, __obf_9c607119691e49da = __obf_e6e4dd1e6a593138(__obf_99ac71c2db085776, __obf_df4a5542fb2bd8f8.Name(), __obf_841443da5c31487a, __obf_a39c26a8085c5a08+__obf_ba8e597c54e63c15); __obf_9c607119691e49da == nil {
			__obf_326214dbcdb702b7 = append(__obf_326214dbcdb702b7, __obf_f86003c7ac3e0dc0{__obf_df4a5542fb2bd8f8, __obf_6d9ef399ff92733e})
			continue
		}
	}
	sort.Sort(__obf_f109286c31d62991(__obf_326214dbcdb702b7))
	return __obf_326214dbcdb702b7, nil
}

type __obf_f86003c7ac3e0dc0 struct {
	os.DirEntry
	__obf_c3d8b258b75a6c0e time.Time
}

type __obf_f109286c31d62991 []__obf_f86003c7ac3e0dc0

func (__obf_646e4c0861064600 __obf_f109286c31d62991) Less(__obf_f4634486c492c662, __obf_65067a7d8168a33b int) bool {
	return __obf_646e4c0861064600[__obf_f4634486c492c662].__obf_c3d8b258b75a6c0e.After(__obf_646e4c0861064600[__obf_65067a7d8168a33b].__obf_c3d8b258b75a6c0e)
}

func (__obf_646e4c0861064600 __obf_f109286c31d62991) Swap(__obf_f4634486c492c662, __obf_65067a7d8168a33b int) {
	__obf_646e4c0861064600[__obf_f4634486c492c662], __obf_646e4c0861064600[__obf_65067a7d8168a33b] = __obf_646e4c0861064600[__obf_65067a7d8168a33b], __obf_646e4c0861064600[__obf_f4634486c492c662]
}

func (__obf_646e4c0861064600 __obf_f109286c31d62991) Len() int {
	return len(__obf_646e4c0861064600)
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
func __obf_e6e4dd1e6a593138(__obf_99ac71c2db085776, __obf_02f956c360c37f5a, __obf_841443da5c31487a, __obf_a39c26a8085c5a08 string) (time.Time, error) {
	if !strings.HasPrefix(__obf_02f956c360c37f5a, __obf_841443da5c31487a) {
		return time.Time{}, fmt.Errorf("mismatched prefix(%s) filename(%s)", __obf_841443da5c31487a, __obf_02f956c360c37f5a)
	}
	if !strings.HasSuffix(__obf_02f956c360c37f5a, __obf_a39c26a8085c5a08) {
		return time.Time{}, fmt.Errorf("mismatched extension(%s) filename(%s)", __obf_a39c26a8085c5a08, __obf_02f956c360c37f5a)
	}
	__obf_e8537d73eff3dbfc := __obf_02f956c360c37f5a[len(__obf_841443da5c31487a) : len(__obf_02f956c360c37f5a)-len(__obf_a39c26a8085c5a08)]
	return time.Parse(__obf_99ac71c2db085776, __obf_e8537d73eff3dbfc)
}
