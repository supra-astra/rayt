package main

import (
	"fmt"
	"log"
)

func RayColor(r *Ray) *Color {
	return &Color{
		E: [3]float64{0.0, 0.0, 0.0},
	}
}

func main() {
	//image
	imageWidth := 400
	//image_height := 256

	aspectRatio := 16.0 / 9.0

	//calc the image height , and ensure that it's atleast 1
	imageHeight := int(imageWidth / int(aspectRatio))
	if imageHeight < 1 {
		imageHeight = 1
	} else {
		imageHeight = imageHeight
	}

	//camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := Point3{[3]float64{0.0, 0.0, 0.0}}

	//calc the vectors across the horizontal and down the vertical viewport edges
	viewportU := NewVec3(viewportWidth, 0.0, 0.0)
	viewportV := NewVec3(0, -viewportHeight, 0.0)

	//calc the horizontal and vertical delta vectors from pixel to pixel
	pixelDeltaU := ScalarDiv(viewportU, float64(imageWidth))
	pixelDeltaV := ScalarDiv(viewportV, float64(imageHeight))

	//calc the location of the upper left pixel
	viewPortUV := VecAdd(ScalarDiv(viewportU, 2), ScalarDiv(viewportV, 2))
	cameraCenterOffset := VecSub(&cameraCenter, NewVec3(0, 0, focalLength))
	viewportUpperLeft := VecSub(cameraCenterOffset, viewPortUV)

	// [0,0,0]'s location
	pixel00Loc := VecAdd(viewportUpperLeft, ScalarMul(VecAdd(pixelDeltaU, pixelDeltaV), 0.5))

	//render
	fmt.Printf("P3\n %d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d ", (imageHeight - j))
		for i := 0; i < imageWidth; i++ {
			pixelCenter := VecAdd(pixel00Loc, VecAdd(ScalarMul(pixelDeltaU, float64(i)), ScalarMul(pixelDeltaV, float64(j))))
			rayDirection := VecSub(pixelCenter, &cameraCenter)

			r := Ray{cameraCenter, *rayDirection}

			pixelColor := RayColor(&r)
			pixelColor.WriteColor()
		}
	}
	log.Printf("\rDone.                  \n")
}
