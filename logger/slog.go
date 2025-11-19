package __obf_dc6fa34d54e9b538

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_d28cacb8529393b9 struct {
	__obf_2d9b03c6a5b82017 slog.Handler
	__obf_535d0d175d663158 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_1259f512f4e88d41 *__obf_d28cacb8529393b9) Handle(__obf_7e1132d5839c8bf0 context.Context, __obf_6522fbb6baff1bd4 slog.Record) error {
	// 获取调用者信息
	_, __obf_577762d05139a620, __obf_6c5236a02107db15, __obf_d988c28fcc121eee := runtime.Caller(__obf_1259f512f4e88d41.__obf_535d0d175d663158)
	if __obf_d988c28fcc121eee {
		// 只保留文件名，不需要完整路径
		__obf_577762d05139a620 = filepath.Base(__obf_577762d05139a620)
		// 添加源代码位置信息
		__obf_6522fbb6baff1bd4.AddAttrs(slog.String("file", __obf_577762d05139a620), slog.Int("line", __obf_6c5236a02107db15))
	}

	// 将记录传递给底层处理器
	return __obf_1259f512f4e88d41.__obf_2d9b03c6a5b82017.Handle(__obf_7e1132d5839c8bf0, __obf_6522fbb6baff1bd4)
}

// 实现 WithAttrs 方法
func (__obf_1259f512f4e88d41 *__obf_d28cacb8529393b9) WithAttrs(__obf_b8867ecf7aac6d8c []slog.Attr) slog.Handler {
	return &__obf_d28cacb8529393b9{
		__obf_2d9b03c6a5b82017: __obf_1259f512f4e88d41.__obf_2d9b03c6a5b82017.WithAttrs(__obf_b8867ecf7aac6d8c),
		__obf_535d0d175d663158: __obf_1259f512f4e88d41.__obf_535d0d175d663158,
	}
}

// 实现 WithGroup 方法
func (__obf_1259f512f4e88d41 *__obf_d28cacb8529393b9) WithGroup(__obf_8b6d36780842caec string) slog.Handler {
	return &__obf_d28cacb8529393b9{
		__obf_2d9b03c6a5b82017: __obf_1259f512f4e88d41.__obf_2d9b03c6a5b82017.WithGroup(__obf_8b6d36780842caec),
		__obf_535d0d175d663158: __obf_1259f512f4e88d41.__obf_535d0d175d663158,
	}
}

// 实现 Enabled 方法
func (__obf_1259f512f4e88d41 *__obf_d28cacb8529393b9) Enabled(__obf_7e1132d5839c8bf0 context.Context, __obf_0486d10368d8b008 slog.Level) bool {
	return __obf_1259f512f4e88d41.__obf_2d9b03c6a5b82017.Enabled(__obf_7e1132d5839c8bf0, __obf_0486d10368d8b008)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_0486d10368d8b008 slog.Leveler, __obf_f210d85c325bd536 io.Writer) *slog.Logger {
	return slog.New(&__obf_d28cacb8529393b9{
		__obf_2d9b03c6a5b82017: slog.NewJSONHandler(
			__obf_f210d85c325bd536,
			&slog.HandlerOptions{
				Level: __obf_0486d10368d8b008,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_a11cee9f3903b5ef []string, __obf_135c9b82f4e7624f slog.Attr) slog.Attr {
					if __obf_135c9b82f4e7624f.Key == slog.TimeKey {
						__obf_135c9b82f4e7624f.Value = slog.StringValue(__obf_135c9b82f4e7624f.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_135c9b82f4e7624f
				},
			},
		),
		__obf_535d0d175d663158: 3,
	})
}
