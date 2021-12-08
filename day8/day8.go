package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type input struct {
	scrambles map[string]int
	tokens    []string
}

func main() {
	inputs := readInput()
	fmt.Println("input", inputs)

	// count
	count := 0
	for _, input := range inputs {
		for _, token := range input.tokens {
			for scramble, v := range input.scrambles {
				if checkIfScrambleMatches(token, scramble) && (v == 1 || v == 4 || v == 7 || v == 8) {
					count++
				}
			}
		}
	}
	fmt.Println("part 1 counted", count)

	sum := 0
	for _, input := range inputs {
		nums := make([]int, 0)
		for _, token := range input.tokens {
			for scramble := range input.scrambles {
				if checkIfScrambleMatches(token, scramble) {
					nums = append(nums, input.scrambles[scramble])
				}
			}
		}
		num, _ := strconv.Atoi(helpers.ConvertIntArr(nums))

		sum += num
	}
	fmt.Println("part 2 counted", sum)
}

func checkIfScrambleMatches(input string, scramble string) bool {
	inputArr := strings.Split(input, "")
	scrambleArr := strings.Split(scramble, "")
	sort.Strings(inputArr)
	sort.Strings(scrambleArr)
	return helpers.ArrEqual(inputArr, scrambleArr)
}

func readInput() []input {
	rawInput := helpers.GetFileArray("./input")

	inputs := make([]input, 0)
	for _, line := range rawInput {
		splittted := strings.Split(line, " | ")
		tokens := strings.Split(splittted[1], " ")
		m := make(map[string]int)
		scrambles := strings.Split(splittted[0], " ")
		sort.Sort(byLength(scrambles))
		rightSideCodes := make([]string, 0)
		fourCodes := make([]string, 0)
		topCode := make([]string, 0)

		for _, token := range scrambles {
			chars := strings.Split(token, "")
			if len(token) == 2 {
				rightSideCodes = strings.Split(token, "")
				m[token] = 1
			} else if len(token) == 3 {
				topCode = helpers.Filter(chars, func(char string) bool {
					return !helpers.ContainsString(rightSideCodes, char)
				})
				m[token] = 7
			} else if len(token) == 7 {
				m[token] = 8
			} else if len(token) == 4 {
				fourCodes = strings.Split(token, "")
				m[token] = 4
			} else if len(token) == 5 {
				testFor3 := len(helpers.Filter(chars, func(char string) bool {
					return !helpers.ContainsString(rightSideCodes, char)
				})) == 3
				testChars := helpers.Filter(fourCodes, func(char string) bool {
					return !helpers.ContainsString(rightSideCodes, char)
				})
				testFor2 := len(helpers.Filter(chars, func(char string) bool {
					return !helpers.ContainsString(testChars, char)
				})) == 4
				if testFor3 {
					m[token] = 3
				} else if testFor2 {
					m[token] = 2
				} else {
					m[token] = 5
				}

			} else if len(token) == 6 {
				testFor9 := len(helpers.Filter(chars, func(char string) bool {
					return !helpers.ContainsString(fourCodes, char) && !helpers.ContainsString(topCode, char)
				})) == 1
				testFor6 := len(helpers.Filter(chars, func(char string) bool {
					return !helpers.ContainsString(rightSideCodes, char)
				})) == 5

				if testFor9 {
					m[token] = 9
				} else if testFor6 {
					m[token] = 6
				} else {
					m[token] = 0
				}

			}
		}
		inputs = append(inputs, input{
			scrambles: m,
			tokens:    tokens,
		})
	}

	return inputs

}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
