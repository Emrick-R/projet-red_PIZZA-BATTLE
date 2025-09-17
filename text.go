package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Chemin spécifique du projet
	projectPath := "C:\\Users\\harol\\Documents\\Cours\\Projet RED\\projet-red_PIZZA-BATTLE"
	outputFile := "project_files.txt"

	// Store all file contents
	var allContents []string

	// Walk through the directory
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-go files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Skip the output file itself
		if info.Name() == outputFile {
			return nil
		}

		// Read file content
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return nil // Continue with other files
		}

		// Add file path and content to collection
		fileEntry := fmt.Sprintf("// FILE: %s\n%s", path, string(content))
		allContents = append(allContents, fileEntry)
		fmt.Printf("Added file: %s\n", path)

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}

	// Join all contents with separator
	joinedContent := strings.Join(allContents, "\n\n///\n\n")

	// Write to output file
	err = ioutil.WriteFile(outputFile, []byte(joinedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Printf("\nSuccess! All Go files from %s combined into %s\n", projectPath, outputFile)
}
