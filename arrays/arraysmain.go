package main

import (
	"fmt"
	"practice/arrays/singleNum"
)

func main() {
	print_singleNum()
}

func print_singleNum() {
	test := [5]int{4, 1, 1, 4, 2}
	fmt.Println(singleNum.SingleNumber(test[:]))
}
