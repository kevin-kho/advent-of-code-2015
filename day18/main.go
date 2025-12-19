package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
)

type Dir struct {
	X int
	Y int
}

func getGrid(data []byte) [][]byte {
	return bytes.Split(data, []byte{10})
}

func checkNeighbors(grid [][]byte, x, y int) int {
	var count int
	X := len(grid[0])
	Y := len(grid)
	dirs := []Dir{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: -1},
	}

	for _, d := range dirs {
		newX := x + d.X
		newY := y + d.Y

		// case: out of bounds
		if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
			continue
		}

		if grid[newY][newX] == 46 {
			count++
		}

	}
	return count
}

func getNextState(grid [][]byte, x, y int) byte {
	onNeighbors := checkNeighbors(grid, x, y)

	// case: on
	if grid[y][x] == 35 {
		if onNeighbors == 2 || onNeighbors == 3 {
			return 35 // stay on
		}
		return 46 // turn off
	}

	// case: off
	if onNeighbors == 3 {
		return 35 // turn on
	}
	return 46

}

func countLightsOn(grid [][]byte) int {
	var count int
	for _, row := range grid {
		for _, lt := range row {
			if lt == 35 {
				count++
			}
		}
	}
	return count
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := getGrid(data)
	fmt.Println(len(grid[0]))

}
