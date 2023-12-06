package day6

import (
	"AdventOfCode/common"
	"log"
	"strings"
)

type race struct {
	time, record int
}

func Part1() {
	lines := common.ReadInputFile("./days/day6/part1.in")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	records := strings.Fields(strings.Split(lines[1], ":")[1])

	var races []race
	for i := 0; i < len(times); i++ {
		races = append(races, race{time: common.StringToInt(times[i]), record: common.StringToInt(records[i])})
	}

	prodWins := 1
	for _, r := range races {
		wins := 0
		for t := 0; t < r.time; t++ {
			distance := t * (r.time - t)
			if distance > r.record {
				wins += 1
			}
		}
		prodWins *= wins
	}
	log.Printf("Final answer: %d", prodWins)
}

type race64 struct {
	time, record int64
}

func Part2() {
	lines := common.ReadInputFile("./days/day6/part1.in")
	timeRaw := strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), "")
	recordRaw := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")

	r := race64{
		time:   common.StringToInt64(timeRaw),
		record: common.StringToInt64(recordRaw),
	}
	wins := 0
	for t := int64(0); t < r.time; t++ {
		distance := t * (r.time - t)
		if distance > r.record {
			wins += 1
		}
	}
	log.Printf("Final answer: %d", wins)
}
