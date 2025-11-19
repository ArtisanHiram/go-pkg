package __obf_7913809aab6c8423

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_cb0ab65220100005 sync.Mutex
	__obf_19dd3f6171079936 map[string]struct{} // 追踪已经生成过的
	__obf_d77c8f4b6a147d22 *rand.Rand          // 基础随机数
	__obf_2986b691e23289ff string              // ID 格式
	__obf_0705ecb79a04ca29 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_439a4cb765b63c55 int, __obf_d0fd21044c58387d []string) (*UniqueRand, error) {
	if __obf_439a4cb765b63c55 > 19 || __obf_439a4cb765b63c55 < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_36fae8fbdda34103 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_0705ecb79a04ca29, __obf_6f36e269383e30fc := stats.IntPow(9, __obf_439a4cb765b63c55), len(__obf_d0fd21044c58387d)

	if __obf_0705ecb79a04ca29 <= __obf_6f36e269383e30fc {
		return nil, ErrIDExhaustied
	}
	__obf_19dd3f6171079936 := make(map[string]struct{}, __obf_6f36e269383e30fc)
	for _, __obf_ce304ea72067f76f := range __obf_d0fd21044c58387d {
		if len(__obf_ce304ea72067f76f) == __obf_439a4cb765b63c55 {
			__obf_19dd3f6171079936[__obf_ce304ea72067f76f] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_19dd3f6171079936: __obf_19dd3f6171079936,
		__obf_d77c8f4b6a147d22: __obf_36fae8fbdda34103,
		__obf_0705ecb79a04ca29: __obf_0705ecb79a04ca29,
		__obf_2986b691e23289ff: fmt.Sprintf("%%0%dd", __obf_439a4cb765b63c55),
	}, nil
}

func (__obf_3d63dae1aea4bf71 *UniqueRand) Generate() string {
	__obf_3d63dae1aea4bf71.__obf_cb0ab65220100005.Lock()
	defer __obf_3d63dae1aea4bf71.__obf_cb0ab65220100005.Unlock()

	if __obf_3d63dae1aea4bf71.__obf_0705ecb79a04ca29 > 0 && len(__obf_3d63dae1aea4bf71.__obf_19dd3f6171079936) >= __obf_3d63dae1aea4bf71.__obf_0705ecb79a04ca29 {
		return ""
	}

	var __obf_c8f04275fcb5c686 bool
	var __obf_c2f487c3e92f8673 string
	for {
		__obf_c2f487c3e92f8673 = fmt.Sprintf(__obf_3d63dae1aea4bf71.__obf_2986b691e23289ff, __obf_3d63dae1aea4bf71.__obf_d77c8f4b6a147d22.Int()%__obf_3d63dae1aea4bf71.__obf_0705ecb79a04ca29)

		if _, __obf_c8f04275fcb5c686 = __obf_3d63dae1aea4bf71.__obf_19dd3f6171079936[__obf_c2f487c3e92f8673]; !__obf_c8f04275fcb5c686 {
			__obf_3d63dae1aea4bf71.__obf_19dd3f6171079936[__obf_c2f487c3e92f8673] = struct{}{}
			return __obf_c2f487c3e92f8673
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_3d63dae1aea4bf71 *UniqueRand) Drop(__obf_10e48c2fa6ea0088 ...string) {
	__obf_3d63dae1aea4bf71.__obf_cb0ab65220100005.Lock()
	defer __obf_3d63dae1aea4bf71.__obf_cb0ab65220100005.Unlock()
	for _, __obf_ce304ea72067f76f := range __obf_10e48c2fa6ea0088 {
		delete(__obf_3d63dae1aea4bf71.__obf_19dd3f6171079936, __obf_ce304ea72067f76f)
	}
}
