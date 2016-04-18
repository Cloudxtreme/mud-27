package test

import (
	"testing"

	"github.com/Norskan/mud/gameworld"
)

func TestDirectionToStringTest(t *testing.T) {
	if !(gameworld.Up.String() == "Up") {
		t.Error("Expected Up for String function")
	}

	if !(gameworld.Right.String() == "Right") {
		t.Error("Expected Right for String function")
	}

	if !(gameworld.Down.String() == "Down") {
		t.Error("Expected Down for String function")
	}

	if !(gameworld.Left.String() == "Left") {
		t.Error("Expected Left for String function")
	}
}
