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

package __obf_ac510735d4d13cdd

import "image"

// Keep value in [0,255] range.
func __obf_8ab15d4a34cce580(__obf_3ccd19e51488be10 int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_3ccd19e51488be10) < 256 {
		return uint8(__obf_3ccd19e51488be10)
	}
	if __obf_3ccd19e51488be10 > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_5fe4185c7d80a557(__obf_3ccd19e51488be10 int64) uint16 {
	if uint64(__obf_3ccd19e51488be10) < 65536 {
		return uint16(__obf_3ccd19e51488be10)
	}
	if __obf_3ccd19e51488be10 > 65535 {
		return 65535
	}
	return 0
}

func __obf_1e8dec540911dfe5(__obf_3ccd19e51488be10 image.Image, __obf_37798f66dad5531e *image.RGBA64, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int32, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]int64
			var __obf_792cfcb46f551313 int64
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case __obf_235133c4493b6fdf < 0:
						__obf_235133c4493b6fdf = 0
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = __obf_859e5632b707c850
					}
					__obf_b4b27f4d455c3be9, __obf_bad790100516d51e, __obf_56457bda45cd368d, __obf_ea0c3716d96a2363 := __obf_3ccd19e51488be10.At(__obf_235133c4493b6fdf+__obf_3ccd19e51488be10.Bounds().Min.X, __obf_21a3146223eb514f+__obf_3ccd19e51488be10.Bounds().Min.Y).RGBA()
					__obf_cc5bf59ec3d2d89c[0] += int64(__obf_2620dbebe2e7a7e3) * int64(__obf_b4b27f4d455c3be9)
					__obf_cc5bf59ec3d2d89c[1] += int64(__obf_2620dbebe2e7a7e3) * int64(__obf_bad790100516d51e)
					__obf_cc5bf59ec3d2d89c[2] += int64(__obf_2620dbebe2e7a7e3) * int64(__obf_56457bda45cd368d)
					__obf_cc5bf59ec3d2d89c[3] += int64(__obf_2620dbebe2e7a7e3) * int64(__obf_ea0c3716d96a2363)
					__obf_792cfcb46f551313 += int64(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_a81ef026a03961f8(__obf_3ccd19e51488be10 *image.RGBA, __obf_37798f66dad5531e *image.RGBA, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int16, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]int32
			var __obf_792cfcb46f551313 int32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 4
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 4 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])
					__obf_cc5bf59ec3d2d89c[1] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])
					__obf_cc5bf59ec3d2d89c[2] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])
					__obf_cc5bf59ec3d2d89c[3] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3])
					__obf_792cfcb46f551313 += int32(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*4
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
		}
	}
}

func __obf_e01bcd08df027a03(__obf_3ccd19e51488be10 *image.NRGBA, __obf_37798f66dad5531e *image.RGBA, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int16, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]int32
			var __obf_792cfcb46f551313 int32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 4
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 4 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_ea0c3716d96a2363 :=// Forward alpha-premultiplication
					int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3])
					__obf_b4b27f4d455c3be9 := int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0]) * __obf_ea0c3716d96a2363
					__obf_b4b27f4d455c3be9 /= 0xff
					__obf_bad790100516d51e := int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]) * __obf_ea0c3716d96a2363
					__obf_bad790100516d51e /= 0xff
					__obf_56457bda45cd368d := int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2]) * __obf_ea0c3716d96a2363
					__obf_56457bda45cd368d /= 0xff
					__obf_cc5bf59ec3d2d89c[0] += int32(__obf_2620dbebe2e7a7e3) * __obf_b4b27f4d455c3be9
					__obf_cc5bf59ec3d2d89c[1] += int32(__obf_2620dbebe2e7a7e3) * __obf_bad790100516d51e
					__obf_cc5bf59ec3d2d89c[2] += int32(__obf_2620dbebe2e7a7e3) * __obf_56457bda45cd368d
					__obf_cc5bf59ec3d2d89c[3] += int32(__obf_2620dbebe2e7a7e3) * __obf_ea0c3716d96a2363
					__obf_792cfcb46f551313 += int32(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*4
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = __obf_8ab15d4a34cce580(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
		}
	}
}

func __obf_fd9e808f1958058e(__obf_3ccd19e51488be10 *image.RGBA64, __obf_37798f66dad5531e *image.RGBA64, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int32, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]int64
			var __obf_792cfcb46f551313 int64
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 8
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 8 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += int64(__obf_2620dbebe2e7a7e3) * (int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8 | int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]))
					__obf_cc5bf59ec3d2d89c[1] += int64(__obf_2620dbebe2e7a7e3) * (int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])<<8 | int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3]))
					__obf_cc5bf59ec3d2d89c[2] += int64(__obf_2620dbebe2e7a7e3) * (int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+4])<<8 | int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+5]))
					__obf_cc5bf59ec3d2d89c[3] += int64(__obf_2620dbebe2e7a7e3) * (int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+6])<<8 | int64(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+7]))
					__obf_792cfcb46f551313 += int64(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_3d8976013120b419(__obf_3ccd19e51488be10 *image.NRGBA64, __obf_37798f66dad5531e *image.RGBA64, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int32, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]int64
			var __obf_792cfcb46f551313 int64
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 8
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 8 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_ea0c3716d96a2363 :=// Forward alpha-premultiplication
					int64(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+6])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+7]))
					__obf_b4b27f4d455c3be9 := int64(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8|uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])) * __obf_ea0c3716d96a2363
					__obf_b4b27f4d455c3be9 /= 0xffff
					__obf_bad790100516d51e := int64(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])<<8|uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3])) * __obf_ea0c3716d96a2363
					__obf_bad790100516d51e /= 0xffff
					__obf_56457bda45cd368d := int64(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+4])<<8|uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+5])) * __obf_ea0c3716d96a2363
					__obf_56457bda45cd368d /= 0xffff
					__obf_cc5bf59ec3d2d89c[0] += int64(__obf_2620dbebe2e7a7e3) * __obf_b4b27f4d455c3be9
					__obf_cc5bf59ec3d2d89c[1] += int64(__obf_2620dbebe2e7a7e3) * __obf_bad790100516d51e
					__obf_cc5bf59ec3d2d89c[2] += int64(__obf_2620dbebe2e7a7e3) * __obf_56457bda45cd368d
					__obf_cc5bf59ec3d2d89c[3] += int64(__obf_2620dbebe2e7a7e3) * __obf_ea0c3716d96a2363
					__obf_792cfcb46f551313 += int64(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_5fe4185c7d80a557(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_e7b82bf93376b1f1(__obf_3ccd19e51488be10 *image.Gray, __obf_37798f66dad5531e *image.Gray, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int16, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[(__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_b0cbe3b481b571fe int32
			var __obf_792cfcb46f551313 int32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case __obf_235133c4493b6fdf < 0:
						__obf_235133c4493b6fdf = 0
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = __obf_859e5632b707c850
					}
					__obf_b0cbe3b481b571fe += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf])
					__obf_792cfcb46f551313 += int32(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f - __obf_91861ff72c438166.Min.X)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b] = __obf_8ab15d4a34cce580(__obf_b0cbe3b481b571fe / __obf_792cfcb46f551313)
		}
	}
}

func __obf_88dc9a250ceb45c4(__obf_3ccd19e51488be10 *image.Gray16, __obf_37798f66dad5531e *image.Gray16, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int32, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_b0cbe3b481b571fe int64
			var __obf_792cfcb46f551313 int64
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 2
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 2 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_b0cbe3b481b571fe += int64(__obf_2620dbebe2e7a7e3) * int64(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8|uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]))
					__obf_792cfcb46f551313 += int64(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*2
			__obf_351a690ac24056c3 := __obf_5fe4185c7d80a557(__obf_b0cbe3b481b571fe / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+1] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_76c3d32722dd59f4(__obf_3ccd19e51488be10 *__obf_9474021cedba6d1a, __obf_37798f66dad5531e *__obf_9474021cedba6d1a, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []int16, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_5fd02db579f190b4 [3]int32
			var __obf_792cfcb46f551313 int32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				__obf_2620dbebe2e7a7e3 := __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e]
				if __obf_2620dbebe2e7a7e3 != 0 {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 3
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 3 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_5fd02db579f190b4[0] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])
					__obf_5fd02db579f190b4[1] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])
					__obf_5fd02db579f190b4[2] += int32(__obf_2620dbebe2e7a7e3) * int32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])
					__obf_792cfcb46f551313 += int32(__obf_2620dbebe2e7a7e3)
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*3
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_8ab15d4a34cce580(__obf_5fd02db579f190b4[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_8ab15d4a34cce580(__obf_5fd02db579f190b4[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_8ab15d4a34cce580(__obf_5fd02db579f190b4[2] / __obf_792cfcb46f551313)
		}
	}
}

func __obf_cb273764563e26fc(__obf_3ccd19e51488be10 *__obf_9474021cedba6d1a, __obf_37798f66dad5531e *__obf_9474021cedba6d1a, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_5fd02db579f190b4 [3]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 3
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 3 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_5fd02db579f190b4[0] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])
					__obf_5fd02db579f190b4[1] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])
					__obf_5fd02db579f190b4[2] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])
					__obf_792cfcb46f551313++
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*3
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_a16106cab6dfb5eb(__obf_5fd02db579f190b4[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_a16106cab6dfb5eb(__obf_5fd02db579f190b4[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_a16106cab6dfb5eb(__obf_5fd02db579f190b4[2] / __obf_792cfcb46f551313)
		}
	}
}
