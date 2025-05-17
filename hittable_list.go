package main

type HittableList struct {
	Objects []Hittable
}

// constructor for it
func NewHittableList() *HittableList {
	return &HittableList{
		Objects: make([]Hittable, 0),
	}
}

// utility methods for add and clear
func (hl *HittableList) Add(obj Hittable) {
	hl.Objects = append(hl.Objects, obj)
}

func (hl *HittableList) Clear() {
	hl.Objects = make([]Hittable, 0)
}

// add a hit method for hittablelist
func (hl *HittableList) Hit(r *Ray, tMin, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax

	var finalRec HitRecord

	for _, obj := range hl.Objects {
		if hit, rec := obj.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = rec.T
			finalRec = rec
		}
	}

	return hitAnything, finalRec
}
