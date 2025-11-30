package __obf_030d39f7a8de96c6

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_0891d5da085d6093(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_2a0baf04ebd1bcd4 := map[string]*Binding{}
	__obf_ef1b847f404bee13 := __obf_1219e3aba7e47ed0(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
		for _, __obf_f29eafe1a2494be5 := range __obf_04ea35bce3978e8f.FromNames {
			__obf_09fa488c9c17c6a9 := __obf_2a0baf04ebd1bcd4[__obf_f29eafe1a2494be5]
			if __obf_09fa488c9c17c6a9 == nil {
				__obf_2a0baf04ebd1bcd4[__obf_f29eafe1a2494be5] = __obf_04ea35bce3978e8f
				continue
			}
			__obf_4229457075af2115, __obf_78181e83750ef80b := __obf_accc8def568f71d8(__obf_a565fbaffca944f0.__obf_81639538813814ff, __obf_09fa488c9c17c6a9, __obf_04ea35bce3978e8f)
			if __obf_4229457075af2115 {
				delete(__obf_2a0baf04ebd1bcd4, __obf_f29eafe1a2494be5)
			}
			if !__obf_78181e83750ef80b {
				__obf_2a0baf04ebd1bcd4[__obf_f29eafe1a2494be5] = __obf_04ea35bce3978e8f
			}
		}
	}
	__obf_dadaf38c30ea4eba := map[string]*__obf_7198db8bbb81224a{}
	for __obf_4c6b7384a71b9c24, __obf_04ea35bce3978e8f := range __obf_2a0baf04ebd1bcd4 {
		__obf_dadaf38c30ea4eba[__obf_4c6b7384a71b9c24] = __obf_04ea35bce3978e8f.Decoder.(*__obf_7198db8bbb81224a)
	}

	if !__obf_a565fbaffca944f0.__obf_af13e3babb780a4e() {
		for __obf_4c6b7384a71b9c24, __obf_04ea35bce3978e8f := range __obf_2a0baf04ebd1bcd4 {
			if _, __obf_ccdc4ee8b9d2f52f := __obf_dadaf38c30ea4eba[strings.ToLower(__obf_4c6b7384a71b9c24)]; !__obf_ccdc4ee8b9d2f52f {
				__obf_dadaf38c30ea4eba[strings.ToLower(__obf_4c6b7384a71b9c24)] = __obf_04ea35bce3978e8f.Decoder.(*__obf_7198db8bbb81224a)
			}
		}
	}

	return __obf_ebe428bb2f728a3a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba)
}

func __obf_ebe428bb2f728a3a(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type, __obf_dadaf38c30ea4eba map[string]*__obf_7198db8bbb81224a) ValDecoder {
	if __obf_a565fbaffca944f0.__obf_5a45cc14ef9b494e {
		return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1: __obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba: __obf_dadaf38c30ea4eba, __obf_5a45cc14ef9b494e: true}
	}
	__obf_e75759e98c5cb01a := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_dadaf38c30ea4eba) {
	case 0:
		return &__obf_16a72fa16e437dbb{__obf_8744d0b8c80ccdc1}
	case 1:
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			return &__obf_371eacac4c292421{__obf_8744d0b8c80ccdc1, __obf_9afc1c5f99993a5e, __obf_e23d8181e71543e6}
		}
	case 2:
		var __obf_c51999a2449e856a int64
		var __obf_569285913319bc7a int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_c51999a2449e856a == 0 {
				__obf_c51999a2449e856a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else {
				__obf_569285913319bc7a = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			}
		}
		return &__obf_a9dfceaad87ed0b3{__obf_8744d0b8c80ccdc1, __obf_c51999a2449e856a, __obf_0f8d27fa45cd3761, __obf_569285913319bc7a, __obf_f30b524436d8ecd8}
	case 3:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			}
		}
		return &__obf_87aa9063c6fa180d{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1}
	case 4:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			}
		}
		return &__obf_8ef4b86f438f1bd4{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea}
	case 5:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			}
		}
		return &__obf_e3c8954bfa7acf68{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed}
	case 6:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_12b48eee32fe7a43 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		var __obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else if __obf_ecae97d74d67e368 == 0 {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			} else {
				__obf_12b48eee32fe7a43 = __obf_9afc1c5f99993a5e
				__obf_0b90998076c47ea5 = __obf_e23d8181e71543e6
			}
		}
		return &__obf_1a2ff03964605af2{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed, __obf_12b48eee32fe7a43, __obf_0b90998076c47ea5}
	case 7:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_12b48eee32fe7a43 int64
		var __obf_413422225123a9a3 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		var __obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
		var __obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else if __obf_ecae97d74d67e368 == 0 {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			} else if __obf_12b48eee32fe7a43 == 0 {
				__obf_12b48eee32fe7a43 = __obf_9afc1c5f99993a5e
				__obf_0b90998076c47ea5 = __obf_e23d8181e71543e6
			} else {
				__obf_413422225123a9a3 = __obf_9afc1c5f99993a5e
				__obf_c6fb15da1dfb26a6 = __obf_e23d8181e71543e6
			}
		}
		return &__obf_a8c5dfe6b7cfb9b1{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed, __obf_12b48eee32fe7a43, __obf_0b90998076c47ea5, __obf_413422225123a9a3, __obf_c6fb15da1dfb26a6}
	case 8:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_12b48eee32fe7a43 int64
		var __obf_413422225123a9a3 int64
		var __obf_4165a8efde98c250 int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		var __obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
		var __obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
		var __obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else if __obf_ecae97d74d67e368 == 0 {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			} else if __obf_12b48eee32fe7a43 == 0 {
				__obf_12b48eee32fe7a43 = __obf_9afc1c5f99993a5e
				__obf_0b90998076c47ea5 = __obf_e23d8181e71543e6
			} else if __obf_413422225123a9a3 == 0 {
				__obf_413422225123a9a3 = __obf_9afc1c5f99993a5e
				__obf_c6fb15da1dfb26a6 = __obf_e23d8181e71543e6
			} else {
				__obf_4165a8efde98c250 = __obf_9afc1c5f99993a5e
				__obf_be1dc2f72b01064a = __obf_e23d8181e71543e6
			}
		}
		return &__obf_6c52c15bc30b18dc{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed, __obf_12b48eee32fe7a43, __obf_0b90998076c47ea5, __obf_413422225123a9a3, __obf_c6fb15da1dfb26a6, __obf_4165a8efde98c250, __obf_be1dc2f72b01064a}
	case 9:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_12b48eee32fe7a43 int64
		var __obf_413422225123a9a3 int64
		var __obf_4165a8efde98c250 int64
		var __obf_e59a52269b341a5e int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		var __obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
		var __obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
		var __obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
		var __obf_31ada8a00f4ff6fc *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else if __obf_ecae97d74d67e368 == 0 {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			} else if __obf_12b48eee32fe7a43 == 0 {
				__obf_12b48eee32fe7a43 = __obf_9afc1c5f99993a5e
				__obf_0b90998076c47ea5 = __obf_e23d8181e71543e6
			} else if __obf_413422225123a9a3 == 0 {
				__obf_413422225123a9a3 = __obf_9afc1c5f99993a5e
				__obf_c6fb15da1dfb26a6 = __obf_e23d8181e71543e6
			} else if __obf_4165a8efde98c250 == 0 {
				__obf_4165a8efde98c250 = __obf_9afc1c5f99993a5e
				__obf_be1dc2f72b01064a = __obf_e23d8181e71543e6
			} else {
				__obf_e59a52269b341a5e = __obf_9afc1c5f99993a5e
				__obf_31ada8a00f4ff6fc = __obf_e23d8181e71543e6
			}
		}
		return &__obf_a361d15f761b2265{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed, __obf_12b48eee32fe7a43, __obf_0b90998076c47ea5, __obf_413422225123a9a3, __obf_c6fb15da1dfb26a6, __obf_4165a8efde98c250, __obf_be1dc2f72b01064a, __obf_e59a52269b341a5e, __obf_31ada8a00f4ff6fc}
	case 10:
		var __obf_405fafbf8107913a int64
		var __obf_6fe4bafb1f235814 int64
		var __obf_587df1f71911f543 int64
		var __obf_91c5db2b5d4a5046 int64
		var __obf_ecae97d74d67e368 int64
		var __obf_12b48eee32fe7a43 int64
		var __obf_413422225123a9a3 int64
		var __obf_4165a8efde98c250 int64
		var __obf_e59a52269b341a5e int64
		var __obf_065528dc30a0c2ed int64
		var __obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
		var __obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
		var __obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
		var __obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
		var __obf_146f68e06750b6ed *__obf_7198db8bbb81224a
		var __obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
		var __obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
		var __obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
		var __obf_31ada8a00f4ff6fc *__obf_7198db8bbb81224a
		var __obf_04520d0a7ef08ed5 *__obf_7198db8bbb81224a
		for __obf_07e0e32c6be0db36, __obf_e23d8181e71543e6 := range __obf_dadaf38c30ea4eba {
			__obf_9afc1c5f99993a5e := __obf_47e34696e8c7ae78(__obf_07e0e32c6be0db36, __obf_a565fbaffca944f0.__obf_af13e3babb780a4e())
			_, __obf_e9ec2ece3f569adf := __obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e]
			if __obf_e9ec2ece3f569adf {
				return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
			}
			__obf_e75759e98c5cb01a[__obf_9afc1c5f99993a5e] = struct{}{}
			if __obf_405fafbf8107913a == 0 {
				__obf_405fafbf8107913a = __obf_9afc1c5f99993a5e
				__obf_0f8d27fa45cd3761 = __obf_e23d8181e71543e6
			} else if __obf_6fe4bafb1f235814 == 0 {
				__obf_6fe4bafb1f235814 = __obf_9afc1c5f99993a5e
				__obf_f30b524436d8ecd8 = __obf_e23d8181e71543e6
			} else if __obf_587df1f71911f543 == 0 {
				__obf_587df1f71911f543 = __obf_9afc1c5f99993a5e
				__obf_219114237ad7b5f1 = __obf_e23d8181e71543e6
			} else if __obf_91c5db2b5d4a5046 == 0 {
				__obf_91c5db2b5d4a5046 = __obf_9afc1c5f99993a5e
				__obf_7daf6f0c8b9a78ea = __obf_e23d8181e71543e6
			} else if __obf_ecae97d74d67e368 == 0 {
				__obf_ecae97d74d67e368 = __obf_9afc1c5f99993a5e
				__obf_146f68e06750b6ed = __obf_e23d8181e71543e6
			} else if __obf_12b48eee32fe7a43 == 0 {
				__obf_12b48eee32fe7a43 = __obf_9afc1c5f99993a5e
				__obf_0b90998076c47ea5 = __obf_e23d8181e71543e6
			} else if __obf_413422225123a9a3 == 0 {
				__obf_413422225123a9a3 = __obf_9afc1c5f99993a5e
				__obf_c6fb15da1dfb26a6 = __obf_e23d8181e71543e6
			} else if __obf_4165a8efde98c250 == 0 {
				__obf_4165a8efde98c250 = __obf_9afc1c5f99993a5e
				__obf_be1dc2f72b01064a = __obf_e23d8181e71543e6
			} else if __obf_e59a52269b341a5e == 0 {
				__obf_e59a52269b341a5e = __obf_9afc1c5f99993a5e
				__obf_31ada8a00f4ff6fc = __obf_e23d8181e71543e6
			} else {
				__obf_065528dc30a0c2ed = __obf_9afc1c5f99993a5e
				__obf_04520d0a7ef08ed5 = __obf_e23d8181e71543e6
			}
		}
		return &__obf_c2713b8bf9032980{__obf_8744d0b8c80ccdc1, __obf_405fafbf8107913a, __obf_0f8d27fa45cd3761, __obf_6fe4bafb1f235814, __obf_f30b524436d8ecd8, __obf_587df1f71911f543, __obf_219114237ad7b5f1, __obf_91c5db2b5d4a5046, __obf_7daf6f0c8b9a78ea, __obf_ecae97d74d67e368, __obf_146f68e06750b6ed, __obf_12b48eee32fe7a43, __obf_0b90998076c47ea5, __obf_413422225123a9a3, __obf_c6fb15da1dfb26a6, __obf_4165a8efde98c250, __obf_be1dc2f72b01064a, __obf_e59a52269b341a5e, __obf_31ada8a00f4ff6fc, __obf_065528dc30a0c2ed, __obf_04520d0a7ef08ed5}
	}
	return &__obf_8131117e7356a2b1{__obf_8744d0b8c80ccdc1, __obf_dadaf38c30ea4eba, false}
}

type __obf_8131117e7356a2b1 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_dadaf38c30ea4eba map[string]*__obf_7198db8bbb81224a
	__obf_5a45cc14ef9b494e bool
}

func (__obf_11a3f28116164d09 *__obf_8131117e7356a2b1) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	var __obf_24309b3b0ff9ba22 byte
	for __obf_24309b3b0ff9ba22 = ','; __obf_24309b3b0ff9ba22 == ','; __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() {
		__obf_11a3f28116164d09.__obf_81da3c07e4697e4b(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	if __obf_24309b3b0ff9ba22 != '}' {
		__obf_4ab56a99965952e7.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

func (__obf_11a3f28116164d09 *__obf_8131117e7356a2b1) __obf_81da3c07e4697e4b(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	var __obf_cd4d02f264b18d55 string
	var __obf_e23d8181e71543e6 *__obf_7198db8bbb81224a
	if __obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_8c3b82d5a5e45b5c {
		__obf_18124272fd265e43 := __obf_4ab56a99965952e7.ReadStringAsSlice()
		__obf_cd4d02f264b18d55 = *(*string)(unsafe.Pointer(&__obf_18124272fd265e43))
		__obf_e23d8181e71543e6 = __obf_11a3f28116164d09.__obf_dadaf38c30ea4eba[__obf_cd4d02f264b18d55]
		if __obf_e23d8181e71543e6 == nil && !__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_af13e3babb780a4e {
			__obf_e23d8181e71543e6 = __obf_11a3f28116164d09.__obf_dadaf38c30ea4eba[strings.ToLower(__obf_cd4d02f264b18d55)]
		}
	} else {
		__obf_cd4d02f264b18d55 = __obf_4ab56a99965952e7.ReadString()
		__obf_e23d8181e71543e6 = __obf_11a3f28116164d09.__obf_dadaf38c30ea4eba[__obf_cd4d02f264b18d55]
		if __obf_e23d8181e71543e6 == nil && !__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_af13e3babb780a4e {
			__obf_e23d8181e71543e6 = __obf_11a3f28116164d09.__obf_dadaf38c30ea4eba[strings.ToLower(__obf_cd4d02f264b18d55)]
		}
	}
	if __obf_e23d8181e71543e6 == nil {
		if __obf_11a3f28116164d09.__obf_5a45cc14ef9b494e {
			__obf_489aa20db03fe286 := "found unknown field: " + __obf_cd4d02f264b18d55
			__obf_4ab56a99965952e7.
				ReportError("ReadObject", __obf_489aa20db03fe286)
		}
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 != ':' {
			__obf_4ab56a99965952e7.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		}
		__obf_4ab56a99965952e7.
			Skip()
		return
	}
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != ':' {
		__obf_4ab56a99965952e7.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
	}
	__obf_e23d8181e71543e6.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
}

type __obf_16a72fa16e437dbb struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
}

func (__obf_11a3f28116164d09 *__obf_16a72fa16e437dbb) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_8afc98a56c05c14e := __obf_4ab56a99965952e7.WhatIsNext()
	if __obf_8afc98a56c05c14e != ObjectValue && __obf_8afc98a56c05c14e != NilValue {
		__obf_4ab56a99965952e7.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_4ab56a99965952e7.
		Skip()
}

type __obf_371eacac4c292421 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_9afc1c5f99993a5e int64
	__obf_e23d8181e71543e6 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_371eacac4c292421) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		if __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() == __obf_11a3f28116164d09.__obf_9afc1c5f99993a5e {
			__obf_11a3f28116164d09.__obf_e23d8181e71543e6.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		} else {
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_a9dfceaad87ed0b3 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_a9dfceaad87ed0b3) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_87aa9063c6fa180d struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_87aa9063c6fa180d) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_8ef4b86f438f1bd4 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_8ef4b86f438f1bd4) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_e3c8954bfa7acf68 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_e3c8954bfa7acf68) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_1a2ff03964605af2 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
	__obf_76a53c275a847bef int64
	__obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_1a2ff03964605af2) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_76a53c275a847bef:
			__obf_11a3f28116164d09.__obf_0b90998076c47ea5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_a8c5dfe6b7cfb9b1 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
	__obf_76a53c275a847bef int64
	__obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
	__obf_468a0f4c60072fc1 int64
	__obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_a8c5dfe6b7cfb9b1) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_76a53c275a847bef:
			__obf_11a3f28116164d09.__obf_0b90998076c47ea5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_468a0f4c60072fc1:
			__obf_11a3f28116164d09.__obf_c6fb15da1dfb26a6.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_6c52c15bc30b18dc struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
	__obf_76a53c275a847bef int64
	__obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
	__obf_468a0f4c60072fc1 int64
	__obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
	__obf_92328418b7eb5f1c int64
	__obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_6c52c15bc30b18dc) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_76a53c275a847bef:
			__obf_11a3f28116164d09.__obf_0b90998076c47ea5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_468a0f4c60072fc1:
			__obf_11a3f28116164d09.__obf_c6fb15da1dfb26a6.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_92328418b7eb5f1c:
			__obf_11a3f28116164d09.__obf_be1dc2f72b01064a.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_a361d15f761b2265 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
	__obf_76a53c275a847bef int64
	__obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
	__obf_468a0f4c60072fc1 int64
	__obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
	__obf_92328418b7eb5f1c int64
	__obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
	__obf_a4d3f7d07fc38cea int64
	__obf_31ada8a00f4ff6fc *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_a361d15f761b2265) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_76a53c275a847bef:
			__obf_11a3f28116164d09.__obf_0b90998076c47ea5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_468a0f4c60072fc1:
			__obf_11a3f28116164d09.__obf_c6fb15da1dfb26a6.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_92328418b7eb5f1c:
			__obf_11a3f28116164d09.__obf_be1dc2f72b01064a.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_a4d3f7d07fc38cea:
			__obf_11a3f28116164d09.__obf_31ada8a00f4ff6fc.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_c2713b8bf9032980 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_c51999a2449e856a int64
	__obf_0f8d27fa45cd3761 *__obf_7198db8bbb81224a
	__obf_569285913319bc7a int64
	__obf_f30b524436d8ecd8 *__obf_7198db8bbb81224a
	__obf_d5a0611ca263ef13 int64
	__obf_219114237ad7b5f1 *__obf_7198db8bbb81224a
	__obf_456d47a3bec00c7d int64
	__obf_7daf6f0c8b9a78ea *__obf_7198db8bbb81224a
	__obf_05b07dbadf878ee7 int64
	__obf_146f68e06750b6ed *__obf_7198db8bbb81224a
	__obf_76a53c275a847bef int64
	__obf_0b90998076c47ea5 *__obf_7198db8bbb81224a
	__obf_468a0f4c60072fc1 int64
	__obf_c6fb15da1dfb26a6 *__obf_7198db8bbb81224a
	__obf_92328418b7eb5f1c int64
	__obf_be1dc2f72b01064a *__obf_7198db8bbb81224a
	__obf_a4d3f7d07fc38cea int64
	__obf_31ada8a00f4ff6fc *__obf_7198db8bbb81224a
	__obf_8bce8a1dee6ddf4d int64
	__obf_04520d0a7ef08ed5 *__obf_7198db8bbb81224a
}

func (__obf_11a3f28116164d09 *__obf_c2713b8bf9032980) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.__obf_2c0c8655dcf5f087() {
		return
	}
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		switch __obf_4ab56a99965952e7.__obf_98be9914bcefeaa7() {
		case __obf_11a3f28116164d09.__obf_c51999a2449e856a:
			__obf_11a3f28116164d09.__obf_0f8d27fa45cd3761.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_569285913319bc7a:
			__obf_11a3f28116164d09.__obf_f30b524436d8ecd8.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_d5a0611ca263ef13:
			__obf_11a3f28116164d09.__obf_219114237ad7b5f1.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_456d47a3bec00c7d:
			__obf_11a3f28116164d09.__obf_7daf6f0c8b9a78ea.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_05b07dbadf878ee7:
			__obf_11a3f28116164d09.__obf_146f68e06750b6ed.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_76a53c275a847bef:
			__obf_11a3f28116164d09.__obf_0b90998076c47ea5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_468a0f4c60072fc1:
			__obf_11a3f28116164d09.__obf_c6fb15da1dfb26a6.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_92328418b7eb5f1c:
			__obf_11a3f28116164d09.__obf_be1dc2f72b01064a.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_a4d3f7d07fc38cea:
			__obf_11a3f28116164d09.__obf_31ada8a00f4ff6fc.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		case __obf_11a3f28116164d09.__obf_8bce8a1dee6ddf4d:
			__obf_11a3f28116164d09.__obf_04520d0a7ef08ed5.
				Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		default:
			__obf_4ab56a99965952e7.
				Skip()
		}
		if __obf_4ab56a99965952e7.__obf_6a883651c29f40bf() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF && len(__obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1.Type1().Name()) != 0 {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v.%s", __obf_11a3f28116164d09.__obf_8744d0b8c80ccdc1, __obf_4ab56a99965952e7.Error.Error())
	}
	__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
}

type __obf_7198db8bbb81224a struct {
	__obf_cd4d02f264b18d55 reflect2.StructField
	__obf_e23d8181e71543e6 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_7198db8bbb81224a) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_ba13c65f6e3a796c := __obf_11a3f28116164d09.__obf_cd4d02f264b18d55.UnsafeGet(__obf_dbbf371b91136494)
	__obf_11a3f28116164d09.__obf_e23d8181e71543e6.
		Decode(__obf_ba13c65f6e3a796c, __obf_4ab56a99965952e7)
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%s: %s", __obf_11a3f28116164d09.__obf_cd4d02f264b18d55.Name(), __obf_4ab56a99965952e7.Error.Error())
	}
}

type __obf_73fb4b35802e3bb0 struct {
	__obf_be23981cae4f81b9 ValDecoder
	__obf_679611dc56ff465b *__obf_81639538813814ff
}

func (__obf_11a3f28116164d09 *__obf_73fb4b35802e3bb0) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	__obf_428c3b4a95725c4a := *((*string)(__obf_dbbf371b91136494))
	__obf_bb43376cdb87a0aa := __obf_11a3f28116164d09.__obf_679611dc56ff465b.BorrowIterator([]byte(__obf_428c3b4a95725c4a))
	defer __obf_11a3f28116164d09.__obf_679611dc56ff465b.ReturnIterator(__obf_bb43376cdb87a0aa)
	*((*string)(__obf_dbbf371b91136494)) = __obf_bb43376cdb87a0aa.ReadString()
}

type __obf_6fee448b50200908 struct {
	__obf_be23981cae4f81b9 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_6fee448b50200908) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.WhatIsNext() == NilValue {
		__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
			Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
		return
	}
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != '"' {
		__obf_4ab56a99965952e7.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	if __obf_4ab56a99965952e7.Error != nil {
		return
	}
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
	if __obf_24309b3b0ff9ba22 != '"' {
		__obf_4ab56a99965952e7.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
}
