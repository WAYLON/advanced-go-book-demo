package main

import (
	"sync"
)

// ćšć±ćé
var counter int64
var m sync.Mutex
var wg sync.WaitGroup
func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}