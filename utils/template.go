package utils

import (
	"strings"
	"text/template"
)

// RenderTemplate renders a template with the provided values and returns the result as a string.
func RenderTemplate(templateContent string, values map[string]interface{}, leftDelim string, rightDelim string) (string, error) {
	// tmpl, err := template.New("yaml-template").Delims("[[", "]]").Parse(templateContent)
	tmpl, err := template.New("yaml-template").Delims(leftDelim, rightDelim).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var outputContent strings.Builder
	err = tmpl.Execute(&outputContent, values)
	if err != nil {
		return "", err
	}

	return outputContent.String(), nil
}
