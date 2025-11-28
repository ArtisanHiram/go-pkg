package __obf_07af88af2ce1a9a4

import (
	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
	"golang.org/x/image/font"
	"image"
)

type DotType int

const (
	Text  DotType = iota // Text content
	Shape                // Shape/Image content
)

// Dot represents a single point (character or shape) in the captcha
type Dot struct {
	Index    int `json:"index"`
	X        int `json:"x"`
	Y        int `json:"y"`
	FontSize int `json:"font-size"`
	// Width    int `json:"width"`
	// Height   int `json:"height"`
	Size int `json:"size"`
	// Type     DotType `json:"type"`
	Angle          int         `json:"angle"`
	PrimaryColor   string      `json:"primary-color"`
	SecondaryColor string      `json:"secondary-color"`
	Text           string      `json:"-" msgpack:"-"`
	Shape          image.Image `json:"-" msgpack:"-"`
	// Content - only one of these should be set based on Type
	// Text  string      `json:"text,omitempty"`
	// Shape image.Image `json:"shape,omitempty"`
	// FontDPI          int         `json:"font-dpi,omitempty"`
	// UseOriginalColor bool `json:"use-original-color,omitempty"`
}

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
	Dots              []*Dot
	ShowShadow        bool
	ShadowPoint       core.Point
	ThumbDisturbAlpha float32
}

type Primary struct {
	Size          core.Size    // Primary.Size
	LenRange      core.Range   // rangeLen
	AnglePosRange []core.Range // rangeAnglePos
	SizeRange     core.Range   // rangeSize
	Alpha         float32      // imageAlpha
	DotPadding    int
}

type Secondary struct {
	Size           core.Size  // Secondary.Size
	VerifyLenRange core.Range // rangeVerifyLen
	SizeRange      core.Range // rangeThumbSize
	BgDistort      int        // thumbBgDistort
	BgCircles      int        // thumbBgCirclesNum
	BgSlimLines    int        // thumbBgSlimLineNum
	NonDeformAble  bool       // isThumbNonDeformAbility
	DisturbAlpha   float32
	DotPadding     int
}

// Options defines the configuration options for the captcha
type Options struct {
	FontDPI     int
	FontHinting font.Hinting
	Primary
	Secondary
	ShowShadow  bool
	ShadowPoint core.Point
	UseRGBA     bool
}

type Option func(*Options)
