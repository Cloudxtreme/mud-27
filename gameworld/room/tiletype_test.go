package room_test

import (
	"testing"

	"github.com/Norskan/mud/gameworld/room"
)

func TestTileTypeToString(t *testing.T) {
	if !(room.Moveable.String() == "Movable") {
		t.Error("Expected Moveable for String function")
	}

	if !(room.NotMoveable.String() == "NotMoveable") {
		t.Error("Expected NotMoveable for String function")
	}
}
