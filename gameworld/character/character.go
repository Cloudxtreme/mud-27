package character

import (
	"github.com/Norskan/mud/gameworld/room"
)

//Character type of the character
type Character struct {
	tilePosition *room.GameTile
}

//NewCharacter create a new default character
func NewCharacter() *Character {
	return &Character{}
}

//GetTilePositon returns the current position of the character
func (character *Character) GetTilePositon() *room.GameTile {
	return character.tilePosition
}

//SetTilePosition sets the current position of the character
func (character *Character) SetTilePosition(gameTile *room.GameTile) {
	character.tilePosition = gameTile
}

//Move moves the character from the current gametile into direction
func (character *Character) Move(direction room.Direction) {
	currentGameTile := character.GetTilePositon()
	currentGameTile.GetConnetionTile(direction).MoveCharacter(character)
}
