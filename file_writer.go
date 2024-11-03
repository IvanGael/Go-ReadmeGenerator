package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeReadmeFile(content, outputPath string) {
	// Clean the output path
	outputPath = cleanWindowsPath(outputPath)

	// Get the directory part of the path
	dir := filepath.Dir(outputPath)

	// Create all necessary directories
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Write the file
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		fmt.Printf("Error writing README file: %v\n", err)
		return
	}

	fmt.Printf("README file successfully generated at: %s\n", outputPath)
}
