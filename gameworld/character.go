package gameworld

//Character type of the character
type Character struct {
	tilePosition *GameTile
}

//NewCharacter create a new default character
func NewCharacter() *Character {
	return &Character{}
}
