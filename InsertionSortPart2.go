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

	for i := 1; i < n; i++ {
		loop(arr, i)
		fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
	}

}

func loop(arr []int, n int) {
	nbr := arr[n]

	for i := n - 1; i > -1; i-- {

		if arr[i] > nbr {
			arr[i+1] = arr[i]
		} else {
			arr[i+1] = nbr
			break
		}

		if i == 0 {
			arr[i] = nbr
		}
	}
}
