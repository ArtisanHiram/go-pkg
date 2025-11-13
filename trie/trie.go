package __obf_dfedaafca2addd44

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
	__obf_3f7cf5325ee67888 = 1024
	__obf_074a9473d87fc36f = 1_000_000 // 1MB
	__obf_42d8b1255425c639 = 3
	__obf_5e574386609778e7 = 100
	__obf_cb3dff0298e0fcb5 = 16
)

// 位图实现
type __obf_34b25fa2095c632d struct {
	__obf_866158ad3ea34bf8 []uint64
}

func __obf_35b1272e36c1ea81(__obf_9160d7d9f06c7a58 int) *__obf_34b25fa2095c632d {
	// 每个 uint64 存储 64 位
	__obf_571556fdb6b3a2ce := (__obf_9160d7d9f06c7a58 + 63) / 64
	return &__obf_34b25fa2095c632d{__obf_866158ad3ea34bf8: make([]uint64, __obf_571556fdb6b3a2ce)}
}

func (__obf_dfb471ec48f46dc3 *__obf_34b25fa2095c632d) __obf_0bbc9f1bf2f310a9(__obf_06bdb1c34eb8fed6 int) {
	__obf_dfb471ec48f46dc3.__obf_866158ad3ea34bf8[__obf_06bdb1c34eb8fed6/64] |= 1 << (__obf_06bdb1c34eb8fed6 % 64)
}

func (__obf_dfb471ec48f46dc3 *__obf_34b25fa2095c632d) __obf_64da578b6e1e10b1(__obf_06bdb1c34eb8fed6 int) bool {
	return __obf_dfb471ec48f46dc3.__obf_866158ad3ea34bf8[__obf_06bdb1c34eb8fed6/64]&(1<<(__obf_06bdb1c34eb8fed6%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_3e3c2daf05f3a57d *DATTrie
	__obf_4c5890dfb1a09af2 *WuManberMatcher
	__obf_973cd598df03dde2 sync.RWMutex
}

func (__obf_b49a6016b6bb025c *SensitiveFilter) Contains(__obf_767392140f1b4453 string) bool {
	__obf_b49a6016b6bb025c.__obf_973cd598df03dde2.Lock()
	defer __obf_b49a6016b6bb025c.__obf_973cd598df03dde2.Unlock()
	if len(__obf_767392140f1b4453) > __obf_5e574386609778e7 {
		if __obf_b49a6016b6bb025c.__obf_4c5890dfb1a09af2.Match(__obf_767392140f1b4453) {
			return true
		} else {
			return false
		}
	}
	return __obf_b49a6016b6bb025c.__obf_3e3c2daf05f3a57d.Contains(__obf_767392140f1b4453)
}

func (__obf_b49a6016b6bb025c *SensitiveFilter) FindAll(__obf_767392140f1b4453 string) []MatchResult {
	if __obf_b49a6016b6bb025c == nil || __obf_b49a6016b6bb025c.__obf_3e3c2daf05f3a57d == nil {
		return nil
	}
	if __obf_767392140f1b4453 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_767392140f1b4453) > __obf_074a9473d87fc36f {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_c53acac2d823dfa9 := __obf_b49a6016b6bb025c.__obf_3e3c2daf05f3a57d.__obf_fcf727ba03b3ff5e([]rune(__obf_767392140f1b4453))
	if len(__obf_c53acac2d823dfa9) < 2 {
		return __obf_c53acac2d823dfa9
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_c53acac2d823dfa9, func(__obf_e7d912988b0b4a9a, __obf_c6addcdffab1583a int) bool {
		__obf_086a895d59cc5c61 := __obf_c53acac2d823dfa9[__obf_e7d912988b0b4a9a].End - __obf_c53acac2d823dfa9[__obf_e7d912988b0b4a9a].Start
		__obf_88f5939e8329cf25 := __obf_c53acac2d823dfa9[__obf_c6addcdffab1583a].End - __obf_c53acac2d823dfa9[__obf_c6addcdffab1583a].Start
		if __obf_086a895d59cc5c61 != __obf_88f5939e8329cf25 {
			return __obf_086a895d59cc5c61 > __obf_88f5939e8329cf25
		}
		return __obf_c53acac2d823dfa9[__obf_e7d912988b0b4a9a].Start < __obf_c53acac2d823dfa9[__obf_c6addcdffab1583a].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_aa030c66f4055235 := make([]MatchResult, 0, len(__obf_c53acac2d823dfa9))
	// occupied := make([]bool, len(runes))
	__obf_4617cf8c1a6c73cf := __obf_35b1272e36c1ea81(len([]rune(__obf_767392140f1b4453)))

	for _, __obf_f83aae4c702f6bea := range __obf_c53acac2d823dfa9 {
		__obf_29e43ec5c5f5039c := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_e7d912988b0b4a9a := __obf_f83aae4c702f6bea.Start; __obf_e7d912988b0b4a9a < __obf_f83aae4c702f6bea.End; __obf_e7d912988b0b4a9a++ {
			// if occupied[i] {
			if __obf_4617cf8c1a6c73cf.__obf_64da578b6e1e10b1(__obf_e7d912988b0b4a9a) {
				__obf_29e43ec5c5f5039c = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_29e43ec5c5f5039c {
			__obf_aa030c66f4055235 = append(__obf_aa030c66f4055235, __obf_f83aae4c702f6bea)
			// 标记该位置已被占用
			for __obf_e7d912988b0b4a9a := __obf_f83aae4c702f6bea.Start; __obf_e7d912988b0b4a9a < __obf_f83aae4c702f6bea.End; __obf_e7d912988b0b4a9a++ {
				// occupied[i] = true
				__obf_4617cf8c1a6c73cf.__obf_0bbc9f1bf2f310a9(__obf_e7d912988b0b4a9a)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_aa030c66f4055235, func(__obf_e7d912988b0b4a9a, __obf_c6addcdffab1583a int) bool {
		return __obf_aa030c66f4055235[__obf_e7d912988b0b4a9a].Start < __obf_aa030c66f4055235[__obf_c6addcdffab1583a].Start
	})

	return __obf_aa030c66f4055235
}

func (__obf_b49a6016b6bb025c *SensitiveFilter) Replace(__obf_767392140f1b4453 string, __obf_b78ddff14d78057e rune) string {
	__obf_d2bf87f0981efafa := __obf_b49a6016b6bb025c.FindAll(__obf_767392140f1b4453)
	if len(__obf_d2bf87f0981efafa) == 0 {
		return __obf_767392140f1b4453
	}

	__obf_a66a98b5236084b4 := []rune(__obf_767392140f1b4453)
	for _, __obf_f83aae4c702f6bea := range __obf_d2bf87f0981efafa {
		for __obf_e7d912988b0b4a9a := __obf_f83aae4c702f6bea.Start; __obf_e7d912988b0b4a9a < __obf_f83aae4c702f6bea.End; __obf_e7d912988b0b4a9a++ {
			__obf_a66a98b5236084b4[__obf_e7d912988b0b4a9a] = __obf_b78ddff14d78057e
		}
	}
	return string(__obf_a66a98b5236084b4)
}

func (__obf_b49a6016b6bb025c *SensitiveFilter) ReplaceWithString(__obf_767392140f1b4453 string, __obf_b78ddff14d78057e string) string {
	__obf_d2bf87f0981efafa := __obf_b49a6016b6bb025c.FindAll(__obf_767392140f1b4453)
	if len(__obf_d2bf87f0981efafa) == 0 {
		return __obf_767392140f1b4453
	}

	var __obf_131a1b75ba9bf550 strings.Builder
	__obf_a66a98b5236084b4 := []rune(__obf_767392140f1b4453)
	__obf_e99f3427fd03afe2 := 0
	for _, __obf_f83aae4c702f6bea := range __obf_d2bf87f0981efafa {
		__obf_131a1b75ba9bf550.WriteString(string(__obf_a66a98b5236084b4[__obf_e99f3427fd03afe2:__obf_f83aae4c702f6bea.Start]))
		__obf_131a1b75ba9bf550.WriteString(__obf_b78ddff14d78057e)
		__obf_e99f3427fd03afe2 = __obf_f83aae4c702f6bea.End
	}
	__obf_131a1b75ba9bf550.WriteString(string(__obf_a66a98b5236084b4[__obf_e99f3427fd03afe2:]))
	return __obf_131a1b75ba9bf550.String()
}

func (__obf_b49a6016b6bb025c *SensitiveFilter) ReplaceWithFunc(__obf_767392140f1b4453 string, __obf_1546fc7d62476ff9 func(__obf_f83aae4c702f6bea MatchResult) string) string {
	__obf_d2bf87f0981efafa := __obf_b49a6016b6bb025c.FindAll(__obf_767392140f1b4453)
	if len(__obf_d2bf87f0981efafa) == 0 {
		return __obf_767392140f1b4453
	}

	var __obf_131a1b75ba9bf550 strings.Builder
	__obf_a66a98b5236084b4 := []rune(__obf_767392140f1b4453)
	__obf_e99f3427fd03afe2 := 0
	for _, __obf_f83aae4c702f6bea := range __obf_d2bf87f0981efafa {
		__obf_131a1b75ba9bf550.WriteString(string(__obf_a66a98b5236084b4[__obf_e99f3427fd03afe2:__obf_f83aae4c702f6bea.Start]))
		__obf_131a1b75ba9bf550.WriteString(__obf_1546fc7d62476ff9(__obf_f83aae4c702f6bea))
		__obf_e99f3427fd03afe2 = __obf_f83aae4c702f6bea.End
	}
	__obf_131a1b75ba9bf550.WriteString(string(__obf_a66a98b5236084b4[__obf_e99f3427fd03afe2:]))
	return __obf_131a1b75ba9bf550.String()
}

// 添加 Close 方法释放资源
func (__obf_b49a6016b6bb025c *SensitiveFilter) Close() {
	__obf_b49a6016b6bb025c.__obf_973cd598df03dde2.Lock()
	defer __obf_b49a6016b6bb025c.__obf_973cd598df03dde2.Unlock()

	__obf_b49a6016b6bb025c.__obf_3e3c2daf05f3a57d = nil
	__obf_b49a6016b6bb025c.__obf_4c5890dfb1a09af2 = nil
}

type FilterBuilder struct {
	__obf_56a4544a77721441 map[string]struct{}
	__obf_ed22f1c9b9abbbb8 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_9160d7d9f06c7a58 int) FilterOption {
	return func(__obf_dfb471ec48f46dc3 *FilterBuilder) {
		if __obf_9160d7d9f06c7a58 > 0 {
			__obf_dfb471ec48f46dc3.__obf_ed22f1c9b9abbbb8 = __obf_9160d7d9f06c7a58
		}
	}
}

func WithWords(__obf_56a4544a77721441 []string) FilterOption {
	return func(__obf_dfb471ec48f46dc3 *FilterBuilder) {
		for _, __obf_ec6123e898a8a713 := range __obf_56a4544a77721441 {
			if __obf_ec6123e898a8a713 != "" {
				__obf_dfb471ec48f46dc3.__obf_56a4544a77721441[__obf_ec6123e898a8a713] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_12baa83bedc79064 ...FilterOption) *FilterBuilder {
	__obf_dfb471ec48f46dc3 := &FilterBuilder{
		__obf_56a4544a77721441: make(map[string]struct{}),
		__obf_ed22f1c9b9abbbb8: __obf_42d8b1255425c639,
	}
	for _, __obf_6ce112e00856d459 := range __obf_12baa83bedc79064 {
		__obf_6ce112e00856d459(__obf_dfb471ec48f46dc3)
	}
	return __obf_dfb471ec48f46dc3
}

func (__obf_dfb471ec48f46dc3 *FilterBuilder) AddWord(__obf_2622dddf6ae3389e string) {
	if __obf_2622dddf6ae3389e != "" {
		__obf_dfb471ec48f46dc3.__obf_56a4544a77721441[__obf_2622dddf6ae3389e] = struct{}{}
	}
}

func (__obf_dfb471ec48f46dc3 *FilterBuilder) AddWords(__obf_56a4544a77721441 []string) {
	for _, __obf_2622dddf6ae3389e := range __obf_56a4544a77721441 {
		__obf_dfb471ec48f46dc3.AddWord(__obf_2622dddf6ae3389e)
	}
}

func (__obf_dfb471ec48f46dc3 *FilterBuilder) LoadFromFile(__obf_883c96fc0e4a1553 string) error {
	__obf_c6508a3346a73976, __obf_832582b4007d4090 := os.Open(__obf_883c96fc0e4a1553)
	if __obf_832582b4007d4090 != nil {
		return __obf_832582b4007d4090
	}
	defer __obf_c6508a3346a73976.Close()

	__obf_f9f894168e030dd6 := bufio.NewReader(__obf_c6508a3346a73976)
	__obf_074c2f36a99c05b2 := []byte{0xEF, 0xBB, 0xBF}
	__obf_13aa7b075bfb9099, __obf_832582b4007d4090 := __obf_f9f894168e030dd6.Peek(3)
	if __obf_832582b4007d4090 == nil && bytes.Equal(__obf_13aa7b075bfb9099, __obf_074c2f36a99c05b2) {
		_, _ = __obf_f9f894168e030dd6.Discard(3)
	}

	__obf_787eba4bbc6663b4 := bufio.NewScanner(__obf_f9f894168e030dd6)
	for __obf_787eba4bbc6663b4.Scan() {
		__obf_0ad149087c852e04 := strings.TrimSpace(__obf_787eba4bbc6663b4.Text())
		__obf_dfb471ec48f46dc3.AddWord(__obf_0ad149087c852e04)
	}
	return __obf_787eba4bbc6663b4.Err()
}

func (__obf_dfb471ec48f46dc3 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_dfb471ec48f46dc3.__obf_56a4544a77721441) == 0 {
		return &SensitiveFilter{__obf_3e3c2daf05f3a57d: NewDATTrie(),
			__obf_4c5890dfb1a09af2: NewWuManberMatcher(__obf_dfb471ec48f46dc3.__obf_ed22f1c9b9abbbb8)}
	}

	__obf_1e446aecf06c1383 := &__obf_104d0560515103b3{__obf_f51cf073181a5daf: make(map[rune]*__obf_104d0560515103b3)}
	__obf_d11b8c2c570b8581 := make([]string, 0, len(__obf_dfb471ec48f46dc3.__obf_56a4544a77721441))
	for __obf_2622dddf6ae3389e := range __obf_dfb471ec48f46dc3.__obf_56a4544a77721441 {
		__obf_d11b8c2c570b8581 = append(__obf_d11b8c2c570b8581, __obf_2622dddf6ae3389e)
		__obf_1e446aecf06c1383.__obf_46b33777ff4c4c4d(__obf_2622dddf6ae3389e)
	}

	__obf_3e3c2daf05f3a57d := __obf_83f660300d1a31f2(__obf_1e446aecf06c1383)
	__obf_4c5890dfb1a09af2 := NewWuManberMatcher(__obf_dfb471ec48f46dc3.__obf_ed22f1c9b9abbbb8)
	__obf_4c5890dfb1a09af2.AddPatterns(__obf_d11b8c2c570b8581)
	__obf_4c5890dfb1a09af2.Build()

	return &SensitiveFilter{
		__obf_3e3c2daf05f3a57d: __obf_3e3c2daf05f3a57d,
		__obf_4c5890dfb1a09af2: __obf_4c5890dfb1a09af2,
	}
}

func (__obf_dfb471ec48f46dc3 *FilterBuilder) Clear() {
	__obf_dfb471ec48f46dc3.__obf_56a4544a77721441 = make(map[string]struct{})
}

type __obf_104d0560515103b3 struct {
	__obf_f51cf073181a5daf map[rune]*__obf_104d0560515103b3
	__obf_d515bbe986da72f4 bool
}

func (__obf_571556fdb6b3a2ce *__obf_104d0560515103b3) __obf_46b33777ff4c4c4d(__obf_2622dddf6ae3389e string) {
	__obf_5adc728d3af4417d := __obf_571556fdb6b3a2ce
	for _, __obf_ba802d835cf7abb0 := range __obf_2622dddf6ae3389e {
		if _, __obf_10bb3dbdc20ef15d := __obf_5adc728d3af4417d.__obf_f51cf073181a5daf[__obf_ba802d835cf7abb0]; !__obf_10bb3dbdc20ef15d {
			__obf_5adc728d3af4417d.__obf_f51cf073181a5daf[__obf_ba802d835cf7abb0] = &__obf_104d0560515103b3{__obf_f51cf073181a5daf: make(map[rune]*__obf_104d0560515103b3)}
		}
		__obf_5adc728d3af4417d = __obf_5adc728d3af4417d.__obf_f51cf073181a5daf[__obf_ba802d835cf7abb0]
	}
	__obf_5adc728d3af4417d.__obf_d515bbe986da72f4 = true
}

type DATTrie struct {
	__obf_574b8dc602e165c9 []int32
	__obf_738ea909577361b1 []int32
	__obf_51f56b042fcc73da []bool
	__obf_9160d7d9f06c7a58 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_574b8dc602e165c9: make([]int32, __obf_3f7cf5325ee67888),
		__obf_738ea909577361b1: make([]int32, __obf_3f7cf5325ee67888),
		__obf_51f56b042fcc73da: make([]bool, __obf_3f7cf5325ee67888),
		__obf_9160d7d9f06c7a58: 1,
	}
}

func __obf_83f660300d1a31f2(__obf_b67735c6745b7749 *__obf_104d0560515103b3) *DATTrie {
	__obf_b5935f8d180cf1cc := NewDATTrie()
	__obf_b5935f8d180cf1cc.__obf_738ea909577361b1[1] = 1

	type __obf_41747c19ec8d0197 struct {
		__obf_5e1c2bdce5a33b5a int32
		__obf_5adc728d3af4417d *__obf_104d0560515103b3
	}

	__obf_45b6b4fed11c7e20 := []__obf_41747c19ec8d0197{{__obf_5e1c2bdce5a33b5a: 1, __obf_5adc728d3af4417d: __obf_b67735c6745b7749}}

	for len(__obf_45b6b4fed11c7e20) > 0 {
		__obf_89530d1c3dd30506 := __obf_45b6b4fed11c7e20[0]
		__obf_45b6b4fed11c7e20 = __obf_45b6b4fed11c7e20[1:]

		if len(__obf_89530d1c3dd30506.__obf_5adc728d3af4417d.__obf_f51cf073181a5daf) == 0 {
			continue
		}

		__obf_59f0fec6e978bf2d := make([]rune, 0, len(__obf_89530d1c3dd30506.__obf_5adc728d3af4417d.__obf_f51cf073181a5daf))
		for __obf_ba802d835cf7abb0 := range __obf_89530d1c3dd30506.__obf_5adc728d3af4417d.__obf_f51cf073181a5daf {
			__obf_59f0fec6e978bf2d = append(__obf_59f0fec6e978bf2d, __obf_ba802d835cf7abb0)
		}
		slices.Sort(__obf_59f0fec6e978bf2d)

		__obf_4b2834cd2bcbe2f0 := __obf_b5935f8d180cf1cc.__obf_ce2846d849db4e44(__obf_59f0fec6e978bf2d)
		__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9[__obf_89530d1c3dd30506.__obf_5e1c2bdce5a33b5a] = __obf_4b2834cd2bcbe2f0

		for _, __obf_a7d82632cddade0d := range __obf_59f0fec6e978bf2d {
			__obf_955d0c532401c893 := __obf_89530d1c3dd30506.__obf_5adc728d3af4417d.__obf_f51cf073181a5daf[__obf_a7d82632cddade0d]
			__obf_df39dfe0d597c9de := __obf_4b2834cd2bcbe2f0 + int32(__obf_a7d82632cddade0d)
			__obf_b5935f8d180cf1cc.__obf_19c5cf06921aed44(__obf_df39dfe0d597c9de + 1)

			__obf_b5935f8d180cf1cc.__obf_738ea909577361b1[__obf_df39dfe0d597c9de] = __obf_89530d1c3dd30506.__obf_5e1c2bdce5a33b5a
			if __obf_955d0c532401c893.__obf_d515bbe986da72f4 {
				__obf_b5935f8d180cf1cc.__obf_51f56b042fcc73da[__obf_df39dfe0d597c9de] = true
			}
			if int32(len(__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9)) <= __obf_df39dfe0d597c9de {
				__obf_b5935f8d180cf1cc.__obf_19c5cf06921aed44(__obf_df39dfe0d597c9de + 1)
			}
			if __obf_b5935f8d180cf1cc.__obf_9160d7d9f06c7a58 <= __obf_df39dfe0d597c9de {
				__obf_b5935f8d180cf1cc.__obf_9160d7d9f06c7a58 = __obf_df39dfe0d597c9de + 1
			}

			__obf_45b6b4fed11c7e20 = append(__obf_45b6b4fed11c7e20, __obf_41747c19ec8d0197{__obf_5e1c2bdce5a33b5a: __obf_df39dfe0d597c9de, __obf_5adc728d3af4417d: __obf_955d0c532401c893})
		}
	}
	return __obf_b5935f8d180cf1cc
}

func (__obf_b5935f8d180cf1cc *DATTrie) __obf_ce2846d849db4e44(__obf_9fffb84f08585cc9 []rune) int32 {
	for __obf_574b8dc602e165c9 := int32(1); ; __obf_574b8dc602e165c9++ {
		__obf_08f48a110d5bcf68 := false
		for _, __obf_40010c475017c5db := range __obf_9fffb84f08585cc9 {
			__obf_522eac98cc1f4817 := __obf_574b8dc602e165c9 + int32(__obf_40010c475017c5db)
			if int32(len(__obf_b5935f8d180cf1cc.__obf_738ea909577361b1)) <= __obf_522eac98cc1f4817 {
				__obf_b5935f8d180cf1cc.__obf_19c5cf06921aed44(__obf_522eac98cc1f4817 + 1)
			}
			if __obf_b5935f8d180cf1cc.__obf_738ea909577361b1[__obf_522eac98cc1f4817] != 0 {
				__obf_08f48a110d5bcf68 = true
				break
			}
		}
		if !__obf_08f48a110d5bcf68 {
			return __obf_574b8dc602e165c9
		}
	}
}

func (__obf_b5935f8d180cf1cc *DATTrie) __obf_19c5cf06921aed44(__obf_14c2644a25cf6079 int32) {
	if __obf_14c2644a25cf6079 < int32(len(__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9)) {
		return
	}
	__obf_3d4a71af522ab37f := max(int32(len(__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9))+int32(len(__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9))/2, __obf_14c2644a25cf6079)
	__obf_ef7c565ca6e263af := make([]int32, __obf_3d4a71af522ab37f)
	__obf_5183acfaeea8a99f := make([]int32, __obf_3d4a71af522ab37f)
	__obf_0a396dc9b5212d96 := make([]bool, __obf_3d4a71af522ab37f)
	copy(__obf_ef7c565ca6e263af, __obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9)
	copy(__obf_5183acfaeea8a99f, __obf_b5935f8d180cf1cc.__obf_738ea909577361b1)
	copy(__obf_0a396dc9b5212d96, __obf_b5935f8d180cf1cc.__obf_51f56b042fcc73da)
	__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9 = __obf_ef7c565ca6e263af
	__obf_b5935f8d180cf1cc.__obf_738ea909577361b1 = __obf_5183acfaeea8a99f
	__obf_b5935f8d180cf1cc.__obf_51f56b042fcc73da = __obf_0a396dc9b5212d96
}

func (__obf_b5935f8d180cf1cc *DATTrie) __obf_56397963b6c5277d(__obf_d729a75feed6d2f8 int32, __obf_40010c475017c5db rune) int32 {
	if __obf_d729a75feed6d2f8 <= 0 || __obf_d729a75feed6d2f8 >= int32(len(__obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9)) || __obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9[__obf_d729a75feed6d2f8] == 0 {
		return -1
	}
	__obf_522eac98cc1f4817 := __obf_b5935f8d180cf1cc.__obf_574b8dc602e165c9[__obf_d729a75feed6d2f8] + int32(__obf_40010c475017c5db)
	if __obf_522eac98cc1f4817 > 0 && __obf_522eac98cc1f4817 < int32(len(__obf_b5935f8d180cf1cc.__obf_738ea909577361b1)) && __obf_b5935f8d180cf1cc.__obf_738ea909577361b1[__obf_522eac98cc1f4817] == __obf_d729a75feed6d2f8 {
		return __obf_522eac98cc1f4817
	}
	return -1
}

func (__obf_b5935f8d180cf1cc *DATTrie) __obf_d515bbe986da72f4(__obf_d729a75feed6d2f8 int32) bool {
	if __obf_d729a75feed6d2f8 > 0 && __obf_d729a75feed6d2f8 < int32(len(__obf_b5935f8d180cf1cc.__obf_51f56b042fcc73da)) {
		return __obf_b5935f8d180cf1cc.__obf_51f56b042fcc73da[__obf_d729a75feed6d2f8]
	}
	return false
}

func (__obf_b5935f8d180cf1cc *DATTrie) Contains(__obf_767392140f1b4453 string) bool {
	__obf_a66a98b5236084b4 := []rune(__obf_767392140f1b4453)
	for __obf_e7d912988b0b4a9a := range __obf_a66a98b5236084b4 {
		__obf_5adc728d3af4417d := int32(1)
		for __obf_c6addcdffab1583a := __obf_e7d912988b0b4a9a; __obf_c6addcdffab1583a < len(__obf_a66a98b5236084b4); __obf_c6addcdffab1583a++ {
			__obf_5adc728d3af4417d = __obf_b5935f8d180cf1cc.__obf_56397963b6c5277d(__obf_5adc728d3af4417d, __obf_a66a98b5236084b4[__obf_c6addcdffab1583a])
			if __obf_5adc728d3af4417d == -1 {
				break
			}
			if __obf_b5935f8d180cf1cc.__obf_d515bbe986da72f4(__obf_5adc728d3af4417d) {
				return true
			}
		}
	}
	return false
}

func (__obf_b5935f8d180cf1cc *DATTrie) __obf_fcf727ba03b3ff5e(__obf_a66a98b5236084b4 []rune) []MatchResult {
	__obf_d2a678c94b4bb112 := make([]MatchResult, 0, __obf_cb3dff0298e0fcb5)
	for __obf_e7d912988b0b4a9a := range __obf_a66a98b5236084b4 {
		__obf_5adc728d3af4417d := int32(1)
		for __obf_c6addcdffab1583a := __obf_e7d912988b0b4a9a; __obf_c6addcdffab1583a < len(__obf_a66a98b5236084b4); __obf_c6addcdffab1583a++ {
			__obf_5adc728d3af4417d = __obf_b5935f8d180cf1cc.__obf_56397963b6c5277d(__obf_5adc728d3af4417d, __obf_a66a98b5236084b4[__obf_c6addcdffab1583a])
			if __obf_5adc728d3af4417d == -1 {
				break
			}
			if __obf_b5935f8d180cf1cc.__obf_d515bbe986da72f4(__obf_5adc728d3af4417d) {
				__obf_d2a678c94b4bb112 = append(__obf_d2a678c94b4bb112, MatchResult{
					Word:  string(__obf_a66a98b5236084b4[__obf_e7d912988b0b4a9a : __obf_c6addcdffab1583a+1]),
					Start: __obf_e7d912988b0b4a9a,
					End:   __obf_c6addcdffab1583a + 1,
				})
			}
		}
	}
	return __obf_d2a678c94b4bb112
}

type WuManberMatcher struct {
	__obf_d171dd8534e22ec7 []string
	__obf_ed22f1c9b9abbbb8 int
	__obf_a463adc655f72ee8 map[string]int32
	__obf_1306cc0e891565aa map[string][]int
	__obf_f3dc7584c9bc3f6b int
}

func NewWuManberMatcher(__obf_ed22f1c9b9abbbb8 int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_d171dd8534e22ec7: make([]string, 0),
		__obf_ed22f1c9b9abbbb8: __obf_ed22f1c9b9abbbb8,
		__obf_a463adc655f72ee8: make(map[string]int32),
		__obf_1306cc0e891565aa: make(map[string][]int),
		__obf_f3dc7584c9bc3f6b: int(^uint(0) >> 1),
	}
}

func (__obf_e6192ff172ca4c36 *WuManberMatcher) AddPatterns(__obf_d171dd8534e22ec7 []string) {
	__obf_e6192ff172ca4c36.__obf_d171dd8534e22ec7 = __obf_d171dd8534e22ec7
	for _, __obf_712d85a2d90d86f4 := range __obf_d171dd8534e22ec7 {
		__obf_1cb6afff2b714745 := utf8.RuneCountInString(__obf_712d85a2d90d86f4)
		if __obf_1cb6afff2b714745 > 0 && __obf_1cb6afff2b714745 < __obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b {
			__obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b = __obf_1cb6afff2b714745
		}
	}
}

func (__obf_e6192ff172ca4c36 *WuManberMatcher) Build() {
	if len(__obf_e6192ff172ca4c36.__obf_d171dd8534e22ec7) == 0 {
		return
	}
	if __obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b < __obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 {
		__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 = __obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b
	}
	if __obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 <= 0 {
		__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 = 1
	}

	for __obf_e7d912988b0b4a9a, __obf_fdc7fee1489132fe := range __obf_e6192ff172ca4c36.__obf_d171dd8534e22ec7 {
		__obf_a66a98b5236084b4 := []rune(__obf_fdc7fee1489132fe)
		__obf_8291b5d725d7e9f3 := len(__obf_a66a98b5236084b4)
		for __obf_c6addcdffab1583a := 0; __obf_c6addcdffab1583a <= __obf_8291b5d725d7e9f3-__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8; __obf_c6addcdffab1583a++ {
			__obf_612e538f7a96f492 := string(__obf_a66a98b5236084b4[__obf_c6addcdffab1583a : __obf_c6addcdffab1583a+__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8])
			__obf_c73fe452939e7bd9 := int32(__obf_8291b5d725d7e9f3 - __obf_c6addcdffab1583a - __obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8)
			if __obf_0b5aea00116cce1b, __obf_f274e8f24782b5a4 := __obf_e6192ff172ca4c36.__obf_a463adc655f72ee8[__obf_612e538f7a96f492]; !__obf_f274e8f24782b5a4 || __obf_c73fe452939e7bd9 < __obf_0b5aea00116cce1b {
				__obf_e6192ff172ca4c36.__obf_a463adc655f72ee8[__obf_612e538f7a96f492] = __obf_c73fe452939e7bd9
			}
		}

		if __obf_8291b5d725d7e9f3 >= __obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 {
			__obf_8bc69cb7d6d38099 := string(__obf_a66a98b5236084b4[__obf_8291b5d725d7e9f3-__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8:])
			__obf_e6192ff172ca4c36.__obf_1306cc0e891565aa[__obf_8bc69cb7d6d38099] = append(__obf_e6192ff172ca4c36.__obf_1306cc0e891565aa[__obf_8bc69cb7d6d38099], __obf_e7d912988b0b4a9a)
		}
	}
}

func (__obf_e6192ff172ca4c36 *WuManberMatcher) Match(__obf_767392140f1b4453 string) bool {
	if len(__obf_e6192ff172ca4c36.__obf_d171dd8534e22ec7) == 0 {
		return false
	}
	__obf_a66a98b5236084b4 := []rune(__obf_767392140f1b4453)
	__obf_4566e081c67c4767 := len(__obf_a66a98b5236084b4)
	if __obf_4566e081c67c4767 < __obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b {
		return false
	}
	__obf_06bdb1c34eb8fed6 := __obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b - 1
	for __obf_06bdb1c34eb8fed6 < __obf_4566e081c67c4767 {
		__obf_f7d932dfc9bb380f := max(__obf_06bdb1c34eb8fed6-__obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8+1, 0)
		__obf_612e538f7a96f492 := string(__obf_a66a98b5236084b4[__obf_f7d932dfc9bb380f : __obf_06bdb1c34eb8fed6+1])
		__obf_a463adc655f72ee8, __obf_f274e8f24782b5a4 := __obf_e6192ff172ca4c36.__obf_a463adc655f72ee8[__obf_612e538f7a96f492]
		if !__obf_f274e8f24782b5a4 {
			__obf_06bdb1c34eb8fed6 += (__obf_e6192ff172ca4c36.__obf_f3dc7584c9bc3f6b - __obf_e6192ff172ca4c36.__obf_ed22f1c9b9abbbb8 + 1)
			continue
		}
		if __obf_a463adc655f72ee8 == 0 {
			if __obf_8f8abfcc5bda8bc8, __obf_10bb3dbdc20ef15d := __obf_e6192ff172ca4c36.__obf_1306cc0e891565aa[__obf_612e538f7a96f492]; __obf_10bb3dbdc20ef15d {
				for _, __obf_f1ffb32dd68e4282 := range __obf_8f8abfcc5bda8bc8 {
					__obf_fdc7fee1489132fe := __obf_e6192ff172ca4c36.__obf_d171dd8534e22ec7[__obf_f1ffb32dd68e4282]
					__obf_2fffa4611c2c74d5 := utf8.RuneCountInString(__obf_fdc7fee1489132fe)
					__obf_91d2b2a337268de0 := __obf_06bdb1c34eb8fed6 - __obf_2fffa4611c2c74d5 + 1
					if __obf_91d2b2a337268de0 >= 0 && string(__obf_a66a98b5236084b4[__obf_91d2b2a337268de0:__obf_91d2b2a337268de0+__obf_2fffa4611c2c74d5]) == __obf_fdc7fee1489132fe {
						return true
					}
				}
			}
			__obf_06bdb1c34eb8fed6++
		} else {
			__obf_06bdb1c34eb8fed6 += int(__obf_a463adc655f72ee8)
		}
	}
	return false
}
