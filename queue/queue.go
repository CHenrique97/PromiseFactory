package queue

import (
	"fmt"
	"sync"
)

type node struct {
	content func()
}

type Queue struct {
	list         []*node
	MaxTasks     int
	runningTasks int
}

func (q *Queue) Push(wg *sync.WaitGroup, chris chan bool, PromiseFactory func(*sync.WaitGroup)) {

	task := func() {
		wg.Add(1)
		q.runningTasks++
		go func() {
			PromiseFactory(wg)
			defer func() {
				q.runningTasks--
				q.tryToExecute()
				chris <- true

			}()
		}()

	}
	n := &node{}
	if q.runningTasks < q.MaxTasks {
		task()
	} else {
		n.content = task
		q.list = append(q.list, n)
	}
	wg.Done()
}

func (q *Queue) tryToExecute() {
	if q.runningTasks >= q.MaxTasks || len(q.list) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Trying to execute")
	q.list[0].content()
	q.list = q.list[1:]

}
