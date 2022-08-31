package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Get(controlMap sync.Map, key string) string {
	res, ok := controlMap.Load(key)
	if ok {
		fmt.Println(res.(string))
		return res.(string)
	}
	return ""
}
func Set(controlMap sync.Map, key string, value string) {
	controlMap.Store(key, value)
}
func Remove(controlMap sync.Map, key string) {
	controlMap.Delete(key)
}

func main() {
	var unique = atomic.Int64{}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			unique.Add(time.Now().Unix())
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("unique = ", unique.Load())
			if unique.Load() != 0 {
				fmt.Println("world")
				return
			}
		}
	}()

	var controlMap sync.Map
	for i := 1; i <= 1000000; i++ {
		go Set(controlMap, fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
		go Get(controlMap, fmt.Sprintf("%d", i))
	}
	var m1 = make(map[string]map[string]interface{})
	m1["hi"] = map[string]interface{}{"hello": "world"}
	fmt.Println(m1)
	for {
		select {}
	}

}
