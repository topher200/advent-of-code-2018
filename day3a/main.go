package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/topher200/advent-of-code-2018/libinput"
)

// the "obvious" way to solve this is to create a double list, representing the
// square board. then for each line of input we use a double 'for' loop to
// mark out the space contained in that line of input. finally, we count the
// number of squares which have been "double booked"

// thinking about it more, we don't need to use the double list to represent
// the board. we can use a hash table to store x,y pairs of coordinates which
// we've added a "spot" onto. so we use a hash table that is
// <(x,y), (numTimesUsed)>. then, we just count the number of values in the
// hash table > 1

type Coordinate struct {
	X int
	Y int
}

var locationRegex = regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)

func parseInputLine(input string) (coordinate Coordinate, length int, width int) {
	// example input line: "#1 @ 1,3: 4x4"
	match := locationRegex.FindStringSubmatch(input)
	if len(match) != 5 {
		panic(fmt.Errorf("unable to parse input string: %v", input))
	}
	x_, err := strconv.ParseInt(match[1], 10, 0)
	if err != nil {
		panic(err)
	}
	x := int(x_)
	y_, err := strconv.ParseInt(match[2], 10, 0)
	if err != nil {
		panic(err)
	}
	y := int(y_)
	coordinate = Coordinate{x, y}
	length_, err := strconv.ParseInt(match[3], 10, 0)
	if err != nil {
		panic(err)
	}
	length = int(length_)
	width_, err := strconv.ParseInt(match[4], 10, 0)
	if err != nil {
		panic(err)
	}
	width = int(width_)
	return coordinate, length, width
}

func CountNumberOfDoubleBookedSquares(inputLines []string) int {
	// for each square of input, add its coordinates to our map
	seenSquares := make(map[Coordinate]int)
	for _, line := range inputLines {
		coordinate, length, width := parseInputLine(line)
		for x := 0; x < length; x++ {
			for y := 0; y < width; y++ {
				seenSquares[Coordinate{coordinate.X + x, coordinate.Y + y}] += 1
			}
		}
	}
	// iterate through the map to determine how many spots were seen twice
	sum := 0
	for _, numHits := range seenSquares {
		if numHits > 1 {
			sum += 1
		}
	}
	return sum
}

func main() {
	res := CountNumberOfDoubleBookedSquares(
		libinput.ReadLinesFromCLIInput("day3a"))
	println(res)
}
