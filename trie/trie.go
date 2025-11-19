package __obf_4188a4c7626a6f4f

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
	__obf_604674d6297e17a8 = 1024
	__obf_c10d74f2037d675e = 1_000_000 // 1MB
	__obf_8d27e9ed243c7511 = 3
	__obf_aff7e517f2c8cb7b = 100
	__obf_de70667a352d20a2 = 16
)

// 位图实现
type __obf_33f4a8d75da03445 struct {
	__obf_d98819ccf10e7220 []uint64
}

func __obf_ad495ca0188dd2f7(__obf_328b601304a26990 int) *__obf_33f4a8d75da03445 {
	// 每个 uint64 存储 64 位
	__obf_928650cbdee804f7 := (__obf_328b601304a26990 + 63) / 64
	return &__obf_33f4a8d75da03445{__obf_d98819ccf10e7220: make([]uint64, __obf_928650cbdee804f7)}
}

func (__obf_e35361b5df990a80 *__obf_33f4a8d75da03445) __obf_3a7e060c2bf16877(__obf_8342c8287962d20b int) {
	__obf_e35361b5df990a80.__obf_d98819ccf10e7220[__obf_8342c8287962d20b/64] |= 1 << (__obf_8342c8287962d20b % 64)
}

func (__obf_e35361b5df990a80 *__obf_33f4a8d75da03445) __obf_1ab1abb3518fb614(__obf_8342c8287962d20b int) bool {
	return __obf_e35361b5df990a80.__obf_d98819ccf10e7220[__obf_8342c8287962d20b/64]&(1<<(__obf_8342c8287962d20b%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_a87ce7253a8d3812 *DATTrie
	__obf_755599ad2d3cbbcb *WuManberMatcher
	__obf_b6a79f1207b44bab sync.RWMutex
}

func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) Contains(__obf_f5d6d3fb367460b8 string) bool {
	__obf_a1832a9eb57e2fb1.__obf_b6a79f1207b44bab.Lock()
	defer __obf_a1832a9eb57e2fb1.__obf_b6a79f1207b44bab.Unlock()
	if len(__obf_f5d6d3fb367460b8) > __obf_aff7e517f2c8cb7b {
		if __obf_a1832a9eb57e2fb1.__obf_755599ad2d3cbbcb.Match(__obf_f5d6d3fb367460b8) {
			return true
		} else {
			return false
		}
	}
	return __obf_a1832a9eb57e2fb1.__obf_a87ce7253a8d3812.Contains(__obf_f5d6d3fb367460b8)
}

func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) FindAll(__obf_f5d6d3fb367460b8 string) []MatchResult {
	if __obf_a1832a9eb57e2fb1 == nil || __obf_a1832a9eb57e2fb1.__obf_a87ce7253a8d3812 == nil {
		return nil
	}
	if __obf_f5d6d3fb367460b8 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_f5d6d3fb367460b8) > __obf_c10d74f2037d675e {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_cba27d55cee7dcd5 := __obf_a1832a9eb57e2fb1.__obf_a87ce7253a8d3812.__obf_e9dc0aee1fdb3eb7([]rune(__obf_f5d6d3fb367460b8))
	if len(__obf_cba27d55cee7dcd5) < 2 {
		return __obf_cba27d55cee7dcd5
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_cba27d55cee7dcd5, func(__obf_57472795d1c165a7, __obf_fd35c01bb10e34ec int) bool {
		__obf_6670ea53cce75451 := __obf_cba27d55cee7dcd5[__obf_57472795d1c165a7].End - __obf_cba27d55cee7dcd5[__obf_57472795d1c165a7].Start
		__obf_c3aa750960c83b5c := __obf_cba27d55cee7dcd5[__obf_fd35c01bb10e34ec].End - __obf_cba27d55cee7dcd5[__obf_fd35c01bb10e34ec].Start
		if __obf_6670ea53cce75451 != __obf_c3aa750960c83b5c {
			return __obf_6670ea53cce75451 > __obf_c3aa750960c83b5c
		}
		return __obf_cba27d55cee7dcd5[__obf_57472795d1c165a7].Start < __obf_cba27d55cee7dcd5[__obf_fd35c01bb10e34ec].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_48f54fa33e6a30ff := make([]MatchResult, 0, len(__obf_cba27d55cee7dcd5))
	// occupied := make([]bool, len(runes))
	__obf_64df25488b6efac2 := __obf_ad495ca0188dd2f7(len([]rune(__obf_f5d6d3fb367460b8)))

	for _, __obf_52c8ecb9bbbb48d4 := range __obf_cba27d55cee7dcd5 {
		__obf_acb42801b62dfc6b := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_57472795d1c165a7 := __obf_52c8ecb9bbbb48d4.Start; __obf_57472795d1c165a7 < __obf_52c8ecb9bbbb48d4.End; __obf_57472795d1c165a7++ {
			// if occupied[i] {
			if __obf_64df25488b6efac2.__obf_1ab1abb3518fb614(__obf_57472795d1c165a7) {
				__obf_acb42801b62dfc6b = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_acb42801b62dfc6b {
			__obf_48f54fa33e6a30ff = append(__obf_48f54fa33e6a30ff, __obf_52c8ecb9bbbb48d4)
			// 标记该位置已被占用
			for __obf_57472795d1c165a7 := __obf_52c8ecb9bbbb48d4.Start; __obf_57472795d1c165a7 < __obf_52c8ecb9bbbb48d4.End; __obf_57472795d1c165a7++ {
				// occupied[i] = true
				__obf_64df25488b6efac2.__obf_3a7e060c2bf16877(__obf_57472795d1c165a7)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_48f54fa33e6a30ff, func(__obf_57472795d1c165a7, __obf_fd35c01bb10e34ec int) bool {
		return __obf_48f54fa33e6a30ff[__obf_57472795d1c165a7].Start < __obf_48f54fa33e6a30ff[__obf_fd35c01bb10e34ec].Start
	})

	return __obf_48f54fa33e6a30ff
}

func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) Replace(__obf_f5d6d3fb367460b8 string, __obf_3162c881b12ceb64 rune) string {
	__obf_60ef62f3e70b4690 := __obf_a1832a9eb57e2fb1.FindAll(__obf_f5d6d3fb367460b8)
	if len(__obf_60ef62f3e70b4690) == 0 {
		return __obf_f5d6d3fb367460b8
	}

	__obf_f3bcd7967ea4e8e0 := []rune(__obf_f5d6d3fb367460b8)
	for _, __obf_52c8ecb9bbbb48d4 := range __obf_60ef62f3e70b4690 {
		for __obf_57472795d1c165a7 := __obf_52c8ecb9bbbb48d4.Start; __obf_57472795d1c165a7 < __obf_52c8ecb9bbbb48d4.End; __obf_57472795d1c165a7++ {
			__obf_f3bcd7967ea4e8e0[__obf_57472795d1c165a7] = __obf_3162c881b12ceb64
		}
	}
	return string(__obf_f3bcd7967ea4e8e0)
}

func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) ReplaceWithString(__obf_f5d6d3fb367460b8 string, __obf_3162c881b12ceb64 string) string {
	__obf_60ef62f3e70b4690 := __obf_a1832a9eb57e2fb1.FindAll(__obf_f5d6d3fb367460b8)
	if len(__obf_60ef62f3e70b4690) == 0 {
		return __obf_f5d6d3fb367460b8
	}

	var __obf_cfd4d1617f675da7 strings.Builder
	__obf_f3bcd7967ea4e8e0 := []rune(__obf_f5d6d3fb367460b8)
	__obf_fd65a3bd4992d08b := 0
	for _, __obf_52c8ecb9bbbb48d4 := range __obf_60ef62f3e70b4690 {
		__obf_cfd4d1617f675da7.WriteString(string(__obf_f3bcd7967ea4e8e0[__obf_fd65a3bd4992d08b:__obf_52c8ecb9bbbb48d4.Start]))
		__obf_cfd4d1617f675da7.WriteString(__obf_3162c881b12ceb64)
		__obf_fd65a3bd4992d08b = __obf_52c8ecb9bbbb48d4.End
	}
	__obf_cfd4d1617f675da7.WriteString(string(__obf_f3bcd7967ea4e8e0[__obf_fd65a3bd4992d08b:]))
	return __obf_cfd4d1617f675da7.String()
}

func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) ReplaceWithFunc(__obf_f5d6d3fb367460b8 string, __obf_52ec3d2376410c94 func(__obf_52c8ecb9bbbb48d4 MatchResult) string) string {
	__obf_60ef62f3e70b4690 := __obf_a1832a9eb57e2fb1.FindAll(__obf_f5d6d3fb367460b8)
	if len(__obf_60ef62f3e70b4690) == 0 {
		return __obf_f5d6d3fb367460b8
	}

	var __obf_cfd4d1617f675da7 strings.Builder
	__obf_f3bcd7967ea4e8e0 := []rune(__obf_f5d6d3fb367460b8)
	__obf_fd65a3bd4992d08b := 0
	for _, __obf_52c8ecb9bbbb48d4 := range __obf_60ef62f3e70b4690 {
		__obf_cfd4d1617f675da7.WriteString(string(__obf_f3bcd7967ea4e8e0[__obf_fd65a3bd4992d08b:__obf_52c8ecb9bbbb48d4.Start]))
		__obf_cfd4d1617f675da7.WriteString(__obf_52ec3d2376410c94(__obf_52c8ecb9bbbb48d4))
		__obf_fd65a3bd4992d08b = __obf_52c8ecb9bbbb48d4.End
	}
	__obf_cfd4d1617f675da7.WriteString(string(__obf_f3bcd7967ea4e8e0[__obf_fd65a3bd4992d08b:]))
	return __obf_cfd4d1617f675da7.String()
}

// 添加 Close 方法释放资源
func (__obf_a1832a9eb57e2fb1 *SensitiveFilter) Close() {
	__obf_a1832a9eb57e2fb1.__obf_b6a79f1207b44bab.Lock()
	defer __obf_a1832a9eb57e2fb1.__obf_b6a79f1207b44bab.Unlock()

	__obf_a1832a9eb57e2fb1.__obf_a87ce7253a8d3812 = nil
	__obf_a1832a9eb57e2fb1.__obf_755599ad2d3cbbcb = nil
}

type FilterBuilder struct {
	__obf_5d3007b91c3520e9 map[string]struct{}
	__obf_a900185b45a028e7 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_328b601304a26990 int) FilterOption {
	return func(__obf_e35361b5df990a80 *FilterBuilder) {
		if __obf_328b601304a26990 > 0 {
			__obf_e35361b5df990a80.__obf_a900185b45a028e7 = __obf_328b601304a26990
		}
	}
}

func WithWords(__obf_5d3007b91c3520e9 []string) FilterOption {
	return func(__obf_e35361b5df990a80 *FilterBuilder) {
		for _, __obf_f650372cd9bedf52 := range __obf_5d3007b91c3520e9 {
			if __obf_f650372cd9bedf52 != "" {
				__obf_e35361b5df990a80.__obf_5d3007b91c3520e9[__obf_f650372cd9bedf52] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_8594eeff2e2b1781 ...FilterOption) *FilterBuilder {
	__obf_e35361b5df990a80 := &FilterBuilder{
		__obf_5d3007b91c3520e9: make(map[string]struct{}),
		__obf_a900185b45a028e7: __obf_8d27e9ed243c7511,
	}
	for _, __obf_ae4a0a9555bc11eb := range __obf_8594eeff2e2b1781 {
		__obf_ae4a0a9555bc11eb(__obf_e35361b5df990a80)
	}
	return __obf_e35361b5df990a80
}

func (__obf_e35361b5df990a80 *FilterBuilder) AddWord(__obf_fecdd7917392b2aa string) {
	if __obf_fecdd7917392b2aa != "" {
		__obf_e35361b5df990a80.__obf_5d3007b91c3520e9[__obf_fecdd7917392b2aa] = struct{}{}
	}
}

func (__obf_e35361b5df990a80 *FilterBuilder) AddWords(__obf_5d3007b91c3520e9 []string) {
	for _, __obf_fecdd7917392b2aa := range __obf_5d3007b91c3520e9 {
		__obf_e35361b5df990a80.AddWord(__obf_fecdd7917392b2aa)
	}
}

func (__obf_e35361b5df990a80 *FilterBuilder) LoadFromFile(__obf_9d2142631d99efe4 string) error {
	__obf_59a07a64dd2ffb4b, __obf_7037b84992919dd0 := os.Open(__obf_9d2142631d99efe4)
	if __obf_7037b84992919dd0 != nil {
		return __obf_7037b84992919dd0
	}
	defer __obf_59a07a64dd2ffb4b.Close()

	__obf_bbde2b91f8fdd72d := bufio.NewReader(__obf_59a07a64dd2ffb4b)
	__obf_d5b31d13e98bb108 := []byte{0xEF, 0xBB, 0xBF}
	__obf_60a6503b62399c77, __obf_7037b84992919dd0 := __obf_bbde2b91f8fdd72d.Peek(3)
	if __obf_7037b84992919dd0 == nil && bytes.Equal(__obf_60a6503b62399c77, __obf_d5b31d13e98bb108) {
		_, _ = __obf_bbde2b91f8fdd72d.Discard(3)
	}

	__obf_bb2e3cbde5420448 := bufio.NewScanner(__obf_bbde2b91f8fdd72d)
	for __obf_bb2e3cbde5420448.Scan() {
		__obf_f0f7e7918ab1da27 := strings.TrimSpace(__obf_bb2e3cbde5420448.Text())
		__obf_e35361b5df990a80.AddWord(__obf_f0f7e7918ab1da27)
	}
	return __obf_bb2e3cbde5420448.Err()
}

func (__obf_e35361b5df990a80 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_e35361b5df990a80.__obf_5d3007b91c3520e9) == 0 {
		return &SensitiveFilter{__obf_a87ce7253a8d3812: NewDATTrie(),
			__obf_755599ad2d3cbbcb: NewWuManberMatcher(__obf_e35361b5df990a80.__obf_a900185b45a028e7)}
	}

	__obf_c9fedfa10f1481f7 := &__obf_1ab22d2cfc439857{__obf_b005fda7398fff20: make(map[rune]*__obf_1ab22d2cfc439857)}
	__obf_e563ffc3344a3f2e := make([]string, 0, len(__obf_e35361b5df990a80.__obf_5d3007b91c3520e9))
	for __obf_fecdd7917392b2aa := range __obf_e35361b5df990a80.__obf_5d3007b91c3520e9 {
		__obf_e563ffc3344a3f2e = append(__obf_e563ffc3344a3f2e, __obf_fecdd7917392b2aa)
		__obf_c9fedfa10f1481f7.__obf_879a57c01752fdf1(__obf_fecdd7917392b2aa)
	}

	__obf_a87ce7253a8d3812 := __obf_3cc919c7fa8806a6(__obf_c9fedfa10f1481f7)
	__obf_755599ad2d3cbbcb := NewWuManberMatcher(__obf_e35361b5df990a80.__obf_a900185b45a028e7)
	__obf_755599ad2d3cbbcb.AddPatterns(__obf_e563ffc3344a3f2e)
	__obf_755599ad2d3cbbcb.Build()

	return &SensitiveFilter{
		__obf_a87ce7253a8d3812: __obf_a87ce7253a8d3812,
		__obf_755599ad2d3cbbcb: __obf_755599ad2d3cbbcb,
	}
}

func (__obf_e35361b5df990a80 *FilterBuilder) Clear() {
	__obf_e35361b5df990a80.__obf_5d3007b91c3520e9 = make(map[string]struct{})
}

type __obf_1ab22d2cfc439857 struct {
	__obf_b005fda7398fff20 map[rune]*__obf_1ab22d2cfc439857
	__obf_97dc747465ee2320 bool
}

func (__obf_928650cbdee804f7 *__obf_1ab22d2cfc439857) __obf_879a57c01752fdf1(__obf_fecdd7917392b2aa string) {
	__obf_c1fc30bb1d976143 := __obf_928650cbdee804f7
	for _, __obf_068f6c2cb163a563 := range __obf_fecdd7917392b2aa {
		if _, __obf_ded101c6b05fe20a := __obf_c1fc30bb1d976143.__obf_b005fda7398fff20[__obf_068f6c2cb163a563]; !__obf_ded101c6b05fe20a {
			__obf_c1fc30bb1d976143.__obf_b005fda7398fff20[__obf_068f6c2cb163a563] = &__obf_1ab22d2cfc439857{__obf_b005fda7398fff20: make(map[rune]*__obf_1ab22d2cfc439857)}
		}
		__obf_c1fc30bb1d976143 = __obf_c1fc30bb1d976143.__obf_b005fda7398fff20[__obf_068f6c2cb163a563]
	}
	__obf_c1fc30bb1d976143.__obf_97dc747465ee2320 = true
}

type DATTrie struct {
	__obf_1939f46302fb76b5 []int32
	__obf_9aa803e376883853 []int32
	__obf_bf5474c7202cbca0 []bool
	__obf_328b601304a26990 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_1939f46302fb76b5: make([]int32, __obf_604674d6297e17a8),
		__obf_9aa803e376883853: make([]int32, __obf_604674d6297e17a8),
		__obf_bf5474c7202cbca0: make([]bool, __obf_604674d6297e17a8),
		__obf_328b601304a26990: 1,
	}
}

func __obf_3cc919c7fa8806a6(__obf_39ab10e9d0e8fdb9 *__obf_1ab22d2cfc439857) *DATTrie {
	__obf_588841a7d480ee60 := NewDATTrie()
	__obf_588841a7d480ee60.__obf_9aa803e376883853[1] = 1

	type __obf_11b38cd77901ec0b struct {
		__obf_7313ba37d9b1e1d5 int32
		__obf_c1fc30bb1d976143 *__obf_1ab22d2cfc439857
	}

	__obf_6a48b58590c1b831 := []__obf_11b38cd77901ec0b{{__obf_7313ba37d9b1e1d5: 1, __obf_c1fc30bb1d976143: __obf_39ab10e9d0e8fdb9}}

	for len(__obf_6a48b58590c1b831) > 0 {
		__obf_fbd29cb6ea7302c6 := __obf_6a48b58590c1b831[0]
		__obf_6a48b58590c1b831 = __obf_6a48b58590c1b831[1:]

		if len(__obf_fbd29cb6ea7302c6.__obf_c1fc30bb1d976143.__obf_b005fda7398fff20) == 0 {
			continue
		}

		__obf_1aeb311c18a20bef := make([]rune, 0, len(__obf_fbd29cb6ea7302c6.__obf_c1fc30bb1d976143.__obf_b005fda7398fff20))
		for __obf_068f6c2cb163a563 := range __obf_fbd29cb6ea7302c6.__obf_c1fc30bb1d976143.__obf_b005fda7398fff20 {
			__obf_1aeb311c18a20bef = append(__obf_1aeb311c18a20bef, __obf_068f6c2cb163a563)
		}
		slices.Sort(__obf_1aeb311c18a20bef)

		__obf_a475b79556a3a261 := __obf_588841a7d480ee60.__obf_8fba1ee5bcb76136(__obf_1aeb311c18a20bef)
		__obf_588841a7d480ee60.__obf_1939f46302fb76b5[__obf_fbd29cb6ea7302c6.__obf_7313ba37d9b1e1d5] = __obf_a475b79556a3a261

		for _, __obf_8795b63f91939794 := range __obf_1aeb311c18a20bef {
			__obf_997c0d748a672018 := __obf_fbd29cb6ea7302c6.__obf_c1fc30bb1d976143.__obf_b005fda7398fff20[__obf_8795b63f91939794]
			__obf_9dcfb04b0c30f096 := __obf_a475b79556a3a261 + int32(__obf_8795b63f91939794)
			__obf_588841a7d480ee60.__obf_498f32d3cbeae1ac(__obf_9dcfb04b0c30f096 + 1)

			__obf_588841a7d480ee60.__obf_9aa803e376883853[__obf_9dcfb04b0c30f096] = __obf_fbd29cb6ea7302c6.__obf_7313ba37d9b1e1d5
			if __obf_997c0d748a672018.__obf_97dc747465ee2320 {
				__obf_588841a7d480ee60.__obf_bf5474c7202cbca0[__obf_9dcfb04b0c30f096] = true
			}
			if int32(len(__obf_588841a7d480ee60.__obf_1939f46302fb76b5)) <= __obf_9dcfb04b0c30f096 {
				__obf_588841a7d480ee60.__obf_498f32d3cbeae1ac(__obf_9dcfb04b0c30f096 + 1)
			}
			if __obf_588841a7d480ee60.__obf_328b601304a26990 <= __obf_9dcfb04b0c30f096 {
				__obf_588841a7d480ee60.__obf_328b601304a26990 = __obf_9dcfb04b0c30f096 + 1
			}

			__obf_6a48b58590c1b831 = append(__obf_6a48b58590c1b831, __obf_11b38cd77901ec0b{__obf_7313ba37d9b1e1d5: __obf_9dcfb04b0c30f096, __obf_c1fc30bb1d976143: __obf_997c0d748a672018})
		}
	}
	return __obf_588841a7d480ee60
}

func (__obf_588841a7d480ee60 *DATTrie) __obf_8fba1ee5bcb76136(__obf_df55979db70aab6c []rune) int32 {
	for __obf_1939f46302fb76b5 := int32(1); ; __obf_1939f46302fb76b5++ {
		__obf_72d2f9431b0ca2f2 := false
		for _, __obf_69461aecf81c69ae := range __obf_df55979db70aab6c {
			__obf_a1265721a43a029a := __obf_1939f46302fb76b5 + int32(__obf_69461aecf81c69ae)
			if int32(len(__obf_588841a7d480ee60.__obf_9aa803e376883853)) <= __obf_a1265721a43a029a {
				__obf_588841a7d480ee60.__obf_498f32d3cbeae1ac(__obf_a1265721a43a029a + 1)
			}
			if __obf_588841a7d480ee60.__obf_9aa803e376883853[__obf_a1265721a43a029a] != 0 {
				__obf_72d2f9431b0ca2f2 = true
				break
			}
		}
		if !__obf_72d2f9431b0ca2f2 {
			return __obf_1939f46302fb76b5
		}
	}
}

func (__obf_588841a7d480ee60 *DATTrie) __obf_498f32d3cbeae1ac(__obf_7b0ccbbafe370051 int32) {
	if __obf_7b0ccbbafe370051 < int32(len(__obf_588841a7d480ee60.__obf_1939f46302fb76b5)) {
		return
	}
	__obf_fb690b1a4f982810 := max(int32(len(__obf_588841a7d480ee60.__obf_1939f46302fb76b5))+int32(len(__obf_588841a7d480ee60.__obf_1939f46302fb76b5))/2, __obf_7b0ccbbafe370051)
	__obf_f5c32f14b50beae0 := make([]int32, __obf_fb690b1a4f982810)
	__obf_46fc6feb91aa85cf := make([]int32, __obf_fb690b1a4f982810)
	__obf_8c2f3ebbb5dd12a1 := make([]bool, __obf_fb690b1a4f982810)
	copy(__obf_f5c32f14b50beae0, __obf_588841a7d480ee60.__obf_1939f46302fb76b5)
	copy(__obf_46fc6feb91aa85cf, __obf_588841a7d480ee60.__obf_9aa803e376883853)
	copy(__obf_8c2f3ebbb5dd12a1, __obf_588841a7d480ee60.__obf_bf5474c7202cbca0)
	__obf_588841a7d480ee60.__obf_1939f46302fb76b5 = __obf_f5c32f14b50beae0
	__obf_588841a7d480ee60.__obf_9aa803e376883853 = __obf_46fc6feb91aa85cf
	__obf_588841a7d480ee60.__obf_bf5474c7202cbca0 = __obf_8c2f3ebbb5dd12a1
}

func (__obf_588841a7d480ee60 *DATTrie) __obf_84bf4dfe3bd536ef(__obf_921fcea44a91af54 int32, __obf_69461aecf81c69ae rune) int32 {
	if __obf_921fcea44a91af54 <= 0 || __obf_921fcea44a91af54 >= int32(len(__obf_588841a7d480ee60.__obf_1939f46302fb76b5)) || __obf_588841a7d480ee60.__obf_1939f46302fb76b5[__obf_921fcea44a91af54] == 0 {
		return -1
	}
	__obf_a1265721a43a029a := __obf_588841a7d480ee60.__obf_1939f46302fb76b5[__obf_921fcea44a91af54] + int32(__obf_69461aecf81c69ae)
	if __obf_a1265721a43a029a > 0 && __obf_a1265721a43a029a < int32(len(__obf_588841a7d480ee60.__obf_9aa803e376883853)) && __obf_588841a7d480ee60.__obf_9aa803e376883853[__obf_a1265721a43a029a] == __obf_921fcea44a91af54 {
		return __obf_a1265721a43a029a
	}
	return -1
}

func (__obf_588841a7d480ee60 *DATTrie) __obf_97dc747465ee2320(__obf_921fcea44a91af54 int32) bool {
	if __obf_921fcea44a91af54 > 0 && __obf_921fcea44a91af54 < int32(len(__obf_588841a7d480ee60.__obf_bf5474c7202cbca0)) {
		return __obf_588841a7d480ee60.__obf_bf5474c7202cbca0[__obf_921fcea44a91af54]
	}
	return false
}

func (__obf_588841a7d480ee60 *DATTrie) Contains(__obf_f5d6d3fb367460b8 string) bool {
	__obf_f3bcd7967ea4e8e0 := []rune(__obf_f5d6d3fb367460b8)
	for __obf_57472795d1c165a7 := range __obf_f3bcd7967ea4e8e0 {
		__obf_c1fc30bb1d976143 := int32(1)
		for __obf_fd35c01bb10e34ec := __obf_57472795d1c165a7; __obf_fd35c01bb10e34ec < len(__obf_f3bcd7967ea4e8e0); __obf_fd35c01bb10e34ec++ {
			__obf_c1fc30bb1d976143 = __obf_588841a7d480ee60.__obf_84bf4dfe3bd536ef(__obf_c1fc30bb1d976143, __obf_f3bcd7967ea4e8e0[__obf_fd35c01bb10e34ec])
			if __obf_c1fc30bb1d976143 == -1 {
				break
			}
			if __obf_588841a7d480ee60.__obf_97dc747465ee2320(__obf_c1fc30bb1d976143) {
				return true
			}
		}
	}
	return false
}

func (__obf_588841a7d480ee60 *DATTrie) __obf_e9dc0aee1fdb3eb7(__obf_f3bcd7967ea4e8e0 []rune) []MatchResult {
	__obf_8bda21dbbf2efcfa := make([]MatchResult, 0, __obf_de70667a352d20a2)
	for __obf_57472795d1c165a7 := range __obf_f3bcd7967ea4e8e0 {
		__obf_c1fc30bb1d976143 := int32(1)
		for __obf_fd35c01bb10e34ec := __obf_57472795d1c165a7; __obf_fd35c01bb10e34ec < len(__obf_f3bcd7967ea4e8e0); __obf_fd35c01bb10e34ec++ {
			__obf_c1fc30bb1d976143 = __obf_588841a7d480ee60.__obf_84bf4dfe3bd536ef(__obf_c1fc30bb1d976143, __obf_f3bcd7967ea4e8e0[__obf_fd35c01bb10e34ec])
			if __obf_c1fc30bb1d976143 == -1 {
				break
			}
			if __obf_588841a7d480ee60.__obf_97dc747465ee2320(__obf_c1fc30bb1d976143) {
				__obf_8bda21dbbf2efcfa = append(__obf_8bda21dbbf2efcfa, MatchResult{
					Word:  string(__obf_f3bcd7967ea4e8e0[__obf_57472795d1c165a7 : __obf_fd35c01bb10e34ec+1]),
					Start: __obf_57472795d1c165a7,
					End:   __obf_fd35c01bb10e34ec + 1,
				})
			}
		}
	}
	return __obf_8bda21dbbf2efcfa
}

type WuManberMatcher struct {
	__obf_09a265f1ea152be0 []string
	__obf_a900185b45a028e7 int
	__obf_5ac34306c38351d6 map[string]int32
	__obf_6c44f042452a792a map[string][]int
	__obf_b46479b3a91290f5 int
}

func NewWuManberMatcher(__obf_a900185b45a028e7 int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_09a265f1ea152be0: make([]string, 0),
		__obf_a900185b45a028e7: __obf_a900185b45a028e7,
		__obf_5ac34306c38351d6: make(map[string]int32),
		__obf_6c44f042452a792a: make(map[string][]int),
		__obf_b46479b3a91290f5: int(^uint(0) >> 1),
	}
}

func (__obf_d664433d6f39a479 *WuManberMatcher) AddPatterns(__obf_09a265f1ea152be0 []string) {
	__obf_d664433d6f39a479.__obf_09a265f1ea152be0 = __obf_09a265f1ea152be0
	for _, __obf_bab8fc057ee93eae := range __obf_09a265f1ea152be0 {
		__obf_e2fa707d379bc328 := utf8.RuneCountInString(__obf_bab8fc057ee93eae)
		if __obf_e2fa707d379bc328 > 0 && __obf_e2fa707d379bc328 < __obf_d664433d6f39a479.__obf_b46479b3a91290f5 {
			__obf_d664433d6f39a479.__obf_b46479b3a91290f5 = __obf_e2fa707d379bc328
		}
	}
}

func (__obf_d664433d6f39a479 *WuManberMatcher) Build() {
	if len(__obf_d664433d6f39a479.__obf_09a265f1ea152be0) == 0 {
		return
	}
	if __obf_d664433d6f39a479.__obf_b46479b3a91290f5 < __obf_d664433d6f39a479.__obf_a900185b45a028e7 {
		__obf_d664433d6f39a479.__obf_a900185b45a028e7 = __obf_d664433d6f39a479.__obf_b46479b3a91290f5
	}
	if __obf_d664433d6f39a479.__obf_a900185b45a028e7 <= 0 {
		__obf_d664433d6f39a479.__obf_a900185b45a028e7 = 1
	}

	for __obf_57472795d1c165a7, __obf_f944645a57d63797 := range __obf_d664433d6f39a479.__obf_09a265f1ea152be0 {
		__obf_f3bcd7967ea4e8e0 := []rune(__obf_f944645a57d63797)
		__obf_1c5a1fed417bfd21 := len(__obf_f3bcd7967ea4e8e0)
		for __obf_fd35c01bb10e34ec := 0; __obf_fd35c01bb10e34ec <= __obf_1c5a1fed417bfd21-__obf_d664433d6f39a479.__obf_a900185b45a028e7; __obf_fd35c01bb10e34ec++ {
			__obf_de82c371655483fa := string(__obf_f3bcd7967ea4e8e0[__obf_fd35c01bb10e34ec : __obf_fd35c01bb10e34ec+__obf_d664433d6f39a479.__obf_a900185b45a028e7])
			__obf_fa949f182bd955f0 := int32(__obf_1c5a1fed417bfd21 - __obf_fd35c01bb10e34ec - __obf_d664433d6f39a479.__obf_a900185b45a028e7)
			if __obf_fb41b71718ac1955, __obf_edf0026ee3f41c4c := __obf_d664433d6f39a479.__obf_5ac34306c38351d6[__obf_de82c371655483fa]; !__obf_edf0026ee3f41c4c || __obf_fa949f182bd955f0 < __obf_fb41b71718ac1955 {
				__obf_d664433d6f39a479.__obf_5ac34306c38351d6[__obf_de82c371655483fa] = __obf_fa949f182bd955f0
			}
		}

		if __obf_1c5a1fed417bfd21 >= __obf_d664433d6f39a479.__obf_a900185b45a028e7 {
			__obf_d1c31cd3cde662e9 := string(__obf_f3bcd7967ea4e8e0[__obf_1c5a1fed417bfd21-__obf_d664433d6f39a479.__obf_a900185b45a028e7:])
			__obf_d664433d6f39a479.__obf_6c44f042452a792a[__obf_d1c31cd3cde662e9] = append(__obf_d664433d6f39a479.__obf_6c44f042452a792a[__obf_d1c31cd3cde662e9], __obf_57472795d1c165a7)
		}
	}
}

func (__obf_d664433d6f39a479 *WuManberMatcher) Match(__obf_f5d6d3fb367460b8 string) bool {
	if len(__obf_d664433d6f39a479.__obf_09a265f1ea152be0) == 0 {
		return false
	}
	__obf_f3bcd7967ea4e8e0 := []rune(__obf_f5d6d3fb367460b8)
	__obf_d240e211e1449169 := len(__obf_f3bcd7967ea4e8e0)
	if __obf_d240e211e1449169 < __obf_d664433d6f39a479.__obf_b46479b3a91290f5 {
		return false
	}
	__obf_8342c8287962d20b := __obf_d664433d6f39a479.__obf_b46479b3a91290f5 - 1
	for __obf_8342c8287962d20b < __obf_d240e211e1449169 {
		__obf_336f1dd1543c0db0 := max(__obf_8342c8287962d20b-__obf_d664433d6f39a479.__obf_a900185b45a028e7+1, 0)
		__obf_de82c371655483fa := string(__obf_f3bcd7967ea4e8e0[__obf_336f1dd1543c0db0 : __obf_8342c8287962d20b+1])
		__obf_5ac34306c38351d6, __obf_edf0026ee3f41c4c := __obf_d664433d6f39a479.__obf_5ac34306c38351d6[__obf_de82c371655483fa]
		if !__obf_edf0026ee3f41c4c {
			__obf_8342c8287962d20b += (__obf_d664433d6f39a479.__obf_b46479b3a91290f5 - __obf_d664433d6f39a479.__obf_a900185b45a028e7 + 1)
			continue
		}
		if __obf_5ac34306c38351d6 == 0 {
			if __obf_64cfd50df3e581a4, __obf_ded101c6b05fe20a := __obf_d664433d6f39a479.__obf_6c44f042452a792a[__obf_de82c371655483fa]; __obf_ded101c6b05fe20a {
				for _, __obf_2570a910ab930792 := range __obf_64cfd50df3e581a4 {
					__obf_f944645a57d63797 := __obf_d664433d6f39a479.__obf_09a265f1ea152be0[__obf_2570a910ab930792]
					__obf_7a997ed2433935ab := utf8.RuneCountInString(__obf_f944645a57d63797)
					__obf_27c7eb56568a9652 := __obf_8342c8287962d20b - __obf_7a997ed2433935ab + 1
					if __obf_27c7eb56568a9652 >= 0 && string(__obf_f3bcd7967ea4e8e0[__obf_27c7eb56568a9652:__obf_27c7eb56568a9652+__obf_7a997ed2433935ab]) == __obf_f944645a57d63797 {
						return true
					}
				}
			}
			__obf_8342c8287962d20b++
		} else {
			__obf_8342c8287962d20b += int(__obf_5ac34306c38351d6)
		}
	}
	return false
}
