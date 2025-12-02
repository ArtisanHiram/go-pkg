/*
Copyright (c) 2014, Charlie Vieth <charlie.vieth@gmail.com>

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

func __obf_45401b8f5adc53e8(__obf_2a2c911b5786aa81 float32) uint8 {
	// Nearest-neighbor values are always
	// positive no need to check lower-bound.
	if __obf_2a2c911b5786aa81 > 0xfe {
		return 0xff
	}
	return uint8(__obf_2a2c911b5786aa81)
}

func __obf_8bdb84597b992660(__obf_2a2c911b5786aa81 float32) uint16 {
	if __obf_2a2c911b5786aa81 > 0xfffe {
		return 0xffff
	}
	return uint16(__obf_2a2c911b5786aa81)
}

func __obf_226995aa8d962794(__obf_a5eab9a5623e4978 image.Image, __obf_dc37e46bcd46a9e8 *image.RGBA64, __obf_eff7440668961f2c float64, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := range __obf_743fac2b8ed0fc32 {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case __obf_82639a84431b18a0 < 0:
						__obf_82639a84431b18a0 = 0
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = __obf_650256bb4b73aacb
					}
					__obf_2be862e501ec49ab, __obf_c0fee89940f20f56, __obf_50b1cb88fb085316, __obf_7fddaa6cb6196f8f := __obf_a5eab9a5623e4978.At(__obf_82639a84431b18a0+__obf_a5eab9a5623e4978.Bounds().Min.X, __obf_2a2c911b5786aa81+__obf_a5eab9a5623e4978.Bounds().Min.Y).RGBA()
					__obf_17f2ca7bc5209adc[0] += float32(__obf_2be862e501ec49ab)
					__obf_17f2ca7bc5209adc[1] += float32(__obf_c0fee89940f20f56)
					__obf_17f2ca7bc5209adc[2] += float32(__obf_50b1cb88fb085316)
					__obf_17f2ca7bc5209adc[3] += float32(__obf_7fddaa6cb6196f8f)
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_6f78347896b0579d(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.RGBA, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 4
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 4 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])
					__obf_17f2ca7bc5209adc[1] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])
					__obf_17f2ca7bc5209adc[2] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])
					__obf_17f2ca7bc5209adc[3] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3])
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*4
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_e973cc62a48a6a19(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.NRGBA, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 4
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 4 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])
					__obf_17f2ca7bc5209adc[1] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1])
					__obf_17f2ca7bc5209adc[2] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])
					__obf_17f2ca7bc5209adc[3] += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3])
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*4
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = __obf_45401b8f5adc53e8(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_fc257f8bf8c94d34(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.RGBA64, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := range __obf_743fac2b8ed0fc32 {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 8
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 8 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]))
					__obf_17f2ca7bc5209adc[1] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3]))
					__obf_17f2ca7bc5209adc[2] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+4])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+5]))
					__obf_17f2ca7bc5209adc[3] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+6])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+7]))
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_0598eeca293592f2(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.NRGBA64, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_17f2ca7bc5209adc [4]float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := range __obf_743fac2b8ed0fc32 {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 8
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 8 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_17f2ca7bc5209adc[0] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]))
					__obf_17f2ca7bc5209adc[1] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+2])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+3]))
					__obf_17f2ca7bc5209adc[2] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+4])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+5]))
					__obf_17f2ca7bc5209adc[3] += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+6])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+7]))
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_6cee7e849e76e232 := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*8
			__obf_8bfaa3dfdb01ed54 := __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[0] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+1] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[1] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+2] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+3] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[2] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+4] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+5] = uint8(__obf_8bfaa3dfdb01ed54)
			__obf_8bfaa3dfdb01ed54 = __obf_8bdb84597b992660(__obf_17f2ca7bc5209adc[3] / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+6] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_6cee7e849e76e232+7] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}

func __obf_271bfd0ed76a5978(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.Gray, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_f3b8147580ea35e0 float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := range __obf_743fac2b8ed0fc32 {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case __obf_82639a84431b18a0 < 0:
						__obf_82639a84431b18a0 = 0
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = __obf_650256bb4b73aacb
					}
					__obf_f3b8147580ea35e0 += float32(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0])
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81 - __obf_dde4217a1d161d53.Min.X)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b] = __obf_45401b8f5adc53e8(__obf_f3b8147580ea35e0 / __obf_01192c7128aa9b6c)
		}
	}
}

func __obf_1d39fd992b17e818(__obf_a5eab9a5623e4978, __obf_dc37e46bcd46a9e8 *image.Gray16, __obf_ed4a0c7a1913ca57 []bool, __obf_3a86ac65e5f0ae4b []int, __obf_743fac2b8ed0fc32 int) {
	__obf_dde4217a1d161d53 := __obf_dc37e46bcd46a9e8.Bounds()
	__obf_650256bb4b73aacb := __obf_a5eab9a5623e4978.Bounds().Dx() - 1

	for __obf_2a2c911b5786aa81 := __obf_dde4217a1d161d53.Min.X; __obf_2a2c911b5786aa81 < __obf_dde4217a1d161d53.Max.X; __obf_2a2c911b5786aa81++ {
		__obf_4bd1d2e3f338b2e9 := __obf_a5eab9a5623e4978.Pix[__obf_2a2c911b5786aa81*__obf_a5eab9a5623e4978.Stride:]
		for __obf_36c97b6a0fef1dad := __obf_dde4217a1d161d53.Min.Y; __obf_36c97b6a0fef1dad < __obf_dde4217a1d161d53.Max.Y; __obf_36c97b6a0fef1dad++ {
			var __obf_f3b8147580ea35e0 float32
			var __obf_01192c7128aa9b6c float32
			__obf_de3222b659094173 := __obf_3a86ac65e5f0ae4b[__obf_36c97b6a0fef1dad]
			__obf_08a465ba68a06bdf := __obf_36c97b6a0fef1dad * __obf_743fac2b8ed0fc32
			for __obf_b92089dde6a4a00b := range __obf_743fac2b8ed0fc32 {
				if __obf_ed4a0c7a1913ca57[__obf_08a465ba68a06bdf+__obf_b92089dde6a4a00b] {
					__obf_82639a84431b18a0 := __obf_de3222b659094173 + __obf_b92089dde6a4a00b
					switch {
					case uint(__obf_82639a84431b18a0) < uint(__obf_650256bb4b73aacb):
						__obf_82639a84431b18a0 *= 2
					case __obf_82639a84431b18a0 >= __obf_650256bb4b73aacb:
						__obf_82639a84431b18a0 = 2 * __obf_650256bb4b73aacb
					default:
						__obf_82639a84431b18a0 = 0
					}
					__obf_f3b8147580ea35e0 += float32(uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+0])<<8 | uint16(__obf_4bd1d2e3f338b2e9[__obf_82639a84431b18a0+1]))
					__obf_01192c7128aa9b6c++
				}
			}
			__obf_3a86ac65e5f0ae4b := (__obf_36c97b6a0fef1dad-__obf_dde4217a1d161d53.Min.Y)*__obf_dc37e46bcd46a9e8.Stride + (__obf_2a2c911b5786aa81-__obf_dde4217a1d161d53.Min.X)*2
			__obf_8bfaa3dfdb01ed54 := __obf_8bdb84597b992660(__obf_f3b8147580ea35e0 / __obf_01192c7128aa9b6c)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+0] = uint8(__obf_8bfaa3dfdb01ed54 >> 8)
			__obf_dc37e46bcd46a9e8.
				Pix[__obf_3a86ac65e5f0ae4b+1] = uint8(__obf_8bfaa3dfdb01ed54)
		}
	}
}
