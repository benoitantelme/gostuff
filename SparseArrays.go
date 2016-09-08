package main

import "fmt"
import "bufio"
import "os"

func main() {
	var strings int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &strings)

	m := make(map[string]int)

	for i := 0; i < strings; i++ {
		var str string
		fmt.Fscan(io, &str)

		value, ok := m[str] // return value if found or ok=false if not found

		if ok {
			m[str] = value + 1
		} else {
			m[str] = 1
		}
	}

	var queries int
	fmt.Fscan(io, &queries)

	for i := 0; i < queries; i++ {
		var q string
		fmt.Fscan(io, &q)
		fmt.Println(m[q])
	}

}
