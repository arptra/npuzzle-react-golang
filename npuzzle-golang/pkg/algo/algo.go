package algo

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"math"
)

var FOUND = false

func AlgoStart(start, goal board.StateOfBoard) int {
	fmt.Println("Algo start")
	threshold := ManhattanDist(start, goal)
	for {
		temp := Search(start, goal, 0, threshold) //function search(node,g score,threshold)
		if temp == -1 {
			return -1
		}
		fmt.Println(threshold)
		if FOUND { // if goal FOUND
			FOUND = false // for SERVER_MODE == true
			return 1
		}
		if temp == math.MaxInt32 { //Threshold larger than maximum possible f value
			return 0
		} //or set Time limit exceeded
		threshold = temp
	}
}

func Search(node, goal board.StateOfBoard, g, threshold float64) float64 {
	f := g + ManhattanDist(node, goal)
	if f > threshold {
		return f
	}
	if board.BoardEquals(node, goal) {
		FOUND = true
		path := getRecursivePath(node)
		if globalvars.SERVER_MODE == false {
			printPath(path)
		} else {
			savePath(path)
			printPath(path)
		}
		return f
	}
	min := math.MaxFloat64
	for _, newNode := range board.GetAllState(node) {
		if globalvars.STOP_CALC == true {
			return -1
		}
		node.To = &newNode
		newNode.From = &node
		temp := Search(newNode, goal, g+1, threshold)
		if board.BoardEquals(newNode, goal) {
			return f
		}
		if temp < min {
			min = temp
		}
	}
	return min
}
