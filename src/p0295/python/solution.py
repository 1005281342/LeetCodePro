import heapq


class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        # 如果测试用例的容量增加，下面 10000 这个数值请大家自行调整
        self.max_heap = list()
        self.min_heap = list()
        heapq.heapify(self.max_heap)
        heapq.heapify(self.min_heap)

    def addNum(self, num: 'int') -> 'None':
        # 大顶堆先进一个元素
        heapq.heappush(self.max_heap, num)
        # 然后从大顶堆里出一个元素到小顶堆
        heapq.heappush(self.min_heap, -heapq.heappop(self.max_heap))
        if len(self.max_heap) < len(self.min_heap):
            # 如果大顶堆的元素少于小顶堆
            # 就要从小顶堆出一个元素到大顶堆
            heapq.heappush(self.max_heap, -heapq.heappop(self.min_heap))

    def findMedian(self) -> 'float':
        if len(self.max_heap) == len(self.min_heap):
            return (self.max_heap[0] - self.min_heap[0]) / 2
        else:
            return self.max_heap[0]

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()

# TML
# class MedianFinder:
#
#     def __init__(self):
#         """
#         initialize your data structure here.
#         """
#         self.ma = list()
#         self.mi = list()
#         heapq.heapify(self.ma)
#         heapq.heapify(self.mi)
#
#     def addNum(self, num: int) -> None:
#         heapq.heappush(self.ma, num)
#         heapq.heappush(self.mi, -num)
#
#     def findMedian(self) -> float:
#         if len(self.ma) == 1:
#             return self.ma[0]
#         if not self.ma:
#             return 0
#         a = heapq.nlargest((len(self.ma) + 1) // 2, self.ma)
#         b = heapq.nlargest((len(self.mi) + 1) // 2, self.mi)
#         return (a[-1] - b[-1]) / 2
# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
