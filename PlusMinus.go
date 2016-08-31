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

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &nums[i])
	}

	var pos, neg, zeros int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			pos += 1
		} else if nums[i] < 0 {
			neg += 1
		} else {
			zeros += 1
		}
	}

	fmt.Printf("%.6f\n", float64(pos)/float64(n))
	fmt.Printf("%.6f\n", float64(neg)/float64(n))
	fmt.Printf("%.6f", float64(zeros)/float64(n))
}
