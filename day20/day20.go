package main

import (
	"fmt"
	"image"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
)

func part1(in string) (rc int) {
	tiles := parse(in)
	keys := tiles.keys()
	sort.Ints(keys)
	rc = 1
	key := keys[0]
	fmt.Println("all orientations of tile ", key)
	for _, q := range tiles[key].orientations() {
		fmt.Println(q)
	}
	return
}

type tiles map[int]tile

func (s tiles) keys() (r []int) {
	for k, _ := range s {
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
	return
}

type tile struct {
	m map[image.Point]string
	n int
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

func (t tile) north() (r string) {
	for i := 0; i < t.n; i++ {
		r += t.m[image.Pt(i, 0)]
	}
	return
}

func (t tile) south() (r string) {
	for i := 0; i < t.n; i++ {
		r += t.m[image.Pt(i, t.n-1)]
	}
	return
}

func (t tile) east() (r string) {
	for i := 0; i < t.n; i++ {
		r += t.m[image.Pt(t.n-1, i)]
	}
	return
}

func (t tile) west() (r string) {
	for i := 0; i < t.n; i++ {
		r += t.m[image.Pt(0, i)]
	}
	return
}

// Flipping a tile produces its mirror image tile.
func (t tile) flip() (q tile) {
	q = tile{m: map[image.Point]string{}, n: t.n}
	for pt, s := range t.m {
		q.m[image.Pt(t.n-1-pt.X, pt.Y)] = s
	}
	return
}

// Rotations returns the tile rotated 90, 180, and 270 degrees.
func (t tile) rotations() (r []tile) {
	r = make([]tile, 3)
	r[0] = tile{m: map[image.Point]string{}, n: t.n}
	r[1] = tile{m: map[image.Point]string{}, n: t.n}
	r[2] = tile{m: map[image.Point]string{}, n: t.n}
	for pt, s := range t.m {
		r[0].m[image.Pt(t.n-1-pt.Y, pt.X)] = s
		r[1].m[image.Pt(t.n-1-pt.X, t.n-1-pt.Y)] = s
		r[2].m[image.Pt(pt.Y, t.n-1-pt.X)] = s
	}
	return
}

// Returns all possible orientations of the tile.
func (t tile) orientations() []tile {
	return append(append([]tile{t, t.flip()}, t.rotations()...), t.flip().rotations()...)
}

func main() {
	fmt.Println(part1(inputs.Day20()))
	// fmt.Println(part2(Input))
}
