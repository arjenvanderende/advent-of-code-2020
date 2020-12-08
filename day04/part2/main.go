package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		hasByr := matches(passport, "\\bbyr:(19[2-9][0-9]|2000|2001|2002)\\b")
		hasIyr := matches(passport, "\\biyr:(201[0-9]|2020)\\b")
		hasEyr := matches(passport, "\\beyr:(202[0-9]|2030)\\b")
		hasHgt := matches(passport, "\\bhgt:(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)\\b")
		hasHcl := matches(passport, "\\bhcl:#[0-9a-f]{6}\\b")
		hasEcl := matches(passport, "\\becl:(amb|blu|brn|gry|grn|hzl|oth)\\b")
		hasPid := matches(passport, "\\bpid:[0-9]{9}\\b")
		if hasByr && hasIyr && hasEyr && hasHgt && hasHcl && hasEcl && hasPid {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid = %d, Invalid = %d", valid, invalid)
	fmt.Println()
}

func matches(line string, pattern string) bool {
	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		log.Fatalf("Regexp %s failed to match on line %s: %v", pattern, line, err)
	}

	return matched
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
