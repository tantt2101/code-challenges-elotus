package main

import (
	"fmt"
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
    graph := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }

    ressult := make([]int, n)
}

func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5},
	}
	res := sumOfDistancesInTree(n, edges);
	fmt.Println(res)
}
