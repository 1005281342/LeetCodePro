package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i := 1; i < len(xors); i++ {
		xors[i] = xors[i-1] ^ arr[i-1]
	}
	ret := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		ret = append(ret, xors[queries[i][0]]^xors[queries[i][1]+1])
	}
	return ret
}

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}
