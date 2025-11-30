package __obf_55b2846f9c90acae

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

// 自定义处理器，用于添加行号信息
type __obf_0c0932fdcc6f9478 struct {
	__obf_d958cc777e1788ab slog.Handler
	__obf_89698fa7fe9f05e0 int // 调用栈跳过的帧数
}

// 实现 Handle 方法
func (__obf_2d6795658eb91c43 *__obf_0c0932fdcc6f9478) Handle(__obf_5626fc6a2c273cc2 context.Context, __obf_af74ffc59ce17d52 slog.Record) error {
	// 获取调用者信息
	_, __obf_52c640a2618c9eec, __obf_d8b9a8897ebe169d, __obf_b2778f427704e408 := runtime.Caller(__obf_2d6795658eb91c43.

		// 只保留文件名，不需要完整路径
		__obf_89698fa7fe9f05e0)
	if __obf_b2778f427704e408 {
		__obf_52c640a2618c9eec = filepath.Base(__obf_52c640a2618c9eec)
		__obf_af74ffc59ce17d52.
			// 添加源代码位置信息
			AddAttrs(slog.String("file", __obf_52c640a2618c9eec), slog.Int("line", __obf_d8b9a8897ebe169d))
	}

	// 将记录传递给底层处理器
	return __obf_2d6795658eb91c43.__obf_d958cc777e1788ab.Handle(__obf_5626fc6a2c273cc2,

		// 实现 WithAttrs 方法
		__obf_af74ffc59ce17d52)
}

func (__obf_2d6795658eb91c43 *__obf_0c0932fdcc6f9478) WithAttrs(__obf_cb4d121d82c0fd02 []slog.Attr) slog.Handler {
	return &__obf_0c0932fdcc6f9478{__obf_d958cc777e1788ab: __obf_2d6795658eb91c43.__obf_d958cc777e1788ab.WithAttrs(__obf_cb4d121d82c0fd02), __obf_89698fa7fe9f05e0: __obf_2d6795658eb91c43.

		// 实现 WithGroup 方法
		__obf_89698fa7fe9f05e0,
	}
}

func (__obf_2d6795658eb91c43 *__obf_0c0932fdcc6f9478) WithGroup(__obf_18363f61ff5eccc1 string) slog.Handler {
	return &__obf_0c0932fdcc6f9478{__obf_d958cc777e1788ab: __obf_2d6795658eb91c43.__obf_d958cc777e1788ab.WithGroup(__obf_18363f61ff5eccc1), __obf_89698fa7fe9f05e0: __obf_2d6795658eb91c43.

		// 实现 Enabled 方法
		__obf_89698fa7fe9f05e0,
	}
}

func (__obf_2d6795658eb91c43 *__obf_0c0932fdcc6f9478) Enabled(__obf_5626fc6a2c273cc2 context.Context, __obf_f894a148e1e70347 slog.Level) bool {
	return __obf_2d6795658eb91c43.__obf_d958cc777e1788ab.Enabled(__obf_5626fc6a2c273cc2,

		// 创建带有行追踪功能的 Logger
		__obf_f894a148e1e70347)
}

func NewSlogger(__obf_f894a148e1e70347 slog.Leveler, __obf_9d0336c7c99cfa2a io.Writer) *slog.Logger {
	return slog.New(&__obf_0c0932fdcc6f9478{__obf_d958cc777e1788ab: slog.NewJSONHandler(__obf_9d0336c7c99cfa2a, &slog.HandlerOptions{
		Level: __obf_f894a148e1e70347,
		// 添加时间格式选项
		ReplaceAttr: func(__obf_3aacef65e523d907 []string, __obf_729db1156eca47b9 slog.Attr) slog.Attr {
			if __obf_729db1156eca47b9.Key == slog.TimeKey {
				__obf_729db1156eca47b9.
					Value = slog.StringValue(__obf_729db1156eca47b9.Value.Time().Format("2006-01-02 15:04:05.000"))
			}
			return __obf_729db1156eca47b9
		},
	},
	), __obf_89698fa7fe9f05e0: 3,
	})
}
