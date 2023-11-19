package utils

import (
	"encoding/json"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// getFileType determines the file type based on the file extension.
func getFileType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".yaml", ".yml":
		return "yaml"
	case ".json":
		return "json"
	default:
		return "unknown"
	}
}

// isValidContent checks if the content is valid based on the file type.
func isValidContent(content []byte, fileType string) bool {
	switch fileType {
	case "yaml":
		return isValidYAML(content)
	case "json":
		return isValidJSON(content)
	default:
		// Unknown types are considered valid by default.
		return true
	}
}

// isValidYAML checks if the YAML content is valid.
func isValidYAML(content []byte) bool {
	var data interface{}
	err := yaml.Unmarshal(content, &data)
	return err == nil
}

// isValidJSON checks if the JSON content is valid.
func isValidJSON(content []byte) bool {
	var data interface{}
	err := json.Unmarshal(content, &data)
	return err == nil
}
