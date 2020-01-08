class Solution:
    def maximumSwap(self, num: int) -> int:

        nums = [int(i) for i in str(num)]
        tn = sorted(nums, reverse=True)
        for i, v in enumerate(nums[:len(nums) - 1]):
            if nums[i] == tn[i]:
                continue
            m = nums[i]
            ti = i
            for j in range(i + 1, len(nums)):
                if nums[j] >= m:
                    ti = j
                    m = nums[j]
            nums[i], nums[ti] = nums[ti], nums[i]
            break

        ans = ''.join([str(c) for c in nums])
        return int(ans)
