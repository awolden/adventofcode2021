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
	curX    int
	curY    int
	total   int
	nextVal int
}

func (path Path) Hash() string {
	return strconv.Itoa(path.curX) + "," + strconv.Itoa(path.curY)
}

var pathVals = map[string]int{}
var visited = map[string]bool{}
var genMap = map[string]int{}
var best = 1<<(UintSize-1) - 1
var achievedEnd = false

func main() {
	input := readInput()
	newInput := multiplyArray(input)
	//fmt.Println(len(newInput), len(newInput[0]))
	takePath(newInput, Path{
		curX:    0,
		curY:    0,
		total:   0,
		nextVal: 0,
	})
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

func checkIfIsBetter(path Path) bool {
	//fmt.Println("mem", memo)

	if v, ok := pathVals[path.Hash()]; ok {
		if v <= path.total && v != 0 {
			//fmt.Println("bailing early cause we've been here", strconv.Itoa(path.curX)+","+strconv.Itoa(path.curY), memo, v, ok)
			return false
		} else {
			pathVals[path.Hash()] = path.total
		}
	} else {
		pathVals[path.Hash()] = path.total
	}
	return true
}

func takePath(input [][]int, path Path) {
	maxX := (len(input[0]) - 1)
	maxY := (len(input) - 1)

	queue := []Path{path}

	i := 0
	for len(queue) > 0 {
		i++
		if i%100000 == 0 {
			fmt.Println("i:", i)
		}
		//fmt.Println("visited", queue)

		path := queue[0]
		queue = queue[1:]
		//fmt.Println("path", path.Hash(), visited[path.Hash()])
		visited[path.Hash()] = true

		curX := path.curX
		curY := path.curY
		currentTotal := path.total

		if currentTotal < best && curX == maxX && curY == maxY {
			best = currentTotal
			hasBest = true
			fmt.Println("new best", best)
			achievedEnd = true
			continue
		}

		nextPaths := []Path{}

		if curY < maxY {
			nextPaths = append(nextPaths, Path{curX: curX, curY: curY + 1, total: pathVals[path.Hash()] + input[curY+1][curX], nextVal: input[curY+1][curX]})
			//takePath(input, curX, curY+1, currentTotal+input[curY+1][curX])
		}
		if curX < maxX {
			nextPaths = append(nextPaths, Path{curX: curX + 1, curY: curY, total: pathVals[path.Hash()] + input[curY][curX+1], nextVal: input[curY][curX+1]})
			//takePath(input, curX+1, curY, currentTotal+input[curY][curX+1])
		}
		if curX > 0 {
			nextPaths = append(nextPaths, Path{curX: curX - 1, curY: curY, total: pathVals[path.Hash()] + input[curY][curX-1], nextVal: input[curY][curX-1]})
			//takePath(input, curX-1, curY, currentTotal+input[curY][curX-1])
		}
		if curY > 0 {
			nextPaths = append(nextPaths, Path{curX: curX, curY: curY - 1, total: pathVals[path.Hash()] + input[curY-1][curX], nextVal: input[curY-1][curX]})
			// takePath(input, curX, curY-1, currentTotal+input[curY-1][curX])
		}

		//fmt.Println("next path", nextPaths)
		for _, nextPath := range nextPaths {
			checkIfIsBetter(nextPath)
			if !visited[nextPath.Hash()] && !ContainsPath(queue, nextPath) {
				//fmt.Println("appending path", nextPath.Hash())
				queue = append(queue, nextPath)
			}
		}

	}
}

func ContainsPath(s []Path, i Path) bool {
	for _, v := range s {
		if v.Hash() == i.Hash() {
			return true
		}
	}

	return false
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
