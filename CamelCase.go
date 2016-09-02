package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var str string
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &str)

	var nbr int
	for _, r := range str {
		if unicode.IsUpper(r) {
			nbr += 1
		}
	}

	fmt.Println(nbr + 1)
}
