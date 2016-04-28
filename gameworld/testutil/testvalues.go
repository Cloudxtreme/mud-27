package testutil

import "github.com/Norskan/mud/gameworld/room"

var TestTileID = 0

var TestTileFloor = room.NewTileTemplate(".", "Floor", room.Moveable)
var TestTileWall = room.NewTileTemplate("w", "Wall", room.NotMoveable)

var TestGameWorldTemplateVertical = [][]string{{"."}, {"."}}
var TestGameWorldTemplateHorizontal = [][]string{{".", "."}}
