package __obf_f99ff5685a4c2bb7

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	"image"
)

// DrawImageParams defines the parameters for drawing the main image
type DrawImageParams struct {
	Size       int
	Background image.Image
	Alpha      float32
}

// DrawCropCircleImageParams defines the parameters for drawing a cropped circle image
type DrawCropCircleImageParams struct {
	ScaleRatioSize int
	Angle          int
	Size           int
	Background     image.Image
	Alpha          float32
}

type SetOption func(*types.RotateOption)
