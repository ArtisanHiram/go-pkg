package __obf_2f51f7d26a2bcdf8

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_54867e7d58cb9039 sync.Mutex
	__obf_cf843cb05731312c map[string]struct{}
	__obf_477df503d01c41d7 * // 追踪已经生成过的
	rand.Rand
	__obf_d14e6a19c031a1ea string // 基础随机数
	__obf_ed62ce100bb7c7ac int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_8e2d436fc3f0566e int, __obf_98c8b6fbbcd74fd8 []string) (*UniqueRand, error) {
	if __obf_8e2d436fc3f0566e > 19 || __obf_8e2d436fc3f0566e < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_0850ae808f7ac83f := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_ed62ce100bb7c7ac, __obf_7e3355f4c3ab00f4 := stats.IntPow(9, __obf_8e2d436fc3f0566e), len(__obf_98c8b6fbbcd74fd8)

	if __obf_ed62ce100bb7c7ac <= __obf_7e3355f4c3ab00f4 {
		return nil, ErrIDExhaustied
	}
	__obf_cf843cb05731312c := make(map[string]struct{}, __obf_7e3355f4c3ab00f4)
	for _, __obf_1caf1c7883b88b75 := range __obf_98c8b6fbbcd74fd8 {
		if len(__obf_1caf1c7883b88b75) == __obf_8e2d436fc3f0566e {
			__obf_cf843cb05731312c[__obf_1caf1c7883b88b75] = struct{}{}
		}
	}
	return &UniqueRand{__obf_cf843cb05731312c: __obf_cf843cb05731312c, __obf_477df503d01c41d7: __obf_0850ae808f7ac83f, __obf_ed62ce100bb7c7ac: __obf_ed62ce100bb7c7ac, __obf_d14e6a19c031a1ea: fmt.Sprintf("%%0%dd", __obf_8e2d436fc3f0566e)}, nil
}

func (__obf_feb0d3963f34862a *UniqueRand) Generate() string {
	__obf_feb0d3963f34862a.__obf_54867e7d58cb9039.
		Lock()
	defer __obf_feb0d3963f34862a.__obf_54867e7d58cb9039.Unlock()

	if __obf_feb0d3963f34862a.__obf_ed62ce100bb7c7ac > 0 && len(__obf_feb0d3963f34862a.__obf_cf843cb05731312c) >= __obf_feb0d3963f34862a.__obf_ed62ce100bb7c7ac {
		return ""
	}

	var __obf_149b4fd56461a8a0 bool
	var __obf_9a4933bb8fba4943 string
	for {
		__obf_9a4933bb8fba4943 = fmt.Sprintf(__obf_feb0d3963f34862a.__obf_d14e6a19c031a1ea, __obf_feb0d3963f34862a.__obf_477df503d01c41d7.Int()%__obf_feb0d3963f34862a.__obf_ed62ce100bb7c7ac)

		if _, __obf_149b4fd56461a8a0 = __obf_feb0d3963f34862a.__obf_cf843cb05731312c[__obf_9a4933bb8fba4943]; !__obf_149b4fd56461a8a0 {
			__obf_feb0d3963f34862a.__obf_cf843cb05731312c[__obf_9a4933bb8fba4943] = struct{}{}
			return __obf_9a4933bb8fba4943
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_feb0d3963f34862a *UniqueRand) Drop(__obf_6e8949ad71eb56e7 ...string) {
	__obf_feb0d3963f34862a.__obf_54867e7d58cb9039.
		Lock()
	defer __obf_feb0d3963f34862a.__obf_54867e7d58cb9039.Unlock()
	for _, __obf_1caf1c7883b88b75 := range __obf_6e8949ad71eb56e7 {
		delete(__obf_feb0d3963f34862a.__obf_cf843cb05731312c, __obf_1caf1c7883b88b75)
	}
}
