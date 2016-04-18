package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

//parameter for test values
var actualID = 0
var actualMark = "test"
var actualDescription = "testDescription"

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := gameworld.NewDefaultTile(actualID, actualMark)

	if !(newDefaultTile.ID() == actualID) {
		t.Errorf("Expected %v for ID", actualID)
	}

	if !(newDefaultTile.Mark() == actualMark) {
		t.Errorf("Expected %v for Mark", actualMark)
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := gameworld.NewTile(actualID, actualMark, actualDescription)

	if !(newTile.ID() == actualID) {
		t.Errorf("Expected %v for ID", actualID)
	}

	if !(newTile.Mark() == actualMark) {
		t.Errorf("Expected %v for Mark", actualMark)
	}

	if !(newTile.Description() == actualDescription) {
		t.Errorf("Expected %v for Description", actualDescription)
	}
}

func TestTileConnection(t *testing.T) {
	tileUp := gameworld.NewTile(actualID, actualMark, actualDescription)
	tileRight := gameworld.NewTile(actualID, actualMark, actualDescription)
	tileDown := gameworld.NewTile(actualID, actualMark, actualDescription)
	tileLeft := gameworld.NewTile(actualID, actualMark, actualDescription)

	tileMiddle := gameworld.NewTile(actualID, actualMark, actualDescription)

	tileMiddle.SetConnetionTile(tileUp, gameworld.Up)
	tileMiddle.SetConnetionTile(tileRight, gameworld.Right)
	tileMiddle.SetConnetionTile(tileDown, gameworld.Down)
	tileMiddle.SetConnetionTile(tileLeft, gameworld.Left)

	if !(tileMiddle.GetConnetionTile(gameworld.Up) == tileUp) ||
		!(tileUp.GetConnetionTile(gameworld.Down) == tileMiddle) {
		t.Error("Invalide tile set for tileUp")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Right) == tileRight) ||
		!(tileRight.GetConnetionTile(gameworld.Left) == tileMiddle) {
		t.Error("Invalide tile set for tileRight")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Down) == tileDown) ||
		!(tileDown.GetConnetionTile(gameworld.Up) == tileMiddle) {
		t.Error("Invalide tile set for tileDown")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Left) == tileLeft) ||
		!(tileLeft.GetConnetionTile(gameworld.Right) == tileMiddle) {
		t.Error("Invalide tile set for tileLeft")
	}
}

func TestTileSetCharacter(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTile := gameworld.NewTile(actualID, actualMark, actualDescription)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.GetCharacter() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacter(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTileA := gameworld.NewTile(actualID, actualMark, actualDescription)
	testTileB := gameworld.NewTile(actualID, actualMark, actualDescription)

	testTileA.SetConnetionTile(testTileB, gameworld.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	testTileB.MoveCharacter(testTileA.GetCharacter())

	if !(testTileA.GetCharacter() == nil) &&
		(testTileB.GetCharacter() == actualCharacter) {
		t.Error("Character did not move correctly from testTileA to testTileB")
	}
}
