package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/awolden/adventofcode2021/helpers"
)

var wg sync.WaitGroup

const UintSize = 32 << (^uint(0) >> 32 & 1)

type Point struct {
	x int
	y int
}

type Path struct {
	points []Point
	total  int
}

var memo = map[string]int{}
var genMap = map[string]int{}
var best = 1<<(UintSize-1) - 1
var achievedEnd = false

func main() {
	input := readInput()
	newInput := multiplyArray(input)
	//fmt.Println(len(newInput), len(newInput[0]))
	takePath(newInput, 0, 0, 0)
	// fmt.Println()
	// for i := 0; i < len(newInput); i++ {
	// 	for j := 0; j < len(newInput[0]); j++ {
	// 		fmt.Print(newInput[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// multiplyArray(input, 5)
	// takePath(input, 0, 0, 0)
}

var hasBest = false

func takePath(input [][]int, curX int, curY int, currentTotal int) {
	maxX := (len(input[0]) - 1)
	maxY := (len(input) - 1)
	if v, ok := memo[strconv.Itoa(curX)+","+strconv.Itoa(curY)]; ok {
		if v < currentTotal {
			// fmt.Println("bailing early cause we've been here", strconv.Itoa(curX)+","+strconv.Itoa(curY))
			return
		} else {
			memo[strconv.Itoa(curX)+","+strconv.Itoa(curY)] = currentTotal
		}
	} else {
		memo[strconv.Itoa(curX)+","+strconv.Itoa(curY)] = currentTotal
	}

	if hasBest && currentTotal > best {
		return
	}

	if currentTotal < best && curX == maxX && curY == maxY {
		best = currentTotal
		hasBest = true
		fmt.Println("new best", best)
		achievedEnd = true
		return
	}

	if curY < maxY {
		takePath(input, curX, curY+1, currentTotal+input[curY+1][curX])
	}
	if curX < maxX {
		takePath(input, curX+1, curY, currentTotal+input[curY][curX+1])
	}
	if curX > 0 {
		takePath(input, curX-1, curY, currentTotal+input[curY][curX-1])
	}
	if curY > 0 {
		takePath(input, curX, curY-1, currentTotal+input[curY-1][curX])
	}

}

func multiplyArray(foo [][]int) [][]int {
	size := 5 * len(foo[0])
	newArr := [][]int{}

	for y := 0; y < size; y++ {
		newRow := []int{}
		for x := 0; x < size; x++ {
			pageX := int(math.Floor(float64(x) / 10))
			pageY := int(math.Floor(float64(y) / 10))
			// fmt.Println("he", x, y, x%10, y%10)
			val := foo[y%10][x%10] + pageX + pageY
			if val > 9 {
				val = val % 9
			}
			newRow = append(newRow, val)
		}
		newArr = append(newArr, newRow)
	}
	return newArr
}

func increaseVals(foo []int, amount int) []int {
	bar := []int{}
	for _, i := range foo {
		bar = append(bar, (i+amount)%9)
	}
	return bar
}

func readInput() [][]int {
	rawInput := helpers.GetFileArray("./input")
	gameboard := [][]int{}

	for _, line := range rawInput {
		split := strings.Split(line, "")
		gameboard = append(gameboard, helpers.ConvertStrArr(split))
	}

	return gameboard
}
