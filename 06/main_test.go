package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchPath(t *testing.T) {
	guardMap = readFile("test_input.txt")
	i, j, direction := findStart()
	walkGuard(i, j, direction)
	printGuardMap()
	paths := countWalkedPath()
	assert.Equal(t, 41, paths)
}
