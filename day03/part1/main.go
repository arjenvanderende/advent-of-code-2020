package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	filename = "input.txt"
)

func main() {
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", filename, err)
	}

	valid, invalid := 0, 0
	x := 0
	for _, line := range lines {
		if line[x] == '.' {
			valid++
		} else {
			invalid++
		}

		x = (x + 3) % len(line)
	}

	fmt.Printf("Free=%d, Trees=%d", valid, invalid)
	fmt.Println()
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
