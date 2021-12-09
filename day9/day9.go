package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Point struct {
	x     int
	y     int
	value int
}

type byLen [][]Point

func (s byLen) Len() int {
	return len(s)
}
func (s byLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLen) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (pnt *Point) equals(otherPoint Point) bool {
	return otherPoint.x == pnt.x && otherPoint.y == pnt.y
}

func main() {
	input := readInput()
	//fmt.Println(input)

	partOne(input)
	partTwo(input)

}
func partTwo(input [][]int) {
	basins := make([][]Point, 0)
	pointMap := make(map[Point]bool)

	//loop over every point
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			point := Point{x: x, y: y, value: input[y][x]}
			if _, ok := pointMap[point]; ok {
				continue
			}
			// discover fullsize of basin in each direction
			basin := collectBasin(input, point, make([]Point, 0), pointMap)
			if len(basin) > 0 {
				basins = append(basins, basin)
			}
		}
	}

	sort.Sort(byLen(basins))

	product := 1
	for i := 0; i < 3; i++ {
		basin := basins[i]
		product *= len(basin)
	}
	fmt.Println("basin sum", product)
}

func collectBasin(layout [][]int, point Point, basin []Point, pointMap map[Point]bool) []Point {
	newBasin := make([]Point, 0)
	if _, ok := pointMap[point]; ok {
		return basin
	}
	if point.value == 9 {
		return basin
	} else {
		pointMap[point] = true
		newBasin = append(newBasin, point)
	}

	// check north
	if point.y != 0 {
		newPoint := Point{x: point.x, y: point.y - 1, value: layout[point.y-1][point.x]}
		newBasin = append(newBasin, collectBasin(layout, newPoint, basin, pointMap)...)
	}
	// check west
	if point.x != len(layout[point.y])-1 {
		newPoint := Point{x: point.x + 1, y: point.y, value: layout[point.y][point.x+1]}
		newBasin = append(newBasin, collectBasin(layout, newPoint, basin, pointMap)...)
	}
	// check south
	if point.y != len(layout)-1 {
		newPoint := Point{x: point.x, y: point.y + 1, value: layout[point.y+1][point.x]}
		newBasin = append(newBasin, collectBasin(layout, newPoint, basin, pointMap)...)
	}
	// check east
	if point.x != 0 {
		newPoint := Point{x: point.x - 1, y: point.y, value: layout[point.y][point.x-1]}
		newBasin = append(newBasin, collectBasin(layout, newPoint, basin, pointMap)...)
	}
	return newBasin
}

func partOne(input [][]int) {
	lowNums := make([]int, 0)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			isLow := true
			curVal := input[y][x]

			// check north
			if y != 0 && input[y-1][x] <= curVal {
				isLow = false
			}
			// check west
			if x != len(input[y])-1 && input[y][x+1] <= curVal {
				isLow = false
			}
			// check south
			if y != len(input)-1 && input[y+1][x] <= curVal {
				isLow = false
			}
			// check east
			if x != 0 && input[y][x-1] <= curVal {
				isLow = false
			}

			if isLow {
				// fmt.Println("isLow", curVal, x, y)
				lowNums = append(lowNums, curVal)
			}
		}
	}
	sum := 0
	for _, num := range lowNums {
		sum += num + 1
	}
	fmt.Println("sum", sum)
}

func readInput() [][]int {
	rawInput := helpers.GetFileArray("./input")
	input := make([][]int, 0)
	for _, line := range rawInput {
		numStrings := strings.Split(line, "")
		input = append(input, helpers.ConvertStrArr(numStrings))
	}
	return input
}
