package gameworld

//Direction enum for direction
type Direction int

//Direction direction of the tile connection
const (
	//Up direction
	Up Direction = iota
	//Right direction
	Right
	//Down direction
	Down
	//Left direction
	Left
)

var stringDirection = [...]string{"Up", "Right", "Down", "Left"}

func (direction Direction) String() string {
	return stringDirection[direction]
}
