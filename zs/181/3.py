from collections import defaultdict


class UndergroundSystem:

    def __init__(self):
        self.dd = defaultdict(list)
        self.flag = defaultdict(bool)
        self.ds = defaultdict(list)

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.dd[id].append((stationName, t))

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        self.dd[id].append((stationName, t))
        self.flag[stationName] = True

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        if self.flag[endStation]:
            self.upd(startStation, endStation)
        v = self.ds[(startStation, endStation)]
        if not v:
            return 0
        return sum(v) / len(v)

    def upd(self, a, b):
        # self.ds[(a, b)] = list()
        for k, v in self.dd.items():
            for i in range(0, len(v) - 1, 2):
                if (v[0][0], v[1][0]) == (a, b):
                    self.ds[(a, b)].append(v[1][1] - v[0][1])
        self.flag[b] = False
