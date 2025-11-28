package __obf_154ce31cd9492d61

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_c3f6882d3032b52f image.Rectangle, __obf_2ce0631e145ef82d bool) *NRGBA {
	__obf_c963037826c31de4 := image.NewNRGBA(__obf_c3f6882d3032b52f)
	for __obf_f8fead3d682d659f := 0; __obf_f8fead3d682d659f < __obf_c3f6882d3032b52f.Max.Y; __obf_f8fead3d682d659f++ {
		for __obf_c87d6c7ba4a30d82 := 0; __obf_c87d6c7ba4a30d82 < __obf_c3f6882d3032b52f.Max.X; __obf_c87d6c7ba4a30d82++ {
			if __obf_2ce0631e145ef82d {
				__obf_c963037826c31de4.Set(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, color.Alpha{A: uint8(0)})
			} else {
				__obf_c963037826c31de4.Set(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{
		__obf_c963037826c31de4,
	}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_3f4810c67144854c *NRGBA) DrawImage(__obf_7e8dc6da215ff393 Palette, __obf_0e93d017f4cc87c1 *PositionRect, __obf_2eb7c830163b0f4d *AreaRect) {
	__obf_d14caaad13625683 := __obf_7e8dc6da215ff393.Bounds().Max.X
	__obf_ff9981d45d8deb98 := __obf_7e8dc6da215ff393.Bounds().Max.Y

	__obf_1c0221c2fa56d79a := __obf_0e93d017f4cc87c1.X
	__obf_8fbf29f58472ce0b := __obf_0e93d017f4cc87c1.Y
	__obf_ef5f6a16754a34e3 := __obf_0e93d017f4cc87c1.Height

	__obf_dd9a9130eced73f8 := __obf_2eb7c830163b0f4d.MinX
	__obf_a01010f938a9b282 := __obf_2eb7c830163b0f4d.MinY
	__obf_ab69088a34360d7f := __obf_2eb7c830163b0f4d.MaxX
	__obf_7bd3edf1e1697cf6 := __obf_2eb7c830163b0f4d.MaxY

	for __obf_c87d6c7ba4a30d82 := range __obf_d14caaad13625683 {
		for __obf_f8fead3d682d659f := range __obf_ff9981d45d8deb98 {
			__obf_12c04581a749b5b3 := __obf_7e8dc6da215ff393.At(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f)
			if _, _, _, __obf_91f28fffbdd83b09 := __obf_12c04581a749b5b3.RGBA(); __obf_91f28fffbdd83b09 > 0 {
				if __obf_c87d6c7ba4a30d82 >= __obf_dd9a9130eced73f8 && __obf_c87d6c7ba4a30d82 <= __obf_ab69088a34360d7f && __obf_f8fead3d682d659f >= __obf_a01010f938a9b282 && __obf_f8fead3d682d659f <= __obf_7bd3edf1e1697cf6 {
					__obf_3f4810c67144854c.Set(__obf_1c0221c2fa56d79a+(__obf_c87d6c7ba4a30d82-__obf_dd9a9130eced73f8), __obf_8fbf29f58472ce0b-__obf_ef5f6a16754a34e3+(__obf_f8fead3d682d659f-__obf_a01010f938a9b282), __obf_7e8dc6da215ff393.At(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_3f4810c67144854c *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_d14caaad13625683 := __obf_3f4810c67144854c.Bounds().Max.X
	__obf_ff9981d45d8deb98 := __obf_3f4810c67144854c.Bounds().Max.Y
	__obf_2112cc534df9dd1a := __obf_d14caaad13625683
	__obf_5bcc151bb94dd6aa := 0
	__obf_f49db9de609cf423 := __obf_ff9981d45d8deb98
	__obf_e6c5d6e5feaf61f3 := 0
	for __obf_c87d6c7ba4a30d82 := range __obf_d14caaad13625683 {
		for __obf_f8fead3d682d659f := range __obf_ff9981d45d8deb98 {
			__obf_12c04581a749b5b3 := __obf_3f4810c67144854c.At(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f)
			if _, _, _, __obf_91f28fffbdd83b09 := __obf_12c04581a749b5b3.RGBA(); __obf_91f28fffbdd83b09 > 0 {
				if __obf_c87d6c7ba4a30d82 < __obf_2112cc534df9dd1a {
					__obf_2112cc534df9dd1a = __obf_c87d6c7ba4a30d82
				}
				if __obf_c87d6c7ba4a30d82 > __obf_5bcc151bb94dd6aa {
					__obf_5bcc151bb94dd6aa = __obf_c87d6c7ba4a30d82
				}

				if __obf_f8fead3d682d659f < __obf_f49db9de609cf423 {
					__obf_f49db9de609cf423 = __obf_f8fead3d682d659f
				}
				if __obf_f8fead3d682d659f > __obf_e6c5d6e5feaf61f3 {
					__obf_e6c5d6e5feaf61f3 = __obf_f8fead3d682d659f
				}
			}
		}
	}

	__obf_2112cc534df9dd1a = int(max(0, float64(__obf_2112cc534df9dd1a-2)))
	__obf_5bcc151bb94dd6aa = int(min(float64(__obf_d14caaad13625683), float64(__obf_5bcc151bb94dd6aa+2)))
	__obf_f49db9de609cf423 = int(max(0, float64(__obf_f49db9de609cf423-2)))
	__obf_e6c5d6e5feaf61f3 = int(min(float64(__obf_ff9981d45d8deb98), float64(__obf_e6c5d6e5feaf61f3+2)))

	return &AreaRect{
		__obf_2112cc534df9dd1a,
		__obf_5bcc151bb94dd6aa,
		__obf_f49db9de609cf423,
		__obf_e6c5d6e5feaf61f3,
	}
}

// Rotate rotates the canvas by any angle
func (__obf_3f4810c67144854c *NRGBA) Rotate(__obf_91f28fffbdd83b09 int, __obf_feccdc0703f37768 bool) {
	if __obf_91f28fffbdd83b09 == 0 {
		return
	}

	__obf_2c1097e2e9ce1f24 := float64(__obf_91f28fffbdd83b09) * math.Pi / 180

	__obf_80c4c4a4d9299b93 := __obf_3f4810c67144854c.Bounds().Dx()
	__obf_e93bb5827b72b466 := __obf_3f4810c67144854c.Bounds().Dy()
	__obf_5c507cf47bf5267e, __obf_f03bf0f36c5833f9 := util.RotatedSize(__obf_80c4c4a4d9299b93, __obf_e93bb5827b72b466, float64(__obf_91f28fffbdd83b09))
	__obf_7e8dc6da215ff393 := image.NewNRGBA(image.Rect(0, 0, __obf_5c507cf47bf5267e, __obf_f03bf0f36c5833f9))

	__obf_0d5e7256517c1808 := float64(__obf_5c507cf47bf5267e) / 2
	__obf_b92c251b6c66c475 := float64(__obf_f03bf0f36c5833f9) / 2

	__obf_03e9e17f4239cd40 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_03e9e17f4239cd40 = __obf_03e9e17f4239cd40.Translate(__obf_0d5e7256517c1808, __obf_b92c251b6c66c475)
	__obf_03e9e17f4239cd40 = __obf_03e9e17f4239cd40.Rotate(__obf_2c1097e2e9ce1f24)
	__obf_03e9e17f4239cd40 = __obf_03e9e17f4239cd40.Translate(-__obf_0d5e7256517c1808, -__obf_b92c251b6c66c475)

	__obf_c87d6c7ba4a30d82 := (__obf_5c507cf47bf5267e - __obf_3f4810c67144854c.Bounds().Dx()) / 2
	__obf_f8fead3d682d659f := (__obf_f03bf0f36c5833f9 - __obf_3f4810c67144854c.Bounds().Dy()) / 2
	__obf_eb502d163673590c, __obf_f55af572487ddde8 := float64(__obf_c87d6c7ba4a30d82), float64(__obf_f8fead3d682d659f)

	__obf_b56c2b7153abcae2 := __obf_03e9e17f4239cd40.Translate(__obf_eb502d163673590c, __obf_f55af572487ddde8)
	__obf_c9488cb5695961ff := f64.Aff3{__obf_b56c2b7153abcae2.XX, __obf_b56c2b7153abcae2.XY, __obf_b56c2b7153abcae2.X0, __obf_b56c2b7153abcae2.YX, __obf_b56c2b7153abcae2.YY, __obf_b56c2b7153abcae2.Y0}

	draw.BiLinear.Transform(__obf_7e8dc6da215ff393, __obf_c9488cb5695961ff, __obf_3f4810c67144854c, __obf_3f4810c67144854c.Bounds(), draw.Over, nil)
	__obf_3f4810c67144854c.NRGBA = __obf_7e8dc6da215ff393

	if __obf_feccdc0703f37768 {
		__obf_3ad7d05fa691070d := __obf_5c507cf47bf5267e - __obf_80c4c4a4d9299b93
		__obf_9aaefe68e0d2c39b := __obf_f03bf0f36c5833f9 - __obf_e93bb5827b72b466
		__obf_2ab04c1711998195 := (__obf_3ad7d05fa691070d / 2) + 1
		__obf_45d338c45fe3468d := (__obf_9aaefe68e0d2c39b / 2) + 1
		__obf_3f4810c67144854c.SubImage(image.Rect(__obf_2ab04c1711998195, __obf_45d338c45fe3468d, __obf_80c4c4a4d9299b93+__obf_2ab04c1711998195, __obf_e93bb5827b72b466+__obf_45d338c45fe3468d))
	}
}

// CropCircle crops a circular area
func (__obf_3f4810c67144854c *NRGBA) CropCircle(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_ac266a6cc46c4953 int) {
	__obf_6486fe380f7edf60 := __obf_3f4810c67144854c.Bounds()
	__obf_7582dbb5962003e0 := image.NewNRGBA(__obf_6486fe380f7edf60)
	for __obf_3aa55514da1d0d93 := __obf_6486fe380f7edf60.Min.Y; __obf_3aa55514da1d0d93 < __obf_6486fe380f7edf60.Max.Y; __obf_3aa55514da1d0d93++ {
		for __obf_e7e6d0305e77b960 := __obf_6486fe380f7edf60.Min.X; __obf_e7e6d0305e77b960 < __obf_6486fe380f7edf60.Max.X; __obf_e7e6d0305e77b960++ {
			__obf_afb07089cdbc8085 := math.Hypot(float64(__obf_e7e6d0305e77b960-__obf_c87d6c7ba4a30d82), float64(__obf_3aa55514da1d0d93-__obf_f8fead3d682d659f))
			if __obf_afb07089cdbc8085 <= float64(__obf_ac266a6cc46c4953) {
				__obf_7582dbb5962003e0.Set(__obf_e7e6d0305e77b960, __obf_3aa55514da1d0d93, color.White)
			} else {
				__obf_7582dbb5962003e0.Set(__obf_e7e6d0305e77b960, __obf_3aa55514da1d0d93, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_7582dbb5962003e0, __obf_7582dbb5962003e0.Bounds(), __obf_3f4810c67144854c, image.Point{X: 0, Y: 0}, __obf_7582dbb5962003e0, image.Point{}, draw.Over)
	__obf_3f4810c67144854c.NRGBA = __obf_7582dbb5962003e0
}

// CropScaleCircle scales and crops a circular area
func (__obf_3f4810c67144854c *NRGBA) CropScaleCircle(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_ac266a6cc46c4953 int, __obf_74c306278c875311 int) {
	__obf_6486fe380f7edf60 := __obf_3f4810c67144854c.Bounds()
	__obf_7582dbb5962003e0 := image.NewNRGBA(__obf_6486fe380f7edf60)

	for __obf_3aa55514da1d0d93 := __obf_6486fe380f7edf60.Min.Y; __obf_3aa55514da1d0d93 < __obf_6486fe380f7edf60.Max.Y; __obf_3aa55514da1d0d93++ {
		for __obf_e7e6d0305e77b960 := __obf_6486fe380f7edf60.Min.X; __obf_e7e6d0305e77b960 < __obf_6486fe380f7edf60.Max.X; __obf_e7e6d0305e77b960++ {
			__obf_afb07089cdbc8085 := math.Hypot(float64(__obf_e7e6d0305e77b960-__obf_c87d6c7ba4a30d82), float64(__obf_3aa55514da1d0d93-__obf_f8fead3d682d659f))
			if __obf_afb07089cdbc8085 <= float64(__obf_ac266a6cc46c4953) {
				__obf_7582dbb5962003e0.Set(__obf_e7e6d0305e77b960, __obf_3aa55514da1d0d93, color.White)
			} else {
				__obf_7582dbb5962003e0.Set(__obf_e7e6d0305e77b960, __obf_3aa55514da1d0d93, color.Transparent)
			}
		}
	}

	if __obf_74c306278c875311 > 0 {
		__obf_06ee349dfe38a9c2 := __obf_74c306278c875311 * 2
		__obf_47e78f92b9cf92e2 := image.NewNRGBA(image.Rect(0, 0, __obf_3f4810c67144854c.Bounds().Dx()-__obf_06ee349dfe38a9c2, __obf_3f4810c67144854c.Bounds().Dy()-__obf_06ee349dfe38a9c2))
		draw.BiLinear.Scale(__obf_47e78f92b9cf92e2, __obf_47e78f92b9cf92e2.Bounds(), __obf_7582dbb5962003e0, __obf_7582dbb5962003e0.Bounds(), draw.Over, nil)
		__obf_7582dbb5962003e0 = __obf_47e78f92b9cf92e2
	}

	draw.DrawMask(__obf_7582dbb5962003e0, __obf_7582dbb5962003e0.Bounds(), __obf_3f4810c67144854c, image.Point{X: __obf_74c306278c875311, Y: __obf_74c306278c875311}, __obf_7582dbb5962003e0, image.Point{}, draw.Over)
	__obf_3f4810c67144854c.NRGBA = __obf_7582dbb5962003e0
}

// Scale scales the canvas
func (__obf_3f4810c67144854c *NRGBA) Scale(__obf_74c306278c875311 int, __obf_1f6386c1ebeeb5f4, __obf_04036ea72a198976 bool) {
	__obf_7e8dc6da215ff393 := __obf_3f4810c67144854c.NRGBA
	if __obf_74c306278c875311 > 0 {
		__obf_06ee349dfe38a9c2 := __obf_74c306278c875311 * 2
		__obf_1987f94682eee223 := __obf_3f4810c67144854c.Bounds().Dx() - __obf_06ee349dfe38a9c2
		__obf_e31a5ebb0f34af57 := __obf_3f4810c67144854c.Bounds().Dy() - __obf_06ee349dfe38a9c2
		__obf_8fc14f152fea517b := image.NewNRGBA(image.Rect(0, 0, __obf_1987f94682eee223, __obf_e31a5ebb0f34af57))

		if !__obf_1f6386c1ebeeb5f4 {
			draw.BiLinear.Scale(__obf_8fc14f152fea517b, __obf_8fc14f152fea517b.Bounds(), __obf_3f4810c67144854c, __obf_3f4810c67144854c.Bounds(), draw.Over, nil)
		} else {
			__obf_16095dc37d26ace2 := util.CalcResizedRect(__obf_3f4810c67144854c.Bounds(), __obf_1987f94682eee223, __obf_e31a5ebb0f34af57, __obf_04036ea72a198976)
			draw.ApproxBiLinear.Scale(__obf_8fc14f152fea517b, __obf_16095dc37d26ace2.Bounds(), __obf_3f4810c67144854c, __obf_3f4810c67144854c.Bounds(), draw.Over, nil)
		}

		__obf_7e8dc6da215ff393 = __obf_8fc14f152fea517b
	}

	__obf_3f4810c67144854c.NRGBA = __obf_7e8dc6da215ff393
}
