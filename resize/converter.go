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

package __obf_42b2ccbdafaee9c3

import "image"

// Keep value in [0,255] range.
func __obf_f4cbeaeddfa5f8f1(__obf_c82b366489672bd7 int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_c82b366489672bd7) < 256 {
		return uint8(__obf_c82b366489672bd7)
	}
	if __obf_c82b366489672bd7 > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_f2dd0bd85be949f0(__obf_c82b366489672bd7 int64) uint16 {
	if uint64(__obf_c82b366489672bd7) < 65536 {
		return uint16(__obf_c82b366489672bd7)
	}
	if __obf_c82b366489672bd7 > 65535 {
		return 65535
	}
	return 0
}

func __obf_17d6ef892ec1014e(__obf_c82b366489672bd7 image.Image, __obf_eab9867e780b63d0 *image.RGBA64, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int32, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_a96897f6a99fe47b [4]int64
			var __obf_ab00e4b1393ed7cc int64
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case __obf_b7b3ec35eb2875ee < 0:
						__obf_b7b3ec35eb2875ee = 0
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = __obf_35b1276e2b34e1d8
					}
					__obf_e02b66492c64e775, __obf_641b456d4f84424d, __obf_0001854db6e58157, __obf_43312d582ce77e14 := __obf_c82b366489672bd7.At(__obf_b7b3ec35eb2875ee+__obf_c82b366489672bd7.Bounds().Min.X, __obf_9ab3da411fe8abc3+__obf_c82b366489672bd7.Bounds().Min.Y).RGBA()
					__obf_a96897f6a99fe47b[0] += int64(__obf_8748f698cbebbcf5) * int64(__obf_e02b66492c64e775)
					__obf_a96897f6a99fe47b[1] += int64(__obf_8748f698cbebbcf5) * int64(__obf_641b456d4f84424d)
					__obf_a96897f6a99fe47b[2] += int64(__obf_8748f698cbebbcf5) * int64(__obf_0001854db6e58157)
					__obf_a96897f6a99fe47b[3] += int64(__obf_8748f698cbebbcf5) * int64(__obf_43312d582ce77e14)
					__obf_ab00e4b1393ed7cc += int64(__obf_8748f698cbebbcf5)
				}
			}
			__obf_c22f1458e7509d32 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*8
			__obf_7ba5ac1ad9178cfe := __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+0] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+1] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+2] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+3] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[2] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+4] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+5] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[3] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+6] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+7] = uint8(__obf_7ba5ac1ad9178cfe)
		}
	}
}

func __obf_efed9ef13011e083(__obf_c82b366489672bd7 *image.RGBA, __obf_eab9867e780b63d0 *image.RGBA, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int16, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_a96897f6a99fe47b [4]int32
			var __obf_ab00e4b1393ed7cc int32
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 4
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 4 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_a96897f6a99fe47b[0] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])
					__obf_a96897f6a99fe47b[1] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1])
					__obf_a96897f6a99fe47b[2] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2])
					__obf_a96897f6a99fe47b[3] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+3])
					__obf_ab00e4b1393ed7cc += int32(__obf_8748f698cbebbcf5)
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*4
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[2] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+3] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[3] / __obf_ab00e4b1393ed7cc)
		}
	}
}

func __obf_bd7f833c4584c433(__obf_c82b366489672bd7 *image.NRGBA, __obf_eab9867e780b63d0 *image.RGBA, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int16, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_a96897f6a99fe47b [4]int32
			var __obf_ab00e4b1393ed7cc int32
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 4
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 4 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_43312d582ce77e14 :=// Forward alpha-premultiplication
					int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+3])
					__obf_e02b66492c64e775 := int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0]) * __obf_43312d582ce77e14
					__obf_e02b66492c64e775 /= 0xff
					__obf_641b456d4f84424d := int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1]) * __obf_43312d582ce77e14
					__obf_641b456d4f84424d /= 0xff
					__obf_0001854db6e58157 := int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2]) * __obf_43312d582ce77e14
					__obf_0001854db6e58157 /= 0xff
					__obf_a96897f6a99fe47b[0] += int32(__obf_8748f698cbebbcf5) * __obf_e02b66492c64e775
					__obf_a96897f6a99fe47b[1] += int32(__obf_8748f698cbebbcf5) * __obf_641b456d4f84424d
					__obf_a96897f6a99fe47b[2] += int32(__obf_8748f698cbebbcf5) * __obf_0001854db6e58157
					__obf_a96897f6a99fe47b[3] += int32(__obf_8748f698cbebbcf5) * __obf_43312d582ce77e14
					__obf_ab00e4b1393ed7cc += int32(__obf_8748f698cbebbcf5)
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*4
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[2] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+3] = __obf_f4cbeaeddfa5f8f1(__obf_a96897f6a99fe47b[3] / __obf_ab00e4b1393ed7cc)
		}
	}
}

func __obf_59ce29d524ed837c(__obf_c82b366489672bd7 *image.RGBA64, __obf_eab9867e780b63d0 *image.RGBA64, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int32, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_a96897f6a99fe47b [4]int64
			var __obf_ab00e4b1393ed7cc int64
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 8
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 8 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_a96897f6a99fe47b[0] += int64(__obf_8748f698cbebbcf5) * (int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])<<8 | int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1]))
					__obf_a96897f6a99fe47b[1] += int64(__obf_8748f698cbebbcf5) * (int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2])<<8 | int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+3]))
					__obf_a96897f6a99fe47b[2] += int64(__obf_8748f698cbebbcf5) * (int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+4])<<8 | int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+5]))
					__obf_a96897f6a99fe47b[3] += int64(__obf_8748f698cbebbcf5) * (int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+6])<<8 | int64(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+7]))
					__obf_ab00e4b1393ed7cc += int64(__obf_8748f698cbebbcf5)
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*8
			__obf_7ba5ac1ad9178cfe := __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+3] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[2] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+4] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+5] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[3] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+6] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+7] = uint8(__obf_7ba5ac1ad9178cfe)
		}
	}
}

func __obf_9fa2c4083eecedb9(__obf_c82b366489672bd7 *image.NRGBA64, __obf_eab9867e780b63d0 *image.RGBA64, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int32, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_a96897f6a99fe47b [4]int64
			var __obf_ab00e4b1393ed7cc int64
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 8
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 8 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_43312d582ce77e14 :=// Forward alpha-premultiplication
					int64(uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+6])<<8 | uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+7]))
					__obf_e02b66492c64e775 := int64(uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])<<8|uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1])) * __obf_43312d582ce77e14
					__obf_e02b66492c64e775 /= 0xffff
					__obf_641b456d4f84424d := int64(uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2])<<8|uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+3])) * __obf_43312d582ce77e14
					__obf_641b456d4f84424d /= 0xffff
					__obf_0001854db6e58157 := int64(uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+4])<<8|uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+5])) * __obf_43312d582ce77e14
					__obf_0001854db6e58157 /= 0xffff
					__obf_a96897f6a99fe47b[0] += int64(__obf_8748f698cbebbcf5) * __obf_e02b66492c64e775
					__obf_a96897f6a99fe47b[1] += int64(__obf_8748f698cbebbcf5) * __obf_641b456d4f84424d
					__obf_a96897f6a99fe47b[2] += int64(__obf_8748f698cbebbcf5) * __obf_0001854db6e58157
					__obf_a96897f6a99fe47b[3] += int64(__obf_8748f698cbebbcf5) * __obf_43312d582ce77e14
					__obf_ab00e4b1393ed7cc += int64(__obf_8748f698cbebbcf5)
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*8
			__obf_7ba5ac1ad9178cfe := __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+3] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[2] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+4] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+5] = uint8(__obf_7ba5ac1ad9178cfe)
			__obf_7ba5ac1ad9178cfe = __obf_f2dd0bd85be949f0(__obf_a96897f6a99fe47b[3] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+6] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+7] = uint8(__obf_7ba5ac1ad9178cfe)
		}
	}
}

func __obf_a5d4dad1cf1a65b6(__obf_c82b366489672bd7 *image.Gray, __obf_eab9867e780b63d0 *image.Gray, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int16, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[(__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_6479165cca1f077b int32
			var __obf_ab00e4b1393ed7cc int32
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case __obf_b7b3ec35eb2875ee < 0:
						__obf_b7b3ec35eb2875ee = 0
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = __obf_35b1276e2b34e1d8
					}
					__obf_6479165cca1f077b += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee])
					__obf_ab00e4b1393ed7cc += int32(__obf_8748f698cbebbcf5)
				}
			}
			__obf_c22f1458e7509d32 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3 - __obf_7cf35fc28dea4bbb.Min.X)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32] = __obf_f4cbeaeddfa5f8f1(__obf_6479165cca1f077b / __obf_ab00e4b1393ed7cc)
		}
	}
}

func __obf_1ff451e772d836c2(__obf_c82b366489672bd7 *image.Gray16, __obf_eab9867e780b63d0 *image.Gray16, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int32, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_6479165cca1f077b int64
			var __obf_ab00e4b1393ed7cc int64
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 2
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 2 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_6479165cca1f077b += int64(__obf_8748f698cbebbcf5) * int64(uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])<<8|uint16(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1]))
					__obf_ab00e4b1393ed7cc += int64(__obf_8748f698cbebbcf5)
				}
			}
			__obf_c22f1458e7509d32 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*2
			__obf_7ba5ac1ad9178cfe := __obf_f2dd0bd85be949f0(__obf_6479165cca1f077b / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+0] = uint8(__obf_7ba5ac1ad9178cfe >> 8)
			__obf_eab9867e780b63d0.
				Pix[__obf_c22f1458e7509d32+1] = uint8(__obf_7ba5ac1ad9178cfe)
		}
	}
}

func __obf_4b9ffd27224f812a(__obf_c82b366489672bd7 *__obf_e4d16584f024f650, __obf_eab9867e780b63d0 *__obf_e4d16584f024f650, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []int16, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_8e9fad7c873cabdb [3]int32
			var __obf_ab00e4b1393ed7cc int32
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				__obf_8748f698cbebbcf5 := __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9]
				if __obf_8748f698cbebbcf5 != 0 {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 3
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 3 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_8e9fad7c873cabdb[0] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])
					__obf_8e9fad7c873cabdb[1] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1])
					__obf_8e9fad7c873cabdb[2] += int32(__obf_8748f698cbebbcf5) * int32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2])
					__obf_ab00e4b1393ed7cc += int32(__obf_8748f698cbebbcf5)
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*3
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = __obf_f4cbeaeddfa5f8f1(__obf_8e9fad7c873cabdb[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = __obf_f4cbeaeddfa5f8f1(__obf_8e9fad7c873cabdb[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = __obf_f4cbeaeddfa5f8f1(__obf_8e9fad7c873cabdb[2] / __obf_ab00e4b1393ed7cc)
		}
	}
}

func __obf_a7249479cb109931(__obf_c82b366489672bd7 *__obf_e4d16584f024f650, __obf_eab9867e780b63d0 *__obf_e4d16584f024f650, __obf_dc6826d61e7dccec float64, __obf_71732cf03a2f4bc1 []bool, __obf_c22f1458e7509d32 []int, __obf_b34b7524c4bfcd41 int) {
	__obf_7cf35fc28dea4bbb := __obf_eab9867e780b63d0.Bounds()
	__obf_35b1276e2b34e1d8 := __obf_c82b366489672bd7.Bounds().Dx() - 1

	for __obf_9ab3da411fe8abc3 := __obf_7cf35fc28dea4bbb.Min.X; __obf_9ab3da411fe8abc3 < __obf_7cf35fc28dea4bbb.Max.X; __obf_9ab3da411fe8abc3++ {
		__obf_59a15792fe631198 := __obf_c82b366489672bd7.Pix[__obf_9ab3da411fe8abc3*__obf_c82b366489672bd7.Stride:]
		for __obf_b876c8539198a4e8 := __obf_7cf35fc28dea4bbb.Min.Y; __obf_b876c8539198a4e8 < __obf_7cf35fc28dea4bbb.Max.Y; __obf_b876c8539198a4e8++ {
			var __obf_8e9fad7c873cabdb [3]float32
			var __obf_ab00e4b1393ed7cc float32
			__obf_3563c6352e74b67e := __obf_c22f1458e7509d32[__obf_b876c8539198a4e8]
			__obf_acde913b01c6d744 := __obf_b876c8539198a4e8 * __obf_b34b7524c4bfcd41
			for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
				if __obf_71732cf03a2f4bc1[__obf_acde913b01c6d744+__obf_e39e622bfcb683d9] {
					__obf_b7b3ec35eb2875ee := __obf_3563c6352e74b67e + __obf_e39e622bfcb683d9
					switch {
					case uint(__obf_b7b3ec35eb2875ee) < uint(__obf_35b1276e2b34e1d8):
						__obf_b7b3ec35eb2875ee *= 3
					case __obf_b7b3ec35eb2875ee >= __obf_35b1276e2b34e1d8:
						__obf_b7b3ec35eb2875ee = 3 * __obf_35b1276e2b34e1d8
					default:
						__obf_b7b3ec35eb2875ee = 0
					}
					__obf_8e9fad7c873cabdb[0] += float32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+0])
					__obf_8e9fad7c873cabdb[1] += float32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+1])
					__obf_8e9fad7c873cabdb[2] += float32(__obf_59a15792fe631198[__obf_b7b3ec35eb2875ee+2])
					__obf_ab00e4b1393ed7cc++
				}
			}
			__obf_a6b1ad7bca234c57 := (__obf_b876c8539198a4e8-__obf_7cf35fc28dea4bbb.Min.Y)*__obf_eab9867e780b63d0.Stride + (__obf_9ab3da411fe8abc3-__obf_7cf35fc28dea4bbb.Min.X)*3
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+0] = __obf_64b96f3e506c4edd(__obf_8e9fad7c873cabdb[0] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+1] = __obf_64b96f3e506c4edd(__obf_8e9fad7c873cabdb[1] / __obf_ab00e4b1393ed7cc)
			__obf_eab9867e780b63d0.
				Pix[__obf_a6b1ad7bca234c57+2] = __obf_64b96f3e506c4edd(__obf_8e9fad7c873cabdb[2] / __obf_ab00e4b1393ed7cc)
		}
	}
}
