package __obf_1fac81d42ad620b3

import (
	"bufio"
	"bytes"
	"os"
	"slices"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

const (
	__obf_97764381e23d9272 = 1024
	__obf_9022334ab5d798da = 1_000_000
	__obf_04837e6f2c750af5 = // 1MB
	3
	__obf_4f0d4a8ae933c06f = 100
	__obf_1fc356c6bc5512cd = 16
)

// 位图实现
type __obf_77a1291079edf805 struct {
	__obf_7d3e586629301158 []uint64
}

func __obf_ce546a3d8b5820b6(__obf_c1ca2e0801da064b int) *__obf_77a1291079edf805 {
	__obf_094101b5cdabd7f1 := // 每个 uint64 存储 64 位
		(__obf_c1ca2e0801da064b + 63) / 64
	return &__obf_77a1291079edf805{__obf_7d3e586629301158: make([]uint64, __obf_094101b5cdabd7f1)}
}

func (__obf_2f9a4a868b63151e *__obf_77a1291079edf805) __obf_481806a73a1d3f69(__obf_36d8c69eb86e8573 int) {
	__obf_2f9a4a868b63151e.__obf_7d3e586629301158[__obf_36d8c69eb86e8573/64] |= 1 << (__obf_36d8c69eb86e8573 % 64)
}

func (__obf_2f9a4a868b63151e *__obf_77a1291079edf805) __obf_0dd4369da2f5eb97(__obf_36d8c69eb86e8573 int) bool {
	return __obf_2f9a4a868b63151e.__obf_7d3e586629301158[__obf_36d8c69eb86e8573/64]&(1<<(__obf_36d8c69eb86e8573%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_41b1dc597079a156 *DATTrie
	__obf_9c7dace349c31e7e *WuManberMatcher
	__obf_68f0f6f85153490c sync.RWMutex
}

func (__obf_c0a010e0d7452ac5 *SensitiveFilter) Contains(__obf_324f64750b1cc9de string) bool {
	__obf_c0a010e0d7452ac5.__obf_68f0f6f85153490c.
		Lock()
	defer __obf_c0a010e0d7452ac5.__obf_68f0f6f85153490c.Unlock()
	if len(__obf_324f64750b1cc9de) > __obf_4f0d4a8ae933c06f {
		if __obf_c0a010e0d7452ac5.__obf_9c7dace349c31e7e.Match(__obf_324f64750b1cc9de) {
			return true
		} else {
			return false
		}
	}
	return __obf_c0a010e0d7452ac5.__obf_41b1dc597079a156.Contains(__obf_324f64750b1cc9de)
}

func (__obf_c0a010e0d7452ac5 *SensitiveFilter) FindAll(__obf_324f64750b1cc9de string) []MatchResult {
	if __obf_c0a010e0d7452ac5 == nil || __obf_c0a010e0d7452ac5.__obf_41b1dc597079a156 == nil {
		return nil
	}
	if __obf_324f64750b1cc9de == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_324f64750b1cc9de) > __obf_9022334ab5d798da {
		return nil // 或返回错误
	}
	__obf_a796ebe9c52eb9c7 := // 1. 找出所有可能重叠的匹配项
		__obf_c0a010e0d7452ac5.__obf_41b1dc597079a156.__obf_3b60730717aba921([]rune(__obf_324f64750b1cc9de))
	if len(__obf_a796ebe9c52eb9c7) < 2 {
		return __obf_a796ebe9c52eb9c7
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_a796ebe9c52eb9c7, func(__obf_433c212716a376c9, __obf_a319b29477f20362 int) bool {
		__obf_9c49e40219af10e8 := __obf_a796ebe9c52eb9c7[__obf_433c212716a376c9].End - __obf_a796ebe9c52eb9c7[__obf_433c212716a376c9].Start
		__obf_943015d37ddcd04f := __obf_a796ebe9c52eb9c7[__obf_a319b29477f20362].End - __obf_a796ebe9c52eb9c7[__obf_a319b29477f20362].Start
		if __obf_9c49e40219af10e8 != __obf_943015d37ddcd04f {
			return __obf_9c49e40219af10e8 > __obf_943015d37ddcd04f
		}
		return __obf_a796ebe9c52eb9c7[__obf_433c212716a376c9].Start < __obf_a796ebe9c52eb9c7[__obf_a319b29477f20362].Start
	})
	__obf_249afe6855e8fbf2 := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_a796ebe9c52eb9c7))
	__obf_aab62c88b15846df := // occupied := make([]bool, len(runes))
		__obf_ce546a3d8b5820b6(len([]rune(__obf_324f64750b1cc9de)))

	for _, __obf_73003db12e2b6cc8 := range __obf_a796ebe9c52eb9c7 {
		__obf_b2c4d79fe39309a1 := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_433c212716a376c9 := __obf_73003db12e2b6cc8.Start; __obf_433c212716a376c9 < __obf_73003db12e2b6cc8.End; __obf_433c212716a376c9++ {
			// if occupied[i] {
			if __obf_aab62c88b15846df.__obf_0dd4369da2f5eb97(__obf_433c212716a376c9) {
				__obf_b2c4d79fe39309a1 = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_b2c4d79fe39309a1 {
			__obf_249afe6855e8fbf2 = append(__obf_249afe6855e8fbf2,
				// 标记该位置已被占用
				__obf_73003db12e2b6cc8)

			for __obf_433c212716a376c9 := __obf_73003db12e2b6cc8.Start; __obf_433c212716a376c9 < __obf_73003db12e2b6cc8.End; __obf_433c212716a376c9++ {
				__obf_aab62c88b15846df.
					// occupied[i] = true
					__obf_481806a73a1d3f69(__obf_433c212716a376c9)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_249afe6855e8fbf2, func(__obf_433c212716a376c9, __obf_a319b29477f20362 int) bool {
		return __obf_249afe6855e8fbf2[__obf_433c212716a376c9].Start < __obf_249afe6855e8fbf2[__obf_a319b29477f20362].Start
	})

	return __obf_249afe6855e8fbf2
}

func (__obf_c0a010e0d7452ac5 *SensitiveFilter) Replace(__obf_324f64750b1cc9de string, __obf_ebe47b467b0fcf18 rune) string {
	__obf_28ad7c5823a9e6d3 := __obf_c0a010e0d7452ac5.FindAll(__obf_324f64750b1cc9de)
	if len(__obf_28ad7c5823a9e6d3) == 0 {
		return __obf_324f64750b1cc9de
	}
	__obf_de9b32348df338b4 := []rune(__obf_324f64750b1cc9de)
	for _, __obf_73003db12e2b6cc8 := range __obf_28ad7c5823a9e6d3 {
		for __obf_433c212716a376c9 := __obf_73003db12e2b6cc8.Start; __obf_433c212716a376c9 < __obf_73003db12e2b6cc8.End; __obf_433c212716a376c9++ {
			__obf_de9b32348df338b4[__obf_433c212716a376c9] = __obf_ebe47b467b0fcf18
		}
	}
	return string(__obf_de9b32348df338b4)
}

func (__obf_c0a010e0d7452ac5 *SensitiveFilter) ReplaceWithString(__obf_324f64750b1cc9de string, __obf_ebe47b467b0fcf18 string) string {
	__obf_28ad7c5823a9e6d3 := __obf_c0a010e0d7452ac5.FindAll(__obf_324f64750b1cc9de)
	if len(__obf_28ad7c5823a9e6d3) == 0 {
		return __obf_324f64750b1cc9de
	}

	var __obf_12472d87ab8c4db4 strings.Builder
	__obf_de9b32348df338b4 := []rune(__obf_324f64750b1cc9de)
	__obf_651cae3e54c9ea06 := 0
	for _, __obf_73003db12e2b6cc8 := range __obf_28ad7c5823a9e6d3 {
		__obf_12472d87ab8c4db4.
			WriteString(string(__obf_de9b32348df338b4[__obf_651cae3e54c9ea06:__obf_73003db12e2b6cc8.Start]))
		__obf_12472d87ab8c4db4.
			WriteString(__obf_ebe47b467b0fcf18)
		__obf_651cae3e54c9ea06 = __obf_73003db12e2b6cc8.End
	}
	__obf_12472d87ab8c4db4.
		WriteString(string(__obf_de9b32348df338b4[__obf_651cae3e54c9ea06:]))
	return __obf_12472d87ab8c4db4.String()
}

func (__obf_c0a010e0d7452ac5 *SensitiveFilter) ReplaceWithFunc(__obf_324f64750b1cc9de string, __obf_c3d4a6f9983e132b func(__obf_73003db12e2b6cc8 MatchResult) string) string {
	__obf_28ad7c5823a9e6d3 := __obf_c0a010e0d7452ac5.FindAll(__obf_324f64750b1cc9de)
	if len(__obf_28ad7c5823a9e6d3) == 0 {
		return __obf_324f64750b1cc9de
	}

	var __obf_12472d87ab8c4db4 strings.Builder
	__obf_de9b32348df338b4 := []rune(__obf_324f64750b1cc9de)
	__obf_651cae3e54c9ea06 := 0
	for _, __obf_73003db12e2b6cc8 := range __obf_28ad7c5823a9e6d3 {
		__obf_12472d87ab8c4db4.
			WriteString(string(__obf_de9b32348df338b4[__obf_651cae3e54c9ea06:__obf_73003db12e2b6cc8.Start]))
		__obf_12472d87ab8c4db4.
			WriteString(__obf_c3d4a6f9983e132b(__obf_73003db12e2b6cc8))
		__obf_651cae3e54c9ea06 = __obf_73003db12e2b6cc8.End
	}
	__obf_12472d87ab8c4db4.
		WriteString(string(__obf_de9b32348df338b4[__obf_651cae3e54c9ea06:]))
	return __obf_12472d87ab8c4db4.String()
}

// 添加 Close 方法释放资源
func (__obf_c0a010e0d7452ac5 *SensitiveFilter) Close() {
	__obf_c0a010e0d7452ac5.__obf_68f0f6f85153490c.
		Lock()
	defer __obf_c0a010e0d7452ac5.__obf_68f0f6f85153490c.Unlock()
	__obf_c0a010e0d7452ac5.__obf_41b1dc597079a156 = nil
	__obf_c0a010e0d7452ac5.__obf_9c7dace349c31e7e = nil
}

type FilterBuilder struct {
	__obf_a6226e9bebc42425 map[string]struct{}
	__obf_7a5d4cdb8fa0b164 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_c1ca2e0801da064b int) FilterOption {
	return func(__obf_2f9a4a868b63151e *FilterBuilder) {
		if __obf_c1ca2e0801da064b > 0 {
			__obf_2f9a4a868b63151e.__obf_7a5d4cdb8fa0b164 = __obf_c1ca2e0801da064b
		}
	}
}

func WithWords(__obf_a6226e9bebc42425 []string) FilterOption {
	return func(__obf_2f9a4a868b63151e *FilterBuilder) {
		for _, __obf_a7373fbad1e6bda7 := range __obf_a6226e9bebc42425 {
			if __obf_a7373fbad1e6bda7 != "" {
				__obf_2f9a4a868b63151e.__obf_a6226e9bebc42425[__obf_a7373fbad1e6bda7] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_11d77ad453af3296 ...FilterOption) *FilterBuilder {
	__obf_2f9a4a868b63151e := &FilterBuilder{__obf_a6226e9bebc42425: make(map[string]struct{}), __obf_7a5d4cdb8fa0b164: __obf_04837e6f2c750af5}
	for _, __obf_17f78dcce27d6470 := range __obf_11d77ad453af3296 {
		__obf_17f78dcce27d6470(__obf_2f9a4a868b63151e)
	}
	return __obf_2f9a4a868b63151e
}

func (__obf_2f9a4a868b63151e *FilterBuilder) AddWord(__obf_35e337a7e5b6f6b4 string) {
	if __obf_35e337a7e5b6f6b4 != "" {
		__obf_2f9a4a868b63151e.__obf_a6226e9bebc42425[__obf_35e337a7e5b6f6b4] = struct{}{}
	}
}

func (__obf_2f9a4a868b63151e *FilterBuilder) AddWords(__obf_a6226e9bebc42425 []string) {
	for _, __obf_35e337a7e5b6f6b4 := range __obf_a6226e9bebc42425 {
		__obf_2f9a4a868b63151e.
			AddWord(__obf_35e337a7e5b6f6b4)
	}
}

func (__obf_2f9a4a868b63151e *FilterBuilder) LoadFromFile(__obf_cd9c81598778a84b string) error {
	__obf_b342e5cd3585aa74, __obf_031332cbba1d1688 := os.Open(__obf_cd9c81598778a84b)
	if __obf_031332cbba1d1688 != nil {
		return __obf_031332cbba1d1688
	}
	defer __obf_b342e5cd3585aa74.Close()
	__obf_3bce40c8189a4703 := bufio.NewReader(__obf_b342e5cd3585aa74)
	__obf_4dd35837e2fd0979 := []byte{0xEF, 0xBB, 0xBF}
	__obf_da7ef48098801e8c, __obf_031332cbba1d1688 := __obf_3bce40c8189a4703.Peek(3)
	if __obf_031332cbba1d1688 == nil && bytes.Equal(__obf_da7ef48098801e8c, __obf_4dd35837e2fd0979) {
		_, _ = __obf_3bce40c8189a4703.Discard(3)
	}
	__obf_f995e1c70f0ad7b3 := bufio.NewScanner(__obf_3bce40c8189a4703)
	for __obf_f995e1c70f0ad7b3.Scan() {
		__obf_4d263d53783e363f := strings.TrimSpace(__obf_f995e1c70f0ad7b3.Text())
		__obf_2f9a4a868b63151e.
			AddWord(__obf_4d263d53783e363f)
	}
	return __obf_f995e1c70f0ad7b3.Err()
}

func (__obf_2f9a4a868b63151e *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_2f9a4a868b63151e.__obf_a6226e9bebc42425) == 0 {
		return &SensitiveFilter{__obf_41b1dc597079a156: NewDATTrie(), __obf_9c7dace349c31e7e: NewWuManberMatcher(__obf_2f9a4a868b63151e.__obf_7a5d4cdb8fa0b164)}
	}
	__obf_5ebb7d1f79397bdc := &__obf_0c3ce44e6d22fabb{__obf_13e9d50886dd9c4a: make(map[rune]*__obf_0c3ce44e6d22fabb)}
	__obf_ddaef4f282464ada := make([]string, 0, len(__obf_2f9a4a868b63151e.__obf_a6226e9bebc42425))
	for __obf_35e337a7e5b6f6b4 := range __obf_2f9a4a868b63151e.__obf_a6226e9bebc42425 {
		__obf_ddaef4f282464ada = append(__obf_ddaef4f282464ada, __obf_35e337a7e5b6f6b4)
		__obf_5ebb7d1f79397bdc.__obf_577f6897a13179c6(__obf_35e337a7e5b6f6b4)
	}
	__obf_41b1dc597079a156 := __obf_6f41e9b4f4ad8d1a(__obf_5ebb7d1f79397bdc)
	__obf_9c7dace349c31e7e := NewWuManberMatcher(__obf_2f9a4a868b63151e.__obf_7a5d4cdb8fa0b164)
	__obf_9c7dace349c31e7e.
		AddPatterns(__obf_ddaef4f282464ada)
	__obf_9c7dace349c31e7e.
		Build()

	return &SensitiveFilter{__obf_41b1dc597079a156: __obf_41b1dc597079a156, __obf_9c7dace349c31e7e: __obf_9c7dace349c31e7e}
}

func (__obf_2f9a4a868b63151e *FilterBuilder) Clear() {
	__obf_2f9a4a868b63151e.__obf_a6226e9bebc42425 = make(map[string]struct{})
}

type __obf_0c3ce44e6d22fabb struct {
	__obf_13e9d50886dd9c4a map[rune]*__obf_0c3ce44e6d22fabb
	__obf_218af756657db028 bool
}

func (__obf_094101b5cdabd7f1 *__obf_0c3ce44e6d22fabb) __obf_577f6897a13179c6(__obf_35e337a7e5b6f6b4 string) {
	__obf_2c668bd60383ec67 := __obf_094101b5cdabd7f1
	for _, __obf_dfbedc1c3c638189 := range __obf_35e337a7e5b6f6b4 {
		if _, __obf_ad62883dfb1e5e56 := __obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a[__obf_dfbedc1c3c638189]; !__obf_ad62883dfb1e5e56 {
			__obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a[__obf_dfbedc1c3c638189] = &__obf_0c3ce44e6d22fabb{__obf_13e9d50886dd9c4a: make(map[rune]*__obf_0c3ce44e6d22fabb)}
		}
		__obf_2c668bd60383ec67 = __obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a[__obf_dfbedc1c3c638189]
	}
	__obf_2c668bd60383ec67.__obf_218af756657db028 = true
}

type DATTrie struct {
	__obf_b9ffc11f3620a9c0 []int32
	__obf_633474aeffb3641a []int32
	__obf_8649c60b00d903ba []bool
	__obf_c1ca2e0801da064b int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_b9ffc11f3620a9c0: make([]int32, __obf_97764381e23d9272), __obf_633474aeffb3641a: make([]int32, __obf_97764381e23d9272), __obf_8649c60b00d903ba: make([]bool, __obf_97764381e23d9272), __obf_c1ca2e0801da064b: 1}
}

func __obf_6f41e9b4f4ad8d1a(__obf_0b47ab94fa35044f *__obf_0c3ce44e6d22fabb) *DATTrie {
	__obf_c8fc4f0214f3464f := NewDATTrie()
	__obf_c8fc4f0214f3464f.__obf_633474aeffb3641a[1] = 1

	type __obf_95471ecf79361374 struct {
		__obf_5a4a6635fda80dfd int32
		__obf_2c668bd60383ec67 *__obf_0c3ce44e6d22fabb
	}
	__obf_4eab74c52145cfeb := []__obf_95471ecf79361374{{__obf_5a4a6635fda80dfd: 1, __obf_2c668bd60383ec67: __obf_0b47ab94fa35044f}}

	for len(__obf_4eab74c52145cfeb) > 0 {
		__obf_7052f7c49fc9cf3a := __obf_4eab74c52145cfeb[0]
		__obf_4eab74c52145cfeb = __obf_4eab74c52145cfeb[1:]

		if len(__obf_7052f7c49fc9cf3a.__obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a) == 0 {
			continue
		}
		__obf_c31d1554975bf7bb := make([]rune, 0, len(__obf_7052f7c49fc9cf3a.__obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a))
		for __obf_dfbedc1c3c638189 := range __obf_7052f7c49fc9cf3a.__obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a {
			__obf_c31d1554975bf7bb = append(__obf_c31d1554975bf7bb, __obf_dfbedc1c3c638189)
		}
		slices.Sort(__obf_c31d1554975bf7bb)
		__obf_d50c443884334761 := __obf_c8fc4f0214f3464f.__obf_65d5bd82ec008358(__obf_c31d1554975bf7bb)
		__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0[__obf_7052f7c49fc9cf3a.__obf_5a4a6635fda80dfd] = __obf_d50c443884334761

		for _, __obf_3c428313cda59850 := range __obf_c31d1554975bf7bb {
			__obf_0039188c81b25e77 := __obf_7052f7c49fc9cf3a.__obf_2c668bd60383ec67.__obf_13e9d50886dd9c4a[__obf_3c428313cda59850]
			__obf_7cf8ac4351058da9 := __obf_d50c443884334761 + int32(__obf_3c428313cda59850)
			__obf_c8fc4f0214f3464f.__obf_f16d065d33e36b74(__obf_7cf8ac4351058da9 + 1)
			__obf_c8fc4f0214f3464f.__obf_633474aeffb3641a[__obf_7cf8ac4351058da9] = __obf_7052f7c49fc9cf3a.__obf_5a4a6635fda80dfd
			if __obf_0039188c81b25e77.__obf_218af756657db028 {
				__obf_c8fc4f0214f3464f.__obf_8649c60b00d903ba[__obf_7cf8ac4351058da9] = true
			}
			if int32(len(__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0)) <= __obf_7cf8ac4351058da9 {
				__obf_c8fc4f0214f3464f.__obf_f16d065d33e36b74(__obf_7cf8ac4351058da9 + 1)
			}
			if __obf_c8fc4f0214f3464f.__obf_c1ca2e0801da064b <= __obf_7cf8ac4351058da9 {
				__obf_c8fc4f0214f3464f.__obf_c1ca2e0801da064b = __obf_7cf8ac4351058da9 + 1
			}
			__obf_4eab74c52145cfeb = append(__obf_4eab74c52145cfeb, __obf_95471ecf79361374{__obf_5a4a6635fda80dfd: __obf_7cf8ac4351058da9, __obf_2c668bd60383ec67: __obf_0039188c81b25e77})
		}
	}
	return __obf_c8fc4f0214f3464f
}

func (__obf_c8fc4f0214f3464f *DATTrie) __obf_65d5bd82ec008358(__obf_577d4f2f18829700 []rune) int32 {
	for __obf_b9ffc11f3620a9c0 := int32(1); ; __obf_b9ffc11f3620a9c0++ {
		__obf_cddb52725e2ab156 := false
		for _, __obf_0f875f97be10ee61 := range __obf_577d4f2f18829700 {
			__obf_963c8c08fa2a4b0e := __obf_b9ffc11f3620a9c0 + int32(__obf_0f875f97be10ee61)
			if int32(len(__obf_c8fc4f0214f3464f.__obf_633474aeffb3641a)) <= __obf_963c8c08fa2a4b0e {
				__obf_c8fc4f0214f3464f.__obf_f16d065d33e36b74(__obf_963c8c08fa2a4b0e + 1)
			}
			if __obf_c8fc4f0214f3464f.__obf_633474aeffb3641a[__obf_963c8c08fa2a4b0e] != 0 {
				__obf_cddb52725e2ab156 = true
				break
			}
		}
		if !__obf_cddb52725e2ab156 {
			return __obf_b9ffc11f3620a9c0
		}
	}
}

func (__obf_c8fc4f0214f3464f *DATTrie) __obf_f16d065d33e36b74(__obf_16943e478fb67805 int32) {
	if __obf_16943e478fb67805 < int32(len(__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0)) {
		return
	}
	__obf_7059cfc5fe2d0273 := max(int32(len(__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0))+int32(len(__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0))/2, __obf_16943e478fb67805)
	__obf_c626901a3695661c := make([]int32, __obf_7059cfc5fe2d0273)
	__obf_a6afc69463c6c37a := make([]int32, __obf_7059cfc5fe2d0273)
	__obf_753c04a86e68bc16 := make([]bool, __obf_7059cfc5fe2d0273)
	copy(__obf_c626901a3695661c, __obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0)
	copy(__obf_a6afc69463c6c37a, __obf_c8fc4f0214f3464f.__obf_633474aeffb3641a)
	copy(__obf_753c04a86e68bc16, __obf_c8fc4f0214f3464f.__obf_8649c60b00d903ba)
	__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0 = __obf_c626901a3695661c
	__obf_c8fc4f0214f3464f.__obf_633474aeffb3641a = __obf_a6afc69463c6c37a
	__obf_c8fc4f0214f3464f.__obf_8649c60b00d903ba = __obf_753c04a86e68bc16
}

func (__obf_c8fc4f0214f3464f *DATTrie) __obf_cd770407cd24ff9e(__obf_7359d9e5bc81fd56 int32, __obf_0f875f97be10ee61 rune) int32 {
	if __obf_7359d9e5bc81fd56 <= 0 || __obf_7359d9e5bc81fd56 >= int32(len(__obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0)) || __obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0[__obf_7359d9e5bc81fd56] == 0 {
		return -1
	}
	__obf_963c8c08fa2a4b0e := __obf_c8fc4f0214f3464f.__obf_b9ffc11f3620a9c0[__obf_7359d9e5bc81fd56] + int32(__obf_0f875f97be10ee61)
	if __obf_963c8c08fa2a4b0e > 0 && __obf_963c8c08fa2a4b0e < int32(len(__obf_c8fc4f0214f3464f.__obf_633474aeffb3641a)) && __obf_c8fc4f0214f3464f.__obf_633474aeffb3641a[__obf_963c8c08fa2a4b0e] == __obf_7359d9e5bc81fd56 {
		return __obf_963c8c08fa2a4b0e
	}
	return -1
}

func (__obf_c8fc4f0214f3464f *DATTrie) __obf_218af756657db028(__obf_7359d9e5bc81fd56 int32) bool {
	if __obf_7359d9e5bc81fd56 > 0 && __obf_7359d9e5bc81fd56 < int32(len(__obf_c8fc4f0214f3464f.__obf_8649c60b00d903ba)) {
		return __obf_c8fc4f0214f3464f.__obf_8649c60b00d903ba[__obf_7359d9e5bc81fd56]
	}
	return false
}

func (__obf_c8fc4f0214f3464f *DATTrie) Contains(__obf_324f64750b1cc9de string) bool {
	__obf_de9b32348df338b4 := []rune(__obf_324f64750b1cc9de)
	for __obf_433c212716a376c9 := range __obf_de9b32348df338b4 {
		__obf_2c668bd60383ec67 := int32(1)
		for __obf_a319b29477f20362 := __obf_433c212716a376c9; __obf_a319b29477f20362 < len(__obf_de9b32348df338b4); __obf_a319b29477f20362++ {
			__obf_2c668bd60383ec67 = __obf_c8fc4f0214f3464f.__obf_cd770407cd24ff9e(__obf_2c668bd60383ec67, __obf_de9b32348df338b4[__obf_a319b29477f20362])
			if __obf_2c668bd60383ec67 == -1 {
				break
			}
			if __obf_c8fc4f0214f3464f.__obf_218af756657db028(__obf_2c668bd60383ec67) {
				return true
			}
		}
	}
	return false
}

func (__obf_c8fc4f0214f3464f *DATTrie) __obf_3b60730717aba921(__obf_de9b32348df338b4 []rune) []MatchResult {
	__obf_3ce9664eed99001f := make([]MatchResult, 0, __obf_1fc356c6bc5512cd)
	for __obf_433c212716a376c9 := range __obf_de9b32348df338b4 {
		__obf_2c668bd60383ec67 := int32(1)
		for __obf_a319b29477f20362 := __obf_433c212716a376c9; __obf_a319b29477f20362 < len(__obf_de9b32348df338b4); __obf_a319b29477f20362++ {
			__obf_2c668bd60383ec67 = __obf_c8fc4f0214f3464f.__obf_cd770407cd24ff9e(__obf_2c668bd60383ec67, __obf_de9b32348df338b4[__obf_a319b29477f20362])
			if __obf_2c668bd60383ec67 == -1 {
				break
			}
			if __obf_c8fc4f0214f3464f.__obf_218af756657db028(__obf_2c668bd60383ec67) {
				__obf_3ce9664eed99001f = append(__obf_3ce9664eed99001f, MatchResult{
					Word:  string(__obf_de9b32348df338b4[__obf_433c212716a376c9 : __obf_a319b29477f20362+1]),
					Start: __obf_433c212716a376c9,
					End:   __obf_a319b29477f20362 + 1,
				})
			}
		}
	}
	return __obf_3ce9664eed99001f
}

type WuManberMatcher struct {
	__obf_c8780955e1592e61 []string
	__obf_7a5d4cdb8fa0b164 int
	__obf_1457c180032a4b4f map[string]int32
	__obf_06860f73d726d463 map[string][]int
	__obf_7d9fb5967569c2a6 int
}

func NewWuManberMatcher(__obf_7a5d4cdb8fa0b164 int) *WuManberMatcher {
	return &WuManberMatcher{__obf_c8780955e1592e61: make([]string, 0), __obf_7a5d4cdb8fa0b164: __obf_7a5d4cdb8fa0b164, __obf_1457c180032a4b4f: make(map[string]int32), __obf_06860f73d726d463: make(map[string][]int), __obf_7d9fb5967569c2a6: int(^uint(0) >> 1)}
}

func (__obf_bec81bad6809e949 *WuManberMatcher) AddPatterns(__obf_c8780955e1592e61 []string) {
	__obf_bec81bad6809e949.__obf_c8780955e1592e61 = __obf_c8780955e1592e61
	for _, __obf_2f0ec858b9df4729 := range __obf_c8780955e1592e61 {
		__obf_b2045ddf5ebb22a6 := utf8.RuneCountInString(__obf_2f0ec858b9df4729)
		if __obf_b2045ddf5ebb22a6 > 0 && __obf_b2045ddf5ebb22a6 < __obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 {
			__obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 = __obf_b2045ddf5ebb22a6
		}
	}
}

func (__obf_bec81bad6809e949 *WuManberMatcher) Build() {
	if len(__obf_bec81bad6809e949.__obf_c8780955e1592e61) == 0 {
		return
	}
	if __obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 < __obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 {
		__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 = __obf_bec81bad6809e949.__obf_7d9fb5967569c2a6
	}
	if __obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 <= 0 {
		__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 = 1
	}

	for __obf_433c212716a376c9, __obf_31fcb9c39608f7e0 := range __obf_bec81bad6809e949.__obf_c8780955e1592e61 {
		__obf_de9b32348df338b4 := []rune(__obf_31fcb9c39608f7e0)
		__obf_5215a35953af7237 := len(__obf_de9b32348df338b4)
		for __obf_a319b29477f20362 := 0; __obf_a319b29477f20362 <= __obf_5215a35953af7237-__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164; __obf_a319b29477f20362++ {
			__obf_04bda4a5f8d88d57 := string(__obf_de9b32348df338b4[__obf_a319b29477f20362 : __obf_a319b29477f20362+__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164])
			__obf_55c27a66dff9b0ba := int32(__obf_5215a35953af7237 - __obf_a319b29477f20362 - __obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164)
			if __obf_227c52e00a88383f, __obf_1c49260530c901f8 := __obf_bec81bad6809e949.__obf_1457c180032a4b4f[__obf_04bda4a5f8d88d57]; !__obf_1c49260530c901f8 || __obf_55c27a66dff9b0ba < __obf_227c52e00a88383f {
				__obf_bec81bad6809e949.__obf_1457c180032a4b4f[__obf_04bda4a5f8d88d57] = __obf_55c27a66dff9b0ba
			}
		}

		if __obf_5215a35953af7237 >= __obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 {
			__obf_c43a7c278f7f7385 := string(__obf_de9b32348df338b4[__obf_5215a35953af7237-__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164:])
			__obf_bec81bad6809e949.__obf_06860f73d726d463[__obf_c43a7c278f7f7385] = append(__obf_bec81bad6809e949.__obf_06860f73d726d463[__obf_c43a7c278f7f7385], __obf_433c212716a376c9)
		}
	}
}

func (__obf_bec81bad6809e949 *WuManberMatcher) Match(__obf_324f64750b1cc9de string) bool {
	if len(__obf_bec81bad6809e949.__obf_c8780955e1592e61) == 0 {
		return false
	}
	__obf_de9b32348df338b4 := []rune(__obf_324f64750b1cc9de)
	__obf_63817e43b3f4e2af := len(__obf_de9b32348df338b4)
	if __obf_63817e43b3f4e2af < __obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 {
		return false
	}
	__obf_36d8c69eb86e8573 := __obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 - 1
	for __obf_36d8c69eb86e8573 < __obf_63817e43b3f4e2af {
		__obf_4fe30dd495239355 := max(__obf_36d8c69eb86e8573-__obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164+1, 0)
		__obf_04bda4a5f8d88d57 := string(__obf_de9b32348df338b4[__obf_4fe30dd495239355 : __obf_36d8c69eb86e8573+1])
		__obf_1457c180032a4b4f, __obf_1c49260530c901f8 := __obf_bec81bad6809e949.__obf_1457c180032a4b4f[__obf_04bda4a5f8d88d57]
		if !__obf_1c49260530c901f8 {
			__obf_36d8c69eb86e8573 += (__obf_bec81bad6809e949.__obf_7d9fb5967569c2a6 - __obf_bec81bad6809e949.__obf_7a5d4cdb8fa0b164 + 1)
			continue
		}
		if __obf_1457c180032a4b4f == 0 {
			if __obf_4752a10f89e199c8, __obf_ad62883dfb1e5e56 := __obf_bec81bad6809e949.__obf_06860f73d726d463[__obf_04bda4a5f8d88d57]; __obf_ad62883dfb1e5e56 {
				for _, __obf_b2d5a857c5823bb1 := range __obf_4752a10f89e199c8 {
					__obf_31fcb9c39608f7e0 := __obf_bec81bad6809e949.__obf_c8780955e1592e61[__obf_b2d5a857c5823bb1]
					__obf_1a39f3c21629bfa7 := utf8.RuneCountInString(__obf_31fcb9c39608f7e0)
					__obf_23ba78b2dbade488 := __obf_36d8c69eb86e8573 - __obf_1a39f3c21629bfa7 + 1
					if __obf_23ba78b2dbade488 >= 0 && string(__obf_de9b32348df338b4[__obf_23ba78b2dbade488:__obf_23ba78b2dbade488+__obf_1a39f3c21629bfa7]) == __obf_31fcb9c39608f7e0 {
						return true
					}
				}
			}
			__obf_36d8c69eb86e8573++
		} else {
			__obf_36d8c69eb86e8573 += int(__obf_1457c180032a4b4f)
		}
	}
	return false
}
