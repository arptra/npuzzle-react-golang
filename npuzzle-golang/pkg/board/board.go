package board

import (
	"fmt"
)

type EmptyTileMove struct {
	Move string
	X    int
	Y    int
}

var MoveUP = EmptyTileMove{"UP", 0, -1}
var MoveDown = EmptyTileMove{"DOWN", 0, 1}
var MoveLeft = EmptyTileMove{"LEFT", -1, 0}
var MoveRight = EmptyTileMove{"RIGHT", 1, 0}
var InitMove = EmptyTileMove{"NULL", 0, 0}

//var MoveUP = EmptyTileMove{"N", 0, -1}
//var MoveDown = EmptyTileMove{"S", 0, 1}
//var MoveLeft = EmptyTileMove{"W", -1, 0}
//var MoveRight = EmptyTileMove{"E", 1, 0}
//var InitMove = EmptyTileMove{"NULL", 0, 0}

// var Moves = []EmptyTileMove{MoveUP, MoveDown, MoveLeft, MoveRight}
var Moves = []EmptyTileMove{MoveUP, MoveRight, MoveDown, MoveLeft}

type StateOfBoard struct {
	Size                  int // board Size
	EmptyTile             EmptyTileMove
	PrevEmptyTilePosition int
	EmptyTilePosition     int
	CurrentBoardState     []int
	From                  *StateOfBoard
	To                    *StateOfBoard
}

func generate(w int, h int, x int, y int) int {
	if y != 0 {
		return w + generate(h-1, w, y-1, w-x-1)
	} else {
		return x + 1
	}
}

func GetGoalState(size int) []int {
	goal := make([]int, 0)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			goal = append(goal, generate(size, size, x, y))
		}
	}
	return goal
}

func shiftTile(position EmptyTileMove, size int) int {
	return position.X + (position.Y * size)
}

func GetState(state StateOfBoard) ([]int, int) {
	nextState := make([]int, state.Size*state.Size)
	copy(nextState, state.CurrentBoardState)
	if (state.EmptyTilePosition%state.Size)+state.EmptyTile.X < 0 || (state.EmptyTilePosition%state.Size)+state.EmptyTile.X >= state.Size ||
		(state.EmptyTilePosition/state.Size)+state.EmptyTile.Y < 0 || (state.EmptyTilePosition/state.Size)+state.EmptyTile.Y >= state.Size {
		return nil, 0
	}
	oldPos := state.EmptyTilePosition                                          //prev position of empty tile
	newPos := state.EmptyTilePosition + shiftTile(state.EmptyTile, state.Size) //next position of empty tile
	oldVal := nextState[oldPos]
	newVal := nextState[newPos]
	nextState[newPos] = oldVal
	nextState[oldPos] = newVal
	return nextState, newPos
}

func CreateNewState(state StateOfBoard, boardState []int, newPos, oldPos int) StateOfBoard {
	newState := new(StateOfBoard)
	newState.Size = state.Size
	newState.CurrentBoardState = boardState
	newState.PrevEmptyTilePosition = oldPos
	newState.EmptyTilePosition = newPos
	newState.EmptyTile = state.EmptyTile
	return *newState
}

func GetAllState(state StateOfBoard) []StateOfBoard {
	allState := make([]StateOfBoard, 0)
	oldPos := state.EmptyTilePosition
	for _, move := range Moves {
		state.EmptyTile = move
		newState, newPos := GetState(state)
		if (newState) == nil || state.PrevEmptyTilePosition == newPos {
			continue
		} else {
			allState = append(allState, CreateNewState(state, newState, newPos, oldPos))
		}
	}
	return allState
}

func BoardEquals(state1, state2 StateOfBoard) bool {
	a := state1.CurrentBoardState
	b := state2.CurrentBoardState
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (state *StateOfBoard) FindN(n int) (int, error) {
	for ii := 0; ii < state.Size*state.Size; ii += 1 {
		if state.CurrentBoardState[ii] == n {
			return ii, nil
		}
	}
	return 0, fmt.Errorf("%d not in board state", n)
}
