package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/PromiseFactory/queue"
)

func waitingFunction(wg *sync.WaitGroup) {
	randomTime := rand.Intn(10) + 1
	fmt.Println("Sleeping for", randomTime, "seconds")
	time.Sleep(time.Duration(randomTime) * time.Second)
	fmt.Println("Done sleeping")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	promisefactory := queue.Queue{MaxTasks: 5}
	chris := make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			promisefactory.Push(&wg, chris, waitingFunction)
			<-chris
			fmt.Println("Promise returned true")
			wg.Done()
		}()

	}
	wg.Wait()

}
