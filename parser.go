package main

import (
	"regexp"
	"strings"
)

func extractFunctions(content, language string) []FunctionInfo {
	switch language {
	case "go":
		return extractGoFunctions(content)
	case "java":
		return extractJavaFunctions(content)
	case "python":
		return extractPythonFunctions(content)
	case "javascript":
		return extractJavaScriptFunctions(content)
	default:
		return []FunctionInfo{}
	}
}

func extractGoFunctions(content string) []FunctionInfo {
	var functions []FunctionInfo
	funcRegex := regexp.MustCompile(`func\s+(\w+)\s*\((.*?)\)\s*(.*?)\s*{`)
	matches := funcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		functionName := match[1]
		parameters := strings.Split(match[2], ",")
		returnType := strings.TrimSpace(match[3])

		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}

		functions = append(functions, FunctionInfo{
			Name:       functionName,
			Parameters: parameters,
			ReturnType: returnType,
		})
	}

	return functions
}

// extractPythonFunctions extracts function information from Python code.
func extractPythonFunctions(content string) []FunctionInfo {
	re := regexp.MustCompile(`def\s+(\w+)\s*\((.*?)\)`)
	matches := re.FindAllStringSubmatch(content, -1)

	var functions []FunctionInfo
	for _, match := range matches {
		name := match[1]
		parameters := strings.Split(match[2], ",")
		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}
		functions = append(functions, FunctionInfo{Name: name, Parameters: parameters})
	}
	return functions
}

// extractJavaFunctions extracts function information from Java code.
func extractJavaFunctions(content string) []FunctionInfo {
	re := regexp.MustCompile(`(public|protected|private|static|\s) +[\w<>\[\]]+\s+(\w+) *\(([^)]*)\)`)
	matches := re.FindAllStringSubmatch(content, -1)

	var functions []FunctionInfo
	for _, match := range matches {
		name := match[2]
		parameters := strings.Split(match[3], ",")
		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}
		functions = append(functions, FunctionInfo{Name: name, Parameters: parameters})
	}
	return functions
}

// extractJavaScriptFunctions extracts function information from JavaScript code.
func extractJavaScriptFunctions(content string) []FunctionInfo {
	re := regexp.MustCompile(`function\s+(\w+)\s*\((.*?)\)|const\s+(\w+)\s*=\s*\((.*?)\)\s*=>`)
	matches := re.FindAllStringSubmatch(content, -1)

	var functions []FunctionInfo
	for _, match := range matches {
		var name string
		var parameters []string
		if match[1] != "" {
			name = match[1]
			parameters = strings.Split(match[2], ",")
		} else {
			name = match[3]
			parameters = strings.Split(match[4], ",")
		}
		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}
		functions = append(functions, FunctionInfo{Name: name, Parameters: parameters})
	}
	return functions
}
