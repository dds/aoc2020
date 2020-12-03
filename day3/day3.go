package main

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
	"github.com/gdamore/tcell/v2"
)

var Input = lib.ParseInput(inputs.Day3(), func(s string) []string { return strings.Split(s, "") })

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
	shred(Input)
}

func part1(input [][]string) (rc int) {
	x := 0
	for _, row := range input {
		n := len(row)
		if row[x%n] == "#" {
			rc++
		}
		x += 3
	}
	return
}

func part2(input [][]string) (rc int) {
	counters := map[image.Point]int{
		image.Point{1, 1}: 0,
		image.Point{3, 1}: 0,
		image.Point{5, 1}: 0,
		image.Point{7, 1}: 0,
		image.Point{1, 2}: 0,
	}
	for y, row := range input {
		n := len(row)
		for p := range counters {
			if p.Y == 2 && y%2 == 1 {
				continue
			}
			if row[y/p.Y*p.X%n] == "#" {
				counters[p]++
			}
		}
	}
	rc = 1
	for _, c := range counters {
		rc *= c
	}
	return
}

const (
	trail     = '\\'
	skier     = '⛷'
	snowbdr   = '🏂'
	tree      = '🌲'
	mntn      = '🏔'
	snowman   = '⛄'
	xmastree  = '🎄'
	snowflake = '❄'
)

func background(s tcell.Screen) {
	colors := []tcell.Color{
		//  - ffffff 80%
		//  - fdfdfd 10%
		//  - fefefe 10%
		tcell.NewRGBColor(0xfe, 0xfe, 0xfe),
		tcell.NewRGBColor(0xfd, 0xfd, 0xfd),
		tcell.NewRGBColor(0xfb, 0xfb, 0xfb),
	}
	w, h := s.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := rand.Int() % 100
			idx := 2
			if p < 80 {
				idx--
			}
			if p < 90 {
				idx--
			}
			_, _, style, _ := s.GetContent(x, y)
			s.SetContent(x, y, ' ', nil, style.Background(colors[idx]))
		}
	}
}

var directions = []image.Point{
	image.Point{0, 1},
	image.Point{0, -1},
	image.Point{-1, 0},
	image.Point{1, 0},
}

func neighbors(p image.Point) (r []image.Point) {
	for _, q := range directions {
		r = append(r, p.Add(q))
	}
	return
}

func taxicab_distance(p, q image.Point) int {
	r := image.Rectangle{p, q}.Canon()
	return r.Dx() + r.Dy()
}

func euclidean_distance(p, q image.Point) float64 {
	r := image.Rectangle{p, q}
	return math.Hypot(float64(r.Dy()), float64(r.Dx()))
}

// BFS path from p to q.
func path(p, q image.Point, cost func(p, q image.Point) float64) (r []image.Point) {
	que := []image.Point{p}
	prevs := map[image.Point]*image.Point{}
	for len(que) > 0 {
		t := que[0]
		if t == q {
			break
		}
		que = que[1:]
		var next image.Point
		minScore := math.Inf(1)
		for _, u := range neighbors(t) {
			if prevs[u] != nil {
				continue
			}
			score := cost(u, q)
			if score < minScore {
				minScore = score
				next = u
			}
		}
		prevs[next] = &t
		que = append(que, next)
		r = append(r, next)
	}
	return
}

// Atar path from p to q.
func astar(p, q image.Point) (r []image.Point) {
	return path(p, q, func(p, q image.Point) float64 {
		if p == q {
			return 0
		}
		return float64(taxicab_distance(p, q)) + euclidean_distance(p, q)
	})
}

// Riders
type rider struct {
	image.Point
	tcell.Color
	glyph rune
}

var riders = map[image.Point]rider{
	image.Point{1, 1}: rider{
		glyph: skier,
		//   - ff8352 // orange
		Color: tcell.NewRGBColor(0xFF, 0x83, 0x52),
	},
	image.Point{3, 1}: rider{
		glyph: snowbdr,
		//   - ffb71c // gold
		Color: tcell.NewRGBColor(0xFF, 0xB7, 0x1C),
	},
	image.Point{5, 1}: rider{
		glyph: skier,
		//   - ff461c // red
		Color: tcell.NewRGBColor(0xFF, 0x46, 0x1c),
	},
	image.Point{7, 1}: rider{
		glyph: skier,
		//   - 91ff1c // neon green
		Color: tcell.NewRGBColor(0x91, 0xFF, 0x1C),
	},
	image.Point{1, 2}: rider{
		glyph: snowbdr,
		//   - ff1c7b // pink
		Color: tcell.NewRGBColor(0xFF, 0x1C, 0x7B),
	},
}

func foreground(s tcell.Screen, scene int, input [][]string) {
	w, h := s.Size()
	m := len(input[0])
	// Trees
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if input[(y+scene)%len(input)][x%m] == "#" {
				_, _, style, _ := s.GetContent(x, y)
				s.SetContent(x, y, tree, nil, style)
			}
		}
	}
	for slope, rider := range riders {
		p := rider.Point
		// Update the trail along the slope and put rider at the end.
		rider.Point.Y = scene
		rider.Point.X = scene * slope.X
		for _, q := range astar(p, rider.Point) {
			if q.Y == 0 {
				continue
			}
			_, _, style, _ := s.GetContent(q.X, q.Y)
			s.SetContent(q.X, q.Y, trail, nil, style.Foreground(rider.Color))
		}
		_, _, style, _ := s.GetContent(rider.Point.X, rider.Point.Y)
		s.SetContent(rider.Point.X, rider.Point.Y, rider.glyph, nil, style)
	}
}

func shred(input [][]string) {
	sc, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := sc.Init(); err != nil {
		panic(err)
	}
	sc.Clear()
	userquit := make(chan struct{})
	go func() {
		for {
			switch ev := sc.PollEvent().(type) {
			case *tcell.EventResize:
				sc.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					userquit <- struct{}{}
				}
			}
		}
	}()
	defer func() {
		sc.Fini()
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	var (
		clock = time.Second / 15
		timer = time.NewTimer(clock)
		scene = 0
	)

	// Loop:
	//  - reset the scene timer
	//  - paint background
	//  - update riders and trails
	//  - pause for refresh or cancel
	//  - display the screen
	//  - increment the row
loop:
	for scene < len(input) {
		timer.Reset(clock)
		background(sc)
		foreground(sc, scene, input)
		sc.Show()
		select {
		case <-userquit:
			break loop
		case <-timer.C:
		}
		scene++
	}
	sc.Fini()
}
