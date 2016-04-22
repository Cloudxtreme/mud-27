package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

var testGameWorldTemplateHorizontal = [][]string{{".", "."}}

func TestNewGameWorldHorizontal(t *testing.T) {
	roomTemplate := gameworld.NewRoomTemplate(testGameWorldTemplateHorizontal)
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) ||
		!(gw.Room().Area()[0][1].Mark() == roomTemplate.TileTemplates()["."].Mark()) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.Room().Area()[0][0].GetConnetionTile(gameworld.Right) == gw.Room().Area()[0][1]) ||
		!(gw.Room().Area()[0][1].GetConnetionTile(gameworld.Left) == gw.Room().Area()[0][0]) {
		t.Error("Tiles were not right connected horizontaly")
	}
}

var testGameWorldTemplateVertical = [][]string{{"."}, {"."}}

func TestNewGameWorldVertical(t *testing.T) {

	roomTemplate := gameworld.NewRoomTemplate(testGameWorldTemplateVertical)
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))

	gw := gameworld.NewGameWorld(roomTemplate)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) ||
		!(gw.Room().Area()[1][0].Mark() == roomTemplate.TileTemplates()["."].Mark()) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.Room().Area()[0][0].GetConnetionTile(gameworld.Down) == gw.Room().Area()[1][0]) ||
		!(gw.Room().Area()[1][0].GetConnetionTile(gameworld.Up) == gw.Room().Area()[0][0]) {
		t.Error("Tiles were not right connected verticaly")
	}
}
