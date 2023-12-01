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
	fmt.Printf("il totale è : %d\n", elaboraInput())
}

func elaboraInput() (tot int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := replaceOverlapping(scanner.Text())
		match := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`).FindAllString(line, -1)
		convertMatch := convertStringNumtoInt(match)
		num1 := convertMatch[0] * 10
		if len(match) == 1 {
			tot += num1 + (num1 / 10)
		} else {
			tot += num1 + convertMatch[len(convertMatch)-1]
		}

	}
	return tot
}

func replaceOverlapping(line string) (ris string) {
	ris = strings.Replace(line, "oneight", "one eight", -1)
	ris = strings.Replace(ris, "nineight", "nine eight", -1)
	ris = strings.Replace(ris, "sevenine", "seven nine", -1)
	ris = strings.Replace(ris, "twone", "two one", -1)
	ris = strings.Replace(ris, "threeight", "three eight", -1)
	ris = strings.Replace(ris, "fiveight", "five eight", -1)
	ris = strings.Replace(ris, "eightree", "eight three", -1)
	ris = strings.Replace(ris, "eightwo", "eight two", -1)
	return ris
}

func inizializzaMappa() map[string]int {
	dict := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	return dict
}

func convertStringNumtoInt(match []string) (ris []int) {
	dict := inizializzaMappa()
	for i := 0; i < len(match); i++ {
		if val, isPresente := dict[match[i]]; isPresente {
			ris = append(ris, val)
		} else {
			n, _ := strconv.Atoi(match[i])
			ris = append(ris, n)
		}
	}
	return ris

}
