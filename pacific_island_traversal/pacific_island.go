package main

import "fmt"

// Problem Leetcode 417
type Pair struct {
	r int
	c int
}

var (
	ROWS = 0
	COLS = 0
)

func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS = len(heights), len(heights[0])
	pacific, atlantic := make(map[Pair]bool), make(map[Pair]bool)

	for c := 0; c < COLS; c++ {
		dfs(0, c, heights, pacific, heights[0][c])
		dfs(ROWS-1, c, heights, atlantic, heights[ROWS-1][c])
	}
	for r := 0; r < ROWS; r++ {
		dfs(r, 0, heights, pacific, heights[r][0])
		dfs(r, COLS-1, heights, atlantic, heights[r][COLS-1])
	}

	res := make([][]int, 0)

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if pacific[Pair{r, c}] && atlantic[Pair{r, c}] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

func dfs(r, c int, heights [][]int, visit map[Pair]bool, prevHeight int) {
	if r < 0 || c < 0 || r == ROWS || c == COLS || visit[Pair{r, c}] || heights[r][c] < prevHeight {
		return
	}
	visit[Pair{r, c}] = true
	dfs(r, c+1, heights, visit, heights[r][c])
	dfs(r, c-1, heights, visit, heights[r][c])
	dfs(r+1, c, heights, visit, heights[r][c])
	dfs(r-1, c, heights, visit, heights[r][c])
}

func main() {
	heights := [][]int{[]int{1, 2, 2, 3, 5}, []int{3, 2, 3, 4, 4}, []int{2, 4, 5, 3, 1}, []int{6, 7, 1, 4, 5}, []int{5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(heights))
}
