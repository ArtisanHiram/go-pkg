package __obf_ad1b8b65c829ec46

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_45cd2e574f62b00e struct {
	__obf_0c6b4783ce3821ff slog.Handler
	__obf_390c6921b97a1c75 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_355cf735d9869b78 *__obf_45cd2e574f62b00e) Handle(__obf_b6bc8945a2dfe1d4 context.Context, __obf_c2a0734b6ac4f901 slog.Record) error {
	// 获取调用者信息
	_, __obf_f669dbd05fc7537a, __obf_635f726b27460257, __obf_8b7e75ba625d8c0e := runtime.Caller(__obf_355cf735d9869b78.

		// 只保留文件名，不需要完整路径
		__obf_390c6921b97a1c75)
	if __obf_8b7e75ba625d8c0e {
		__obf_f669dbd05fc7537a = filepath.Base(__obf_f669dbd05fc7537a)
		__obf_c2a0734b6ac4f901.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_f669dbd05fc7537a), slog.Int("line", __obf_635f726b27460257))
	}

	// 将记录传递给底层处理器
	return __obf_355cf735d9869b78.__obf_0c6b4783ce3821ff.Handle(__obf_b6bc8945a2dfe1d4,

		// 实现 WithAttrs 方法
		__obf_c2a0734b6ac4f901)
}

func (__obf_355cf735d9869b78 *__obf_45cd2e574f62b00e) WithAttrs(__obf_5b78bfc45884c50a []slog.Attr) slog.Handler {
	return &__obf_45cd2e574f62b00e{__obf_0c6b4783ce3821ff: __obf_355cf735d9869b78.__obf_0c6b4783ce3821ff.WithAttrs(__obf_5b78bfc45884c50a), __obf_390c6921b97a1c75: __obf_355cf735d9869b78.

		// 实现 WithGroup 方法
		__obf_390c6921b97a1c75,
	}
}

func (__obf_355cf735d9869b78 *__obf_45cd2e574f62b00e) WithGroup(__obf_c35be3640bde6e62 string) slog.Handler {
	return &__obf_45cd2e574f62b00e{__obf_0c6b4783ce3821ff: __obf_355cf735d9869b78.__obf_0c6b4783ce3821ff.WithGroup(__obf_c35be3640bde6e62), __obf_390c6921b97a1c75: __obf_355cf735d9869b78.

		// 实现 Enabled 方法
		__obf_390c6921b97a1c75,
	}
}

func (__obf_355cf735d9869b78 *__obf_45cd2e574f62b00e) Enabled(__obf_b6bc8945a2dfe1d4 context.Context, __obf_5b43585b455c6b54 slog.Level) bool {
	return __obf_355cf735d9869b78.__obf_0c6b4783ce3821ff.Enabled(__obf_b6bc8945a2dfe1d4,

		// 创建带有行追踪功能的 Logger
		__obf_5b43585b455c6b54)
}

func NewSlogger(__obf_5b43585b455c6b54 slog.Leveler, __obf_025aab1df039d9be io.Writer) *slog.Logger {
	return slog.New(&__obf_45cd2e574f62b00e{__obf_0c6b4783ce3821ff: slog.NewJSONHandler(__obf_025aab1df039d9be, &slog.HandlerOptions{
		Level: __obf_5b43585b455c6b54,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_373a4a9d5baa1c8a []string, __obf_6f3cbaca302447ba slog.Attr) slog.Attr {
			if __obf_6f3cbaca302447ba.Key == slog.TimeKey {
				__obf_6f3cbaca302447ba.
					Value = slog.StringValue(__obf_6f3cbaca302447ba.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_6f3cbaca302447ba
		},
	},
	), __obf_390c6921b97a1c75: 3,
	})
}
