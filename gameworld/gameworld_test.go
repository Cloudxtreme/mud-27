package gameworld_test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
	"github.com/Norskan/mud/gameworld/room"
	"github.com/Norskan/mud/gameworld/testutil"
)

func TestNewGameWorldHorizontal(t *testing.T) {
	roomTemplate := room.NewRoomTemplate(testutil.TestGameWorldTemplateHorizontal)
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) ||
		!(gw.Room().Area()[0][1].Mark() == roomTemplate.TileTemplates()["."].Mark()) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.Room().Area()[0][0].GetConnetionTile(room.Right) == gw.Room().Area()[0][1]) ||
		!(gw.Room().Area()[0][1].GetConnetionTile(room.Left) == gw.Room().Area()[0][0]) {
		t.Error("Tiles were not right connected horizontaly")
	}
}

func TestNewGameWorldVertical(t *testing.T) {

	roomTemplate := room.NewRoomTemplate(testutil.TestGameWorldTemplateVertical)
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) ||
		!(gw.Room().Area()[1][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.Room().Area()[0][0].GetConnetionTile(room.Down) == gw.Room().Area()[1][0]) ||
		!(gw.Room().Area()[1][0].GetConnetionTile(room.Up) == gw.Room().Area()[0][0]) {
		t.Error("Tiles were not right connected verticaly")
	}
}
