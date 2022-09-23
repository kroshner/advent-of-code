package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getRepeatsCount() []int {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	repeatsCount := make([]int, 12)

	for scanner.Scan() {
		currentTextValue := scanner.Text()
		for charIndex, charValue := range currentTextValue {
			if charValue == '1' {
				repeatsCount[charIndex]++
			}
			if charValue == '0' {
				repeatsCount[charIndex]--
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return repeatsCount
}

func repeatsCountToBinaryNumber(repeatsCount []int) string {
	str := ""
	for _, repeat := range repeatsCount {
		if repeat > 0 {
			str = str + "1"
		} else if repeat < 0 {
			str = str + "0"
		}
	}
	return str
}

func binaryStringNumberToNumber(binaryStringNumber string) int64 {
	output, err := strconv.ParseInt(binaryStringNumber, 2, 64)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	repeatsCount := getRepeatsCount()
	gammaRateString := repeatsCountToBinaryNumber(repeatsCount)
	gammaRateNumber := binaryStringNumberToNumber(gammaRateString)
	epsilon := gammaRateNumber ^ (0b111111111111)
	result := gammaRateNumber * epsilon
	fmt.Println(result)
}
