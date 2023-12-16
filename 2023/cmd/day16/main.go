package main

import (
	"fmt"
	"slices"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
	"github.com/gammazero/deque"
)

var directions = map[string][]int{
	"N": {-1, 0},
	"E": {0, 1},
	"S": {1, 0},
	"W": {0, -1},
}

var MaxLength = 0
var MaxWidth = 0

var visitedSet = lib.Set[lib.MazeLocation]{}
func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init(true)
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	MaxWidth = len(input.Lines[0]) -1
	MaxLength = len(input.Lines)-1

	start,_ := getStartPosition(0,0,input.Lines, "E")
	puzzle1Res = energize(input.Lines, start)

	results := []int{}
	startPositions := []*lib.MazeLocation{}

	for i := range input.Lines[0] {
		startPos, secondStart :=  getStartPosition(0,i,input.Lines, "S")
		if secondStart != nil && secondStart.Direction != "" {
			startPositions = append(startPositions, secondStart)
		}
		startPositions = append(startPositions, startPos)
	}
	for i := range input.Lines[MaxLength-1] {
		startPos, secondStart :=  getStartPosition(MaxLength ,i,input.Lines, "N")
		if secondStart != nil && secondStart.Direction != "" {
			startPositions = append(startPositions, secondStart)
		}
		startPositions = append(startPositions, startPos)
	}


	for i := range input.Lines {
		startPos, secondStart :=  getStartPosition(i,0,input.Lines, "E")
		if secondStart != nil && secondStart.Direction != "" {
			startPositions = append(startPositions, secondStart)
		}
		startPositions = append(startPositions, startPos)
	}

	for i := range input.Lines {
		startPos, secondStart :=  getStartPosition(i,MaxWidth,input.Lines, "W")
		if secondStart != nil && secondStart.Direction != "" {
			startPositions = append(startPositions, secondStart)
		}
		startPositions = append(startPositions, startPos)
	}

	for _, startPos := range startPositions {
		visitedSet = lib.Set[lib.MazeLocation]{}
		res := energize(input.Lines, startPos)
		results = append(results, res)
	}

	puzzle2Res = slices.Max(results)
}

func energize(input [][]string, start *lib.MazeLocation) int {
	energizedSet := lib.Set[lib.Coordinate]{}

	var q deque.Deque[*lib.MazeLocation]
	q.PushBack(start)
	for {
		queueLength := q.Len()

		if queueLength == 0 {
			break
		}
		for i := 0; i < queueLength; i++ {
			current := q.PopFront()
			energizedSet.Add(current.Coordinate)

			next, err := getNext(current)
			if err != nil {
				continue
			}
			secondDirection := determineNewDirections(current, next, input)

			if !visitedSet.Has(*next) {
				q.PushBack(next)
			}

			if secondDirection != nil && secondDirection.Direction != ""{
				if !visitedSet.Has(*secondDirection) {
					q.PushBack(secondDirection)
				}
			}
			visitedSet.Add(*current)
		}
	}
	return len(energizedSet)
}

var backslashMap = map[string]string {
	"N"  : "W",
	"E"  : "S",
	"S"  : "E",
	"W" : "N",
}

var forwardslashMap = map[string]string {
	"N"  : "E",
	"E"  : "N",
	"S"  : "W",
	"W" : "S",
}

func determineNewDirections(current *lib.MazeLocation, next *lib.MazeLocation, maze [][]string) *lib.MazeLocation{
	char := maze[next.Coordinate.Row][next.Coordinate.Col]
	copy := &lib.MazeLocation{}

	switch char {
	case ".":
			next.Direction = current.Direction
	case "\\":
			next.Direction = backslashMap[current.Direction]

	case "/":
			next.Direction = forwardslashMap[current.Direction]

	case "-":
			if current.Direction == "W" || current.Direction == "E" {
					next.Direction = current.Direction
			} else {
				next.Direction = "W"
				copy = lib.CopyMazeLocation(next)
				copy.Direction = "E"
			}

	case "|":
			if current.Direction == "N" || current.Direction == "S" {
					next.Direction = current.Direction
			} else {
					next.Direction = "N"
					copy = lib.CopyMazeLocation(next)
					copy.Direction = "S"
			}
	}

	return copy
}

func getNext(current *lib.MazeLocation) (*lib.MazeLocation, error) {
	direction := directions[current.Direction]
	newCoord := &lib.Coordinate{}
	newCoord.Row = current.Coordinate.Row + direction[0]
	newCoord.Col = current.Coordinate.Col + direction[1]

	if newCoord.Row < 0 || newCoord.Row > MaxLength ||
			newCoord.Col < 0  || newCoord.Col >  MaxWidth {
				return &lib.MazeLocation{},fmt.Errorf("index out of range")
	}

	return &lib.MazeLocation{
		Coordinate: *newCoord,
	}, nil
}

func getStartPosition(row, col int, maze [][]string, startPosition string) (*lib.MazeLocation,*lib.MazeLocation) {
	current := &lib.MazeLocation{
		Coordinate: lib.Coordinate{
			Row: row,
			Col: col,
		},
		Direction: startPosition,
	}

	copy := determineNewDirections(current, current, maze)
	if copy.Direction != "" {
		return current,copy
	}
	return current, nil
}
