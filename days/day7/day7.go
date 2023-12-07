package day7

import (
	"AdventOfCode/common"
	"log"
	"slices"
	"strconv"
	"strings"
)

var strengthMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func (h hand) isFiveOfAKind() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	for _, count := range counts {
		if count == 5 {
			return true
		}
	}

	return false
}

func (h hand) isFourOfAKind() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	for _, count := range counts {
		if count == 4 {
			return true
		}
	}

	return false
}

func (h hand) isFullHouse() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	threeOfAKind := false
	pair := false
	for _, count := range counts {
		if count == 3 {
			threeOfAKind = true
		}
		if count == 2 {
			pair = true
		}
	}

	return threeOfAKind && pair
}

func (h hand) isThreeOfAKind() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	for _, count := range counts {
		if count == 3 {
			return true
		}
	}

	return false
}

func (h hand) isTwoPair() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	pairCount := 0
	for _, count := range counts {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func (h hand) isOnePair() bool {
	counts := make(map[rune]int)
	for _, card := range h.cards {
		counts[card]++
	}

	for _, count := range counts {
		if count == 2 {
			return true
		}
	}

	return false
}

type hand struct {
	cards []rune
	bid   int
}

// types:
// Five of a kind: 	7
// Four of a kind: 	6
// Full House: 		5
// Three of a kind: 4
// Two pair: 		3
// One pair: 		2
// High card: 		1

func (h hand) getType() int {
	if h.isFiveOfAKind() {
		return 7
	} else if h.isFourOfAKind() {
		return 6
	} else if h.isFullHouse() {
		return 5
	} else if h.isThreeOfAKind() {
		return 4
	} else if h.isTwoPair() {
		return 3
	} else if h.isOnePair() {
		return 2
	}
	return 1 // High card if no other hand type is detected
}

func compareHands(this, other hand) int {
	thisType := this.getType()
	otherType := other.getType()

	if thisType > otherType {
		return 1
	} else if thisType < otherType {
		return -1
	}

	for i := 0; i < len(this.cards); i++ {
		if strengthMap[this.cards[i]] > strengthMap[other.cards[i]] {
			return 1
		} else if strengthMap[this.cards[i]] < strengthMap[other.cards[i]] {
			return -1
		}
	}

	return 0
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")

	cards := []rune(parts[0])

	bid := 0
	bid, _ = strconv.Atoi(parts[1])

	return hand{
		cards: cards,
		bid:   bid,
	}
}

func Part1() {
	handsRaw := common.ReadInputFile("./days/day7/part1.in")
	var hands []hand
	for _, handRaw := range handsRaw {
		hands = append(hands, parseHand(handRaw))
	}

	slices.SortFunc(hands, compareHands)

	totalWinnings := 0
	for rank, hand := range hands {
		totalWinnings += (rank + 1) * hand.bid
	}
	log.Printf("Final answer: %d", totalWinnings)
}

func emptyCount() map[rune]int {
	counts := make(map[rune]int)
	for r, _ := range strengthMap {
		counts[r] = 0
	}
	return counts
}

func (h hand) isFiveOfAKindP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}

	for card, count := range counts {
		// added condition to make J act as a wildcard while checking hand type
		if count+counts['J'] == 5 && card != 'J' {
			return true
		}
	}

	return false
}

func (h hand) isFourOfAKindP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}

	for card, count := range counts {
		if count+counts['J'] == 4 && card != 'J' {
			return true
		}
	}
	return false
}

func (h hand) isFullHouseP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}
	pair := 0
	threeOfAKind := 0
	for card, count := range counts {
		if card == 'J' {
			continue
		}
		if count+counts['J'] >= 3 {
			threeOfAKind = 1
			counts['J'] -= 3 - count
		} else if count+counts['J'] >= 2 && pair == 0 {
			pair = 1
			counts['J'] -= 2 - count
		}

	}
	return pair+threeOfAKind == 2

}

func (h hand) isThreeOfAKindP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}

	for card, count := range counts {
		if count+counts['J'] == 3 && card != 'J' {
			return true
		}
	}
	return false
}

func (h hand) isTwoPairP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}

	pairCount := 0
	for card, count := range counts {
		if count+counts['J'] == 2 && card != 'J' {
			pairCount++
			// deduct from joker count only if it acted as wildcard
			counts['J'] = 2 - counts['J'] - count
		}
	}
	// Should have 2 pairs
	return pairCount == 2
}

func (h hand) isOnePairP2() bool {
	counts := emptyCount()
	for _, card := range h.cards {
		counts[card]++
	}

	for card, count := range counts {
		if count+counts['J'] == 2 && card != 'J' {
			return true
		}
	}
	return false
}

func compareHandsP2(this, other hand) int {
	thisType := this.getTypeP2()
	otherType := other.getTypeP2()

	if thisType > otherType {
		return 1
	} else if thisType < otherType {
		return -1
	}

	for i := 0; i < 5; i++ {
		thisCard := this.cards[i]
		otherCard := other.cards[i]

		if thisCard == 'J' && otherCard != 'J' {
			return -1
		} else if thisCard != 'J' && otherCard == 'J' {
			return 1
		} else if strengthMap[thisCard] > strengthMap[otherCard] {
			return 1
		} else if strengthMap[thisCard] < strengthMap[otherCard] {
			return -1
		}
	}
	return 0
}

func (h hand) getTypeP2() int {
	if h.isFiveOfAKindP2() {
		return 7
	} else if h.isFourOfAKindP2() {
		return 6
	} else if h.isFullHouseP2() {
		return 5
	} else if h.isThreeOfAKindP2() {
		return 4
	} else if h.isTwoPairP2() {
		return 3
	} else if h.isOnePairP2() {
		return 2
	}
	return 1 // High card if no other hand type is detected
}

func Part2() {
	handsRaw := common.ReadInputFile("./days/day7/part1.in")
	var hands []hand
	for _, handRaw := range handsRaw {
		hands = append(hands, parseHand(handRaw))
	}

	slices.SortFunc(hands, compareHandsP2)

	var totalWinnings int64
	for rank, hand := range hands {
		totalWinnings += int64(rank+1) * int64(hand.bid)
	}
	log.Printf("Final answer: %d", totalWinnings)
}
