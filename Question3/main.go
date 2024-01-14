package main

import (
	"fmt"
	"strings"
)

func mostRepeated(arr []string) string {
	counts := make(map[string]int)

	for _, val := range arr {
		counts[val]++
	}

	var mostRepeated string
	maxCount := 0

	for key, count := range counts {
		if count > maxCount {
			maxCount = count
			mostRepeated = key
		}
	}

	return mostRepeated
}

func main() {
	var userInput string
	fmt.Println("Enter a Array (Separate with commas and do not use square brackets): ")
	fmt.Scanln(&userInput)
	inputArray := strings.Split(userInput, ",")
	for i := range inputArray {
		inputArray[i] = strings.TrimSpace(inputArray[i])
	}
	output := mostRepeated(inputArray)
	fmt.Println(output)
}
