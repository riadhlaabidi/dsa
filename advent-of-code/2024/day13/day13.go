package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Machine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func getMachines(lines [][]byte) []Machine {
	machines := make([]Machine, 0)

	for i := 0; i < len(lines)-2; i += 4 {
		var machine Machine

		buttonA := strings.Split(string(lines[i]), ": ")[1]
		axy := strings.Split(buttonA, ", ")
		machine.ax = util.GetInt(strings.Split(axy[0], "+")[1])
		machine.ay = util.GetInt(strings.Split(axy[1], "+")[1])

		buttonB := strings.Split(string(lines[i+1]), ": ")[1]
		bxy := strings.Split(buttonB, ", ")
		machine.bx = util.GetInt(strings.Split(bxy[0], "+")[1])
		machine.by = util.GetInt(strings.Split(bxy[1], "+")[1])

		prize := strings.Split(string(lines[i+2]), ": ")[1]
		pxy := strings.Split(prize, ", ")
		machine.px = util.GetInt(strings.Split(pxy[0], "=")[1])
		machine.py = util.GetInt(strings.Split(pxy[1], "=")[1])

		machines = append(machines, machine)
	}

	return machines
}

func solveSystem(first, second [3]int) (int, int) {
	coeff := first[1]*second[0] - second[1]*first[0]
	res := first[2]*second[0] - second[2]*first[0]

	y := res / coeff
	x := (first[2] - y*first[1]) / first[0]

	return x, y
}

func solvePartI(lines [][]byte) {
	ans := 0
	machines := getMachines(lines)

	for _, machine := range machines {
		eq1 := [3]int{machine.ax, machine.bx, machine.px}
		eq2 := [3]int{machine.ay, machine.by, machine.py}

		x, y := solveSystem(eq1, eq2)

		px := (machine.ax * x) + (machine.bx * y)
		py := (machine.ay * x) + (machine.by * y)

		if px == machine.px && py == machine.py {
			ans += x*3 + y
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0

	machines := getMachines(lines)
	diff := 10000000000000

	for _, machine := range machines {
		machine.px += diff
		machine.py += diff

		eq1 := [3]int{machine.ax, machine.bx, machine.px}
		eq2 := [3]int{machine.ay, machine.by, machine.py}

		x, y := solveSystem(eq1, eq2)

		px := (machine.ax * x) + (machine.bx * y)
		py := (machine.ay * x) + (machine.by * y)

		if px == machine.px && py == machine.py {
			ans += x*3 + y
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day13.test")
	inputLines := util.ReadInput("day13.input")
	fmt.Printf("---- Day 13 Test ----\n")
	solvePartI(testLines)  /* 480 */
	solvePartII(testLines) /* 875318608908 */
	fmt.Printf("---- Day 13 Input ----\n")
	solvePartI(inputLines)  /* 27157 */
	solvePartII(inputLines) /* 104015411578548 */
}
