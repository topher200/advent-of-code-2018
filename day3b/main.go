package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/topher200/advent-of-code-2018/libinput"
)

// for this one, i think we need to make to make a map of coordinates again.
// this time, we should have a map of <coordinate, [array of ids which cover
// this coordinate]>. then, we can look through all the values of our map and
// pick out which ever one has a length of 1

type Coordinate struct {
	X int
	Y int
}

type Claim struct {
	Coor   Coordinate
	Length int
	Width  int
	Id     int
}

var locationRegex = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func parseInputLine(input string) (claim Claim) {
	// example input line: "#1 @ 1,3: 4x4"
	match := locationRegex.FindStringSubmatch(input)
	if len(match) != 6 {
		panic(fmt.Errorf("unable to parse input string: %v", input))
	}
	id_, err := strconv.ParseInt(match[1], 10, 0)
	if err != nil {
		panic(err)
	}
	id := int(id_)
	x_, err := strconv.ParseInt(match[2], 10, 0)
	if err != nil {
		panic(err)
	}
	x := int(x_)
	y_, err := strconv.ParseInt(match[3], 10, 0)
	if err != nil {
		panic(err)
	}
	y := int(y_)
	coordinate := &Coordinate{x, y}
	length_, err := strconv.ParseInt(match[4], 10, 0)
	if err != nil {
		panic(err)
	}
	length := int(length_)
	if length <= 0 {
		panic(fmt.Errorf("bad input string: %v", input))
	}
	width_, err := strconv.ParseInt(match[5], 10, 0)
	if err != nil {
		panic(err)
	}
	width := int(width_)
	if width <= 0 {
		panic(fmt.Errorf("bad input string: %v", input))
	}
	return Claim{*coordinate, length, width, id}
}

func FindIdOfClaimWithNoOverlap(inputLines []string) int {
	// for each square of input, add its coordinates to our map
	seenSquares := make(map[Coordinate][]Claim)
	allClaims := make(map[Claim]bool)
	for _, line := range inputLines {
		claim := parseInputLine(line)
		allClaims[claim] = true
		for x := 0; x < claim.Length; x++ {
			for y := 0; y < claim.Width; y++ {
				thisCoordinate := Coordinate{claim.Coor.X + x, claim.Coor.Y + y}
				seenSquares[thisCoordinate] = append(
					seenSquares[thisCoordinate], claim)
			}
		}
	}
	// iterate through the map to find the claim who never shares a spot with
	// anyone else
	claimsWithSharedSquares := make(map[Claim]bool)
	for _, claims := range seenSquares {
		if len(claims) > 1 {
			for _, claim := range claims {
				claimsWithSharedSquares[claim] = true
			}
		}
	}
	claimsWithNoSharedSquares := findMapDiff(allClaims, claimsWithSharedSquares)
	if len(claimsWithNoSharedSquares) != 1 {
		panic("no matching claims found")
	}
	return claimsWithNoSharedSquares[0].Id
}

func findMapDiff(map1 map[Claim]bool, map2 map[Claim]bool) (
	claimsNotInMap1 []Claim) {
	// returns any Claim found in map1 that are not in map2
	for claim, _ := range map1 {
		if !map2[claim] {
			claimsNotInMap1 = append(claimsNotInMap1, claim)
		}
	}
	return claimsNotInMap1
}

func main() {
	res := FindIdOfClaimWithNoOverlap(
		libinput.ReadLinesFromCLIInput("day3a"))
	println(res)
}
