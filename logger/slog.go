package __obf_0e25611156600e69

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_a68e7e9c8755112d struct {
	__obf_006aba450ffbea62 slog.Handler
	__obf_43f40065fe38ec15 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_44add25f405dca15 *__obf_a68e7e9c8755112d) Handle(__obf_81e9d6c6a93edad5 context.Context, __obf_0d2d0bcaee09d3f2 slog.Record) error {
	// 获取调用者信息
	_, __obf_3ff7d69d4eeecb80, __obf_e9b252947651cfc5, __obf_9961801cee5402be := runtime.Caller(__obf_44add25f405dca15.__obf_43f40065fe38ec15)
	if __obf_9961801cee5402be {
		// 只保留文件名，不需要完整路径
		__obf_3ff7d69d4eeecb80 = filepath.Base(__obf_3ff7d69d4eeecb80)
		// 添加源代码位置信息
		__obf_0d2d0bcaee09d3f2.AddAttrs(slog.String("file", __obf_3ff7d69d4eeecb80), slog.Int("line", __obf_e9b252947651cfc5))
	}

	// 将记录传递给底层处理器
	return __obf_44add25f405dca15.__obf_006aba450ffbea62.Handle(__obf_81e9d6c6a93edad5, __obf_0d2d0bcaee09d3f2)
}

// 实现 WithAttrs 方法
func (__obf_44add25f405dca15 *__obf_a68e7e9c8755112d) WithAttrs(__obf_4e62cb6198b8aac9 []slog.Attr) slog.Handler {
	return &__obf_a68e7e9c8755112d{
		__obf_006aba450ffbea62: __obf_44add25f405dca15.__obf_006aba450ffbea62.WithAttrs(__obf_4e62cb6198b8aac9),
		__obf_43f40065fe38ec15: __obf_44add25f405dca15.__obf_43f40065fe38ec15,
	}
}

// 实现 WithGroup 方法
func (__obf_44add25f405dca15 *__obf_a68e7e9c8755112d) WithGroup(__obf_69f9cb1f34a954bc string) slog.Handler {
	return &__obf_a68e7e9c8755112d{
		__obf_006aba450ffbea62: __obf_44add25f405dca15.__obf_006aba450ffbea62.WithGroup(__obf_69f9cb1f34a954bc),
		__obf_43f40065fe38ec15: __obf_44add25f405dca15.__obf_43f40065fe38ec15,
	}
}

// 实现 Enabled 方法
func (__obf_44add25f405dca15 *__obf_a68e7e9c8755112d) Enabled(__obf_81e9d6c6a93edad5 context.Context, __obf_a48f5c2c619f6e71 slog.Level) bool {
	return __obf_44add25f405dca15.__obf_006aba450ffbea62.Enabled(__obf_81e9d6c6a93edad5, __obf_a48f5c2c619f6e71)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_a48f5c2c619f6e71 slog.Leveler, __obf_b95351e6905708e0 io.Writer) *slog.Logger {
	return slog.New(&__obf_a68e7e9c8755112d{
		__obf_006aba450ffbea62: slog.NewJSONHandler(
			__obf_b95351e6905708e0,
			&slog.HandlerOptions{
				Level: __obf_a48f5c2c619f6e71,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_58d017d222670e8d []string, __obf_5dddbc19c11015f5 slog.Attr) slog.Attr {
					if __obf_5dddbc19c11015f5.Key == slog.TimeKey {
						__obf_5dddbc19c11015f5.Value = slog.StringValue(__obf_5dddbc19c11015f5.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_5dddbc19c11015f5
				},
			},
		),
		__obf_43f40065fe38ec15: 3,
	})
}
