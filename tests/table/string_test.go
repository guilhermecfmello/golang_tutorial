package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - Indexes: Expected (%d) <> Found (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Havia uma pedra no meio do caminho", "pedra", 10},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Guilherme", "u", 1},
	}
	for _, test := range tests {
		t.Logf("Massa: %v", test)
		current := strings.Index(test.text, test.part)
		if current != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}
