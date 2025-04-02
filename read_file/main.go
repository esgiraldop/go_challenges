package main

import (
	// "fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	// file, err := os.ReadFile(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// bs := make([]byte, 9999)
	// file.Read(bs)
	// fmt.Println("file: ", string(file))
	io.Copy(os.Stdout, file)
}
