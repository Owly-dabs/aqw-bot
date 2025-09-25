package window

import (
	"fmt"
	"os/exec"
	"strings"
)

// getWindowID searches for a window by title (partial match allowed)
// LEGACY
func GetWindowID() (string, error) {
	cmd := exec.Command("xdotool", "search", "--name", "Artix Game Launcher")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not find window: %v", err)
	}
	// Return first window match
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines[0], nil
}

// activateWindow brings the window to foreground
func (w *Window) Activate() error {
	return exec.Command("xdotool", "windowactivate", "--sync", w.ID).Run()
}

// LEGACY
func GetWindowGeometry(windowID string) (int, int, int, int, error) {
	cmd := exec.Command("xdotool", "getwindowgeometry", windowID)
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("could not get window geometry: %v", err)
	}

	var x, y, w, h int
	_, err = fmt.Sscanf(string(out), "Window %s:\n  Position: %d,%d\n  Geometry: %dx%d", &x, &y, &w, &h)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("could not parse window geometry: %v", err)
	}

	return x, y, w, h, nil
}
