package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

func TestTileTypeToString(t *testing.T) {
	if !(gameworld.Moveable.String() == "Movable") {
		t.Error("Expected Moveable for String function")
	}

	if !(gameworld.NotMoveable.String() == "NotMoveable") {
		t.Error("Expected NotMoveable for String function")
	}
}
