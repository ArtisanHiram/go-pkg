package __obf_ddb9362593c4afda

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_ccca19e67c464db3 struct {
	__obf_be0b0016436d72de slog.Handler
	__obf_efe2b5df04aef2a1 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_93b46eef0d66c35f *__obf_ccca19e67c464db3) Handle(__obf_5db46511d434be85 context.Context, __obf_6987ccdd35368d85 slog.Record) error {
	// 获取调用者信息
	_, __obf_546b139714a6d61f, __obf_899c41640655b181, __obf_3fcb7a46ac2668c9 := runtime.Caller(__obf_93b46eef0d66c35f.

		// 只保留文件名，不需要完整路径
		__obf_efe2b5df04aef2a1)
	if __obf_3fcb7a46ac2668c9 {
		__obf_546b139714a6d61f = filepath.Base(__obf_546b139714a6d61f)
		__obf_6987ccdd35368d85.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_546b139714a6d61f), slog.Int("line", __obf_899c41640655b181))
	}

	// 将记录传递给底层处理器
	return __obf_93b46eef0d66c35f.__obf_be0b0016436d72de.Handle(__obf_5db46511d434be85,

		// 实现 WithAttrs 方法
		__obf_6987ccdd35368d85)
}

func (__obf_93b46eef0d66c35f *__obf_ccca19e67c464db3) WithAttrs(__obf_e648cc2ea6418aa8 []slog.Attr) slog.Handler {
	return &__obf_ccca19e67c464db3{__obf_be0b0016436d72de: __obf_93b46eef0d66c35f.__obf_be0b0016436d72de.WithAttrs(__obf_e648cc2ea6418aa8), __obf_efe2b5df04aef2a1: __obf_93b46eef0d66c35f.

		// 实现 WithGroup 方法
		__obf_efe2b5df04aef2a1,
	}
}

func (__obf_93b46eef0d66c35f *__obf_ccca19e67c464db3) WithGroup(__obf_3af07cce60786893 string) slog.Handler {
	return &__obf_ccca19e67c464db3{__obf_be0b0016436d72de: __obf_93b46eef0d66c35f.__obf_be0b0016436d72de.WithGroup(__obf_3af07cce60786893), __obf_efe2b5df04aef2a1: __obf_93b46eef0d66c35f.

		// 实现 Enabled 方法
		__obf_efe2b5df04aef2a1,
	}
}

func (__obf_93b46eef0d66c35f *__obf_ccca19e67c464db3) Enabled(__obf_5db46511d434be85 context.Context, __obf_f5b23006f4018170 slog.Level) bool {
	return __obf_93b46eef0d66c35f.__obf_be0b0016436d72de.Enabled(__obf_5db46511d434be85,

		// 创建带有行追踪功能的 Logger
		__obf_f5b23006f4018170)
}

func NewSlogger(__obf_f5b23006f4018170 slog.Leveler, __obf_9b959bece22934bd io.Writer) *slog.Logger {
	return slog.New(&__obf_ccca19e67c464db3{__obf_be0b0016436d72de: slog.NewJSONHandler(__obf_9b959bece22934bd, &slog.HandlerOptions{
		Level: __obf_f5b23006f4018170,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_68e38ce534311499 []string, __obf_9ac5b3353ed05a24 slog.Attr) slog.Attr {
			if __obf_9ac5b3353ed05a24.Key == slog.TimeKey {
				__obf_9ac5b3353ed05a24.
					Value = slog.StringValue(__obf_9ac5b3353ed05a24.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_9ac5b3353ed05a24
		},
	},
	), __obf_efe2b5df04aef2a1: 3,
	})
}
