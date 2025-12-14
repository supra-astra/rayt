package main

// Ray represents a line in 3D space defined by an origin and a direction.
type Ray struct {
	Origin    Point3
	Direction Vec3
}

// At returns the point along the ray at parameter t: P(t) = A + t*b
func (r Ray) At(t float64) Point3 {
	// P(t) = Origin + Direction * t
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}
