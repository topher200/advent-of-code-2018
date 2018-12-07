package libinput

import (
	"bufio"
	"flag"
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

func ReadLinesFromCLIInput() []string {
	inputFilenamePtr := flag.String("input", "", "input filename")

	flag.Parse()
	inputFilename := *inputFilenamePtr

	inputValues := LoadInputFile(inputFilename)
	return inputValues
}
