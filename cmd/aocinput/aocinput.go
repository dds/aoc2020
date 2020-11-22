package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
	"github.com/atotto/clipboard"
	"github.com/dds/aoc2020/util"
)

var CLI struct {
	Clipboard bool
	Day       int    `kong:"arg,required"`
	Session   string `kong:"required"`
	Timeout   time.Duration
}

func main() {
	kong.Parse(&CLI)
	deadline := time.Now().Add(CLI.Timeout)
	var (
		s   string
		err error
	)
	for s, err = util.GetInput(CLI.Day, CLI.Session); time.Now().Before(deadline); s, err = util.GetInput(1, CLI.Session) {
		if err != nil {
			panic(err)
		}
		if s == "" {
			time.Sleep(3 * time.Second)
			continue
		}
		break
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
