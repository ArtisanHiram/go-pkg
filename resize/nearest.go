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

package __obf_ac510735d4d13cdd

import "image"

func __obf_a16106cab6dfb5eb(__obf_21a3146223eb514f float32) uint8 {
	// Nearest-neighbor values are always
	// positive no need to check lower-bound.
	if __obf_21a3146223eb514f > 0xfe {
		return 0xff
	}
	return uint8(__obf_21a3146223eb514f)
}

func __obf_869edae50e32a65e(__obf_21a3146223eb514f float32) uint16 {
	if __obf_21a3146223eb514f > 0xfffe {
		return 0xffff
	}
	return uint16(__obf_21a3146223eb514f)
}

func __obf_d4c3da6972dfb241(__obf_3ccd19e51488be10 image.Image, __obf_37798f66dad5531e *image.RGBA64, __obf_6f7dfba2f58bb527 float64, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := range __obf_052b80a7daaa449d {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case __obf_235133c4493b6fdf < 0:
						__obf_235133c4493b6fdf = 0
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = __obf_859e5632b707c850
					}
					__obf_b4b27f4d455c3be9, __obf_bad790100516d51e, __obf_56457bda45cd368d, __obf_ea0c3716d96a2363 := __obf_3ccd19e51488be10.At(__obf_235133c4493b6fdf+__obf_3ccd19e51488be10.Bounds().Min.X, __obf_21a3146223eb514f+__obf_3ccd19e51488be10.Bounds().Min.Y).RGBA()
					__obf_cc5bf59ec3d2d89c[0] += float32(__obf_b4b27f4d455c3be9)
					__obf_cc5bf59ec3d2d89c[1] += float32(__obf_bad790100516d51e)
					__obf_cc5bf59ec3d2d89c[2] += float32(__obf_56457bda45cd368d)
					__obf_cc5bf59ec3d2d89c[3] += float32(__obf_ea0c3716d96a2363)
					__obf_792cfcb46f551313++
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_9dfa6bb18b3e659e(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.RGBA, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 4
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 4 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])
					__obf_cc5bf59ec3d2d89c[1] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])
					__obf_cc5bf59ec3d2d89c[2] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])
					__obf_cc5bf59ec3d2d89c[3] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3])
					__obf_792cfcb46f551313++
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*4
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
		}
	}
}

func __obf_d3929eb1ff5d9a81(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.NRGBA, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 4
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 4 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])
					__obf_cc5bf59ec3d2d89c[1] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1])
					__obf_cc5bf59ec3d2d89c[2] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])
					__obf_cc5bf59ec3d2d89c[3] += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3])
					__obf_792cfcb46f551313++
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*4
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = __obf_a16106cab6dfb5eb(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
		}
	}
}

func __obf_e7f8b39bc0d4199a(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.RGBA64, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := range __obf_052b80a7daaa449d {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 8
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 8 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]))
					__obf_cc5bf59ec3d2d89c[1] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3]))
					__obf_cc5bf59ec3d2d89c[2] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+4])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+5]))
					__obf_cc5bf59ec3d2d89c[3] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+6])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+7]))
					__obf_792cfcb46f551313++
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_2cf1e988bdf22f40(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.NRGBA64, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_cc5bf59ec3d2d89c [4]float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := range __obf_052b80a7daaa449d {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 8
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 8 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_cc5bf59ec3d2d89c[0] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]))
					__obf_cc5bf59ec3d2d89c[1] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+2])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+3]))
					__obf_cc5bf59ec3d2d89c[2] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+4])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+5]))
					__obf_cc5bf59ec3d2d89c[3] += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+6])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+7]))
					__obf_792cfcb46f551313++
				}
			}
			__obf_5b1a6e2dbf7ddf24 := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*8
			__obf_351a690ac24056c3 := __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[0] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+1] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[1] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+2] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+3] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[2] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+4] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+5] = uint8(__obf_351a690ac24056c3)
			__obf_351a690ac24056c3 = __obf_869edae50e32a65e(__obf_cc5bf59ec3d2d89c[3] / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+6] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_5b1a6e2dbf7ddf24+7] = uint8(__obf_351a690ac24056c3)
		}
	}
}

func __obf_220524a8a738e132(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.Gray, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_b0cbe3b481b571fe float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := range __obf_052b80a7daaa449d {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case __obf_235133c4493b6fdf < 0:
						__obf_235133c4493b6fdf = 0
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = __obf_859e5632b707c850
					}
					__obf_b0cbe3b481b571fe += float32(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf])
					__obf_792cfcb46f551313++
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f - __obf_91861ff72c438166.Min.X)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b] = __obf_a16106cab6dfb5eb(__obf_b0cbe3b481b571fe / __obf_792cfcb46f551313)
		}
	}
}

func __obf_b46a659aaf06ae03(__obf_3ccd19e51488be10, __obf_37798f66dad5531e *image.Gray16, __obf_46e0cb0caad4034a []bool, __obf_cf2c82255bf7ee2b []int, __obf_052b80a7daaa449d int) {
	__obf_91861ff72c438166 := __obf_37798f66dad5531e.Bounds()
	__obf_859e5632b707c850 := __obf_3ccd19e51488be10.Bounds().Dx() - 1

	for __obf_21a3146223eb514f := __obf_91861ff72c438166.Min.X; __obf_21a3146223eb514f < __obf_91861ff72c438166.Max.X; __obf_21a3146223eb514f++ {
		__obf_a1af34dea5cda2f0 := __obf_3ccd19e51488be10.Pix[__obf_21a3146223eb514f*__obf_3ccd19e51488be10.Stride:]
		for __obf_4811266080d7795e := __obf_91861ff72c438166.Min.Y; __obf_4811266080d7795e < __obf_91861ff72c438166.Max.Y; __obf_4811266080d7795e++ {
			var __obf_b0cbe3b481b571fe float32
			var __obf_792cfcb46f551313 float32
			__obf_807c674c91e487a0 := __obf_cf2c82255bf7ee2b[__obf_4811266080d7795e]
			__obf_5a46f1764400ebbd := __obf_4811266080d7795e * __obf_052b80a7daaa449d
			for __obf_7326873273fab46e := range __obf_052b80a7daaa449d {
				if __obf_46e0cb0caad4034a[__obf_5a46f1764400ebbd+__obf_7326873273fab46e] {
					__obf_235133c4493b6fdf := __obf_807c674c91e487a0 + __obf_7326873273fab46e
					switch {
					case uint(__obf_235133c4493b6fdf) < uint(__obf_859e5632b707c850):
						__obf_235133c4493b6fdf *= 2
					case __obf_235133c4493b6fdf >= __obf_859e5632b707c850:
						__obf_235133c4493b6fdf = 2 * __obf_859e5632b707c850
					default:
						__obf_235133c4493b6fdf = 0
					}
					__obf_b0cbe3b481b571fe += float32(uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+0])<<8 | uint16(__obf_a1af34dea5cda2f0[__obf_235133c4493b6fdf+1]))
					__obf_792cfcb46f551313++
				}
			}
			__obf_cf2c82255bf7ee2b := (__obf_4811266080d7795e-__obf_91861ff72c438166.Min.Y)*__obf_37798f66dad5531e.Stride + (__obf_21a3146223eb514f-__obf_91861ff72c438166.Min.X)*2
			__obf_351a690ac24056c3 := __obf_869edae50e32a65e(__obf_b0cbe3b481b571fe / __obf_792cfcb46f551313)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+0] = uint8(__obf_351a690ac24056c3 >> 8)
			__obf_37798f66dad5531e.
				Pix[__obf_cf2c82255bf7ee2b+1] = uint8(__obf_351a690ac24056c3)
		}
	}
}
