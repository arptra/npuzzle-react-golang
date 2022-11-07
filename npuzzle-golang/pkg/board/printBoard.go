package board

import "fmt"

func PrintBoard(board []int, size int) {
	x := 0
	for y := 0; y < (size * size); y++ {
		x++
		fmt.Printf("%d ", board[y])
		if x == size {
			fmt.Println()
			x = 0
		}
	}
	fmt.Println()
}

func PrintState(state StateOfBoard) {
	//fmt.Printf("%+v\n", state)
	PrintBoard(state.CurrentBoardState, state.Size)
}
