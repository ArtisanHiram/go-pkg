package __obf_7d8ac56be2e11a40

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_36f601912d1c90f0 sync.Mutex
	__obf_40425dda65d2d12f map[string]struct{} // 追踪已经生成过的
	__obf_03f20a63f1ecffc1 *rand.Rand          // 基础随机数
	__obf_483540ad095b3df7 string              // ID 格式
	__obf_4bd4babb44bef9b5 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_0c891168f7223a13 int, __obf_01ea039b15c850f4 []string) (*UniqueRand, error) {
	if __obf_0c891168f7223a13 > 19 || __obf_0c891168f7223a13 < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_6e0637b91c9ca30e := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_4bd4babb44bef9b5, __obf_24ee02177ee0ed66 := stats.IntPow(9, __obf_0c891168f7223a13), len(__obf_01ea039b15c850f4)

	if __obf_4bd4babb44bef9b5 <= __obf_24ee02177ee0ed66 {
		return nil, ErrIDExhaustied
	}
	__obf_40425dda65d2d12f := make(map[string]struct{}, __obf_24ee02177ee0ed66)
	for _, __obf_88dfec73463ba77a := range __obf_01ea039b15c850f4 {
		if len(__obf_88dfec73463ba77a) == __obf_0c891168f7223a13 {
			__obf_40425dda65d2d12f[__obf_88dfec73463ba77a] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_40425dda65d2d12f: __obf_40425dda65d2d12f,
		__obf_03f20a63f1ecffc1: __obf_6e0637b91c9ca30e,
		__obf_4bd4babb44bef9b5: __obf_4bd4babb44bef9b5,
		__obf_483540ad095b3df7: fmt.Sprintf("%%0%dd", __obf_0c891168f7223a13),
	}, nil
}

func (__obf_841c11dc9eb6471a *UniqueRand) Generate() string {
	__obf_841c11dc9eb6471a.__obf_36f601912d1c90f0.Lock()
	defer __obf_841c11dc9eb6471a.__obf_36f601912d1c90f0.Unlock()

	if __obf_841c11dc9eb6471a.__obf_4bd4babb44bef9b5 > 0 && len(__obf_841c11dc9eb6471a.__obf_40425dda65d2d12f) >= __obf_841c11dc9eb6471a.__obf_4bd4babb44bef9b5 {
		return ""
	}

	var __obf_4f3a1b70936ef093 bool
	var __obf_5a3113a41c0f69a6 string
	for {
		__obf_5a3113a41c0f69a6 = fmt.Sprintf(__obf_841c11dc9eb6471a.__obf_483540ad095b3df7, __obf_841c11dc9eb6471a.__obf_03f20a63f1ecffc1.Int()%__obf_841c11dc9eb6471a.__obf_4bd4babb44bef9b5)

		if _, __obf_4f3a1b70936ef093 = __obf_841c11dc9eb6471a.__obf_40425dda65d2d12f[__obf_5a3113a41c0f69a6]; !__obf_4f3a1b70936ef093 {
			__obf_841c11dc9eb6471a.__obf_40425dda65d2d12f[__obf_5a3113a41c0f69a6] = struct{}{}
			return __obf_5a3113a41c0f69a6
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_841c11dc9eb6471a *UniqueRand) Drop(__obf_9f6cebe65b5deddc ...string) {
	__obf_841c11dc9eb6471a.__obf_36f601912d1c90f0.Lock()
	defer __obf_841c11dc9eb6471a.__obf_36f601912d1c90f0.Unlock()
	for _, __obf_88dfec73463ba77a := range __obf_9f6cebe65b5deddc {
		delete(__obf_841c11dc9eb6471a.__obf_40425dda65d2d12f, __obf_88dfec73463ba77a)
	}
}
