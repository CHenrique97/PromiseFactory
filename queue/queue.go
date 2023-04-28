package queue

import (
	"fmt"
	"sync"
)

type node struct {
	next     *node
	previous *node
	content  func()
}

type Queue struct {
	head         *node
	tail         *node
	maxTasks     int
	runningTasks int
}

func (q *Queue) Push(wg *sync.WaitGroup, PromiseFactory func(*sync.WaitGroup) ){
	fmt.Println("Pushing")
	task := func() {
		fmt.Println("Executing")
		q.runningTasks++
		wg.Add(1)
		go PromiseFactory(wg)
		defer func() {
			q.runningTasks--
			q.tryToExecute(wg)
		}()
	}
	n := &node{}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		n.previous = q.tail
		q.tail = n
	}
	if q.runningTasks < 2 {
		task()
	} else {
		q.head.content = task
	}
}

func (q *Queue) Pop() func() {
	if q.head == nil {
		return nil
	} else {
		n := q.head
		q.head = q.head.next
		return n.content
	}
}

func (q *Queue) tryToExecute(wg *sync.WaitGroup) {
	if q.runningTasks >= 2 && q.head == nil {
		fmt.Println("Queue is empty")
		return 
	}
	q.Pop()

}
