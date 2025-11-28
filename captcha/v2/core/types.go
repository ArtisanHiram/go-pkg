package __obf_0e73be9e4c3ea3fd

import (
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// AreaRect struct for defining a rectangular area
type AreaRect struct {
	MinX, MaxX, MinY, MaxY int
}

// MakeAreaRect creates an area rectangle
func MakeAreaRect(__obf_d58117fac2c9257c, __obf_ca5e679331176b01, __obf_b78ac2d2e1cf7363, __obf_cdc91d61a35557f1 int) *AreaRect {
	return &AreaRect{
		MinX: __obf_d58117fac2c9257c,
		MaxX: __obf_b78ac2d2e1cf7363,
		MinY: __obf_ca5e679331176b01,
		MaxY: __obf_cdc91d61a35557f1,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_462a1d8e0f4b0353, __obf_6e1fff01dcc62820 int) *PositionRect {
	return &PositionRect{
		X:      __obf_fe086773828b5535,
		Y:      __obf_44ef04afc5c8241a,
		Height: __obf_462a1d8e0f4b0353,
		Width:  __obf_6e1fff01dcc62820,
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
