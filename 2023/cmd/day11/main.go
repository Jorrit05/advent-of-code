package main

import (
	"math"
	"strconv"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

var coordinateMap = make(map[int]lib.Coordinate)
var baseMultiplier = 1

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init(true)
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	rowsToExpand := getEmptyRows(input.Lines)
	columnsToExpand := getEmptyRows(input.TransposedLines)
	nrOfGalaxies := numberGalaxy(input.Lines)
	galaxyPairs := lib.GenerateIntPairs(nrOfGalaxies)

	puzzle1Res = getGalaxyDistances(galaxyPairs, rowsToExpand, columnsToExpand)

	//puzzle2
	baseMultiplier = 999999
	puzzle2Res = getGalaxyDistances(galaxyPairs, rowsToExpand, columnsToExpand)
}

func getGalaxyDistances(galaxyPairs []lib.Pair[lib.IntCompare], rowSet lib.IntSet, columnSet lib.IntSet) (result int) {
	for _, pair := range galaxyPairs {
		coordinate1 := coordinateMap[int(pair.Fst)]
		coordinate2 := coordinateMap[int(pair.Snd)]
		extraDistanceRow := checkInBetweenExpansions(coordinate1.Row, coordinate2.Row, rowSet)
		extraDistanceCol := checkInBetweenExpansions(coordinate1.Col, coordinate2.Col, columnSet)

		result += calculateDistance(coordinate1, coordinate2) + extraDistanceRow + extraDistanceCol
	}
	return
}

func checkInBetweenExpansions(i1, i2 int, set lib.IntSet) (count int) {
	for i := min(i1, i2); i <= max(i1, i2); i++ {
		if set.Has(i) {
			count += baseMultiplier
		}
	}
	return
}

// calculateDistance calculates the Manhattan distance between two coordinates on a grid.
func calculateDistance(coordinate1, coordinate2 lib.Coordinate) int {
	// Calculate the absolute difference in rows and columns
	rowDiff := math.Abs(float64(coordinate1.Row - coordinate2.Row))
	colDiff := math.Abs(float64(coordinate1.Col - coordinate2.Col))

	// Sum the differences to get the Manhattan distance
	distance := int(rowDiff + colDiff)
	return distance
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

func getEmptyRows(lines [][]string) lib.IntSet {
	rowNumbers := make(lib.IntSet)
	for i, line := range lines {
		indexes := findCharOccurrences(line, "#")
		if len(indexes) == 0 {
			rowNumbers.Add(i)
		}
	}
	return rowNumbers
}

func setNumbers(count int, indexes []int, row int, lines [][]string) int {
	newCount := count
	for _, index := range indexes {
		newCount++
		lines[row][index] = strconv.Itoa(newCount)
		coordinateMap[newCount] = lib.Coordinate{Row: row, Col: index}
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
