package main

import "math"

// Sphere is a basic Hittable object.
type Sphere struct {
	Center Point3
	Radius float64
}

// Hit implements the Hittable interface for a Sphere.
// It uses the quadratic formula solution for the ray-sphere intersection.
// Ray equation: P(t) = O + t*D
// Sphere equation: |P - C|^2 = r^2
// Substituting P: |(O + t*D) - C|^2 = r^2
// This expands to a quadratic equation in t: a*t^2 + 2*h*t + c = 0
// where:
// a = D . D (Direction.LengthSquared())
// h = D . (O - C)
// c = |O - C|^2 - r^2
func (s Sphere) Hit(r Ray, rayTMin, rayTMax float64) (HitRecord, bool) {
	oc := r.Origin.Subtract(s.Center) // O - C
	a := r.Direction.LengthSquared()
	h := DotProduct(r.Direction, oc)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return HitRecord{}, false // No real roots, no hit
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range [rayTMin, rayTMax]

	// Try the first root (t0)
	root := (h - sqrtd) / a
	if root <= rayTMin || rayTMax <= root {
		// If t0 is outside the range, try the second root (t1)
		root = (h + sqrtd) / a
		if root <= rayTMin || rayTMax <= root {
			return HitRecord{}, false // Both roots are outside the range
		}
	}

	// We found an acceptable hit at parameter 'root' (t)
	rec := HitRecord{}
	rec.T = root
	rec.P = r.At(rec.T)

	// Calculate the normal and set the face normal
	outwardNormal := rec.P.Subtract(s.Center).DivideScalar(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)

	return rec, true
}
