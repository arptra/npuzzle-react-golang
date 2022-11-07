package main

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/apiserver"
	"N-puzzle-GO/pkg/board"
	"log"
	"os"
)

func manual() {
	//firstState := []int{6, 9, 8, 3, 1, 5, 7, 4, 2} // doesnt-work
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1} // work
	//firstState := []int{2, 9, 7, 1, 3, 5, 8, 4, 6} // work too
	//firstState := []int{2, 7, 3, 6, 4, 8, 5, 1, 9} // index out of range (if last index is equal 9 - error exception)

	firstState := []int{4, 2, 8, 6, 5, 1, 9, 7, 3}
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}
	//firstState := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}

	State := board.StateOfBoard{
		3,
		board.InitMove,
		8,
		8,
		firstState,
		nil,
		nil,
	}
	//test := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	//fmt.Println(firstState[len(firstState)-1])
	//firstState[len(firstState)-1],
	goalState := board.StateOfBoard{
		3,
		board.InitMove,
		8,
		8,
		board.GetGoalState(3),
		//test,
		nil,
		nil,
	}
	board.PrintState(State)
	board.PrintState(goalState)
	//all := board.GetAllState(State)
	//for _, v := range all {
	//	board.PrintState(v)
	//}
	algo.AlgoStart(State, goalState)
}

func main() {
	if len(os.Args) > 2 || (len(os.Args) == 2 && os.Args[1] != "manual") {
		log.Fatal("usage: ./npuzzle manual | ./npuzzle")
	}
	if len(os.Args) == 1 {
		globalvars.SERVER_MODE = true
		apiserver.ApiServerStart()
	} else {
		manual()
	}
}
