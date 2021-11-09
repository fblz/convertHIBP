package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputPath string
	flag.StringVar(&inputPath, "InputFile", "none", "Specify the ascii hibp file")
	var outputPath string
	flag.StringVar(&outputPath, "OutputFile", "none", "Specify the name for the binary target file")
	flag.Parse()

	if inputPath == "none" || outputPath == "none" {
		fmt.Println(os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		stringHash := line[:40]

		data, err := hex.DecodeString(stringHash)
		if err != nil {
			log.Fatal(err)
		}

		bytesWritten, err := writer.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		if bytesWritten != 20 {
			log.Fatal("Write did not return 20 bytes. That is bad. Aborting!")
		}
	}

	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
