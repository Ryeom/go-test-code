package queue

import "sync"

type MutexChan struct {
	mutex sync.RWMutex
	List  map[string]chan bool
}

func NewMutexChan() *MutexChan {
	return &MutexChan{List: map[string]chan bool{}}
}

func (m *MutexChan) Add(key string) {
	m.mutex.Lock()
	m.List[key] = make(chan bool, 1)
	m.mutex.Unlock()
}

func (m *MutexChan) Delete(key string) {
	m.mutex.Lock()
	_, ok := m.List[key]
	if ok {
		close(m.List[key])
		delete(m.List, key)
	}
	m.mutex.Unlock()
}

func (m *MutexChan) Get(key string) chan bool {
	return m.List[key]
}

func (m *MutexChan) GetListAll() []string {
	chanList := []string{}
	m.mutex.Lock()
	for key, _ := range m.List {
		chanList = append(chanList, key)
	}
	m.mutex.Unlock()
	return chanList
}
