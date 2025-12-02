package __obf_34ce7ee87a5aa6b7

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_ca950ed529feb0aa sync.Mutex
	__obf_1fcf58928447e389 map[string]struct{}
	__obf_427e7770a91e7063 * // 追踪已经生成过的
	rand.Rand
	__obf_614ea3d51a4c4b66 string // 基础随机数
	__obf_765fa9d15dbef551 int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_6007fa86825567e7 int, __obf_4519e2d5d5d14c6f []string) (*UniqueRand, error) {
	if __obf_6007fa86825567e7 > 19 || __obf_6007fa86825567e7 < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_915d69697aec0ced := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_765fa9d15dbef551, __obf_70dbcd262cb8b14f := stats.IntPow(9, __obf_6007fa86825567e7), len(__obf_4519e2d5d5d14c6f)

	if __obf_765fa9d15dbef551 <= __obf_70dbcd262cb8b14f {
		return nil, ErrIDExhaustied
	}
	__obf_1fcf58928447e389 := make(map[string]struct{}, __obf_70dbcd262cb8b14f)
	for _, __obf_29b927ea91ca599d := range __obf_4519e2d5d5d14c6f {
		if len(__obf_29b927ea91ca599d) == __obf_6007fa86825567e7 {
			__obf_1fcf58928447e389[__obf_29b927ea91ca599d] = struct{}{}
		}
	}
	return &UniqueRand{__obf_1fcf58928447e389: __obf_1fcf58928447e389, __obf_427e7770a91e7063: __obf_915d69697aec0ced, __obf_765fa9d15dbef551: __obf_765fa9d15dbef551, __obf_614ea3d51a4c4b66: fmt.Sprintf("%%0%dd", __obf_6007fa86825567e7)}, nil
}

func (__obf_e722acd2b7e3a316 *UniqueRand) Generate() string {
	__obf_e722acd2b7e3a316.__obf_ca950ed529feb0aa.
		Lock()
	defer __obf_e722acd2b7e3a316.__obf_ca950ed529feb0aa.Unlock()

	if __obf_e722acd2b7e3a316.__obf_765fa9d15dbef551 > 0 && len(__obf_e722acd2b7e3a316.__obf_1fcf58928447e389) >= __obf_e722acd2b7e3a316.__obf_765fa9d15dbef551 {
		return ""
	}

	var __obf_04415e5e48fee761 bool
	var __obf_5c1ed62ff2bd2f1a string
	for {
		__obf_5c1ed62ff2bd2f1a = fmt.Sprintf(__obf_e722acd2b7e3a316.__obf_614ea3d51a4c4b66, __obf_e722acd2b7e3a316.__obf_427e7770a91e7063.Int()%__obf_e722acd2b7e3a316.__obf_765fa9d15dbef551)

		if _, __obf_04415e5e48fee761 = __obf_e722acd2b7e3a316.__obf_1fcf58928447e389[__obf_5c1ed62ff2bd2f1a]; !__obf_04415e5e48fee761 {
			__obf_e722acd2b7e3a316.__obf_1fcf58928447e389[__obf_5c1ed62ff2bd2f1a] = struct{}{}
			return __obf_5c1ed62ff2bd2f1a
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_e722acd2b7e3a316 *UniqueRand) Drop(__obf_1c7bcf8738840124 ...string) {
	__obf_e722acd2b7e3a316.__obf_ca950ed529feb0aa.
		Lock()
	defer __obf_e722acd2b7e3a316.__obf_ca950ed529feb0aa.Unlock()
	for _, __obf_29b927ea91ca599d := range __obf_1c7bcf8738840124 {
		delete(__obf_e722acd2b7e3a316.__obf_1fcf58928447e389, __obf_29b927ea91ca599d)
	}
}
