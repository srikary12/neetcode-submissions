func maxAreaOfIsland(grid [][]int) int {
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    rows, cols := len(grid), len(grid[0])
	maxArea := 0
	currentArea := 0

    var dfs func(r, c int, area *int)
	dfs = func(r, c int, area *int) {
        if r < 0 || c < 0 || r >= rows ||
           c >= cols || grid[r][c] == 0 {
            return
        }
        grid[r][c] = 0
		*area += 1
		if *area>maxArea{
			maxArea = *area
		}
        for _, dir := range directions {
            dfs(r+dir[0], c+dir[1], area)
        }
    }
	for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
				currentArea = 0
                dfs(r, c, &currentArea)
            }
        }
    }
    return maxArea
}
