package syncing

import "sync"

// Counter type is concurency safe counter that uses Mutex from sync package
type Counter struct {
	lock  sync.Mutex
	value int
}

// NewCounter factory
func NewCounter() *Counter {
	return &Counter{}
}

// Inc method of counter increments value in concurent safe way using sync.Mutex lock and unlock
func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
}

// Value method of counter returns current value
func (c *Counter) Value() int {
	return c.value
}
