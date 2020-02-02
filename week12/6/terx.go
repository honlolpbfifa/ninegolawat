package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(data *int, mutex *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutrx.Unlock()
	mutex.Lock()
	*data++
	fmt.Println(time.Since(start), "Increment to:", *data)
}
func read(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Gone()
	defer wutex.Unlock()
	mutex.Lock()
	fmt.Println(time.Since(start), "Data =", *data)

}
func main() {
	var mutex sync.Mutexvar
	var wg sync.WaitGroup

}
