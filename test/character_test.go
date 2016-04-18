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
