package day2

import (
	"AdventOfCode/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func parseGame(line string) game {
	var parsedGame game
	gameRe := regexp.MustCompile("Game (\\d+): (.*)")
	gameMatches := gameRe.FindStringSubmatch(line)
	parsedGame.id, _ = strconv.Atoi(gameMatches[1])

	cubesRe := regexp.MustCompile("(\\d+) (blue|red|green)")
	for _, setRaw := range strings.Split(gameMatches[2], ";") {
		var s set
		for _, cubeRaw := range strings.Split(setRaw, ",") {
			cubeMatches := cubesRe.FindStringSubmatch(cubeRaw)
			switch cubeMatches[2] {
			case "red":
				s.red, _ = strconv.Atoi(cubeMatches[1])
			case "green":
				s.green, _ = strconv.Atoi(cubeMatches[1])
			case "blue":
				s.blue, _ = strconv.Atoi(cubeMatches[1])
			default:
				panic(cubeMatches)
			}
		}
		parsedGame.sets = append(parsedGame.sets, s)
	}
	return parsedGame
}

func Part1() {
	lines := common.ReadInputFile("./days/day2/part1.in")

	var sumOfIds = 0
	for _, line := range lines {
		game := parseGame(line)
		possible := true
		for _, s := range game.sets {
			if s.red > 12 || s.green > 13 || s.blue > 14 {
				possible = false
			}
		}
		if possible {
			sumOfIds += game.id
		}
	}
	log.Printf("Final answer: %d", sumOfIds)

}

func Part2() {
	lines := common.ReadInputFile("./days/day2/part1.in")
	sumOfPowers := 0
	for _, line := range lines {
		var maxRed, maxGreen, maxBlue int
		game := parseGame(line)
		for _, s := range game.sets {
			maxRed = max(s.red, maxRed)
			maxGreen = max(s.green, maxGreen)
			maxBlue = max(s.blue, maxBlue)
		}
		sumOfPowers += maxRed * maxGreen * maxBlue
	}
	log.Printf("Final answer: %d", sumOfPowers)
}
