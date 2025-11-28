package __obf_220a18b254a8ad8c

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

// Primary defines the main image configuration
type Primary struct {
	Size          int           // Image size
	Alpha         float32       // Image alpha
	AnglePosRange []types.Range // Angle position range
}

// Secondary defines the thumbnail image configuration
type Secondary struct {
	Alpha     float32 // Thumbnail alpha
	SizeRange []int   // Thumbnail size range
}

// Options defines the configuration options for the captcha
type Options struct {
	Primary
	Secondary
}

type Option func(*Options)
