package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
	"github.com/atotto/clipboard"
	"github.com/dds/aoc2020/lib"
)

var CLI struct {
	Clipboard bool
	Day       int `kong:"arg,required"`
	Year      int
	Session   string `kong:"required"`
	Timeout   time.Duration
}

func main() {
	kong.Parse(&CLI)
	if CLI.Year == 0 {
		CLI.Year = time.Now().Year()
	}
	deadline := time.Now().Add(CLI.Timeout)
	s, err := lib.GetInput(CLI.Year, CLI.Day, CLI.Session)
	for s == "" && time.Now().Before(deadline) {
		time.Sleep(3 * time.Second)
		s, err = lib.GetInput(CLI.Year, CLI.Day, CLI.Session)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	if CLI.Clipboard {
		if err := clipboard.WriteAll(s); err != nil {
			panic(err)
		}
	}
}
