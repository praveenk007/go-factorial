package main

import (
	"fmt"
)

func main() {
	n := 4
	//series := make(chan int, n)
	fmt.Println(fibo(n))
	/*o fibo_start(series, n)
	for i := range series {
		fmt.Println(i)
	}*/
}

func fibo_start(series chan int, n int) {

	close(series)
}

func fibo(n int) int {
	if n == 0 {
		return 1
	}
	return n * fibo(n-1)
}
