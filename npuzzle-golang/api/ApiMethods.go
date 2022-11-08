package api

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"strconv"
)

func GetStop(context *gin.Context) {
	if globalvars.ALGO_END == false {
		globalvars.STOP_CALC = true
	} else {
		globalvars.STOP_CALC = false
	}
}

func GetPath(context *gin.Context) {
	if globalvars.ALGO_END == true {
		context.JSON(200, globalvars.SuccessPath)
	} else {
		context.AbortWithStatus(202)
	}
}

func GetState(context *gin.Context) {
	//arr := []int{8, 2, 6, 3, 9, 4, 7, 5, 1} // work
	arr := []int{6, 9, 8, 3, 1, 5, 7, 4, 2} // also work
	//arr := []int{1, 2, 3, 8, 9, 7, 4, 5, 6} // dont work (unsolvable)

	for i := range arr { // some crutch because I do not want to debug react, sorry
		arr[i] = arr[i] - 1
	}
	globalvars.ALGO_END = false // if previous GET request not 200

	tilesNum := context.Param("tilesNum")
	num, err := strconv.Atoi(tilesNum)
	if err != nil {
		log.Println(err)
	}
	//if num ...
	num = num
	he := context.Param("he")
	he = he
	//if he ...
	solvable := context.Param("solvable")
	solvable = solvable
	//if solvable ...
	context.JSON(200, arr)
	for i := range arr { // some crutch because I do not want to debug react, sorry
		arr[i] = arr[i] + 1
	}
	globalvars.GenState = arr
}

func StartAlgo(context *gin.Context) {
	arr := globalvars.GenState

	fmt.Println(arr)
	state := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		len(arr) - 3,
		len(arr) - 3,
		arr,
		nil,
		nil,
	}
	goalState := board.StateOfBoard{
		int(math.Sqrt(float64(len(arr)))),
		board.InitMove,
		len(arr) - 3,
		len(arr) - 3,
		board.GetGoalState(int(math.Sqrt(float64(len(arr))))),
		nil,
		nil,
	}
	board.PrintState(state)
	board.PrintState(goalState)

	status := algo.AlgoStart(state, goalState)
	if status == -1 {
		globalvars.STOP_CALC = false
	}
}
