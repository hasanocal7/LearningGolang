package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}
		return countA_i > countA_j
	})

	return words
}

func main() {
	var userInput string
	fmt.Println("Enter a Array (Separate with commas and do not use square brackets): ")
	fmt.Scanln(&userInput)
	inputArray := strings.Split(userInput, ",")
	for i := range inputArray {
		inputArray[i] = strings.TrimSpace(inputArray[i])
	}
	output := sortByA(inputArray)
	fmt.Println("Sorted Array: ", output)
}
