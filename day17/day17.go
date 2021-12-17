package main

import (
	"fmt"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Area struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func main() {
	input := readInput()
	simulate(input)

}

func simulate(area Area) {
	fmt.Println(area)
	superHighestY := area.yMin
	hits := 0
	for xVel := area.xMax + 1000; xVel > 0; xVel-- {
		for yVel := area.yMin - 1000; yVel < 4000; yVel++ {
			xv := xVel
			yv := yVel

			x := 0
			y := 0
			highestY := area.yMin
			for x < area.xMax+1 && y > area.yMax {
				x += xv
				y += yv
				if y > highestY {
					highestY = y


				if x >= area.xMin && x <= area.xMax && y >= area.yMax && y <= area.yMin {
					if highestY > superHighestY {
						superHighestY = highestY
					}
					hits++
					break
				}

				//adjust velocities
				yv = yv - 1
				if xv > 0 {
					xv--
				} else if xv < 0 {
					xv++
				}

			}
		}
	}

	fmt.Println("super highest", superHighestY, hits)
}

func readInput() Area {
	rawInput := helpers.GetFileArray("./input")
	split1 := strings.Split(strings.Replace(rawInput[0], "target area: ", "", 1), ", ")
	xBoundary := helpers.ConvertStrArr(strings.Split(strings.Replace(split1[0], "x=", "", 1), ".."))
	yBoundary := helpers.ConvertStrArr(strings.Split(strings.Replace(split1[1], "y=", "", 1), ".."))

	return Area{
		xMin: xBoundary[0],
		xMax: xBoundary[1],
		yMin: yBoundary[1],
		yMax: yBoundary[0],
	}
}
