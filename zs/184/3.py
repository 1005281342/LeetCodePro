P = [
    ("&frasl;", "/"),
    ("&quot;", '"'),
    ("&apos;", "'"),
    ("&gt;", ">"),
    ("&lt;", "<"),
    ("&amp;", "&")
]


class Solution:
    def entityParser(self, text: str) -> str:
        for a, b in P:
            text = text.replace(a, b)

        return text


if __name__ == '__main__':
    S=Solution()
    print(S.entityParser(text = "&amp; is an HTML entity but &ambassador; is not."))
    print(S.entityParser(text = "and I quote: &quot;...&quot;"))
    print(S.entityParser(text = "Stay home! Practice on Leetcode :)"))
    print(S.entityParser(text = "x &gt; y &amp;&amp; x &lt; y is always false"))
    print(S.entityParser(text = "leetcode.com&frasl;problemset&frasl;all"))

