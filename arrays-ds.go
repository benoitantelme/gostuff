package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &arr[i])
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}
