package lib

import (
	"image"
	"math"
)

type Dim int

func (n Dim) Index(p image.Point) int {
	return int(n)*p.X + p.Y
}
func (n Dim) Point(i int) image.Point {
	return image.Point{i / int(n), i % int(n)}
}

func Euclid(p, q image.Point) float64 {
	r := image.Rectangle{p, q}
	return math.Hypot(float64(r.Dx()), float64(r.Dy()))
}

func Taxi(p, q image.Point) float64 {
	r := image.Rectangle{p, q}.Canon()
	return float64(r.Dx() + r.Dy())
}

var Directions = []image.Point{
	image.Point{0, 1},
	image.Point{0, -1},
	image.Point{-1, 0},
	image.Point{1, 0},
}

// Returns the neighbors up, down, left, and right of P.
func Neighbors4(p image.Point) (r []image.Point) {
	for _, q := range Directions {
		r = append(r, p.Add(q))
	}
	return
}

var Diagnals = []image.Point{
	image.Point{1, 1},
	image.Point{1, -1},
	image.Point{-1, -1},
	image.Point{-1, 1},
}

// Returns the neighbors UDLR of P and the diagnals.
func Neighbors8(p image.Point) (r []image.Point) {
	r = Neighbors4(p)
	for _, q := range Diagnals {
		r = append(r, p.Add(q))
	}
	return
}

// Returns 1, 2, 3, or 4 if q is the up, down, left, or right neighbor of P, or
// 0 otherwise.
func UDLR(p, q image.Point) int {
	for i, u := range Directions {
		if q.Sub(p) == u {
			return i + 1
		}
	}
	return 0
}
