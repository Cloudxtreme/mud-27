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
	gameTile := gameworld.NewDefaultTile(testTileID, testTileMark)

	newCharacter.SetTilePosition(gameTile)

	if !(newCharacter.GetTilePositon() == gameTile) {
		t.Error("Expected a tile on tilePosition")
	}
}

func TestMoveCharacter(t *testing.T) {
	character := gameworld.NewCharacter()
	gw := gameworld.NewGameWorld(testGameWorldTemplateHorizontal)
	gw.SetCharacter(character, gw.GameArea()[0][0])

	character.Move(gameworld.Right)

	if !(character.GetTilePositon() == gw.GameArea()[0][1]) {
		t.Error("Character was not moved")
	}

	if !(gw.GameArea()[0][0].Character() == nil) {
		t.Error("Character is still on original position")
	}

	if !(gw.GameArea()[0][1].Character() == character) {
		t.Error("Character is not on new position")
	}

}
