package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PuzzleInput struct {
	Lines           [][]string
	TransposedLines [][]string
	StringLines     []string
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
