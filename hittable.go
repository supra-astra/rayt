package main

type HitRecord struct {
	P      Point3
	Normal Vec3
	T      float64
}

// using interfaces in go to complement the virtual functions
type Hittable interface {
	Hit(r *Ray, tMin, tMax float64) (bool, HitRecord)
}
