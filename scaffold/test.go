package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	input := util.Inputs[0]

	type test struct {
		// ...
	}

	tests := []test{
		test{
			// ...
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var err error
			// ...
			require.NoError(t, err)
		})
	}
}
