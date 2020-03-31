package main

func judgeCircle(moves string) bool {
	d := make(map[int32]int, 4)
	for _, v := range moves {
		d[v-'A'] += 1
	}
	return d['U'-'A'] == d['D'-'A'] && d['L'-'A'] == d['R'-'A']
}
