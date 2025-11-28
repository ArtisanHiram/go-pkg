package __obf_8fe28252c1b01dfe

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_26212263d98d9893 sync.Mutex
	__obf_c77bc476a039bd1a map[string]struct{} // 追踪已经生成过的
	__obf_be567146b6a23b53 *rand.Rand          // 基础随机数
	__obf_52e80bcde9b46afa string              // ID 格式
	__obf_f4276f4fdcc61367 int                 // 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_cc79f69d55c771d6 int, __obf_170fe6f489890d2b []string) (*UniqueRand, error) {
	if __obf_cc79f69d55c771d6 > 19 || __obf_cc79f69d55c771d6 < 0 {
		return nil, ErrInvalidIDLength
	}

	__obf_392a5c26963c1117 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_f4276f4fdcc61367, __obf_5ceda09ca95e585a := stats.IntPow(9, __obf_cc79f69d55c771d6), len(__obf_170fe6f489890d2b)

	if __obf_f4276f4fdcc61367 <= __obf_5ceda09ca95e585a {
		return nil, ErrIDExhaustied
	}
	__obf_c77bc476a039bd1a := make(map[string]struct{}, __obf_5ceda09ca95e585a)
	for _, __obf_3b548eea5a5f82ce := range __obf_170fe6f489890d2b {
		if len(__obf_3b548eea5a5f82ce) == __obf_cc79f69d55c771d6 {
			__obf_c77bc476a039bd1a[__obf_3b548eea5a5f82ce] = struct{}{}
		}
	}
	return &UniqueRand{
		__obf_c77bc476a039bd1a: __obf_c77bc476a039bd1a,
		__obf_be567146b6a23b53: __obf_392a5c26963c1117,
		__obf_f4276f4fdcc61367: __obf_f4276f4fdcc61367,
		__obf_52e80bcde9b46afa: fmt.Sprintf("%%0%dd", __obf_cc79f69d55c771d6),
	}, nil
}

func (__obf_6a86d57e2bf1634f *UniqueRand) Generate() string {
	__obf_6a86d57e2bf1634f.__obf_26212263d98d9893.Lock()
	defer __obf_6a86d57e2bf1634f.__obf_26212263d98d9893.Unlock()

	if __obf_6a86d57e2bf1634f.__obf_f4276f4fdcc61367 > 0 && len(__obf_6a86d57e2bf1634f.__obf_c77bc476a039bd1a) >= __obf_6a86d57e2bf1634f.__obf_f4276f4fdcc61367 {
		return ""
	}

	var __obf_a3279a70b6b813b3 bool
	var __obf_a6dd0a7023f6fe3f string
	for {
		__obf_a6dd0a7023f6fe3f = fmt.Sprintf(__obf_6a86d57e2bf1634f.__obf_52e80bcde9b46afa, __obf_6a86d57e2bf1634f.__obf_be567146b6a23b53.Int()%__obf_6a86d57e2bf1634f.__obf_f4276f4fdcc61367)

		if _, __obf_a3279a70b6b813b3 = __obf_6a86d57e2bf1634f.__obf_c77bc476a039bd1a[__obf_a6dd0a7023f6fe3f]; !__obf_a3279a70b6b813b3 {
			__obf_6a86d57e2bf1634f.__obf_c77bc476a039bd1a[__obf_a6dd0a7023f6fe3f] = struct{}{}
			return __obf_a6dd0a7023f6fe3f
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_6a86d57e2bf1634f *UniqueRand) Drop(__obf_b01a990908319d26 ...string) {
	__obf_6a86d57e2bf1634f.__obf_26212263d98d9893.Lock()
	defer __obf_6a86d57e2bf1634f.__obf_26212263d98d9893.Unlock()
	for _, __obf_3b548eea5a5f82ce := range __obf_b01a990908319d26 {
		delete(__obf_6a86d57e2bf1634f.__obf_c77bc476a039bd1a, __obf_3b548eea5a5f82ce)
	}
}
