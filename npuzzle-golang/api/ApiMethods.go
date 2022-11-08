package api

import (
	"N-puzzle-GO/globalvars"
	"N-puzzle-GO/pkg/algo"
	"N-puzzle-GO/pkg/board"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var Directions = map[string]int{"UP": 0, "DOWN": 1, "LEFT": 2, "RIGHT": 3}

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
	//arr := []int{6, 9, 8, 3, 1, 5, 7, 4, 2} // also work
	//arr := []int{1, 2, 3, 8, 9, 7, 4, 5, 6} // dont work (unsolvable)
	//for i := range arr { // some crutch because I do not want to debug react, sorry
	//	arr[i] = arr[i] - 1
	//}
	globalvars.ALGO_END = false // if previous GET request not 200

	tilesNum := context.Param("tilesNum")
	size, err := strconv.Atoi(tilesNum)
	log.Println(size)
	if err != nil {
		log.Println(err)
	}

	size = int(math.Sqrt(float64(size)))
	he := context.Param("he")
	solvable, err := strconv.ParseBool(context.Param("solvable"))
	if err != nil {
		log.Println(err)
	}
	iterations := 100
	arr := makePuzzle(size, solvable, iterations)
	result := BuildCorrectResult(size)
	isSolvable := CheckSolvable(arr, result, size)
	if !isSolvable {
		fmt.Fprintf(os.Stderr, "This puzzle IS NOT solvable\n")
		context.AbortWithStatus(202)
	}
	replaceDigit(arr, 0, 9)
	log.Println(arr)
	for i := range arr { // some crutch because I do not want to debug react, sorry
		arr[i] = arr[i] - 1
	}
	context.JSON(200, arr)
	for i := range arr { // some crutch because I do not want to debug react, sorry
		arr[i] = arr[i] + 1
	}
	globalvars.GenState = arr
	globalvars.Heuristic = he
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

func CheckSolvable(givenTab []int, result []int, size int) bool {
	if size%2 == 1 {
		if countInversions(givenTab)%2 == countInversions(result)%2 {
			return true
		}
	} else {
		rowFromBottom := size - (getEmptyTile(givenTab) / size)
		if (size % 4) == 2 {
			rowFromBottom -= 1
		}
		if rowFromBottom%2 == 1 {

			if countInversions(givenTab)%2 != countInversions(result)%2 {
				return true
			}
		} else {

			if countInversions(givenTab)%2 == countInversions(result)%2 {
				return true
			}
		}
	}
	return false
}

func countInversions(tab []int) int {
	inversions := 0
	for i, _ := range tab {
		for a := i + 1; a < len(tab); a++ {
			if tab[a] != 0 && tab[i] != 0 && tab[i] > tab[a] {
				inversions++
			}
		}
	}
	return int(inversions)
}

func getEmptyTile(tab []int) int {
	size := int(math.Sqrt(float64(len(tab))))
	var i int = 0
	for i = 0; i < size*size; i++ {
		if tab[i] == int(0) {
			return (i)
		}
	}
	return -1
}

func BuildCorrectResult(size int) []int {
	tab := make([]int, size*size)

	var start int = 0
	step_len := size
	for step_len > 0 {
		start = buildCrown(tab, start, step_len, size)
		step_len -= 2
	}
	i := 0
	for tab[i] < size*size {
		i++
	}
	tab[i] = 0
	return (tab)
}

func buildCrown(tab []int, step_start int, step_len int, total_len int) int {
	// Fill top line
	offset := (total_len - step_len) / 2
	block_start := offset*total_len + offset
	val := step_start + 1
	for i := block_start; i < block_start+step_len; i++ {
		tab[i] = val
		val++
	}
	// Fill right column
	val--
	step_start = block_start + step_len - 1
	for i := step_start; i < step_start+total_len*step_len; i += total_len {
		tab[i] = val
		val++
	}
	// Fill bottom line
	val--
	step_start = step_start + total_len*(step_len-1)
	for i := step_start; i > step_start-step_len; i-- {
		tab[i] = val
		val++
	}
	// Fill left line
	val--
	step_start = step_start - (step_len - 1)
	for i := step_start; i > block_start; i -= total_len {
		tab[i] = val
		val++
	}
	val--
	return (val)
}

func makePuzzle(size int, solvable bool, iterations int) []int {
	tab := BuildCorrectResult(size)
	for iter := 0; iter < iterations; iter++ {
		r := GetRandomNumber(4)
		if MoveIsValid(tab, r) {
			tab = Move(tab, r)
		} else {
			iter--
		}
	}
	if solvable == false {
		if tab[0] != 0 && tab[1] != 0 {
			tab[0], tab[1] = tab[1], tab[0]
		} else {
			tab[len(tab)-1], tab[len(tab)-2] = tab[len(tab)-2], tab[len(tab)-1]
		}
	}
	return tab
}

func GetRandomNumber(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return (r1.Intn(max))
}

func MoveIsValid(tab []int, dir int) bool {
	size := int(math.Sqrt(float64(len(tab))))
	empty := GetEmptyTile(tab)

	if dir == Directions["UP"] && empty >= size {
		return (true)
	} else if dir == Directions["DOWN"] && empty < size*(size-1) {
		return (true)
	} else if dir == Directions["LEFT"] && empty%size != 0 {
		return (true)
	} else if dir == Directions["RIGHT"] && empty%size != size-1 {
		return (true)
	}
	return (false)
}

func Move(tab []int, dir int) []int {
	size := int(math.Sqrt(float64(len(tab))))
	var dst = 0
	new := make([]int, len(tab))
	copy(new, tab)
	src := GetEmptyTile(new)
	if dir == Directions["UP"] {
		dst = src - size
	} else if dir == Directions["DOWN"] {
		dst = src + size
	} else if dir == Directions["LEFT"] {
		dst = src - 1
	} else if dir == Directions["RIGHT"] {
		dst = src + 1
	}
	new[src], new[dst] = new[dst], new[src]
	return (new)
}

func GetEmptyTile(tab []int) int {
	size := int(math.Sqrt(float64(len(tab))))
	var i = 0
	for i = 0; i < size*size; i++ {
		if tab[i] == 0 {
			return (i)
		}
	}
	return -1
}

func replaceDigit(arr []int, digitToFind int, digitToReplace int) {
	for k, v := range arr {
		if v == digitToFind {
			arr[k] = digitToReplace
		}
	}
}
