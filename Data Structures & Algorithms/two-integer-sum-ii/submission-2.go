func twoSum(numbers []int, target int) []int {
    hash := make(map[int]int)
    for i,v := range numbers {
        if _, ok := hash[target-v]; !ok {
            hash[v] = i
        } else {
            return []int{hash[target-v] + 1, i + 1}
        }
    }
    return []int{}
}
