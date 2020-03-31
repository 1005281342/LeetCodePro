package main

type WordsFrequency struct {
	mp map[string]int
}

func Constructor(book []string) WordsFrequency {
	mp := make(map[string]int, 16)
	for _, v := range book {
		mp[v] += 1
	}
	return WordsFrequency{mp: mp}
}

func (this *WordsFrequency) Get(word string) int {
	v := this.mp[word]
	return v
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
