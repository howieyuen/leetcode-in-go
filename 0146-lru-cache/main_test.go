package _146_lru_cache

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 3)
	lruCache.Put(4, 1)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
}
