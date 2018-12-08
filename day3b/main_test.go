package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseInputLine(t *testing.T) {
	claim := parseInputLine("#1 @ 1,3: 4x4")
	assert.Equal(t, Coordinate{1, 3}, claim.Coor)
	assert.Equal(t, 1, claim.Id)
	assert.Equal(t, 4, claim.Length)
	assert.Equal(t, 4, claim.Width)

	claim = parseInputLine("#2 @ 3,1: 4x4")
	assert.Equal(t, Coordinate{3, 1}, claim.Coor)
	assert.Equal(t, 2, claim.Id)
	assert.Equal(t, 4, claim.Length)
	assert.Equal(t, 4, claim.Width)

	claim = parseInputLine("#3 @ 5,5: 2x2")
	assert.Equal(t, 3, claim.Id)
	assert.Equal(t, Coordinate{5, 5}, claim.Coor)
	assert.Equal(t, 2, claim.Length)
	assert.Equal(t, 2, claim.Width)
}

func TestGivenExample(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2"}
	res := FindIdOfClaimWithNoOverlap(input)
	assert.Equal(t, 3, res)
}

func TestClaimsNotInMap(t *testing.T) {
	map1 := map[Claim]bool{
		Claim{
			Id: 1,
		}: true,
		Claim{
			Id: 2,
		}: true,
		Claim{
			Id: 3,
		}: true,
	}
	map2 := map[Claim]bool{
		Claim{
			Id: 2,
		}: true,
		Claim{
			Id: 3,
		}: true,
	}
	diff := findMapDiff(map1, map2)
	expected := []Claim{
		Claim{Id: 1}}
	assert.ElementsMatch(t, expected, diff)
}
