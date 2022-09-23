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
	previousValue := 0

	scanner.Scan()
	firstLine := scanner.Text()
	firstLineValue, firstLineErr := strconv.Atoi(firstLine)
	if firstLineErr != nil {
		panic(firstLineErr)
	}
	previousValue = firstLineValue

	for scanner.Scan() {
		currentTextValue := scanner.Text()
		currentValue, err := strconv.Atoi(currentTextValue)
		if err != nil {
			panic(err)
		}
		if currentValue > previousValue {
			timesMore++
		}
		previousValue = currentValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return timesMore
}

func main() {
	inputNumbers := readInputToArr()
	fmt.Println(inputNumbers)
}
