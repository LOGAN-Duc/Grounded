package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	i    int
)

func CreateInstance() {
	once.Do(func() {
		fmt.Println("Running initialization...")
		i = 123 // chỉ chạy 1 lần
	})
}

func main() {
	wg := sync.WaitGroup{}

	// gọi CreateInstance từ nhiều goroutine
	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CreateInstance()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of i =", i)
}
