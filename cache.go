package cache

import (
	"sync"
	"time"
)
//Cache have cache string and requestd, timeout for duration.
type Cache struct {
	requestd time.Duration
	timeout  time.Duration
	value    string
	sync.RWMutex
}
//Set sets cache string
func (c *Cache) set(s string) {
	c.Lock()
	c.value = s
	c.Unlock()
}

func (c *Cache) getCached() string {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

//Query makes request with duration
func (c *Cache) Query() string {
	ch := make(chan string, 1)

	go func() { ch <- c.request() }()

	select {
	case <-time.After(c.timeout):
		return c.getCached()
	case r := <-ch:
		return r
	}
}

func (c *Cache) request() string {
	time.Sleep(c.requestd)
	response := "QueryValue"
	defer c.set(response)
	return response
}
