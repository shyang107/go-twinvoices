package main

import "fmt"

func main() {
	test1()
}

func test1() {
	var a [3][5]string
	var b [3][5][7]string
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("len(b) = %d\n", len(b))
	for i := 0; i < len(a); i++ {
		fmt.Printf("> len(a[%d]) = %d\n", i, len(a[i]))
		fmt.Printf("> len(b) = %d\n", len(b))
		for j := 0; j < len(b[i]); j++ {
			fmt.Printf(">> len(b[%d][%d]) = %d\n", i, j, len(b[i][j]))
		}
	}
}
