package __obf_1844c112ef6a527c

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_cb1cc7f19a2e1cb9 struct {
	__obf_a990aa981c7cad98 slog.Handler
	__obf_e65a8eae0e1c95a8 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_5a3bd0d4b44cb27c *__obf_cb1cc7f19a2e1cb9) Handle(__obf_b96819be9e49b713 context.Context, __obf_86c0c91bd587d770 slog.Record) error {
	// 获取调用者信息
	_, __obf_982af539c63d8b14, __obf_33091478b617d382, __obf_c57632c58d1b3ca8 := runtime.Caller(__obf_5a3bd0d4b44cb27c.__obf_e65a8eae0e1c95a8)
	if __obf_c57632c58d1b3ca8 {
		// 只保留文件名，不需要完整路径
		__obf_982af539c63d8b14 = filepath.Base(__obf_982af539c63d8b14)
		// 添加源代码位置信息
		__obf_86c0c91bd587d770.AddAttrs(slog.String("file", __obf_982af539c63d8b14), slog.Int("line", __obf_33091478b617d382))
	}

	// 将记录传递给底层处理器
	return __obf_5a3bd0d4b44cb27c.__obf_a990aa981c7cad98.Handle(__obf_b96819be9e49b713, __obf_86c0c91bd587d770)
}

// 实现 WithAttrs 方法
func (__obf_5a3bd0d4b44cb27c *__obf_cb1cc7f19a2e1cb9) WithAttrs(__obf_9a3d5435a3e3cd64 []slog.Attr) slog.Handler {
	return &__obf_cb1cc7f19a2e1cb9{
		__obf_a990aa981c7cad98: __obf_5a3bd0d4b44cb27c.__obf_a990aa981c7cad98.WithAttrs(__obf_9a3d5435a3e3cd64),
		__obf_e65a8eae0e1c95a8: __obf_5a3bd0d4b44cb27c.__obf_e65a8eae0e1c95a8,
	}
}

// 实现 WithGroup 方法
func (__obf_5a3bd0d4b44cb27c *__obf_cb1cc7f19a2e1cb9) WithGroup(__obf_4c6dfac2cc499632 string) slog.Handler {
	return &__obf_cb1cc7f19a2e1cb9{
		__obf_a990aa981c7cad98: __obf_5a3bd0d4b44cb27c.__obf_a990aa981c7cad98.WithGroup(__obf_4c6dfac2cc499632),
		__obf_e65a8eae0e1c95a8: __obf_5a3bd0d4b44cb27c.__obf_e65a8eae0e1c95a8,
	}
}

// 实现 Enabled 方法
func (__obf_5a3bd0d4b44cb27c *__obf_cb1cc7f19a2e1cb9) Enabled(__obf_b96819be9e49b713 context.Context, __obf_786e59c1c90ef760 slog.Level) bool {
	return __obf_5a3bd0d4b44cb27c.__obf_a990aa981c7cad98.Enabled(__obf_b96819be9e49b713, __obf_786e59c1c90ef760)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_786e59c1c90ef760 slog.Leveler, __obf_9bb1cd5776d3686a io.Writer) *slog.Logger {
	return slog.New(&__obf_cb1cc7f19a2e1cb9{
		__obf_a990aa981c7cad98: slog.NewJSONHandler(
			__obf_9bb1cd5776d3686a,
			&slog.HandlerOptions{
				Level: __obf_786e59c1c90ef760,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_3cd5076ae6dca316 []string, __obf_f04a808cb04c0e9f slog.Attr) slog.Attr {
					if __obf_f04a808cb04c0e9f.Key == slog.TimeKey {
						__obf_f04a808cb04c0e9f.Value = slog.StringValue(__obf_f04a808cb04c0e9f.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_f04a808cb04c0e9f
				},
			},
		),
		__obf_e65a8eae0e1c95a8: 3,
	})
}
