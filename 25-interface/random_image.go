package main

import (
	"image"
	"image/color"
	"math/rand"
)

type randomImage image.Point

// ColorModel returns the Image's color model.
func (r randomImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (r randomImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: r.X,
			Y: r.Y,
		},
	}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (r randomImage) At(x int, y int) color.Color {
	return color.RGBA{
		R: uint8(rand.Uint32() & (uint32(x) ^ uint32(y))),
		G: uint8(rand.Uint32() & (uint32(x) ^ uint32(y))),
		B: uint8(rand.Uint32() & (uint32(x) ^ uint32(y))),
		A: 255,
	}
}
