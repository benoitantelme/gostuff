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

	nbr := arr[n-1]

	for i := n - 2; i > -1; i-- {
		if arr[i] > nbr {
			arr[i+1] = arr[i]
			fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
		} else {
			arr[i+1] = nbr
			fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
			break
		}

		if i == 0 {
			arr[i] = nbr
			fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
		}
	}

}
