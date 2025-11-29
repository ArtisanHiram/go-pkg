package __obf_f075990603c6fd45

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

type Primary struct {
	Size          types.Size    // Primary.Size
	LenRange      types.Range   // rangeLen
	AnglePosRange []types.Range // rangeAnglePos
	SizeRange     types.Range   // rangeSize
	Alpha         float32       // imageAlpha
	DotPadding    int
}

type Secondary struct {
	Size           types.Size  // Secondary.Size
	VerifyLenRange types.Range // rangeVerifyLen
	SizeRange      types.Range // rangeThumbSize
	BgDistort      int         // thumbBgDistort
	BgCircles      int         // thumbBgCirclesNum
	BgSlimLines    int         // thumbBgSlimLineNum
	NonDeformAble  bool        // isThumbNonDeformAbility
	DisturbAlpha   float32
	DotPadding     int
}

// Options defines the configuration options for the captcha
type Options struct {
	Primary
	Secondary
	FontDPI     int
	FontHinting font.Hinting
	ShowShadow  bool
	ShadowPoint types.Point
	UseRGBA     bool
}

type Option func(*Options)
