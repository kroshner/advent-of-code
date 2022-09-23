package game

import boardModule "example/hello/board"

type Game struct {
	Boards []*boardModule.Board
}

func CreateGame() *Game {
	return &Game{Boards: make([]*boardModule.Board, 0)}
}

func (game *Game) AddBoard(board *boardModule.Board) {
	game.Boards = append(game.Boards, board)
}

func (game *Game) CheckNumberAndAssertWin(numberToCheck int) int {
	for _, board := range game.Boards {
		board.CheckNumber(numberToCheck)
		if board.AssertWin() {
			return board.SumUnchecked * numberToCheck
		}
	}
	return -1
}
