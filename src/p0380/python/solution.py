from random import randint


class RandomizedSet:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.se = set()
        self.cnt = 0

    def search(self, val):
        return val in self.se

    def insert(self, val: int) -> bool:
        """
        Inserts a value to the set. Returns true if the set did not already contain the specified element.
        """
        if self.search(val):
            return False
        self.se.add(val)
        return True

    def remove(self, val: int) -> bool:
        """
        Removes a value from the set. Returns true if the set contained the specified element.
        """
        if self.search(val):
            self.se -= {val, }
            return True
        return False

    def getRandom(self) -> int:
        """
        Get a random element from the set.
        """
        return list(self.se)[randint(0, len(self.se) - 1)]

# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
