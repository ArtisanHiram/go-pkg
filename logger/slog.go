package __obf_b8a611fece80b60a

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_c91ead9c2858a9ff struct {
	__obf_bb7d41905482e26d slog.Handler
	__obf_b7dc19f3dc37604f int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_91cf5caa62deaccd *__obf_c91ead9c2858a9ff) Handle(__obf_20486abd98484928 context.Context, __obf_34d1428a80b78146 slog.Record) error {
	// 获取调用者信息
	_, __obf_64056d01978f85ff, __obf_d035776e0bcaef93, __obf_af9001773ef0c4ab := runtime.Caller(__obf_91cf5caa62deaccd.__obf_b7dc19f3dc37604f)
	if __obf_af9001773ef0c4ab {
		// 只保留文件名，不需要完整路径
		__obf_64056d01978f85ff = filepath.Base(__obf_64056d01978f85ff)
		// 添加源代码位置信息
		__obf_34d1428a80b78146.AddAttrs(slog.String("file", __obf_64056d01978f85ff), slog.Int("line", __obf_d035776e0bcaef93))
	}

	// 将记录传递给底层处理器
	return __obf_91cf5caa62deaccd.__obf_bb7d41905482e26d.Handle(__obf_20486abd98484928, __obf_34d1428a80b78146)
}

// 实现 WithAttrs 方法
func (__obf_91cf5caa62deaccd *__obf_c91ead9c2858a9ff) WithAttrs(__obf_6263a6a38ca64833 []slog.Attr) slog.Handler {
	return &__obf_c91ead9c2858a9ff{
		__obf_bb7d41905482e26d: __obf_91cf5caa62deaccd.__obf_bb7d41905482e26d.WithAttrs(__obf_6263a6a38ca64833),
		__obf_b7dc19f3dc37604f: __obf_91cf5caa62deaccd.__obf_b7dc19f3dc37604f,
	}
}

// 实现 WithGroup 方法
func (__obf_91cf5caa62deaccd *__obf_c91ead9c2858a9ff) WithGroup(__obf_46c5bc753fee3bc1 string) slog.Handler {
	return &__obf_c91ead9c2858a9ff{
		__obf_bb7d41905482e26d: __obf_91cf5caa62deaccd.__obf_bb7d41905482e26d.WithGroup(__obf_46c5bc753fee3bc1),
		__obf_b7dc19f3dc37604f: __obf_91cf5caa62deaccd.__obf_b7dc19f3dc37604f,
	}
}

// 实现 Enabled 方法
func (__obf_91cf5caa62deaccd *__obf_c91ead9c2858a9ff) Enabled(__obf_20486abd98484928 context.Context, __obf_951861844cea0608 slog.Level) bool {
	return __obf_91cf5caa62deaccd.__obf_bb7d41905482e26d.Enabled(__obf_20486abd98484928, __obf_951861844cea0608)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_951861844cea0608 slog.Leveler, __obf_cf0b3e22ab1aa655 io.Writer) *slog.Logger {
	return slog.New(&__obf_c91ead9c2858a9ff{
		__obf_bb7d41905482e26d: slog.NewJSONHandler(
			__obf_cf0b3e22ab1aa655,
			&slog.HandlerOptions{
				Level: __obf_951861844cea0608,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_f400ef298b515bbb []string, __obf_7a803ee5a16bb56b slog.Attr) slog.Attr {
					if __obf_7a803ee5a16bb56b.Key == slog.TimeKey {
						__obf_7a803ee5a16bb56b.Value = slog.StringValue(__obf_7a803ee5a16bb56b.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_7a803ee5a16bb56b
				},
			},
		),
		__obf_b7dc19f3dc37604f: 3,
	})
}
