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

package __obf_3858dbfa2ccfdbe9

import "image"

func __obf_5f411479157375dc(__obf_15a5faef6b57be64 float32) uint8 {
	// Nearest-neighbor values are always
	// positive no need to check lower-bound.
	if __obf_15a5faef6b57be64 > 0xfe {
		return 0xff
	}
	return uint8(__obf_15a5faef6b57be64)
}

func __obf_5eaed3e82e09e743(__obf_15a5faef6b57be64 float32) uint16 {
	if __obf_15a5faef6b57be64 > 0xfffe {
		return 0xffff
	}
	return uint16(__obf_15a5faef6b57be64)
}

func __obf_9224993e7b70f1c7(__obf_8da81bf3c5df7f0d image.Image, __obf_c499a65912cac0c7 *image.RGBA64, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := range __obf_52dfb4be4832ebbd {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case __obf_0b5cf3f67563c51f < 0:
						__obf_0b5cf3f67563c51f = 0
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = __obf_3be4ed2a751d70c1
					}
					__obf_c3d2c2fc3dcef772, __obf_c8ad4e923dd632ca, __obf_19f96ad503f181fd, __obf_f6cabba4b7ec43fa := __obf_8da81bf3c5df7f0d.At(__obf_0b5cf3f67563c51f+__obf_8da81bf3c5df7f0d.Bounds().Min.X, __obf_15a5faef6b57be64+__obf_8da81bf3c5df7f0d.Bounds().Min.Y).RGBA()
					__obf_b7f694425525b116[0] += float32(__obf_c3d2c2fc3dcef772)
					__obf_b7f694425525b116[1] += float32(__obf_c8ad4e923dd632ca)
					__obf_b7f694425525b116[2] += float32(__obf_19f96ad503f181fd)
					__obf_b7f694425525b116[3] += float32(__obf_f6cabba4b7ec43fa)
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_5eaed3e82e09e743(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_1d1b25a5bcd2a73c(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.RGBA, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 4
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 4 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])
					__obf_b7f694425525b116[1] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])
					__obf_b7f694425525b116[2] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])
					__obf_b7f694425525b116[3] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3])
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*4
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_5f411479157375dc(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_5f411479157375dc(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_5f411479157375dc(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = __obf_5f411479157375dc(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_bedef5bbf1790511(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.NRGBA, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 4
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 4 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])
					__obf_b7f694425525b116[1] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])
					__obf_b7f694425525b116[2] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])
					__obf_b7f694425525b116[3] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3])
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*4
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_5f411479157375dc(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_5f411479157375dc(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_5f411479157375dc(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = __obf_5f411479157375dc(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_7a7f84fcfd267669(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.RGBA64, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := range __obf_52dfb4be4832ebbd {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 8
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 8 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]))
					__obf_b7f694425525b116[1] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3]))
					__obf_b7f694425525b116[2] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+4])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+5]))
					__obf_b7f694425525b116[3] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+6])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+7]))
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_5eaed3e82e09e743(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_5864fea3a63745f1(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.NRGBA64, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := range __obf_52dfb4be4832ebbd {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 8
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 8 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]))
					__obf_b7f694425525b116[1] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3]))
					__obf_b7f694425525b116[2] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+4])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+5]))
					__obf_b7f694425525b116[3] += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+6])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+7]))
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_5eaed3e82e09e743(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_5eaed3e82e09e743(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_c0cb06ed5d8c6356(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.Gray, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_42867cae7271e71c float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := range __obf_52dfb4be4832ebbd {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case __obf_0b5cf3f67563c51f < 0:
						__obf_0b5cf3f67563c51f = 0
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = __obf_3be4ed2a751d70c1
					}
					__obf_42867cae7271e71c += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f])
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64 - __obf_596992dcd1ca6b8d.Min.X)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8] = __obf_5f411479157375dc(__obf_42867cae7271e71c / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_9803aa82d6ad8fd8(__obf_8da81bf3c5df7f0d, __obf_c499a65912cac0c7 *image.Gray16, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_42867cae7271e71c float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := range __obf_52dfb4be4832ebbd {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 2
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 2 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_42867cae7271e71c += float32(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]))
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*2
			__obf_71de0dc978a1c220 := __obf_5eaed3e82e09e743(__obf_42867cae7271e71c / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+1] = uint8(__obf_71de0dc978a1c220)
		}
	}
}
