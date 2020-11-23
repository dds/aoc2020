package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	type test struct {
		input  int
		expect int
	}

	tests := []test{
		test{
			// ...
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, test.input)
		})
	}
}
