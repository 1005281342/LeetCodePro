# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
def isBadVersion(version):
    pass


class Solution:

    def bin_search(self, start, end):
        while start < end:
            mid = ((end - start) >> 1) + start
            if isBadVersion(mid) is False:
                start = mid + 1
            else:
                end = mid
        return end

    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        return self.bin_search(1, n + 1)
