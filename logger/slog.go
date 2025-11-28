package __obf_290168e300ae1c1e

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_e9d26c980ef13bff struct {
	__obf_50c9e26f247bbae9 slog.Handler
	__obf_8a0b9e506e54a3bc int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_4b45982338a4946b *__obf_e9d26c980ef13bff) Handle(__obf_f05990338313c974 context.Context, __obf_1b50fa1747e275a5 slog.Record) error {
	// 获取调用者信息
	_, __obf_0da9f7ddb1f44779, __obf_81616b8a6e55d698, __obf_5ef2f0c561222ebe := runtime.Caller(__obf_4b45982338a4946b.__obf_8a0b9e506e54a3bc)
	if __obf_5ef2f0c561222ebe {
		// 只保留文件名，不需要完整路径
		__obf_0da9f7ddb1f44779 = filepath.Base(__obf_0da9f7ddb1f44779)
		// 添加源代码位置信息
		__obf_1b50fa1747e275a5.AddAttrs(slog.String("file", __obf_0da9f7ddb1f44779), slog.Int("line", __obf_81616b8a6e55d698))
	}

	// 将记录传递给底层处理器
	return __obf_4b45982338a4946b.__obf_50c9e26f247bbae9.Handle(__obf_f05990338313c974, __obf_1b50fa1747e275a5)
}

// 实现 WithAttrs 方法
func (__obf_4b45982338a4946b *__obf_e9d26c980ef13bff) WithAttrs(__obf_ff6d5198dedec1d2 []slog.Attr) slog.Handler {
	return &__obf_e9d26c980ef13bff{
		__obf_50c9e26f247bbae9: __obf_4b45982338a4946b.__obf_50c9e26f247bbae9.WithAttrs(__obf_ff6d5198dedec1d2),
		__obf_8a0b9e506e54a3bc: __obf_4b45982338a4946b.__obf_8a0b9e506e54a3bc,
	}
}

// 实现 WithGroup 方法
func (__obf_4b45982338a4946b *__obf_e9d26c980ef13bff) WithGroup(__obf_5d263750ece655d5 string) slog.Handler {
	return &__obf_e9d26c980ef13bff{
		__obf_50c9e26f247bbae9: __obf_4b45982338a4946b.__obf_50c9e26f247bbae9.WithGroup(__obf_5d263750ece655d5),
		__obf_8a0b9e506e54a3bc: __obf_4b45982338a4946b.__obf_8a0b9e506e54a3bc,
	}
}

// 实现 Enabled 方法
func (__obf_4b45982338a4946b *__obf_e9d26c980ef13bff) Enabled(__obf_f05990338313c974 context.Context, __obf_20fb1cc05fd7f304 slog.Level) bool {
	return __obf_4b45982338a4946b.__obf_50c9e26f247bbae9.Enabled(__obf_f05990338313c974, __obf_20fb1cc05fd7f304)
}

// 创建带有行追踪功能的 Logger
func NewSlogger(__obf_20fb1cc05fd7f304 slog.Leveler, __obf_2c11932736b40451 io.Writer) *slog.Logger {
	return slog.New(&__obf_e9d26c980ef13bff{
		__obf_50c9e26f247bbae9: slog.NewJSONHandler(
			__obf_2c11932736b40451,
			&slog.HandlerOptions{
				Level: __obf_20fb1cc05fd7f304,
				// 添加时间格式选项
				ReplaceAttr: func(__obf_a85fa8b1c322573f []string, __obf_4becbe1a7600dcdb slog.Attr) slog.Attr {
					if __obf_4becbe1a7600dcdb.Key == slog.TimeKey {
						__obf_4becbe1a7600dcdb.Value = slog.StringValue(__obf_4becbe1a7600dcdb.Value.Time().Format("2006-01-02 15:04:05.000"))
					}
					return __obf_4becbe1a7600dcdb
				},
			},
		),
		__obf_8a0b9e506e54a3bc: 3,
	})
}
