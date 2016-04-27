package character

import (
	"testing"

	"github.com/Norskan/mud/gameworld/room"
)

var testTileID = 0

var testTileFloor = room.NewTileTemplate(".", "Floor", room.Moveable)
var testTileWall = room.NewTileTemplate("w", "Wall", room.NotMoveable)

func TestCharacterNewCharacter(t *testing.T) {
	newCharacter := NewCharacter()

	if !(newCharacter != nil) {
		t.Error("Expected a new Character")
	}
}

func TestGetTilePosition(t *testing.T) {
	newCharacter := NewCharacter()
	gameTile := room.NewTile(testTileID, testTileWall)

	newCharacter.SetTilePosition(gameTile)

	if !(newCharacter.GetTilePositon() == gameTile) {
		t.Error("Expected a tile on tilePosition")
	}
}

func TestMoveCharacter(t *testing.T) {
	//TODO fix test
	/*character := NewCharacter()

	roomTemplate := room.NewRoomTemplate(testGameWorldTemplateHorizontal)
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)
	gw.SetCharacter(character, gw.Room().Area()[0][0])

	Move(room.Right)

	if !(character.GetTilePositon() == gw.Room().Area()[0][1]) {
		t.Error("Character was not moved")
	}

	if !(gw.Room().Area()[0][0].Character() == nil) {
		t.Error("Character is still on original position")
	}

	if !(gw.Room().Area()[0][1].Character() == character) {
		t.Error("Character is not on new position")
	}*/

}
