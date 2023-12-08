package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

type CubeData struct {
	GameID      int
	ColorValues map[string][]int
}

func main() {
	startTime, input := lib.Init()
	defer lib.Close(startTime)

	totalsPuzzle1 := 0
	totalsPuzzle2 := 0

	nrOfCubesMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	resultMap := getCubeNumbers(*input)

	for _, cubeData := range resultMap {
		// Puzzle 1
		tooManyCubes := false

		for color, valueSlice := range cubeData.ColorValues {
			if slices.Max(valueSlice) > nrOfCubesMap[color] {
				tooManyCubes = true
				break
			}
		}
		if !tooManyCubes {
			totalsPuzzle1 += cubeData.GameID
		}

		// Puzzle 2
		maxRed := slices.Max(cubeData.ColorValues["red"])
		maxBlue := slices.Max(cubeData.ColorValues["blue"])
		maxGreen := slices.Max(cubeData.ColorValues["green"])

		totalsPuzzle2 += maxRed * maxBlue * maxGreen
	}

	fmt.Printf("Puzzle 1: %d\n", totalsPuzzle1)
	fmt.Printf("Puzzle 2: %d\n", totalsPuzzle2)
}

func getCubeNumbers(input lib.PuzzleInput) []CubeData {
	var results []CubeData

	for _, game := range input.StringLines {

		var result CubeData
		result.ColorValues = make(map[string][]int)
		splitGame := strings.Split(game, ":")
		result.GameID, _ = strconv.Atoi(strings.Split(splitGame[0], " ")[1])
		gameValues := strings.Split(splitGame[1], ";")

		for _, shownCubes := range gameValues {
			values := strings.Split(shownCubes, ",")
			for _, cubes := range values {
				value := strings.Split(strings.Trim(cubes, " "), " ")
				nrOfCubes, _ := strconv.Atoi(value[0])
				currentColor := value[1]

				result.ColorValues[currentColor] = append(result.ColorValues[currentColor], nrOfCubes)
			}
		}
		results = append(results, result)
	}

	return results
}
