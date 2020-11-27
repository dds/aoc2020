package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplyReduce(t *testing.T) {
	type test struct {
		applyer func(int) int
		reducer func(int, int) int
		input   []int
		expect  int
	}
	tests := []test{
		test{
			applyer: func(x int) int { return x * x },
			reducer: func(x, y int) int { return x + y },
			input:   []int{1, -2, 3},
			expect:  1 + 4 + 9,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			nums := make(chan int)
			go func() {
				defer close(nums)
				for _, i := range test.input {
					nums <- i
				}
			}()
			require.Equal(t, test.expect, Reduce(test.reducer, Apply(test.applyer, nums)))
		})
	}
}
