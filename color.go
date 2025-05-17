package main

import (
	"fmt"
	"os"
)

//utility to write a single pixel's color out to the std output

type Color = Vec3

func NewColor(l, b, h float64) *Color {
	return &Color{
		E: [3]float64{l, b, h},
	}
}

func (pixelColor *Color) WriteColor() {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	//translate the [0,1] component values to the byte range [0,255]

	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	//write the result
	fmt.Fprintf(os.Stdout, "%d %d %d\n", rbyte, gbyte, bbyte)
}
