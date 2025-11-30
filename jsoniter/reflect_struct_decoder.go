package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_a2a209dd7fe7e0d3(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_e52c8293b0cb7ead := map[string]*Binding{}
	__obf_832a553f1a21f913 := __obf_0446b148b9ab4187(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
		for _, __obf_8a7b137df000d098 := range __obf_e94d87391e53ee63.FromNames {
			__obf_b2f3b4aadfe3e860 := __obf_e52c8293b0cb7ead[__obf_8a7b137df000d098]
			if __obf_b2f3b4aadfe3e860 == nil {
				__obf_e52c8293b0cb7ead[__obf_8a7b137df000d098] = __obf_e94d87391e53ee63
				continue
			}
			__obf_265bf31f944f5bd2, __obf_60cf2537b91343db := __obf_f52233c898eec9eb(__obf_62435d89ebefd5aa.__obf_3a25fb4c9a8342bb, __obf_b2f3b4aadfe3e860, __obf_e94d87391e53ee63)
			if __obf_265bf31f944f5bd2 {
				delete(__obf_e52c8293b0cb7ead, __obf_8a7b137df000d098)
			}
			if !__obf_60cf2537b91343db {
				__obf_e52c8293b0cb7ead[__obf_8a7b137df000d098] = __obf_e94d87391e53ee63
			}
		}
	}
	__obf_ec908eb5c5599686 := map[string]*__obf_ca6a50aada57ad9b{}
	for __obf_7814f735b261cef8, __obf_e94d87391e53ee63 := range __obf_e52c8293b0cb7ead {
		__obf_ec908eb5c5599686[__obf_7814f735b261cef8] = __obf_e94d87391e53ee63.Decoder.(*__obf_ca6a50aada57ad9b)
	}

	if !__obf_62435d89ebefd5aa.__obf_b600271a247f99b8() {
		for __obf_7814f735b261cef8, __obf_e94d87391e53ee63 := range __obf_e52c8293b0cb7ead {
			if _, __obf_da9e40840d5dbfdf := __obf_ec908eb5c5599686[strings.ToLower(__obf_7814f735b261cef8)]; !__obf_da9e40840d5dbfdf {
				__obf_ec908eb5c5599686[strings.ToLower(__obf_7814f735b261cef8)] = __obf_e94d87391e53ee63.Decoder.(*__obf_ca6a50aada57ad9b)
			}
		}
	}

	return __obf_3ea60b046895c5a8(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686)
}

func __obf_3ea60b046895c5a8(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type, __obf_ec908eb5c5599686 map[string]*__obf_ca6a50aada57ad9b) ValDecoder {
	if __obf_62435d89ebefd5aa.__obf_97cd803c79c71ab5 {
		return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7: __obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686: __obf_ec908eb5c5599686, __obf_97cd803c79c71ab5: true}
	}
	__obf_83ae60110136c2d5 := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_ec908eb5c5599686) {
	case 0:
		return &__obf_1fcae6b0a6f8dac0{__obf_8667d4fc2a95b0d7}
	case 1:
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			return &__obf_394a691f0c2a0074{__obf_8667d4fc2a95b0d7, __obf_77787447179deab1, __obf_0d55b1266cf7ca0d}
		}
	case 2:
		var __obf_80ba3d097a22adc3 int64
		var __obf_f542833a7f7ab361 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_80ba3d097a22adc3 == 0 {
				__obf_80ba3d097a22adc3 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else {
				__obf_f542833a7f7ab361 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_ae44fb0644269343{__obf_8667d4fc2a95b0d7, __obf_80ba3d097a22adc3, __obf_adeeb8511fb851cf, __obf_f542833a7f7ab361, __obf_1531e0a9b2cac3b5}
	case 3:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_a611a8eb3931970b{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489}
	case 4:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_6ed603c8c6b3225b{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a}
	case 5:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_39da693ba866eed0{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42}
	case 6:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_08df7824a595b962 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		var __obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else if __obf_c50e0205131fb177 == 0 {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_08df7824a595b962 = __obf_77787447179deab1
				__obf_b865d3b543e5432e = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_4efeab422beaa103{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42, __obf_08df7824a595b962, __obf_b865d3b543e5432e}
	case 7:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_08df7824a595b962 int64
		var __obf_20362d6a24e302dd int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		var __obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
		var __obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else if __obf_c50e0205131fb177 == 0 {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			} else if __obf_08df7824a595b962 == 0 {
				__obf_08df7824a595b962 = __obf_77787447179deab1
				__obf_b865d3b543e5432e = __obf_0d55b1266cf7ca0d
			} else {
				__obf_20362d6a24e302dd = __obf_77787447179deab1
				__obf_03749b5b4c0af342 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_199466caa2de845c{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42, __obf_08df7824a595b962, __obf_b865d3b543e5432e, __obf_20362d6a24e302dd, __obf_03749b5b4c0af342}
	case 8:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_08df7824a595b962 int64
		var __obf_20362d6a24e302dd int64
		var __obf_8b5a8703ac859899 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		var __obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
		var __obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
		var __obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else if __obf_c50e0205131fb177 == 0 {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			} else if __obf_08df7824a595b962 == 0 {
				__obf_08df7824a595b962 = __obf_77787447179deab1
				__obf_b865d3b543e5432e = __obf_0d55b1266cf7ca0d
			} else if __obf_20362d6a24e302dd == 0 {
				__obf_20362d6a24e302dd = __obf_77787447179deab1
				__obf_03749b5b4c0af342 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_8b5a8703ac859899 = __obf_77787447179deab1
				__obf_c3cb54eefedae2d6 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_4af35830845f5846{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42, __obf_08df7824a595b962, __obf_b865d3b543e5432e, __obf_20362d6a24e302dd, __obf_03749b5b4c0af342, __obf_8b5a8703ac859899, __obf_c3cb54eefedae2d6}
	case 9:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_08df7824a595b962 int64
		var __obf_20362d6a24e302dd int64
		var __obf_8b5a8703ac859899 int64
		var __obf_1442ab6c8e0abe98 int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		var __obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
		var __obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
		var __obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
		var __obf_03385efce74aa190 *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else if __obf_c50e0205131fb177 == 0 {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			} else if __obf_08df7824a595b962 == 0 {
				__obf_08df7824a595b962 = __obf_77787447179deab1
				__obf_b865d3b543e5432e = __obf_0d55b1266cf7ca0d
			} else if __obf_20362d6a24e302dd == 0 {
				__obf_20362d6a24e302dd = __obf_77787447179deab1
				__obf_03749b5b4c0af342 = __obf_0d55b1266cf7ca0d
			} else if __obf_8b5a8703ac859899 == 0 {
				__obf_8b5a8703ac859899 = __obf_77787447179deab1
				__obf_c3cb54eefedae2d6 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_1442ab6c8e0abe98 = __obf_77787447179deab1
				__obf_03385efce74aa190 = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_e6c600f953a00dc8{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42, __obf_08df7824a595b962, __obf_b865d3b543e5432e, __obf_20362d6a24e302dd, __obf_03749b5b4c0af342, __obf_8b5a8703ac859899, __obf_c3cb54eefedae2d6, __obf_1442ab6c8e0abe98, __obf_03385efce74aa190}
	case 10:
		var __obf_e7ed10634c0c04f1 int64
		var __obf_e52c309437628c60 int64
		var __obf_36922032f49049a3 int64
		var __obf_cff4e5d236883149 int64
		var __obf_c50e0205131fb177 int64
		var __obf_08df7824a595b962 int64
		var __obf_20362d6a24e302dd int64
		var __obf_8b5a8703ac859899 int64
		var __obf_1442ab6c8e0abe98 int64
		var __obf_8048d8eddccf451e int64
		var __obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
		var __obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
		var __obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
		var __obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
		var __obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
		var __obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
		var __obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
		var __obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
		var __obf_03385efce74aa190 *__obf_ca6a50aada57ad9b
		var __obf_c4bdf19cb264e78b *__obf_ca6a50aada57ad9b
		for __obf_35302245dbc18a32, __obf_0d55b1266cf7ca0d := range __obf_ec908eb5c5599686 {
			__obf_77787447179deab1 := __obf_5651d62f410f485f(__obf_35302245dbc18a32, __obf_62435d89ebefd5aa.__obf_b600271a247f99b8())
			_, __obf_cab734fd24f472a8 := __obf_83ae60110136c2d5[__obf_77787447179deab1]
			if __obf_cab734fd24f472a8 {
				return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
			}
			__obf_83ae60110136c2d5[__obf_77787447179deab1] = struct{}{}
			if __obf_e7ed10634c0c04f1 == 0 {
				__obf_e7ed10634c0c04f1 = __obf_77787447179deab1
				__obf_adeeb8511fb851cf = __obf_0d55b1266cf7ca0d
			} else if __obf_e52c309437628c60 == 0 {
				__obf_e52c309437628c60 = __obf_77787447179deab1
				__obf_1531e0a9b2cac3b5 = __obf_0d55b1266cf7ca0d
			} else if __obf_36922032f49049a3 == 0 {
				__obf_36922032f49049a3 = __obf_77787447179deab1
				__obf_4ed52a0566e9f489 = __obf_0d55b1266cf7ca0d
			} else if __obf_cff4e5d236883149 == 0 {
				__obf_cff4e5d236883149 = __obf_77787447179deab1
				__obf_1a44ad95b1157a5a = __obf_0d55b1266cf7ca0d
			} else if __obf_c50e0205131fb177 == 0 {
				__obf_c50e0205131fb177 = __obf_77787447179deab1
				__obf_ec363ed5fa0eac42 = __obf_0d55b1266cf7ca0d
			} else if __obf_08df7824a595b962 == 0 {
				__obf_08df7824a595b962 = __obf_77787447179deab1
				__obf_b865d3b543e5432e = __obf_0d55b1266cf7ca0d
			} else if __obf_20362d6a24e302dd == 0 {
				__obf_20362d6a24e302dd = __obf_77787447179deab1
				__obf_03749b5b4c0af342 = __obf_0d55b1266cf7ca0d
			} else if __obf_8b5a8703ac859899 == 0 {
				__obf_8b5a8703ac859899 = __obf_77787447179deab1
				__obf_c3cb54eefedae2d6 = __obf_0d55b1266cf7ca0d
			} else if __obf_1442ab6c8e0abe98 == 0 {
				__obf_1442ab6c8e0abe98 = __obf_77787447179deab1
				__obf_03385efce74aa190 = __obf_0d55b1266cf7ca0d
			} else {
				__obf_8048d8eddccf451e = __obf_77787447179deab1
				__obf_c4bdf19cb264e78b = __obf_0d55b1266cf7ca0d
			}
		}
		return &__obf_db9a44c878af7fdc{__obf_8667d4fc2a95b0d7, __obf_e7ed10634c0c04f1, __obf_adeeb8511fb851cf, __obf_e52c309437628c60, __obf_1531e0a9b2cac3b5, __obf_36922032f49049a3, __obf_4ed52a0566e9f489, __obf_cff4e5d236883149, __obf_1a44ad95b1157a5a, __obf_c50e0205131fb177, __obf_ec363ed5fa0eac42, __obf_08df7824a595b962, __obf_b865d3b543e5432e, __obf_20362d6a24e302dd, __obf_03749b5b4c0af342, __obf_8b5a8703ac859899, __obf_c3cb54eefedae2d6, __obf_1442ab6c8e0abe98, __obf_03385efce74aa190, __obf_8048d8eddccf451e, __obf_c4bdf19cb264e78b}
	}
	return &__obf_d10d42839a7a2650{__obf_8667d4fc2a95b0d7, __obf_ec908eb5c5599686, false}
}

type __obf_d10d42839a7a2650 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_ec908eb5c5599686 map[string]*__obf_ca6a50aada57ad9b
	__obf_97cd803c79c71ab5 bool
}

func (__obf_924cc7ef585cdfb0 *__obf_d10d42839a7a2650) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	var __obf_0c1bc1e511a43120 byte
	for __obf_0c1bc1e511a43120 = ','; __obf_0c1bc1e511a43120 == ','; __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() {
		__obf_924cc7ef585cdfb0.__obf_c3f2f9658b320249(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	if __obf_0c1bc1e511a43120 != '}' {
		__obf_edd9fe48d39445e4.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

func (__obf_924cc7ef585cdfb0 *__obf_d10d42839a7a2650) __obf_c3f2f9658b320249(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	var __obf_f992c91036745ca0 string
	var __obf_0d55b1266cf7ca0d *__obf_ca6a50aada57ad9b
	if __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_ccb7c8499611626f {
		__obf_88503e6d5ff428ea := __obf_edd9fe48d39445e4.ReadStringAsSlice()
		__obf_f992c91036745ca0 = *(*string)(unsafe.Pointer(&__obf_88503e6d5ff428ea))
		__obf_0d55b1266cf7ca0d = __obf_924cc7ef585cdfb0.__obf_ec908eb5c5599686[__obf_f992c91036745ca0]
		if __obf_0d55b1266cf7ca0d == nil && !__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_b600271a247f99b8 {
			__obf_0d55b1266cf7ca0d = __obf_924cc7ef585cdfb0.__obf_ec908eb5c5599686[strings.ToLower(__obf_f992c91036745ca0)]
		}
	} else {
		__obf_f992c91036745ca0 = __obf_edd9fe48d39445e4.ReadString()
		__obf_0d55b1266cf7ca0d = __obf_924cc7ef585cdfb0.__obf_ec908eb5c5599686[__obf_f992c91036745ca0]
		if __obf_0d55b1266cf7ca0d == nil && !__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_b600271a247f99b8 {
			__obf_0d55b1266cf7ca0d = __obf_924cc7ef585cdfb0.__obf_ec908eb5c5599686[strings.ToLower(__obf_f992c91036745ca0)]
		}
	}
	if __obf_0d55b1266cf7ca0d == nil {
		if __obf_924cc7ef585cdfb0.__obf_97cd803c79c71ab5 {
			__obf_864e3235428d31a6 := "found unknown field: " + __obf_f992c91036745ca0
			__obf_edd9fe48d39445e4.
				ReportError("ReadObject", __obf_864e3235428d31a6)
		}
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 != ':' {
			__obf_edd9fe48d39445e4.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		}
		__obf_edd9fe48d39445e4.
			Skip()
		return
	}
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != ':' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
	}
	__obf_0d55b1266cf7ca0d.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
}

type __obf_1fcae6b0a6f8dac0 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
}

func (__obf_924cc7ef585cdfb0 *__obf_1fcae6b0a6f8dac0) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_64fa966d17d74b1c := __obf_edd9fe48d39445e4.WhatIsNext()
	if __obf_64fa966d17d74b1c != ObjectValue && __obf_64fa966d17d74b1c != NilValue {
		__obf_edd9fe48d39445e4.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_edd9fe48d39445e4.
		Skip()
}

type __obf_394a691f0c2a0074 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_77787447179deab1 int64
	__obf_0d55b1266cf7ca0d *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_394a691f0c2a0074) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		if __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() == __obf_924cc7ef585cdfb0.__obf_77787447179deab1 {
			__obf_924cc7ef585cdfb0.__obf_0d55b1266cf7ca0d.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		} else {
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_ae44fb0644269343 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_ae44fb0644269343) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_a611a8eb3931970b struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_a611a8eb3931970b) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_6ed603c8c6b3225b struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_6ed603c8c6b3225b) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_39da693ba866eed0 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_39da693ba866eed0) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_4efeab422beaa103 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
	__obf_9f3edf84e5aec8b6 int64
	__obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_4efeab422beaa103) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_9f3edf84e5aec8b6:
			__obf_924cc7ef585cdfb0.__obf_b865d3b543e5432e.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_199466caa2de845c struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
	__obf_9f3edf84e5aec8b6 int64
	__obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
	__obf_7413be796031f5af int64
	__obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_199466caa2de845c) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_9f3edf84e5aec8b6:
			__obf_924cc7ef585cdfb0.__obf_b865d3b543e5432e.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_7413be796031f5af:
			__obf_924cc7ef585cdfb0.__obf_03749b5b4c0af342.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_4af35830845f5846 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
	__obf_9f3edf84e5aec8b6 int64
	__obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
	__obf_7413be796031f5af int64
	__obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
	__obf_ef1fd2d286316b52 int64
	__obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_4af35830845f5846) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_9f3edf84e5aec8b6:
			__obf_924cc7ef585cdfb0.__obf_b865d3b543e5432e.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_7413be796031f5af:
			__obf_924cc7ef585cdfb0.__obf_03749b5b4c0af342.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_ef1fd2d286316b52:
			__obf_924cc7ef585cdfb0.__obf_c3cb54eefedae2d6.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_e6c600f953a00dc8 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
	__obf_9f3edf84e5aec8b6 int64
	__obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
	__obf_7413be796031f5af int64
	__obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
	__obf_ef1fd2d286316b52 int64
	__obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
	__obf_91f875818953541c int64
	__obf_03385efce74aa190 *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_e6c600f953a00dc8) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_9f3edf84e5aec8b6:
			__obf_924cc7ef585cdfb0.__obf_b865d3b543e5432e.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_7413be796031f5af:
			__obf_924cc7ef585cdfb0.__obf_03749b5b4c0af342.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_ef1fd2d286316b52:
			__obf_924cc7ef585cdfb0.__obf_c3cb54eefedae2d6.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_91f875818953541c:
			__obf_924cc7ef585cdfb0.__obf_03385efce74aa190.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_db9a44c878af7fdc struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_80ba3d097a22adc3 int64
	__obf_adeeb8511fb851cf *__obf_ca6a50aada57ad9b
	__obf_f542833a7f7ab361 int64
	__obf_1531e0a9b2cac3b5 *__obf_ca6a50aada57ad9b
	__obf_09b21731ae601b06 int64
	__obf_4ed52a0566e9f489 *__obf_ca6a50aada57ad9b
	__obf_f24c5b79fa32cd3b int64
	__obf_1a44ad95b1157a5a *__obf_ca6a50aada57ad9b
	__obf_cb6185bf7bec1ccb int64
	__obf_ec363ed5fa0eac42 *__obf_ca6a50aada57ad9b
	__obf_9f3edf84e5aec8b6 int64
	__obf_b865d3b543e5432e *__obf_ca6a50aada57ad9b
	__obf_7413be796031f5af int64
	__obf_03749b5b4c0af342 *__obf_ca6a50aada57ad9b
	__obf_ef1fd2d286316b52 int64
	__obf_c3cb54eefedae2d6 *__obf_ca6a50aada57ad9b
	__obf_91f875818953541c int64
	__obf_03385efce74aa190 *__obf_ca6a50aada57ad9b
	__obf_afb9e70deeefe1b4 int64
	__obf_c4bdf19cb264e78b *__obf_ca6a50aada57ad9b
}

func (__obf_924cc7ef585cdfb0 *__obf_db9a44c878af7fdc) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.__obf_870a32c6708136fd() {
		return
	}
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		switch __obf_edd9fe48d39445e4.__obf_18f2978e30c0e899() {
		case __obf_924cc7ef585cdfb0.__obf_80ba3d097a22adc3:
			__obf_924cc7ef585cdfb0.__obf_adeeb8511fb851cf.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f542833a7f7ab361:
			__obf_924cc7ef585cdfb0.__obf_1531e0a9b2cac3b5.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_09b21731ae601b06:
			__obf_924cc7ef585cdfb0.__obf_4ed52a0566e9f489.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_f24c5b79fa32cd3b:
			__obf_924cc7ef585cdfb0.__obf_1a44ad95b1157a5a.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_cb6185bf7bec1ccb:
			__obf_924cc7ef585cdfb0.__obf_ec363ed5fa0eac42.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_9f3edf84e5aec8b6:
			__obf_924cc7ef585cdfb0.__obf_b865d3b543e5432e.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_7413be796031f5af:
			__obf_924cc7ef585cdfb0.__obf_03749b5b4c0af342.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_ef1fd2d286316b52:
			__obf_924cc7ef585cdfb0.__obf_c3cb54eefedae2d6.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_91f875818953541c:
			__obf_924cc7ef585cdfb0.__obf_03385efce74aa190.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		case __obf_924cc7ef585cdfb0.__obf_afb9e70deeefe1b4:
			__obf_924cc7ef585cdfb0.__obf_c4bdf19cb264e78b.
				Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		default:
			__obf_edd9fe48d39445e4.
				Skip()
		}
		if __obf_edd9fe48d39445e4.__obf_9edfbea94670bd7a() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF && len(__obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7.Type1().Name()) != 0 {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v.%s", __obf_924cc7ef585cdfb0.__obf_8667d4fc2a95b0d7, __obf_edd9fe48d39445e4.Error.Error())
	}
	__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
}

type __obf_ca6a50aada57ad9b struct {
	__obf_f992c91036745ca0 reflect2.StructField
	__obf_0d55b1266cf7ca0d ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_ca6a50aada57ad9b) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_f4916a65ae1a5a5a := __obf_924cc7ef585cdfb0.__obf_f992c91036745ca0.UnsafeGet(__obf_753ab3fb72cdd659)
	__obf_924cc7ef585cdfb0.__obf_0d55b1266cf7ca0d.
		Decode(__obf_f4916a65ae1a5a5a, __obf_edd9fe48d39445e4)
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%s: %s", __obf_924cc7ef585cdfb0.__obf_f992c91036745ca0.Name(), __obf_edd9fe48d39445e4.Error.Error())
	}
}

type __obf_ed19b45f4bdc0c0f struct {
	__obf_fb5db223a2c09df6 ValDecoder
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
}

func (__obf_924cc7ef585cdfb0 *__obf_ed19b45f4bdc0c0f) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	__obf_2d944450d48e5810 := *((*string)(__obf_753ab3fb72cdd659))
	__obf_bc00b1b9eb11ae2b := __obf_924cc7ef585cdfb0.__obf_0c8065c219ccf0ab.BorrowIterator([]byte(__obf_2d944450d48e5810))
	defer __obf_924cc7ef585cdfb0.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_bc00b1b9eb11ae2b)
	*((*string)(__obf_753ab3fb72cdd659)) = __obf_bc00b1b9eb11ae2b.ReadString()
}

type __obf_cabfece1b2ec841f struct {
	__obf_fb5db223a2c09df6 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_cabfece1b2ec841f) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.WhatIsNext() == NilValue {
		__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
			Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
		return
	}
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != '"' {
		__obf_edd9fe48d39445e4.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	if __obf_edd9fe48d39445e4.Error != nil {
		return
	}
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
	if __obf_0c1bc1e511a43120 != '"' {
		__obf_edd9fe48d39445e4.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
}
