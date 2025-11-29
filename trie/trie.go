package __obf_7835dc7fc5eb2142

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
	__obf_4bf4a19d9834e88a = 1024
	__obf_1c4cf1624975e748 = 1_000_000
	__obf_398eebf92202c3d1 = // 1MB
	3
	__obf_1ab0b5e6c9e44533 = 100
	__obf_b1462bf48eca6147 = 16
)

// 位图实现
type __obf_4f9594fd8ac9be45 struct {
	__obf_71b1864061145367 []uint64
}

func __obf_7277ed1d5d840270(__obf_3c0899c6f8ed2d7d int) *__obf_4f9594fd8ac9be45 {
	__obf_67843b25c7ef5722 := // 每个 uint64 存储 64 位
		(__obf_3c0899c6f8ed2d7d + 63) / 64
	return &__obf_4f9594fd8ac9be45{__obf_71b1864061145367: make([]uint64, __obf_67843b25c7ef5722)}
}

func (__obf_2dfd6f91963d2f45 *__obf_4f9594fd8ac9be45) __obf_e9484a2ac2eaee12(__obf_03793896c47f8d30 int) {
	__obf_2dfd6f91963d2f45.__obf_71b1864061145367[__obf_03793896c47f8d30/64] |= 1 << (__obf_03793896c47f8d30 % 64)
}

func (__obf_2dfd6f91963d2f45 *__obf_4f9594fd8ac9be45) __obf_bd892b0121c27eac(__obf_03793896c47f8d30 int) bool {
	return __obf_2dfd6f91963d2f45.__obf_71b1864061145367[__obf_03793896c47f8d30/64]&(1<<(__obf_03793896c47f8d30%64)) != 0
}

type MatchResult struct {
	Word  string
	Start int
	End   int
}

type SensitiveFilter struct {
	__obf_50e83274a4fe755a *DATTrie
	__obf_5cce43980a89c035 *WuManberMatcher
	__obf_d46168803f411f26 sync.RWMutex
}

func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) Contains(__obf_8fd45bd0d265008c string) bool {
	__obf_e8b4dadf2a35b1d3.__obf_d46168803f411f26.
		Lock()
	defer __obf_e8b4dadf2a35b1d3.__obf_d46168803f411f26.Unlock()
	if len(__obf_8fd45bd0d265008c) > __obf_1ab0b5e6c9e44533 {
		if __obf_e8b4dadf2a35b1d3.__obf_5cce43980a89c035.Match(__obf_8fd45bd0d265008c) {
			return true
		} else {
			return false
		}
	}
	return __obf_e8b4dadf2a35b1d3.__obf_50e83274a4fe755a.Contains(__obf_8fd45bd0d265008c)
}

func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) FindAll(__obf_8fd45bd0d265008c string) []MatchResult {
	if __obf_e8b4dadf2a35b1d3 == nil || __obf_e8b4dadf2a35b1d3.__obf_50e83274a4fe755a == nil {
		return nil
	}
	if __obf_8fd45bd0d265008c == "" {
		return []MatchResult{}
	}
	// 添加最大文本长度限制
	if len(__obf_8fd45bd0d265008c) > __obf_1c4cf1624975e748 {
		return nil // 或返回错误
	}
	__obf_94dd07967f7a0179 := // 1. 找出所有可能重叠的匹配项
		__obf_e8b4dadf2a35b1d3.__obf_50e83274a4fe755a.__obf_6ac3be995625316d([]rune(__obf_8fd45bd0d265008c))
	if len(__obf_94dd07967f7a0179) < 2 {
		return __obf_94dd07967f7a0179
	}

	// 2. 按长度降序排序，最长的词优先。如果长度相同，则按起点升序。
	sort.Slice(__obf_94dd07967f7a0179, func(__obf_74e4c62962a40491, __obf_ec5c7f3e67ff13d2 int) bool {
		__obf_9a3e673b6f6091f9 := __obf_94dd07967f7a0179[__obf_74e4c62962a40491].End - __obf_94dd07967f7a0179[__obf_74e4c62962a40491].Start
		__obf_1c24892b066ba9ff := __obf_94dd07967f7a0179[__obf_ec5c7f3e67ff13d2].End - __obf_94dd07967f7a0179[__obf_ec5c7f3e67ff13d2].Start
		if __obf_9a3e673b6f6091f9 != __obf_1c24892b066ba9ff {
			return __obf_9a3e673b6f6091f9 > __obf_1c24892b066ba9ff
		}
		return __obf_94dd07967f7a0179[__obf_74e4c62962a40491].Start < __obf_94dd07967f7a0179[__obf_ec5c7f3e67ff13d2].Start
	})
	__obf_ca57418c490e9c53 := // 3. 使用贪心算法，选择不重叠的最长匹配
		make([]MatchResult, 0, len(__obf_94dd07967f7a0179))
	__obf_6d51ababf5d5307a := // occupied := make([]bool, len(runes))
		__obf_7277ed1d5d840270(len([]rune(__obf_8fd45bd0d265008c)))

	for _, __obf_92fb913565a85298 := range __obf_94dd07967f7a0179 {
		__obf_6519b5e53c79566b := false
		// 检查该词占据的位置是否已被更长的词占用
		for __obf_74e4c62962a40491 := __obf_92fb913565a85298.Start; __obf_74e4c62962a40491 < __obf_92fb913565a85298.End; __obf_74e4c62962a40491++ {
			// if occupied[i] {
			if __obf_6d51ababf5d5307a.__obf_bd892b0121c27eac(__obf_74e4c62962a40491) {
				__obf_6519b5e53c79566b = true
				break
			}
		}

		// 如果未被占用，则接受该匹配
		if !__obf_6519b5e53c79566b {
			__obf_ca57418c490e9c53 = append(__obf_ca57418c490e9c53,
				// 标记该位置已被占用
				__obf_92fb913565a85298)

			for __obf_74e4c62962a40491 := __obf_92fb913565a85298.Start; __obf_74e4c62962a40491 < __obf_92fb913565a85298.End; __obf_74e4c62962a40491++ {
				__obf_6d51ababf5d5307a.
					// occupied[i] = true
					__obf_e9484a2ac2eaee12(__obf_74e4c62962a40491)
			}
		}
	}

	// 4. 为了方便 Replace 函数处理，将结果按起点排序后返回
	sort.Slice(__obf_ca57418c490e9c53, func(__obf_74e4c62962a40491, __obf_ec5c7f3e67ff13d2 int) bool {
		return __obf_ca57418c490e9c53[__obf_74e4c62962a40491].Start < __obf_ca57418c490e9c53[__obf_ec5c7f3e67ff13d2].Start
	})

	return __obf_ca57418c490e9c53
}

func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) Replace(__obf_8fd45bd0d265008c string, __obf_9f61e15faac45890 rune) string {
	__obf_a9f58197f688732d := __obf_e8b4dadf2a35b1d3.FindAll(__obf_8fd45bd0d265008c)
	if len(__obf_a9f58197f688732d) == 0 {
		return __obf_8fd45bd0d265008c
	}
	__obf_0e2c9fed14041ce5 := []rune(__obf_8fd45bd0d265008c)
	for _, __obf_92fb913565a85298 := range __obf_a9f58197f688732d {
		for __obf_74e4c62962a40491 := __obf_92fb913565a85298.Start; __obf_74e4c62962a40491 < __obf_92fb913565a85298.End; __obf_74e4c62962a40491++ {
			__obf_0e2c9fed14041ce5[__obf_74e4c62962a40491] = __obf_9f61e15faac45890
		}
	}
	return string(__obf_0e2c9fed14041ce5)
}

func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) ReplaceWithString(__obf_8fd45bd0d265008c string, __obf_9f61e15faac45890 string) string {
	__obf_a9f58197f688732d := __obf_e8b4dadf2a35b1d3.FindAll(__obf_8fd45bd0d265008c)
	if len(__obf_a9f58197f688732d) == 0 {
		return __obf_8fd45bd0d265008c
	}

	var __obf_f36021775f19eb30 strings.Builder
	__obf_0e2c9fed14041ce5 := []rune(__obf_8fd45bd0d265008c)
	__obf_28d4518c36cdcfc9 := 0
	for _, __obf_92fb913565a85298 := range __obf_a9f58197f688732d {
		__obf_f36021775f19eb30.
			WriteString(string(__obf_0e2c9fed14041ce5[__obf_28d4518c36cdcfc9:__obf_92fb913565a85298.Start]))
		__obf_f36021775f19eb30.
			WriteString(__obf_9f61e15faac45890)
		__obf_28d4518c36cdcfc9 = __obf_92fb913565a85298.End
	}
	__obf_f36021775f19eb30.
		WriteString(string(__obf_0e2c9fed14041ce5[__obf_28d4518c36cdcfc9:]))
	return __obf_f36021775f19eb30.String()
}

func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) ReplaceWithFunc(__obf_8fd45bd0d265008c string, __obf_d9050ebc727558c1 func(__obf_92fb913565a85298 MatchResult) string) string {
	__obf_a9f58197f688732d := __obf_e8b4dadf2a35b1d3.FindAll(__obf_8fd45bd0d265008c)
	if len(__obf_a9f58197f688732d) == 0 {
		return __obf_8fd45bd0d265008c
	}

	var __obf_f36021775f19eb30 strings.Builder
	__obf_0e2c9fed14041ce5 := []rune(__obf_8fd45bd0d265008c)
	__obf_28d4518c36cdcfc9 := 0
	for _, __obf_92fb913565a85298 := range __obf_a9f58197f688732d {
		__obf_f36021775f19eb30.
			WriteString(string(__obf_0e2c9fed14041ce5[__obf_28d4518c36cdcfc9:__obf_92fb913565a85298.Start]))
		__obf_f36021775f19eb30.
			WriteString(__obf_d9050ebc727558c1(__obf_92fb913565a85298))
		__obf_28d4518c36cdcfc9 = __obf_92fb913565a85298.End
	}
	__obf_f36021775f19eb30.
		WriteString(string(__obf_0e2c9fed14041ce5[__obf_28d4518c36cdcfc9:]))
	return __obf_f36021775f19eb30.String()
}

// 添加 Close 方法释放资源
func (__obf_e8b4dadf2a35b1d3 *SensitiveFilter) Close() {
	__obf_e8b4dadf2a35b1d3.__obf_d46168803f411f26.
		Lock()
	defer __obf_e8b4dadf2a35b1d3.__obf_d46168803f411f26.Unlock()
	__obf_e8b4dadf2a35b1d3.__obf_50e83274a4fe755a = nil
	__obf_e8b4dadf2a35b1d3.__obf_5cce43980a89c035 = nil
}

type FilterBuilder struct {
	__obf_d703a722e273e59c map[string]struct{}
	__obf_2569a480ce883fcb int
}

type FilterOption func(*FilterBuilder)

func WithBlockSize(__obf_3c0899c6f8ed2d7d int) FilterOption {
	return func(__obf_2dfd6f91963d2f45 *FilterBuilder) {
		if __obf_3c0899c6f8ed2d7d > 0 {
			__obf_2dfd6f91963d2f45.__obf_2569a480ce883fcb = __obf_3c0899c6f8ed2d7d
		}
	}
}

func WithWords(__obf_d703a722e273e59c []string) FilterOption {
	return func(__obf_2dfd6f91963d2f45 *FilterBuilder) {
		for _, __obf_5a086cc16b4aa89b := range __obf_d703a722e273e59c {
			if __obf_5a086cc16b4aa89b != "" {
				__obf_2dfd6f91963d2f45.__obf_d703a722e273e59c[__obf_5a086cc16b4aa89b] = struct{}{}
			}
		}
	}
}

func NewFilterBuilder(__obf_70877c70273e0436 ...FilterOption) *FilterBuilder {
	__obf_2dfd6f91963d2f45 := &FilterBuilder{__obf_d703a722e273e59c: make(map[string]struct{}), __obf_2569a480ce883fcb: __obf_398eebf92202c3d1}
	for _, __obf_6b46c05936b1460d := range __obf_70877c70273e0436 {
		__obf_6b46c05936b1460d(__obf_2dfd6f91963d2f45)
	}
	return __obf_2dfd6f91963d2f45
}

func (__obf_2dfd6f91963d2f45 *FilterBuilder) AddWord(__obf_803effbcd3198ea6 string) {
	if __obf_803effbcd3198ea6 != "" {
		__obf_2dfd6f91963d2f45.__obf_d703a722e273e59c[__obf_803effbcd3198ea6] = struct{}{}
	}
}

func (__obf_2dfd6f91963d2f45 *FilterBuilder) AddWords(__obf_d703a722e273e59c []string) {
	for _, __obf_803effbcd3198ea6 := range __obf_d703a722e273e59c {
		__obf_2dfd6f91963d2f45.
			AddWord(__obf_803effbcd3198ea6)
	}
}

func (__obf_2dfd6f91963d2f45 *FilterBuilder) LoadFromFile(__obf_138c6ddb796bb3c8 string) error {
	__obf_ecfc306cae325dfb, __obf_24edb27520aab478 := os.Open(__obf_138c6ddb796bb3c8)
	if __obf_24edb27520aab478 != nil {
		return __obf_24edb27520aab478
	}
	defer __obf_ecfc306cae325dfb.Close()
	__obf_5ebbac3b0d4801e5 := bufio.NewReader(__obf_ecfc306cae325dfb)
	__obf_d03c507e81f35b1d := []byte{0xEF, 0xBB, 0xBF}
	__obf_5b8b205a66101ba8, __obf_24edb27520aab478 := __obf_5ebbac3b0d4801e5.Peek(3)
	if __obf_24edb27520aab478 == nil && bytes.Equal(__obf_5b8b205a66101ba8, __obf_d03c507e81f35b1d) {
		_, _ = __obf_5ebbac3b0d4801e5.Discard(3)
	}
	__obf_48416c620ebbcbdb := bufio.NewScanner(__obf_5ebbac3b0d4801e5)
	for __obf_48416c620ebbcbdb.Scan() {
		__obf_4fdf9fa0e9555ea9 := strings.TrimSpace(__obf_48416c620ebbcbdb.Text())
		__obf_2dfd6f91963d2f45.
			AddWord(__obf_4fdf9fa0e9555ea9)
	}
	return __obf_48416c620ebbcbdb.Err()
}

func (__obf_2dfd6f91963d2f45 *FilterBuilder) Build() *SensitiveFilter {
	if len(__obf_2dfd6f91963d2f45.__obf_d703a722e273e59c) == 0 {
		return &SensitiveFilter{__obf_50e83274a4fe755a: NewDATTrie(), __obf_5cce43980a89c035: NewWuManberMatcher(__obf_2dfd6f91963d2f45.__obf_2569a480ce883fcb)}
	}
	__obf_b769105ac3ded9ee := &__obf_5ff12cd2cc7cfa19{__obf_14e88d925d6fa53a: make(map[rune]*__obf_5ff12cd2cc7cfa19)}
	__obf_abd2e7d038288d8d := make([]string, 0, len(__obf_2dfd6f91963d2f45.__obf_d703a722e273e59c))
	for __obf_803effbcd3198ea6 := range __obf_2dfd6f91963d2f45.__obf_d703a722e273e59c {
		__obf_abd2e7d038288d8d = append(__obf_abd2e7d038288d8d, __obf_803effbcd3198ea6)
		__obf_b769105ac3ded9ee.__obf_e25cc745772af301(__obf_803effbcd3198ea6)
	}
	__obf_50e83274a4fe755a := __obf_8bb0c199230eb484(__obf_b769105ac3ded9ee)
	__obf_5cce43980a89c035 := NewWuManberMatcher(__obf_2dfd6f91963d2f45.__obf_2569a480ce883fcb)
	__obf_5cce43980a89c035.
		AddPatterns(__obf_abd2e7d038288d8d)
	__obf_5cce43980a89c035.
		Build()

	return &SensitiveFilter{__obf_50e83274a4fe755a: __obf_50e83274a4fe755a, __obf_5cce43980a89c035: __obf_5cce43980a89c035}
}

func (__obf_2dfd6f91963d2f45 *FilterBuilder) Clear() {
	__obf_2dfd6f91963d2f45.__obf_d703a722e273e59c = make(map[string]struct{})
}

type __obf_5ff12cd2cc7cfa19 struct {
	__obf_14e88d925d6fa53a map[rune]*__obf_5ff12cd2cc7cfa19
	__obf_ef5ed10d7c6a1194 bool
}

func (__obf_67843b25c7ef5722 *__obf_5ff12cd2cc7cfa19) __obf_e25cc745772af301(__obf_803effbcd3198ea6 string) {
	__obf_6deee49b8dd788d8 := __obf_67843b25c7ef5722
	for _, __obf_d5a6d0a1d55f42e3 := range __obf_803effbcd3198ea6 {
		if _, __obf_de4222e243defcfd := __obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a[__obf_d5a6d0a1d55f42e3]; !__obf_de4222e243defcfd {
			__obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a[__obf_d5a6d0a1d55f42e3] = &__obf_5ff12cd2cc7cfa19{__obf_14e88d925d6fa53a: make(map[rune]*__obf_5ff12cd2cc7cfa19)}
		}
		__obf_6deee49b8dd788d8 = __obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a[__obf_d5a6d0a1d55f42e3]
	}
	__obf_6deee49b8dd788d8.__obf_ef5ed10d7c6a1194 = true
}

type DATTrie struct {
	__obf_da7ce5419e7c2ab3 []int32
	__obf_00026af41d7782b2 []int32
	__obf_97e93c9eccbe5a21 []bool
	__obf_3c0899c6f8ed2d7d int32
}

func NewDATTrie() *DATTrie {
	return &DATTrie{__obf_da7ce5419e7c2ab3: make([]int32, __obf_4bf4a19d9834e88a), __obf_00026af41d7782b2: make([]int32, __obf_4bf4a19d9834e88a), __obf_97e93c9eccbe5a21: make([]bool, __obf_4bf4a19d9834e88a), __obf_3c0899c6f8ed2d7d: 1}
}

func __obf_8bb0c199230eb484(__obf_4f0ce1c7548b1bc0 *__obf_5ff12cd2cc7cfa19) *DATTrie {
	__obf_5868836158514e2c := NewDATTrie()
	__obf_5868836158514e2c.__obf_00026af41d7782b2[1] = 1

	type __obf_ecc42234bc4ab0ef struct {
		__obf_5c19df52406936fd int32
		__obf_6deee49b8dd788d8 *__obf_5ff12cd2cc7cfa19
	}
	__obf_bfd639f9c3be7e80 := []__obf_ecc42234bc4ab0ef{{__obf_5c19df52406936fd: 1, __obf_6deee49b8dd788d8: __obf_4f0ce1c7548b1bc0}}

	for len(__obf_bfd639f9c3be7e80) > 0 {
		__obf_488431258beaf3e9 := __obf_bfd639f9c3be7e80[0]
		__obf_bfd639f9c3be7e80 = __obf_bfd639f9c3be7e80[1:]

		if len(__obf_488431258beaf3e9.__obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a) == 0 {
			continue
		}
		__obf_8ba5c725cacd5f3d := make([]rune, 0, len(__obf_488431258beaf3e9.__obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a))
		for __obf_d5a6d0a1d55f42e3 := range __obf_488431258beaf3e9.__obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a {
			__obf_8ba5c725cacd5f3d = append(__obf_8ba5c725cacd5f3d, __obf_d5a6d0a1d55f42e3)
		}
		slices.Sort(__obf_8ba5c725cacd5f3d)
		__obf_cbf8722190f9f75b := __obf_5868836158514e2c.__obf_f8e2472f3b89a228(__obf_8ba5c725cacd5f3d)
		__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3[__obf_488431258beaf3e9.__obf_5c19df52406936fd] = __obf_cbf8722190f9f75b

		for _, __obf_452697cd389e3134 := range __obf_8ba5c725cacd5f3d {
			__obf_ceb17d4721b21808 := __obf_488431258beaf3e9.__obf_6deee49b8dd788d8.__obf_14e88d925d6fa53a[__obf_452697cd389e3134]
			__obf_b97394cad50e288c := __obf_cbf8722190f9f75b + int32(__obf_452697cd389e3134)
			__obf_5868836158514e2c.__obf_2e7a4d69845854bf(__obf_b97394cad50e288c + 1)
			__obf_5868836158514e2c.__obf_00026af41d7782b2[__obf_b97394cad50e288c] = __obf_488431258beaf3e9.__obf_5c19df52406936fd
			if __obf_ceb17d4721b21808.__obf_ef5ed10d7c6a1194 {
				__obf_5868836158514e2c.__obf_97e93c9eccbe5a21[__obf_b97394cad50e288c] = true
			}
			if int32(len(__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3)) <= __obf_b97394cad50e288c {
				__obf_5868836158514e2c.__obf_2e7a4d69845854bf(__obf_b97394cad50e288c + 1)
			}
			if __obf_5868836158514e2c.__obf_3c0899c6f8ed2d7d <= __obf_b97394cad50e288c {
				__obf_5868836158514e2c.__obf_3c0899c6f8ed2d7d = __obf_b97394cad50e288c + 1
			}
			__obf_bfd639f9c3be7e80 = append(__obf_bfd639f9c3be7e80, __obf_ecc42234bc4ab0ef{__obf_5c19df52406936fd: __obf_b97394cad50e288c, __obf_6deee49b8dd788d8: __obf_ceb17d4721b21808})
		}
	}
	return __obf_5868836158514e2c
}

func (__obf_5868836158514e2c *DATTrie) __obf_f8e2472f3b89a228(__obf_2d131eb3a30e7fc5 []rune) int32 {
	for __obf_da7ce5419e7c2ab3 := int32(1); ; __obf_da7ce5419e7c2ab3++ {
		__obf_5c2b0ea80ad819f1 := false
		for _, __obf_8c6f7d98a27e343f := range __obf_2d131eb3a30e7fc5 {
			__obf_ed1ddb116cf50801 := __obf_da7ce5419e7c2ab3 + int32(__obf_8c6f7d98a27e343f)
			if int32(len(__obf_5868836158514e2c.__obf_00026af41d7782b2)) <= __obf_ed1ddb116cf50801 {
				__obf_5868836158514e2c.__obf_2e7a4d69845854bf(__obf_ed1ddb116cf50801 + 1)
			}
			if __obf_5868836158514e2c.__obf_00026af41d7782b2[__obf_ed1ddb116cf50801] != 0 {
				__obf_5c2b0ea80ad819f1 = true
				break
			}
		}
		if !__obf_5c2b0ea80ad819f1 {
			return __obf_da7ce5419e7c2ab3
		}
	}
}

func (__obf_5868836158514e2c *DATTrie) __obf_2e7a4d69845854bf(__obf_1df64e12a3a5182f int32) {
	if __obf_1df64e12a3a5182f < int32(len(__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3)) {
		return
	}
	__obf_7b2c2bf38599b5cb := max(int32(len(__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3))+int32(len(__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3))/2, __obf_1df64e12a3a5182f)
	__obf_524a9ec0ee88eae3 := make([]int32, __obf_7b2c2bf38599b5cb)
	__obf_7011fff19d30c04a := make([]int32, __obf_7b2c2bf38599b5cb)
	__obf_9dd39681eeaed6d0 := make([]bool, __obf_7b2c2bf38599b5cb)
	copy(__obf_524a9ec0ee88eae3, __obf_5868836158514e2c.__obf_da7ce5419e7c2ab3)
	copy(__obf_7011fff19d30c04a, __obf_5868836158514e2c.__obf_00026af41d7782b2)
	copy(__obf_9dd39681eeaed6d0, __obf_5868836158514e2c.__obf_97e93c9eccbe5a21)
	__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3 = __obf_524a9ec0ee88eae3
	__obf_5868836158514e2c.__obf_00026af41d7782b2 = __obf_7011fff19d30c04a
	__obf_5868836158514e2c.__obf_97e93c9eccbe5a21 = __obf_9dd39681eeaed6d0
}

func (__obf_5868836158514e2c *DATTrie) __obf_74d639d3361fde50(__obf_1b0cd6b59e2eb7eb int32, __obf_8c6f7d98a27e343f rune) int32 {
	if __obf_1b0cd6b59e2eb7eb <= 0 || __obf_1b0cd6b59e2eb7eb >= int32(len(__obf_5868836158514e2c.__obf_da7ce5419e7c2ab3)) || __obf_5868836158514e2c.__obf_da7ce5419e7c2ab3[__obf_1b0cd6b59e2eb7eb] == 0 {
		return -1
	}
	__obf_ed1ddb116cf50801 := __obf_5868836158514e2c.__obf_da7ce5419e7c2ab3[__obf_1b0cd6b59e2eb7eb] + int32(__obf_8c6f7d98a27e343f)
	if __obf_ed1ddb116cf50801 > 0 && __obf_ed1ddb116cf50801 < int32(len(__obf_5868836158514e2c.__obf_00026af41d7782b2)) && __obf_5868836158514e2c.__obf_00026af41d7782b2[__obf_ed1ddb116cf50801] == __obf_1b0cd6b59e2eb7eb {
		return __obf_ed1ddb116cf50801
	}
	return -1
}

func (__obf_5868836158514e2c *DATTrie) __obf_ef5ed10d7c6a1194(__obf_1b0cd6b59e2eb7eb int32) bool {
	if __obf_1b0cd6b59e2eb7eb > 0 && __obf_1b0cd6b59e2eb7eb < int32(len(__obf_5868836158514e2c.__obf_97e93c9eccbe5a21)) {
		return __obf_5868836158514e2c.__obf_97e93c9eccbe5a21[__obf_1b0cd6b59e2eb7eb]
	}
	return false
}

func (__obf_5868836158514e2c *DATTrie) Contains(__obf_8fd45bd0d265008c string) bool {
	__obf_0e2c9fed14041ce5 := []rune(__obf_8fd45bd0d265008c)
	for __obf_74e4c62962a40491 := range __obf_0e2c9fed14041ce5 {
		__obf_6deee49b8dd788d8 := int32(1)
		for __obf_ec5c7f3e67ff13d2 := __obf_74e4c62962a40491; __obf_ec5c7f3e67ff13d2 < len(__obf_0e2c9fed14041ce5); __obf_ec5c7f3e67ff13d2++ {
			__obf_6deee49b8dd788d8 = __obf_5868836158514e2c.__obf_74d639d3361fde50(__obf_6deee49b8dd788d8, __obf_0e2c9fed14041ce5[__obf_ec5c7f3e67ff13d2])
			if __obf_6deee49b8dd788d8 == -1 {
				break
			}
			if __obf_5868836158514e2c.__obf_ef5ed10d7c6a1194(__obf_6deee49b8dd788d8) {
				return true
			}
		}
	}
	return false
}

func (__obf_5868836158514e2c *DATTrie) __obf_6ac3be995625316d(__obf_0e2c9fed14041ce5 []rune) []MatchResult {
	__obf_89f2c8dbe6260ce8 := make([]MatchResult, 0, __obf_b1462bf48eca6147)
	for __obf_74e4c62962a40491 := range __obf_0e2c9fed14041ce5 {
		__obf_6deee49b8dd788d8 := int32(1)
		for __obf_ec5c7f3e67ff13d2 := __obf_74e4c62962a40491; __obf_ec5c7f3e67ff13d2 < len(__obf_0e2c9fed14041ce5); __obf_ec5c7f3e67ff13d2++ {
			__obf_6deee49b8dd788d8 = __obf_5868836158514e2c.__obf_74d639d3361fde50(__obf_6deee49b8dd788d8, __obf_0e2c9fed14041ce5[__obf_ec5c7f3e67ff13d2])
			if __obf_6deee49b8dd788d8 == -1 {
				break
			}
			if __obf_5868836158514e2c.__obf_ef5ed10d7c6a1194(__obf_6deee49b8dd788d8) {
				__obf_89f2c8dbe6260ce8 = append(__obf_89f2c8dbe6260ce8, MatchResult{
					Word:  string(__obf_0e2c9fed14041ce5[__obf_74e4c62962a40491 : __obf_ec5c7f3e67ff13d2+1]),
					Start: __obf_74e4c62962a40491,
					End:   __obf_ec5c7f3e67ff13d2 + 1,
				})
			}
		}
	}
	return __obf_89f2c8dbe6260ce8
}

type WuManberMatcher struct {
	__obf_33207bb7dda88884 []string
	__obf_2569a480ce883fcb int
	__obf_693546b9cb409350 map[string]int32
	__obf_033004856de206c0 map[string][]int
	__obf_4642d9d31ba5ea32 int
}

func NewWuManberMatcher(__obf_2569a480ce883fcb int) *WuManberMatcher {
	return &WuManberMatcher{__obf_33207bb7dda88884: make([]string, 0), __obf_2569a480ce883fcb: __obf_2569a480ce883fcb, __obf_693546b9cb409350: make(map[string]int32), __obf_033004856de206c0: make(map[string][]int), __obf_4642d9d31ba5ea32: int(^uint(0) >> 1)}
}

func (__obf_b3b2ce14beb9c7ee *WuManberMatcher) AddPatterns(__obf_33207bb7dda88884 []string) {
	__obf_b3b2ce14beb9c7ee.__obf_33207bb7dda88884 = __obf_33207bb7dda88884
	for _, __obf_8f49d9c682b3f0da := range __obf_33207bb7dda88884 {
		__obf_e04796fdada640f9 := utf8.RuneCountInString(__obf_8f49d9c682b3f0da)
		if __obf_e04796fdada640f9 > 0 && __obf_e04796fdada640f9 < __obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 {
			__obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 = __obf_e04796fdada640f9
		}
	}
}

func (__obf_b3b2ce14beb9c7ee *WuManberMatcher) Build() {
	if len(__obf_b3b2ce14beb9c7ee.__obf_33207bb7dda88884) == 0 {
		return
	}
	if __obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 < __obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb {
		__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb = __obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32
	}
	if __obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb <= 0 {
		__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb = 1
	}

	for __obf_74e4c62962a40491, __obf_aeb38b7ab5c45dbb := range __obf_b3b2ce14beb9c7ee.__obf_33207bb7dda88884 {
		__obf_0e2c9fed14041ce5 := []rune(__obf_aeb38b7ab5c45dbb)
		__obf_85d67f13d49591ab := len(__obf_0e2c9fed14041ce5)
		for __obf_ec5c7f3e67ff13d2 := 0; __obf_ec5c7f3e67ff13d2 <= __obf_85d67f13d49591ab-__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb; __obf_ec5c7f3e67ff13d2++ {
			__obf_ec28b331372b9a42 := string(__obf_0e2c9fed14041ce5[__obf_ec5c7f3e67ff13d2 : __obf_ec5c7f3e67ff13d2+__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb])
			__obf_a8978473506acb08 := int32(__obf_85d67f13d49591ab - __obf_ec5c7f3e67ff13d2 - __obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb)
			if __obf_6bd0f348f93114f0, __obf_e9f356cc30ce1244 := __obf_b3b2ce14beb9c7ee.__obf_693546b9cb409350[__obf_ec28b331372b9a42]; !__obf_e9f356cc30ce1244 || __obf_a8978473506acb08 < __obf_6bd0f348f93114f0 {
				__obf_b3b2ce14beb9c7ee.__obf_693546b9cb409350[__obf_ec28b331372b9a42] = __obf_a8978473506acb08
			}
		}

		if __obf_85d67f13d49591ab >= __obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb {
			__obf_e05bcca3f4a38c43 := string(__obf_0e2c9fed14041ce5[__obf_85d67f13d49591ab-__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb:])
			__obf_b3b2ce14beb9c7ee.__obf_033004856de206c0[__obf_e05bcca3f4a38c43] = append(__obf_b3b2ce14beb9c7ee.__obf_033004856de206c0[__obf_e05bcca3f4a38c43], __obf_74e4c62962a40491)
		}
	}
}

func (__obf_b3b2ce14beb9c7ee *WuManberMatcher) Match(__obf_8fd45bd0d265008c string) bool {
	if len(__obf_b3b2ce14beb9c7ee.__obf_33207bb7dda88884) == 0 {
		return false
	}
	__obf_0e2c9fed14041ce5 := []rune(__obf_8fd45bd0d265008c)
	__obf_097c9ee75c2d3dba := len(__obf_0e2c9fed14041ce5)
	if __obf_097c9ee75c2d3dba < __obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 {
		return false
	}
	__obf_03793896c47f8d30 := __obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 - 1
	for __obf_03793896c47f8d30 < __obf_097c9ee75c2d3dba {
		__obf_f020b87544fd5832 := max(__obf_03793896c47f8d30-__obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb+1, 0)
		__obf_ec28b331372b9a42 := string(__obf_0e2c9fed14041ce5[__obf_f020b87544fd5832 : __obf_03793896c47f8d30+1])
		__obf_693546b9cb409350, __obf_e9f356cc30ce1244 := __obf_b3b2ce14beb9c7ee.__obf_693546b9cb409350[__obf_ec28b331372b9a42]
		if !__obf_e9f356cc30ce1244 {
			__obf_03793896c47f8d30 += (__obf_b3b2ce14beb9c7ee.__obf_4642d9d31ba5ea32 - __obf_b3b2ce14beb9c7ee.__obf_2569a480ce883fcb + 1)
			continue
		}
		if __obf_693546b9cb409350 == 0 {
			if __obf_38c4744101e4e0a9, __obf_de4222e243defcfd := __obf_b3b2ce14beb9c7ee.__obf_033004856de206c0[__obf_ec28b331372b9a42]; __obf_de4222e243defcfd {
				for _, __obf_33302d7af91ae0ae := range __obf_38c4744101e4e0a9 {
					__obf_aeb38b7ab5c45dbb := __obf_b3b2ce14beb9c7ee.__obf_33207bb7dda88884[__obf_33302d7af91ae0ae]
					__obf_89953d2883e95bd7 := utf8.RuneCountInString(__obf_aeb38b7ab5c45dbb)
					__obf_893eb2049b9e2d6c := __obf_03793896c47f8d30 - __obf_89953d2883e95bd7 + 1
					if __obf_893eb2049b9e2d6c >= 0 && string(__obf_0e2c9fed14041ce5[__obf_893eb2049b9e2d6c:__obf_893eb2049b9e2d6c+__obf_89953d2883e95bd7]) == __obf_aeb38b7ab5c45dbb {
						return true
					}
				}
			}
			__obf_03793896c47f8d30++
		} else {
			__obf_03793896c47f8d30 += int(__obf_693546b9cb409350)
		}
	}
	return false
}
