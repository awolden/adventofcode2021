package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type gameSetup struct {
	numbers []int
	boards  [][][]int
}

func main() {
	game := readInput()
	fmt.Println(game)
	getAnswer(game)
	get2ndAnswer(game)
}

func getAnswer(game gameSetup) {
	for i := 1; i < len(game.numbers); i++ {
		for y := 0; y < len(game.boards); y++ {
			result := isBoardComplete(game.numbers[:i], game.boards[y])
			if result {
				fmt.Println("Found winning board", game.boards[y])
				sumProduct(game.numbers[:i], game.boards[y])
				return
			}
		}
	}
}

func get2ndAnswer(game gameSetup) {
	for i := (len(game.numbers)); i >= 0; i-- {
		for y := 0; y < len(game.boards); y++ {
			result := isBoardComplete(game.numbers[:i], game.boards[y])
			if !result {
				fmt.Println("Found losing board", game.boards[y])
				sumProduct(game.numbers[:i+1], game.boards[y])
				return
			}
		}
	}
}

func sumProduct(drawnNums []int, board [][]int) {
	sum := 0
	for _, row := range board {
		for _, num := range row {
			if !helpers.Contains(drawnNums, num) {
				sum += num
			}
		}
	}
	fmt.Println("answer", sum, drawnNums[len(drawnNums)-1], sum*drawnNums[len(drawnNums)-1])
}

func isBoardComplete(drawnNums []int, board [][]int) bool {
	// check horizontal
	for _, row := range board {
		if rowIsComplete(row, drawnNums) {
			return true
		}
	}
	//check vert
	for x := range board[0] {
		vertRow := make([]int, 0)
		for y := 0; y < len(board); y++ {
			vertRow = append(vertRow, board[y][x])
		}
		if rowIsComplete(vertRow, drawnNums) {
			return true
		}
	}
	return false
}

func rowIsComplete(row []int, drawnNums []int) bool {
	counter := 0
	for _, num := range drawnNums {
		if helpers.Contains(row, num) {
			counter++
			if counter >= len(row) {
				return true
			}
		}
	}

	return false
}

func readInput() gameSetup {
	rawInput := helpers.GetFileArray("./input")
	game := gameSetup{numbers: make([]int, 0), boards: make([][][]int, 0)}
	game.numbers = append(game.numbers, stringArrayToNum(strings.Split(rawInput[0], ","))...)

	currentBoard := make([][]int, 0)
	for index := 2; index < len(rawInput); index++ {
		line := rawInput[index]
		if (index >= len(rawInput)-1) || (len(line) == 0 && len(currentBoard) != 0) {
			game.boards = append(game.boards, currentBoard)
			currentBoard = make([][]int, 0)
		} else {
			currentBoard = append(currentBoard, stringArrayToNum(strings.Split(strings.Replace(line, "  ", " ", 1), " ")))
		}
	}
	return game
}

func stringArrayToNum(arr []string) []int {
	newArr := make([]int, 0)
	for _, s := range arr {
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		newArr = append(newArr, i)
	}
	return newArr
}
