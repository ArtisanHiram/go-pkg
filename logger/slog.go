package __obf_6664e1d521d1e765

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_137c1aaa8db1b9a6 struct {
	__obf_ff5b088cb4f38aa3 slog.Handler
	__obf_27740c52a270edeb int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_2d3313359418c213 *__obf_137c1aaa8db1b9a6) Handle(__obf_2d680bdb77f841e8 context.Context, __obf_2ce2f7faad11bbd3 slog.Record) error {
	// 获取调用者信息
	_, __obf_5459f006a6731245, __obf_c2a90964df3136e3, __obf_3c32a351e83eac58 := runtime.Caller(__obf_2d3313359418c213.__obf_27740c52a270edeb)
	if __obf_3c32a351e83eac58 {
		// 只保留文件名，不需要完整路径
		__obf_5459f006a6731245 = filepath.Base(__obf_5459f006a6731245)
		// 添加源代码位置信息
		__obf_2ce2f7faad11bbd3.AddAttrs(slog.String("file", __obf_5459f006a6731245), slog.Int("line", __obf_c2a90964df3136e3))
	}

	// 将记录传递给底层处理器
	return __obf_2d3313359418c213.__obf_ff5b088cb4f38aa3.Handle(__obf_2d680bdb77f841e8, __obf_2ce2f7faad11bbd3)
}

// 实现 WithAttrs 方法
func (__obf_2d3313359418c213 *__obf_137c1aaa8db1b9a6) WithAttrs(__obf_8721703beaa45b57 []slog.Attr) slog.Handler {
	return &__obf_137c1aaa8db1b9a6{
		__obf_ff5b088cb4f38aa3: __obf_2d3313359418c213.__obf_ff5b088cb4f38aa3.WithAttrs(__obf_8721703beaa45b57),
		__obf_27740c52a270edeb: __obf_2d3313359418c213.__obf_27740c52a270edeb,
	}
}

// 实现 WithGroup 方法
func (__obf_2d3313359418c213 *__obf_137c1aaa8db1b9a6) WithGroup(__obf_5f4fae4ad2289e4f string) slog.Handler {
	return &__obf_137c1aaa8db1b9a6{
		__obf_ff5b088cb4f38aa3: __obf_2d3313359418c213.__obf_ff5b088cb4f38aa3.WithGroup(__obf_5f4fae4ad2289e4f),
		__obf_27740c52a270edeb: __obf_2d3313359418c213.__obf_27740c52a270edeb,
	}
}

// 实现 Enabled 方法
func (__obf_2d3313359418c213 *__obf_137c1aaa8db1b9a6) Enabled(__obf_2d680bdb77f841e8 context.Context, __obf_2d025ad4da223d05 slog.Level) bool {
	return __obf_2d3313359418c213.__obf_ff5b088cb4f38aa3.Enabled(__obf_2d680bdb77f841e8, __obf_2d025ad4da223d05)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_2d025ad4da223d05 slog.Leveler, __obf_b1f5f3c1ab69ed8c io.Writer) *slog.Logger {
	return slog.New(&__obf_137c1aaa8db1b9a6{
		__obf_ff5b088cb4f38aa3: slog.NewJSONHandler(
			__obf_b1f5f3c1ab69ed8c,
			&slog.HandlerOptions{
				Level: __obf_2d025ad4da223d05,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_0a307634795fea6a []string, __obf_971707c9c02a13af slog.Attr) slog.Attr {
					if __obf_971707c9c02a13af.Key == slog.TimeKey {
						__obf_971707c9c02a13af.Value = slog.StringValue(__obf_971707c9c02a13af.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_971707c9c02a13af
				},
			},
		),
		__obf_27740c52a270edeb: 3,
	})
}
