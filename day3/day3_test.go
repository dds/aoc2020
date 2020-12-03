package main

import (
	"fmt"
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	type test struct {
		p, q   image.Point
		expect []image.Point
	}

	tests := []test{
		test{
			p: image.Point{0, 0}, q: image.Point{1, 0},
			expect: []image.Point{image.Point{0, 0}, image.Point{1, 0}},
		},
		test{
			p: image.Point{0, 0}, q: image.Point{0, 1},
			expect: []image.Point{image.Point{0, 0}, image.Point{0, 1}},
		},
		test{
			p: image.Point{0, 0}, q: image.Point{-1, 0},
			expect: []image.Point{image.Point{0, 0}, image.Point{-1, 0}},
		},
		test{
			p: image.Point{0, 0}, q: image.Point{0, -1},
			expect: []image.Point{image.Point{0, 0}, image.Point{0, -1}},
		},
		test{
			p: image.Point{0, 0}, q: image.Point{3, 7},
			expect: []image.Point{
				image.Point{0, 0},
				image.Point{0, 1},
				image.Point{0, 2},
				image.Point{0, 3},
				image.Point{0, 4},
				image.Point{0, 5},
				image.Point{1, 5},
				image.Point{1, 6},
				image.Point{2, 6},
				image.Point{2, 7},
				image.Point{3, 7},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, path(test.p, test.q))
		})
	}
}
