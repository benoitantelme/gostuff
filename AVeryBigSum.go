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

	var sum int64
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i])
	}

	fmt.Println(sum)
}
