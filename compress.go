package main

import (
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputFileName := "input.txt"
	outputFileName := "output.txt"

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	compressor, err := flate.NewWriter(outputFile, flate.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	defer compressor.Close() // Ensure compressor is closed at the end of the function

	_, err = io.Copy(compressor, inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Compression completed. Output file:", outputFileName)
}

