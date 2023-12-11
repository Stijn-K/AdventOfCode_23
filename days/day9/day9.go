package day9

import (
	"AdventOfCode/common"
	"log"
	"strconv"
	"strings"
)

func parseHistories(lines []string) [][]int64 {
	var histories [][]int64
	for _, line := range lines {
		raw := strings.Fields(line)
		history := make([]int64, len(raw))
		for i, s := range raw {
			history[i], _ = strconv.ParseInt(s, 10, 64)
		}
		histories = append(histories, history)
	}
	return histories
}

func allZeroes(sequence []int64) bool {
	for _, value := range sequence {
		if value != 0 {
			return false
		}
	}
	return true
}

func predictNextValue(history []int64) int64 {
	var sequences [][]int64
	sequences = append(sequences, history)

	curSequence := history
	for !allZeroes(curSequence) {
		differences := make([]int64, len(curSequence)-1)
		for i := 0; i < len(curSequence)-1; i++ {
			differences[i] = curSequence[i+1] - curSequence[i]
		}
		sequences = append(sequences, differences)
		curSequence = differences
	}
	nextValue := int64(0)
	for _, sequence := range sequences {
		nextValue += sequence[len(sequence)-1]
	}
	return nextValue
}

func Part1() {
	lines := common.ReadInputFile("./days/day9/part1.in")
	histories := parseHistories(lines)

	nextValues := make([]int64, len(histories))
	for i, history := range histories {
		nextValues[i] = predictNextValue(history)
	}

	sumNextValues := int64(0)
	for _, value := range nextValues {
		sumNextValues += value
	}

	log.Printf("Final answer: %d", sumNextValues)
}

func predictPreviousValue(history []int64) int64 {
	var sequences [][]int64
	sequences = append(sequences, history)

	curSequence := history
	for !allZeroes(curSequence) {
		differences := make([]int64, len(curSequence)-1)
		for i := 0; i < len(curSequence)-1; i++ {
			differences[i] = curSequence[i+1] - curSequence[i]
		}
		sequences = append(sequences, differences)
		curSequence = differences
	}
	previousValue := int64(0)
	for i := len(sequences) - 1; i >= 0; i-- {
		previousValue = sequences[i][0] - previousValue
	}
	return previousValue
}

func Part2() {
	lines := common.ReadInputFile("./days/day9/part1.in")
	histories := parseHistories(lines)

	previousValues := make([]int64, len(histories))
	for i, history := range histories {
		previousValues[i] = predictPreviousValue(history)
	}

	sumPreviousValues := int64(0)
	for _, value := range previousValues {
		sumPreviousValues += value
	}
	log.Printf("Final answer: %d", sumPreviousValues)
}
