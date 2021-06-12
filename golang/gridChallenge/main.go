package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func grid_order(grid [][]string) [][]string {
	for i := 0; i < len(grid); i++ {
		sort.Strings(grid[i])
	}
	return grid
}

func grid_checker(grid [][]string) bool {
	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid)-1; j++ {
			if !(grid[i][j] <= grid[i][j+1] && grid[i][j] <= grid[i+1][j]) {
				return false
			}
		}
		if !(grid[i][len(grid)-1] <= grid[i+1][len(grid)-1]) {
			return false
		}
	}
	return true
}

func main() {
	var T int
	var N int
	var buffer string
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &T)
	fmt.Fscan(io, &N)
	fmt.Fscan(io, &buffer)
	for k := 0; k < T; k++ {
		var N int
		G := make([][]string, N)
		for i := 0; i < N; i++ {
			var buffer string
			G[i] = strings.Split(buffer, "")
		}
		if grid_checker(grid_order(G)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
