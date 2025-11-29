package __obf_e583e2630f62a7bc

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
	__obf_ea666664aed9d0d7 = 1024
	__obf_ca10e198f26ff511 = 1_000_000
	__obf_a336d1e49f3afa9a = // 1MB
	3
	__obf_1b6de80e1d075a56 = 100
	__obf_b5584c9e1f4e82bf = 16
)

// 位图实现
type __obf_7c4cd9c117552a2a struct {
	__obf_eb9a4fbd2c30fd31 []uint64
}

func __obf_aaa9eb14da5295df(__obf_914ad28fcf67ea88 int) *__obf_7c4cd9c117552a2a {
	__obf_533959f4b300364a := // 每个 uint64 存储 64 位
		(__obf_914ad28fcf67ea88 + 63) / 64
	return &__obf_7c4cd9c117552a2a{__obf_eb9a4fbd2c30fd31: make([]uint64, __obf_533959f4b300364a)}
}

func (__obf_49851572551c0ee5 *__obf_7c4cd9c117552a2a) __obf_e3160267df481474(__obf_e821b24c447a8d6a int) {
	__obf_49851572551c0ee5.__obf_eb9a4fbd2c30fd31[__obf_e821b24c447a8d6a/64] |= 1 << (__obf_e821b24c447a8d6a % 64)
}

func (__obf_49851572551c0ee5 *__obf_7c4cd9c117552a2a) __obf_ef4fb51149c2589b(__obf_e821b24c447a8d6a int) bool {
	return __obf_49851572551c0ee5.__obf_eb9a4fbd2c30fd31[__obf_e821b24c447a8d6a/64]&(1<<(__obf_e821b24c447a8d6a%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_7527f4f9a2ce135b *DATTrie
	__obf_a0c045b4ddfac43d *WuManberMatcher
	__obf_e06313f010e75b7b sync.RWMutex
}

func (__obf_68366884350bf834 *SensitiveFilter) Contains(__obf_2f9bfc5f616e4c2e string) bool {
	__obf_68366884350bf834.__obf_e06313f010e75b7b.
		Lock()
	defer __obf_68366884350bf834.__obf_e06313f010e75b7b.Unlock()
	if len(__obf_2f9bfc5f616e4c2e) > __obf_1b6de80e1d075a56 {
		if __obf_68366884350bf834.__obf_a0c045b4ddfac43d.Match(__obf_2f9bfc5f616e4c2e) {
			return true
		} else {
			return false
		}
	}
	return __obf_68366884350bf834.__obf_7527f4f9a2ce135b.Contains(__obf_2f9bfc5f616e4c2e)
}

func (__obf_68366884350bf834 *SensitiveFilter) FindAll(__obf_2f9bfc5f616e4c2e string) []MatchResult {
	if __obf_68366884350bf834 == nil || __obf_68366884350bf834.__obf_7527f4f9a2ce135b == nil {
		return nil
	}
	if __obf_2f9bfc5f616e4c2e == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_2f9bfc5f616e4c2e) > __obf_ca10e198f26ff511 {
		return nil // 或返回错误
	}
	__obf_1adc2820070ea2bc := // 1. 找出所有可能重叠的匹配项
		__obf_68366884350bf834.__obf_7527f4f9a2ce135b.__obf_a3e898872b98d664([]rune(__obf_2f9bfc5f616e4c2e))
	if len(__obf_1adc2820070ea2bc) < 2 {
		return __obf_1adc2820070ea2bc
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_1adc2820070ea2bc, func(__obf_09c4704e98f82128, __obf_8c2ffe965b3e6421 int) bool {
		__obf_df42ecc7c0a92c48 := __obf_1adc2820070ea2bc[__obf_09c4704e98f82128].End - __obf_1adc2820070ea2bc[__obf_09c4704e98f82128].Start
		__obf_07cb3860ba6b96b1 := __obf_1adc2820070ea2bc[__obf_8c2ffe965b3e6421].End - __obf_1adc2820070ea2bc[__obf_8c2ffe965b3e6421].Start
		if __obf_df42ecc7c0a92c48 != __obf_07cb3860ba6b96b1 {
			return __obf_df42ecc7c0a92c48 > __obf_07cb3860ba6b96b1
		}
		return __obf_1adc2820070ea2bc[__obf_09c4704e98f82128].Start < __obf_1adc2820070ea2bc[__obf_8c2ffe965b3e6421].Start
	})
	__obf_dff8bac0ab4de84d := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_1adc2820070ea2bc))
	__obf_268cba24e31abaeb := // occupied := make([]bool, len(runes))
		__obf_aaa9eb14da5295df(len([]rune(__obf_2f9bfc5f616e4c2e)))

	for _, __obf_58d9d03c190988b1 := range __obf_1adc2820070ea2bc {
		__obf_32b1197e0d5af7c5 := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_09c4704e98f82128 := __obf_58d9d03c190988b1.Start; __obf_09c4704e98f82128 < __obf_58d9d03c190988b1.End; __obf_09c4704e98f82128++ {
			// if occupied[i] {
			if __obf_268cba24e31abaeb.__obf_ef4fb51149c2589b(__obf_09c4704e98f82128) {
				__obf_32b1197e0d5af7c5 = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_32b1197e0d5af7c5 {
			__obf_dff8bac0ab4de84d = append(__obf_dff8bac0ab4de84d,
				// 标记该位置已被占用
				__obf_58d9d03c190988b1)

			for __obf_09c4704e98f82128 := __obf_58d9d03c190988b1.Start; __obf_09c4704e98f82128 < __obf_58d9d03c190988b1.End; __obf_09c4704e98f82128++ {
				__obf_268cba24e31abaeb.
					// occupied[i] = true
					__obf_e3160267df481474(__obf_09c4704e98f82128)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_dff8bac0ab4de84d, func(__obf_09c4704e98f82128, __obf_8c2ffe965b3e6421 int) bool {
		return __obf_dff8bac0ab4de84d[__obf_09c4704e98f82128].Start < __obf_dff8bac0ab4de84d[__obf_8c2ffe965b3e6421].Start
	})

	return __obf_dff8bac0ab4de84d
}

func (__obf_68366884350bf834 *SensitiveFilter) Replace(__obf_2f9bfc5f616e4c2e string, __obf_0f3a84651c2341ba rune) string {
	__obf_4611708a6d6f0a23 := __obf_68366884350bf834.FindAll(__obf_2f9bfc5f616e4c2e)
	if len(__obf_4611708a6d6f0a23) == 0 {
		return __obf_2f9bfc5f616e4c2e
	}
	__obf_c6569bea94c0c483 := []rune(__obf_2f9bfc5f616e4c2e)
	for _, __obf_58d9d03c190988b1 := range __obf_4611708a6d6f0a23 {
		for __obf_09c4704e98f82128 := __obf_58d9d03c190988b1.Start; __obf_09c4704e98f82128 < __obf_58d9d03c190988b1.End; __obf_09c4704e98f82128++ {
			__obf_c6569bea94c0c483[__obf_09c4704e98f82128] = __obf_0f3a84651c2341ba
		}
	}
	return string(__obf_c6569bea94c0c483)
}

func (__obf_68366884350bf834 *SensitiveFilter) ReplaceWithString(__obf_2f9bfc5f616e4c2e string, __obf_0f3a84651c2341ba string) string {
	__obf_4611708a6d6f0a23 := __obf_68366884350bf834.FindAll(__obf_2f9bfc5f616e4c2e)
	if len(__obf_4611708a6d6f0a23) == 0 {
		return __obf_2f9bfc5f616e4c2e
	}

	var __obf_31cd6f1ad50b1472 strings.Builder
	__obf_c6569bea94c0c483 := []rune(__obf_2f9bfc5f616e4c2e)
	__obf_b92619b87bbd26e5 := 0
	for _, __obf_58d9d03c190988b1 := range __obf_4611708a6d6f0a23 {
		__obf_31cd6f1ad50b1472.
			WriteString(string(__obf_c6569bea94c0c483[__obf_b92619b87bbd26e5:__obf_58d9d03c190988b1.Start]))
		__obf_31cd6f1ad50b1472.
			WriteString(__obf_0f3a84651c2341ba)
		__obf_b92619b87bbd26e5 = __obf_58d9d03c190988b1.End
	}
	__obf_31cd6f1ad50b1472.
		WriteString(string(__obf_c6569bea94c0c483[__obf_b92619b87bbd26e5:]))
	return __obf_31cd6f1ad50b1472.String()
}

func (__obf_68366884350bf834 *SensitiveFilter) ReplaceWithFunc(__obf_2f9bfc5f616e4c2e string, __obf_f0145afe0d39988a func(__obf_58d9d03c190988b1 MatchResult) string) string {
	__obf_4611708a6d6f0a23 := __obf_68366884350bf834.FindAll(__obf_2f9bfc5f616e4c2e)
	if len(__obf_4611708a6d6f0a23) == 0 {
		return __obf_2f9bfc5f616e4c2e
	}

	var __obf_31cd6f1ad50b1472 strings.Builder
	__obf_c6569bea94c0c483 := []rune(__obf_2f9bfc5f616e4c2e)
	__obf_b92619b87bbd26e5 := 0
	for _, __obf_58d9d03c190988b1 := range __obf_4611708a6d6f0a23 {
		__obf_31cd6f1ad50b1472.
			WriteString(string(__obf_c6569bea94c0c483[__obf_b92619b87bbd26e5:__obf_58d9d03c190988b1.Start]))
		__obf_31cd6f1ad50b1472.
			WriteString(__obf_f0145afe0d39988a(__obf_58d9d03c190988b1))
		__obf_b92619b87bbd26e5 = __obf_58d9d03c190988b1.End
	}
	__obf_31cd6f1ad50b1472.
		WriteString(string(__obf_c6569bea94c0c483[__obf_b92619b87bbd26e5:]))
	return __obf_31cd6f1ad50b1472.String()
}

// 添加 Close 方法释放资源
func (__obf_68366884350bf834 *SensitiveFilter) Close() {
	__obf_68366884350bf834.__obf_e06313f010e75b7b.
		Lock()
	defer __obf_68366884350bf834.__obf_e06313f010e75b7b.Unlock()
	__obf_68366884350bf834.__obf_7527f4f9a2ce135b = nil
	__obf_68366884350bf834.__obf_a0c045b4ddfac43d = nil
}

type FilterBuilder struct {
	__obf_48321fda5b414a8a map[string]struct{}
	__obf_5e405c455150a363 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_914ad28fcf67ea88 int) FilterOption {
	return func(__obf_49851572551c0ee5 *FilterBuilder) {
		if __obf_914ad28fcf67ea88 > 0 {
			__obf_49851572551c0ee5.__obf_5e405c455150a363 = __obf_914ad28fcf67ea88
		}
	}
}

func WithWords(__obf_48321fda5b414a8a []string) FilterOption {
	return func(__obf_49851572551c0ee5 *FilterBuilder) {
		for _, __obf_9f840f939469d0dd := range __obf_48321fda5b414a8a {
			if __obf_9f840f939469d0dd != "" {
				__obf_49851572551c0ee5.__obf_48321fda5b414a8a[__obf_9f840f939469d0dd] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_c63dfbefaeba8018 ...FilterOption) *FilterBuilder {
	__obf_49851572551c0ee5 := &FilterBuilder{__obf_48321fda5b414a8a: make(map[string]struct{}), __obf_5e405c455150a363: __obf_a336d1e49f3afa9a}
	for _, __obf_6c2eba25ad4ffe58 := range __obf_c63dfbefaeba8018 {
		__obf_6c2eba25ad4ffe58(__obf_49851572551c0ee5)
	}
	return __obf_49851572551c0ee5
}

func (__obf_49851572551c0ee5 *FilterBuilder) AddWord(__obf_d68115d58b63f3f5 string) {
	if __obf_d68115d58b63f3f5 != "" {
		__obf_49851572551c0ee5.__obf_48321fda5b414a8a[__obf_d68115d58b63f3f5] = struct{}{}
	}
}

func (__obf_49851572551c0ee5 *FilterBuilder) AddWords(__obf_48321fda5b414a8a []string) {
	for _, __obf_d68115d58b63f3f5 := range __obf_48321fda5b414a8a {
		__obf_49851572551c0ee5.
			AddWord(__obf_d68115d58b63f3f5)
	}
}

func (__obf_49851572551c0ee5 *FilterBuilder) LoadFromFile(__obf_63d3f4956ee55ec7 string) error {
	__obf_0b187175dadb84d1, __obf_afafc45c4edea8e0 := os.Open(__obf_63d3f4956ee55ec7)
	if __obf_afafc45c4edea8e0 != nil {
		return __obf_afafc45c4edea8e0
	}
	defer __obf_0b187175dadb84d1.Close()
	__obf_7ff1b65897e4f757 := bufio.NewReader(__obf_0b187175dadb84d1)
	__obf_0eb572c122f2eddc := []byte{0xEF, 0xBB, 0xBF}
	__obf_b020035bcdfe70a6, __obf_afafc45c4edea8e0 := __obf_7ff1b65897e4f757.Peek(3)
	if __obf_afafc45c4edea8e0 == nil && bytes.Equal(__obf_b020035bcdfe70a6, __obf_0eb572c122f2eddc) {
		_, _ = __obf_7ff1b65897e4f757.Discard(3)
	}
	__obf_878d013614e9b457 := bufio.NewScanner(__obf_7ff1b65897e4f757)
	for __obf_878d013614e9b457.Scan() {
		__obf_c22bc91659961b80 := strings.TrimSpace(__obf_878d013614e9b457.Text())
		__obf_49851572551c0ee5.
			AddWord(__obf_c22bc91659961b80)
	}
	return __obf_878d013614e9b457.Err()
}

func (__obf_49851572551c0ee5 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_49851572551c0ee5.__obf_48321fda5b414a8a) == 0 {
		return &SensitiveFilter{__obf_7527f4f9a2ce135b: NewDATTrie(), __obf_a0c045b4ddfac43d: NewWuManberMatcher(__obf_49851572551c0ee5.__obf_5e405c455150a363)}
	}
	__obf_785fefbaebb15956 := &__obf_9bffe14594ccb425{__obf_fcc5fb5d4c9e35ba: make(map[rune]*__obf_9bffe14594ccb425)}
	__obf_e40c7e0ba76db350 := make([]string, 0, len(__obf_49851572551c0ee5.__obf_48321fda5b414a8a))
	for __obf_d68115d58b63f3f5 := range __obf_49851572551c0ee5.__obf_48321fda5b414a8a {
		__obf_e40c7e0ba76db350 = append(__obf_e40c7e0ba76db350, __obf_d68115d58b63f3f5)
		__obf_785fefbaebb15956.__obf_7e3d3393aac1f60a(__obf_d68115d58b63f3f5)
	}
	__obf_7527f4f9a2ce135b := __obf_af34cbbbd81a522b(__obf_785fefbaebb15956)
	__obf_a0c045b4ddfac43d := NewWuManberMatcher(__obf_49851572551c0ee5.__obf_5e405c455150a363)
	__obf_a0c045b4ddfac43d.
		AddPatterns(__obf_e40c7e0ba76db350)
	__obf_a0c045b4ddfac43d.
		Build()

	return &SensitiveFilter{__obf_7527f4f9a2ce135b: __obf_7527f4f9a2ce135b, __obf_a0c045b4ddfac43d: __obf_a0c045b4ddfac43d}
}

func (__obf_49851572551c0ee5 *FilterBuilder) Clear() {
	__obf_49851572551c0ee5.__obf_48321fda5b414a8a = make(map[string]struct{})
}

type __obf_9bffe14594ccb425 struct {
	__obf_fcc5fb5d4c9e35ba map[rune]*__obf_9bffe14594ccb425
	__obf_f7281b41d9f98ac5 bool
}

func (__obf_533959f4b300364a *__obf_9bffe14594ccb425) __obf_7e3d3393aac1f60a(__obf_d68115d58b63f3f5 string) {
	__obf_6b3585156999efd8 := __obf_533959f4b300364a
	for _, __obf_bb5c99c8ba8010c0 := range __obf_d68115d58b63f3f5 {
		if _, __obf_97c4a66ba64905ca := __obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba[__obf_bb5c99c8ba8010c0]; !__obf_97c4a66ba64905ca {
			__obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba[__obf_bb5c99c8ba8010c0] = &__obf_9bffe14594ccb425{__obf_fcc5fb5d4c9e35ba: make(map[rune]*__obf_9bffe14594ccb425)}
		}
		__obf_6b3585156999efd8 = __obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba[__obf_bb5c99c8ba8010c0]
	}
	__obf_6b3585156999efd8.__obf_f7281b41d9f98ac5 = true
}

type DATTrie struct {
	__obf_21e4a0c991098223 []int32
	__obf_bdde4c8ff394b211 []int32
	__obf_90e9f9cd57864c5a []bool
	__obf_914ad28fcf67ea88 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_21e4a0c991098223: make([]int32, __obf_ea666664aed9d0d7), __obf_bdde4c8ff394b211: make([]int32, __obf_ea666664aed9d0d7), __obf_90e9f9cd57864c5a: make([]bool, __obf_ea666664aed9d0d7), __obf_914ad28fcf67ea88: 1}
}

func __obf_af34cbbbd81a522b(__obf_44298884ebf65902 *__obf_9bffe14594ccb425) *DATTrie {
	__obf_c858f7714608ec5c := NewDATTrie()
	__obf_c858f7714608ec5c.__obf_bdde4c8ff394b211[1] = 1

	type __obf_7875de6b999ca66f struct {
		__obf_278b706c401b5ddb int32
		__obf_6b3585156999efd8 *__obf_9bffe14594ccb425
	}
	__obf_d0aa8f017b2b3732 := []__obf_7875de6b999ca66f{{__obf_278b706c401b5ddb: 1, __obf_6b3585156999efd8: __obf_44298884ebf65902}}

	for len(__obf_d0aa8f017b2b3732) > 0 {
		__obf_33f0987f7d28cef0 := __obf_d0aa8f017b2b3732[0]
		__obf_d0aa8f017b2b3732 = __obf_d0aa8f017b2b3732[1:]

		if len(__obf_33f0987f7d28cef0.__obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba) == 0 {
			continue
		}
		__obf_b9a23a6874a06e96 := make([]rune, 0, len(__obf_33f0987f7d28cef0.__obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba))
		for __obf_bb5c99c8ba8010c0 := range __obf_33f0987f7d28cef0.__obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba {
			__obf_b9a23a6874a06e96 = append(__obf_b9a23a6874a06e96, __obf_bb5c99c8ba8010c0)
		}
		slices.Sort(__obf_b9a23a6874a06e96)
		__obf_73d1975284885f8e := __obf_c858f7714608ec5c.__obf_12a5075cdf951790(__obf_b9a23a6874a06e96)
		__obf_c858f7714608ec5c.__obf_21e4a0c991098223[__obf_33f0987f7d28cef0.__obf_278b706c401b5ddb] = __obf_73d1975284885f8e

		for _, __obf_8bbd017e129d91a4 := range __obf_b9a23a6874a06e96 {
			__obf_959d45d8a85493a7 := __obf_33f0987f7d28cef0.__obf_6b3585156999efd8.__obf_fcc5fb5d4c9e35ba[__obf_8bbd017e129d91a4]
			__obf_22c0ffbc9cbe43ff := __obf_73d1975284885f8e + int32(__obf_8bbd017e129d91a4)
			__obf_c858f7714608ec5c.__obf_17efc3b49e94c9d5(__obf_22c0ffbc9cbe43ff + 1)
			__obf_c858f7714608ec5c.__obf_bdde4c8ff394b211[__obf_22c0ffbc9cbe43ff] = __obf_33f0987f7d28cef0.__obf_278b706c401b5ddb
			if __obf_959d45d8a85493a7.__obf_f7281b41d9f98ac5 {
				__obf_c858f7714608ec5c.__obf_90e9f9cd57864c5a[__obf_22c0ffbc9cbe43ff] = true
			}
			if int32(len(__obf_c858f7714608ec5c.__obf_21e4a0c991098223)) <= __obf_22c0ffbc9cbe43ff {
				__obf_c858f7714608ec5c.__obf_17efc3b49e94c9d5(__obf_22c0ffbc9cbe43ff + 1)
			}
			if __obf_c858f7714608ec5c.__obf_914ad28fcf67ea88 <= __obf_22c0ffbc9cbe43ff {
				__obf_c858f7714608ec5c.__obf_914ad28fcf67ea88 = __obf_22c0ffbc9cbe43ff + 1
			}
			__obf_d0aa8f017b2b3732 = append(__obf_d0aa8f017b2b3732, __obf_7875de6b999ca66f{__obf_278b706c401b5ddb: __obf_22c0ffbc9cbe43ff, __obf_6b3585156999efd8: __obf_959d45d8a85493a7})
		}
	}
	return __obf_c858f7714608ec5c
}

func (__obf_c858f7714608ec5c *DATTrie) __obf_12a5075cdf951790(__obf_9dd002dfb645415b []rune) int32 {
	for __obf_21e4a0c991098223 := int32(1); ; __obf_21e4a0c991098223++ {
		__obf_e061af02650a336b := false
		for _, __obf_e4aafdbcb1b13739 := range __obf_9dd002dfb645415b {
			__obf_9a7f80ff75ced90f := __obf_21e4a0c991098223 + int32(__obf_e4aafdbcb1b13739)
			if int32(len(__obf_c858f7714608ec5c.__obf_bdde4c8ff394b211)) <= __obf_9a7f80ff75ced90f {
				__obf_c858f7714608ec5c.__obf_17efc3b49e94c9d5(__obf_9a7f80ff75ced90f + 1)
			}
			if __obf_c858f7714608ec5c.__obf_bdde4c8ff394b211[__obf_9a7f80ff75ced90f] != 0 {
				__obf_e061af02650a336b = true
				break
			}
		}
		if !__obf_e061af02650a336b {
			return __obf_21e4a0c991098223
		}
	}
}

func (__obf_c858f7714608ec5c *DATTrie) __obf_17efc3b49e94c9d5(__obf_4d12a09f61c200fb int32) {
	if __obf_4d12a09f61c200fb < int32(len(__obf_c858f7714608ec5c.__obf_21e4a0c991098223)) {
		return
	}
	__obf_b061b612945ad714 := max(int32(len(__obf_c858f7714608ec5c.__obf_21e4a0c991098223))+int32(len(__obf_c858f7714608ec5c.__obf_21e4a0c991098223))/2, __obf_4d12a09f61c200fb)
	__obf_90dc784192571ec1 := make([]int32, __obf_b061b612945ad714)
	__obf_d484d5f5b7eaf144 := make([]int32, __obf_b061b612945ad714)
	__obf_c67d5df962c2e81a := make([]bool, __obf_b061b612945ad714)
	copy(__obf_90dc784192571ec1, __obf_c858f7714608ec5c.__obf_21e4a0c991098223)
	copy(__obf_d484d5f5b7eaf144, __obf_c858f7714608ec5c.__obf_bdde4c8ff394b211)
	copy(__obf_c67d5df962c2e81a, __obf_c858f7714608ec5c.__obf_90e9f9cd57864c5a)
	__obf_c858f7714608ec5c.__obf_21e4a0c991098223 = __obf_90dc784192571ec1
	__obf_c858f7714608ec5c.__obf_bdde4c8ff394b211 = __obf_d484d5f5b7eaf144
	__obf_c858f7714608ec5c.__obf_90e9f9cd57864c5a = __obf_c67d5df962c2e81a
}

func (__obf_c858f7714608ec5c *DATTrie) __obf_7426f753a0b18cca(__obf_97fab2dfcc31983d int32, __obf_e4aafdbcb1b13739 rune) int32 {
	if __obf_97fab2dfcc31983d <= 0 || __obf_97fab2dfcc31983d >= int32(len(__obf_c858f7714608ec5c.__obf_21e4a0c991098223)) || __obf_c858f7714608ec5c.__obf_21e4a0c991098223[__obf_97fab2dfcc31983d] == 0 {
		return -1
	}
	__obf_9a7f80ff75ced90f := __obf_c858f7714608ec5c.__obf_21e4a0c991098223[__obf_97fab2dfcc31983d] + int32(__obf_e4aafdbcb1b13739)
	if __obf_9a7f80ff75ced90f > 0 && __obf_9a7f80ff75ced90f < int32(len(__obf_c858f7714608ec5c.__obf_bdde4c8ff394b211)) && __obf_c858f7714608ec5c.__obf_bdde4c8ff394b211[__obf_9a7f80ff75ced90f] == __obf_97fab2dfcc31983d {
		return __obf_9a7f80ff75ced90f
	}
	return -1
}

func (__obf_c858f7714608ec5c *DATTrie) __obf_f7281b41d9f98ac5(__obf_97fab2dfcc31983d int32) bool {
	if __obf_97fab2dfcc31983d > 0 && __obf_97fab2dfcc31983d < int32(len(__obf_c858f7714608ec5c.__obf_90e9f9cd57864c5a)) {
		return __obf_c858f7714608ec5c.__obf_90e9f9cd57864c5a[__obf_97fab2dfcc31983d]
	}
	return false
}

func (__obf_c858f7714608ec5c *DATTrie) Contains(__obf_2f9bfc5f616e4c2e string) bool {
	__obf_c6569bea94c0c483 := []rune(__obf_2f9bfc5f616e4c2e)
	for __obf_09c4704e98f82128 := range __obf_c6569bea94c0c483 {
		__obf_6b3585156999efd8 := int32(1)
		for __obf_8c2ffe965b3e6421 := __obf_09c4704e98f82128; __obf_8c2ffe965b3e6421 < len(__obf_c6569bea94c0c483); __obf_8c2ffe965b3e6421++ {
			__obf_6b3585156999efd8 = __obf_c858f7714608ec5c.__obf_7426f753a0b18cca(__obf_6b3585156999efd8, __obf_c6569bea94c0c483[__obf_8c2ffe965b3e6421])
			if __obf_6b3585156999efd8 == -1 {
				break
			}
			if __obf_c858f7714608ec5c.__obf_f7281b41d9f98ac5(__obf_6b3585156999efd8) {
				return true
			}
		}
	}
	return false
}

func (__obf_c858f7714608ec5c *DATTrie) __obf_a3e898872b98d664(__obf_c6569bea94c0c483 []rune) []MatchResult {
	__obf_92125008916bb172 := make([]MatchResult, 0, __obf_b5584c9e1f4e82bf)
	for __obf_09c4704e98f82128 := range __obf_c6569bea94c0c483 {
		__obf_6b3585156999efd8 := int32(1)
		for __obf_8c2ffe965b3e6421 := __obf_09c4704e98f82128; __obf_8c2ffe965b3e6421 < len(__obf_c6569bea94c0c483); __obf_8c2ffe965b3e6421++ {
			__obf_6b3585156999efd8 = __obf_c858f7714608ec5c.__obf_7426f753a0b18cca(__obf_6b3585156999efd8, __obf_c6569bea94c0c483[__obf_8c2ffe965b3e6421])
			if __obf_6b3585156999efd8 == -1 {
				break
			}
			if __obf_c858f7714608ec5c.__obf_f7281b41d9f98ac5(__obf_6b3585156999efd8) {
				__obf_92125008916bb172 = append(__obf_92125008916bb172, MatchResult{
					Word:  string(__obf_c6569bea94c0c483[__obf_09c4704e98f82128 : __obf_8c2ffe965b3e6421+1]),
					Start: __obf_09c4704e98f82128,
					End:   __obf_8c2ffe965b3e6421 + 1,
				})
			}
		}
	}
	return __obf_92125008916bb172
}

type WuManberMatcher struct {
	__obf_0d256bcc98a70c22 []string
	__obf_5e405c455150a363 int
	__obf_ecab505238341be1 map[string]int32
	__obf_2e308ab2dce5055a map[string][]int
	__obf_7197038eb37c915f int
}

func NewWuManberMatcher(__obf_5e405c455150a363 int) *WuManberMatcher {
	return &WuManberMatcher{__obf_0d256bcc98a70c22: make([]string, 0), __obf_5e405c455150a363: __obf_5e405c455150a363, __obf_ecab505238341be1: make(map[string]int32), __obf_2e308ab2dce5055a: make(map[string][]int), __obf_7197038eb37c915f: int(^uint(0) >> 1)}
}

func (__obf_eb0bdcb309ca17f2 *WuManberMatcher) AddPatterns(__obf_0d256bcc98a70c22 []string) {
	__obf_eb0bdcb309ca17f2.__obf_0d256bcc98a70c22 = __obf_0d256bcc98a70c22
	for _, __obf_6035c2599a071317 := range __obf_0d256bcc98a70c22 {
		__obf_31c98cbf675c8e03 := utf8.RuneCountInString(__obf_6035c2599a071317)
		if __obf_31c98cbf675c8e03 > 0 && __obf_31c98cbf675c8e03 < __obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f {
			__obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f = __obf_31c98cbf675c8e03
		}
	}
}

func (__obf_eb0bdcb309ca17f2 *WuManberMatcher) Build() {
	if len(__obf_eb0bdcb309ca17f2.__obf_0d256bcc98a70c22) == 0 {
		return
	}
	if __obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f < __obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 {
		__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 = __obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f
	}
	if __obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 <= 0 {
		__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 = 1
	}

	for __obf_09c4704e98f82128, __obf_f139de00b7562515 := range __obf_eb0bdcb309ca17f2.__obf_0d256bcc98a70c22 {
		__obf_c6569bea94c0c483 := []rune(__obf_f139de00b7562515)
		__obf_468675188c263d34 := len(__obf_c6569bea94c0c483)
		for __obf_8c2ffe965b3e6421 := 0; __obf_8c2ffe965b3e6421 <= __obf_468675188c263d34-__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363; __obf_8c2ffe965b3e6421++ {
			__obf_58d90450001daaac := string(__obf_c6569bea94c0c483[__obf_8c2ffe965b3e6421 : __obf_8c2ffe965b3e6421+__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363])
			__obf_cc429112e67f1d4e := int32(__obf_468675188c263d34 - __obf_8c2ffe965b3e6421 - __obf_eb0bdcb309ca17f2.__obf_5e405c455150a363)
			if __obf_c9b6499c56fcac00, __obf_161bbfc9bfb30dc2 := __obf_eb0bdcb309ca17f2.__obf_ecab505238341be1[__obf_58d90450001daaac]; !__obf_161bbfc9bfb30dc2 || __obf_cc429112e67f1d4e < __obf_c9b6499c56fcac00 {
				__obf_eb0bdcb309ca17f2.__obf_ecab505238341be1[__obf_58d90450001daaac] = __obf_cc429112e67f1d4e
			}
		}

		if __obf_468675188c263d34 >= __obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 {
			__obf_2bc3793ec54c8dc4 := string(__obf_c6569bea94c0c483[__obf_468675188c263d34-__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363:])
			__obf_eb0bdcb309ca17f2.__obf_2e308ab2dce5055a[__obf_2bc3793ec54c8dc4] = append(__obf_eb0bdcb309ca17f2.__obf_2e308ab2dce5055a[__obf_2bc3793ec54c8dc4], __obf_09c4704e98f82128)
		}
	}
}

func (__obf_eb0bdcb309ca17f2 *WuManberMatcher) Match(__obf_2f9bfc5f616e4c2e string) bool {
	if len(__obf_eb0bdcb309ca17f2.__obf_0d256bcc98a70c22) == 0 {
		return false
	}
	__obf_c6569bea94c0c483 := []rune(__obf_2f9bfc5f616e4c2e)
	__obf_8d54d637e1c142b6 := len(__obf_c6569bea94c0c483)
	if __obf_8d54d637e1c142b6 < __obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f {
		return false
	}
	__obf_e821b24c447a8d6a := __obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f - 1
	for __obf_e821b24c447a8d6a < __obf_8d54d637e1c142b6 {
		__obf_72595a93817cdaa2 := max(__obf_e821b24c447a8d6a-__obf_eb0bdcb309ca17f2.__obf_5e405c455150a363+1, 0)
		__obf_58d90450001daaac := string(__obf_c6569bea94c0c483[__obf_72595a93817cdaa2 : __obf_e821b24c447a8d6a+1])
		__obf_ecab505238341be1, __obf_161bbfc9bfb30dc2 := __obf_eb0bdcb309ca17f2.__obf_ecab505238341be1[__obf_58d90450001daaac]
		if !__obf_161bbfc9bfb30dc2 {
			__obf_e821b24c447a8d6a += (__obf_eb0bdcb309ca17f2.__obf_7197038eb37c915f - __obf_eb0bdcb309ca17f2.__obf_5e405c455150a363 + 1)
			continue
		}
		if __obf_ecab505238341be1 == 0 {
			if __obf_7f17fdd3da1fdaec, __obf_97c4a66ba64905ca := __obf_eb0bdcb309ca17f2.__obf_2e308ab2dce5055a[__obf_58d90450001daaac]; __obf_97c4a66ba64905ca {
				for _, __obf_a3f2e1d560217eae := range __obf_7f17fdd3da1fdaec {
					__obf_f139de00b7562515 := __obf_eb0bdcb309ca17f2.__obf_0d256bcc98a70c22[__obf_a3f2e1d560217eae]
					__obf_5c460035be1bcb7d := utf8.RuneCountInString(__obf_f139de00b7562515)
					__obf_e92287697e7d7a89 := __obf_e821b24c447a8d6a - __obf_5c460035be1bcb7d + 1
					if __obf_e92287697e7d7a89 >= 0 && string(__obf_c6569bea94c0c483[__obf_e92287697e7d7a89:__obf_e92287697e7d7a89+__obf_5c460035be1bcb7d]) == __obf_f139de00b7562515 {
						return true
					}
				}
			}
			__obf_e821b24c447a8d6a++
		} else {
			__obf_e821b24c447a8d6a += int(__obf_ecab505238341be1)
		}
	}
	return false
}
