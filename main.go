package main

import "fmt"

func main() {
	//image
	image_width := 256
	image_height := 256

	//render
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)

	for j := range image_height {
		for i := range image_width {
			r := float64(i) / (float64(image_width) - 1)
			g := float64(j) / (float64(image_height) - 1)
			b := 0.0

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
