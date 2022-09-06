package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"task/scheduler"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(10)

	var manager = scheduler.New()

	manager.SetLimit(2)

	var mux sync.RWMutex
	var taskMap = make(map[string]*scheduler.Worker)

	for i := 0; i < 10; i++ {

		var index = i

		go func() {

			var f1 = func() {
				//time.Sleep(time.Second * 1)
				mux.Lock()
				if taskMap[fmt.Sprintf("%v", index)] == nil {
					//log.Println("hello task", index)
					mux.Unlock()
					return
				}
				mux.Unlock()
				run(index)

			}

			mux.Lock()
			var worker = manager.Add(f1)
			taskMap[fmt.Sprintf("%d", index)] = worker
			mux.Unlock()

			var err = worker.Wait()
			if err != nil {
				log.Println("err:", err)
			}

			wait.Done()
			mux.Lock()
			delete(taskMap, fmt.Sprintf("%+v", index))
			mux.Unlock()
		}()
	}

	time.AfterFunc(time.Millisecond, func() {
		mux.Lock()
		task := taskMap["5"]
		delete(taskMap, "5")
		mux.Unlock()
		task.Stop()
	})

	wait.Wait()
	log.Println("taskMap = ", taskMap)
}

func run(i int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("running", i)
}
