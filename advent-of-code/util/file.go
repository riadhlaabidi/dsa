package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInput(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := make([][]byte, 0)

	for {
		line, err := reader.ReadBytes('\n')
		n := len(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lines = append(lines, line[:n-1])
	}

	return lines
}

func GetInt(str string) int {
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse int from \"%s\"", str))
	}

	return int(n)
}
