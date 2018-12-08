package main

import (
	"strconv"

	"github.com/topher200/advent-of-code-2018/libinput"
)

func FindFirstRepeatFrequency(listOfvalues []string) int {
	// loop through the list forever
	// use a map to keep track of sums which we reach
	// once we hit the same sum twice, return it
	sum := 0
	seenSums := make(map[int]bool)
	for {
		for _, valueString := range listOfvalues {
			// this could be more efficient by only doing Atoi once
			value, err := strconv.Atoi(valueString)
			if err != nil {
				panic(err)
			}
			sum += value
			if seenSums[sum] {
				return sum
			} else {
				seenSums[sum] = true
			}
		}
	}
}

func main() {
	println(FindFirstRepeatFrequency(libinput.ReadLinesFromCLIInput("day1b")))
}
