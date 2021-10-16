package base

import (
	"image"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	Title       string  = "Image Viewer"
	AspectRatio float32 = 16.0 / 9.0
)

var (
	BackgroundColor mgl32.Vec4 = mgl32.Vec4{0.55, 0.55, 0.55, 0.0} // background color
)

func UpdateAspectRatio(size image.Point) image.Point {
	var newSize image.Point
	originAspectRatio := float32(size.X) / float32(size.Y)

	if originAspectRatio >= AspectRatio {
		newSize.X = size.X
		newSize.Y = int(float32(size.X) / AspectRatio)
	} else {
		newSize.X = int(float32(size.Y) * AspectRatio)
		newSize.Y = size.Y
	}
	return newSize

}
