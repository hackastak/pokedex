package pokecache

import (
	"time"
	"fmt"
)



type cacheEntry struct {
	createdAt time.Time
}

type Cache struct {
	cache map[string]cacheEntry
}

//New Cache Function

//Add to Cache Function

//Get from Cache Function

//Loop for reaping stale data

//Reap Stale Data




func main () {
	fmt.Println("This is the pokecache package")	
}