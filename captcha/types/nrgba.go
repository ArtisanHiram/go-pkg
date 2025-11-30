package __obf_54406b1a1de84196

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_e73d974dafb71af7 image.Rectangle, __obf_51a128146961ca9b bool) *NRGBA {
	__obf_d1a18961c3dc0c48 := image.NewNRGBA(__obf_e73d974dafb71af7)
	for __obf_33a1f511b09ac2af := 0; __obf_33a1f511b09ac2af < __obf_e73d974dafb71af7.Max.Y; __obf_33a1f511b09ac2af++ {
		for __obf_48afb3005cd4a35c := 0; __obf_48afb3005cd4a35c < __obf_e73d974dafb71af7.Max.X; __obf_48afb3005cd4a35c++ {
			if __obf_51a128146961ca9b {
				__obf_d1a18961c3dc0c48.
					Set(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, color.Alpha{A: uint8(0)})
			} else {
				__obf_d1a18961c3dc0c48.
					Set(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_d1a18961c3dc0c48}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_cbeabe4ae9e49861 *NRGBA) DrawImage(__obf_10d21d6285312279 Palette, __obf_2aaedf4fbe252830 *PositionRect, __obf_a2015708abfe68b9 *AreaRect) {
	__obf_5883e9285a516ece := __obf_10d21d6285312279.Bounds().Max.X
	__obf_0f8e181d560d05a0 := __obf_10d21d6285312279.Bounds().Max.Y
	__obf_f15799c0ea07c056 := __obf_2aaedf4fbe252830.X
	__obf_7170273a7e80d213 := __obf_2aaedf4fbe252830.Y
	__obf_18a0732091ffa1cb := __obf_2aaedf4fbe252830.Height
	__obf_cb9e1ba19af05291 := __obf_a2015708abfe68b9.MinX
	__obf_11c91e88acee59b7 := __obf_a2015708abfe68b9.MinY
	__obf_bbeece3ba11da587 := __obf_a2015708abfe68b9.MaxX
	__obf_7b60088f78f03634 := __obf_a2015708abfe68b9.MaxY

	for __obf_48afb3005cd4a35c := range __obf_5883e9285a516ece {
		for __obf_33a1f511b09ac2af := range __obf_0f8e181d560d05a0 {
			__obf_bbad58a05ae9b2bc := __obf_10d21d6285312279.At(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af)
			if _, _, _, __obf_23448ffa2f9aadbc := __obf_bbad58a05ae9b2bc.RGBA(); __obf_23448ffa2f9aadbc > 0 {
				if __obf_48afb3005cd4a35c >= __obf_cb9e1ba19af05291 && __obf_48afb3005cd4a35c <= __obf_bbeece3ba11da587 && __obf_33a1f511b09ac2af >= __obf_11c91e88acee59b7 && __obf_33a1f511b09ac2af <= __obf_7b60088f78f03634 {
					__obf_cbeabe4ae9e49861.
						Set(__obf_f15799c0ea07c056+(__obf_48afb3005cd4a35c-__obf_cb9e1ba19af05291), __obf_7170273a7e80d213-__obf_18a0732091ffa1cb+(__obf_33a1f511b09ac2af-__obf_11c91e88acee59b7), __obf_10d21d6285312279.At(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_cbeabe4ae9e49861 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_5883e9285a516ece := __obf_cbeabe4ae9e49861.Bounds().Max.X
	__obf_0f8e181d560d05a0 := __obf_cbeabe4ae9e49861.Bounds().Max.Y
	__obf_907f96d0fe0868f4 := __obf_5883e9285a516ece
	__obf_f3f182f1c2fb0d54 := 0
	__obf_bf482b3ecc79f58c := __obf_0f8e181d560d05a0
	__obf_bcd72b5e7c0080de := 0
	for __obf_48afb3005cd4a35c := range __obf_5883e9285a516ece {
		for __obf_33a1f511b09ac2af := range __obf_0f8e181d560d05a0 {
			__obf_bbad58a05ae9b2bc := __obf_cbeabe4ae9e49861.At(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af)
			if _, _, _, __obf_23448ffa2f9aadbc := __obf_bbad58a05ae9b2bc.RGBA(); __obf_23448ffa2f9aadbc > 0 {
				if __obf_48afb3005cd4a35c < __obf_907f96d0fe0868f4 {
					__obf_907f96d0fe0868f4 = __obf_48afb3005cd4a35c
				}
				if __obf_48afb3005cd4a35c > __obf_f3f182f1c2fb0d54 {
					__obf_f3f182f1c2fb0d54 = __obf_48afb3005cd4a35c
				}

				if __obf_33a1f511b09ac2af < __obf_bf482b3ecc79f58c {
					__obf_bf482b3ecc79f58c = __obf_33a1f511b09ac2af
				}
				if __obf_33a1f511b09ac2af > __obf_bcd72b5e7c0080de {
					__obf_bcd72b5e7c0080de = __obf_33a1f511b09ac2af
				}
			}
		}
	}
	__obf_907f96d0fe0868f4 = int(max(0, float64(__obf_907f96d0fe0868f4-2)))
	__obf_f3f182f1c2fb0d54 = int(min(float64(__obf_5883e9285a516ece), float64(__obf_f3f182f1c2fb0d54+2)))
	__obf_bf482b3ecc79f58c = int(max(0, float64(__obf_bf482b3ecc79f58c-2)))
	__obf_bcd72b5e7c0080de = int(min(float64(__obf_0f8e181d560d05a0), float64(__obf_bcd72b5e7c0080de+2)))

	return &AreaRect{__obf_907f96d0fe0868f4, __obf_f3f182f1c2fb0d54,

		// Rotate rotates the canvas by any angle
		__obf_bf482b3ecc79f58c, __obf_bcd72b5e7c0080de,
	}
}

func (__obf_cbeabe4ae9e49861 *NRGBA) Rotate(__obf_23448ffa2f9aadbc int, __obf_2f0fe0f43907e24d bool) {
	if __obf_23448ffa2f9aadbc == 0 {
		return
	}
	__obf_27c01b4675c92ab0 := float64(__obf_23448ffa2f9aadbc) * math.Pi / 180
	__obf_ecc7d755268f10f0 := __obf_cbeabe4ae9e49861.Bounds().Dx()
	__obf_0a20c78b8422649a := __obf_cbeabe4ae9e49861.Bounds().Dy()
	__obf_2366a8ab85b2ae47, __obf_daaeaddeff0e74f7 := util.RotatedSize(__obf_ecc7d755268f10f0, __obf_0a20c78b8422649a, float64(__obf_23448ffa2f9aadbc))
	__obf_10d21d6285312279 := image.NewNRGBA(image.Rect(0, 0, __obf_2366a8ab85b2ae47, __obf_daaeaddeff0e74f7))
	__obf_27f743b51ed7c59a := float64(__obf_2366a8ab85b2ae47) / 2
	__obf_f7cf0a5616e6a051 := float64(__obf_daaeaddeff0e74f7) / 2
	__obf_1c5e4a64e4ba09fe := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_1c5e4a64e4ba09fe = __obf_1c5e4a64e4ba09fe.Translate(__obf_27f743b51ed7c59a, __obf_f7cf0a5616e6a051)
	__obf_1c5e4a64e4ba09fe = __obf_1c5e4a64e4ba09fe.Rotate(__obf_27c01b4675c92ab0)
	__obf_1c5e4a64e4ba09fe = __obf_1c5e4a64e4ba09fe.Translate(-__obf_27f743b51ed7c59a, -__obf_f7cf0a5616e6a051)
	__obf_48afb3005cd4a35c := (__obf_2366a8ab85b2ae47 - __obf_cbeabe4ae9e49861.Bounds().Dx()) / 2
	__obf_33a1f511b09ac2af := (__obf_daaeaddeff0e74f7 - __obf_cbeabe4ae9e49861.Bounds().Dy()) / 2
	__obf_0866c3b92f5f395f, __obf_6301216cf0360ba2 := float64(__obf_48afb3005cd4a35c), float64(__obf_33a1f511b09ac2af)
	__obf_93daf9680a0f770e := __obf_1c5e4a64e4ba09fe.Translate(__obf_0866c3b92f5f395f, __obf_6301216cf0360ba2)
	__obf_1b8b05889c05494a := f64.Aff3{__obf_93daf9680a0f770e.XX, __obf_93daf9680a0f770e.XY, __obf_93daf9680a0f770e.X0, __obf_93daf9680a0f770e.YX, __obf_93daf9680a0f770e.YY, __obf_93daf9680a0f770e.Y0}

	draw.BiLinear.Transform(__obf_10d21d6285312279, __obf_1b8b05889c05494a, __obf_cbeabe4ae9e49861, __obf_cbeabe4ae9e49861.Bounds(), draw.Over, nil)
	__obf_cbeabe4ae9e49861.
		NRGBA = __obf_10d21d6285312279

	if __obf_2f0fe0f43907e24d {
		__obf_1f26d2447b187389 := __obf_2366a8ab85b2ae47 - __obf_ecc7d755268f10f0
		__obf_05da0c2bfc50f0e2 := __obf_daaeaddeff0e74f7 - __obf_0a20c78b8422649a
		__obf_21594ae268a2eb66 := (__obf_1f26d2447b187389 / 2) + 1
		__obf_8c43a22828ae34a2 := (__obf_05da0c2bfc50f0e2 / 2) + 1
		__obf_cbeabe4ae9e49861.
			SubImage(image.Rect(__obf_21594ae268a2eb66, __obf_8c43a22828ae34a2,

				// CropCircle crops a circular area
				__obf_ecc7d755268f10f0+__obf_21594ae268a2eb66, __obf_0a20c78b8422649a+__obf_8c43a22828ae34a2))
	}
}

func (__obf_cbeabe4ae9e49861 *NRGBA) CropCircle(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_9dba054a785bce28 int) {
	__obf_0b934908b11a6f63 := __obf_cbeabe4ae9e49861.Bounds()
	__obf_81fd9df55ede76fb := image.NewNRGBA(__obf_0b934908b11a6f63)
	for __obf_9aa12d47f69296a8 := __obf_0b934908b11a6f63.Min.Y; __obf_9aa12d47f69296a8 < __obf_0b934908b11a6f63.Max.Y; __obf_9aa12d47f69296a8++ {
		for __obf_3e26d844e26658db := __obf_0b934908b11a6f63.Min.X; __obf_3e26d844e26658db < __obf_0b934908b11a6f63.Max.X; __obf_3e26d844e26658db++ {
			__obf_b95967467beb5e15 := math.Hypot(float64(__obf_3e26d844e26658db-__obf_48afb3005cd4a35c), float64(__obf_9aa12d47f69296a8-__obf_33a1f511b09ac2af))
			if __obf_b95967467beb5e15 <= float64(__obf_9dba054a785bce28) {
				__obf_81fd9df55ede76fb.
					Set(__obf_3e26d844e26658db, __obf_9aa12d47f69296a8, color.White)
			} else {
				__obf_81fd9df55ede76fb.
					Set(__obf_3e26d844e26658db, __obf_9aa12d47f69296a8, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_81fd9df55ede76fb, __obf_81fd9df55ede76fb.Bounds(), __obf_cbeabe4ae9e49861, image.Point{X: 0, Y: 0}, __obf_81fd9df55ede76fb, image.Point{}, draw.Over)
	__obf_cbeabe4ae9e49861.
		NRGBA = __obf_81fd9df55ede76fb
}

// CropScaleCircle scales and crops a circular area
func (__obf_cbeabe4ae9e49861 *NRGBA) CropScaleCircle(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_9dba054a785bce28 int, __obf_97b9ba5dfd7cfec2 int) {
	__obf_0b934908b11a6f63 := __obf_cbeabe4ae9e49861.Bounds()
	__obf_81fd9df55ede76fb := image.NewNRGBA(__obf_0b934908b11a6f63)

	for __obf_9aa12d47f69296a8 := __obf_0b934908b11a6f63.Min.Y; __obf_9aa12d47f69296a8 < __obf_0b934908b11a6f63.Max.Y; __obf_9aa12d47f69296a8++ {
		for __obf_3e26d844e26658db := __obf_0b934908b11a6f63.Min.X; __obf_3e26d844e26658db < __obf_0b934908b11a6f63.Max.X; __obf_3e26d844e26658db++ {
			__obf_b95967467beb5e15 := math.Hypot(float64(__obf_3e26d844e26658db-__obf_48afb3005cd4a35c), float64(__obf_9aa12d47f69296a8-__obf_33a1f511b09ac2af))
			if __obf_b95967467beb5e15 <= float64(__obf_9dba054a785bce28) {
				__obf_81fd9df55ede76fb.
					Set(__obf_3e26d844e26658db, __obf_9aa12d47f69296a8, color.White)
			} else {
				__obf_81fd9df55ede76fb.
					Set(__obf_3e26d844e26658db, __obf_9aa12d47f69296a8, color.Transparent)
			}
		}
	}

	if __obf_97b9ba5dfd7cfec2 > 0 {
		__obf_586c3d7db20a27e0 := __obf_97b9ba5dfd7cfec2 * 2
		__obf_60fe907884237f74 := image.NewNRGBA(image.Rect(0, 0, __obf_cbeabe4ae9e49861.Bounds().Dx()-__obf_586c3d7db20a27e0, __obf_cbeabe4ae9e49861.Bounds().Dy()-__obf_586c3d7db20a27e0))
		draw.BiLinear.Scale(__obf_60fe907884237f74, __obf_60fe907884237f74.Bounds(), __obf_81fd9df55ede76fb, __obf_81fd9df55ede76fb.Bounds(), draw.Over, nil)
		__obf_81fd9df55ede76fb = __obf_60fe907884237f74
	}

	draw.DrawMask(__obf_81fd9df55ede76fb, __obf_81fd9df55ede76fb.Bounds(), __obf_cbeabe4ae9e49861, image.Point{X: __obf_97b9ba5dfd7cfec2, Y: __obf_97b9ba5dfd7cfec2}, __obf_81fd9df55ede76fb, image.Point{}, draw.Over)
	__obf_cbeabe4ae9e49861.
		NRGBA = __obf_81fd9df55ede76fb
}

// Scale scales the canvas
func (__obf_cbeabe4ae9e49861 *NRGBA) Scale(__obf_97b9ba5dfd7cfec2 int, __obf_c671268a5e208079, __obf_e072103bd3d49152 bool) {
	__obf_10d21d6285312279 := __obf_cbeabe4ae9e49861.NRGBA
	if __obf_97b9ba5dfd7cfec2 > 0 {
		__obf_586c3d7db20a27e0 := __obf_97b9ba5dfd7cfec2 * 2
		__obf_4e29246903c3938f := __obf_cbeabe4ae9e49861.Bounds().Dx() - __obf_586c3d7db20a27e0
		__obf_ae489a729c4f9d98 := __obf_cbeabe4ae9e49861.Bounds().Dy() - __obf_586c3d7db20a27e0
		__obf_8a2207a04ed493d5 := image.NewNRGBA(image.Rect(0, 0, __obf_4e29246903c3938f, __obf_ae489a729c4f9d98))

		if !__obf_c671268a5e208079 {
			draw.BiLinear.Scale(__obf_8a2207a04ed493d5, __obf_8a2207a04ed493d5.Bounds(), __obf_cbeabe4ae9e49861, __obf_cbeabe4ae9e49861.Bounds(), draw.Over, nil)
		} else {
			__obf_e8ddb8cf8eeb6425 := util.CalcResizedRect(__obf_cbeabe4ae9e49861.Bounds(), __obf_4e29246903c3938f, __obf_ae489a729c4f9d98, __obf_e072103bd3d49152)
			draw.ApproxBiLinear.Scale(__obf_8a2207a04ed493d5, __obf_e8ddb8cf8eeb6425.Bounds(), __obf_cbeabe4ae9e49861, __obf_cbeabe4ae9e49861.Bounds(), draw.Over, nil)
		}
		__obf_10d21d6285312279 = __obf_8a2207a04ed493d5
	}
	__obf_cbeabe4ae9e49861.
		NRGBA = __obf_10d21d6285312279
}
