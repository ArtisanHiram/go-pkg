package __obf_b0f3614e51931450

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_698deef70980b293 image.Rectangle, __obf_555e40bb748b48c9 bool) *NRGBA {
	__obf_73d09a6a149c4def := image.NewNRGBA(__obf_698deef70980b293)
	for __obf_431a48e7317ece6d := 0; __obf_431a48e7317ece6d < __obf_698deef70980b293.Max.Y; __obf_431a48e7317ece6d++ {
		for __obf_80fae96c5947b2de := 0; __obf_80fae96c5947b2de < __obf_698deef70980b293.Max.X; __obf_80fae96c5947b2de++ {
			if __obf_555e40bb748b48c9 {
				__obf_73d09a6a149c4def.Set(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, color.Alpha{A: uint8(0)})
			} else {
				__obf_73d09a6a149c4def.Set(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{
		__obf_73d09a6a149c4def,
	}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_c263d022ccedf9ef *NRGBA) DrawImage(__obf_71f0387c46e538dc Palette, __obf_496a09259abf4375 *PositionRect, __obf_c2f46c13562c1258 *AreaRect) {
	__obf_e61ea68b462e25a1 := __obf_71f0387c46e538dc.Bounds().Max.X
	__obf_5cc37636829791d6 := __obf_71f0387c46e538dc.Bounds().Max.Y

	__obf_3cb3cefd9e6069a6 := __obf_496a09259abf4375.X
	__obf_9041003a8e179884 := __obf_496a09259abf4375.Y
	__obf_f6e77f24316edfe3 := __obf_496a09259abf4375.Height

	__obf_0e2ccf475f5c0fc4 := __obf_c2f46c13562c1258.MinX
	__obf_f4ec86d75ac10909 := __obf_c2f46c13562c1258.MinY
	__obf_ef727839c0f5d6c8 := __obf_c2f46c13562c1258.MaxX
	__obf_5e98270467f34ac9 := __obf_c2f46c13562c1258.MaxY

	for __obf_80fae96c5947b2de := range __obf_e61ea68b462e25a1 {
		for __obf_431a48e7317ece6d := range __obf_5cc37636829791d6 {
			__obf_ddc50ca41a79bca6 := __obf_71f0387c46e538dc.At(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d)
			if _, _, _, __obf_bd896093c58f47ac := __obf_ddc50ca41a79bca6.RGBA(); __obf_bd896093c58f47ac > 0 {
				if __obf_80fae96c5947b2de >= __obf_0e2ccf475f5c0fc4 && __obf_80fae96c5947b2de <= __obf_ef727839c0f5d6c8 && __obf_431a48e7317ece6d >= __obf_f4ec86d75ac10909 && __obf_431a48e7317ece6d <= __obf_5e98270467f34ac9 {
					__obf_c263d022ccedf9ef.Set(__obf_3cb3cefd9e6069a6+(__obf_80fae96c5947b2de-__obf_0e2ccf475f5c0fc4), __obf_9041003a8e179884-__obf_f6e77f24316edfe3+(__obf_431a48e7317ece6d-__obf_f4ec86d75ac10909), __obf_71f0387c46e538dc.At(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_c263d022ccedf9ef *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_e61ea68b462e25a1 := __obf_c263d022ccedf9ef.Bounds().Max.X
	__obf_5cc37636829791d6 := __obf_c263d022ccedf9ef.Bounds().Max.Y
	__obf_00491d31e8fc60c8 := __obf_e61ea68b462e25a1
	__obf_7c3283c789f60bf3 := 0
	__obf_4995721af19641b6 := __obf_5cc37636829791d6
	__obf_c0a3fea089ed7c65 := 0
	for __obf_80fae96c5947b2de := range __obf_e61ea68b462e25a1 {
		for __obf_431a48e7317ece6d := range __obf_5cc37636829791d6 {
			__obf_ddc50ca41a79bca6 := __obf_c263d022ccedf9ef.At(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d)
			if _, _, _, __obf_bd896093c58f47ac := __obf_ddc50ca41a79bca6.RGBA(); __obf_bd896093c58f47ac > 0 {
				if __obf_80fae96c5947b2de < __obf_00491d31e8fc60c8 {
					__obf_00491d31e8fc60c8 = __obf_80fae96c5947b2de
				}
				if __obf_80fae96c5947b2de > __obf_7c3283c789f60bf3 {
					__obf_7c3283c789f60bf3 = __obf_80fae96c5947b2de
				}

				if __obf_431a48e7317ece6d < __obf_4995721af19641b6 {
					__obf_4995721af19641b6 = __obf_431a48e7317ece6d
				}
				if __obf_431a48e7317ece6d > __obf_c0a3fea089ed7c65 {
					__obf_c0a3fea089ed7c65 = __obf_431a48e7317ece6d
				}
			}
		}
	}

	__obf_00491d31e8fc60c8 = int(max(0, float64(__obf_00491d31e8fc60c8-2)))
	__obf_7c3283c789f60bf3 = int(min(float64(__obf_e61ea68b462e25a1), float64(__obf_7c3283c789f60bf3+2)))
	__obf_4995721af19641b6 = int(max(0, float64(__obf_4995721af19641b6-2)))
	__obf_c0a3fea089ed7c65 = int(min(float64(__obf_5cc37636829791d6), float64(__obf_c0a3fea089ed7c65+2)))

	return &AreaRect{
		__obf_00491d31e8fc60c8,
		__obf_7c3283c789f60bf3,
		__obf_4995721af19641b6,
		__obf_c0a3fea089ed7c65,
	}
}

// Rotate rotates the canvas by any angle
func (__obf_c263d022ccedf9ef *NRGBA) Rotate(__obf_bd896093c58f47ac int, __obf_335f79eb2cb18bba bool) {
	if __obf_bd896093c58f47ac == 0 {
		return
	}

	__obf_31d599a870027db7 := float64(__obf_bd896093c58f47ac) * math.Pi / 180

	__obf_de52c489f5336494 := __obf_c263d022ccedf9ef.Bounds().Dx()
	__obf_5efd21b62fd26e68 := __obf_c263d022ccedf9ef.Bounds().Dy()
	__obf_ac67bef5710e786a, __obf_e455662263d832db := util.RotatedSize(__obf_de52c489f5336494, __obf_5efd21b62fd26e68, float64(__obf_bd896093c58f47ac))
	__obf_71f0387c46e538dc := image.NewNRGBA(image.Rect(0, 0, __obf_ac67bef5710e786a, __obf_e455662263d832db))

	__obf_1b99d214cd2bfbd1 := float64(__obf_ac67bef5710e786a) / 2
	__obf_3b7710d3187dd4b3 := float64(__obf_e455662263d832db) / 2

	__obf_415742d04bd36878 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_415742d04bd36878 = __obf_415742d04bd36878.Translate(__obf_1b99d214cd2bfbd1, __obf_3b7710d3187dd4b3)
	__obf_415742d04bd36878 = __obf_415742d04bd36878.Rotate(__obf_31d599a870027db7)
	__obf_415742d04bd36878 = __obf_415742d04bd36878.Translate(-__obf_1b99d214cd2bfbd1, -__obf_3b7710d3187dd4b3)

	__obf_80fae96c5947b2de := (__obf_ac67bef5710e786a - __obf_c263d022ccedf9ef.Bounds().Dx()) / 2
	__obf_431a48e7317ece6d := (__obf_e455662263d832db - __obf_c263d022ccedf9ef.Bounds().Dy()) / 2
	__obf_49981259e794ccde, __obf_c4e32409976dca90 := float64(__obf_80fae96c5947b2de), float64(__obf_431a48e7317ece6d)

	__obf_b72b2bb410c8c093 := __obf_415742d04bd36878.Translate(__obf_49981259e794ccde, __obf_c4e32409976dca90)
	__obf_ca8490bd1939f470 := f64.Aff3{__obf_b72b2bb410c8c093.XX, __obf_b72b2bb410c8c093.XY, __obf_b72b2bb410c8c093.X0, __obf_b72b2bb410c8c093.YX, __obf_b72b2bb410c8c093.YY, __obf_b72b2bb410c8c093.Y0}

	draw.BiLinear.Transform(__obf_71f0387c46e538dc, __obf_ca8490bd1939f470, __obf_c263d022ccedf9ef, __obf_c263d022ccedf9ef.Bounds(), draw.Over, nil)
	__obf_c263d022ccedf9ef.NRGBA = __obf_71f0387c46e538dc

	if __obf_335f79eb2cb18bba {
		__obf_4f8a26cb6b532810 := __obf_ac67bef5710e786a - __obf_de52c489f5336494
		__obf_e984e5ff27154922 := __obf_e455662263d832db - __obf_5efd21b62fd26e68
		__obf_5a4d5259564716ed := (__obf_4f8a26cb6b532810 / 2) + 1
		__obf_1a34e4081dd5892b := (__obf_e984e5ff27154922 / 2) + 1
		__obf_c263d022ccedf9ef.SubImage(image.Rect(__obf_5a4d5259564716ed, __obf_1a34e4081dd5892b, __obf_de52c489f5336494+__obf_5a4d5259564716ed, __obf_5efd21b62fd26e68+__obf_1a34e4081dd5892b))
	}
}

// CropCircle crops a circular area
func (__obf_c263d022ccedf9ef *NRGBA) CropCircle(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_b290e37944a6ece3 int) {
	__obf_6dc455be3103c65e := __obf_c263d022ccedf9ef.Bounds()
	__obf_0f1398799e6e7846 := image.NewNRGBA(__obf_6dc455be3103c65e)
	for __obf_ab7647cc0c07f11c := __obf_6dc455be3103c65e.Min.Y; __obf_ab7647cc0c07f11c < __obf_6dc455be3103c65e.Max.Y; __obf_ab7647cc0c07f11c++ {
		for __obf_8d678fbb17742ef6 := __obf_6dc455be3103c65e.Min.X; __obf_8d678fbb17742ef6 < __obf_6dc455be3103c65e.Max.X; __obf_8d678fbb17742ef6++ {
			__obf_1032b003cda4a2b2 := math.Hypot(float64(__obf_8d678fbb17742ef6-__obf_80fae96c5947b2de), float64(__obf_ab7647cc0c07f11c-__obf_431a48e7317ece6d))
			if __obf_1032b003cda4a2b2 <= float64(__obf_b290e37944a6ece3) {
				__obf_0f1398799e6e7846.Set(__obf_8d678fbb17742ef6, __obf_ab7647cc0c07f11c, color.White)
			} else {
				__obf_0f1398799e6e7846.Set(__obf_8d678fbb17742ef6, __obf_ab7647cc0c07f11c, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_0f1398799e6e7846, __obf_0f1398799e6e7846.Bounds(), __obf_c263d022ccedf9ef, image.Point{X: 0, Y: 0}, __obf_0f1398799e6e7846, image.Point{}, draw.Over)
	__obf_c263d022ccedf9ef.NRGBA = __obf_0f1398799e6e7846
}

// CropScaleCircle scales and crops a circular area
func (__obf_c263d022ccedf9ef *NRGBA) CropScaleCircle(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_b290e37944a6ece3 int, __obf_11692ea0873041cf int) {
	__obf_6dc455be3103c65e := __obf_c263d022ccedf9ef.Bounds()
	__obf_0f1398799e6e7846 := image.NewNRGBA(__obf_6dc455be3103c65e)

	for __obf_ab7647cc0c07f11c := __obf_6dc455be3103c65e.Min.Y; __obf_ab7647cc0c07f11c < __obf_6dc455be3103c65e.Max.Y; __obf_ab7647cc0c07f11c++ {
		for __obf_8d678fbb17742ef6 := __obf_6dc455be3103c65e.Min.X; __obf_8d678fbb17742ef6 < __obf_6dc455be3103c65e.Max.X; __obf_8d678fbb17742ef6++ {
			__obf_1032b003cda4a2b2 := math.Hypot(float64(__obf_8d678fbb17742ef6-__obf_80fae96c5947b2de), float64(__obf_ab7647cc0c07f11c-__obf_431a48e7317ece6d))
			if __obf_1032b003cda4a2b2 <= float64(__obf_b290e37944a6ece3) {
				__obf_0f1398799e6e7846.Set(__obf_8d678fbb17742ef6, __obf_ab7647cc0c07f11c, color.White)
			} else {
				__obf_0f1398799e6e7846.Set(__obf_8d678fbb17742ef6, __obf_ab7647cc0c07f11c, color.Transparent)
			}
		}
	}

	if __obf_11692ea0873041cf > 0 {
		__obf_6acceae67ea7a6e2 := __obf_11692ea0873041cf * 2
		__obf_bf3bac509ba0e5f3 := image.NewNRGBA(image.Rect(0, 0, __obf_c263d022ccedf9ef.Bounds().Dx()-__obf_6acceae67ea7a6e2, __obf_c263d022ccedf9ef.Bounds().Dy()-__obf_6acceae67ea7a6e2))
		draw.BiLinear.Scale(__obf_bf3bac509ba0e5f3, __obf_bf3bac509ba0e5f3.Bounds(), __obf_0f1398799e6e7846, __obf_0f1398799e6e7846.Bounds(), draw.Over, nil)
		__obf_0f1398799e6e7846 = __obf_bf3bac509ba0e5f3
	}

	draw.DrawMask(__obf_0f1398799e6e7846, __obf_0f1398799e6e7846.Bounds(), __obf_c263d022ccedf9ef, image.Point{X: __obf_11692ea0873041cf, Y: __obf_11692ea0873041cf}, __obf_0f1398799e6e7846, image.Point{}, draw.Over)
	__obf_c263d022ccedf9ef.NRGBA = __obf_0f1398799e6e7846
}

// Scale scales the canvas
func (__obf_c263d022ccedf9ef *NRGBA) Scale(__obf_11692ea0873041cf int, __obf_341fec743dd3154c, __obf_992a051f1ea20d54 bool) {
	__obf_71f0387c46e538dc := __obf_c263d022ccedf9ef.NRGBA
	if __obf_11692ea0873041cf > 0 {
		__obf_6acceae67ea7a6e2 := __obf_11692ea0873041cf * 2
		__obf_040ab8ca741c209c := __obf_c263d022ccedf9ef.Bounds().Dx() - __obf_6acceae67ea7a6e2
		__obf_da2c34c7bf21135f := __obf_c263d022ccedf9ef.Bounds().Dy() - __obf_6acceae67ea7a6e2
		__obf_24cb14cc3a2939b9 := image.NewNRGBA(image.Rect(0, 0, __obf_040ab8ca741c209c, __obf_da2c34c7bf21135f))

		if !__obf_341fec743dd3154c {
			draw.BiLinear.Scale(__obf_24cb14cc3a2939b9, __obf_24cb14cc3a2939b9.Bounds(), __obf_c263d022ccedf9ef, __obf_c263d022ccedf9ef.Bounds(), draw.Over, nil)
		} else {
			__obf_9a0fadfcc7402285 := util.CalcResizedRect(__obf_c263d022ccedf9ef.Bounds(), __obf_040ab8ca741c209c, __obf_da2c34c7bf21135f, __obf_992a051f1ea20d54)
			draw.ApproxBiLinear.Scale(__obf_24cb14cc3a2939b9, __obf_9a0fadfcc7402285.Bounds(), __obf_c263d022ccedf9ef, __obf_c263d022ccedf9ef.Bounds(), draw.Over, nil)
		}

		__obf_71f0387c46e538dc = __obf_24cb14cc3a2939b9
	}

	__obf_c263d022ccedf9ef.NRGBA = __obf_71f0387c46e538dc
}
