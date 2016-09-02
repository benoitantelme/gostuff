package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	first := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &first[i])
	}

	var m int
	fmt.Fscan(io, &m)

	second := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(io, &second[i])
	}

	missing := make(map[int]int)

	for i := 0; i < n; i++ {
		missing[first[i]] = missing[first[i]] + 1
	}

	for i := 0; i < m; i++ {
		missing[second[i]] = missing[second[i]] - 1
	}

	var keys []int
	for k := range missing {
		if missing[k] < 0 {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)
	fmt.Println(strings.Trim(fmt.Sprint(keys), "[]"))

}
