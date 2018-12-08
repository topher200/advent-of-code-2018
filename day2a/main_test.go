package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestCountLettersTwo(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	numLetters := 2
	expectedResult := 4
	assert.Equal(t, expectedResult, GetNumMatches(input, numLetters))
}

func TestCountLettersThree(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	numLetters := 3
	expectedResult := 3
	assert.Equal(t, expectedResult, GetNumMatches(input, numLetters))
}

func TestGetChecksum(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	assert.Equal(t, 12, GetCheckSum(input))
}
