package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pt1, pt2 := leggiInput()
	fmt.Printf("pt1 points: %d\npt2 points: %d\n", pt1, pt2)
}

func leggiInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	copy := make(map[int]int)
	re := regexp.MustCompile(`\d+`)
	var totPunti int

	for scanner.Scan() {
		splitedCard := strings.Split(strings.Replace(scanner.Text(), ":", "|", 1), "|")
		winner := strings.ReplaceAll(strings.TrimSpace(splitedCard[2]), "  ", " ")
		possiblyWinner := strings.Split(strings.ReplaceAll(strings.TrimSpace(splitedCard[1]), "  ", " "), " ")
		winCards := make(map[string]bool)
		for _, val := range strings.Split(winner, " ") {
			winCards[val] = false
		}
		partialPoint, numMatch := matchPoints(possiblyWinner, winCards)
		cardNum, _ := strconv.Atoi(re.FindAllString(splitedCard[0], -1)[0])
		updateCopy(copy, cardNum, numMatch)
		totPunti += partialPoint
	}
	return totPunti, pointsPt2(copy)
}

func updateCopy(copy map[int]int, cardNum int, numMatch int) {
	copy[cardNum] += 1
	for i := cardNum + 1; i <= numMatch+cardNum; i++ {
		copy[i] += copy[cardNum]
	}
}

func pointsPt2(copy map[int]int) (pt int) {
	for _, value := range copy {
		pt += value
	}
	return
}

func matchPoints(possiblyWinner []string, winCards map[string]bool) (points int, numMatch int) {
	first := true
	for _, val := range possiblyWinner {
		if _, ok := winCards[val]; ok {
			points *= 2
			if first {
				points += 1
				first = false
			}
			numMatch++
		}
	}
	return points, numMatch
}
