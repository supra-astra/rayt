package main

type Point3 = Vec3

func NewPoint3(l, b, h float64) *Point3 {
	return &Point3{
		E: [3]float64{l, b, h},
	}
}

type Ray struct {
	Dir  Vec3
	Orig Point3
}

func NewRay(orig *Point3, dir *Vec3) *Ray {
	return &Ray{
		Orig: *orig, Dir: *dir,
	}
}

func (r Ray) Origin() *Point3 {
	return &r.Orig
}

func (r Ray) Direction() *Vec3 {
	return &r.Dir
}

func (r Ray) At(t float64) Point3 {
	return *VecAdd(&r.Orig, ScalarMul(&r.Dir, t))
}
