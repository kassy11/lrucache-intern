package lrucache

import "sync"

type LRUCache struct {
	limit   int
	values     map[int]*item
	currentAge int
	mutex      *sync.Mutex
}

func NewLRU(limit int) *LRUCache{
	if limit < 1 {
		panic("nonsensical LRU cache size specified")
	}

	return &LRUCache{
		limit: limit,
		values: make(map[int]*item, limit),
		currentAge: 0,
		mutex: new(sync.Mutex),
	}

}
func (c *LRUCache)IsEmpty()bool{
	if len(c.values) == 0{
		return true
	}
	return false
}