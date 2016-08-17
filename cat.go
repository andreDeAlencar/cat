package main

import (
	"fmt"
	"os"
	"io"
	"log"
)
func main() {
	file, err := os.Open(os.Args[1]) 
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
		if _, err := fmt.Printf(string(buffer[:n])); err != nil {
            log.Fatal(err)
        }
	}
	
}