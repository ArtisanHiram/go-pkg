package __obf_244ef50d49151021

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_662a4428339c8e5b image.Rectangle, __obf_1e7042c2817d5eba bool) *NRGBA {
	__obf_78f6078a984d78ab := image.NewNRGBA(__obf_662a4428339c8e5b)
	for __obf_8cbd87ef15403513 := 0; __obf_8cbd87ef15403513 < __obf_662a4428339c8e5b.Max.Y; __obf_8cbd87ef15403513++ {
		for __obf_e487b3618d36e102 := 0; __obf_e487b3618d36e102 < __obf_662a4428339c8e5b.Max.X; __obf_e487b3618d36e102++ {
			if __obf_1e7042c2817d5eba {
				__obf_78f6078a984d78ab.
					Set(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, color.Alpha{A: uint8(0)})
			} else {
				__obf_78f6078a984d78ab.
					Set(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_78f6078a984d78ab}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_623e64214a71cd35 *NRGBA) DrawImage(__obf_4b8ec8035512b29f Palette, __obf_877d57dbfccc40cb *PositionRect, __obf_3c0fd844398c77e7 *AreaRect) {
	__obf_a7ac3c0c472a6787 := __obf_4b8ec8035512b29f.Bounds().Max.X
	__obf_e6eb35f8a1f9e292 := __obf_4b8ec8035512b29f.Bounds().Max.Y
	__obf_90a43b08cae9d210 := __obf_877d57dbfccc40cb.X
	__obf_838ab6c3fd178c64 := __obf_877d57dbfccc40cb.Y
	__obf_e76a75b7d41d29d6 := __obf_877d57dbfccc40cb.Height
	__obf_6caca97f55b0210f := __obf_3c0fd844398c77e7.MinX
	__obf_e21104c6aae1e8bc := __obf_3c0fd844398c77e7.MinY
	__obf_efbc28e8396d5130 := __obf_3c0fd844398c77e7.MaxX
	__obf_6509f1081234fc72 := __obf_3c0fd844398c77e7.MaxY

	for __obf_e487b3618d36e102 := range __obf_a7ac3c0c472a6787 {
		for __obf_8cbd87ef15403513 := range __obf_e6eb35f8a1f9e292 {
			__obf_15fd876741357eb2 := __obf_4b8ec8035512b29f.At(__obf_e487b3618d36e102, __obf_8cbd87ef15403513)
			if _, _, _, __obf_2b033369da82f821 := __obf_15fd876741357eb2.RGBA(); __obf_2b033369da82f821 > 0 {
				if __obf_e487b3618d36e102 >= __obf_6caca97f55b0210f && __obf_e487b3618d36e102 <= __obf_efbc28e8396d5130 && __obf_8cbd87ef15403513 >= __obf_e21104c6aae1e8bc && __obf_8cbd87ef15403513 <= __obf_6509f1081234fc72 {
					__obf_623e64214a71cd35.
						Set(__obf_90a43b08cae9d210+(__obf_e487b3618d36e102-__obf_6caca97f55b0210f), __obf_838ab6c3fd178c64-__obf_e76a75b7d41d29d6+(__obf_8cbd87ef15403513-__obf_e21104c6aae1e8bc), __obf_4b8ec8035512b29f.At(__obf_e487b3618d36e102, __obf_8cbd87ef15403513))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_623e64214a71cd35 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_a7ac3c0c472a6787 := __obf_623e64214a71cd35.Bounds().Max.X
	__obf_e6eb35f8a1f9e292 := __obf_623e64214a71cd35.Bounds().Max.Y
	__obf_4a5b8d8ec177e616 := __obf_a7ac3c0c472a6787
	__obf_37b7b15db85179cd := 0
	__obf_f5116c14efec20ac := __obf_e6eb35f8a1f9e292
	__obf_436aa8c9e2895e2a := 0
	for __obf_e487b3618d36e102 := range __obf_a7ac3c0c472a6787 {
		for __obf_8cbd87ef15403513 := range __obf_e6eb35f8a1f9e292 {
			__obf_15fd876741357eb2 := __obf_623e64214a71cd35.At(__obf_e487b3618d36e102, __obf_8cbd87ef15403513)
			if _, _, _, __obf_2b033369da82f821 := __obf_15fd876741357eb2.RGBA(); __obf_2b033369da82f821 > 0 {
				if __obf_e487b3618d36e102 < __obf_4a5b8d8ec177e616 {
					__obf_4a5b8d8ec177e616 = __obf_e487b3618d36e102
				}
				if __obf_e487b3618d36e102 > __obf_37b7b15db85179cd {
					__obf_37b7b15db85179cd = __obf_e487b3618d36e102
				}

				if __obf_8cbd87ef15403513 < __obf_f5116c14efec20ac {
					__obf_f5116c14efec20ac = __obf_8cbd87ef15403513
				}
				if __obf_8cbd87ef15403513 > __obf_436aa8c9e2895e2a {
					__obf_436aa8c9e2895e2a = __obf_8cbd87ef15403513
				}
			}
		}
	}
	__obf_4a5b8d8ec177e616 = int(max(0, float64(__obf_4a5b8d8ec177e616-2)))
	__obf_37b7b15db85179cd = int(min(float64(__obf_a7ac3c0c472a6787), float64(__obf_37b7b15db85179cd+2)))
	__obf_f5116c14efec20ac = int(max(0, float64(__obf_f5116c14efec20ac-2)))
	__obf_436aa8c9e2895e2a = int(min(float64(__obf_e6eb35f8a1f9e292), float64(__obf_436aa8c9e2895e2a+2)))

	return &AreaRect{__obf_4a5b8d8ec177e616, __obf_37b7b15db85179cd,

		// Rotate rotates the canvas by any angle
		__obf_f5116c14efec20ac, __obf_436aa8c9e2895e2a,
	}
}

func (__obf_623e64214a71cd35 *NRGBA) Rotate(__obf_2b033369da82f821 int, __obf_49cac67673103e20 bool) {
	if __obf_2b033369da82f821 == 0 {
		return
	}
	__obf_2b5072b762a61186 := float64(__obf_2b033369da82f821) * math.Pi / 180
	__obf_0a0ead4cc7165313 := __obf_623e64214a71cd35.Bounds().Dx()
	__obf_e79bd01a1212d97d := __obf_623e64214a71cd35.Bounds().Dy()
	__obf_b0bdb05fb01e5fd8, __obf_e0d45be06f8e29d9 := util.RotatedSize(__obf_0a0ead4cc7165313, __obf_e79bd01a1212d97d, float64(__obf_2b033369da82f821))
	__obf_4b8ec8035512b29f := image.NewNRGBA(image.Rect(0, 0, __obf_b0bdb05fb01e5fd8, __obf_e0d45be06f8e29d9))
	__obf_5d3733b3e3ab252b := float64(__obf_b0bdb05fb01e5fd8) / 2
	__obf_99d774eebaf67bfb := float64(__obf_e0d45be06f8e29d9) / 2
	__obf_42fe649b5f96a75f := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_42fe649b5f96a75f = __obf_42fe649b5f96a75f.Translate(__obf_5d3733b3e3ab252b, __obf_99d774eebaf67bfb)
	__obf_42fe649b5f96a75f = __obf_42fe649b5f96a75f.Rotate(__obf_2b5072b762a61186)
	__obf_42fe649b5f96a75f = __obf_42fe649b5f96a75f.Translate(-__obf_5d3733b3e3ab252b, -__obf_99d774eebaf67bfb)
	__obf_e487b3618d36e102 := (__obf_b0bdb05fb01e5fd8 - __obf_623e64214a71cd35.Bounds().Dx()) / 2
	__obf_8cbd87ef15403513 := (__obf_e0d45be06f8e29d9 - __obf_623e64214a71cd35.Bounds().Dy()) / 2
	__obf_d6e06ff51771d839, __obf_6e876aff5aaac50b := float64(__obf_e487b3618d36e102), float64(__obf_8cbd87ef15403513)
	__obf_d9209fb7830c5a01 := __obf_42fe649b5f96a75f.Translate(__obf_d6e06ff51771d839, __obf_6e876aff5aaac50b)
	__obf_fc88369c1393388e := f64.Aff3{__obf_d9209fb7830c5a01.XX, __obf_d9209fb7830c5a01.XY, __obf_d9209fb7830c5a01.X0, __obf_d9209fb7830c5a01.YX, __obf_d9209fb7830c5a01.YY, __obf_d9209fb7830c5a01.Y0}

	draw.BiLinear.Transform(__obf_4b8ec8035512b29f, __obf_fc88369c1393388e, __obf_623e64214a71cd35, __obf_623e64214a71cd35.Bounds(), draw.Over, nil)
	__obf_623e64214a71cd35.
		NRGBA = __obf_4b8ec8035512b29f

	if __obf_49cac67673103e20 {
		__obf_c05f3d139f17d1e5 := __obf_b0bdb05fb01e5fd8 - __obf_0a0ead4cc7165313
		__obf_7bbaeca861c502a7 := __obf_e0d45be06f8e29d9 - __obf_e79bd01a1212d97d
		__obf_ed508a8173048d4d := (__obf_c05f3d139f17d1e5 / 2) + 1
		__obf_c8ea7f83d06e4eb7 := (__obf_7bbaeca861c502a7 / 2) + 1
		__obf_623e64214a71cd35.
			SubImage(image.Rect(__obf_ed508a8173048d4d, __obf_c8ea7f83d06e4eb7,

				// CropCircle crops a circular area
				__obf_0a0ead4cc7165313+__obf_ed508a8173048d4d, __obf_e79bd01a1212d97d+__obf_c8ea7f83d06e4eb7))
	}
}

func (__obf_623e64214a71cd35 *NRGBA) CropCircle(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_2ff57e13a5c6715c int) {
	__obf_fa7ffec980b4842d := __obf_623e64214a71cd35.Bounds()
	__obf_7407cc102de3315e := image.NewNRGBA(__obf_fa7ffec980b4842d)
	for __obf_1ba6ff6b4903a822 := __obf_fa7ffec980b4842d.Min.Y; __obf_1ba6ff6b4903a822 < __obf_fa7ffec980b4842d.Max.Y; __obf_1ba6ff6b4903a822++ {
		for __obf_5bc987f242e0d607 := __obf_fa7ffec980b4842d.Min.X; __obf_5bc987f242e0d607 < __obf_fa7ffec980b4842d.Max.X; __obf_5bc987f242e0d607++ {
			__obf_2fb2ea25ae2a5368 := math.Hypot(float64(__obf_5bc987f242e0d607-__obf_e487b3618d36e102), float64(__obf_1ba6ff6b4903a822-__obf_8cbd87ef15403513))
			if __obf_2fb2ea25ae2a5368 <= float64(__obf_2ff57e13a5c6715c) {
				__obf_7407cc102de3315e.
					Set(__obf_5bc987f242e0d607, __obf_1ba6ff6b4903a822, color.White)
			} else {
				__obf_7407cc102de3315e.
					Set(__obf_5bc987f242e0d607, __obf_1ba6ff6b4903a822, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_7407cc102de3315e, __obf_7407cc102de3315e.Bounds(), __obf_623e64214a71cd35, image.Point{X: 0, Y: 0}, __obf_7407cc102de3315e, image.Point{}, draw.Over)
	__obf_623e64214a71cd35.
		NRGBA = __obf_7407cc102de3315e
}

// CropScaleCircle scales and crops a circular area
func (__obf_623e64214a71cd35 *NRGBA) CropScaleCircle(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_2ff57e13a5c6715c int, __obf_4a1b6a6612949cf8 int) {
	__obf_fa7ffec980b4842d := __obf_623e64214a71cd35.Bounds()
	__obf_7407cc102de3315e := image.NewNRGBA(__obf_fa7ffec980b4842d)

	for __obf_1ba6ff6b4903a822 := __obf_fa7ffec980b4842d.Min.Y; __obf_1ba6ff6b4903a822 < __obf_fa7ffec980b4842d.Max.Y; __obf_1ba6ff6b4903a822++ {
		for __obf_5bc987f242e0d607 := __obf_fa7ffec980b4842d.Min.X; __obf_5bc987f242e0d607 < __obf_fa7ffec980b4842d.Max.X; __obf_5bc987f242e0d607++ {
			__obf_2fb2ea25ae2a5368 := math.Hypot(float64(__obf_5bc987f242e0d607-__obf_e487b3618d36e102), float64(__obf_1ba6ff6b4903a822-__obf_8cbd87ef15403513))
			if __obf_2fb2ea25ae2a5368 <= float64(__obf_2ff57e13a5c6715c) {
				__obf_7407cc102de3315e.
					Set(__obf_5bc987f242e0d607, __obf_1ba6ff6b4903a822, color.White)
			} else {
				__obf_7407cc102de3315e.
					Set(__obf_5bc987f242e0d607, __obf_1ba6ff6b4903a822, color.Transparent)
			}
		}
	}

	if __obf_4a1b6a6612949cf8 > 0 {
		__obf_965b088f04ff9730 := __obf_4a1b6a6612949cf8 * 2
		__obf_846b2356ffb4a2db := image.NewNRGBA(image.Rect(0, 0, __obf_623e64214a71cd35.Bounds().Dx()-__obf_965b088f04ff9730, __obf_623e64214a71cd35.Bounds().Dy()-__obf_965b088f04ff9730))
		draw.BiLinear.Scale(__obf_846b2356ffb4a2db, __obf_846b2356ffb4a2db.Bounds(), __obf_7407cc102de3315e, __obf_7407cc102de3315e.Bounds(), draw.Over, nil)
		__obf_7407cc102de3315e = __obf_846b2356ffb4a2db
	}

	draw.DrawMask(__obf_7407cc102de3315e, __obf_7407cc102de3315e.Bounds(), __obf_623e64214a71cd35, image.Point{X: __obf_4a1b6a6612949cf8, Y: __obf_4a1b6a6612949cf8}, __obf_7407cc102de3315e, image.Point{}, draw.Over)
	__obf_623e64214a71cd35.
		NRGBA = __obf_7407cc102de3315e
}

// Scale scales the canvas
func (__obf_623e64214a71cd35 *NRGBA) Scale(__obf_4a1b6a6612949cf8 int, __obf_3a90f0b63d1c65b8, __obf_ee1057905e4c23c2 bool) {
	__obf_4b8ec8035512b29f := __obf_623e64214a71cd35.NRGBA
	if __obf_4a1b6a6612949cf8 > 0 {
		__obf_965b088f04ff9730 := __obf_4a1b6a6612949cf8 * 2
		__obf_617e3f73bf024321 := __obf_623e64214a71cd35.Bounds().Dx() - __obf_965b088f04ff9730
		__obf_20a4b6bf54eb22de := __obf_623e64214a71cd35.Bounds().Dy() - __obf_965b088f04ff9730
		__obf_b11be0f64d462d26 := image.NewNRGBA(image.Rect(0, 0, __obf_617e3f73bf024321, __obf_20a4b6bf54eb22de))

		if !__obf_3a90f0b63d1c65b8 {
			draw.BiLinear.Scale(__obf_b11be0f64d462d26, __obf_b11be0f64d462d26.Bounds(), __obf_623e64214a71cd35, __obf_623e64214a71cd35.Bounds(), draw.Over, nil)
		} else {
			__obf_566934613d62146b := util.CalcResizedRect(__obf_623e64214a71cd35.Bounds(), __obf_617e3f73bf024321, __obf_20a4b6bf54eb22de, __obf_ee1057905e4c23c2)
			draw.ApproxBiLinear.Scale(__obf_b11be0f64d462d26, __obf_566934613d62146b.Bounds(), __obf_623e64214a71cd35, __obf_623e64214a71cd35.Bounds(), draw.Over, nil)
		}
		__obf_4b8ec8035512b29f = __obf_b11be0f64d462d26
	}
	__obf_623e64214a71cd35.
		NRGBA = __obf_4b8ec8035512b29f
}
