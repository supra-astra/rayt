package main

type Point3 = Vec3

type Ray struct {
	Dir  Vec3
	Orig Point3
}

func NewRay(dir *Vec3, orig *Point3) *Ray {
	return &Ray{
		Dir:  *dir,
		Orig: *orig,
	}
}

func (r Ray) Origin() Point3 {
	return r.Orig
}

func (r Ray) Direction() Vec3 {
	return r.Dir
}

func (r Ray) At(t float64) Point3 {
	return *VecAdd(&r.Orig, &*ScalarMul(&r.Dir, t))
}
