package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type direction struct {
	direction string
	value     int
}

func main() {
	input := readInput()
	partOne(input)
	partTwo(input)
}

func partOne(input []direction) {
	position := 0
	depth := 0

	for _, command := range input {
		if command.direction == "forward" {
			position += command.value
		} else if command.direction == "down" {
			depth += command.value
		} else if command.direction == "up" {
			depth -= command.value
		}
	}

	fmt.Println("part one", position*depth)
}

func partTwo(input []direction) {
	position := 0
	depth := 0
	aim := 0

	for _, command := range input {
		if command.direction == "forward" {
			position += command.value
			depth += aim * command.value
		} else if command.direction == "down" {
			aim += command.value
		} else if command.direction == "up" {
			aim -= command.value
		}
	}

	fmt.Println("part one", position*depth)
}

func readInput() []direction {
	rawInput := helpers.GetFileArray("./input")
	arr := make([]direction, 0)

	for _, line := range rawInput {
		pair := strings.Split(string(line), " ")
		inputInt, _ := strconv.Atoi(pair[1])
		arr = append(arr, direction{direction: pair[0], value: inputInt})
	}
	return arr
}
