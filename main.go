package main

import (
	"fmt"
)

func main() {
	n := 4
	series := make(chan int, n)
	go fact_start(series, n)
	for i := range series {
		fmt.Println(i)
	}
}

func fact_start(series chan int, n int) {

	close(series)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
