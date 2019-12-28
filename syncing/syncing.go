package syncing

import "sync"

type Counter struct {
	lock  sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
