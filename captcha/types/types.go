package __obf_6dcb1d06bd949756

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
	Verify(__obf_21a09cfd5c7d4748 any, __obf_7fb86c457aa6cf23 int) bool
}

type Captcha interface {
	Generate() (CaptchaData, error)
}

// RangeVal .
type Range struct {
	Min, Max int
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
func MakeAreaRect(__obf_00f34bb6aaa3ea80, __obf_f1a0ff8eaf634306, __obf_cf2653fb9a7ca32d, __obf_db9ce805f0c741cf int) *AreaRect {
	return &AreaRect{
		MinX: __obf_00f34bb6aaa3ea80,
		MaxX: __obf_cf2653fb9a7ca32d,
		MinY: __obf_f1a0ff8eaf634306,
		MaxY: __obf_db9ce805f0c741cf,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_ac11522e37274a9b, __obf_1f037513843595bc int) *PositionRect {
	return &PositionRect{
		X:      __obf_57d99e8bd0a3444e,
		Y:      __obf_3376ebf4b35235fa,
		Height: __obf_ac11522e37274a9b,
		Width:  __obf_1f037513843595bc,
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
	Width  int
	Height int
	Alpha  float32 // Image alpha
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
	Width         int
	Height        int
	LenRange      Range   // rangeLen
	AnglePosRange []Range // rangeAnglePos
	SizeRange     Range   // rangeSize
	Alpha         float32 // imageAlpha
	DotPadding    int
}

type ClickSecondary struct {
	Width          int
	Height         int
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
