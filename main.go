package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/locustbaby/stt/utils"
)

func main() {
	// Define command-line flags
	valuesFile := flag.String("v", "", "Values file")
	templateFile := flag.String("t", "", "Template file or directory path")
	outputDir := flag.String("o", "", "Output Directory path, only [Dir]")

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *templateFile == "" {
		flag.PrintDefaults()
		return
	}

	// Read values.yaml file
	values, err := utils.ReadFile(*valuesFile)
	if err != nil {
		utils.HandleError("Error reading", *valuesFile, err)
		return
	}

	// Parse values.yaml file and convert its content to a Go map
	valuesMap, err := utils.ParseYAML(values)
	if err != nil {
		utils.HandleError("Error parsing", *valuesFile, err)
		return
	}

	// Create the output directory
	if *outputDir != "" {
		err = utils.CreateDirectory(*outputDir)
		if err != nil {
			utils.HandleError("Error creating", *outputDir, err)
			return
		}
	}

	// Traverse all files in the templates directory
	err = filepath.Walk(*templateFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Ignore directories
		if info.IsDir() {
			return nil
		}

		// Read the template file
		templateContent, err := utils.ReadFile(path)
		if err != nil {
			utils.HandleError("Error reading template", path, err)
			return err
		}

		// Render the template
		outputContent, err := utils.RenderTemplate(string(templateContent), valuesMap)
		if err != nil {
			utils.HandleError("Error rendering template", path, err)
			return err
		}

		// Write to file
		if *outputDir != "" {
			err = utils.WriteFile(filepath.Join(*outputDir, info.Name()), outputContent)
			if err != nil {
				utils.HandleError("Error writing to", filepath.Join(*outputDir, info.Name()), err)
				return err
			}
		}

		return nil
	})

	utils.HandleError("Error processing templates", *templateFile, err)
}
