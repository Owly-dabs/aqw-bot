package window

import (
	"testing"
)

func TestGetWindowID(t *testing.T) {
	id, err := GetWindowID()
	if err != nil {
		t.Fatalf("Failed to get window ID: %v", err)
	}
	t.Logf("Window ID: %s", id)
}
