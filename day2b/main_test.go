package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetIdentialLettersBetweenCorrectBoxes(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	result, err := GetIdenticalLettersBetweenCorrectBoxes(input)
	assert.Nil(t, err)
	assert.Equal(t, "fgij", result)
}
