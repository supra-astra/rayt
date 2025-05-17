package main

import (
	"fmt"
	"log"
	"math"
)

func hit_sphere(center *Point3, radius float64, r *Ray) float64 {
	oc := VecSub(&r.Orig, center)
	a := DotProduct(&r.Dir, &r.Dir)
	b := DotProduct(&r.Dir, oc)
	c := DotProduct(oc, oc) - (radius * radius)
	discriminant := b*b - (a * c)
	// return (discriminant >= 0)
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / a

}

func RayColor(r *Ray) *Color {
	t := hit_sphere(NewPoint3(0, 0, -1), 0.5, r)
	if t > 0.0 {
		point := r.At(t)
		normal := UnitVector(VecSub(&point, NewPoint3(0, 0, -1)))
		return ScalarMul(VecAdd(normal, NewVec3(1, 1, 1)), 0.5)
	}
	unitDirection := UnitVector(&r.Dir)
	a := 0.5 * (unitDirection.Y() + 1.0)
	return VecAdd(
		ScalarMul(NewColor(1.0, 1.0, 1.0), 1.0-a),
		ScalarMul(NewColor(0.5, 0.7, 1.0), a),
	)
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
	}

	//camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := NewPoint3(0, 0, 0)

	//calc the vectors across the horizontal and down the vertical viewport edges
	viewportU := NewVec3(viewportWidth, 0.0, 0.0)
	viewportV := NewVec3(0.0, -viewportHeight, 0.0)

	//calc the horizontal and vertical delta vectors from pixel to pixel
	pixelDeltaU := ScalarDiv(viewportU, float64(imageWidth))
	pixelDeltaV := ScalarDiv(viewportV, float64(imageHeight))

	//calc the location of the upper left pixel
	viewportUpperLeft := VecSub(VecSub(VecSub(cameraCenter, NewVec3(0.0, 0.0, focalLength)), ScalarDiv(viewportU, 2)), ScalarDiv(viewportV, 2))

	// [0,0,0]'s location
	pixel00Loc := VecAdd(viewportUpperLeft, ScalarMul(VecAdd(pixelDeltaU, pixelDeltaV), 0.5))

	//render
	fmt.Printf("P3\n %d %d\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		log.Printf("\rScanlines remaining: %d ", (imageHeight - j))
		for i := 0; i < imageWidth; i++ {
			pixelCenter := VecAdd(pixel00Loc, VecAdd(ScalarMul(pixelDeltaU, float64(i)), ScalarMul(pixelDeltaV, float64(j))))
			rayDirection := VecSub(pixelCenter, cameraCenter)

			r := NewRay(cameraCenter, rayDirection)

			pixelColor := RayColor(r)
			pixelColor.WriteColor()
		}
	}
	log.Printf("\rDone.                  \n")
}
