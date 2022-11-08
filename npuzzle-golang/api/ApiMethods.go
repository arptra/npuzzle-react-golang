package api

import (
	"N-puzzle-GO/globalvars"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	arr := []int{8, 2, 6, 3, 9, 4, 7, 5, 1}

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
}

//func StartAlgo(context *gin.Context) {
//	var arr []int
//	var emptyTileIndex int
//
//	context.AbortWithStatus(200)
//	globalvars.ALGO_END = false // if previous GET request not 200
//	arr = globalvars.InputState[globalvars.InputStateKey]
//	emptyTile := globalvars.InputState[globalvars.EmptyTileKey][0]
//	for i := range arr {
//		if emptyTile == arr[i] {
//			emptyTileIndex = i
//		}
//	}
//	for i := range arr { // some crutch because I do not want to debug react, sorry
//		arr[i]++
//	}
//	fmt.Println(arr)
//	state := board.StateOfBoard{
//		int(math.Sqrt(float64(len(arr)))),
//		board.InitMove,
//		emptyTileIndex,
//		emptyTileIndex,
//		arr,
//		nil,
//		nil,
//	}
//	goalState := board.StateOfBoard{
//		int(math.Sqrt(float64(len(arr)))),
//		board.InitMove,
//		emptyTileIndex,
//		emptyTileIndex,
//		board.GetGoalState(int(math.Sqrt(float64(len(arr))))),
//		nil,
//		nil,
//	}
//	board.PrintState(state)
//	board.PrintState(goalState)
//
//	status := algo.AlgoStart(state, goalState)
//	if status == -1 {
//		globalvars.STOP_CALC = false
//	}
//}

func PutState(context *gin.Context) {
	if err := context.ShouldBindJSON(&globalvars.InputState); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.AbortWithStatus(200)
		return
	}
}
