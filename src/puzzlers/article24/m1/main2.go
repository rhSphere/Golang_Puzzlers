package main

import (
	"fmt"
	"time"
)

func main() {
	var res, i uint64
	start := time.Now()

	for i = 0; i < 10000000000; i++ {
		res += i
	}
	elapsed := time.Since(start)

	fmt.Printf("执行消耗的时间为:", elapsed)
	fmt.Println(", result:", res)

	start = time.Now()
	ch1 := calc(1, 2500000000)
	ch2 := calc(2500000001, 5000000000)
	ch3 := calc(5000000001, 7500000000)
	ch4 := calc(7500000001, 10000000000)
	res = <-ch1 + <-ch2 + <-ch3 + <-ch4
	elapsed = time.Since(start)
	fmt.Printf("执行消耗的时间为:", elapsed)
	fmt.Println(", result:", res)
}

func calc(from, to uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		res := from
		for i := from; i < to; i++ {
			res += i
		}
		ch <- res
	}()

	return ch
}
