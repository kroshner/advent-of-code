package coordinate

type Coordinate struct {
	X int
	Y int
}

func (coordinate *Coordinate) ToPosition() int {
	return coordinate.X + coordinate.Y*1000
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
		coordinates := ComplexCase(*ventline)
		return coordinates, true
	}
}

type VentLineDirectionChange struct {
	X, Y int
}

func (ventline *VentLine) GetDirection() VentLineDirectionChange {
	x := 0
	y := 0
	if ventline.End.X > ventline.Start.X {
		x = 1
	} else {
		x = -1
	}
	if ventline.End.Y > ventline.Start.Y {
		y = 1
	} else {
		y = -1
	}
	return VentLineDirectionChange{X: x, Y: y}
}

func ComplexCase(ventline VentLine) []Coordinate {
	xDiffAbs := abs(ventline.Start.X - ventline.End.X)
	ventLineDirectionChange := ventline.GetDirection()
	coordinates := make([]Coordinate, 0)
	for i := 0; i <= xDiffAbs; i++ {
		baseXCoord := 0
		xCoordDelta := 0
		if i == 0 {
			baseXCoord = ventline.Start.X
		} else {
			previousCoord := coordinates[i-1]
			baseXCoord = previousCoord.X
			xCoordDelta = ventLineDirectionChange.X
		}
		baseYCoord := 0
		yCoordDelta := 0
		if i == 0 {
			baseYCoord = ventline.Start.Y
		} else {
			previousCoord := coordinates[i-1]
			baseYCoord = previousCoord.Y
			yCoordDelta = ventLineDirectionChange.Y
		}
		coordinates = append(coordinates, Coordinate{X: baseXCoord + xCoordDelta, Y: baseYCoord + yCoordDelta})
	}
	return coordinates
}
