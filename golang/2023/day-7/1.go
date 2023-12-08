package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

type hand struct {
	hand   string
	bid    int
	cards  [5]byte
	points byte
}

const HIGHCARD byte = 1
const ONEPAIR byte = 2
const TWOPAIR byte = 3
const THREEOFAKIND byte = 4
const FULLHOUSE byte = 5
const FOUROFAKIND byte = 6
const FIVEOFAFIND byte = 7

var cardToPoint map[byte]byte

func init() {
	cardToPoint = map[byte]byte{
		byte('2'): 2,
		byte('3'): 3,
		byte('4'): 4,
		byte('5'): 5,
		byte('6'): 6,
		byte('7'): 7,
		byte('8'): 8,
		byte('9'): 9,
		byte('T'): 10,
		byte('J'): 12,
		byte('Q'): 13,
		byte('K'): 14,
		byte('A'): 15,
	}
}

func getCardNumbers(cards string) [5]byte {
	c1 := cards[0]
	c2 := cards[1]
	c3 := cards[2]
	c4 := cards[3]
	c5 := cards[4]

	return [5]byte{
		cardToPoint[c1],
		cardToPoint[c2],
		cardToPoint[c3],
		cardToPoint[c4],
		cardToPoint[c5],
	}
}

func cardNumbersToMap(cards [5]byte) map[byte]byte {
	m := map[byte]byte{}

	for _, card := range cards {
		m[card] = m[card] + 1
	}

	return m
}

func getPoints(cards [5]byte) byte {
	m := cardNumbersToMap(cards)

	if len(m) == 1 {
		return FIVEOFAFIND
	}

	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return FOUROFAKIND
			}
		}
		return FULLHOUSE
	}

	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return THREEOFAKIND
			}
		}
		return TWOPAIR
	}

	if len(m) == 4 {
		return ONEPAIR
	}
	if len(m) == 5 {
		return HIGHCARD
	}

	panic("Length of map:" + string(len(m)))
}

func NewHand(cards string, bid int) hand {
	cardsNums := getCardNumbers(cards)

	points := getPoints(cardsNums)

	return hand{
		hand:   cards,
		bid:    bid,
		cards:  getCardNumbers(cards),
		points: points,
	}
}

type listOfHands []hand

func (l listOfHands) Len() int {
	return len(l)
}

func (l listOfHands) Less(idx1, idx2 int) bool {
	h1 := l[idx1]
	h2 := l[idx2]

	if h1.points != h2.points {
		return h1.points < h2.points
	}

	for idx, card := range h1.cards {
		if card != h2.cards[idx] {

			return card < h2.cards[idx]
		}
	}

	return false
}

func (l listOfHands) Swap(idx1, idx2 int) {
	l[idx1], l[idx2] = l[idx2], l[idx1]

}

func main() {

	hands := listOfHands{}

	for _, line := range strings.Split(test, "\n") {

		parts := strings.Split(line, " ")

		bid, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		hands = append(hands, NewHand(parts[0], bid))

	}
	sort.Sort(hands)
	// slices.SortFunc(hands, compare)
	totalWinnings := 0
	for idx, hand := range hands {
		totalWinnings += (idx + 1) * hand.bid
		fmt.Println(idx+1, " - ", hand.hand, "(", hand.points, ")")
	}

	fmt.Println("Total winings:", totalWinnings)
}
