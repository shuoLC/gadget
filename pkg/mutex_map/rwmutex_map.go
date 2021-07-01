package mutex_map

import (
	"sync"
)

type MutexMap struct {
	Mp map[string]interface{}
	*sync.RWMutex
}

func (M *MutexMap) Data() map[string]interface{} {
	return M.Mp
}

func (M *MutexMap) Get(k string) interface{} {
	M.RLock()
	defer M.RUnlock()
	return M.Mp[k]
}

func (M *MutexMap) Set(k string, v interface{}) {
	M.Lock()
	defer M.Unlock()
	M.Mp[k] = v
}

func (M *MutexMap) Delete(k string) {
	M.Lock()
	defer M.Unlock()
	if _, ok := M.Mp[k]; !ok {
		return
	}else {
		delete(M.Mp,k)
	}
}