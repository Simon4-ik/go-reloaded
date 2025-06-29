package main

import (
	"bufio"
	"fmt"
	"os"

	"go-reloaded/internal"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer outFile.Close()

	// Set up scanner for reading the input file and buffered writer for output
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)

	// Process each line from the input file
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := internal.ProcessText(line) // Assuming you have a ProcessText function in your internal package
		bytesWritten, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			return
		}
		if bytesWritten < len(processedLine)+1 {
			fmt.Println("Warning: Not all bytes were written to the output file.")
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// Flush the buffered writer to ensure all data is written to the file
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error flushing writer: %v\n", err)
	}

	fmt.Println("Processing complete. Check", outputFile)
}
