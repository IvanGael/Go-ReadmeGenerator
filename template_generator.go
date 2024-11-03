package main

import (
	"fmt"
	"strings"
	"text/template"
)

// generateReadmeTemplate generates a README template based on the provided files.
func generateReadmeTemplate(files []FileInfo) string {
	const readmeTemplate = `# Project Name
## Description
Add project description here.

## Usage
Add usage instructions here.

## Files
{{range .}}
### {{.Name}}
{{range .FunctionInfo}}
- **{{.Name}}**({{join .Parameters ", "}}) {{.ReturnType}}
{{end}}
{{end}}
`

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl, err := template.New("readme").Funcs(funcMap).Parse(readmeTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return ""
	}

	var result strings.Builder
	err = tmpl.Execute(&result, files)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return ""
	}

	return result.String()
}
