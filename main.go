package main

import (
	"fmt"
	"math"
	"os"
)

const (
	// These constants match the values hardcoded in your Python main.py,
	// overriding the values from the Python constants at the top.
	IMAGE_WIDTH  = 300
	IMAGE_HEIGHT = 300
)

// RayColor determines the color of a ray. This function is simplified
// to match the single-sphere logic in your main.py.
func RayColor(r Ray) Color {
	// World definition: A single sphere at (0, 0, -1) with radius 0.5
	center := Point3{0, 0, -1}
	radius := 0.5
	sphere := Sphere{Center: center, Radius: radius}

	// Ray-Sphere intersection logic from hit_sphere(center, radius, r)

	// oc = center - r.origin() -> oc = r.origin() - center
	oc := r.Origin.Subtract(center)
	a := r.Direction.LengthSquared()
	h := DotProduct(r.Direction, oc)
	c := oc.LengthSquared() - (radius * radius)
	discriminant := (h * h) - (a * c)

	t := -1.0
	if discriminant >= 0 {
		// Only take the closer root, similar to your simplified hit_sphere
		t = (h - math.Sqrt(discriminant)) / (a)
	}

	// Check if the hit is valid (t > 0)
	if t > 0 {
		// Shading based on the surface normal
		N := UnitVector(r.At(t).Subtract(center))
		// Color(N.X()+1, N.Y()+1, N.Z()+1) * 0.5
		return Color{N.X + 1, N.Y + 1, N.Z + 1}.MultiplyScalar(0.5)
	}
	//

	// Background color (sky gradient)
	unitDirection := UnitVector(r.Direction)
	aSky := (unitDirection.Y + 1.0) * 0.5

	// return Color(1.0, 1.0, 1.0) * (1.0 - a) + Color(0.5, 0.7, 1.0) * a
	white := Color{1.0, 1.0, 1.0}
	blue := Color{0.5, 0.7, 1.0}

	// (1.0 - aSky) * white + aSky * blue
	return white.MultiplyScalar(1.0 - aSky).Add(blue.MultiplyScalar(aSky))
}

func main() {
	// --- Image Properties (using hardcoded constants from main.py) ---
	// Note: Your main.py calculates a custom width/height but then uses
	// hardcoded constants IMAGE_WIDTH=300 and IMAGE_HEIGHT=300 for the loops.
	// We will follow the loop constraints, but use the correct camera setup.

	// Image calculated properties (used for camera setup)
	const aspectRatio = 16.0 / 9.0
	const imageWidthCalc = 400
	imageHeightCalc := float64(imageWidthCalc) / aspectRatio
	if imageHeightCalc < 1 {
		imageHeightCalc = 1
	}

	// --- Camera Properties ---
	const focalLength = 1.0
	const viewportHeight = 2.0
	viewportWidth := viewportHeight * (float64(imageWidthCalc) / imageHeightCalc)
	cameraCenter := Vec3{0.0, 0.0, 0.0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges
	viewportU := Vec3{viewportWidth, 0, 0}
	viewportV := Vec3{0, -viewportHeight, 0} // Y points down

	// Calculate the horizontal and vertical delta vectors from pixel to pixel
	// NOTE: We use the loop constants IMAGE_WIDTH and IMAGE_HEIGHT here
	pixelDeltaU := viewportU.DivideScalar(IMAGE_WIDTH)
	pixelDeltaV := viewportV.DivideScalar(IMAGE_HEIGHT)

	// Calculate the location of the upper left pixel
	// viewport_upper_left = camera_center - Vec3(0, 0, focal_length) - viewport_u / 2 - viewport_v / 2
	viewportUpperLeft := cameraCenter.Subtract(Vec3{0, 0, focalLength}).
		Subtract(viewportU.DivideScalar(2)).
		Subtract(viewportV.DivideScalar(2))

	// pixel00_location = viewport_upper_left + (pixel_delta_u + pixel_delta_v) * 0.5
	pixel00Location := viewportUpperLeft.
		Add(pixelDeltaU.Add(pixelDeltaV).MultiplyScalar(0.5))

	// --- Render ---
	// Write the PPM header
	fmt.Fprintf(os.Stdout, "P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	// Main rendering loop
	for j := 0; j < IMAGE_HEIGHT; j++ {
		// Progress indicator (writing to stderr)
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", IMAGE_HEIGHT-j)

		for i := 0; i < IMAGE_WIDTH; i++ {
			// pixel_center = pixel00_location + (pixel_delta_u * i) + (pixel_delta_v * j)
			pixelCenter := pixel00Location.
				Add(pixelDeltaU.MultiplyScalar(float64(i))).
				Add(pixelDeltaV.MultiplyScalar(float64(j)))

			// ray_direction = pixel_center - camera_center
			rayDirection := pixelCenter.Subtract(cameraCenter)

			// r = Ray(camera_center, ray_direction)
			r := Ray{Origin: cameraCenter, Direction: rayDirection}

			pixelColor := RayColor(r)
			WriteColor(pixelColor)
		}
	}

	fmt.Fprintln(os.Stderr, "\nDone.")
}
