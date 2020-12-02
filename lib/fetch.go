package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const YEAR = 2020

const startTime = "2020-12-01T00:00:00-05:00"

var start time.Time

func init() {
	var err error
	start, err = time.Parse(time.RFC3339, startTime)
	if err != nil {
		panic(err)
	}
}

// GetInput returns the puzzle input for the given day.
func GetInput(day int, sessionCookie string) (r string, err error) {
	start := start.Add(time.Duration(24*(day-1)) * time.Hour)
	if time.Now().Before(start) {
		return "", fmt.Errorf("AoC %v Day %v has not yet begun (%v)", YEAR, day, time.Until(start))
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
