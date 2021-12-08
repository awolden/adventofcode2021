package helpers

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)

func GetFileArray(file string) []string {
	fileHandle, _ := os.Open(file)
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	var arr = make([]string, 0)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		arr = append(arr, string(line))
	}
	return arr
}

func ArrEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ConvertIntArr(nums []int) string {

	var arr = make([]string, 0)
	for _, num := range nums {
		arr = append(arr, strconv.Itoa(num))
	}
	return strings.Join(arr, "")
}

func Contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

func ContainsString(s []string, i string) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

func Min(arr []int) int {
	min := 1<<(UintSize-1) - 1
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func Max(arr []int) int {
	max := 0
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
