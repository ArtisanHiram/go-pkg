/*
Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package __obf_d24101647651f4c3

import "image"

// Keep value in [0,255] range.
func __obf_eb593ac00ffa365d(__obf_2a764779fc1401ce int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_2a764779fc1401ce) < 256 {
		return uint8(__obf_2a764779fc1401ce)
	}
	if __obf_2a764779fc1401ce > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_a2d20f4c701b7468(__obf_2a764779fc1401ce int64) uint16 {
	if uint64(__obf_2a764779fc1401ce) < 65536 {
		return uint16(__obf_2a764779fc1401ce)
	}
	if __obf_2a764779fc1401ce > 65535 {
		return 65535
	}
	return 0
}

func __obf_989e6431922ef0ad(__obf_2a764779fc1401ce image.Image, __obf_d5c26403b97b0f9a *image.RGBA64, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int32, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_9544fde7c0f8af7b [4]int64
			var __obf_d129597213b1cd02 int64
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case __obf_b6f83e69eabfe073 < 0:
						__obf_b6f83e69eabfe073 = 0
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = __obf_0fb6ee372c8d553c
					}
					__obf_13d76cd467248ddd, __obf_810c89f32765412d, __obf_87ac9dbaa8f66b85, __obf_637dc46ce0f08c96 := __obf_2a764779fc1401ce.At(__obf_b6f83e69eabfe073+__obf_2a764779fc1401ce.Bounds().Min.X, __obf_88d1a0c3dc734ee1+__obf_2a764779fc1401ce.Bounds().Min.Y).RGBA()
					__obf_9544fde7c0f8af7b[0] += int64(__obf_b83b1b352940ee14) * int64(__obf_13d76cd467248ddd)
					__obf_9544fde7c0f8af7b[1] += int64(__obf_b83b1b352940ee14) * int64(__obf_810c89f32765412d)
					__obf_9544fde7c0f8af7b[2] += int64(__obf_b83b1b352940ee14) * int64(__obf_87ac9dbaa8f66b85)
					__obf_9544fde7c0f8af7b[3] += int64(__obf_b83b1b352940ee14) * int64(__obf_637dc46ce0f08c96)
					__obf_d129597213b1cd02 += int64(__obf_b83b1b352940ee14)
				}
			}
			__obf_04772264944a3f8f := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*8
			__obf_c0e12478b15335b8 := __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+0] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+1] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+2] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+3] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[2] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+4] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+5] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[3] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+6] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+7] = uint8(__obf_c0e12478b15335b8)
		}
	}
}

func __obf_8ec8d60aa6453477(__obf_2a764779fc1401ce *image.RGBA, __obf_d5c26403b97b0f9a *image.RGBA, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int16, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_9544fde7c0f8af7b [4]int32
			var __obf_d129597213b1cd02 int32
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 4
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 4 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_9544fde7c0f8af7b[0] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])
					__obf_9544fde7c0f8af7b[1] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1])
					__obf_9544fde7c0f8af7b[2] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2])
					__obf_9544fde7c0f8af7b[3] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+3])
					__obf_d129597213b1cd02 += int32(__obf_b83b1b352940ee14)
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*4
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[2] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+3] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[3] / __obf_d129597213b1cd02)
		}
	}
}

func __obf_f38f8e4e927a9b1f(__obf_2a764779fc1401ce *image.NRGBA, __obf_d5c26403b97b0f9a *image.RGBA, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int16, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_9544fde7c0f8af7b [4]int32
			var __obf_d129597213b1cd02 int32
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 4
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 4 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_637dc46ce0f08c96 :=// Forward alpha-premultiplication
					int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+3])
					__obf_13d76cd467248ddd := int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0]) * __obf_637dc46ce0f08c96
					__obf_13d76cd467248ddd /= 0xff
					__obf_810c89f32765412d := int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1]) * __obf_637dc46ce0f08c96
					__obf_810c89f32765412d /= 0xff
					__obf_87ac9dbaa8f66b85 := int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2]) * __obf_637dc46ce0f08c96
					__obf_87ac9dbaa8f66b85 /= 0xff
					__obf_9544fde7c0f8af7b[0] += int32(__obf_b83b1b352940ee14) * __obf_13d76cd467248ddd
					__obf_9544fde7c0f8af7b[1] += int32(__obf_b83b1b352940ee14) * __obf_810c89f32765412d
					__obf_9544fde7c0f8af7b[2] += int32(__obf_b83b1b352940ee14) * __obf_87ac9dbaa8f66b85
					__obf_9544fde7c0f8af7b[3] += int32(__obf_b83b1b352940ee14) * __obf_637dc46ce0f08c96
					__obf_d129597213b1cd02 += int32(__obf_b83b1b352940ee14)
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*4
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[2] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+3] = __obf_eb593ac00ffa365d(__obf_9544fde7c0f8af7b[3] / __obf_d129597213b1cd02)
		}
	}
}

func __obf_454a0ae091c9479b(__obf_2a764779fc1401ce *image.RGBA64, __obf_d5c26403b97b0f9a *image.RGBA64, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int32, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_9544fde7c0f8af7b [4]int64
			var __obf_d129597213b1cd02 int64
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 8
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 8 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_9544fde7c0f8af7b[0] += int64(__obf_b83b1b352940ee14) * (int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])<<8 | int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1]))
					__obf_9544fde7c0f8af7b[1] += int64(__obf_b83b1b352940ee14) * (int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2])<<8 | int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+3]))
					__obf_9544fde7c0f8af7b[2] += int64(__obf_b83b1b352940ee14) * (int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+4])<<8 | int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+5]))
					__obf_9544fde7c0f8af7b[3] += int64(__obf_b83b1b352940ee14) * (int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+6])<<8 | int64(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+7]))
					__obf_d129597213b1cd02 += int64(__obf_b83b1b352940ee14)
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*8
			__obf_c0e12478b15335b8 := __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+3] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[2] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+4] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+5] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[3] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+6] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+7] = uint8(__obf_c0e12478b15335b8)
		}
	}
}

func __obf_a658dbd4fd6775a7(__obf_2a764779fc1401ce *image.NRGBA64, __obf_d5c26403b97b0f9a *image.RGBA64, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int32, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_9544fde7c0f8af7b [4]int64
			var __obf_d129597213b1cd02 int64
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 8
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 8 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_637dc46ce0f08c96 :=// Forward alpha-premultiplication
					int64(uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+6])<<8 | uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+7]))
					__obf_13d76cd467248ddd := int64(uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])<<8|uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1])) * __obf_637dc46ce0f08c96
					__obf_13d76cd467248ddd /= 0xffff
					__obf_810c89f32765412d := int64(uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2])<<8|uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+3])) * __obf_637dc46ce0f08c96
					__obf_810c89f32765412d /= 0xffff
					__obf_87ac9dbaa8f66b85 := int64(uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+4])<<8|uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+5])) * __obf_637dc46ce0f08c96
					__obf_87ac9dbaa8f66b85 /= 0xffff
					__obf_9544fde7c0f8af7b[0] += int64(__obf_b83b1b352940ee14) * __obf_13d76cd467248ddd
					__obf_9544fde7c0f8af7b[1] += int64(__obf_b83b1b352940ee14) * __obf_810c89f32765412d
					__obf_9544fde7c0f8af7b[2] += int64(__obf_b83b1b352940ee14) * __obf_87ac9dbaa8f66b85
					__obf_9544fde7c0f8af7b[3] += int64(__obf_b83b1b352940ee14) * __obf_637dc46ce0f08c96
					__obf_d129597213b1cd02 += int64(__obf_b83b1b352940ee14)
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*8
			__obf_c0e12478b15335b8 := __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+3] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[2] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+4] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+5] = uint8(__obf_c0e12478b15335b8)
			__obf_c0e12478b15335b8 = __obf_a2d20f4c701b7468(__obf_9544fde7c0f8af7b[3] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+6] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+7] = uint8(__obf_c0e12478b15335b8)
		}
	}
}

func __obf_1eb0358183feb59b(__obf_2a764779fc1401ce *image.Gray, __obf_d5c26403b97b0f9a *image.Gray, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int16, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[(__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_4deb34f548d8238d int32
			var __obf_d129597213b1cd02 int32
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case __obf_b6f83e69eabfe073 < 0:
						__obf_b6f83e69eabfe073 = 0
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = __obf_0fb6ee372c8d553c
					}
					__obf_4deb34f548d8238d += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073])
					__obf_d129597213b1cd02 += int32(__obf_b83b1b352940ee14)
				}
			}
			__obf_04772264944a3f8f := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1 - __obf_89dbd054048c93be.Min.X)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f] = __obf_eb593ac00ffa365d(__obf_4deb34f548d8238d / __obf_d129597213b1cd02)
		}
	}
}

func __obf_a51435ed23b2f46f(__obf_2a764779fc1401ce *image.Gray16, __obf_d5c26403b97b0f9a *image.Gray16, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int32, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_4deb34f548d8238d int64
			var __obf_d129597213b1cd02 int64
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 2
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 2 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_4deb34f548d8238d += int64(__obf_b83b1b352940ee14) * int64(uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])<<8|uint16(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1]))
					__obf_d129597213b1cd02 += int64(__obf_b83b1b352940ee14)
				}
			}
			__obf_04772264944a3f8f := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*2
			__obf_c0e12478b15335b8 := __obf_a2d20f4c701b7468(__obf_4deb34f548d8238d / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+0] = uint8(__obf_c0e12478b15335b8 >> 8)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_04772264944a3f8f+1] = uint8(__obf_c0e12478b15335b8)
		}
	}
}

func __obf_5df6e1d64c485b9f(__obf_2a764779fc1401ce *__obf_489a9befc4898050, __obf_d5c26403b97b0f9a *__obf_489a9befc4898050, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []int16, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_a78e262cc2a46e75 [3]int32
			var __obf_d129597213b1cd02 int32
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				__obf_b83b1b352940ee14 := __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8]
				if __obf_b83b1b352940ee14 != 0 {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 3
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 3 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_a78e262cc2a46e75[0] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])
					__obf_a78e262cc2a46e75[1] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1])
					__obf_a78e262cc2a46e75[2] += int32(__obf_b83b1b352940ee14) * int32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2])
					__obf_d129597213b1cd02 += int32(__obf_b83b1b352940ee14)
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*3
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = __obf_eb593ac00ffa365d(__obf_a78e262cc2a46e75[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = __obf_eb593ac00ffa365d(__obf_a78e262cc2a46e75[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = __obf_eb593ac00ffa365d(__obf_a78e262cc2a46e75[2] / __obf_d129597213b1cd02)
		}
	}
}

func __obf_8b9375bb53904eb4(__obf_2a764779fc1401ce *__obf_489a9befc4898050, __obf_d5c26403b97b0f9a *__obf_489a9befc4898050, __obf_b65438e8a147153b float64, __obf_779a7f81421e09ef []bool, __obf_04772264944a3f8f []int, __obf_277f8fd7946d0e2e int) {
	__obf_89dbd054048c93be := __obf_d5c26403b97b0f9a.Bounds()
	__obf_0fb6ee372c8d553c := __obf_2a764779fc1401ce.Bounds().Dx() - 1

	for __obf_88d1a0c3dc734ee1 := __obf_89dbd054048c93be.Min.X; __obf_88d1a0c3dc734ee1 < __obf_89dbd054048c93be.Max.X; __obf_88d1a0c3dc734ee1++ {
		__obf_414b17b06de01f62 := __obf_2a764779fc1401ce.Pix[__obf_88d1a0c3dc734ee1*__obf_2a764779fc1401ce.Stride:]
		for __obf_097022f910960674 := __obf_89dbd054048c93be.Min.Y; __obf_097022f910960674 < __obf_89dbd054048c93be.Max.Y; __obf_097022f910960674++ {
			var __obf_a78e262cc2a46e75 [3]float32
			var __obf_d129597213b1cd02 float32
			__obf_1eb9b282d993da7c := __obf_04772264944a3f8f[__obf_097022f910960674]
			__obf_c3764290f5437417 := __obf_097022f910960674 * __obf_277f8fd7946d0e2e
			for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
				if __obf_779a7f81421e09ef[__obf_c3764290f5437417+__obf_41043e9a68df34f8] {
					__obf_b6f83e69eabfe073 := __obf_1eb9b282d993da7c + __obf_41043e9a68df34f8
					switch {
					case uint(__obf_b6f83e69eabfe073) < uint(__obf_0fb6ee372c8d553c):
						__obf_b6f83e69eabfe073 *= 3
					case __obf_b6f83e69eabfe073 >= __obf_0fb6ee372c8d553c:
						__obf_b6f83e69eabfe073 = 3 * __obf_0fb6ee372c8d553c
					default:
						__obf_b6f83e69eabfe073 = 0
					}
					__obf_a78e262cc2a46e75[0] += float32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+0])
					__obf_a78e262cc2a46e75[1] += float32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+1])
					__obf_a78e262cc2a46e75[2] += float32(__obf_414b17b06de01f62[__obf_b6f83e69eabfe073+2])
					__obf_d129597213b1cd02++
				}
			}
			__obf_a2513444a19e1397 := (__obf_097022f910960674-__obf_89dbd054048c93be.Min.Y)*__obf_d5c26403b97b0f9a.Stride + (__obf_88d1a0c3dc734ee1-__obf_89dbd054048c93be.Min.X)*3
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+0] = __obf_d72ba056d092a878(__obf_a78e262cc2a46e75[0] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+1] = __obf_d72ba056d092a878(__obf_a78e262cc2a46e75[1] / __obf_d129597213b1cd02)
			__obf_d5c26403b97b0f9a.
				Pix[__obf_a2513444a19e1397+2] = __obf_d72ba056d092a878(__obf_a78e262cc2a46e75[2] / __obf_d129597213b1cd02)
		}
	}
}
