package main

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	length := len(A)
	j := 0
	for i := length - n; i < length; i++ {
		A[i] = B[j]
		j += 1
	}
	qSort(A)
}

func qSort(A []int) {
	if len(A) <= 1 {
		return
	}
	mid := A[0]
	head, tail := 0, len(A)-1
	for i := 1; i <= tail; {
		if A[i] > mid {
			A[i], A[tail] = A[tail], A[i]
			tail--
		} else {
			A[i], A[head] = A[head], A[i]
			i++
			head++
		}
	}
	qSort(A[:head])
	qSort(A[head+1:])
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
	fmt.Println(a)
}
