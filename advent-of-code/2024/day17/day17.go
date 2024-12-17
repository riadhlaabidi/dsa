package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

var registerA int
var registerB int
var registerC int

func getRegisters(lines [][]byte) (int, int, int) {
	rA := util.GetInt(strings.Split(string(lines[0]), ": ")[1])
	rB := util.GetInt(strings.Split(string(lines[1]), ": ")[1])
	rC := util.GetInt(strings.Split(string(lines[2]), ": ")[1])

	return rA, rB, rC
}

func getProgram(lines [][]byte) []int {
	program := make([]int, 0)
	p := strings.Split(string(lines[4]), ": ")[1]
	pp := strings.Split(p, ",")

	for _, op := range pp {
		program = append(program, util.GetInt(op))
	}

	return program
}

func getComboOperand(operand int) int {
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return operand
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	default:
		panic(fmt.Sprintf("Invalid combo operand %d", operand))
	}
}

func solvePartI(lines [][]byte) {
	registerA, registerB, registerC = getRegisters(lines)
	program := getProgram(lines)
	output := make([]string, 0)

	for i := 0; i < len(program); {
		operand := program[i+1]
		switch program[i] {
		case 0:
			registerA /= 1 << getComboOperand(operand)
		case 1:
			registerB ^= operand
		case 2:
			registerB = (getComboOperand(operand) % 8) & 7
		case 3:
			if registerA != 0 {
				i = operand
				continue
			}
		case 4:
			registerB ^= registerC
		case 5:
			output = append(output, fmt.Sprintf("%d", getComboOperand(operand)%8))
		case 6:
			registerB = registerA / (1 << getComboOperand(operand))
		case 7:
			registerC = registerA / (1 << getComboOperand(operand))
		default:
			panic(fmt.Sprintf("Invalid operation with opcode %d", program[i]))
		}

		i += 2
	}

	fmt.Printf("Part 1: %v\n", strings.Join(output, ","))
}

func solvePartII(lines [][]byte) {
	ans := 0
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines1 := util.ReadInput("day17-1.test")
	testLines2 := util.ReadInput("day17-2.test")
	inputLines := util.ReadInput("day17.input")
	fmt.Printf("---- Day 17 Test ----\n")
	solvePartI(testLines1) /* 4,6,3,5,6,3,5,2,1,0 */
	solvePartII(testLines2)
	fmt.Printf("---- Day 17 Input ----\n")
	solvePartI(inputLines) /* 2,7,6,5,6,0,2,3,1 */
	solvePartII(inputLines)
}
