package __obf_49e32fb15fc6cbca

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
	__obf_b3af71592ffe7b45 = 1024
	__obf_097f529acdc0ec7f = 1_000_000 // 1MB
	__obf_958c2f08f7fc34ff = 3
	__obf_749596bd6c58520b = 100
	__obf_502304427eed492e = 16
)

// 位图实现
type __obf_8c974a24a812c5fd struct {
	__obf_b1e0355499eb168c []uint64
}

func __obf_d3baa6565d9b2bf1(__obf_37a27ddf449840f6 int) *__obf_8c974a24a812c5fd {
	// 每个 uint64 存储 64 位
	__obf_207a6a05c815e798 := (__obf_37a27ddf449840f6 + 63) / 64
	return &__obf_8c974a24a812c5fd{__obf_b1e0355499eb168c: make([]uint64, __obf_207a6a05c815e798)}
}

func (__obf_c87f0e7965811aa1 *__obf_8c974a24a812c5fd) __obf_f04e6ce9d0de59ad(__obf_993a16dcd21f3f67 int) {
	__obf_c87f0e7965811aa1.__obf_b1e0355499eb168c[__obf_993a16dcd21f3f67/64] |= 1 << (__obf_993a16dcd21f3f67 % 64)
}

func (__obf_c87f0e7965811aa1 *__obf_8c974a24a812c5fd) __obf_d06ff645fb68a1ba(__obf_993a16dcd21f3f67 int) bool {
	return __obf_c87f0e7965811aa1.__obf_b1e0355499eb168c[__obf_993a16dcd21f3f67/64]&(1<<(__obf_993a16dcd21f3f67%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_e84dc0e2aeaade32 *DATTrie
	__obf_fb4e836e84f3ac94 *WuManberMatcher
	__obf_b5c49117d8b42ac7 sync.RWMutex
}

func (__obf_263cd2356f0163b1 *SensitiveFilter) Contains(__obf_97f6376cbdb6230d string) bool {
	__obf_263cd2356f0163b1.__obf_b5c49117d8b42ac7.Lock()
	defer __obf_263cd2356f0163b1.__obf_b5c49117d8b42ac7.Unlock()
	if len(__obf_97f6376cbdb6230d) > __obf_749596bd6c58520b {
		if __obf_263cd2356f0163b1.__obf_fb4e836e84f3ac94.Match(__obf_97f6376cbdb6230d) {
			return true
		} else {
			return false
		}
	}
	return __obf_263cd2356f0163b1.__obf_e84dc0e2aeaade32.Contains(__obf_97f6376cbdb6230d)
}

func (__obf_263cd2356f0163b1 *SensitiveFilter) FindAll(__obf_97f6376cbdb6230d string) []MatchResult {
	if __obf_263cd2356f0163b1 == nil || __obf_263cd2356f0163b1.__obf_e84dc0e2aeaade32 == nil {
		return nil
	}
	if __obf_97f6376cbdb6230d == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_97f6376cbdb6230d) > __obf_097f529acdc0ec7f {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_435403b653bc1363 := __obf_263cd2356f0163b1.__obf_e84dc0e2aeaade32.__obf_1697909b29002289([]rune(__obf_97f6376cbdb6230d))
	if len(__obf_435403b653bc1363) < 2 {
		return __obf_435403b653bc1363
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_435403b653bc1363, func(__obf_1ec68bf22e83bbaa, __obf_6ef7d5c3ea22af66 int) bool {
		__obf_7e6d342639823351 := __obf_435403b653bc1363[__obf_1ec68bf22e83bbaa].End - __obf_435403b653bc1363[__obf_1ec68bf22e83bbaa].Start
		__obf_0cf07070b1217de7 := __obf_435403b653bc1363[__obf_6ef7d5c3ea22af66].End - __obf_435403b653bc1363[__obf_6ef7d5c3ea22af66].Start
		if __obf_7e6d342639823351 != __obf_0cf07070b1217de7 {
			return __obf_7e6d342639823351 > __obf_0cf07070b1217de7
		}
		return __obf_435403b653bc1363[__obf_1ec68bf22e83bbaa].Start < __obf_435403b653bc1363[__obf_6ef7d5c3ea22af66].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_09967e0d498e29f1 := make([]MatchResult, 0, len(__obf_435403b653bc1363))
	// occupied := make([]bool, len(runes))
	__obf_6b8667f36288bede := __obf_d3baa6565d9b2bf1(len([]rune(__obf_97f6376cbdb6230d)))

	for _, __obf_46099bf46f877831 := range __obf_435403b653bc1363 {
		__obf_9c53d70cacb23c1f := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_1ec68bf22e83bbaa := __obf_46099bf46f877831.Start; __obf_1ec68bf22e83bbaa < __obf_46099bf46f877831.End; __obf_1ec68bf22e83bbaa++ {
			// if occupied[i] {
			if __obf_6b8667f36288bede.__obf_d06ff645fb68a1ba(__obf_1ec68bf22e83bbaa) {
				__obf_9c53d70cacb23c1f = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_9c53d70cacb23c1f {
			__obf_09967e0d498e29f1 = append(__obf_09967e0d498e29f1, __obf_46099bf46f877831)
			// 标记该位置已被占用
			for __obf_1ec68bf22e83bbaa := __obf_46099bf46f877831.Start; __obf_1ec68bf22e83bbaa < __obf_46099bf46f877831.End; __obf_1ec68bf22e83bbaa++ {
				// occupied[i] = true
				__obf_6b8667f36288bede.__obf_f04e6ce9d0de59ad(__obf_1ec68bf22e83bbaa)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_09967e0d498e29f1, func(__obf_1ec68bf22e83bbaa, __obf_6ef7d5c3ea22af66 int) bool {
		return __obf_09967e0d498e29f1[__obf_1ec68bf22e83bbaa].Start < __obf_09967e0d498e29f1[__obf_6ef7d5c3ea22af66].Start
	})

	return __obf_09967e0d498e29f1
}

func (__obf_263cd2356f0163b1 *SensitiveFilter) Replace(__obf_97f6376cbdb6230d string, __obf_8e0e1df3210afe11 rune) string {
	__obf_238f36331dfb72f9 := __obf_263cd2356f0163b1.FindAll(__obf_97f6376cbdb6230d)
	if len(__obf_238f36331dfb72f9) == 0 {
		return __obf_97f6376cbdb6230d
	}

	__obf_05976448b987f949 := []rune(__obf_97f6376cbdb6230d)
	for _, __obf_46099bf46f877831 := range __obf_238f36331dfb72f9 {
		for __obf_1ec68bf22e83bbaa := __obf_46099bf46f877831.Start; __obf_1ec68bf22e83bbaa < __obf_46099bf46f877831.End; __obf_1ec68bf22e83bbaa++ {
			__obf_05976448b987f949[__obf_1ec68bf22e83bbaa] = __obf_8e0e1df3210afe11
		}
	}
	return string(__obf_05976448b987f949)
}

func (__obf_263cd2356f0163b1 *SensitiveFilter) ReplaceWithString(__obf_97f6376cbdb6230d string, __obf_8e0e1df3210afe11 string) string {
	__obf_238f36331dfb72f9 := __obf_263cd2356f0163b1.FindAll(__obf_97f6376cbdb6230d)
	if len(__obf_238f36331dfb72f9) == 0 {
		return __obf_97f6376cbdb6230d
	}

	var __obf_d89402c873c6ba57 strings.Builder
	__obf_05976448b987f949 := []rune(__obf_97f6376cbdb6230d)
	__obf_362deb6cfa8bfdcc := 0
	for _, __obf_46099bf46f877831 := range __obf_238f36331dfb72f9 {
		__obf_d89402c873c6ba57.WriteString(string(__obf_05976448b987f949[__obf_362deb6cfa8bfdcc:__obf_46099bf46f877831.Start]))
		__obf_d89402c873c6ba57.WriteString(__obf_8e0e1df3210afe11)
		__obf_362deb6cfa8bfdcc = __obf_46099bf46f877831.End
	}
	__obf_d89402c873c6ba57.WriteString(string(__obf_05976448b987f949[__obf_362deb6cfa8bfdcc:]))
	return __obf_d89402c873c6ba57.String()
}

func (__obf_263cd2356f0163b1 *SensitiveFilter) ReplaceWithFunc(__obf_97f6376cbdb6230d string, __obf_0201329f4ae9e6c9 func(__obf_46099bf46f877831 MatchResult) string) string {
	__obf_238f36331dfb72f9 := __obf_263cd2356f0163b1.FindAll(__obf_97f6376cbdb6230d)
	if len(__obf_238f36331dfb72f9) == 0 {
		return __obf_97f6376cbdb6230d
	}

	var __obf_d89402c873c6ba57 strings.Builder
	__obf_05976448b987f949 := []rune(__obf_97f6376cbdb6230d)
	__obf_362deb6cfa8bfdcc := 0
	for _, __obf_46099bf46f877831 := range __obf_238f36331dfb72f9 {
		__obf_d89402c873c6ba57.WriteString(string(__obf_05976448b987f949[__obf_362deb6cfa8bfdcc:__obf_46099bf46f877831.Start]))
		__obf_d89402c873c6ba57.WriteString(__obf_0201329f4ae9e6c9(__obf_46099bf46f877831))
		__obf_362deb6cfa8bfdcc = __obf_46099bf46f877831.End
	}
	__obf_d89402c873c6ba57.WriteString(string(__obf_05976448b987f949[__obf_362deb6cfa8bfdcc:]))
	return __obf_d89402c873c6ba57.String()
}

// 添加 Close 方法释放资源
func (__obf_263cd2356f0163b1 *SensitiveFilter) Close() {
	__obf_263cd2356f0163b1.__obf_b5c49117d8b42ac7.Lock()
	defer __obf_263cd2356f0163b1.__obf_b5c49117d8b42ac7.Unlock()

	__obf_263cd2356f0163b1.__obf_e84dc0e2aeaade32 = nil
	__obf_263cd2356f0163b1.__obf_fb4e836e84f3ac94 = nil
}

type FilterBuilder struct {
	__obf_91accb9d4ab4e5f0 map[string]struct{}
	__obf_c081dfcaef8c67ef int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_37a27ddf449840f6 int) FilterOption {
	return func(__obf_c87f0e7965811aa1 *FilterBuilder) {
		if __obf_37a27ddf449840f6 > 0 {
			__obf_c87f0e7965811aa1.__obf_c081dfcaef8c67ef = __obf_37a27ddf449840f6
		}
	}
}

func WithWords(__obf_91accb9d4ab4e5f0 []string) FilterOption {
	return func(__obf_c87f0e7965811aa1 *FilterBuilder) {
		for _, __obf_beae256b56724464 := range __obf_91accb9d4ab4e5f0 {
			if __obf_beae256b56724464 != "" {
				__obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0[__obf_beae256b56724464] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_c52e242802c1c96c ...FilterOption) *FilterBuilder {
	__obf_c87f0e7965811aa1 := &FilterBuilder{
		__obf_91accb9d4ab4e5f0: make(map[string]struct{}),
		__obf_c081dfcaef8c67ef: __obf_958c2f08f7fc34ff,
	}
	for _, __obf_9007d8fb635b6a58 := range __obf_c52e242802c1c96c {
		__obf_9007d8fb635b6a58(__obf_c87f0e7965811aa1)
	}
	return __obf_c87f0e7965811aa1
}

func (__obf_c87f0e7965811aa1 *FilterBuilder) AddWord(__obf_130ba2bc34059d2b string) {
	if __obf_130ba2bc34059d2b != "" {
		__obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0[__obf_130ba2bc34059d2b] = struct{}{}
	}
}

func (__obf_c87f0e7965811aa1 *FilterBuilder) AddWords(__obf_91accb9d4ab4e5f0 []string) {
	for _, __obf_130ba2bc34059d2b := range __obf_91accb9d4ab4e5f0 {
		__obf_c87f0e7965811aa1.AddWord(__obf_130ba2bc34059d2b)
	}
}

func (__obf_c87f0e7965811aa1 *FilterBuilder) LoadFromFile(__obf_bd589e4977713680 string) error {
	__obf_15b79167220164da, __obf_b63bfff967e4c831 := os.Open(__obf_bd589e4977713680)
	if __obf_b63bfff967e4c831 != nil {
		return __obf_b63bfff967e4c831
	}
	defer __obf_15b79167220164da.Close()

	__obf_37c86ae8faa33e9d := bufio.NewReader(__obf_15b79167220164da)
	__obf_e767aeadcf0ec265 := []byte{0xEF, 0xBB, 0xBF}
	__obf_708a6f58483933dc, __obf_b63bfff967e4c831 := __obf_37c86ae8faa33e9d.Peek(3)
	if __obf_b63bfff967e4c831 == nil && bytes.Equal(__obf_708a6f58483933dc, __obf_e767aeadcf0ec265) {
		_, _ = __obf_37c86ae8faa33e9d.Discard(3)
	}

	__obf_775e5f63d2ba2ad4 := bufio.NewScanner(__obf_37c86ae8faa33e9d)
	for __obf_775e5f63d2ba2ad4.Scan() {
		__obf_f1cfd4013cb9ea3e := strings.TrimSpace(__obf_775e5f63d2ba2ad4.Text())
		__obf_c87f0e7965811aa1.AddWord(__obf_f1cfd4013cb9ea3e)
	}
	return __obf_775e5f63d2ba2ad4.Err()
}

func (__obf_c87f0e7965811aa1 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0) == 0 {
		return &SensitiveFilter{__obf_e84dc0e2aeaade32: NewDATTrie(),
			__obf_fb4e836e84f3ac94: NewWuManberMatcher(__obf_c87f0e7965811aa1.__obf_c081dfcaef8c67ef)}
	}

	__obf_64900ed7aa06a94b := &__obf_145596435449a272{__obf_e61e70bd2b51b124: make(map[rune]*__obf_145596435449a272)}
	__obf_c2c3de9b3daa9fd0 := make([]string, 0, len(__obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0))
	for __obf_130ba2bc34059d2b := range __obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0 {
		__obf_c2c3de9b3daa9fd0 = append(__obf_c2c3de9b3daa9fd0, __obf_130ba2bc34059d2b)
		__obf_64900ed7aa06a94b.__obf_ba3ed4a4c79b5cf7(__obf_130ba2bc34059d2b)
	}

	__obf_e84dc0e2aeaade32 := __obf_533e3cdb22840795(__obf_64900ed7aa06a94b)
	__obf_fb4e836e84f3ac94 := NewWuManberMatcher(__obf_c87f0e7965811aa1.__obf_c081dfcaef8c67ef)
	__obf_fb4e836e84f3ac94.AddPatterns(__obf_c2c3de9b3daa9fd0)
	__obf_fb4e836e84f3ac94.Build()

	return &SensitiveFilter{
		__obf_e84dc0e2aeaade32: __obf_e84dc0e2aeaade32,
		__obf_fb4e836e84f3ac94: __obf_fb4e836e84f3ac94,
	}
}

func (__obf_c87f0e7965811aa1 *FilterBuilder) Clear() {
	__obf_c87f0e7965811aa1.__obf_91accb9d4ab4e5f0 = make(map[string]struct{})
}

type __obf_145596435449a272 struct {
	__obf_e61e70bd2b51b124 map[rune]*__obf_145596435449a272
	__obf_bb04ebcd034a854d bool
}

func (__obf_207a6a05c815e798 *__obf_145596435449a272) __obf_ba3ed4a4c79b5cf7(__obf_130ba2bc34059d2b string) {
	__obf_0d8ae933c2fc5fa6 := __obf_207a6a05c815e798
	for _, __obf_ed308405d66e7e0b := range __obf_130ba2bc34059d2b {
		if _, __obf_73740bb85c6c299c := __obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124[__obf_ed308405d66e7e0b]; !__obf_73740bb85c6c299c {
			__obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124[__obf_ed308405d66e7e0b] = &__obf_145596435449a272{__obf_e61e70bd2b51b124: make(map[rune]*__obf_145596435449a272)}
		}
		__obf_0d8ae933c2fc5fa6 = __obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124[__obf_ed308405d66e7e0b]
	}
	__obf_0d8ae933c2fc5fa6.__obf_bb04ebcd034a854d = true
}

type DATTrie struct {
	__obf_ee9e97f203c47010 []int32
	__obf_13414edfa91f5d90 []int32
	__obf_3b84c9b7cde0957a []bool
	__obf_37a27ddf449840f6 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_ee9e97f203c47010: make([]int32, __obf_b3af71592ffe7b45),
		__obf_13414edfa91f5d90: make([]int32, __obf_b3af71592ffe7b45),
		__obf_3b84c9b7cde0957a: make([]bool, __obf_b3af71592ffe7b45),
		__obf_37a27ddf449840f6: 1,
	}
}

func __obf_533e3cdb22840795(__obf_2fc69aa1d67c6ae5 *__obf_145596435449a272) *DATTrie {
	__obf_c8c8c1b65c718057 := NewDATTrie()
	__obf_c8c8c1b65c718057.__obf_13414edfa91f5d90[1] = 1

	type __obf_1d9fb45a4effe3e0 struct {
		__obf_4661e888ec51df02 int32
		__obf_0d8ae933c2fc5fa6 *__obf_145596435449a272
	}

	__obf_e7bf96645f1feb1f := []__obf_1d9fb45a4effe3e0{{__obf_4661e888ec51df02: 1, __obf_0d8ae933c2fc5fa6: __obf_2fc69aa1d67c6ae5}}

	for len(__obf_e7bf96645f1feb1f) > 0 {
		__obf_0d5f6a8ae0de4bc8 := __obf_e7bf96645f1feb1f[0]
		__obf_e7bf96645f1feb1f = __obf_e7bf96645f1feb1f[1:]

		if len(__obf_0d5f6a8ae0de4bc8.__obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124) == 0 {
			continue
		}

		__obf_db17c30c4b067299 := make([]rune, 0, len(__obf_0d5f6a8ae0de4bc8.__obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124))
		for __obf_ed308405d66e7e0b := range __obf_0d5f6a8ae0de4bc8.__obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124 {
			__obf_db17c30c4b067299 = append(__obf_db17c30c4b067299, __obf_ed308405d66e7e0b)
		}
		slices.Sort(__obf_db17c30c4b067299)

		__obf_4d9713d8586c4711 := __obf_c8c8c1b65c718057.__obf_d8b7ffebc042b96e(__obf_db17c30c4b067299)
		__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010[__obf_0d5f6a8ae0de4bc8.__obf_4661e888ec51df02] = __obf_4d9713d8586c4711

		for _, __obf_a3248292255c6ed5 := range __obf_db17c30c4b067299 {
			__obf_3173803c76df667b := __obf_0d5f6a8ae0de4bc8.__obf_0d8ae933c2fc5fa6.__obf_e61e70bd2b51b124[__obf_a3248292255c6ed5]
			__obf_2d016292afdbf307 := __obf_4d9713d8586c4711 + int32(__obf_a3248292255c6ed5)
			__obf_c8c8c1b65c718057.__obf_6eb8cada251e675f(__obf_2d016292afdbf307 + 1)

			__obf_c8c8c1b65c718057.__obf_13414edfa91f5d90[__obf_2d016292afdbf307] = __obf_0d5f6a8ae0de4bc8.__obf_4661e888ec51df02
			if __obf_3173803c76df667b.__obf_bb04ebcd034a854d {
				__obf_c8c8c1b65c718057.__obf_3b84c9b7cde0957a[__obf_2d016292afdbf307] = true
			}
			if int32(len(__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010)) <= __obf_2d016292afdbf307 {
				__obf_c8c8c1b65c718057.__obf_6eb8cada251e675f(__obf_2d016292afdbf307 + 1)
			}
			if __obf_c8c8c1b65c718057.__obf_37a27ddf449840f6 <= __obf_2d016292afdbf307 {
				__obf_c8c8c1b65c718057.__obf_37a27ddf449840f6 = __obf_2d016292afdbf307 + 1
			}

			__obf_e7bf96645f1feb1f = append(__obf_e7bf96645f1feb1f, __obf_1d9fb45a4effe3e0{__obf_4661e888ec51df02: __obf_2d016292afdbf307, __obf_0d8ae933c2fc5fa6: __obf_3173803c76df667b})
		}
	}
	return __obf_c8c8c1b65c718057
}

func (__obf_c8c8c1b65c718057 *DATTrie) __obf_d8b7ffebc042b96e(__obf_43f2f60b762e05e2 []rune) int32 {
	for __obf_ee9e97f203c47010 := int32(1); ; __obf_ee9e97f203c47010++ {
		__obf_78f7e6b2d6652374 := false
		for _, __obf_22a5d992c757b5ff := range __obf_43f2f60b762e05e2 {
			__obf_09c120d0da7ae09a := __obf_ee9e97f203c47010 + int32(__obf_22a5d992c757b5ff)
			if int32(len(__obf_c8c8c1b65c718057.__obf_13414edfa91f5d90)) <= __obf_09c120d0da7ae09a {
				__obf_c8c8c1b65c718057.__obf_6eb8cada251e675f(__obf_09c120d0da7ae09a + 1)
			}
			if __obf_c8c8c1b65c718057.__obf_13414edfa91f5d90[__obf_09c120d0da7ae09a] != 0 {
				__obf_78f7e6b2d6652374 = true
				break
			}
		}
		if !__obf_78f7e6b2d6652374 {
			return __obf_ee9e97f203c47010
		}
	}
}

func (__obf_c8c8c1b65c718057 *DATTrie) __obf_6eb8cada251e675f(__obf_f45a261e9adde871 int32) {
	if __obf_f45a261e9adde871 < int32(len(__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010)) {
		return
	}
	__obf_e9dc8d3916ca246e := max(int32(len(__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010))+int32(len(__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010))/2, __obf_f45a261e9adde871)
	__obf_1e258f352032bfed := make([]int32, __obf_e9dc8d3916ca246e)
	__obf_97be5ffb47e4f915 := make([]int32, __obf_e9dc8d3916ca246e)
	__obf_abdab9507350db41 := make([]bool, __obf_e9dc8d3916ca246e)
	copy(__obf_1e258f352032bfed, __obf_c8c8c1b65c718057.__obf_ee9e97f203c47010)
	copy(__obf_97be5ffb47e4f915, __obf_c8c8c1b65c718057.__obf_13414edfa91f5d90)
	copy(__obf_abdab9507350db41, __obf_c8c8c1b65c718057.__obf_3b84c9b7cde0957a)
	__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010 = __obf_1e258f352032bfed
	__obf_c8c8c1b65c718057.__obf_13414edfa91f5d90 = __obf_97be5ffb47e4f915
	__obf_c8c8c1b65c718057.__obf_3b84c9b7cde0957a = __obf_abdab9507350db41
}

func (__obf_c8c8c1b65c718057 *DATTrie) __obf_d33f0b8d739e13c7(__obf_79943565671ba9cc int32, __obf_22a5d992c757b5ff rune) int32 {
	if __obf_79943565671ba9cc <= 0 || __obf_79943565671ba9cc >= int32(len(__obf_c8c8c1b65c718057.__obf_ee9e97f203c47010)) || __obf_c8c8c1b65c718057.__obf_ee9e97f203c47010[__obf_79943565671ba9cc] == 0 {
		return -1
	}
	__obf_09c120d0da7ae09a := __obf_c8c8c1b65c718057.__obf_ee9e97f203c47010[__obf_79943565671ba9cc] + int32(__obf_22a5d992c757b5ff)
	if __obf_09c120d0da7ae09a > 0 && __obf_09c120d0da7ae09a < int32(len(__obf_c8c8c1b65c718057.__obf_13414edfa91f5d90)) && __obf_c8c8c1b65c718057.__obf_13414edfa91f5d90[__obf_09c120d0da7ae09a] == __obf_79943565671ba9cc {
		return __obf_09c120d0da7ae09a
	}
	return -1
}

func (__obf_c8c8c1b65c718057 *DATTrie) __obf_bb04ebcd034a854d(__obf_79943565671ba9cc int32) bool {
	if __obf_79943565671ba9cc > 0 && __obf_79943565671ba9cc < int32(len(__obf_c8c8c1b65c718057.__obf_3b84c9b7cde0957a)) {
		return __obf_c8c8c1b65c718057.__obf_3b84c9b7cde0957a[__obf_79943565671ba9cc]
	}
	return false
}

func (__obf_c8c8c1b65c718057 *DATTrie) Contains(__obf_97f6376cbdb6230d string) bool {
	__obf_05976448b987f949 := []rune(__obf_97f6376cbdb6230d)
	for __obf_1ec68bf22e83bbaa := range __obf_05976448b987f949 {
		__obf_0d8ae933c2fc5fa6 := int32(1)
		for __obf_6ef7d5c3ea22af66 := __obf_1ec68bf22e83bbaa; __obf_6ef7d5c3ea22af66 < len(__obf_05976448b987f949); __obf_6ef7d5c3ea22af66++ {
			__obf_0d8ae933c2fc5fa6 = __obf_c8c8c1b65c718057.__obf_d33f0b8d739e13c7(__obf_0d8ae933c2fc5fa6, __obf_05976448b987f949[__obf_6ef7d5c3ea22af66])
			if __obf_0d8ae933c2fc5fa6 == -1 {
				break
			}
			if __obf_c8c8c1b65c718057.__obf_bb04ebcd034a854d(__obf_0d8ae933c2fc5fa6) {
				return true
			}
		}
	}
	return false
}

func (__obf_c8c8c1b65c718057 *DATTrie) __obf_1697909b29002289(__obf_05976448b987f949 []rune) []MatchResult {
	__obf_8056ad475c2ff239 := make([]MatchResult, 0, __obf_502304427eed492e)
	for __obf_1ec68bf22e83bbaa := range __obf_05976448b987f949 {
		__obf_0d8ae933c2fc5fa6 := int32(1)
		for __obf_6ef7d5c3ea22af66 := __obf_1ec68bf22e83bbaa; __obf_6ef7d5c3ea22af66 < len(__obf_05976448b987f949); __obf_6ef7d5c3ea22af66++ {
			__obf_0d8ae933c2fc5fa6 = __obf_c8c8c1b65c718057.__obf_d33f0b8d739e13c7(__obf_0d8ae933c2fc5fa6, __obf_05976448b987f949[__obf_6ef7d5c3ea22af66])
			if __obf_0d8ae933c2fc5fa6 == -1 {
				break
			}
			if __obf_c8c8c1b65c718057.__obf_bb04ebcd034a854d(__obf_0d8ae933c2fc5fa6) {
				__obf_8056ad475c2ff239 = append(__obf_8056ad475c2ff239, MatchResult{
					Word:  string(__obf_05976448b987f949[__obf_1ec68bf22e83bbaa : __obf_6ef7d5c3ea22af66+1]),
					Start: __obf_1ec68bf22e83bbaa,
					End:   __obf_6ef7d5c3ea22af66 + 1,
				})
			}
		}
	}
	return __obf_8056ad475c2ff239
}

type WuManberMatcher struct {
	__obf_a7cb8bf5e5ab3dbc []string
	__obf_c081dfcaef8c67ef int
	__obf_d4788b1999a21fb8 map[string]int32
	__obf_c9270bd40803ca8a map[string][]int
	__obf_15d89da8f5382126 int
}

func NewWuManberMatcher(__obf_c081dfcaef8c67ef int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_a7cb8bf5e5ab3dbc: make([]string, 0),
		__obf_c081dfcaef8c67ef: __obf_c081dfcaef8c67ef,
		__obf_d4788b1999a21fb8: make(map[string]int32),
		__obf_c9270bd40803ca8a: make(map[string][]int),
		__obf_15d89da8f5382126: int(^uint(0) >> 1),
	}
}

func (__obf_947c9f9939368cd9 *WuManberMatcher) AddPatterns(__obf_a7cb8bf5e5ab3dbc []string) {
	__obf_947c9f9939368cd9.__obf_a7cb8bf5e5ab3dbc = __obf_a7cb8bf5e5ab3dbc
	for _, __obf_70c17e5eb7309c6b := range __obf_a7cb8bf5e5ab3dbc {
		__obf_b1c3190e833d59c8 := utf8.RuneCountInString(__obf_70c17e5eb7309c6b)
		if __obf_b1c3190e833d59c8 > 0 && __obf_b1c3190e833d59c8 < __obf_947c9f9939368cd9.__obf_15d89da8f5382126 {
			__obf_947c9f9939368cd9.__obf_15d89da8f5382126 = __obf_b1c3190e833d59c8
		}
	}
}

func (__obf_947c9f9939368cd9 *WuManberMatcher) Build() {
	if len(__obf_947c9f9939368cd9.__obf_a7cb8bf5e5ab3dbc) == 0 {
		return
	}
	if __obf_947c9f9939368cd9.__obf_15d89da8f5382126 < __obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef {
		__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef = __obf_947c9f9939368cd9.__obf_15d89da8f5382126
	}
	if __obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef <= 0 {
		__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef = 1
	}

	for __obf_1ec68bf22e83bbaa, __obf_285763c6b50b66ab := range __obf_947c9f9939368cd9.__obf_a7cb8bf5e5ab3dbc {
		__obf_05976448b987f949 := []rune(__obf_285763c6b50b66ab)
		__obf_6dfb3df7ba655b78 := len(__obf_05976448b987f949)
		for __obf_6ef7d5c3ea22af66 := 0; __obf_6ef7d5c3ea22af66 <= __obf_6dfb3df7ba655b78-__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef; __obf_6ef7d5c3ea22af66++ {
			__obf_4025ac22a012f7bc := string(__obf_05976448b987f949[__obf_6ef7d5c3ea22af66 : __obf_6ef7d5c3ea22af66+__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef])
			__obf_675dabc782472496 := int32(__obf_6dfb3df7ba655b78 - __obf_6ef7d5c3ea22af66 - __obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef)
			if __obf_53d288d6257046e8, __obf_ce591b077d7afa96 := __obf_947c9f9939368cd9.__obf_d4788b1999a21fb8[__obf_4025ac22a012f7bc]; !__obf_ce591b077d7afa96 || __obf_675dabc782472496 < __obf_53d288d6257046e8 {
				__obf_947c9f9939368cd9.__obf_d4788b1999a21fb8[__obf_4025ac22a012f7bc] = __obf_675dabc782472496
			}
		}

		if __obf_6dfb3df7ba655b78 >= __obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef {
			__obf_0d483267b7ed0457 := string(__obf_05976448b987f949[__obf_6dfb3df7ba655b78-__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef:])
			__obf_947c9f9939368cd9.__obf_c9270bd40803ca8a[__obf_0d483267b7ed0457] = append(__obf_947c9f9939368cd9.__obf_c9270bd40803ca8a[__obf_0d483267b7ed0457], __obf_1ec68bf22e83bbaa)
		}
	}
}

func (__obf_947c9f9939368cd9 *WuManberMatcher) Match(__obf_97f6376cbdb6230d string) bool {
	if len(__obf_947c9f9939368cd9.__obf_a7cb8bf5e5ab3dbc) == 0 {
		return false
	}
	__obf_05976448b987f949 := []rune(__obf_97f6376cbdb6230d)
	__obf_ced8ee6752c158b0 := len(__obf_05976448b987f949)
	if __obf_ced8ee6752c158b0 < __obf_947c9f9939368cd9.__obf_15d89da8f5382126 {
		return false
	}
	__obf_993a16dcd21f3f67 := __obf_947c9f9939368cd9.__obf_15d89da8f5382126 - 1
	for __obf_993a16dcd21f3f67 < __obf_ced8ee6752c158b0 {
		__obf_88f2b1974bb1aa33 := max(__obf_993a16dcd21f3f67-__obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef+1, 0)
		__obf_4025ac22a012f7bc := string(__obf_05976448b987f949[__obf_88f2b1974bb1aa33 : __obf_993a16dcd21f3f67+1])
		__obf_d4788b1999a21fb8, __obf_ce591b077d7afa96 := __obf_947c9f9939368cd9.__obf_d4788b1999a21fb8[__obf_4025ac22a012f7bc]
		if !__obf_ce591b077d7afa96 {
			__obf_993a16dcd21f3f67 += (__obf_947c9f9939368cd9.__obf_15d89da8f5382126 - __obf_947c9f9939368cd9.__obf_c081dfcaef8c67ef + 1)
			continue
		}
		if __obf_d4788b1999a21fb8 == 0 {
			if __obf_3fa9b366efeac267, __obf_73740bb85c6c299c := __obf_947c9f9939368cd9.__obf_c9270bd40803ca8a[__obf_4025ac22a012f7bc]; __obf_73740bb85c6c299c {
				for _, __obf_240fcf8ff1b9f0b1 := range __obf_3fa9b366efeac267 {
					__obf_285763c6b50b66ab := __obf_947c9f9939368cd9.__obf_a7cb8bf5e5ab3dbc[__obf_240fcf8ff1b9f0b1]
					__obf_5ddb29739f3152b3 := utf8.RuneCountInString(__obf_285763c6b50b66ab)
					__obf_60ae600a46178da4 := __obf_993a16dcd21f3f67 - __obf_5ddb29739f3152b3 + 1
					if __obf_60ae600a46178da4 >= 0 && string(__obf_05976448b987f949[__obf_60ae600a46178da4:__obf_60ae600a46178da4+__obf_5ddb29739f3152b3]) == __obf_285763c6b50b66ab {
						return true
					}
				}
			}
			__obf_993a16dcd21f3f67++
		} else {
			__obf_993a16dcd21f3f67 += int(__obf_d4788b1999a21fb8)
		}
	}
	return false
}
