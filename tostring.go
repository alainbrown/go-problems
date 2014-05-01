package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	tests := []int64{123, -123, 0, math.MaxInt64, math.MinInt64}
	for _, i := range tests {
		if test(i) {
			fmt.Println("test:", i, ": Passed")
		} else {
			fmt.Println("test:", i, ": Failed")
		}
	}
}

func test(i int64) bool {
	return strconv.FormatInt(i, 10) == tostring(i)
}

func tostring(n int64) string {
	if n > 0 {
		return solve(n)
	} else if n == 0 {
		return "0"
	} else if n == math.MinInt64 {
		return "-" + solve(math.MaxInt64/10) +
			ch(math.MaxInt64%10+1)
	} else {
		return "-" + solve(-n)
	}
}

func solve(n int64) string {
	if n <= 0 {
		return ""
	}
	return solve(n/10) + ch(n%10)
}

var m = map[int64]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func ch(n int64) string {
	return m[n]
}
