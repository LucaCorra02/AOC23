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
		match := regexp.MustCompile(`\d`).FindAllString(scanner.Text(), -1)
		num1, _ := strconv.Atoi(match[0])
		num1 *= 10
		if len(match) == 1 {
			tot += num1 + (num1 / 10)
		} else {
			num2, _ := strconv.Atoi(match[len(match)-1])
			tot += num1 + num2
		}
	}
	return tot
}
