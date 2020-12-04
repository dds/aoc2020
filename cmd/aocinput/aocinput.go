// Install with `go get github.com/dds/aoc2020/cmd/aocinput`.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/atotto/clipboard"
	"github.com/dds/aoc2020/lib"
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/allbrowsers"
	"github.com/zellyn/kooky/chrome"
)

var CLI struct {
	Clipboard  bool          `short:"c" help:"Copy to window system clipboard on success."`
	CookieFile string        `short:"f" help:"File path of Chrome/Chromium cookie jar file."`
	Day        int           `kong:"arg,required"`
	Session    string        `short:"s" help:"Your personal session cookie from your browser."`
	Timeout    time.Duration `short:"t" help:"Retry up to timeout. Examples: 8h,20s."`
	Year       int           `short:"y" help:"Year. Default is current year."`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("aocinput"),
		kong.Description("Download Advent of Code Puzzle Inputs"),
		kong.UsageOnError(),
	)
	if CLI.Year == 0 {
		CLI.Year = time.Now().Year()
	}
	if CLI.CookieFile == "" {
		userdir, _ := os.UserConfigDir()
		CLI.CookieFile = userdir + "/Google/Chrome/Default/Cookies"
	}
	if CLI.Session == "" {
		stores := kooky.FindAllCookieStores()
		cookieOpts := []kooky.Filter{kooky.Valid, kooky.Name("session")}
		// Try Firefox first because otherwise it prompts for keychain for
		// Chrome, which I must click to skip to get to Firefox.
		for _, store := range stores {
			if store.Browser() != "firefox" {
				continue
			}
			cookies, err := store.ReadCookies(append(cookieOpts, kooky.Domain("adventofcode.com"))...)
			if err != nil {
				continue
			}
			if len(cookies) < 1 {
				continue
			}
			CLI.Session = cookies[0].Value
		}
		if CLI.Session == "" {
			for _, store := range stores {
				cookies, err := store.ReadCookies(append(cookieOpts, kooky.Domain(".adventofcode.com"))...)
				if err != nil {
					continue
				}
				if len(cookies) < 1 {
					continue
				}
				CLI.Session = cookies[0].Value
			}
		}
		if CLI.Session == "" {
			cookies, err := chrome.ReadCookies(CLI.CookieFile, append(cookieOpts, kooky.Domain(".adventofcode.com"))...)
			if err == nil && len(cookies) >= 1 {
				CLI.Session = cookies[0].Value
			}
		}
		if CLI.Session == "" {
			ctx.FatalIfErrorf(fmt.Errorf("No session cookie. Must set --session or --cookie-file argument."))
		}
	}
	deadline := time.Now().Add(CLI.Timeout)
	s, err := lib.GetInput(CLI.Year, CLI.Day, CLI.Session)
	for s == "" && time.Now().Before(deadline) {
		time.Sleep(3 * time.Second)
		s, err = lib.GetInput(CLI.Year, CLI.Day, CLI.Session)
	}
	if err != nil {
		ctx.FatalIfErrorf(err)
	}
	fmt.Print(s)
	if CLI.Clipboard {
		if err := clipboard.WriteAll(s); err != nil {
			ctx.FatalIfErrorf(err)
		}
	}
}
