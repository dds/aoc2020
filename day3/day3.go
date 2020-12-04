package main

import (
	"fmt"
	"image"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/dds/aoc2020/lib"
	"github.com/dds/aoc2020/lib/inputs"
	"github.com/gdamore/tcell/v2"
)

var Input = lib.ParseInput(inputs.Day3(), func(s string) []string { return strings.Split(s, "") })

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}

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
	trail     = '╍'
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
		// Update the trail along the slope and put rider at the end.
		rider.Point.Y = scene
		rider.Point.X = scene * slope.X
		_, _, style, _ := s.GetContent(rider.Point.X, rider.Point.Y)
		s.SetContent(rider.Point.X, rider.Point.Y, rider.glyph, nil, style.Foreground(rider.Color))
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
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
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
		clock = time.Second / 10
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
