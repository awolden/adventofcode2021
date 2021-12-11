package main

import (
	"fmt"

	"github.com/awolden/adventofcode2021/helpers"
)

func main() {
	input := readInput()
	fmt.Println(input)

}

func readInput() []string {
	rawInput := helpers.GetFileArray("./input")

	return rawInput
}
