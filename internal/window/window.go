package window

import (
	"fmt"
	"os/exec"
	"strings"
)

type Window struct {
	ID     string
	X      int
	Y      int
	Width  int
	Height int
}

// NewWindowFromTitle searches for a window by title and builds a Window struct
func NewWindow() (*Window, error) {
	// Find window ID
	out, err := exec.Command("xdotool", "search", "--name", "Artix Game Launcher").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to find window: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no window found for title: %s", "Artix Game Launcher")
	}
	id := lines[0]

	// Get window geometry
	cmd := exec.Command("xdotool", "getwindowgeometry", id)
	geomOut, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get geometry: %v", err)
	}

	var x, y, w, h int
	lines = strings.Split(string(geomOut), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Position:") {
			fmt.Sscanf(line, "Position: %d,%d", &x, &y)
		}
		if strings.HasPrefix(line, "Geometry:") {
			fmt.Sscanf(line, "Geometry: %dx%d", &w, &h)
		}
	}

	return &Window{
		ID:     id,
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}, nil
}
