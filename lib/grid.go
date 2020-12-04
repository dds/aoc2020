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
