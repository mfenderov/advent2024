package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	actual, masses := findAllXmasInFile("test_input.txt")
	assert.Equal(t, 18, actual)
	assert.Equal(t, 9, masses)
}
