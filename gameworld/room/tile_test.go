package room_test

import (
	"testing"

	"github.com/Norskan/mud/gameworld/room"
	"github.com/Norskan/mud/gameworld/testutil"
)

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

	if !(newDefaultTile.ID() == testutil.TestTileID) {
		t.Errorf("Expected %v for ID", testutil.TestTileID)
	}

	if !(newDefaultTile.Mark() == testutil.TestTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testutil.TestTileFloor.Mark())
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

	if !(newTile.ID() == testutil.TestTileID) {
		t.Errorf("Expected %v for ID", testutil.TestTileID)
	}

	if !(newTile.Mark() == testutil.TestTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testutil.TestTileFloor.Mark())
	}

	if !(newTile.Description() == testutil.TestTileFloor.Description()) {
		t.Errorf("Expected %v for Description", testutil.TestTileFloor.Description())
	}
}

func TestTileConnection(t *testing.T) {
	tileUp := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)
	tileRight := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)
	tileDown := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)
	tileLeft := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

	tileMiddle := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

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
	actualCharacter := testutil.NewCharacterMokup()
	testTile := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.Character() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacterMovalbe(t *testing.T) {
	actualCharacter := testutil.NewCharacterMokup()
	testTileA := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)
	testTileB := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)

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
	actualCharacter := testutil.NewCharacterMokup()
	testTileA := room.NewTile(testutil.TestTileID, testutil.TestTileFloor)
	testTileB := room.NewTile(testutil.TestTileID, testutil.TestTileWall)

	testTileA.SetConnetionTile(testTileB, room.Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	if testTileB.MoveCharacter(actualCharacter) ||
		!(testTileA.Character() == actualCharacter) ||
		!(testTileB.Character() == nil) {
		t.Error("Character should not be able to move")
	}
}
