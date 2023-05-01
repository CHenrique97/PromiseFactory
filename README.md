# Queue Package

This is a package that implements a queue with a specified maximum number of concurrent tasks.

## How to Use

First, import the package:

import "github.com/PromiseFactory/queue"

Then, create a new queue with a specified maximum number of concurrent tasks:

queueOfPromises := queue.Queue{MaxTasks: 2}

Next, create a function that you want to execute concurrently. For example:

func awaitingFunction(wg *sync.WaitGroup) {
    randomTime := rand.Intn(1) + 1
    fmt.Println("Sleeping for", randomTime, "seconds")
    time.Sleep(time.Duration(randomTime) * time.Second)
    fmt.Println("Done sleeping")
    wg.Done()
}

Finally, add the task to the queue using the `Push` method:

var wg sync.WaitGroup
chris := make(chan bool)

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        queueOfPromises.Push(&wg, chris, awaitingFunction)
        <-chris
        fmt.Println("Promise completed")
    }()
}
wg.Wait()

This will execute `awaitingFunction` concurrently with a maximum of 2 tasks at a time. Once all tasks have completed, the program will exit.
