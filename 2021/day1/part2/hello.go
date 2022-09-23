package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInputToArr() int {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	timesMore := 0
	valuesWindow := make([]int, 0)

	// initialize first window
	for i := 0; i < 3; i++ {
		scanner.Scan()
		textValue := scanner.Text()
		textValueAsNumber, textValueAsNumberError := strconv.Atoi(textValue)
		if textValueAsNumberError != nil {
			panic(textValueAsNumberError)
		}
		valuesWindow = append(valuesWindow, textValueAsNumber)
	}

	for scanner.Scan() {
		currentTextValue := scanner.Text()
		currentValue, err := strconv.Atoi(currentTextValue)
		if err != nil {
			panic(err)
		}
		if currentValue > valuesWindow[0] {
			timesMore++
		}
		valuesWindow = append(valuesWindow[1:], currentValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return timesMore
}

func main() {
	test := readInputToArr()
	fmt.Println(test)
}
