package algo

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"math"
)

var manhattan = "Manhattan"
var euclidean = "Euclidean"
var hamming = "Hamming"

func ManhattanDist(state, goal board.StateOfBoard) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var total float64
	total = 0
	for ii := range state.CurrentBoardState {
		if state.CurrentBoardState[ii] != 9 {
			n, err := goal.FindN(state.CurrentBoardState[ii])
			if err != nil {
				panic(err)
			}
			total += math.Abs(float64(ii%state.Size-n%state.Size)) +
				math.Abs(float64(ii/state.Size-n/state.Size))
		}
	}
	return total
}

func EuclideanDistance(tab []int, result []int) int {
	var dist int
	var i int
	var destIndex int
	var distRow float64
	var distCol float64

	inverseGoal := invert(result)
	size := int(math.Sqrt(float64(len(tab))))
	for i = 0; i < int(len(tab)); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			distRow = math.Pow(float64(i/size-destIndex/size), 2)
			distCol = math.Pow(float64(i%size-destIndex%size), 2)
			dist += int(math.Sqrt(distRow + distCol))
		}
	}
	return dist
}

func HammingDistance(tab []int, result []int) int {
	inverseGoal := invert(result)
	var i int
	var misplaced int
	var destIndex int
	for i = 0; i < len(tab); i++ {
		if tab[i] != 0 {
			destIndex = inverseGoal[tab[i]]
			if destIndex != i {
				misplaced++
			}
		}
	}
	return misplaced
}

func invert(tab []int) []int {
	result := append(tab[:0:0], tab...)
	var i = 0
	for i = 0; i < len(result); i++ {
		result[tab[i]] = i
	}
	return result
}

func getScore(state, goal board.StateOfBoard, heuristic string) float64 {
	score := 0.0
	if heuristic == manhattan {
		return ManhattanDist(state, goal)
	}
	replaceDigit(state.CurrentBoardState, globalvars.EmptyTileNumber, 0)
	replaceDigit(goal.CurrentBoardState, globalvars.EmptyTileNumber, 0)
	if heuristic == euclidean {
		score = float64(EuclideanDistance(state.CurrentBoardState, goal.CurrentBoardState))
	}
	if heuristic == hamming {
		score = float64(HammingDistance(state.CurrentBoardState, goal.CurrentBoardState))
	}
	replaceDigit(state.CurrentBoardState, 0, globalvars.EmptyTileNumber)
	replaceDigit(goal.CurrentBoardState, 0, globalvars.EmptyTileNumber)
	return score
}

func replaceDigit(arr []int, digitToFind int, digitToReplace int) {
	for k, v := range arr {
		if v == digitToFind {
			arr[k] = digitToReplace
		}
	}
}
