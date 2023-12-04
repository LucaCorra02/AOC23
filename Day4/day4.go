package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	leggiInput()
}

func leggiInput() {
	scanner := bufio.NewScanner(os.Stdin)
	//re := regexp.MustCompile(`(\|)|(\d*)`)
	var tot int

	for scanner.Scan() {
		splitedCard := strings.Split(strings.Replace(scanner.Text(), ":", "|", 1), "|")
		winner := strings.ReplaceAll(strings.TrimSpace(splitedCard[2]), "  ", " ")
		possiblyWinner := strings.Split(strings.ReplaceAll(strings.TrimSpace(splitedCard[1]), "  ", " "), " ")

		winCards := make(map[string]bool)
		for _, val := range strings.Split(winner, " ") {
			winCards[val] = false
		}
		points := 0
		first := true
		for _, val := range possiblyWinner {
			if _, ok := winCards[val]; ok {
				//fmt.Println(val)
				if first {
					points += 1
					first = false
				} else {
					points *= 2
				}
			}
		}
		fmt.Println(points)
		tot += points
	}
	fmt.Println(tot)
}
