package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	filename = "input.txt"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", filename, err)
	}
	defer file.Close()

	valid := 0
	invalid := 0
	for {
		var min, max int
		var character, password string

		_, err = fmt.Fscanf(file, "%d-%d %1s: %s", &min, &max, &character, &password)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read line from %s: %v", filename, err)
		}

		// check valid
		actual := strings.Count(password, character)
		if actual >= min && actual <= max {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid=%d, Invalid=%d", valid, invalid)
	fmt.Println()
}
