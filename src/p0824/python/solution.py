H = {'a', 'e', 'i', 'o', 'u',
     'A', 'E', 'I', 'O', 'U'}
MA = 'ma'


class Solution:

    def handle_h(self, s: str, i: int):
        s += MA + i * 'a'
        return s

    def handle_not_h(self, s, i):

        s = s[1:] + s[0]
        s += MA + i * 'a'
        return s

    def toGoatLatin(self, S: str) -> str:
        slst = S.split(' ')
        for i, v in enumerate(slst):
            if v[0] in H:
                slst[i] = self.handle_h(v, i+1)
            else:
                slst[i] = self.handle_not_h(v, i+1)

        return ' '.join(slst)