package main

import (
	"fmt"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Instruction struct {
	first     string
	second    string
	insertion string
}

type GameInput struct {
	polymer      []string
	instructions map[string]Instruction
}

const UintSize = 32 << (^uint(0) >> 32 & 1)

func main() {
	input := readInput()
	runSteps(input, 40)
}

func runSteps(gameInput GameInput, steps int) {
	pairs := map[string]int{}
	letterBuffer := map[string]int{}

	for i := 0; i < len(gameInput.polymer)-1; i++ {
		currentFirst := gameInput.polymer[i]
		currentSecond := gameInput.polymer[i+1]
		pairs[currentFirst+currentSecond]++
	}

	for key, value := range pairs {
		letters := strings.Split(key, "")
		letterBuffer[letters[0]] += value
		letterBuffer[letters[1]] += value
	}

	for i := 0; i < steps; i++ {
		newPairs := map[string]int{}
		for key, value := range pairs {
			newPairs[key] = value
		}
		for _, instruction := range gameInput.instructions {
			if _, ok := pairs[instruction.first+instruction.second]; ok {
				newPairs[instruction.first+instruction.insertion] += pairs[instruction.first+instruction.second]
				newPairs[instruction.insertion+instruction.second] += pairs[instruction.first+instruction.second]
				newPairs[instruction.first+instruction.second] -= pairs[instruction.first+instruction.second]
				letterBuffer[instruction.insertion] += pairs[instruction.first+instruction.second]
			}
		}
		pairs = newPairs
	}
	min := 1<<(UintSize-1) - 1
	max := 0

	for _, v := range letterBuffer {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func readInput() GameInput {
	rawInput := helpers.GetFileArray("./input")

	polymer := strings.Split(rawInput[0], "")
	instructions := map[string]Instruction{}

	for i := 2; i < len(rawInput); i++ {
		tokens := strings.Split(rawInput[i], " -> ")
		pair := strings.Split(tokens[0], "")
		instructions[pair[0]+pair[1]] = Instruction{first: pair[0], second: pair[1], insertion: tokens[1]}
	}

	return GameInput{
		instructions: instructions,
		polymer:      polymer,
	}
}
