package main

import "math"
import "fmt"

type Vec3 struct {
	E [3]float64
}

func NewVec3(l, b, h float64) *Vec3 {
	return &Vec3{
		E: [3]float64{l, b, h},
	}
}

func (v Vec3) X() float64 {
	return v.E[0]
}

func (v Vec3) Y() float64 {
	return v.E[1]
}

func (v Vec3) Z() float64 {
	return v.E[2]
}

func (v Vec3) Complement() *Vec3 {
	return NewVec3(-1*v.E[0], -1*v.E[1], -1*v.E[2])
}

func (v Vec3) ItemAt(i int) float64 {
	return v.E[i]
}

func PlusEquals(v *Vec3, u *Vec3) *Vec3 {
	v.E[0] += u.E[0]
	v.E[1] += u.E[1]
	v.E[2] += u.E[2]
	return v
}

// *= operator
func MultiplyConst(v *Vec3, t float64) *Vec3 {
	v.E[0] *= t
	v.E[1] *= t
	v.E[2] *= t
	return v
}

// /= operator
func DivideConst(v *Vec3, t float64) *Vec3 {
	return MultiplyConst(v, 1/t)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}

// vector utility functions
func (v Vec3) ToString() string {
	return fmt.Sprintf("%f %f %f", v.E[0], v.E[1], v.E[2])
}

func VecAdd(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(u.E[0]+v.E[0], u.E[1]+v.E[1], u.E[2]+v.E[2])
}

func VecSub(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(u.E[0]-v.E[0], u.E[1]-v.E[1], u.E[2]-v.E[2])
}

func VecMul(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(u.E[0]*v.E[0], u.E[1]*v.E[1], u.E[2]*v.E[2])
}

func ScalarMul(u *Vec3, t float64) *Vec3 {
	return NewVec3(t*u.E[0], t*u.E[1], t*u.E[2])
}

func ScalarDiv(u *Vec3, t float64) *Vec3 {
	return ScalarMul(u, 1/t)
}

func DotProduct(u *Vec3, v *Vec3) float64 {
	return u.E[0]*v.E[0] + u.E[1]*v.E[1] + u.E[2]*v.E[2]
}

func CrossProduct(u *Vec3, v *Vec3) *Vec3 {
	return NewVec3(u.E[1]*v.E[2]-u.E[2]*v.E[1],
		u.E[2]*v.E[0]-u.E[0]*v.E[2],
		u.E[0]*v.E[1]-u.E[1]*v.E[0])
}

func UnitVector(v *Vec3) *Vec3 {
	return DivideConst(v, v.Length())
}
