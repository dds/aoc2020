package lib

import (
	"fmt"
	"image"

	"github.com/yourbasic/graph"
	"github.com/yourbasic/graph/build"
)

func ExampleGrid() {
	// Build a 100x100 grid.
	n := 100
	d := Dim(n)

	g := build.Grid(n, n).AddCostFunc(func(v, w int) int64 {
		// Distance to three decimal places.
		return int64(1000 * Euclid(d.Point(v), d.Point(w)))
	})

	// Find a shortest path from (0,0) to (3,7).
	path, dist := graph.ShortestPath(g, d.Index(image.Point{0, 0}), d.Index(image.Point{3, 7}))
	fmt.Println("path:", path, "length:", float64(dist)/1000)
	for _, p := range path {
		fmt.Println(d.Point(p))
	}
	// Output:
	// path: [0 1 2 3 4 5 6 106 107 207 307] length: 10
	// (0,0)
	// (0,1)
	// (0,2)
	// (0,3)
	// (0,4)
	// (0,5)
	// (0,6)
	// (1,6)
	// (1,7)
	// (2,7)
	// (3,7)
}
