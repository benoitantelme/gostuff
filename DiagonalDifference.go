package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(io, &row[j])
		}
		matrix[i] = row
	}

	var first int
	var second int

	for i := 0; i < n; i++ {
		first += matrix[i][i]
		second += matrix[i][n-i-1]
	}

	fmt.Println(math.Abs(float64(first) - float64(second)))
}
