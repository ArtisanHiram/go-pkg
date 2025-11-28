package __obf_42695b1bf73991c0

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
	__obf_e0e166ce44bc42b0 = 1024
	__obf_3a88c9b588a197f7 = 1_000_000 // 1MB
	__obf_6ef7a738aae79b28 = 3
	__obf_a3c5ffa5edbc2c45 = 100
	__obf_ad212734531787e5 = 16
)

// 位图实现
type __obf_1bc75846f7dbcd73 struct {
	__obf_0d9dc3ce04c53b7e []uint64
}

func __obf_e1db7c879c9c715c(__obf_77148cd28a144387 int) *__obf_1bc75846f7dbcd73 {
	// 每个 uint64 存储 64 位
	__obf_279c7cc327e79da3 := (__obf_77148cd28a144387 + 63) / 64
	return &__obf_1bc75846f7dbcd73{__obf_0d9dc3ce04c53b7e: make([]uint64, __obf_279c7cc327e79da3)}
}

func (__obf_aa2e78feeb44897f *__obf_1bc75846f7dbcd73) __obf_973b277b350f1232(__obf_62d67604be0b3fe8 int) {
	__obf_aa2e78feeb44897f.__obf_0d9dc3ce04c53b7e[__obf_62d67604be0b3fe8/64] |= 1 << (__obf_62d67604be0b3fe8 % 64)
}

func (__obf_aa2e78feeb44897f *__obf_1bc75846f7dbcd73) __obf_1e5300c99808ac8e(__obf_62d67604be0b3fe8 int) bool {
	return __obf_aa2e78feeb44897f.__obf_0d9dc3ce04c53b7e[__obf_62d67604be0b3fe8/64]&(1<<(__obf_62d67604be0b3fe8%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_24c6cd346be81337 *DATTrie
	__obf_ba399dd3d896cc03 *WuManberMatcher
	__obf_f336bbfffd346999 sync.RWMutex
}

func (__obf_efed0950529e7498 *SensitiveFilter) Contains(__obf_c5e81538f7c752f6 string) bool {
	__obf_efed0950529e7498.__obf_f336bbfffd346999.Lock()
	defer __obf_efed0950529e7498.__obf_f336bbfffd346999.Unlock()
	if len(__obf_c5e81538f7c752f6) > __obf_a3c5ffa5edbc2c45 {
		if __obf_efed0950529e7498.__obf_ba399dd3d896cc03.Match(__obf_c5e81538f7c752f6) {
			return true
		} else {
			return false
		}
	}
	return __obf_efed0950529e7498.__obf_24c6cd346be81337.Contains(__obf_c5e81538f7c752f6)
}

func (__obf_efed0950529e7498 *SensitiveFilter) FindAll(__obf_c5e81538f7c752f6 string) []MatchResult {
	if __obf_efed0950529e7498 == nil || __obf_efed0950529e7498.__obf_24c6cd346be81337 == nil {
		return nil
	}
	if __obf_c5e81538f7c752f6 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_c5e81538f7c752f6) > __obf_3a88c9b588a197f7 {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_335caf0cf4094a1f := __obf_efed0950529e7498.__obf_24c6cd346be81337.__obf_b2e02f30ad64606a([]rune(__obf_c5e81538f7c752f6))
	if len(__obf_335caf0cf4094a1f) < 2 {
		return __obf_335caf0cf4094a1f
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_335caf0cf4094a1f, func(__obf_1a42297cf5681be2, __obf_a8753e2db05116b7 int) bool {
		__obf_1e60e8005e79bffc := __obf_335caf0cf4094a1f[__obf_1a42297cf5681be2].End - __obf_335caf0cf4094a1f[__obf_1a42297cf5681be2].Start
		__obf_be16b2339b58d64a := __obf_335caf0cf4094a1f[__obf_a8753e2db05116b7].End - __obf_335caf0cf4094a1f[__obf_a8753e2db05116b7].Start
		if __obf_1e60e8005e79bffc != __obf_be16b2339b58d64a {
			return __obf_1e60e8005e79bffc > __obf_be16b2339b58d64a
		}
		return __obf_335caf0cf4094a1f[__obf_1a42297cf5681be2].Start < __obf_335caf0cf4094a1f[__obf_a8753e2db05116b7].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_4dc48b6c7c00dcb7 := make([]MatchResult, 0, len(__obf_335caf0cf4094a1f))
	// occupied := make([]bool, len(runes))
	__obf_ac1e4fa6174775b4 := __obf_e1db7c879c9c715c(len([]rune(__obf_c5e81538f7c752f6)))

	for _, __obf_997e64eb1cb075b5 := range __obf_335caf0cf4094a1f {
		__obf_321a84ec10645b60 := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_1a42297cf5681be2 := __obf_997e64eb1cb075b5.Start; __obf_1a42297cf5681be2 < __obf_997e64eb1cb075b5.End; __obf_1a42297cf5681be2++ {
			// if occupied[i] {
			if __obf_ac1e4fa6174775b4.__obf_1e5300c99808ac8e(__obf_1a42297cf5681be2) {
				__obf_321a84ec10645b60 = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_321a84ec10645b60 {
			__obf_4dc48b6c7c00dcb7 = append(__obf_4dc48b6c7c00dcb7, __obf_997e64eb1cb075b5)
			// 标记该位置已被占用
			for __obf_1a42297cf5681be2 := __obf_997e64eb1cb075b5.Start; __obf_1a42297cf5681be2 < __obf_997e64eb1cb075b5.End; __obf_1a42297cf5681be2++ {
				// occupied[i] = true
				__obf_ac1e4fa6174775b4.__obf_973b277b350f1232(__obf_1a42297cf5681be2)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_4dc48b6c7c00dcb7, func(__obf_1a42297cf5681be2, __obf_a8753e2db05116b7 int) bool {
		return __obf_4dc48b6c7c00dcb7[__obf_1a42297cf5681be2].Start < __obf_4dc48b6c7c00dcb7[__obf_a8753e2db05116b7].Start
	})

	return __obf_4dc48b6c7c00dcb7
}

func (__obf_efed0950529e7498 *SensitiveFilter) Replace(__obf_c5e81538f7c752f6 string, __obf_9f11b66bb5594266 rune) string {
	__obf_686753faf97ee586 := __obf_efed0950529e7498.FindAll(__obf_c5e81538f7c752f6)
	if len(__obf_686753faf97ee586) == 0 {
		return __obf_c5e81538f7c752f6
	}

	__obf_bbebb446a4b21552 := []rune(__obf_c5e81538f7c752f6)
	for _, __obf_997e64eb1cb075b5 := range __obf_686753faf97ee586 {
		for __obf_1a42297cf5681be2 := __obf_997e64eb1cb075b5.Start; __obf_1a42297cf5681be2 < __obf_997e64eb1cb075b5.End; __obf_1a42297cf5681be2++ {
			__obf_bbebb446a4b21552[__obf_1a42297cf5681be2] = __obf_9f11b66bb5594266
		}
	}
	return string(__obf_bbebb446a4b21552)
}

func (__obf_efed0950529e7498 *SensitiveFilter) ReplaceWithString(__obf_c5e81538f7c752f6 string, __obf_9f11b66bb5594266 string) string {
	__obf_686753faf97ee586 := __obf_efed0950529e7498.FindAll(__obf_c5e81538f7c752f6)
	if len(__obf_686753faf97ee586) == 0 {
		return __obf_c5e81538f7c752f6
	}

	var __obf_a0719e0defa7839d strings.Builder
	__obf_bbebb446a4b21552 := []rune(__obf_c5e81538f7c752f6)
	__obf_19d48774e782b716 := 0
	for _, __obf_997e64eb1cb075b5 := range __obf_686753faf97ee586 {
		__obf_a0719e0defa7839d.WriteString(string(__obf_bbebb446a4b21552[__obf_19d48774e782b716:__obf_997e64eb1cb075b5.Start]))
		__obf_a0719e0defa7839d.WriteString(__obf_9f11b66bb5594266)
		__obf_19d48774e782b716 = __obf_997e64eb1cb075b5.End
	}
	__obf_a0719e0defa7839d.WriteString(string(__obf_bbebb446a4b21552[__obf_19d48774e782b716:]))
	return __obf_a0719e0defa7839d.String()
}

func (__obf_efed0950529e7498 *SensitiveFilter) ReplaceWithFunc(__obf_c5e81538f7c752f6 string, __obf_536c36cd0f2a2e30 func(__obf_997e64eb1cb075b5 MatchResult) string) string {
	__obf_686753faf97ee586 := __obf_efed0950529e7498.FindAll(__obf_c5e81538f7c752f6)
	if len(__obf_686753faf97ee586) == 0 {
		return __obf_c5e81538f7c752f6
	}

	var __obf_a0719e0defa7839d strings.Builder
	__obf_bbebb446a4b21552 := []rune(__obf_c5e81538f7c752f6)
	__obf_19d48774e782b716 := 0
	for _, __obf_997e64eb1cb075b5 := range __obf_686753faf97ee586 {
		__obf_a0719e0defa7839d.WriteString(string(__obf_bbebb446a4b21552[__obf_19d48774e782b716:__obf_997e64eb1cb075b5.Start]))
		__obf_a0719e0defa7839d.WriteString(__obf_536c36cd0f2a2e30(__obf_997e64eb1cb075b5))
		__obf_19d48774e782b716 = __obf_997e64eb1cb075b5.End
	}
	__obf_a0719e0defa7839d.WriteString(string(__obf_bbebb446a4b21552[__obf_19d48774e782b716:]))
	return __obf_a0719e0defa7839d.String()
}

// 添加 Close 方法释放资源
func (__obf_efed0950529e7498 *SensitiveFilter) Close() {
	__obf_efed0950529e7498.__obf_f336bbfffd346999.Lock()
	defer __obf_efed0950529e7498.__obf_f336bbfffd346999.Unlock()

	__obf_efed0950529e7498.__obf_24c6cd346be81337 = nil
	__obf_efed0950529e7498.__obf_ba399dd3d896cc03 = nil
}

type FilterBuilder struct {
	__obf_cc3f315fd295b345 map[string]struct{}
	__obf_c2522539269e55c0 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_77148cd28a144387 int) FilterOption {
	return func(__obf_aa2e78feeb44897f *FilterBuilder) {
		if __obf_77148cd28a144387 > 0 {
			__obf_aa2e78feeb44897f.__obf_c2522539269e55c0 = __obf_77148cd28a144387
		}
	}
}

func WithWords(__obf_cc3f315fd295b345 []string) FilterOption {
	return func(__obf_aa2e78feeb44897f *FilterBuilder) {
		for _, __obf_8c4365cb807dfed2 := range __obf_cc3f315fd295b345 {
			if __obf_8c4365cb807dfed2 != "" {
				__obf_aa2e78feeb44897f.__obf_cc3f315fd295b345[__obf_8c4365cb807dfed2] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_e7b781813ccb4785 ...FilterOption) *FilterBuilder {
	__obf_aa2e78feeb44897f := &FilterBuilder{
		__obf_cc3f315fd295b345: make(map[string]struct{}),
		__obf_c2522539269e55c0: __obf_6ef7a738aae79b28,
	}
	for _, __obf_60d30722c307d11d := range __obf_e7b781813ccb4785 {
		__obf_60d30722c307d11d(__obf_aa2e78feeb44897f)
	}
	return __obf_aa2e78feeb44897f
}

func (__obf_aa2e78feeb44897f *FilterBuilder) AddWord(__obf_03a71fd9b1b33303 string) {
	if __obf_03a71fd9b1b33303 != "" {
		__obf_aa2e78feeb44897f.__obf_cc3f315fd295b345[__obf_03a71fd9b1b33303] = struct{}{}
	}
}

func (__obf_aa2e78feeb44897f *FilterBuilder) AddWords(__obf_cc3f315fd295b345 []string) {
	for _, __obf_03a71fd9b1b33303 := range __obf_cc3f315fd295b345 {
		__obf_aa2e78feeb44897f.AddWord(__obf_03a71fd9b1b33303)
	}
}

func (__obf_aa2e78feeb44897f *FilterBuilder) LoadFromFile(__obf_c3fa06df5627e35b string) error {
	__obf_9dfaf1fd2ebe473e, __obf_bd29a0c1ef4bbd9e := os.Open(__obf_c3fa06df5627e35b)
	if __obf_bd29a0c1ef4bbd9e != nil {
		return __obf_bd29a0c1ef4bbd9e
	}
	defer __obf_9dfaf1fd2ebe473e.Close()

	__obf_3e85e170309ffa7e := bufio.NewReader(__obf_9dfaf1fd2ebe473e)
	__obf_3df0822a1fa5f018 := []byte{0xEF, 0xBB, 0xBF}
	__obf_ddee6118370b4413, __obf_bd29a0c1ef4bbd9e := __obf_3e85e170309ffa7e.Peek(3)
	if __obf_bd29a0c1ef4bbd9e == nil && bytes.Equal(__obf_ddee6118370b4413, __obf_3df0822a1fa5f018) {
		_, _ = __obf_3e85e170309ffa7e.Discard(3)
	}

	__obf_ce3f71eba51aaeee := bufio.NewScanner(__obf_3e85e170309ffa7e)
	for __obf_ce3f71eba51aaeee.Scan() {
		__obf_ddd8ad53acc05c2a := strings.TrimSpace(__obf_ce3f71eba51aaeee.Text())
		__obf_aa2e78feeb44897f.AddWord(__obf_ddd8ad53acc05c2a)
	}
	return __obf_ce3f71eba51aaeee.Err()
}

func (__obf_aa2e78feeb44897f *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_aa2e78feeb44897f.__obf_cc3f315fd295b345) == 0 {
		return &SensitiveFilter{__obf_24c6cd346be81337: NewDATTrie(),
			__obf_ba399dd3d896cc03: NewWuManberMatcher(__obf_aa2e78feeb44897f.__obf_c2522539269e55c0)}
	}

	__obf_15836a85c8036924 := &__obf_70d56f953f18fc29{__obf_6cbd299ab218e93a: make(map[rune]*__obf_70d56f953f18fc29)}
	__obf_ad259f74b084e0e0 := make([]string, 0, len(__obf_aa2e78feeb44897f.__obf_cc3f315fd295b345))
	for __obf_03a71fd9b1b33303 := range __obf_aa2e78feeb44897f.__obf_cc3f315fd295b345 {
		__obf_ad259f74b084e0e0 = append(__obf_ad259f74b084e0e0, __obf_03a71fd9b1b33303)
		__obf_15836a85c8036924.__obf_3d16e13da4b5d7ff(__obf_03a71fd9b1b33303)
	}

	__obf_24c6cd346be81337 := __obf_5b3aa877f63f70d8(__obf_15836a85c8036924)
	__obf_ba399dd3d896cc03 := NewWuManberMatcher(__obf_aa2e78feeb44897f.__obf_c2522539269e55c0)
	__obf_ba399dd3d896cc03.AddPatterns(__obf_ad259f74b084e0e0)
	__obf_ba399dd3d896cc03.Build()

	return &SensitiveFilter{
		__obf_24c6cd346be81337: __obf_24c6cd346be81337,
		__obf_ba399dd3d896cc03: __obf_ba399dd3d896cc03,
	}
}

func (__obf_aa2e78feeb44897f *FilterBuilder) Clear() {
	__obf_aa2e78feeb44897f.__obf_cc3f315fd295b345 = make(map[string]struct{})
}

type __obf_70d56f953f18fc29 struct {
	__obf_6cbd299ab218e93a map[rune]*__obf_70d56f953f18fc29
	__obf_c90cec341103836a bool
}

func (__obf_279c7cc327e79da3 *__obf_70d56f953f18fc29) __obf_3d16e13da4b5d7ff(__obf_03a71fd9b1b33303 string) {
	__obf_3dc258a80b4bf26c := __obf_279c7cc327e79da3
	for _, __obf_300c537f443a3fcd := range __obf_03a71fd9b1b33303 {
		if _, __obf_d8b944328567cd86 := __obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a[__obf_300c537f443a3fcd]; !__obf_d8b944328567cd86 {
			__obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a[__obf_300c537f443a3fcd] = &__obf_70d56f953f18fc29{__obf_6cbd299ab218e93a: make(map[rune]*__obf_70d56f953f18fc29)}
		}
		__obf_3dc258a80b4bf26c = __obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a[__obf_300c537f443a3fcd]
	}
	__obf_3dc258a80b4bf26c.__obf_c90cec341103836a = true
}

type DATTrie struct {
	__obf_e232d9fe5ae998f0 []int32
	__obf_308adf244162f98f []int32
	__obf_e455d80801940cb6 []bool
	__obf_77148cd28a144387 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_e232d9fe5ae998f0: make([]int32, __obf_e0e166ce44bc42b0),
		__obf_308adf244162f98f: make([]int32, __obf_e0e166ce44bc42b0),
		__obf_e455d80801940cb6: make([]bool, __obf_e0e166ce44bc42b0),
		__obf_77148cd28a144387: 1,
	}
}

func __obf_5b3aa877f63f70d8(__obf_bc0203839deffee0 *__obf_70d56f953f18fc29) *DATTrie {
	__obf_e190518f28bbec8a := NewDATTrie()
	__obf_e190518f28bbec8a.__obf_308adf244162f98f[1] = 1

	type __obf_7f4567c87761901d struct {
		__obf_ec75cbb08b3548cd int32
		__obf_3dc258a80b4bf26c *__obf_70d56f953f18fc29
	}

	__obf_9a9ebc5074c946f1 := []__obf_7f4567c87761901d{{__obf_ec75cbb08b3548cd: 1, __obf_3dc258a80b4bf26c: __obf_bc0203839deffee0}}

	for len(__obf_9a9ebc5074c946f1) > 0 {
		__obf_65a00f07539d1c81 := __obf_9a9ebc5074c946f1[0]
		__obf_9a9ebc5074c946f1 = __obf_9a9ebc5074c946f1[1:]

		if len(__obf_65a00f07539d1c81.__obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a) == 0 {
			continue
		}

		__obf_d92a75a568cbabec := make([]rune, 0, len(__obf_65a00f07539d1c81.__obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a))
		for __obf_300c537f443a3fcd := range __obf_65a00f07539d1c81.__obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a {
			__obf_d92a75a568cbabec = append(__obf_d92a75a568cbabec, __obf_300c537f443a3fcd)
		}
		slices.Sort(__obf_d92a75a568cbabec)

		__obf_98b4eccc31df81b6 := __obf_e190518f28bbec8a.__obf_1c9351e34711565e(__obf_d92a75a568cbabec)
		__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0[__obf_65a00f07539d1c81.__obf_ec75cbb08b3548cd] = __obf_98b4eccc31df81b6

		for _, __obf_31357d8b66cfa19d := range __obf_d92a75a568cbabec {
			__obf_111e4d8c56636566 := __obf_65a00f07539d1c81.__obf_3dc258a80b4bf26c.__obf_6cbd299ab218e93a[__obf_31357d8b66cfa19d]
			__obf_c01957f748607d8c := __obf_98b4eccc31df81b6 + int32(__obf_31357d8b66cfa19d)
			__obf_e190518f28bbec8a.__obf_8068d219cfc05ae9(__obf_c01957f748607d8c + 1)

			__obf_e190518f28bbec8a.__obf_308adf244162f98f[__obf_c01957f748607d8c] = __obf_65a00f07539d1c81.__obf_ec75cbb08b3548cd
			if __obf_111e4d8c56636566.__obf_c90cec341103836a {
				__obf_e190518f28bbec8a.__obf_e455d80801940cb6[__obf_c01957f748607d8c] = true
			}
			if int32(len(__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0)) <= __obf_c01957f748607d8c {
				__obf_e190518f28bbec8a.__obf_8068d219cfc05ae9(__obf_c01957f748607d8c + 1)
			}
			if __obf_e190518f28bbec8a.__obf_77148cd28a144387 <= __obf_c01957f748607d8c {
				__obf_e190518f28bbec8a.__obf_77148cd28a144387 = __obf_c01957f748607d8c + 1
			}

			__obf_9a9ebc5074c946f1 = append(__obf_9a9ebc5074c946f1, __obf_7f4567c87761901d{__obf_ec75cbb08b3548cd: __obf_c01957f748607d8c, __obf_3dc258a80b4bf26c: __obf_111e4d8c56636566})
		}
	}
	return __obf_e190518f28bbec8a
}

func (__obf_e190518f28bbec8a *DATTrie) __obf_1c9351e34711565e(__obf_25f540586aa5cf6c []rune) int32 {
	for __obf_e232d9fe5ae998f0 := int32(1); ; __obf_e232d9fe5ae998f0++ {
		__obf_fe6aad8ffc069d8a := false
		for _, __obf_2c189edf4a9ea297 := range __obf_25f540586aa5cf6c {
			__obf_e814970dc28f2033 := __obf_e232d9fe5ae998f0 + int32(__obf_2c189edf4a9ea297)
			if int32(len(__obf_e190518f28bbec8a.__obf_308adf244162f98f)) <= __obf_e814970dc28f2033 {
				__obf_e190518f28bbec8a.__obf_8068d219cfc05ae9(__obf_e814970dc28f2033 + 1)
			}
			if __obf_e190518f28bbec8a.__obf_308adf244162f98f[__obf_e814970dc28f2033] != 0 {
				__obf_fe6aad8ffc069d8a = true
				break
			}
		}
		if !__obf_fe6aad8ffc069d8a {
			return __obf_e232d9fe5ae998f0
		}
	}
}

func (__obf_e190518f28bbec8a *DATTrie) __obf_8068d219cfc05ae9(__obf_03ad90b14b7db06e int32) {
	if __obf_03ad90b14b7db06e < int32(len(__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0)) {
		return
	}
	__obf_98942c68891f887f := max(int32(len(__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0))+int32(len(__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0))/2, __obf_03ad90b14b7db06e)
	__obf_3fd33a8b4c8dd1b0 := make([]int32, __obf_98942c68891f887f)
	__obf_8c65a136ea1e69f6 := make([]int32, __obf_98942c68891f887f)
	__obf_e14f3c6c126faee9 := make([]bool, __obf_98942c68891f887f)
	copy(__obf_3fd33a8b4c8dd1b0, __obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0)
	copy(__obf_8c65a136ea1e69f6, __obf_e190518f28bbec8a.__obf_308adf244162f98f)
	copy(__obf_e14f3c6c126faee9, __obf_e190518f28bbec8a.__obf_e455d80801940cb6)
	__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0 = __obf_3fd33a8b4c8dd1b0
	__obf_e190518f28bbec8a.__obf_308adf244162f98f = __obf_8c65a136ea1e69f6
	__obf_e190518f28bbec8a.__obf_e455d80801940cb6 = __obf_e14f3c6c126faee9
}

func (__obf_e190518f28bbec8a *DATTrie) __obf_36693c214b3d67f0(__obf_e5a884b023b49c1b int32, __obf_2c189edf4a9ea297 rune) int32 {
	if __obf_e5a884b023b49c1b <= 0 || __obf_e5a884b023b49c1b >= int32(len(__obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0)) || __obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0[__obf_e5a884b023b49c1b] == 0 {
		return -1
	}
	__obf_e814970dc28f2033 := __obf_e190518f28bbec8a.__obf_e232d9fe5ae998f0[__obf_e5a884b023b49c1b] + int32(__obf_2c189edf4a9ea297)
	if __obf_e814970dc28f2033 > 0 && __obf_e814970dc28f2033 < int32(len(__obf_e190518f28bbec8a.__obf_308adf244162f98f)) && __obf_e190518f28bbec8a.__obf_308adf244162f98f[__obf_e814970dc28f2033] == __obf_e5a884b023b49c1b {
		return __obf_e814970dc28f2033
	}
	return -1
}

func (__obf_e190518f28bbec8a *DATTrie) __obf_c90cec341103836a(__obf_e5a884b023b49c1b int32) bool {
	if __obf_e5a884b023b49c1b > 0 && __obf_e5a884b023b49c1b < int32(len(__obf_e190518f28bbec8a.__obf_e455d80801940cb6)) {
		return __obf_e190518f28bbec8a.__obf_e455d80801940cb6[__obf_e5a884b023b49c1b]
	}
	return false
}

func (__obf_e190518f28bbec8a *DATTrie) Contains(__obf_c5e81538f7c752f6 string) bool {
	__obf_bbebb446a4b21552 := []rune(__obf_c5e81538f7c752f6)
	for __obf_1a42297cf5681be2 := range __obf_bbebb446a4b21552 {
		__obf_3dc258a80b4bf26c := int32(1)
		for __obf_a8753e2db05116b7 := __obf_1a42297cf5681be2; __obf_a8753e2db05116b7 < len(__obf_bbebb446a4b21552); __obf_a8753e2db05116b7++ {
			__obf_3dc258a80b4bf26c = __obf_e190518f28bbec8a.__obf_36693c214b3d67f0(__obf_3dc258a80b4bf26c, __obf_bbebb446a4b21552[__obf_a8753e2db05116b7])
			if __obf_3dc258a80b4bf26c == -1 {
				break
			}
			if __obf_e190518f28bbec8a.__obf_c90cec341103836a(__obf_3dc258a80b4bf26c) {
				return true
			}
		}
	}
	return false
}

func (__obf_e190518f28bbec8a *DATTrie) __obf_b2e02f30ad64606a(__obf_bbebb446a4b21552 []rune) []MatchResult {
	__obf_2c2157aaf0123f13 := make([]MatchResult, 0, __obf_ad212734531787e5)
	for __obf_1a42297cf5681be2 := range __obf_bbebb446a4b21552 {
		__obf_3dc258a80b4bf26c := int32(1)
		for __obf_a8753e2db05116b7 := __obf_1a42297cf5681be2; __obf_a8753e2db05116b7 < len(__obf_bbebb446a4b21552); __obf_a8753e2db05116b7++ {
			__obf_3dc258a80b4bf26c = __obf_e190518f28bbec8a.__obf_36693c214b3d67f0(__obf_3dc258a80b4bf26c, __obf_bbebb446a4b21552[__obf_a8753e2db05116b7])
			if __obf_3dc258a80b4bf26c == -1 {
				break
			}
			if __obf_e190518f28bbec8a.__obf_c90cec341103836a(__obf_3dc258a80b4bf26c) {
				__obf_2c2157aaf0123f13 = append(__obf_2c2157aaf0123f13, MatchResult{
					Word:  string(__obf_bbebb446a4b21552[__obf_1a42297cf5681be2 : __obf_a8753e2db05116b7+1]),
					Start: __obf_1a42297cf5681be2,
					End:   __obf_a8753e2db05116b7 + 1,
				})
			}
		}
	}
	return __obf_2c2157aaf0123f13
}

type WuManberMatcher struct {
	__obf_b903a599ae1e0620 []string
	__obf_c2522539269e55c0 int
	__obf_4dd452a5b56cca5f map[string]int32
	__obf_a6bc6d9a9f012ccd map[string][]int
	__obf_672eaa62f3e54f8b int
}

func NewWuManberMatcher(__obf_c2522539269e55c0 int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_b903a599ae1e0620: make([]string, 0),
		__obf_c2522539269e55c0: __obf_c2522539269e55c0,
		__obf_4dd452a5b56cca5f: make(map[string]int32),
		__obf_a6bc6d9a9f012ccd: make(map[string][]int),
		__obf_672eaa62f3e54f8b: int(^uint(0) >> 1),
	}
}

func (__obf_0e74a0f01677fac8 *WuManberMatcher) AddPatterns(__obf_b903a599ae1e0620 []string) {
	__obf_0e74a0f01677fac8.__obf_b903a599ae1e0620 = __obf_b903a599ae1e0620
	for _, __obf_14892a27b66c28f1 := range __obf_b903a599ae1e0620 {
		__obf_6c0dd739b76b0e18 := utf8.RuneCountInString(__obf_14892a27b66c28f1)
		if __obf_6c0dd739b76b0e18 > 0 && __obf_6c0dd739b76b0e18 < __obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b {
			__obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b = __obf_6c0dd739b76b0e18
		}
	}
}

func (__obf_0e74a0f01677fac8 *WuManberMatcher) Build() {
	if len(__obf_0e74a0f01677fac8.__obf_b903a599ae1e0620) == 0 {
		return
	}
	if __obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b < __obf_0e74a0f01677fac8.__obf_c2522539269e55c0 {
		__obf_0e74a0f01677fac8.__obf_c2522539269e55c0 = __obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b
	}
	if __obf_0e74a0f01677fac8.__obf_c2522539269e55c0 <= 0 {
		__obf_0e74a0f01677fac8.__obf_c2522539269e55c0 = 1
	}

	for __obf_1a42297cf5681be2, __obf_078616414ff3fcab := range __obf_0e74a0f01677fac8.__obf_b903a599ae1e0620 {
		__obf_bbebb446a4b21552 := []rune(__obf_078616414ff3fcab)
		__obf_2a0f6089b51a14c0 := len(__obf_bbebb446a4b21552)
		for __obf_a8753e2db05116b7 := 0; __obf_a8753e2db05116b7 <= __obf_2a0f6089b51a14c0-__obf_0e74a0f01677fac8.__obf_c2522539269e55c0; __obf_a8753e2db05116b7++ {
			__obf_63e526fe1cd8adac := string(__obf_bbebb446a4b21552[__obf_a8753e2db05116b7 : __obf_a8753e2db05116b7+__obf_0e74a0f01677fac8.__obf_c2522539269e55c0])
			__obf_d569e3b7db9ba08f := int32(__obf_2a0f6089b51a14c0 - __obf_a8753e2db05116b7 - __obf_0e74a0f01677fac8.__obf_c2522539269e55c0)
			if __obf_815fe8e96bbd942f, __obf_c02a132e507f86d8 := __obf_0e74a0f01677fac8.__obf_4dd452a5b56cca5f[__obf_63e526fe1cd8adac]; !__obf_c02a132e507f86d8 || __obf_d569e3b7db9ba08f < __obf_815fe8e96bbd942f {
				__obf_0e74a0f01677fac8.__obf_4dd452a5b56cca5f[__obf_63e526fe1cd8adac] = __obf_d569e3b7db9ba08f
			}
		}

		if __obf_2a0f6089b51a14c0 >= __obf_0e74a0f01677fac8.__obf_c2522539269e55c0 {
			__obf_900bca6f774a623e := string(__obf_bbebb446a4b21552[__obf_2a0f6089b51a14c0-__obf_0e74a0f01677fac8.__obf_c2522539269e55c0:])
			__obf_0e74a0f01677fac8.__obf_a6bc6d9a9f012ccd[__obf_900bca6f774a623e] = append(__obf_0e74a0f01677fac8.__obf_a6bc6d9a9f012ccd[__obf_900bca6f774a623e], __obf_1a42297cf5681be2)
		}
	}
}

func (__obf_0e74a0f01677fac8 *WuManberMatcher) Match(__obf_c5e81538f7c752f6 string) bool {
	if len(__obf_0e74a0f01677fac8.__obf_b903a599ae1e0620) == 0 {
		return false
	}
	__obf_bbebb446a4b21552 := []rune(__obf_c5e81538f7c752f6)
	__obf_d3ccc758a39fb176 := len(__obf_bbebb446a4b21552)
	if __obf_d3ccc758a39fb176 < __obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b {
		return false
	}
	__obf_62d67604be0b3fe8 := __obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b - 1
	for __obf_62d67604be0b3fe8 < __obf_d3ccc758a39fb176 {
		__obf_57e6925686297e69 := max(__obf_62d67604be0b3fe8-__obf_0e74a0f01677fac8.__obf_c2522539269e55c0+1, 0)
		__obf_63e526fe1cd8adac := string(__obf_bbebb446a4b21552[__obf_57e6925686297e69 : __obf_62d67604be0b3fe8+1])
		__obf_4dd452a5b56cca5f, __obf_c02a132e507f86d8 := __obf_0e74a0f01677fac8.__obf_4dd452a5b56cca5f[__obf_63e526fe1cd8adac]
		if !__obf_c02a132e507f86d8 {
			__obf_62d67604be0b3fe8 += (__obf_0e74a0f01677fac8.__obf_672eaa62f3e54f8b - __obf_0e74a0f01677fac8.__obf_c2522539269e55c0 + 1)
			continue
		}
		if __obf_4dd452a5b56cca5f == 0 {
			if __obf_ec3b1497cd9a9616, __obf_d8b944328567cd86 := __obf_0e74a0f01677fac8.__obf_a6bc6d9a9f012ccd[__obf_63e526fe1cd8adac]; __obf_d8b944328567cd86 {
				for _, __obf_257444100adc6119 := range __obf_ec3b1497cd9a9616 {
					__obf_078616414ff3fcab := __obf_0e74a0f01677fac8.__obf_b903a599ae1e0620[__obf_257444100adc6119]
					__obf_93267bd25ffbc109 := utf8.RuneCountInString(__obf_078616414ff3fcab)
					__obf_e19273dbf5115528 := __obf_62d67604be0b3fe8 - __obf_93267bd25ffbc109 + 1
					if __obf_e19273dbf5115528 >= 0 && string(__obf_bbebb446a4b21552[__obf_e19273dbf5115528:__obf_e19273dbf5115528+__obf_93267bd25ffbc109]) == __obf_078616414ff3fcab {
						return true
					}
				}
			}
			__obf_62d67604be0b3fe8++
		} else {
			__obf_62d67604be0b3fe8 += int(__obf_4dd452a5b56cca5f)
		}
	}
	return false
}
