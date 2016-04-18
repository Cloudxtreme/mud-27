package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

var testGameWorldTemplateHorizontal21 = [][]string{{testTileMark, testTileMark}}

func TestNewGameWorldHorizontal(t *testing.T) {
	gw := gameworld.NewGameWorld(testGameWorldTemplateHorizontal21)

	//test for tilemarks
	if !(gw.GameArea()[0][0].Mark() == testTileMark) ||
		!(gw.GameArea()[0][1].Mark() == testTileMark) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.GameArea()[0][0].GetConnetionTile(gameworld.Right) == gw.GameArea()[0][1]) ||
		!(gw.GameArea()[0][1].GetConnetionTile(gameworld.Left) == gw.GameArea()[0][0]) {
		t.Error("Tiles were not right connected horizontaly")
	}
}

var testGameWorldTemplateHorizontal12 = [][]string{{testTileMark}, {testTileMark}}

func TestNewGameWorldVertical(t *testing.T) {
	gw := gameworld.NewGameWorld(testGameWorldTemplateHorizontal12)

	//test for tilemarks
	if !(gw.GameArea()[0][0].Mark() == testTileMark) ||
		!(gw.GameArea()[1][0].Mark() == testTileMark) {
		t.Error("Invalide tile marks")
	}

	//test for connection
	if !(gw.GameArea()[0][0].GetConnetionTile(gameworld.Down) == gw.GameArea()[1][0]) ||
		!(gw.GameArea()[1][0].GetConnetionTile(gameworld.Up) == gw.GameArea()[0][0]) {
		t.Error("Tiles were not right connected verticaly")
	}
}
