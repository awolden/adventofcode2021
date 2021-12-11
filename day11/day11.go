package main

import (
	"fmt"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)

func main() {
	input := readInput()
	funSteps(input, 1<<(UintSize-1)-1)
}

type Point struct {
	x int
	y int
}

func funSteps(octos [][]int, steps int) {
	//step 1
	totalFlashes := 0
	for i := 0; i < steps; i++ {
		flashes := []Point{}
		for y, row := range octos {
			for x, cell := range row {
				octos[y][x] = cell + 1
				if octos[y][x] == 10 {
					flashes = append(flashes, Point{x: x, y: y})
					octos[y][x] = 0
					totalFlashes++
				}
			}
		}
		totalFlashes += processSecondaryFlashes(octos, flashes)
		if checkAllOctos(octos) {
			fmt.Println("we have the end", i)
			return
		}
	}
	printOctos(octos, "")
	fmt.Println("flashes", steps, totalFlashes)
}

func checkAllOctos(octos [][]int) bool {
	for _, row := range octos {
		for _, cell := range row {
			if cell != 0 {
				return false
			}
		}
	}
	return true
}

func printOctos(octos [][]int, i string) {
	fmt.Println("octos:", i)
	fmt.Println("===========")
	for _, row := range octos {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func processSecondaryFlashes(octos [][]int, flashes []Point) int {
	upperXBound := len(octos[0]) - 1
	upperYBound := len(octos) - 1
	totalFlashes := 0
	for len(flashes) > 0 {
		flash := flashes[0]
		flashes = flashes[1:]
		x := flash.x
		y := flash.y

		isYEdge := y >= upperYBound
		isXEdge := x >= upperXBound

		if y != 0 && octos[y-1][x] > 0 {
			octos[y-1][x] += 1
			if octos[y-1][x] > 9 {
				totalFlashes++
				octos[y-1][x] = 0
				flashes = append(flashes, Point{y: y - 1, x: x})
			}
		}
		if y != 0 && !isXEdge && octos[y-1][x+1] > 0 {
			octos[y-1][x+1] += 1
			if octos[y-1][x+1] > 9 {
				totalFlashes++
				octos[y-1][x+1] = 0
				flashes = append(flashes, Point{y: y - 1, x: x + 1})
			}
		}
		if !isXEdge && octos[y][x+1] > 0 {
			octos[y][x+1] += 1
			if octos[y][x+1] > 9 {
				totalFlashes++
				octos[y][x+1] = 0
				flashes = append(flashes, Point{y: y, x: x + 1})
			}
		}
		if !isYEdge && !isXEdge && octos[y+1][x+1] > 0 {
			octos[y+1][x+1] += 1
			if octos[y+1][x+1] > 9 {
				totalFlashes++
				octos[y+1][x+1] = 0
				flashes = append(flashes, Point{y: y + 1, x: x + 1})
			}
		}

		if !isYEdge && octos[y+1][x] > 0 {
			octos[y+1][x] += 1
			if octos[y+1][x] > 9 {
				totalFlashes++
				octos[y+1][x] = 0
				flashes = append(flashes, Point{y: y + 1, x: x})
			}
		}

		if !isYEdge && x != 0 && octos[y+1][x-1] > 0 {
			octos[y+1][x-1] += 1
			if octos[y+1][x-1] > 9 {
				totalFlashes++
				octos[y+1][x-1] = 0
				flashes = append(flashes, Point{y: y + 1, x: x - 1})
			}
		}

		if y != 0 && x != 0 && octos[y-1][x-1] > 0 {
			octos[y-1][x-1] += 1
			if octos[y-1][x-1] > 9 {
				totalFlashes++
				octos[y-1][x-1] = 0
				flashes = append(flashes, Point{y: y - 1, x: x - 1})
			}
		}
		if x != 0 && octos[y][x-1] > 0 {
			octos[y][x-1] += 1
			if octos[y][x-1] > 9 {
				totalFlashes++
				octos[y][x-1] = 0
				flashes = append(flashes, Point{y: y, x: x - 1})
			}
		}
	}
	return totalFlashes
}

func readInput() [][]int {
	rawInput := helpers.GetFileArray("./input")
	allOctos := [][]int{}
	for _, line := range rawInput {
		octos := helpers.ConvertStrArr(strings.Split(line, ""))
		allOctos = append(allOctos, octos)
	}

	return allOctos
}
