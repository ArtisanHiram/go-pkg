package __obf_738b46210fdb4199

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_c09045abdc5d3619 image.Rectangle, __obf_4fd62f435078a950 bool) *NRGBA {
	__obf_fcf9917e57a78dde := image.NewNRGBA(__obf_c09045abdc5d3619)
	for __obf_a69b4d01937bfbfc := 0; __obf_a69b4d01937bfbfc < __obf_c09045abdc5d3619.Max.Y; __obf_a69b4d01937bfbfc++ {
		for __obf_bd372f9dc9e16d75 := 0; __obf_bd372f9dc9e16d75 < __obf_c09045abdc5d3619.Max.X; __obf_bd372f9dc9e16d75++ {
			if __obf_4fd62f435078a950 {
				__obf_fcf9917e57a78dde.
					Set(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, color.Alpha{A: uint8(0)})
			} else {
				__obf_fcf9917e57a78dde.
					Set(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_fcf9917e57a78dde}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_106eeae4b6afaaf2 *NRGBA) DrawImage(__obf_0912be1955a21fd1 Palette, __obf_e96f68ddd4911047 *PositionRect, __obf_3db57b163baf9558 *AreaRect) {
	__obf_d17e3b2fb9b48db3 := __obf_0912be1955a21fd1.Bounds().Max.X
	__obf_e0ad449835d092e9 := __obf_0912be1955a21fd1.Bounds().Max.Y
	__obf_01996b16e0baaf3e := __obf_e96f68ddd4911047.X
	__obf_dd85267ddd3b467b := __obf_e96f68ddd4911047.Y
	__obf_2f78cda9221b6a31 := __obf_e96f68ddd4911047.Height
	__obf_83d427569ab9b8a2 := __obf_3db57b163baf9558.MinX
	__obf_922ffdbc3b430129 := __obf_3db57b163baf9558.MinY
	__obf_2c15510d8b6e7835 := __obf_3db57b163baf9558.MaxX
	__obf_b22b5a86f8bb3635 := __obf_3db57b163baf9558.MaxY

	for __obf_bd372f9dc9e16d75 := range __obf_d17e3b2fb9b48db3 {
		for __obf_a69b4d01937bfbfc := range __obf_e0ad449835d092e9 {
			__obf_e4c726f405e92dc2 := __obf_0912be1955a21fd1.At(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc)
			if _, _, _, __obf_5eeb26a04aed8320 := __obf_e4c726f405e92dc2.RGBA(); __obf_5eeb26a04aed8320 > 0 {
				if __obf_bd372f9dc9e16d75 >= __obf_83d427569ab9b8a2 && __obf_bd372f9dc9e16d75 <= __obf_2c15510d8b6e7835 && __obf_a69b4d01937bfbfc >= __obf_922ffdbc3b430129 && __obf_a69b4d01937bfbfc <= __obf_b22b5a86f8bb3635 {
					__obf_106eeae4b6afaaf2.
						Set(__obf_01996b16e0baaf3e+(__obf_bd372f9dc9e16d75-__obf_83d427569ab9b8a2), __obf_dd85267ddd3b467b-__obf_2f78cda9221b6a31+(__obf_a69b4d01937bfbfc-__obf_922ffdbc3b430129), __obf_0912be1955a21fd1.At(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_106eeae4b6afaaf2 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_d17e3b2fb9b48db3 := __obf_106eeae4b6afaaf2.Bounds().Max.X
	__obf_e0ad449835d092e9 := __obf_106eeae4b6afaaf2.Bounds().Max.Y
	__obf_ba97c5a5606ddd7e := __obf_d17e3b2fb9b48db3
	__obf_c6b746bbb5d1f238 := 0
	__obf_36f26cd4c2083cc6 := __obf_e0ad449835d092e9
	__obf_7d40f40942b754c0 := 0
	for __obf_bd372f9dc9e16d75 := range __obf_d17e3b2fb9b48db3 {
		for __obf_a69b4d01937bfbfc := range __obf_e0ad449835d092e9 {
			__obf_e4c726f405e92dc2 := __obf_106eeae4b6afaaf2.At(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc)
			if _, _, _, __obf_5eeb26a04aed8320 := __obf_e4c726f405e92dc2.RGBA(); __obf_5eeb26a04aed8320 > 0 {
				if __obf_bd372f9dc9e16d75 < __obf_ba97c5a5606ddd7e {
					__obf_ba97c5a5606ddd7e = __obf_bd372f9dc9e16d75
				}
				if __obf_bd372f9dc9e16d75 > __obf_c6b746bbb5d1f238 {
					__obf_c6b746bbb5d1f238 = __obf_bd372f9dc9e16d75
				}

				if __obf_a69b4d01937bfbfc < __obf_36f26cd4c2083cc6 {
					__obf_36f26cd4c2083cc6 = __obf_a69b4d01937bfbfc
				}
				if __obf_a69b4d01937bfbfc > __obf_7d40f40942b754c0 {
					__obf_7d40f40942b754c0 = __obf_a69b4d01937bfbfc
				}
			}
		}
	}
	__obf_ba97c5a5606ddd7e = int(max(0, float64(__obf_ba97c5a5606ddd7e-2)))
	__obf_c6b746bbb5d1f238 = int(min(float64(__obf_d17e3b2fb9b48db3), float64(__obf_c6b746bbb5d1f238+2)))
	__obf_36f26cd4c2083cc6 = int(max(0, float64(__obf_36f26cd4c2083cc6-2)))
	__obf_7d40f40942b754c0 = int(min(float64(__obf_e0ad449835d092e9), float64(__obf_7d40f40942b754c0+2)))

	return &AreaRect{__obf_ba97c5a5606ddd7e, __obf_c6b746bbb5d1f238,

		// Rotate rotates the canvas by any angle
		__obf_36f26cd4c2083cc6, __obf_7d40f40942b754c0,
	}
}

func (__obf_106eeae4b6afaaf2 *NRGBA) Rotate(__obf_5eeb26a04aed8320 int, __obf_6562df73619e4f32 bool) {
	if __obf_5eeb26a04aed8320 == 0 {
		return
	}
	__obf_3a2dad55c354b1d6 := float64(__obf_5eeb26a04aed8320) * math.Pi / 180
	__obf_256ef905b77b961d := __obf_106eeae4b6afaaf2.Bounds().Dx()
	__obf_8db0e99c5de01893 := __obf_106eeae4b6afaaf2.Bounds().Dy()
	__obf_ccd0c8ac1f16e3bd, __obf_90a966c355c00010 := util.RotatedSize(__obf_256ef905b77b961d, __obf_8db0e99c5de01893, float64(__obf_5eeb26a04aed8320))
	__obf_0912be1955a21fd1 := image.NewNRGBA(image.Rect(0, 0, __obf_ccd0c8ac1f16e3bd, __obf_90a966c355c00010))
	__obf_e0be35b6ff05a6cb := float64(__obf_ccd0c8ac1f16e3bd) / 2
	__obf_20166c1fac2cf6ec := float64(__obf_90a966c355c00010) / 2
	__obf_cbb89fd7b04f7ec6 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_cbb89fd7b04f7ec6 = __obf_cbb89fd7b04f7ec6.Translate(__obf_e0be35b6ff05a6cb, __obf_20166c1fac2cf6ec)
	__obf_cbb89fd7b04f7ec6 = __obf_cbb89fd7b04f7ec6.Rotate(__obf_3a2dad55c354b1d6)
	__obf_cbb89fd7b04f7ec6 = __obf_cbb89fd7b04f7ec6.Translate(-__obf_e0be35b6ff05a6cb, -__obf_20166c1fac2cf6ec)
	__obf_bd372f9dc9e16d75 := (__obf_ccd0c8ac1f16e3bd - __obf_106eeae4b6afaaf2.Bounds().Dx()) / 2
	__obf_a69b4d01937bfbfc := (__obf_90a966c355c00010 - __obf_106eeae4b6afaaf2.Bounds().Dy()) / 2
	__obf_4740fb0bc3ed9a47, __obf_7f330bbdccf6a7a2 := float64(__obf_bd372f9dc9e16d75), float64(__obf_a69b4d01937bfbfc)
	__obf_8fee99affc23d01a := __obf_cbb89fd7b04f7ec6.Translate(__obf_4740fb0bc3ed9a47, __obf_7f330bbdccf6a7a2)
	__obf_8d4cdf6457eacc6e := f64.Aff3{__obf_8fee99affc23d01a.XX, __obf_8fee99affc23d01a.XY, __obf_8fee99affc23d01a.X0, __obf_8fee99affc23d01a.YX, __obf_8fee99affc23d01a.YY, __obf_8fee99affc23d01a.Y0}

	draw.BiLinear.Transform(__obf_0912be1955a21fd1, __obf_8d4cdf6457eacc6e, __obf_106eeae4b6afaaf2, __obf_106eeae4b6afaaf2.Bounds(), draw.Over, nil)
	__obf_106eeae4b6afaaf2.
		NRGBA = __obf_0912be1955a21fd1

	if __obf_6562df73619e4f32 {
		__obf_75abb26d19e4fe26 := __obf_ccd0c8ac1f16e3bd - __obf_256ef905b77b961d
		__obf_89fd261ea92ca04b := __obf_90a966c355c00010 - __obf_8db0e99c5de01893
		__obf_936088d1b4a69899 := (__obf_75abb26d19e4fe26 / 2) + 1
		__obf_141144958ec0757c := (__obf_89fd261ea92ca04b / 2) + 1
		__obf_106eeae4b6afaaf2.
			SubImage(image.Rect(__obf_936088d1b4a69899, __obf_141144958ec0757c,

				// CropCircle crops a circular area
				__obf_256ef905b77b961d+__obf_936088d1b4a69899, __obf_8db0e99c5de01893+__obf_141144958ec0757c))
	}
}

func (__obf_106eeae4b6afaaf2 *NRGBA) CropCircle(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_6b93bc026df65fff int) {
	__obf_9c71f470ce859f64 := __obf_106eeae4b6afaaf2.Bounds()
	__obf_1704409e361bc977 := image.NewNRGBA(__obf_9c71f470ce859f64)
	for __obf_62c582c6d89a6880 := __obf_9c71f470ce859f64.Min.Y; __obf_62c582c6d89a6880 < __obf_9c71f470ce859f64.Max.Y; __obf_62c582c6d89a6880++ {
		for __obf_3d5d9f5d275b53cc := __obf_9c71f470ce859f64.Min.X; __obf_3d5d9f5d275b53cc < __obf_9c71f470ce859f64.Max.X; __obf_3d5d9f5d275b53cc++ {
			__obf_6a5131e41ef7c9a7 := math.Hypot(float64(__obf_3d5d9f5d275b53cc-__obf_bd372f9dc9e16d75), float64(__obf_62c582c6d89a6880-__obf_a69b4d01937bfbfc))
			if __obf_6a5131e41ef7c9a7 <= float64(__obf_6b93bc026df65fff) {
				__obf_1704409e361bc977.
					Set(__obf_3d5d9f5d275b53cc, __obf_62c582c6d89a6880, color.White)
			} else {
				__obf_1704409e361bc977.
					Set(__obf_3d5d9f5d275b53cc, __obf_62c582c6d89a6880, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_1704409e361bc977, __obf_1704409e361bc977.Bounds(), __obf_106eeae4b6afaaf2, image.Point{X: 0, Y: 0}, __obf_1704409e361bc977, image.Point{}, draw.Over)
	__obf_106eeae4b6afaaf2.
		NRGBA = __obf_1704409e361bc977
}

// CropScaleCircle scales and crops a circular area
func (__obf_106eeae4b6afaaf2 *NRGBA) CropScaleCircle(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_6b93bc026df65fff int, __obf_d591e7f05b1c54d9 int) {
	__obf_9c71f470ce859f64 := __obf_106eeae4b6afaaf2.Bounds()
	__obf_1704409e361bc977 := image.NewNRGBA(__obf_9c71f470ce859f64)

	for __obf_62c582c6d89a6880 := __obf_9c71f470ce859f64.Min.Y; __obf_62c582c6d89a6880 < __obf_9c71f470ce859f64.Max.Y; __obf_62c582c6d89a6880++ {
		for __obf_3d5d9f5d275b53cc := __obf_9c71f470ce859f64.Min.X; __obf_3d5d9f5d275b53cc < __obf_9c71f470ce859f64.Max.X; __obf_3d5d9f5d275b53cc++ {
			__obf_6a5131e41ef7c9a7 := math.Hypot(float64(__obf_3d5d9f5d275b53cc-__obf_bd372f9dc9e16d75), float64(__obf_62c582c6d89a6880-__obf_a69b4d01937bfbfc))
			if __obf_6a5131e41ef7c9a7 <= float64(__obf_6b93bc026df65fff) {
				__obf_1704409e361bc977.
					Set(__obf_3d5d9f5d275b53cc, __obf_62c582c6d89a6880, color.White)
			} else {
				__obf_1704409e361bc977.
					Set(__obf_3d5d9f5d275b53cc, __obf_62c582c6d89a6880, color.Transparent)
			}
		}
	}

	if __obf_d591e7f05b1c54d9 > 0 {
		__obf_d4d588cd8a11c009 := __obf_d591e7f05b1c54d9 * 2
		__obf_780d6afa4c16771f := image.NewNRGBA(image.Rect(0, 0, __obf_106eeae4b6afaaf2.Bounds().Dx()-__obf_d4d588cd8a11c009, __obf_106eeae4b6afaaf2.Bounds().Dy()-__obf_d4d588cd8a11c009))
		draw.BiLinear.Scale(__obf_780d6afa4c16771f, __obf_780d6afa4c16771f.Bounds(), __obf_1704409e361bc977, __obf_1704409e361bc977.Bounds(), draw.Over, nil)
		__obf_1704409e361bc977 = __obf_780d6afa4c16771f
	}

	draw.DrawMask(__obf_1704409e361bc977, __obf_1704409e361bc977.Bounds(), __obf_106eeae4b6afaaf2, image.Point{X: __obf_d591e7f05b1c54d9, Y: __obf_d591e7f05b1c54d9}, __obf_1704409e361bc977, image.Point{}, draw.Over)
	__obf_106eeae4b6afaaf2.
		NRGBA = __obf_1704409e361bc977
}

// Scale scales the canvas
func (__obf_106eeae4b6afaaf2 *NRGBA) Scale(__obf_d591e7f05b1c54d9 int, __obf_533571087bc944bb, __obf_5771891454be5a0a bool) {
	__obf_0912be1955a21fd1 := __obf_106eeae4b6afaaf2.NRGBA
	if __obf_d591e7f05b1c54d9 > 0 {
		__obf_d4d588cd8a11c009 := __obf_d591e7f05b1c54d9 * 2
		__obf_44a8bffdcd371bf3 := __obf_106eeae4b6afaaf2.Bounds().Dx() - __obf_d4d588cd8a11c009
		__obf_4388af138cf0ba1b := __obf_106eeae4b6afaaf2.Bounds().Dy() - __obf_d4d588cd8a11c009
		__obf_09fbe41990570987 := image.NewNRGBA(image.Rect(0, 0, __obf_44a8bffdcd371bf3, __obf_4388af138cf0ba1b))

		if !__obf_533571087bc944bb {
			draw.BiLinear.Scale(__obf_09fbe41990570987, __obf_09fbe41990570987.Bounds(), __obf_106eeae4b6afaaf2, __obf_106eeae4b6afaaf2.Bounds(), draw.Over, nil)
		} else {
			__obf_efd28ef94e99a215 := util.CalcResizedRect(__obf_106eeae4b6afaaf2.Bounds(), __obf_44a8bffdcd371bf3, __obf_4388af138cf0ba1b, __obf_5771891454be5a0a)
			draw.ApproxBiLinear.Scale(__obf_09fbe41990570987, __obf_efd28ef94e99a215.Bounds(), __obf_106eeae4b6afaaf2, __obf_106eeae4b6afaaf2.Bounds(), draw.Over, nil)
		}
		__obf_0912be1955a21fd1 = __obf_09fbe41990570987
	}
	__obf_106eeae4b6afaaf2.
		NRGBA = __obf_0912be1955a21fd1
}
