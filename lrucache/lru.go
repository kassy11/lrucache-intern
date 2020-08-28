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

func (c *LRUCache)Put(key int, value int){

}

func (c *LRUCache)Get(key int)int{

}
