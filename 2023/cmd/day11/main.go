package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

var coordinateMap = make(map[int]lib.Coordinate)

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	nrOfGalaxies, galaxyMap := prepareGalaxy(input.Lines)

	galaxyPairs := lib.GenerateIntPairs(nrOfGalaxies)

	puzzle1Res = getGalaxyDistances(galaxyPairs, galaxyMap)

	//puzzle2
	// probably instead of doing the array transformation, just
	// get all coordinates without expanding the galaxy, then add
	// expansions manually. (so in rowDiff, is there a column that needs to be expanded in between)
	// _, galaxyMap := prepareGalaxy(input.Lines)

}

func getGalaxyDistances(galaxyPairs []lib.Pair[lib.IntCompare], galaxyMap [][]int) (result int) {
	for _, pair := range galaxyPairs {
		coordinate1 := coordinateMap[int(pair.Fst)]
		coordinate2 := coordinateMap[int(pair.Snd)]
		result += calculateDistance(coordinate1, coordinate2, galaxyMap)
	}
	return
}

// calculateDistance calculates the Manhattan distance between two coordinates on a grid.
func calculateDistance(coordinate1, coordinate2 lib.Coordinate, galaxyMap [][]int) int {
	// Calculate the absolute difference in rows and columns
	rowDiff := math.Abs(float64(coordinate1.Row - coordinate2.Row))
	colDiff := math.Abs(float64(coordinate1.Col - coordinate2.Col))

	// Sum the differences to get the Manhattan distance
	distance := int(rowDiff + colDiff)
	return distance
}

func prepareGalaxy(lines [][]string) (int, [][]int) {
	rowNumbers := getEmptyRows(lines)
	galaxyMap := expandGalaxy(lines, rowNumbers)
	galaxyMap = lib.Transpose(galaxyMap)
	rowNumbers = getEmptyRows(galaxyMap)

	galaxyMap = expandGalaxy(galaxyMap, rowNumbers)

	galaxyMap = lib.Transpose(galaxyMap)
	count := numberGalaxy(galaxyMap)

	intMap := make([][]int, len(galaxyMap))
	for i := range intMap {
		intMap[i] = make([]int, len(galaxyMap[i]))
	}

	for row, line := range galaxyMap {
		for col, char := range line {
			if char == "." {
				intMap[row][col] = 0
			} else {
				nr, err := strconv.Atoi(char)
				if err != nil {
					fmt.Println(err)
				}
				intMap[row][col] = nr
				coordinateMap[nr] = lib.Coordinate{Row: row, Col: col}
			}
		}
	}

	return count, intMap
}

func numberGalaxy(galaxyMap [][]string) int {
	count := 0
	for i, line := range galaxyMap {
		indexes := findCharOccurrences(line, "#")
		if len(indexes) > 0 {
			count = setNumbers(count, indexes, i, galaxyMap)
		}
	}
	return count
}

func getEmptyRows(lines [][]string) (rowNumbers []int) {
	for i, line := range lines {
		indexes := findCharOccurrences(line, "#")
		if len(indexes) == 0 {
			rowNumbers = append(rowNumbers, i)
		}
	}
	return
}

func expandGalaxy(lines [][]string, rowNumbers []int) [][]string {
	tmp := lines
	emptyLine := lib.InitStringSlice(len(lines[0]), ".")
	count := 0

	for _, rowNumber := range rowNumbers {
		tmp = lib.InsertIntoSliceAtIndex(tmp, emptyLine, (rowNumber + count))
		count++
	}
	return tmp
}

func setNumbers(count int, indexes []int, row int, lines [][]string) int {
	newCount := count
	for _, index := range indexes {
		newCount++
		lines[row][index] = strconv.Itoa(newCount)
	}
	return newCount
}

func findCharOccurrences(slice []string, char string) (indexes []int) {
	for i, v := range slice {
		if v == char {
			indexes = append(indexes, i)
		}
	}
	return
}
