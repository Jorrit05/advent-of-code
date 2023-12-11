package main

import (
	"fmt"
	"strings"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

var directions = map[int][]int{
	0: {-1, 0},  // "N"
	1: {0, 1},   // "E"
	2: {1, 0},   // "S"
	3: {0, -1},  // "W"
	4: {-1, -1}, // "NW"
	5: {-1, 1},  // "NE"
	6: {1, 1},   // "SE"
	7: {1, -1},  // "SW"
}
var cameFrom = map[int]int{
	0: 2, // "N"
	1: 3, // "E"
	2: 0, // "S"
	3: 1, // "W"
}

var validDirections = map[int]string{
	0: "|7FS",
	1: "-J7S",
	2: "|JLS",
	3: "-FLS",
}

var inValidDirections = map[string]lib.IntSet{
	"|": {1: struct{}{}, 3: struct{}{}},
	"7": {0: struct{}{}, 1: struct{}{}},
	"F": {0: struct{}{}, 3: struct{}{}},
	"-": {0: struct{}{}, 2: struct{}{}},
	"L": {2: struct{}{}, 3: struct{}{}},
	"J": {1: struct{}{}, 2: struct{}{}},
}

type Coordinate struct {
	Row int
	Col int
}

type Location struct {
	Coordinate Coordinate
	CameFrom   int
	Steps      int
	Pipe       string
	MaxWidth   int
	MaxLength  int
}

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	currentPosition := getStartPosition(input.Lines)

	// Puzzle1
	for i := 1; i > 0; {
		currentPosition = moveToNextPos(currentPosition, input.Lines)

		if currentPosition.Pipe == "S" {
			break
		}
	}
	puzzle1Res = currentPosition.Steps / 2

	// Puzzle2
	sanitizedInput := fillEdges(input.StringLines)
	lib.PrettyPrint(sanitizedInput)

}

func fillEdges(maze []string) [][]string {
	maze[0] = strings.Replace(maze[0], ".", "0", -1)
	maze[len(maze)-1] = strings.Replace(maze[len(maze)-1], ".", "0", -1)
	var lines [][]string

	for _, line := range maze {
		lineSlice := strings.Split(line, "")
		if lineSlice[0] == "." {
			lineSlice[0] = "0"
		}
		if lineSlice[len(lineSlice)-1] == "." {
			lineSlice[len(lineSlice)-1] = "0"
		}
		lines = append(lines, lineSlice)
	}

	maxWidth := len(lines[0]) - 1
	maxLength := len(lines) - 1

	for row, line := range lines {
		for col, char := range line {
			if char == "0" {
				markAllNeighbors(row, col, lines, maxWidth, maxLength)
			}
		}
	}

	return lines
}

func markAllNeighbors(row, col int, lines [][]string, maxWidth, maxLength int) {
	currentPosition := Coordinate{
		Row: row,
		Col: col,
	}

	possibleDirections := lib.IntSet{}
	for i := 0; i < 8; i++ {
		possibleDirections.Add(i)
	}

	inValidDirection := lib.IntSet{}
	if currentPosition.Col == 0 {
		inValidDirection.Add(3)
		inValidDirection.Add(4)
		inValidDirection.Add(7)
	}

	if currentPosition.Col == maxWidth {
		inValidDirection.Add(1)
		inValidDirection.Add(5)
		inValidDirection.Add(6)
	}

	if currentPosition.Row == 0 {
		inValidDirection.Add(0)
		inValidDirection.Add(4)
		inValidDirection.Add(5)
	}

	if currentPosition.Row == maxLength {
		inValidDirection.Add(2)
		inValidDirection.Add(6)
		inValidDirection.Add(7)
	}
	possibleDirections = possibleDirections.Difference(inValidDirection)

	// if row == 8 && col == 6 {
	// 	fmt.Println("break")
	// }
	for direction := range possibleDirections {
		move := directions[direction]
		newRow := currentPosition.Row + move[0]
		newCol := currentPosition.Col + move[1]
		pipe := lines[newRow][newCol]
		if pipe == "." {
			lines[newRow][newCol] = "0"

			markAllNeighbors(newRow, newCol, lines, maxWidth, maxLength)
		}
	}
}

func moveToNextPos(currentPosition *Location, maze [][]string) *Location {
	getValidDirection(currentPosition, maze)
	return currentPosition
}

func getValidDirection(currentPosition *Location, maze [][]string) {
	possibleDirections := lib.IntSet{}
	for i := 0; i < 4; i++ {
		possibleDirections.Add(i)
	}

	inValidDirection := lib.IntSet{}
	if currentPosition.Coordinate.Col == 0 {
		inValidDirection.Add(3)
	}

	if currentPosition.Coordinate.Col == currentPosition.MaxWidth {
		inValidDirection.Add(1)
	}

	if currentPosition.Coordinate.Row == 0 {
		inValidDirection.Add(0)
	}

	if currentPosition.Coordinate.Row == currentPosition.MaxLength {
		inValidDirection.Add(2)
	}

	if currentPosition.CameFrom != -1 {
		inValidDirection.Add(currentPosition.CameFrom)
	}

	// fmt.Println(fmt.Sprintf("%v,  %v", currentPosition.Pipe, inValidDirections[currentPosition.Pipe]))

	inValidDirection = inValidDirection.Union(inValidDirections[currentPosition.Pipe])
	possibleDirections = possibleDirections.Difference(inValidDirection)

	for direction := range possibleDirections {
		move := directions[direction]
		row := currentPosition.Coordinate.Row + move[0]
		col := currentPosition.Coordinate.Col + move[1]
		pipe := maze[row][col]
		if strings.Contains(validDirections[direction], pipe) {
			// correct direction
			currentPosition.Coordinate.Col = col
			currentPosition.Coordinate.Row = row
			currentPosition.Steps += 1
			currentPosition.Pipe = pipe
			currentPosition.CameFrom = cameFrom[direction]
			return
		}
	}
	panic(fmt.Sprintf("Stuck: %v", currentPosition))
}

func getStartPosition(input [][]string) *Location {
	for i, line := range input {
		for j, char := range line {
			if char == "S" {
				return &Location{
					Coordinate: Coordinate{
						Row: i,
						Col: j,
					},
					CameFrom:  -1,
					Steps:     0,
					Pipe:      "S",
					MaxWidth:  len(line) - 1,
					MaxLength: len(input) - 1,
				}
			}
		}
	}
	return &Location{}
}
