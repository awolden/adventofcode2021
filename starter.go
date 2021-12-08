package main

import (
	"fmt"

	"github.com/awolden/adventofcode2021/helpers"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)

func main() {
	input := readInput()
	fmt.Println(input)

}

func readInput() []string {
	rawInput := helpers.GetFileArray("./input")

	return rawInput
}
