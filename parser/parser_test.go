package parser

import (
	"testing"
)

func TestHigh(t *testing.T) {
	t.Run("highest_score", func(t *testing.T) {
		name := High()
		if name != "Prisha" {
			t.Error("test failed: ", name, "Prisha")
		}
	})
}

func TestLow(t *testing.T) {
	t.Run("highest_score", func(t *testing.T) {
		name := Low()
		if name != "Charlie" {
			t.Error("test failed", name, "Charlie")
		}
	})
}
