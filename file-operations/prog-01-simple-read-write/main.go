package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CreateFile function create a text file
func CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString("Hey this is my first file read/write operation program..")
	if err != nil {
		log.Fatalf("failed to write to file: %s: %s", file.Name(), err)
	}

	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile name: %s", file.Name())
}

// ReadFile reads from the text file created using CreateFile func
func ReadFile() {
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Panicf("failed to read from file: %s", err)
	}

	fmt.Printf("\nLength: %d", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v\n", err)
}

func main() {
	fmt.Printf("Creating a file to write into\n")
	CreateFile()

	fmt.Printf("\n\nReading file...")
	ReadFile()
}
