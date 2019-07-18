package singleton

import (
	"sync"
)

var once = sync.Once{}
var rwlock = sync.RWMutex{}

var inst *instance

type instance struct {
	name string
}

func (instance instance) GetName() string {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return instance.name
}

func (instance *instance) SetName(name string) {
	rwlock.Lock()
	defer rwlock.Unlock()
	instance.name = name
}

func GetInstance() *instance {
	once.Do(func() {
		inst = &instance{"test"}
	})

	return inst
}
