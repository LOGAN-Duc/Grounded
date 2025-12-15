package main

import (
	"sync"
)

var num int = 0

func add(lc *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i = i + 1 {
		lc.Lock()
		num = num + 1
		lc.Unlock()
	}
}
func minus(lc *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i = i + 1 {
		lc.Lock()
		num = num - 1
		lc.Unlock()
	}
}
func main() {
	var mutex *sync.Mutex = new(sync.Mutex)
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(2)
	go add(mutex, wg)
	go minus(mutex, wg)
	wg.Wait()

	println(num) // 0
}
