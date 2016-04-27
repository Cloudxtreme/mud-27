package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld/room"
)

func TestDirectionToString(t *testing.T) {
	if !(room.Up.String() == "Up") {
		t.Error("Expected Up for String function")
	}

	if !(room.Right.String() == "Right") {
		t.Error("Expected Right for String function")
	}

	if !(room.Down.String() == "Down") {
		t.Error("Expected Down for String function")
	}

	if !(room.Left.String() == "Left") {
		t.Error("Expected Left for String function")
	}
}
