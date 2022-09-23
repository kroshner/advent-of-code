package coordinate

type Coordinate struct {
	X int
	Y int
}

func (coordinate *Coordinate) ToPosition() int {
	return coordinate.X + coordinate.Y*10
}

type VentLine struct {
	Start Coordinate
	End   Coordinate
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (ventline *VentLine) ToCoordinates() ([]Coordinate, bool) {
	if ventline.Start.X == ventline.End.X {
		coordinates := make([]Coordinate, 0)
		for i := min(ventline.Start.Y, ventline.End.Y); i <= max(ventline.Start.Y, ventline.End.Y); i++ {
			coordinates = append(coordinates, Coordinate{X: ventline.Start.X, Y: i})
		}
		return coordinates, true
	} else if ventline.Start.Y == ventline.End.Y {
		coordinates := make([]Coordinate, 0)
		for i := min(ventline.Start.X, ventline.End.X); i <= max(ventline.Start.X, ventline.End.X); i++ {
			coordinates = append(coordinates, Coordinate{X: i, Y: ventline.Start.Y})
		}
		return coordinates, true
	} else {
		coordinates := make([]Coordinate, 0)
		minX := min(ventline.Start.X, ventline.End.X)
		minY := min(ventline.Start.Y, ventline.End.Y)
		xDiffAbs := abs(ventline.Start.X - ventline.End.X)
		for i := 0; i <= xDiffAbs; i++ {
			coordinates = append(coordinates, Coordinate{X: minX + i, Y: minY + i})
		}
		return coordinates, true
	}
}
