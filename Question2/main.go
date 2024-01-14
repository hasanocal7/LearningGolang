package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func reverseArray(arr []int) []int{
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
	   arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
 }

func recursive_output(n int) []int {
	var output []int
	output = append(output, n)
	if n > 2 {
		result := recursive_output(int(math.Floor(float64(n) / 2)))
		output = append(output, result...)
	}
	return output
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	textNumber, _ := reader.ReadString('\n')
	textNumber = strings.TrimSpace(textNumber)

	n, err := strconv.Atoi(textNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := recursive_output(n)

	for i := 0; i < len(result); i++ {
		fmt.Println(reverseArray(result)[i])
	}
}
