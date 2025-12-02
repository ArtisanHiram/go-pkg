package __obf_5efebb632ee092eb

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_2b2b7fd5cbd26093 struct {
	__obf_4d99658ef1ac6559 slog.Handler
	__obf_df5bdae140f201bf int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_6bdf86600ff00b0f *__obf_2b2b7fd5cbd26093) Handle(__obf_c60a0d69a66caf25 context.Context, __obf_c0411ac768140b17 slog.Record) error {
	// 获取调用者信息
	_, __obf_74a95368ac22e29f, __obf_8eece89742f2098d, __obf_2b91da7d2b18670c := runtime.Caller(__obf_6bdf86600ff00b0f.

		// 只保留文件名，不需要完整路径
		__obf_df5bdae140f201bf)
	if __obf_2b91da7d2b18670c {
		__obf_74a95368ac22e29f = filepath.Base(__obf_74a95368ac22e29f)
		__obf_c0411ac768140b17.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_74a95368ac22e29f), slog.Int("line", __obf_8eece89742f2098d))
	}

	// 将记录传递给底层处理器
	return __obf_6bdf86600ff00b0f.__obf_4d99658ef1ac6559.Handle(__obf_c60a0d69a66caf25,

		// 实现 WithAttrs 方法
		__obf_c0411ac768140b17)
}

func (__obf_6bdf86600ff00b0f *__obf_2b2b7fd5cbd26093) WithAttrs(__obf_bb7a964b6d2333f7 []slog.Attr) slog.Handler {
	return &__obf_2b2b7fd5cbd26093{__obf_4d99658ef1ac6559: __obf_6bdf86600ff00b0f.__obf_4d99658ef1ac6559.WithAttrs(__obf_bb7a964b6d2333f7), __obf_df5bdae140f201bf: __obf_6bdf86600ff00b0f.

		// 实现 WithGroup 方法
		__obf_df5bdae140f201bf,
	}
}

func (__obf_6bdf86600ff00b0f *__obf_2b2b7fd5cbd26093) WithGroup(__obf_08567634b89f4c18 string) slog.Handler {
	return &__obf_2b2b7fd5cbd26093{__obf_4d99658ef1ac6559: __obf_6bdf86600ff00b0f.__obf_4d99658ef1ac6559.WithGroup(__obf_08567634b89f4c18), __obf_df5bdae140f201bf: __obf_6bdf86600ff00b0f.

		// 实现 Enabled 方法
		__obf_df5bdae140f201bf,
	}
}

func (__obf_6bdf86600ff00b0f *__obf_2b2b7fd5cbd26093) Enabled(__obf_c60a0d69a66caf25 context.Context, __obf_e16dc43e54a8fc55 slog.Level) bool {
	return __obf_6bdf86600ff00b0f.__obf_4d99658ef1ac6559.Enabled(__obf_c60a0d69a66caf25,

		// 创建带有行追踪功能的 Logger
		__obf_e16dc43e54a8fc55)
}

func NewSlogger(__obf_e16dc43e54a8fc55 slog.Leveler, __obf_145b3a4dda2e0b6c io.Writer) *slog.Logger {
	return slog.New(&__obf_2b2b7fd5cbd26093{__obf_4d99658ef1ac6559: slog.NewJSONHandler(__obf_145b3a4dda2e0b6c, &slog.HandlerOptions{
		Level: __obf_e16dc43e54a8fc55,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_fdc53ed64cb69e07 []string, __obf_6cfb56303ab59e68 slog.Attr) slog.Attr {
			if __obf_6cfb56303ab59e68.Key == slog.TimeKey {
				__obf_6cfb56303ab59e68.
					Value = slog.StringValue(__obf_6cfb56303ab59e68.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_6cfb56303ab59e68
		},
	},
	), __obf_df5bdae140f201bf: 3,
	})
}
