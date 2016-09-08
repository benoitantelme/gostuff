package main

import "bufio"
import "fmt"
import "os"

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &arr[i])
	}

	count := 0

	for i := 1; i < n; i++ {
		loopcount(arr, i, &count)

	}

	fmt.Println(count)

}

func loopcount(arr []int, n int, count *int) {
	nbr := arr[n]

	for i := n - 1; i > -1; i-- {

		if arr[i] > nbr {
			arr[i+1] = arr[i]
			*count += 1
		} else {
			arr[i+1] = nbr
			break
		}

		if i == 0 {
			arr[i] = nbr
		}
	}
}
