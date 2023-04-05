package main

import "fmt"

// # 200
// problem statement
// https://leetcode.com/problems/number-of-islands

func main() {
	input := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(input) == 1)
	input = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(input) == 3)
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if dfs(grid, i, j) {
				count += 1
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) bool {
	isIsland := false
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return false
	}
	if grid[i][j] != '1' {
		return false
	}
	grid[i][j] = '0'
	isIsland = true
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
	return isIsland
}
