package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseSimple(t *testing.T) {
	file := []byte(`mul(1,2) add(3,4)`)
	expected := [][]byte{[]byte("mul(1,2)")}
	actual := parse(file)
	assert.Equal(t, expected, actual)
}

func Test_ParseMultiple(t *testing.T) {
	file := []byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)
	expected := [][]byte{
		[]byte("mul(2,4)"),
		[]byte("mul(5,5)"),
		[]byte("mul(11,8)"),
		[]byte("mul(8,5)"),
	}
	actual := parse(file)
	assert.Equal(t, expected, actual)
}

func Test_ParseNumbers(t *testing.T) {
	mul := []byte("mul(123,234)")
	expected := [][]byte{[]byte("123"), []byte("234")}
	actual := parseNumbers(mul)
	assert.Equal(t, expected, actual)
}
