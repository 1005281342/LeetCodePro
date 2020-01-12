from typing import List


class Solution:

    # 返回 x 在 arr 中的索引，如果不存在返回 -1
    def bin_search(self, arr, l, r, x):

        # 基本判断
        if r >= l:

            mid = int(l + (r - l) // 2)

            # 元素整好的中间位置
            if arr[mid] == x:
                return mid

                # 元素小于中间位置的元素，只需要再比较左边的元素
            elif arr[mid] > x:
                return self.bin_search(arr, l, mid - 1, x)

                # 元素大于中间位置的元素，只需要再比较右边的元素
            else:
                return self.bin_search(arr, mid + 1, r, x)

        else:
            # 不存在
            return -1

    def fairCandySwap(self, A: List[int], B: List[int]) -> List[int]:

        d = (sum(A) - sum(B)) // 2
        B.sort()

        for a in A:

            if self.bin_search(B, 0, len(B) - 1, a - d) >= 0:
                return [a, a - d]
