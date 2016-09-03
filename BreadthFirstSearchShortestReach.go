package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var graphs int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &graphs)

	for i := 0; i < graphs; i++ {
		var nbrnodes int
		fmt.Fscan(io, &nbrnodes)

		var nbredges int
		fmt.Fscan(io, &nbredges)

		//build edges array
		edges := make([]int, nbredges*2)
		for j := 0; j < nbredges*2; j++ {
			var val int
			fmt.Fscan(io, &val)
			edges[j] = val
		}

		matrix := buildMatrix(nbrnodes, edges)

		var start int
		fmt.Fscan(io, &start)

		result := checkneighbours(matrix, start)

		//remove start index
		result = append(result[:start-1], result[start:]...)

		fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
	}
}

func buildMatrix(size int, edges []int) [][]int {
	matrix := make([][]int, size)

	//initialise
	for i := 0; i < size; i++ {
		row := make([]int, size)
		matrix[i] = row
	}

	//fill
	for i := 0; i < len(edges); i += 2 {
		x := edges[i] - 1
		y := edges[i+1] - 1

		matrix[x][y] = 1
		matrix[y][x] = 1
	}
	return matrix
}

func checkneighbours(matrix [][]int, start int) []int {
	//initialize result array to -1
	result := make([]int, len(matrix))
	for i := range result {
		result[i] = -1
	}

	//initialize values to check array with the first node
	tocheck := make([]int, 0)
	tocheck = append(tocheck, start-1)

	distance := 1

	for len(tocheck) > 0 {
		result, tocheck = reccheck(result, tocheck, distance, start-1, matrix)
		distance += 1
	}

	return result
}

func reccheck(result []int, tocheck []int, distance int, start int, matrix [][]int) ([]int, []int) {
	checklength := len(tocheck)
	for k := 0; k < checklength; k++ {
		//set current node
		node := tocheck[k]
		for i := 0; i < len(matrix); i++ {
			// if not starting or current node and new link
			if i != node && i != start && matrix[node][i] == 1 && result[i] == -1 {
				result[i] = distance * 6
				tocheck = append(tocheck, i)
			}
		}
	}

	//clean check list
	tocheck = tocheck[checklength:]

	return result, tocheck
}
