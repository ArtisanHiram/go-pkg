package __obf_4d7a51f8e2f0494a

import (
	"errors"
	"fmt"
	stats "github.com/ArtisanHiram/go-pkg/stats"
	"math/rand"
	"sync"
	"time"
)

type UniqueRand struct {
	__obf_6d0dc3fbe6f4289e sync.Mutex
	__obf_d4f2b802b42da4ac map[string]struct{}
	__obf_c2e9304dc1461d7c * // 追踪已经生成过的
	rand.Rand
	__obf_ef6adb335fbbb512 string // 基础随机数
	__obf_7ef5b11e66e83b08 int    // ID 格式
	// 有效ID数
}

var (
	ErrIDExhaustied    = errors.New("ID have been exhausted")
	ErrInvalidIDLength = errors.New("ID length is invalid")
)

func NewUniqueRand(__obf_2eb779e3a021aa3e int, __obf_a0ee58603967c7af []string) (*UniqueRand, error) {
	if __obf_2eb779e3a021aa3e > 19 || __obf_2eb779e3a021aa3e < 0 {
		return nil, ErrInvalidIDLength
	}
	__obf_57f129fecc995228 := rand.New(rand.NewSource(time.Now().UnixNano()))
	__obf_7ef5b11e66e83b08, __obf_5f403f4f516d7a24 := stats.IntPow(9, __obf_2eb779e3a021aa3e), len(__obf_a0ee58603967c7af)

	if __obf_7ef5b11e66e83b08 <= __obf_5f403f4f516d7a24 {
		return nil, ErrIDExhaustied
	}
	__obf_d4f2b802b42da4ac := make(map[string]struct{}, __obf_5f403f4f516d7a24)
	for _, __obf_ec9ea9e2ee15569a := range __obf_a0ee58603967c7af {
		if len(__obf_ec9ea9e2ee15569a) == __obf_2eb779e3a021aa3e {
			__obf_d4f2b802b42da4ac[__obf_ec9ea9e2ee15569a] = struct{}{}
		}
	}
	return &UniqueRand{__obf_d4f2b802b42da4ac: __obf_d4f2b802b42da4ac, __obf_c2e9304dc1461d7c: __obf_57f129fecc995228, __obf_7ef5b11e66e83b08: __obf_7ef5b11e66e83b08, __obf_ef6adb335fbbb512: fmt.Sprintf("%%0%dd", __obf_2eb779e3a021aa3e)}, nil
}

func (__obf_51362cca9c9190a3 *UniqueRand) Generate() string {
	__obf_51362cca9c9190a3.__obf_6d0dc3fbe6f4289e.
		Lock()
	defer __obf_51362cca9c9190a3.__obf_6d0dc3fbe6f4289e.Unlock()

	if __obf_51362cca9c9190a3.__obf_7ef5b11e66e83b08 > 0 && len(__obf_51362cca9c9190a3.__obf_d4f2b802b42da4ac) >= __obf_51362cca9c9190a3.__obf_7ef5b11e66e83b08 {
		return ""
	}

	var __obf_a596699de7db8dcb bool
	var __obf_d43488f90f0960a8 string
	for {
		__obf_d43488f90f0960a8 = fmt.Sprintf(__obf_51362cca9c9190a3.__obf_ef6adb335fbbb512, __obf_51362cca9c9190a3.__obf_c2e9304dc1461d7c.Int()%__obf_51362cca9c9190a3.__obf_7ef5b11e66e83b08)

		if _, __obf_a596699de7db8dcb = __obf_51362cca9c9190a3.__obf_d4f2b802b42da4ac[__obf_d43488f90f0960a8]; !__obf_a596699de7db8dcb {
			__obf_51362cca9c9190a3.__obf_d4f2b802b42da4ac[__obf_d43488f90f0960a8] = struct{}{}
			return __obf_d43488f90f0960a8
		}
	}
}

// Drop 对释放的元素也在generated中销毁
func (__obf_51362cca9c9190a3 *UniqueRand) Drop(__obf_ff3de3d32e391733 ...string) {
	__obf_51362cca9c9190a3.__obf_6d0dc3fbe6f4289e.
		Lock()
	defer __obf_51362cca9c9190a3.__obf_6d0dc3fbe6f4289e.Unlock()
	for _, __obf_ec9ea9e2ee15569a := range __obf_ff3de3d32e391733 {
		delete(__obf_51362cca9c9190a3.__obf_d4f2b802b42da4ac, __obf_ec9ea9e2ee15569a)
	}
}
