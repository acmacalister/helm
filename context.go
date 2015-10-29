package helm

import (
	"net/http"
	"sync"
)

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]map[interface{}]interface{})
)

func Set(r *http.Request, key, val interface{}) {
	mutex.Lock()
	if data[r] == nil {
		data[r] = make(map[interface{}]interface{})
	}

	data[r][key] = val
	mutex.Unlock()
}

func Get(r *http.Request, key interface{}) interface{} {
	mutex.RLock()
	if ctx := data[r]; ctx != nil {
		value := ctx[key]
		mutex.RUnlock()
		return value
	}
	mutex.RUnlock()
	return nil
}

func clear(r *http.Request) {
	mutex.Lock()
	delete(data, r)
	mutex.Unlock()
}
