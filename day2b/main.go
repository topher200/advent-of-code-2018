package main

import (
	"errors"

	"github.com/topher200/advent-of-code-2018/libinput"
)

// we need to find the correct two boxes  <-- hard
// then we need to find the identical letters between the two <-- easy

// how do we find the correct boxes? we need to find the two strings who have
// the same characters at each position other than exactly one position

// start simple: how do we find the two strings which have the exact same
// characters? i think i'd use some kind of hash and find the two with the same
// hash value. if we're looking for ones with one difference... i'm not as sure.
// we could do something like "hash many versions of each string, each version
// missing a letter". that seems crazy though

// i do just need to know the matching characters, so maybe it isn't crazy! for
// each input string, remove a each character in-turn and add the neutered
// version of the string to a map. if we we're ever attempting to add a string
// to the map that already exists, return it instead!

func GetIdenticalLettersBetweenCorrectBoxes(input []string) (string, error) {
	stringCollection := make(map[string]string)
	for _, inputString := range input {
		// split the input into runes. make strings dropping one rune at a time
		runes := []rune(inputString)
		for i := range runes {
			st := string(runes[:i]) + string(runes[i+1:])
			// if we've seen this substring before, check to make sure it's
			// from a different input string
			if stringCollection[st] != "" && stringCollection[st] != inputString {
				return st, nil
			}
			stringCollection[st] = inputString
		}
	}
	return "", errors.New("could not find match")
}

func main() {
	res, err := GetIdenticalLettersBetweenCorrectBoxes(
		libinput.ReadLinesFromCLIInput("day2b"))
	if err != nil {
		panic(err)
	}
	println(res)
}
