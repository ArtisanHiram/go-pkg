package __obf_3c3e7d138194bb97

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
	__obf_0ec1a3ce022c6de9 = 1024
	__obf_cc3ed5aad57c4c9b = 1_000_000
	__obf_67ad8873d0e41bba = // 1MB
	3
	__obf_e061fd7d8b055db2 = 100
	__obf_60dfb52fbdddaaa2 = 16
)

// 位图实现
type __obf_87a183005970194b struct {
	__obf_2757c744dd238bb1 []uint64
}

func __obf_f40ed7ed1815502e(__obf_10b7c68ad99f9a21 int) *__obf_87a183005970194b {
	__obf_aae68ceb6e17dedb := // 每个 uint64 存储 64 位
		(__obf_10b7c68ad99f9a21 + 63) / 64
	return &__obf_87a183005970194b{__obf_2757c744dd238bb1: make([]uint64, __obf_aae68ceb6e17dedb)}
}

func (__obf_60fe9c1843022346 *__obf_87a183005970194b) __obf_6035a9da58561cad(__obf_f8a590cb6583c3c4 int) {
	__obf_60fe9c1843022346.__obf_2757c744dd238bb1[__obf_f8a590cb6583c3c4/64] |= 1 << (__obf_f8a590cb6583c3c4 % 64)
}

func (__obf_60fe9c1843022346 *__obf_87a183005970194b) __obf_0de331218b182aa1(__obf_f8a590cb6583c3c4 int) bool {
	return __obf_60fe9c1843022346.__obf_2757c744dd238bb1[__obf_f8a590cb6583c3c4/64]&(1<<(__obf_f8a590cb6583c3c4%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_326662376bbcd008 *DATTrie
	__obf_55c70a58b803dc60 *WuManberMatcher
	__obf_1dcbc6d988b39138 sync.RWMutex
}

func (__obf_a0dfaa331f04ec5f *SensitiveFilter) Contains(__obf_00d587cd0df9a816 string) bool {
	__obf_a0dfaa331f04ec5f.__obf_1dcbc6d988b39138.
		Lock()
	defer __obf_a0dfaa331f04ec5f.__obf_1dcbc6d988b39138.Unlock()
	if len(__obf_00d587cd0df9a816) > __obf_e061fd7d8b055db2 {
		if __obf_a0dfaa331f04ec5f.__obf_55c70a58b803dc60.Match(__obf_00d587cd0df9a816) {
			return true
		} else {
			return false
		}
	}
	return __obf_a0dfaa331f04ec5f.__obf_326662376bbcd008.Contains(__obf_00d587cd0df9a816)
}

func (__obf_a0dfaa331f04ec5f *SensitiveFilter) FindAll(__obf_00d587cd0df9a816 string) []MatchResult {
	if __obf_a0dfaa331f04ec5f == nil || __obf_a0dfaa331f04ec5f.__obf_326662376bbcd008 == nil {
		return nil
	}
	if __obf_00d587cd0df9a816 == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_00d587cd0df9a816) > __obf_cc3ed5aad57c4c9b {
		return nil // 或返回错误
	}
	__obf_17214b301dd09ccc := // 1. 找出所有可能重叠的匹配项
		__obf_a0dfaa331f04ec5f.__obf_326662376bbcd008.__obf_689f5625f72cb479([]rune(__obf_00d587cd0df9a816))
	if len(__obf_17214b301dd09ccc) < 2 {
		return __obf_17214b301dd09ccc
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_17214b301dd09ccc, func(__obf_67d3deb431c1a5a7, __obf_609d9bd53b5268f9 int) bool {
		__obf_295875b1849af7f3 := __obf_17214b301dd09ccc[__obf_67d3deb431c1a5a7].End - __obf_17214b301dd09ccc[__obf_67d3deb431c1a5a7].Start
		__obf_715cd46fd123cd6b := __obf_17214b301dd09ccc[__obf_609d9bd53b5268f9].End - __obf_17214b301dd09ccc[__obf_609d9bd53b5268f9].Start
		if __obf_295875b1849af7f3 != __obf_715cd46fd123cd6b {
			return __obf_295875b1849af7f3 > __obf_715cd46fd123cd6b
		}
		return __obf_17214b301dd09ccc[__obf_67d3deb431c1a5a7].Start < __obf_17214b301dd09ccc[__obf_609d9bd53b5268f9].Start
	})
	__obf_c7289b3f330885d7 := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_17214b301dd09ccc))
	__obf_ee30044d25fe1d1c := // occupied := make([]bool, len(runes))
		__obf_f40ed7ed1815502e(len([]rune(__obf_00d587cd0df9a816)))

	for _, __obf_ffed2cb5dbc87bfe := range __obf_17214b301dd09ccc {
		__obf_ad24766362028b2b := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_67d3deb431c1a5a7 := __obf_ffed2cb5dbc87bfe.Start; __obf_67d3deb431c1a5a7 < __obf_ffed2cb5dbc87bfe.End; __obf_67d3deb431c1a5a7++ {
			// if occupied[i] {
			if __obf_ee30044d25fe1d1c.__obf_0de331218b182aa1(__obf_67d3deb431c1a5a7) {
				__obf_ad24766362028b2b = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_ad24766362028b2b {
			__obf_c7289b3f330885d7 = append(__obf_c7289b3f330885d7,
				// 标记该位置已被占用
				__obf_ffed2cb5dbc87bfe)

			for __obf_67d3deb431c1a5a7 := __obf_ffed2cb5dbc87bfe.Start; __obf_67d3deb431c1a5a7 < __obf_ffed2cb5dbc87bfe.End; __obf_67d3deb431c1a5a7++ {
				__obf_ee30044d25fe1d1c.
					// occupied[i] = true
					__obf_6035a9da58561cad(__obf_67d3deb431c1a5a7)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_c7289b3f330885d7, func(__obf_67d3deb431c1a5a7, __obf_609d9bd53b5268f9 int) bool {
		return __obf_c7289b3f330885d7[__obf_67d3deb431c1a5a7].Start < __obf_c7289b3f330885d7[__obf_609d9bd53b5268f9].Start
	})

	return __obf_c7289b3f330885d7
}

func (__obf_a0dfaa331f04ec5f *SensitiveFilter) Replace(__obf_00d587cd0df9a816 string, __obf_401d17ed70c71f63 rune) string {
	__obf_6386c60c3e65939c := __obf_a0dfaa331f04ec5f.FindAll(__obf_00d587cd0df9a816)
	if len(__obf_6386c60c3e65939c) == 0 {
		return __obf_00d587cd0df9a816
	}
	__obf_aae47c13ace4f6c7 := []rune(__obf_00d587cd0df9a816)
	for _, __obf_ffed2cb5dbc87bfe := range __obf_6386c60c3e65939c {
		for __obf_67d3deb431c1a5a7 := __obf_ffed2cb5dbc87bfe.Start; __obf_67d3deb431c1a5a7 < __obf_ffed2cb5dbc87bfe.End; __obf_67d3deb431c1a5a7++ {
			__obf_aae47c13ace4f6c7[__obf_67d3deb431c1a5a7] = __obf_401d17ed70c71f63
		}
	}
	return string(__obf_aae47c13ace4f6c7)
}

func (__obf_a0dfaa331f04ec5f *SensitiveFilter) ReplaceWithString(__obf_00d587cd0df9a816 string, __obf_401d17ed70c71f63 string) string {
	__obf_6386c60c3e65939c := __obf_a0dfaa331f04ec5f.FindAll(__obf_00d587cd0df9a816)
	if len(__obf_6386c60c3e65939c) == 0 {
		return __obf_00d587cd0df9a816
	}

	var __obf_7275d2c8ae939254 strings.Builder
	__obf_aae47c13ace4f6c7 := []rune(__obf_00d587cd0df9a816)
	__obf_16bcc2839f3a5737 := 0
	for _, __obf_ffed2cb5dbc87bfe := range __obf_6386c60c3e65939c {
		__obf_7275d2c8ae939254.
			WriteString(string(__obf_aae47c13ace4f6c7[__obf_16bcc2839f3a5737:__obf_ffed2cb5dbc87bfe.Start]))
		__obf_7275d2c8ae939254.
			WriteString(__obf_401d17ed70c71f63)
		__obf_16bcc2839f3a5737 = __obf_ffed2cb5dbc87bfe.End
	}
	__obf_7275d2c8ae939254.
		WriteString(string(__obf_aae47c13ace4f6c7[__obf_16bcc2839f3a5737:]))
	return __obf_7275d2c8ae939254.String()
}

func (__obf_a0dfaa331f04ec5f *SensitiveFilter) ReplaceWithFunc(__obf_00d587cd0df9a816 string, __obf_25c0c17aa27b91e1 func(__obf_ffed2cb5dbc87bfe MatchResult) string) string {
	__obf_6386c60c3e65939c := __obf_a0dfaa331f04ec5f.FindAll(__obf_00d587cd0df9a816)
	if len(__obf_6386c60c3e65939c) == 0 {
		return __obf_00d587cd0df9a816
	}

	var __obf_7275d2c8ae939254 strings.Builder
	__obf_aae47c13ace4f6c7 := []rune(__obf_00d587cd0df9a816)
	__obf_16bcc2839f3a5737 := 0
	for _, __obf_ffed2cb5dbc87bfe := range __obf_6386c60c3e65939c {
		__obf_7275d2c8ae939254.
			WriteString(string(__obf_aae47c13ace4f6c7[__obf_16bcc2839f3a5737:__obf_ffed2cb5dbc87bfe.Start]))
		__obf_7275d2c8ae939254.
			WriteString(__obf_25c0c17aa27b91e1(__obf_ffed2cb5dbc87bfe))
		__obf_16bcc2839f3a5737 = __obf_ffed2cb5dbc87bfe.End
	}
	__obf_7275d2c8ae939254.
		WriteString(string(__obf_aae47c13ace4f6c7[__obf_16bcc2839f3a5737:]))
	return __obf_7275d2c8ae939254.String()
}

// 添加 Close 方法释放资源
func (__obf_a0dfaa331f04ec5f *SensitiveFilter) Close() {
	__obf_a0dfaa331f04ec5f.__obf_1dcbc6d988b39138.
		Lock()
	defer __obf_a0dfaa331f04ec5f.__obf_1dcbc6d988b39138.Unlock()
	__obf_a0dfaa331f04ec5f.__obf_326662376bbcd008 = nil
	__obf_a0dfaa331f04ec5f.__obf_55c70a58b803dc60 = nil
}

type FilterBuilder struct {
	__obf_004ef383c11aa7b0 map[string]struct{}
	__obf_6cdc374f44cda29e int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_10b7c68ad99f9a21 int) FilterOption {
	return func(__obf_60fe9c1843022346 *FilterBuilder) {
		if __obf_10b7c68ad99f9a21 > 0 {
			__obf_60fe9c1843022346.__obf_6cdc374f44cda29e = __obf_10b7c68ad99f9a21
		}
	}
}

func WithWords(__obf_004ef383c11aa7b0 []string) FilterOption {
	return func(__obf_60fe9c1843022346 *FilterBuilder) {
		for _, __obf_deada875cb3504cd := range __obf_004ef383c11aa7b0 {
			if __obf_deada875cb3504cd != "" {
				__obf_60fe9c1843022346.__obf_004ef383c11aa7b0[__obf_deada875cb3504cd] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_842bab54600a483e ...FilterOption) *FilterBuilder {
	__obf_60fe9c1843022346 := &FilterBuilder{__obf_004ef383c11aa7b0: make(map[string]struct{}), __obf_6cdc374f44cda29e: __obf_67ad8873d0e41bba}
	for _, __obf_822492d9c41422e9 := range __obf_842bab54600a483e {
		__obf_822492d9c41422e9(__obf_60fe9c1843022346)
	}
	return __obf_60fe9c1843022346
}

func (__obf_60fe9c1843022346 *FilterBuilder) AddWord(__obf_b7314748a099a8d9 string) {
	if __obf_b7314748a099a8d9 != "" {
		__obf_60fe9c1843022346.__obf_004ef383c11aa7b0[__obf_b7314748a099a8d9] = struct{}{}
	}
}

func (__obf_60fe9c1843022346 *FilterBuilder) AddWords(__obf_004ef383c11aa7b0 []string) {
	for _, __obf_b7314748a099a8d9 := range __obf_004ef383c11aa7b0 {
		__obf_60fe9c1843022346.
			AddWord(__obf_b7314748a099a8d9)
	}
}

func (__obf_60fe9c1843022346 *FilterBuilder) LoadFromFile(__obf_6f5278f703639b3c string) error {
	__obf_28a527d5096c8539, __obf_bbfab29165a25abf := os.Open(__obf_6f5278f703639b3c)
	if __obf_bbfab29165a25abf != nil {
		return __obf_bbfab29165a25abf
	}
	defer __obf_28a527d5096c8539.Close()
	__obf_21bad4d6b6f0e373 := bufio.NewReader(__obf_28a527d5096c8539)
	__obf_53f507725c746c81 := []byte{0xEF, 0xBB, 0xBF}
	__obf_fb69faeb52e09c85, __obf_bbfab29165a25abf := __obf_21bad4d6b6f0e373.Peek(3)
	if __obf_bbfab29165a25abf == nil && bytes.Equal(__obf_fb69faeb52e09c85, __obf_53f507725c746c81) {
		_, _ = __obf_21bad4d6b6f0e373.Discard(3)
	}
	__obf_01aa09c9457ef14a := bufio.NewScanner(__obf_21bad4d6b6f0e373)
	for __obf_01aa09c9457ef14a.Scan() {
		__obf_ece93023d11b0fbc := strings.TrimSpace(__obf_01aa09c9457ef14a.Text())
		__obf_60fe9c1843022346.
			AddWord(__obf_ece93023d11b0fbc)
	}
	return __obf_01aa09c9457ef14a.Err()
}

func (__obf_60fe9c1843022346 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_60fe9c1843022346.__obf_004ef383c11aa7b0) == 0 {
		return &SensitiveFilter{__obf_326662376bbcd008: NewDATTrie(), __obf_55c70a58b803dc60: NewWuManberMatcher(__obf_60fe9c1843022346.__obf_6cdc374f44cda29e)}
	}
	__obf_85f4903fa8c4026b := &__obf_f1f3daf130c5cd24{__obf_17e23cc19753cd5c: make(map[rune]*__obf_f1f3daf130c5cd24)}
	__obf_f55cc4d7bf7a87f3 := make([]string, 0, len(__obf_60fe9c1843022346.__obf_004ef383c11aa7b0))
	for __obf_b7314748a099a8d9 := range __obf_60fe9c1843022346.__obf_004ef383c11aa7b0 {
		__obf_f55cc4d7bf7a87f3 = append(__obf_f55cc4d7bf7a87f3, __obf_b7314748a099a8d9)
		__obf_85f4903fa8c4026b.__obf_0bf60045b6e86538(__obf_b7314748a099a8d9)
	}
	__obf_326662376bbcd008 := __obf_729c96b7a90972b5(__obf_85f4903fa8c4026b)
	__obf_55c70a58b803dc60 := NewWuManberMatcher(__obf_60fe9c1843022346.__obf_6cdc374f44cda29e)
	__obf_55c70a58b803dc60.
		AddPatterns(__obf_f55cc4d7bf7a87f3)
	__obf_55c70a58b803dc60.
		Build()

	return &SensitiveFilter{__obf_326662376bbcd008: __obf_326662376bbcd008, __obf_55c70a58b803dc60: __obf_55c70a58b803dc60}
}

func (__obf_60fe9c1843022346 *FilterBuilder) Clear() {
	__obf_60fe9c1843022346.__obf_004ef383c11aa7b0 = make(map[string]struct{})
}

type __obf_f1f3daf130c5cd24 struct {
	__obf_17e23cc19753cd5c map[rune]*__obf_f1f3daf130c5cd24
	__obf_af520c5629e27122 bool
}

func (__obf_aae68ceb6e17dedb *__obf_f1f3daf130c5cd24) __obf_0bf60045b6e86538(__obf_b7314748a099a8d9 string) {
	__obf_b142b4cbfcd63cb6 := __obf_aae68ceb6e17dedb
	for _, __obf_8998379c7f6909c1 := range __obf_b7314748a099a8d9 {
		if _, __obf_46662dd7cc038814 := __obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c[__obf_8998379c7f6909c1]; !__obf_46662dd7cc038814 {
			__obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c[__obf_8998379c7f6909c1] = &__obf_f1f3daf130c5cd24{__obf_17e23cc19753cd5c: make(map[rune]*__obf_f1f3daf130c5cd24)}
		}
		__obf_b142b4cbfcd63cb6 = __obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c[__obf_8998379c7f6909c1]
	}
	__obf_b142b4cbfcd63cb6.__obf_af520c5629e27122 = true
}

type DATTrie struct {
	__obf_6653d08e2cbc7783 []int32
	__obf_4a0e1a4eff2518db []int32
	__obf_078297a345ff097a []bool
	__obf_10b7c68ad99f9a21 int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_6653d08e2cbc7783: make([]int32, __obf_0ec1a3ce022c6de9), __obf_4a0e1a4eff2518db: make([]int32, __obf_0ec1a3ce022c6de9), __obf_078297a345ff097a: make([]bool, __obf_0ec1a3ce022c6de9), __obf_10b7c68ad99f9a21: 1}
}

func __obf_729c96b7a90972b5(__obf_600c8986f73db818 *__obf_f1f3daf130c5cd24) *DATTrie {
	__obf_fb99e8f966b0509e := NewDATTrie()
	__obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db[1] = 1

	type __obf_3a5707a01c63ac4d struct {
		__obf_5f7ea57ccc512104 int32
		__obf_b142b4cbfcd63cb6 *__obf_f1f3daf130c5cd24
	}
	__obf_b022cb0807c5ba8b := []__obf_3a5707a01c63ac4d{{__obf_5f7ea57ccc512104: 1, __obf_b142b4cbfcd63cb6: __obf_600c8986f73db818}}

	for len(__obf_b022cb0807c5ba8b) > 0 {
		__obf_8bd37e3ec948ce82 := __obf_b022cb0807c5ba8b[0]
		__obf_b022cb0807c5ba8b = __obf_b022cb0807c5ba8b[1:]

		if len(__obf_8bd37e3ec948ce82.__obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c) == 0 {
			continue
		}
		__obf_042b67450ef53990 := make([]rune, 0, len(__obf_8bd37e3ec948ce82.__obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c))
		for __obf_8998379c7f6909c1 := range __obf_8bd37e3ec948ce82.__obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c {
			__obf_042b67450ef53990 = append(__obf_042b67450ef53990, __obf_8998379c7f6909c1)
		}
		slices.Sort(__obf_042b67450ef53990)
		__obf_057d1f3d841fd1ed := __obf_fb99e8f966b0509e.__obf_a3077474aec6d927(__obf_042b67450ef53990)
		__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783[__obf_8bd37e3ec948ce82.__obf_5f7ea57ccc512104] = __obf_057d1f3d841fd1ed

		for _, __obf_6633e1017b6c4990 := range __obf_042b67450ef53990 {
			__obf_fefaed1b96967cae := __obf_8bd37e3ec948ce82.__obf_b142b4cbfcd63cb6.__obf_17e23cc19753cd5c[__obf_6633e1017b6c4990]
			__obf_15671e69feef325c := __obf_057d1f3d841fd1ed + int32(__obf_6633e1017b6c4990)
			__obf_fb99e8f966b0509e.__obf_64f935e4a85e9041(__obf_15671e69feef325c + 1)
			__obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db[__obf_15671e69feef325c] = __obf_8bd37e3ec948ce82.__obf_5f7ea57ccc512104
			if __obf_fefaed1b96967cae.__obf_af520c5629e27122 {
				__obf_fb99e8f966b0509e.__obf_078297a345ff097a[__obf_15671e69feef325c] = true
			}
			if int32(len(__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783)) <= __obf_15671e69feef325c {
				__obf_fb99e8f966b0509e.__obf_64f935e4a85e9041(__obf_15671e69feef325c + 1)
			}
			if __obf_fb99e8f966b0509e.__obf_10b7c68ad99f9a21 <= __obf_15671e69feef325c {
				__obf_fb99e8f966b0509e.__obf_10b7c68ad99f9a21 = __obf_15671e69feef325c + 1
			}
			__obf_b022cb0807c5ba8b = append(__obf_b022cb0807c5ba8b, __obf_3a5707a01c63ac4d{__obf_5f7ea57ccc512104: __obf_15671e69feef325c, __obf_b142b4cbfcd63cb6: __obf_fefaed1b96967cae})
		}
	}
	return __obf_fb99e8f966b0509e
}

func (__obf_fb99e8f966b0509e *DATTrie) __obf_a3077474aec6d927(__obf_8575f76132dcd5db []rune) int32 {
	for __obf_6653d08e2cbc7783 := int32(1); ; __obf_6653d08e2cbc7783++ {
		__obf_d8b288e70d59e49d := false
		for _, __obf_bc60fe0da3c7e87f := range __obf_8575f76132dcd5db {
			__obf_ea6bcf9e57307382 := __obf_6653d08e2cbc7783 + int32(__obf_bc60fe0da3c7e87f)
			if int32(len(__obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db)) <= __obf_ea6bcf9e57307382 {
				__obf_fb99e8f966b0509e.__obf_64f935e4a85e9041(__obf_ea6bcf9e57307382 + 1)
			}
			if __obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db[__obf_ea6bcf9e57307382] != 0 {
				__obf_d8b288e70d59e49d = true
				break
			}
		}
		if !__obf_d8b288e70d59e49d {
			return __obf_6653d08e2cbc7783
		}
	}
}

func (__obf_fb99e8f966b0509e *DATTrie) __obf_64f935e4a85e9041(__obf_6d54be4d7b07b753 int32) {
	if __obf_6d54be4d7b07b753 < int32(len(__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783)) {
		return
	}
	__obf_75d8bef07d890269 := max(int32(len(__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783))+int32(len(__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783))/2, __obf_6d54be4d7b07b753)
	__obf_9d0f6880f795f84c := make([]int32, __obf_75d8bef07d890269)
	__obf_babeebd9fcc086a3 := make([]int32, __obf_75d8bef07d890269)
	__obf_4bdd655465640757 := make([]bool, __obf_75d8bef07d890269)
	copy(__obf_9d0f6880f795f84c, __obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783)
	copy(__obf_babeebd9fcc086a3, __obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db)
	copy(__obf_4bdd655465640757, __obf_fb99e8f966b0509e.__obf_078297a345ff097a)
	__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783 = __obf_9d0f6880f795f84c
	__obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db = __obf_babeebd9fcc086a3
	__obf_fb99e8f966b0509e.__obf_078297a345ff097a = __obf_4bdd655465640757
}

func (__obf_fb99e8f966b0509e *DATTrie) __obf_2ab55a683a73cafc(__obf_b4413f47ca0f9e06 int32, __obf_bc60fe0da3c7e87f rune) int32 {
	if __obf_b4413f47ca0f9e06 <= 0 || __obf_b4413f47ca0f9e06 >= int32(len(__obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783)) || __obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783[__obf_b4413f47ca0f9e06] == 0 {
		return -1
	}
	__obf_ea6bcf9e57307382 := __obf_fb99e8f966b0509e.__obf_6653d08e2cbc7783[__obf_b4413f47ca0f9e06] + int32(__obf_bc60fe0da3c7e87f)
	if __obf_ea6bcf9e57307382 > 0 && __obf_ea6bcf9e57307382 < int32(len(__obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db)) && __obf_fb99e8f966b0509e.__obf_4a0e1a4eff2518db[__obf_ea6bcf9e57307382] == __obf_b4413f47ca0f9e06 {
		return __obf_ea6bcf9e57307382
	}
	return -1
}

func (__obf_fb99e8f966b0509e *DATTrie) __obf_af520c5629e27122(__obf_b4413f47ca0f9e06 int32) bool {
	if __obf_b4413f47ca0f9e06 > 0 && __obf_b4413f47ca0f9e06 < int32(len(__obf_fb99e8f966b0509e.__obf_078297a345ff097a)) {
		return __obf_fb99e8f966b0509e.__obf_078297a345ff097a[__obf_b4413f47ca0f9e06]
	}
	return false
}

func (__obf_fb99e8f966b0509e *DATTrie) Contains(__obf_00d587cd0df9a816 string) bool {
	__obf_aae47c13ace4f6c7 := []rune(__obf_00d587cd0df9a816)
	for __obf_67d3deb431c1a5a7 := range __obf_aae47c13ace4f6c7 {
		__obf_b142b4cbfcd63cb6 := int32(1)
		for __obf_609d9bd53b5268f9 := __obf_67d3deb431c1a5a7; __obf_609d9bd53b5268f9 < len(__obf_aae47c13ace4f6c7); __obf_609d9bd53b5268f9++ {
			__obf_b142b4cbfcd63cb6 = __obf_fb99e8f966b0509e.__obf_2ab55a683a73cafc(__obf_b142b4cbfcd63cb6, __obf_aae47c13ace4f6c7[__obf_609d9bd53b5268f9])
			if __obf_b142b4cbfcd63cb6 == -1 {
				break
			}
			if __obf_fb99e8f966b0509e.__obf_af520c5629e27122(__obf_b142b4cbfcd63cb6) {
				return true
			}
		}
	}
	return false
}

func (__obf_fb99e8f966b0509e *DATTrie) __obf_689f5625f72cb479(__obf_aae47c13ace4f6c7 []rune) []MatchResult {
	__obf_2ce2f6bb30630869 := make([]MatchResult, 0, __obf_60dfb52fbdddaaa2)
	for __obf_67d3deb431c1a5a7 := range __obf_aae47c13ace4f6c7 {
		__obf_b142b4cbfcd63cb6 := int32(1)
		for __obf_609d9bd53b5268f9 := __obf_67d3deb431c1a5a7; __obf_609d9bd53b5268f9 < len(__obf_aae47c13ace4f6c7); __obf_609d9bd53b5268f9++ {
			__obf_b142b4cbfcd63cb6 = __obf_fb99e8f966b0509e.__obf_2ab55a683a73cafc(__obf_b142b4cbfcd63cb6, __obf_aae47c13ace4f6c7[__obf_609d9bd53b5268f9])
			if __obf_b142b4cbfcd63cb6 == -1 {
				break
			}
			if __obf_fb99e8f966b0509e.__obf_af520c5629e27122(__obf_b142b4cbfcd63cb6) {
				__obf_2ce2f6bb30630869 = append(__obf_2ce2f6bb30630869, MatchResult{
					Word:  string(__obf_aae47c13ace4f6c7[__obf_67d3deb431c1a5a7 : __obf_609d9bd53b5268f9+1]),
					Start: __obf_67d3deb431c1a5a7,
					End:   __obf_609d9bd53b5268f9 + 1,
				})
			}
		}
	}
	return __obf_2ce2f6bb30630869
}

type WuManberMatcher struct {
	__obf_023fb5f094c42036 []string
	__obf_6cdc374f44cda29e int
	__obf_12e1087bd5f38286 map[string]int32
	__obf_30a79b3319bfe42b map[string][]int
	__obf_66216b854072a814 int
}

func NewWuManberMatcher(__obf_6cdc374f44cda29e int) *WuManberMatcher {
	return &WuManberMatcher{__obf_023fb5f094c42036: make([]string, 0), __obf_6cdc374f44cda29e: __obf_6cdc374f44cda29e, __obf_12e1087bd5f38286: make(map[string]int32), __obf_30a79b3319bfe42b: make(map[string][]int), __obf_66216b854072a814: int(^uint(0) >> 1)}
}

func (__obf_5526e4f8e424c261 *WuManberMatcher) AddPatterns(__obf_023fb5f094c42036 []string) {
	__obf_5526e4f8e424c261.__obf_023fb5f094c42036 = __obf_023fb5f094c42036
	for _, __obf_3b1ee108a7e3baaf := range __obf_023fb5f094c42036 {
		__obf_9ca63a0be24fb6a2 := utf8.RuneCountInString(__obf_3b1ee108a7e3baaf)
		if __obf_9ca63a0be24fb6a2 > 0 && __obf_9ca63a0be24fb6a2 < __obf_5526e4f8e424c261.__obf_66216b854072a814 {
			__obf_5526e4f8e424c261.__obf_66216b854072a814 = __obf_9ca63a0be24fb6a2
		}
	}
}

func (__obf_5526e4f8e424c261 *WuManberMatcher) Build() {
	if len(__obf_5526e4f8e424c261.__obf_023fb5f094c42036) == 0 {
		return
	}
	if __obf_5526e4f8e424c261.__obf_66216b854072a814 < __obf_5526e4f8e424c261.__obf_6cdc374f44cda29e {
		__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e = __obf_5526e4f8e424c261.__obf_66216b854072a814
	}
	if __obf_5526e4f8e424c261.__obf_6cdc374f44cda29e <= 0 {
		__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e = 1
	}

	for __obf_67d3deb431c1a5a7, __obf_60e69afeeb8727f9 := range __obf_5526e4f8e424c261.__obf_023fb5f094c42036 {
		__obf_aae47c13ace4f6c7 := []rune(__obf_60e69afeeb8727f9)
		__obf_32840c2625c97d6e := len(__obf_aae47c13ace4f6c7)
		for __obf_609d9bd53b5268f9 := 0; __obf_609d9bd53b5268f9 <= __obf_32840c2625c97d6e-__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e; __obf_609d9bd53b5268f9++ {
			__obf_0da6e71bbb703e75 := string(__obf_aae47c13ace4f6c7[__obf_609d9bd53b5268f9 : __obf_609d9bd53b5268f9+__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e])
			__obf_8e85d8c469f9e5ce := int32(__obf_32840c2625c97d6e - __obf_609d9bd53b5268f9 - __obf_5526e4f8e424c261.__obf_6cdc374f44cda29e)
			if __obf_8b72d52bf8ef6860, __obf_0addeb44cce851ff := __obf_5526e4f8e424c261.__obf_12e1087bd5f38286[__obf_0da6e71bbb703e75]; !__obf_0addeb44cce851ff || __obf_8e85d8c469f9e5ce < __obf_8b72d52bf8ef6860 {
				__obf_5526e4f8e424c261.__obf_12e1087bd5f38286[__obf_0da6e71bbb703e75] = __obf_8e85d8c469f9e5ce
			}
		}

		if __obf_32840c2625c97d6e >= __obf_5526e4f8e424c261.__obf_6cdc374f44cda29e {
			__obf_c14d54e8f28c09cc := string(__obf_aae47c13ace4f6c7[__obf_32840c2625c97d6e-__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e:])
			__obf_5526e4f8e424c261.__obf_30a79b3319bfe42b[__obf_c14d54e8f28c09cc] = append(__obf_5526e4f8e424c261.__obf_30a79b3319bfe42b[__obf_c14d54e8f28c09cc], __obf_67d3deb431c1a5a7)
		}
	}
}

func (__obf_5526e4f8e424c261 *WuManberMatcher) Match(__obf_00d587cd0df9a816 string) bool {
	if len(__obf_5526e4f8e424c261.__obf_023fb5f094c42036) == 0 {
		return false
	}
	__obf_aae47c13ace4f6c7 := []rune(__obf_00d587cd0df9a816)
	__obf_6c7739f901cf5cce := len(__obf_aae47c13ace4f6c7)
	if __obf_6c7739f901cf5cce < __obf_5526e4f8e424c261.__obf_66216b854072a814 {
		return false
	}
	__obf_f8a590cb6583c3c4 := __obf_5526e4f8e424c261.__obf_66216b854072a814 - 1
	for __obf_f8a590cb6583c3c4 < __obf_6c7739f901cf5cce {
		__obf_34d0e21bb28bbb7f := max(__obf_f8a590cb6583c3c4-__obf_5526e4f8e424c261.__obf_6cdc374f44cda29e+1, 0)
		__obf_0da6e71bbb703e75 := string(__obf_aae47c13ace4f6c7[__obf_34d0e21bb28bbb7f : __obf_f8a590cb6583c3c4+1])
		__obf_12e1087bd5f38286, __obf_0addeb44cce851ff := __obf_5526e4f8e424c261.__obf_12e1087bd5f38286[__obf_0da6e71bbb703e75]
		if !__obf_0addeb44cce851ff {
			__obf_f8a590cb6583c3c4 += (__obf_5526e4f8e424c261.__obf_66216b854072a814 - __obf_5526e4f8e424c261.__obf_6cdc374f44cda29e + 1)
			continue
		}
		if __obf_12e1087bd5f38286 == 0 {
			if __obf_11a093d923da2221, __obf_46662dd7cc038814 := __obf_5526e4f8e424c261.__obf_30a79b3319bfe42b[__obf_0da6e71bbb703e75]; __obf_46662dd7cc038814 {
				for _, __obf_d3a0bdb81381df8c := range __obf_11a093d923da2221 {
					__obf_60e69afeeb8727f9 := __obf_5526e4f8e424c261.__obf_023fb5f094c42036[__obf_d3a0bdb81381df8c]
					__obf_2a265c120fe65165 := utf8.RuneCountInString(__obf_60e69afeeb8727f9)
					__obf_7363ca20fefbc63a := __obf_f8a590cb6583c3c4 - __obf_2a265c120fe65165 + 1
					if __obf_7363ca20fefbc63a >= 0 && string(__obf_aae47c13ace4f6c7[__obf_7363ca20fefbc63a:__obf_7363ca20fefbc63a+__obf_2a265c120fe65165]) == __obf_60e69afeeb8727f9 {
						return true
					}
				}
			}
			__obf_f8a590cb6583c3c4++
		} else {
			__obf_f8a590cb6583c3c4 += int(__obf_12e1087bd5f38286)
		}
	}
	return false
}
