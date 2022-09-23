package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInputToArr() int {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	horizontalPosition := 0
	verticalPosition := 0

	r := regexp.MustCompile(`(forward|up|down) (\d+)`)

	for scanner.Scan() {
		currentTextValue := scanner.Text()
		fmt.Println(currentTextValue)
		test := r.FindStringSubmatch(currentTextValue)
		direction := test[1]
		value, valueErr := strconv.Atoi(test[2])
		if valueErr != nil {
			panic(valueErr)
		}
		switch direction {
		case "forward":
			{
				fmt.Println("going forward")
				horizontalPosition += value
			}
		case "up":
			{
				fmt.Println("going up")
				verticalPosition -= value
			}
		case "down":
			{
				fmt.Println("going down")
				verticalPosition += value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return horizontalPosition * verticalPosition
}

func main() {
	test := readInputToArr()
	fmt.Println(test)
}
