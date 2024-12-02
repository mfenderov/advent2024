package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validation(t *testing.T) {
	testCases := []struct {
		parts    []string
		expected bool
	}{
		{
			[]string{"10", "1", "3", "5"},
			true,
		},
		{
			[]string{"1", "3", "5", "10"},
			true,
		},
		{
			[]string{"1", "3", "5", "10", "20"},
			false,
		},
		{
			[]string{"5", "7", "3", "2", "1"},
			true,
		},
		{
			[]string{"5", "7", "8", "2", "1"},
			false,
		},
		{
			[]string{"5", "7", "8", "1", "2"},
			false,
		},
		{
			[]string{"10", "7", "8", "6", "5"},
			true,
		},
		{
			[]string{"83", "86", "85", "82", "80"},
			true,
		},
		{
			[]string{"35", "38", "37", "34", "33", "30", "28"},
			true,
		},
		{
			[]string{"22", "21", "22", "24", "25", "26", "27"},
			true,
		},
	}
	for _, tc := range testCases {
		t.Run("Test validation", func(t *testing.T) {
			actual := validate(tc.parts, 0)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
