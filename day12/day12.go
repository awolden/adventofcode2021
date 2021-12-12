package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/awolden/adventofcode2021/helpers"
)

func main() {
	input := readInput()
	fmt.Println(input)

	paths := takePath(input, "start", []string{})
	fmt.Println("paths", len(paths))
}

func takePath(theMap map[string][]string, currentRoom string, visited []string) [][]string {
	allPaths := [][]string{}
	newlyVisited := make([]string, len(visited))
	copy(newlyVisited, visited)
	newlyVisited = append(newlyVisited, currentRoom)
	nextRooms := theMap[currentRoom]

	for _, nextRoom := range nextRooms {
		if nextRoom == "end" {
			allPaths = append(allPaths, newlyVisited)
			continue
		}
		if nextRoom == "start" {
			continue
		}
		if (IsLower(nextRoom) && !helpers.ContainsString(newlyVisited, nextRoom)) ||
			(IsLower(nextRoom) && !checkIfArrHasSmallDupes(newlyVisited)) ||
			!IsLower(nextRoom) {
			allPaths = append(allPaths, takePath(theMap, nextRoom, newlyVisited)...)
		}
	}
	return allPaths
}

func checkIfArrHasSmallDupes(arr []string) bool {

	m := map[string]int{}
	for _, item := range arr {
		if !IsLower(item) {
			continue
		}
		m[item] += 1
		if m[item] == 2 {
			return true
		}
	}
	return false
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readInput() map[string][]string {
	rawInput := helpers.GetFileArray("./input")
	theMap := make(map[string][]string)

	for _, line := range rawInput {
		part := strings.Split(line, "-")
		destinations := theMap[part[0]]
		destinations2 := theMap[part[1]]
		theMap[part[0]] = append(destinations, part[1])
		theMap[part[1]] = append(destinations2, part[0])
	}

	return theMap
}
