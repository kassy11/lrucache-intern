package lrucache_test

import (
	"testing"
	"github.com/kassy11/mylrucache/lrucache"
)


func TestNewEmptyLRU(t *testing.T)  {
	lru_limit := 3
	cache := lrucache.NewLRU(lru_limit)

	if !cache.IsEmpty(){
		t.Errorf(`Cannot generate empty LRU cache`)
	}
}