package helpers

import (
	"bufio"
	"io"
	"os"
)

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
