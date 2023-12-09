package main

import (
	"fmt"
	"strings"

	"github.com/gammazero/deque"

	"github.com/Jorrit05/advent-of-code/pkg/lib"
)

func Graphs(input *lib.PuzzleInput, startPositions []string, directions []string) {
	graph := lib.NewGraph()

	for i := 2; i < len(input.StringLines); i++ {
		fields := strings.Fields(input.StringLines[i])
		graph.AddNode(fields[0])
	}

	for i := 2; i < len(input.StringLines); i++ {
		fields := strings.Fields(input.StringLines[i])
		fst := strings.Trim(fields[2], "(,")
		snd := strings.Trim(fields[3], "(,)")
		graph.AddEdge(fields[0], fst, true)
		graph.AddEdge(fields[0], snd, false)
	}
	res := bfs(graph, startPositions, directions)
	fmt.Println(res)
}

func bfs(graph *lib.Graph, startPositions []string, directions []string) int {
	var currentPositions []*lib.Node
	stateMap := make(map[string]int)

	for _, startPos := range startPositions {
		currentPositions = append(currentPositions, graph.Nodes[startPos])
		stateMap[graph.Nodes[startPos].Name] = 0
	}

	step := 0
	maxInstructionLength := len(directions)
	instructionCounter := 0

	for len(currentPositions) > 0 {
		nextPositions := []*lib.Node{}

		for _, current := range currentPositions {
			var next *lib.Node
			if directions[instructionCounter] == "L" {
				next = current.Left
			} else {
				next = current.Right
			}

			if next != nil {
				nextPositions = append(nextPositions, next)
				stateMap[next.Name] = step + 1
			}
		}

		if nextPositions[0].Name[2] == 'Z' && checkAllNodesEndWithZ(nextPositions) {
			return step + 1
		}

		currentPositions = nextPositions
		step++
		instructionCounter = (instructionCounter + 1) % maxInstructionLength
	}

	return -1
}

func checkAllNodesEndWithZ(nodes []*lib.Node) bool {
	for _, node := range nodes {
		if node.Name[2] != 'Z' {
			return false
		}
	}
	return true
}

func bfs2(graph *lib.Graph, startPositions []string, directions []string) int {
	var q deque.Deque[*lib.Node]
	stateMap := make(map[string]int)

	for _, startPos := range startPositions {
		q.PushBack(graph.Nodes[startPos])
		stateMap[graph.Nodes[startPos].Name] = 0
	}
	instructionCounter := 0
	maxInstructionLength := len(directions)

	for {
		queueLength := q.Len()
		if queueLength == 0 {
			break
		}

		for i := 0; i < queueLength; i++ {
			current := q.PopFront()
			currentStep := stateMap[current.Name]

			if current.Name[2] == 'Z' && checkAllNodes(stateMap) {
				return currentStep
			}
			if directions[instructionCounter] == "L" {
				_, ok := stateMap[current.Left.Name]
				if !ok || stateMap[current.Left.Name] > currentStep+1 {
					stateMap[current.Left.Name] = currentStep + 1
					q.PushBack(current.Left)
				}
			} else {
				_, ok := stateMap[current.Right.Name]
				if !ok || stateMap[current.Right.Name] > currentStep+1 {
					stateMap[current.Right.Name] = currentStep + 1
					q.PushBack(current.Right)
				}
			}
			if instructionCounter == maxInstructionLength-1 {
				instructionCounter = 0
			}
		}
	}
	return -1
}

// function ConcurrentBFS(graph, startNodes)
//     create a queue Q
//     create a map M to store the state of each path (node and step count)
//     for each node in startNodes
//         add node to Q
//         add node to M with step count 0

//     while Q is not empty
//         size = length of Q
//         for i from 0 to size
//             current = Q.dequeue()
//             currentStep = M[current]

//             if current ends with 'Z' and all nodes in M end with 'Z'
//                 return currentStep

//             for each node n that is adjacent to current (based on the instruction)
//                 if n is not in M or M[n] > currentStep + 1
//                     M[n] = currentStep + 1
//                     enqueue n onto Q

//     return failure

func checkAllNodes(stateMap map[string]int) bool {
	for k, _ := range stateMap {
		if k[2] != 'Z' {
			return false
		}
	}
	return true
}
