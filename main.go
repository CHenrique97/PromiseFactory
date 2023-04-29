package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/PromiseFactory/queue"
)

func awaitingFunction(wg *sync.WaitGroup) {
	randomTime := rand.Intn(1,10)
	fmt.Println("Sleeping for", randomTime, "seconds")
	time.Sleep(time.Duration(randomTime) * time.Second)
	fmt.Println("Done sleeping")
	wg.Done()
}

func main() {
	queueOfPromises:= queue.Queue{MaxTasks: 1}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		queueOfPromises.Push(&wg, awaitingFunction)
	}
	wg.Wait()
}
