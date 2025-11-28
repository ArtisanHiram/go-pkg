package __obf_deb7e65d40f46bd0

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

type Captcha[T any] interface {
	GetPrimary() CaptchaImage
	GetSecondary() CaptchaImage
	GetData() T
	Type() CaptchaType
	Verify(__obf_7c619d340d1ba5c3 T, __obf_15100489c67cad53 int) bool
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
func MakeAreaRect(__obf_d55200013970c320, __obf_34df54cbb9c2c28d, __obf_96224a87efdb578a, __obf_65a377fe49b085e7 int) *AreaRect {
	return &AreaRect{
		MinX: __obf_d55200013970c320,
		MaxX: __obf_96224a87efdb578a,
		MinY: __obf_34df54cbb9c2c28d,
		MaxY: __obf_65a377fe49b085e7,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_7113bad0317e187f, __obf_3fbfb4aba2bc08f3 int) *PositionRect {
	return &PositionRect{
		X:      __obf_ce61f234cbe42808,
		Y:      __obf_7a88a48a3d7db1aa,
		Height: __obf_7113bad0317e187f,
		Width:  __obf_3fbfb4aba2bc08f3,
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
