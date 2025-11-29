package __obf_bda21a78cb74016a

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_49fbaf091ac38b90 image.Rectangle, __obf_6acd5040e93d3dc1 bool) *NRGBA {
	__obf_625109fd824d6873 := image.NewNRGBA(__obf_49fbaf091ac38b90)
	for __obf_a4dd1bd05990217f := 0; __obf_a4dd1bd05990217f < __obf_49fbaf091ac38b90.Max.Y; __obf_a4dd1bd05990217f++ {
		for __obf_e22e7e43df48995e := 0; __obf_e22e7e43df48995e < __obf_49fbaf091ac38b90.Max.X; __obf_e22e7e43df48995e++ {
			if __obf_6acd5040e93d3dc1 {
				__obf_625109fd824d6873.
					Set(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, color.Alpha{A: uint8(0)})
			} else {
				__obf_625109fd824d6873.
					Set(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_625109fd824d6873}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_d94f6c073549a0ae *NRGBA) DrawImage(__obf_65fb97e6a605e786 Palette, __obf_fab13560644983fd *PositionRect, __obf_e3ca4bb26b15a246 *AreaRect) {
	__obf_86fb1adf8186600b := __obf_65fb97e6a605e786.Bounds().Max.X
	__obf_151c7628f33db4f3 := __obf_65fb97e6a605e786.Bounds().Max.Y
	__obf_b2d9011915b0a72c := __obf_fab13560644983fd.X
	__obf_6c108242961414c7 := __obf_fab13560644983fd.Y
	__obf_28899029d1b609a6 := __obf_fab13560644983fd.Height
	__obf_7ca0063bd82b512e := __obf_e3ca4bb26b15a246.MinX
	__obf_6225d6a1ae58a9d8 := __obf_e3ca4bb26b15a246.MinY
	__obf_016fe4d7ba1b9cec := __obf_e3ca4bb26b15a246.MaxX
	__obf_6f4f0d7ef7944ee1 := __obf_e3ca4bb26b15a246.MaxY

	for __obf_e22e7e43df48995e := range __obf_86fb1adf8186600b {
		for __obf_a4dd1bd05990217f := range __obf_151c7628f33db4f3 {
			__obf_e7796a53fdcf1f25 := __obf_65fb97e6a605e786.At(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f)
			if _, _, _, __obf_6ed92d728f70797c := __obf_e7796a53fdcf1f25.RGBA(); __obf_6ed92d728f70797c > 0 {
				if __obf_e22e7e43df48995e >= __obf_7ca0063bd82b512e && __obf_e22e7e43df48995e <= __obf_016fe4d7ba1b9cec && __obf_a4dd1bd05990217f >= __obf_6225d6a1ae58a9d8 && __obf_a4dd1bd05990217f <= __obf_6f4f0d7ef7944ee1 {
					__obf_d94f6c073549a0ae.
						Set(__obf_b2d9011915b0a72c+(__obf_e22e7e43df48995e-__obf_7ca0063bd82b512e), __obf_6c108242961414c7-__obf_28899029d1b609a6+(__obf_a4dd1bd05990217f-__obf_6225d6a1ae58a9d8), __obf_65fb97e6a605e786.At(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_d94f6c073549a0ae *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_86fb1adf8186600b := __obf_d94f6c073549a0ae.Bounds().Max.X
	__obf_151c7628f33db4f3 := __obf_d94f6c073549a0ae.Bounds().Max.Y
	__obf_cd41514d006d0679 := __obf_86fb1adf8186600b
	__obf_a8629efeda1a9547 := 0
	__obf_2d89945f425dcbd1 := __obf_151c7628f33db4f3
	__obf_3d1794aad27ceed1 := 0
	for __obf_e22e7e43df48995e := range __obf_86fb1adf8186600b {
		for __obf_a4dd1bd05990217f := range __obf_151c7628f33db4f3 {
			__obf_e7796a53fdcf1f25 := __obf_d94f6c073549a0ae.At(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f)
			if _, _, _, __obf_6ed92d728f70797c := __obf_e7796a53fdcf1f25.RGBA(); __obf_6ed92d728f70797c > 0 {
				if __obf_e22e7e43df48995e < __obf_cd41514d006d0679 {
					__obf_cd41514d006d0679 = __obf_e22e7e43df48995e
				}
				if __obf_e22e7e43df48995e > __obf_a8629efeda1a9547 {
					__obf_a8629efeda1a9547 = __obf_e22e7e43df48995e
				}

				if __obf_a4dd1bd05990217f < __obf_2d89945f425dcbd1 {
					__obf_2d89945f425dcbd1 = __obf_a4dd1bd05990217f
				}
				if __obf_a4dd1bd05990217f > __obf_3d1794aad27ceed1 {
					__obf_3d1794aad27ceed1 = __obf_a4dd1bd05990217f
				}
			}
		}
	}
	__obf_cd41514d006d0679 = int(max(0, float64(__obf_cd41514d006d0679-2)))
	__obf_a8629efeda1a9547 = int(min(float64(__obf_86fb1adf8186600b), float64(__obf_a8629efeda1a9547+2)))
	__obf_2d89945f425dcbd1 = int(max(0, float64(__obf_2d89945f425dcbd1-2)))
	__obf_3d1794aad27ceed1 = int(min(float64(__obf_151c7628f33db4f3), float64(__obf_3d1794aad27ceed1+2)))

	return &AreaRect{__obf_cd41514d006d0679, __obf_a8629efeda1a9547,

		// Rotate rotates the canvas by any angle
		__obf_2d89945f425dcbd1, __obf_3d1794aad27ceed1,
	}
}

func (__obf_d94f6c073549a0ae *NRGBA) Rotate(__obf_6ed92d728f70797c int, __obf_70461785f2def14c bool) {
	if __obf_6ed92d728f70797c == 0 {
		return
	}
	__obf_1ed21c16a991a4c4 := float64(__obf_6ed92d728f70797c) * math.Pi / 180
	__obf_15746c634842ee1c := __obf_d94f6c073549a0ae.Bounds().Dx()
	__obf_50377fb510be7ac5 := __obf_d94f6c073549a0ae.Bounds().Dy()
	__obf_6f8e99209c36dd76, __obf_b02fcc956b08e1ae := util.RotatedSize(__obf_15746c634842ee1c, __obf_50377fb510be7ac5, float64(__obf_6ed92d728f70797c))
	__obf_65fb97e6a605e786 := image.NewNRGBA(image.Rect(0, 0, __obf_6f8e99209c36dd76, __obf_b02fcc956b08e1ae))
	__obf_31f53a46247fe7e2 := float64(__obf_6f8e99209c36dd76) / 2
	__obf_8552415cafbd3ff0 := float64(__obf_b02fcc956b08e1ae) / 2
	__obf_120d0ca47292ade5 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_120d0ca47292ade5 = __obf_120d0ca47292ade5.Translate(__obf_31f53a46247fe7e2, __obf_8552415cafbd3ff0)
	__obf_120d0ca47292ade5 = __obf_120d0ca47292ade5.Rotate(__obf_1ed21c16a991a4c4)
	__obf_120d0ca47292ade5 = __obf_120d0ca47292ade5.Translate(-__obf_31f53a46247fe7e2, -__obf_8552415cafbd3ff0)
	__obf_e22e7e43df48995e := (__obf_6f8e99209c36dd76 - __obf_d94f6c073549a0ae.Bounds().Dx()) / 2
	__obf_a4dd1bd05990217f := (__obf_b02fcc956b08e1ae - __obf_d94f6c073549a0ae.Bounds().Dy()) / 2
	__obf_35ef85fdb9a66115, __obf_d069bdbced4ea4a4 := float64(__obf_e22e7e43df48995e), float64(__obf_a4dd1bd05990217f)
	__obf_a46e16fd1d616085 := __obf_120d0ca47292ade5.Translate(__obf_35ef85fdb9a66115, __obf_d069bdbced4ea4a4)
	__obf_834837877601fc9c := f64.Aff3{__obf_a46e16fd1d616085.XX, __obf_a46e16fd1d616085.XY, __obf_a46e16fd1d616085.X0, __obf_a46e16fd1d616085.YX, __obf_a46e16fd1d616085.YY, __obf_a46e16fd1d616085.Y0}

	draw.BiLinear.Transform(__obf_65fb97e6a605e786, __obf_834837877601fc9c, __obf_d94f6c073549a0ae, __obf_d94f6c073549a0ae.Bounds(), draw.Over, nil)
	__obf_d94f6c073549a0ae.
		NRGBA = __obf_65fb97e6a605e786

	if __obf_70461785f2def14c {
		__obf_76374c29fdf6ed54 := __obf_6f8e99209c36dd76 - __obf_15746c634842ee1c
		__obf_148ee71bbd9c7be2 := __obf_b02fcc956b08e1ae - __obf_50377fb510be7ac5
		__obf_2d0838a79d72d9c2 := (__obf_76374c29fdf6ed54 / 2) + 1
		__obf_68c9cbdb41f80cb3 := (__obf_148ee71bbd9c7be2 / 2) + 1
		__obf_d94f6c073549a0ae.
			SubImage(image.Rect(__obf_2d0838a79d72d9c2, __obf_68c9cbdb41f80cb3,

				// CropCircle crops a circular area
				__obf_15746c634842ee1c+__obf_2d0838a79d72d9c2, __obf_50377fb510be7ac5+__obf_68c9cbdb41f80cb3))
	}
}

func (__obf_d94f6c073549a0ae *NRGBA) CropCircle(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_f0b1976bf7767947 int) {
	__obf_8b05c243f860fe12 := __obf_d94f6c073549a0ae.Bounds()
	__obf_502b91d9d4397210 := image.NewNRGBA(__obf_8b05c243f860fe12)
	for __obf_ba58cd7e91b5411b := __obf_8b05c243f860fe12.Min.Y; __obf_ba58cd7e91b5411b < __obf_8b05c243f860fe12.Max.Y; __obf_ba58cd7e91b5411b++ {
		for __obf_61be968c0d023547 := __obf_8b05c243f860fe12.Min.X; __obf_61be968c0d023547 < __obf_8b05c243f860fe12.Max.X; __obf_61be968c0d023547++ {
			__obf_3402f1d2fcb063b2 := math.Hypot(float64(__obf_61be968c0d023547-__obf_e22e7e43df48995e), float64(__obf_ba58cd7e91b5411b-__obf_a4dd1bd05990217f))
			if __obf_3402f1d2fcb063b2 <= float64(__obf_f0b1976bf7767947) {
				__obf_502b91d9d4397210.
					Set(__obf_61be968c0d023547, __obf_ba58cd7e91b5411b, color.White)
			} else {
				__obf_502b91d9d4397210.
					Set(__obf_61be968c0d023547, __obf_ba58cd7e91b5411b, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_502b91d9d4397210, __obf_502b91d9d4397210.Bounds(), __obf_d94f6c073549a0ae, image.Point{X: 0, Y: 0}, __obf_502b91d9d4397210, image.Point{}, draw.Over)
	__obf_d94f6c073549a0ae.
		NRGBA = __obf_502b91d9d4397210
}

// CropScaleCircle scales and crops a circular area
func (__obf_d94f6c073549a0ae *NRGBA) CropScaleCircle(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_f0b1976bf7767947 int, __obf_94d5cd53464d75c6 int) {
	__obf_8b05c243f860fe12 := __obf_d94f6c073549a0ae.Bounds()
	__obf_502b91d9d4397210 := image.NewNRGBA(__obf_8b05c243f860fe12)

	for __obf_ba58cd7e91b5411b := __obf_8b05c243f860fe12.Min.Y; __obf_ba58cd7e91b5411b < __obf_8b05c243f860fe12.Max.Y; __obf_ba58cd7e91b5411b++ {
		for __obf_61be968c0d023547 := __obf_8b05c243f860fe12.Min.X; __obf_61be968c0d023547 < __obf_8b05c243f860fe12.Max.X; __obf_61be968c0d023547++ {
			__obf_3402f1d2fcb063b2 := math.Hypot(float64(__obf_61be968c0d023547-__obf_e22e7e43df48995e), float64(__obf_ba58cd7e91b5411b-__obf_a4dd1bd05990217f))
			if __obf_3402f1d2fcb063b2 <= float64(__obf_f0b1976bf7767947) {
				__obf_502b91d9d4397210.
					Set(__obf_61be968c0d023547, __obf_ba58cd7e91b5411b, color.White)
			} else {
				__obf_502b91d9d4397210.
					Set(__obf_61be968c0d023547, __obf_ba58cd7e91b5411b, color.Transparent)
			}
		}
	}

	if __obf_94d5cd53464d75c6 > 0 {
		__obf_1273b3135d29d24d := __obf_94d5cd53464d75c6 * 2
		__obf_44f32130c3ed88b1 := image.NewNRGBA(image.Rect(0, 0, __obf_d94f6c073549a0ae.Bounds().Dx()-__obf_1273b3135d29d24d, __obf_d94f6c073549a0ae.Bounds().Dy()-__obf_1273b3135d29d24d))
		draw.BiLinear.Scale(__obf_44f32130c3ed88b1, __obf_44f32130c3ed88b1.Bounds(), __obf_502b91d9d4397210, __obf_502b91d9d4397210.Bounds(), draw.Over, nil)
		__obf_502b91d9d4397210 = __obf_44f32130c3ed88b1
	}

	draw.DrawMask(__obf_502b91d9d4397210, __obf_502b91d9d4397210.Bounds(), __obf_d94f6c073549a0ae, image.Point{X: __obf_94d5cd53464d75c6, Y: __obf_94d5cd53464d75c6}, __obf_502b91d9d4397210, image.Point{}, draw.Over)
	__obf_d94f6c073549a0ae.
		NRGBA = __obf_502b91d9d4397210
}

// Scale scales the canvas
func (__obf_d94f6c073549a0ae *NRGBA) Scale(__obf_94d5cd53464d75c6 int, __obf_89fdbdbc95977aec, __obf_6436dda4e8616050 bool) {
	__obf_65fb97e6a605e786 := __obf_d94f6c073549a0ae.NRGBA
	if __obf_94d5cd53464d75c6 > 0 {
		__obf_1273b3135d29d24d := __obf_94d5cd53464d75c6 * 2
		__obf_632869d2455df647 := __obf_d94f6c073549a0ae.Bounds().Dx() - __obf_1273b3135d29d24d
		__obf_7ca04051c83e8ca9 := __obf_d94f6c073549a0ae.Bounds().Dy() - __obf_1273b3135d29d24d
		__obf_2462ade647f83408 := image.NewNRGBA(image.Rect(0, 0, __obf_632869d2455df647, __obf_7ca04051c83e8ca9))

		if !__obf_89fdbdbc95977aec {
			draw.BiLinear.Scale(__obf_2462ade647f83408, __obf_2462ade647f83408.Bounds(), __obf_d94f6c073549a0ae, __obf_d94f6c073549a0ae.Bounds(), draw.Over, nil)
		} else {
			__obf_53ea589f0311d0ac := util.CalcResizedRect(__obf_d94f6c073549a0ae.Bounds(), __obf_632869d2455df647, __obf_7ca04051c83e8ca9, __obf_6436dda4e8616050)
			draw.ApproxBiLinear.Scale(__obf_2462ade647f83408, __obf_53ea589f0311d0ac.Bounds(), __obf_d94f6c073549a0ae, __obf_d94f6c073549a0ae.Bounds(), draw.Over, nil)
		}
		__obf_65fb97e6a605e786 = __obf_2462ade647f83408
	}
	__obf_d94f6c073549a0ae.
		NRGBA = __obf_65fb97e6a605e786
}
