package __obf_3747a7e09ff475ee

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_1d87674837946b18 sync.Mutex
	__obf_43c5aef97381d02a map[string]struct{}
	__obf_52673c5d7080ff8a * // 追踪已经生成过的
	rand.Rand
	__obf_562651da80e97ff6 string // 基础随机数
	__obf_62fbdc826f412aa0 int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_a4e2162ac7a584e4 int, __obf_ec612b9003ac9421 []string) (*UniqueRand, error) {
	if __obf_a4e2162ac7a584e4 > 19 || __obf_a4e2162ac7a584e4 < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_0586d2e6de52cce2 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_62fbdc826f412aa0, __obf_ecd824e1688b4abe := stats.IntPow(9, __obf_a4e2162ac7a584e4), len(__obf_ec612b9003ac9421)

	if __obf_62fbdc826f412aa0 <= __obf_ecd824e1688b4abe {
		return nil, ErrIDExhaustied
	}
	__obf_43c5aef97381d02a := make(map[string]struct{}, __obf_ecd824e1688b4abe)
	for _, __obf_5ed0238ba74cbb91 := range __obf_ec612b9003ac9421 {
		if len(__obf_5ed0238ba74cbb91) == __obf_a4e2162ac7a584e4 {
			__obf_43c5aef97381d02a[__obf_5ed0238ba74cbb91] = struct{}{}
		}
	}
	return &UniqueRand{__obf_43c5aef97381d02a: __obf_43c5aef97381d02a, __obf_52673c5d7080ff8a: __obf_0586d2e6de52cce2, __obf_62fbdc826f412aa0: __obf_62fbdc826f412aa0, __obf_562651da80e97ff6: fmt.Sprintf("%%0%dd", __obf_a4e2162ac7a584e4)}, nil
}

func (__obf_b4dcaafcadeeef58 *UniqueRand) Generate() string {
	__obf_b4dcaafcadeeef58.__obf_1d87674837946b18.
		Lock()
	defer __obf_b4dcaafcadeeef58.__obf_1d87674837946b18.Unlock()

	if __obf_b4dcaafcadeeef58.__obf_62fbdc826f412aa0 > 0 && len(__obf_b4dcaafcadeeef58.__obf_43c5aef97381d02a) >= __obf_b4dcaafcadeeef58.__obf_62fbdc826f412aa0 {
		return ""
	}

	var __obf_0941b7586b921890 bool
	var __obf_a4042898d22f95f6 string
	for {
		__obf_a4042898d22f95f6 = fmt.Sprintf(__obf_b4dcaafcadeeef58.__obf_562651da80e97ff6, __obf_b4dcaafcadeeef58.__obf_52673c5d7080ff8a.Int()%__obf_b4dcaafcadeeef58.__obf_62fbdc826f412aa0)

		if _, __obf_0941b7586b921890 = __obf_b4dcaafcadeeef58.__obf_43c5aef97381d02a[__obf_a4042898d22f95f6]; !__obf_0941b7586b921890 {
			__obf_b4dcaafcadeeef58.__obf_43c5aef97381d02a[__obf_a4042898d22f95f6] = struct{}{}
			return __obf_a4042898d22f95f6
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_b4dcaafcadeeef58 *UniqueRand) Drop(__obf_370dac7a83e0c78f ...string) {
	__obf_b4dcaafcadeeef58.__obf_1d87674837946b18.
		Lock()
	defer __obf_b4dcaafcadeeef58.__obf_1d87674837946b18.Unlock()
	for _, __obf_5ed0238ba74cbb91 := range __obf_370dac7a83e0c78f {
		delete(__obf_b4dcaafcadeeef58.__obf_43c5aef97381d02a, __obf_5ed0238ba74cbb91)
	}
}
