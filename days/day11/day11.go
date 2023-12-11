package day11

import (
	"AdventOfCode/common"
	"log"
	"slices"
)

type galaxy struct {
	x, y int
}

func findGalaxies(image [][]rune) (galaxies []galaxy) {
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[y]); x++ {
			if image[y][x] == '#' {
				galaxies = append(galaxies, galaxy{
					x: x,
					y: y,
				})
			}
		}
	}
	return
}

func findExpandedRowsAndColumns(image [][]rune) (rows, columns []int) {
	for i := 0; i < len(image); i++ {
		empty := true
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] != '.' {
				empty = false
				break
			}
		}
		if empty {
			rows = append(rows, i)
		}
	}
	for i := 0; i < len(image[0]); i++ {
		empty := true
		for j := 0; j < len(image); j++ {
			if image[j][i] != '.' {
				empty = false
				break
			}
		}
		if empty {
			columns = append(columns, i)
		}
	}
	return
}

func findShortestPath(start, end galaxy, expandRows, expandColumns []int, expandRatio int) (length int) {
	for i := min(start.x, end.x) + 1; i <= max(start.x, end.x); i++ {
		if slices.Contains(expandColumns, i) {
			length += expandRatio
		} else {
			length++
		}
	}
	for i := min(start.y, end.y); i < max(start.y, end.y); i++ {
		if slices.Contains(expandRows, i) {
			length += expandRatio
		} else {
			length++
		}
	}
	return
}

func Part1() {
	image := common.ReadInputFileAsRunes("./days/day11/part1.in")
	galaxies := findGalaxies(image)
	expandRows, expandColumns := findExpandedRowsAndColumns(image)

	expandRatio := 2
	var shortestPaths []int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			shortestPaths = append(shortestPaths, findShortestPath(galaxies[i], galaxies[j], expandRows, expandColumns, expandRatio))
		}
	}
	sumShortestPaths := 0
	for _, shortestPath := range shortestPaths {
		sumShortestPaths += shortestPath
	}

	log.Printf("Final answer: %d", sumShortestPaths)
}

func Part2() {
	image := common.ReadInputFileAsRunes("./days/day11/part1.in")
	galaxies := findGalaxies(image)
	expandRows, expandColumns := findExpandedRowsAndColumns(image)

	expandRatio := 1_000_000
	var shortestPaths []int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			shortestPaths = append(shortestPaths, findShortestPath(galaxies[i], galaxies[j], expandRows, expandColumns, expandRatio))
		}
	}
	sumShortestPaths := 0
	for _, shortestPath := range shortestPaths {
		sumShortestPaths += shortestPath
	}

	log.Printf("Final answer: %d", sumShortestPaths)
}
