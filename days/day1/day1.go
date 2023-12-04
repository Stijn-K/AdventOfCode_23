package day1

import (
	"AdventOfCode/common"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Part1() {
	data := common.ReadInputFile("./days/day1/part1.in")
	var sumOfCalibrationValues int64 = 0
	for _, line := range data {
		var first rune = 0
		var last rune = 0
		for _, c := range line {
			if unicode.IsDigit(c) {
				if first == 0 {
					first = rune(c)
				}
				last = rune(c)
			}
		}
		log.Printf("first: %c, last: %c", first, last)
		calibrationValue, _ := strconv.ParseInt(fmt.Sprintf("%c%c", first, last), 10, 32)
		log.Printf("calibration value for line: %s = %d", line, calibrationValue)
		sumOfCalibrationValues += calibrationValue
	}
	log.Printf("Final answer: %d", sumOfCalibrationValues)
}

func getDigit(substr string) (rune, error) {
	if unicode.IsDigit(rune(substr[0])) {
		return rune(substr[0]), nil
	}
	prefixes := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for spelled, r := range prefixes {
		if strings.HasPrefix(substr, spelled) {
			return r, nil
		}
	}
	return 0, errors.New("not a digit")
}

func Part2() {
	data := common.ReadInputFile("./days/day1/part1.in")
	var sumOfCalibrationValues int64 = 0
	for _, line := range data {
		var first rune = 0
		var last rune = 0
		for idx, _ := range line {
			r, err := getDigit(line[idx:])
			if err != nil {
				continue
			}
			if first == 0 {
				first = r
			}
			last = r
		}
		log.Printf("first: %c, last: %c", first, last)
		calibrationValue, _ := strconv.ParseInt(fmt.Sprintf("%c%c", first, last), 10, 32)
		log.Printf("calibration value for line: %s = %d", line, calibrationValue)
		sumOfCalibrationValues += calibrationValue
	}
	log.Printf("Final answer: %d", sumOfCalibrationValues)
}
