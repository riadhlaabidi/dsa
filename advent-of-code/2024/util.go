package main

import (
	"bufio"
	"io"
	"os"
)

func ReadInput(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

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

	file.Close()
	return lines
}
