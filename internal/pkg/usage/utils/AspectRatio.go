package utils

import "image"

const (
	AspectRatio float32 = 16.0 / 9.0
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
