package mutex_map

import "sync"

type MutexMapIntMap struct {
	Mp map[int]map[string]string
	*sync.RWMutex
}

func (M *MutexMapIntMap) Data() map[int]map[string]string {
	return M.Mp
}

func (M *MutexMapIntMap) Get(k int) map[string]string {
	M.RLock()
	defer M.RUnlock()
	return M.Mp[k]
}

func (M *MutexMapIntMap) Set(k int,v map[string]string) {
	M.Lock()
	defer M.Unlock()
	M.Mp[k] = v
}

func (M *MutexMapIntMap) Delete(k int) {
	M.Lock()
	defer M.RUnlock()
	if _, ok := M.Mp[k]; !ok {
		return
	}else {
		delete(M.Mp,k)
	}
}
