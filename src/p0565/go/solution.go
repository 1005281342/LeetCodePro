func arrayNesting(nums []int) int {
    ans := 1
    length := len(nums)
    for _, v := range(nums) {
        if ans > length / 2{
            return ans
        }
        t := v
        cnt := 1
        for nums[t] != v{
            t = nums[t]
            cnt += 1
        }
        if cnt > ans{
            ans = cnt
        }
    }

    return ans
}