package __obf_a334e8eb146672ab

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	"golang.org/x/image/font"
	"image"
)

type DotType int

const (
	Text  DotType = iota // Text content
	Shape                // Shape/Image content
)

// DrawImageParams defines the parameters for drawing images
type DrawImageParams struct {
	Width             int
	Height            int
	Background        image.Image
	BackgroundDistort int
	BgCircles         int
	BgSlimLines       int
	Alpha             float32
	FontHinting       font.Hinting
	Dots              []*types.Dot
	ShowShadow        bool
	ShadowPoint       types.Point
	ThumbDisturbAlpha float32
}

type SetOption func(*types.ClickOption)
