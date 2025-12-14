package main

import (
	"fmt"
	"math"
)

// Vec3 is the fundamental structure for colors (RGB) and points/directions (XYZ).
type Vec3 struct {
	X, Y, Z float64
}

// Add returns the vector sum of u and v.
func (u Vec3) Add(v Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

// Subtract returns the vector difference of u and v.
func (u Vec3) Subtract(v Vec3) Vec3 {
	return Vec3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

// MultiplyScalar returns the vector scaled by a factor t.
func (u Vec3) MultiplyScalar(t float64) Vec3 {
	return Vec3{u.X * t, u.Y * t, u.Z * t}
}

// DivideScalar returns the vector divided by a factor t.
func (u Vec3) DivideScalar(t float64) Vec3 {
	return u.MultiplyScalar(1.0 / t)
}

// Negate returns the negative of the vector.
func (u Vec3) Negate() Vec3 {
	return Vec3{-u.X, -u.Y, -u.Z}
}

// LengthSquared returns the squared length of the vector.
func (u Vec3) LengthSquared() float64 {
	return u.X*u.X + u.Y*u.Y + u.Z*u.Z
}

// Length returns the length (magnitude) of the vector.
func (u Vec3) Length() float64 {
	return math.Sqrt(u.LengthSquared())
}

// String returns a string representation of the vector.
func (u Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", u.X, u.Y, u.Z)
}

// Utility functions (for consistency with Python's utils)

// DotProduct returns the dot product of u and v.
func DotProduct(u, v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// UnitVector returns the unit vector of v.
func UnitVector(v Vec3) Vec3 {
	return v.DivideScalar(v.Length())
}

// Point3 is an alias for Vec3, representing a 3D point.
type Point3 = Vec3

// Color is an alias for Vec3, representing an RGB color.
type Color = Vec3
