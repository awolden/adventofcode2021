package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	fmt.Println("increases: window size: 1", countIncreases(input, 1))
	fmt.Println("increases: window size: 3", countIncreases(input, 3))
}

func countIncreases(input []int, windowSize int) int {
	counter := 0
	for i := windowSize + 1; i <= len(input); i += 1 {
		previousWindowSum := sum(input[i-windowSize-1 : i-1])
		currentWindowSum := sum(input[i-windowSize : i])
		if currentWindowSum > previousWindowSum {
			counter++
		}
	}
	return counter
}

func sum(input []int) int {
	sum := 0
	for _, value := range input {
		sum += value
	}
	return sum
}

func readInput() []int {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReader(file)

	var arr = make([]int, 0)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		inputInt, _ := strconv.Atoi(string(line))
		arr = append(arr, inputInt)
	}
	return arr
}
