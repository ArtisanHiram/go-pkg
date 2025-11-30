package __obf_703d23ba520c3295

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_be37b1a24a3a47e8(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_655058327e5c073a := map[string]*Binding{}
	__obf_ab1ac896d4aa86ee := __obf_969b50c34ac12a33(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
		for _, __obf_64f91a0cf4b12d6c := range __obf_af9543e035a21def.FromNames {
			__obf_3dafc9b4a2f87f34 := __obf_655058327e5c073a[__obf_64f91a0cf4b12d6c]
			if __obf_3dafc9b4a2f87f34 == nil {
				__obf_655058327e5c073a[__obf_64f91a0cf4b12d6c] = __obf_af9543e035a21def
				continue
			}
			__obf_a63abce1e9e94c66, __obf_8f285a2701f37644 := __obf_9de3265131b75a41(__obf_2aaf7367de3ff86d.__obf_ef74248d8ae9ba78, __obf_3dafc9b4a2f87f34, __obf_af9543e035a21def)
			if __obf_a63abce1e9e94c66 {
				delete(__obf_655058327e5c073a, __obf_64f91a0cf4b12d6c)
			}
			if !__obf_8f285a2701f37644 {
				__obf_655058327e5c073a[__obf_64f91a0cf4b12d6c] = __obf_af9543e035a21def
			}
		}
	}
	__obf_f4d348cda1281225 := map[string]*__obf_5b7a5815a547f104{}
	for __obf_a95ada6cc83385bf, __obf_af9543e035a21def := range __obf_655058327e5c073a {
		__obf_f4d348cda1281225[__obf_a95ada6cc83385bf] = __obf_af9543e035a21def.Decoder.(*__obf_5b7a5815a547f104)
	}

	if !__obf_2aaf7367de3ff86d.__obf_9332366e75c2e023() {
		for __obf_a95ada6cc83385bf, __obf_af9543e035a21def := range __obf_655058327e5c073a {
			if _, __obf_87a836f73d31233b := __obf_f4d348cda1281225[strings.ToLower(__obf_a95ada6cc83385bf)]; !__obf_87a836f73d31233b {
				__obf_f4d348cda1281225[strings.ToLower(__obf_a95ada6cc83385bf)] = __obf_af9543e035a21def.Decoder.(*__obf_5b7a5815a547f104)
			}
		}
	}

	return __obf_88f7137bf8e70471(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e, __obf_f4d348cda1281225)
}

func __obf_88f7137bf8e70471(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type, __obf_f4d348cda1281225 map[string]*__obf_5b7a5815a547f104) ValDecoder {
	if __obf_2aaf7367de3ff86d.__obf_f3aed2fd747dcb8e {
		return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e: __obf_3b7f6abbae19451e, __obf_f4d348cda1281225: __obf_f4d348cda1281225, __obf_f3aed2fd747dcb8e: true}
	}
	__obf_f84eed5fee1614c2 := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_f4d348cda1281225) {
	case 0:
		return &__obf_87c01baa0aecd1de{__obf_3b7f6abbae19451e}
	case 1:
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			return &__obf_9d5b95b9ade9b61b{__obf_3b7f6abbae19451e, __obf_904d17a1db8bff22, __obf_5e70cce5680d61b6}
		}
	case 2:
		var __obf_218d368357bba003 int64
		var __obf_fca16fa0ada58d69 int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_218d368357bba003 == 0 {
				__obf_218d368357bba003 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else {
				__obf_fca16fa0ada58d69 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_62cdf53760ce8958{__obf_3b7f6abbae19451e, __obf_218d368357bba003, __obf_c068c4315afd5733, __obf_fca16fa0ada58d69, __obf_079fb380254cbbda}
	case 3:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_6f5b316f8a3ea115{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d}
	case 4:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_42adb9ad608cad61{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7}
	case 5:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_e4692cbb7ed95cb5{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0}
	case 6:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_84eba590dc004df9 int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		var __obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else if __obf_e48b8f9d0e271abf == 0 {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			} else {
				__obf_84eba590dc004df9 = __obf_904d17a1db8bff22
				__obf_f4a053b4639e9514 = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_9deaed8c9e04f68f{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0, __obf_84eba590dc004df9, __obf_f4a053b4639e9514}
	case 7:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_84eba590dc004df9 int64
		var __obf_da01ead606d2eb72 int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		var __obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
		var __obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else if __obf_e48b8f9d0e271abf == 0 {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			} else if __obf_84eba590dc004df9 == 0 {
				__obf_84eba590dc004df9 = __obf_904d17a1db8bff22
				__obf_f4a053b4639e9514 = __obf_5e70cce5680d61b6
			} else {
				__obf_da01ead606d2eb72 = __obf_904d17a1db8bff22
				__obf_eb37d3b9c8a6969f = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_b4e105b6dad7055e{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0, __obf_84eba590dc004df9, __obf_f4a053b4639e9514, __obf_da01ead606d2eb72, __obf_eb37d3b9c8a6969f}
	case 8:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_84eba590dc004df9 int64
		var __obf_da01ead606d2eb72 int64
		var __obf_c1a809415d03a8de int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		var __obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
		var __obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
		var __obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else if __obf_e48b8f9d0e271abf == 0 {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			} else if __obf_84eba590dc004df9 == 0 {
				__obf_84eba590dc004df9 = __obf_904d17a1db8bff22
				__obf_f4a053b4639e9514 = __obf_5e70cce5680d61b6
			} else if __obf_da01ead606d2eb72 == 0 {
				__obf_da01ead606d2eb72 = __obf_904d17a1db8bff22
				__obf_eb37d3b9c8a6969f = __obf_5e70cce5680d61b6
			} else {
				__obf_c1a809415d03a8de = __obf_904d17a1db8bff22
				__obf_67871318f1fb2d29 = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_308e238a1d242ed4{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0, __obf_84eba590dc004df9, __obf_f4a053b4639e9514, __obf_da01ead606d2eb72, __obf_eb37d3b9c8a6969f, __obf_c1a809415d03a8de, __obf_67871318f1fb2d29}
	case 9:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_84eba590dc004df9 int64
		var __obf_da01ead606d2eb72 int64
		var __obf_c1a809415d03a8de int64
		var __obf_0c610b08f69141da int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		var __obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
		var __obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
		var __obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
		var __obf_de8ef2e6dc722a8a *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else if __obf_e48b8f9d0e271abf == 0 {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			} else if __obf_84eba590dc004df9 == 0 {
				__obf_84eba590dc004df9 = __obf_904d17a1db8bff22
				__obf_f4a053b4639e9514 = __obf_5e70cce5680d61b6
			} else if __obf_da01ead606d2eb72 == 0 {
				__obf_da01ead606d2eb72 = __obf_904d17a1db8bff22
				__obf_eb37d3b9c8a6969f = __obf_5e70cce5680d61b6
			} else if __obf_c1a809415d03a8de == 0 {
				__obf_c1a809415d03a8de = __obf_904d17a1db8bff22
				__obf_67871318f1fb2d29 = __obf_5e70cce5680d61b6
			} else {
				__obf_0c610b08f69141da = __obf_904d17a1db8bff22
				__obf_de8ef2e6dc722a8a = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_185ac0b6849f5150{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0, __obf_84eba590dc004df9, __obf_f4a053b4639e9514, __obf_da01ead606d2eb72, __obf_eb37d3b9c8a6969f, __obf_c1a809415d03a8de, __obf_67871318f1fb2d29, __obf_0c610b08f69141da, __obf_de8ef2e6dc722a8a}
	case 10:
		var __obf_77d71040cea0dc43 int64
		var __obf_778f0f04a1b033a6 int64
		var __obf_c601d0a45044d515 int64
		var __obf_6353927f3367305f int64
		var __obf_e48b8f9d0e271abf int64
		var __obf_84eba590dc004df9 int64
		var __obf_da01ead606d2eb72 int64
		var __obf_c1a809415d03a8de int64
		var __obf_0c610b08f69141da int64
		var __obf_00de187f71c32af8 int64
		var __obf_c068c4315afd5733 *__obf_5b7a5815a547f104
		var __obf_079fb380254cbbda *__obf_5b7a5815a547f104
		var __obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
		var __obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
		var __obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
		var __obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
		var __obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
		var __obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
		var __obf_de8ef2e6dc722a8a *__obf_5b7a5815a547f104
		var __obf_94405688d8ad81bc *__obf_5b7a5815a547f104
		for __obf_f8b76c1746a1e630, __obf_5e70cce5680d61b6 := range __obf_f4d348cda1281225 {
			__obf_904d17a1db8bff22 := __obf_818f762e1b421403(__obf_f8b76c1746a1e630, __obf_2aaf7367de3ff86d.__obf_9332366e75c2e023())
			_, __obf_d950ed5c40449d9b := __obf_f84eed5fee1614c2[__obf_904d17a1db8bff22]
			if __obf_d950ed5c40449d9b {
				return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
			}
			__obf_f84eed5fee1614c2[__obf_904d17a1db8bff22] = struct{}{}
			if __obf_77d71040cea0dc43 == 0 {
				__obf_77d71040cea0dc43 = __obf_904d17a1db8bff22
				__obf_c068c4315afd5733 = __obf_5e70cce5680d61b6
			} else if __obf_778f0f04a1b033a6 == 0 {
				__obf_778f0f04a1b033a6 = __obf_904d17a1db8bff22
				__obf_079fb380254cbbda = __obf_5e70cce5680d61b6
			} else if __obf_c601d0a45044d515 == 0 {
				__obf_c601d0a45044d515 = __obf_904d17a1db8bff22
				__obf_388ade5cdd5ca15d = __obf_5e70cce5680d61b6
			} else if __obf_6353927f3367305f == 0 {
				__obf_6353927f3367305f = __obf_904d17a1db8bff22
				__obf_52a24b5a0d6105f7 = __obf_5e70cce5680d61b6
			} else if __obf_e48b8f9d0e271abf == 0 {
				__obf_e48b8f9d0e271abf = __obf_904d17a1db8bff22
				__obf_b7955ba8aad92db0 = __obf_5e70cce5680d61b6
			} else if __obf_84eba590dc004df9 == 0 {
				__obf_84eba590dc004df9 = __obf_904d17a1db8bff22
				__obf_f4a053b4639e9514 = __obf_5e70cce5680d61b6
			} else if __obf_da01ead606d2eb72 == 0 {
				__obf_da01ead606d2eb72 = __obf_904d17a1db8bff22
				__obf_eb37d3b9c8a6969f = __obf_5e70cce5680d61b6
			} else if __obf_c1a809415d03a8de == 0 {
				__obf_c1a809415d03a8de = __obf_904d17a1db8bff22
				__obf_67871318f1fb2d29 = __obf_5e70cce5680d61b6
			} else if __obf_0c610b08f69141da == 0 {
				__obf_0c610b08f69141da = __obf_904d17a1db8bff22
				__obf_de8ef2e6dc722a8a = __obf_5e70cce5680d61b6
			} else {
				__obf_00de187f71c32af8 = __obf_904d17a1db8bff22
				__obf_94405688d8ad81bc = __obf_5e70cce5680d61b6
			}
		}
		return &__obf_b23dd4c081ae0650{__obf_3b7f6abbae19451e, __obf_77d71040cea0dc43, __obf_c068c4315afd5733, __obf_778f0f04a1b033a6, __obf_079fb380254cbbda, __obf_c601d0a45044d515, __obf_388ade5cdd5ca15d, __obf_6353927f3367305f, __obf_52a24b5a0d6105f7, __obf_e48b8f9d0e271abf, __obf_b7955ba8aad92db0, __obf_84eba590dc004df9, __obf_f4a053b4639e9514, __obf_da01ead606d2eb72, __obf_eb37d3b9c8a6969f, __obf_c1a809415d03a8de, __obf_67871318f1fb2d29, __obf_0c610b08f69141da, __obf_de8ef2e6dc722a8a, __obf_00de187f71c32af8, __obf_94405688d8ad81bc}
	}
	return &__obf_ca6e5494c85f479a{__obf_3b7f6abbae19451e, __obf_f4d348cda1281225, false}
}

type __obf_ca6e5494c85f479a struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_f4d348cda1281225 map[string]*__obf_5b7a5815a547f104
	__obf_f3aed2fd747dcb8e bool
}

func (__obf_c9719e68325f7d44 *__obf_ca6e5494c85f479a) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	var __obf_bd08f5161fff294a byte
	for __obf_bd08f5161fff294a = ','; __obf_bd08f5161fff294a == ','; __obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec() {
		__obf_c9719e68325f7d44.__obf_d7248237c99a0d10(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	if __obf_bd08f5161fff294a != '}' {
		__obf_47edab4c16a2d88d.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_bd08f5161fff294a}))
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

func (__obf_c9719e68325f7d44 *__obf_ca6e5494c85f479a) __obf_d7248237c99a0d10(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	var __obf_13f6440419533990 string
	var __obf_5e70cce5680d61b6 *__obf_5b7a5815a547f104
	if __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_c8260e5b2e39ae59 {
		__obf_2d81b0dfac4d4c30 := __obf_47edab4c16a2d88d.ReadStringAsSlice()
		__obf_13f6440419533990 = *(*string)(unsafe.Pointer(&__obf_2d81b0dfac4d4c30))
		__obf_5e70cce5680d61b6 = __obf_c9719e68325f7d44.__obf_f4d348cda1281225[__obf_13f6440419533990]
		if __obf_5e70cce5680d61b6 == nil && !__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_9332366e75c2e023 {
			__obf_5e70cce5680d61b6 = __obf_c9719e68325f7d44.__obf_f4d348cda1281225[strings.ToLower(__obf_13f6440419533990)]
		}
	} else {
		__obf_13f6440419533990 = __obf_47edab4c16a2d88d.ReadString()
		__obf_5e70cce5680d61b6 = __obf_c9719e68325f7d44.__obf_f4d348cda1281225[__obf_13f6440419533990]
		if __obf_5e70cce5680d61b6 == nil && !__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_9332366e75c2e023 {
			__obf_5e70cce5680d61b6 = __obf_c9719e68325f7d44.__obf_f4d348cda1281225[strings.ToLower(__obf_13f6440419533990)]
		}
	}
	if __obf_5e70cce5680d61b6 == nil {
		if __obf_c9719e68325f7d44.__obf_f3aed2fd747dcb8e {
			__obf_c6f61ce4298691bb := "found unknown field: " + __obf_13f6440419533990
			__obf_47edab4c16a2d88d.
				ReportError("ReadObject", __obf_c6f61ce4298691bb)
		}
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a != ':' {
			__obf_47edab4c16a2d88d.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
		}
		__obf_47edab4c16a2d88d.
			Skip()
		return
	}
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != ':' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
	}
	__obf_5e70cce5680d61b6.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
}

type __obf_87c01baa0aecd1de struct {
	__obf_3b7f6abbae19451e reflect2.Type
}

func (__obf_c9719e68325f7d44 *__obf_87c01baa0aecd1de) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_835102d74dbf01c6 := __obf_47edab4c16a2d88d.WhatIsNext()
	if __obf_835102d74dbf01c6 != ObjectValue && __obf_835102d74dbf01c6 != NilValue {
		__obf_47edab4c16a2d88d.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_47edab4c16a2d88d.
		Skip()
}

type __obf_9d5b95b9ade9b61b struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_904d17a1db8bff22 int64
	__obf_5e70cce5680d61b6 *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_9d5b95b9ade9b61b) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		if __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() == __obf_c9719e68325f7d44.__obf_904d17a1db8bff22 {
			__obf_c9719e68325f7d44.__obf_5e70cce5680d61b6.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		} else {
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_62cdf53760ce8958 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_62cdf53760ce8958) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_6f5b316f8a3ea115 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_6f5b316f8a3ea115) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_42adb9ad608cad61 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_42adb9ad608cad61) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_e4692cbb7ed95cb5 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_e4692cbb7ed95cb5) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_9deaed8c9e04f68f struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
	__obf_f37130b9a811c838 int64
	__obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_9deaed8c9e04f68f) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f37130b9a811c838:
			__obf_c9719e68325f7d44.__obf_f4a053b4639e9514.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_b4e105b6dad7055e struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
	__obf_f37130b9a811c838 int64
	__obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
	__obf_dd3d3f9213da10bd int64
	__obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_b4e105b6dad7055e) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f37130b9a811c838:
			__obf_c9719e68325f7d44.__obf_f4a053b4639e9514.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_dd3d3f9213da10bd:
			__obf_c9719e68325f7d44.__obf_eb37d3b9c8a6969f.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_308e238a1d242ed4 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
	__obf_f37130b9a811c838 int64
	__obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
	__obf_dd3d3f9213da10bd int64
	__obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
	__obf_ff6eeeffb66df370 int64
	__obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_308e238a1d242ed4) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f37130b9a811c838:
			__obf_c9719e68325f7d44.__obf_f4a053b4639e9514.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_dd3d3f9213da10bd:
			__obf_c9719e68325f7d44.__obf_eb37d3b9c8a6969f.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_ff6eeeffb66df370:
			__obf_c9719e68325f7d44.__obf_67871318f1fb2d29.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_185ac0b6849f5150 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
	__obf_f37130b9a811c838 int64
	__obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
	__obf_dd3d3f9213da10bd int64
	__obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
	__obf_ff6eeeffb66df370 int64
	__obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
	__obf_ed091618163fc9d2 int64
	__obf_de8ef2e6dc722a8a *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_185ac0b6849f5150) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f37130b9a811c838:
			__obf_c9719e68325f7d44.__obf_f4a053b4639e9514.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_dd3d3f9213da10bd:
			__obf_c9719e68325f7d44.__obf_eb37d3b9c8a6969f.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_ff6eeeffb66df370:
			__obf_c9719e68325f7d44.__obf_67871318f1fb2d29.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_ed091618163fc9d2:
			__obf_c9719e68325f7d44.__obf_de8ef2e6dc722a8a.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_b23dd4c081ae0650 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_218d368357bba003 int64
	__obf_c068c4315afd5733 *__obf_5b7a5815a547f104
	__obf_fca16fa0ada58d69 int64
	__obf_079fb380254cbbda *__obf_5b7a5815a547f104
	__obf_c06c05566acf5436 int64
	__obf_388ade5cdd5ca15d *__obf_5b7a5815a547f104
	__obf_199ea3123c647e59 int64
	__obf_52a24b5a0d6105f7 *__obf_5b7a5815a547f104
	__obf_d69abc8aa62025ef int64
	__obf_b7955ba8aad92db0 *__obf_5b7a5815a547f104
	__obf_f37130b9a811c838 int64
	__obf_f4a053b4639e9514 *__obf_5b7a5815a547f104
	__obf_dd3d3f9213da10bd int64
	__obf_eb37d3b9c8a6969f *__obf_5b7a5815a547f104
	__obf_ff6eeeffb66df370 int64
	__obf_67871318f1fb2d29 *__obf_5b7a5815a547f104
	__obf_ed091618163fc9d2 int64
	__obf_de8ef2e6dc722a8a *__obf_5b7a5815a547f104
	__obf_f89cd799d0d9d731 int64
	__obf_94405688d8ad81bc *__obf_5b7a5815a547f104
}

func (__obf_c9719e68325f7d44 *__obf_b23dd4c081ae0650) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.__obf_000080914c527b80() {
		return
	}
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		switch __obf_47edab4c16a2d88d.__obf_0ede172a7a3a6a62() {
		case __obf_c9719e68325f7d44.__obf_218d368357bba003:
			__obf_c9719e68325f7d44.__obf_c068c4315afd5733.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_fca16fa0ada58d69:
			__obf_c9719e68325f7d44.__obf_079fb380254cbbda.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_c06c05566acf5436:
			__obf_c9719e68325f7d44.__obf_388ade5cdd5ca15d.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_199ea3123c647e59:
			__obf_c9719e68325f7d44.__obf_52a24b5a0d6105f7.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_d69abc8aa62025ef:
			__obf_c9719e68325f7d44.__obf_b7955ba8aad92db0.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f37130b9a811c838:
			__obf_c9719e68325f7d44.__obf_f4a053b4639e9514.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_dd3d3f9213da10bd:
			__obf_c9719e68325f7d44.__obf_eb37d3b9c8a6969f.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_ff6eeeffb66df370:
			__obf_c9719e68325f7d44.__obf_67871318f1fb2d29.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_ed091618163fc9d2:
			__obf_c9719e68325f7d44.__obf_de8ef2e6dc722a8a.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		case __obf_c9719e68325f7d44.__obf_f89cd799d0d9d731:
			__obf_c9719e68325f7d44.__obf_94405688d8ad81bc.
				Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		default:
			__obf_47edab4c16a2d88d.
				Skip()
		}
		if __obf_47edab4c16a2d88d.__obf_ceba151f5486ca04() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF && len(__obf_c9719e68325f7d44.__obf_3b7f6abbae19451e.Type1().Name()) != 0 {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v.%s", __obf_c9719e68325f7d44.__obf_3b7f6abbae19451e, __obf_47edab4c16a2d88d.Error.Error())
	}
	__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
}

type __obf_5b7a5815a547f104 struct {
	__obf_13f6440419533990 reflect2.StructField
	__obf_5e70cce5680d61b6 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_5b7a5815a547f104) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_94be1fdf2ec0d598 := __obf_c9719e68325f7d44.__obf_13f6440419533990.UnsafeGet(__obf_47fa53f3d161f45c)
	__obf_c9719e68325f7d44.__obf_5e70cce5680d61b6.
		Decode(__obf_94be1fdf2ec0d598, __obf_47edab4c16a2d88d)
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%s: %s", __obf_c9719e68325f7d44.__obf_13f6440419533990.Name(), __obf_47edab4c16a2d88d.Error.Error())
	}
}

type __obf_bfcd3685f73f8364 struct {
	__obf_0db873ef26515049 ValDecoder
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
}

func (__obf_c9719e68325f7d44 *__obf_bfcd3685f73f8364) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_c9719e68325f7d44.__obf_0db873ef26515049.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	__obf_44b48c26051f8033 := *((*string)(__obf_47fa53f3d161f45c))
	__obf_30e4c83763be2016 := __obf_c9719e68325f7d44.__obf_25bd4f754a37b862.BorrowIterator([]byte(__obf_44b48c26051f8033))
	defer __obf_c9719e68325f7d44.__obf_25bd4f754a37b862.ReturnIterator(__obf_30e4c83763be2016)
	*((*string)(__obf_47fa53f3d161f45c)) = __obf_30e4c83763be2016.ReadString()
}

type __obf_e2eefa9baff51ea8 struct {
	__obf_0db873ef26515049 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_e2eefa9baff51ea8) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.WhatIsNext() == NilValue {
		__obf_c9719e68325f7d44.__obf_0db873ef26515049.
			Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
		return
	}
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != '"' {
		__obf_47edab4c16a2d88d.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
	__obf_c9719e68325f7d44.__obf_0db873ef26515049.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	if __obf_47edab4c16a2d88d.Error != nil {
		return
	}
	__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
	if __obf_bd08f5161fff294a != '"' {
		__obf_47edab4c16a2d88d.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
}
