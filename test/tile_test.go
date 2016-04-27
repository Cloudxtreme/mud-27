package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld/character"
	"github.com/Norskan/mud/gameworld/room"
)

//parameter for tile test values
var testTileID = 0

var testTileFloor = room.NewTileTemplate(".", "Floor", room.Moveable)
var testTileWall = room.NewTileTemplate("w", "Wall", room.NotMoveable)

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := room.NewTile(testTileID, testTileFloor)

	if !(newDefaultTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newDefaultTile.Mark() == testTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testTileFloor.Mark())
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := room.NewTile(testTileID, testTileFloor)

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
	tileUp := room.NewTile(testTileID, testTileFloor)
	tileRight := room.NewTile(testTileID, testTileFloor)
	tileDown := room.NewTile(testTileID, testTileFloor)
	tileLeft := room.NewTile(testTileID, testTileFloor)

	tileMiddle := room.NewTile(testTileID, testTileFloor)

	tileMiddle.SetConnetionTile(tileUp, room.Up)
	tileMiddle.SetConnetionTile(tileRight, room.Right)
	tileMiddle.SetConnetionTile(tileDown, room.Down)
	tileMiddle.SetConnetionTile(tileLeft, room.Left)

	if !(tileMiddle.GetConnetionTile(room.Up) == tileUp) {
		t.Error("Invalide tile set for tileUp")
	}

	if !(tileMiddle.GetConnetionTile(room.Right) == tileRight) {
		t.Error("Invalide tile set for tileRight")
	}

	if !(tileMiddle.GetConnetionTile(room.Down) == tileDown) {
		t.Error("Invalide tile set for tileDown")
	}

	if !(tileMiddle.GetConnetionTile(room.Left) == tileLeft) {
		t.Error("Invalide tile set for tileLeft")
	}
}

func TestTileSetCharacter(t *testing.T) {
	actualCharacter := character.NewCharacter()
	testTile := room.NewTile(testTileID, testTileFloor)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.Character() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacterMovalbe(t *testing.T) {
	actualCharacter := character.NewCharacter()
	testTileA := room.NewTile(testTileID, testTileFloor)
	testTileB := room.NewTile(testTileID, testTileFloor)

	testTileA.SetConnetionTile(testTileB, room.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	testTileB.MoveCharacter(testTileA.Character())

	if !(testTileA.Character() == nil) ||
		!(testTileB.Character() == actualCharacter) {
		t.Error("Character did not move correctly from testTileA to testTileB")
	}
}

func TestTileMoveCharacterNotMovable(t *testing.T) {
	actualCharacter := character.NewCharacter()
	testTileA := room.NewTile(testTileID, testTileFloor)
	testTileB := room.NewTile(testTileID, testTileWall)

	testTileA.SetConnetionTile(testTileB, room.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	if testTileB.MoveCharacter(actualCharacter) ||
		!(testTileA.Character() == actualCharacter) ||
		!(testTileB.Character() == nil) {
		t.Error("Character should not be able to move")
	}
}
