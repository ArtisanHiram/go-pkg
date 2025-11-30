package __obf_54406b1a1de84196

import (
	"errors"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	ImageEmptyErr       = errors.New("image is empty")
	ImageMissingDataErr = errors.New("missing image data")
)

const (
	DistortNone = iota
	DistortLevel1
	DistortLevel2
	DistortLevel3
	DistortLevel4
	DistortLevel5
)

type CaptchaType int

const (
	ClickCaptchat CaptchaType = iota
	RotateCaptchat
	MoveCaptchat
	DragCaptchat
)

const (
	QualityNone   = 100
	QualityLevel1 = 95
	QualityLevel2 = 85
	QualityLevel3 = 75
	QualityLevel4 = 65
	QualityLevel5 = 55
)

type CaptchaImage interface {
	Get() image.Image
	ToBytes() ([]byte, error)
	ToBase64() (string, error)
}

// Ensure JPEGImage and PNGImage implement CaptchaImage
var _ CaptchaImage = (*JPEGImage)(nil)
var _ CaptchaImage = (*PNGImage)(nil)

type CaptchaData interface {
	GetPrimary() CaptchaImage
	GetSecondary() CaptchaImage
	GetData() any
	Type() CaptchaType
	Verify(__obf_36c313afd5cadd12 any, __obf_ddf5266a6f4224d9 int) bool
}

type Captcha interface {
	Generate() (CaptchaData, error)
}

// RangeVal .
type Range struct {
	Min, Max int
}

// Size .
type Size struct {
	Width, Height int
}

// Point .
type Point struct {
	X, Y int
}

// AreaRect struct for defining a rectangular area
type AreaRect struct {
	MinX, MaxX, MinY, MaxY int
}

// MakeAreaRect creates an area rectangle
func MakeAreaRect(__obf_e8cf2dd52fedfa9e, __obf_81e79046ae17fba8, __obf_ecdc860838e3baf1, __obf_3093f2ae86219c18 int) *AreaRect {
	return &AreaRect{
		MinX: __obf_e8cf2dd52fedfa9e,
		MaxX: __obf_ecdc860838e3baf1,
		MinY: __obf_81e79046ae17fba8,
		MaxY: __obf_3093f2ae86219c18,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_daaeaddeff0e74f7, __obf_2366a8ab85b2ae47 int) *PositionRect {
	return &PositionRect{
		X:      __obf_48afb3005cd4a35c,
		Y:      __obf_33a1f511b09ac2af,
		Height: __obf_daaeaddeff0e74f7,
		Width:  __obf_2366a8ab85b2ae47,
	}
}

// DrawStringParams struct for string drawing parameters
type DrawStringParam struct {
	Color       color.Color
	FontSize    int
	Size        int
	FontDPI     int
	Text        string
	Font        *opentype.Font
	FontHinting font.Hinting
}

// Dot represents a single point (character or shape) in the captcha
type Dot struct {
	Index          int         `json:"index"`
	X              int         `json:"x"`
	Y              int         `json:"y"`
	Size           int         `json:"size"`
	Angle          int         `json:"angle" msgpack:"-"`
	Text           string      `json:"text" msgpack:"-"`
	Shape          image.Image `json:"shape" msgpack:"-"`
	FontSize       int         `json:"font-size" msgpack:"-"`
	PrimaryColor   string      `json:"primary-color" msgpack:"-"`
	SecondaryColor string      `json:"secondary-color" msgpack:"-"`
}

// Block defines the block data for the slide CAPTCHA
type Block struct {
	X int `json:"x"`
	Y int `json:"y"`
	// Display x,y
	DX     int         `json:"dx"`
	DY     int         `json:"dy"`
	Width  int         `json:"width" msgpack:"-"`
	Height int         `json:"height" msgpack:"-"`
	Angle  int         `json:"angle" msgpack:"-"`
	Shape  image.Image `json:"shape" msgpack:"-"`
}

// Mode defines the slide CAPTCHA mode
type SlideType int

const (
	MoveSlide SlideType = iota // Move mode - slide to left
	DragSlide                  // Drag mode - drag to any direction
)

// DeadZoneDirectionType defines the dead zone direction
type DeadZoneDirectionType int

const (
	DeadZoneDirectionTypeLeft DeadZoneDirectionType = iota
	DeadZoneDirectionTypeRight
	DeadZoneDirectionTypeTop
	DeadZoneDirectionTypeBottom
)

// Primary defines the main image configuration
type SlidePrimary struct {
	Size  Size    // Image size
	Alpha float32 // Image alpha
}

// GraphConfig defines the graph configuration
type SlideSecondary struct {
	CountRange           Range                   // Number of graphs to generate
	SizeRange            Range                   // Graph size range
	AnglePosRange        []Range                 // Angle position range
	DeadZoneDirections   []DeadZoneDirectionType // Dead zone directions
	EnableVerticalRandom bool                    // Enable vertical random positioning
}

// Options defines the configuration options for the captcha
type SlideOption struct {
	Primary   SlidePrimary
	Secondary SlideSecondary
	Type      SlideType
}

// Primary defines the main image configuration
type RotatePrimary struct {
	Size          int     // Image size
	Alpha         float32 // Image alpha
	AnglePosRange []Range // Angle position range
}

// Secondary defines the thumbnail image configuration
type RotateSecondary struct {
	Alpha     float32 // Thumbnail alpha
	SizeRange []int   // Thumbnail size range
}

// Options defines the configuration options for the captcha
type RotateOption struct {
	Primary   RotatePrimary
	Secondary RotateSecondary
}

type ClickPrimary struct {
	Size          Size    // Primary.Size
	LenRange      Range   // rangeLen
	AnglePosRange []Range // rangeAnglePos
	SizeRange     Range   // rangeSize
	Alpha         float32 // imageAlpha
	DotPadding    int
}

type ClickSecondary struct {
	Size           Size  // Secondary.Size
	VerifyLenRange Range // rangeVerifyLen
	SizeRange      Range // rangeThumbSize
	BgDistort      int   // thumbBgDistort
	BgCircles      int   // thumbBgCirclesNum
	BgSlimLines    int   // thumbBgSlimLineNum
	NonDeformAble  bool  // isThumbNonDeformAbility
	DisturbAlpha   float32
	DotPadding     int
}

type ClickOption struct {
	Primary     ClickPrimary
	Secondary   ClickSecondary
	FontDPI     int
	FontHinting font.Hinting
	ShowShadow  bool
	ShadowPoint Point
	UseRGBA     bool
}
