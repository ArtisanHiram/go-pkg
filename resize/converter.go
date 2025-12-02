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

package __obf_4f71249d3f545aa0

import "image"

// Keep value in [0,255] range.
func __obf_5a8863417ff3c9da(__obf_a5eab9a5623e4978 int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_a5eab9a5623e4978) < 256 {
		return uint8(__obf_a5eab9a5623e4978)
	}
	if __obf_a5eab9a5623e4978 > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_0bba1b151ef50bff(__obf_a5eab9a5623e4978 int64) uint16 {
	if uint64(__obf_a5eab9a5623e4978) < 65536 {
		return uint16(__obf_a5eab9a5623e4978)
	}
	if __obf_a5eab9a5623e4978 > 65535 {
		return 65535
	}
	return 0
}

func __obf_ccde39b12bafba9e(__obf_a5eab9a5623e4978 image.Image, __obf_dc37e46bcd46a9e8 *image.RGBA64, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int32, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]int64
			var __obf_01192c7128aa9b6c int64
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case __obf_82639a84431b18a0 < 0:
						__obf_82639a84431b18a0 = 0
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = __obf_650256bb4b73aacb
					}
					__obf_2be862e501ec49ab, __obf_c0fee89940f20f56, __obf_50b1cb88fb085316, __obf_7fddaa6cb6196f8f := __obf_a5eab9a5623e4978.At(__obf_82639a84431b18a0+__obf_a5eab9a5623e4978.Bounds().Min.X, __obf_2a2c911b5786aa81+__obf_a5eab9a5623e4978.Bounds().Min.Y).RGBA()
					__obf_17f2ca7bc5209adc[0] += int64(__obf_be3808c2782f8a12) * int64(__obf_2be862e501ec49ab)
					__obf_17f2ca7bc5209adc[1] += int64(__obf_be3808c2782f8a12) * int64(__obf_c0fee89940f20f56)
					__obf_17f2ca7bc5209adc[2] += int64(__obf_be3808c2782f8a12) * int64(__obf_50b1cb88fb085316)
					__obf_17f2ca7bc5209adc[3] += int64(__obf_be3808c2782f8a12) * int64(__obf_7fddaa6cb6196f8f)
					__obf_01192c7128aa9b6c += int64(__obf_be3808c2782f8a12)
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_b2d4489701a105b7(__obf_a5eab9a5623e4978 *image.RGBA, __obf_dc37e46bcd46a9e8 *image.RGBA, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int16, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]int32
			var __obf_01192c7128aa9b6c int32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 4
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 4 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])
					__obf_17f2ca7bc5209adc[1] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])
					__obf_17f2ca7bc5209adc[2] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])
					__obf_17f2ca7bc5209adc[3] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3])
					__obf_01192c7128aa9b6c += int32(__obf_be3808c2782f8a12)
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*4
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_e8e29c8f236bdc6f(__obf_a5eab9a5623e4978 *image.NRGBA, __obf_dc37e46bcd46a9e8 *image.RGBA, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int16, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]int32
			var __obf_01192c7128aa9b6c int32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 4
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 4 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_7fddaa6cb6196f8f :=// Forward alpha-premultiplication
					int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3])
					__obf_2be862e501ec49ab := int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0]) * __obf_7fddaa6cb6196f8f
					__obf_2be862e501ec49ab /= 0xff
					__obf_c0fee89940f20f56 := int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]) * __obf_7fddaa6cb6196f8f
					__obf_c0fee89940f20f56 /= 0xff
					__obf_50b1cb88fb085316 := int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2]) * __obf_7fddaa6cb6196f8f
					__obf_50b1cb88fb085316 /= 0xff
					__obf_17f2ca7bc5209adc[0] += int32(__obf_be3808c2782f8a12) * __obf_2be862e501ec49ab
					__obf_17f2ca7bc5209adc[1] += int32(__obf_be3808c2782f8a12) * __obf_c0fee89940f20f56
					__obf_17f2ca7bc5209adc[2] += int32(__obf_be3808c2782f8a12) * __obf_50b1cb88fb085316
					__obf_17f2ca7bc5209adc[3] += int32(__obf_be3808c2782f8a12) * __obf_7fddaa6cb6196f8f
					__obf_01192c7128aa9b6c += int32(__obf_be3808c2782f8a12)
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*4
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = __obf_5a8863417ff3c9da(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_dd75986f2d7d6f93(__obf_a5eab9a5623e4978 *image.RGBA64, __obf_dc37e46bcd46a9e8 *image.RGBA64, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int32, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]int64
			var __obf_01192c7128aa9b6c int64
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 8
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 8 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += int64(__obf_be3808c2782f8a12) * (int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8 | int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]))
					__obf_17f2ca7bc5209adc[1] += int64(__obf_be3808c2782f8a12) * (int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])<<8 | int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3]))
					__obf_17f2ca7bc5209adc[2] += int64(__obf_be3808c2782f8a12) * (int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+4])<<8 | int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+5]))
					__obf_17f2ca7bc5209adc[3] += int64(__obf_be3808c2782f8a12) * (int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+6])<<8 | int64(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+7]))
					__obf_01192c7128aa9b6c += int64(__obf_be3808c2782f8a12)
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_7dc22bccd4944b8d(__obf_a5eab9a5623e4978 *image.NRGBA64, __obf_dc37e46bcd46a9e8 *image.RGBA64, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int32, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]int64
			var __obf_01192c7128aa9b6c int64
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 8
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 8 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_7fddaa6cb6196f8f :=// Forward alpha-premultiplication
					int64(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+6])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+7]))
					__obf_2be862e501ec49ab := int64(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8|uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])) * __obf_7fddaa6cb6196f8f
					__obf_2be862e501ec49ab /= 0xffff
					__obf_c0fee89940f20f56 := int64(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])<<8|uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3])) * __obf_7fddaa6cb6196f8f
					__obf_c0fee89940f20f56 /= 0xffff
					__obf_50b1cb88fb085316 := int64(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+4])<<8|uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+5])) * __obf_7fddaa6cb6196f8f
					__obf_50b1cb88fb085316 /= 0xffff
					__obf_17f2ca7bc5209adc[0] += int64(__obf_be3808c2782f8a12) * __obf_2be862e501ec49ab
					__obf_17f2ca7bc5209adc[1] += int64(__obf_be3808c2782f8a12) * __obf_c0fee89940f20f56
					__obf_17f2ca7bc5209adc[2] += int64(__obf_be3808c2782f8a12) * __obf_50b1cb88fb085316
					__obf_17f2ca7bc5209adc[3] += int64(__obf_be3808c2782f8a12) * __obf_7fddaa6cb6196f8f
					__obf_01192c7128aa9b6c += int64(__obf_be3808c2782f8a12)
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_0bba1b151ef50bff(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_e05bbe664e33033f(__obf_a5eab9a5623e4978 *image.Gray, __obf_dc37e46bcd46a9e8 *image.Gray, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int16, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[(__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_f3b8147580ea35e0 int32
			var __obf_01192c7128aa9b6c int32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case __obf_82639a84431b18a0 < 0:
						__obf_82639a84431b18a0 = 0
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = __obf_650256bb4b73aacb
					}
					__obf_f3b8147580ea35e0 += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0])
					__obf_01192c7128aa9b6c += int32(__obf_be3808c2782f8a12)
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81 - __obf_dde4217a1d161d53.Min.X)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b] = __obf_5a8863417ff3c9da(__obf_f3b8147580ea35e0 / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_909dd2d677e37880(__obf_a5eab9a5623e4978 *image.Gray16, __obf_dc37e46bcd46a9e8 *image.Gray16, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int32, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_f3b8147580ea35e0 int64
			var __obf_01192c7128aa9b6c int64
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 2
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 2 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_f3b8147580ea35e0 += int64(__obf_be3808c2782f8a12) * int64(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8|uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]))
					__obf_01192c7128aa9b6c += int64(__obf_be3808c2782f8a12)
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*2
			__obf_8bfaa3dfdb01ed54 := __obf_0bba1b151ef50bff(__obf_f3b8147580ea35e0 / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+1] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_65afc53989013e4c(__obf_a5eab9a5623e4978 *__obf_7b573c09d347a918, __obf_dc37e46bcd46a9e8 *__obf_7b573c09d347a918, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []int16, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_6f27b861509320fe [3]int32
			var __obf_01192c7128aa9b6c int32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				__obf_be3808c2782f8a12 := __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b]
				if __obf_be3808c2782f8a12 != 0 {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 3
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 3 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_6f27b861509320fe[0] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])
					__obf_6f27b861509320fe[1] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])
					__obf_6f27b861509320fe[2] += int32(__obf_be3808c2782f8a12) * int32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])
					__obf_01192c7128aa9b6c += int32(__obf_be3808c2782f8a12)
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*3
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_5a8863417ff3c9da(__obf_6f27b861509320fe[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_5a8863417ff3c9da(__obf_6f27b861509320fe[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_5a8863417ff3c9da(__obf_6f27b861509320fe[2] / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_dee53dca87365a2a(__obf_a5eab9a5623e4978 *__obf_7b573c09d347a918, __obf_dc37e46bcd46a9e8 *__obf_7b573c09d347a918, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_6f27b861509320fe [3]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 3
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 3 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_6f27b861509320fe[0] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])
					__obf_6f27b861509320fe[1] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])
					__obf_6f27b861509320fe[2] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*3
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_45401b8f5adc53e8(__obf_6f27b861509320fe[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_45401b8f5adc53e8(__obf_6f27b861509320fe[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_45401b8f5adc53e8(__obf_6f27b861509320fe[2] / __obf_01192c7128aa9b6c)
		}
	}
}
