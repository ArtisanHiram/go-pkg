package __obf_22f1b0ffe180f6be

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
	__obf_a178e7e88f8b73dd = 1024
	__obf_25e33ab2b640605d = 1_000_000
	__obf_9925bc4b5558b8a9 = // 1MB
	3
	__obf_758717872097a8e6 = 100
	__obf_2d64cc667c4a31a6 = 16
)

// 位图实现
type __obf_d9523d4c35540e5a struct {
	__obf_59d6dda9781ee62e []uint64
}

func __obf_150f2e71a723dfa6(__obf_6b8673cb2adcc5ca int) *__obf_d9523d4c35540e5a {
	__obf_3c54b85276dbca40 := // 每个 uint64 存储 64 位
		(__obf_6b8673cb2adcc5ca + 63) / 64
	return &__obf_d9523d4c35540e5a{__obf_59d6dda9781ee62e: make([]uint64, __obf_3c54b85276dbca40)}
}

func (__obf_06b93b022f5607b0 *__obf_d9523d4c35540e5a) __obf_6d8113e5855afe6a(__obf_96544632c03330e8 int) {
	__obf_06b93b022f5607b0.__obf_59d6dda9781ee62e[__obf_96544632c03330e8/64] |= 1 << (__obf_96544632c03330e8 % 64)
}

func (__obf_06b93b022f5607b0 *__obf_d9523d4c35540e5a) __obf_4939102778525da1(__obf_96544632c03330e8 int) bool {
	return __obf_06b93b022f5607b0.__obf_59d6dda9781ee62e[__obf_96544632c03330e8/64]&(1<<(__obf_96544632c03330e8%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_149268cc63a4a5ef *DATTrie
	__obf_d47ee72164e7063b *WuManberMatcher
	__obf_25ef03ff9fa3c18a sync.RWMutex
}

func (__obf_f44350334ea861e0 *SensitiveFilter) Contains(__obf_a1b5965ae1f72bca string) bool {
	__obf_f44350334ea861e0.__obf_25ef03ff9fa3c18a.
		Lock()
	defer __obf_f44350334ea861e0.__obf_25ef03ff9fa3c18a.Unlock()
	if len(__obf_a1b5965ae1f72bca) > __obf_758717872097a8e6 {
		if __obf_f44350334ea861e0.__obf_d47ee72164e7063b.Match(__obf_a1b5965ae1f72bca) {
			return true
		} else {
			return false
		}
	}
	return __obf_f44350334ea861e0.__obf_149268cc63a4a5ef.Contains(__obf_a1b5965ae1f72bca)
}

func (__obf_f44350334ea861e0 *SensitiveFilter) FindAll(__obf_a1b5965ae1f72bca string) []MatchResult {
	if __obf_f44350334ea861e0 == nil || __obf_f44350334ea861e0.__obf_149268cc63a4a5ef == nil {
		return nil
	}
	if __obf_a1b5965ae1f72bca == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_a1b5965ae1f72bca) > __obf_25e33ab2b640605d {
		return nil // 或返回错误
	}
	__obf_acfb34643eebfd95 := // 1. 找出所有可能重叠的匹配项
		__obf_f44350334ea861e0.__obf_149268cc63a4a5ef.__obf_88de8b9f7dfc747b([]rune(__obf_a1b5965ae1f72bca))
	if len(__obf_acfb34643eebfd95) < 2 {
		return __obf_acfb34643eebfd95
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_acfb34643eebfd95, func(__obf_90eab7913c086801, __obf_e789441fbb5c7599 int) bool {
		__obf_a73946f84a4a9d39 := __obf_acfb34643eebfd95[__obf_90eab7913c086801].End - __obf_acfb34643eebfd95[__obf_90eab7913c086801].Start
		__obf_b70fe7d18ad624bf := __obf_acfb34643eebfd95[__obf_e789441fbb5c7599].End - __obf_acfb34643eebfd95[__obf_e789441fbb5c7599].Start
		if __obf_a73946f84a4a9d39 != __obf_b70fe7d18ad624bf {
			return __obf_a73946f84a4a9d39 > __obf_b70fe7d18ad624bf
		}
		return __obf_acfb34643eebfd95[__obf_90eab7913c086801].Start < __obf_acfb34643eebfd95[__obf_e789441fbb5c7599].Start
	})
	__obf_5b26d525d3c5339e := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_acfb34643eebfd95))
	__obf_4bab427574ea1e16 := // occupied := make([]bool, len(runes))
		__obf_150f2e71a723dfa6(len([]rune(__obf_a1b5965ae1f72bca)))

	for _, __obf_9814464d2062400b := range __obf_acfb34643eebfd95 {
		__obf_17bb693d223fd98d := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_90eab7913c086801 := __obf_9814464d2062400b.Start; __obf_90eab7913c086801 < __obf_9814464d2062400b.End; __obf_90eab7913c086801++ {
			// if occupied[i] {
			if __obf_4bab427574ea1e16.__obf_4939102778525da1(__obf_90eab7913c086801) {
				__obf_17bb693d223fd98d = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_17bb693d223fd98d {
			__obf_5b26d525d3c5339e = append(__obf_5b26d525d3c5339e,
				// 标记该位置已被占用
				__obf_9814464d2062400b)

			for __obf_90eab7913c086801 := __obf_9814464d2062400b.Start; __obf_90eab7913c086801 < __obf_9814464d2062400b.End; __obf_90eab7913c086801++ {
				__obf_4bab427574ea1e16.
					// occupied[i] = true
					__obf_6d8113e5855afe6a(__obf_90eab7913c086801)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_5b26d525d3c5339e, func(__obf_90eab7913c086801, __obf_e789441fbb5c7599 int) bool {
		return __obf_5b26d525d3c5339e[__obf_90eab7913c086801].Start < __obf_5b26d525d3c5339e[__obf_e789441fbb5c7599].Start
	})

	return __obf_5b26d525d3c5339e
}

func (__obf_f44350334ea861e0 *SensitiveFilter) Replace(__obf_a1b5965ae1f72bca string, __obf_f33a8d3331aede3b rune) string {
	__obf_1ac543ad08f631d8 := __obf_f44350334ea861e0.FindAll(__obf_a1b5965ae1f72bca)
	if len(__obf_1ac543ad08f631d8) == 0 {
		return __obf_a1b5965ae1f72bca
	}
	__obf_c1661e43c591b057 := []rune(__obf_a1b5965ae1f72bca)
	for _, __obf_9814464d2062400b := range __obf_1ac543ad08f631d8 {
		for __obf_90eab7913c086801 := __obf_9814464d2062400b.Start; __obf_90eab7913c086801 < __obf_9814464d2062400b.End; __obf_90eab7913c086801++ {
			__obf_c1661e43c591b057[__obf_90eab7913c086801] = __obf_f33a8d3331aede3b
		}
	}
	return string(__obf_c1661e43c591b057)
}

func (__obf_f44350334ea861e0 *SensitiveFilter) ReplaceWithString(__obf_a1b5965ae1f72bca string, __obf_f33a8d3331aede3b string) string {
	__obf_1ac543ad08f631d8 := __obf_f44350334ea861e0.FindAll(__obf_a1b5965ae1f72bca)
	if len(__obf_1ac543ad08f631d8) == 0 {
		return __obf_a1b5965ae1f72bca
	}

	var __obf_832f8b77a4160156 strings.Builder
	__obf_c1661e43c591b057 := []rune(__obf_a1b5965ae1f72bca)
	__obf_85a154d45b9634a2 := 0
	for _, __obf_9814464d2062400b := range __obf_1ac543ad08f631d8 {
		__obf_832f8b77a4160156.
			WriteString(string(__obf_c1661e43c591b057[__obf_85a154d45b9634a2:__obf_9814464d2062400b.Start]))
		__obf_832f8b77a4160156.
			WriteString(__obf_f33a8d3331aede3b)
		__obf_85a154d45b9634a2 = __obf_9814464d2062400b.End
	}
	__obf_832f8b77a4160156.
		WriteString(string(__obf_c1661e43c591b057[__obf_85a154d45b9634a2:]))
	return __obf_832f8b77a4160156.String()
}

func (__obf_f44350334ea861e0 *SensitiveFilter) ReplaceWithFunc(__obf_a1b5965ae1f72bca string, __obf_fed52046bbc4cfea func(__obf_9814464d2062400b MatchResult) string) string {
	__obf_1ac543ad08f631d8 := __obf_f44350334ea861e0.FindAll(__obf_a1b5965ae1f72bca)
	if len(__obf_1ac543ad08f631d8) == 0 {
		return __obf_a1b5965ae1f72bca
	}

	var __obf_832f8b77a4160156 strings.Builder
	__obf_c1661e43c591b057 := []rune(__obf_a1b5965ae1f72bca)
	__obf_85a154d45b9634a2 := 0
	for _, __obf_9814464d2062400b := range __obf_1ac543ad08f631d8 {
		__obf_832f8b77a4160156.
			WriteString(string(__obf_c1661e43c591b057[__obf_85a154d45b9634a2:__obf_9814464d2062400b.Start]))
		__obf_832f8b77a4160156.
			WriteString(__obf_fed52046bbc4cfea(__obf_9814464d2062400b))
		__obf_85a154d45b9634a2 = __obf_9814464d2062400b.End
	}
	__obf_832f8b77a4160156.
		WriteString(string(__obf_c1661e43c591b057[__obf_85a154d45b9634a2:]))
	return __obf_832f8b77a4160156.String()
}

// 添加 Close 方法释放资源
func (__obf_f44350334ea861e0 *SensitiveFilter) Close() {
	__obf_f44350334ea861e0.__obf_25ef03ff9fa3c18a.
		Lock()
	defer __obf_f44350334ea861e0.__obf_25ef03ff9fa3c18a.Unlock()
	__obf_f44350334ea861e0.__obf_149268cc63a4a5ef = nil
	__obf_f44350334ea861e0.__obf_d47ee72164e7063b = nil
}

type FilterBuilder struct {
	__obf_c5992a597bc545e9 map[string]struct{}
	__obf_b57e0987b8094cd8 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_6b8673cb2adcc5ca int) FilterOption {
	return func(__obf_06b93b022f5607b0 *FilterBuilder) {
		if __obf_6b8673cb2adcc5ca > 0 {
			__obf_06b93b022f5607b0.__obf_b57e0987b8094cd8 = __obf_6b8673cb2adcc5ca
		}
	}
}

func WithWords(__obf_c5992a597bc545e9 []string) FilterOption {
	return func(__obf_06b93b022f5607b0 *FilterBuilder) {
		for _, __obf_6615e440a321a469 := range __obf_c5992a597bc545e9 {
			if __obf_6615e440a321a469 != "" {
				__obf_06b93b022f5607b0.__obf_c5992a597bc545e9[__obf_6615e440a321a469] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_04b05aeb94379043 ...FilterOption) *FilterBuilder {
	__obf_06b93b022f5607b0 := &FilterBuilder{__obf_c5992a597bc545e9: make(map[string]struct{}), __obf_b57e0987b8094cd8: __obf_9925bc4b5558b8a9}
	for _, __obf_015878c83093f903 := range __obf_04b05aeb94379043 {
		__obf_015878c83093f903(__obf_06b93b022f5607b0)
	}
	return __obf_06b93b022f5607b0
}

func (__obf_06b93b022f5607b0 *FilterBuilder) AddWord(__obf_636278d651eb5ff7 string) {
	if __obf_636278d651eb5ff7 != "" {
		__obf_06b93b022f5607b0.__obf_c5992a597bc545e9[__obf_636278d651eb5ff7] = struct{}{}
	}
}

func (__obf_06b93b022f5607b0 *FilterBuilder) AddWords(__obf_c5992a597bc545e9 []string) {
	for _, __obf_636278d651eb5ff7 := range __obf_c5992a597bc545e9 {
		__obf_06b93b022f5607b0.
			AddWord(__obf_636278d651eb5ff7)
	}
}

func (__obf_06b93b022f5607b0 *FilterBuilder) LoadFromFile(__obf_84ebefc42a39f344 string) error {
	__obf_0e3db7f2a767029f, __obf_59790140096babe6 := os.Open(__obf_84ebefc42a39f344)
	if __obf_59790140096babe6 != nil {
		return __obf_59790140096babe6
	}
	defer __obf_0e3db7f2a767029f.Close()
	__obf_b437df655dfe020a := bufio.NewReader(__obf_0e3db7f2a767029f)
	__obf_50c4db890b3708aa := []byte{0xEF, 0xBB, 0xBF}
	__obf_11669613aea05df8, __obf_59790140096babe6 := __obf_b437df655dfe020a.Peek(3)
	if __obf_59790140096babe6 == nil && bytes.Equal(__obf_11669613aea05df8, __obf_50c4db890b3708aa) {
		_, _ = __obf_b437df655dfe020a.Discard(3)
	}
	__obf_f0c8737f5cd71985 := bufio.NewScanner(__obf_b437df655dfe020a)
	for __obf_f0c8737f5cd71985.Scan() {
		__obf_e425585dc02e7f01 := strings.TrimSpace(__obf_f0c8737f5cd71985.Text())
		__obf_06b93b022f5607b0.
			AddWord(__obf_e425585dc02e7f01)
	}
	return __obf_f0c8737f5cd71985.Err()
}

func (__obf_06b93b022f5607b0 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_06b93b022f5607b0.__obf_c5992a597bc545e9) == 0 {
		return &SensitiveFilter{__obf_149268cc63a4a5ef: NewDATTrie(), __obf_d47ee72164e7063b: NewWuManberMatcher(__obf_06b93b022f5607b0.__obf_b57e0987b8094cd8)}
	}
	__obf_ac89e17a7bd412b6 := &__obf_2924eb10a9f3934d{__obf_fe05284a4938a737: make(map[rune]*__obf_2924eb10a9f3934d)}
	__obf_eb08e60b11d806b4 := make([]string, 0, len(__obf_06b93b022f5607b0.__obf_c5992a597bc545e9))
	for __obf_636278d651eb5ff7 := range __obf_06b93b022f5607b0.__obf_c5992a597bc545e9 {
		__obf_eb08e60b11d806b4 = append(__obf_eb08e60b11d806b4, __obf_636278d651eb5ff7)
		__obf_ac89e17a7bd412b6.__obf_a6415adecc865108(__obf_636278d651eb5ff7)
	}
	__obf_149268cc63a4a5ef := __obf_c4975ebfbad41128(__obf_ac89e17a7bd412b6)
	__obf_d47ee72164e7063b := NewWuManberMatcher(__obf_06b93b022f5607b0.__obf_b57e0987b8094cd8)
	__obf_d47ee72164e7063b.
		AddPatterns(__obf_eb08e60b11d806b4)
	__obf_d47ee72164e7063b.
		Build()

	return &SensitiveFilter{__obf_149268cc63a4a5ef: __obf_149268cc63a4a5ef, __obf_d47ee72164e7063b: __obf_d47ee72164e7063b}
}

func (__obf_06b93b022f5607b0 *FilterBuilder) Clear() {
	__obf_06b93b022f5607b0.__obf_c5992a597bc545e9 = make(map[string]struct{})
}

type __obf_2924eb10a9f3934d struct {
	__obf_fe05284a4938a737 map[rune]*__obf_2924eb10a9f3934d
	__obf_57585e2a322709e2 bool
}

func (__obf_3c54b85276dbca40 *__obf_2924eb10a9f3934d) __obf_a6415adecc865108(__obf_636278d651eb5ff7 string) {
	__obf_7523f0540667c4b9 := __obf_3c54b85276dbca40
	for _, __obf_1642751db5d7c379 := range __obf_636278d651eb5ff7 {
		if _, __obf_0d9767cd504e5421 := __obf_7523f0540667c4b9.__obf_fe05284a4938a737[__obf_1642751db5d7c379]; !__obf_0d9767cd504e5421 {
			__obf_7523f0540667c4b9.__obf_fe05284a4938a737[__obf_1642751db5d7c379] = &__obf_2924eb10a9f3934d{__obf_fe05284a4938a737: make(map[rune]*__obf_2924eb10a9f3934d)}
		}
		__obf_7523f0540667c4b9 = __obf_7523f0540667c4b9.__obf_fe05284a4938a737[__obf_1642751db5d7c379]
	}
	__obf_7523f0540667c4b9.__obf_57585e2a322709e2 = true
}

type DATTrie struct {
	__obf_89d3fa262329516e []int32
	__obf_4eb42ec8fd3a2a49 []int32
	__obf_c73a079cbf96dc29 []bool
	__obf_6b8673cb2adcc5ca int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_89d3fa262329516e: make([]int32, __obf_a178e7e88f8b73dd), __obf_4eb42ec8fd3a2a49: make([]int32, __obf_a178e7e88f8b73dd), __obf_c73a079cbf96dc29: make([]bool, __obf_a178e7e88f8b73dd), __obf_6b8673cb2adcc5ca: 1}
}

func __obf_c4975ebfbad41128(__obf_4a0b336d5b6f3140 *__obf_2924eb10a9f3934d) *DATTrie {
	__obf_88eb68dc4dfc8868 := NewDATTrie()
	__obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49[1] = 1

	type __obf_1a1a568f39a94818 struct {
		__obf_3253008d96b5fb18 int32
		__obf_7523f0540667c4b9 *__obf_2924eb10a9f3934d
	}
	__obf_fdb9831588f43443 := []__obf_1a1a568f39a94818{{__obf_3253008d96b5fb18: 1, __obf_7523f0540667c4b9: __obf_4a0b336d5b6f3140}}

	for len(__obf_fdb9831588f43443) > 0 {
		__obf_694ea1d1daa72c83 := __obf_fdb9831588f43443[0]
		__obf_fdb9831588f43443 = __obf_fdb9831588f43443[1:]

		if len(__obf_694ea1d1daa72c83.__obf_7523f0540667c4b9.__obf_fe05284a4938a737) == 0 {
			continue
		}
		__obf_799a588b5e306d55 := make([]rune, 0, len(__obf_694ea1d1daa72c83.__obf_7523f0540667c4b9.__obf_fe05284a4938a737))
		for __obf_1642751db5d7c379 := range __obf_694ea1d1daa72c83.__obf_7523f0540667c4b9.__obf_fe05284a4938a737 {
			__obf_799a588b5e306d55 = append(__obf_799a588b5e306d55, __obf_1642751db5d7c379)
		}
		slices.Sort(__obf_799a588b5e306d55)
		__obf_4583c4439954a929 := __obf_88eb68dc4dfc8868.__obf_e6a05fd89d420acf(__obf_799a588b5e306d55)
		__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e[__obf_694ea1d1daa72c83.__obf_3253008d96b5fb18] = __obf_4583c4439954a929

		for _, __obf_b0c09eab0c214931 := range __obf_799a588b5e306d55 {
			__obf_6a14537885c2af39 := __obf_694ea1d1daa72c83.__obf_7523f0540667c4b9.__obf_fe05284a4938a737[__obf_b0c09eab0c214931]
			__obf_5c0cea033c93c10e := __obf_4583c4439954a929 + int32(__obf_b0c09eab0c214931)
			__obf_88eb68dc4dfc8868.__obf_c24aaffd3b3f111f(__obf_5c0cea033c93c10e + 1)
			__obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49[__obf_5c0cea033c93c10e] = __obf_694ea1d1daa72c83.__obf_3253008d96b5fb18
			if __obf_6a14537885c2af39.__obf_57585e2a322709e2 {
				__obf_88eb68dc4dfc8868.__obf_c73a079cbf96dc29[__obf_5c0cea033c93c10e] = true
			}
			if int32(len(__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e)) <= __obf_5c0cea033c93c10e {
				__obf_88eb68dc4dfc8868.__obf_c24aaffd3b3f111f(__obf_5c0cea033c93c10e + 1)
			}
			if __obf_88eb68dc4dfc8868.__obf_6b8673cb2adcc5ca <= __obf_5c0cea033c93c10e {
				__obf_88eb68dc4dfc8868.__obf_6b8673cb2adcc5ca = __obf_5c0cea033c93c10e + 1
			}
			__obf_fdb9831588f43443 = append(__obf_fdb9831588f43443, __obf_1a1a568f39a94818{__obf_3253008d96b5fb18: __obf_5c0cea033c93c10e, __obf_7523f0540667c4b9: __obf_6a14537885c2af39})
		}
	}
	return __obf_88eb68dc4dfc8868
}

func (__obf_88eb68dc4dfc8868 *DATTrie) __obf_e6a05fd89d420acf(__obf_71bb010d56328df3 []rune) int32 {
	for __obf_89d3fa262329516e := int32(1); ; __obf_89d3fa262329516e++ {
		__obf_30da92f1cca861ee := false
		for _, __obf_6da5d7c57ce8ba1e := range __obf_71bb010d56328df3 {
			__obf_fa880660aaccd7da := __obf_89d3fa262329516e + int32(__obf_6da5d7c57ce8ba1e)
			if int32(len(__obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49)) <= __obf_fa880660aaccd7da {
				__obf_88eb68dc4dfc8868.__obf_c24aaffd3b3f111f(__obf_fa880660aaccd7da + 1)
			}
			if __obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49[__obf_fa880660aaccd7da] != 0 {
				__obf_30da92f1cca861ee = true
				break
			}
		}
		if !__obf_30da92f1cca861ee {
			return __obf_89d3fa262329516e
		}
	}
}

func (__obf_88eb68dc4dfc8868 *DATTrie) __obf_c24aaffd3b3f111f(__obf_bee60065d4f9788c int32) {
	if __obf_bee60065d4f9788c < int32(len(__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e)) {
		return
	}
	__obf_18422dc187a30618 := max(int32(len(__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e))+int32(len(__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e))/2, __obf_bee60065d4f9788c)
	__obf_fdf8e5471c62e339 := make([]int32, __obf_18422dc187a30618)
	__obf_11baf44b6e07c786 := make([]int32, __obf_18422dc187a30618)
	__obf_535f51e0cfc5daf9 := make([]bool, __obf_18422dc187a30618)
	copy(__obf_fdf8e5471c62e339, __obf_88eb68dc4dfc8868.__obf_89d3fa262329516e)
	copy(__obf_11baf44b6e07c786, __obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49)
	copy(__obf_535f51e0cfc5daf9, __obf_88eb68dc4dfc8868.__obf_c73a079cbf96dc29)
	__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e = __obf_fdf8e5471c62e339
	__obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49 = __obf_11baf44b6e07c786
	__obf_88eb68dc4dfc8868.__obf_c73a079cbf96dc29 = __obf_535f51e0cfc5daf9
}

func (__obf_88eb68dc4dfc8868 *DATTrie) __obf_334f2b7b677d1b94(__obf_3539ef774937fce8 int32, __obf_6da5d7c57ce8ba1e rune) int32 {
	if __obf_3539ef774937fce8 <= 0 || __obf_3539ef774937fce8 >= int32(len(__obf_88eb68dc4dfc8868.__obf_89d3fa262329516e)) || __obf_88eb68dc4dfc8868.__obf_89d3fa262329516e[__obf_3539ef774937fce8] == 0 {
		return -1
	}
	__obf_fa880660aaccd7da := __obf_88eb68dc4dfc8868.__obf_89d3fa262329516e[__obf_3539ef774937fce8] + int32(__obf_6da5d7c57ce8ba1e)
	if __obf_fa880660aaccd7da > 0 && __obf_fa880660aaccd7da < int32(len(__obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49)) && __obf_88eb68dc4dfc8868.__obf_4eb42ec8fd3a2a49[__obf_fa880660aaccd7da] == __obf_3539ef774937fce8 {
		return __obf_fa880660aaccd7da
	}
	return -1
}

func (__obf_88eb68dc4dfc8868 *DATTrie) __obf_57585e2a322709e2(__obf_3539ef774937fce8 int32) bool {
	if __obf_3539ef774937fce8 > 0 && __obf_3539ef774937fce8 < int32(len(__obf_88eb68dc4dfc8868.__obf_c73a079cbf96dc29)) {
		return __obf_88eb68dc4dfc8868.__obf_c73a079cbf96dc29[__obf_3539ef774937fce8]
	}
	return false
}

func (__obf_88eb68dc4dfc8868 *DATTrie) Contains(__obf_a1b5965ae1f72bca string) bool {
	__obf_c1661e43c591b057 := []rune(__obf_a1b5965ae1f72bca)
	for __obf_90eab7913c086801 := range __obf_c1661e43c591b057 {
		__obf_7523f0540667c4b9 := int32(1)
		for __obf_e789441fbb5c7599 := __obf_90eab7913c086801; __obf_e789441fbb5c7599 < len(__obf_c1661e43c591b057); __obf_e789441fbb5c7599++ {
			__obf_7523f0540667c4b9 = __obf_88eb68dc4dfc8868.__obf_334f2b7b677d1b94(__obf_7523f0540667c4b9, __obf_c1661e43c591b057[__obf_e789441fbb5c7599])
			if __obf_7523f0540667c4b9 == -1 {
				break
			}
			if __obf_88eb68dc4dfc8868.__obf_57585e2a322709e2(__obf_7523f0540667c4b9) {
				return true
			}
		}
	}
	return false
}

func (__obf_88eb68dc4dfc8868 *DATTrie) __obf_88de8b9f7dfc747b(__obf_c1661e43c591b057 []rune) []MatchResult {
	__obf_47a7d10a5a6057e0 := make([]MatchResult, 0, __obf_2d64cc667c4a31a6)
	for __obf_90eab7913c086801 := range __obf_c1661e43c591b057 {
		__obf_7523f0540667c4b9 := int32(1)
		for __obf_e789441fbb5c7599 := __obf_90eab7913c086801; __obf_e789441fbb5c7599 < len(__obf_c1661e43c591b057); __obf_e789441fbb5c7599++ {
			__obf_7523f0540667c4b9 = __obf_88eb68dc4dfc8868.__obf_334f2b7b677d1b94(__obf_7523f0540667c4b9, __obf_c1661e43c591b057[__obf_e789441fbb5c7599])
			if __obf_7523f0540667c4b9 == -1 {
				break
			}
			if __obf_88eb68dc4dfc8868.__obf_57585e2a322709e2(__obf_7523f0540667c4b9) {
				__obf_47a7d10a5a6057e0 = append(__obf_47a7d10a5a6057e0, MatchResult{
					Word:  string(__obf_c1661e43c591b057[__obf_90eab7913c086801 : __obf_e789441fbb5c7599+1]),
					Start: __obf_90eab7913c086801,
					End:   __obf_e789441fbb5c7599 + 1,
				})
			}
		}
	}
	return __obf_47a7d10a5a6057e0
}

type WuManberMatcher struct {
	__obf_eb254f0167c8cc0e []string
	__obf_b57e0987b8094cd8 int
	__obf_3a16805c3ef1fbec map[string]int32
	__obf_88c4064b51353c57 map[string][]int
	__obf_294a583c56c23cec int
}

func NewWuManberMatcher(__obf_b57e0987b8094cd8 int) *WuManberMatcher {
	return &WuManberMatcher{__obf_eb254f0167c8cc0e: make([]string, 0), __obf_b57e0987b8094cd8: __obf_b57e0987b8094cd8, __obf_3a16805c3ef1fbec: make(map[string]int32), __obf_88c4064b51353c57: make(map[string][]int), __obf_294a583c56c23cec: int(^uint(0) >> 1)}
}

func (__obf_ff8e7687e6f8d238 *WuManberMatcher) AddPatterns(__obf_eb254f0167c8cc0e []string) {
	__obf_ff8e7687e6f8d238.__obf_eb254f0167c8cc0e = __obf_eb254f0167c8cc0e
	for _, __obf_a14babd08ad7c8bb := range __obf_eb254f0167c8cc0e {
		__obf_3fdc48117b031ff2 := utf8.RuneCountInString(__obf_a14babd08ad7c8bb)
		if __obf_3fdc48117b031ff2 > 0 && __obf_3fdc48117b031ff2 < __obf_ff8e7687e6f8d238.__obf_294a583c56c23cec {
			__obf_ff8e7687e6f8d238.__obf_294a583c56c23cec = __obf_3fdc48117b031ff2
		}
	}
}

func (__obf_ff8e7687e6f8d238 *WuManberMatcher) Build() {
	if len(__obf_ff8e7687e6f8d238.__obf_eb254f0167c8cc0e) == 0 {
		return
	}
	if __obf_ff8e7687e6f8d238.__obf_294a583c56c23cec < __obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 {
		__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 = __obf_ff8e7687e6f8d238.__obf_294a583c56c23cec
	}
	if __obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 <= 0 {
		__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 = 1
	}

	for __obf_90eab7913c086801, __obf_870d444a8ef54595 := range __obf_ff8e7687e6f8d238.__obf_eb254f0167c8cc0e {
		__obf_c1661e43c591b057 := []rune(__obf_870d444a8ef54595)
		__obf_37d09f1602ff98f3 := len(__obf_c1661e43c591b057)
		for __obf_e789441fbb5c7599 := 0; __obf_e789441fbb5c7599 <= __obf_37d09f1602ff98f3-__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8; __obf_e789441fbb5c7599++ {
			__obf_64ad7b2b1a11124a := string(__obf_c1661e43c591b057[__obf_e789441fbb5c7599 : __obf_e789441fbb5c7599+__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8])
			__obf_01c48539d692eab0 := int32(__obf_37d09f1602ff98f3 - __obf_e789441fbb5c7599 - __obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8)
			if __obf_2cfa7b05532ae108, __obf_875e73d8c6d10821 := __obf_ff8e7687e6f8d238.__obf_3a16805c3ef1fbec[__obf_64ad7b2b1a11124a]; !__obf_875e73d8c6d10821 || __obf_01c48539d692eab0 < __obf_2cfa7b05532ae108 {
				__obf_ff8e7687e6f8d238.__obf_3a16805c3ef1fbec[__obf_64ad7b2b1a11124a] = __obf_01c48539d692eab0
			}
		}

		if __obf_37d09f1602ff98f3 >= __obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 {
			__obf_b859ef9ff8b842c8 := string(__obf_c1661e43c591b057[__obf_37d09f1602ff98f3-__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8:])
			__obf_ff8e7687e6f8d238.__obf_88c4064b51353c57[__obf_b859ef9ff8b842c8] = append(__obf_ff8e7687e6f8d238.__obf_88c4064b51353c57[__obf_b859ef9ff8b842c8], __obf_90eab7913c086801)
		}
	}
}

func (__obf_ff8e7687e6f8d238 *WuManberMatcher) Match(__obf_a1b5965ae1f72bca string) bool {
	if len(__obf_ff8e7687e6f8d238.__obf_eb254f0167c8cc0e) == 0 {
		return false
	}
	__obf_c1661e43c591b057 := []rune(__obf_a1b5965ae1f72bca)
	__obf_6df84f00b5a925c6 := len(__obf_c1661e43c591b057)
	if __obf_6df84f00b5a925c6 < __obf_ff8e7687e6f8d238.__obf_294a583c56c23cec {
		return false
	}
	__obf_96544632c03330e8 := __obf_ff8e7687e6f8d238.__obf_294a583c56c23cec - 1
	for __obf_96544632c03330e8 < __obf_6df84f00b5a925c6 {
		__obf_191e0a011d7c9264 := max(__obf_96544632c03330e8-__obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8+1, 0)
		__obf_64ad7b2b1a11124a := string(__obf_c1661e43c591b057[__obf_191e0a011d7c9264 : __obf_96544632c03330e8+1])
		__obf_3a16805c3ef1fbec, __obf_875e73d8c6d10821 := __obf_ff8e7687e6f8d238.__obf_3a16805c3ef1fbec[__obf_64ad7b2b1a11124a]
		if !__obf_875e73d8c6d10821 {
			__obf_96544632c03330e8 += (__obf_ff8e7687e6f8d238.__obf_294a583c56c23cec - __obf_ff8e7687e6f8d238.__obf_b57e0987b8094cd8 + 1)
			continue
		}
		if __obf_3a16805c3ef1fbec == 0 {
			if __obf_1927e54a154dcfdc, __obf_0d9767cd504e5421 := __obf_ff8e7687e6f8d238.__obf_88c4064b51353c57[__obf_64ad7b2b1a11124a]; __obf_0d9767cd504e5421 {
				for _, __obf_ac4e22622c16b928 := range __obf_1927e54a154dcfdc {
					__obf_870d444a8ef54595 := __obf_ff8e7687e6f8d238.__obf_eb254f0167c8cc0e[__obf_ac4e22622c16b928]
					__obf_d9bfca9727e194ac := utf8.RuneCountInString(__obf_870d444a8ef54595)
					__obf_e976370197910668 := __obf_96544632c03330e8 - __obf_d9bfca9727e194ac + 1
					if __obf_e976370197910668 >= 0 && string(__obf_c1661e43c591b057[__obf_e976370197910668:__obf_e976370197910668+__obf_d9bfca9727e194ac]) == __obf_870d444a8ef54595 {
						return true
					}
				}
			}
			__obf_96544632c03330e8++
		} else {
			__obf_96544632c03330e8 += int(__obf_3a16805c3ef1fbec)
		}
	}
	return false
}
