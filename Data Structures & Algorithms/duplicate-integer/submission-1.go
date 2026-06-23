func hasDuplicate(nums []int) bool {
    temp := make(map[int]bool)
    for _, v := range nums {
        _, ok := temp[v]
        if ok {
            return true
        }
        if !ok {
            temp[v] = true
        }
    }
    return false
}
