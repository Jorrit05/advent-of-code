package main

import (
	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func main() {
	startTime, input, puzzle1Res, puzzle2Res := lib.Init()
	defer lib.Close(startTime, &puzzle1Res, &puzzle2Res)

}
