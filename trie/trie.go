package __obf_3b1f5fc614878e4b

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
	__obf_0064b9f6194e16fc = 1024
	__obf_48d1060f485316df = 1_000_000
	__obf_845c801915dc7ccb = // 1MB
	3
	__obf_bce5884e1d0ba5a4 = 100
	__obf_50512e5426718993 = 16
)

// 位图实现
type __obf_e61d5ea51e70e4e1 struct {
	__obf_9fa920ff616983f5 []uint64
}

func __obf_46184d633beef21a(__obf_7ae05edf1bc1043a int) *__obf_e61d5ea51e70e4e1 {
	__obf_440a652a115457e7 := // 每个 uint64 存储 64 位
		(__obf_7ae05edf1bc1043a + 63) / 64
	return &__obf_e61d5ea51e70e4e1{__obf_9fa920ff616983f5: make([]uint64, __obf_440a652a115457e7)}
}

func (__obf_1d6b4e33ca9b1775 *__obf_e61d5ea51e70e4e1) __obf_f1c033ed6f7addde(__obf_331c7073204a90da int) {
	__obf_1d6b4e33ca9b1775.__obf_9fa920ff616983f5[__obf_331c7073204a90da/64] |= 1 << (__obf_331c7073204a90da % 64)
}

func (__obf_1d6b4e33ca9b1775 *__obf_e61d5ea51e70e4e1) __obf_ec49e25047332e1d(__obf_331c7073204a90da int) bool {
	return __obf_1d6b4e33ca9b1775.__obf_9fa920ff616983f5[__obf_331c7073204a90da/64]&(1<<(__obf_331c7073204a90da%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_1c42e26cde8af5ad *DATTrie
	__obf_d8da60170d15c427 *WuManberMatcher
	__obf_3c1b367f563cfff8 sync.RWMutex
}

func (__obf_34705bf6010f6bc7 *SensitiveFilter) Contains(__obf_a84e33958a5dcb66 string) bool {
	__obf_34705bf6010f6bc7.__obf_3c1b367f563cfff8.
		Lock()
	defer __obf_34705bf6010f6bc7.__obf_3c1b367f563cfff8.Unlock()
	if len(__obf_a84e33958a5dcb66) > __obf_bce5884e1d0ba5a4 {
		if __obf_34705bf6010f6bc7.__obf_d8da60170d15c427.Match(__obf_a84e33958a5dcb66) {
			return true
		} else {
			return false
		}
	}
	return __obf_34705bf6010f6bc7.__obf_1c42e26cde8af5ad.Contains(__obf_a84e33958a5dcb66)
}

func (__obf_34705bf6010f6bc7 *SensitiveFilter) FindAll(__obf_a84e33958a5dcb66 string) []MatchResult {
	if __obf_34705bf6010f6bc7 == nil || __obf_34705bf6010f6bc7.__obf_1c42e26cde8af5ad == nil {
		return nil
	}
	if __obf_a84e33958a5dcb66 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_a84e33958a5dcb66) > __obf_48d1060f485316df {
		return nil // 或返回错误
	}
	__obf_f3150b4047f1f308 := // 1. 找出所有可能重叠的匹配项
		__obf_34705bf6010f6bc7.__obf_1c42e26cde8af5ad.__obf_65e96160765a434d([]rune(__obf_a84e33958a5dcb66))
	if len(__obf_f3150b4047f1f308) < 2 {
		return __obf_f3150b4047f1f308
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_f3150b4047f1f308, func(__obf_d8c1bf5a4eb8a1ab, __obf_c9e569a1236eabde int) bool {
		__obf_155bfd6703a10d85 := __obf_f3150b4047f1f308[__obf_d8c1bf5a4eb8a1ab].End - __obf_f3150b4047f1f308[__obf_d8c1bf5a4eb8a1ab].Start
		__obf_10cc4a6ea0827a7d := __obf_f3150b4047f1f308[__obf_c9e569a1236eabde].End - __obf_f3150b4047f1f308[__obf_c9e569a1236eabde].Start
		if __obf_155bfd6703a10d85 != __obf_10cc4a6ea0827a7d {
			return __obf_155bfd6703a10d85 > __obf_10cc4a6ea0827a7d
		}
		return __obf_f3150b4047f1f308[__obf_d8c1bf5a4eb8a1ab].Start < __obf_f3150b4047f1f308[__obf_c9e569a1236eabde].Start
	})
	__obf_5a3415778ac3837e := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_f3150b4047f1f308))
	__obf_252add1dd6a1ce50 := // occupied := make([]bool, len(runes))
		__obf_46184d633beef21a(len([]rune(__obf_a84e33958a5dcb66)))

	for _, __obf_281f36801766c5c0 := range __obf_f3150b4047f1f308 {
		__obf_cc585e41abb03b26 := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_d8c1bf5a4eb8a1ab := __obf_281f36801766c5c0.Start; __obf_d8c1bf5a4eb8a1ab < __obf_281f36801766c5c0.End; __obf_d8c1bf5a4eb8a1ab++ {
			// if occupied[i] {
			if __obf_252add1dd6a1ce50.__obf_ec49e25047332e1d(__obf_d8c1bf5a4eb8a1ab) {
				__obf_cc585e41abb03b26 = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_cc585e41abb03b26 {
			__obf_5a3415778ac3837e = append(__obf_5a3415778ac3837e,
				// 标记该位置已被占用
				__obf_281f36801766c5c0)

			for __obf_d8c1bf5a4eb8a1ab := __obf_281f36801766c5c0.Start; __obf_d8c1bf5a4eb8a1ab < __obf_281f36801766c5c0.End; __obf_d8c1bf5a4eb8a1ab++ {
				__obf_252add1dd6a1ce50.
					// occupied[i] = true
					__obf_f1c033ed6f7addde(__obf_d8c1bf5a4eb8a1ab)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_5a3415778ac3837e, func(__obf_d8c1bf5a4eb8a1ab, __obf_c9e569a1236eabde int) bool {
		return __obf_5a3415778ac3837e[__obf_d8c1bf5a4eb8a1ab].Start < __obf_5a3415778ac3837e[__obf_c9e569a1236eabde].Start
	})

	return __obf_5a3415778ac3837e
}

func (__obf_34705bf6010f6bc7 *SensitiveFilter) Replace(__obf_a84e33958a5dcb66 string, __obf_d576b09909df5393 rune) string {
	__obf_730323256455161e := __obf_34705bf6010f6bc7.FindAll(__obf_a84e33958a5dcb66)
	if len(__obf_730323256455161e) == 0 {
		return __obf_a84e33958a5dcb66
	}
	__obf_62a595cec6b2a4d0 := []rune(__obf_a84e33958a5dcb66)
	for _, __obf_281f36801766c5c0 := range __obf_730323256455161e {
		for __obf_d8c1bf5a4eb8a1ab := __obf_281f36801766c5c0.Start; __obf_d8c1bf5a4eb8a1ab < __obf_281f36801766c5c0.End; __obf_d8c1bf5a4eb8a1ab++ {
			__obf_62a595cec6b2a4d0[__obf_d8c1bf5a4eb8a1ab] = __obf_d576b09909df5393
		}
	}
	return string(__obf_62a595cec6b2a4d0)
}

func (__obf_34705bf6010f6bc7 *SensitiveFilter) ReplaceWithString(__obf_a84e33958a5dcb66 string, __obf_d576b09909df5393 string) string {
	__obf_730323256455161e := __obf_34705bf6010f6bc7.FindAll(__obf_a84e33958a5dcb66)
	if len(__obf_730323256455161e) == 0 {
		return __obf_a84e33958a5dcb66
	}

	var __obf_066fce58679855fd strings.Builder
	__obf_62a595cec6b2a4d0 := []rune(__obf_a84e33958a5dcb66)
	__obf_2d6c635e17d17505 := 0
	for _, __obf_281f36801766c5c0 := range __obf_730323256455161e {
		__obf_066fce58679855fd.
			WriteString(string(__obf_62a595cec6b2a4d0[__obf_2d6c635e17d17505:__obf_281f36801766c5c0.Start]))
		__obf_066fce58679855fd.
			WriteString(__obf_d576b09909df5393)
		__obf_2d6c635e17d17505 = __obf_281f36801766c5c0.End
	}
	__obf_066fce58679855fd.
		WriteString(string(__obf_62a595cec6b2a4d0[__obf_2d6c635e17d17505:]))
	return __obf_066fce58679855fd.String()
}

func (__obf_34705bf6010f6bc7 *SensitiveFilter) ReplaceWithFunc(__obf_a84e33958a5dcb66 string, __obf_4d811f515f6ba26b func(__obf_281f36801766c5c0 MatchResult) string) string {
	__obf_730323256455161e := __obf_34705bf6010f6bc7.FindAll(__obf_a84e33958a5dcb66)
	if len(__obf_730323256455161e) == 0 {
		return __obf_a84e33958a5dcb66
	}

	var __obf_066fce58679855fd strings.Builder
	__obf_62a595cec6b2a4d0 := []rune(__obf_a84e33958a5dcb66)
	__obf_2d6c635e17d17505 := 0
	for _, __obf_281f36801766c5c0 := range __obf_730323256455161e {
		__obf_066fce58679855fd.
			WriteString(string(__obf_62a595cec6b2a4d0[__obf_2d6c635e17d17505:__obf_281f36801766c5c0.Start]))
		__obf_066fce58679855fd.
			WriteString(__obf_4d811f515f6ba26b(__obf_281f36801766c5c0))
		__obf_2d6c635e17d17505 = __obf_281f36801766c5c0.End
	}
	__obf_066fce58679855fd.
		WriteString(string(__obf_62a595cec6b2a4d0[__obf_2d6c635e17d17505:]))
	return __obf_066fce58679855fd.String()
}

// 添加 Close 方法释放资源
func (__obf_34705bf6010f6bc7 *SensitiveFilter) Close() {
	__obf_34705bf6010f6bc7.__obf_3c1b367f563cfff8.
		Lock()
	defer __obf_34705bf6010f6bc7.__obf_3c1b367f563cfff8.Unlock()
	__obf_34705bf6010f6bc7.__obf_1c42e26cde8af5ad = nil
	__obf_34705bf6010f6bc7.__obf_d8da60170d15c427 = nil
}

type FilterBuilder struct {
	__obf_bf8118fabc964ae1 map[string]struct{}
	__obf_38ad2ecacff6be83 int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_7ae05edf1bc1043a int) FilterOption {
	return func(__obf_1d6b4e33ca9b1775 *FilterBuilder) {
		if __obf_7ae05edf1bc1043a > 0 {
			__obf_1d6b4e33ca9b1775.__obf_38ad2ecacff6be83 = __obf_7ae05edf1bc1043a
		}
	}
}

func WithWords(__obf_bf8118fabc964ae1 []string) FilterOption {
	return func(__obf_1d6b4e33ca9b1775 *FilterBuilder) {
		for _, __obf_d222e26cef289ef2 := range __obf_bf8118fabc964ae1 {
			if __obf_d222e26cef289ef2 != "" {
				__obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1[__obf_d222e26cef289ef2] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_0fbdadec36c22222 ...FilterOption) *FilterBuilder {
	__obf_1d6b4e33ca9b1775 := &FilterBuilder{__obf_bf8118fabc964ae1: make(map[string]struct{}), __obf_38ad2ecacff6be83: __obf_845c801915dc7ccb}
	for _, __obf_cb52bc05c6aad706 := range __obf_0fbdadec36c22222 {
		__obf_cb52bc05c6aad706(__obf_1d6b4e33ca9b1775)
	}
	return __obf_1d6b4e33ca9b1775
}

func (__obf_1d6b4e33ca9b1775 *FilterBuilder) AddWord(__obf_51933e724941c4af string) {
	if __obf_51933e724941c4af != "" {
		__obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1[__obf_51933e724941c4af] = struct{}{}
	}
}

func (__obf_1d6b4e33ca9b1775 *FilterBuilder) AddWords(__obf_bf8118fabc964ae1 []string) {
	for _, __obf_51933e724941c4af := range __obf_bf8118fabc964ae1 {
		__obf_1d6b4e33ca9b1775.
			AddWord(__obf_51933e724941c4af)
	}
}

func (__obf_1d6b4e33ca9b1775 *FilterBuilder) LoadFromFile(__obf_30255a651006679d string) error {
	__obf_880e333490fa150e, __obf_227e3823cac5e67b := os.Open(__obf_30255a651006679d)
	if __obf_227e3823cac5e67b != nil {
		return __obf_227e3823cac5e67b
	}
	defer __obf_880e333490fa150e.Close()
	__obf_9f201e6b1b6f0ffb := bufio.NewReader(__obf_880e333490fa150e)
	__obf_0e3cd5f4d8e4b87b := []byte{0xEF, 0xBB, 0xBF}
	__obf_ccc50e7dce0f4e1f, __obf_227e3823cac5e67b := __obf_9f201e6b1b6f0ffb.Peek(3)
	if __obf_227e3823cac5e67b == nil && bytes.Equal(__obf_ccc50e7dce0f4e1f, __obf_0e3cd5f4d8e4b87b) {
		_, _ = __obf_9f201e6b1b6f0ffb.Discard(3)
	}
	__obf_4c1dffbee1e9cc43 := bufio.NewScanner(__obf_9f201e6b1b6f0ffb)
	for __obf_4c1dffbee1e9cc43.Scan() {
		__obf_2774775bd41a0e0f := strings.TrimSpace(__obf_4c1dffbee1e9cc43.Text())
		__obf_1d6b4e33ca9b1775.
			AddWord(__obf_2774775bd41a0e0f)
	}
	return __obf_4c1dffbee1e9cc43.Err()
}

func (__obf_1d6b4e33ca9b1775 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1) == 0 {
		return &SensitiveFilter{__obf_1c42e26cde8af5ad: NewDATTrie(), __obf_d8da60170d15c427: NewWuManberMatcher(__obf_1d6b4e33ca9b1775.__obf_38ad2ecacff6be83)}
	}
	__obf_d6f6b9400e0e2226 := &__obf_a550e7f523918cda{__obf_466f09a4a50246a4: make(map[rune]*__obf_a550e7f523918cda)}
	__obf_81d37d8b2cd9b736 := make([]string, 0, len(__obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1))
	for __obf_51933e724941c4af := range __obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1 {
		__obf_81d37d8b2cd9b736 = append(__obf_81d37d8b2cd9b736, __obf_51933e724941c4af)
		__obf_d6f6b9400e0e2226.__obf_23c144eb9d84846b(__obf_51933e724941c4af)
	}
	__obf_1c42e26cde8af5ad := __obf_d38084f74bf9e902(__obf_d6f6b9400e0e2226)
	__obf_d8da60170d15c427 := NewWuManberMatcher(__obf_1d6b4e33ca9b1775.__obf_38ad2ecacff6be83)
	__obf_d8da60170d15c427.
		AddPatterns(__obf_81d37d8b2cd9b736)
	__obf_d8da60170d15c427.
		Build()

	return &SensitiveFilter{__obf_1c42e26cde8af5ad: __obf_1c42e26cde8af5ad, __obf_d8da60170d15c427: __obf_d8da60170d15c427}
}

func (__obf_1d6b4e33ca9b1775 *FilterBuilder) Clear() {
	__obf_1d6b4e33ca9b1775.__obf_bf8118fabc964ae1 = make(map[string]struct{})
}

type __obf_a550e7f523918cda struct {
	__obf_466f09a4a50246a4 map[rune]*__obf_a550e7f523918cda
	__obf_106bae0a60192ecc bool
}

func (__obf_440a652a115457e7 *__obf_a550e7f523918cda) __obf_23c144eb9d84846b(__obf_51933e724941c4af string) {
	__obf_c1fc72b4b42552d9 := __obf_440a652a115457e7
	for _, __obf_8a0bf00287f73646 := range __obf_51933e724941c4af {
		if _, __obf_47430ca261d42cfb := __obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4[__obf_8a0bf00287f73646]; !__obf_47430ca261d42cfb {
			__obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4[__obf_8a0bf00287f73646] = &__obf_a550e7f523918cda{__obf_466f09a4a50246a4: make(map[rune]*__obf_a550e7f523918cda)}
		}
		__obf_c1fc72b4b42552d9 = __obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4[__obf_8a0bf00287f73646]
	}
	__obf_c1fc72b4b42552d9.__obf_106bae0a60192ecc = true
}

type DATTrie struct {
	__obf_bb8bf071860242a9 []int32
	__obf_72308bd8e731fd7b []int32
	__obf_7e544757149c5c74 []bool
	__obf_7ae05edf1bc1043a int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_bb8bf071860242a9: make([]int32, __obf_0064b9f6194e16fc), __obf_72308bd8e731fd7b: make([]int32, __obf_0064b9f6194e16fc), __obf_7e544757149c5c74: make([]bool, __obf_0064b9f6194e16fc), __obf_7ae05edf1bc1043a: 1}
}

func __obf_d38084f74bf9e902(__obf_a5c963d7c7fd5739 *__obf_a550e7f523918cda) *DATTrie {
	__obf_ca9d265da57c2b0a := NewDATTrie()
	__obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b[1] = 1

	type __obf_a1285923ea1686e9 struct {
		__obf_84b871fc19d8f6b1 int32
		__obf_c1fc72b4b42552d9 *__obf_a550e7f523918cda
	}
	__obf_c465d148c1de2556 := []__obf_a1285923ea1686e9{{__obf_84b871fc19d8f6b1: 1, __obf_c1fc72b4b42552d9: __obf_a5c963d7c7fd5739}}

	for len(__obf_c465d148c1de2556) > 0 {
		__obf_f394dd778be71329 := __obf_c465d148c1de2556[0]
		__obf_c465d148c1de2556 = __obf_c465d148c1de2556[1:]

		if len(__obf_f394dd778be71329.__obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4) == 0 {
			continue
		}
		__obf_5fb85f7930c28081 := make([]rune, 0, len(__obf_f394dd778be71329.__obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4))
		for __obf_8a0bf00287f73646 := range __obf_f394dd778be71329.__obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4 {
			__obf_5fb85f7930c28081 = append(__obf_5fb85f7930c28081, __obf_8a0bf00287f73646)
		}
		slices.Sort(__obf_5fb85f7930c28081)
		__obf_5e069cfc5dc6cdbe := __obf_ca9d265da57c2b0a.__obf_59ac51df1c49ba98(__obf_5fb85f7930c28081)
		__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9[__obf_f394dd778be71329.__obf_84b871fc19d8f6b1] = __obf_5e069cfc5dc6cdbe

		for _, __obf_d21a096495e9dc30 := range __obf_5fb85f7930c28081 {
			__obf_6d88119fdd13fc8a := __obf_f394dd778be71329.__obf_c1fc72b4b42552d9.__obf_466f09a4a50246a4[__obf_d21a096495e9dc30]
			__obf_4fc2e172f4e6271a := __obf_5e069cfc5dc6cdbe + int32(__obf_d21a096495e9dc30)
			__obf_ca9d265da57c2b0a.__obf_dd2132c27e1ba4b8(__obf_4fc2e172f4e6271a + 1)
			__obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b[__obf_4fc2e172f4e6271a] = __obf_f394dd778be71329.__obf_84b871fc19d8f6b1
			if __obf_6d88119fdd13fc8a.__obf_106bae0a60192ecc {
				__obf_ca9d265da57c2b0a.__obf_7e544757149c5c74[__obf_4fc2e172f4e6271a] = true
			}
			if int32(len(__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9)) <= __obf_4fc2e172f4e6271a {
				__obf_ca9d265da57c2b0a.__obf_dd2132c27e1ba4b8(__obf_4fc2e172f4e6271a + 1)
			}
			if __obf_ca9d265da57c2b0a.__obf_7ae05edf1bc1043a <= __obf_4fc2e172f4e6271a {
				__obf_ca9d265da57c2b0a.__obf_7ae05edf1bc1043a = __obf_4fc2e172f4e6271a + 1
			}
			__obf_c465d148c1de2556 = append(__obf_c465d148c1de2556, __obf_a1285923ea1686e9{__obf_84b871fc19d8f6b1: __obf_4fc2e172f4e6271a, __obf_c1fc72b4b42552d9: __obf_6d88119fdd13fc8a})
		}
	}
	return __obf_ca9d265da57c2b0a
}

func (__obf_ca9d265da57c2b0a *DATTrie) __obf_59ac51df1c49ba98(__obf_218e0d097a75fadc []rune) int32 {
	for __obf_bb8bf071860242a9 := int32(1); ; __obf_bb8bf071860242a9++ {
		__obf_0ec929fc179ea97c := false
		for _, __obf_546564c99169da76 := range __obf_218e0d097a75fadc {
			__obf_d251ae39891f28a6 := __obf_bb8bf071860242a9 + int32(__obf_546564c99169da76)
			if int32(len(__obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b)) <= __obf_d251ae39891f28a6 {
				__obf_ca9d265da57c2b0a.__obf_dd2132c27e1ba4b8(__obf_d251ae39891f28a6 + 1)
			}
			if __obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b[__obf_d251ae39891f28a6] != 0 {
				__obf_0ec929fc179ea97c = true
				break
			}
		}
		if !__obf_0ec929fc179ea97c {
			return __obf_bb8bf071860242a9
		}
	}
}

func (__obf_ca9d265da57c2b0a *DATTrie) __obf_dd2132c27e1ba4b8(__obf_2f5b3f1a051d8a92 int32) {
	if __obf_2f5b3f1a051d8a92 < int32(len(__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9)) {
		return
	}
	__obf_c21903691f9a14ac := max(int32(len(__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9))+int32(len(__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9))/2, __obf_2f5b3f1a051d8a92)
	__obf_8e5e65c3a1295900 := make([]int32, __obf_c21903691f9a14ac)
	__obf_82997476e14ed9ec := make([]int32, __obf_c21903691f9a14ac)
	__obf_7e6722a8f8ab4ddf := make([]bool, __obf_c21903691f9a14ac)
	copy(__obf_8e5e65c3a1295900, __obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9)
	copy(__obf_82997476e14ed9ec, __obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b)
	copy(__obf_7e6722a8f8ab4ddf, __obf_ca9d265da57c2b0a.__obf_7e544757149c5c74)
	__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9 = __obf_8e5e65c3a1295900
	__obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b = __obf_82997476e14ed9ec
	__obf_ca9d265da57c2b0a.__obf_7e544757149c5c74 = __obf_7e6722a8f8ab4ddf
}

func (__obf_ca9d265da57c2b0a *DATTrie) __obf_7646ce7f8ed946e1(__obf_71b8a2496e23de4a int32, __obf_546564c99169da76 rune) int32 {
	if __obf_71b8a2496e23de4a <= 0 || __obf_71b8a2496e23de4a >= int32(len(__obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9)) || __obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9[__obf_71b8a2496e23de4a] == 0 {
		return -1
	}
	__obf_d251ae39891f28a6 := __obf_ca9d265da57c2b0a.__obf_bb8bf071860242a9[__obf_71b8a2496e23de4a] + int32(__obf_546564c99169da76)
	if __obf_d251ae39891f28a6 > 0 && __obf_d251ae39891f28a6 < int32(len(__obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b)) && __obf_ca9d265da57c2b0a.__obf_72308bd8e731fd7b[__obf_d251ae39891f28a6] == __obf_71b8a2496e23de4a {
		return __obf_d251ae39891f28a6
	}
	return -1
}

func (__obf_ca9d265da57c2b0a *DATTrie) __obf_106bae0a60192ecc(__obf_71b8a2496e23de4a int32) bool {
	if __obf_71b8a2496e23de4a > 0 && __obf_71b8a2496e23de4a < int32(len(__obf_ca9d265da57c2b0a.__obf_7e544757149c5c74)) {
		return __obf_ca9d265da57c2b0a.__obf_7e544757149c5c74[__obf_71b8a2496e23de4a]
	}
	return false
}

func (__obf_ca9d265da57c2b0a *DATTrie) Contains(__obf_a84e33958a5dcb66 string) bool {
	__obf_62a595cec6b2a4d0 := []rune(__obf_a84e33958a5dcb66)
	for __obf_d8c1bf5a4eb8a1ab := range __obf_62a595cec6b2a4d0 {
		__obf_c1fc72b4b42552d9 := int32(1)
		for __obf_c9e569a1236eabde := __obf_d8c1bf5a4eb8a1ab; __obf_c9e569a1236eabde < len(__obf_62a595cec6b2a4d0); __obf_c9e569a1236eabde++ {
			__obf_c1fc72b4b42552d9 = __obf_ca9d265da57c2b0a.__obf_7646ce7f8ed946e1(__obf_c1fc72b4b42552d9, __obf_62a595cec6b2a4d0[__obf_c9e569a1236eabde])
			if __obf_c1fc72b4b42552d9 == -1 {
				break
			}
			if __obf_ca9d265da57c2b0a.__obf_106bae0a60192ecc(__obf_c1fc72b4b42552d9) {
				return true
			}
		}
	}
	return false
}

func (__obf_ca9d265da57c2b0a *DATTrie) __obf_65e96160765a434d(__obf_62a595cec6b2a4d0 []rune) []MatchResult {
	__obf_4614aed34f02d66f := make([]MatchResult, 0, __obf_50512e5426718993)
	for __obf_d8c1bf5a4eb8a1ab := range __obf_62a595cec6b2a4d0 {
		__obf_c1fc72b4b42552d9 := int32(1)
		for __obf_c9e569a1236eabde := __obf_d8c1bf5a4eb8a1ab; __obf_c9e569a1236eabde < len(__obf_62a595cec6b2a4d0); __obf_c9e569a1236eabde++ {
			__obf_c1fc72b4b42552d9 = __obf_ca9d265da57c2b0a.__obf_7646ce7f8ed946e1(__obf_c1fc72b4b42552d9, __obf_62a595cec6b2a4d0[__obf_c9e569a1236eabde])
			if __obf_c1fc72b4b42552d9 == -1 {
				break
			}
			if __obf_ca9d265da57c2b0a.__obf_106bae0a60192ecc(__obf_c1fc72b4b42552d9) {
				__obf_4614aed34f02d66f = append(__obf_4614aed34f02d66f, MatchResult{
					Word:  string(__obf_62a595cec6b2a4d0[__obf_d8c1bf5a4eb8a1ab : __obf_c9e569a1236eabde+1]),
					Start: __obf_d8c1bf5a4eb8a1ab,
					End:   __obf_c9e569a1236eabde + 1,
				})
			}
		}
	}
	return __obf_4614aed34f02d66f
}

type WuManberMatcher struct {
	__obf_4be39de029ed6923 []string
	__obf_38ad2ecacff6be83 int
	__obf_7e4ff0109a0bc79c map[string]int32
	__obf_e99d97b7c70ebfa9 map[string][]int
	__obf_dbfe63a63c5efef1 int
}

func NewWuManberMatcher(__obf_38ad2ecacff6be83 int) *WuManberMatcher {
	return &WuManberMatcher{__obf_4be39de029ed6923: make([]string, 0), __obf_38ad2ecacff6be83: __obf_38ad2ecacff6be83, __obf_7e4ff0109a0bc79c: make(map[string]int32), __obf_e99d97b7c70ebfa9: make(map[string][]int), __obf_dbfe63a63c5efef1: int(^uint(0) >> 1)}
}

func (__obf_525ef50a7fa87ba2 *WuManberMatcher) AddPatterns(__obf_4be39de029ed6923 []string) {
	__obf_525ef50a7fa87ba2.__obf_4be39de029ed6923 = __obf_4be39de029ed6923
	for _, __obf_f5a9b3614efbf368 := range __obf_4be39de029ed6923 {
		__obf_fbd6ba0aabbe0c71 := utf8.RuneCountInString(__obf_f5a9b3614efbf368)
		if __obf_fbd6ba0aabbe0c71 > 0 && __obf_fbd6ba0aabbe0c71 < __obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 {
			__obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 = __obf_fbd6ba0aabbe0c71
		}
	}
}

func (__obf_525ef50a7fa87ba2 *WuManberMatcher) Build() {
	if len(__obf_525ef50a7fa87ba2.__obf_4be39de029ed6923) == 0 {
		return
	}
	if __obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 < __obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 {
		__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 = __obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1
	}
	if __obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 <= 0 {
		__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 = 1
	}

	for __obf_d8c1bf5a4eb8a1ab, __obf_16780741d516622b := range __obf_525ef50a7fa87ba2.__obf_4be39de029ed6923 {
		__obf_62a595cec6b2a4d0 := []rune(__obf_16780741d516622b)
		__obf_4485326d284697b5 := len(__obf_62a595cec6b2a4d0)
		for __obf_c9e569a1236eabde := 0; __obf_c9e569a1236eabde <= __obf_4485326d284697b5-__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83; __obf_c9e569a1236eabde++ {
			__obf_9d9567a8250d4003 := string(__obf_62a595cec6b2a4d0[__obf_c9e569a1236eabde : __obf_c9e569a1236eabde+__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83])
			__obf_4ebc06d7eb303d7c := int32(__obf_4485326d284697b5 - __obf_c9e569a1236eabde - __obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83)
			if __obf_1b70afa72640a5ab, __obf_6d47731ed262eae3 := __obf_525ef50a7fa87ba2.__obf_7e4ff0109a0bc79c[__obf_9d9567a8250d4003]; !__obf_6d47731ed262eae3 || __obf_4ebc06d7eb303d7c < __obf_1b70afa72640a5ab {
				__obf_525ef50a7fa87ba2.__obf_7e4ff0109a0bc79c[__obf_9d9567a8250d4003] = __obf_4ebc06d7eb303d7c
			}
		}

		if __obf_4485326d284697b5 >= __obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 {
			__obf_005f5b2fd450c6ef := string(__obf_62a595cec6b2a4d0[__obf_4485326d284697b5-__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83:])
			__obf_525ef50a7fa87ba2.__obf_e99d97b7c70ebfa9[__obf_005f5b2fd450c6ef] = append(__obf_525ef50a7fa87ba2.__obf_e99d97b7c70ebfa9[__obf_005f5b2fd450c6ef], __obf_d8c1bf5a4eb8a1ab)
		}
	}
}

func (__obf_525ef50a7fa87ba2 *WuManberMatcher) Match(__obf_a84e33958a5dcb66 string) bool {
	if len(__obf_525ef50a7fa87ba2.__obf_4be39de029ed6923) == 0 {
		return false
	}
	__obf_62a595cec6b2a4d0 := []rune(__obf_a84e33958a5dcb66)
	__obf_a06c4db6a7daefb3 := len(__obf_62a595cec6b2a4d0)
	if __obf_a06c4db6a7daefb3 < __obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 {
		return false
	}
	__obf_331c7073204a90da := __obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 - 1
	for __obf_331c7073204a90da < __obf_a06c4db6a7daefb3 {
		__obf_c058c791ef6e2df0 := max(__obf_331c7073204a90da-__obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83+1, 0)
		__obf_9d9567a8250d4003 := string(__obf_62a595cec6b2a4d0[__obf_c058c791ef6e2df0 : __obf_331c7073204a90da+1])
		__obf_7e4ff0109a0bc79c, __obf_6d47731ed262eae3 := __obf_525ef50a7fa87ba2.__obf_7e4ff0109a0bc79c[__obf_9d9567a8250d4003]
		if !__obf_6d47731ed262eae3 {
			__obf_331c7073204a90da += (__obf_525ef50a7fa87ba2.__obf_dbfe63a63c5efef1 - __obf_525ef50a7fa87ba2.__obf_38ad2ecacff6be83 + 1)
			continue
		}
		if __obf_7e4ff0109a0bc79c == 0 {
			if __obf_d6c6d06385c398c7, __obf_47430ca261d42cfb := __obf_525ef50a7fa87ba2.__obf_e99d97b7c70ebfa9[__obf_9d9567a8250d4003]; __obf_47430ca261d42cfb {
				for _, __obf_92482e6d7761285b := range __obf_d6c6d06385c398c7 {
					__obf_16780741d516622b := __obf_525ef50a7fa87ba2.__obf_4be39de029ed6923[__obf_92482e6d7761285b]
					__obf_3f9d74d714180007 := utf8.RuneCountInString(__obf_16780741d516622b)
					__obf_3dc35cbb595d92da := __obf_331c7073204a90da - __obf_3f9d74d714180007 + 1
					if __obf_3dc35cbb595d92da >= 0 && string(__obf_62a595cec6b2a4d0[__obf_3dc35cbb595d92da:__obf_3dc35cbb595d92da+__obf_3f9d74d714180007]) == __obf_16780741d516622b {
						return true
					}
				}
			}
			__obf_331c7073204a90da++
		} else {
			__obf_331c7073204a90da += int(__obf_7e4ff0109a0bc79c)
		}
	}
	return false
}
