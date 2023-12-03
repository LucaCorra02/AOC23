package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const COL = 140

func main() {
	field := leggiFile()
	//printField(field)
	fmt.Printf("punteggio: %d \n", calcolaPunteggio(field))
}

func calcolaPunteggio(field [][]rune) int {
	var tot int
	for i := 1; i < len(field)-1; i++ {
		for j := 1; j < len(field[i])-1; j++ {
			if unicode.IsDigit(field[i][j]) {
				if foundSimbol(field[i][j+1]) || foundSimbol(field[i][j-1]) ||
					foundSimbol(field[i+1][j+1]) || foundSimbol(field[i+1][j-1]) ||
					foundSimbol(field[i-1][j+1]) || foundSimbol(field[i-1][j-1]) {
					n, k := calcolaNum(field, i, j)
					j = k
					tot += n
				}
			}
		}
	}
	return tot
}

func calcolaNum(field [][]rune, i int, start int) (n int, k int) {
	numStr := string(field[i][start])
	k = start
	for {
		k--
		if !unicode.IsDigit(field[i][k]) {
			break
		}
		numStr = string(field[i][k]) + numStr
	}
	k = start
	for {
		k++
		if !unicode.IsDigit(field[i][k]) {
			break
		}
		numStr += string(field[i][k])
	}
	n, _ = strconv.Atoi(numStr)
	return
}

func foundSimbol(obj rune) bool {
	if obj == '.' {
		return false
	}
	return !unicode.IsDigit(obj)
}

func leggiFile() [][]rune {
	scanner := bufio.NewScanner(os.Stdin)
	field := make([][]rune, COL+2)
	var lunghezzaRiga int
	var cont = 1
	for scanner.Scan() {
		line := scanner.Text()
		lunghezzaRiga = len(line)
		field[cont] = make([]rune, lunghezzaRiga)
		field[cont] = []rune("." + line + ".")
		cont++
	}
	field[0] = []rune(strings.Repeat(".", lunghezzaRiga+2))
	field[COL+1] = []rune(strings.Repeat(".", lunghezzaRiga+2))
	return field
}

func printField(field [][]rune) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			fmt.Print(string(field[i][j]) + " ")
		}
		fmt.Println()
	}
}
