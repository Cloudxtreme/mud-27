package room

import (
	"testing"
)

//parameter for tile test values
var testTileID = 0

var testTileFloor = NewTileTemplate(".", "Floor", Moveable)
var testTileWall = NewTileTemplate("w", "Wall", NotMoveable)

type CharacterMokup struct {
	tilePosition *GameTile
}

func NewCharacter() *CharacterMokup {
	return &CharacterMokup{}
}

func (character *CharacterMokup) GetTilePositon() *GameTile {
	return character.tilePosition
}

func (character *CharacterMokup) SetTilePosition(gameTile *GameTile) {
	character.tilePosition = gameTile
}

func (character *CharacterMokup) Move(direction Direction) {
	currentGameTile := character.GetTilePositon()
	currentGameTile.GetConnetionTile(direction).MoveCharacter(character)
}

func TestTileNewDefaultTile(t *testing.T) {
	newDefaultTile := NewTile(testTileID, testTileFloor)

	if !(newDefaultTile.ID() == testTileID) {
		t.Errorf("Expected %v for ID", testTileID)
	}

	if !(newDefaultTile.Mark() == testTileFloor.Mark()) {
		t.Errorf("Expected %v for Mark", testTileFloor.Mark())
	}
}

func TestTileNewTile(t *testing.T) {
	newTile := NewTile(testTileID, testTileFloor)

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
	tileUp := NewTile(testTileID, testTileFloor)
	tileRight := NewTile(testTileID, testTileFloor)
	tileDown := NewTile(testTileID, testTileFloor)
	tileLeft := NewTile(testTileID, testTileFloor)

	tileMiddle := NewTile(testTileID, testTileFloor)

	tileMiddle.SetConnetionTile(tileUp, Up)
	tileMiddle.SetConnetionTile(tileRight, Right)
	tileMiddle.SetConnetionTile(tileDown, Down)
	tileMiddle.SetConnetionTile(tileLeft, Left)

	if !(tileMiddle.GetConnetionTile(Up) == tileUp) {
		t.Error("Invalide tile set for tileUp")
	}

	if !(tileMiddle.GetConnetionTile(Right) == tileRight) {
		t.Error("Invalide tile set for tileRight")
	}

	if !(tileMiddle.GetConnetionTile(Down) == tileDown) {
		t.Error("Invalide tile set for tileDown")
	}

	if !(tileMiddle.GetConnetionTile(Left) == tileLeft) {
		t.Error("Invalide tile set for tileLeft")
	}
}

func TestTileSetCharacter(t *testing.T) {
	actualCharacter := NewCharacter()
	testTile := NewTile(testTileID, testTileFloor)

	testTile.SetCharacter(actualCharacter)

	if !(testTile.Character() == actualCharacter) {
		t.Error("Invalide character set in testTile")
	}
}

func TestTileMoveCharacterMovalbe(t *testing.T) {
	actualCharacter := NewCharacter()
	testTileA := NewTile(testTileID, testTileFloor)
	testTileB := NewTile(testTileID, testTileFloor)

	testTileA.SetConnetionTile(testTileB, Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	testTileB.MoveCharacter(testTileA.Character())

	if !(testTileA.Character() == nil) ||
		!(testTileB.Character() == actualCharacter) {
		t.Error("Character did not move correctly from testTileA to testTileB")
	}
}

func TestTileMoveCharacterNotMovable(t *testing.T) {
	actualCharacter := NewCharacter()
	testTileA := NewTile(testTileID, testTileFloor)
	testTileB := NewTile(testTileID, testTileWall)

	testTileA.SetConnetionTile(testTileB, Left)
	testTileA.SetCharacter(actualCharacter)
	actualCharacter.SetTilePosition(testTileA)

	if testTileB.MoveCharacter(actualCharacter) ||
		!(testTileA.Character() == actualCharacter) ||
		!(testTileB.Character() == nil) {
		t.Error("Character should not be able to move")
	}
}
