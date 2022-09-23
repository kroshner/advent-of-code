package main

import (
	boardModule "example/hello/board"
	gameModule "example/hello/game"
	inputModule "example/hello/input"
	"fmt"
)

func main() {
	inputNumbers := inputModule.ReadInputNumbers()
	inputBoardsNumbersContainers := inputModule.ReadInputBoards()
	game := gameModule.CreateGame()
	for _, inputBoardNumbersContainer := range inputBoardsNumbersContainers {
		newBoard := boardModule.CreateBoard(inputBoardNumbersContainer)
		game.AddBoard(newBoard)
	}
	for _, inputNumber := range inputNumbers {
		game.CheckNumber(inputNumber)
	}
	if game.LastWinningBoardIndex == -1 || game.WinningNumber == -1 {
		panic("No answer found")
	}
	result := game.Boards[game.LastWinningBoardIndex].SumUnchecked * game.WinningNumber
	fmt.Println(result)
}
