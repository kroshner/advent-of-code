package board

type NumbersContainer = []int

type Board struct {
	SumUnchecked        int
	SunUncheckedInitial int
	numbersContainer    []NumbersContainer
}

func (board *Board) CheckNumber(numberToCheck int) {
	matchedNumber := false
	for _, container := range board.numbersContainer {
		for numberIndex, number := range container {
			if number == numberToCheck {
				matchedNumber = true
				container[numberIndex] = 0
			}
		}
	}
	if matchedNumber {
		board.SumUnchecked -= numberToCheck
	}
}

func (board *Board) AssertWin() bool {
	for _, container := range board.numbersContainer {
		containerSum := 0
		for _, number := range container {
			containerSum += number
		}
		if containerSum == 0 {
			return true
		}
	}
	return false
}

func CreateBoard(numbersInputs [][]int) *Board {
	sumUnchecked := 0
	for _, container := range numbersInputs {
		for _, number := range container {
			sumUnchecked += number
		}
	}
	numbersContainer := make([]NumbersContainer, 0)
	inputDimensionLength := len(numbersInputs)
	currentContainer := make(NumbersContainer, 0)
	for i := 0; i < inputDimensionLength; i++ {
		currentContainer = make(NumbersContainer, 0)
		for j := 0; j < inputDimensionLength; j++ {
			currentContainer = append(currentContainer, numbersInputs[i][j])
		}
		numbersContainer = append(numbersContainer, currentContainer)
		currentContainer = make(NumbersContainer, 0)
		for j := 0; j < inputDimensionLength; j++ {
			currentContainer = append(currentContainer, numbersInputs[j][i])
		}
		numbersContainer = append(numbersContainer, currentContainer)
		// currentContainer = append(currentContainer, numbersInputs)
	}
	return &Board{SumUnchecked: sumUnchecked, SunUncheckedInitial: sumUnchecked, numbersContainer: numbersContainer}
}
