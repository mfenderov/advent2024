package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchPath(t *testing.T) {
	guardMap = readFile("test_input.txt")
	i, j, direction := findStart()
	walkGuard(i, j, direction)
	paths := countWalkedPath()
	searchForLoops(i, j, direction)
	assert.Equal(t, 6, loopsFounds)
	assert.Equal(t, 41, paths)
}
