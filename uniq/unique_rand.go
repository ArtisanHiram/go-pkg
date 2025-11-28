package __obf_e2239f4853c61591

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_19980aa09194e26b sync.Mutex
	__obf_c537817dd3670843 map[string]struct{} // 追踪已经生成过的
	__obf_6547fdd8df243eca *rand.Rand          // 基础随机数
	__obf_f355d56a8e3116bf string              // ID 格式
	__obf_5c2fbf6ee723cd97 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_9e7527fdc8f56dea int, __obf_751c01bf569b0320 []string) (*UniqueRand, error) {
	if __obf_9e7527fdc8f56dea > 19 || __obf_9e7527fdc8f56dea < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_8e99b4261f870e38 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_5c2fbf6ee723cd97, __obf_dc5c58bf81bb657e := stats.IntPow(9, __obf_9e7527fdc8f56dea), len(__obf_751c01bf569b0320)

	if __obf_5c2fbf6ee723cd97 <= __obf_dc5c58bf81bb657e {
		return nil, ErrIDExhaustied
	}
	__obf_c537817dd3670843 := make(map[string]struct{}, __obf_dc5c58bf81bb657e)
	for _, __obf_1ab57eab018cfbd0 := range __obf_751c01bf569b0320 {
		if len(__obf_1ab57eab018cfbd0) == __obf_9e7527fdc8f56dea {
			__obf_c537817dd3670843[__obf_1ab57eab018cfbd0] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_c537817dd3670843: __obf_c537817dd3670843,
		__obf_6547fdd8df243eca: __obf_8e99b4261f870e38,
		__obf_5c2fbf6ee723cd97: __obf_5c2fbf6ee723cd97,
		__obf_f355d56a8e3116bf: fmt.Sprintf("%%0%dd", __obf_9e7527fdc8f56dea),
	}, nil
}

func (__obf_ac971c48d174c3f7 *UniqueRand) Generate() string {
	__obf_ac971c48d174c3f7.__obf_19980aa09194e26b.Lock()
	defer __obf_ac971c48d174c3f7.__obf_19980aa09194e26b.Unlock()

	if __obf_ac971c48d174c3f7.__obf_5c2fbf6ee723cd97 > 0 && len(__obf_ac971c48d174c3f7.__obf_c537817dd3670843) >= __obf_ac971c48d174c3f7.__obf_5c2fbf6ee723cd97 {
		return ""
	}

	var __obf_c8f6ef41918857c4 bool
	var __obf_38a57c7e430e9465 string
	for {
		__obf_38a57c7e430e9465 = fmt.Sprintf(__obf_ac971c48d174c3f7.__obf_f355d56a8e3116bf, __obf_ac971c48d174c3f7.__obf_6547fdd8df243eca.Int()%__obf_ac971c48d174c3f7.__obf_5c2fbf6ee723cd97)

		if _, __obf_c8f6ef41918857c4 = __obf_ac971c48d174c3f7.__obf_c537817dd3670843[__obf_38a57c7e430e9465]; !__obf_c8f6ef41918857c4 {
			__obf_ac971c48d174c3f7.__obf_c537817dd3670843[__obf_38a57c7e430e9465] = struct{}{}
			return __obf_38a57c7e430e9465
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_ac971c48d174c3f7 *UniqueRand) Drop(__obf_052c44fc336f6520 ...string) {
	__obf_ac971c48d174c3f7.__obf_19980aa09194e26b.Lock()
	defer __obf_ac971c48d174c3f7.__obf_19980aa09194e26b.Unlock()
	for _, __obf_1ab57eab018cfbd0 := range __obf_052c44fc336f6520 {
		delete(__obf_ac971c48d174c3f7.__obf_c537817dd3670843, __obf_1ab57eab018cfbd0)
	}
}
