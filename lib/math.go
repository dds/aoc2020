package lib

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Reduce(f func(int, int) int, c <-chan int) (r int) {
	r = f(<-c, <-c)
	for i := range c {
		r = f(r, i)
	}
	return
}

func Apply(f func(int) int, c <-chan int) <-chan int {
	r := make(chan int)
	go func() {
		defer close(r)
		for i := range c {
			r <- f(i)
		}
	}()
	return r
}

func Sign(a int) int {
	if a >= 0 {
		return 1
	}
	return -1
}
