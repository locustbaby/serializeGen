package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/locustbaby/serializeGen/utils"
)

func main() {
	// Define command-line flags
	valuesFile := flag.String("values", "", "values file")
	templateFile := flag.String("template", "", "template file or directory path")
	outputDir := flag.String("output", "", "output directory path")

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *templateFile == "" {
		flag.PrintDefaults()
		return
	}

	// Print flag values
	fmt.Println("Template file path:", *templateFile)
	fmt.Println("Output directory path:", *outputDir)

	// Read values.yaml file
	values, err := utils.ReadFile(*valuesFile)
	if err != nil {
		errorMessage := fmt.Sprintf("Error reading [%s]: %v", *valuesFile, err)
		fmt.Println(errorMessage)
		os.Exit(1)
		return
	}

	// Parse values.yaml file and convert its content to a Go map
	valuesMap, err := utils.ParseYAML(values)
	if err != nil {
		errorMessage := fmt.Sprintf("Error parsing [%s]: %v", *valuesFile, err)
		fmt.Println(errorMessage)
		os.Exit(1)
		return
	}

	// Create the output directory
	if *outputDir != "" {
		err = utils.CreateDirectory(*outputDir)
		if err != nil {
			errorMessage := fmt.Sprintf("Error parsing [%s]: %v", *outputDir, err)
			fmt.Println(errorMessage)
			os.Exit(1)
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
			return err
		}

		// Render the template
		outputContent, err := utils.RenderTemplate(string(templateContent), valuesMap)
		if err != nil {
			return err
		}

		// Print to standard output
		fmt.Println(outputContent)

		// Write to file
		if *outputDir != "" {
			err = utils.WriteFile(filepath.Join(*outputDir, info.Name()), outputContent)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error processing templates:", err)
		os.Exit(1)
		return
	}
}
