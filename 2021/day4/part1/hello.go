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
		checkResult := game.CheckNumberAndAssertWin(inputNumber)
		if checkResult != -1 {
			fmt.Println(checkResult)
			return
		}
	}
	panic("No answer found")
}
