package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	comand, graph := leggiInput()
	fmt.Printf("mosse pt1: %d\n", visitGraph(comand, graph))
}

func visitGraph(comand []int, graph map[string][]string) (move int) {
	prev := "AAA"
	for i := 0; i < len(comand); i++ {
		curr := graph[prev][comand[i]]
		if prev == "ZZZ" {
			break
		}
		prev = curr
		move++
		if i+1 == len(comand) {
			i = -1
		}
	}
	return move
}

func leggiInput() (comand []int, graph map[string][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	graph = make(map[string][]string)

	scanner.Scan()
	comand = parseComand([]rune(scanner.Text()))
	scanner.Scan()
	scanner.Text()

	for scanner.Scan() {
		line := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "(", ""), ")", "")
		splitted := strings.Split(line, " = ")
		graph[splitted[0]] = strings.Split(splitted[1], ", ")
	}
	return
}

func parseComand(comand []rune) (parsed []int) {
	for _, val := range comand {
		if val == 'R' {
			parsed = append(parsed, 1)
		} else {
			parsed = append(parsed, 0)
		}
	}
	return
}
