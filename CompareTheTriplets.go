package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)

	alice := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &alice[i])
	}

	var score int
	var alScore int
	var bobScore int

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &score)

		if alice[i] > score {
			alScore++
		} else if score > alice[i] {
			bobScore++
		}

	}

	fmt.Printf("%d %d", alScore, bobScore)
}
