package ocr

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// allow mocking in tests
var getWorkingDir = os.Getwd

type OCRResult struct {
	Text  []string `json:"text"`
	Error string   `json:"error,omitempty"`
}

func RunEasyOCR(imagePath string) ([]string, error) {
	exePath, err := getWorkingDir()
	if err != nil {
		return nil, fmt.Errorf("could not determine working directory: %w", err)
	}

	scriptPath := filepath.Join(exePath, "scripts", "easyocr_runner.py")
	venvPython := filepath.Join(exePath, "scripts", ".venv", "bin", "python")

	cmd := exec.Command(venvPython, scriptPath, imagePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("python error: %v - output: %s", err, out.String())
	}

	// 🔍 Extract first line that looks like JSON
	scanner := bufio.NewScanner(&out)
	var jsonLine string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}") {
			jsonLine = line
			break
		}
	}

	if jsonLine == "" {
		return nil, fmt.Errorf("no valid JSON found in output:\n%s", out.String())
	}

	var result OCRResult
	if err := json.Unmarshal([]byte(jsonLine), &result); err != nil {
		return nil, fmt.Errorf("failed to parse OCR output: %w", err)
	}

	if result.Error != "" {
		return nil, fmt.Errorf("easyocr error: %s", result.Error)
	}

	return result.Text, nil
}
