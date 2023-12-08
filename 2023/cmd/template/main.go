package main

import (
	"fmt"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input := lib.Init()
	defer lib.Close(startTime)
	puzzle1Res := 0
	puzzle2Res := 0

	fmt.Printf("Puzzle 1: %d\n", puzzle1Res)
	fmt.Printf("Puzzle 2: %d\n", puzzle2Res)
}
