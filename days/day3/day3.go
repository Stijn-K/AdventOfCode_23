package day3

import (
	"AdventOfCode/common"
	"log"
	"reflect"
	"strconv"
	"unicode"
)

type coord struct {
	x, y int
}

func isSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}

func findPartNumber(schematic []string, c coord) []coord {
	var partNumber []coord
	x, y := c.x, c.y
	for {
		if x >= len(schematic[y]) {
			break
		}
		if unicode.IsDigit(rune(schematic[y][x])) {
			partNumber = append(partNumber, coord{
				x: x,
				y: y,
			})
		} else {
			break
		}
		x += 1
	}
	return partNumber
}

func adjacentToSymbol(schematic []string, partNumber []coord) bool {
	delta := []int{-1, 0, 1}
	for _, pn := range partNumber {
		for _, dy := range delta {
			for _, dx := range delta {
				if pn.y+dy < 0 || pn.y+dy >= len(schematic) || pn.x+dx < 0 || pn.x+dx >= len(schematic[pn.y+dy]) {
					continue
				}
				if isSymbol(rune(schematic[pn.y+dy][pn.x+dx])) {
					return true
				}
			}
		}
	}
	return false
}

func partNumberToNumber(schematic []string, partNumber []coord) int {
	var numberRaw []rune
	for _, c := range partNumber {
		numberRaw = append(numberRaw, rune(schematic[c.y][c.x]))
	}
	number, _ := strconv.ParseInt(string(numberRaw), 10, 32)
	return int(number)
}

func Part1() {
	schematic := common.ReadInputFile("./days/day3/part1.in")

	sumOfPartNumbers := 0

	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			c := rune(schematic[y][x])
			if !unicode.IsDigit(c) {
				continue
			}
			partNumber := findPartNumber(schematic, coord{x: x, y: y})
			if adjacentToSymbol(schematic, partNumber) {
				sumOfPartNumbers += partNumberToNumber(schematic, partNumber)
			}
			x = partNumber[len(partNumber)-1].x
		}
	}
	log.Printf("Final answer: %d", sumOfPartNumbers)
}

type gearPart struct {
	partNumber []coord
	gearSymbol coord
}

func adjacentToGearSymbol(schematic []string, partNumber []coord) (bool, coord) {
	delta := []int{-1, 0, 1}
	for _, pn := range partNumber {
		for _, dy := range delta {
			for _, dx := range delta {
				if pn.y+dy < 0 || pn.y+dy >= len(schematic) || pn.x+dx < 0 || pn.x+dx >= len(schematic[pn.y+dy]) {
					continue
				}
				if rune(schematic[pn.y+dy][pn.x+dx]) == '*' {
					return true, coord{
						x: pn.y + dy,
						y: pn.x + dx,
					}
				}
			}
		}
	}
	return false, coord{}
}

func Part2() {
	schematic := common.ReadInputFile("./days/day3/part1.in")
	sumOfGearRatios := 0
	var gearParts []gearPart

	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			c := rune(schematic[y][x])
			if !unicode.IsDigit(c) {
				continue
			}
			partNumber := findPartNumber(schematic, coord{x: x, y: y})
			if ok, gearSymbol := adjacentToGearSymbol(schematic, partNumber); ok == true {
				gearParts = append(gearParts, gearPart{
					partNumber: partNumber,
					gearSymbol: gearSymbol,
				})
			}
			x = partNumber[len(partNumber)-1].x
		}
	}

	for i := 0; i < len(gearParts); i++ {
		for j := i + 1; j < len(gearParts); j++ {
			this := gearParts[i]
			other := gearParts[j]
			if reflect.DeepEqual(this.gearSymbol, other.gearSymbol) {
				sumOfGearRatios += partNumberToNumber(schematic, this.partNumber) * partNumberToNumber(schematic, other.partNumber)
			}
		}
	}
	log.Printf("Final answer: %d", sumOfGearRatios)
}
