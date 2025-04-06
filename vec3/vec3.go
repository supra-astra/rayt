package vec3

import (
	"fmt"
	"math"
)

//classes for storing geometric vectors and colors.
// for our purposes the three corrdinates suffice.
// using this same class vec3 for colors, locations,directions,
//offsets ,whatever.

type Vec3 struct {
	E [3]float64
}

func NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{
		E: [3]float64{e0, e1, e2},
	}
}

func DefaultVec3() *Vec3 {
	return &Vec3{
		E: [3]float64{0.0, 0.0, 0.0},
	}
}

// get the X,Y,Z coordinates
func (v Vec3) X() float64 {
	return v.E[0]
}

func (v Vec3) Y() float64 {
	return v.E[1]
}

func (v Vec3) Z() float64 {
	return v.E[2]
}

// operator-()
func (v Vec3) Negate() Vec3 {
	return Vec3{
		E: [3]float64{-v.E[0], -v.E[1], -v.E[2]},
	}
}

// operator[]
func (v Vec3) At(i int) float64 {
	return v.E[i]
}

// set / assign
func (v *Vec3) Set(i int, val float64) {
	v.E[i] = val
}

// operator+=
func (v *Vec3) AddAssign(other Vec3) {
	v.E[0] += other.E[0]
	v.E[1] += other.E[1]
	v.E[2] += other.E[2]
}

// operator*=
func (v *Vec3) MulAssign(t float64) {
	v.E[0] *= t
	v.E[1] *= t
	v.E[2] *= t
}

// operator /=
func (v *Vec3) DivAssign(t float64) {
	inv := 1.0 / t
	v.MulAssign(inv)
}

// length
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}

//vector utility functions

func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.E[0], v.E[1], v.E[2])
}

func VecAdd(v Vec3, u Vec3) Vec3 {
	return Vec3{
		E: [3]float64{v.E[0] + u.E[0], u.E[1] + v.E[1], u.E[2] + v.E[2]},
	}
}

func VecSub(u Vec3, v Vec3) Vec3 {
	return Vec3{
		E: [3]float64{u.E[0] - v.E[0], u.E[1] - v.E[1], u.E[2] - v.E[2]},
	}
}

func VecMul(u Vec3, v Vec3) Vec3 {
	return Vec3{
		E: [3]float64{u.E[0] * v.E[0], u.E[1] * v.E[1], u.E[2] * v.E[2]},
	}
}

func ScaleUp(t float64, v Vec3) Vec3 {
	return Vec3{
		E: [3]float64{t * v.E[0], t * v.E[1], t * v.E[2]},
	}
}

func ScaleDown(t float64, v Vec3) Vec3 {
	return ScaleUp(1/t, v)
}

func DotProduct(u Vec3, v Vec3) float64 {
	return u.E[0]*v.E[0] + u.E[1]*v.E[1] + u.E[2]*v.E[2]
}

func CrossProduct(u Vec3, v Vec3) Vec3 {
	return Vec3{
		E: [3]float64{u.E[1]*v.E[2] - u.E[2]*v.E[1],
			u.E[2]*v.E[0] - u.E[0]*v.E[2],
			u.E[0]*v.E[1] - u.E[1]*v.E[0]},
	}
}

func (v Vec3) UnitVector() Vec3 {
	return ScaleDown(v.Length(), v)
}
