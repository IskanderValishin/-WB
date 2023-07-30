package main

import (
	"fmt"
	"sync"
)

func sumSqrt() (sum int) {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			sum += num * num
			wg.Done()
		}(num)

	}
	wg.Wait()
	return sum
}
func main() {

	fmt.Println(sumSqrt())
}
