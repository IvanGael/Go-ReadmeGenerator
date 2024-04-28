package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
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

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <project_directory_path> <language>")
		return
	}

	projectDir := os.Args[1]
	language := os.Args[2]

	files := extractFiles(projectDir, language)
	readmeTemplate := generateReadmeTemplate(files)
	writeReadmeFile(readmeTemplate)
}

// extractFiles traverses the project directory and returns information about files
// extractFiles traverses the project directory and returns information about files with their tree structures
func extractFiles(dirPath, language string) []FileInfo {
	var files []FileInfo

	// Walk through the project directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Exclude files related to the .git folder and the go.sum file for Go projects
			if language == "go" && (strings.Contains(path, ".git") || strings.HasSuffix(path, "go.sum")) {
				return nil
			}

			// Calculate the relative path of the file
			relPath, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err
			}

			// Read the file
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Extract function information based on the language
			var functions []FunctionInfo
			switch language {
			case "go":
				functions = extractGoFunctions(string(content))
			case "java":
				// Add parsing logic for Java functions
			default:
				// Handle unsupported languages or provide a default parsing method
			}

			files = append(files, FileInfo{
				Name:         relPath, // Use relative path as file name
				FunctionInfo: functions,
			})
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return files
}

// extractGoFunctions extracts Go function information from file content
func extractGoFunctions(content string) []FunctionInfo {
	var functions []FunctionInfo

	// Regular expression to match Go function declarations
	funcRegex := regexp.MustCompile(`func\s+(\w+)\s*\((.*?)\)\s*(.*?)\s*{`)

	// Find all matches in the content
	matches := funcRegex.FindAllStringSubmatch(content, -1)

	// Process each match
	for _, match := range matches {
		functionName := match[1]
		parameters := strings.Split(match[2], ",")
		returnType := strings.TrimSpace(match[3])

		// Trim spaces from parameters
		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}

		// Add the function info to the list
		functions = append(functions, FunctionInfo{
			Name:       functionName,
			Parameters: parameters,
			ReturnType: returnType,
		})
	}

	return functions
}

// generateReadmeTemplate generates a README.md template
func generateReadmeTemplate(files []FileInfo) string {
	// Template for the README.md
	readmeTemplate := `
# Project Name

## Description

Add project description here.

## Usage

Add usage instructions here.

## Files

{{range .}}
### {{.Name}}

{{range .FunctionInfo}}
- {{.Name}}({{range $index, $element := .Parameters}}{{if $index}}, {{end}}{{$element}}{{end}}) - {{.ReturnType}}
{{end}}

{{end}}
`
	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var templateBuffer strings.Builder
	err = tmpl.Execute(&templateBuffer, files)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return templateBuffer.String()
}

// writeReadmeFile writes the README.md file
func writeReadmeFile(content string) {
	err := os.WriteFile("README.md", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// go run main.go "C:/Users/user/OneDrive/Bureau/Mes projets/Learning Go/AssoConnect" go
