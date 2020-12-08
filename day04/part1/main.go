package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filename = "input.txt"
)

func main() {
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", filename, err)
	}

	passports := groupPassports(lines)

	valid, invalid := 0, 0
	for _, passport := range passports {
		fmt.Println(passport)

		hasByr := strings.Index(passport, "byr:") >= 0
		hasIyr := strings.Index(passport, "iyr:") >= 0
		hasEyr := strings.Index(passport, "eyr:") >= 0
		hasHgt := strings.Index(passport, "hgt:") >= 0
		hasHcl := strings.Index(passport, "hcl:") >= 0
		hasEcl := strings.Index(passport, "ecl:") >= 0
		hasPid := strings.Index(passport, "pid:") >= 0
		// hasCid := strings.Index(passport, "cid:") >= 0
		if hasByr && hasIyr && hasEyr && hasHgt && hasHcl && hasEcl && hasPid {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid = %d, Invalid = %d", valid, invalid)
	fmt.Println()
}

func groupPassports(lines []string) []string {
	var passports []string
	var current string
	for _, line := range lines {
		if len(line) == 0 {
			passports = append(passports, current)
			current = ""
		} else {
			current += line + " "
		}
	}
	passports = append(passports, current)
	return passports
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
