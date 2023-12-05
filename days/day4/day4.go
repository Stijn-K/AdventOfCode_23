package day4

import (
	"AdventOfCode/common"
	"log"
	"math"
	"strconv"
	"strings"
)

func isWinningNumber(winningNumbers []string, number string) bool {
	for _, winningNumber := range winningNumbers {
		if winningNumber == number {
			return true
		}
	}
	return false
}

func Part1() {
	cardsRaw := common.ReadInputFile("./days/day4/part1.in")

	sumCardValue := 0

	for _, cardRaw := range cardsRaw {
		numbersRaw := strings.Split(strings.Split(cardRaw, ":")[1], "|")
		winningNumbers := strings.Fields(numbersRaw[0])
		ownedNumbers := strings.Fields(numbersRaw[1])

		matches := 0
		for _, ownedNumber := range ownedNumbers {
			if isWinningNumber(winningNumbers, ownedNumber) {
				matches += 1
			}
		}
		sumCardValue += int(math.Pow(float64(2), float64(matches-1)))
	}
	log.Printf("Final answer: %d", sumCardValue)
}

func Part2() {
	cardsRaw := common.ReadInputFile("./days/day4/part1.in")

	cardWins := make([]int, len(cardsRaw))

	totalProcessedCards := 0
	for _, cardRaw := range cardsRaw {
		cardNumber, _ := strconv.ParseInt(strings.Fields(strings.Split(cardRaw, ":")[0])[1], 10, 32)
		numbersRaw := strings.Split(strings.Split(cardRaw, ":")[1], "|")
		winningNumbers := strings.Fields(numbersRaw[0])
		ownedNumbers := strings.Fields(numbersRaw[1])

		matches := 0
		for _, ownedNumber := range ownedNumbers {
			if isWinningNumber(winningNumbers, ownedNumber) {
				matches += 1
			}
		}
		cardWins[cardNumber-1] = matches
	}
	cards := make([]int, len(cardsRaw))
	for i := range cards {
		cards[i] = 1
	}
	for card := range cards {
		for newCard := 1; newCard <= cardWins[card]; newCard++ {
			cards[card+newCard] += cards[card]
		}
	}
	for _, card := range cards {
		totalProcessedCards += card
	}

	log.Printf("Final answer: %d", totalProcessedCards)
}
