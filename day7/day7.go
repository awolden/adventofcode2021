package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)

func main() {
	crabs := readInput()
	// calc(crabs, false)
	calc(crabs, true)

}

func calc(crabs []int, exp bool) {
	min := 1<<(UintSize-1) - 1
	for i := helpers.Min(crabs); i < helpers.Max(crabs); i++ {
		fuel := getFuelCost(i, crabs, exp)
		if fuel < min {
			fmt.Println("fuel cost", i, fuel)
			min = fuel
		}
	}
	fmt.Println("final fuel cost", min)
}

func getExp(num int) int {
	exp := 0
	for i := 0; i < num; i++ {
		exp += (i + 1)
	}
	return exp
}

func getFuelCost(num int, crabs []int, exp bool) int {
	sum := 0
	for _, i := range crabs {
		distance := 0
		if i > num {
			distance = i - num
		} else {
			distance = num - i
		}
		if exp {
			distance = getExp(distance)
		}
		sum += distance
	}
	return sum
}

func readInput() []int {
	rawInput := helpers.GetFileArray("./input")
	splited := strings.Split(rawInput[0], ",")
	intCrabs := make([]int, 0)
	for _, fishString := range splited {
		i, _ := strconv.Atoi(fishString)
		intCrabs = append(intCrabs, i)
	}
	return intCrabs
}
