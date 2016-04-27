package room

import (
	"testing"
)

func TestTileTypeToString(t *testing.T) {
	if !(Moveable.String() == "Movable") {
		t.Error("Expected Moveable for String function")
	}

	if !(NotMoveable.String() == "NotMoveable") {
		t.Error("Expected NotMoveable for String function")
	}
}
