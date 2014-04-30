package main

import "fmt"

func main() {
	fmt.Println("123" == tostring(123))
	fmt.Println("-123" == tostring(-123))
	fmt.Println("0" == tostring(0))
}

func tostring(n int) string {
	if n > 0 {
		return solve(n)
	} else if n == 0 {
		return "0"
	} else {
		return "-" + solve(-n)
	}
}

func solve(n int) string {
	if n <= 0 {
		return ""
	}
	return solve(n/10) + ch(n%10)
}

var m = map[int]string{
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

func ch(n int) string {
	return m[n]
}
