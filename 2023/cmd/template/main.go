package main

import (
	"fmt"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

	fmt.Printf("Puzzle 1: %d\n", puzzle1Res)
	fmt.Printf("Puzzle 2: %d\n", puzzle2Res)
}
