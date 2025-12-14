package main

import (
	"fmt"
	"math"
	"os"
)

// WriteColor writes the color components to stdout in the PPM format.
// It translates the [0, 1] color components to the byte range [0, 255].
func WriteColor(pixelColor Color) {
	r := pixelColor.X
	g := pixelColor.Y
	b := pixelColor.Z

	// Translate the [0,1] component values to the byte range [0,255]

	// Use math.Floor to ensure it stays in the range [0, 255] and handles the conversion to int.
	rbyte := int(math.Floor(256 * r))
	gbyte := int(math.Floor(256 * g))
	bbyte := int(math.Floor(256 * b))

	// Write out the pixel color components
	// Using fmt.Fprintf(os.Stdout, ...) is a standard way to write to standard output.
	// We'll trust Go's buffering is efficient enough, but for extreme performance in
	// a real application, you'd use a bufio.Writer.
	fmt.Fprintf(os.Stdout, "%d %d %d\n", rbyte, gbyte, bbyte)
}
