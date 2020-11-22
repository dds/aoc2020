package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const YEAR = 2020
const start = "01 Dec 20 00:00 -0000"

// GetInput returns the puzzle input for the given day.
func GetInput(day int, sessionCookie string) (r string, err error) {
	start, err := time.Parse(time.RFC822Z, start)
	if err != nil {
		return
	}
	if time.Now().Before(start) {
		return "", fmt.Errorf("AoC %v does not begin until %v (-%v)", YEAR, start, time.Until(start))
	}
	var (
		res *http.Response
		req *http.Request
	)
	req, err = http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%v/day/%d/input", YEAR, day), nil)
	if err != nil {
		return
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	var b []byte
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	res.Body.Close()
	r = string(b)
	return
}
