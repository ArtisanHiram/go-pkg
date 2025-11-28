package __obf_0f0e0d1ff72f3ff0

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_8acdbb3d62aaf7ee sync.Mutex
	__obf_0b4afe7e3481ddca map[string]struct{} // 追踪已经生成过的
	__obf_b5931328e9a79d42 *rand.Rand          // 基础随机数
	__obf_ff67db0d480bd97a string              // ID 格式
	__obf_aea8aec9c6ab3e89 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_ec72a0dddc05cff8 int, __obf_d58360b14c50595d []string) (*UniqueRand, error) {
	if __obf_ec72a0dddc05cff8 > 19 || __obf_ec72a0dddc05cff8 < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_bdeb5b714f331ec6 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_aea8aec9c6ab3e89, __obf_14f79114b751cdfd := stats.IntPow(9, __obf_ec72a0dddc05cff8), len(__obf_d58360b14c50595d)

	if __obf_aea8aec9c6ab3e89 <= __obf_14f79114b751cdfd {
		return nil, ErrIDExhaustied
	}
	__obf_0b4afe7e3481ddca := make(map[string]struct{}, __obf_14f79114b751cdfd)
	for _, __obf_ffc5d2d602752af4 := range __obf_d58360b14c50595d {
		if len(__obf_ffc5d2d602752af4) == __obf_ec72a0dddc05cff8 {
			__obf_0b4afe7e3481ddca[__obf_ffc5d2d602752af4] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_0b4afe7e3481ddca: __obf_0b4afe7e3481ddca,
		__obf_b5931328e9a79d42: __obf_bdeb5b714f331ec6,
		__obf_aea8aec9c6ab3e89: __obf_aea8aec9c6ab3e89,
		__obf_ff67db0d480bd97a: fmt.Sprintf("%%0%dd", __obf_ec72a0dddc05cff8),
	}, nil
}

func (__obf_c5c91f357c087ad4 *UniqueRand) Generate() string {
	__obf_c5c91f357c087ad4.__obf_8acdbb3d62aaf7ee.Lock()
	defer __obf_c5c91f357c087ad4.__obf_8acdbb3d62aaf7ee.Unlock()

	if __obf_c5c91f357c087ad4.__obf_aea8aec9c6ab3e89 > 0 && len(__obf_c5c91f357c087ad4.__obf_0b4afe7e3481ddca) >= __obf_c5c91f357c087ad4.__obf_aea8aec9c6ab3e89 {
		return ""
	}

	var __obf_b08eebaa4bd2dee3 bool
	var __obf_0725dd8f5f5e1058 string
	for {
		__obf_0725dd8f5f5e1058 = fmt.Sprintf(__obf_c5c91f357c087ad4.__obf_ff67db0d480bd97a, __obf_c5c91f357c087ad4.__obf_b5931328e9a79d42.Int()%__obf_c5c91f357c087ad4.__obf_aea8aec9c6ab3e89)

		if _, __obf_b08eebaa4bd2dee3 = __obf_c5c91f357c087ad4.__obf_0b4afe7e3481ddca[__obf_0725dd8f5f5e1058]; !__obf_b08eebaa4bd2dee3 {
			__obf_c5c91f357c087ad4.__obf_0b4afe7e3481ddca[__obf_0725dd8f5f5e1058] = struct{}{}
			return __obf_0725dd8f5f5e1058
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_c5c91f357c087ad4 *UniqueRand) Drop(__obf_cdcb9c5694652477 ...string) {
	__obf_c5c91f357c087ad4.__obf_8acdbb3d62aaf7ee.Lock()
	defer __obf_c5c91f357c087ad4.__obf_8acdbb3d62aaf7ee.Unlock()
	for _, __obf_ffc5d2d602752af4 := range __obf_cdcb9c5694652477 {
		delete(__obf_c5c91f357c087ad4.__obf_0b4afe7e3481ddca, __obf_ffc5d2d602752af4)
	}
}
