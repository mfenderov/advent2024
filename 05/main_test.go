package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validation(t *testing.T) {
	totalValid, totalFixed := readFile("test_input.txt")
	assert.Equal(t, 143, totalValid)
	assert.Equal(t, 123, totalFixed)
}
