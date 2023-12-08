package lib

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type PuzzleInput struct {
	Lines           [][]string
	TransposedLines [][]string
	StringLines     []string
}

func Init() (time.Time, *PuzzleInput) {
	iFlag := flag.Bool("i", false, "When passed the input.txt file is taken")

	tm := time.Now()
	flag.Parse()

	inputFile := "sample.txt"
	if *iFlag {
		inputFile = "input.txt"
	}

	input, err := GetInput(inputFile, false)
	if err != nil {
		fmt.Println(err)
	}

	return tm, input
}

func Close(startTime time.Time) {
	duration := time.Since(startTime)
	fmt.Println("Execution time: ", duration)
}

func NewPuzzleInputFromFile(filePath string, doTranspose bool) (*PuzzleInput, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	var stringLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringLines = append(stringLines, line)
		lineSlice := strings.Split(line, "")
		lines = append(lines, lineSlice)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var transposedLines [][]string
	if doTranspose {
		transposedLines = transpose(lines)
	}

	return &PuzzleInput{
		Lines:           lines,
		TransposedLines: transposedLines,
		StringLines:     stringLines,
	}, nil
}

func transpose(slice [][]string) [][]string {
	if len(slice) == 0 {
		return slice
	}

	transposed := make([][]string, len(slice[0]))
	for i := range transposed {
		transposed[i] = make([]string, len(slice))
	}

	for i, row := range slice {
		for j, val := range row {
			transposed[j][i] = val
		}
	}

	return transposed
}

func GetInput(fileName string, transpose bool) (*PuzzleInput, error) {
	puzzleInput, err := NewPuzzleInputFromFile(fmt.Sprintf("./%s", fileName), transpose)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return &PuzzleInput{}, err
	}
	return puzzleInput, nil
}
