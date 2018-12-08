package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseInputLine(t *testing.T) {
	coordinate, length, width := parseInputLine("#1 @ 1,3: 4x4")
	assert.Equal(t, Coordinate{1, 3}, coordinate)
	assert.Equal(t, 4, length)
	assert.Equal(t, 4, width)

	coordinate, length, width = parseInputLine("#2 @ 3,1: 4x4")
	assert.Equal(t, Coordinate{3, 1}, coordinate)
	assert.Equal(t, 4, length)
	assert.Equal(t, 4, width)

	coordinate, length, width = parseInputLine("#3 @ 5,5: 2x2")
	assert.Equal(t, Coordinate{5, 5}, coordinate)
	assert.Equal(t, 2, length)
	assert.Equal(t, 2, width)
}

func TestGivenExample(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2"}
	res := CountNumberOfDoubleBookedSquares(input)
	assert.Equal(t, 4, res)
}

func TestNoOverlap(t *testing.T) {
	input := []string{
		"#1 @ 1,1: 4x4",
		"#3 @ 5,5: 2x2"}
	res := CountNumberOfDoubleBookedSquares(input)
	assert.Equal(t, 0, res)
}

func TestFullOverlap(t *testing.T) {
	input := []string{
		"#1 @ 1,1: 4x4",
		"#3 @ 1,1: 4x4"}
	res := CountNumberOfDoubleBookedSquares(input)
	assert.Equal(t, 16, res)
}

func TestSomeOverlap(t *testing.T) {
	input := []string{
		"#1 @ 1,1: 4x4",
		"#3 @ 1,1: 2x2"}
	res := CountNumberOfDoubleBookedSquares(input)
	assert.Equal(t, 4, res)
}
