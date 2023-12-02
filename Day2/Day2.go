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
	leggiFile()

}

func leggiFile() {
	scanner := bufio.NewScanner(os.Stdin)
	var pt1, pt2, game int
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ";")
		regex := regexp.MustCompile(`\d* r|\d* g|\d* b`)

		dict := map[byte]int{'r': 0, 'b': 0, 'g': 0}
		for _, str := range line {
			match := regex.FindAllString(str, -1)
			for _, play := range match { //singole giocate nel match
				color := len(play) - 1
				n, _ := strconv.Atoi(play[:color-1])
				dict[play[color]] = max(dict[play[color]], n)
			}
		}
		if dict['b'] <= 14 && dict['r'] <= 12 && dict['g'] <= 13 {
			pt1 += game + 1
		}
		pt2 += dict['r'] * dict['g'] * dict['b']
		game++
	}

	fmt.Printf("Parte 1 : %d, Parte 2 : %d\n", pt1, pt2)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
