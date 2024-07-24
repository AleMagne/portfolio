package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Inserisci un testo per vedere\n")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	words := strings.Split(text, " ")
	for _, word := range words {
		fmt.Println(word)
	}
}
