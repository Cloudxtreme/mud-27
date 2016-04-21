package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

var testGameWorldTemplateHorizontal = [][]string{{".", "."}}

func TestNewGameWorldHorizontal(t *testing.T) {
	tileTemplates := make(map[string]*gameworld.TileTemplate)
	tileTemplates["."] = gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable)
	tileTemplates["w"] = gameworld.NewTileTemplate(".", "Wall", gameworld.NotMoveable)

	gw := gameworld.NewGameWorld(testGameWorldTemplateHorizontal, tileTemplates)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == tileTemplates["."].Mark()) ||
		!(gw.Room().Area()[0][1].Mark() == tileTemplates["."].Mark()) {
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
	tileTemplates := make(map[string]*gameworld.TileTemplate)
	tileTemplates["."] = gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable)
	tileTemplates["w"] = gameworld.NewTileTemplate("w", "Wall", gameworld.NotMoveable)

	gw := gameworld.NewGameWorld(testGameWorldTemplateVertical, tileTemplates)

	//test for tilemarks
	if !(gw.Room().Area()[0][0].Mark() == tileTemplates["."].Mark()) ||
		!(gw.Room().Area()[1][0].Mark() == tileTemplates["."].Mark()) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.Room().Area()[0][0].GetConnetionTile(gameworld.Down) == gw.Room().Area()[1][0]) ||
		!(gw.Room().Area()[1][0].GetConnetionTile(gameworld.Up) == gw.Room().Area()[0][0]) {
		t.Error("Tiles were not right connected verticaly")
	}
}
