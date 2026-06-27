func twoSum(nums []int, target int) []int {
    visited := make(map[int]int)
    for j,v := range nums {
        if i,ok := visited[target-v]; ok {
            return []int{i, j}
        } else {
            visited[v] = j
        }
    }
    return []int{0,0}
}
