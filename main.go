package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/PromiseFactory/queue"
)

func awaitingFunction(wg *sync.WaitGroup) {
	randomTime := rand.Intn(1)+1
	fmt.Println("Sleeping for", randomTime, "seconds")
	time.Sleep(time.Duration(randomTime) * time.Second)
	fmt.Println("Done sleeping")
	wg.Done()
}

func main() {
	queueOfPromises:= queue.Queue{MaxTasks: 2}
   chris := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (){
		queueOfPromises.Push(&wg, chris, awaitingFunction)
		<-chris
		fmt.Println("Promise completed")
		}()
		
	}
	wg.Wait()
}
