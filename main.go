package main

import (
	"fmt"
	"github.com/kassy11/mylrucache/lrucache"
)

// プロンプトでLRUの初期化と入力を繰り返す？
func main() {
	cache, _ := lrucache.NewLRU(3)

	fmt.Println("IsEmpty(): %t", cache.IsEmpty()) // true
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	fmt.Printf("Get(3): %d\n", cache.Get(3)) // 3
	fmt.Printf("Get(2): %d\n", cache.Get(2)) // 2
	fmt.Printf("Get(10): %d\n", cache.Get(10)) //
	cache.Put(4, 40)
	fmt.Printf("Get(1): %d", cache.Get(1)) // -1
	fmt.Printf("Get(4): %d", cache.Get(4)) // 4
}