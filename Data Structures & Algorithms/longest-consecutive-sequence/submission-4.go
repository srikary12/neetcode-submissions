import(
    "slices"
)

func longestConsecutive(nums []int) int {
    if len(nums) == 0 { return 0 }
    l,r, count := 0,0,1
    slices.Sort(nums)
    nums = slices.Compact(nums)
    for i := 1; i< len(nums);i++ {
        if nums[i-1] + 1 == nums[i]{
            r = i
        } else {
            count = max(count,r-l+1)
            l = i
            r = i
        }
    }
    return max(count, r-l+1)
}
