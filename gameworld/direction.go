package gameworld

//Direction enum for direction
type Direction int

//Direction direction of the tile connection
const (
	//Up up direction
	Up Direction = iota
	//Right right direction
	Right
	//Down down direction
	Down
	//Left left direction
	Left
)

var stringDirection = [...]string{"Up", "Right", "Down", "Left"}

func (direction Direction) String() string {
	return stringDirection[direction]
}
