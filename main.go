package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// FunctionInfo represents information about a function
type FunctionInfo struct {
	Name       string
	Parameters []string
	ReturnType string
}

// FileInfo represents information about a file
type FileInfo struct {
	Name         string
	FunctionInfo []FunctionInfo
}

type ProjectConfig struct {
	Directory  string
	Language   string
	OutputPath string
}

func main() {
	config, err := promptConfiguration()
	if err != nil {
		fmt.Printf("Error during configuration: %v\n", err)
		return
	}

	files := extractFiles(config.Directory, config.Language)
	readmeTemplate := generateReadmeTemplate(files)
	writeReadmeFile(readmeTemplate, config.OutputPath)
}

func promptConfiguration() (ProjectConfig, error) {
	config := ProjectConfig{}

	questions := []*survey.Question{
		{
			Name: "directory",
			Prompt: &survey.Input{
				Message: "Enter the project directory path:",
				Default: ".",
			},
			Validate: func(val interface{}) error {
				str, ok := val.(string)
				if !ok {
					return fmt.Errorf("invalid input type")
				}
				if str == "" {
					return fmt.Errorf("directory path cannot be empty")
				}
				// Clean the path and check if it exists
				cleanPath := cleanWindowsPath(str)
				if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
					return fmt.Errorf("directory does not exist: %s", cleanPath)
				}
				return nil
			},
		},
		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Choose the programming language:",
				Options: []string{"go", "java", "python", "javascript"},
				Default: "go",
			},
		},
		{
			Name: "outputPath",
			Prompt: &survey.Input{
				Message: "Enter the README.md output path:",
				Default: "README.md",
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(questions, &config)
	if err != nil {
		return ProjectConfig{}, err
	}

	// Clean and handle paths
	config.Directory = cleanWindowsPath(config.Directory)
	config.OutputPath = handleOutputPath(cleanWindowsPath(config.OutputPath), config.Directory)

	// Print the paths for verification
	fmt.Printf("\nUsing directory: %s\n", config.Directory)

	return config, nil
}

func cleanWindowsPath(path string) string {
	// Remove any surrounding quotes
	path = strings.Trim(path, "\"'")

	// Convert forward slashes to backslashes for Windows
	path = strings.ReplaceAll(path, "/", "\\")

	// Clean the path
	path = filepath.Clean(path)

	return path
}

func handleOutputPath(outputPath, projectDir string) string {
	// If output path is absolute, use it as is
	if filepath.IsAbs(outputPath) {
		return outputPath
	}

	// If the output path is just a filename, place it in the current directory
	if !strings.Contains(outputPath, "\\") && !strings.Contains(outputPath, "/") {
		currentDir, err := os.Getwd()
		if err == nil {
			return filepath.Join(currentDir, outputPath)
		}
	}

	// Otherwise, resolve relative to project directory
	return filepath.Join(projectDir, outputPath)
}
