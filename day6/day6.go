package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

func main() {
	fishes := readInput()
	calc(fishes, 80)
	calc(fishes, 256)
}

func calc(fishes []int, days int) {
	m := make(map[int]int)
	for _, fish := range fishes {
		m[fish]++
	}
	fmt.Println(m)

	for day := 0; day < days; day++ {
		m = map[int]int{
			0: m[1], 1: m[2], 2: m[3], 3: m[4], 4: m[5], 5: m[6], 6: m[7] + m[0], 7: m[8], 8: m[0],
		}
	}

	sum := 0
	for _, fish := range m {
		sum += fish
	}
	fmt.Println("fishes:", sum)
}

func readInput() []int {
	rawInput := helpers.GetFileArray("./input")
	rawFish := strings.Split(rawInput[0], ",")
	intFish := make([]int, 0)
	for _, fishString := range rawFish {
		i, _ := strconv.Atoi(fishString)
		intFish = append(intFish, i)
	}

	return intFish
}
