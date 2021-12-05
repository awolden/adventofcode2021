package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Lines struct {
	lines []Line
	xMax  int
	yMax  int
}

func main() {
	line := readInput()
	part1(line)
	part2(line)
}

func part1(lines Lines) {
	counter := 0
	for x := 0; x <= lines.xMax; x++ {
		for y := 0; y <= lines.yMax; y++ {
			//testing all points
			if testAllLines(x, y, lines, true) >= 2 {
				counter++
			}
		}
	}
	fmt.Println("detectedpart 1:", counter)
}

func part2(lines Lines) {
	counter := 0
	for x := 0; x <= lines.xMax; x++ {
		for y := 0; y <= lines.yMax; y++ {
			//testing all points
			if testAllLines(x, y, lines, false) >= 2 {
				counter++
			}
		}
	}
	fmt.Println("detected part 2: ", counter)
}

func testAllLines(x int, y int, lines Lines, filterDiag bool) int {
	counter := 0
	for _, line := range lines.lines {
		if filterDiag && (line.x1 != line.x2 && line.y1 != line.y2) {
			continue
		}
		if inLine(line.x1, line.y1, line.x2, line.y2, x, y) {
			counter++
		}
	}
	return counter
}

// cross product shit
func inLine(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) bool {
	minX := int(math.Min(float64(x1), float64(x2)))
	maxX := int(math.Max(float64(x1), float64(x2)))
	minY := int(math.Min(float64(y1), float64(y2)))
	maxY := int(math.Max(float64(y1), float64(y2)))

	dxc := x3 - x1
	dyc := y3 - y1

	dxl := x2 - x1
	dyl := y2 - y1
	cross := dxc*dyl - dyc*dxl

	if cross != 0 {
		return false
	}
	if math.Abs(float64(dxl)) >= math.Abs(float64(dyl)) {
		return minX <= x3 && x3 <= maxX
	} else {
		return minY <= y3 && y3 <= maxY
	}
}

func readInput() Lines {
	rawInput := helpers.GetFileArray("./input")
	yMax := 0
	xMax := 0
	lines := make([]Line, 0)
	for _, line := range rawInput {
		pairs := strings.Split(line, " -> ")
		pair1 := strings.Split(pairs[0], ",")
		pair2 := strings.Split(pairs[1], ",")
		x1, _ := strconv.Atoi(pair1[0])
		y1, _ := strconv.Atoi(pair1[1])
		x2, _ := strconv.Atoi(pair2[0])
		y2, _ := strconv.Atoi(pair2[1])
		if x1 > xMax {
			xMax = x1
		}
		if y1 > yMax {
			yMax = y1
		}
		if x2 > xMax {
			xMax = x2
		}
		if y2 > yMax {
			yMax = y2
		}
		newLine := Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		lines = append(lines, newLine)
	}
	return Lines{
		lines: lines,
		yMax:  yMax,
		xMax:  xMax,
	}
}
