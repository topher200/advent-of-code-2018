package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
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

func LoadInputFile(filename string) (res []string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return
}

func main() {
	inputFilenamePtr := flag.String("input", "", "input filename")

	flag.Parse()
	inputFilename := *inputFilenamePtr

	inputValues := LoadInputFile(inputFilename)
	println(ApplyListOfValues(inputValues))
}
