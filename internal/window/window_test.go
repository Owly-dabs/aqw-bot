package window

import (
	"testing"
)

func TestNewWindow(t *testing.T) {
	win, err := NewWindow()
	if err != nil {
		t.Fatalf("Failed to create new window: %v", err)
	}

	t.Logf("Window ID: %s", win.ID)
	t.Logf("Position: (%d, %d)", win.X, win.Y)
	t.Logf("Size: %dx%d", win.Width, win.Height)
}
