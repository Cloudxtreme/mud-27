package room

import (
	"testing"
)

func TestDirectionToString(t *testing.T) {
	if !(Up.String() == "Up") {
		t.Error("Expected Up for String function")
	}

	if !(Right.String() == "Right") {
		t.Error("Expected Right for String function")
	}

	if !(Down.String() == "Down") {
		t.Error("Expected Down for String function")
	}

	if !(Left.String() == "Left") {
		t.Error("Expected Left for String function")
	}
}
