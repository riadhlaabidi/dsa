package main

import (
	"aoc/util"
	"fmt"
)

type Cell struct {
	x int
	y int
}

var offset = map[byte][]int{
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
	'^': {-1, 0},
}

func getInput(lines [][]byte, double bool) ([][]byte, []byte) {
	warehouse := make([][]byte, 0)
	directions := make([]byte, 0)

	sep := 0
	for i, line := range lines {
		if len(line) == 0 {
			sep = i
			break
		}

		l := make([]byte, 0)
		if double {
			for _, ch := range line {
				if ch == '#' {
					l = append(l, '#', '#')
				} else if ch == 'O' {
					l = append(l, '[', ']')
				} else if ch == '.' {
					l = append(l, '.', '.')
				} else if ch == '@' {
					l = append(l, '@', '.')
				}
			}
		} else {
			for _, ch := range line {
				l = append(l, ch)
			}
		}
		warehouse = append(warehouse, l)
	}

	for _, line := range lines[sep+1:] {
		directions = append(directions, line...)
	}

	return warehouse, directions
}

func getRobotInitialPos(warehouse [][]byte) (int, int) {
	x := -1
	y := -1

	for i, line := range warehouse {
		for j, ch := range line {
			if ch == '@' {
				x, y = i, j
			}
		}
	}

	if x < 0 || y < 0 {
		panic("Robot not found")
	}

	return x, y
}

func getCoordinatesSum(warehouse [][]byte, char byte) int {
	ans := 0

	for i, line := range warehouse {
		for j, ch := range line {
			if ch == char {
				ans += 100*i + j
			}
		}
	}

	return ans
}

func moveBoxes(warehouse [][]byte, directions []byte, rx, ry int) {
	for _, dir := range directions {
		offX := offset[dir][0]
		offY := offset[dir][1]
		visited := make(map[Cell]bool)
		boxes := make([]Cell, 0)
		q := []Cell{{rx, ry}}
		canMove := true

		for len(q) != 0 {
			c := q[0]
			boxes = append(boxes, c)
			q = q[1:]
			nextX := c.x + offX
			nextY := c.y + offY

			if warehouse[nextX][nextY] == '#' {
				canMove = false
				break
			}

			if warehouse[nextX][nextY] == '.' {
				continue
			}

			if !visited[Cell{nextX, nextY}] {
				visited[Cell{nextX, nextY}] = true
				q = append(q, Cell{nextX, nextY})

				if warehouse[nextX][nextY] == '[' {
					visited[Cell{nextX, nextY + 1}] = true
					q = append(q, Cell{nextX, nextY + 1})
				} else if warehouse[nextX][nextY] == ']' {
					visited[Cell{nextX, nextY - 1}] = true
					q = append(q, Cell{nextX, nextY - 1})
				}
			}
		}

		if canMove {
			// move all boxes in reverse order except the robot @ index 0
			for i := len(boxes) - 1; i > 0; i-- {
				oldX := boxes[i].x
				oldY := boxes[i].y
				newX := oldX + offX
				newY := oldY + offY
				warehouse[newX][newY], warehouse[oldX][oldY] = warehouse[oldX][oldY], '.'
			}

			// move robot
			warehouse[rx][ry] = '.'
			rx += offX
			ry += offY
			warehouse[rx][ry] = '@'
		}
	}
}

func solvePartI(lines [][]byte) {
	warehouse, directions := getInput(lines, false)
	rx, ry := getRobotInitialPos(warehouse)

	moveBoxes(warehouse, directions, rx, ry)

	ans := getCoordinatesSum(warehouse, 'O')
	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	warehouse, directions := getInput(lines, true)
	rx, ry := getRobotInitialPos(warehouse)

	moveBoxes(warehouse, directions, rx, ry)

	ans := getCoordinatesSum(warehouse, '[')
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day15.test")
	inputLines := util.ReadInput("day15.input")
	fmt.Printf("---- Day 15 Test ----\n")
	solvePartI(testLines)  /* 10092 */
	solvePartII(testLines) /* 9021 */
	fmt.Printf("---- Day 15 Input ----\n")
	solvePartI(inputLines)  /* 1371036 */
	solvePartII(inputLines) /* 1392847 */
}
