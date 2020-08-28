package lrucache

// LRUのインタフェース
type LRU interface {
	Get(key int) int
	Put(key, value int)
}

type item struct {
	value int
	age   int
}