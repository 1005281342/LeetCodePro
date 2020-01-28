from math import gcd


class Solution:
    def canMeasureWater(self, x: int, y: int, z: int) -> bool:
        if x + y < z:
            return False
        if z in {x, y, x+y}:
            return True
        return not (z % gcd(x, y))
