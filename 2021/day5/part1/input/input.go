package input

import (
	"bufio"
	"example/hello/coordinate"
	"os"
	"regexp"
	"strconv"
)

func ReadVentLines() []coordinate.VentLine {
	file, err := os.Open("./exampleNumbers.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	ventLineRegExp := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	ventLines := make([]coordinate.VentLine, 0)
	for scanner.Scan() {
		currentTextLine := scanner.Text()
		ventLineRegExpMatch := ventLineRegExp.FindStringSubmatch(currentTextLine)
		ventLineXStart, ventLineXStartValueErr := strconv.Atoi(ventLineRegExpMatch[1])
		if ventLineXStartValueErr != nil {
			panic(ventLineXStartValueErr)
		}
		ventLineYStart, ventLineYStartValueErr := strconv.Atoi(ventLineRegExpMatch[2])
		if ventLineYStartValueErr != nil {
			panic(ventLineYStartValueErr)
		}
		ventLineXEnd, ventLineXEndValueErr := strconv.Atoi(ventLineRegExpMatch[3])
		if ventLineXEndValueErr != nil {
			panic(ventLineXEndValueErr)
		}
		ventLineYEnd, ventLineYEndValueErr := strconv.Atoi(ventLineRegExpMatch[4])
		if ventLineYEndValueErr != nil {
			panic(ventLineYEndValueErr)
		}
		ventLines = append(ventLines, coordinate.VentLine{Start: coordinate.Coordinate{X: ventLineXStart, Y: ventLineYStart}, End: coordinate.Coordinate{X: ventLineXEnd, Y: ventLineYEnd}})
	}

	return ventLines
}
