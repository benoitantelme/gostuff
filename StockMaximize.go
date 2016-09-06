package main

import "fmt"
import "bufio"
import "os"

type max struct {
	index int
	sum   int
}

func main() {
	var cases int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &cases)

	for i := 0; i < cases; i++ {
		var prices int
		fmt.Fscan(io, &prices)

		//build prices array
		pricesarr := make([]int, prices)
		for j := 0; j < prices; j++ {
			var val int
			fmt.Fscan(io, &val)
			pricesarr[j] = val
		}

		total := 0
		max := 0

		for j := prices - 1; j > -1; j-- {
			price := pricesarr[j]
			if max <= price {
				max = price
			}
			total += max - price
		}

		fmt.Println(total)
	}
}
