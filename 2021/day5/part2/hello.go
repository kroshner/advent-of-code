package main

import (
	inputModule "example/hello/input"
	"fmt"
)

func main() {
	ventLines := inputModule.ReadVentLines()
	fmt.Println(len(ventLines))
	coordidatesMap := make(map[int]int)
	for _, ventLine := range ventLines {
		ventLineCoordinates, ventLineIsSupported := ventLine.ToCoordinates()
		if !ventLineIsSupported {
			continue
		}
		for _, coordinate := range ventLineCoordinates {
			coordinatePosition := coordinate.ToPosition()
			coordinateInMapValue, coordinateInMapFound := coordidatesMap[coordinatePosition]
			if !coordinateInMapFound {
				coordidatesMap[coordinatePosition] = 1
			} else {
				coordidatesMap[coordinatePosition] = coordinateInMapValue + 1
			}
		}
	}
	results := 0
	for _, coordinatesMapValue := range coordidatesMap {
		if coordinatesMapValue > 1 {
			results++
		}
	}
	fmt.Println(results)
}
