package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

func main() {
	input := readInput()
	getGammaAndEpsilon(input)
	oxy := getLifeSupport(input, 0, "oxy")
	co2 := getLifeSupport(input, 0, "co2")
	fmt.Println("part2", oxy*co2)
}

func getLifeSupport(input [][]string, starting int, typeOfSupport string) int64 {
	ones := make([][]string, 0)
	zeros := make([][]string, 0)
	if len(input) == 1 {
		num, _ := strconv.ParseInt(strings.Join(input[0][:], ""), 2, 32)
		return num
	}

	for j := 0; j < len(input); j++ {
		if input[j][starting] == "1" {
			ones = append(ones, input[j])
		} else {
			zeros = append(zeros, input[j])
		}
	}
	if typeOfSupport == "oxy" {
		if len(zeros) > len(ones) {
			return getLifeSupport(zeros, starting+1, typeOfSupport)
		} else {
			return getLifeSupport(ones, starting+1, typeOfSupport)
		}
	}
	if typeOfSupport == "co2" {
		if len(zeros) > len(ones) {
			return getLifeSupport(ones, starting+1, typeOfSupport)
		} else {
			return getLifeSupport(zeros, starting+1, typeOfSupport)
		}
	}
	// should never be hit, but wonky
	return 0
}

func getGammaAndEpsilon(input [][]string) {
	limit := len(input[0])
	gamma := ""
	epsilon := ""
	for i := 0; i < limit; i++ {
		counter := 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == "1" {
				counter++
			}
		}
		if counter > len(input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gNum, _ := strconv.ParseInt(gamma, 2, 32)
	eNum, _ := strconv.ParseInt(epsilon, 2, 32)
	fmt.Println("part1", gNum*eNum)
}

func readInput() [][]string {
	rawInput := helpers.GetFileArray("./input")
	arr := make([][]string, 0)

	for _, line := range rawInput {
		digits := strings.Split(string(line), "")
		digitArr := make([]string, 0)
		for _, digit := range digits {
			digitArr = append(digitArr, digit)
		}
		arr = append(arr, digitArr)
	}
	return arr
}
