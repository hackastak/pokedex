package pokecache

import (
	"time"
	"fmt"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	value []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

//New Cache Function
func NewCache(interval time.Duration) Cache{
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	
	return c
}

//Add to Cache Function
func (c *Cache) Add(key string, value []byte) {
	fmt.Println("Add a new entry to the cache")
	fmt.Println(key)
	fmt.Println(value)
}

//Get from Cache Function
func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Println("Get an entry from the cache")
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.value, ok
}

//Loop for reaping stale data
func (c *Cache) reapLoop(interval time.Duration) {
	fmt.Println("Start reaping loop for Cache entries to remove data after specific time interval")
}

//Reap Stale Data
func (c *Cache) reap(now time.Time, last time.Duration) {
	fmt.Println("Reap stale data entries from the cache")
}



func main () {
	fmt.Println("This is the pokecache package")	
}