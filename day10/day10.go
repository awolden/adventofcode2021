package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

var matchingPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func main() {

	chunks := readInput()
	part1(chunks)
	part2(chunks)

}

func part2(chunks [][]string) {

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	vals := []int{}
	for _, chunk := range chunks {
		stack := []string{}
		isValid := true
		for _, char := range chunk {
			if _, ok := matchingPairs[char]; ok {
				stack = append(stack, char)
				continue
			}
			topOfStack := stack[len(stack)-1]
			stack = RemoveIndex(stack, len(stack)-1)
			if char != matchingPairs[topOfStack] {
				isValid = false
				break
			}
		}
		if isValid {
			sum := 0
			for _, leftOverChar := range helpers.ReverseString(stack) {
				sum = sum * 5
				sum += points[matchingPairs[leftOverChar]]
			}
			vals = append(vals, sum)
		}
	}
	sort.Ints(vals)
	fmt.Println("sum", vals[int(math.Ceil(float64(len(vals)/2.0)))])
}

func part1(chunks [][]string) {

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	sum := 0
	for _, chunk := range chunks {
		fmt.Println("chunk", chunk)
		stack := []string{}
		for _, char := range chunk {
			//fmt.Println("going", char, stack)
			if _, ok := matchingPairs[char]; ok {
				stack = append(stack, char)
				continue
			}
			topOfStack := stack[len(stack)-1]
			stack = RemoveIndex(stack, len(stack)-1)
			if char != matchingPairs[topOfStack] {
				sum += points[char]
				break
			}
		}
	}
	fmt.Println("sum", sum)
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func readInput() [][]string {
	rawInput := helpers.GetFileArray("./input")
	arr := [][]string{}
	for _, line := range rawInput {
		chunks := strings.Split(line, "")
		arr = append(arr, chunks)
	}
	return arr
}
