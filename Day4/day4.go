package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("pt1 points: %d\n", leggiInput())
}

func leggiInput() (totPunti int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		splitedCard := strings.Split(strings.Replace(scanner.Text(), ":", "|", 1), "|")
		winner := strings.ReplaceAll(strings.TrimSpace(splitedCard[2]), "  ", " ")
		possiblyWinner := strings.Split(strings.ReplaceAll(strings.TrimSpace(splitedCard[1]), "  ", " "), " ")
		winCards := make(map[string]bool)
		for _, val := range strings.Split(winner, " ") {
			winCards[val] = false
		}
		totPunti += matchPoints(possiblyWinner, winCards)
	}
	return totPunti
}

func matchPoints(possiblyWinner []string, winCards map[string]bool) (points int) {
	first := true
	for _, val := range possiblyWinner {
		if _, ok := winCards[val]; ok {
			if first {
				points += 1
				first = false
			} else {
				points *= 2
			}
		}
	}
	return points
}
