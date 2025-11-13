package __obf_417508f5ae215d0a

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_d87314b4d756b35e sync.Mutex
	__obf_308b7bd03a91519c map[string]struct{} // 追踪已经生成过的
	__obf_0a92aebefacfe4a3 *rand.Rand          // 基础随机数
	__obf_d4bb5d68698de2e1 string              // ID 格式
	__obf_87542ba969cf7322 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_1e077a7771d6af3a int, __obf_2c5f6a694027c639 []string) (*UniqueRand, error) {
	if __obf_1e077a7771d6af3a > 19 || __obf_1e077a7771d6af3a < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_e70a7e9e64987893 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_87542ba969cf7322, __obf_244f98b217fca15b := stats.IntPow(9, __obf_1e077a7771d6af3a), len(__obf_2c5f6a694027c639)

	if __obf_87542ba969cf7322 <= __obf_244f98b217fca15b {
		return nil, ErrIDExhaustied
	}
	__obf_308b7bd03a91519c := make(map[string]struct{}, __obf_244f98b217fca15b)
	for _, __obf_a9527a68ae1dfba0 := range __obf_2c5f6a694027c639 {
		if len(__obf_a9527a68ae1dfba0) == __obf_1e077a7771d6af3a {
			__obf_308b7bd03a91519c[__obf_a9527a68ae1dfba0] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_308b7bd03a91519c: __obf_308b7bd03a91519c,
		__obf_0a92aebefacfe4a3: __obf_e70a7e9e64987893,
		__obf_87542ba969cf7322: __obf_87542ba969cf7322,
		__obf_d4bb5d68698de2e1: fmt.Sprintf("%%0%dd", __obf_1e077a7771d6af3a),
	}, nil
}

func (__obf_feb8e5a89f005fb4 *UniqueRand) Generate() string {
	__obf_feb8e5a89f005fb4.__obf_d87314b4d756b35e.Lock()
	defer __obf_feb8e5a89f005fb4.__obf_d87314b4d756b35e.Unlock()

	if __obf_feb8e5a89f005fb4.__obf_87542ba969cf7322 > 0 && len(__obf_feb8e5a89f005fb4.__obf_308b7bd03a91519c) >= __obf_feb8e5a89f005fb4.__obf_87542ba969cf7322 {
		return ""
	}

	var __obf_0bd9d9196038949b bool
	var __obf_6e993b9e217ddd26 string
	for {
		__obf_6e993b9e217ddd26 = fmt.Sprintf(__obf_feb8e5a89f005fb4.__obf_d4bb5d68698de2e1, __obf_feb8e5a89f005fb4.__obf_0a92aebefacfe4a3.Int()%__obf_feb8e5a89f005fb4.__obf_87542ba969cf7322)

		if _, __obf_0bd9d9196038949b = __obf_feb8e5a89f005fb4.__obf_308b7bd03a91519c[__obf_6e993b9e217ddd26]; !__obf_0bd9d9196038949b {
			__obf_feb8e5a89f005fb4.__obf_308b7bd03a91519c[__obf_6e993b9e217ddd26] = struct{}{}
			return __obf_6e993b9e217ddd26
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_feb8e5a89f005fb4 *UniqueRand) Drop(__obf_c08be9797e182535 ...string) {
	__obf_feb8e5a89f005fb4.__obf_d87314b4d756b35e.Lock()
	defer __obf_feb8e5a89f005fb4.__obf_d87314b4d756b35e.Unlock()
	for _, __obf_a9527a68ae1dfba0 := range __obf_c08be9797e182535 {
		delete(__obf_feb8e5a89f005fb4.__obf_308b7bd03a91519c, __obf_a9527a68ae1dfba0)
	}
}
