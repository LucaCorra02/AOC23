package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	comand, graph := leggiInput()
	fmt.Printf("mosse pt1: %d\n", visitGraph(comand, graph))
	movesPt2 := visitParallel(comand, graph)
	fmt.Printf("mosse pt2: %d\n", LCM(movesPt2, 0))
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

func visitParallel(comand []int, graph map[string][]string) (moves []int) {
	starts := searchStartNode(graph)
	ends := searchEndingNode(graph)

	for _, startingNode := range starts {
		node := startingNode
		move := 0
		for !slices.Contains(ends, node) {
			frMove := comand[move%len(comand)]
			node = graph[node][frMove] //mossa vera, 0 sinistra 1 destra
			move++
		}
		moves = append(moves, move)
	}
	return
}

func searchStartNode(graph map[string][]string) (starts []string) {
	for key := range graph {
		if strings.HasSuffix(key, "A") {
			starts = append(starts, key)
		}
	}
	return
}

func searchEndingNode(graph map[string][]string) (ends []string) {
	for key := range graph {
		if strings.HasSuffix(key, "Z") {
			ends = append(ends, key)
		}
	}
	return
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(numbers []int, index int) int {
	if index == len(numbers)-1 {
		return numbers[index]
	}
	a := numbers[index]
	b := LCM(numbers, index+1)
	return (a * b) / GCD(a, b)
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
