package lrucache_test

import (
	"testing"
	"github.com/kassy11/mylrucache/lrucache"
)

// LRU生成部分についてのUnitTest
func TestNewLRU(t *testing.T) {
	t.Run("can create LRU with limit 1 or more", func(t *testing.T) {
		lru_limit_ok := []int{1, 3, 5, 7}
		for i:=0; i<len(lru_limit_ok); i++{
			cache, _ := lrucache.NewLRU(lru_limit_ok[i])
			if !cache.IsEmpty(){
				t.Errorf("Case[%d]: Cannot generate empty LRU cache", i)
			}
		}
	})

	t.Run("cannot create LRU with limit less than 1", func(t *testing.T) {
		lru_limit_ng := []int{-1, 0, -100, -5}
		for i:=0; i<len(lru_limit_ng); i++{
			cache, err := lrucache.NewLRU(lru_limit_ng[i])
			if err == nil || cache != nil{
				t.Errorf("Case[%d]: should fail", i)
			}
		}
	})
}

// GetとPutのUnitTest&&Integrationtest
func TestLRUCache_Get_and_Put(t *testing.T) {
	limit_num := 3
	t.Run("can Put and Get", func(t *testing.T) {
		cache, _ := lrucache.NewLRU(limit_num)
		cache.Put(1, 3)
		if got, want := cache.Get(1), 3; got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("update element with the same key", func(t *testing.T) {
		cache, _ := lrucache.NewLRU(limit_num)
		// keyが同じものは更新される
		cache.Put(2, 1)
		cache.Put(2, 2)
		if got, want := cache.Get(2), 2; got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("cannnot Get non-existent key", func(t *testing.T) {
		cache, _ := lrucache.NewLRU(limit_num)
		cache.Put(1, 1)
		cache.Put(4, 1)
		if got, want := cache.Get(2), -1; got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("replace least referenced cache if exceeds the limit", func(t *testing.T) {
		cache, _ := lrucache.NewLRU(limit_num)
		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(3, 30)
		cache.Get(1)
		cache.Get(2)
		cache.Put(4, 40)
		got, want := cache.Get(4), 40
		got_deleted, want_deleted := cache.Get(3), -1
		if got_deleted != want_deleted || got != want{
			t.Errorf("got_deleted %v but want_deleted %v", got_deleted, want_deleted)
			t.Errorf("got %v but want %v", got, want)
		}
	})
}