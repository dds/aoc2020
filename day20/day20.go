package main

import (
	"fmt"
	"image"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func main() {
	fmt.Println(part1(inputs.Day20()))
	fmt.Println(part2(inputs.Day20()))
}

// Solution based on http://chenlab.ece.cornell.edu/people/Andy/publications/Andy_files/Gallagher_cvpr2012_puzzleAssembly.pdf
func part1(in string) (rc int) {
	tiles := parse(in)
	keys := tiles.keys()
	sort.Ints(keys)

	rc = 1
	dim := lib.Dim(math.Sqrt(float64(len(tiles))))
	fmt.Println("Building square image with side length", dim)

	for a, m := range tiles.allComparisons() {
		// Corner pieces have two matches each with at least one edge.
		matches := 0
		for _, c := range m {
			if len(c.toMap()) >= 1 {
				matches++
			}
		}
		if matches == 2 {
			rc *= a
		}
	}
	return
}

func part2(in string) (rc int) {
	tiles := parse(in)
	keys := tiles.keys()
	sort.Ints(keys)

	rc = 1
	return
}

func (s tiles) allComparisons() (r map[int]map[int]comparison) {
	r = map[int]map[int]comparison{}
	for id, t := range s {
		r[id] = map[int]comparison{}
		for oid, q := range s {
			if id == oid {
				continue
			}
			r[id][oid] = t.compare(q)
		}
	}
	return
}

// Score how well tile q fits tile t by comparing all possible connections.
// Returns a map of int representing a possible orientation between p and q to a
// bool of whether that orientation fits.
func (t tile) compare(q tile) (r comparison) {
	r = comparison{}
	// The tile can fit to the right, left, top, or bottom. The other tile can
	// be arbitrarily rotated or flipped. For each border of this tile, for each
	// permutation of the other tile, compare if the corresponding border in the
	// other tile matches. Return a map of orientation possibility to bool of
	// match or no match.
	borderPairs := [][]int{
		[]int{1, 0},
		[]int{0, 1},
		[]int{3, 2},
		[]int{2, 3},
	}
	b := t.borders()
	i := 0
	for _, q := range q.orientations() {
		d := q.borders()
		for _, p := range borderPairs {
			r[i] = b[p[0]] == d[p[1]]
			i++
		}
	}
	return
}

type comparison [48]bool

func (c comparison) toMap() (r map[int]bool) {
	r = map[int]bool{}
	for i, c := range c {
		if !c {
			continue
		}
		r[i] = c
	}
	return
}

type tiles map[int]tile

func (s tiles) keys() (r []int) {
	for k := range s {
		r = append(r, k)
	}
	return
}

func parse(in string) (r tiles) {
	r = tiles{}
	for _, s := range strings.Split(in, "\n\n") {
		id, tile := parseTile(s)
		r[id] = tile
	}
	return
}

var tilehdr = regexp.MustCompile(`^Tile (\d+):`)

func parseTile(s string) (id int, t tile) {
	if !tilehdr.MatchString(s) {
		panic(fmt.Errorf("no tile header: %v", s))
	}
	ids := tilehdr.FindStringSubmatch(s)[1]
	id, _ = strconv.Atoi(ids)
	t = tile{m: map[image.Point]string{}}
	rows := strings.Split(s, "\n")[1:]
	n := 0
	for y, row := range rows {
		if len(row) == 0 {
			continue
		}
		for x, c := range row {
			t.m[image.Pt(x, y)] = string(c)
		}
		n = lib.Max(n, y)
	}
	t.n = n + 1
	t.id = id
	return
}

type tile struct {
	m     map[image.Point]string
	n, id int
	o     orientation
}

type orientation struct {
	flip int // 0, 1, 2: none, horizontal, vertical
	rot  int // 0, 1, 2, 3: none, 90, 180, 270
}

func (t tile) String() (r string) {
	for i := 0; i < t.n; i++ {
		for j := 0; j < t.n; j++ {
			r += t.m[image.Pt(j, i)]
		}
		r += "\n"
	}
	return
}

// Returns the north, south, east, and west borders as strings.
func (t tile) borders() (r []string) {
	r = make([]string, 4)
	for i := 0; i < t.n; i++ {
		r[0] += t.m[image.Pt(i, 0)]
		r[1] += t.m[image.Pt(i, t.n-1)]
		r[2] += t.m[image.Pt(t.n-1, i)]
		r[3] += t.m[image.Pt(0, i)]
	}
	return
}

// Flipping a tile produces flips it horizontally and verticaly.
func (t tile) flips() (r []tile) {
	r = make([]tile, 2)
	r[0] = tile{m: map[image.Point]string{}, n: t.n, id: t.id, o: orientation{flip: 1, rot: t.o.rot}}
	r[1] = tile{m: map[image.Point]string{}, n: t.n, id: t.id, o: orientation{flip: 2, rot: t.o.rot}}
	for pt, s := range t.m {
		r[0].m[image.Pt(t.n-1-pt.X, pt.Y)] = s
		r[1].m[image.Pt(pt.X, t.n-1-pt.Y)] = s
	}
	return
}

// Rotations returns the tile rotated 90, 180, and 270 degrees.
func (t tile) rotations() (r []tile) {
	r = make([]tile, 3)
	r[0] = tile{m: map[image.Point]string{}, n: t.n, id: t.id, o: orientation{flip: t.o.flip, rot: 1}}
	r[1] = tile{m: map[image.Point]string{}, n: t.n, id: t.id, o: orientation{flip: t.o.flip, rot: 2}}
	r[2] = tile{m: map[image.Point]string{}, n: t.n, id: t.id, o: orientation{flip: t.o.flip, rot: 3}}
	for pt, s := range t.m {
		r[0].m[image.Pt(t.n-1-pt.Y, pt.X)] = s
		r[1].m[image.Pt(t.n-1-pt.X, t.n-1-pt.Y)] = s
		r[2].m[image.Pt(pt.Y, t.n-1-pt.X)] = s
	}
	return
}

// Returns all possible orientations of the tile.
func (t tile) orientations() (r []tile) {
	r = append(append([]tile{t}, t.flips()...), t.rotations()...)
	for _, t := range t.flips() {
		r = append(r, t.rotations()...)
	}
	return
}
