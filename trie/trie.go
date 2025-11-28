package __obf_b365f1baf6b057e4

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
	__obf_8dcbf981c918df95 = 1024
	__obf_9d27ed3184dfd2a9 = 1_000_000 // 1MB
	__obf_5b418c06239c49a5 = 3
	__obf_c1e34bbf9f86df73 = 100
	__obf_1c44ebf5d8dd34a9 = 16
)

// 位图实现
type __obf_d5bcc986c9bd71b7 struct {
	__obf_243a3c4b0da4e389 []uint64
}

func __obf_222761502a59f547(__obf_01f4e78b15b572f5 int) *__obf_d5bcc986c9bd71b7 {
	// 每个 uint64 存储 64 位
	__obf_30c9ef6807c898ec := (__obf_01f4e78b15b572f5 + 63) / 64
	return &__obf_d5bcc986c9bd71b7{__obf_243a3c4b0da4e389: make([]uint64, __obf_30c9ef6807c898ec)}
}

func (__obf_4be7f978fb1fb712 *__obf_d5bcc986c9bd71b7) __obf_7a4d6bbc88af0ec4(__obf_5229e703e575004b int) {
	__obf_4be7f978fb1fb712.__obf_243a3c4b0da4e389[__obf_5229e703e575004b/64] |= 1 << (__obf_5229e703e575004b % 64)
}

func (__obf_4be7f978fb1fb712 *__obf_d5bcc986c9bd71b7) __obf_1a675aa0a0b3bdb9(__obf_5229e703e575004b int) bool {
	return __obf_4be7f978fb1fb712.__obf_243a3c4b0da4e389[__obf_5229e703e575004b/64]&(1<<(__obf_5229e703e575004b%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_96af5de324f3c6dc *DATTrie
	__obf_f15098afe629cf43 *WuManberMatcher
	__obf_8b3e01fac045f562 sync.RWMutex
}

func (__obf_3d61186146037b37 *SensitiveFilter) Contains(__obf_100ac6ab033077ae string) bool {
	__obf_3d61186146037b37.__obf_8b3e01fac045f562.Lock()
	defer __obf_3d61186146037b37.__obf_8b3e01fac045f562.Unlock()
	if len(__obf_100ac6ab033077ae) > __obf_c1e34bbf9f86df73 {
		if __obf_3d61186146037b37.__obf_f15098afe629cf43.Match(__obf_100ac6ab033077ae) {
			return true
		} else {
			return false
		}
	}
	return __obf_3d61186146037b37.__obf_96af5de324f3c6dc.Contains(__obf_100ac6ab033077ae)
}

func (__obf_3d61186146037b37 *SensitiveFilter) FindAll(__obf_100ac6ab033077ae string) []MatchResult {
	if __obf_3d61186146037b37 == nil || __obf_3d61186146037b37.__obf_96af5de324f3c6dc == nil {
		return nil
	}
	if __obf_100ac6ab033077ae == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_100ac6ab033077ae) > __obf_9d27ed3184dfd2a9 {
		return nil // 或返回错误
	}
	// 1. 找出所有可能重叠的匹配项
	__obf_12d8086fd2247eda := __obf_3d61186146037b37.__obf_96af5de324f3c6dc.__obf_da1d7f0396ff91f7([]rune(__obf_100ac6ab033077ae))
	if len(__obf_12d8086fd2247eda) < 2 {
		return __obf_12d8086fd2247eda
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_12d8086fd2247eda, func(__obf_481c1e102be52ccd, __obf_547930925e67b72f int) bool {
		__obf_95bde51e2feb0528 := __obf_12d8086fd2247eda[__obf_481c1e102be52ccd].End - __obf_12d8086fd2247eda[__obf_481c1e102be52ccd].Start
		__obf_bb4a1a4022140e95 := __obf_12d8086fd2247eda[__obf_547930925e67b72f].End - __obf_12d8086fd2247eda[__obf_547930925e67b72f].Start
		if __obf_95bde51e2feb0528 != __obf_bb4a1a4022140e95 {
			return __obf_95bde51e2feb0528 > __obf_bb4a1a4022140e95
		}
		return __obf_12d8086fd2247eda[__obf_481c1e102be52ccd].Start < __obf_12d8086fd2247eda[__obf_547930925e67b72f].Start
	})

	// 3. 使用贪心算法，选择不重叠的最长匹配
	__obf_6dfe391eb6f61283 := make([]MatchResult, 0, len(__obf_12d8086fd2247eda))
	// occupied := make([]bool, len(runes))
	__obf_1b86f25be274c828 := __obf_222761502a59f547(len([]rune(__obf_100ac6ab033077ae)))

	for _, __obf_93e32ba1363e21d7 := range __obf_12d8086fd2247eda {
		__obf_2e777508e5dec5bf := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_481c1e102be52ccd := __obf_93e32ba1363e21d7.Start; __obf_481c1e102be52ccd < __obf_93e32ba1363e21d7.End; __obf_481c1e102be52ccd++ {
			// if occupied[i] {
			if __obf_1b86f25be274c828.__obf_1a675aa0a0b3bdb9(__obf_481c1e102be52ccd) {
				__obf_2e777508e5dec5bf = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_2e777508e5dec5bf {
			__obf_6dfe391eb6f61283 = append(__obf_6dfe391eb6f61283, __obf_93e32ba1363e21d7)
			// 标记该位置已被占用
			for __obf_481c1e102be52ccd := __obf_93e32ba1363e21d7.Start; __obf_481c1e102be52ccd < __obf_93e32ba1363e21d7.End; __obf_481c1e102be52ccd++ {
				// occupied[i] = true
				__obf_1b86f25be274c828.__obf_7a4d6bbc88af0ec4(__obf_481c1e102be52ccd)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_6dfe391eb6f61283, func(__obf_481c1e102be52ccd, __obf_547930925e67b72f int) bool {
		return __obf_6dfe391eb6f61283[__obf_481c1e102be52ccd].Start < __obf_6dfe391eb6f61283[__obf_547930925e67b72f].Start
	})

	return __obf_6dfe391eb6f61283
}

func (__obf_3d61186146037b37 *SensitiveFilter) Replace(__obf_100ac6ab033077ae string, __obf_0f67480b5247e36c rune) string {
	__obf_b25a126b5b80f378 := __obf_3d61186146037b37.FindAll(__obf_100ac6ab033077ae)
	if len(__obf_b25a126b5b80f378) == 0 {
		return __obf_100ac6ab033077ae
	}

	__obf_0dc92fd760b3ad1d := []rune(__obf_100ac6ab033077ae)
	for _, __obf_93e32ba1363e21d7 := range __obf_b25a126b5b80f378 {
		for __obf_481c1e102be52ccd := __obf_93e32ba1363e21d7.Start; __obf_481c1e102be52ccd < __obf_93e32ba1363e21d7.End; __obf_481c1e102be52ccd++ {
			__obf_0dc92fd760b3ad1d[__obf_481c1e102be52ccd] = __obf_0f67480b5247e36c
		}
	}
	return string(__obf_0dc92fd760b3ad1d)
}

func (__obf_3d61186146037b37 *SensitiveFilter) ReplaceWithString(__obf_100ac6ab033077ae string, __obf_0f67480b5247e36c string) string {
	__obf_b25a126b5b80f378 := __obf_3d61186146037b37.FindAll(__obf_100ac6ab033077ae)
	if len(__obf_b25a126b5b80f378) == 0 {
		return __obf_100ac6ab033077ae
	}

	var __obf_42542da3d236005b strings.Builder
	__obf_0dc92fd760b3ad1d := []rune(__obf_100ac6ab033077ae)
	__obf_80144dbdac075282 := 0
	for _, __obf_93e32ba1363e21d7 := range __obf_b25a126b5b80f378 {
		__obf_42542da3d236005b.WriteString(string(__obf_0dc92fd760b3ad1d[__obf_80144dbdac075282:__obf_93e32ba1363e21d7.Start]))
		__obf_42542da3d236005b.WriteString(__obf_0f67480b5247e36c)
		__obf_80144dbdac075282 = __obf_93e32ba1363e21d7.End
	}
	__obf_42542da3d236005b.WriteString(string(__obf_0dc92fd760b3ad1d[__obf_80144dbdac075282:]))
	return __obf_42542da3d236005b.String()
}

func (__obf_3d61186146037b37 *SensitiveFilter) ReplaceWithFunc(__obf_100ac6ab033077ae string, __obf_bbae8100a6331ef3 func(__obf_93e32ba1363e21d7 MatchResult) string) string {
	__obf_b25a126b5b80f378 := __obf_3d61186146037b37.FindAll(__obf_100ac6ab033077ae)
	if len(__obf_b25a126b5b80f378) == 0 {
		return __obf_100ac6ab033077ae
	}

	var __obf_42542da3d236005b strings.Builder
	__obf_0dc92fd760b3ad1d := []rune(__obf_100ac6ab033077ae)
	__obf_80144dbdac075282 := 0
	for _, __obf_93e32ba1363e21d7 := range __obf_b25a126b5b80f378 {
		__obf_42542da3d236005b.WriteString(string(__obf_0dc92fd760b3ad1d[__obf_80144dbdac075282:__obf_93e32ba1363e21d7.Start]))
		__obf_42542da3d236005b.WriteString(__obf_bbae8100a6331ef3(__obf_93e32ba1363e21d7))
		__obf_80144dbdac075282 = __obf_93e32ba1363e21d7.End
	}
	__obf_42542da3d236005b.WriteString(string(__obf_0dc92fd760b3ad1d[__obf_80144dbdac075282:]))
	return __obf_42542da3d236005b.String()
}

// 添加 Close 方法释放资源
func (__obf_3d61186146037b37 *SensitiveFilter) Close() {
	__obf_3d61186146037b37.__obf_8b3e01fac045f562.Lock()
	defer __obf_3d61186146037b37.__obf_8b3e01fac045f562.Unlock()

	__obf_3d61186146037b37.__obf_96af5de324f3c6dc = nil
	__obf_3d61186146037b37.__obf_f15098afe629cf43 = nil
}

type FilterBuilder struct {
	__obf_bc064d3874e98403 map[string]struct{}
	__obf_8c1c46cc1eb8fdcf int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_01f4e78b15b572f5 int) FilterOption {
	return func(__obf_4be7f978fb1fb712 *FilterBuilder) {
		if __obf_01f4e78b15b572f5 > 0 {
			__obf_4be7f978fb1fb712.__obf_8c1c46cc1eb8fdcf = __obf_01f4e78b15b572f5
		}
	}
}

func WithWords(__obf_bc064d3874e98403 []string) FilterOption {
	return func(__obf_4be7f978fb1fb712 *FilterBuilder) {
		for _, __obf_8a0b1a85399b11a3 := range __obf_bc064d3874e98403 {
			if __obf_8a0b1a85399b11a3 != "" {
				__obf_4be7f978fb1fb712.__obf_bc064d3874e98403[__obf_8a0b1a85399b11a3] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_952ac3fa780b16f5 ...FilterOption) *FilterBuilder {
	__obf_4be7f978fb1fb712 := &FilterBuilder{
		__obf_bc064d3874e98403: make(map[string]struct{}),
		__obf_8c1c46cc1eb8fdcf: __obf_5b418c06239c49a5,
	}
	for _, __obf_8baad3d5073dd1a3 := range __obf_952ac3fa780b16f5 {
		__obf_8baad3d5073dd1a3(__obf_4be7f978fb1fb712)
	}
	return __obf_4be7f978fb1fb712
}

func (__obf_4be7f978fb1fb712 *FilterBuilder) AddWord(__obf_7cc62ee662820472 string) {
	if __obf_7cc62ee662820472 != "" {
		__obf_4be7f978fb1fb712.__obf_bc064d3874e98403[__obf_7cc62ee662820472] = struct{}{}
	}
}

func (__obf_4be7f978fb1fb712 *FilterBuilder) AddWords(__obf_bc064d3874e98403 []string) {
	for _, __obf_7cc62ee662820472 := range __obf_bc064d3874e98403 {
		__obf_4be7f978fb1fb712.AddWord(__obf_7cc62ee662820472)
	}
}

func (__obf_4be7f978fb1fb712 *FilterBuilder) LoadFromFile(__obf_8c362f1988d3ac0e string) error {
	__obf_e8826377b278cc14, __obf_500a8f2182664af3 := os.Open(__obf_8c362f1988d3ac0e)
	if __obf_500a8f2182664af3 != nil {
		return __obf_500a8f2182664af3
	}
	defer __obf_e8826377b278cc14.Close()

	__obf_796b0bdcd4b22a84 := bufio.NewReader(__obf_e8826377b278cc14)
	__obf_5282e3ef5b62d177 := []byte{0xEF, 0xBB, 0xBF}
	__obf_37ad3428efd43982, __obf_500a8f2182664af3 := __obf_796b0bdcd4b22a84.Peek(3)
	if __obf_500a8f2182664af3 == nil && bytes.Equal(__obf_37ad3428efd43982, __obf_5282e3ef5b62d177) {
		_, _ = __obf_796b0bdcd4b22a84.Discard(3)
	}

	__obf_072c78c8acfe4d89 := bufio.NewScanner(__obf_796b0bdcd4b22a84)
	for __obf_072c78c8acfe4d89.Scan() {
		__obf_276a5ed2207b0fa0 := strings.TrimSpace(__obf_072c78c8acfe4d89.Text())
		__obf_4be7f978fb1fb712.AddWord(__obf_276a5ed2207b0fa0)
	}
	return __obf_072c78c8acfe4d89.Err()
}

func (__obf_4be7f978fb1fb712 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_4be7f978fb1fb712.__obf_bc064d3874e98403) == 0 {
		return &SensitiveFilter{__obf_96af5de324f3c6dc: NewDATTrie(),
			__obf_f15098afe629cf43: NewWuManberMatcher(__obf_4be7f978fb1fb712.__obf_8c1c46cc1eb8fdcf)}
	}

	__obf_4a7754d3a37fdc97 := &__obf_6d235e2ecffc87cb{__obf_564314fc8e971646: make(map[rune]*__obf_6d235e2ecffc87cb)}
	__obf_162c1845203971c3 := make([]string, 0, len(__obf_4be7f978fb1fb712.__obf_bc064d3874e98403))
	for __obf_7cc62ee662820472 := range __obf_4be7f978fb1fb712.__obf_bc064d3874e98403 {
		__obf_162c1845203971c3 = append(__obf_162c1845203971c3, __obf_7cc62ee662820472)
		__obf_4a7754d3a37fdc97.__obf_89d6e1fbd3e6fcff(__obf_7cc62ee662820472)
	}

	__obf_96af5de324f3c6dc := __obf_23293f431ef8715b(__obf_4a7754d3a37fdc97)
	__obf_f15098afe629cf43 := NewWuManberMatcher(__obf_4be7f978fb1fb712.__obf_8c1c46cc1eb8fdcf)
	__obf_f15098afe629cf43.AddPatterns(__obf_162c1845203971c3)
	__obf_f15098afe629cf43.Build()

	return &SensitiveFilter{
		__obf_96af5de324f3c6dc: __obf_96af5de324f3c6dc,
		__obf_f15098afe629cf43: __obf_f15098afe629cf43,
	}
}

func (__obf_4be7f978fb1fb712 *FilterBuilder) Clear() {
	__obf_4be7f978fb1fb712.__obf_bc064d3874e98403 = make(map[string]struct{})
}

type __obf_6d235e2ecffc87cb struct {
	__obf_564314fc8e971646 map[rune]*__obf_6d235e2ecffc87cb
	__obf_6778034c36c7f7c8 bool
}

func (__obf_30c9ef6807c898ec *__obf_6d235e2ecffc87cb) __obf_89d6e1fbd3e6fcff(__obf_7cc62ee662820472 string) {
	__obf_bd61e0edf4e82645 := __obf_30c9ef6807c898ec
	for _, __obf_582cf66952790466 := range __obf_7cc62ee662820472 {
		if _, __obf_1256bcce11518b0c := __obf_bd61e0edf4e82645.__obf_564314fc8e971646[__obf_582cf66952790466]; !__obf_1256bcce11518b0c {
			__obf_bd61e0edf4e82645.__obf_564314fc8e971646[__obf_582cf66952790466] = &__obf_6d235e2ecffc87cb{__obf_564314fc8e971646: make(map[rune]*__obf_6d235e2ecffc87cb)}
		}
		__obf_bd61e0edf4e82645 = __obf_bd61e0edf4e82645.__obf_564314fc8e971646[__obf_582cf66952790466]
	}
	__obf_bd61e0edf4e82645.__obf_6778034c36c7f7c8 = true
}

type DATTrie struct {
	__obf_22ac97ae6458a322 []int32
	__obf_637cf2049e4fbf69 []int32
	__obf_9bc0c5da352e4e2b []bool
	__obf_01f4e78b15b572f5 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{
		__obf_22ac97ae6458a322: make([]int32, __obf_8dcbf981c918df95),
		__obf_637cf2049e4fbf69: make([]int32, __obf_8dcbf981c918df95),
		__obf_9bc0c5da352e4e2b: make([]bool, __obf_8dcbf981c918df95),
		__obf_01f4e78b15b572f5: 1,
	}
}

func __obf_23293f431ef8715b(__obf_39fe5d17818af665 *__obf_6d235e2ecffc87cb) *DATTrie {
	__obf_59c627d756486595 := NewDATTrie()
	__obf_59c627d756486595.__obf_637cf2049e4fbf69[1] = 1

	type __obf_8207d67cd8ff19b4 struct {
		__obf_68cbf53b787e72da int32
		__obf_bd61e0edf4e82645 *__obf_6d235e2ecffc87cb
	}

	__obf_f02b262bac03c070 := []__obf_8207d67cd8ff19b4{{__obf_68cbf53b787e72da: 1, __obf_bd61e0edf4e82645: __obf_39fe5d17818af665}}

	for len(__obf_f02b262bac03c070) > 0 {
		__obf_02eecbc1c91b5727 := __obf_f02b262bac03c070[0]
		__obf_f02b262bac03c070 = __obf_f02b262bac03c070[1:]

		if len(__obf_02eecbc1c91b5727.__obf_bd61e0edf4e82645.__obf_564314fc8e971646) == 0 {
			continue
		}

		__obf_7bf5bf0ea3c571fd := make([]rune, 0, len(__obf_02eecbc1c91b5727.__obf_bd61e0edf4e82645.__obf_564314fc8e971646))
		for __obf_582cf66952790466 := range __obf_02eecbc1c91b5727.__obf_bd61e0edf4e82645.__obf_564314fc8e971646 {
			__obf_7bf5bf0ea3c571fd = append(__obf_7bf5bf0ea3c571fd, __obf_582cf66952790466)
		}
		slices.Sort(__obf_7bf5bf0ea3c571fd)

		__obf_b1c087e95af77648 := __obf_59c627d756486595.__obf_5208ca3ce6fe2bbc(__obf_7bf5bf0ea3c571fd)
		__obf_59c627d756486595.__obf_22ac97ae6458a322[__obf_02eecbc1c91b5727.__obf_68cbf53b787e72da] = __obf_b1c087e95af77648

		for _, __obf_c7be2cb2604d35ec := range __obf_7bf5bf0ea3c571fd {
			__obf_0e6755ef936efefe := __obf_02eecbc1c91b5727.__obf_bd61e0edf4e82645.__obf_564314fc8e971646[__obf_c7be2cb2604d35ec]
			__obf_4e2c98da27bb4507 := __obf_b1c087e95af77648 + int32(__obf_c7be2cb2604d35ec)
			__obf_59c627d756486595.__obf_8ded9723e71599b4(__obf_4e2c98da27bb4507 + 1)

			__obf_59c627d756486595.__obf_637cf2049e4fbf69[__obf_4e2c98da27bb4507] = __obf_02eecbc1c91b5727.__obf_68cbf53b787e72da
			if __obf_0e6755ef936efefe.__obf_6778034c36c7f7c8 {
				__obf_59c627d756486595.__obf_9bc0c5da352e4e2b[__obf_4e2c98da27bb4507] = true
			}
			if int32(len(__obf_59c627d756486595.__obf_22ac97ae6458a322)) <= __obf_4e2c98da27bb4507 {
				__obf_59c627d756486595.__obf_8ded9723e71599b4(__obf_4e2c98da27bb4507 + 1)
			}
			if __obf_59c627d756486595.__obf_01f4e78b15b572f5 <= __obf_4e2c98da27bb4507 {
				__obf_59c627d756486595.__obf_01f4e78b15b572f5 = __obf_4e2c98da27bb4507 + 1
			}

			__obf_f02b262bac03c070 = append(__obf_f02b262bac03c070, __obf_8207d67cd8ff19b4{__obf_68cbf53b787e72da: __obf_4e2c98da27bb4507, __obf_bd61e0edf4e82645: __obf_0e6755ef936efefe})
		}
	}
	return __obf_59c627d756486595
}

func (__obf_59c627d756486595 *DATTrie) __obf_5208ca3ce6fe2bbc(__obf_af4d09e6eaabe62a []rune) int32 {
	for __obf_22ac97ae6458a322 := int32(1); ; __obf_22ac97ae6458a322++ {
		__obf_6483d1e500971b8f := false
		for _, __obf_8909273538900418 := range __obf_af4d09e6eaabe62a {
			__obf_d97c3ae373d6a91d := __obf_22ac97ae6458a322 + int32(__obf_8909273538900418)
			if int32(len(__obf_59c627d756486595.__obf_637cf2049e4fbf69)) <= __obf_d97c3ae373d6a91d {
				__obf_59c627d756486595.__obf_8ded9723e71599b4(__obf_d97c3ae373d6a91d + 1)
			}
			if __obf_59c627d756486595.__obf_637cf2049e4fbf69[__obf_d97c3ae373d6a91d] != 0 {
				__obf_6483d1e500971b8f = true
				break
			}
		}
		if !__obf_6483d1e500971b8f {
			return __obf_22ac97ae6458a322
		}
	}
}

func (__obf_59c627d756486595 *DATTrie) __obf_8ded9723e71599b4(__obf_51eb9421da7497c9 int32) {
	if __obf_51eb9421da7497c9 < int32(len(__obf_59c627d756486595.__obf_22ac97ae6458a322)) {
		return
	}
	__obf_016b34bec0d980ff := max(int32(len(__obf_59c627d756486595.__obf_22ac97ae6458a322))+int32(len(__obf_59c627d756486595.__obf_22ac97ae6458a322))/2, __obf_51eb9421da7497c9)
	__obf_e6ddd63be68c9c9f := make([]int32, __obf_016b34bec0d980ff)
	__obf_ceb7b9162e6b4c03 := make([]int32, __obf_016b34bec0d980ff)
	__obf_197c670754d42999 := make([]bool, __obf_016b34bec0d980ff)
	copy(__obf_e6ddd63be68c9c9f, __obf_59c627d756486595.__obf_22ac97ae6458a322)
	copy(__obf_ceb7b9162e6b4c03, __obf_59c627d756486595.__obf_637cf2049e4fbf69)
	copy(__obf_197c670754d42999, __obf_59c627d756486595.__obf_9bc0c5da352e4e2b)
	__obf_59c627d756486595.__obf_22ac97ae6458a322 = __obf_e6ddd63be68c9c9f
	__obf_59c627d756486595.__obf_637cf2049e4fbf69 = __obf_ceb7b9162e6b4c03
	__obf_59c627d756486595.__obf_9bc0c5da352e4e2b = __obf_197c670754d42999
}

func (__obf_59c627d756486595 *DATTrie) __obf_aa0774e419c68810(__obf_3f320056c510ddca int32, __obf_8909273538900418 rune) int32 {
	if __obf_3f320056c510ddca <= 0 || __obf_3f320056c510ddca >= int32(len(__obf_59c627d756486595.__obf_22ac97ae6458a322)) || __obf_59c627d756486595.__obf_22ac97ae6458a322[__obf_3f320056c510ddca] == 0 {
		return -1
	}
	__obf_d97c3ae373d6a91d := __obf_59c627d756486595.__obf_22ac97ae6458a322[__obf_3f320056c510ddca] + int32(__obf_8909273538900418)
	if __obf_d97c3ae373d6a91d > 0 && __obf_d97c3ae373d6a91d < int32(len(__obf_59c627d756486595.__obf_637cf2049e4fbf69)) && __obf_59c627d756486595.__obf_637cf2049e4fbf69[__obf_d97c3ae373d6a91d] == __obf_3f320056c510ddca {
		return __obf_d97c3ae373d6a91d
	}
	return -1
}

func (__obf_59c627d756486595 *DATTrie) __obf_6778034c36c7f7c8(__obf_3f320056c510ddca int32) bool {
	if __obf_3f320056c510ddca > 0 && __obf_3f320056c510ddca < int32(len(__obf_59c627d756486595.__obf_9bc0c5da352e4e2b)) {
		return __obf_59c627d756486595.__obf_9bc0c5da352e4e2b[__obf_3f320056c510ddca]
	}
	return false
}

func (__obf_59c627d756486595 *DATTrie) Contains(__obf_100ac6ab033077ae string) bool {
	__obf_0dc92fd760b3ad1d := []rune(__obf_100ac6ab033077ae)
	for __obf_481c1e102be52ccd := range __obf_0dc92fd760b3ad1d {
		__obf_bd61e0edf4e82645 := int32(1)
		for __obf_547930925e67b72f := __obf_481c1e102be52ccd; __obf_547930925e67b72f < len(__obf_0dc92fd760b3ad1d); __obf_547930925e67b72f++ {
			__obf_bd61e0edf4e82645 = __obf_59c627d756486595.__obf_aa0774e419c68810(__obf_bd61e0edf4e82645, __obf_0dc92fd760b3ad1d[__obf_547930925e67b72f])
			if __obf_bd61e0edf4e82645 == -1 {
				break
			}
			if __obf_59c627d756486595.__obf_6778034c36c7f7c8(__obf_bd61e0edf4e82645) {
				return true
			}
		}
	}
	return false
}

func (__obf_59c627d756486595 *DATTrie) __obf_da1d7f0396ff91f7(__obf_0dc92fd760b3ad1d []rune) []MatchResult {
	__obf_0dd3b0f1cecc4d06 := make([]MatchResult, 0, __obf_1c44ebf5d8dd34a9)
	for __obf_481c1e102be52ccd := range __obf_0dc92fd760b3ad1d {
		__obf_bd61e0edf4e82645 := int32(1)
		for __obf_547930925e67b72f := __obf_481c1e102be52ccd; __obf_547930925e67b72f < len(__obf_0dc92fd760b3ad1d); __obf_547930925e67b72f++ {
			__obf_bd61e0edf4e82645 = __obf_59c627d756486595.__obf_aa0774e419c68810(__obf_bd61e0edf4e82645, __obf_0dc92fd760b3ad1d[__obf_547930925e67b72f])
			if __obf_bd61e0edf4e82645 == -1 {
				break
			}
			if __obf_59c627d756486595.__obf_6778034c36c7f7c8(__obf_bd61e0edf4e82645) {
				__obf_0dd3b0f1cecc4d06 = append(__obf_0dd3b0f1cecc4d06, MatchResult{
					Word:  string(__obf_0dc92fd760b3ad1d[__obf_481c1e102be52ccd : __obf_547930925e67b72f+1]),
					Start: __obf_481c1e102be52ccd,
					End:   __obf_547930925e67b72f + 1,
				})
			}
		}
	}
	return __obf_0dd3b0f1cecc4d06
}

type WuManberMatcher struct {
	__obf_14d47c0ed341ce4c []string
	__obf_8c1c46cc1eb8fdcf int
	__obf_3d4e769143806edf map[string]int32
	__obf_28c1abf132d3f4a4 map[string][]int
	__obf_7e98e71276eac4da int
}

func NewWuManberMatcher(__obf_8c1c46cc1eb8fdcf int) *WuManberMatcher {
	return &WuManberMatcher{
		__obf_14d47c0ed341ce4c: make([]string, 0),
		__obf_8c1c46cc1eb8fdcf: __obf_8c1c46cc1eb8fdcf,
		__obf_3d4e769143806edf: make(map[string]int32),
		__obf_28c1abf132d3f4a4: make(map[string][]int),
		__obf_7e98e71276eac4da: int(^uint(0) >> 1),
	}
}

func (__obf_6fdee0c2817626ff *WuManberMatcher) AddPatterns(__obf_14d47c0ed341ce4c []string) {
	__obf_6fdee0c2817626ff.__obf_14d47c0ed341ce4c = __obf_14d47c0ed341ce4c
	for _, __obf_d805c1ceb2043f6b := range __obf_14d47c0ed341ce4c {
		__obf_d84d7c93d043c5f0 := utf8.RuneCountInString(__obf_d805c1ceb2043f6b)
		if __obf_d84d7c93d043c5f0 > 0 && __obf_d84d7c93d043c5f0 < __obf_6fdee0c2817626ff.__obf_7e98e71276eac4da {
			__obf_6fdee0c2817626ff.__obf_7e98e71276eac4da = __obf_d84d7c93d043c5f0
		}
	}
}

func (__obf_6fdee0c2817626ff *WuManberMatcher) Build() {
	if len(__obf_6fdee0c2817626ff.__obf_14d47c0ed341ce4c) == 0 {
		return
	}
	if __obf_6fdee0c2817626ff.__obf_7e98e71276eac4da < __obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf {
		__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf = __obf_6fdee0c2817626ff.__obf_7e98e71276eac4da
	}
	if __obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf <= 0 {
		__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf = 1
	}

	for __obf_481c1e102be52ccd, __obf_15038adae8b06461 := range __obf_6fdee0c2817626ff.__obf_14d47c0ed341ce4c {
		__obf_0dc92fd760b3ad1d := []rune(__obf_15038adae8b06461)
		__obf_ef0040a4b29b071e := len(__obf_0dc92fd760b3ad1d)
		for __obf_547930925e67b72f := 0; __obf_547930925e67b72f <= __obf_ef0040a4b29b071e-__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf; __obf_547930925e67b72f++ {
			__obf_66a03c7d5cc2796d := string(__obf_0dc92fd760b3ad1d[__obf_547930925e67b72f : __obf_547930925e67b72f+__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf])
			__obf_a339011ee0eb5b3c := int32(__obf_ef0040a4b29b071e - __obf_547930925e67b72f - __obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf)
			if __obf_3e7ecfe3b51cbcec, __obf_73a2d4d0fe8cd73d := __obf_6fdee0c2817626ff.__obf_3d4e769143806edf[__obf_66a03c7d5cc2796d]; !__obf_73a2d4d0fe8cd73d || __obf_a339011ee0eb5b3c < __obf_3e7ecfe3b51cbcec {
				__obf_6fdee0c2817626ff.__obf_3d4e769143806edf[__obf_66a03c7d5cc2796d] = __obf_a339011ee0eb5b3c
			}
		}

		if __obf_ef0040a4b29b071e >= __obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf {
			__obf_5608b5aee07e9679 := string(__obf_0dc92fd760b3ad1d[__obf_ef0040a4b29b071e-__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf:])
			__obf_6fdee0c2817626ff.__obf_28c1abf132d3f4a4[__obf_5608b5aee07e9679] = append(__obf_6fdee0c2817626ff.__obf_28c1abf132d3f4a4[__obf_5608b5aee07e9679], __obf_481c1e102be52ccd)
		}
	}
}

func (__obf_6fdee0c2817626ff *WuManberMatcher) Match(__obf_100ac6ab033077ae string) bool {
	if len(__obf_6fdee0c2817626ff.__obf_14d47c0ed341ce4c) == 0 {
		return false
	}
	__obf_0dc92fd760b3ad1d := []rune(__obf_100ac6ab033077ae)
	__obf_6eb9a99200f54561 := len(__obf_0dc92fd760b3ad1d)
	if __obf_6eb9a99200f54561 < __obf_6fdee0c2817626ff.__obf_7e98e71276eac4da {
		return false
	}
	__obf_5229e703e575004b := __obf_6fdee0c2817626ff.__obf_7e98e71276eac4da - 1
	for __obf_5229e703e575004b < __obf_6eb9a99200f54561 {
		__obf_0e9e88e8f353dde5 := max(__obf_5229e703e575004b-__obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf+1, 0)
		__obf_66a03c7d5cc2796d := string(__obf_0dc92fd760b3ad1d[__obf_0e9e88e8f353dde5 : __obf_5229e703e575004b+1])
		__obf_3d4e769143806edf, __obf_73a2d4d0fe8cd73d := __obf_6fdee0c2817626ff.__obf_3d4e769143806edf[__obf_66a03c7d5cc2796d]
		if !__obf_73a2d4d0fe8cd73d {
			__obf_5229e703e575004b += (__obf_6fdee0c2817626ff.__obf_7e98e71276eac4da - __obf_6fdee0c2817626ff.__obf_8c1c46cc1eb8fdcf + 1)
			continue
		}
		if __obf_3d4e769143806edf == 0 {
			if __obf_ee71f74f9bed6460, __obf_1256bcce11518b0c := __obf_6fdee0c2817626ff.__obf_28c1abf132d3f4a4[__obf_66a03c7d5cc2796d]; __obf_1256bcce11518b0c {
				for _, __obf_6ddb9a8cd2076508 := range __obf_ee71f74f9bed6460 {
					__obf_15038adae8b06461 := __obf_6fdee0c2817626ff.__obf_14d47c0ed341ce4c[__obf_6ddb9a8cd2076508]
					__obf_f3e2c9e4f0dea061 := utf8.RuneCountInString(__obf_15038adae8b06461)
					__obf_dbc8bce9cfa625cc := __obf_5229e703e575004b - __obf_f3e2c9e4f0dea061 + 1
					if __obf_dbc8bce9cfa625cc >= 0 && string(__obf_0dc92fd760b3ad1d[__obf_dbc8bce9cfa625cc:__obf_dbc8bce9cfa625cc+__obf_f3e2c9e4f0dea061]) == __obf_15038adae8b06461 {
						return true
					}
				}
			}
			__obf_5229e703e575004b++
		} else {
			__obf_5229e703e575004b += int(__obf_3d4e769143806edf)
		}
	}
	return false
}
