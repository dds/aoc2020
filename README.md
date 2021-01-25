# Advent of Code 2020 (Golang) 

[![Build Status](https://github.com/dds/aoc2020/workflows/ci/badge.svg)](https://github.com/dds/aoc2020/actions?query=workflow%3Aci)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/dds/aoc2020)](https://goreportcard.com/report/github.com/dds/aoc2020)

This repo contains my solutions to [Advent of Code, 2020](https://adventofcode.com/2020).

## Prerequisites

You will need `go` installed. Recommend installing [`brew` ](https://brew.sh/) and running `brew install go`.

## Running the Solutions

Download, build, and install the solutions like so:

```sh
# Download, build, and install solutions from this repo
go get github.com/dds/aoc2020/...

# Run solution for day 1
day1

# Get input for day 20, wait up to 8h, and copy to clipboard if successful
aocinput 20 --timeout 8h --clipboard

# Watch the time count down until AoC 2022 Day 1, then display the input and stop.
watch -n 0.3 -e '! aocinput 1 -y 2022'
```
