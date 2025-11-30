package __obf_0f05083f0caeeb4d

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_98045df2cc8cb1f6 struct {
	__obf_92643c6dd8996174 slog.Handler
	__obf_9c733d281b7a126a int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_fee344f8daea7c2b *__obf_98045df2cc8cb1f6) Handle(__obf_26bde9a094df0274 context.Context, __obf_495b6b3373985f7b slog.Record) error {
	// 获取调用者信息
	_, __obf_70f2d96a7c56d8e7, __obf_917b87fcbac61290, __obf_c97474f5a6894426 := runtime.Caller(__obf_fee344f8daea7c2b.

		// 只保留文件名，不需要完整路径
		__obf_9c733d281b7a126a)
	if __obf_c97474f5a6894426 {
		__obf_70f2d96a7c56d8e7 = filepath.Base(__obf_70f2d96a7c56d8e7)
		__obf_495b6b3373985f7b.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_70f2d96a7c56d8e7), slog.Int("line", __obf_917b87fcbac61290))
	}

	// 将记录传递给底层处理器
	return __obf_fee344f8daea7c2b.__obf_92643c6dd8996174.Handle(__obf_26bde9a094df0274,

		// 实现 WithAttrs 方法
		__obf_495b6b3373985f7b)
}

func (__obf_fee344f8daea7c2b *__obf_98045df2cc8cb1f6) WithAttrs(__obf_64539d49a98aaf9f []slog.Attr) slog.Handler {
	return &__obf_98045df2cc8cb1f6{__obf_92643c6dd8996174: __obf_fee344f8daea7c2b.__obf_92643c6dd8996174.WithAttrs(__obf_64539d49a98aaf9f), __obf_9c733d281b7a126a: __obf_fee344f8daea7c2b.

		// 实现 WithGroup 方法
		__obf_9c733d281b7a126a,
	}
}

func (__obf_fee344f8daea7c2b *__obf_98045df2cc8cb1f6) WithGroup(__obf_9d7072afc9c6c54b string) slog.Handler {
	return &__obf_98045df2cc8cb1f6{__obf_92643c6dd8996174: __obf_fee344f8daea7c2b.__obf_92643c6dd8996174.WithGroup(__obf_9d7072afc9c6c54b), __obf_9c733d281b7a126a: __obf_fee344f8daea7c2b.

		// 实现 Enabled 方法
		__obf_9c733d281b7a126a,
	}
}

func (__obf_fee344f8daea7c2b *__obf_98045df2cc8cb1f6) Enabled(__obf_26bde9a094df0274 context.Context, __obf_8e7bac9b0e753fe7 slog.Level) bool {
	return __obf_fee344f8daea7c2b.__obf_92643c6dd8996174.Enabled(__obf_26bde9a094df0274,

		// 创建带有行追踪功能的 Logger
		__obf_8e7bac9b0e753fe7)
}

func NewSlogger(__obf_8e7bac9b0e753fe7 slog.Leveler, __obf_7a49a6d7efbc3977 io.Writer) *slog.Logger {
	return slog.New(&__obf_98045df2cc8cb1f6{__obf_92643c6dd8996174: slog.NewJSONHandler(__obf_7a49a6d7efbc3977, &slog.HandlerOptions{
		Level: __obf_8e7bac9b0e753fe7,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_7f80bbb4930f1e53 []string, __obf_029e861f13207ac8 slog.Attr) slog.Attr {
			if __obf_029e861f13207ac8.Key == slog.TimeKey {
				__obf_029e861f13207ac8.
					Value = slog.StringValue(__obf_029e861f13207ac8.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_029e861f13207ac8
		},
	},
	), __obf_9c733d281b7a126a: 3,
	})
}
