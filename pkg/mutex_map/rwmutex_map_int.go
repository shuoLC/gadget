package mutex_map

import "sync"

type MutexMapInt struct {
	Mp map[int]string
	*sync.RWMutex
}

func (M *MutexMapInt) Data() map[int]string {
	return M.Mp
}

func (M *MutexMapInt) Get(k int) string {
	M.RLock()
	defer M.RUnlock()
	return M.Mp[k]
}

func (M *MutexMapInt) Set(k int, v string) {
	M.Lock()
	defer M.Unlock()
	M.Mp[k] = v
}

func (M *MutexMapInt) Delete(k int) {
	M.Lock()
	defer M.Unlock()
	if _, ok := M.Mp[k]; !ok {
		return
	}else {
		delete(M.Mp,k)
	}
}
