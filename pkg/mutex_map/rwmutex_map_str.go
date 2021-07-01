package mutex_map

import "sync"

type MutexMapStr struct {
	Mp map[string]string
	*sync.RWMutex
}

func (M *MutexMapStr) Data() map[string]string {
	return M.Mp
}

func (M *MutexMapStr) Get(k string) string {
	M.RLock()
	defer M.RUnlock()
	return M.Mp[k]
}

func (M *MutexMapStr) Set(k, v string) {
	M.Lock()
	defer M.Unlock()
	M.Mp[k] = v
}

func (M *MutexMapStr) Delete(k string) {
	M.Lock()
	defer M.Unlock()
	if _, ok := M.Mp[k]; !ok {
		return
	}else {
		delete(M.Mp,k)
	}
}
