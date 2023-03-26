package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {

	}()
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

}
