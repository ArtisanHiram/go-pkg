package __obf_ae0bbdaffaa1f242

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	"image"
)

// DrawImageParams defines the parameters for drawing the main image
type DrawImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Alpha      float32
	DrawBlocks []*types.Block
}

// DrawTplImageParams defines the parameters for drawing the template image (tile)
type DrawTplImageParam struct {
	Width      int
	Height     int
	Background image.Image
	Mask       image.Image
	Alpha      float32
	Block      *types.Block
}

type SetOption func(*types.SlideOption)
