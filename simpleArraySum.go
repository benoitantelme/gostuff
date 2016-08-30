package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &arr[i])
	}

	var sum int

	for i := 0; i < len(arr); i = i + 1 {
		sum = sum + arr[i]
	}

	fmt.Println(sum)
}
