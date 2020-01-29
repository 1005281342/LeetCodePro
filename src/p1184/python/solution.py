from typing import List


class Solution:
    def distanceBetweenBusStops(self, distance: List[int], start: int, destination: int) -> int:
        su = sum(distance)
        if start > destination:
            start, destination = destination, start
        a1 = sum(distance[start: destination])
        return min(a1, su - a1)

    # def distanceBetweenBusStops(self, distance: List[int], start: int, destination: int) -> int:
    #     n = len(distance)
    #     distance = distance + distance
    #     if start > destination:
    #         start, destination = destination, start
    #     return min(sum(distance[start:destination]), sum(distance[destination:n + start]))
