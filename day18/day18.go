package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2020/lib/inputs"
)

// Parse input into an array of bytes: an expr.
type expr []byte

func (e expr) String() (r string) {
	for i := 0; i < len(e); i++ {
		r += string(e[i])
	}
	return
}

func parse(l string) (r expr) {
	for i := 0; i < len(l); i++ {
		switch l[i] {
		case ' ':
		default:
			r = append(r, l[i])
		}
	}
	return
}

var NumberRE = regexp.MustCompile(`^\d+`)

func num(b []byte) (r int) {
	if !NumberRE.Match(b) {
		return -1
	}
	r, _ = strconv.Atoi(string(NumberRE.Find(b)))
	return
}

func (e expr) eval1() (rc int) {
	orig := make(expr, len(e))
	copy(orig, e)
	var out, ops []byte
	for len(e) > 0 {
		a := num(e[0:])
		if a != -1 {
			out = append(out, e[0])
			e = e[1:]
			continue
		}
		if e[0] == '+' || e[0] == '*' {
			for {
				n := len(ops)
				if n == 0 || ops[n-1] == '(' {
					break
				}
				out = append(out, ops[n-1])
				ops = ops[:n-1]
			}
			ops = append(ops, e[0])
			e = e[1:]
			continue
		}
		if e[0] == '(' {
			ops = append(ops, e[0])
			e = e[1:]
			continue
		}
		if e[0] == ')' {
			for {
				n := len(ops)
				if n <= 0 {
					panic(fmt.Errorf("mismatched parens: %v", orig))
				}
				if ops[n-1] != '(' {
					out = append(out, ops[n-1])
					ops = ops[:n-1]
					continue
				} else {
					ops = ops[:n-1]
					e = e[1:]
					break
				}
			}
		}
	}
	for len(ops) > 0 {
		if ops[0] == '(' || ops[0] == ')' {
			panic(fmt.Errorf("mismatched parens: %v", orig))
		}
		out = append(out, ops[0])
		ops = ops[1:]
	}

	stk := []int{}
	for i := 0; i < len(out); i++ {
		n := len(stk)
		if out[i] == '+' {
			if n < 2 {
				panic(fmt.Errorf("invalid expression: %v (=> %v)", orig, expr(out)))
			}
			t := stk[n-2] + stk[n-1]
			stk[n-2] = t
			stk = stk[:n-1]
			continue
		}
		if out[i] == '*' {
			if n < 2 {
				panic(fmt.Errorf("invalid expression: %v (=> %v)", orig, out))
			}
			t := stk[n-2] * stk[n-1]
			stk[n-2] = t
			stk = stk[:n-1]
			continue
		}
		stk = append(stk, num(out[i:i+1]))
	}
	return stk[0]
}

func part1(in string) (rc int) {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for _, line := range lines {
		expr := parse(line)
		rc += expr.eval1()
	}
	return
}

func (e expr) eval2() (rc int) {
	orig := make(expr, len(e))
	copy(orig, e)
	var out, ops []byte
	for len(e) > 0 {
		a := num(e[0:])
		if a != -1 {
			out = append(out, e[0])
			e = e[1:]
			continue
		}
		if e[0] == '+' || e[0] == '*' {
			for {
				n := len(ops)
				if n == 0 {
					break
				}
				if ops[n-1] == '(' {
					break
				}
				if ops[n-1] == '+' {
					out = append(out, ops[n-1])
					ops = ops[:n-1]
					continue
				}
				break
			}
			ops = append(ops, e[0])
			e = e[1:]
			continue
		}
		if e[0] == '(' {
			ops = append(ops, e[0])
			e = e[1:]
			continue
		}
		if e[0] == ')' {
			for {
				n := len(ops)
				if n <= 0 {
					panic(fmt.Errorf("mismatched parens: %v", orig))
				}
				if ops[n-1] != '(' {
					out = append(out, ops[n-1])
					ops = ops[:n-1]
					continue
				} else {
					ops = ops[:n-1]
					e = e[1:]
					break
				}
			}
		}
	}

	for len(ops) > 0 {
		n := len(ops) - 1
		if ops[n] == '(' || ops[n] == ')' {
			panic(fmt.Errorf("mismatched parens: %v", orig))
		}
		out = append(out, ops[n])
		ops = ops[:n]
	}

	stk := []int{}
	for i := 0; i < len(out); i++ {
		n := len(stk)
		if out[i] == '+' {
			if n < 2 {
				panic(fmt.Errorf("invalid expression: %v (=> %v)", orig, expr(out)))
			}
			t := stk[n-2] + stk[n-1]
			stk[n-2] = t
			stk = stk[:n-1]
			continue
		}
		if out[i] == '*' {
			if n < 2 {
				panic(fmt.Errorf("invalid expression: %v (=> %v)", orig, out))
			}
			t := stk[n-2] * stk[n-1]
			stk[n-2] = t
			stk = stk[:n-1]
			continue
		}
		stk = append(stk, num(out[i:i+1]))
	}
	return stk[0]
}

func part2(in string) (rc int) {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for _, line := range lines {
		expr := parse(line)
		rc += expr.eval2()
	}
	return
}

var Input = inputs.Day18()

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}
