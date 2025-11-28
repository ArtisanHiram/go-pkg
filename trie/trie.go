package __obf_27b512d76308b7ee

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
	__obf_7a4a385b6bea4800 = 1024
	__obf_b3611601dd0859d7 = 1_000_000 // 1MB
	__obf_7e66a81f0200bf6a = 3
	__obf_352243a431368eb6 = 100
	__obf_faaf03dbc409c3c6 = 16
)

// 位图实现
type __obf_0ca77b826bd5a20e struct {
	__obf_ef068b91db0067d4 []uint64
}

func __obf_54e54e32704ecaec(__obf_078939f09d2ea144 int) *__obf_0ca77b826bd5a20e {
	// 每个 uint64 存储 64 位
	__obf_10030d0844dac628 := (__obf_078939f09d2ea144 + 63) / 64
	return &__obf_0ca77b826bd5a20e{__obf_ef068b91db0067d4: make([]uint64, __obf_10030d0844dac628)}
}

func (__obf_35376aacc995ddfa *__obf_0ca77b826bd5a20e) __obf_cac9f584f044f79d(__obf_4f77dc8fbf7391fb int) {
	__obf_35376aacc995ddfa.__obf_ef068b91db0067d4[__obf_4f77dc8fbf7391fb/64] |= 1 << (__obf_4f77dc8fbf7391fb % 64)
}

func (__obf_35376aacc995ddfa *__obf_0ca77b826bd5a20e) __obf_8f38c952e3a2c86a(__obf_4f77dc8fbf7391fb int) bool {
	return __obf_35376aacc995ddfa.__obf_ef068b91db0067d4[__obf_4f77dc8fbf7391fb/64]&(1<<(__obf_4f77dc8fbf7391fb%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_fbcf66c806431122 *DATTrie
	__obf_283f7257ff4ff5a3 *WuManberMatcher
	__obf_9bfe402713051802 sync.RWMutex
}

func (__obf_090e852b40ac43d8 *SensitiveFilter) Contains(__obf_4337d7361cd777c5 string) bool {
	__obf_090e852b40ac43d8.__obf_9bfe402713051802.Lock()
	defer __obf_090e852b40ac43d8.__obf_9bfe402713051802.Unlock()
	if len(__obf_4337d7361cd777c5) > __obf_352243a431368eb6 {
		if __obf_090e852b40ac43d8.__obf_283f7257ff4ff5a3.Match(__obf_4337d7361cd777c5) {
			return true
		} else {
			return false
		}
	}
	return __obf_090e852b40ac43d8.__obf_fbcf66c806431122.Contains(__obf_4337d7361cd777c5)
}

func (__obf_090e852b40ac43d8 *SensitiveFilter) FindAll(__obf_4337d7361cd777c5 string) []MatchResult {
	if __obf_090e852b40ac43d8 == nil || __obf_090e852b40ac43d8.__obf_fbcf66c806431122 == nil {
		return nil
	}
	if __obf_4337d7361cd777c5 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_4337d7361cd777c5) > __obf_b3611601dd0859d7 {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_e6976629562c79b8 := __obf_090e852b40ac43d8.__obf_fbcf66c806431122.__obf_47d97ae362ed90ab([]rune(__obf_4337d7361cd777c5))
	if len(__obf_e6976629562c79b8) < 2 {
		return __obf_e6976629562c79b8
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_e6976629562c79b8, func(__obf_a656e554f3f4f366, __obf_5dc3a5ef6f3e4c0f int) bool {
		__obf_cf0963390bbfa58c := __obf_e6976629562c79b8[__obf_a656e554f3f4f366].End - __obf_e6976629562c79b8[__obf_a656e554f3f4f366].Start
		__obf_effebb6ab0b7ef8e := __obf_e6976629562c79b8[__obf_5dc3a5ef6f3e4c0f].End - __obf_e6976629562c79b8[__obf_5dc3a5ef6f3e4c0f].Start
		if __obf_cf0963390bbfa58c != __obf_effebb6ab0b7ef8e {
			return __obf_cf0963390bbfa58c > __obf_effebb6ab0b7ef8e
		}
		return __obf_e6976629562c79b8[__obf_a656e554f3f4f366].Start < __obf_e6976629562c79b8[__obf_5dc3a5ef6f3e4c0f].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_ffca087eb4bb6797 := make([]MatchResult, 0, len(__obf_e6976629562c79b8))
	// occupied := make([]bool, len(runes))
	__obf_251efd41c943989c := __obf_54e54e32704ecaec(len([]rune(__obf_4337d7361cd777c5)))

	for _, __obf_007869c26d0f4a4a := range __obf_e6976629562c79b8 {
		__obf_74bf6ce5e637d69e := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_a656e554f3f4f366 := __obf_007869c26d0f4a4a.Start; __obf_a656e554f3f4f366 < __obf_007869c26d0f4a4a.End; __obf_a656e554f3f4f366++ {
			// if occupied[i] {
			if __obf_251efd41c943989c.__obf_8f38c952e3a2c86a(__obf_a656e554f3f4f366) {
				__obf_74bf6ce5e637d69e = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_74bf6ce5e637d69e {
			__obf_ffca087eb4bb6797 = append(__obf_ffca087eb4bb6797, __obf_007869c26d0f4a4a)
			// 标记该位置已被占用
			for __obf_a656e554f3f4f366 := __obf_007869c26d0f4a4a.Start; __obf_a656e554f3f4f366 < __obf_007869c26d0f4a4a.End; __obf_a656e554f3f4f366++ {
				// occupied[i] = true
				__obf_251efd41c943989c.__obf_cac9f584f044f79d(__obf_a656e554f3f4f366)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_ffca087eb4bb6797, func(__obf_a656e554f3f4f366, __obf_5dc3a5ef6f3e4c0f int) bool {
		return __obf_ffca087eb4bb6797[__obf_a656e554f3f4f366].Start < __obf_ffca087eb4bb6797[__obf_5dc3a5ef6f3e4c0f].Start
	})

	return __obf_ffca087eb4bb6797
}

func (__obf_090e852b40ac43d8 *SensitiveFilter) Replace(__obf_4337d7361cd777c5 string, __obf_63d57aff584c9d5f rune) string {
	__obf_a8af48579e90e295 := __obf_090e852b40ac43d8.FindAll(__obf_4337d7361cd777c5)
	if len(__obf_a8af48579e90e295) == 0 {
		return __obf_4337d7361cd777c5
	}

	__obf_95d5040152389175 := []rune(__obf_4337d7361cd777c5)
	for _, __obf_007869c26d0f4a4a := range __obf_a8af48579e90e295 {
		for __obf_a656e554f3f4f366 := __obf_007869c26d0f4a4a.Start; __obf_a656e554f3f4f366 < __obf_007869c26d0f4a4a.End; __obf_a656e554f3f4f366++ {
			__obf_95d5040152389175[__obf_a656e554f3f4f366] = __obf_63d57aff584c9d5f
		}
	}
	return string(__obf_95d5040152389175)
}

func (__obf_090e852b40ac43d8 *SensitiveFilter) ReplaceWithString(__obf_4337d7361cd777c5 string, __obf_63d57aff584c9d5f string) string {
	__obf_a8af48579e90e295 := __obf_090e852b40ac43d8.FindAll(__obf_4337d7361cd777c5)
	if len(__obf_a8af48579e90e295) == 0 {
		return __obf_4337d7361cd777c5
	}

	var __obf_667c5cdf905d1bde strings.Builder
	__obf_95d5040152389175 := []rune(__obf_4337d7361cd777c5)
	__obf_42400ad1e555c0b1 := 0
	for _, __obf_007869c26d0f4a4a := range __obf_a8af48579e90e295 {
		__obf_667c5cdf905d1bde.WriteString(string(__obf_95d5040152389175[__obf_42400ad1e555c0b1:__obf_007869c26d0f4a4a.Start]))
		__obf_667c5cdf905d1bde.WriteString(__obf_63d57aff584c9d5f)
		__obf_42400ad1e555c0b1 = __obf_007869c26d0f4a4a.End
	}
	__obf_667c5cdf905d1bde.WriteString(string(__obf_95d5040152389175[__obf_42400ad1e555c0b1:]))
	return __obf_667c5cdf905d1bde.String()
}

func (__obf_090e852b40ac43d8 *SensitiveFilter) ReplaceWithFunc(__obf_4337d7361cd777c5 string, __obf_0f5657a283c134c8 func(__obf_007869c26d0f4a4a MatchResult) string) string {
	__obf_a8af48579e90e295 := __obf_090e852b40ac43d8.FindAll(__obf_4337d7361cd777c5)
	if len(__obf_a8af48579e90e295) == 0 {
		return __obf_4337d7361cd777c5
	}

	var __obf_667c5cdf905d1bde strings.Builder
	__obf_95d5040152389175 := []rune(__obf_4337d7361cd777c5)
	__obf_42400ad1e555c0b1 := 0
	for _, __obf_007869c26d0f4a4a := range __obf_a8af48579e90e295 {
		__obf_667c5cdf905d1bde.WriteString(string(__obf_95d5040152389175[__obf_42400ad1e555c0b1:__obf_007869c26d0f4a4a.Start]))
		__obf_667c5cdf905d1bde.WriteString(__obf_0f5657a283c134c8(__obf_007869c26d0f4a4a))
		__obf_42400ad1e555c0b1 = __obf_007869c26d0f4a4a.End
	}
	__obf_667c5cdf905d1bde.WriteString(string(__obf_95d5040152389175[__obf_42400ad1e555c0b1:]))
	return __obf_667c5cdf905d1bde.String()
}

// 添加 Close 方法释放资源
func (__obf_090e852b40ac43d8 *SensitiveFilter) Close() {
	__obf_090e852b40ac43d8.__obf_9bfe402713051802.Lock()
	defer __obf_090e852b40ac43d8.__obf_9bfe402713051802.Unlock()

	__obf_090e852b40ac43d8.__obf_fbcf66c806431122 = nil
	__obf_090e852b40ac43d8.__obf_283f7257ff4ff5a3 = nil
}

type FilterBuilder struct {
	__obf_79d9981c4d8614f9 map[string]struct{}
	__obf_15f94d91d34c8114 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_078939f09d2ea144 int) FilterOption {
	return func(__obf_35376aacc995ddfa *FilterBuilder) {
		if __obf_078939f09d2ea144 > 0 {
			__obf_35376aacc995ddfa.__obf_15f94d91d34c8114 = __obf_078939f09d2ea144
		}
	}
}

func WithWords(__obf_79d9981c4d8614f9 []string) FilterOption {
	return func(__obf_35376aacc995ddfa *FilterBuilder) {
		for _, __obf_39ab00e8deaa44bb := range __obf_79d9981c4d8614f9 {
			if __obf_39ab00e8deaa44bb != "" {
				__obf_35376aacc995ddfa.__obf_79d9981c4d8614f9[__obf_39ab00e8deaa44bb] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_96d51eead1f7b1e8 ...FilterOption) *FilterBuilder {
	__obf_35376aacc995ddfa := &FilterBuilder{
		__obf_79d9981c4d8614f9: make(map[string]struct{}),
		__obf_15f94d91d34c8114: __obf_7e66a81f0200bf6a,
	}
	for _, __obf_04dd664bb9ab4ef1 := range __obf_96d51eead1f7b1e8 {
		__obf_04dd664bb9ab4ef1(__obf_35376aacc995ddfa)
	}
	return __obf_35376aacc995ddfa
}

func (__obf_35376aacc995ddfa *FilterBuilder) AddWord(__obf_05f7f13c2062b386 string) {
	if __obf_05f7f13c2062b386 != "" {
		__obf_35376aacc995ddfa.__obf_79d9981c4d8614f9[__obf_05f7f13c2062b386] = struct{}{}
	}
}

func (__obf_35376aacc995ddfa *FilterBuilder) AddWords(__obf_79d9981c4d8614f9 []string) {
	for _, __obf_05f7f13c2062b386 := range __obf_79d9981c4d8614f9 {
		__obf_35376aacc995ddfa.AddWord(__obf_05f7f13c2062b386)
	}
}

func (__obf_35376aacc995ddfa *FilterBuilder) LoadFromFile(__obf_c01568e9af486259 string) error {
	__obf_cc85080e7a30de6d, __obf_aee107d0e187fc7d := os.Open(__obf_c01568e9af486259)
	if __obf_aee107d0e187fc7d != nil {
		return __obf_aee107d0e187fc7d
	}
	defer __obf_cc85080e7a30de6d.Close()

	__obf_c9e517c2097a2069 := bufio.NewReader(__obf_cc85080e7a30de6d)
	__obf_669fd87aeb4e4fe7 := []byte{0xEF, 0xBB, 0xBF}
	__obf_531790e08e710cc6, __obf_aee107d0e187fc7d := __obf_c9e517c2097a2069.Peek(3)
	if __obf_aee107d0e187fc7d == nil && bytes.Equal(__obf_531790e08e710cc6, __obf_669fd87aeb4e4fe7) {
		_, _ = __obf_c9e517c2097a2069.Discard(3)
	}

	__obf_0547ec2aaf984649 := bufio.NewScanner(__obf_c9e517c2097a2069)
	for __obf_0547ec2aaf984649.Scan() {
		__obf_b2dff9deb00a9060 := strings.TrimSpace(__obf_0547ec2aaf984649.Text())
		__obf_35376aacc995ddfa.AddWord(__obf_b2dff9deb00a9060)
	}
	return __obf_0547ec2aaf984649.Err()
}

func (__obf_35376aacc995ddfa *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_35376aacc995ddfa.__obf_79d9981c4d8614f9) == 0 {
		return &SensitiveFilter{__obf_fbcf66c806431122: NewDATTrie(),
			__obf_283f7257ff4ff5a3: NewWuManberMatcher(__obf_35376aacc995ddfa.__obf_15f94d91d34c8114)}
	}

	__obf_dd217062cd329bcf := &__obf_11d9e315e7f74407{__obf_3133c6cb010f2426: make(map[rune]*__obf_11d9e315e7f74407)}
	__obf_b4e6724f735d7909 := make([]string, 0, len(__obf_35376aacc995ddfa.__obf_79d9981c4d8614f9))
	for __obf_05f7f13c2062b386 := range __obf_35376aacc995ddfa.__obf_79d9981c4d8614f9 {
		__obf_b4e6724f735d7909 = append(__obf_b4e6724f735d7909, __obf_05f7f13c2062b386)
		__obf_dd217062cd329bcf.__obf_47ec6843063d10aa(__obf_05f7f13c2062b386)
	}

	__obf_fbcf66c806431122 := __obf_fbd49da586817452(__obf_dd217062cd329bcf)
	__obf_283f7257ff4ff5a3 := NewWuManberMatcher(__obf_35376aacc995ddfa.__obf_15f94d91d34c8114)
	__obf_283f7257ff4ff5a3.AddPatterns(__obf_b4e6724f735d7909)
	__obf_283f7257ff4ff5a3.Build()

	return &SensitiveFilter{
		__obf_fbcf66c806431122: __obf_fbcf66c806431122,
		__obf_283f7257ff4ff5a3: __obf_283f7257ff4ff5a3,
	}
}

func (__obf_35376aacc995ddfa *FilterBuilder) Clear() {
	__obf_35376aacc995ddfa.__obf_79d9981c4d8614f9 = make(map[string]struct{})
}

type __obf_11d9e315e7f74407 struct {
	__obf_3133c6cb010f2426 map[rune]*__obf_11d9e315e7f74407
	__obf_59358b6ba2654ab9 bool
}

func (__obf_10030d0844dac628 *__obf_11d9e315e7f74407) __obf_47ec6843063d10aa(__obf_05f7f13c2062b386 string) {
	__obf_e13b9ec989e92579 := __obf_10030d0844dac628
	for _, __obf_44483a7b936b2483 := range __obf_05f7f13c2062b386 {
		if _, __obf_d2c35e349c65d172 := __obf_e13b9ec989e92579.__obf_3133c6cb010f2426[__obf_44483a7b936b2483]; !__obf_d2c35e349c65d172 {
			__obf_e13b9ec989e92579.__obf_3133c6cb010f2426[__obf_44483a7b936b2483] = &__obf_11d9e315e7f74407{__obf_3133c6cb010f2426: make(map[rune]*__obf_11d9e315e7f74407)}
		}
		__obf_e13b9ec989e92579 = __obf_e13b9ec989e92579.__obf_3133c6cb010f2426[__obf_44483a7b936b2483]
	}
	__obf_e13b9ec989e92579.__obf_59358b6ba2654ab9 = true
}

type DATTrie struct {
	__obf_38c9bbfc8092cbf6 []int32
	__obf_75934e10b6a857c0 []int32
	__obf_8865fd29a1e4d86e []bool
	__obf_078939f09d2ea144 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_38c9bbfc8092cbf6: make([]int32, __obf_7a4a385b6bea4800),
		__obf_75934e10b6a857c0: make([]int32, __obf_7a4a385b6bea4800),
		__obf_8865fd29a1e4d86e: make([]bool, __obf_7a4a385b6bea4800),
		__obf_078939f09d2ea144: 1,
	}
}

func __obf_fbd49da586817452(__obf_c536eb162be1e2d4 *__obf_11d9e315e7f74407) *DATTrie {
	__obf_070f2bd7c6dd1c51 := NewDATTrie()
	__obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0[1] = 1

	type __obf_b4bf807e81cb1b48 struct {
		__obf_9832904a420eab16 int32
		__obf_e13b9ec989e92579 *__obf_11d9e315e7f74407
	}

	__obf_5348b3b82f95cc08 := []__obf_b4bf807e81cb1b48{{__obf_9832904a420eab16: 1, __obf_e13b9ec989e92579: __obf_c536eb162be1e2d4}}

	for len(__obf_5348b3b82f95cc08) > 0 {
		__obf_a94485d3c1dafdf6 := __obf_5348b3b82f95cc08[0]
		__obf_5348b3b82f95cc08 = __obf_5348b3b82f95cc08[1:]

		if len(__obf_a94485d3c1dafdf6.__obf_e13b9ec989e92579.__obf_3133c6cb010f2426) == 0 {
			continue
		}

		__obf_4037e8efe49a1ba6 := make([]rune, 0, len(__obf_a94485d3c1dafdf6.__obf_e13b9ec989e92579.__obf_3133c6cb010f2426))
		for __obf_44483a7b936b2483 := range __obf_a94485d3c1dafdf6.__obf_e13b9ec989e92579.__obf_3133c6cb010f2426 {
			__obf_4037e8efe49a1ba6 = append(__obf_4037e8efe49a1ba6, __obf_44483a7b936b2483)
		}
		slices.Sort(__obf_4037e8efe49a1ba6)

		__obf_804c4af7c382199b := __obf_070f2bd7c6dd1c51.__obf_a7a1ed47c30e92c4(__obf_4037e8efe49a1ba6)
		__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6[__obf_a94485d3c1dafdf6.__obf_9832904a420eab16] = __obf_804c4af7c382199b

		for _, __obf_6718c9b4dcf47ced := range __obf_4037e8efe49a1ba6 {
			__obf_de1822cc2a9637dd := __obf_a94485d3c1dafdf6.__obf_e13b9ec989e92579.__obf_3133c6cb010f2426[__obf_6718c9b4dcf47ced]
			__obf_c6df9042c919d27b := __obf_804c4af7c382199b + int32(__obf_6718c9b4dcf47ced)
			__obf_070f2bd7c6dd1c51.__obf_39c5c42c4009afd3(__obf_c6df9042c919d27b + 1)

			__obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0[__obf_c6df9042c919d27b] = __obf_a94485d3c1dafdf6.__obf_9832904a420eab16
			if __obf_de1822cc2a9637dd.__obf_59358b6ba2654ab9 {
				__obf_070f2bd7c6dd1c51.__obf_8865fd29a1e4d86e[__obf_c6df9042c919d27b] = true
			}
			if int32(len(__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6)) <= __obf_c6df9042c919d27b {
				__obf_070f2bd7c6dd1c51.__obf_39c5c42c4009afd3(__obf_c6df9042c919d27b + 1)
			}
			if __obf_070f2bd7c6dd1c51.__obf_078939f09d2ea144 <= __obf_c6df9042c919d27b {
				__obf_070f2bd7c6dd1c51.__obf_078939f09d2ea144 = __obf_c6df9042c919d27b + 1
			}

			__obf_5348b3b82f95cc08 = append(__obf_5348b3b82f95cc08, __obf_b4bf807e81cb1b48{__obf_9832904a420eab16: __obf_c6df9042c919d27b, __obf_e13b9ec989e92579: __obf_de1822cc2a9637dd})
		}
	}
	return __obf_070f2bd7c6dd1c51
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) __obf_a7a1ed47c30e92c4(__obf_9c4b4db8ab7430da []rune) int32 {
	for __obf_38c9bbfc8092cbf6 := int32(1); ; __obf_38c9bbfc8092cbf6++ {
		__obf_3bb4fac0e75bd5a4 := false
		for _, __obf_0e9696c1a8ca7e07 := range __obf_9c4b4db8ab7430da {
			__obf_c1eb53bb86fab477 := __obf_38c9bbfc8092cbf6 + int32(__obf_0e9696c1a8ca7e07)
			if int32(len(__obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0)) <= __obf_c1eb53bb86fab477 {
				__obf_070f2bd7c6dd1c51.__obf_39c5c42c4009afd3(__obf_c1eb53bb86fab477 + 1)
			}
			if __obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0[__obf_c1eb53bb86fab477] != 0 {
				__obf_3bb4fac0e75bd5a4 = true
				break
			}
		}
		if !__obf_3bb4fac0e75bd5a4 {
			return __obf_38c9bbfc8092cbf6
		}
	}
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) __obf_39c5c42c4009afd3(__obf_2d3f161625c97563 int32) {
	if __obf_2d3f161625c97563 < int32(len(__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6)) {
		return
	}
	__obf_dae0dc4ff4183480 := max(int32(len(__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6))+int32(len(__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6))/2, __obf_2d3f161625c97563)
	__obf_0953450419cec18c := make([]int32, __obf_dae0dc4ff4183480)
	__obf_133d676f0332ea46 := make([]int32, __obf_dae0dc4ff4183480)
	__obf_fb3ea5baca23c6ab := make([]bool, __obf_dae0dc4ff4183480)
	copy(__obf_0953450419cec18c, __obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6)
	copy(__obf_133d676f0332ea46, __obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0)
	copy(__obf_fb3ea5baca23c6ab, __obf_070f2bd7c6dd1c51.__obf_8865fd29a1e4d86e)
	__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6 = __obf_0953450419cec18c
	__obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0 = __obf_133d676f0332ea46
	__obf_070f2bd7c6dd1c51.__obf_8865fd29a1e4d86e = __obf_fb3ea5baca23c6ab
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) __obf_7800f73b79889e11(__obf_3f7e49c728a5fe6e int32, __obf_0e9696c1a8ca7e07 rune) int32 {
	if __obf_3f7e49c728a5fe6e <= 0 || __obf_3f7e49c728a5fe6e >= int32(len(__obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6)) || __obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6[__obf_3f7e49c728a5fe6e] == 0 {
		return -1
	}
	__obf_c1eb53bb86fab477 := __obf_070f2bd7c6dd1c51.__obf_38c9bbfc8092cbf6[__obf_3f7e49c728a5fe6e] + int32(__obf_0e9696c1a8ca7e07)
	if __obf_c1eb53bb86fab477 > 0 && __obf_c1eb53bb86fab477 < int32(len(__obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0)) && __obf_070f2bd7c6dd1c51.__obf_75934e10b6a857c0[__obf_c1eb53bb86fab477] == __obf_3f7e49c728a5fe6e {
		return __obf_c1eb53bb86fab477
	}
	return -1
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) __obf_59358b6ba2654ab9(__obf_3f7e49c728a5fe6e int32) bool {
	if __obf_3f7e49c728a5fe6e > 0 && __obf_3f7e49c728a5fe6e < int32(len(__obf_070f2bd7c6dd1c51.__obf_8865fd29a1e4d86e)) {
		return __obf_070f2bd7c6dd1c51.__obf_8865fd29a1e4d86e[__obf_3f7e49c728a5fe6e]
	}
	return false
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) Contains(__obf_4337d7361cd777c5 string) bool {
	__obf_95d5040152389175 := []rune(__obf_4337d7361cd777c5)
	for __obf_a656e554f3f4f366 := range __obf_95d5040152389175 {
		__obf_e13b9ec989e92579 := int32(1)
		for __obf_5dc3a5ef6f3e4c0f := __obf_a656e554f3f4f366; __obf_5dc3a5ef6f3e4c0f < len(__obf_95d5040152389175); __obf_5dc3a5ef6f3e4c0f++ {
			__obf_e13b9ec989e92579 = __obf_070f2bd7c6dd1c51.__obf_7800f73b79889e11(__obf_e13b9ec989e92579, __obf_95d5040152389175[__obf_5dc3a5ef6f3e4c0f])
			if __obf_e13b9ec989e92579 == -1 {
				break
			}
			if __obf_070f2bd7c6dd1c51.__obf_59358b6ba2654ab9(__obf_e13b9ec989e92579) {
				return true
			}
		}
	}
	return false
}

func (__obf_070f2bd7c6dd1c51 *DATTrie) __obf_47d97ae362ed90ab(__obf_95d5040152389175 []rune) []MatchResult {
	__obf_b7ae09fef9928452 := make([]MatchResult, 0, __obf_faaf03dbc409c3c6)
	for __obf_a656e554f3f4f366 := range __obf_95d5040152389175 {
		__obf_e13b9ec989e92579 := int32(1)
		for __obf_5dc3a5ef6f3e4c0f := __obf_a656e554f3f4f366; __obf_5dc3a5ef6f3e4c0f < len(__obf_95d5040152389175); __obf_5dc3a5ef6f3e4c0f++ {
			__obf_e13b9ec989e92579 = __obf_070f2bd7c6dd1c51.__obf_7800f73b79889e11(__obf_e13b9ec989e92579, __obf_95d5040152389175[__obf_5dc3a5ef6f3e4c0f])
			if __obf_e13b9ec989e92579 == -1 {
				break
			}
			if __obf_070f2bd7c6dd1c51.__obf_59358b6ba2654ab9(__obf_e13b9ec989e92579) {
				__obf_b7ae09fef9928452 = append(__obf_b7ae09fef9928452, MatchResult{
					Word:  string(__obf_95d5040152389175[__obf_a656e554f3f4f366 : __obf_5dc3a5ef6f3e4c0f+1]),
					Start: __obf_a656e554f3f4f366,
					End:   __obf_5dc3a5ef6f3e4c0f + 1,
				})
			}
		}
	}
	return __obf_b7ae09fef9928452
}

type WuManberMatcher struct {
	__obf_122e9615dad4302a []string
	__obf_15f94d91d34c8114 int
	__obf_262135b3b25b4aa8 map[string]int32
	__obf_e048d8cc62b9d318 map[string][]int
	__obf_740d4f16579c50b7 int
}

func NewWuManberMatcher(__obf_15f94d91d34c8114 int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_122e9615dad4302a: make([]string, 0),
		__obf_15f94d91d34c8114: __obf_15f94d91d34c8114,
		__obf_262135b3b25b4aa8: make(map[string]int32),
		__obf_e048d8cc62b9d318: make(map[string][]int),
		__obf_740d4f16579c50b7: int(^uint(0) >> 1),
	}
}

func (__obf_61b9b6346ec1b42e *WuManberMatcher) AddPatterns(__obf_122e9615dad4302a []string) {
	__obf_61b9b6346ec1b42e.__obf_122e9615dad4302a = __obf_122e9615dad4302a
	for _, __obf_549f793c017c43f3 := range __obf_122e9615dad4302a {
		__obf_b653791e5f952a32 := utf8.RuneCountInString(__obf_549f793c017c43f3)
		if __obf_b653791e5f952a32 > 0 && __obf_b653791e5f952a32 < __obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 {
			__obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 = __obf_b653791e5f952a32
		}
	}
}

func (__obf_61b9b6346ec1b42e *WuManberMatcher) Build() {
	if len(__obf_61b9b6346ec1b42e.__obf_122e9615dad4302a) == 0 {
		return
	}
	if __obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 < __obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 {
		__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 = __obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7
	}
	if __obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 <= 0 {
		__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 = 1
	}

	for __obf_a656e554f3f4f366, __obf_ce56c5ea843d5ac6 := range __obf_61b9b6346ec1b42e.__obf_122e9615dad4302a {
		__obf_95d5040152389175 := []rune(__obf_ce56c5ea843d5ac6)
		__obf_b721a91ef6ca7878 := len(__obf_95d5040152389175)
		for __obf_5dc3a5ef6f3e4c0f := 0; __obf_5dc3a5ef6f3e4c0f <= __obf_b721a91ef6ca7878-__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114; __obf_5dc3a5ef6f3e4c0f++ {
			__obf_f277d147a87e6d02 := string(__obf_95d5040152389175[__obf_5dc3a5ef6f3e4c0f : __obf_5dc3a5ef6f3e4c0f+__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114])
			__obf_8a8a9ebfb37e65a4 := int32(__obf_b721a91ef6ca7878 - __obf_5dc3a5ef6f3e4c0f - __obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114)
			if __obf_cfb028c58b9eb28b, __obf_046112dc3284f34d := __obf_61b9b6346ec1b42e.__obf_262135b3b25b4aa8[__obf_f277d147a87e6d02]; !__obf_046112dc3284f34d || __obf_8a8a9ebfb37e65a4 < __obf_cfb028c58b9eb28b {
				__obf_61b9b6346ec1b42e.__obf_262135b3b25b4aa8[__obf_f277d147a87e6d02] = __obf_8a8a9ebfb37e65a4
			}
		}

		if __obf_b721a91ef6ca7878 >= __obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 {
			__obf_ef842c3af48abaf9 := string(__obf_95d5040152389175[__obf_b721a91ef6ca7878-__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114:])
			__obf_61b9b6346ec1b42e.__obf_e048d8cc62b9d318[__obf_ef842c3af48abaf9] = append(__obf_61b9b6346ec1b42e.__obf_e048d8cc62b9d318[__obf_ef842c3af48abaf9], __obf_a656e554f3f4f366)
		}
	}
}

func (__obf_61b9b6346ec1b42e *WuManberMatcher) Match(__obf_4337d7361cd777c5 string) bool {
	if len(__obf_61b9b6346ec1b42e.__obf_122e9615dad4302a) == 0 {
		return false
	}
	__obf_95d5040152389175 := []rune(__obf_4337d7361cd777c5)
	__obf_916cbcc0e88e4292 := len(__obf_95d5040152389175)
	if __obf_916cbcc0e88e4292 < __obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 {
		return false
	}
	__obf_4f77dc8fbf7391fb := __obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 - 1
	for __obf_4f77dc8fbf7391fb < __obf_916cbcc0e88e4292 {
		__obf_4c5b1413d25c662f := max(__obf_4f77dc8fbf7391fb-__obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114+1, 0)
		__obf_f277d147a87e6d02 := string(__obf_95d5040152389175[__obf_4c5b1413d25c662f : __obf_4f77dc8fbf7391fb+1])
		__obf_262135b3b25b4aa8, __obf_046112dc3284f34d := __obf_61b9b6346ec1b42e.__obf_262135b3b25b4aa8[__obf_f277d147a87e6d02]
		if !__obf_046112dc3284f34d {
			__obf_4f77dc8fbf7391fb += (__obf_61b9b6346ec1b42e.__obf_740d4f16579c50b7 - __obf_61b9b6346ec1b42e.__obf_15f94d91d34c8114 + 1)
			continue
		}
		if __obf_262135b3b25b4aa8 == 0 {
			if __obf_b0a7532d44c337aa, __obf_d2c35e349c65d172 := __obf_61b9b6346ec1b42e.__obf_e048d8cc62b9d318[__obf_f277d147a87e6d02]; __obf_d2c35e349c65d172 {
				for _, __obf_bb1750ba28979a1e := range __obf_b0a7532d44c337aa {
					__obf_ce56c5ea843d5ac6 := __obf_61b9b6346ec1b42e.__obf_122e9615dad4302a[__obf_bb1750ba28979a1e]
					__obf_598460ef7654520e := utf8.RuneCountInString(__obf_ce56c5ea843d5ac6)
					__obf_2c206e94ed2017d4 := __obf_4f77dc8fbf7391fb - __obf_598460ef7654520e + 1
					if __obf_2c206e94ed2017d4 >= 0 && string(__obf_95d5040152389175[__obf_2c206e94ed2017d4:__obf_2c206e94ed2017d4+__obf_598460ef7654520e]) == __obf_ce56c5ea843d5ac6 {
						return true
					}
				}
			}
			__obf_4f77dc8fbf7391fb++
		} else {
			__obf_4f77dc8fbf7391fb += int(__obf_262135b3b25b4aa8)
		}
	}
	return false
}
