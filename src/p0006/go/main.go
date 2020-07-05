package main

import "strings"

func convert(s string, numRows int) string {

	if numRows <= 1 {
		return s
	}

	var charNums [][]byte = make([][]byte, numRows, numRows)
	var i int
	for ; i < numRows; i++ {
		charNums[i] = make([]byte, 0)
	}

	i = 0
	var flag = -1
	for index := 0; index < len(s); index++ {
		charNums[i] = append(charNums[i], s[index])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	var stringList []string = make([]string, numRows, numRows)
	for i := 0; i < numRows; i++ {
		stringList[i] = string(charNums[i])
	}

	return strings.Join(stringList, "")
}
