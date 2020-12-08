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

	trees1 := checkSlope(lines, 1, 1)
	trees2 := checkSlope(lines, 3, 1)
	trees3 := checkSlope(lines, 5, 1)
	trees4 := checkSlope(lines, 7, 1)
	trees5 := checkSlope(lines, 1, 2)

	fmt.Printf("Trees = %d %d %d %d %d", trees1, trees2, trees3, trees4, trees5)
	fmt.Println()
	fmt.Printf("Total = %d", trees1*trees2*trees3*trees4*trees5)
	fmt.Println()
}

func checkSlope(lines []string, offsetX int, offsetY int) int {
	trees := 0
	x := 0
	for y := 0; y < len(lines); {
		line := lines[y]
		if line[x] != '.' {
			trees++
		}

		x = (x + offsetX) % len(line)
		y += offsetY
	}
	return trees
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
