package main

import (
	"github.com/topher200/advent-of-code-2018/libinput"
)

func GetNumMatches(listOfvalues []string, numIdenticalLetters int) int {
	// loop through the input. count each time we have a string that has
	// numIdenticalLetters identical letters
	sumMatchingWords := 0
	for _, valueString := range listOfvalues {
		seenLetters := make(map[rune]int)
		for _, letter := range valueString {
			seenLetters[letter] += 1
		}
		for _, numTimesSeen := range seenLetters {
			if numTimesSeen == numIdenticalLetters {
				sumMatchingWords += 1
				break
			}
		}
	}
	return sumMatchingWords
}

func GetCheckSum(listOfvalues []string) int {
	return GetNumMatches(listOfvalues, 2) * GetNumMatches(listOfvalues, 3)
}

func main() {
	println(GetCheckSum(libinput.ReadLinesFromCLIInput("day2a")))
}
