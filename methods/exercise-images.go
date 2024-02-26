package main

import (
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{255, 255},
	}
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{0, 255, 255, 255}
}
