package __obf_10b299a6084cd5a7

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_6e1c1d21fa758b20 struct {
	__obf_18b9926912e0b30f slog.Handler
	__obf_b7df722733c8ead7 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_bb4242b03591434e *__obf_6e1c1d21fa758b20) Handle(__obf_ce1fd7f3c3ead082 context.Context, __obf_1ccaeea252bff4c8 slog.Record) error {
	// 获取调用者信息
	_, __obf_a142fe1a109f3f49, __obf_ea52975dcc892982, __obf_dddab163f3920dd1 := runtime.Caller(__obf_bb4242b03591434e.

		// 只保留文件名，不需要完整路径
		__obf_b7df722733c8ead7)
	if __obf_dddab163f3920dd1 {
		__obf_a142fe1a109f3f49 = filepath.Base(__obf_a142fe1a109f3f49)
		__obf_1ccaeea252bff4c8.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_a142fe1a109f3f49), slog.Int("line", __obf_ea52975dcc892982))
	}

	// 将记录传递给底层处理器
	return __obf_bb4242b03591434e.__obf_18b9926912e0b30f.Handle(__obf_ce1fd7f3c3ead082,

		// 实现 WithAttrs 方法
		__obf_1ccaeea252bff4c8)
}

func (__obf_bb4242b03591434e *__obf_6e1c1d21fa758b20) WithAttrs(__obf_0d5feb7c145c6fcb []slog.Attr) slog.Handler {
	return &__obf_6e1c1d21fa758b20{__obf_18b9926912e0b30f: __obf_bb4242b03591434e.__obf_18b9926912e0b30f.WithAttrs(__obf_0d5feb7c145c6fcb), __obf_b7df722733c8ead7: __obf_bb4242b03591434e.

		// 实现 WithGroup 方法
		__obf_b7df722733c8ead7,
	}
}

func (__obf_bb4242b03591434e *__obf_6e1c1d21fa758b20) WithGroup(__obf_67032a27bbbf2a0e string) slog.Handler {
	return &__obf_6e1c1d21fa758b20{__obf_18b9926912e0b30f: __obf_bb4242b03591434e.__obf_18b9926912e0b30f.WithGroup(__obf_67032a27bbbf2a0e), __obf_b7df722733c8ead7: __obf_bb4242b03591434e.

		// 实现 Enabled 方法
		__obf_b7df722733c8ead7,
	}
}

func (__obf_bb4242b03591434e *__obf_6e1c1d21fa758b20) Enabled(__obf_ce1fd7f3c3ead082 context.Context, __obf_a4fc493b1adbfe50 slog.Level) bool {
	return __obf_bb4242b03591434e.__obf_18b9926912e0b30f.Enabled(__obf_ce1fd7f3c3ead082,

		// 创建带有行追踪功能的 Logger
		__obf_a4fc493b1adbfe50)
}

func NewSlogger(__obf_a4fc493b1adbfe50 slog.Leveler, __obf_314efd80329f0e2d io.Writer) *slog.Logger {
	return slog.New(&__obf_6e1c1d21fa758b20{__obf_18b9926912e0b30f: slog.NewJSONHandler(__obf_314efd80329f0e2d, &slog.HandlerOptions{
		Level: __obf_a4fc493b1adbfe50,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_3053adcdea8f13c1 []string, __obf_e534ea04415f78f7 slog.Attr) slog.Attr {
			if __obf_e534ea04415f78f7.Key == slog.TimeKey {
				__obf_e534ea04415f78f7.
					Value = slog.StringValue(__obf_e534ea04415f78f7.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_e534ea04415f78f7
		},
	},
	), __obf_b7df722733c8ead7: 3,
	})
}
