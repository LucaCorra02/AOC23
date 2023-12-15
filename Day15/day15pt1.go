package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sumHash := leggiFile()
	fmt.Printf("Part1: %d\n", sumHash)
}

func leggiFile() (sumHash int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ",")
		for _, word := range words {
			sumHash += calcHash(word)
		}
	}
	return
}

func calcHash(str string) (curVal int) {
	for _, char := range str {
		curVal = ((int(byte(char)) + curVal) * 17) % 256
	}
	return
}
