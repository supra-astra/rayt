package main

import "math"

type Sphere struct {
	Center Point3
	Radius float64
}

func (s *Sphere) Hit(r *Ray, tMin, tMax float64) (bool, HitRecord) {
	oc := VecSub(&r.Orig, &s.Center)
	a := DotProduct(&r.Dir, &r.Dir)
	halfB := DotProduct(&r.Dir, oc)
	c := DotProduct(oc, oc) - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return false, HitRecord{}
	}

	sqrtD := math.Sqrt(discriminant)

	//nearest root within the acceptable range
	root := (-halfB - sqrtD) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrtD) / a
		if root < tMin || root > tMax {
			return false, HitRecord{}
		}
	}

	p := r.At(root)
	normal := ScalarDiv(VecSub(&p, &s.Center), s.Radius)

	rec := HitRecord{
		P:      p,
		Normal: *normal,
		T:      root,
	}

	outwardNormal := ScalarDiv(VecSub(&rec.P, &s.Center), s.Radius)
	rec.SetFaceNormal(r, outwardNormal)

	return true, rec
}
