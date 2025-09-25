package input

import (
	"os/exec"
)

// sendEnter sends an Enter key (using newline type for better compatibility)
func SendEnter(windowID string) error {
	return exec.Command("xdotool", "key", "--window", windowID, "KP_Enter").Run()
}

// sendText types a message into the window
func SendText(windowID, text string) error {
	return exec.Command("xdotool", "type", "--window", windowID, text).Run()
}
