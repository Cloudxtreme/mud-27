package gameworld

import (
	"fmt"
)

//GameTile ....
type GameTile struct {
	//tile description
	id          int
	mark        string
	description string

	//connetions to other GameTiles
	tileUp    *GameTile
	tileRight *GameTile
	tileDown  *GameTile
	tileLeft  *GameTile

	//Character on the tile
	// TODO: refactor to multible characters
	character *Character
}

//NewDefaultTile create a new default tile
func NewDefaultTile(id int, mark string) *GameTile {
	return NewTile(id, mark, "Default")
}

//NewTile create new generic tile
func NewTile(id int, mark string, description string) *GameTile {
	return &GameTile{id: id, mark: mark, description: description}
}

//PrintDebug returns a string that represents the state of the Tile
func (tile *GameTile) PrintDebug() {
	fmt.Printf("[Tile id:%v, name:%v, con[%v, %v, %v, %v]]", tile.ID(), tile.description, tile.tileUp.ID(), tile.tileRight.ID(), tile.tileDown.ID(), tile.tileLeft.ID())
}

//ID getter for id property
func (tile *GameTile) ID() int {
	if tile == nil {
		return -1
	}
	return tile.id
}

//Mark getter for mark property
func (tile *GameTile) Mark() string {
	if tile.character != nil {
		return "c"
	}
	return tile.mark
}

//Description getter for the description property
func (tile *GameTile) Description() string {
	return tile.description
}

//GetConnetionTile sets the connection in the direction
func (tile *GameTile) GetConnetionTile(direction Direction) *GameTile {
	switch direction {
	case Up:
		return tile.tileUp
	case Right:
		return tile.tileRight
	case Down:
		return tile.tileDown
	case Left:
		return tile.tileLeft
	default:
		return nil
	}
}

//SetConnetionTile sets the connection in the direction
func (tile *GameTile) SetConnetionTile(connectedTile *GameTile, direction Direction) {
	switch direction {
	case Up:
		tile.tileUp = connectedTile
	case Right:
		tile.tileRight = connectedTile
	case Down:
		tile.tileDown = connectedTile
	case Left:
		tile.tileLeft = connectedTile
	}
}

//SetCharacter set a character on the tile
func (tile *GameTile) SetCharacter(character *Character) {
	tile.character = character
}

//MoveCharacter moves a character on the tile
func (tile *GameTile) MoveCharacter(character *Character) bool {
	// TODO: add logic for testing if the character can move onto the tile
	character.tilePosition.character = nil
	tile.character = character
	character.tilePosition = tile
	return true
}
