package algo

import (
	"N-puzzle-GO/pkg/board"
	"fmt"
	"math"
)

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
