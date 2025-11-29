package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_f6a69b5d2ae02314(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_bde6dee98bfb656b := map[string]*Binding{}
	__obf_d31a49f9cbbb8716 := __obf_3b77508c34fb1648(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
		for _, __obf_42f1adaa610cc1a2 := range __obf_55bb392806565443.FromNames {
			__obf_3668b60d03008ac5 := __obf_bde6dee98bfb656b[__obf_42f1adaa610cc1a2]
			if __obf_3668b60d03008ac5 == nil {
				__obf_bde6dee98bfb656b[__obf_42f1adaa610cc1a2] = __obf_55bb392806565443
				continue
			}
			__obf_e29aa548946ea7a6, __obf_95fe2be3f6af92bc := __obf_5902434c0363a1d9(__obf_08da24be66d0d13c.__obf_5d13d7f3bd06c6cf, __obf_3668b60d03008ac5, __obf_55bb392806565443)
			if __obf_e29aa548946ea7a6 {
				delete(__obf_bde6dee98bfb656b, __obf_42f1adaa610cc1a2)
			}
			if !__obf_95fe2be3f6af92bc {
				__obf_bde6dee98bfb656b[__obf_42f1adaa610cc1a2] = __obf_55bb392806565443
			}
		}
	}
	__obf_db858406797b4036 := map[string]*__obf_2afabbd5d9366177{}
	for __obf_3a6754e3ef0a86a0, __obf_55bb392806565443 := range __obf_bde6dee98bfb656b {
		__obf_db858406797b4036[__obf_3a6754e3ef0a86a0] = __obf_55bb392806565443.Decoder.(*__obf_2afabbd5d9366177)
	}

	if !__obf_08da24be66d0d13c.__obf_f48e3198f571baa9() {
		for __obf_3a6754e3ef0a86a0, __obf_55bb392806565443 := range __obf_bde6dee98bfb656b {
			if _, __obf_98a3eefc4c187186 := __obf_db858406797b4036[strings.ToLower(__obf_3a6754e3ef0a86a0)]; !__obf_98a3eefc4c187186 {
				__obf_db858406797b4036[strings.ToLower(__obf_3a6754e3ef0a86a0)] = __obf_55bb392806565443.Decoder.(*__obf_2afabbd5d9366177)
			}
		}
	}

	return __obf_273ed6746ea197b1(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133, __obf_db858406797b4036)
}

func __obf_273ed6746ea197b1(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type, __obf_db858406797b4036 map[string]*__obf_2afabbd5d9366177) ValDecoder {
	if __obf_08da24be66d0d13c.__obf_c395639ca05d31b5 {
		return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133: __obf_5efc66d8c338c133, __obf_db858406797b4036: __obf_db858406797b4036, __obf_c395639ca05d31b5: true}
	}
	__obf_bf98e3765005f117 := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_db858406797b4036) {
	case 0:
		return &__obf_0dd7af3b2984b4c9{__obf_5efc66d8c338c133}
	case 1:
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			return &__obf_d33a6078253496a3{__obf_5efc66d8c338c133, __obf_545773ac0ee62760, __obf_a9e04e7d0fe27e22}
		}
	case 2:
		var __obf_4c3c654d50bddd40 int64
		var __obf_78de1725f89e072e int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_4c3c654d50bddd40 == 0 {
				__obf_4c3c654d50bddd40 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else {
				__obf_78de1725f89e072e = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_c8971356553f1a66{__obf_5efc66d8c338c133, __obf_4c3c654d50bddd40, __obf_3df72ca96c7c013b, __obf_78de1725f89e072e, __obf_93f28a01ca302684}
	case 3:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_d59745d5652ff9c0{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36}
	case 4:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_e5da5edd727c721b{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529}
	case 5:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_52198bec7e39050b{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc}
	case 6:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_15dede546e753c13 int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		var __obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else if __obf_ebc281a1fd3ee3b0 == 0 {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			} else {
				__obf_15dede546e753c13 = __obf_545773ac0ee62760
				__obf_ca35ba59b432c267 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_f454fe52aafe747b{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc, __obf_15dede546e753c13, __obf_ca35ba59b432c267}
	case 7:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_15dede546e753c13 int64
		var __obf_4fdcbe49d4a39a9e int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		var __obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
		var __obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else if __obf_ebc281a1fd3ee3b0 == 0 {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			} else if __obf_15dede546e753c13 == 0 {
				__obf_15dede546e753c13 = __obf_545773ac0ee62760
				__obf_ca35ba59b432c267 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_4fdcbe49d4a39a9e = __obf_545773ac0ee62760
				__obf_d4b6ace4e6431bd2 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_afaef975c4d36468{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc, __obf_15dede546e753c13, __obf_ca35ba59b432c267, __obf_4fdcbe49d4a39a9e, __obf_d4b6ace4e6431bd2}
	case 8:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_15dede546e753c13 int64
		var __obf_4fdcbe49d4a39a9e int64
		var __obf_3aa49820e178687c int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		var __obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
		var __obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
		var __obf_3e665dd563de382d *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else if __obf_ebc281a1fd3ee3b0 == 0 {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			} else if __obf_15dede546e753c13 == 0 {
				__obf_15dede546e753c13 = __obf_545773ac0ee62760
				__obf_ca35ba59b432c267 = __obf_a9e04e7d0fe27e22
			} else if __obf_4fdcbe49d4a39a9e == 0 {
				__obf_4fdcbe49d4a39a9e = __obf_545773ac0ee62760
				__obf_d4b6ace4e6431bd2 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_3aa49820e178687c = __obf_545773ac0ee62760
				__obf_3e665dd563de382d = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_800b1fafc62885e8{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc, __obf_15dede546e753c13, __obf_ca35ba59b432c267, __obf_4fdcbe49d4a39a9e, __obf_d4b6ace4e6431bd2, __obf_3aa49820e178687c, __obf_3e665dd563de382d}
	case 9:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_15dede546e753c13 int64
		var __obf_4fdcbe49d4a39a9e int64
		var __obf_3aa49820e178687c int64
		var __obf_bf11880038521597 int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		var __obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
		var __obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
		var __obf_3e665dd563de382d *__obf_2afabbd5d9366177
		var __obf_1fb42a03fce1f8d5 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else if __obf_ebc281a1fd3ee3b0 == 0 {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			} else if __obf_15dede546e753c13 == 0 {
				__obf_15dede546e753c13 = __obf_545773ac0ee62760
				__obf_ca35ba59b432c267 = __obf_a9e04e7d0fe27e22
			} else if __obf_4fdcbe49d4a39a9e == 0 {
				__obf_4fdcbe49d4a39a9e = __obf_545773ac0ee62760
				__obf_d4b6ace4e6431bd2 = __obf_a9e04e7d0fe27e22
			} else if __obf_3aa49820e178687c == 0 {
				__obf_3aa49820e178687c = __obf_545773ac0ee62760
				__obf_3e665dd563de382d = __obf_a9e04e7d0fe27e22
			} else {
				__obf_bf11880038521597 = __obf_545773ac0ee62760
				__obf_1fb42a03fce1f8d5 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_0e5eee2b03962bc2{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc, __obf_15dede546e753c13, __obf_ca35ba59b432c267, __obf_4fdcbe49d4a39a9e, __obf_d4b6ace4e6431bd2, __obf_3aa49820e178687c, __obf_3e665dd563de382d, __obf_bf11880038521597, __obf_1fb42a03fce1f8d5}
	case 10:
		var __obf_b0f524a8d48ebcf7 int64
		var __obf_96debf7735b42c01 int64
		var __obf_9bc0c65543f24b26 int64
		var __obf_e8d92e959a208e15 int64
		var __obf_ebc281a1fd3ee3b0 int64
		var __obf_15dede546e753c13 int64
		var __obf_4fdcbe49d4a39a9e int64
		var __obf_3aa49820e178687c int64
		var __obf_bf11880038521597 int64
		var __obf_92ba9ee5d0f66b3d int64
		var __obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
		var __obf_93f28a01ca302684 *__obf_2afabbd5d9366177
		var __obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
		var __obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
		var __obf_cc84aff481e666cc *__obf_2afabbd5d9366177
		var __obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
		var __obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
		var __obf_3e665dd563de382d *__obf_2afabbd5d9366177
		var __obf_1fb42a03fce1f8d5 *__obf_2afabbd5d9366177
		var __obf_23e956684e6ad6d3 *__obf_2afabbd5d9366177
		for __obf_d776b95c6ab08d8c, __obf_a9e04e7d0fe27e22 := range __obf_db858406797b4036 {
			__obf_545773ac0ee62760 := __obf_7dc2f08aee807385(__obf_d776b95c6ab08d8c, __obf_08da24be66d0d13c.__obf_f48e3198f571baa9())
			_, __obf_d065f21b32308a00 := __obf_bf98e3765005f117[__obf_545773ac0ee62760]
			if __obf_d065f21b32308a00 {
				return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
			}
			__obf_bf98e3765005f117[__obf_545773ac0ee62760] = struct{}{}
			if __obf_b0f524a8d48ebcf7 == 0 {
				__obf_b0f524a8d48ebcf7 = __obf_545773ac0ee62760
				__obf_3df72ca96c7c013b = __obf_a9e04e7d0fe27e22
			} else if __obf_96debf7735b42c01 == 0 {
				__obf_96debf7735b42c01 = __obf_545773ac0ee62760
				__obf_93f28a01ca302684 = __obf_a9e04e7d0fe27e22
			} else if __obf_9bc0c65543f24b26 == 0 {
				__obf_9bc0c65543f24b26 = __obf_545773ac0ee62760
				__obf_e7c837df1ca4ef36 = __obf_a9e04e7d0fe27e22
			} else if __obf_e8d92e959a208e15 == 0 {
				__obf_e8d92e959a208e15 = __obf_545773ac0ee62760
				__obf_c07b964d99eeb529 = __obf_a9e04e7d0fe27e22
			} else if __obf_ebc281a1fd3ee3b0 == 0 {
				__obf_ebc281a1fd3ee3b0 = __obf_545773ac0ee62760
				__obf_cc84aff481e666cc = __obf_a9e04e7d0fe27e22
			} else if __obf_15dede546e753c13 == 0 {
				__obf_15dede546e753c13 = __obf_545773ac0ee62760
				__obf_ca35ba59b432c267 = __obf_a9e04e7d0fe27e22
			} else if __obf_4fdcbe49d4a39a9e == 0 {
				__obf_4fdcbe49d4a39a9e = __obf_545773ac0ee62760
				__obf_d4b6ace4e6431bd2 = __obf_a9e04e7d0fe27e22
			} else if __obf_3aa49820e178687c == 0 {
				__obf_3aa49820e178687c = __obf_545773ac0ee62760
				__obf_3e665dd563de382d = __obf_a9e04e7d0fe27e22
			} else if __obf_bf11880038521597 == 0 {
				__obf_bf11880038521597 = __obf_545773ac0ee62760
				__obf_1fb42a03fce1f8d5 = __obf_a9e04e7d0fe27e22
			} else {
				__obf_92ba9ee5d0f66b3d = __obf_545773ac0ee62760
				__obf_23e956684e6ad6d3 = __obf_a9e04e7d0fe27e22
			}
		}
		return &__obf_882c308aef13eabb{__obf_5efc66d8c338c133, __obf_b0f524a8d48ebcf7, __obf_3df72ca96c7c013b, __obf_96debf7735b42c01, __obf_93f28a01ca302684, __obf_9bc0c65543f24b26, __obf_e7c837df1ca4ef36, __obf_e8d92e959a208e15, __obf_c07b964d99eeb529, __obf_ebc281a1fd3ee3b0, __obf_cc84aff481e666cc, __obf_15dede546e753c13, __obf_ca35ba59b432c267, __obf_4fdcbe49d4a39a9e, __obf_d4b6ace4e6431bd2, __obf_3aa49820e178687c, __obf_3e665dd563de382d, __obf_bf11880038521597, __obf_1fb42a03fce1f8d5, __obf_92ba9ee5d0f66b3d, __obf_23e956684e6ad6d3}
	}
	return &__obf_54bf1fd12a2537b4{__obf_5efc66d8c338c133, __obf_db858406797b4036, false}
}

type __obf_54bf1fd12a2537b4 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_db858406797b4036 map[string]*__obf_2afabbd5d9366177
	__obf_c395639ca05d31b5 bool
}

func (__obf_18f746d7b5b440ee *__obf_54bf1fd12a2537b4) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	var __obf_dab9baaadfa7c8c2 byte
	for __obf_dab9baaadfa7c8c2 = ','; __obf_dab9baaadfa7c8c2 == ','; __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() {
		__obf_18f746d7b5b440ee.__obf_677dae316fac1ac3(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	if __obf_dab9baaadfa7c8c2 != '}' {
		__obf_67008a6a9e5ba828.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

func (__obf_18f746d7b5b440ee *__obf_54bf1fd12a2537b4) __obf_677dae316fac1ac3(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	var __obf_61998fb083387c3c string
	var __obf_a9e04e7d0fe27e22 *__obf_2afabbd5d9366177
	if __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_35f74bd8c14f5694 {
		__obf_e74553e37b233334 := __obf_67008a6a9e5ba828.ReadStringAsSlice()
		__obf_61998fb083387c3c = *(*string)(unsafe.Pointer(&__obf_e74553e37b233334))
		__obf_a9e04e7d0fe27e22 = __obf_18f746d7b5b440ee.__obf_db858406797b4036[__obf_61998fb083387c3c]
		if __obf_a9e04e7d0fe27e22 == nil && !__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_f48e3198f571baa9 {
			__obf_a9e04e7d0fe27e22 = __obf_18f746d7b5b440ee.__obf_db858406797b4036[strings.ToLower(__obf_61998fb083387c3c)]
		}
	} else {
		__obf_61998fb083387c3c = __obf_67008a6a9e5ba828.ReadString()
		__obf_a9e04e7d0fe27e22 = __obf_18f746d7b5b440ee.__obf_db858406797b4036[__obf_61998fb083387c3c]
		if __obf_a9e04e7d0fe27e22 == nil && !__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_f48e3198f571baa9 {
			__obf_a9e04e7d0fe27e22 = __obf_18f746d7b5b440ee.__obf_db858406797b4036[strings.ToLower(__obf_61998fb083387c3c)]
		}
	}
	if __obf_a9e04e7d0fe27e22 == nil {
		if __obf_18f746d7b5b440ee.__obf_c395639ca05d31b5 {
			__obf_599756f70dc37f5f := "found unknown field: " + __obf_61998fb083387c3c
			__obf_67008a6a9e5ba828.
				ReportError("ReadObject", __obf_599756f70dc37f5f)
		}
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 != ':' {
			__obf_67008a6a9e5ba828.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		}
		__obf_67008a6a9e5ba828.
			Skip()
		return
	}
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != ':' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
	}
	__obf_a9e04e7d0fe27e22.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
}

type __obf_0dd7af3b2984b4c9 struct {
	__obf_5efc66d8c338c133 reflect2.Type
}

func (__obf_18f746d7b5b440ee *__obf_0dd7af3b2984b4c9) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_9ad2e389a2a6fdb8 := __obf_67008a6a9e5ba828.WhatIsNext()
	if __obf_9ad2e389a2a6fdb8 != ObjectValue && __obf_9ad2e389a2a6fdb8 != NilValue {
		__obf_67008a6a9e5ba828.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_67008a6a9e5ba828.
		Skip()
}

type __obf_d33a6078253496a3 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_545773ac0ee62760 int64
	__obf_a9e04e7d0fe27e22 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_d33a6078253496a3) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		if __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() == __obf_18f746d7b5b440ee.__obf_545773ac0ee62760 {
			__obf_18f746d7b5b440ee.__obf_a9e04e7d0fe27e22.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		} else {
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_c8971356553f1a66 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_c8971356553f1a66) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_d59745d5652ff9c0 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_d59745d5652ff9c0) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_e5da5edd727c721b struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_e5da5edd727c721b) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_52198bec7e39050b struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_52198bec7e39050b) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_f454fe52aafe747b struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
	__obf_d568364d3b979a08 int64
	__obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_f454fe52aafe747b) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d568364d3b979a08:
			__obf_18f746d7b5b440ee.__obf_ca35ba59b432c267.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_afaef975c4d36468 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
	__obf_d568364d3b979a08 int64
	__obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
	__obf_a0de3592366add5f int64
	__obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_afaef975c4d36468) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d568364d3b979a08:
			__obf_18f746d7b5b440ee.__obf_ca35ba59b432c267.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_a0de3592366add5f:
			__obf_18f746d7b5b440ee.__obf_d4b6ace4e6431bd2.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_800b1fafc62885e8 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
	__obf_d568364d3b979a08 int64
	__obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
	__obf_a0de3592366add5f int64
	__obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
	__obf_d741bcb3c2e7745d int64
	__obf_3e665dd563de382d *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_800b1fafc62885e8) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d568364d3b979a08:
			__obf_18f746d7b5b440ee.__obf_ca35ba59b432c267.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_a0de3592366add5f:
			__obf_18f746d7b5b440ee.__obf_d4b6ace4e6431bd2.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d741bcb3c2e7745d:
			__obf_18f746d7b5b440ee.__obf_3e665dd563de382d.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_0e5eee2b03962bc2 struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
	__obf_d568364d3b979a08 int64
	__obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
	__obf_a0de3592366add5f int64
	__obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
	__obf_d741bcb3c2e7745d int64
	__obf_3e665dd563de382d *__obf_2afabbd5d9366177
	__obf_cf51f6a9581662b2 int64
	__obf_1fb42a03fce1f8d5 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_0e5eee2b03962bc2) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d568364d3b979a08:
			__obf_18f746d7b5b440ee.__obf_ca35ba59b432c267.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_a0de3592366add5f:
			__obf_18f746d7b5b440ee.__obf_d4b6ace4e6431bd2.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d741bcb3c2e7745d:
			__obf_18f746d7b5b440ee.__obf_3e665dd563de382d.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_cf51f6a9581662b2:
			__obf_18f746d7b5b440ee.__obf_1fb42a03fce1f8d5.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_882c308aef13eabb struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_4c3c654d50bddd40 int64
	__obf_3df72ca96c7c013b *__obf_2afabbd5d9366177
	__obf_78de1725f89e072e int64
	__obf_93f28a01ca302684 *__obf_2afabbd5d9366177
	__obf_b8f73fcb6a19f274 int64
	__obf_e7c837df1ca4ef36 *__obf_2afabbd5d9366177
	__obf_4f09ca9b11a0a6dd int64
	__obf_c07b964d99eeb529 *__obf_2afabbd5d9366177
	__obf_f58bad7faa387f48 int64
	__obf_cc84aff481e666cc *__obf_2afabbd5d9366177
	__obf_d568364d3b979a08 int64
	__obf_ca35ba59b432c267 *__obf_2afabbd5d9366177
	__obf_a0de3592366add5f int64
	__obf_d4b6ace4e6431bd2 *__obf_2afabbd5d9366177
	__obf_d741bcb3c2e7745d int64
	__obf_3e665dd563de382d *__obf_2afabbd5d9366177
	__obf_cf51f6a9581662b2 int64
	__obf_1fb42a03fce1f8d5 *__obf_2afabbd5d9366177
	__obf_3ff1efa5e33f1a42 int64
	__obf_23e956684e6ad6d3 *__obf_2afabbd5d9366177
}

func (__obf_18f746d7b5b440ee *__obf_882c308aef13eabb) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.__obf_b8169430213031d6() {
		return
	}
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		switch __obf_67008a6a9e5ba828.__obf_64030aba1e95f95b() {
		case __obf_18f746d7b5b440ee.__obf_4c3c654d50bddd40:
			__obf_18f746d7b5b440ee.__obf_3df72ca96c7c013b.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_78de1725f89e072e:
			__obf_18f746d7b5b440ee.__obf_93f28a01ca302684.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_b8f73fcb6a19f274:
			__obf_18f746d7b5b440ee.__obf_e7c837df1ca4ef36.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_4f09ca9b11a0a6dd:
			__obf_18f746d7b5b440ee.__obf_c07b964d99eeb529.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_f58bad7faa387f48:
			__obf_18f746d7b5b440ee.__obf_cc84aff481e666cc.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d568364d3b979a08:
			__obf_18f746d7b5b440ee.__obf_ca35ba59b432c267.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_a0de3592366add5f:
			__obf_18f746d7b5b440ee.__obf_d4b6ace4e6431bd2.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_d741bcb3c2e7745d:
			__obf_18f746d7b5b440ee.__obf_3e665dd563de382d.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_cf51f6a9581662b2:
			__obf_18f746d7b5b440ee.__obf_1fb42a03fce1f8d5.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		case __obf_18f746d7b5b440ee.__obf_3ff1efa5e33f1a42:
			__obf_18f746d7b5b440ee.__obf_23e956684e6ad6d3.
				Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		default:
			__obf_67008a6a9e5ba828.
				Skip()
		}
		if __obf_67008a6a9e5ba828.__obf_cddc9a4e84baa583() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF && len(__obf_18f746d7b5b440ee.__obf_5efc66d8c338c133.Type1().Name()) != 0 {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v.%s", __obf_18f746d7b5b440ee.__obf_5efc66d8c338c133, __obf_67008a6a9e5ba828.Error.Error())
	}
	__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
}

type __obf_2afabbd5d9366177 struct {
	__obf_61998fb083387c3c reflect2.StructField
	__obf_a9e04e7d0fe27e22 ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_2afabbd5d9366177) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_a6502f11c1c12aab := __obf_18f746d7b5b440ee.__obf_61998fb083387c3c.UnsafeGet(__obf_d3c919547bf7e47a)
	__obf_18f746d7b5b440ee.__obf_a9e04e7d0fe27e22.
		Decode(__obf_a6502f11c1c12aab, __obf_67008a6a9e5ba828)
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%s: %s", __obf_18f746d7b5b440ee.__obf_61998fb083387c3c.Name(), __obf_67008a6a9e5ba828.Error.Error())
	}
}

type __obf_76d674defb583cf0 struct {
	__obf_45aaf5212411947b ValDecoder
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
}

func (__obf_18f746d7b5b440ee *__obf_76d674defb583cf0) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	__obf_12c21b79fa86dcba := *((*string)(__obf_d3c919547bf7e47a))
	__obf_7c43ffb9668925d3 := __obf_18f746d7b5b440ee.__obf_dca106e1cdf85392.BorrowIterator([]byte(__obf_12c21b79fa86dcba))
	defer __obf_18f746d7b5b440ee.__obf_dca106e1cdf85392.ReturnIterator(__obf_7c43ffb9668925d3)
	*((*string)(__obf_d3c919547bf7e47a)) = __obf_7c43ffb9668925d3.ReadString()
}

type __obf_dcbce58ec66b632d struct {
	__obf_45aaf5212411947b ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_dcbce58ec66b632d) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.WhatIsNext() == NilValue {
		__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
			Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
		return
	}
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != '"' {
		__obf_67008a6a9e5ba828.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	if __obf_67008a6a9e5ba828.Error != nil {
		return
	}
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
	if __obf_dab9baaadfa7c8c2 != '"' {
		__obf_67008a6a9e5ba828.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
}
