package logic

import (
	"aqw/internal/window"
	"testing"
)

func TestFarm(t *testing.T) {
	win, err := window.NewWindow()
	if err != nil {
		t.Fatalf("Failed to create new window: %v", err)
	}

	err = FarmWarrior(win)
	if err != nil {
		t.Fatalf("Farm failed: %v", err)
	}

	t.Logf("Farming completed successfully on window ID: %s", win.ID)
}
