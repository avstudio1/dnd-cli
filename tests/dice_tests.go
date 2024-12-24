package dice_test

import (
	"testing"

	"github.com/avstudio1/dnd-cli/dice"
)

func TestRoll(t *testing.T) {
	tests := []struct {
		input       string
		percentage  bool
		average     bool
		expectError bool
	}{
		{"2d6", false, false, false},
		{"1d20", false, true, false},
		{"2d10", true, false, false},
		{"5d8", false, true, false},
		{"invalid", false, false, true},
	}

	for _, test := range tests {
		_, err := dice.Roll(test.input, test.percentage, test.average)
		if (err != nil) != test.expectError {
			t.Errorf("Unexpected error state for input %s: %v", test.input, err)
		}
	}
}
