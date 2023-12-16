package lib

type Coordinate struct {
	Row int
	Col int
}

type Location struct {
	Coordinate Coordinate
	CameFrom   int
	Steps      int
	Pipe       string
	MaxWidth   int
	MaxLength  int
}

type MazeLocation struct {
	Coordinate Coordinate
	Direction  string
}


func CopyMazeLocation(origin *MazeLocation) *MazeLocation {
	return &MazeLocation{
		Coordinate : origin.Coordinate,
		Direction : origin.Direction,
	}
}