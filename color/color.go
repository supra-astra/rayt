package color

import (
	"fmt"

	"github.com/supra-astra/rayt/vec3"
)

type Color = vec3.Vec3

func WriteColor(pixelColor Color) {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	//translate [0,1] component values to the byte range [0,255]
	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	//write the pixel color
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
