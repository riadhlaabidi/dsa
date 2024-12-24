package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Gate struct {
	inWire1   string
	inWire2   string
	out       string
	operation string
}

func solvePartI(lines [][]byte) {
	wires := make(map[string]int)
	q := make([]Gate, 0)

	i := 0
	for ; len(lines[i]) > 0; i++ {
		l := strings.Split(string(lines[i]), ": ")
		wires[l[0]] = util.GetInt(l[1])
	}

	zs := 0
	for _, line := range lines[i+1:] {
		l := strings.Split(string(line), " -> ")
		ins := strings.Split(l[0], " ")
		out := l[1]
		if strings.HasPrefix(out, "z") {
			zs++
		}
		gate := Gate{ins[0], ins[2], out, ins[1]}
		q = append(q, gate)
	}

	res := make([]byte, zs)
	for len(q) > 0 {
		gate := q[0]
		q = q[1:]

		in1, exists1 := wires[gate.inWire1]
		in2, exists2 := wires[gate.inWire2]

		if exists1 && exists2 {
			switch gate.operation {
			case "XOR":
				wires[gate.out] = in1 ^ in2
			case "AND":
				wires[gate.out] = in1 & in2
			case "OR":
				wires[gate.out] = in1 | in2
			default:
				panic("Invalid operation")
			}

			if strings.HasPrefix(gate.out, "z") {
				idx := util.GetInt(gate.out[1:])
				res[zs-idx-1] = byte(wires[gate.out] + '0')
			}
			continue
		}

		q = append(q, gate)
	}

	ans, err := strconv.ParseInt(string(res), 2, 0)

	if err != nil {
		panic("Failed to convert number")
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day24.test")
	inputLines := util.ReadInput("day24.input")
	fmt.Printf("---- Day 24 Test ----\n")
	solvePartI(testLines) /* 2024 */
	solvePartII(testLines)
	fmt.Printf("---- Day 24 Input ----\n")
	solvePartI(inputLines) /* 42049478636360 */
	solvePartII(inputLines)
}
