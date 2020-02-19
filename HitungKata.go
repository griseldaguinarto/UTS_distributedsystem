package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountWord(s string) map[string]int {
	words := strings.Fields(s)
	countWord := make(map[string]int)
	for i := range words {
		countWord[words[i]]++
	}
	return countWord
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(CountWord(scanner.Text()))
	}
}
