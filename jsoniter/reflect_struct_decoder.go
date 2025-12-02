package __obf_c7b79b12b33d8238

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_00d9b9fc4fc9dd14(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_a61cdc4b92dc1d63 := map[string]*Binding{}
	__obf_31531899b0d36123 := __obf_0572835446c88a68(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
		for _, __obf_377a35e286c5e577 := range __obf_56b28257c2a1cc8d.FromNames {
			__obf_ffbea95e31df6856 := __obf_a61cdc4b92dc1d63[__obf_377a35e286c5e577]
			if __obf_ffbea95e31df6856 == nil {
				__obf_a61cdc4b92dc1d63[__obf_377a35e286c5e577] = __obf_56b28257c2a1cc8d
				continue
			}
			__obf_81ed0382471e779e, __obf_5019fe5e38747cb9 := __obf_169b1277fbacdc3f(__obf_99ec45908bceabdb.__obf_30fe5c95cabd69c0, __obf_ffbea95e31df6856, __obf_56b28257c2a1cc8d)
			if __obf_81ed0382471e779e {
				delete(__obf_a61cdc4b92dc1d63, __obf_377a35e286c5e577)
			}
			if !__obf_5019fe5e38747cb9 {
				__obf_a61cdc4b92dc1d63[__obf_377a35e286c5e577] = __obf_56b28257c2a1cc8d
			}
		}
	}
	__obf_c397376f03d6d1bb := map[string]*__obf_7ca3f82b6bdb44f7{}
	for __obf_60f36f72f8a4c230, __obf_56b28257c2a1cc8d := range __obf_a61cdc4b92dc1d63 {
		__obf_c397376f03d6d1bb[__obf_60f36f72f8a4c230] = __obf_56b28257c2a1cc8d.Decoder.(*__obf_7ca3f82b6bdb44f7)
	}

	if !__obf_99ec45908bceabdb.__obf_3704f04b7ac67609() {
		for __obf_60f36f72f8a4c230, __obf_56b28257c2a1cc8d := range __obf_a61cdc4b92dc1d63 {
			if _, __obf_22405d874667f998 := __obf_c397376f03d6d1bb[strings.ToLower(__obf_60f36f72f8a4c230)]; !__obf_22405d874667f998 {
				__obf_c397376f03d6d1bb[strings.ToLower(__obf_60f36f72f8a4c230)] = __obf_56b28257c2a1cc8d.Decoder.(*__obf_7ca3f82b6bdb44f7)
			}
		}
	}

	return __obf_2c41cf2b3da00b1b(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb)
}

func __obf_2c41cf2b3da00b1b(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type, __obf_c397376f03d6d1bb map[string]*__obf_7ca3f82b6bdb44f7) ValDecoder {
	if __obf_99ec45908bceabdb.__obf_b1dc099c042d4d26 {
		return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c: __obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb: __obf_c397376f03d6d1bb, __obf_b1dc099c042d4d26: true}
	}
	__obf_7585cd2caeaa548b := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_c397376f03d6d1bb) {
	case 0:
		return &__obf_966d8a7edd331c70{__obf_edcded11af6ebc4c}
	case 1:
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			return &__obf_13a378ace127c1a7{__obf_edcded11af6ebc4c, __obf_f9abf27a7d87d6ae, __obf_fbbd552d8d48f5f2}
		}
	case 2:
		var __obf_08b36383be79e662 int64
		var __obf_d1ce14cb84761def int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_08b36383be79e662 == 0 {
				__obf_08b36383be79e662 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else {
				__obf_d1ce14cb84761def = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_37dd7161a98d005f{__obf_edcded11af6ebc4c, __obf_08b36383be79e662, __obf_07c5ed10b4f8401d, __obf_d1ce14cb84761def, __obf_39d063eb592af548}
	case 3:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_3cdbf5fc3dccb593{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925}
	case 4:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_bf2641561b45c065{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895}
	case 5:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_3803a5a30fcafd24{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe}
	case 6:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_7564cf98b8cb5615 int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		var __obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else if __obf_eead0593ff2b927f == 0 {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			} else {
				__obf_7564cf98b8cb5615 = __obf_f9abf27a7d87d6ae
				__obf_b938864c5f5e8dbd = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_046e872add4a19f4{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe, __obf_7564cf98b8cb5615, __obf_b938864c5f5e8dbd}
	case 7:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_7564cf98b8cb5615 int64
		var __obf_edcc1711692f1cbc int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		var __obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
		var __obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else if __obf_eead0593ff2b927f == 0 {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			} else if __obf_7564cf98b8cb5615 == 0 {
				__obf_7564cf98b8cb5615 = __obf_f9abf27a7d87d6ae
				__obf_b938864c5f5e8dbd = __obf_fbbd552d8d48f5f2
			} else {
				__obf_edcc1711692f1cbc = __obf_f9abf27a7d87d6ae
				__obf_95e559f2e7b06d72 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_be3b01ef821c1fe6{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe, __obf_7564cf98b8cb5615, __obf_b938864c5f5e8dbd, __obf_edcc1711692f1cbc, __obf_95e559f2e7b06d72}
	case 8:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_7564cf98b8cb5615 int64
		var __obf_edcc1711692f1cbc int64
		var __obf_5a7dfb949ef88283 int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		var __obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
		var __obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
		var __obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else if __obf_eead0593ff2b927f == 0 {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			} else if __obf_7564cf98b8cb5615 == 0 {
				__obf_7564cf98b8cb5615 = __obf_f9abf27a7d87d6ae
				__obf_b938864c5f5e8dbd = __obf_fbbd552d8d48f5f2
			} else if __obf_edcc1711692f1cbc == 0 {
				__obf_edcc1711692f1cbc = __obf_f9abf27a7d87d6ae
				__obf_95e559f2e7b06d72 = __obf_fbbd552d8d48f5f2
			} else {
				__obf_5a7dfb949ef88283 = __obf_f9abf27a7d87d6ae
				__obf_e24e7b3ac2913a05 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_8239535c6e562d8f{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe, __obf_7564cf98b8cb5615, __obf_b938864c5f5e8dbd, __obf_edcc1711692f1cbc, __obf_95e559f2e7b06d72, __obf_5a7dfb949ef88283, __obf_e24e7b3ac2913a05}
	case 9:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_7564cf98b8cb5615 int64
		var __obf_edcc1711692f1cbc int64
		var __obf_5a7dfb949ef88283 int64
		var __obf_1737630a914e6577 int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		var __obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
		var __obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
		var __obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
		var __obf_6549f536b94c945d *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else if __obf_eead0593ff2b927f == 0 {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			} else if __obf_7564cf98b8cb5615 == 0 {
				__obf_7564cf98b8cb5615 = __obf_f9abf27a7d87d6ae
				__obf_b938864c5f5e8dbd = __obf_fbbd552d8d48f5f2
			} else if __obf_edcc1711692f1cbc == 0 {
				__obf_edcc1711692f1cbc = __obf_f9abf27a7d87d6ae
				__obf_95e559f2e7b06d72 = __obf_fbbd552d8d48f5f2
			} else if __obf_5a7dfb949ef88283 == 0 {
				__obf_5a7dfb949ef88283 = __obf_f9abf27a7d87d6ae
				__obf_e24e7b3ac2913a05 = __obf_fbbd552d8d48f5f2
			} else {
				__obf_1737630a914e6577 = __obf_f9abf27a7d87d6ae
				__obf_6549f536b94c945d = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_c1bdac930f0bb76f{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe, __obf_7564cf98b8cb5615, __obf_b938864c5f5e8dbd, __obf_edcc1711692f1cbc, __obf_95e559f2e7b06d72, __obf_5a7dfb949ef88283, __obf_e24e7b3ac2913a05, __obf_1737630a914e6577, __obf_6549f536b94c945d}
	case 10:
		var __obf_17a87ca452046558 int64
		var __obf_44abbae2c72963d6 int64
		var __obf_9f0db67392d20e48 int64
		var __obf_70e8d883a59fd21e int64
		var __obf_eead0593ff2b927f int64
		var __obf_7564cf98b8cb5615 int64
		var __obf_edcc1711692f1cbc int64
		var __obf_5a7dfb949ef88283 int64
		var __obf_1737630a914e6577 int64
		var __obf_277b0405c1c0e7cc int64
		var __obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
		var __obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
		var __obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
		var __obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
		var __obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
		var __obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
		var __obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
		var __obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
		var __obf_6549f536b94c945d *__obf_7ca3f82b6bdb44f7
		var __obf_b4ad13b78d729536 *__obf_7ca3f82b6bdb44f7
		for __obf_dd29aaaf63f43a4a, __obf_fbbd552d8d48f5f2 := range __obf_c397376f03d6d1bb {
			__obf_f9abf27a7d87d6ae := __obf_cedc170066626295(__obf_dd29aaaf63f43a4a, __obf_99ec45908bceabdb.__obf_3704f04b7ac67609())
			_, __obf_dabb21aea1980073 := __obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae]
			if __obf_dabb21aea1980073 {
				return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
			}
			__obf_7585cd2caeaa548b[__obf_f9abf27a7d87d6ae] = struct{}{}
			if __obf_17a87ca452046558 == 0 {
				__obf_17a87ca452046558 = __obf_f9abf27a7d87d6ae
				__obf_07c5ed10b4f8401d = __obf_fbbd552d8d48f5f2
			} else if __obf_44abbae2c72963d6 == 0 {
				__obf_44abbae2c72963d6 = __obf_f9abf27a7d87d6ae
				__obf_39d063eb592af548 = __obf_fbbd552d8d48f5f2
			} else if __obf_9f0db67392d20e48 == 0 {
				__obf_9f0db67392d20e48 = __obf_f9abf27a7d87d6ae
				__obf_ace5a68d76e4d925 = __obf_fbbd552d8d48f5f2
			} else if __obf_70e8d883a59fd21e == 0 {
				__obf_70e8d883a59fd21e = __obf_f9abf27a7d87d6ae
				__obf_a27ca7588c1f1895 = __obf_fbbd552d8d48f5f2
			} else if __obf_eead0593ff2b927f == 0 {
				__obf_eead0593ff2b927f = __obf_f9abf27a7d87d6ae
				__obf_0c438d080933cdbe = __obf_fbbd552d8d48f5f2
			} else if __obf_7564cf98b8cb5615 == 0 {
				__obf_7564cf98b8cb5615 = __obf_f9abf27a7d87d6ae
				__obf_b938864c5f5e8dbd = __obf_fbbd552d8d48f5f2
			} else if __obf_edcc1711692f1cbc == 0 {
				__obf_edcc1711692f1cbc = __obf_f9abf27a7d87d6ae
				__obf_95e559f2e7b06d72 = __obf_fbbd552d8d48f5f2
			} else if __obf_5a7dfb949ef88283 == 0 {
				__obf_5a7dfb949ef88283 = __obf_f9abf27a7d87d6ae
				__obf_e24e7b3ac2913a05 = __obf_fbbd552d8d48f5f2
			} else if __obf_1737630a914e6577 == 0 {
				__obf_1737630a914e6577 = __obf_f9abf27a7d87d6ae
				__obf_6549f536b94c945d = __obf_fbbd552d8d48f5f2
			} else {
				__obf_277b0405c1c0e7cc = __obf_f9abf27a7d87d6ae
				__obf_b4ad13b78d729536 = __obf_fbbd552d8d48f5f2
			}
		}
		return &__obf_554a4f65c9d0ea9b{__obf_edcded11af6ebc4c, __obf_17a87ca452046558, __obf_07c5ed10b4f8401d, __obf_44abbae2c72963d6, __obf_39d063eb592af548, __obf_9f0db67392d20e48, __obf_ace5a68d76e4d925, __obf_70e8d883a59fd21e, __obf_a27ca7588c1f1895, __obf_eead0593ff2b927f, __obf_0c438d080933cdbe, __obf_7564cf98b8cb5615, __obf_b938864c5f5e8dbd, __obf_edcc1711692f1cbc, __obf_95e559f2e7b06d72, __obf_5a7dfb949ef88283, __obf_e24e7b3ac2913a05, __obf_1737630a914e6577, __obf_6549f536b94c945d, __obf_277b0405c1c0e7cc, __obf_b4ad13b78d729536}
	}
	return &__obf_82a8336e5e8f0eec{__obf_edcded11af6ebc4c, __obf_c397376f03d6d1bb, false}
}

type __obf_82a8336e5e8f0eec struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_c397376f03d6d1bb map[string]*__obf_7ca3f82b6bdb44f7
	__obf_b1dc099c042d4d26 bool
}

func (__obf_801f19702638809c *__obf_82a8336e5e8f0eec) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	var __obf_12577c03a12f0d2c byte
	for __obf_12577c03a12f0d2c = ','; __obf_12577c03a12f0d2c == ','; __obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d() {
		__obf_801f19702638809c.__obf_56696f90bb4bb9f1(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	if __obf_12577c03a12f0d2c != '}' {
		__obf_0da8c843dad13139.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

func (__obf_801f19702638809c *__obf_82a8336e5e8f0eec) __obf_56696f90bb4bb9f1(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	var __obf_ad34f8a6de2b01cb string
	var __obf_fbbd552d8d48f5f2 *__obf_7ca3f82b6bdb44f7
	if __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_82ca3cd044b095ae {
		__obf_bd092ea2efa33534 := __obf_0da8c843dad13139.ReadStringAsSlice()
		__obf_ad34f8a6de2b01cb = *(*string)(unsafe.Pointer(&__obf_bd092ea2efa33534))
		__obf_fbbd552d8d48f5f2 = __obf_801f19702638809c.__obf_c397376f03d6d1bb[__obf_ad34f8a6de2b01cb]
		if __obf_fbbd552d8d48f5f2 == nil && !__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_3704f04b7ac67609 {
			__obf_fbbd552d8d48f5f2 = __obf_801f19702638809c.__obf_c397376f03d6d1bb[strings.ToLower(__obf_ad34f8a6de2b01cb)]
		}
	} else {
		__obf_ad34f8a6de2b01cb = __obf_0da8c843dad13139.ReadString()
		__obf_fbbd552d8d48f5f2 = __obf_801f19702638809c.__obf_c397376f03d6d1bb[__obf_ad34f8a6de2b01cb]
		if __obf_fbbd552d8d48f5f2 == nil && !__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_3704f04b7ac67609 {
			__obf_fbbd552d8d48f5f2 = __obf_801f19702638809c.__obf_c397376f03d6d1bb[strings.ToLower(__obf_ad34f8a6de2b01cb)]
		}
	}
	if __obf_fbbd552d8d48f5f2 == nil {
		if __obf_801f19702638809c.__obf_b1dc099c042d4d26 {
			__obf_e2c48e9bb87a8f24 := "found unknown field: " + __obf_ad34f8a6de2b01cb
			__obf_0da8c843dad13139.
				ReportError("ReadObject", __obf_e2c48e9bb87a8f24)
		}
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c != ':' {
			__obf_0da8c843dad13139.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
		}
		__obf_0da8c843dad13139.
			Skip()
		return
	}
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != ':' {
		__obf_0da8c843dad13139.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
	}
	__obf_fbbd552d8d48f5f2.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
}

type __obf_966d8a7edd331c70 struct {
	__obf_edcded11af6ebc4c reflect2.Type
}

func (__obf_801f19702638809c *__obf_966d8a7edd331c70) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_2cd064ad0a3f5c35 := __obf_0da8c843dad13139.WhatIsNext()
	if __obf_2cd064ad0a3f5c35 != ObjectValue && __obf_2cd064ad0a3f5c35 != NilValue {
		__obf_0da8c843dad13139.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_0da8c843dad13139.
		Skip()
}

type __obf_13a378ace127c1a7 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_f9abf27a7d87d6ae int64
	__obf_fbbd552d8d48f5f2 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_13a378ace127c1a7) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		if __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() == __obf_801f19702638809c.__obf_f9abf27a7d87d6ae {
			__obf_801f19702638809c.__obf_fbbd552d8d48f5f2.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		} else {
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_37dd7161a98d005f struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_37dd7161a98d005f) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_3cdbf5fc3dccb593 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_3cdbf5fc3dccb593) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_bf2641561b45c065 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_bf2641561b45c065) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_3803a5a30fcafd24 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_3803a5a30fcafd24) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_046e872add4a19f4 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
	__obf_784afe9dbdec2f98 int64
	__obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_046e872add4a19f4) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_784afe9dbdec2f98:
			__obf_801f19702638809c.__obf_b938864c5f5e8dbd.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_be3b01ef821c1fe6 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
	__obf_784afe9dbdec2f98 int64
	__obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
	__obf_bcf3cd5cd1946ca6 int64
	__obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_be3b01ef821c1fe6) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_784afe9dbdec2f98:
			__obf_801f19702638809c.__obf_b938864c5f5e8dbd.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_bcf3cd5cd1946ca6:
			__obf_801f19702638809c.__obf_95e559f2e7b06d72.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_8239535c6e562d8f struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
	__obf_784afe9dbdec2f98 int64
	__obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
	__obf_bcf3cd5cd1946ca6 int64
	__obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
	__obf_f31e5669da3b679f int64
	__obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_8239535c6e562d8f) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_784afe9dbdec2f98:
			__obf_801f19702638809c.__obf_b938864c5f5e8dbd.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_bcf3cd5cd1946ca6:
			__obf_801f19702638809c.__obf_95e559f2e7b06d72.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_f31e5669da3b679f:
			__obf_801f19702638809c.__obf_e24e7b3ac2913a05.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_c1bdac930f0bb76f struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
	__obf_784afe9dbdec2f98 int64
	__obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
	__obf_bcf3cd5cd1946ca6 int64
	__obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
	__obf_f31e5669da3b679f int64
	__obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
	__obf_3c50d966f76dbc94 int64
	__obf_6549f536b94c945d *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_c1bdac930f0bb76f) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_784afe9dbdec2f98:
			__obf_801f19702638809c.__obf_b938864c5f5e8dbd.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_bcf3cd5cd1946ca6:
			__obf_801f19702638809c.__obf_95e559f2e7b06d72.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_f31e5669da3b679f:
			__obf_801f19702638809c.__obf_e24e7b3ac2913a05.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3c50d966f76dbc94:
			__obf_801f19702638809c.__obf_6549f536b94c945d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_554a4f65c9d0ea9b struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_08b36383be79e662 int64
	__obf_07c5ed10b4f8401d *__obf_7ca3f82b6bdb44f7
	__obf_d1ce14cb84761def int64
	__obf_39d063eb592af548 *__obf_7ca3f82b6bdb44f7
	__obf_7a7c68adffeb0f0b int64
	__obf_ace5a68d76e4d925 *__obf_7ca3f82b6bdb44f7
	__obf_3538c6ab82297763 int64
	__obf_a27ca7588c1f1895 *__obf_7ca3f82b6bdb44f7
	__obf_a8a1d03617db22a4 int64
	__obf_0c438d080933cdbe *__obf_7ca3f82b6bdb44f7
	__obf_784afe9dbdec2f98 int64
	__obf_b938864c5f5e8dbd *__obf_7ca3f82b6bdb44f7
	__obf_bcf3cd5cd1946ca6 int64
	__obf_95e559f2e7b06d72 *__obf_7ca3f82b6bdb44f7
	__obf_f31e5669da3b679f int64
	__obf_e24e7b3ac2913a05 *__obf_7ca3f82b6bdb44f7
	__obf_3c50d966f76dbc94 int64
	__obf_6549f536b94c945d *__obf_7ca3f82b6bdb44f7
	__obf_2cd3b104921b349e int64
	__obf_b4ad13b78d729536 *__obf_7ca3f82b6bdb44f7
}

func (__obf_801f19702638809c *__obf_554a4f65c9d0ea9b) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.__obf_0454a225ebb329d1() {
		return
	}
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		switch __obf_0da8c843dad13139.__obf_a4ae174d25635aaf() {
		case __obf_801f19702638809c.__obf_08b36383be79e662:
			__obf_801f19702638809c.__obf_07c5ed10b4f8401d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_d1ce14cb84761def:
			__obf_801f19702638809c.__obf_39d063eb592af548.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_7a7c68adffeb0f0b:
			__obf_801f19702638809c.__obf_ace5a68d76e4d925.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3538c6ab82297763:
			__obf_801f19702638809c.__obf_a27ca7588c1f1895.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_a8a1d03617db22a4:
			__obf_801f19702638809c.__obf_0c438d080933cdbe.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_784afe9dbdec2f98:
			__obf_801f19702638809c.__obf_b938864c5f5e8dbd.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_bcf3cd5cd1946ca6:
			__obf_801f19702638809c.__obf_95e559f2e7b06d72.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_f31e5669da3b679f:
			__obf_801f19702638809c.__obf_e24e7b3ac2913a05.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_3c50d966f76dbc94:
			__obf_801f19702638809c.__obf_6549f536b94c945d.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		case __obf_801f19702638809c.__obf_2cd3b104921b349e:
			__obf_801f19702638809c.__obf_b4ad13b78d729536.
				Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		default:
			__obf_0da8c843dad13139.
				Skip()
		}
		if __obf_0da8c843dad13139.__obf_5059d0524ee888e7() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF && len(__obf_801f19702638809c.__obf_edcded11af6ebc4c.Type1().Name()) != 0 {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v.%s", __obf_801f19702638809c.__obf_edcded11af6ebc4c, __obf_0da8c843dad13139.Error.Error())
	}
	__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
}

type __obf_7ca3f82b6bdb44f7 struct {
	__obf_ad34f8a6de2b01cb reflect2.StructField
	__obf_fbbd552d8d48f5f2 ValDecoder
}

func (__obf_801f19702638809c *__obf_7ca3f82b6bdb44f7) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_c34985751cfe0592 := __obf_801f19702638809c.__obf_ad34f8a6de2b01cb.UnsafeGet(__obf_575c04f2b9d91315)
	__obf_801f19702638809c.__obf_fbbd552d8d48f5f2.
		Decode(__obf_c34985751cfe0592, __obf_0da8c843dad13139)
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%s: %s", __obf_801f19702638809c.__obf_ad34f8a6de2b01cb.Name(), __obf_0da8c843dad13139.Error.Error())
	}
}

type __obf_df799b1b6fef897a struct {
	__obf_5b1e51b09d209694 ValDecoder
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
}

func (__obf_801f19702638809c *__obf_df799b1b6fef897a) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_801f19702638809c.__obf_5b1e51b09d209694.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	__obf_a3eaafc22faf1644 := *((*string)(__obf_575c04f2b9d91315))
	__obf_7cce46f505b0dcfb := __obf_801f19702638809c.__obf_c52e0bcfad4b8b71.BorrowIterator([]byte(__obf_a3eaafc22faf1644))
	defer __obf_801f19702638809c.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_7cce46f505b0dcfb)
	*((*string)(__obf_575c04f2b9d91315)) = __obf_7cce46f505b0dcfb.ReadString()
}

type __obf_cdd0c9c1c11eaa23 struct {
	__obf_5b1e51b09d209694 ValDecoder
}

func (__obf_801f19702638809c *__obf_cdd0c9c1c11eaa23) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.WhatIsNext() == NilValue {
		__obf_801f19702638809c.__obf_5b1e51b09d209694.
			Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
		return
	}
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != '"' {
		__obf_0da8c843dad13139.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
	__obf_801f19702638809c.__obf_5b1e51b09d209694.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	if __obf_0da8c843dad13139.Error != nil {
		return
	}
	__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
	if __obf_12577c03a12f0d2c != '"' {
		__obf_0da8c843dad13139.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
}
