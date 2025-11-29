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

package __obf_3858dbfa2ccfdbe9

import "image"

// Keep value in [0,255] range.
func __obf_cf865a084441317b(__obf_8da81bf3c5df7f0d int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_8da81bf3c5df7f0d) < 256 {
		return uint8(__obf_8da81bf3c5df7f0d)
	}
	if __obf_8da81bf3c5df7f0d > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_f02d250ba870435c(__obf_8da81bf3c5df7f0d int64) uint16 {
	if uint64(__obf_8da81bf3c5df7f0d) < 65536 {
		return uint16(__obf_8da81bf3c5df7f0d)
	}
	if __obf_8da81bf3c5df7f0d > 65535 {
		return 65535
	}
	return 0
}

func __obf_d42ce57ba503f1f8(__obf_8da81bf3c5df7f0d image.Image, __obf_c499a65912cac0c7 *image.RGBA64, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int32, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]int64
			var __obf_30d00e6256dd4ca7 int64
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case __obf_0b5cf3f67563c51f < 0:
						__obf_0b5cf3f67563c51f = 0
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = __obf_3be4ed2a751d70c1
					}
					__obf_c3d2c2fc3dcef772, __obf_c8ad4e923dd632ca, __obf_19f96ad503f181fd, __obf_f6cabba4b7ec43fa := __obf_8da81bf3c5df7f0d.At(__obf_0b5cf3f67563c51f+__obf_8da81bf3c5df7f0d.Bounds().Min.X, __obf_15a5faef6b57be64+__obf_8da81bf3c5df7f0d.Bounds().Min.Y).RGBA()
					__obf_b7f694425525b116[0] += int64(__obf_8e4ada777d217ea9) * int64(__obf_c3d2c2fc3dcef772)
					__obf_b7f694425525b116[1] += int64(__obf_8e4ada777d217ea9) * int64(__obf_c8ad4e923dd632ca)
					__obf_b7f694425525b116[2] += int64(__obf_8e4ada777d217ea9) * int64(__obf_19f96ad503f181fd)
					__obf_b7f694425525b116[3] += int64(__obf_8e4ada777d217ea9) * int64(__obf_f6cabba4b7ec43fa)
					__obf_30d00e6256dd4ca7 += int64(__obf_8e4ada777d217ea9)
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_f02d250ba870435c(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_9c2deddd470bf9c3(__obf_8da81bf3c5df7f0d *image.RGBA, __obf_c499a65912cac0c7 *image.RGBA, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int16, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]int32
			var __obf_30d00e6256dd4ca7 int32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 4
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 4 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])
					__obf_b7f694425525b116[1] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])
					__obf_b7f694425525b116[2] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])
					__obf_b7f694425525b116[3] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3])
					__obf_30d00e6256dd4ca7 += int32(__obf_8e4ada777d217ea9)
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*4
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_cf865a084441317b(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_cf865a084441317b(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_cf865a084441317b(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = __obf_cf865a084441317b(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_1ef8f3d1fec5a3c3(__obf_8da81bf3c5df7f0d *image.NRGBA, __obf_c499a65912cac0c7 *image.RGBA, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int16, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]int32
			var __obf_30d00e6256dd4ca7 int32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 4
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 4 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_f6cabba4b7ec43fa :=// Forward alpha-premultiplication
					int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3])
					__obf_c3d2c2fc3dcef772 := int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0]) * __obf_f6cabba4b7ec43fa
					__obf_c3d2c2fc3dcef772 /= 0xff
					__obf_c8ad4e923dd632ca := int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]) * __obf_f6cabba4b7ec43fa
					__obf_c8ad4e923dd632ca /= 0xff
					__obf_19f96ad503f181fd := int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2]) * __obf_f6cabba4b7ec43fa
					__obf_19f96ad503f181fd /= 0xff
					__obf_b7f694425525b116[0] += int32(__obf_8e4ada777d217ea9) * __obf_c3d2c2fc3dcef772
					__obf_b7f694425525b116[1] += int32(__obf_8e4ada777d217ea9) * __obf_c8ad4e923dd632ca
					__obf_b7f694425525b116[2] += int32(__obf_8e4ada777d217ea9) * __obf_19f96ad503f181fd
					__obf_b7f694425525b116[3] += int32(__obf_8e4ada777d217ea9) * __obf_f6cabba4b7ec43fa
					__obf_30d00e6256dd4ca7 += int32(__obf_8e4ada777d217ea9)
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*4
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_cf865a084441317b(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_cf865a084441317b(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_cf865a084441317b(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = __obf_cf865a084441317b(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_eb72b995d24df161(__obf_8da81bf3c5df7f0d *image.RGBA64, __obf_c499a65912cac0c7 *image.RGBA64, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int32, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]int64
			var __obf_30d00e6256dd4ca7 int64
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 8
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 8 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b7f694425525b116[0] += int64(__obf_8e4ada777d217ea9) * (int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8 | int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]))
					__obf_b7f694425525b116[1] += int64(__obf_8e4ada777d217ea9) * (int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])<<8 | int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3]))
					__obf_b7f694425525b116[2] += int64(__obf_8e4ada777d217ea9) * (int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+4])<<8 | int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+5]))
					__obf_b7f694425525b116[3] += int64(__obf_8e4ada777d217ea9) * (int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+6])<<8 | int64(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+7]))
					__obf_30d00e6256dd4ca7 += int64(__obf_8e4ada777d217ea9)
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_f02d250ba870435c(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_c1c846751d43d0f4(__obf_8da81bf3c5df7f0d *image.NRGBA64, __obf_c499a65912cac0c7 *image.RGBA64, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int32, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b7f694425525b116 [4]int64
			var __obf_30d00e6256dd4ca7 int64
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 8
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 8 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_f6cabba4b7ec43fa :=// Forward alpha-premultiplication
					int64(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+6])<<8 | uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+7]))
					__obf_c3d2c2fc3dcef772 := int64(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8|uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])) * __obf_f6cabba4b7ec43fa
					__obf_c3d2c2fc3dcef772 /= 0xffff
					__obf_c8ad4e923dd632ca := int64(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])<<8|uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+3])) * __obf_f6cabba4b7ec43fa
					__obf_c8ad4e923dd632ca /= 0xffff
					__obf_19f96ad503f181fd := int64(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+4])<<8|uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+5])) * __obf_f6cabba4b7ec43fa
					__obf_19f96ad503f181fd /= 0xffff
					__obf_b7f694425525b116[0] += int64(__obf_8e4ada777d217ea9) * __obf_c3d2c2fc3dcef772
					__obf_b7f694425525b116[1] += int64(__obf_8e4ada777d217ea9) * __obf_c8ad4e923dd632ca
					__obf_b7f694425525b116[2] += int64(__obf_8e4ada777d217ea9) * __obf_19f96ad503f181fd
					__obf_b7f694425525b116[3] += int64(__obf_8e4ada777d217ea9) * __obf_f6cabba4b7ec43fa
					__obf_30d00e6256dd4ca7 += int64(__obf_8e4ada777d217ea9)
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*8
			__obf_71de0dc978a1c220 := __obf_f02d250ba870435c(__obf_b7f694425525b116[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+3] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[2] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+4] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+5] = uint8(__obf_71de0dc978a1c220)
			__obf_71de0dc978a1c220 = __obf_f02d250ba870435c(__obf_b7f694425525b116[3] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+6] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+7] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_44079ea63696618c(__obf_8da81bf3c5df7f0d *image.Gray, __obf_c499a65912cac0c7 *image.Gray, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int16, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[(__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_42867cae7271e71c int32
			var __obf_30d00e6256dd4ca7 int32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case __obf_0b5cf3f67563c51f < 0:
						__obf_0b5cf3f67563c51f = 0
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = __obf_3be4ed2a751d70c1
					}
					__obf_42867cae7271e71c += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f])
					__obf_30d00e6256dd4ca7 += int32(__obf_8e4ada777d217ea9)
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64 - __obf_596992dcd1ca6b8d.Min.X)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8] = __obf_cf865a084441317b(__obf_42867cae7271e71c / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_010f5d83554658c7(__obf_8da81bf3c5df7f0d *image.Gray16, __obf_c499a65912cac0c7 *image.Gray16, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int32, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_42867cae7271e71c int64
			var __obf_30d00e6256dd4ca7 int64
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 2
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 2 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_42867cae7271e71c += int64(__obf_8e4ada777d217ea9) * int64(uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])<<8|uint16(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1]))
					__obf_30d00e6256dd4ca7 += int64(__obf_8e4ada777d217ea9)
				}
			}
			__obf_c032dbb648d056a8 := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*2
			__obf_71de0dc978a1c220 := __obf_f02d250ba870435c(__obf_42867cae7271e71c / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+0] = uint8(__obf_71de0dc978a1c220 >> 8)
			__obf_c499a65912cac0c7.
				Pix[__obf_c032dbb648d056a8+1] = uint8(__obf_71de0dc978a1c220)
		}
	}
}

func __obf_ab9eaf73a4199f9c(__obf_8da81bf3c5df7f0d *__obf_1f271b19471a7add, __obf_c499a65912cac0c7 *__obf_1f271b19471a7add, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []int16, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b6ea2276e9a14426 [3]int32
			var __obf_30d00e6256dd4ca7 int32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				__obf_8e4ada777d217ea9 := __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f]
				if __obf_8e4ada777d217ea9 != 0 {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 3
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 3 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b6ea2276e9a14426[0] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])
					__obf_b6ea2276e9a14426[1] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])
					__obf_b6ea2276e9a14426[2] += int32(__obf_8e4ada777d217ea9) * int32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])
					__obf_30d00e6256dd4ca7 += int32(__obf_8e4ada777d217ea9)
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*3
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_cf865a084441317b(__obf_b6ea2276e9a14426[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_cf865a084441317b(__obf_b6ea2276e9a14426[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_cf865a084441317b(__obf_b6ea2276e9a14426[2] / __obf_30d00e6256dd4ca7)
		}
	}
}

func __obf_de165637ccf50364(__obf_8da81bf3c5df7f0d *__obf_1f271b19471a7add, __obf_c499a65912cac0c7 *__obf_1f271b19471a7add, __obf_0d0fd6056291c76b float64, __obf_3fabfc59cdeeb8b4 []bool, __obf_c032dbb648d056a8 []int, __obf_52dfb4be4832ebbd int) {
	__obf_596992dcd1ca6b8d := __obf_c499a65912cac0c7.Bounds()
	__obf_3be4ed2a751d70c1 := __obf_8da81bf3c5df7f0d.Bounds().Dx() - 1

	for __obf_15a5faef6b57be64 := __obf_596992dcd1ca6b8d.Min.X; __obf_15a5faef6b57be64 < __obf_596992dcd1ca6b8d.Max.X; __obf_15a5faef6b57be64++ {
		__obf_476731dea5e9746e := __obf_8da81bf3c5df7f0d.Pix[__obf_15a5faef6b57be64*__obf_8da81bf3c5df7f0d.Stride:]
		for __obf_523cb3ec667c6a54 := __obf_596992dcd1ca6b8d.Min.Y; __obf_523cb3ec667c6a54 < __obf_596992dcd1ca6b8d.Max.Y; __obf_523cb3ec667c6a54++ {
			var __obf_b6ea2276e9a14426 [3]float32
			var __obf_30d00e6256dd4ca7 float32
			__obf_5f2d417327929024 := __obf_c032dbb648d056a8[__obf_523cb3ec667c6a54]
			__obf_69d5eccaa41f8a07 := __obf_523cb3ec667c6a54 * __obf_52dfb4be4832ebbd
			for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
				if __obf_3fabfc59cdeeb8b4[__obf_69d5eccaa41f8a07+__obf_9d9681e93d5a567f] {
					__obf_0b5cf3f67563c51f := __obf_5f2d417327929024 + __obf_9d9681e93d5a567f
					switch {
					case uint(__obf_0b5cf3f67563c51f) < uint(__obf_3be4ed2a751d70c1):
						__obf_0b5cf3f67563c51f *= 3
					case __obf_0b5cf3f67563c51f >= __obf_3be4ed2a751d70c1:
						__obf_0b5cf3f67563c51f = 3 * __obf_3be4ed2a751d70c1
					default:
						__obf_0b5cf3f67563c51f = 0
					}
					__obf_b6ea2276e9a14426[0] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+0])
					__obf_b6ea2276e9a14426[1] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+1])
					__obf_b6ea2276e9a14426[2] += float32(__obf_476731dea5e9746e[__obf_0b5cf3f67563c51f+2])
					__obf_30d00e6256dd4ca7++
				}
			}
			__obf_a3c377f4f1e92f6e := (__obf_523cb3ec667c6a54-__obf_596992dcd1ca6b8d.Min.Y)*__obf_c499a65912cac0c7.Stride + (__obf_15a5faef6b57be64-__obf_596992dcd1ca6b8d.Min.X)*3
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+0] = __obf_5f411479157375dc(__obf_b6ea2276e9a14426[0] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+1] = __obf_5f411479157375dc(__obf_b6ea2276e9a14426[1] / __obf_30d00e6256dd4ca7)
			__obf_c499a65912cac0c7.
				Pix[__obf_a3c377f4f1e92f6e+2] = __obf_5f411479157375dc(__obf_b6ea2276e9a14426[2] / __obf_30d00e6256dd4ca7)
		}
	}
}
