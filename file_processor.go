package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func extractFiles(dirPath, language string) []FileInfo {
	var files []FileInfo

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && shouldProcessFile(path, language) {
			relPath, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			functions := extractFunctions(string(content), language)
			files = append(files, FileInfo{
				Name:         relPath,
				FunctionInfo: functions,
			})
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

	return files
}

func shouldProcessFile(path, language string) bool {
	excludedPaths := map[string][]string{
		"go":         {".git", "go.sum", "go.mod", ".idea", ".iml"},
		"java":       {".git", "target", "build", ".idea"},
		"python":     {".git", "__pycache__", ".pyc", ".env"},
		"javascript": {".git", "node_modules", "package-lock.json"},
	}

	if excluded, ok := excludedPaths[language]; ok {
		for _, exclude := range excluded {
			if strings.Contains(path, exclude) {
				return false
			}
		}
	}

	return true
}
