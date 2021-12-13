package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	direction string
	location  int
}

type GameInput struct {
	points map[Point]bool
	folds  []Fold
}

func main() {
	input := readInput()
	//fmt.Println(input)
	// printPaper(input)
	//calc(input, 1)
	calc(input, len(input.folds))
}

func calc(input GameInput, loops int) {
	for i, fold := range input.folds {
		if i >= loops {
			break
		}
		newPoints := map[Point]bool{}
		for point := range input.points {
			if fold.direction == "vertical" {
				if point.x == fold.location {
					//noop
				} else if point.x > fold.location {
					newPoints[Point{x: fold.location - (point.x - fold.location), y: point.y}] = true
				} else {
					newPoints[point] = true
				}
			} else {
				if point.y == fold.location {
					// noop
				} else if point.y > fold.location {
					newPoints[Point{x: point.x, y: fold.location - (point.y - fold.location)}] = true
				} else {
					newPoints[point] = true
				}
			}
		}
		input.points = newPoints

	}
	fmt.Println("done", len(input.points))
	printPaper(input)
}

func printPaper(input GameInput) {
	maxX := 0
	maxY := 0
	for point := range input.points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	fmt.Println("")
	fmt.Println("Printing Paper", maxX, maxY)
	fmt.Println("=======================")
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if (input.points[Point{x: x, y: y}]) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func readInput() GameInput {
	rawInput := helpers.GetFileArray("./input")
	folds := []Fold{}
	points := map[Point]bool{}

	for _, line := range rawInput {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "fold") {

			parts := strings.Split(line, " ")
			subParts := strings.Split(parts[2], "=")
			num, _ := strconv.Atoi(subParts[1])
			if subParts[0] == "x" {
				folds = append(folds, Fold{direction: "vertical", location: num})
			} else {
				folds = append(folds, Fold{direction: "horizontal", location: num})
			}
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points[Point{x: x, y: y}] = true
		}
	}

	return GameInput{
		folds:  folds,
		points: points,
	}
}
