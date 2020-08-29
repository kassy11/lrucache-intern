package lrucache

// LRUのインタフェース
type LRU interface {
	Get(key int) int
	Put(key, value int)
}

// LRUCashe構造体のmapの値を管理
type item struct {
	value int
	age   int
}