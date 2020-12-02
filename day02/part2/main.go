package main

import (
	"fmt"
	"io"
	"log"
	"os"
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
		var idx1, idx2 int
		var expected, password string

		_, err = fmt.Fscanf(file, "%d-%d %1s: %s", &idx1, &idx2, &expected, &password)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read line from %s: %v", filename, err)
		}

		// check valid
		char1, char2 := password[idx1-1:idx1], password[idx2-1:idx2]
		if (char1 == expected && char2 != expected) || (char1 != expected && char2 == expected) {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid=%d, Invalid=%d", valid, invalid)
	fmt.Println()
}
