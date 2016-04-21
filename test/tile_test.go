package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

//parameter for tile test values
var testTileID = 0

var testTileFloor = gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable)
var testTileWall = gameworld.NewTileTemplate("w", "Wall", gameworld.NotMoveable)

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := gameworld.NewTile(testTileID, testTileFloor)

	if !(newDefaultTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newDefaultTile.Mark() == testTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testTileFloor.Mark())
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := gameworld.NewTile(testTileID, testTileFloor)

	if !(newTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newTile.Mark() == testTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testTileFloor.Mark())
	}

	if !(newTile.Description() == testTileFloor.Description()) {
		t.Errorf("Expected %v for Description", testTileFloor.Description())
	}
}

func TestTileConnection(t *testing.T) {
	tileUp := gameworld.NewTile(testTileID, testTileFloor)
	tileRight := gameworld.NewTile(testTileID, testTileFloor)
	tileDown := gameworld.NewTile(testTileID, testTileFloor)
	tileLeft := gameworld.NewTile(testTileID, testTileFloor)

	tileMiddle := gameworld.NewTile(testTileID, testTileFloor)

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
	testTile := gameworld.NewTile(testTileID, testTileFloor)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.Character() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacterMovalbe(t *testing.T) {
	actualCharacter := gameworld.NewCharacter()
	testTileA := gameworld.NewTile(testTileID, testTileFloor)
	testTileB := gameworld.NewTile(testTileID, testTileFloor)

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
	testTileA := gameworld.NewTile(testTileID, testTileFloor)
	testTileB := gameworld.NewTile(testTileID, testTileWall)

	testTileA.SetConnetionTile(testTileB, gameworld.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	if testTileB.MoveCharacter(actualCharacter) ||
		!(testTileA.Character() == actualCharacter) ||
		!(testTileB.Character() == nil) {
		t.Error("Character should not be able to move")
	}
}
