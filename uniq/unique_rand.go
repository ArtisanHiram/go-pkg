package __obf_07f0876faa0cf68e

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_15719f79f826bdc4 sync.Mutex
	__obf_0dddbc6b117ed615 map[string]struct{}
	__obf_e67e8888d48d964d * // 追踪已经生成过的
	rand.Rand
	__obf_9e80bd3b1d9a7224 string // 基础随机数
	__obf_ef955105287ffc55 int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_ddbb4cb554aa04a0 int, __obf_a453ec6cd1cbcdce []string) (*UniqueRand, error) {
	if __obf_ddbb4cb554aa04a0 > 19 || __obf_ddbb4cb554aa04a0 < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_b26d73d753cfcc70 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_ef955105287ffc55, __obf_9ad047387e1f4b21 := stats.IntPow(9, __obf_ddbb4cb554aa04a0), len(__obf_a453ec6cd1cbcdce)

	if __obf_ef955105287ffc55 <= __obf_9ad047387e1f4b21 {
		return nil, ErrIDExhaustied
	}
	__obf_0dddbc6b117ed615 := make(map[string]struct{}, __obf_9ad047387e1f4b21)
	for _, __obf_a300f1fefd6652b6 := range __obf_a453ec6cd1cbcdce {
		if len(__obf_a300f1fefd6652b6) == __obf_ddbb4cb554aa04a0 {
			__obf_0dddbc6b117ed615[__obf_a300f1fefd6652b6] = struct{}{}
		}
	}
	return &UniqueRand{__obf_0dddbc6b117ed615: __obf_0dddbc6b117ed615, __obf_e67e8888d48d964d: __obf_b26d73d753cfcc70, __obf_ef955105287ffc55: __obf_ef955105287ffc55, __obf_9e80bd3b1d9a7224: fmt.Sprintf("%%0%dd", __obf_ddbb4cb554aa04a0)}, nil
}

func (__obf_0962c376eb75c625 *UniqueRand) Generate() string {
	__obf_0962c376eb75c625.__obf_15719f79f826bdc4.
		Lock()
	defer __obf_0962c376eb75c625.__obf_15719f79f826bdc4.Unlock()

	if __obf_0962c376eb75c625.__obf_ef955105287ffc55 > 0 && len(__obf_0962c376eb75c625.__obf_0dddbc6b117ed615) >= __obf_0962c376eb75c625.__obf_ef955105287ffc55 {
		return ""
	}

	var __obf_820762e8a994bf2b bool
	var __obf_1ef2f5d835c97294 string
	for {
		__obf_1ef2f5d835c97294 = fmt.Sprintf(__obf_0962c376eb75c625.__obf_9e80bd3b1d9a7224, __obf_0962c376eb75c625.__obf_e67e8888d48d964d.Int()%__obf_0962c376eb75c625.__obf_ef955105287ffc55)

		if _, __obf_820762e8a994bf2b = __obf_0962c376eb75c625.__obf_0dddbc6b117ed615[__obf_1ef2f5d835c97294]; !__obf_820762e8a994bf2b {
			__obf_0962c376eb75c625.__obf_0dddbc6b117ed615[__obf_1ef2f5d835c97294] = struct{}{}
			return __obf_1ef2f5d835c97294
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_0962c376eb75c625 *UniqueRand) Drop(__obf_0457192af94f8356 ...string) {
	__obf_0962c376eb75c625.__obf_15719f79f826bdc4.
		Lock()
	defer __obf_0962c376eb75c625.__obf_15719f79f826bdc4.Unlock()
	for _, __obf_a300f1fefd6652b6 := range __obf_0457192af94f8356 {
		delete(__obf_0962c376eb75c625.__obf_0dddbc6b117ed615, __obf_a300f1fefd6652b6)
	}
}
