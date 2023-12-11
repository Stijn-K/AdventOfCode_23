package day10

import (
	"AdventOfCode/common"
	"log"
)

type gridSquare struct {
	x, y int
	tile rune
}

func findStart(grid [][]rune) (x, y int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'S' {
				return x, y
			}
		}
	}
	panic("Start not found")
}

func findStartingDirections(grid [][]rune, x, y int) [][]int {
	var startingDirections [][]int
	if x-1 > 0 && (grid[y][x-1] == 'L' || grid[y][x-1] == 'F' || grid[y][x-1] == '-') {
		startingDirections = append(startingDirections, []int{-1, 0})
	}
	if x+1 < len(grid[y]) && (grid[y][x+1] == 'J' || grid[y][x+1] == '7' || grid[y][x+1] == '-') {
		startingDirections = append(startingDirections, []int{1, 0})
	}
	if y-1 > 0 && (grid[y-1][x] == '7' || grid[y-1][x] == 'F' || grid[y-1][x] == '|') {
		startingDirections = append(startingDirections, []int{0, -1})
	}
	if y+1 < len(grid) && (grid[y+1][x] == 'L' || grid[y+1][x] == 'J' || grid[y+1][x] == '|') {
		startingDirections = append(startingDirections, []int{0, 1})
	}
	return startingDirections
}

func findLoop(grid [][]rune) []gridSquare {
	x, y := findStart(grid)
	startingDirections := findStartingDirections(grid, x, y)
	dx, dy := startingDirections[0][0], startingDirections[0][1]
	var loop []gridSquare
	loop = append(loop, gridSquare{
		x:    x,
		y:    y,
		tile: 'S',
	})
	for {
		x, y = x+dx, y+dy
		if grid[y][x] == 'S' {
			break
		}
		tile := grid[y][x]
		if dx == 1 {
			if tile == '-' {
				dx, dy = 1, 0
			} else if tile == '7' {
				dx, dy = 0, 1
			} else if tile == 'J' {
				dx, dy = 0, -1
			}
		} else if dx == -1 {
			if tile == '-' {
				dx, dy = -1, 0
			} else if tile == 'F' {
				dx, dy = 0, 1
			} else if tile == 'L' {
				dx, dy = 0, -1
			}
		} else if dy == 1 {
			if tile == '|' {
				dx, dy = 0, 1
			} else if tile == 'L' {
				dx, dy = 1, 0
			} else if tile == 'J' {
				dx, dy = -1, 0
			}
		} else if dy == -1 {
			if tile == '|' {
				dx, dy = 0, -1
			} else if tile == '7' {
				dx, dy = -1, 0
			} else if tile == 'F' {
				dx, dy = 1, 0
			}
		}

		loop = append(loop, gridSquare{
			x:    x,
			y:    y,
			tile: tile,
		})

	}
	return loop
}

func Part1() {
	grid := common.ReadInputFileAsRunes("./days/day10/part1.in")
	loop := findLoop(grid)
	log.Printf("Final answer: %d", len(loop)/2)
}

func isPartOfLoop(loop []gridSquare, x, y int) bool {
	for _, square := range loop {
		if square.x == x && square.y == y {
			return true
		}
	}
	return false
}

func findStartReplacement(startingDirections [][]int) rune {
	if startingDirections[0][0] == 1 && startingDirections[1][0] == -1 {
		return '-'
	}
	if startingDirections[0][1] == -1 && startingDirections[1][1] == 1 {
		return '|'
	}
	if startingDirections[0][0] == 1 && startingDirections[1][1] == -1 {
		return 'L'
	}
	if startingDirections[0][0] == -1 && startingDirections[1][1] == -1 {
		return 'J'
	}
	if startingDirections[0][0] == -1 && startingDirections[1][1] == 1 {
		return '7'
	}
	if startingDirections[0][0] == 1 && startingDirections[1][1] == 1 {
		return 'F'
	}
	panic("No start replacement found")
}

func Part2() {
	grid := common.ReadInputFileAsRunes("./days/day10/part1.in")
	startX, startY := findStart(grid)
	startingDirections := findStartingDirections(grid, startX, startY)
	loop := findLoop(grid)
	grid[startY][startX] = findStartReplacement(startingDirections)
	insideTiles := 0
	for y := 0; y < len(grid); y++ {
		cnt := 0
		for x := 0; x < len(grid[y]); x++ {
			if isPartOfLoop(loop, x, y) {
				if grid[y][x] == 'F' || grid[y][x] == '7' || grid[y][x] == '|' {
					cnt++
				}
			} else if cnt%2 == 1 {
				insideTiles++
			}
		}
	}

	log.Printf("Final answer: %d", insideTiles)
}
