package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("il totale è : %d\n", elaboraInput())
}

func elaboraInput() (tot int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		match := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`).FindAllString(scanner.Text(), -1)
		convertMatch := convertStringNumtoInt(match)
		num1 := convertMatch[0]
		num1 *= 10
		if len(match) == 1 {
			tot += num1 + (num1 / 10)
		} else {
			tot += num1 + convertMatch[len(convertMatch)-1]
		}

	}
	return tot
}

func convertStringNumtoInt(match []string) (ris []int) {
	for i := 0; i < len(match); i++ {
		ris = append(ris, numStringToInt(match[i]))
	}
	return ris
}

func numStringToInt(str string) int {
	switch str {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		n, _ := strconv.Atoi(str)
		return n
	}
}
