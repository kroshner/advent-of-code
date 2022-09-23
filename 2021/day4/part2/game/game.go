package game

import boardModule "example/hello/board"

type Game struct {
	Boards                      []*boardModule.Board
	boardsIndexesAlreadyChecked []int
	LastWinningBoardIndex       int
	WinningNumber               int
}

func CreateGame() *Game {
	return &Game{Boards: make([]*boardModule.Board, 0), LastWinningBoardIndex: -1, WinningNumber: -1, boardsIndexesAlreadyChecked: make([]int, 0)}
}

func (game *Game) AddBoard(board *boardModule.Board) {
	game.Boards = append(game.Boards, board)
}

func (game *Game) CheckNumber(numberToCheck int) {
	for boardIndex, board := range game.Boards {
		if sliceHasElement(game.boardsIndexesAlreadyChecked, boardIndex) {
			continue
		}
		board.CheckNumber(numberToCheck)
		if board.AssertWin() {
			game.WinningNumber = numberToCheck
			game.boardsIndexesAlreadyChecked = append(game.boardsIndexesAlreadyChecked, boardIndex)
			game.LastWinningBoardIndex = boardIndex
		}
	}
}

func sliceHasElement(slice []int, element int) bool {
	for _, sliceElement := range slice {
		if sliceElement == element {
			return true
		}
	}
	return false
}
