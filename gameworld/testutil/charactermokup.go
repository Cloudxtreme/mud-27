package testutil

import "github.com/Norskan/mud/gameworld/room"

type CharacterMokup struct {
	tilePosition *room.GameTile
}

func NewCharacterMokup() *CharacterMokup {
	return &CharacterMokup{}
}

func (character *CharacterMokup) GetTilePositon() *room.GameTile {
	return character.tilePosition
}

func (character *CharacterMokup) SetTilePosition(gameTile *room.GameTile) {
	character.tilePosition = gameTile
}

func (character *CharacterMokup) Move(direction room.Direction) {
	currentGameTile := character.GetTilePositon()
	currentGameTile.GetConnetionTile(direction).MoveCharacter(character)
}
