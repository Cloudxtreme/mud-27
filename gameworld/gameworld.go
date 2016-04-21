package gameworld

//GameWorld hold information about the current world
type GameWorld struct {
	room *Room
}

//NewGameWorld creates a new gameWorld instance
func NewGameWorld(gameWorldTemplate [][]string, tileTemplates map[string]TileTemplate) *GameWorld {
	gw := GameWorld{}

	gw.room = NewRoom(gameWorldTemplate, tileTemplates map[string]TileTemplate)

	return &gw
}

//Room gets the room
func (gw *GameWorld) Room() *Room {
	return gw.room
}

//SetCharacter sets a character on a tile
func (gw *GameWorld) SetCharacter(character *Character, gameTile *GameTile) {
	gameTile.SetCharacter(character)
	character.tilePosition = gameTile
}
