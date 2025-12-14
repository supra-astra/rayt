package main

// HitRecord stores information about a ray-object intersection.
type HitRecord struct {
	P         Point3  // The intersection point
	Normal    Vec3    // The surface normal at the intersection point
	T         float64 // The ray parameter (t) at the intersection
	FrontFace bool    // True if the ray is outside the object
}

// SetFaceNormal determines the correct normal orientation.
// The outward_normal is assumed to be a unit vector pointing out from the object.
// We set the record normal to point against the ray if it's a front face.
func (rec *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	// A front face is when the ray direction and the outward normal are pointing
	// in opposite directions (Dot Product < 0).
	rec.FrontFace = DotProduct(r.Direction, outwardNormal) < 0.0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Negate()
	}
}

// Hittable is an interface for any object a ray can potentially hit.
type Hittable interface {
	Hit(r Ray, rayTMin, rayTMax float64) (HitRecord, bool)
}
