package main

type HitRecord struct {
	P         Point3
	Normal    Vec3
	T         float64
	FrontFace bool
}

// using interfaces in go to complement the virtual functions
type Hittable interface {
	Hit(r *Ray, tMin, tMax float64) (bool, HitRecord)
}

func (h *HitRecord) SetFaceNormal(r *Ray, outwardNormal *Vec3) {
	// set the hit record normal vector
	// outwardNormal is assumed to have unit length
	frontFace := DotProduct(r.Direction(), outwardNormal) < 0
	if frontFace {
		h.Normal = *outwardNormal
	} else {
		h.Normal = *ScalarMul(outwardNormal, -1)
	}
}
