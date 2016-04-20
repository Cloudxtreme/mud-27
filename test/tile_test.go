package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

//parameter for tile test values
var testTileID = 0
var testTileMark = "testTileMark"
var testTileDescription = "testDescription"

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := gameworld.NewDefaultTile(testTileID, testTileMark)

	if !(newDefaultTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newDefaultTile.Mark() == testTileMark) {
		t.Errorf("Expected %v for Mark", testTileMark)
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)

	if !(newTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newTile.Mark() == testTileMark) {
		t.Errorf("Expected %v for Mark", testTileMark)
	}

	if !(newTile.Description() == testTileDescription) {
		t.Errorf("Expected %v for Description", testTileDescription)
	}
}

func TestTileConnection(t *testing.T) {
	tileUp := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)
	tileRight := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)
	tileDown := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)
	tileLeft := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)

	tileMiddle := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)

	tileMiddle.SetConnetionTile(tileUp, gameworld.Up)
	tileMiddle.SetConnetionTile(tileRight, gameworld.Right)
	tileMiddle.SetConnetionTile(tileDown, gameworld.Down)
	tileMiddle.SetConnetionTile(tileLeft, gameworld.Left)

	if !(tileMiddle.GetConnetionTile(gameworld.Up) == tileUp) {
		t.Error("Invalide tile set for tileUp")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Right) == tileRight) {
		t.Error("Invalide tile set for tileRight")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Down) == tileDown) {
		t.Error("Invalide tile set for tileDown")
	}

	if !(tileMiddle.GetConnetionTile(gameworld.Left) == tileLeft) {
		t.Error("Invalide tile set for tileLeft")
	}
}

func TestTileSetCharacter(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTile := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.Character() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacterMovalbe(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTileA := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)
	testTileB := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)

	testTileA.SetConnetionTile(testTileB, gameworld.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	testTileB.MoveCharacter(testTileA.Character())

	if !(testTileA.Character() == nil) ||
		!(testTileB.Character() == actualCharacter) {
		t.Error("Character did not move correctly from testTileA to testTileB")
	}
}

func TestTileMoveCharacterNotMovable(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTileA := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.Moveable)
	testTileB := gameworld.NewTile(testTileID, testTileMark, testTileDescription, gameworld.NotMoveable)

	testTileA.SetConnetionTile(testTileB, gameworld.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	if testTileB.MoveCharacter(actualCharacter) ||
		!(testTileA.Character() == actualCharacter) ||
		!(testTileB.Character() == nil) {
		t.Error("Character should not be able to move")
	}
}
