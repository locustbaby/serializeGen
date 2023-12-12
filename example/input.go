package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// Read content from standard input
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Unable to read from standard input:", err)
		return
	}

	// Parse the input content and render it in YAML format
	var data interface{}
	err = yaml.Unmarshal(inputBytes, &data)
	if err != nil {
		fmt.Println("Unable to parse input content:", err)
		return
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("Unable to render as YAML:", err)
		return
	}

	// Output the rendered YAML content
	fmt.Println(string(yamlData))
}
