package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	filename    = "input.txt"
	targetValue = 2020
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", filename, err)
	}
	defer file.Close()

	values := make([]int, 0)
	for {
		var value int
		_, err = fmt.Fscanf(file, "%d", &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read line from %s: %v", filename, err)
		}

		values = append(values, value)
	}

	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			value1, value2 := values[i], values[j]
			if value1+value2 == targetValue {
				fmt.Printf("Found: %d + %d = %d", value1, value2, targetValue)
				fmt.Println()
				fmt.Printf("Product: %d", value1*value2)
				fmt.Println()
			}
		}
	}
}
