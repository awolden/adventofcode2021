package helpers

import (
	"bufio"
	"io"
	"os"
	"reflect"
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

func ConvertStrArr(nums []string) []int {

	var arr = make([]int, 0)
	for _, numAsString := range nums {
		num, _ := strconv.Atoi(numAsString)
		arr = append(arr, num)
	}
	return arr
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

func ReverseString(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
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

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func FindArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func FindArrayProduct(arr []int) int {
	res := 1
	for i := 0; i < len(arr); i++ {
		res *= arr[i]
	}
	return res
}
