package lrucache_test

import (
	"testing"
	"github.com/kassy11/mylrucache/lrucache"
)


func TestSuccessNewLRU (t *testing.T) {
	lru_limit_ok := []int{1, 3, 5, 7}
	for i:=0; i<len(lru_limit_ok); i++{
		cache, _ := lrucache.NewLRU(lru_limit_ok[i])
		if !cache.IsEmpty(){
			t.Errorf("Case[%d]: Cannot generate empty LRU cache", i)
		}
	}
}

func TestFailNewLRU(t *testing.T){
	lru_limit_ng := []int{-1, 0, -100, -5}
	for i:=0; i<len(lru_limit_ng); i++{
		cache, err := lrucache.NewLRU(lru_limit_ng[i])
		if err == nil || cache != nil{
			t.Errorf("Case[%d]: should fail", i)
		}
	}

}