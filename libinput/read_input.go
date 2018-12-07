package libinput

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

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

func ReadLinesFromCLIInput(directory string) []string {
	pathToDefaultInputFile := fmt.Sprintf("%s/input.txt", directory)
	inputFilenamePtr := flag.String("input", pathToDefaultInputFile, "input filename")

	flag.Parse()
	inputFilename := *inputFilenamePtr

	inputValues := LoadInputFile(inputFilename)
	return inputValues
}
