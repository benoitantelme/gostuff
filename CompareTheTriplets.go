package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	io := bufio.NewReader(os.Stdin)

	al := make([]int, 3)
	bob := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &al[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Fscan(io, &bob[i])
	}

	var alScore int
	var bobScore int

	for i := 0; i < 3; i++ {
		if al[i] > bob[i] {
			alScore++
		} else if bob[i] > al[i] {
			bobScore++
		}
	}
	fmt.Printf("%d %d", alScore, bobScore)

}
