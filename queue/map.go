package queue

import "sync"

func Get(controlMap sync.Map, key string) string {
	res, ok := controlMap.Load(key)
	if ok {
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
