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

	tileType TileType
}

//NewTile create new generic tile
func NewTile(id int, tileTemplate *TileTemplate) *GameTile {
	return &GameTile{id: id, mark: tileTemplate.Mark(), description: tileTemplate.Description(), tileType: tileTemplate.TileType()}
}

func newTile(id int, mark string, description string, tileType TileType) *GameTile {
	return &GameTile{id: id, mark: mark, description: description, tileType: tileType}
}

//PrintDebug returns a string that represents the state of the Tile
func (gameTile *GameTile) PrintDebug() {
	fmt.Printf("[Tile id:%v, name:%v, con[%v, %v, %v, %v]]", gameTile.ID(), gameTile.description, gameTile.tileUp.ID(), gameTile.tileRight.ID(), gameTile.tileDown.ID(), gameTile.tileLeft.ID())
}

//ID getter for id property
func (gameTile *GameTile) ID() int {
	if gameTile == nil {
		return -1
	}
	return gameTile.id
}

//Mark getter for mark property
func (gameTile *GameTile) Mark() string {
	if gameTile.character != nil {
		return "c"
	}
	return gameTile.mark
}

//Description getter for the description property
func (gameTile *GameTile) Description() string {
	return gameTile.description
}

//GetConnetionTile sets the connection in the direction
func (gameTile *GameTile) GetConnetionTile(direction Direction) *GameTile {
	switch direction {
	case Up:
		return gameTile.tileUp
	case Right:
		return gameTile.tileRight
	case Down:
		return gameTile.tileDown
	case Left:
		return gameTile.tileLeft
	default:
		return nil
	}
}

//SetConnetionTile sets the connection in the direction of both tiles
func (gameTile *GameTile) SetConnetionTile(connectionTile *GameTile, direction Direction) {
	switch direction {
	case Up:
		gameTile.tileUp = connectionTile
	case Right:
		gameTile.tileRight = connectionTile
	case Down:
		gameTile.tileDown = connectionTile
	case Left:
		gameTile.tileLeft = connectionTile
	}
}

//Character returns the character set on the tile or nil if non is set
func (gameTile *GameTile) Character() *Character {
	return gameTile.character
}

//SetCharacter set a character on the tile
func (gameTile *GameTile) SetCharacter(character *Character) {
	gameTile.character = character
}

//TileType gets the type of the gameTile
func (gameTile *GameTile) TileType() TileType {
	return gameTile.tileType
}

//MoveCharacter moves a character on the tile
func (gameTile *GameTile) MoveCharacter(character *Character) bool {
	if gameTile.TileType() == Moveable {
		character.tilePosition.character = nil
		gameTile.character = character
		character.tilePosition = gameTile
		return true
	}
	return false
}
