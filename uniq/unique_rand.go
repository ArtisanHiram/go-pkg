package __obf_a51a64e21f311927

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_7017cfd2525a0d6a sync.Mutex
	__obf_e33a08c1807f922c map[string]struct{}
	__obf_2f0dd74324bcccf3 * // 追踪已经生成过的
	rand.Rand
	__obf_dbf91e26a7514edb string // 基础随机数
	__obf_b4a2c4aa6e7b7d78 int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_41f853eb812daa0f int, __obf_55f91dc5d963eb6f []string) (*UniqueRand, error) {
	if __obf_41f853eb812daa0f > 19 || __obf_41f853eb812daa0f < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_0ed5f7a88929d5ab := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_b4a2c4aa6e7b7d78, __obf_1ab4380c3e72b1b0 := stats.IntPow(9, __obf_41f853eb812daa0f), len(__obf_55f91dc5d963eb6f)

	if __obf_b4a2c4aa6e7b7d78 <= __obf_1ab4380c3e72b1b0 {
		return nil, ErrIDExhaustied
	}
	__obf_e33a08c1807f922c := make(map[string]struct{}, __obf_1ab4380c3e72b1b0)
	for _, __obf_8c882dce5ca83eeb := range __obf_55f91dc5d963eb6f {
		if len(__obf_8c882dce5ca83eeb) == __obf_41f853eb812daa0f {
			__obf_e33a08c1807f922c[__obf_8c882dce5ca83eeb] = struct{}{}
		}
	}
	return &UniqueRand{__obf_e33a08c1807f922c: __obf_e33a08c1807f922c, __obf_2f0dd74324bcccf3: __obf_0ed5f7a88929d5ab, __obf_b4a2c4aa6e7b7d78: __obf_b4a2c4aa6e7b7d78, __obf_dbf91e26a7514edb: fmt.Sprintf("%%0%dd", __obf_41f853eb812daa0f)}, nil
}

func (__obf_d5f5a3a4687208bc *UniqueRand) Generate() string {
	__obf_d5f5a3a4687208bc.__obf_7017cfd2525a0d6a.
		Lock()
	defer __obf_d5f5a3a4687208bc.__obf_7017cfd2525a0d6a.Unlock()

	if __obf_d5f5a3a4687208bc.__obf_b4a2c4aa6e7b7d78 > 0 && len(__obf_d5f5a3a4687208bc.__obf_e33a08c1807f922c) >= __obf_d5f5a3a4687208bc.__obf_b4a2c4aa6e7b7d78 {
		return ""
	}

	var __obf_80eec96aa83c2efa bool
	var __obf_c022f98255ac9ddc string
	for {
		__obf_c022f98255ac9ddc = fmt.Sprintf(__obf_d5f5a3a4687208bc.__obf_dbf91e26a7514edb, __obf_d5f5a3a4687208bc.__obf_2f0dd74324bcccf3.Int()%__obf_d5f5a3a4687208bc.__obf_b4a2c4aa6e7b7d78)

		if _, __obf_80eec96aa83c2efa = __obf_d5f5a3a4687208bc.__obf_e33a08c1807f922c[__obf_c022f98255ac9ddc]; !__obf_80eec96aa83c2efa {
			__obf_d5f5a3a4687208bc.__obf_e33a08c1807f922c[__obf_c022f98255ac9ddc] = struct{}{}
			return __obf_c022f98255ac9ddc
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_d5f5a3a4687208bc *UniqueRand) Drop(__obf_f609cc0264590f39 ...string) {
	__obf_d5f5a3a4687208bc.__obf_7017cfd2525a0d6a.
		Lock()
	defer __obf_d5f5a3a4687208bc.__obf_7017cfd2525a0d6a.Unlock()
	for _, __obf_8c882dce5ca83eeb := range __obf_f609cc0264590f39 {
		delete(__obf_d5f5a3a4687208bc.__obf_e33a08c1807f922c, __obf_8c882dce5ca83eeb)
	}
}
