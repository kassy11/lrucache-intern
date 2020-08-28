package lrucache

import (
	"fmt"
	"sync"
)

type LRUCache struct {
	limit   int
	values     map[int]*item
	currentAge int
	mutex      *sync.Mutex
}

func NewLRU(limit int) (*LRUCache, error){
	if limit < 1 {
		return nil, fmt.Errorf("nonsensical LRU cache size specified\n")
	}

	return &LRUCache{
		limit: limit,
		values: make(map[int]*item, limit),
		currentAge: 0,
		mutex: new(sync.Mutex),
	}, nil

}
func (c *LRUCache)IsEmpty()bool{
	if len(c.values) == 0{
		return true
	}
	return false
}

func (c *LRUCache) Get(key int) int {
	i, ok := c.values[key]
	if !ok {
		return -1
	}
	c.mutex.Lock()
	i.age = c.currentAge
	c.currentAge++
	c.mutex.Unlock()
	return i.value
}

func (c *LRUCache) Put(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	i, ok := c.values[key]
	// キーが存在する時は更新する
	if ok {
		i.value = value
		// TODO: Putしたときは使われたとみなす？
		i.age = c.currentAge
		c.currentAge++
	}else {
		c.values[key] = &item{
			value: value,
			age:   c.currentAge,
		}
		c.currentAge++
	}
}