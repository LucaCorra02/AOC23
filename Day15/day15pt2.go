package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type slot struct {
	label     string
	focusLent int
}

func main() {
	hashmap := leggiFile()
	printHashMap(hashmap)
	fmt.Printf("Pt2: %d\n", calcFocusingPower(hashmap))
}

func leggiFile() [][]slot {
	scanner := bufio.NewScanner(os.Stdin)
	hashmap := createMap(256)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ",")
		for _, word := range words {
			if strings.Contains(word, "=") {
				label := strings.Split(word, "=")[0]
				focusLent, _ := strconv.Atoi(strings.Split(word, "=")[1])
				searchLabel(hashmap, label, focusLent)
			} else {
				label := strings.Split(word, "-")[0]
				removeLabel(hashmap, label)
			}
		}
	}
	return hashmap
}

func removeLabel(hashmap [][]slot, label string) {
	indice := calcHash(label)
	for i, s := range hashmap[indice] {
		if s.label == label {
			newLine := append(hashmap[indice][:i], hashmap[indice][i+1:]...)
			hashmap[indice] = newLine
		}
	}
}

func searchLabel(hashmap [][]slot, label string, focusLent int) {
	indice := calcHash(label)
	for i, s := range hashmap[indice] {
		if s.label == label {
			hashmap[indice][i] = slot{label, focusLent}
			return
		}
	}
	hashmap[indice] = append(hashmap[indice], slot{label, focusLent})
}

func createMap(dim int) [][]slot {
	hashmap := make([][]slot, dim)
	for i := 0; i < dim; i++ {
		hashmap[i] = make([]slot, 0)
	}
	return hashmap
}

func calcHash(str string) (curVal int) {
	for _, char := range str {
		curVal = ((int(byte(char)) + curVal) * 17) % 256
	}
	return
}

func calcFocusingPower(hashmap [][]slot) (sum int) {
	for box, slots := range hashmap {
		for numslot, slot := range slots {
			sum += (box + 1) * (numslot + 1) * slot.focusLent
		}
	}
	return
}

func printHashMap(hashmap [][]slot) {
	for box, slots := range hashmap {
		fmt.Printf("Box %d: %v\n", box, slots)
	}
}
