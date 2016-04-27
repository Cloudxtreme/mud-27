package gameworld

import (
	"github.com/Norskan/mud/gameworld/character"
	"github.com/Norskan/mud/gameworld/room"
)

//GameWorld hold information about the current world
type GameWorld struct {
	room *room.Room
}

//NewGameWorld creates a new gameWorld instance
func NewGameWorld(roomTemplate *room.Template) *GameWorld {
	gw := GameWorld{}

	gw.room = room.NewRoom(roomTemplate)

	return &gw
}

//Room gets the room
func (gw *GameWorld) Room() *room.Room {
	return gw.room
}

//SetCharacter sets a character on a tile
func (gw *GameWorld) SetCharacter(character *character.Character, gameTile *room.GameTile) {
	gameTile.SetCharacter(character)
	character.SetTilePosition(gameTile)
}
