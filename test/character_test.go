package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

func TestCharacterNewCharacter(t *testing.T) {
	newCharacter := gameworld.NewCharacter()

	if !(newCharacter != nil) {
		t.Error("Expected a new Character")
	}
}

func TestGetTilePosition(t *testing.T) {
	newCharacter := gameworld.NewCharacter()
	gameTile := gameworld.NewTile(testTileID, testTileWall)

	newCharacter.SetTilePosition(gameTile)

	if !(newCharacter.GetTilePositon() == gameTile) {
		t.Error("Expected a tile on tilePosition")
	}
}

func TestMoveCharacter(t *testing.T) {
	character := gameworld.NewCharacter()

	roomTemplate := gameworld.NewRoomTemplate(testGameWorldTemplateHorizontal)
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)
	gw.SetCharacter(character, gw.Room().Area()[0][0])

	character.Move(gameworld.Right)

	if !(character.GetTilePositon() == gw.Room().Area()[0][1]) {
		t.Error("Character was not moved")
	}

	if !(gw.Room().Area()[0][0].Character() == nil) {
		t.Error("Character is still on original position")
	}

	if !(gw.Room().Area()[0][1].Character() == character) {
		t.Error("Character is not on new position")
	}

}
