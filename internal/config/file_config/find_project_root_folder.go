package file_config

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindProjectRoot finds the root directory of the project by looking for a marker file or directory
func FindProjectRoot(marker string) (string, error) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse the directory tree upwards until the marker file or directory is found
	for {
		// Check if the marker file or directory exists in the current directory
		markerPath := filepath.Join(dir, marker)
		if _, err := os.Stat(markerPath); err == nil {
			return markerPath, nil
		}

		// Move up one directory level
		parentDir := filepath.Dir(dir)
		// If we've reached the root directory, stop
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("marker %s not found in any parent directory", marker)
}
