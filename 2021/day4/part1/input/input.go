package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputNumbers() []int {
	file, err := os.Open("./inputNumbers.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	currentTextValue := scanner.Text()
	currentTextValueSplitted := strings.Split(currentTextValue, ",")

	var inputNumbers = make([]int, 0)

	for _, i := range currentTextValueSplitted {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputNumbers = append(inputNumbers, j)
	}

	return inputNumbers
}

func ReadInputBoards() [][][]int {
	file, err := os.Open("./inputBoards.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	boardsNumbersContainers := make([][][]int, 0)
	currentBoardNumbersContainers := make([][]int, 0)
	for scanner.Scan() {
		currentTextLine := scanner.Text()
		if strings.Contains(currentTextLine, "52 29 23 97 14") {
			fmt.Println("123")
		}
		if currentTextLine == "" {
			boardsNumbersContainers = append(boardsNumbersContainers, currentBoardNumbersContainers)
			currentBoardNumbersContainers = make([][]int, 0)
			continue
		}
		currentRowNumbers := make([]int, 0)
		currentTextLineTrimmed := strings.TrimSpace(currentTextLine)
		currentTextLineTrimmedAndSingleSpaced := strings.ReplaceAll(currentTextLineTrimmed, "  ", " ")
		currentTextLineSplitted := strings.Split(currentTextLineTrimmedAndSingleSpaced, " ")
		for _, test := range currentTextLineSplitted {
			currentValue, err := strconv.Atoi(test)
			if err != nil {
				panic(err)
			}
			currentRowNumbers = append(currentRowNumbers, currentValue)
		}
		currentBoardNumbersContainers = append(currentBoardNumbersContainers, currentRowNumbers)
	}
	boardsNumbersContainers = append(boardsNumbersContainers, currentBoardNumbersContainers)
	currentBoardNumbersContainers = make([][]int, 0)
	return boardsNumbersContainers
}
