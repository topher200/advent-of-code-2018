package main

import (
	"strconv"

	"github.com/topher200/advent-of-code-2018/libinput"
)

func ApplyListOfValues(listOfvalues []string) int {
	sum := 0
	for _, valueString := range listOfvalues {
		value, err := strconv.Atoi(valueString)
		if err != nil {
			panic(err)
		}
		sum += value
	}
	return sum
}

func main() {
	println(ApplyListOfValues(libinput.ReadLinesFromCLIInput()))
}
