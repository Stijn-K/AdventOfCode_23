package day8

import (
	"AdventOfCode/common"
	"log"
	"regexp"
)

func parseNetwork(networkRaw []string) map[string][2]string {
	network := make(map[string][2]string)
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	for _, line := range networkRaw {
		m := re.FindAllSubmatch([]byte(line), -1)
		network[string(m[0][1])] = [2]string{string(m[0][2]), string(m[0][3])}
	}
	return network
}

var actionMap = map[rune]int{
	'L': 0,
	'R': 1,
}

func Part1() {
	lines := common.ReadInputFile("./days/day8/part1.in")
	actions := lines[0]
	network := parseNetwork(lines[2:])
	curPos := "AAA"

	var steps int64 = 0
	for {
		action := steps % int64(len(actions))
		curPos = network[curPos][actionMap[rune(actions[action])]]
		steps += 1
		if curPos == "ZZZ" {
			break
		}
	}
	log.Printf("Final answer: %d", steps)
}

func Part2() {
	lines := common.ReadInputFile("./days/day8/part1.in")
	actions := lines[0]
	network := parseNetwork(lines[2:])

	var starts []string
	for node := range network {
		if node[2] == 'A' {
			starts = append(starts, node)
		}
	}
	allLengths := make([]int, len(starts))
	for idx, start := range starts {
		curPos := start
		steps := 0
		for {
			action := steps % len(actions)
			curPos = network[curPos][actionMap[rune(actions[action])]]
			steps += 1
			if curPos[2] == 'Z' {
				allLengths[idx] = steps
				break
			}
		}
	}
	log.Print(allLengths)
	log.Printf("Final answer: %d", common.LCM(1, 1, allLengths...))
}
