package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	leggiFile()
}

func leggiFile() {
	scanner := bufio.NewScanner(os.Stdin)
	var pt1, pt2, game int
	for scanner.Scan() {
		match := regexp.MustCompile(`\d* r|\d* g|\d* b`).FindAllString(scanner.Text(), -1)
		dict := map[byte]int{'r': 0, 'b': 0, 'g': 0}
		createRoundMap(dict, match)

		if dict['b'] <= 14 && dict['r'] <= 12 && dict['g'] <= 13 {
			pt1 += game + 1
		}
		pt2 += dict['r'] * dict['g'] * dict['b']
		game++
	}
	fmt.Printf("Parte 1 : %d, Parte 2 : %d\n", pt1, pt2)
}

func createRoundMap(dict map[byte]int, match []string) {
	for _, round := range match {
		color := len(round) - 1
		n, _ := strconv.Atoi(round[:color-1])
		dict[round[color]] = max(dict[round[color]], n)
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
