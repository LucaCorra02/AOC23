package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand   string
	points int
	factor int
}

var powers = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '5': 4, '4': 3, '3': 2, '2': 1}

func main() {
	match := leggiFile()
	fmt.Printf("Punti := %d\n", finalScore(match))
}

func leggiFile() (match []Hand) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		hand := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(hand[1])
		match = append(match, Hand{hand[0], getHandPoints(hand[0]), n})
	}
	return
}

func finalScore(match []Hand) (points int) {
	sort.SliceStable(match, func(i, j int) bool {
		c1 := match[i].points
		c2 := match[j].points
		return c1 < c2
	})

	for i := 0; i < len(match)-1; i++ {
		if match[i].points == match[i+1].points && getStrongestHand(match[i], match[i+1]) {
			match[i], match[i+1] = match[i+1], match[i]
		}
	}
	for i, hand := range match {
		points += (hand.factor) * (i + 1)
	}
	return
}

func getStrongestHand(hand1 Hand, hand2 Hand) bool {
	cards1 := hand1.hand
	cards2 := hand2.hand
	for i := 0; i < len(cards1); i++ {
		if powers[rune(cards1[i])] > powers[rune(cards2[i])] {
			return true
		}
	}
	return false
}

func getHandPoints(hand string) (points int) {
	occurency := make(map[rune]int)
	for _, card := range hand {
		occurency[card] += 1
	}
	lun := len(occurency)
	switch lun {
	case 1:
		return 7
	case 2:
		for _, value := range occurency {
			if value == 4 {
				return 6
			}
		}
		return 5
	case 3:
		for _, value := range occurency {
			if value == 3 {
				return 4
			}
		}
		return 3
	case 4:
		return 2
	default:
		return 1
	}
}
