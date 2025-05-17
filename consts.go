//some common constants

package main

import "math"

// pos infinity
var Infinity float64 = math.Inf(1)

// pi
var Pi float64 = math.Pi

//some utility functions

func DegreesToRadians(degrees float64) float64 {
	return (degrees * Pi) / 180.0
}


